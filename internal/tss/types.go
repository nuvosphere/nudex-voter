package tss

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"sync"
	"time"
)

type TSSService struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address

	libp2p *p2p.LibP2PService
	state  *state.State

	keygenReqCh     chan interface{}
	keygenReceiveCh chan interface{}

	sigStartCh   chan interface{}
	sigReceiveCh chan interface{}

	sigFailChan    chan interface{}
	sigFinishChan  chan interface{}
	sigTimeoutChan chan interface{}

	// [request_id][vote_address]MsgSign
	sigMap        map[string]map[string]interface{}
	sigTimeoutMap map[string]time.Time
	sigMu         sync.RWMutex

	once sync.Once
}
