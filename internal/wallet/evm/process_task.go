package evm

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
						c.taskQueue.Add(v)
						// todo
					case db.Pending:
						// todo withdraw
						c.taskQueue.Add(v)

					case db.Completed, db.Failed:
						c.taskQueue.Remove(v.TaskID())
						// todo
					default:
						log.Errorf("taskID: %d, invalid task state : %v", v.TaskID(), v.Status())
					}
				}
			}
		}
	}()
}

func (c *WalletClient) processTask(detailTask pool.Task[uint64]) {
	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		// userAddress := c.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
		_ = c.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)

		// send to evm operation
		// c.submitTask()

	case *db.DepositTask:
		// todo
		// c.submitTask()

	case *db.WithdrawalTask:
		// todo
		// c.submitTask()

	default:
		log.Errorf("unhandled default case")
	}
}

func (c *WalletClient) submitTask(detailTask pool.Task[uint64]) {
	c.event.Publish(eventbus.EventSubmitTask{}, detailTask)
}
