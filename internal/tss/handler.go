package tss

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func (t *TSSService) handleSessionMsg(msg SessionMessage[int32]) error {
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
			int32(helper.SenateTaskID),
			new(big.Int).SetBytes(txHash.Bytes()),
			t.proposer,
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

func (t *TSSService) sendTssMsg(ctx context.Context, dataType string, event tsslib.Message) (*p2p.Message[types.TssMessage], error) {
	if event.GetFrom().Id != t.localParty.PartyID().Id {
		return nil, fmt.Errorf("sendTssMsg error, event %v, not self", event)
	}

	msgWireBytes, _, err := event.WireBytes()
	if err != nil {
		return nil, fmt.Errorf("sendTssMsg parse wire bytes error, event %v", event)
	}

	msg := types.TssMessage{
		FromPartyId:  event.GetFrom().GetId(),
		ToPartyIds:   extractToIds(event),
		IsBroadcast:  event.IsBroadcast(),
		MsgWireBytes: msgWireBytes,
	}

	requestId := fmt.Sprintf("TSS_UPDATE:%s", event.GetFrom().GetId())

	p2pMsg := p2p.Message[types.TssMessage]{
		MessageType: p2p.MessageTypeTssMsg,
		RequestId:   requestId,
		DataType:    dataType,
		Data:        msg,
	}

	return &p2pMsg, t.p2p.PublishMessage(ctx, p2pMsg)
}

func (t *TSSService) IsGenesis() bool {
	if t.localPartySaveData != nil && t.localPartySaveData.ECDSAPub != nil {
		return false
	}

	localPartySaveData, err := LoadTSSData()
	if err != nil {
		log.Errorf("Failed to load TSS data: %v", err)
	}

	if localPartySaveData == nil {
		return true
	}

	t.localPartySaveData = localPartySaveData

	return false
}

func (t *TSSService) Genesis(ctx context.Context) {
	// todo
	log.Info("TSS keygen process started")
}

func (t *TSSService) Partners() []common.Address {
	// todo online contact get address list
	return lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}
