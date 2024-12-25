package evm

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
)

func (w *WalletClient) ChainType() uint8 {
	return types.ChainEthereum
}

func (w *WalletClient) Verify(reqId uint64, signDigest string, ExtraData []byte) error {
	ctx, ok := w.pendingTx.Load(signDigest)
	if !ok {
		return fmt.Errorf("tx id %d is not found", reqId)
	}
	txCtx, is := ctx.(*TxContext)
	if !is {
		return fmt.Errorf("tx id %d is not TxContext", reqId)
	}

	if txCtx.TxHash().String() != signDigest {
		return fmt.Errorf("tx id %d hash does not match", reqId)
	}

	return nil
}

func (w *WalletClient) ReceiveSignature(res *suite.SignRes) {
	switch res.Type {
	case types.SignOperationSessionType:
		op := w.operationsQueue.Get(res.SeqId)
		if op != nil {
			operations := op.(*Operations)
			operations.Signature = res.Signature
			w.processOperationSignResult(operations)
			lo.ForEach(operations.Operation, func(item contracts.TaskOperation, _ int) { w.AddDiscussedTask(item.TaskId) })
			w.operationsQueue.RemoveTopN(operations.TaskID() - 1)
		}

	case types.SignTxSessionType:
		w.processTxSignResult(res)
	}
}

func (w *WalletClient) signTx(ctx *TxContext) error {
	switch ctx.dbTX.Type {
	case db.TaskTypeWithdrawal, db.TaskTypeConsolidation:
		// todo
		req := &suite.SignReq{
			SeqId:      ctx.SeqID(),
			Type:       types.SignTxSessionType,
			ChainType:  w.ChainType(),
			Signer:     ctx.dbTX.Sender.String(),
			DataDigest: ctx.TxHash().String(),
			SignData:   ctx.TxHash().Bytes(),
			ExtraData:  nil,
		}
		err := w.tss.Sign(req)
		if err != nil {
			return err
		}
		<-ctx.notify

	case db.TaskTypeOperations:
		sig, err := w.SignOperationNewTx(ctx.UnSignTx())
		if err != nil {
			return err
		}
		ctx.sig = sig
	default:
		panic("unhandled default case")
	}
	return fmt.Errorf("unknown task:%d, type %d", ctx.SeqID(), ctx.dbTX.Type)
}
