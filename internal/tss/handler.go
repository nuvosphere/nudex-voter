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
)

// handleSessionMsg handler received msg from other node.
func (t *Service) handleSessionMsg(msg SessionMessage[TaskId, Msg]) error {
	// todo
	if t.IsCompleted(msg.TaskID) {
		return fmt.Errorf("task already completed")
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
	// todo to validator msg field
	txHash := common.Hash{} // todo

	switch msg.Type {
	case GenKeySessionType:
		_ = t.scheduler.NewGenerateKeySession(
			t.proposer,
			msg.TaskID,
			new(big.Int).SetBytes(txHash.Bytes()),
			int(t.threshold.Load()),
			t.partners,
		)
	case ReShareGroupSessionType:
		// todo
		newThreshold := 0

		var newPartners []common.Address
		_ = t.scheduler.NewReShareGroupSession(
			t.localAddress,
			t.proposer,
			helper.SenateTaskID,
			new(big.Int).SetBytes(txHash.Bytes()),
			int(t.threshold.Load()),
			t.partners,
			newThreshold,
			newPartners,
		)
	case SignSessionType:
		keyDerivationDelta := &big.Int{} // todo
		localPartySaveData, err := LoadTSSData()
		utils.Assert(err)

		_ = t.scheduler.NewSignSession(
			msg.GroupID,
			msg.SessionID,
			msg.Proposer,
			t.localAddress,
			msg.TaskID,
			new(big.Int).SetBytes(txHash.Bytes()),
			int(t.threshold.Load()),
			t.partners,
			*localPartySaveData,
			keyDerivationDelta,
		)
	default:
		return fmt.Errorf("unknown msg type: %v, msg: %v", msg.Type, msg)
	}

	return nil
}

func (t *Service) Partners() []common.Address {
	// todo online contact get address list
	return lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}

func (t *Service) proposalSignTaskSession(dbTask db.Task) error {
	task := db.DecodeTask(dbTask.TaskId, dbTask.Context)

	nonce, err := t.layer2Listener.ContractVotingManager().TssNonce(nil)
	if err != nil {
		return fmt.Errorf("get nonce error for task %x, error: %v", dbTask.Context, err)
	}

	switch taskRequest := task.(type) {
	case *db.CreateWalletTask:
		coinType := getCoinTypeByChain(taskRequest.Chain)

		var result contracts.TaskPayloadContractWalletCreationResult

		_ = wallet.GenerateAddressByPath(
			*(t.scheduler.MasterPublicKey()),
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
		keyDerivationDelta, extendedChildPk, err := wallet.DerivingPubKeyFromPath(*t.scheduler.MasterPublicKey(), param.Indexes())
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
			new(big.Int),
			int(t.threshold.Load()),
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
