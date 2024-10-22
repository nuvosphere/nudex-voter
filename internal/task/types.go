package task

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss"
)

type TaskService struct {
	state       *state.State
	dbm         *db.DatabaseManager
	Tss         *tss.TSSService
	currentTask *db.Task
}

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdraw
)

type CreateWalletTask struct {
	taskId  uint64
	user    string
	account uint64
	chain   string
	index   string
}
