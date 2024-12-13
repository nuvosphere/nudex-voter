package wallet

import (
	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func HotAddressOfBtc(masterPubKey *crypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeBTC, 0, 0)
}

func HotAddressOfEth(masterPubKey *crypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeEVM, 0, 0)
}

func HotAddressOfSolana(masterPubKey *crypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeSOL, 0, 0)
}

func HotAddressOfSui(masterPubKey *crypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeSUI, 0, 0)
}
