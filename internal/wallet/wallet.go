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

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const defaultGasLimit = 1000000

type Wallet struct {
	client              *ethclient.Client
	tssPublicKey        ecdsa.PublicKey
	submitterPrivateKey ecdsa.PrivateKey
	submitter           common.Address
	pendingTx           sync.Map // txHash: bool
	chainID             atomic.Int64
	evmState            *state.EvmState
}

func NewWallet(l2url string, submitterPrivateKey ecdsa.PrivateKey) *Wallet {
	client, err := ethclient.Dial(l2url)
	utils.Assert(err)

	return &Wallet{
		client:              client,
		submitterPrivateKey: submitterPrivateKey,
		submitter:           crypto.PubkeyToAddress(submitterPrivateKey.PublicKey),
		pendingTx:           sync.Map{},
		chainID:             atomic.Int64{},
	}
}

func (s *Wallet) SetTssPublicKey(tssPublicKey ecdsa.PublicKey) {
	s.tssPublicKey = tssPublicKey
}

func (s *Wallet) Address(coinType, account uint32, index uint8) common.Address {
	return GenerateEthAddressByPath(ECPoint(&s.tssPublicKey), coinType, account, index)
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
	err := s.client.SendTransaction(ctx, tx)
	if err != nil {
		err = errors.Join(ErrSendTransaction, wrapError(err))
	}

	return err
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

func (s *Wallet) BuildUnsignTx(ctx context.Context, account, contractAddress common.Address, value *big.Int, calldata []byte) (*types.Transaction, error) {
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

	err = s.evmState.CreateTx(nil, account, decimal.NewFromUint64(nextNonce), jsonData, calldata, tx.Hash(), head.Number.Uint64())

	return tx, err
}

func (s *Wallet) RawTxBytes(ctx context.Context, tx *types.Transaction) []byte {
	chainID, _ := s.ChainID(ctx)
	signer := types.LatestSignerForChainID(chainID)

	return signer.Hash(tx).Bytes()
}

func (s *Wallet) WaitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	s.pendingTx.Store(txHash, true)
	defer s.pendingTx.Delete(txHash)

	return s.waitTxSuccess(ctx, txHash)
}

func (s *Wallet) waitTxSuccess(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	begin := time.Now()
	defer func() {
		log.Info("waitTxSuccess", "duration_ms", time.Since(begin).Milliseconds())
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
			return r, err
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
		if errors.Is(err, wrap) {
			return errors.Join(err, wrap, ErrWallet)
		}
	}

	return err
}
