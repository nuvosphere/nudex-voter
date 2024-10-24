package tss

import (
	"crypto/ecdsa"
	tssCommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"sync"
	"time"
)

type TSSService struct {
	privateKey *ecdsa.PrivateKey
	Address    common.Address

	libp2p *p2p.LibP2PService
	state  *state.State

	Party              *keygen.LocalParty
	LocalPartySaveData *keygen.LocalPartySaveData
	partyIdMap         map[string]*tsslib.PartyID

	setupTime              time.Time
	round1P2pMessage       *p2p.Message
	round1MessageSendTimes int

	tssUpdateCh chan interface{}

	keyOutCh    chan tsslib.Message
	keygenEndCh chan *keygen.LocalPartySaveData
	signEndCh   chan *tssCommon.SignatureData

	sigStartCh   chan interface{}
	sigReceiveCh chan interface{}

	sigFailChan    chan interface{}
	sigFinishChan  chan interface{}
	sigTimeoutChan chan interface{}

	sigPartyMap   map[string]*signing.LocalParty
	sigTimeoutMap map[string]time.Time
	sigMu         sync.RWMutex

	once sync.Once
}
