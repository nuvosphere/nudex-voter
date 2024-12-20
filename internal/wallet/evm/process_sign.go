package evm

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
)

func (w *WalletClient) ChainType() uint8 {
	return types.ChainEthereum
}

func (w *WalletClient) Verify(reqId *big.Int, signDigest string, ExtraData []byte) error {
	// TODO implement me
	panic("implement me")
}

func (w *WalletClient) ReceiveSignature(res *suite.SignRes) {
	switch res.Type {
	case types.SignOperationSessionType:
		op := w.operationsQueue.Get(res.SeqId)
		if op != nil {
			operations := op.(*Operations)
			operations.Signature = res.Signature
			w.processOperationSignResult(operations)
			lo.ForEach(operations.Operation, func(item contracts.Operation, _ int) { w.AddDiscussedTask(item.TaskId) })
			w.operationsQueue.RemoveTopN(operations.TaskID() - 1)
		}

	case types.SignTxSessionType:
		w.processTxSignResult(res)
	}
}
