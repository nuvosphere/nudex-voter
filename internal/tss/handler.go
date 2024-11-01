package tss

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"
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

func (tss *TSSService) handleTssMsg(event interface{}) error {
	if tss.LocalParty == nil {
		return fmt.Errorf("handleTssUpdate error, tss local party not init")
	}

	message, ok := event.(types.TssMessage)
	if !ok {
		return fmt.Errorf("handleTssUpdate error, event %v, is not tss update message", event)
	}

	fromPartyID := tss.partyIdMap[message.FromPartyId]
	if fromPartyID == nil {
		return fmt.Errorf("fromPartyID %s not found", message.FromPartyId)
	}

	if !message.IsBroadcast && !slices.Contains(message.ToPartyIds, tss.LocalParty.PartyID().Id) {
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
		if strings.HasPrefix(msg.Type(), "binance.tsslib.ecdsa.keygen.") {
			if _, err := tss.LocalParty.Update(msg); err != nil {
				log.Errorf("Failed to update keygen party: FromPartyID=%v, error=%v", message.FromPartyId, err)
				return
			} else {
				log.Infof("Keygen party updated: FromPartyID=%v, type=%v", message.FromPartyId, msg.Type())
			}
		} else if strings.HasPrefix(msg.Type(), "binance.tsslib.ecdsa.signing.") {
			if tss.state.TssState.CurrentTask != nil {
				requestId := getRequestId(tss.state.TssState.CurrentTask)

				partyMap := tss.sigMap[requestId]
				if partyMap != nil {
					party := partyMap[int32(tss.state.TssState.CurrentTask.TaskId)]
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
		}
	}()

	return nil
}

func (tss *TSSService) sendTssMsg(ctx context.Context, dataType string, event tsslib.Message) (*p2p.Message[types.TssMessage], error) {
	if tss.LocalParty == nil {
		return nil, fmt.Errorf("sendTssMsg error, event %v, self not init", event)
	}

	if event.GetFrom().Id != tss.LocalParty.PartyID().Id {
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

	return &p2pMsg, tss.p2p.PublishMessage(ctx, p2pMsg)
}

func (tss *TSSService) handleTssKeyOut(ctx context.Context, msg tsslib.Message) (err error) {
	tss.keygenRound1P2pMessage, err = tss.sendTssMsg(ctx, DataTypeTssKeygenMsg, msg)
	return err
}

func (tss *TSSService) handleTssSigOut(ctx context.Context, msg tsslib.Message) error {
	p2pMsg, err := tss.sendTssMsg(ctx, DataTypeTssSignMsg, msg)
	if err != nil {
		return fmt.Errorf("handleTssSigOut error, %w", err)
	}

	tss.sigRound1P2pMessageMap[p2pMsg.RequestId] = p2pMsg

	return nil
}

func (tss *TSSService) checkTimeouts() {
	tss.rw.Lock()
	now := time.Now()
	expiredRequests := make([]string, 0)

	for requestId, expireTime := range tss.sigTimeoutMap {
		if now.After(expireTime) {
			log.Debugf("Request %s has timed out, removing from sigMap", requestId)
			expiredRequests = append(expiredRequests, requestId)
		}
	}
	tss.rw.Unlock()

	for _, requestId := range expiredRequests {
		tss.removeSigMap(requestId, true)
	}
}

func (tss *TSSService) checkParty(ctx context.Context) {
	if tss.LocalParty == nil {
		log.Debug("Party not init, start to setup")
		tss.setup()

		return
	}

	localPartySaveData, err := LoadTSSData()
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

	party := reflect.ValueOf(tss.LocalParty.BaseParty).Elem()

	round := party.FieldByName("rnd")
	if !round.CanInterface() {
		round = reflect.NewAt(round.Type(), unsafe.Pointer(round.UnsafeAddr())).Elem()
	}

	rnd, ok := round.Interface().(tsslib.Round)
	if ok {
		if rnd.RoundNumber() == 1 {
			if tss.keygenRound1P2pMessage != nil {
				log.Debug("Party set up timeout, send first round p2p message again")

				err = tss.p2p.PublishMessage(ctx, *tss.keygenRound1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", tss.keygenRound1P2pMessage.RequestId)
				}
			}
		} else if rnd.RoundNumber() == 2 {
			if tss.keygenRound1P2pMessage != nil && tss.round1MessageSendTimes < 3 {
				tss.round1MessageSendTimes++
				log.Debugf("Reached round2, send first round p2p message the %d time", tss.round1MessageSendTimes)

				err = tss.p2p.PublishMessage(ctx, *tss.keygenRound1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", tss.keygenRound1P2pMessage.RequestId)
				}
			}
		}
	}

	if time.Now().After(tss.setupTime.Add(config.AppConfig.TssSigTimeout)) {
		if err := tss.LocalParty.FirstRound().Start(); err != nil {
			log.Errorf("TSS keygen process failed to start: %v, start to setup again", err)
			tss.setup()

			return
		}

		log.Debug("Party set up timeout, start local party first round again")

		tss.setupTime = time.Now()
	}
}

func (tss *TSSService) checkSign(ctx context.Context) {
	tss.rw.Lock()
	defer tss.rw.Unlock()

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

						err := tss.p2p.PublishMessage(ctx, *(tss.sigRound1P2pMessageMap[requestId]))
						if err == nil {
							log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s, TaskId:%d", requestId, taskId)
						}
					}
				} else if rnd.RoundNumber() == 2 {
					if tss.sigRound1P2pMessageMap[requestId] != nil && tss.sigRound1MessageSendTimesMap[requestId] < 3 {
						tss.sigRound1MessageSendTimesMap[requestId]++
						log.Debugf("Tss sign reached round2, send first round p2p message the %d time", tss.sigRound1MessageSendTimesMap[requestId])

						err := tss.p2p.PublishMessage(ctx, *(tss.sigRound1P2pMessageMap[requestId]))
						if err == nil {
							log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s, TaskId:%d", requestId, taskId)
						}
					}
				}
			}
		}
	}
}

func (tss *TSSService) sigExists(requestId string) (map[int32]*signing.LocalParty, bool) {
	tss.rw.RLock()
	defer tss.rw.RUnlock()
	data, ok := tss.sigMap[requestId]

	return data, ok
}

func (tss *TSSService) removeSigMap(requestId string, reportTimeout bool) {
	tss.rw.Lock()
	defer tss.rw.Unlock()

	if reportTimeout {
		taskPartyMap := tss.sigMap[requestId]
		tss.state.EventBus.Publish(state.EventSigTimeout{}, taskPartyMap)
	}

	delete(tss.sigMap, requestId)
	delete(tss.sigTimeoutMap, requestId)
	delete(tss.sigRound1P2pMessageMap, requestId)
	delete(tss.sigRound1MessageSendTimesMap, requestId)
}
