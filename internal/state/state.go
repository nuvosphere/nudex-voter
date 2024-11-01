package state

import (
	"errors"
	"github.com/nuvosphere/nudex-voter/internal/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type State struct {
	EventBus Bus

	dbm *db.DatabaseManager

	btcHeadMu sync.RWMutex

	btcHeadState BtcHeadState

	TssState TssState
}

func (s *State) Bus() Bus {
	return s.EventBus
}

// InitializeState initializes the state by reading from the DB
func InitializeState(dbm *db.DatabaseManager) *State {
	// Load states from db when start up
	var (
		latestBtcBlock    db.BtcBlock
		unconfirmBtcQueue []*db.BtcBlock
		sigBtcQueue       []*db.BtcBlock

		currentSubmitter     string
		participantAddresses []string
		L2BlockNumber        uint64
	)
	btcLightDb := dbm.GetBtcLightDB()
	relayerDb := dbm.GetRelayerDB()

	loadData := func(db *gorm.DB, dest interface{}, query string, args ...interface{}) {
		if err := db.Where(query, args...).Find(dest).Error; err != nil {
			log.Warnf("Failed to load data: %v", err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(5)

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

	go func() {
		defer wg.Done()

		var submitterChosen db.SubmitterChosen
		err := relayerDb.Order("block_number DESC").First(&submitterChosen).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Warnf("Failed to load latest submitter rotation: %v", err)
			}
		} else {
			L2BlockNumber = submitterChosen.BlockNumber
			currentSubmitter = submitterChosen.Submitter
		}
	}()

	go func() {
		defer wg.Done()

		var participants []db.Participant

		err := relayerDb.Find(&participants).Error
		if err != nil {
			for _, participant := range participants {
				participantAddresses = append(participantAddresses, participant.Address)
			}
		}
	}()

	wg.Wait()

	return &State{
		EventBus: NewBus(),
		dbm:      dbm,

		btcHeadState: BtcHeadState{
			Latest:         latestBtcBlock,
			UnconfirmQueue: unconfirmBtcQueue,
			SigQueue:       sigBtcQueue,
		},
		TssState: TssState{
			CurrentSubmitter: currentSubmitter,
			Participants:     participantAddresses,
			BlockNumber:      L2BlockNumber,
		},
	}
}
