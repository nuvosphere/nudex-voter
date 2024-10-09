package tss

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewTssService(libp2p *p2p.LibP2PService, state *state.State) *TSSService {
	return &TSSService{
		privateKey: config.AppConfig.L2PrivateKey,
		address:    crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		libp2p:     libp2p,
		state:      state,

		keygenReqCh:     make(chan interface{}, 10),
		keygenReceiveCh: make(chan interface{}, 10),

		sigStartCh:   make(chan interface{}, 256),
		sigReceiveCh: make(chan interface{}, 1024),

		sigFailChan:    make(chan interface{}, 10),
		sigFinishChan:  make(chan interface{}, 10),
		sigTimeoutChan: make(chan interface{}, 10),

		sigMap:        make(map[string]map[string]interface{}),
		sigTimeoutMap: make(map[string]time.Time),
	}
}

func (tss *TSSService) Start(ctx context.Context) {
	tss.keygen(ctx)

	go tss.signLoop(ctx)

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	tss.Stop()
}

func (tss *TSSService) keygen(ctx context.Context) {
	tss.sigMu.Lock()
	defer tss.sigMu.Unlock()

	requestId := fmt.Sprintf("KEYGEN:%s", crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey).Hex())
	keygenReqMessage := types.KeygenReqMessage{
		RequestId:    requestId,
		VoterAddress: tss.address.Hex(),
		CreateTime:   time.Now().Unix(),
		PublicKeys:   PublicKeysToHex(config.AppConfig.TssPublicKeys),
		Threshold:    config.AppConfig.TssThreshold,
	}

	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeKeygenReq,
		RequestId:   requestId,
		DataType:    p2p.DataTypeKeygenReq,
		Data:        keygenReqMessage,
	}

	err := tss.libp2p.PublishMessage(ctx, p2pMsg)
	if err != nil {
		log.Errorf("Error publishing keygenReqMessage message: %v", err)
		return
	}
	log.Debugf("Publish p2p message keygenReqMessage: RequestId=%s, Key Length=%d, Threshold=%d, Keys=%v",
		requestId, len(keygenReqMessage.PublicKeys), keygenReqMessage.Threshold,
		keygenReqMessage.PublicKeys)

	tss.sigMap[requestId] = make(map[string]interface{})
	tss.sigMap[requestId][tss.address.Hex()] = true
	timeoutDuration := config.AppConfig.TssSigTimeout
	tss.sigTimeoutMap[requestId] = time.Now().Add(timeoutDuration)
	log.Infof("KeygenReq broadcast ok, request id: %s", requestId)
}

func (tss *TSSService) setup(ctx context.Context) {

}

func (tss *TSSService) signLoop(ctx context.Context) {
	tss.state.EventBus.Subscribe(state.KeygenStart, tss.keygenReqCh)
	tss.state.EventBus.Subscribe(state.KeygenReceive, tss.keygenReceiveCh)

	tss.state.EventBus.Subscribe(state.SigStart, tss.sigStartCh)
	tss.state.EventBus.Subscribe(state.SigReceive, tss.sigReceiveCh)

	tss.state.EventBus.Subscribe(state.SigFailed, tss.sigFailChan)
	tss.state.EventBus.Subscribe(state.SigFinish, tss.sigFinishChan)
	tss.state.EventBus.Subscribe(state.SigTimeout, tss.sigTimeoutChan)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Info("Signer stoping...")
			return
		case event := <-tss.keygenReqCh:
			log.Debugf("Received keygenReq event: %v", event)
			tss.handleKeygenReq(ctx, event)
		case event := <-tss.keygenReceiveCh:
			log.Debugf("Received keygenReveive event: %v", event)
			tss.handleKeygenReceive(ctx, event)
		case event := <-tss.sigStartCh:
			log.Debugf("Received sigStart event: %v", event)
			tss.handleSigStart(ctx, event)
		case event := <-tss.sigReceiveCh:
			log.Debugf("Received sigReceive event: %v", event)
			tss.handleSigReceive(ctx, event)
		case sigFail := <-tss.sigFailChan:
			tss.handleSigFailed(ctx, sigFail, "failed")
		case sigTimeout := <-tss.sigTimeoutChan:
			tss.handleSigFailed(ctx, sigTimeout, "timeout")
		case sigFinish := <-tss.sigFinishChan:
			tss.handleSigFinish(ctx, sigFinish)
		}
	}

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info("Timeout checker stopping...")
				return
			case <-ticker.C:
				tss.checkTimeouts()
			}
		}
	}()
}

func (tss *TSSService) Stop() {
	tss.once.Do(func() {
		close(tss.keygenReqCh)
		close(tss.keygenReceiveCh)

		close(tss.sigStartCh)
		close(tss.sigReceiveCh)

		close(tss.sigFailChan)
		close(tss.sigFinishChan)
		close(tss.sigTimeoutChan)
	})
}
