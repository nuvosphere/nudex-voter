package tss

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/ethereum/go-ethereum/common"
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
	isProd             bool
	p2p                p2p.P2PService
	bus                eventbus.Bus
	ctx                context.Context
	cancel             context.CancelFunc
	grw                sync.RWMutex
	groups             map[helper.GroupID]*helper.Group
	srw                sync.RWMutex
	sessions           map[helper.SessionID]Session[ProposalID]
	proposalSession    map[ProposalID]Session[ProposalID]
	sigInToOut         chan *SessionResult[ProposalID, *tsscommon.SignatureData]
	senateInToOut      chan *SessionResult[ProposalID, *helper.LocalPartySaveData]
	partyData          *PartyData
	localSubmitter     common.Address // submit = partyID
	proposer           *atomic.Value  // current submitter
	partners           *atomic.Value  // types.Participants
	discussedTaskCache *cache.Cache
	voterContract      layer2.VoterContract
	stateDB            *gorm.DB
	submitterChosen    *db.SubmitterChosen
	pendingProposal    chan any
	notify             chan struct{}
	tssMsgCh           <-chan any   // eventbus channel
	pendingTask        <-chan any   // eventbus channel
	newGroup           atomic.Value // todo NewGroup
}

func NewScheduler(isProd bool, p p2p.P2PService, bus eventbus.Bus, stateDB *gorm.DB, voterContract layer2.VoterContract, localSubmitter common.Address) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	proposer, err := voterContract.Proposer()
	utils.Assert(err)

	pp := atomic.Value{}
	pp.Store(proposer)

	partners, err := voterContract.Participants()
	utils.Assert(err)

	ps := atomic.Value{}
	ps.Store(partners)

	return &Scheduler{
		isProd:             isProd,
		p2p:                p,
		bus:                bus,
		srw:                sync.RWMutex{},
		grw:                sync.RWMutex{},
		groups:             make(map[helper.GroupID]*helper.Group),
		sessions:           make(map[helper.SessionID]Session[ProposalID]),
		proposalSession:    make(map[ProposalID]Session[ProposalID]),
		sigInToOut:         make(chan *SessionResult[ProposalID, *tsscommon.SignatureData], 1024),
		senateInToOut:      make(chan *SessionResult[ProposalID, *helper.LocalPartySaveData]),
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
		partyData:          NewPartyData(config.AppConfig.DbDir),
	}
}

func (m *Scheduler) Start() {
	m.p2pLoop()
	m.proposalLoop()
	m.BlockDetectionThreshold()

	is := m.IsGenesis()
	if is {
		if m.isCanProposal() {
			log.Info("TSS keygen process started ", "leader:", m.LocalSubmitter(), "proposer: ", m.Proposer())
			// leader
			m.Genesis() // build senate session
		} else {
			log.Info("TSS keygen process started ", "Candidate:", m.LocalSubmitter(), "proposer: ", m.Proposer())
		}

		m.saveSenateData()
		log.Info("TSS keygen success!", "localSubmitter:", m.LocalSubmitter(), "proposer: ", m.Proposer())
	}

	// loop approveProposal
	m.loopApproveProposal()
	m.reGroupResultLoop()
}

func (m *Scheduler) SaveSenateSessionResult(sessionResult *SessionResult[ProposalID, *helper.LocalPartySaveData]) {
	if sessionResult.Err != nil {
		panic(sessionResult.Err)
	}

	err := m.partyData.SaveLocalData(sessionResult.Data)
	utils.Assert(err)
}

func (m *Scheduler) Stop() {
	m.cancel()
}

func (m *Scheduler) Genesis() {
	_ = m.NewGenerateKeySession(
		helper.ECDSA,
		helper.SenateProposalIDOfECDSA,
		helper.SenateSessionIDOfECDSA,
		common.Address{},
		helper.SenateProposal,
	)
	_ = m.NewGenerateKeySession(
		helper.EDDSA,
		helper.SenateProposalIDOfEDDSA,
		helper.SenateSessionIDOfEDDSA,
		common.Address{},
		helper.SenateProposal,
	)
}

func (m *Scheduler) saveSenateData() {
	sessionResult := <-m.senateInToOut
	m.SaveSenateSessionResult(sessionResult)
	sessionResult = <-m.senateInToOut
	m.SaveSenateSessionResult(sessionResult)
}

func (m *Scheduler) IsGenesis() bool {
	return !m.partyData.LoadData()
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

func (m *Scheduler) AddSession(session Session[ProposalID]) bool {
	//m.grw.Lock()
	//_, ok := m.groups[session.GroupID()] // todo
	//m.grw.Unlock()
	//
	//if ok {
	//	m.srw.Lock()
	//	m.sessions[session.SessionID()] = session
	//	m.proposalSession[session.ProposalID()] = session
	//	m.srw.Unlock()
	//}
	m.srw.Lock()
	m.sessions[session.SessionID()] = session
	m.proposalSession[session.ProposalID()] = session
	m.srw.Unlock()

	return true
}

func (m *Scheduler) GetGroup(groupID helper.GroupID) *helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Scheduler) GetSession(sessionID helper.SessionID) Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Scheduler) IsMeeting(taskID ProposalID) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.proposalSession[taskID]

	return ok
}

func (m *Scheduler) GetGroups() []*helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ helper.GroupID, group *helper.Group) *helper.Group { return group })
}

func (m *Scheduler) GetSessions() []Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ helper.SessionID, session Session[ProposalID]) Session[ProposalID] { return session })
}

func (m *Scheduler) ReleaseGroup(groupID helper.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Scheduler) SessionRelease(sessionID helper.SessionID) {
	m.srw.Lock()
	defer m.srw.Unlock()

	s, ok := m.sessions[sessionID]
	if ok {
		delete(m.sessions, sessionID)
		delete(m.proposalSession, s.ProposalID())
		s.Release()
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

	m.sessions = make(map[helper.SessionID]Session[ProposalID])
	m.proposalSession = make(map[ProposalID]Session[ProposalID])
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
		case proposal := <-m.pendingProposal:
			// todo
			// signature proposal
			log.Info("doing approve proposal", proposal)
			m.BlockDetectionThreshold()
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

func (m *Scheduler) p2pLoop() {
	m.p2p.Bind(p2p.MessageTypeTssMsg, eventbus.EventTssMsg{})
	m.tssMsgCh = m.bus.Subscribe(eventbus.EventTssMsg{})

	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-m.tssMsgCh: // from p2p network
				log.Debugf("Received m msg event")

				e := event.(p2p.Message[json.RawMessage])
				proposal := convertMsgData(e).(SessionMessage[ProposalID, Proposal])

				err := m.processReceivedProposal(proposal)
				if err != nil {
					log.Warnf("handle session msg error, %v", err)
				}
			}
		}
	}()
	log.Info("p2p loop started")
}

func (m *Scheduler) proposalLoop() {
	m.pendingTask = m.bus.Subscribe(eventbus.EventTask{})

	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("proposal loop stopping...")
				return
			case data := <-m.pendingTask: // from layer2 log scan
				log.Info("received task from layer2 log scan: ", data)

				switch v := data.(type) {
				case db.ITask:
					if m.isCanProposal() {
						log.Info("proposal task", v)
						m.processTaskProposal(v)
					}
				case *db.ParticipantEvent:
					m.processReGroupProposal(v)

				case *db.SubmitterChosen:
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))

				case *db.TaskCompletedEvent: // todo
					log.Infof("taskID: %d completed on blockchain", v.TaskId)
				}
			}
		}
	}()
	log.Info("proposal loop started")
}

func (m *Scheduler) processReGroupProposal(v *db.ParticipantEvent) {
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
				helper.ECDSA,
				helper.SenateProposalIDOfECDSA,
				helper.SenateProposal,
				m.Participants(),
				newParts,
			)
			_ = m.NewReShareGroupSession(
				helper.EDDSA,
				helper.SenateProposalIDOfEDDSA,
				helper.SenateProposal,
				m.Participants(),
				newParts,
			)

			newGroup := m.newGroup.Load().(*NewGroup)
			m.partners.Store(newGroup.NewParts)
			newGroup = nil // todo
			m.newGroup.Store(newGroup)
		}
	}
}

func (m *Scheduler) reGroupResultLoop() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("reGroup result loop stopping...")
			case sessionResult := <-m.senateInToOut:
				m.SaveSenateSessionResult(sessionResult)
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
