package btc

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
)

type TxContext struct {
	c txClient
}

func (w *WalletClient) processWithdrawTxSign(task *db.WithdrawalTask) {
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
}
