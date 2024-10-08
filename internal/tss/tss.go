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

		sigStartCh:   make(chan interface{}, 256),
		sigReceiveCh: make(chan interface{}, 1024),

		sigMap:        make(map[string]map[string]interface{}),
		sigTimeoutMap: make(map[string]time.Time),
	}
}

func (tss *TSSService) Start(ctx context.Context) {
	keygenPrepareMessage := types.KeyGenPrepareMessage{
		PublicKeys:  PublicKeysToHex(config.AppConfig.TssPublicKeys),
		FromAddress: tss.address.Hex(),
		Threshold:   config.AppConfig.TssThreshold,
		Timestamp:   time.Now().Unix(),
	}

	requestId := fmt.Sprintf("KEYGEN:%s", crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey).Hex())
	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeSigReq,
		RequestId:   requestId,
		DataType:    p2p.DataTypeKeygenPrepare,
		Data:        keygenPrepareMessage,
	}

	tss.libp2p.PublishMessage(ctx, p2pMsg)

	log.Debugf("Publish p2p message keygenPrepare: RequestId=%s, Key Length=%d, Threshold=%d, Keys=%v",
		requestId, len(keygenPrepareMessage.PublicKeys), keygenPrepareMessage.Threshold,
		keygenPrepareMessage.PublicKeys)

	tss.state.EventBus.Subscribe(state.SigStart, tss.sigStartCh)
	tss.state.EventBus.Subscribe(state.SigReceive, tss.sigReceiveCh)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stoping...")
				return
			case event := <-tss.sigStartCh:
				log.Debugf("Received sigStart event: %v", event)
				tss.handleSigStart(ctx, event)
			case event := <-tss.sigReceiveCh:
				log.Debugf("Received sigReceive event: %v", event)
				//tss.handleSigReceive(ctx, event)
			}
		}
	}()

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
