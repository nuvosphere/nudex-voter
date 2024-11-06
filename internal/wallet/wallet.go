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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
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
	nonce               atomic.Uint64
}

func NewWallet(l2url string, tssPublicKey ecdsa.PublicKey, submitterPrivateKey ecdsa.PrivateKey) *Wallet {
	client, err := ethclient.Dial(l2url)
	utils.Assert(err)

	return &Wallet{
		client:              client,
		tssPublicKey:        tssPublicKey,
		submitterPrivateKey: submitterPrivateKey,
		submitter:           crypto.PubkeyToAddress(submitterPrivateKey.PublicKey),
		pendingTx:           sync.Map{},
		chainID:             atomic.Int64{},
		nonce:               atomic.Uint64{},
	}
}

func Bip44DerivationPath(coinType, user, account uint32) DerivePath {
	// https://learnblockchain.cn/2018/09/28/hdwallet
	// m / purpose' / coin' / account' / change / address_index
	// coin list https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	return DerivePath(fmt.Sprintf("m/44/%d/%d/0/%d", coinType, user, account))
}

func GenerateAddressByPath(masterPubKey ecdsa.PublicKey, coinType, user, account uint32) common.Address {
	bip44Path := Bip44DerivationPath(coinType, user, account)

	p, err := bip44Path.ToParams()
	utils.Assert(err)

	_, extendKey, err := DerivingPubKeyFromPath(masterPubKey, p.Indexes())
	utils.Assert(err)

	return crypto.PubkeyToAddress(extendKey.PublicKey)
}

func (s *Wallet) Address(coinType, user, account uint32) common.Address {
	return GenerateAddressByPath(s.tssPublicKey, coinType, user, account)
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
		chainID, err := s.client.ChainID(ctx)
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

func (s *Wallet) EstimateGas(ctx context.Context, contractAddress common.Address, data []byte) (uint64, error) {
	account := s.submitter
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

func (s *Wallet) BuildUnsignTx(ctx context.Context, contractAddress common.Address, value *big.Int, calldata []byte) (*types.Transaction, error) {
	head, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, wrapError(err)
	}

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
		From:      s.submitter,
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

	nextNonce, err := s.client.PendingNonceAt(ctx, s.submitter)
	if err != nil {
		return nil, wrapError(err)
	}

	latestNonce := s.nonce.Load()
	if latestNonce >= nextNonce {
		nextNonce = latestNonce + 1
		s.nonce.Store(nextNonce)
	}

	baseTx := &types.DynamicFeeTx{
		To:        &contractAddress,
		Nonce:     nextNonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit.BigInt().Uint64(),
		Value:     value,
		Data:      calldata,
	}

	return types.NewTx(baseTx), nil
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
