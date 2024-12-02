package types

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/samber/lo"
)

func DefaultPartners() Participants {
	// todo online contact get address list
	return lo.Map(
		config.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}

func PartyKey(ec CurveType, participants Participants, address common.Address) *big.Int {
	key := new(big.Int).Add(address.Big(), big.NewInt(int64(ec)))
	return key.Add(key, participants.GroupID().Big())
}

func PartyID(ec CurveType, participants Participants, address common.Address) string {
	return PartyKey(ec, participants, address).Text(16)
}
