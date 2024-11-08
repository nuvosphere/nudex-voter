package tss

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	tssCommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type TSSService struct {
	isPrepared   atomic.Bool
	privateKey   *ecdsa.PrivateKey // submit
	localAddress common.Address    // submit = partyID

	p2p     p2p.P2PService
	state   *state.State
	manager *Scheduler[int32]
	cache   *cache.Cache

	layer2Listener *layer2.Layer2Listener
	dbm            *db.DatabaseManager
	taskReceive    chan any // task

	threshold          *atomic.Int64
	partners           []common.Address
	localParty         *keygen.LocalParty
	localPartySaveData *keygen.LocalPartySaveData
	partyIdMap         map[string]*tsslib.PartyID

	setupTime              time.Time
	keygenRound1P2pMessage *p2p.Message[types.TssMessage]
	round1MessageSendTimes int

	// tss keygen
	keyOutCh chan tsslib.Message
	keyEndCh chan *keygen.LocalPartySaveData

	// resharing channel
	reSharingOutCh chan tsslib.Message
	reSharingEndCh chan *keygen.LocalPartySaveData
	reLocalParty   *resharing.LocalParty

	// tss signature channel
	sigOutCh chan tsslib.Message
	sigEndCh chan *tssCommon.SignatureData

	// eventbus channel
	tssMsgCh       <-chan any
	partyAddOrRmCh <-chan any
	sigStartCh     <-chan any
	// sigReceiveCh   <-chan any
	sigFailChan    <-chan any
	sigTimeoutChan <-chan any

	sigMap                       map[string]map[int32]*signing.LocalParty
	sigRound1P2pMessageMap       map[string]*p2p.Message[types.TssMessage]
	sigRound1MessageSendTimesMap map[string]int
	sigTimeoutMap                map[string]time.Time

	rw sync.RWMutex

	once sync.Once
}

func (t *TSSService) IsCompleted(taskID int32) bool {
	_, ok := t.cache.Get(fmt.Sprintf("%d", taskID))
	return ok
}

func (t *TSSService) AddCompletedTask(taskID int32) {
	t.cache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (t *TSSService) LocalSubmitter() common.Address {
	return t.localAddress
}

func (t *TSSService) IsPrepared() bool {
	return t.isPrepared.Load()
}

func (t *TSSService) PostTask(task any) {
	t.taskReceive <- task
}

func (t *TSSService) taskLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info("TSS task loop stopped")
		case task := <-t.taskReceive:
			log.Info("TSS task receive", task)
		}
	}
}

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, state *state.State, layer2Listener *layer2.Layer2Listener) *TSSService {
	return &TSSService{
		isPrepared:     atomic.Bool{},
		privateKey:     config.AppConfig.L2PrivateKey,
		localAddress:   crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		p2p:            p,
		dbm:            dbm,
		state:          state,
		layer2Listener: layer2Listener,
		manager:        NewManager[int32](),
		cache:          cache.New(time.Minute*10, time.Minute),

		partyIdMap: make(map[string]*tsslib.PartyID),

		keyOutCh:       make(chan tsslib.Message),
		keyEndCh:       make(chan *keygen.LocalPartySaveData),
		reSharingOutCh: make(chan tsslib.Message),
		reSharingEndCh: make(chan *keygen.LocalPartySaveData),
		sigOutCh:       make(chan tsslib.Message),
		sigEndCh:       make(chan *tssCommon.SignatureData),

		sigMap:                       make(map[string]map[int32]*signing.LocalParty),
		sigRound1P2pMessageMap:       make(map[string]*p2p.Message[types.TssMessage]),
		sigRound1MessageSendTimesMap: make(map[string]int),
		sigTimeoutMap:                make(map[string]time.Time),
		taskReceive:                  make(chan any, 256),
	}
}

func (t *TSSService) Start(ctx context.Context) {
	t.waitForThreshold(ctx)

	is := t.IsGenesis()
	if is {
		t.Genesis(ctx)
	}

	t.initCheckGenKey(ctx)
	t.loop(ctx)
	t.check(ctx)

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	t.Stop()
}

func (t *TSSService) loop(ctx context.Context) {
	t.eventLoop(ctx)
}

func (t *TSSService) check(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info("Timeout checker stopping...")
				return
			case <-ticker.C:
				t.checkSigTimeouts()
			}
		}
	}()
}

func (t *TSSService) initCheckGenKey(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

	L:
		for {
			if t.localParty != nil && t.localPartySaveData != nil {
				break L
			}
			select {
			case <-ctx.Done():
				log.Info("Tss keygen checker stopping...")
				return
			case <-ticker.C:
				t.checkParty(ctx)
			}
		}
		t.tssLoop(ctx)
	}()
}

func (t *TSSService) genKeyLoop(ctx context.Context) {
	// generate tss key
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("t loop stopping...")
				return
			case msg := <-t.keyOutCh:
				log.Debugf("Received t keyOut event")

				err := t.handleTssKeyOut(ctx, msg)
				if err != nil {
					log.Warnf("handle t keyOut error, msg: %v, %v", msg, err)
				}
			case event := <-t.keyEndCh:
				log.Debugf("Received t keygenEnd event")

				err := t.handleTssKeyEnd(event)
				if err != nil {
					log.Warnf("handle t keygenEnd error, event: %v, %v", event, err)
				} else {
					t.isPrepared.Store(true)
					return
				}
			}
		}
	}()
}

func (t *TSSService) tssLoop(ctx context.Context) {
	// t re-share
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("t loop stopping...")
				return
			case msg := <-t.reSharingOutCh:
				log.Debugf("Received t re-sharing event")

				err := t.handleTssReSharingOut(ctx, msg)
				if err != nil {
					log.Warnf("handle t keyOut error, msg: %v, %v", msg, err)
				}
			case event := <-t.reSharingEndCh:
				log.Debugf("Received t re-sharing event")

				err := t.handleTssKeyEnd(event)
				if err != nil {
					log.Warnf("handle t re-sharing error, event: %v, %v", event, err)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("t loop stopping...")
				return
			case msg := <-t.sigOutCh:
				err := t.handleTssSigOut(ctx, msg)
				if err != nil {
					log.Warnf("handle t signature out error, msg: %v, %v", msg, err)
				}
			case sigFinish := <-t.sigEndCh:
				t.handleSigFinish(ctx, sigFinish)
			}
		}
	}()
}

func (t *TSSService) eventLoop(ctx context.Context) {
	t.p2p.Bind(p2p.MessageTypeTssMsg, state.EventTssMsg{})
	t.tssMsgCh = t.state.EventBus.Subscribe(state.EventTssMsg{})
	t.sigStartCh = t.state.EventBus.Subscribe(state.EventSigStart{})
	t.sigFailChan = t.state.EventBus.Subscribe(state.EventSigFailed{})
	t.sigTimeoutChan = t.state.EventBus.Subscribe(state.EventSigTimeout{})
	t.partyAddOrRmCh = t.state.EventBus.Subscribe(state.EventParticipantAddedOrRemoved{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Signer stopping...")
				return
			case event := <-t.tssMsgCh: // from p2p network
				log.Debugf("Received t msg event")

				e := event.(p2p.Message[json.RawMessage])

				err := t.handleTssMsg(e.DataType, convertMsgData(e))
				if err != nil {
					log.Warnf("handle t msg error, %v", err)
				}
			case <-t.partyAddOrRmCh: // from layer2 log scan
				log.Debugf("Received t add or remove participant event")

				err := t.startReSharing(t.state.TssState.Participants, config.AppConfig.TssThreshold) // todo
				if err != nil {
					log.Warnf("handle t add or remove participant event error, %v", err)
				}
			case event := <-t.sigStartCh:
				log.Debugf("Received sigStart event: %v", event)
				t.handleSigStart(ctx, event)
			case sigFail := <-t.sigFailChan:
				t.handleSigFailed(ctx, sigFail, "failed")
			case sigTimeout := <-t.sigTimeoutChan:
				t.handleSigFailed(ctx, sigTimeout, "timeout") // from self ??? todo
			}
		}
	}()
}

func (t *TSSService) Stop() {
	t.once.Do(func() {
		// close(t.keyOutCh)
		// close(t.keyEndCh)
		close(t.reSharingEndCh)
		close(t.reSharingOutCh)
		close(t.sigOutCh)
		close(t.sigEndCh)
	})
}

func (t *TSSService) cleanAllSigInfo() {
	t.sigMap = make(map[string]map[int32]*signing.LocalParty)
	t.sigRound1P2pMessageMap = make(map[string]*p2p.Message[types.TssMessage])
	t.sigRound1MessageSendTimesMap = make(map[string]int)
	t.sigTimeoutMap = make(map[string]time.Time)
}

func (t *TSSService) waitForThreshold(ctx context.Context) {
	count := t.p2p.OnlinePeerCount()
L:
	for {
		select {
		case <-ctx.Done():
			log.Info("waitForThreshold context done")
		default:
			if int64(count) >= t.threshold.Load() {
				break L
			}
			time.Sleep(time.Second)
		}
	}
}
