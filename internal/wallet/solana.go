package wallet

import (
	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
)

// https://github.com/gagliardetto/solana-go

func SolanaAddress(p *crypto.ECPoint) string {
	pubkey := edwards.NewPublicKey(p.X(), p.Y())
	return base58.Encode(pubkey.Serialize())
}
