package tss

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"sync"
	"time"

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
	keygenRound1P2pMessage *p2p.Message[types.TssMessage]
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

//type MsgSign struct {
//	RequestId    string `json:"request_id"`
//	VoterAddress string `json:"voter_address"`
//	IsProposer   bool   `json:"is_proposer"`
//	CreateTime   int64  `json:"create_time"`
//}
//
//type MsgSignCreateWalletMessage struct {
//	MsgSign
//	Task types.CreateWalletTask `json:"task"`
//}
//
//func (m *MsgSignCreateWalletMessage) Unmarshal(msg p2p.Message[json.RawMessage]) error {
//	if msg.DataType == DataTypeSignCreateWallet {
//		return json.Unmarshal(msg.Data, m)
//	}
//
//	return ErrDataType
//}
//
//type TssMessage struct {
//	FromPartyId  string   `json:"from_party_id"`
//	ToPartyIds   []string `json:"to_party_ids"`
//	IsBroadcast  bool     `json:"is_broadcast"`
//	MsgWireBytes []byte   `json:"msg_wire_bytes"`
//}
//
//func (t *TssMessage) Unmarshal(msg p2p.Message[json.RawMessage]) error {
//	switch msg.DataType {
//	case DataTypeTssKeygenMsg, DataTypeTssSignMsg, DataTypeTssReSharingMsg:
//		return json.Unmarshal(msg.Data, t)
//	}
//
//	return ErrDataType
//}

var ErrDataType = errors.New("error data type")

// convertMsgData converts the message data to the corresponding struct.
func convertMsgData(msg p2p.Message[json.RawMessage]) any {
	var rawData any

	switch msg.DataType {
	case DataTypeTssKeygenMsg, DataTypeTssSignMsg, DataTypeTssReSharingMsg:
		rawData = &types.TssMessage{}
	case DataTypeSignCreateWallet:
		rawData = &types.SignMessage{}
	}

	if rawData != nil {
		err := json.Unmarshal(msg.Data, &rawData)
		utils.Assert(err)
	} else {
		utils.Assert(ErrDataType)
	}

	return rawData
}
