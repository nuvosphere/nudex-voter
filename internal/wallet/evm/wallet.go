package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const defaultGasLimit = 1000000

type WalletClient struct {
	ctx                 context.Context
	cancel              context.CancelFunc
	event               eventbus.Bus
	client              *ethclient.Client
	submitterPrivateKey *ecdsa.PrivateKey
	submitter           common.Address
	pendingTx           sync.Map // txHash: bool
	chainID             atomic.Int64
	state               *state.EvmWalletState
	tss                 suite.TssService
	notify              chan struct{}
	voterContract       layer2.VoterContract
	currentVoterNonce   *atomic.Uint64
	taskQueue           *pool.Pool[uint64] // l2 task
	submitTaskQueue     *pool.Pool[uint64] // submit task
	operationsQueue     *pool.Pool[uint64] // pending batch task
}

func (c *WalletClient) Start(context.Context) {
	log.Info("evm wallet client is stopping...")
	c.loopApproveProposal()
}

func (c *WalletClient) Stop(context.Context) {
	c.cancel()
}

func NewWallet(event eventbus.Bus, tss suite.TssService, voterContract layer2.VoterContract, state *state.EvmWalletState) *WalletClient {
	client, err := ethclient.Dial(config.AppConfig.L2Rpc)
	utils.Assert(err)

	currentNonce := &atomic.Uint64{}
	nonce, _ := voterContract.TssNonce()
	if nonce != nil {
		currentNonce.Store(nonce.Uint64())
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &WalletClient{
		ctx:                 ctx,
		cancel:              cancel,
		event:               event,
		client:              client,
		submitterPrivateKey: config.L2PrivateKey,
		submitter:           crypto.PubkeyToAddress(config.L2PrivateKey.PublicKey),
		pendingTx:           sync.Map{},
		chainID:             atomic.Int64{},
		state:               state,
		tss:                 tss,
		notify:              make(chan struct{}, 1024),
		voterContract:       voterContract,
		currentVoterNonce:   currentNonce,
		taskQueue:           pool.NewTaskPool[uint64](),
		submitTaskQueue:     pool.NewTaskPool[uint64](),
		operationsQueue:     pool.NewTaskPool[uint64](),
	}
}

func (c *WalletClient) BalanceOf(erc20Token, owner common.Address) (*big.Int, error) {
	erc20, err := contracts.NewERC20(erc20Token, c.client)
	if err != nil {
		return nil, err
	}

	return erc20.BalanceOf(nil, owner)
}

func (c *WalletClient) BalanceAt(owner common.Address) (*big.Int, error) {
	return c.client.BalanceAt(context.Background(), owner, nil)
}

func (c *WalletClient) ChainID(ctx context.Context) (*big.Int, error) {
	if c.chainID.Load() == 0 {
		var (
			chainID *big.Int
			err     error
		)

		err = backoff.Retry(
			func() error {
				chainID, err = c.client.ChainID(ctx)
				return err
			},
			backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 3),
		)
		if err != nil {
			return nil, fmt.Errorf("chainID error: %c", err)
		}

		c.chainID.Store(chainID.Int64())
	}

	return big.NewInt(c.chainID.Load()), nil
}

func (c *WalletClient) SendSingedTx(ctx context.Context, tx *types.Transaction) error {
	return c.sendTransaction(ctx, tx)
}

func (c *WalletClient) EstimateGasAPI(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := c.client.EstimateGas(ctx, msg)
	if err != nil {
		err = errors.Join(ErrEstimateGas, wrapError(err))
		gasLimit = defaultGasLimit

		if errors.Is(err, ErrInsufficientFunds) {
			err = errors.Join(fmt.Errorf("insufficient funds for sender: %v", msg.From), err)
		}
	}

	return gasLimit, err
}

func (c *WalletClient) EstimateGas(ctx context.Context, account, contractAddress common.Address, data []byte) (uint64, error) {
	// Estimate GasPrice
	// gasPrice := opt.GasPrice
	price, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return 0, fmt.Errorf("SuggestGasPrice error: %c, account address: %v", err, account)
	}

	msg := ethereum.CallMsg{
		From:     account,
		To:       &contractAddress,
		Data:     data,
		GasPrice: price,
	}

	return c.EstimateGasAPI(ctx, msg)
}

func (c *WalletClient) BuildUnsignTx(
	ctx context.Context,
	account, contractAddress common.Address,
	value *big.Int, calldata []byte,
	Operations *db.Operations,
	EvmWithdraw *db.EvmWithdraw,
	EvmConsolidation *db.EvmConsolidation,
) (*types.Transaction, error) {
	head, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, wrapError(err)
	}

	// EIP1559
	// https://learnblockchain.cn/article/8593
	// https://blog.csdn.net/vigor2323/article/details/122817104
	// https://metamask.io/1559/
	// https://www.blocknative.com/blog/eip-1559-fees
	// 1、Transaction Fee=GasUsed*GasPrice=GasUsed*(BaseFee+MaxPriorityFee)；
	// 2、MaxFee = MaxGasPrice = (2*BaseFee)+MaxPriorityFee
	// 3、Burnt=BaseFee*GasUsed
	// 4、TxSavingsFees=MaxFee*GasUsed−(BaseFee+MaxPriorityFee)*GasUsed = (MaxFee-(BaseFee+MaxPriorityFee))*GasUsed
	// 5、Tip(minter get fee) = Min(Max fee - Base fee, Max priority fee)
	// 6、usedGasPrice = min(MaxPriorityFeePerGas + basefee, MaxFeePerGas)
	gasTipCap, err := c.client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, wrapError(err)
	}

	gasFeeCap := new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
	)
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		return nil, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", gasFeeCap, gasTipCap)
	}

	// Estimate GasLimit
	msg := ethereum.CallMsg{
		From:      account,
		To:        &contractAddress,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      calldata,
		Value:     value,
	}

	gas, err := c.EstimateGasAPI(ctx, msg)
	if err != nil {
		return nil, wrapError(err)
	}

	// extend gas limit 20%
	gasLimit := decimal.NewFromUint64(gas).Mul(decimal.NewFromFloat(1.2))

	nextNonce, err := c.client.PendingNonceAt(ctx, account)
	if err != nil {
		return nil, wrapError(err)
	}

	//latestNonce := c.nonce.Load()
	//if latestNonce >= nextNonce {
	//	nextNonce = latestNonce + 1
	//	c.nonce.Store(nextNonce)
	//}

	baseTx := &types.DynamicFeeTx{
		To:        &contractAddress,
		Nonce:     nextNonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit.BigInt().Uint64(),
		Value:     value,
		Data:      calldata,
	}

	tx := types.NewTx(baseTx)

	jsonData, err := tx.MarshalJSON()
	if err != nil {
		return nil, wrapError(err)
	}

	err = c.state.CreateTx(
		nil,
		account,
		decimal.NewFromUint64(nextNonce),
		jsonData,
		calldata,
		tx.Hash(),
		head.Number.Uint64(),
		Operations,
		EvmWithdraw,
		EvmConsolidation,
	)

	return tx, err
}

func (c *WalletClient) newTx(tx *db.EvmTransaction) (*types.Transaction, error) {
	oldTx := types.Transaction{}

	err := oldTx.UnmarshalJSON(tx.TxJsonData)
	if err != nil {
		return nil, err
	}

	gasLimit := oldTx.Gas()
	if gasLimit < defaultGasLimit {
		gasLimit = defaultGasLimit
	}

	// extend 20%
	gasTipCap := decimal.NewFromFloat(1.2).Mul(decimal.NewFromBigInt(oldTx.GasTipCap(), 0)).BigInt()
	baseTx := &types.DynamicFeeTx{
		To:        oldTx.To(),
		Nonce:     oldTx.Nonce(),
		GasTipCap: gasTipCap,
		GasFeeCap: oldTx.GasFeeCap(),
		Gas:       gasLimit * 2,
		Value:     oldTx.Value(),
		Data:      oldTx.Data(),
	}
	unSignedTx := types.NewTx(baseTx)
	signer := types.LatestSignerForChainID(oldTx.ChainId())

	signature, err := crypto.Sign(signer.Hash(unSignedTx).Bytes(), c.submitterPrivateKey)
	if err != nil {
		return nil, err
	}

	return unSignedTx.WithSignature(signer, signature)
}

// SpeedSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (c *WalletClient) speedSendOrderTx(ctx context.Context, oldOrderTx *db.EvmTransaction) (*types.Transaction, error) {
	signedTx, err := c.newTx(oldOrderTx)
	if err != nil {
		return nil, wrapError(err)
	}

	newTxHash := signedTx.Hash()

	c.pendingTx.Store(newTxHash, true)
	defer c.pendingTx.Delete(newTxHash)

	jsonData, err := signedTx.MarshalJSON()
	if err != nil {
		return signedTx, wrapError(err)
	}

	head, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return signedTx, wrapError(err)
	}

	err = c.state.CreateTx(nil,
		oldOrderTx.Sender,
		oldOrderTx.TxNonce,
		jsonData,
		oldOrderTx.Calldata,
		newTxHash,
		head.Number.Uint64(),
		oldOrderTx.Operations,
		oldOrderTx.EvmWithdraw,
		oldOrderTx.EvmConsolidation,
	)
	if err != nil {
		return signedTx, wrapError(err)
	}

	err = c.sendTransaction(ctx, signedTx)

	return signedTx, err
}

func (c *WalletClient) sendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := c.client.SendTransaction(ctx, tx)
	if err != nil {
		err = errors.Join(ErrSendTransaction, wrapError(err))
	}

	dbErr := c.state.UpdatePendingTx(tx.Hash())

	time.Sleep(2 * time.Second)

	return errors.Join(err, dbErr)
}

// SpeedSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (c *WalletClient) SpeedSendOrderTx(ctx context.Context, oldOrderTx *db.EvmTransaction) (common.Hash, *types.Receipt, error) {
	if oldOrderTx.Status == db.Pending {
		signedTx, err := c.speedSendOrderTx(ctx, oldOrderTx)
		if err != nil {
			return common.Hash{}, nil, err
		}

		r, err := c.waitTxSuccess(ctx, signedTx.Hash())

		return signedTx.Hash(), r, err
	}

	return common.Hash{}, nil, nil
}

// AgainSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (c *WalletClient) AgainSendOrderTx(ctx context.Context, tx *db.EvmTransaction) (common.Hash, *types.Receipt, error) {
	orderTx := &types.Transaction{}

	err := orderTx.UnmarshalJSON(tx.TxJsonData)
	if err != nil {
		return common.Hash{}, nil, wrapError(err)
	}

	txHash := orderTx.Hash()
	_, ok := c.pendingTx.Load(txHash)

	if ok {
		return txHash, nil, ErrTxPending
	}

	is := c.IsOnline(ctx, txHash)

	if is {
		return txHash, nil, fmt.Errorf("tx already online: %c or %c", ErrTxPending, ErrTxCompleted)
	}

	c.pendingTx.Store(txHash, true)
	defer c.pendingTx.Delete(txHash)

	err = c.sendTransaction(ctx, orderTx)
	if err != nil {
		if errors.Is(err, ErrIntrinsicGasTooLow) || errors.Is(err, ErrReplacement) || errors.Is(err, ErrAlreadyKnown) {
			_ = c.state.UpdateFailTx(txHash, err)
			return c.SpeedSendOrderTx(ctx, tx)
		}

		if errors.Is(err, ErrNonceTooLow) {
			_ = c.state.UpdateFailTx(txHash, err) // todo
		}

		return txHash, nil, err
	}

	r, err := c.waitTxSuccess(ctx, txHash)

	return txHash, r, err
}

func (c *WalletClient) RawTxBytes(ctx context.Context, tx *types.Transaction) []byte {
	chainID, _ := c.ChainID(ctx)
	signer := types.LatestSignerForChainID(chainID)

	return signer.Hash(tx).Bytes()
}

func (c *WalletClient) TransactionWithSignature(ctx context.Context, tx *types.Transaction, signature []byte) (*types.Transaction, error) {
	chainID, err := c.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer := types.LatestSignerForChainID(chainID)

	return tx.WithSignature(signer, signature)
}

func (c *WalletClient) SendTransactionWithSignature(ctx context.Context, tx *types.Transaction, signature []byte) error {
	chainID, err := c.ChainID(ctx)
	if err != nil {
		return err
	}
	signer := types.LatestSignerForChainID(chainID)

	tx, err = tx.WithSignature(signer, signature)
	if err != nil {
		return fmt.Errorf("tx with signature: %c", err)
	}

	return c.sendTransaction(ctx, tx)
}

func (c *WalletClient) WaitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	c.pendingTx.Store(txHash, true)
	defer c.pendingTx.Delete(txHash)

	return c.waitTxSuccess(ctx, txHash)
}

func (c *WalletClient) waitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	begin := time.Now()
	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		r, err := c.client.TransactionReceipt(ctx, txHash)
		// TransactionReceipt can is nil
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(time.Second)

				count--
			} else {
				return nil, errors.Join(fmt.Errorf("call TransactionReceipt method error, txHash: %s", txHash), err, ErrWallet)
			}
		} else {
			dbErr := c.state.UpdateBookedTx(txHash)
			return r, dbErr
		}
	}

	return nil, ErrTxFoundTimeOut
}

func (c *WalletClient) IsOnChain(ctx context.Context, txHash common.Hash) bool {
	receipt, err := c.client.TransactionReceipt(ctx, txHash)
	if receipt != nil && err == nil {
		return true
	}

	return false
}

func (c *WalletClient) IsPending(ctx context.Context, txHash common.Hash) bool {
	_, ok := c.pendingTx.Load(txHash)
	if ok {
		return ok
	}

	_, isPending, _ := c.client.TransactionByHash(ctx, txHash)

	return isPending
}

// IsOnline = IsPending + IsOnChain.
func (c *WalletClient) IsOnline(ctx context.Context, txHash common.Hash) bool {
	tx, isPending, _ := c.client.TransactionByHash(ctx, txHash)
	return isPending || tx != nil
}

func wrapError(err error) error {
	for _, wrap := range wrapErrorList {
		if errors.Is(err, wrap) {
			return errors.Join(err, wrap, ErrWallet)
		}
	}

	return err
}
