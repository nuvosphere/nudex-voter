package task

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (ts *TaskService) checkTasks(ctx context.Context) {
	if ts.Tss.LocalParty == nil || ts.Tss.LocalPartySaveData == nil || ts.Tss.LocalPartySaveData.ECDSAPub == nil {
		localPartySaveData, err := tss.LoadTSSData()
		if err != nil && localPartySaveData != nil {
			ts.Tss.LocalPartySaveData = localPartySaveData
		}
		log.Debug("Party not init, skip task check")
		return
	}

	if ts.state.TssState.CurrentSubmitter != ts.Tss.Address {
		log.Debugf("Current submitter is %v, not self %v", ts.state.TssState.CurrentSubmitter, ts.Tss.Address)
		return
	}

	if ts.state.TssState.CurrentTask != nil {
		log.Debugf("Current task %d is already running", ts.state.TssState.CurrentTask.TaskId)
		return
	}

	var dbTask db.Task
	err := ts.dbm.GetRelayerDB().Order("created_at DESC").First(&dbTask).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Debug("No local task found")
			return
		} else {
			log.Errorf("Get task from db error: %v", err)
			return
		}
	}
	buf := bytes.NewReader(dbTask.Context)
	var taskType int32
	if err := binary.Read(buf, binary.LittleEndian, &taskType); err != nil {
		log.Errorf("Parse task %d error, content: %x, ", dbTask.TaskId, dbTask.Context)
		return
	}

	switch taskType {
	case types.TaskTypeUnknown:
		log.Warnf("Parse task  %d type error, not know task type, context: %s", dbTask.TaskId, dbTask.Context)
		return
	case types.TaskTypeCreateWallet:
		ts.state.TssState.CurrentTask = &dbTask
		err := ts.handleCreateWalletTask(ctx, dbTask)
		if err != nil {
			log.Errorf("Handle create wallet task %d error, description: %s, %v", dbTask.TaskId, dbTask.Context, err)
		}
	case types.TaskTypeDeposit:
		ts.state.TssState.CurrentTask = &dbTask
		err := ts.handleDepositTask(ctx, dbTask)
		if err != nil {
			log.Errorf("Handle deposit task %d error, description: %s, %v", dbTask.TaskId, dbTask.Context, err)
		}
	case types.TaskTypeWithdraw:
		err := ts.handleWithdrawTask(dbTask)
		if err != nil {
			log.Errorf("Handle withdraw task %d error, description: %s, %v", dbTask.TaskId, dbTask.Context, err)
		}
	default:
		log.Warnf("Parse task  %d error, not know task type, description: %s", dbTask.TaskId, dbTask.Context)
	}

}

func (ts *TaskService) handleCreateWalletTask(ctx context.Context, dbTask db.Task) error {
	buf := bytes.NewReader(dbTask.Context)
	var taskType int32
	if err := binary.Read(buf, binary.LittleEndian, &taskType); err != nil {
		return err
	}
	createWalletTask := types.CreateWalletTask{TaskId: int32(dbTask.TaskId)}

	if err := binary.Read(buf, binary.LittleEndian, &createWalletTask.User); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &createWalletTask.Account); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &createWalletTask.Chain); err != nil {
		return err
	}
	return ts.Tss.HandleSignCreateAccount(ctx, createWalletTask)
}

func (ts *TaskService) handleDepositTask(ctx context.Context, dbTask db.Task) error {
	buf := bytes.NewReader(dbTask.Context)
	var taskType int32
	if err := binary.Read(buf, binary.LittleEndian, &taskType); err != nil {
		return err
	}
	depositTask := types.DepositTask{TaskId: int32(dbTask.TaskId)}

	if err := binary.Read(buf, binary.LittleEndian, &depositTask.Address); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &depositTask.Amount); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &depositTask.ChainId); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &depositTask.Token); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &depositTask.TxInfo); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &depositTask.ExtraInfo); err != nil {
		return err
	}
	return ts.Tss.HandleSignDeposit(ctx, depositTask)
}

func (ts *TaskService) handleWithdrawTask(dbTask db.Task) error {
	return nil
}
