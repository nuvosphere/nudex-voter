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

		case detailTask := <-taskEvent:
			val, ok := detailTask.(db.DetailTask)
			if ok {
				if val.ChainType() == c.ChainType() {
					c.taskQueue.Add(val)
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

	case *db.DepositTask:
		// todo

	case *db.WithdrawalTask:
		// todo

	default:
		log.Errorf("unhandled default case")
	}
}

func (c *WalletClient) submitTask(detailTask pool.Task[uint64]) {
	c.event.Publish(eventbus.EventSubmitTask{}, detailTask)
}
