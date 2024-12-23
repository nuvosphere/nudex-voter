package btc

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func (w *WalletClient) ChainType() uint8 {
	return types.ChainBitcoin
}

func (w *WalletClient) Verify(reqId uint64, signDigest string, ExtraData []byte) error {
	ctx, ok := w.txContext.Load(reqId)
	if !ok {
		return fmt.Errorf("tx id %d is not found", reqId)
	}
	// txCtx, is := ctx.(*TxContext)
	_, is := ctx.(*TxContext)
	if !is {
		return fmt.Errorf("tx id %d is not TxContext", reqId)
	}

	// todo
	return nil
}

func (w *WalletClient) ReceiveSignature(res *suite.SignRes) {
	if res.Type == types.SignTxSessionType {
		w.processTxSignResult(res)
	}
}
