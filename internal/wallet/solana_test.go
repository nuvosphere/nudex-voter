package wallet

import (
	"strings"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/stretchr/testify/assert"
)

func TestSolanaAddress(t *testing.T) {
	data := base58.Decode("31ry6kknaZvLyi4yf1ZCVvmoMXtvg8FkpZFtKqJtV6xmmkgii3cKuuBXS8FWgYxW6R5eUDgeDaZuxo2mEVaEriwr")
	pk, pubKey := edwards.PrivKeyFromBytes(data)
	assert.NotNil(t, pk)
	assert.NotNil(t, pubKey)

	address := strings.ToLower("12bLXhcf6gzHHp6pzWyGHUE2NUpXVpbLva8ikTCd6AtG")
	t.Log(base58.Encode(pubKey.Serialize()))
	assert.Equal(t, address, strings.ToLower(base58.Encode(pubKey.Serialize())))

	point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
	solanaAddress := SolanaAddress(point)
	assert.Equal(t, address, strings.ToLower(solanaAddress))
}
