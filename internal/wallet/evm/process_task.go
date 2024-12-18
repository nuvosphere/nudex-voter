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
					c.taskQueue.Add(v)
				}
			case *db.TaskUpdatedEvent: // todo
				switch v.State {
				case db.Pending:
					// todo withdraw
				case db.Completed, db.Failed:
					// todo
				default:
					log.Errorf("invalid task state : %v", v.State)
				}
				log.Infof("taskID: %d completed on blockchain", v.TaskId)
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

		//send to evm operation
		//c.submitTask()

	case *db.DepositTask:
		// todo
		//c.submitTask()

	case *db.WithdrawalTask:
		// todo
		//c.submitTask()

	default:
		log.Errorf("unhandled default case")
	}
}

func (c *WalletClient) submitTask(detailTask pool.Task[uint64]) {
	c.event.Publish(eventbus.EventSubmitTask{}, detailTask)
}
