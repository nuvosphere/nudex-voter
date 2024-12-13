package tss

import (
	"math/big"
)

type SignerContext struct {
	chainType          uint8
	localData          LocalPartySaveData
	keyDerivationDelta *big.Int
}

func (c *SignerContext) Address() string {
	return c.localData.Address(c.ChainType())
}

func (c *SignerContext) ChainType() uint8 {
	return c.chainType
}
