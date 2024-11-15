package state

import (
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// BtcHeadState to manage BTC head.
type BtcHeadState struct {
	Latest         db.BtcBlock
	UnconfirmQueue []*db.BtcBlock // status in 'unconfirm', 'confirmed'
	SigQueue       []*db.BtcBlock // status in 'signing', 'pending'
}

type State struct {
	eventBus     eventbus.Bus
	dbm          *db.DatabaseManager
	btcHeadMu    sync.RWMutex
	btcHeadState BtcHeadState
}

func (s *State) Bus() eventbus.Bus {
	return s.eventBus
}

// InitializeState initializes the state by reading from the DB.
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

	if err := btcLightDb.Where("status = ?", "processed").Order("height desc").First(&latestBtcBlock).Error; err != nil {
		log.Warnf("Failed to load latest processed Btc Block: %v", err)
	}

	loadData(btcLightDb, &unconfirmBtcQueue, "status in (?)", []string{"unconfirm", "confirmed"})
	loadData(btcLightDb, &sigBtcQueue, "status in (?)", []string{"signing", "pending"})

	return &State{
		eventBus: eventbus.NewBus(),
		dbm:      dbm,

		btcHeadState: BtcHeadState{
			Latest:         latestBtcBlock,
			UnconfirmQueue: unconfirmBtcQueue,
			SigQueue:       sigBtcQueue,
		},
	}
}
