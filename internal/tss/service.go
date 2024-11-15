package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"slices"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	isPrepared      atomic.Bool
	privateKey      *ecdsa.PrivateKey // submit
	localAddress    common.Address    // submit = partyID
	proposer        common.Address    // current submitter
	p2p             p2p.P2PService
	state           *state.State
	scheduler       *Scheduler[TaskId]
	cache           *cache.Cache
	layer2Listener  *layer2.Layer2Listener
	dbm             *db.DatabaseManager
	partners        types.Participants
	submitterChosen *db.SubmitterChosen
	// eventbus channel
	tssMsgCh    <-chan any
	pendingTask <-chan any
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
	// proposer := layer2Listener.Proposer()
	threshold := config.AppConfig.TssThreshold

	var parts types.Participants = layer2Listener.Participants()

	if parts.Threshold() > 0 {
		threshold = parts.Threshold()
	}

	return &Service{
		isPrepared:     atomic.Bool{},
		privateKey:     config.AppConfig.L2PrivateKey,
		localAddress:   crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		p2p:            p,
		dbm:            dbm,
		state:          state,
		layer2Listener: layer2Listener,
		scheduler:      NewScheduler[int64](p, int64(threshold), localAddress),
		cache:          cache.New(time.Minute*10, time.Minute),
		partners:       parts,
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
	t.p2p.Bind(p2p.MessageTypeTssMsg, eventbus.EventTssMsg{})
	t.tssMsgCh = t.state.Bus().Subscribe(eventbus.EventTssMsg{})
	t.pendingTask = t.state.Bus().Subscribe(eventbus.EventTask{})

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

				switch v := data.(type) {
				case db.ITask:
					if t.isCanProposal() {
						log.Info("proposal task", v)

						err := t.proposalTaskSession(v)
						if err != nil {
							log.Warnf("handle session msg error, %v", err)
						}
					}
				case *db.ParticipantEvent:
					party := common.HexToAddress(v.Address)

					var newNewParts types.Participants

					switch v.EventName {
					case layer2.ParticipantAdded:
						if !slices.Contains(t.partners, party) {
							newNewParts = append(t.partners, party)
						}
					case layer2.ParticipantRemoved:
						newNewParts = lo.Filter(t.partners, func(item common.Address, index int) bool { return item != party })
					}

					if t.isCanProposal() && len(newNewParts) > 0 && len(newNewParts) != len(t.partners) {
						_ = t.scheduler.NewReShareGroupSession(
							t.localAddress,
							t.proposer,
							helper.SenateTaskID,
							helper.SenateTaskMsg,
							t.partners,
							newNewParts,
						)
					}

				case *db.SubmitterChosen:
					t.submitterChosen = v
					t.proposer = common.HexToAddress(v.Submitter)

				case *db.TaskCompletedEvent:
					log.Infof("taskID: %d completed on blockchain", v.TaskId)
				}
			}
		}
	}()
}

func (t *Service) isCanProposal() bool {
	t.scheduler.BlockDetectionThreshold()
	return t.localAddress == t.proposer
}

func (t *Service) Stop() {}

func (t *Service) Threshold() int {
	return t.partners.Threshold()
}
