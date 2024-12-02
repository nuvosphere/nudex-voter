package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Scheduler struct {
	isProd             bool
	p2p                p2p.P2PService
	bus                eventbus.Bus
	ctx                context.Context
	cancel             context.CancelFunc
	grw                sync.RWMutex
	groups             map[types.GroupID]*types.Group
	srw                sync.RWMutex
	sessions           map[types.SessionID]Session[ProposalID]
	proposalSession    map[ProposalID]Session[ProposalID]
	sigInToOut         chan *SessionResult[ProposalID, *tsscommon.SignatureData]
	senateInToOut      chan *SessionResult[ProposalID, *types.LocalPartySaveData]
	partyData          *PartyData
	localSubmitter     common.Address
	proposer           *atomic.Value // current submitter
	partners           *atomic.Value // types.Participants
	ecCount            *atomic.Int64
	newGroup           *atomic.Value          // *NewGroup
	taskQueue          *pool.Pool[uint64]     // created state task
	pendingStateTasks  *pool.Pool[uint64]     // pending state task
	operationsQueue    *pool.Pool[ProposalID] // pending batch task
	discussedTaskCache *cache.Cache
	voterContract      layer2.VoterContract
	stateDB            *gorm.DB
	submitterChosen    *db.SubmitterChosen
	notify             chan struct{}
	tssMsgCh           <-chan any // eventbus channel
	pendingTask        <-chan any // eventbus channel
	handleSigFinish    func(*Operations)
	currentNonce       *atomic.Uint64
}

func NewScheduler(isProd bool, p p2p.P2PService, bus eventbus.Bus, stateDB *gorm.DB, voterContract layer2.VoterContract, localSubmitter common.Address) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	pp := atomic.Value{}

	proposer, err := voterContract.Proposer()
	if err != nil {
		log.Warnf("get proposer error, %s", err.Error())
		log.Infof("TssPublicKeys: %v", len(config.TssPublicKeys))
		proposer = crypto.PubkeyToAddress(*config.TssPublicKeys[0]) // genesis
		pp.Store(proposer)
	} else {
		pp.Store(proposer)
	}

	ps := atomic.Value{}

	partners, err := voterContract.Participants()
	if err != nil {
		log.Warnf("get partners error, %s", err.Error())
		partners = lo.Map(config.TssPublicKeys, func(item *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*item) })
		ps.Store(partners)
	} else {
		ps.Store(partners)
	}
	log.Infof("partners: %v", partners)
	p.UpdateParticipants(partners)
	currentNonce := &atomic.Uint64{}
	nonce, _ := voterContract.TssNonce()
	if nonce != nil {
		currentNonce.Store(nonce.Uint64())
	}

	newGroup := &atomic.Value{}
	newGroup.Store(nullNewGroup)

	return &Scheduler{
		isProd:             isProd,
		p2p:                p,
		bus:                bus,
		srw:                sync.RWMutex{},
		grw:                sync.RWMutex{},
		groups:             make(map[types.GroupID]*types.Group),
		sessions:           make(map[types.SessionID]Session[ProposalID]),
		proposalSession:    make(map[ProposalID]Session[ProposalID]),
		sigInToOut:         make(chan *SessionResult[ProposalID, *tsscommon.SignatureData], 1024),
		senateInToOut:      make(chan *SessionResult[ProposalID, *types.LocalPartySaveData], 1024),
		ctx:                ctx,
		cancel:             cancel,
		localSubmitter:     localSubmitter,
		proposer:           &pp,
		partners:           &ps,
		newGroup:           newGroup,
		taskQueue:          pool.NewTxPool[uint64](),
		pendingStateTasks:  pool.NewTxPool[uint64](),
		operationsQueue:    pool.NewTxPool[ProposalID](),
		discussedTaskCache: cache.New(time.Minute*10, time.Minute),
		notify:             make(chan struct{}, 1024),
		stateDB:            stateDB,
		voterContract:      voterContract,
		partyData:          NewPartyData(config.AppConfig.DbDir),
		currentNonce:       currentNonce,
	}
}

func (m *Scheduler) Start() {
	m.p2pLoop()
	m.proposalLoop()
	m.BlockDetectionThreshold()

	if m.IsGenesis() {
		if m.isCanProposal() {
			log.Info("TSS keygen process started ", "leader:", m.LocalSubmitter(), " proposer: ", m.Proposer())
			// leader
			m.Genesis() // build senate session
		} else {
			log.Info("TSS keygen process started ", "Candidate:", m.LocalSubmitter(), " proposer: ", m.Proposer())
		}

		m.saveSenateData()
		log.Info("TSS keygen success!", "localSubmitter:", m.LocalSubmitter(), " proposer: ", m.Proposer(), " ECDSA PublicKey: ", m.partyData.ECDSALocalData().PublicKeyBase58(), " EDDSA PublicKey: ", m.partyData.EDDSALocalData().PublicKeyBase58())
	} else {
		log.Info("local data already exists: scheduler begin running")
	}

	log.Infof("********Scheduler master tss ecdsa address********: %v", m.partyData.GetData(types.ECDSA).Address())
	log.Infof("localSubmitter: %v, proposer: %v", m.LocalSubmitter(), m.Proposer())
	// loop approveProposal
	m.loopApproveProposal()
	m.reGroupResultLoop()
	m.loopSigInToOut()
	m.loopDetectionCondition()
	log.Info("Scheduler stared success!")
}

func (m *Scheduler) SaveSenateSessionResult(sessionResult *SessionResult[ProposalID, *types.LocalPartySaveData]) {
	if sessionResult.Err != nil {
		panic(sessionResult.Err)
	}

	err := m.partyData.SaveLocalData(sessionResult.Data)
	utils.Assert(err)
	log.Info("TSS keygen success! SaveSenateSessionResult: ", "localSubmitter:", m.LocalSubmitter())
}

func (m *Scheduler) Stop() {
	m.cancel()
}

func (m *Scheduler) Genesis() {
	_ = m.NewGenerateKeySession(
		types.ECDSA,
		types.SenateProposalIDOfECDSA,
		types.SenateSessionIDOfECDSA,
		common.Address{},
		types.SenateProposal,
	)
	_ = m.NewGenerateKeySession(
		types.EDDSA,
		types.SenateProposalIDOfEDDSA,
		types.SenateSessionIDOfEDDSA,
		common.Address{},
		types.SenateProposal,
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
			threshold := m.Threshold()
			if count > 0 && threshold > 0 && count > threshold {
				break L
			}
			log.Infof("detection online peer count:%d, threshold:%d", count, threshold)
			time.Sleep(time.Second)
		}
	}
}

func (m *Scheduler) Threshold() int {
	return m.Participants().Threshold()
}

func (m *Scheduler) AddGroup(group *types.Group) {
	m.grw.Lock()
	defer m.grw.Unlock()
	m.groups[group.GroupID()] = group
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

func (m *Scheduler) GetGroup(groupID types.GroupID) *types.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Scheduler) GetSession(sessionID types.SessionID) Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Scheduler) IsMeeting(proposalID ProposalID) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.proposalSession[proposalID]

	return ok
}

func (m *Scheduler) GetGroups() []*types.Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ types.GroupID, group *types.Group) *types.Group { return group })
}

func (m *Scheduler) GetSessions() []Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ types.SessionID, session Session[ProposalID]) Session[ProposalID] { return session })
}

func (m *Scheduler) ReleaseGroup(groupID types.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Scheduler) SessionRelease(sessionID types.SessionID) {
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
	m.groups = make(map[types.GroupID]*types.Group)
	m.grw.Unlock()
	m.srw.Lock()
	for _, s := range m.sessions {
		s.Release()
	}

	m.sessions = make(map[types.SessionID]Session[ProposalID])
	m.proposalSession = make(map[ProposalID]Session[ProposalID])
	m.srw.Unlock()
	close(m.sigInToOut)
}

func (m *Scheduler) loopApproveProposal() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		select {
		case <-m.ctx.Done():
			log.Info("approve proposal done")

		case <-ticker.C:
			m.BatchTask()

		case <-m.notify:
			m.BatchTask()
		}
	}()
}

func (m *Scheduler) BatchTask() {
	if m.isCanProposal() && m.isCanNext() {
		log.Info("batch proposal")
		tasks := m.taskQueue.GetTopN(TopN)
		operations := lo.Map(tasks, func(item pool.Task[uint64], index int) contracts.Operation { return *m.Operation(item) })
		if len(operations) == 0 {
			log.Warnf("operationsQueue is empty")
			return
		}
		nonce, dataHash, msg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(operations)
		if err != nil {
			log.Errorf("batch task generate verify task unsign msg err:%v", err)
			return
		}
		log.Infof("nonce: %v, dataHash: %v, msg: %v", nonce, dataHash, msg)

		// only ecdsa batch
		m.NewMasterSignBatchSession(
			types.ZeroSessionID,
			nonce.Uint64(), // ProposalID
			msg.Big(),
			lo.Map(tasks, func(item pool.Task[uint64], index int) ProposalID { return item.TaskID() }),
		)
		m.saveOperations(nonce, operations, dataHash, msg)
	}
}

func (m *Scheduler) isCanNext() bool {
	op := m.operationsQueue.Last()
	if op == nil {
		return true
	}
	if op.(*Operations).Signature != nil {
		return true
	}
	return false
}

func (m *Scheduler) IsDiscussed(taskID uint64) bool {
	_, ok := m.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	if !ok {
		ok, _ = m.voterContract.IsTaskCompleted(taskID)
	}

	return ok
}

func (m *Scheduler) AddDiscussedTask(taskID uint64) {
	m.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (m *Scheduler) LocalSubmitter() common.Address {
	return m.localSubmitter
}

var zeroAddress = common.Address{}

func (m *Scheduler) Proposer() common.Address {
	p := m.proposer.Load()
	if p != nil {
		return p.(common.Address)
	}
	proposer, err := m.voterContract.Proposer()
	if proposer != zeroAddress && err == nil {
		m.proposer.Store(proposer)
	}
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

const TopN = 20

// from layer2 log event
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
				case pool.Task[uint64]: // task create
					if m.IsDiscussed(v.TaskID()) {
						log.Errorf("received task from layer2 is discussed : %v", v.TaskID())
					} else {
						m.taskQueue.Add(v)

						if m.taskQueue.Len() >= TopN {
							m.notify <- struct{}{}
						}
					}
				case *db.ParticipantEvent: // regroup
					m.processReGroupProposal(v)

				case *db.SubmitterChosen: // charge proposer
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))

				case *db.TaskUpdatedEvent: // todo
					switch v.State {
					case db.Pending:
						// todo withdraw
						task := &db.Task{}
						err := m.stateDB.
							Model(task).
							Preload(clause.Associations).
							Where("task_id = ?", v.TaskId).
							First(task).
							Error
						if err != nil {
							log.Errorf("get task err:%v", err)
						} else {
							// todo
							m.pendingStateTasks.Add(task)
						}

					case db.Completed, db.Failed:
						m.taskQueue.Remove(v.TaskId)
						m.AddDiscussedTask(v.TaskId)
					default:
						log.Errorf("invalid task state : %v", v.State)
					}

					log.Infof("taskID: %d completed on blockchain", v.TaskId)
				}
			}
		}
	}()

	testPendingTask := m.bus.Subscribe(eventbus.EventTestTask{})

	// test branch
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("proposal loop stopping...")
				return
			case data := <-testPendingTask: // from test task
				log.Info("received task from layer2 log scan: ", data)

				switch v := data.(type) {
				case pool.Task[uint64]:
					m.taskQueue.Add(v)

					if m.isCanProposal() {
						log.Info("proposal task", v)
						m.processTaskProposal(v)
					}
				case *db.ParticipantEvent: // regroup
					m.processReGroupProposal(v)

				case *db.SubmitterChosen: // charge proposer
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))

				case *db.TaskUpdatedEvent: // todo
					log.Infof("taskID: %d completed on blockchain", v.TaskId)
					if v.State == db.Completed {
						m.taskQueue.Remove(v.TaskId)
						m.AddDiscussedTask(v.TaskId)
					}
				}
			}
		}
	}()
	log.Info("proposal loop started")
}

func (m *Scheduler) processReGroupProposal(v *db.ParticipantEvent) {
	joinAddress := common.HexToAddress(v.Address)
	newParts := types.Participants{}
	oldParts := m.Participants()

	log.Debugf("ParticipantEvent: %v, address: %v", v.EventName, v.Address)

	switch v.EventName {
	case layer2.ParticipantAdded:
		if !slices.Contains(oldParts, joinAddress) {
			newParts = append(oldParts, joinAddress)
		}
	case layer2.ParticipantRemoved:
		newParts = lo.Filter(oldParts, func(item common.Address, index int) bool { return item != joinAddress })
	}

	log.Debugf("newParts: %v, oldParts: %v, joinAddress:%v ", newParts, oldParts, joinAddress)

	if len(newParts) > 0 && newParts.GroupID() != oldParts.GroupID() {
		m.newGroup.Store(&NewGroup{
			Event:    v,
			NewParts: newParts,
			OldParts: oldParts,
		})

		if m.isCanProposal() {
			for {
				if m.p2p.IsOnline(strings.ToLower(joinAddress.String())) {
					break
				}
				time.Sleep(1 * time.Second)
			}

			_ = m.NewReShareGroupSession(
				types.ECDSA,
				types.SenateSessionIDOfECDSA,
				types.SenateProposalIDOfECDSA,
				types.SenateProposal,
				oldParts,
				newParts,
			)
			_ = m.NewReShareGroupSession(
				types.EDDSA,
				types.SenateSessionIDOfEDDSA,
				types.SenateProposalIDOfEDDSA,
				types.SenateProposal,
				oldParts,
				newParts,
			)

			log.Info("Leader NewReShareGroupSession stared")
		} else {
			log.Info("Candidate NewReShareGroupSession stared")
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
				m.ecCount.Add(-1)

				if m.ecCount.Load() == 0 {
					newGroup := m.newGroup.Swap(nullNewGroup).(*NewGroup)
					m.partners.Store(newGroup.NewParts)
					m.p2p.UpdateParticipants(newGroup.NewParts)
					log.Infof("regroup success!!!: new groupID: %v", newGroup.NewParts.GroupID())
				}
			}
		}
	}()
}

func (m *Scheduler) loopSigInToOut() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("tss signature read result loop stopped")
			case result := <-m.sigInToOut:
				log.Infof("finish consensus success, sessionID:%s", result.SessionID)
				info := fmt.Sprintf("tss signature sessionID=%v, groupID=%v, ProposalID=%v", result.SessionID, result.GroupID, result.ProposalID)

				if result.Err != nil {
					log.Errorf("%s, result error:%v", info, result.Err)
				} else {
					switch result.Type {
					case SignBatchTaskSessionType:
						ops := m.operationsQueue.Get(result.ProposalID).(*Operations)
						ops.Signature = secp256k1Signature(result.Data)
						log.Infof("result.Data.Signature: len: %d, result.Data.Signature: %x", len(result.Data.Signature), result.Data.Signature)
						log.Infof("result.Data.SignatureRecovery: len: %d, result.Data.SignatureRecovery: %x", len(result.Data.SignatureRecovery), result.Data.SignatureRecovery)
						log.Infof("ops.Signature: len: %d, ops.Signature: %x, Hash: %v,dataHash: %v", len(ops.Signature), ops.Signature, ops.Hash, ops.DataHash)
						if m.handleSigFinish != nil {
							m.handleSigFinish(ops)
						}
						lo.ForEach(ops.Operation, func(item contracts.Operation, _ int) { m.AddDiscussedTask(item.TaskId) })
						m.operationsQueue.RemoveTopN(ops.TaskID() - 1)
					default:
						log.Infof("tss signature result: %v", result)
					}
				}
			}
		}
	}()
}

func (m *Scheduler) loopDetectionCondition() {
	ticker := time.NewTicker(20 * time.Second)
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("detection condition loop stopped")
			case <-ticker.C:
				latestProposer, err := m.voterContract.Proposer()
				if err != nil {
					log.Errorf("voterContract.Proposer err: %v", err)
				} else {
					proposer := m.Proposer()
					if proposer != latestProposer {
						m.proposer.Store(latestProposer)
					}
				}
			}
		}
	}()
}

func (m *Scheduler) GetDiscussedOperation(id uint64) *Operations {
	ops := m.operationsQueue.Get(id)
	if ops == nil {
		return nil
	}
	return ops.(*Operations)
}

func (m *Scheduler) isCanProposal() bool {
	m.BlockDetectionThreshold()
	proposer, err := m.voterContract.Proposer()
	if err != nil || proposer == zeroAddress {
		proposer = m.Proposer()
	}
	return m.LocalSubmitter() == proposer && m.isJoined()
}

func (m *Scheduler) isJoined() bool {
	return m.Participants().Contains(m.LocalSubmitter())
}

func (m *Scheduler) IsNewJoined() bool {
	return m.newGroup.Load().(*NewGroup).IsNewJoined(m.LocalSubmitter())
}

func (m *Scheduler) Participants() types.Participants {
	if val := m.partners.Load(); val != nil {
		return val.(types.Participants)
	}
	return types.Participants{}
}

type NewGroup struct {
	Event    *db.ParticipantEvent
	NewParts types.Participants
	OldParts types.Participants
}

func (g *NewGroup) IsNewJoined(address common.Address) bool {
	return g.NewParts.Contains(address)
}

var nullNewGroup *NewGroup
