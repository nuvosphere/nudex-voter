package tss

import (
	"context"
	"fmt"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"slices"
)

func (tss *TSSService) handleTssKeyOut(ctx context.Context, event interface{}) error {
	message, ok := event.(tsslib.Message)
	if !ok {
		return fmt.Errorf("handleTssKeyOut error, event %v, is not tss lib message", event)
	}
	if message.GetFrom().Id != tss.party.PartyID().Id {
		return fmt.Errorf("handleTssKeyOut error, event %v, not self", event)
	}

	tss.sigMu.Lock()
	defer tss.sigMu.Unlock()

	msgWireBytes, _, err := message.WireBytes()
	if err != nil {
		return fmt.Errorf("handleTssKeyOut parse wire bytes error, event %v", event)
	}

	tssUpdateMessage := types.TssUpdateMessage{
		FromPartyId:  message.GetFrom().GetId(),
		ToPartyIds:   extractToIds(message),
		IsBroadcast:  message.IsBroadcast(),
		MsgWireBytes: msgWireBytes,
	}

	requestId := fmt.Sprintf("TSS_UPDATE:%s", message.GetFrom().GetId())

	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeTssUpdate,
		RequestId:   requestId,
		DataType:    p2p.DataTypeTssUpdateMessage,
		Data:        tssUpdateMessage,
	}

	err = tss.libp2p.PublishMessage(ctx, p2pMsg)
	if err != nil {
		log.Debugf("Publish p2p message tssUpdateMessage: RequestId=%s, Data=%v",
			requestId, tssUpdateMessage)
	}
	return err
}

func (tss *TSSService) handleTssUpdate(event interface{}) error {
	message, ok := event.(types.TssUpdateMessage)
	if !ok {
		return fmt.Errorf("handleTssUpdate error, event %v, is not tss update message", event)
	}

	fromPartyID := tss.partyIdMap[message.FromPartyId]
	if fromPartyID == nil {
		return fmt.Errorf("fromPartyID %s not found", message.FromPartyId)
	}

	if !message.IsBroadcast && slices.Contains(message.ToPartyIds, tss.party.PartyID().Id) {
		log.Debugf("PartyId not one of p2p message receiver: %v", message.ToPartyIds)
		return nil
	}

	tss.sigMu.Lock()

	ok, err := tss.party.UpdateFromBytes(message.MsgWireBytes, fromPartyID, message.IsBroadcast)
	if err != nil && !ok {
		tss.sigMu.Unlock()
		return err
	}

	log.Infof("party updated: FromPartyID=%v", message.FromPartyId)

	tss.sigMu.Unlock()
	return nil
}

func extractToIds(message tsslib.Message) []string {
	recipients := message.GetTo()

	ids := make([]string, len(recipients))

	for i, recipient := range recipients {
		ids[i] = recipient.GetId()
	}

	return ids
}
