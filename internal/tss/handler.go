package tss

import (
	"context"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"time"
)

func (tss *TSSService) handleSigStart(ctx context.Context, event interface{}) {
	switch e := event.(type) {
	case types.MsgSignKeyPrepareMessage:
		log.Debugf("Event handleSigStart is of type MsgSignKeyPrepareMessage, request id %s", e.RequestId)
		if err := tss.handleSigStartKeyPrepare(ctx, e); err != nil {
			log.Errorf("Error handleSigStart MsgSignKeyPrepareMessage, %v", err)
			tss.state.EventBus.Publish(state.SigFailed, e)
		}
	default:
		log.Debug("Unknown event handleSigStart type")
	}
}

func (tss *TSSService) handleSigReceive(ctx context.Context, event interface{}) {
}

func (tss *TSSService) handleSigFailed(ctx context.Context, event interface{}, reason string) {
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

func (tss *TSSService) checkKeygen() {
	tss.sigMu.Lock()
	tss.sigMu.Unlock()

	if tss.party == nil {
		log.Debug("Party not init, start to setup")
		tss.setup()
		return
	}

	_, err := loadTSSData()
	if err == nil {
		return
	}

	if tss.setupTime.IsZero() {
		tss.setup()
		return
	}

	if time.Now().After(tss.setupTime.Add(config.AppConfig.TssSigTimeout)) {
		log.Debug("Party set up timeout, start to setup again")
		tss.setup()
		return
	}
}

func (tss *TSSService) sigExists(requestId string) (map[string]interface{}, bool) {
	tss.sigMu.RLock()
	defer tss.sigMu.RUnlock()
	data, ok := tss.sigMap[requestId]
	return data, ok
}

func (tss *TSSService) removeSigMap(requestId string, reportTimeout bool) {
	tss.sigMu.Lock()
	defer tss.sigMu.Unlock()
	if reportTimeout {
		if voteMap, ok := tss.sigMap[requestId]; ok {
			if voteMsg, ok := voteMap[tss.address.Hex()]; ok {
				log.Debugf("Report timeout when remove sig map, found msg, request id %s, proposer %s",
					requestId, tss.address.Hex())
				tss.state.EventBus.Publish(state.SigTimeout, voteMsg)
			}
		}
	}
	delete(tss.sigMap, requestId)
	delete(tss.sigTimeoutMap, requestId)
}
