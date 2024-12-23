package wallet

import (
	"fmt"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/codec"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/patrickmn/go-cache"
)

type BaseWallet struct {
	stateDB            *state.ContractState
	voterContract      layer2.VoterContract
	discussedTaskCache *cache.Cache
	taskQueue          *pool.Pool[uint64] // receive l2 task
}

func NewBaseWallet(stateDB *state.ContractState, voterContract layer2.VoterContract) *BaseWallet {
	return &BaseWallet{
		stateDB:            stateDB,
		voterContract:      voterContract,
		discussedTaskCache: cache.New(5*time.Minute, 10*time.Minute),
		taskQueue:          pool.NewTaskPool[uint64](),
	}
}

func (w *BaseWallet) IsProd() bool {
	return false
}

func (w *BaseWallet) VoterContract() layer2.VoterContract {
	return w.voterContract
}

func (w *BaseWallet) AddTask(task pool.Task[uint64]) {
	w.taskQueue.Add(task)
}

func (w *BaseWallet) RemoveTask(taskId uint64) {
	w.taskQueue.Remove(taskId)
}

func (w *BaseWallet) GetTask(taskID uint64) (pool.Task[uint64], error) {
	t := w.taskQueue.Get(taskID)
	if t != nil {
		return t, nil
	}

	task, err := w.stateDB.GetUnCompletedTask(taskID)
	//todo
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return w.GetOnlineTask(taskID)
	//}
	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task.DetailTask(), err
}

func (w *BaseWallet) GetOnlineTask(taskId uint64) (pool.Task[uint64], error) {
	t, err := w.voterContract.Tasks(taskId)
	if err != nil {
		return nil, err
	}

	detailTask := codec.DecodeTask(t.Id, t.Result)

	baseTask := db.Task{
		TaskId:    t.Id,
		TaskType:  detailTask.Type(),
		Context:   t.Result,
		Submitter: t.Submitter.Hex(),
		State:     int(t.State),
	}
	detailTask.SetBaseTask(baseTask)

	return detailTask, nil
}

func (w *BaseWallet) IsDiscussed(taskID uint64) bool {
	_, ok := w.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	if !ok {
		ok, _ = w.voterContract.IsTaskCompleted(taskID)
	}

	return ok
}

func (w *BaseWallet) AddDiscussedTask(taskID uint64) {
	w.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (w *BaseWallet) GetAddressBalance(chainID uint64, minAmount uint64) []db.AddressBalance {
	address, _ := w.stateDB.GetAddressBalanceByCondition(chainID, minAmount)
	return address
}
