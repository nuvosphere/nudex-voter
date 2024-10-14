package tss

import (
	"context"
	"fmt"
	"slices"

	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (tss *TSSService) handleTssKeyOut(ctx context.Context, event tsslib.Message) error {
	if tss.party == nil {
		return fmt.Errorf("handleTssKeyOut error, event %v, self not init", event)
	}

	if event.GetFrom().Id != tss.party.PartyID().Id {
		return fmt.Errorf("handleTssKeyOut error, event %v, not self", event)
	}

	msgWireBytes, _, err := event.WireBytes()
	if err != nil {
		return fmt.Errorf("handleTssKeyOut parse wire bytes error, event %v", event)
	}

	tssUpdateMessage := types.TssUpdateMessage{
		FromPartyId:  event.GetFrom().GetId(),
		ToPartyIds:   extractToIds(event),
		IsBroadcast:  event.IsBroadcast(),
		MsgWireBytes: msgWireBytes,
	}

	requestId := fmt.Sprintf("TSS_UPDATE:%s", event.GetFrom().GetId())

	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeTssUpdate,
		RequestId:   requestId,
		DataType:    p2p.DataTypeTssUpdateMessage,
		Data:        tssUpdateMessage,
	}

	err = tss.libp2p.PublishMessage(ctx, p2pMsg)
	if err == nil {
		log.Debugf("Publish p2p message tssUpdateMessage: RequestId=%s, IsBroadcast=%v, ToPartyIds=%v",
			requestId, tssUpdateMessage.IsBroadcast, tssUpdateMessage.ToPartyIds)
	}
	return err
}

func (tss *TSSService) handleTssUpdate(event interface{}) error {
	if tss.party == nil {
		return fmt.Errorf("handleTssUpdate error, tss local party not int")
	}
	message, ok := event.(types.TssUpdateMessage)
	if !ok {
		return fmt.Errorf("handleTssUpdate error, event %v, is not tss update message", event)
	}

	fromPartyID := tss.partyIdMap[message.FromPartyId]
	if fromPartyID == nil {
		return fmt.Errorf("fromPartyID %s not found", message.FromPartyId)
	}

	if !message.IsBroadcast && !slices.Contains(message.ToPartyIds, tss.party.PartyID().Id) {
		log.Debugf("PartyId not one of p2p message receiver: %v", message.ToPartyIds)
		return nil
	}

	msg, err := tsslib.ParseWireMessage(
		message.MsgWireBytes,
		fromPartyID,
		message.IsBroadcast)
	if err != nil {
		return err
	}

	ok, tssErr := tss.party.Update(msg)
	if !ok && tssErr != nil {
		return tssErr.Cause()
	}

	log.Infof("party updated: FromPartyID=%v, type=%v", message.FromPartyId, msg.Type())

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
