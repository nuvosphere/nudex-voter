package tss

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/types"
)

type SignContext struct {
	chainType          uint8
	localData          LocalPartySaveData
	keyDerivationDelta *big.Int
}

func (c *SignContext) Address() string {
	return c.localData.Address(c.ChainType())
}

func (c *SignContext) ChainType() uint8 {
	return c.chainType
}

func (m *Scheduler) Sign(requester types.Requester, msg []byte) error {
	panic("todo")
}
