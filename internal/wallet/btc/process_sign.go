package btc

import (
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func (w *WalletClient) ChainType() uint8 {
	return types.ChainBitcoin
}

func (w *WalletClient) Verify(reqId uint64, signDigest string, ExtraData []byte) error {
	// TODO implement me
	panic("implement me")
}

func (w *WalletClient) ReceiveSignature(res *suite.SignRes) {
	if res.Type == types.SignTxSessionType {
		w.processTxSignResult(res)
	}
}
