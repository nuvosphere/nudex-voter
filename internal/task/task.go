package task

import (
	"context"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss"
	log "github.com/sirupsen/logrus"
)

func NewTaskService(state *state.State, dbm *db.DatabaseManager, tss *tss.Service) *Service {
	return &Service{
		state: state,
		dbm:   dbm,
		Tss:   tss,
	}
}

func (ts *Service) Start(ctx context.Context) {
	go ts.loop(ctx)

	<-ctx.Done()
	log.Info("TaskService is stopping...")
}

func (ts *Service) loop(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info("Tss keygen checker stopping...")
				return
			case <-ticker.C:
				ts.checkTasks(ctx)
			}
		}
	}()
}
