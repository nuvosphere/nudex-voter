package state

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
)

type State struct {
	EventBus *EventBus

	dbm *db.DatabaseManager
}

// InitializeState initializes the state by reading from the DB
func InitializeState(dbm *db.DatabaseManager) *State {
	return &State{
		EventBus: NewEventBus(),
		dbm:      dbm,
	}
}
