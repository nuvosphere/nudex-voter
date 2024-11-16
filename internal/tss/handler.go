package tss

import (
	"crypto/ecdsa"
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

func (m *Scheduler) Validate(msg SessionMessage[TaskId, Msg]) error {
	if m.IsCompleted(msg.TaskID) {
		return fmt.Errorf("taskID:%v, %w", msg.TaskID, ErrTaskCompleted)
	}

	if msg.Proposer != m.proposer {
		return fmt.Errorf("proposer:(%v, %v), %w", msg.Proposer, m.proposer, ErrProposerWrong)
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
	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task, err
}

func (m *Scheduler) GenKeyUnSignMsg(task *db.Task) Msg {
	return *helper.SenateTaskMsg
}

func (m *Scheduler) ReShareGroupUnSignMsg(task *db.Task) Msg {
	return *helper.SenateTaskMsg
}

func (m *Scheduler) TaskUnSignMsg(task *db.Task) Msg {
	// todo
	return big.Int{}
}

// handleSessionMsg handler received msg from other node.
func (m *Scheduler) handleSessionMsg(msg SessionMessage[TaskId, Msg]) error {
	task, err := m.GetTask(msg.TaskID)
	if err != nil {
		return err
	}

	//if msg.TaskID < m.currentDoingTaskID {
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
		if msg.GroupID != helper.SenateGroupID {
			return fmt.Errorf("GenKeySessionType: %w", ErrGroupIdWrong)
		}
		// check msg
		unSignMsg := m.GenKeyUnSignMsg(task)
		if unSignMsg.String() != msg.Msg.String() {
			return fmt.Errorf("GenKeyUnSignMsg: %w", ErrTaskSignatureMsgWrong)
		}

		_ = m.NewGenerateKeySession(
			m.proposer,
			msg.TaskID,
			&msg.Msg,
			m.partners,
		)
	case ReShareGroupSessionType:
		// todo How find new part?
		var newPartners types.Participants // todo
		// check groupID
		if msg.GroupID != helper.SenateGroupID {
			return fmt.Errorf("ReShareGroupSessionType: %w", ErrGroupIdWrong)
		}
		// check msg
		unSignMsg := m.ReShareGroupUnSignMsg(task)
		if unSignMsg.String() != msg.Msg.String() {
			return fmt.Errorf("ReShareGroupUnSignMsg: %w", ErrTaskSignatureMsgWrong)
		}

		_ = m.NewReShareGroupSession(
			m.LocalSubmitter(),
			m.proposer,
			helper.SenateTaskID,
			&msg.Msg,
			m.partners,
			newPartners,
		)
	case SignTaskSessionType:
		masterLocalPartySaveData, err := LoadTSSData()
		utils.Assert(err)

		unSignMsg := m.TaskUnSignMsg(task)
		if unSignMsg.String() != msg.Msg.String() {
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
			msg.GroupID,
			msg.SessionID,
			msg.Proposer,
			m.LocalSubmitter(),
			msg.TaskID,
			&msg.Msg,
			m.partners,
			*masterLocalPartySaveData,
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

func (m *Scheduler) proposalTaskSession(task db.ITask) error {
	switch taskData := task.(type) {
	case *db.CreateWalletTask:
		coinType := getCoinTypeByChain(taskData.Chain)

		_ = wallet.GenerateAddressByPath(
			m.MasterPublicKey(),
			uint32(coinType),
			taskData.Account,
			taskData.Index,
		)

		path := wallet.Bip44DerivationPath(uint32(coinType), taskData.Account, taskData.Index)
		param, err := path.ToParams()
		utils.Assert(err)
		keyDerivationDelta, extendedChildPk, err := wallet.DerivingPubKeyFromPath(m.MasterPublicKey(), param.Indexes())
		utils.Assert(err)
		localPartySaveData, err := LoadTSSData()
		utils.Assert(err)
		err = wallet.UpdatePublicKeyAndAdjustBigXj(keyDerivationDelta, localPartySaveData, &extendedChildPk.PublicKey, tss.S256())
		utils.Assert(err)

		m.NewSignSession(
			helper.SenateGroupID,
			helper.ZeroSessionID,
			m.proposer,
			m.LocalSubmitter(),
			helper.SenateTaskID,
			new(big.Int), // todo
			m.partners,
			*localPartySaveData,
			keyDerivationDelta,
		)

		return err
	case *db.DepositTask:

	case *db.WithdrawalTask:
	}

	return nil
}
