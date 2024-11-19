package tss

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
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

func (m *Scheduler) GenKeyProposal() Proposal {
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
		return m.JoinGenKeySession(msg)
	case ReShareGroupSessionType:
		return m.JoinReShareGroupSession(msg, task)
	case SignTaskSessionType:
		return m.JoinSignTaskSession(msg, task)
	case TxSignatureSessionType: // blockchain wallet tx signature
		return m.JoinTxSignatureSession(msg, task)
	default:
		return fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}
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

	ec := helper.ECDSA

	switch msg.ProposalID {
	case helper.SenateProposalIDOfEDDSA:
		ec = helper.EDDSA
	}

	_ = m.NewGenerateKeySession(
		ec,
		msg.ProposalID,
		msg.SessionID,
		msg.Signer,
		&msg.Proposal,
	)

	return nil
}

func (m *Scheduler) JoinReShareGroupSession(msg SessionMessage[ProposalID, Proposal], task *db.Task) error {
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

	ec := helper.ECDSA

	switch msg.ProposalID {
	case helper.SenateProposalIDOfEDDSA:
		ec = helper.EDDSA
	}

	_ = m.NewReShareGroupSession(
		ec,
		msg.ProposalID,
		&msg.Proposal,
		m.Participants(),
		newPartners,
	)

	return nil
}

func (m *Scheduler) JoinSignTaskSession(msg SessionMessage[ProposalID, Proposal], task *db.Task) error {
	//localPartySaveData := m.partyData.GetData(ec)
	//unSignMsg := m.TaskProposal(task)
	//if unSignMsg.String() != msg.Proposal.String() {
	//	return fmt.Errorf("SignTaskSessionType: %w", ErrTaskSignatureMsgWrong)
	//}
	ec := m.CurveType(task)

	switch task.TaskType {
	case db.TaskTypeCreateWallet:
		localPartySaveData, keyDerivationDelta, unSignMsg := m.GenerateCreateWalletProposal(task.CreateWalletTask)
		if unSignMsg.String() != msg.Proposal.String() {
			return fmt.Errorf("SignTaskSessionType: %w", ErrTaskSignatureMsgWrong)
		}

		m.NewSignSession(
			ec,
			msg.SessionID,
			ProposalID(task.TaskId),
			unSignMsg,
			localPartySaveData,
			keyDerivationDelta,
		)

	case db.TaskTypeDeposit:

	case db.TaskTypeWithdrawal:

	default:
		return fmt.Errorf("taskID %d: %w: %v", task.TaskId, ErrTaskIdWrong, task.TaskType)
	}

	//_ = m.NewSignSession(
	//	ec,
	//	msg.SessionID,
	//	msg.ProposalID,
	//	&msg.Proposal,
	//	*localPartySaveData,
	//	keyDerivationDelta,
	//)

	return nil
}

func (m *Scheduler) JoinTxSignatureSession(msg SessionMessage[ProposalID, Proposal], task *db.Task) error {
	return nil
}

func (m *Scheduler) CurveType(task *db.Task) helper.CurveType {
	// todo
	return helper.ECDSA
}

func (m *Scheduler) GenerateCreateWalletProposal(task *db.CreateWalletTask) (helper.LocalPartySaveData, *big.Int, *big.Int) {
	//taskId := big.NewInt(int64(task.TaskId))
	//var (
	//	contractAddress common.Address
	//	calldata        []byte
	//)
	//unSignMsg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(contractAddress, calldata, taskId)
	//if err != nil {
	//	log.Error("GenerateVerifyTaskUnSignMsg error", err)
	//	return
	//}
	coinType := getCoinTypeByChain(task.Chain)
	path := wallet.Bip44DerivationPath(uint32(coinType), task.Account, task.Index)
	param, err := path.ToParams()
	utils.Assert(err)

	ec := m.CurveType(&task.Task)
	localPartySaveData := m.partyData.GetData(ec)

	l := *localPartySaveData

	switch ec {
	case helper.ECDSA:
		keyDerivationDelta, extendedChildPk, err := wallet.DerivingPubKeyFromPath(*l.ECDSAData().ECDSAPub.ToECDSAPubKey(), param.Indexes())
		utils.Assert(err)

		err = wallet.UpdatePublicKeyAndAdjustBigXj(keyDerivationDelta, l.ECDSAData(), &extendedChildPk.PublicKey, ec.EC())
		utils.Assert(err)

		return l, keyDerivationDelta, big.NewInt(100) // todo
	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

func (m *Scheduler) processTaskProposal(task db.ITask) {
	switch taskData := task.(type) {
	case *db.CreateWalletTask:
		ec := m.CurveType(&taskData.Task)
		localPartySaveData, keyDerivationDelta, unSignMsg := m.GenerateCreateWalletProposal(taskData)

		m.NewSignSession(
			ec,
			helper.ZeroSessionID,
			ProposalID(taskData.TaskId),
			unSignMsg,
			localPartySaveData,
			keyDerivationDelta,
		)
	case *db.DepositTask:
		account := &db.Account{}

		err := m.stateDB.
			Preload(clause.Associations).
			Where("chain_id = ? AND address = ?", taskData.ChainId, taskData.TargetAddress).
			Last(account).
			Error
		if err != nil {
			log.Error("db.DepositTask get account error", err)
			return
		}

		switch taskData.AssetType {
		case db.AssetTypeMain:
		case db.AssetTypeErc20:
		default:
			log.Errorf("unknown asset type: %v", taskData.AssetType)
		}

	case *db.WithdrawalTask:
	}
}
