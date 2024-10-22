package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func (ts *TaskService) checkTasks(ctx context.Context) {
	if ts.Tss.Party == nil || ts.Tss.LocalPartySaveData == nil || ts.Tss.LocalPartySaveData.ECDSAPub == nil {
		log.Debug("Party not init, skip task check")
		return
	}

	if ts.state.TssState.CurrentSubmitter != ts.Tss.Address.Hex() {
		log.Debugf("Current submitter is %s, not self %s", ts.state.TssState.CurrentSubmitter, ts.Tss.Address.Hex())
		return
	}

	var dbTask db.Task
	err := ts.dbm.GetRelayerDB().Order("block_number DESC").First(&dbTask).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Debug("No local task found")
			return
		} else {
			log.Errorf("Get task from db error: %v", err)
			return
		}
	}
	parts := strings.Split(dbTask.Description, ":")
	if len(parts) < 1 {
		log.Errorf("Parse task %d error, description: %s", dbTask.TaskId, dbTask.Description)
		return
	}
	taskType, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Errorf("Parse task  %d type error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
		return
	}

	switch taskType {
	case TaskTypeUnknown:
		log.Warnf("Parse task  %d type error, not know task type, description: %s", dbTask.TaskId, dbTask.Description)
		return
	case TaskTypeCreateWallet:
		err := ts.handleCreateWalletTask(dbTask)
		if err != nil {
			log.Errorf("Handle create wallet task %d error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
		}
	case TaskTypeDeposit:
		err := ts.handleDepositTask(dbTask)
		if err != nil {
			log.Errorf("Handle deposit task %d error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
		}
	case TaskTypeWithdraw:
		err := ts.handleWithdrawTask(dbTask)
		if err != nil {
			log.Errorf("Handle withdraw task %d error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
		}
	default:
		log.Warnf("Parse task  %d error, not know task type, description: %s", dbTask.TaskId, dbTask.Description)
	}

}

func (ts *TaskService) handleCreateWalletTask(dbTask db.Task) error {
	parts := strings.Split(dbTask.Description, ":")
	if len(parts) < 4 {
		return fmt.Errorf("parse task %d to create wallet task error, description: %s", dbTask.TaskId, dbTask.Description)
	}
	account, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return fmt.Errorf("parse task %d to create wallet task error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
	}
	createWalletTask := CreateWalletTask{
		taskId:  dbTask.TaskId,
		user:    parts[1],
		account: account,
		chain:   parts[3],
		index:   parts[4],
	}
	log.Debugf("Parse create task: %v", createWalletTask)
	return nil
}

func (ts *TaskService) handleDepositTask(dbTask db.Task) error {
	return nil
}

func (ts *TaskService) handleWithdrawTask(dbTask db.Task) error {
	return nil
}
