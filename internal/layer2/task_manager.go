package layer2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

type TaskManager interface {
	GetLatestTask() (contracts.ITaskManagerTask, error)
	IsTaskCompleted(taskId uint64) (bool, error)
	GetTaskState(taskId uint64) (uint8, error)
	GetUncompletedTasks() ([]contracts.ITaskManagerTask, error)
	EncodeSubmitTask(submitter common.Address, context []byte) []byte
	NextTaskId() (uint64, error)
	Tasks(taskId *big.Int) (contracts.ITaskManagerTask, error)
	TaskSubmitter() (common.Address, error)
	EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte
}

func (l *Layer2Listener) GetLatestTask() (contracts.ITaskManagerTask, error) {
	return l.taskManager.GetLatestTask(nil)
}

func (l *Layer2Listener) IsTaskCompleted(taskId uint64) (bool, error) {
	state, err := l.GetTaskState(taskId)
	return state == db.Completed, err
}

func (l *Layer2Listener) GetTaskState(taskId uint64) (uint8, error) {
	return l.taskManager.GetTaskState(nil, taskId)
}

func (l *Layer2Listener) GetUncompletedTasks() ([]contracts.ITaskManagerTask, error) {
	return l.taskManager.GetUncompletedTasks(nil)
}

func (l *Layer2Listener) NextTaskId() (uint64, error) {
	return l.taskManager.NextTaskId(nil)
}

func (l *Layer2Listener) Tasks(taskId *big.Int) (contracts.ITaskManagerTask, error) {
	return l.taskManager.Tasks(nil, taskId)
}

func (l *Layer2Listener) TaskSubmitter() (common.Address, error) {
	return l.taskManager.TaskSubmitter(nil)
}

func (l *Layer2Listener) EncodeSubmitTask(submitter common.Address, context []byte) []byte {
	return contracts.EncodeFun(contracts.TaskManagerContractMetaData.ABI, "submitTask", submitter, context)
}

func (l *Layer2Listener) EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte {
	return contracts.EncodeFun(contracts.TaskManagerContractMetaData.ABI, "markTaskCompleted", taskId, result)
}

func (l *Layer2Listener) EncodeMarkTaskCompletedBatch(taskIds []*big.Int, results [][]byte) []byte {
	return contracts.EncodeFun(contracts.TaskManagerContractMetaData.ABI, "markTaskCompleted_Batch", taskIds, results)
}
