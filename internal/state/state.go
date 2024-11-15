package state

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/samber/lo"
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
	TssState     TssState
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

		currentSubmitter     string
		participantAddresses []common.Address
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
			participantAddresses = lo.Map(
				participants,
				func(item db.Participant, _ int) common.Address {
					return common.HexToAddress(item.Address)
				},
			)
		}
	}()

	wg.Wait()

	return &State{
		eventBus: eventbus.NewBus(),
		dbm:      dbm,

		btcHeadState: BtcHeadState{
			Latest:         latestBtcBlock,
			UnconfirmQueue: unconfirmBtcQueue,
			SigQueue:       sigBtcQueue,
		},
		TssState: TssState{
			CurrentSubmitter: common.HexToAddress(currentSubmitter),
			Participants:     participantAddresses,
			BlockNumber:      L2BlockNumber,
		},
	}
}
