package tss

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/bnb-chain/tss-lib/v2/crypto/ckd"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/chenzhijie/go-web3/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/codec"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound          = gorm.ErrRecordNotFound
	ErrTaskCompleted         = fmt.Errorf("task already completed")
	ErrTaskOrderInconsistent = fmt.Errorf("order of the task is inconsistent")
	ErrTaskIdWrong           = fmt.Errorf("taskId is wrong")
	ErrTaskSignatureMsgWrong = fmt.Errorf("task signature msg is wrong")
	ErrGroupIdWrong          = fmt.Errorf("groupId is wrong")
	ErrSessionIdWrong        = fmt.Errorf("sessionId is wrong")
	ErrProposerWrong         = fmt.Errorf("task proposer is wrong")
)

func (m *Scheduler) Validate(msg SessionMessage[ProposalID, Proposal]) error {
	if m.IsDiscussed(msg.ProposalID) {
		return fmt.Errorf("taskID:%v, %w", msg.ProposalID, ErrTaskCompleted)
	}

	if msg.Proposer != m.Proposer() {
		return fmt.Errorf("proposer:(%v, %v), %w", msg.Proposer, m.Proposer(), ErrProposerWrong)
	}

	return nil
}

func (m *Scheduler) GetTask(taskID uint64) (pool.Task[uint64], error) {
	t := m.taskQueue.Get(taskID)
	if t != nil {
		return t, nil
	}

	task, err := m.stateDB.Task(taskID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m.GetOnlineTask(taskID)
	}

	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task.DetailTask(), err
}

func (m *Scheduler) GetOnlineTask(taskId uint64) (pool.Task[uint64], error) {
	t, err := m.voterContract.Tasks(taskId)
	if err != nil {
		return nil, err
	}

	detailTask := codec.DecodeTask(t.Id, t.Context)

	baseTask := db.Task{
		TaskId:    t.Id,
		TaskType:  detailTask.Type(),
		Context:   t.Context,
		Submitter: t.Submitter.Hex(),
		Status:    int(t.State),
	}
	detailTask.SetBaseTask(baseTask)

	return detailTask, nil
}

func (m *Scheduler) GenKeyProposal() Proposal {
	return *types.SenateProposal
}

func (m *Scheduler) ReShareGroupProposal() Proposal {
	return *types.SenateProposal
}

func (m *Scheduler) isSenateSession(sessionID types.SessionID) bool {
	return sessionID == types.SenateSessionIDOfECDSA || sessionID == types.SenateSessionIDOfEDDSA
}

func (m *Scheduler) OpenSession(msg SessionMessage[ProposalID, Proposal]) bool {
	session := m.GetSession(msg.SessionID)
	if session != nil {
		if !session.Equal(msg.FromPartyId) { // not from self
			from := session.PartyID(msg.FromPartyId)
			if from != nil {
				session.Post(msg.State(from))
			} else {
				if session.Included(msg.ToPartyIds) {
					log.Errorf("session is nil, but included: %v", msg.SessionID)
				}

				if !m.isSenateSession(msg.SessionID) {
					// panic
					panic(fmt.Errorf("session from not is exist:%v basePath: %v", msg.FromPartyId, m.partyData.basePath))
				}

				log.Errorf("session from not is exist:%v basePath: %v", msg.FromPartyId, m.partyData.basePath)
			}
		}

		return true
	}

	log.Debug("session not is exist")

	return false
}

// processReceivedProposal handler received msg from other node.
func (m *Scheduler) processReceivedProposal(msg SessionMessage[ProposalID, Proposal]) error {
	log.Debugf("process received proposal id: %v, basePath: %v", msg.ProposalID, m.partyData.basePath)

	ok := m.OpenSession(msg)
	if ok {
		log.Debugf("open session success, sessionID: %v", msg.SessionID)
		return nil
	}

	log.Debugf("open session fail: session id: %v, msg type: %v,", msg.SessionID, msg.Type)

	var err error
	// build new session
	switch msg.Type {
	case GenKeySessionType:
		err = m.JoinGenKeySession(msg)
	case ReShareGroupSessionType:
		return m.JoinReShareGroupSession(msg)
	case SignTaskSessionType:
		task, errTask := m.GetTask(msg.ProposalID)
		if errTask != nil {
			return errTask
		}

		err = m.JoinSignTaskSession(msg, task)
	case SignBatchTaskSessionType:
		err = m.JoinSignBatchTaskSession(msg)

	case TxSignatureSessionType: // blockchain wallet tx signature
		task, errTask := m.GetTask(msg.ProposalID)
		if errTask != nil {
			return errTask
		}
		m.JoinTxSignatureSession(msg, task)
	default:
		err = fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}

	if err != nil {
		return err
	}

	_ = m.OpenSession(msg)

	return nil
}

func (m *Scheduler) JoinGenKeySession(msg SessionMessage[ProposalID, Proposal]) error {
	// check groupID
	if msg.GroupID != m.Participants().GroupID() {
		return fmt.Errorf("GenKeySessionType: %w", ErrGroupIdWrong)
	}
	// check msg
	unSignMsg := m.GenKeyProposal()
	if unSignMsg.String() != msg.Proposal.String() {
		return fmt.Errorf("GenKeyUnSignMsg: %w", ErrTaskSignatureMsgWrong)
	}

	ec := m.CurveTypeBySenateSession(msg.SessionID)

	_ = m.NewGenerateKeySession(
		ec,
		msg.ProposalID,
		msg.SessionID,
		msg.Signer,
		&msg.Proposal,
	)

	m.OpenSession(msg)

	return nil
}

func (m *Scheduler) isReShareGroup() bool {
	newGroup := m.newGroup.Load().(*NewGroup)
	if newGroup != nullNewGroup {
		return true
	}

	// get latest participants compare local participants
	partners, err := m.voterContract.Participants()
	utils.Assert(err)

	old := m.Participants()
	if old.GroupID() != partners.GroupID() {
		g := &NewGroup{
			NewParts: partners,
			OldParts: old,
		}
		m.newGroup.Store(g)

		return true
	}

	return false
}

func (m *Scheduler) CurveTypeBySenateSession(sessionID types.SessionID) types.CurveType {
	switch sessionID {
	case types.SenateSessionIDOfEDDSA:
		return types.EDDSA
	case types.SenateSessionIDOfECDSA:
		return types.ECDSA
	default:
		panic("unimplemented")
	}
}

func (m *Scheduler) JoinReShareGroupSession(msg SessionMessage[ProposalID, Proposal]) error {
	// todo How find new part?
	is := m.isReShareGroup()
	if !is {
		return fmt.Errorf("not new group")
	}

	newGroup := m.newGroup.Load().(*NewGroup)
	// check groupID
	if msg.GroupID != newGroup.NewParts.GroupID() {
		return fmt.Errorf("JoinReShareGroupSession: session id: %v, msg.SessionID, %w", msg.SessionID, ErrGroupIdWrong)
	}
	// check msg
	unSignMsg := m.ReShareGroupProposal()
	if unSignMsg.Cmp(&msg.Proposal) != 0 {
		return fmt.Errorf("JoinReShareGroupSession: proposal error, session id: %v", msg.SessionID)
	}

	ec := m.CurveTypeBySenateSession(msg.SessionID)

	_ = m.NewReShareGroupSession(
		ec,
		msg.SessionID,
		msg.ProposalID,
		&msg.Proposal,
		newGroup.OldParts,
		newGroup.NewParts,
	)

	m.OpenSession(msg)

	return nil
}

func (m *Scheduler) saveOperations(nonce *big.Int, ops []contracts.Operation, dataHash, hash common.Hash) {
	operations := &Operations{
		Nonce:     nonce,
		Operation: ops,
		Hash:      hash,
		DataHash:  dataHash,
	}
	m.operationsQueue.Add(operations)
	m.currentVoterNonce.Store(nonce.Uint64())
}

func (m *Scheduler) JoinSignBatchTaskSession(msg SessionMessage[ProposalID, Proposal]) error {
	log.Debugf("JoinSignBatchTaskSession: session id: %v, tss nonce(proposalID):%v", msg.SessionID, msg.ProposalID)

	tasks := m.taskQueue.BatchGet(msg.Data)
	operations := lo.Map(tasks, func(item pool.Task[uint64], index int) contracts.Operation { return *m.Operation(item) })

	nonce, dataHash, unSignMsg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(operations)
	if err != nil {
		return fmt.Errorf("batch task generate verify task unsign msg err:%v", err)
	}

	if nonce.Uint64() != msg.ProposalID {
		return fmt.Errorf("nonce error: %v", nonce.Uint64())
	}

	if msg.Proposal.Cmp(unSignMsg.Big()) != 0 {
		return fmt.Errorf("proposal error: %v", msg.Proposal.Text(16))
	}

	// only ecdsa batch
	m.NewMasterSignBatchSession(
		msg.SessionID,
		msg.ProposalID,
		&msg.Proposal,
		msg.Data,
	)
	m.saveOperations(nonce, operations, dataHash, unSignMsg)

	return nil
}

func (m *Scheduler) JoinSignTaskSession(msg SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) error {
	log.Debugf("JoinSignTaskSession: session id: %v, task id:%v, task type: %v", msg.SessionID, task.TaskID(), task.Type())

	switch v := task.(type) {
	case *db.CreateWalletTask:
		localPartySaveData, unSignMsg := m.CreateWalletProposal(v)
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("SignTaskSessionType: %w", ErrTaskSignatureMsgWrong)
		}

		m.NewSignSession(
			msg.SessionID,
			task.TaskID(),
			unSignMsg,
			localPartySaveData,
			nil,
		)
	case *db.DepositTask:
	case *db.WithdrawalTask:

	default:
		return fmt.Errorf("taskID %d: %w: %v", task.TaskID(), ErrTaskIdWrong, task.Type())
	}

	return nil
}

func (m *Scheduler) JoinTxSignatureSession(msg SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) {
	m.processPendingTaskSign(&msg, task)
}

func (m *Scheduler) CurveType(task pool.Task[uint64]) types.CurveType {
	// todo
	return types.ECDSA
}

func (m *Scheduler) CreateWalletProposal(task *db.CreateWalletTask) (types.LocalPartySaveData, *big.Int) {
	coinType := types.GetCoinTypeByChain(task.Chain)

	ec := types.GetCurveTypeByCoinType(coinType)

	switch ec {
	case types.ECDSA:
		localPartySaveData := m.partyData.GetData(ec)
		userAddress := wallet.GenerateAddressByPath(localPartySaveData.ECPoint(), uint32(coinType), task.Account, task.Index)
		msg := m.voterContract.EncodeRegisterNewAddress(big.NewInt(int64(task.Account)), task.Chain, big.NewInt(int64(task.Index)), strings.ToLower(userAddress))
		msg = crypto.Keccak256Hash(msg)

		return *localPartySaveData, new(big.Int).SetBytes(msg)
	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

func (m *Scheduler) GenerateDerivationWalletProposal(coinType, account uint32, index uint8) (types.LocalPartySaveData, *big.Int) {
	// coinType := types.GetCoinTypeByChain(coinType)
	path := wallet.Bip44DerivationPath(coinType, account, index)
	param, err := path.ToParams()
	utils.Assert(err)
	ec := types.GetCurveTypeByCoinType(int(coinType))
	localPartySaveData := m.partyData.GetData(ec)
	l := *localPartySaveData
	var chainCode []byte // todo
	switch ec {
	case types.ECDSA:
		keyDerivationDelta, extendedChildPk, err := ckd.DerivingPubkeyFromPath(l.ECDSAData().ECDSAPub, chainCode, param.Indexes(), ec.EC())
		utils.Assert(err)
		err = ecdsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			[]ecdsaKeygen.LocalPartySaveData{*l.ECDSAData()},
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)

		return l, keyDerivationDelta
	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

func (m *Scheduler) processTaskProposal(task pool.Task[uint64]) {
	switch taskData := task.(type) {
	case *db.CreateWalletTask:
		localPartySaveData, unSignMsg := m.CreateWalletProposal(taskData)

		m.NewSignSession(
			types.ZeroSessionID,
			taskData.TaskId,
			unSignMsg,
			localPartySaveData,
			nil,
		)
	case *db.DepositTask:
		_, err := m.stateDB.Account(taskData.TargetAddress)
		if err != nil {
			log.Error("db.DepositTask get account error", err)
			return
		}

		switch taskData.AssetType {
		case types.AssetTypeMain:
		case types.AssetTypeErc20:
		default:
			log.Errorf("unknown asset type: %v", taskData.AssetType)
		}

	case *db.WithdrawalTask:
	}
}
