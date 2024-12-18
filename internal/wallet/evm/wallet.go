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
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const defaultGasLimit = 1000000

type WalletClient struct {
	*wallet.BaseWallet
	ctx                 context.Context
	cancel              context.CancelFunc
	event               eventbus.Bus
	client              *ethclient.Client
	submitterPrivateKey *ecdsa.PrivateKey
	submitter           common.Address
	pendingTx           sync.Map // txHash: bool
	chainID             atomic.Int64
	walletState         *state.EvmWalletState
	tss                 suite.TssService
	notify              chan struct{}
	currentVoterNonce   *atomic.Uint64
	submitTaskQueue     *pool.Pool[uint64] // submit task
	operationsQueue     *pool.Pool[uint64] // pending batch task
	txContext           sync.Map           // taskID:TxContext
}

func (w *WalletClient) Start(context.Context) {
	log.Info("evm wallet client is stopping...")
	w.loopApproveProposal()
}

func (w *WalletClient) Stop(context.Context) {
	w.cancel()
}

func NewWallet(event eventbus.Bus, tss suite.TssService, voterContract layer2.VoterContract, stateDB *state.ContractState, walletState *state.EvmWalletState) *WalletClient {
	client, err := ethclient.Dial(config.AppConfig.L2Rpc)
	utils.Assert(err)

	currentNonce := &atomic.Uint64{}
	nonce, _ := voterContract.TssNonce()
	if nonce != nil {
		currentNonce.Store(nonce.Uint64())
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &WalletClient{
		BaseWallet:          wallet.NewBaseWallet(stateDB, voterContract),
		ctx:                 ctx,
		cancel:              cancel,
		event:               event,
		client:              client,
		submitterPrivateKey: config.L2PrivateKey,
		submitter:           crypto.PubkeyToAddress(config.L2PrivateKey.PublicKey),
		pendingTx:           sync.Map{},
		chainID:             atomic.Int64{},
		walletState:         walletState,
		tss:                 tss,
		notify:              make(chan struct{}, 1024),
		currentVoterNonce:   currentNonce,
		submitTaskQueue:     pool.NewTaskPool[uint64](),
		operationsQueue:     pool.NewTaskPool[uint64](),
		txContext:           sync.Map{},
	}
}

func (w *WalletClient) BalanceOf(erc20Token, owner common.Address) (*big.Int, error) {
	erc20, err := contracts.NewERC20(erc20Token, w.client)
	if err != nil {
		return nil, err
	}

	return erc20.BalanceOf(nil, owner)
}

func (w *WalletClient) BalanceAt(owner common.Address) (*big.Int, error) {
	return w.client.BalanceAt(context.Background(), owner, nil)
}

func (w *WalletClient) ChainID(ctx context.Context) (*big.Int, error) {
	if w.chainID.Load() == 0 {
		var (
			chainID *big.Int
			err     error
		)

		err = backoff.Retry(
			func() error {
				chainID, err = w.client.ChainID(ctx)
				return err
			},
			backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 3),
		)
		if err != nil {
			return nil, fmt.Errorf("chainID error: %w", err)
		}

		w.chainID.Store(chainID.Int64())
	}

	return big.NewInt(w.chainID.Load()), nil
}

func (w *WalletClient) SendSingedTx(ctx context.Context, tx *types.Transaction) error {
	return w.sendTransaction(ctx, tx)
}

func (w *WalletClient) EstimateGasAPI(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := w.client.EstimateGas(ctx, msg)
	if err != nil {
		err = errors.Join(ErrEstimateGas, wrapError(err))
		gasLimit = defaultGasLimit

		if errors.Is(err, ErrInsufficientFunds) {
			err = errors.Join(fmt.Errorf("insufficient funds for sender: %v", msg.From), err)
		}
	}

	return gasLimit, err
}

func (w *WalletClient) EstimateGas(ctx context.Context, account, contractAddress common.Address, data []byte) (uint64, error) {
	// Estimate GasPrice
	// gasPrice := opt.GasPrice
	price, err := w.client.SuggestGasPrice(ctx)
	if err != nil {
		return 0, fmt.Errorf("SuggestGasPrice error: %w, account address: %v", err, account)
	}

	msg := ethereum.CallMsg{
		From:     account,
		To:       &contractAddress,
		Data:     data,
		GasPrice: price,
	}

	return w.EstimateGasAPI(ctx, msg)
}

func (w *WalletClient) BuildUnsignTx(
	ctx context.Context,
	account, contractAddress common.Address,
	value *big.Int, calldata []byte,
	Operations *db.Operations,
	EvmWithdraw *db.EvmWithdraw,
	EvmConsolidation *db.EvmConsolidation,
) (*types.Transaction, error) {
	head, err := w.client.HeaderByNumber(ctx, nil)
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
	gasTipCap, err := w.client.SuggestGasTipCap(ctx)
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

	gas, err := w.EstimateGasAPI(ctx, msg)
	if err != nil {
		return nil, wrapError(err)
	}

	// extend gas limit 20%
	gasLimit := decimal.NewFromUint64(gas).Mul(decimal.NewFromFloat(1.2))

	nextNonce, err := w.client.PendingNonceAt(ctx, account)
	if err != nil {
		return nil, wrapError(err)
	}

	//latestNonce := w.nonce.Load()
	//if latestNonce >= nextNonce {
	//	nextNonce = latestNonce + 1
	//	w.nonce.Store(nextNonce)
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

	err = w.walletState.CreateTx(
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

func (w *WalletClient) newTx(tx *db.EvmTransaction) (*types.Transaction, error) {
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

	signature, err := crypto.Sign(signer.Hash(unSignedTx).Bytes(), w.submitterPrivateKey)
	if err != nil {
		return nil, err
	}

	return unSignedTx.WithSignature(signer, signature)
}

// SpeedSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) speedSendOrderTx(ctx context.Context, oldOrderTx *db.EvmTransaction) (*types.Transaction, error) {
	signedTx, err := w.newTx(oldOrderTx)
	if err != nil {
		return nil, wrapError(err)
	}

	newTxHash := signedTx.Hash()

	w.pendingTx.Store(newTxHash, true)
	defer w.pendingTx.Delete(newTxHash)

	jsonData, err := signedTx.MarshalJSON()
	if err != nil {
		return signedTx, wrapError(err)
	}

	head, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return signedTx, wrapError(err)
	}

	err = w.walletState.CreateTx(nil,
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

	err = w.sendTransaction(ctx, signedTx)

	return signedTx, err
}

func (w *WalletClient) sendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := w.client.SendTransaction(ctx, tx)
	if err != nil {
		err = errors.Join(ErrSendTransaction, wrapError(err))
	}

	dbErr := w.walletState.UpdatePendingTx(tx.Hash())

	time.Sleep(2 * time.Second)

	return errors.Join(err, dbErr)
}

// SpeedSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) SpeedSendOrderTx(ctx context.Context, oldOrderTx *db.EvmTransaction) (common.Hash, *types.Receipt, error) {
	if oldOrderTx.Status == db.Pending {
		signedTx, err := w.speedSendOrderTx(ctx, oldOrderTx)
		if err != nil {
			return common.Hash{}, nil, err
		}

		r, err := w.waitTxSuccess(ctx, signedTx.Hash())

		return signedTx.Hash(), r, err
	}

	return common.Hash{}, nil, nil
}

// AgainSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) AgainSendOrderTx(ctx context.Context, tx *db.EvmTransaction) (common.Hash, *types.Receipt, error) {
	orderTx := &types.Transaction{}

	err := orderTx.UnmarshalJSON(tx.TxJsonData)
	if err != nil {
		return common.Hash{}, nil, wrapError(err)
	}

	txHash := orderTx.Hash()
	_, ok := w.pendingTx.Load(txHash)

	if ok {
		return txHash, nil, ErrTxPending
	}

	is := w.IsOnline(ctx, txHash)

	if is {
		return txHash, nil, fmt.Errorf("tx already online: %w or %w", ErrTxPending, ErrTxCompleted)
	}

	w.pendingTx.Store(txHash, true)
	defer w.pendingTx.Delete(txHash)

	err = w.sendTransaction(ctx, orderTx)
	if err != nil {
		if errors.Is(err, ErrIntrinsicGasTooLow) || errors.Is(err, ErrReplacement) || errors.Is(err, ErrAlreadyKnown) {
			_ = w.walletState.UpdateFailTx(txHash, err)
			return w.SpeedSendOrderTx(ctx, tx)
		}

		if errors.Is(err, ErrNonceTooLow) {
			_ = w.walletState.UpdateFailTx(txHash, err) // todo
		}

		return txHash, nil, err
	}

	r, err := w.waitTxSuccess(ctx, txHash)

	return txHash, r, err
}

func (w *WalletClient) RawTxBytes(ctx context.Context, tx *types.Transaction) []byte {
	chainID, _ := w.ChainID(ctx)
	signer := types.LatestSignerForChainID(chainID)

	return signer.Hash(tx).Bytes()
}

func (w *WalletClient) TransactionWithSignature(ctx context.Context, tx *types.Transaction, signature []byte) (*types.Transaction, error) {
	chainID, err := w.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer := types.LatestSignerForChainID(chainID)

	return tx.WithSignature(signer, signature)
}

func (w *WalletClient) SendTransactionWithSignature(ctx context.Context, tx *types.Transaction, signature []byte) error {
	chainID, err := w.ChainID(ctx)
	if err != nil {
		return err
	}
	signer := types.LatestSignerForChainID(chainID)

	tx, err = tx.WithSignature(signer, signature)
	if err != nil {
		return fmt.Errorf("tx with signature: %w", err)
	}

	return w.sendTransaction(ctx, tx)
}

func (w *WalletClient) WaitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	w.pendingTx.Store(txHash, true)
	defer w.pendingTx.Delete(txHash)

	return w.waitTxSuccess(ctx, txHash)
}

func (w *WalletClient) waitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	begin := time.Now()
	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		r, err := w.client.TransactionReceipt(ctx, txHash)
		// TransactionReceipt can is nil
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(time.Second)

				count--
			} else {
				return nil, errors.Join(fmt.Errorf("call TransactionReceipt method error, txHash: %s", txHash), err, ErrWallet)
			}
		} else {
			dbErr := w.walletState.UpdateBookedTx(txHash)
			return r, dbErr
		}
	}

	return nil, ErrTxFoundTimeOut
}

func (w *WalletClient) IsOnChain(ctx context.Context, txHash common.Hash) bool {
	receipt, err := w.client.TransactionReceipt(ctx, txHash)
	if receipt != nil && err == nil {
		return true
	}

	return false
}

func (w *WalletClient) IsPending(ctx context.Context, txHash common.Hash) bool {
	_, ok := w.pendingTx.Load(txHash)
	if ok {
		return ok
	}

	_, isPending, _ := w.client.TransactionByHash(ctx, txHash)

	return isPending
}

// IsOnline = IsPending + IsOnChain.
func (w *WalletClient) IsOnline(ctx context.Context, txHash common.Hash) bool {
	tx, isPending, _ := w.client.TransactionByHash(ctx, txHash)
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
