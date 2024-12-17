package wallet

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/codec"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/patrickmn/go-cache"
)

type BaseWallet struct {
	taskQueue          *pool.Pool[uint64]
	voterContract      layer2.VoterContract
	discussedTaskCache *cache.Cache
	stateDB            *state.ContractState
}

func (m *BaseWallet) GetTask(taskID uint64) (pool.Task[uint64], error) {
	t := m.taskQueue.Get(taskID)
	if t != nil {
		return t, nil
	}

	task, err := m.stateDB.GetUnCompletedTask(taskID)
	//todo
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return m.GetOnlineTask(taskID)
	//}
	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task.DetailTask(), err
}

func (m *BaseWallet) GetOnlineTask(taskId uint64) (pool.Task[uint64], error) {
	t, err := m.voterContract.Tasks(taskId)
	if err != nil {
		return nil, err
	}

	detailTask := codec.DecodeTask(t.Id, t.Context)

	baseTask := db.Task{
		TaskId:    t.Id,
		TaskType:  detailTask.Type(),
		Context:   t.Context,
		Submitter: t.Submitter.Hex(),
		Status:    int(t.State),
	}
	detailTask.SetBaseTask(baseTask)

	return detailTask, nil
}

func (m *BaseWallet) IsDiscussed(taskID uint64) bool {
	_, ok := m.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	if !ok {
		ok, _ = m.voterContract.IsTaskCompleted(taskID)
	}

	return ok
}

func (m *BaseWallet) AddDiscussedTask(taskID uint64) {
	m.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}
