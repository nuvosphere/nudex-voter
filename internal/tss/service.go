package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/task"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	isPrepared   atomic.Bool
	privateKey   *ecdsa.PrivateKey // submit
	localAddress common.Address    // submit = partyID
	proposer     common.Address    // current submitter

	p2p                p2p.P2PService
	state              *state.State
	scheduler          *Scheduler[TaskId]
	cache              *cache.Cache
	currentDoingTaskID int64

	layer2Listener *layer2.Layer2Listener
	dbm            *db.DatabaseManager
	threshold      *atomic.Int64
	partners       []common.Address
	// eventbus channel
	tssMsgCh    <-chan any
	pendingTask <-chan any

	rw sync.RWMutex
}

func (t *Service) IsCompleted(taskID int64) bool {
	_, ok := t.cache.Get(fmt.Sprintf("%d", taskID))
	return ok
}

func (t *Service) AddCompletedTask(taskID int64) {
	t.cache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (t *Service) LocalSubmitter() common.Address {
	return t.localAddress
}

func (t *Service) IsPrepared() bool {
	return t.isPrepared.Load()
}

func (t *Service) sigEndLoop(ctx context.Context) {
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
				t.AddCompletedTask(result.TaskID)
				t.handleSigFinish(ctx, result.Data)
			}
		}
	}
}

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, state *state.State, layer2Listener *layer2.Layer2Listener) *Service {
	localAddress := crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey)
	proposer := common.Address{} // todo

	return &Service{
		isPrepared:     atomic.Bool{},
		privateKey:     config.AppConfig.L2PrivateKey,
		localAddress:   crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		p2p:            p,
		dbm:            dbm,
		state:          state,
		layer2Listener: layer2Listener,
		scheduler:      NewScheduler[int64](p, int64(config.AppConfig.TssThreshold), localAddress, proposer),
		cache:          cache.New(time.Minute*10, time.Minute),
	}
}

func (t *Service) Start(ctx context.Context) {
	t.eventLoop(ctx)
	t.scheduler.Start()

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	t.Stop()
}

func (t *Service) eventLoop(ctx context.Context) {
	t.p2p.Bind(p2p.MessageTypeTssMsg, state.EventTssMsg{})
	t.tssMsgCh = t.state.EventBus.Subscribe(state.EventTssMsg{})
	t.pendingTask = t.state.EventBus.Subscribe(state.EventTask{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-t.tssMsgCh: // from p2p network
				log.Debugf("Received t msg event")

				e := event.(p2p.Message[json.RawMessage])

				err := t.handleSessionMsg(convertMsgData(e).(SessionMessage[TaskId, Msg]))
				if err != nil {
					log.Warnf("handle session msg error, %v", err)
				}
			case data := <-t.pendingTask: // from layer2 log scan
				log.Info("task", data)
				// todo tss task

				switch v := data.(type) {
				case *db.Task:
					if t.isCanProposal() {
						// todo
						log.Info(v)

						err := t.proposalSignTaskSession(*v)
						if err != nil {
							log.Warnf("handle session msg error, %v", err)
						}
					}
				case *task.SubmitterChosenPair:
					// todo
				case *task.ParticipantPair:
					if t.isCanProposal() {
						// todo
						newThreshold := 0

						var newPartners []common.Address
						_ = t.scheduler.NewReShareGroupSession(
							t.localAddress,
							t.proposer,
							helper.SenateTaskID,
							helper.SenateSessionID.Big(),
							int(t.threshold.Load()),
							t.partners,
							newThreshold,
							newPartners,
						)
					}
				}
			}
		}
	}()
}

func (t *Service) isCanProposal() bool {
	return t.localAddress == t.proposer
}

func (t *Service) Stop() {}

func (t *Service) oldPartners() []common.Address {
	return t.partners
}
