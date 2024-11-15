package tss

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
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

func (t *Service) Validate(msg SessionMessage[TaskId, Msg]) error {
	if t.IsCompleted(msg.TaskID) {
		return fmt.Errorf("taskID:%v, %w", msg.TaskID, ErrTaskCompleted)
	}

	if msg.Proposer != t.proposer {
		return fmt.Errorf("proposer:(%v, %v), %w", msg.Proposer, t.proposer, ErrProposerWrong)
	}

	return nil
}

func (t *Service) GetTask(taskID int64) (*db.Task, error) {
	task := &db.Task{}

	err := t.dbm.GetRelayerDB().
		Preload(clause.Associations).
		Where("task_id", taskID).
		Last(task).
		Error
	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task, err
}

func (t *Service) GenKeyUnSignMsg(task *db.Task) Msg {
	return *helper.SenateTaskMsg
}

func (t *Service) ReShareGroupUnSignMsg(task *db.Task) Msg {
	return *helper.SenateTaskMsg
}

func (t *Service) TaskUnSignMsg(task *db.Task) Msg {
	// todo
	return big.Int{}
}

// handleSessionMsg handler received msg from other node.
func (t *Service) handleSessionMsg(msg SessionMessage[TaskId, Msg]) error {
	task, err := t.GetTask(msg.TaskID)
	if err != nil {
		return err
	}

	//if msg.TaskID < t.currentDoingTaskID {
	//	return fmt.Errorf("task already in progress")
	//}

	session := t.scheduler.GetSession(msg.SessionID)
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
		unSignMsg := t.GenKeyUnSignMsg(task)
		if unSignMsg.String() != msg.Msg.String() {
			return fmt.Errorf("GenKeyUnSignMsg: %w", ErrTaskSignatureMsgWrong)
		}

		_ = t.scheduler.NewGenerateKeySession(
			t.proposer,
			msg.TaskID,
			&msg.Msg,
			t.partners,
		)
	case ReShareGroupSessionType:
		var newPartners Participants // todo
		// check groupID
		if msg.GroupID != helper.SenateGroupID {
			return fmt.Errorf("ReShareGroupSessionType: %w", ErrGroupIdWrong)
		}
		// check msg
		unSignMsg := t.ReShareGroupUnSignMsg(task)
		if unSignMsg.String() != msg.Msg.String() {
			return fmt.Errorf("ReShareGroupUnSignMsg: %w", ErrTaskSignatureMsgWrong)
		}

		_ = t.scheduler.NewReShareGroupSession(
			t.localAddress,
			t.proposer,
			helper.SenateTaskID,
			&msg.Msg,
			t.partners,
			newPartners,
		)
	case SignSessionType:
		localPartySaveData, err := LoadTSSData()
		utils.Assert(err)

		keyDerivationDelta := &big.Int{} // todo

		unSignMsg := t.TaskUnSignMsg(task)
		if unSignMsg.String() != msg.Msg.String() {
			return fmt.Errorf("SignSessionType: %w", ErrTaskSignatureMsgWrong)
		}

		_ = t.scheduler.NewSignSession(
			msg.GroupID,
			msg.SessionID,
			msg.Proposer,
			t.localAddress,
			msg.TaskID,
			&msg.Msg,
			t.partners,
			*localPartySaveData,
			keyDerivationDelta,
		)
	default:
		return fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}

	return nil
}

func (t *Service) Partners() Participants {
	// todo online contact get address list
	return lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}

func (t *Service) proposalTaskSession(task db.ITask) error {
	nonce, err := t.layer2Listener.ContractVotingManager().TssNonce(nil)
	if err != nil {
		return fmt.Errorf("get nonce error: %v", err)
	}

	switch taskRequest := task.(type) {
	case *db.CreateWalletTask:
		coinType := getCoinTypeByChain(taskRequest.Chain)

		var result contracts.TaskPayloadContractWalletCreationResult

		_ = wallet.GenerateAddressByPath(
			t.scheduler.MasterPublicKey(),
			uint32(coinType),
			taskRequest.Account,
			taskRequest.Index,
		)

		resultBytes, err := db.EncodeTaskResult(db.TaskTypeCreateWallet, result)
		if err != nil {
			return err
		}

		// todo
		_, err = serializeMessageToBeSigned(nonce.Uint64(), resultBytes)
		if err != nil {
			return err
		}

		path := wallet.Bip44DerivationPath(uint32(coinType), taskRequest.Account, taskRequest.Index)
		param, err := path.ToParams()
		utils.Assert(err)
		keyDerivationDelta, extendedChildPk, err := wallet.DerivingPubKeyFromPath(t.scheduler.MasterPublicKey(), param.Indexes())
		utils.Assert(err)
		localPartySaveData, err := LoadTSSData()
		utils.Assert(err)
		err = wallet.UpdatePublicKeyAndAdjustBigXj(keyDerivationDelta, localPartySaveData, &extendedChildPk.PublicKey, tss.S256())
		utils.Assert(err)

		t.scheduler.NewSignSession(
			helper.SenateGroupID,
			helper.ZeroSessionID,
			t.proposer,
			t.localAddress,
			helper.SenateTaskID,
			new(big.Int), // todo
			t.partners,
			*localPartySaveData,
			keyDerivationDelta,
		)

		return err
	case *db.DepositTask:

	case *db.WithdrawalTask:
	}

	return nil
}

func (t *Service) handleSigFinish(ctx context.Context, signatureData *tsscommon.SignatureData) {
	//t.rw.Lock()
	//
	//log.Infof("sig finish, taskId:%d, R:%x, S:%x, V:%x", t.state.TssState.CurrentTask.TaskId, signatureData.R, signatureData.S, signatureData.SignatureRecovery)
	//
	//if t.state.TssState.CurrentTask.Submitter == t.localAddress.Hex() {
	//	buf := bytes.NewReader(t.state.TssState.CurrentTask.Context)
	//
	//		// @todo
	//		// generate wallet and send to chain
	//		address := wallet.GenerateAddressByPath(
	//			*(t.scheduler.MasterPublicKey()),
	//			uint32(coinType),
	//			createWalletTask.Account,
	//			createWalletTask.Index,
	//		)
	//		log.Infof("user account address: %s", address)
	//	}
	//}
	//
	//t.rw.Unlock()
}
