package evm

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (m *WalletClient) receiveL2TaskLoop() {
	taskEvent := m.event.Subscribe(eventbus.EventTask{})

	go func() {
		select {
		case <-m.ctx.Done():
			log.Info("evm wallet receive task event done")

		case detailTask := <-taskEvent:
			val, ok := detailTask.(db.DetailTask)
			if ok {
				if val.ChainType() == m.ChainType() {
					m.taskQueue.Add(val)
				}
			}
		}
	}()
}

func (m *WalletClient) processTask(detailTask pool.Task[uint64]) {
	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		// userAddress := m.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
		_ = m.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)

	case *db.DepositTask:
		// todo

	case *db.WithdrawalTask:
		// todo

	default:
		log.Errorf("unhandled default case")
	}
}

func (m *WalletClient) submitTask(detailTask pool.Task[uint64]) {
	m.event.Publish(eventbus.EventSubmitTask{}, detailTask)
}
