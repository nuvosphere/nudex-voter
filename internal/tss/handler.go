package tss

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/utils"
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
