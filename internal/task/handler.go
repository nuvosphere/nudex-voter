package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/types"
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

	if ts.currentTask != nil {
		log.Debugf("Current task %d is already running", ts.currentTask.TaskId)
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
		ts.currentTask = &dbTask
		err := ts.handleCreateWalletTask(ctx, dbTask)
		if err != nil {
			log.Errorf("Handle create wallet task %d error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
			ts.currentTask = nil
		}
	case TaskTypeDeposit:
		err := ts.handleDepositTask(dbTask)
		if err != nil {
			log.Errorf("Handle deposit task %d error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
			ts.currentTask = nil
		}
	case TaskTypeWithdraw:
		err := ts.handleWithdrawTask(dbTask)
		if err != nil {
			log.Errorf("Handle withdraw task %d error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
			ts.currentTask = nil
		}
	default:
		log.Warnf("Parse task  %d error, not know task type, description: %s", dbTask.TaskId, dbTask.Description)
		ts.currentTask = nil
	}

}

func (ts *TaskService) handleCreateWalletTask(ctx context.Context, dbTask db.Task) error {
	parts := strings.Split(dbTask.Description, ":")
	if len(parts) < 4 {
		return fmt.Errorf("parse task %d to create wallet task error, description: %s", dbTask.TaskId, dbTask.Description)
	}
	account, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return fmt.Errorf("parse task %d to create wallet task error, description: %s, %v", dbTask.TaskId, dbTask.Description, err)
	}
	createWalletTask := types.CreateWalletTask{
		TaskId:  dbTask.TaskId,
		User:    parts[1],
		Account: account,
		Chain:   parts[3],
	}
	return ts.Tss.HandleSignCreateAccount(ctx, createWalletTask)
}

func (ts *TaskService) handleDepositTask(dbTask db.Task) error {
	return nil
}

func (ts *TaskService) handleWithdrawTask(dbTask db.Task) error {
	return nil
}
