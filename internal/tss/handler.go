package tss

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (m *Scheduler) GetTask(taskID int64) (*db.Task, error) {
	task := &db.Task{}

	err := m.stateDB.
		Preload(clause.Associations).
		Where("task_id", taskID).
		Last(task).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		itask, err := m.voterContract.Tasks(big.NewInt(taskID))
		if err != nil {
			return nil, err
		}

		return &db.Task{
			TaskId:    uint32(itask.Id.Uint64()),
			Context:   itask.Context,
			Submitter: itask.Submitter.Hex(),
		}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task, err
}

func (m *Scheduler) GenKeyProposal(task *db.Task) Proposal {
	return *helper.SenateProposal
}

func (m *Scheduler) ReShareGroupProposal(task *db.Task) Proposal {
	return *helper.SenateProposal
}

func (m *Scheduler) TaskProposal(task *db.Task) Proposal {
	// todo
	return big.Int{}
}

// processReceivedProposal handler received msg from other node.
func (m *Scheduler) processReceivedProposal(msg SessionMessage[ProposalID, Proposal]) error {
	task, err := m.GetTask(msg.ProposalID)
	if err != nil {
		return err
	}

	//if msg.ProposalID < m.currentDoingTaskID {
	//	return fmt.Errorf("task already in progress")
	//}

	session := m.GetSession(msg.SessionID)
	if session != nil {
		from := session.PartyID(msg.FromPartyId)
		if from != nil && !session.Equal(from.Id) && (msg.IsBroadcast || session.Included(msg.ToPartyIds)) {
			session.Post(msg.State(from))
		}

		return nil
	}
	// build new session
	switch msg.Type {
	case GenKeySessionType:
		// check groupID
		if msg.GroupID != m.Participants().GroupID() {
			return fmt.Errorf("GenKeySessionType: %w", ErrGroupIdWrong)
		}
		// check msg
		unSignMsg := m.GenKeyProposal(task)
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("GenKeyUnSignMsg: %w", ErrTaskSignatureMsgWrong)
		}

		_ = m.NewGenerateKeySession(
			m.CurveType(task),
			msg.ProposalID,
			&msg.Proposal,
		)

	case ReShareGroupSessionType:
		// todo How find new part?
		ng := m.newGroup.Load()
		if ng == nil {
			return fmt.Errorf("newGroup: %w", ErrGroupIdWrong)
		}

		newGroup := ng.(*NewGroup)
		newPartners := newGroup.NewParts
		// check groupID
		if msg.GroupID != m.Participants().GroupID() {
			return fmt.Errorf("ReShareGroupSessionType: %w", ErrGroupIdWrong)
		}
		// check msg
		unSignMsg := m.ReShareGroupProposal(task) // todo add or remove address
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("ReShareGroupUnSignMsg: %w", ErrTaskSignatureMsgWrong)
		}

		_ = m.NewReShareGroupSession(
			m.CurveType(task),
			msg.ProposalID,
			&msg.Proposal,
			m.Participants(),
			newPartners,
		)
	case SignTaskSessionType:
		localPartySaveData, err := LoadTSSData(m.CurveType(task))
		utils.Assert(err)

		unSignMsg := m.TaskProposal(task)
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("SignTaskSessionType: %w", ErrTaskSignatureMsgWrong)
		}

		var keyDerivationDelta *big.Int

		switch task.TaskType {
		case db.TaskTypeCreateWallet:

		case db.TaskTypeDeposit:

		case db.TaskTypeWithdrawal:

		default:
			return fmt.Errorf("taskID %d: %w: %v", task.TaskId, ErrTaskIdWrong, task.TaskType)
		}

		_ = m.NewSignSession(
			m.CurveType(task),
			msg.SessionID,
			msg.ProposalID,
			&msg.Proposal,
			*localPartySaveData,
			keyDerivationDelta,
		)
	case TxSignatureSessionType: // blockchain wallet tx signature
	// todo

	default:
		return fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}

	return nil
}

func Partners() types.Participants {
	// todo online contact get address list
	return lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}

func (m *Scheduler) CurveType(task *db.Task) helper.CurveType {
	// todo
	return helper.ECDSA
}

func (m *Scheduler) processTaskProposal(task db.ITask) {
	switch taskData := task.(type) {
	case *db.CreateWalletTask:
		coinType := getCoinTypeByChain(taskData.Chain)
		taskId := big.NewInt(int64(taskData.TaskId))
		// todo
		var contractAddress common.Address

		var calldata []byte

		unSignMsg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(contractAddress, calldata, taskId)
		if err != nil {
			log.Error("GenerateVerifyTaskUnSignMsg error", err)
			return
		}

		path := wallet.Bip44DerivationPath(uint32(coinType), taskData.Account, taskData.Index)
		param, err := path.ToParams()
		utils.Assert(err)
		localPartySaveData, err := LoadTSSData(m.CurveType(&taskData.Task))
		utils.Assert(err)

		keyDerivationDelta, extendedChildPk, err := wallet.DerivingPubKeyFromPath(*localPartySaveData.ECDSAData().ECDSAPub.ToECDSAPubKey(), param.Indexes())
		utils.Assert(err)

		// todo
		err = wallet.UpdatePublicKeyAndAdjustBigXj(keyDerivationDelta, localPartySaveData.ECDSAData(), &extendedChildPk.PublicKey, tss.S256())
		utils.Assert(err)

		m.NewSignSession(
			m.CurveType(&taskData.Task),
			helper.ZeroSessionID,
			ProposalID(taskData.TaskId),
			unSignMsg.Big(),
			*localPartySaveData,
			keyDerivationDelta,
		)
	case *db.DepositTask:

	case *db.WithdrawalTask:
	}
}
