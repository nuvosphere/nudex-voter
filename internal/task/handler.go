package task

import (
	"context"
	"errors"

	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (ts *TaskService) checkTasks(ctx context.Context) {
	if !ts.Tss.IsPrepared() {
		log.Debug("Party not init, skip task check")
		return
	}

	if ts.state.TssState.CurrentSubmitter != ts.Tss.LocalSubmitter() {
		log.Debugf("Current submitter is %v, not self %v", ts.state.TssState.CurrentSubmitter, ts.Tss.LocalSubmitter())
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

	ts.state.TssState.CurrentTask = &dbTask

	err = ts.Tss.HandleSignPrepare(ctx, dbTask)
	if err != nil {
		log.Errorf("Handle sign prepare error for task %x err: %v", dbTask.Context, err)

		ts.state.TssState.CurrentTask = nil
	}
}
