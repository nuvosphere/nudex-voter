package wallet

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/davecgh/go-spew/spew"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/nuvosphere/nudex-voter/internal/types"
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

func TestGenerateAddress(t *testing.T) {
	data := base58.Decode("5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc")
	pk, pubKey := edwards.PrivKeyFromBytes(data)
	assert.NotNil(t, pk)
	assert.NotNil(t, pubKey)
	point := crypto.NewECPointNoCurveCheck(tss.Edwards(), pubKey.X, pubKey.Y)
	address := GenerateAddressByPath(point, types.CoinTypeSOL, 1, 1)
	t.Log("address", address)
	assert.Equal(t, strings.ToLower("jxK4DrMrDevCn7UXGhiJPjT36e4XP12cJLFDvP9uvxX"), strings.ToLower(address))

	client := rpc.New(rpc.DevNet_RPC)
	pubkey := solana.MustPublicKeyFromBase58("jxK4DrMrDevCn7UXGhiJPjT36e4XP12cJLFDvP9uvxX")
	out, err := client.GetBalance(
		context.Background(),
		pubkey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	lamportsOnAccount := new(big.Float).SetUint64(out.Value)
	// Convert lamports to sol:
	solBalance := new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

	// WARNING: this is not a precise conversion.
	fmt.Println("◎", solBalance.Text('f', 10))
}
