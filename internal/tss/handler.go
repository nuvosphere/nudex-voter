package tss

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"time"
	"unsafe"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (t *TSSService) handleTssMsg(dataType string, event interface{}) error {
	if t.LocalParty == nil {
		return fmt.Errorf("handleTssUpdate error, t local party not init")
	}

	message, ok := event.(types.TssMessage)
	if !ok {
		return fmt.Errorf("handleTssUpdate error, event %v, is not t update message", event)
	}

	fromPartyID := t.partyIdMap[message.FromPartyId]
	if fromPartyID == nil {
		return fmt.Errorf("fromPartyID %s not found", message.FromPartyId)
	}

	if !message.IsBroadcast && !slices.Contains(message.ToPartyIds, t.LocalParty.PartyID().Id) {
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
		switch dataType {
		case DataTypeTssKeygenMsg:
			if _, err := t.LocalParty.Update(msg); err != nil {
				log.Errorf("Failed to update keygen party: FromPartyID=%v, error=%v", message.FromPartyId, err)
				return
			} else {
				log.Infof("Keygen party updated: FromPartyID=%v, type=%v", message.FromPartyId, msg.Type())
			}
		case DataTypeTssReSharingMsg:
			if _, err := t.reLocalParty.Update(msg); err != nil {
				log.Errorf("Failed to update resharing party: FromPartyID=%v, error=%v", message.FromPartyId, err)
				return
			} else {
				log.Infof("resharing party updated: FromPartyID=%v, type=%v", message.FromPartyId, msg.Type())
			}

		case DataTypeTssSignMsg:
			if t.state.TssState.CurrentTask != nil {
				requestId := getRequestId(t.state.TssState.CurrentTask)

				partyMap := t.sigMap[requestId]
				if partyMap != nil {
					party := partyMap[t.state.TssState.CurrentTask.TaskId]
					if party != nil {
						if _, err := party.Update(msg); err != nil {
							log.Errorf("Failed to update sign party: FromPartyID=%v, error=%v", message.FromPartyId, err)
							return
						} else {
							log.Infof("Sign party updated: FromPartyID=%v, type=%v", message.FromPartyId, msg.Type())
						}
					}
				}
			} else {
				log.Errorf("Failed to update sign party: FromPartyID=%v, current task not found", message.FromPartyId)
			}
		default:
			panic(fmt.Sprintf("unknown data type %v, event %v", dataType, event))
		}
	}()

	return nil
}

func (t *TSSService) sendTssMsg(ctx context.Context, dataType string, event tsslib.Message) (*p2p.Message[types.TssMessage], error) {
	if t.LocalParty == nil {
		return nil, fmt.Errorf("sendTssMsg error, event %v, self not init", event)
	}

	if event.GetFrom().Id != t.LocalParty.PartyID().Id {
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

func (t *TSSService) handleTssKeyOut(ctx context.Context, msg tsslib.Message) (err error) {
	t.keygenRound1P2pMessage, err = t.sendTssMsg(ctx, DataTypeTssKeygenMsg, msg)
	return err
}

func (t *TSSService) handleTssSigOut(ctx context.Context, msg tsslib.Message) error {
	p2pMsg, err := t.sendTssMsg(ctx, DataTypeTssSignMsg, msg)
	if err != nil {
		return fmt.Errorf("handleTssSigOut error, %w", err)
	}

	t.sigRound1P2pMessageMap[p2pMsg.RequestId] = p2pMsg

	return nil
}

func (t *TSSService) checkTimeouts() {
	t.rw.Lock()
	now := time.Now()
	expiredRequests := make([]string, 0)

	for requestId, expireTime := range t.sigTimeoutMap {
		if now.After(expireTime) {
			log.Debugf("Request %s has timed out, removing from sigMap", requestId)
			expiredRequests = append(expiredRequests, requestId)
		}
	}
	t.rw.Unlock()

	for _, requestId := range expiredRequests {
		t.removeSigMap(requestId, true)
	}
}

func (t *TSSService) checkParty(ctx context.Context) {
	if t.LocalParty == nil {
		log.Debug("Party not init, start to setup")
		t.setup()

		return
	}

	localPartySaveData, err := LoadTSSData()
	if err == nil {
		if localPartySaveData.ECDSAPub != nil {
			t.checkSign(ctx)
			return
		}
	}

	if t.setupTime.IsZero() {
		t.setup()
		return
	}

	party := reflect.ValueOf(t.LocalParty.BaseParty).Elem()

	round := party.FieldByName("rnd")
	if !round.CanInterface() {
		round = reflect.NewAt(round.Type(), unsafe.Pointer(round.UnsafeAddr())).Elem()
	}

	rnd, ok := round.Interface().(tsslib.Round)
	if ok {
		if rnd.RoundNumber() == 1 {
			if t.keygenRound1P2pMessage != nil {
				log.Debug("Party set up timeout, send first round p2p message again")

				err = t.p2p.PublishMessage(ctx, *t.keygenRound1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", t.keygenRound1P2pMessage.RequestId)
				}
			}
		} else if rnd.RoundNumber() == 2 {
			if t.keygenRound1P2pMessage != nil && t.round1MessageSendTimes < 3 {
				t.round1MessageSendTimes++
				log.Debugf("Reached round2, send first round p2p message the %d time", t.round1MessageSendTimes)

				err = t.p2p.PublishMessage(ctx, *t.keygenRound1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", t.keygenRound1P2pMessage.RequestId)
				}
			}
		}
	}

	if time.Now().After(t.setupTime.Add(config.AppConfig.TssSigTimeout)) {
		if err := t.LocalParty.FirstRound().Start(); err != nil {
			log.Errorf("TSS keygen process failed to start: %v, start to setup again", err)
			t.setup()

			return
		}

		log.Debug("Party set up timeout, start local party first round again")

		t.setupTime = time.Now()
	}
}

func (t *TSSService) checkSign(ctx context.Context) {
	t.rw.Lock()
	defer t.rw.Unlock()

	for _, partyMap := range t.sigMap {
		for taskId, localParty := range partyMap {
			party := reflect.ValueOf(localParty.BaseParty).Elem()

			round := party.FieldByName("rnd")
			if !round.CanInterface() {
				round = reflect.NewAt(round.Type(), unsafe.Pointer(round.UnsafeAddr())).Elem()
			}

			rnd, ok := round.Interface().(tsslib.Round)
			if ok {
				requestId := "TSS_UPDATE:" + t.localAddress.Hex()
				if rnd.RoundNumber() == 1 {
					if t.sigRound1P2pMessageMap[requestId] != nil {
						log.Debug("Party sign timeout, send first round p2p message again")

						err := t.p2p.PublishMessage(ctx, *(t.sigRound1P2pMessageMap[requestId]))
						if err == nil {
							log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s, TaskId:%d", requestId, taskId)
						}
					}
				} else if rnd.RoundNumber() == 2 {
					if t.sigRound1P2pMessageMap[requestId] != nil && t.sigRound1MessageSendTimesMap[requestId] < 3 {
						t.sigRound1MessageSendTimesMap[requestId]++
						log.Debugf("Tss sign reached round2, send first round p2p message the %d time", t.sigRound1MessageSendTimesMap[requestId])

						err := t.p2p.PublishMessage(ctx, *(t.sigRound1P2pMessageMap[requestId]))
						if err == nil {
							log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s, TaskId:%d", requestId, taskId)
						}
					}
				}
			}
		}
	}
}

func (t *TSSService) sigExists(requestId string) (map[uint32]*signing.LocalParty, bool) {
	t.rw.RLock()
	defer t.rw.RUnlock()
	data, ok := t.sigMap[requestId]

	return data, ok
}

func (t *TSSService) removeSigMap(requestId string, reportTimeout bool) {
	t.rw.Lock()
	defer t.rw.Unlock()

	if reportTimeout {
		taskPartyMap := t.sigMap[requestId]
		t.state.EventBus.Publish(state.EventSigTimeout{}, taskPartyMap)
	}

	delete(t.sigMap, requestId)
	delete(t.sigTimeoutMap, requestId)
	delete(t.sigRound1P2pMessageMap, requestId)
	delete(t.sigRound1MessageSendTimesMap, requestId)
}
