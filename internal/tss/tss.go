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

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, state *state.State) *TSSService {
	return &TSSService{
		privateKey: config.AppConfig.L2PrivateKey,
		Address:    crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		p2p:        p,
		dbm:        dbm,
		state:      state,

		partyIdMap: make(map[string]*tsslib.PartyID),

		keyOutCh:       make(chan tsslib.Message),
		keyEndCh:       make(chan *keygen.LocalPartySaveData),
		reSharingOutCh: make(chan tsslib.Message),
		reSharingEndCh: make(chan *keygen.LocalPartySaveData),
		sigOutCh:       make(chan tsslib.Message),
		sigEndCh:       make(chan *common.SignatureData),

		//tssMsgCh:       make(chan interface{}, 10),
		//sigStartCh:     make(chan interface{}, 256),
		//sigReceiveCh:   make(chan interface{}, 1024),
		//sigFailChan:    make(chan interface{}, 10),
		//sigTimeoutChan: make(chan interface{}, 10),

		sigMap:                       make(map[string]map[int32]*signing.LocalParty),
		sigRound1P2pMessageMap:       make(map[string]*p2p.Message),
		sigRound1MessageSendTimesMap: make(map[string]int),
		sigTimeoutMap:                make(map[string]time.Time),
	}
}

func (tss *TSSService) Start(ctx context.Context) {
	tss.loop(ctx)
	tss.check(ctx)

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	tss.Stop()
}

func (tss *TSSService) loop(ctx context.Context) {
	tss.eventLoop(ctx)
	tss.tssLoop(ctx)
}

func (tss *TSSService) check(ctx context.Context) {
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
				tss.checkParty(ctx)
			}
		}
	}()
}

func (tss *TSSService) tssLoop(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("tss loop stopping...")
				return
			case msg := <-tss.keyOutCh:
				log.Debugf("Received tss keyOut event")
				err := tss.handleTssKeyOut(ctx, msg)
				if err != nil {
					log.Warnf("handle tss keyOut error, msg: %v, %v", msg, err)
				}
			case event := <-tss.keyEndCh:
				log.Debugf("Received tss keygenEnd event")
				err := tss.handleTssKeyEnd(event)
				if err != nil {
					log.Warnf("handle tss keygenEnd error, event: %v, %v", event, err)
				}
			case msg := <-tss.reSharingOutCh:
				log.Debugf("Received tss re-sharing event")
				err := tss.handleTssReSharingOut(ctx, msg)
				if err != nil {
					log.Warnf("handle tss keyOut error, msg: %v, %v", msg, err)
				}
			case event := <-tss.reSharingEndCh:
				log.Debugf("Received tss re-sharing event")
				err := tss.handleTssKeyEnd(event)
				if err != nil {
					log.Warnf("handle tss re-sharing error, event: %v, %v", event, err)
				}
			case msg := <-tss.sigOutCh:
				err := tss.handleTssSigOut(ctx, msg)
				if err != nil {
					log.Warnf("handle tss signature out error, msg: %v, %v", msg, err)
				}
			case sigFinish := <-tss.sigEndCh:
				tss.handleSigFinish(ctx, sigFinish)
			}
		}
	}()
}

func (tss *TSSService) eventLoop(ctx context.Context) {
	tss.tssMsgCh = tss.state.EventBus.Subscribe(state.EventTssMsg{})
	tss.sigStartCh = tss.state.EventBus.Subscribe(state.EventSigStart{})
	tss.sigReceiveCh = tss.state.EventBus.Subscribe(state.EventSigReceive{})
	tss.sigFailChan = tss.state.EventBus.Subscribe(state.EventSigFailed{})
	tss.sigTimeoutChan = tss.state.EventBus.Subscribe(state.EventSigTimeout{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-tss.tssMsgCh:
				log.Debugf("Received tssUpdated event")
				err := tss.handleTssMsg(event)
				if err != nil {
					log.Warnf("handle tssUpdate error, %v", err)
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
			}
		}
	}()
}

func (tss *TSSService) Stop() {
	tss.once.Do(func() {
		close(tss.keyOutCh)
		close(tss.keyEndCh)
		close(tss.reSharingEndCh)
		close(tss.reSharingOutCh)
		close(tss.sigOutCh)
		close(tss.sigEndCh)
	})
}

func (tss *TSSService) CleanAllSigInfo() {
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
