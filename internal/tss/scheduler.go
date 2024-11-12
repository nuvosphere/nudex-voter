package tss

import (
	"context"
	"crypto/ecdsa"
	"sync"
	"sync/atomic"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Scheduler[T comparable] struct {
	p2p                      p2p.P2PService
	grw                      sync.RWMutex
	groups                   map[helper.GroupID]*helper.Group
	srw                      sync.RWMutex
	sessions                 map[helper.SessionID]Session[T]
	sessionTasks             map[T]Session[T]
	sigInToOut               chan *SessionResult[T, *tsscommon.SignatureData]
	senateInToOut            chan *SessionResult[T, *keygen.LocalPartySaveData]
	ctx                      context.Context
	cancel                   context.CancelFunc
	masterLocalPartySaveData *keygen.LocalPartySaveData
	masterPubKey             *ecdsa.PublicKey
	threshold                *atomic.Int64
	localSubmitter           common.Address
	proposer                 common.Address // current proposer(submitter)
	pendingProposal          chan any
	notify                   chan struct{}
}

func NewScheduler[T comparable](p p2p.P2PService, threshold int64, localSubmitter, proposer common.Address) *Scheduler[T] {
	t := &atomic.Int64{}
	t.Store(threshold)

	ctx, cancel := context.WithCancel(context.Background())

	return &Scheduler[T]{
		p2p:             p,
		srw:             sync.RWMutex{},
		grw:             sync.RWMutex{},
		groups:          make(map[helper.GroupID]*helper.Group),
		sessions:        make(map[helper.SessionID]Session[T]),
		sessionTasks:    make(map[T]Session[T]),
		sigInToOut:      make(chan *SessionResult[T, *tsscommon.SignatureData], 1024),
		senateInToOut:   make(chan *SessionResult[T, *keygen.LocalPartySaveData]),
		threshold:       t,
		ctx:             ctx,
		cancel:          cancel,
		localSubmitter:  localSubmitter,
		proposer:        proposer,
		pendingProposal: make(chan any, 1024),
		notify:          make(chan struct{}, 1024),
	}
}

func (m *Scheduler[T]) Start() {
	m.DetectionThreshold()

	is := m.IsGenesis()
	if is {
		m.Genesis() // build senate session

		sessionResult := <-m.senateInToOut
		if sessionResult.Err != nil {
			panic(sessionResult.Err)
		}

		err := saveTSSData(sessionResult.Data)
		utils.Assert(err)

		m.masterLocalPartySaveData = sessionResult.Data
	}

	m.masterPubKey = m.masterLocalPartySaveData.ECDSAPub.ToECDSAPubKey()

	// loop approveProposal
	m.loopApproveProposal()
}

func (m *Scheduler[T]) Stop() {
	m.cancel()
}

func (m *Scheduler[T]) Genesis() {
	// todo
	log.Info("TSS keygen process started")
}

func (m *Scheduler[T]) IsGenesis() bool {
	if m.masterLocalPartySaveData != nil && m.masterLocalPartySaveData.ECDSAPub != nil {
		return false
	}

	localPartySaveData, err := LoadTSSData()
	if err != nil {
		log.Errorf("Failed to load TSS data: %v", err)
	}

	if localPartySaveData == nil {
		return true
	}

	m.masterLocalPartySaveData = localPartySaveData

	return false
}

func (m *Scheduler[T]) DetectionThreshold() {
L:
	for {
		select {
		case <-m.ctx.Done():
			log.Info("DetectionThreshold context done")
		default:
			count := m.p2p.OnlinePeerCount()
			if int64(count) >= m.threshold.Load() {
				break L
			}
			time.Sleep(time.Second)
		}
	}
}

func (m *Scheduler[T]) Threshold() int64 {
	return m.threshold.Load()
}

func (m *Scheduler[T]) SetThreshold(threshold int64) {
	m.threshold.Store(threshold)
}

func (m *Scheduler[T]) AddGroup(group *helper.Group) {
	m.grw.Lock()
	defer m.grw.Unlock()
	m.groups[group.GroupID] = group
}

func (m *Scheduler[T]) AddSession(session Session[T]) bool {
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

func (m *Scheduler[T]) GetGroup(groupID helper.GroupID) *helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Scheduler[T]) GetSession(sessionID helper.SessionID) Session[T] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Scheduler[T]) IsMeeting(taskID T) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.sessionTasks[taskID]

	return ok
}

func (m *Scheduler[T]) GetGroups() []*helper.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ helper.GroupID, group *helper.Group) *helper.Group { return group })
}

func (m *Scheduler[T]) GetSessions() []Session[T] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ helper.SessionID, session Session[T]) Session[T] { return session })
}

func (m *Scheduler[T]) ReleaseGroup(groupID helper.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Scheduler[T]) SessionRelease(session helper.SessionID) {
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

func (m *Scheduler[T]) Release() {
	m.grw.Lock()
	m.groups = make(map[helper.GroupID]*helper.Group)
	m.grw.Unlock()
	m.srw.Lock()
	for _, s := range m.sessions {
		s.Release()
	}

	m.sessions = make(map[helper.SessionID]Session[T])
	m.sessionTasks = make(map[T]Session[T])
	m.srw.Unlock()
	close(m.sigInToOut)
}

func (m *Scheduler[T]) SubmitProposal(proposal any) {
	m.pendingProposal <- proposal
}

func (m *Scheduler[T]) loopApproveProposal() {
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
				m.DetectionThreshold()
			}
		}
	}()
}

func (m *Scheduler[T]) MasterPublicKey() *ecdsa.PublicKey {
	return m.masterPubKey
}
