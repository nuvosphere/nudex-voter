package evm

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func (c *WalletClient) ChainType() uint8 {
	return types.ChainEthereum
}

func (c *WalletClient) Verify(reqId *big.Int, signDigest string, ExtraData []byte) error {
	// TODO implement me
	panic("implement me")
}

func (c *WalletClient) ReceiveSignature(res *suite.SignRes) {
	// TODO implement me
	panic("implement me")
}
