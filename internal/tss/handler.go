package tss

import (
	"fmt"
	"math/big"
	"strings"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
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
	if m.IsDiscussed(msg.SeqId) {
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

	task, err := m.stateDB.GetUnCompletedTask(taskID)
	//todo
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return m.GetOnlineTask(taskID)
	//}
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

	detailTask := layer2.DecodeTask(t.Id, t.Result)

	baseTask := db.Task{
		TaskId:    t.Id,
		TaskType:  detailTask.Type(),
		Context:   t.Result,
		Submitter: t.Submitter.Hex(),
		State:     int(t.State),
	}
	detailTask.SetBaseTask(baseTask)

	return detailTask, nil
}

func (m *Scheduler) GenKeyProposal() Proposal {
	return *SenateProposal
}

func (m *Scheduler) ReShareGroupProposal() Proposal {
	return *SenateProposal
}

func (m *Scheduler) isSenateSession(sessionID party.SessionID) bool {
	return sessionID == SenateSessionIDOfECDSA || sessionID == SenateSessionIDOfEDDSA
}

func (m *Scheduler) curveTypeBySenateSession(sessionID party.SessionID) crypto.CurveType {
	switch sessionID {
	case SenateSessionIDOfEDDSA:
		return crypto.EDDSA
	case SenateSessionIDOfECDSA:
		return crypto.ECDSA
	default:
		panic("unimplemented")
	}
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
	log.Debugf("process received proposal id: %v, basePath: %v", msg.SeqId, m.partyData.basePath)

	// todo
	//err := m.Validate(msg)
	//if err != nil {
	//	return err
	//}

	ok := m.OpenSession(msg)
	if ok {
		log.Debugf("open session success, sessionID: %v", msg.SessionID)
		return nil
	}

	log.Debugf("open session fail: session id: %v, msg type: %v,", msg.SessionID, msg.Type)

	var err error
	// build new session
	switch msg.Type {
	case types.GenKeySessionType:
		err = m.JoinGenKeySession(msg)
	case types.ReShareGroupSessionType:
		err = m.JoinReShareGroupSession(msg)

	case types.SignOperationSessionType, types.SignTxSessionType:
		err = m.processTxSign(&msg)

	case types.SignTestOperationSessionType:
		err = m.joinSignOperationSession(msg)
	case types.SignTestTxSessionType: // blockchain wallet tx signature
		task := m.pendingStateTasks.Get(msg.SeqId) // todo
		if task != nil {
			m.joinTxSignatureSession(msg, task)
		} else {
			err = fmt.Errorf("pending task is not exsit")
		}
	case types.SignTaskSessionType: // only used test
		task, errTask := m.GetTask(msg.SeqId)
		if errTask != nil {
			return errTask
		}
		err = m.joinSignTaskSession(msg, task)
	default:
		err = fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}

	if err != nil {
		return err
	}

	_ = m.OpenSession(msg)

	return nil
}

// only used test
func (m *Scheduler) joinTxSignatureSession(msg SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) {
	m.processTxSignForTest(&msg, task)
}

// only used test
func (m *Scheduler) createUserAddressProposal(task *db.CreateWalletTask) (LocalPartySaveData, *big.Int) {
	coinType := types.GetCoinTypeByChain(task.Chain)

	ec := types.GetCurveTypeByCoinType(coinType)

	switch ec {
	case crypto.ECDSA:
		localPartySaveData := m.partyData.GetData(ec)
		userAddress := address.GenerateAddressByPath(localPartySaveData.ECPoint(), uint32(coinType), task.Account, task.Index)
		msg := m.voterContract.EncodeRegisterNewAddress(big.NewInt(int64(task.Account)), task.Chain, big.NewInt(int64(task.Index)), strings.ToLower(userAddress))
		hash := ethCrypto.Keccak256Hash(msg)
		return *localPartySaveData, hash.Big()
	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

// only used test
func (m *Scheduler) joinSignTaskSession(msg SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) error {
	log.Debugf("JoinSignTaskSession: session id: %v, task id:%v, task type: %v", msg.SessionID, task.TaskID(), task.Type())

	switch v := task.(type) {
	case *db.CreateWalletTask:
		_, unSignMsg := m.createUserAddressProposal(v)
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("SignTaskSessionType: %w", ErrTaskSignatureMsgWrong)
		}

		m.NewSignSession(
			msg.SessionID,
			task.TaskID(),
			unSignMsg.String(),
			unSignMsg,
			m.tssSigner(),
		)
	case *db.DepositTask:
	case *db.WithdrawalTask:

	default:
		return fmt.Errorf("taskID %d: %w: %v", task.TaskID(), ErrTaskIdWrong, task.Type())
	}

	return nil
}

// only used test
func (m *Scheduler) processTaskProposal(task pool.Task[uint64]) {
	switch taskData := task.(type) {
	case *db.CreateWalletTask:
		_, unSignMsg := m.createUserAddressProposal(taskData)

		m.NewSignSession(
			ZeroSessionID,
			taskData.TaskId,
			unSignMsg.String(),
			unSignMsg,
			m.tssSigner(),
		)
	case *db.DepositTask:
		_, err := m.stateDB.Account(taskData.TargetAddress)
		if err != nil {
			log.Error("db.DepositTask get account error", err)
			return
		}

		//switch taskData.AssetType {
		//case types.AssetTypeMain:
		//case types.AssetTypeErc20:
		//default:
		//	log.Errorf("unknown asset type: %v", taskData.AssetType)
		//}
	case *db.WithdrawalTask:
		m.processTxSignForTest(nil, taskData)
	}
}
