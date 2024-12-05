package wallet

import (
	"fmt"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func GenerateCompressedBTCAddress(p *crypto.ECPoint) (string, error) {
	return btcAddress(NewPublicKeyOfBtc(p).SerializeCompressed())
}

func GenerateUnCompressedBTCAddress(p *crypto.ECPoint) (string, error) {
	return btcAddress(NewPublicKeyOfBtc(p).SerializeUncompressed())
}

func btcAddress(serializedPubKey []byte) (string, error) {
	addr, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}
	return addr.EncodeAddress(), nil
}

func NewPublicKeyOfBtc(p *crypto.ECPoint) *btcec.PublicKey {
	var (
		x = &btcec.FieldVal{}
		y = &btcec.FieldVal{}
	)
	x.SetByteSlice(p.X().Bytes())
	y.SetByteSlice(p.Y().Bytes())
	return btcec.NewPublicKey(x, y)
}
