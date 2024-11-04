package tss

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"sync"
	"time"

	tssCommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

type TSSService struct {
	privateKey *ecdsa.PrivateKey
	Address    common.Address

	p2p   p2p.P2PService
	state *state.State

	layer2Listener *layer2.Layer2Listener
	dbm            *db.DatabaseManager

	addressList        []common.Address
	LocalParty         *keygen.LocalParty
	LocalPartySaveData *keygen.LocalPartySaveData
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

const (
	DataTypeTssKeygenMsg     = "TssKeygenMsg"
	DataTypeTssSignMsg       = "TssSignMsg"
	DataTypeTssReSharingMsg  = "TssReSharingMsg"
	DataTypeSignCreateWallet = "SignCreateWallet"
	DataTypeSignDeposit      = "SignDeposit"
	DataTypeSignWithdrawal   = "SignWithdrawal"
)

// convertMsgData converts the message data to the corresponding struct.
func convertMsgData(msg p2p.Message[json.RawMessage]) any {
	switch msg.DataType {
	case DataTypeTssKeygenMsg, DataTypeTssSignMsg, DataTypeTssReSharingMsg:
		return unmarshal[types.TssMessage](msg.Data)
	case DataTypeSignCreateWallet:
		return unmarshal[types.SignMessage](msg.Data)
	}

	return unmarshal[any](msg.Data)
}

func unmarshal[T any](data json.RawMessage) T {
	var obj T

	err := json.Unmarshal(data, &obj)
	if err != nil || data == nil {
		panic(fmt.Errorf("unmarshal data:%v, error: %w", data, err))
	}

	return obj
}
