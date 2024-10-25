package tss

import (
	"context"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"reflect"
	"slices"
	"time"
	"unsafe"

	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (tss *TSSService) handleTssUpdate(event interface{}) error {
	if tss.Party == nil {
		return fmt.Errorf("handleTssUpdate error, tss local party not init")
	}
	message, ok := event.(types.TssUpdateMessage)
	if !ok {
		return fmt.Errorf("handleTssUpdate error, event %v, is not tss update message", event)
	}

	fromPartyID := tss.partyIdMap[message.FromPartyId]
	if fromPartyID == nil {
		return fmt.Errorf("fromPartyID %s not found", message.FromPartyId)
	}

	if !message.IsBroadcast && !slices.Contains(message.ToPartyIds, tss.Party.PartyID().Id) {
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

	go func() {
		if _, err := tss.Party.Update(msg); err != nil {
			log.Errorf("Failed to update party: FromPartyID=%v, error=%v", message.FromPartyId, err)
			return
		} else {
			log.Infof("Party updated: FromPartyID=%v, type=%v", message.FromPartyId, msg.Type())
		}
	}()

	return nil
}

func (tss *TSSService) handleTssKeyOut(ctx context.Context, event tsslib.Message) error {
	if tss.Party == nil {
		return fmt.Errorf("handleTssKeyOut error, event %v, self not init", event)
	}
	if event.GetFrom().Id != tss.Party.PartyID().Id {
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
	if event.Type() == "binance.tsslib.ecdsa.keygen.KGRound1Message" {
		tss.keygenRound1P2pMessage = &p2pMsg
	} else if event.Type() == "binance.tsslib.ecdsa.signing.SignRound1Message1" {
		tss.sigRound1P2pMessageMap[requestId] = &p2pMsg
	}

	return err
}

func (tss *TSSService) checkTimeouts() {
	tss.sigMu.Lock()
	now := time.Now()
	expiredRequests := make([]string, 0)

	for requestId, expireTime := range tss.sigTimeoutMap {
		if now.After(expireTime) {
			log.Debugf("Request %s has timed out, removing from sigMap", requestId)
			expiredRequests = append(expiredRequests, requestId)
		}
	}
	tss.sigMu.Unlock()

	for _, requestId := range expiredRequests {
		tss.removeSigMap(requestId, true)
	}
}

func (tss *TSSService) check(ctx context.Context) {
	if tss.Party == nil {
		log.Debug("Party not init, start to setup")
		tss.setup()
		return
	}

	localPartySaveData, err := loadTSSData()
	if err == nil {
		if localPartySaveData.ECDSAPub != nil {
			tss.checkSign(ctx)
			return
		}
	}

	if tss.setupTime.IsZero() {
		tss.setup()
		return
	}

	party := reflect.ValueOf(tss.Party.BaseParty).Elem()
	round := party.FieldByName("rnd")
	if !round.CanInterface() {
		round = reflect.NewAt(round.Type(), unsafe.Pointer(round.UnsafeAddr())).Elem()
	}
	rnd, ok := round.Interface().(tsslib.Round)
	if ok {
		if rnd.RoundNumber() == 1 {
			if tss.keygenRound1P2pMessage != nil {
				log.Debug("Party set up timeout, send first round p2p message again")
				err = tss.libp2p.PublishMessage(ctx, *tss.keygenRound1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", tss.keygenRound1P2pMessage.RequestId)
				}
			}
		} else if rnd.RoundNumber() == 2 {
			if tss.keygenRound1P2pMessage != nil && tss.round1MessageSendTimes < 3 {
				tss.round1MessageSendTimes++
				log.Debugf("Reached round2, send first round p2p message the %d time", tss.round1MessageSendTimes)
				err = tss.libp2p.PublishMessage(ctx, *tss.keygenRound1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", tss.keygenRound1P2pMessage.RequestId)
				}
			}
		}
	}

	if time.Now().After(tss.setupTime.Add(config.AppConfig.TssSigTimeout)) {
		if err := tss.Party.FirstRound().Start(); err != nil {
			log.Errorf("TSS keygen process failed to start: %v, start to setup again", err)
			tss.setup()
			return
		}
		log.Debug("Party set up timeout, start local party first round again")
		tss.setupTime = time.Now()
	}
}

func (tss *TSSService) checkSign(ctx context.Context) {
	tss.sigMu.Lock()

	for _, partyMap := range tss.sigMap {
		for taskId, localParty := range partyMap {
			party := reflect.ValueOf(localParty.BaseParty).Elem()
			round := party.FieldByName("rnd")
			if !round.CanInterface() {
				round = reflect.NewAt(round.Type(), unsafe.Pointer(round.UnsafeAddr())).Elem()
			}
			rnd, ok := round.Interface().(tsslib.Round)
			if ok {
				requestId := "TSS_UPDATE:" + tss.Address.Hex()
				if rnd.RoundNumber() == 1 {
					if tss.sigRound1P2pMessageMap[requestId] != nil {
						log.Debug("Party sign timeout, send first round p2p message again")
						err := tss.libp2p.PublishMessage(ctx, *(tss.sigRound1P2pMessageMap[requestId]))
						if err == nil {
							log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s, TaskId:%d", requestId, taskId)
						}
					}
				} else if rnd.RoundNumber() == 2 {
					if tss.sigRound1P2pMessageMap[requestId] != nil && tss.sigRound1MessageSendTimesMap[requestId] < 3 {
						tss.sigRound1MessageSendTimesMap[requestId]++
						log.Debugf("Tss sign reached round2, send first round p2p message the %d time", tss.sigRound1MessageSendTimesMap[requestId])
						err := tss.libp2p.PublishMessage(ctx, *(tss.sigRound1P2pMessageMap[requestId]))
						if err == nil {
							log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s, TaskId:%d", requestId, taskId)
						}
					}
				}
			}
		}
	}

	tss.sigMu.Unlock()
}

func (tss *TSSService) sigExists(requestId string) (map[uint64]*signing.LocalParty, bool) {
	tss.sigMu.RLock()
	defer tss.sigMu.RUnlock()
	data, ok := tss.sigMap[requestId]
	return data, ok
}

func (tss *TSSService) removeSigMap(requestId string, reportTimeout bool) {
	tss.sigMu.Lock()
	defer tss.sigMu.Unlock()
	if reportTimeout {
		taskPartyMap := tss.sigMap[requestId]
		tss.state.EventBus.Publish(state.SigTimeout, taskPartyMap)
	}
	delete(tss.sigMap, requestId)
	delete(tss.sigTimeoutMap, requestId)
	delete(tss.sigRound1P2pMessageMap, requestId)
	delete(tss.sigRound1MessageSendTimesMap, requestId)
}
