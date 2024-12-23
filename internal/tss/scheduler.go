package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	. "github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/crypto/ckd"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	eddsaSigning "github.com/bnb-chain/tss-lib/v2/eddsa/signing"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet/bip44"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Scheduler struct {
	isProd          bool
	p2p             p2p.P2PService
	bus             eventbus.Bus
	ctx             context.Context
	cancel          context.CancelFunc
	grw             sync.RWMutex
	groups          map[party.GroupID]*Group
	srw             sync.RWMutex
	sessions        map[party.SessionID]Session[ProposalID]
	proposalSession map[ProposalID]Session[ProposalID]
	crw             sync.RWMutex
	tssClients      map[uint8]suite.TssClient
	sigrw           sync.RWMutex
	sigContext      map[string]*SignerContext // address:SigContext
	sigInToOut      chan *SessionResult[ProposalID, *tsscommon.SignatureData]
	senateInToOut   chan *SessionResult[ProposalID, *LocalPartySaveData]
	partyData       *PartyData
	localSubmitter  common.Address
	submitterChosen *db.SubmitterChosen
	proposer        *atomic.Value // current submitter
	partners        *atomic.Value // types.Participants
	ecCount         *atomic.Int64
	newGroup        *atomic.Value // *NewGroup
	voterContract   layer2.VoterContract
	// only used test
	stateDB            *state.ContractState
	taskQueue          *pool.Pool[uint64] // created state task
	pendingStateTasks  *pool.Pool[uint64] // pending state task
	operationsQueue    *pool.Pool[uint64] // pending batch task
	discussedTaskCache *cache.Cache
	notify             chan struct{}
	currentVoterNonce  *atomic.Uint64
	txContext          sync.Map // taskID:TxContext
}

func (m *Scheduler) TssSigner() common.Address {
	return common.HexToAddress(m.partyData.ECDSALocalData().TssSigner())
}

func (m *Scheduler) tssSigner() *SignerContext {
	return m.GetSigner(m.partyData.ECDSALocalData().TssSigner())
}

func (m *Scheduler) IsMeeting(signDigest string) bool {
	m.srw.RLock()
	defer m.srw.RUnlock()
	_, ok := m.proposalSession[signDigest]

	return ok
}

func (m *Scheduler) GetPublicKey(address string) crypto.PublicKey {
	signer := m.GetSigner(address)
	if signer == nil {
		return nil
	}
	localData := signer.LocalData()
	return localData.PublicKey()
}

func (m *Scheduler) RegisterTssClient(client suite.TssClient) {
	defer m.crw.Unlock()
	m.crw.Lock()
	m.tssClients[client.ChainType()] = client
}

func NewScheduler(isProd bool, p p2p.P2PService, bus eventbus.Bus, stateDB *state.ContractState, voterContract layer2.VoterContract, localSubmitter common.Address) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	pp := atomic.Value{}

	proposer, err := voterContract.Proposer()
	if err != nil {
		log.Warnf("get proposer error, %s", err.Error())
		log.Infof("TssPublicKeys: %v", len(config.TssPublicKeys))
		proposer = ethcrypto.PubkeyToAddress(*config.TssPublicKeys[0]) // genesis
		pp.Store(proposer)
	} else {
		pp.Store(proposer)
	}

	ps := atomic.Value{}

	partners, err := voterContract.Participants()
	if err != nil {
		log.Warnf("get partners error, %s", err.Error())
		partners = lo.Map(config.TssPublicKeys, func(item *ecdsa.PublicKey, _ int) common.Address { return ethcrypto.PubkeyToAddress(*item) })
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
		ctx:                ctx,
		cancel:             cancel,
		isProd:             isProd,
		p2p:                p,
		bus:                bus,
		srw:                sync.RWMutex{},
		grw:                sync.RWMutex{},
		groups:             make(map[party.GroupID]*Group),
		sessions:           make(map[party.SessionID]Session[ProposalID]),
		proposalSession:    make(map[ProposalID]Session[ProposalID]),
		crw:                sync.RWMutex{},
		tssClients:         make(map[uint8]suite.TssClient),
		sigrw:              sync.RWMutex{},
		sigContext:         make(map[string]*SignerContext),
		sigInToOut:         make(chan *SessionResult[ProposalID, *tsscommon.SignatureData], 1024),
		senateInToOut:      make(chan *SessionResult[ProposalID, *LocalPartySaveData], 1024),
		localSubmitter:     localSubmitter,
		proposer:           &pp,
		partners:           &ps,
		newGroup:           newGroup,
		taskQueue:          pool.NewTaskPool[uint64](),
		pendingStateTasks:  pool.NewTaskPool[uint64](),
		operationsQueue:    pool.NewTaskPool[uint64](),
		discussedTaskCache: cache.New(time.Minute*10, time.Minute),
		notify:             make(chan struct{}, 1024),
		stateDB:            stateDB,
		voterContract:      voterContract,
		partyData:          NewPartyData(config.AppConfig.DbDir),
		currentVoterNonce:  currentNonce,
		txContext:          sync.Map{},
	}
}

func (m *Scheduler) Start() {
	m.p2pLoop()
	m.systemProposalLoop()
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
		log.Info("ECDSA PublicKey: ", m.partyData.ECDSALocalData().PublicKeyBase58(), " EDDSA PublicKey: ", m.partyData.EDDSALocalData().PublicKeyBase58())
	}

	log.Infof("********Scheduler master tss ecdsa address********: %v", m.partyData.GetData(crypto.ECDSA).TssSigner())
	log.Infof("localSubmitter: %v, proposer: %v", m.LocalSubmitter(), m.Proposer())
	m.initKnownSigner()
	m.reGroupResultLoop()
	m.loopSigInToOut()
	m.loopDetectionCondition()
	// loop approveProposal
	m.loopApproveProposal()
	log.Info("Scheduler stared success!")
}

func (m *Scheduler) BlockDetectionLatestState() {
	for m.voterContract.IsSyncing() {
		time.Sleep(1 * time.Second)
		log.Warn("l2 info is syncing")
	}
}

func (m *Scheduler) SaveSenateSessionResult(sessionResult *SessionResult[ProposalID, *LocalPartySaveData]) {
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
		crypto.ECDSA,
		SenateProposalIDOfECDSA,
		SenateSessionIDOfECDSA,
		SenateProposal,
	)
	_ = m.NewGenerateKeySession(
		crypto.EDDSA,
		SenateProposalIDOfEDDSA,
		SenateSessionIDOfEDDSA,
		SenateProposal,
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

func (m *Scheduler) GetUserAddress(coinType, account uint32, index uint8) string {
	return address.GenerateAddressByPath(m.partyData.GetData(types.GetCurveTypeByCoinType(int(coinType))).ECPoint(), coinType, account, index)
}

func (m *Scheduler) initKnownSigner() {
	// tss signer
	m.AddSigner(&SignerContext{
		chainType:          types.ChainEthereum,
		localData:          *m.partyData.ECDSALocalData(),
		keyDerivationDelta: nil,
	})

	// add hot address
	coins := []int{types.CoinTypeBTC, types.CoinTypeEVM, types.CoinTypeSOL, types.CoinTypeSUI}
	for _, coin := range coins {
		localPartySaveData, key := m.GenerateDerivationWalletProposal(uint32(coin), 0, 0)
		m.AddSigner(&SignerContext{
			chainType:          types.GetChainByCoinType(coin),
			localData:          localPartySaveData,
			keyDerivationDelta: key,
		})
	}
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
				if m.IsGenesis() {
					if count >= m.Participants().Len() {
						break L
					}
				} else {
					break L
				}
			}
			log.Infof("detection online peer count:%d, threshold:%d", count, threshold)
			time.Sleep(time.Second)
		}
	}
}

func (m *Scheduler) Threshold() int {
	return m.Participants().Threshold()
}

func (m *Scheduler) ECPoint(chainType uint8) *ECPoint {
	return m.partyData.GetDataByChain(chainType).ECPoint()
}

func (m *Scheduler) AddSigner(signer *SignerContext) {
	defer m.sigrw.Unlock()
	m.sigrw.Lock()
	m.sigContext[signer.Address()] = signer
}

func (m *Scheduler) GetSigner(address string) *SignerContext {
	defer m.sigrw.RUnlock()
	m.sigrw.RLock()

	signer := m.sigContext[strings.ToLower(address)]
	if signer != nil {
		return signer
	}

	account, err := m.stateDB.Account(address)
	if err != nil {
		return nil
	}

	local, keyDerivationDelta := m.GenerateDerivationWalletProposal(uint32(types.GetCoinTypeByChain(account.Chain)), uint32(account.Account), uint8(account.Index))

	signer = &SignerContext{
		chainType:          account.Chain,
		localData:          local,
		keyDerivationDelta: keyDerivationDelta,
	}

	m.sigContext[signer.Address()] = signer

	return signer
}

func (m *Scheduler) AddGroup(group *Group) {
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
	defer m.srw.Unlock()
	if _, ok := m.proposalSession[session.ProposalID()]; ok {
		return false
	}

	m.sessions[session.SessionID()] = session
	m.proposalSession[session.ProposalID()] = session

	return true
}

func (m *Scheduler) GetGroup(groupID party.GroupID) *Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return m.groups[groupID]
}

func (m *Scheduler) GetSession(sessionID party.SessionID) Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return m.sessions[sessionID]
}

func (m *Scheduler) GetGroups() []*Group {
	m.grw.RLock()
	defer m.grw.RUnlock()

	return lo.MapToSlice(m.groups, func(_ party.GroupID, group *Group) *Group { return group })
}

func (m *Scheduler) GetSessions() []Session[ProposalID] {
	m.srw.RLock()
	defer m.srw.RUnlock()

	return lo.MapToSlice(m.sessions, func(_ party.SessionID, session Session[ProposalID]) Session[ProposalID] { return session })
}

func (m *Scheduler) ReleaseGroup(groupID party.GroupID) {
	m.grw.Lock()
	defer m.grw.Unlock()
	delete(m.groups, groupID)
}

func (m *Scheduler) SessionRelease(sessionID party.SessionID) {
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
	m.groups = make(map[party.GroupID]*Group)
	m.grw.Unlock()
	m.srw.Lock()
	for _, s := range m.sessions {
		s.Release()
	}

	m.sessions = make(map[party.SessionID]Session[ProposalID])
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
			m.ProcessOperation()

		case <-m.notify:
			m.ProcessOperation()
		}
	}()
}

func (m *Scheduler) ProcessOperation() {
	if m.isCanProposal() && m.isCanNextOperation() {
		log.Info("batch proposal")
		tasks := m.taskQueue.GetTopN(TopN)
		operations := lo.Map(tasks, func(item pool.Task[uint64], index int) contracts.TaskOperation { return *m.operation(item) })
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

		data := lo.Map(tasks, func(item pool.Task[uint64], index int) uint64 { return item.TaskID() })
		batchData := types.BatchData{Ids: data}

		// only ecdsa batch
		m.NewSignOperationSession(
			ZeroSessionID,
			nonce.Uint64(),
			msg.String(), // ProposalID
			msg.Big(),
			batchData.Bytes(),
		)
		m.saveOperations(nonce, operations, dataHash, msg)
	}
}

func (m *Scheduler) isCanNextOperation() bool {
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
	m.p2p.Bind(p2p.MessageTypeTxStatusUpdate, eventbus.EventTxStatusUpdate{})
	m.p2p.Bind(p2p.MessageTypeTxReSign, eventbus.EventTxReSign{})
	tssMsgCh := m.bus.Subscribe(eventbus.EventTssMsg{})
	eventTxStatusUpdate := m.bus.Subscribe(eventbus.EventTxStatusUpdate{})
	eventTxReSign := m.bus.Subscribe(eventbus.EventTxReSign{})

	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-tssMsgCh: // from p2p network
				log.Debugf("Received m msg event")

				e := event.(p2p.Message[json.RawMessage])
				proposal := ConvertP2PMsgData(e).(SessionMessage[ProposalID, Proposal])

				err := m.processReceivedProposal(proposal)
				if err != nil {
					log.Warnf("handle session msg error, %v", err)
				}
			case <-eventTxStatusUpdate:
				// updated task status
			case <-eventTxReSign:
				// rebuild signature
			}
		}
	}()
	log.Info("p2p loop started")
}

// from layer2 log event
func (m *Scheduler) systemProposalLoop() {
	go func() {
		pendingTask := m.bus.Subscribe(eventbus.EventTask{})
		for {
			select {
			case <-m.ctx.Done():
				log.Info("proposal loop stopping...")
				return
			case data := <-pendingTask: // from layer2 log scan
				log.Info("received task from layer2 log scan: ", data)

				switch v := data.(type) {
				case *db.ParticipantEvent: // regroup
					m.processReGroupProposal(v)

				case *db.SubmitterChosen: // charge proposer
					m.submitterChosen = v
					m.proposer.Store(common.HexToAddress(v.Submitter))
				}
			}
		}
	}()

	log.Info("proposal loop started")
}

const TopN = 20

func (m *Scheduler) proposalLoop() {
	go func() {
		pendingTask := m.bus.Subscribe(eventbus.EventTask{})
		for {
			select {
			case <-m.ctx.Done():
				log.Info("proposal loop stopping...")
				return
			case data := <-pendingTask: // from layer2 log scan
				log.Info("received task from layer2 log scan: ", data)

				switch v := data.(type) {
				case db.DetailTask: // task create
					switch v.Status() {
					case db.Completed, db.Failed:
						m.taskQueue.Remove(v.TaskID())
						m.pendingStateTasks.Remove(v.TaskID())
						m.AddDiscussedTask(v.TaskID())
					case db.Pending:
						// todo withdraw
						task, err := m.stateDB.GetUnCompletedTask(v.TaskID())
						if err != nil {
							log.Errorf("get task err:%v", err)
						} else {
							// todo
							m.pendingStateTasks.Add(task)
							// pending task
							if m.isCanProposal() {
								m.processTxSignForTest(nil, task)
							}
						}
					case db.Created:
						if m.IsDiscussed(v.TaskID()) {
							log.Errorf("received task from layer2 is discussed : %v", v.TaskID())
						} else {
							m.taskQueue.Add(v)
							if m.taskQueue.Len() >= TopN {
								m.notify <- struct{}{}
							}
						}
					}
					log.Infof("taskID: %d completed on blockchain", v.TaskID())
				}
			}
		}
	}()

	// test branch
	go m.proposalLoopForTest()
	log.Info("proposal loop started")
}

// only test
func (m *Scheduler) proposalLoopForTest() {
	testPendingTask := m.bus.Subscribe(eventbus.EventTestTask{})
	for {
		select {
		case <-m.ctx.Done():
			log.Info("proposal loop stopping...")
			return
		case data := <-testPendingTask: // from test task
			log.Info("received task from layer2 log scan: ", data)

			switch v := data.(type) {
			case *db.ParticipantEvent: // regroup
				m.processReGroupProposal(v)

			case *db.SubmitterChosen: // charge proposer
				m.submitterChosen = v
				m.proposer.Store(common.HexToAddress(v.Submitter))

			case db.DetailTask:
				if v.Status() == db.Completed || v.Status() == db.Failed {
					log.Infof("taskID: %d completed on blockchain", v.TaskID())
					m.taskQueue.Remove(v.TaskID())
					m.AddDiscussedTask(v.TaskID())
				} else {
					// m.taskQueue.Add(v)
					m.pendingStateTasks.Add(v)
					if m.isCanProposal() {
						log.Info("proposal task", v)
						m.processTaskProposal(v)
					}
				}
			}
		}
	}
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
	m.BlockDetectionLatestState()
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

func (m *Scheduler) GenerateDerivationWalletProposal(coinType, account uint32, index uint8) (LocalPartySaveData, *big.Int) {
	// coinType := types.GetCoinTypeByChain(coinType)
	path := bip44.Bip44DerivationPath(coinType, account, index)
	param, err := path.ToParams()
	utils.Assert(err)
	ec := types.GetCurveTypeByCoinType(int(coinType))
	localPartySaveData := m.partyData.GetData(ec)
	l := *localPartySaveData

	chainCode := big.NewInt(int64(coinType)).Bytes() // todo
	keyDerivationDelta, extendedChildPk, err := ckd.DerivingPubkeyFromPath(l.ECPoint(), chainCode, param.Indexes(), ec.EC())
	utils.Assert(err)
	switch ec {
	case crypto.ECDSA:
		data := []ecdsaKeygen.LocalPartySaveData{*l.ECDSAData()}
		err = ecdsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			data,
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)
		l.SetData(&data[0])
		return l, keyDerivationDelta
	case crypto.EDDSA:
		data := []eddsaKeygen.LocalPartySaveData{*l.EDDSAData()}
		err = eddsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			data,
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)
		l.SetData(&data[0])
		return l, keyDerivationDelta

	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
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
