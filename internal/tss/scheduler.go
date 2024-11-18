package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Scheduler struct {
	p2p                      p2p.P2PService
	bus                      eventbus.Bus
	ctx                      context.Context
	cancel                   context.CancelFunc
	grw                      sync.RWMutex
	groups                   map[helper.GroupID]*helper.Group
	srw                      sync.RWMutex
	sessions                 map[helper.SessionID]Session[TaskId]
	sessionTasks             map[TaskId]Session[TaskId]
	sigInToOut               chan *SessionResult[TaskId, *tsscommon.SignatureData]
	senateInToOut            chan *SessionResult[TaskId, *keygen.LocalPartySaveData]
	masterLocalPartySaveData keygen.LocalPartySaveData
	localSubmitter           common.Address // submit = partyID
	proposer                 *atomic.Value  // current submitter
	partners                 *atomic.Value  // types.Participants
	discussedTaskCache       *cache.Cache
	voterContract            layer2.VoterContract
	stateDB                  *gorm.DB
	submitterChosen          *db.SubmitterChosen
	pendingProposal          chan any
	notify                   chan struct{}
	tssMsgCh                 <-chan any   // eventbus channel
	pendingTask              <-chan any   // eventbus channel
	newGroup                 atomic.Value // todo NewGroup
}

func NewScheduler(p p2p.P2PService, bus eventbus.Bus, stateDB *gorm.DB, voterContract layer2.VoterContract) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	localSubmitter := crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey)
	proposer, err := voterContract.Proposer()
	utils.Assert(err)

	pp := atomic.Value{}
	pp.Store(proposer)

	partners, err := voterContract.Participants()
	utils.Assert(err)

	ps := atomic.Value{}
	ps.Store(partners)

	return &Scheduler{
		p2p:                p,
		bus:                bus,
		srw:                sync.RWMutex{},
		grw:                sync.RWMutex{},
		groups:             make(map[helper.GroupID]*helper.Group),
		sessions:           make(map[helper.SessionID]Session[TaskId]),
		sessionTasks:       make(map[TaskId]Session[TaskId]),
		sigInToOut:         make(chan *SessionResult[TaskId, *tsscommon.SignatureData], 1024),
		senateInToOut:      make(chan *SessionResult[TaskId, *keygen.LocalPartySaveData]),
		ctx:                ctx,
		cancel:             cancel,
		localSubmitter:     localSubmitter,
		proposer:           &pp,
		partners:           &ps,
		discussedTaskCache: cache.New(time.Minute*10, time.Minute),
		pendingProposal:    make(chan any, 1024),
		notify:             make(chan struct{}, 1024),
		stateDB:            stateDB,
		voterContract:      voterContract,
	}
}

func (m *Scheduler) Start() {
	m.eventLoop(m.ctx)
	m.BlockDetectionThreshold()

	is := m.IsGenesis()
	if is {
		m.Genesis() // build senate session

		sessionResult := <-m.senateInToOut
		m.SaveSenateSessionResult(sessionResult)
	}

	m.LoopReShareGroupResult()
	// loop approveProposal
	m.loopApproveProposal()
}

func (m *Scheduler) SaveSenateSessionResult(sessionResult *SessionResult[TaskId, *keygen.LocalPartySaveData]) {
	if sessionResult.Err != nil {
		panic(sessionResult.Err)
	}

	err := saveTSSData(sessionResult.Data)
	utils.Assert(err)

	m.masterLocalPartySaveData = *sessionResult.Data
}

func (m *Scheduler) Stop() {
	m.cancel()
}

func (m *Scheduler) Genesis() {
	_ = m.NewGenerateKeySession(
		helper.SenateTaskID,
		helper.SenateTaskMsg,
	)

	log.Info("TSS keygen process started")
}

func (m *Scheduler) IsGenesis() bool {
	if m.masterLocalPartySaveData.ECDSAPub != nil {
		return false
	}

	localPartySaveData, err := LoadTSSData()
	if err != nil {
		log.Errorf("Failed to load TSS data: %v", err)
	}

	if localPartySaveData == nil {
		return true
	}

	m.masterLocalPartySaveData = *localPartySaveData

	return false
}

func (m *Scheduler) BlockDetectionThreshold() {
L:
	for {
		select {
		case <-m.ctx.Done():
			log.Info("DetectionThreshold context done")
		default:
			count := m.p2p.OnlinePeerCount()
			if count >= m.Threshold() {
				break L
			}
			time.Sleep(time.Second)
		}
	}
}

func (m *Scheduler) Threshold() int {
	return m.Participants().Threshold()
}

func (m *Scheduler) AddGroup(group *helper.Group) {
	m.grw.Lock()
	defer m.grw.Unlock()
	m.groups[group.GroupID] = group
}

func (m *Scheduler) AddSession(session Session[TaskId]) bool {
	m.grw.Lock()
	_, ok := m.groups[session.GroupID()] // todo
	m.grw.Unlock()

	if ok {
		m.srw.Lock()
		m.sessions[session.SessionID()] = session
		m.sessionTasks[session.TaskID()] = session
		m.srw.Unlock()
	}

	return ok
}

func (m *Scheduler) GetGroup(groupID helper.GroupID) *helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Scheduler) GetSession(sessionID helper.SessionID) Session[TaskId] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Scheduler) IsMeeting(taskID TaskId) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.sessionTasks[taskID]

	return ok
}

func (m *Scheduler) GetGroups() []*helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ helper.GroupID, group *helper.Group) *helper.Group { return group })
}

func (m *Scheduler) GetSessions() []Session[TaskId] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ helper.SessionID, session Session[TaskId]) Session[TaskId] { return session })
}

func (m *Scheduler) ReleaseGroup(groupID helper.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Scheduler) SessionRelease(session helper.SessionID) {
	m.srw.Lock()
	defer m.srw.Unlock()

	s, ok := m.sessions[session]
	if ok {
		delete(m.sessions, session)
		delete(m.sessionTasks, s.TaskID())
		s.Release()

		if len(m.sessions) == 0 {
			m.notify <- struct{}{}
		}
	}
}

func (m *Scheduler) Release() {
	m.grw.Lock()
	m.groups = make(map[helper.GroupID]*helper.Group)
	m.grw.Unlock()
	m.srw.Lock()
	for _, s := range m.sessions {
		s.Release()
	}

	m.sessions = make(map[helper.SessionID]Session[TaskId])
	m.sessionTasks = make(map[TaskId]Session[TaskId])
	m.srw.Unlock()
	close(m.sigInToOut)
}

func (m *Scheduler) SubmitProposal(proposal any) {
	m.pendingProposal <- proposal
}

func (m *Scheduler) loopApproveProposal() {
	go func() {
		select {
		case <-m.ctx.Done():
			log.Info("approve proposal done")
		case <-m.notify:
			select {
			case <-m.ctx.Done():
				log.Info("approve proposal done")
			case proposal := <-m.pendingProposal:
				// todo
				// signature and re-share
				log.Info("doing approve proposal", proposal)
				m.BlockDetectionThreshold()
			}
		}
	}()
}

func (m *Scheduler) MasterPublicKey() ecdsa.PublicKey {
	return *m.masterLocalPartySaveData.ECDSAPub.ToECDSAPubKey()
}

func (m *Scheduler) LoopReShareGroupResult() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("loopReShareGroupResult done")
			case sessionResult := <-m.senateInToOut:
				m.SaveSenateSessionResult(sessionResult)
				newGroup := m.newGroup.Load().(*NewGroup)
				m.partners.Store(newGroup.NewParts)
				newGroup = nil // todo
				m.newGroup.Store(newGroup)
			}
		}
	}()
}

func (m *Scheduler) IsDiscussed(taskID int64) bool {
	_, ok := m.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	if !ok {
		ok, _ = m.voterContract.IsTaskCompleted(big.NewInt(taskID))
	}

	return ok
}

func (m *Scheduler) AddDiscussedTask(taskID int64) {
	m.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (m *Scheduler) LocalSubmitter() common.Address {
	return m.localSubmitter
}

func (m *Scheduler) Proposer() common.Address {
	proposer, err := m.voterContract.Proposer()
	utils.Assert(err)

	return proposer
}

func (m *Scheduler) IsProposer() bool {
	return m.Proposer() == m.LocalSubmitter()
}

func (m *Scheduler) eventLoop(ctx context.Context) {
	m.p2p.Bind(p2p.MessageTypeTssMsg, eventbus.EventTssMsg{})
	m.tssMsgCh = m.bus.Subscribe(eventbus.EventTssMsg{})
	m.pendingTask = m.bus.Subscribe(eventbus.EventTask{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-m.tssMsgCh: // from p2p network
				log.Debugf("Received m msg event")

				e := event.(p2p.Message[json.RawMessage])

				err := m.handleSessionMsg(convertMsgData(e).(SessionMessage[TaskId, Msg]))
				if err != nil {
					log.Warnf("handle session msg error, %v", err)
				}
			case data := <-m.pendingTask: // from layer2 log scan
				log.Info("received task: ", data)

				switch v := data.(type) {
				case db.ITask:
					if m.isCanProposal() {
						log.Info("proposal task", v)

						err := m.proposalTaskSession(v)
						if err != nil {
							log.Warnf("handle session msg error, %v", err)
						}
					}
				case *db.ParticipantEvent:
					party := common.HexToAddress(v.Address)

					var newParts types.Participants

					switch v.EventName {
					case layer2.ParticipantAdded:
						if !slices.Contains(m.Participants(), party) {
							newParts = append(m.Participants(), party)
						}
					case layer2.ParticipantRemoved:
						newParts = lo.Filter(m.Participants(), func(item common.Address, index int) bool { return item != party })
					}

					if len(newParts) > 0 && len(newParts) != len(m.Participants()) {
						m.newGroup.Store(&NewGroup{
							Event:    v,
							NewParts: m.Participants(),
							OldParts: newParts,
						})

						if m.isCanProposal() {
							_ = m.NewReShareGroupSession(
								helper.SenateTaskID,
								helper.SenateTaskMsg,
								m.Participants(),
								newParts,
							)
						}
					}

				case *db.SubmitterChosen:
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))

				case *db.TaskCompletedEvent: // todo
					log.Infof("taskID: %d completed on blockchain", v.TaskId)
				}
			}
		}
	}()
}

func (m *Scheduler) isCanProposal() bool {
	m.BlockDetectionThreshold()
	return m.LocalSubmitter() == m.Proposer()
}

func (m *Scheduler) Participants() types.Participants {
	if m.IsProposer() {
		return m.partners.Load().(types.Participants)
	}

	partners, err := m.voterContract.Participants()
	utils.Assert(err)
	m.partners.Store(partners)

	return partners
}

type NewGroup struct {
	Event    *db.ParticipantEvent
	NewParts types.Participants
	OldParts types.Participants
}

type Draft struct {
	Type    int
	DraftID int32
	Extra   any
}
