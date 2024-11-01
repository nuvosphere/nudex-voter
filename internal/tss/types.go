package tss

import (
	"crypto/ecdsa"
	tssCommon "github.com/bnb-chain/tss-lib/v2/common"
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

	p2p   p2p.P2PService
	state *state.State
	dbm   *db.DatabaseManager

	LocalParty         *keygen.LocalParty
	LocalPartySaveData *keygen.LocalPartySaveData
	partyIdMap         map[string]*tsslib.PartyID

	setupTime              time.Time
	keygenRound1P2pMessage *p2p.Message
	round1MessageSendTimes int

	// tss keygen
	keyOutCh chan tsslib.Message
	keyEndCh chan *keygen.LocalPartySaveData

	// resharing channel
	reSharingOutCh chan tsslib.Message
	reSharingEndCh chan *keygen.LocalPartySaveData

	// tss signature channel
	sigOutCh chan tsslib.Message
	sigEndCh chan *tssCommon.SignatureData

	// eventbus channel
	tssMsgCh       chan interface{}
	sigStartCh     chan interface{}
	sigReceiveCh   chan interface{}
	sigFailChan    chan interface{}
	sigTimeoutChan chan interface{}

	sigMap                       map[string]map[int32]*signing.LocalParty
	sigRound1P2pMessageMap       map[string]*p2p.Message
	sigRound1MessageSendTimesMap map[string]int
	sigTimeoutMap                map[string]time.Time

	rw sync.RWMutex

	once sync.Once
}
