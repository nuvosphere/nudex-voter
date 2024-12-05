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
	type Args struct {
		PK              string
		ExpectedAddress string
	}

	args := []Args{
		{
			PK:              "31ry6kknaZvLyi4yf1ZCVvmoMXtvg8FkpZFtKqJtV6xmmkgii3cKuuBXS8FWgYxW6R5eUDgeDaZuxo2mEVaEriwr",
			ExpectedAddress: "12bLXhcf6gzHHp6pzWyGHUE2NUpXVpbLva8ikTCd6AtG",
		},
		{
			PK:              "47pXBoR4PA6K6VgQY8VCfBbtukffQhfedXYdqZ27TPB6MQdEhF74znuRrvbe3VjP2pFqByhgxZNbeGZ3v7H1emNe",
			ExpectedAddress: "1LPWKkn3KYD2kFp3dZquTBb7hFabKM9KEFAQ2wi76mt",
		},
		{
			PK:              "5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc",
			ExpectedAddress: "12nV3gkyFCAYmScaAsBgU1LyyEFQer3f9x2FRvirDSJ",
		},
	}

	for _, arg := range args {
		data := base58.Decode(arg.PK)
		pk, pubKey := edwards.PrivKeyFromBytes(data)
		assert.NotNil(t, pk)
		assert.NotNil(t, pubKey)

		assert.Equal(t, strings.ToLower(arg.ExpectedAddress), strings.ToLower(base58.Encode(pubKey.Serialize())))

		point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
		solanaAddress := SolanaAddress(point)
		assert.Equal(t, strings.ToLower(arg.ExpectedAddress), strings.ToLower(solanaAddress))
	}
}
