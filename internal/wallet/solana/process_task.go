package solana

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) evenLoop() {
	taskEvent := w.event.Subscribe(eventbus.EventTask{})

	go func() {
		select {
		case <-w.ctx.Done():
			log.Info("evm wallet receive task event done")

		case detailTask := <-taskEvent:
			val, ok := detailTask.(db.DetailTask)
			if ok {
				if val.ChainType() == w.ChainType() {
					w.AddTask(val)
				}
			}
		}
	}()
}

func (w *WalletClient) processTask(detailTask pool.Task[uint64]) {
	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		// userAddress := w.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
		_ = w.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)

	case *db.DepositTask:
		// todo

	case *db.WithdrawalTask:
		// todo

	default:
		log.Errorf("unhandled default case")
	}
}

func (w *WalletClient) submitTask(detailTask pool.Task[uint64]) {
	w.event.Publish(eventbus.EventSubmitTask{}, detailTask)
}
