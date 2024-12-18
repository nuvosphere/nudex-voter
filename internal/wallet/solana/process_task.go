package solana

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (c *WalletClient) receiveL2TaskLoop() {
	taskEvent := c.event.Subscribe(eventbus.EventTask{})

	go func() {
		select {
		case <-c.ctx.Done():
			log.Info("evm wallet receive task event done")
		case data := <-taskEvent: // from layer2 log scan
			log.Info("received task from layer2 log scan: ", data)
			switch v := data.(type) {
			case db.DetailTask:
				if v.ChainType() == c.ChainType() {
					switch v.Status() {
					case db.Created:
						c.AddTask(v)
						// todo
					case db.Pending:
						// todo withdraw
						c.AddTask(v)

					case db.Completed, db.Failed:
						c.RemoveTask(v.TaskID())
						// todo
					default:
						log.Errorf("taskID: %d, invalid task state : %v", v.TaskID(), v.Status())
					}
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
