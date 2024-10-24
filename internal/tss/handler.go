package tss

import (
	"context"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"reflect"
	"time"
	"unsafe"

	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (tss *TSSService) handleSigStart(ctx context.Context, event interface{}) {
	switch e := event.(type) {
	case types.MsgSignCreateWalletMessage:
		log.Debugf("Event handleSigStart is of type MsgSignCreateWalletMessage, request id %s", e.RequestId)
		if err := tss.handleSignCreateWalletStart(ctx, e); err != nil {
			log.Errorf("Error handleSigStart MsgSignCreateWalletMessage, %v", err)
			tss.state.EventBus.Publish(state.SigFailed, e)
		}
	default:
		log.Debug("Unknown event handleSigStart type")
	}
}

func (tss *TSSService) handleSigReceive(ctx context.Context, event interface{}) {
}

func (tss *TSSService) handleSigFailed(ctx context.Context, event interface{}, reason string) {
	if e, ok := event.(map[uint64]*signing.LocalParty); ok {
		taskId := tss.state.TssState.CurrentTask.TaskId
		if _, exists := e[taskId]; exists {
			tss.state.TssState.CurrentTask = nil
		}
		for key := range e {
			log.Infof("handle sig failed, taskId:%d, reason:%s", key, reason)
			break
		}
	} else {
		log.Warnf("event is not sign type, actual type: %T", event)
	}
}

func (tss *TSSService) handleSigFinish(ctx context.Context, event interface{}) {
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

func (tss *TSSService) checkKeygen(ctx context.Context) {
	if tss.Party == nil {
		log.Debug("Party not init, start to setup")
		tss.setup()
		return
	}

	localPartySaveData, err := loadTSSData()
	if err == nil {
		if localPartySaveData.ECDSAPub != nil {
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
			if tss.round1P2pMessage != nil {
				log.Debug("Party set up timeout, send first round p2p message again")
				err = tss.libp2p.PublishMessage(ctx, *tss.round1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", tss.round1P2pMessage.RequestId)
				}
			}
		} else if rnd.RoundNumber() == 2 {
			if tss.round1P2pMessage != nil && tss.round1MessageSendTimes < 3 {
				tss.round1MessageSendTimes++
				log.Debugf("Reached round2, send first round p2p message the %d time", tss.round1MessageSendTimes)
				err = tss.libp2p.PublishMessage(ctx, *tss.round1P2pMessage)
				if err == nil {
					log.Debugf("Publish p2p message tssUpdateMessage again: RequestId=%s", tss.round1P2pMessage.RequestId)
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
}
