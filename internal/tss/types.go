package tss

import (
	"crypto/ecdsa"
	common2 "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
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
	dbm    *db.DatabaseManager

	Party              *keygen.LocalParty
	LocalPartySaveData *keygen.LocalPartySaveData
	partyIdMap         map[string]*tsslib.PartyID

	setupTime              time.Time
	keygenRound1P2pMessage *p2p.Message
	round1MessageSendTimes int

	tssUpdateCh chan interface{}

	keyOutCh    chan tsslib.Message
	keygenEndCh chan *keygen.LocalPartySaveData

	sigFinishChan chan *common2.SignatureData

	sigStartCh   chan interface{}
	sigReceiveCh chan interface{}

	sigFailChan    chan interface{}
	sigTimeoutChan chan interface{}

	sigMap                       map[string]map[int32]*signing.LocalParty
	sigRound1P2pMessageMap       map[string]*p2p.Message
	sigRound1MessageSendTimesMap map[string]int
	sigTimeoutMap                map[string]time.Time

	sigMu sync.RWMutex

	once sync.Once
}
