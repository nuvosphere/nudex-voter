package tss

import (
	"context"
	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
)

func NewTssService(libp2p *p2p.LibP2PService, dbm *db.DatabaseManager, state *state.State) *TSSService {
	return &TSSService{
		privateKey: config.AppConfig.L2PrivateKey,
		Address:    crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		libp2p:     libp2p,
		dbm:        dbm,
		state:      state,

		tssUpdateCh: make(chan interface{}, 10),

		keyOutCh:    make(chan tsslib.Message),
		keygenEndCh: make(chan *keygen.LocalPartySaveData),

		sigFinishChan: make(chan *common.SignatureData),

		sigStartCh:   make(chan interface{}, 256),
		sigReceiveCh: make(chan interface{}, 1024),

		sigFailChan:    make(chan interface{}, 10),
		sigTimeoutChan: make(chan interface{}, 10),

		sigMap:                       make(map[string]map[int32]*signing.LocalParty),
		sigRound1P2pMessageMap:       make(map[string]*p2p.Message),
		sigRound1MessageSendTimesMap: make(map[string]int),
		sigTimeoutMap:                make(map[string]time.Time),
	}
}

func (tss *TSSService) Start(ctx context.Context) {
	go tss.signLoop(ctx)

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	tss.Stop()
}

func (tss *TSSService) signLoop(ctx context.Context) {
	tss.state.EventBus.Subscribe(state.TssUpdate, tss.tssUpdateCh)

	tss.state.EventBus.Subscribe(state.SigStart, tss.sigStartCh)
	tss.state.EventBus.Subscribe(state.SigReceive, tss.sigReceiveCh)

	tss.state.EventBus.Subscribe(state.SigFailed, tss.sigFailChan)
	tss.state.EventBus.Subscribe(state.SigTimeout, tss.sigTimeoutChan)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-tss.tssUpdateCh:
				log.Debugf("Received tssUpdated event")
				err := tss.handleTssUpdate(event)
				if err != nil {
					log.Warnf("handle tssUpdate error, %v", err)
				}
			case event := <-tss.keyOutCh:
				log.Debugf("Received tss keyOut event")
				err := tss.handleTssKeyOut(ctx, event)
				if err != nil {
					log.Warnf("handle tss keyOut error, event: %v, %v", event, err)
				}
			case event := <-tss.keygenEndCh:
				log.Debugf("Received tss keygenEnd event")
				err := tss.handleTssKeyEnd(event)
				if err != nil {
					log.Warnf("handle tss keygenEnd error, event: %v, %v", event, err)
				}
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

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info("Tss keygen checker stopping...")
				return
			case <-ticker.C:
				tss.check(ctx)
			}
		}
	}()
}

func (tss *TSSService) Stop() {
	tss.once.Do(func() {
		close(tss.tssUpdateCh)

		close(tss.keyOutCh)
		close(tss.keygenEndCh)
		close(tss.sigFinishChan)

		close(tss.sigStartCh)
		close(tss.sigReceiveCh)

		close(tss.sigFailChan)
		close(tss.sigTimeoutChan)
	})
}

func (tss *TSSService) CleanAll() {
	for k := range tss.sigMap {
		delete(tss.sigMap, k)
	}
	for k := range tss.sigTimeoutMap {
		delete(tss.sigTimeoutMap, k)
	}
	for k := range tss.sigRound1P2pMessageMap {
		delete(tss.sigRound1P2pMessageMap, k)
	}
	for k := range tss.sigRound1MessageSendTimesMap {
		delete(tss.sigRound1MessageSendTimesMap, k)
	}
}
