package task

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss"
)

type Service struct {
	state *state.State
	dbm   *db.DatabaseManager
	Tss   *tss.Service
}
