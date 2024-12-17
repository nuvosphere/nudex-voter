package types

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/crypto"
)

const (
	ChainBitcoin = iota
	ChainEthereum
	ChainSolana
	ChainSui
)

// https://docs.expand.network/ids/chain-ids
const (
	ChainIdBitcoin  = 1800
	ChainIdEthereum = 1
	ChainIdSolana   = 900
	ChainIdSui      = 101
)

const (
	CoinTypeBTC = 0
	CoinTypeEVM = 60
	CoinTypeSOL = 501
	CoinTypeSUI = 784
)

const (
	AssetTypeMain = iota
	AssetTypeErc20
)

var ErrCoinType = fmt.Errorf("error coin type")

func GetCoinTypeByChain(chain uint8) int {
	switch chain {
	case ChainBitcoin:
		return CoinTypeBTC
	case ChainEthereum:
		return CoinTypeEVM
	case ChainSolana:
		return CoinTypeSOL
	case ChainSui:
		return CoinTypeSUI
	default:
		panic(ErrCoinType)
	}
}

func GetChainByCoinType(coinType int) uint8 {
	switch coinType {
	case CoinTypeBTC:
		return ChainBitcoin
	case CoinTypeEVM:
		return ChainEthereum
	case CoinTypeSOL:
		return ChainSolana
	case CoinTypeSUI:
		return ChainSui
	default:
		panic(ErrCoinType)
	}
}

func GetCurveTypeByChain(chain uint8) crypto.CurveType {
	switch chain {
	case ChainBitcoin, ChainEthereum:
		return crypto.ECDSA
	case ChainSolana, ChainSui:
		return crypto.EDDSA
	default:
		panic("implement me")
	}
}

func GetCurveTypeByCoinType(coinType int) crypto.CurveType {
	return GetCurveTypeByChain(GetChainByCoinType(coinType))
}

const (
	DataTypeTssKeygenMsg    = "TssKeygenMsg"
	DataTypeTssSignMsg      = "TssSignMsg"
	DataTypeTssReSharingMsg = "TssReSharingMsg"
	DataTypeSignDeposit     = "SignDeposit"
	DataTypeSignWithdrawal  = "SignWithdrawal"
)

const (
	GenKeySessionType            = "GenerateKeySession"
	ReShareGroupSessionType      = "ReShareGroupSession"
	SignTaskSessionType          = "SignTaskSession"
	SignTestOperationSessionType = "SignTestOperationSession"
	SignTestTxSessionType        = "SignTestTxSession"
	SignOperationSessionType     = "SignOperationSession"
	SignTxSessionType            = "SignTxSession"
)
