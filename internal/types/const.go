package types

import (
	"fmt"
)

const (
	ChainBitcoin = iota
	ChainEthereum
	ChainSolana
	ChainSui
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
