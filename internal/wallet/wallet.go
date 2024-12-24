package wallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/state"
	ty "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet/bip44"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func init() {
	address.RegisterAddress(ty.CoinTypeEVM, EthAddress)
}

func EthAddress(p *tssCrypto.ECPoint) string {
	return crypto.PubkeyToAddress(*p.ToECDSAPubKey()).String()
}

const defaultGasLimit = 1000000

type Wallet struct {
	client              *ethclient.Client
	tssPublicKey        ecdsa.PublicKey
	submitterPrivateKey *ecdsa.PrivateKey
	submitter           common.Address
	pendingTx           sync.Map // txHash: bool
	chainID             atomic.Int64
	state               *state.EvmWalletState
}

func NewWallet() *Wallet {
	client, err := ethclient.Dial(config.AppConfig.L2Rpc)
	utils.Assert(err)

	return &Wallet{
		client:              client,
		submitterPrivateKey: config.L2PrivateKey,
		submitter:           crypto.PubkeyToAddress(config.L2PrivateKey.PublicKey),
		pendingTx:           sync.Map{},
		chainID:             atomic.Int64{},
	}
}

func (s *Wallet) SetTssPublicKey(tssPublicKey ecdsa.PublicKey) {
	s.tssPublicKey = tssPublicKey
}

func (s *Wallet) Address(coinType, account uint32, index uint8) common.Address {
	address := address.GenerateAddressByPath(bip44.ECPoint(&s.tssPublicKey), coinType, account, index)
	return common.HexToAddress(address)
}

func (s *Wallet) HotAddressOfCoin(coinType uint32) common.Address {
	return s.Address(coinType, 0, 0)
}

func (s *Wallet) BalanceOf(erc20Token, owner common.Address) (*big.Int, error) {
	erc20, err := contracts.NewERC20(erc20Token, s.client)
	if err != nil {
		return nil, err
	}

	return erc20.BalanceOf(nil, owner)
}

func (s *Wallet) BalanceAt(owner common.Address) (*big.Int, error) {
	return s.client.BalanceAt(context.Background(), owner, nil)
}

func (s *Wallet) ChainID(ctx context.Context) (*big.Int, error) {
	if s.chainID.Load() == 0 {
		var (
			chainID *big.Int
			err     error
		)

		err = backoff.Retry(
			func() error {
				chainID, err = s.client.ChainID(ctx)
				return err
			},
			backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 3),
		)
		if err != nil {
			return nil, fmt.Errorf("chainID error: %w", err)
		}

		s.chainID.Store(chainID.Int64())
	}

	return big.NewInt(s.chainID.Load()), nil
}

func (s *Wallet) SendSingedTx(ctx context.Context, tx *types.Transaction) error {
	return s.sendTransaction(ctx, tx)
}

func (s *Wallet) EstimateGasAPI(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := s.client.EstimateGas(ctx, msg)
	if err != nil {
		err = errors.Join(ErrEstimateGas, wrapError(err))
		gasLimit = defaultGasLimit

		if errors.Is(err, ErrInsufficientFunds) {
			err = errors.Join(fmt.Errorf("insufficient funds for sender: %v", msg.From), err)
		}
	}

	return gasLimit, err
}

func (s *Wallet) EstimateGas(ctx context.Context, account, contractAddress common.Address, data []byte) (uint64, error) {
	// Estimate GasPrice
	// gasPrice := opt.GasPrice
	price, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return 0, fmt.Errorf("SuggestGasPrice error: %w, account address: %v", err, account)
	}

	msg := ethereum.CallMsg{
		From:     account,
		To:       &contractAddress,
		Data:     data,
		GasPrice: price,
	}

	return s.EstimateGasAPI(ctx, msg)
}

func (s *Wallet) BuildUnsignTx(
	ctx context.Context,
	account, contractAddress common.Address,
	value *big.Int, calldata []byte,
	Operations *db.Operations,
	EvmWithdraw *db.EvmWithdraw,
	EvmConsolidation *db.EvmConsolidation,
) (*types.Transaction, error) {
	head, err := s.client.HeaderByNumber(ctx, nil)
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
	gasTipCap, err := s.client.SuggestGasTipCap(ctx)
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

	gas, err := s.EstimateGasAPI(ctx, msg)
	if err != nil {
		return nil, wrapError(err)
	}

	// extend gas limit 20%
	gasLimit := decimal.NewFromUint64(gas).Mul(decimal.NewFromFloat(1.2))

	nextNonce, err := s.client.PendingNonceAt(ctx, account)
	if err != nil {
		return nil, wrapError(err)
	}

	//latestNonce := s.nonce.Load()
	//if latestNonce >= nextNonce {
	//	nextNonce = latestNonce + 1
	//	s.nonce.Store(nextNonce)
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

	err = s.state.CreateTx(
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

func (s *Wallet) newTx(tx *db.EvmTransaction) (*types.Transaction, error) {
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

	signature, err := crypto.Sign(signer.Hash(unSignedTx).Bytes(), s.submitterPrivateKey)
	if err != nil {
		return nil, err
	}

	return unSignedTx.WithSignature(signer, signature)
}

// SpeedSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (s *Wallet) speedSendOrderTx(ctx context.Context, oldOrderTx *db.EvmTransaction) (*types.Transaction, error) {
	signedTx, err := s.newTx(oldOrderTx)
	if err != nil {
		return nil, wrapError(err)
	}

	newTxHash := signedTx.Hash()

	s.pendingTx.Store(newTxHash, true)
	defer s.pendingTx.Delete(newTxHash)

	jsonData, err := signedTx.MarshalJSON()
	if err != nil {
		return signedTx, wrapError(err)
	}

	head, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return signedTx, wrapError(err)
	}

	err = s.state.CreateTx(nil,
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

	err = s.sendTransaction(ctx, signedTx)

	return signedTx, err
}

func (s *Wallet) sendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := s.client.SendTransaction(ctx, tx)
	if err != nil {
		err = errors.Join(ErrSendTransaction, wrapError(err))
	}

	dbErr := s.state.UpdatePendingTx(tx.Hash())

	time.Sleep(2 * time.Second)

	return errors.Join(err, dbErr)
}

// SpeedSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (s *Wallet) SpeedSendOrderTx(ctx context.Context, oldOrderTx *db.EvmTransaction) (common.Hash, *types.Receipt, error) {
	if oldOrderTx.Status == db.Pending {
		signedTx, err := s.speedSendOrderTx(ctx, oldOrderTx)
		if err != nil {
			return common.Hash{}, nil, err
		}

		r, err := s.waitTxSuccess(ctx, signedTx.Hash())

		return signedTx.Hash(), r, err
	}

	return common.Hash{}, nil, nil
}

// AgainSendOrderTx
// 1. resend order tx
// 2. speed resend.
func (s *Wallet) AgainSendOrderTx(ctx context.Context, tx *db.EvmTransaction) (common.Hash, *types.Receipt, error) {
	orderTx := &types.Transaction{}

	err := orderTx.UnmarshalJSON(tx.TxJsonData)
	if err != nil {
		return common.Hash{}, nil, wrapError(err)
	}

	txHash := orderTx.Hash()
	_, ok := s.pendingTx.Load(txHash)

	if ok {
		return txHash, nil, ErrTxPending
	}

	is := s.IsOnline(ctx, txHash)

	if is {
		return txHash, nil, fmt.Errorf("tx already online: %w or %w", ErrTxPending, ErrTxCompleted)
	}

	s.pendingTx.Store(txHash, true)
	defer s.pendingTx.Delete(txHash)

	err = s.sendTransaction(ctx, orderTx)
	if err != nil {
		if errors.Is(err, ErrIntrinsicGasTooLow) || errors.Is(err, ErrReplacement) || errors.Is(err, ErrAlreadyKnown) {
			_ = s.state.UpdateFailTx(txHash, err)
			return s.SpeedSendOrderTx(ctx, tx)
		}

		if errors.Is(err, ErrNonceTooLow) {
			_ = s.state.UpdateFailTx(txHash, err) // todo
		}

		return txHash, nil, err
	}

	r, err := s.waitTxSuccess(ctx, txHash)

	return txHash, r, err
}

func (s *Wallet) RawTxBytes(ctx context.Context, tx *types.Transaction) []byte {
	chainID, _ := s.ChainID(ctx)
	signer := types.LatestSignerForChainID(chainID)

	return signer.Hash(tx).Bytes()
}

func (s *Wallet) TransactionWithSignature(ctx context.Context, tx *types.Transaction, signature []byte) (*types.Transaction, error) {
	chainID, err := s.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer := types.LatestSignerForChainID(chainID)

	return tx.WithSignature(signer, signature)
}

func (s *Wallet) SendTransactionWithSignature(ctx context.Context, tx *types.Transaction, signature []byte) error {
	chainID, err := s.ChainID(ctx)
	if err != nil {
		return err
	}
	signer := types.LatestSignerForChainID(chainID)

	tx, err = tx.WithSignature(signer, signature)
	if err != nil {
		return fmt.Errorf("tx with signature: %w", err)
	}

	return s.sendTransaction(ctx, tx)
}

func (s *Wallet) WaitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	s.pendingTx.Store(txHash, true)
	defer s.pendingTx.Delete(txHash)

	return s.waitTxSuccess(ctx, txHash)
}

func (s *Wallet) waitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	begin := time.Now()
	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		r, err := s.client.TransactionReceipt(ctx, txHash)
		// TransactionReceipt can is nil
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(time.Second)

				count--
			} else {
				return nil, errors.Join(fmt.Errorf("call TransactionReceipt method error, txHash: %s", txHash), err, ErrWallet)
			}
		} else {
			dbErr := s.state.UpdateBookedTx(txHash)
			return r, dbErr
		}
	}

	return nil, ErrTxFoundTimeOut
}

func (s *Wallet) IsOnChain(ctx context.Context, txHash common.Hash) bool {
	receipt, err := s.client.TransactionReceipt(ctx, txHash)
	if receipt != nil && err == nil {
		return true
	}

	return false
}

func (s *Wallet) IsPending(ctx context.Context, txHash common.Hash) bool {
	_, ok := s.pendingTx.Load(txHash)
	if ok {
		return ok
	}

	_, isPending, _ := s.client.TransactionByHash(ctx, txHash)

	return isPending
}

// IsOnline = IsPending + IsOnChain.
func (s *Wallet) IsOnline(ctx context.Context, txHash common.Hash) bool {
	tx, isPending, _ := s.client.TransactionByHash(ctx, txHash)
	return isPending || tx != nil
}

func wrapError(err error) error {
	for _, wrap := range wrapErrorList {
		if utils.ContainErr(err, wrap) {
			return errors.Join(err, wrap, ErrWallet)
		}
	}

	return err
}
