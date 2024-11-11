package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type TSSService struct {
	isPrepared   atomic.Bool
	privateKey   *ecdsa.PrivateKey // submit
	localAddress common.Address    // submit = partyID
	proposer     common.Address

	p2p       p2p.P2PService
	state     *state.State
	scheduler *Scheduler[int64]
	cache     *cache.Cache

	layer2Listener *layer2.Layer2Listener
	dbm            *db.DatabaseManager
	taskReceive    chan any // task

	threshold          *atomic.Int64
	partners           []common.Address
	localParty         *keygen.LocalParty
	localPartySaveData *keygen.LocalPartySaveData

	// eventbus channel
	tssMsgCh       <-chan any
	partyAddOrRmCh <-chan any

	rw sync.RWMutex
}

func (t *TSSService) IsCompleted(taskID int32) bool {
	_, ok := t.cache.Get(fmt.Sprintf("%d", taskID))
	return ok
}

func (t *TSSService) AddCompletedTask(taskID int32) {
	t.cache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (t *TSSService) LocalSubmitter() common.Address {
	return t.localAddress
}

func (t *TSSService) IsPrepared() bool {
	return t.isPrepared.Load()
}

func (t *TSSService) PostTask(task any) {
	t.taskReceive <- task
}

func (t *TSSService) sigEndLoop(ctx context.Context) {
	out := t.scheduler.sigInToOut

	for {
		select {
		case <-ctx.Done():
			log.Info("tss signature read result loop stopped")
		case result := <-out:
			info := fmt.Sprintf("tss signature sessionID=%v, groupID=%v, taskID=%v", result.SessionID, result.GroupID, result.TaskID)
			if result.Err != nil {
				log.Errorf("%s, result error:%v", info, result.Err)
			} else {
				t.handleSigFinish(ctx, result.Data)
			}
		}
	}
}

func (t *TSSService) taskLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info("TSS task loop stopped")
		case task := <-t.taskReceive:
			log.Info("TSS task receive", task)
		}
	}
}

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, state *state.State, layer2Listener *layer2.Layer2Listener) *TSSService {
	return &TSSService{
		isPrepared:     atomic.Bool{},
		privateKey:     config.AppConfig.L2PrivateKey,
		localAddress:   crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		p2p:            p,
		dbm:            dbm,
		state:          state,
		layer2Listener: layer2Listener,
		scheduler:      NewScheduler[int64](p, int64(config.AppConfig.TssThreshold)),
		cache:          cache.New(time.Minute*10, time.Minute),
		taskReceive:    make(chan any, 256),
	}
}

func (t *TSSService) Start(ctx context.Context) {
	t.eventLoop(ctx)
	t.scheduler.Start()

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	t.Stop()
}

func (t *TSSService) eventLoop(ctx context.Context) {
	t.p2p.Bind(p2p.MessageTypeTssMsg, state.EventTssMsg{})
	t.tssMsgCh = t.state.EventBus.Subscribe(state.EventTssMsg{})
	t.partyAddOrRmCh = t.state.EventBus.Subscribe(state.EventParticipantAddedOrRemoved{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-t.tssMsgCh: // from p2p network
				log.Debugf("Received t msg event")

				e := event.(p2p.Message[json.RawMessage])

				err := t.handleSessionMsg(convertMsgData(e).(SessionMessage[int64]))
				if err != nil {
					log.Warnf("handle session msg error, %v", err)
				}
			case <-t.partyAddOrRmCh: // from layer2 log scan
				log.Debugf("Received t add or remove participant event")

				if t.localAddress == t.proposer {
					// todo
					newThreshold := 0

					var newPartners []common.Address
					_ = t.scheduler.NewReShareGroupSession(
						t.localAddress,
						helper.SenateTaskID,
						helper.SenateSessionID.Big(),
						t.proposer,
						int(t.threshold.Load()),
						t.partners,
						newThreshold,
						newPartners,
					)
				}
			}
		}
	}()
}

func (t *TSSService) Stop() {}

func (t *TSSService) waitForThreshold(ctx context.Context) {
	count := t.p2p.OnlinePeerCount()
L:
	for {
		select {
		case <-ctx.Done():
			log.Info("waitForThreshold context done")
		default:
			if int64(count) >= t.threshold.Load() {
				break L
			}
			time.Sleep(time.Second)
		}
	}
}

func (t *TSSService) oldPartners() []common.Address {
	return t.partners
}

func (t *TSSService) newPartners() []common.Address {
	return nil // todo
}
