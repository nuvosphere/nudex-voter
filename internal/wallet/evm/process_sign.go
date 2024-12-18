package evm

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
)

func (c *WalletClient) ChainType() uint8 {
	return types.ChainEthereum
}

func (c *WalletClient) Verify(reqId *big.Int, signDigest string, ExtraData []byte) error {
	// TODO implement me
	panic("implement me")
}

func (c *WalletClient) ReceiveSignature(res *suite.SignRes) {
	switch res.Type {
	case types.SignOperationSessionType:
		op := c.operationsQueue.Get(res.SeqId)
		if op != nil {
			operations := op.(*Operations)
			operations.Signature = res.Signature
			c.processOperationSignResult(operations)
			lo.ForEach(operations.Operation, func(item contracts.Operation, _ int) { c.AddDiscussedTask(item.TaskId) })
			c.operationsQueue.RemoveTopN(operations.TaskID() - 1)
		}

	case types.SignTxSessionType:
		// todo
	}
}
