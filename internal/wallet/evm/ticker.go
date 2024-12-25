package evm

import (
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
)

func (w *WalletClient) tickerLoopUnCompletedTasks() {
	if !w.VoterContract().IsSyncing() {
		tasks, _ := w.ContractState().GetUnCompletedTasks()
		for _, task := range tasks {
			w.event.Publish(eventbus.EventTask{}, task.DetailTask())
		}
	}
}
