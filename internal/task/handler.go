package task

import (
	"context"
	"errors"
	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

}
