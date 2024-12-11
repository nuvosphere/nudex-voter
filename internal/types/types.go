package types

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/samber/lo"
)

func DefaultPartners() Participants {
	// todo online contact get address list
	return lo.Map(
		config.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return ethCrypto.PubkeyToAddress(*pubKey) },
	)
}

func PartyKey(ec crypto.CurveType, participants Participants, address common.Address) *big.Int {
	key := new(big.Int).Add(address.Big(), big.NewInt(int64(ec)))
	return key.Add(key, participants.GroupID().Big())
}

func PartyID(ec crypto.CurveType, participants Participants, address common.Address) string {
	return PartyKey(ec, participants, address).Text(16)
}

type Requester interface {
	Post(hash, signature []byte)
	ChainType() int
	SendTransaction(to string, amount int64, allowHighFees bool) ([]byte, error)
}

type Signer interface {
	Sign(requester Requester, msg []byte) error
}
