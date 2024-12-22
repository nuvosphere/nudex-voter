package solana

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) receiveL2TaskLoop() {
	taskEvent := w.event.Subscribe(eventbus.EventTask{})

	go func() {
		select {
		case <-w.ctx.Done():
			log.Info("evm wallet receive task event done")
		case data := <-taskEvent: // from layer2 log scan
			log.Info("received task from layer2 log scan: ", data)
			switch v := data.(type) {
			case db.DetailTask:
				if v.ChainType() == w.ChainType() {
					switch v.Status() {
					case db.Created:
						w.AddTask(v)
						w.processCreatedTask(v)
					case db.Pending:
						w.AddTask(v)
						w.processPendingTask(v)

					case db.Completed, db.Failed:
						w.RemoveTask(v.TaskID())
						w.txContext.Delete(v.TaskID())
					default:
						log.Errorf("taskID: %d, invalid task walletState : %v", v.TaskID(), v.Status())
					}
				}
			}
		}
	}()
}

func (w *WalletClient) processCreatedTask(detailTask pool.Task[uint64]) {
	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.AddressType)
		// userAddress := w.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
		_ = w.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)

		// send to evm operation
		// w.submitTask()

	case *db.DepositTask:
		// todo
		// w.submitTask()

	case *db.WithdrawalTask:
		w.processWithdrawTxSign(task)
	case *db.ConsolidationTask:
		// todo
		// w.submitTask()
	default:
		log.Errorf("unhandled default case")
	}
}

func (w *WalletClient) processPendingTask(detailTask pool.Task[uint64]) {
	switch task := detailTask.(type) {
	case *db.WithdrawalTask:
		// todo
		w.submitTask(task)
	case *db.ConsolidationTask:
		// todo
		w.submitTask(task)
	default:
		log.Errorf("unhandled default case")
	}
}

func (w *WalletClient) submitTask(detailTask pool.Task[uint64]) {
	w.event.Publish(eventbus.EventSubmitTask{}, detailTask)
}
