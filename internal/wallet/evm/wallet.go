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
	"github.com/samber/lo"
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
	pendingTx           sync.Map // txHash: EvmTx
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
	log.Info("evm wallet client is starting...")
	w.tss.RegisterTssClient(w)
	w.receiveL2TaskLoop()
	w.receiveSubmitTaskLoop()
	w.loopProcessOperation()
	w.tickerRetrySendTx()
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
	return w.client.BalanceAt(w.ctx, owner, nil)
}

func (w *WalletClient) ChainID() (*big.Int, error) {
	if w.chainID.Load() == 0 {
		var (
			chainID *big.Int
			err     error
		)

		err = backoff.Retry(
			func() error {
				chainID, err = w.client.ChainID(w.ctx)
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

func (w *WalletClient) SendSingedTx(tx *types.Transaction) error {
	err := w.sendTransaction(tx)

	if errors.Is(err, ErrIntrinsicGasTooLow) || errors.Is(err, ErrReplacement) {
		dbTX, is := w.pendingTx.Load(tx.Hash())
		if is {
			_, err = w.SpeedSendTx(dbTX.(*db.EvmTransaction))
		}
	}

	if errors.Is(err, ErrNonceTooLow) {
		dbTX, is := w.pendingTx.Load(tx.Hash())
		if is {
			_, err = w.BuildNewDbTx(dbTX.(*db.EvmTransaction))
		}
	}

	return err
}

func (w *WalletClient) EstimateGasAPI(msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := w.client.EstimateGas(w.ctx, msg)
	if err != nil {
		err = errors.Join(ErrEstimateGas, wrapError(err))
		gasLimit = defaultGasLimit

		if errors.Is(err, ErrInsufficientFunds) {
			err = errors.Join(fmt.Errorf("insufficient funds for sender: %v", msg.From), err)
		}
	}

	return gasLimit, err
}

func (w *WalletClient) EstimateGas(account, contractAddress common.Address, data []byte) (uint64, error) {
	// Estimate GasPrice
	// gasPrice := opt.GasPrice
	price, err := w.client.SuggestGasPrice(w.ctx)
	if err != nil {
		return 0, fmt.Errorf("SuggestGasPrice error: %w, account address: %v", err, account)
	}

	msg := ethereum.CallMsg{
		From:     account,
		To:       &contractAddress,
		Data:     data,
		GasPrice: price,
	}

	return w.EstimateGasAPI(msg)
}

func (w *WalletClient) BuildUnsignTx(
	account, contractAddress common.Address,
	value *big.Int, calldata []byte,
	Operations *db.Operations,
	EvmWithdraw *db.EvmWithdraw,
	EvmConsolidation *db.EvmConsolidation,
) (*types.Transaction, error) {
	head, err := w.client.HeaderByNumber(w.ctx, nil)
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
	gasTipCap, err := w.client.SuggestGasTipCap(w.ctx)
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

	gas, err := w.EstimateGasAPI(msg)
	if err != nil {
		return nil, wrapError(err)
	}

	// extend gas limit 20%
	gasLimit := decimal.NewFromUint64(gas).Mul(decimal.NewFromFloat(1.2))
	var nextNonce uint64 = 0
	for {
		nextNonce, err = w.client.PendingNonceAt(w.ctx, account)
		if err != nil {
			return nil, wrapError(err)
		}

		latestNonce, err := w.walletState.LatestNonce(nil, account)
		if err != nil {
			return nil, err
		}

		if nextNonce > latestNonce.BigInt().Uint64() {
			break
		}
		time.Sleep(200 * time.Millisecond)
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

	dbTX, err := w.walletState.CreateTx(
		nil,
		tx.Hash(),
		jsonData,
		decimal.NewFromUint64(nextNonce),
		head.Number.Uint64(),
		Operations,
		EvmWithdraw,
		EvmConsolidation,
	)
	if err != nil {
		w.pendingTx.Store(tx.Hash(), dbTX)
	}

	return tx, err
}

func (w *WalletClient) BuildNewDbTx(oldTx *db.EvmTransaction) (*types.Transaction, error) {
	tx := oldTx.Tx()
	sender := w.From(tx)

	return w.BuildUnsignTx(sender, *tx.To(), tx.Value(), tx.Data(), oldTx.Operations, oldTx.EvmWithdraw, oldTx.EvmConsolidation)
}

func (w *WalletClient) ExtendTxGas(tx *db.EvmTransaction) (*types.Transaction, error) {
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

	return unSignedTx, nil
}

func (w *WalletClient) SignOperationNewTx(unSignedTx *types.Transaction) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(unSignedTx.ChainId())
	signature, err := crypto.Sign(signer.Hash(unSignedTx).Bytes(), w.submitterPrivateKey)
	if err != nil {
		return nil, err
	}
	return unSignedTx.WithSignature(signer, signature)
}

// BuildSpeedSendTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) BuildSpeedSendTx(oldOrderTx *db.EvmTransaction) (*types.Transaction, error) {
	UnSignedTx, err := w.ExtendTxGas(oldOrderTx)
	if err != nil {
		return nil, wrapError(err)
	}

	w.pendingTx.Store(UnSignedTx.Hash(), nil)
	signedTx, err := w.sign(oldOrderTx.Type, UnSignedTx) // todo
	if err != nil {
		return nil, wrapError(err)
	}

	newTxHash := signedTx.Hash()

	jsonData, err := signedTx.MarshalJSON()
	if err != nil {
		return signedTx, wrapError(err)
	}

	head, err := w.client.HeaderByNumber(w.ctx, nil)
	if err != nil {
		return signedTx, wrapError(err)
	}

	dbTX, err := w.walletState.CreateTx(nil,
		newTxHash,
		jsonData,
		oldOrderTx.TxNonce,
		head.Number.Uint64(),
		oldOrderTx.Operations,
		oldOrderTx.EvmWithdraw,
		oldOrderTx.EvmConsolidation,
	)
	if err != nil {
		w.pendingTx.Store(newTxHash, dbTX)
	}

	return signedTx, wrapError(err)
}

func (w *WalletClient) sign(ty int, tx *types.Transaction) (*types.Transaction, error) {
	switch ty {
	case db.TaskTypeWithdrawal:
		// todo
	case db.TaskTypeConsolidation:
		// todo
	case db.TaskTypeOperations:
		return w.SignOperationNewTx(tx)
	default:
		panic("unhandled default case")
	}
	return nil, fmt.Errorf("unknown task type %d", ty)
}

func (w *WalletClient) sendTransaction(tx *types.Transaction) error {
	err := w.client.SendTransaction(w.ctx, tx)
	if err != nil {
		err = errors.Join(ErrSendTransaction, wrapError(err))
		_, is := lo.Find(failErrorList, func(item error) bool { return errors.Is(err, item) })
		if !is {
			dbErr := w.walletState.UpdateFailTx(tx.Hash(), err)
			err = errors.Join(err, dbErr)
		}
		return err
	}

	return w.walletState.UpdatePendingTx(tx.Hash())
}

// SpeedSendTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) SpeedSendTx(oldOrderTx *db.EvmTransaction) (common.Hash, error) {
	signedTx, err := w.BuildSpeedSendTx(oldOrderTx)
	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), w.SendSingedTx(signedTx)
}

// AgainSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (w *WalletClient) AgainSendOrderTx(tx *db.EvmTransaction) (common.Hash, error) {
	orderTx := &types.Transaction{}
	utils.Assert(orderTx.UnmarshalJSON(tx.TxJsonData))
	txHash := tx.TxHash
	err := w.SendSingedTx(orderTx)
	if err != nil {
		return txHash, err
	}

	_, err = w.WaitTxSuccess(txHash)

	return txHash, err
}

func (w *WalletClient) RawTxBytes(tx *types.Transaction) []byte {
	chainID, err := w.ChainID()
	utils.Assert(err)
	signer := types.LatestSignerForChainID(chainID)

	return signer.Hash(tx).Bytes()
}

func (w *WalletClient) From(tx *types.Transaction) common.Address {
	chainID, err := w.ChainID()
	utils.Assert(err)
	signer := types.LatestSignerForChainID(chainID)
	address, err := signer.Sender(tx)
	utils.Assert(err)
	return address
}

func (w *WalletClient) TransactionWithSignature(tx *types.Transaction, signature []byte) (*types.Transaction, error) {
	chainID, err := w.ChainID()
	if err != nil {
		return nil, err
	}
	signer := types.LatestSignerForChainID(chainID)

	return tx.WithSignature(signer, signature)
}

func (w *WalletClient) SendTransactionWithSignature(tx *types.Transaction, signature []byte) error {
	chainID, err := w.ChainID()
	if err != nil {
		return err
	}
	signer := types.LatestSignerForChainID(chainID)

	tx, err = tx.WithSignature(signer, signature)
	if err != nil {
		return fmt.Errorf("tx with signature: %w", err)
	}

	return w.SendSingedTx(tx)
}

func (w *WalletClient) WaitTxSuccess(txHash common.Hash) (*types.Receipt, error) {
	begin := time.Now()
	defer func() {
		log.Infof("WaitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		r, err := w.client.TransactionReceipt(w.ctx, txHash)
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

func (w *WalletClient) IsOnChain(txHash common.Hash) bool {
	receipt, err := w.client.TransactionReceipt(w.ctx, txHash)
	if receipt != nil && err == nil {
		return true
	}

	return false
}

func (w *WalletClient) IsPending(txHash common.Hash) bool {
	_, ok := w.pendingTx.Load(txHash)
	if ok {
		return ok
	}

	_, isPending, _ := w.client.TransactionByHash(w.ctx, txHash)

	return isPending
}

// IsOnline = IsPending + IsOnChain.
func (w *WalletClient) IsOnline(txHash common.Hash) bool {
	tx, isPending, _ := w.client.TransactionByHash(w.ctx, txHash)
	return isPending || tx != nil
}

func (w *WalletClient) IsCanProcess(txHash common.Hash) bool {
	_, isPending := w.pendingTx.Load(txHash)
	if !isPending {
		tx, _, _ := w.client.TransactionByHash(w.ctx, txHash)
		return tx == nil
	}

	return false
}

func wrapError(err error) error {
	for _, wrap := range wrapErrorList {
		if utils.ContainErr(err, wrap) {
			return errors.Join(err, wrap, ErrWallet)
		}
	}

	return err
}

func (w *WalletClient) tickerRetrySendTx() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("evm wallet tickerRetrySendTx done")
			case <-ticker.C:
				w.tickerRetryUpdateTx()
			}
		}
	}()
}

func (w *WalletClient) tickerRetryUpdateTx() {
	txs, _ := w.walletState.PendingBlockchainTransactions(nil)
	group := sync.WaitGroup{}
	for _, tx := range txs {
		group.Add(1)
		go func() {
			if w.IsCanProcess(tx.TxHash) {
				w.pendingTx.Store(tx.TxHash, &tx)
				_, err := w.AgainSendOrderTx(&tx)
				if err != nil {
					log.Infof("evm wallet AgainSendOrderTx: tx hash:%v, err:%v", tx.TxHash, err)
				}
			}
			group.Done()
		}()
	}
	group.Wait()
}
