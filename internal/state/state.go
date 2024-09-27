package state

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type State struct {
	EventBus *EventBus

	dbm *db.DatabaseManager

	btcHeadMu sync.RWMutex

	btcHeadState BtcHeadState
}

// InitializeState initializes the state by reading from the DB
func InitializeState(dbm *db.DatabaseManager) *State {
	// Load states from db when start up
	var (
		latestBtcBlock    db.BtcBlock
		unconfirmBtcQueue []*db.BtcBlock
		sigBtcQueue       []*db.BtcBlock
	)
	btcLightDb := dbm.GetBtcLightDB()

	loadData := func(db *gorm.DB, dest interface{}, query string, args ...interface{}) {
		if err := db.Where(query, args...).Find(dest).Error; err != nil {
			log.Warnf("Failed to load data: %v", err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(11)

	go func() {
		defer wg.Done()
		if err := btcLightDb.Where("status = ?", "processed").Order("height desc").First(&latestBtcBlock).Error; err != nil {
			log.Warnf("Failed to load latest processed Btc Block: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		loadData(btcLightDb, &unconfirmBtcQueue, "status in (?)", []string{"unconfirm", "confirmed"})
	}()

	go func() {
		defer wg.Done()
		loadData(btcLightDb, &sigBtcQueue, "status in (?)", []string{"signing", "pending"})
	}()

	return &State{
		EventBus: NewEventBus(),
		dbm:      dbm,

		btcHeadState: BtcHeadState{
			Latest:         latestBtcBlock,
			UnconfirmQueue: unconfirmBtcQueue,
			SigQueue:       sigBtcQueue,
		},
	}
}
