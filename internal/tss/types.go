package tss

import (
	"crypto/ecdsa"
	"encoding/json"
	tssCommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"sync"
	"time"
)

type TSSService struct {
	privateKey *ecdsa.PrivateKey
	Address    common.Address

	p2p   p2p.P2PService
	state *state.State
	dbm   *db.DatabaseManager

	addressList        []common.Address
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
	tssMsgCh       <-chan any
	partyAddOrRmCh <-chan any
	sigStartCh     <-chan any
	sigReceiveCh   <-chan any
	sigFailChan    <-chan any
	sigTimeoutChan <-chan any

	sigMap                       map[string]map[int32]*signing.LocalParty
	sigRound1P2pMessageMap       map[string]*p2p.Message
	sigRound1MessageSendTimesMap map[string]int
	sigTimeoutMap                map[string]time.Time

	rw sync.RWMutex

	once sync.Once
}

const (
	DataTypeTssKeygenMsg     = "TssKeygenMsg"
	DataTypeTssSignMsg       = "TssSignMsg"
	DataTypeTssReSharingMsg  = "TssReSharingMsg"
	DataTypeSignCreateWallet = "SignCreateWallet"
)

// convertMsgData converts the message data to the corresponding struct
// TODO: use reflector to optimize this function
func convertMsgData(msg p2p.Message) any {
	switch msg.DataType {
	case DataTypeTssKeygenMsg, DataTypeTssSignMsg, DataTypeTssReSharingMsg:
		jsonBytes, _ := json.Marshal(msg.Data)
		var rawData types.TssMessage
		err := json.Unmarshal(jsonBytes, &rawData)
		utils.Assert(err)
		return rawData
	case DataTypeSignCreateWallet:
		jsonBytes, _ := json.Marshal(msg.Data)
		var rawData types.MsgSignCreateWalletMessage
		err := json.Unmarshal(jsonBytes, &rawData)
		utils.Assert(err)
		return rawData
	}
	return msg.Data
}
