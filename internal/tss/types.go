package tss

import (
	"encoding/json"

	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

const (
	DataTypeTssKeygenMsg    = "TssKeygenMsg"
	DataTypeTssSignMsg      = "TssSignMsg"
	DataTypeTssReSharingMsg = "TssReSharingMsg"
	DataTypeSignDeposit     = "SignDeposit"
	DataTypeSignWithdrawal  = "SignWithdrawal"
)

const (
	GenKeySessionType        = "GenerateKeySession"
	ReShareGroupSessionType  = "ReShareGroupSession"
	SignTaskSessionType      = "SignTaskSession"
	SignBatchTaskSessionType = "SignBatchTaskSessionType"
	TxSignatureSessionType   = "TxSignatureSession"
)

// ConvertP2PMsgData converts the message data to the corresponding struct.
func ConvertP2PMsgData(msg p2p.Message[json.RawMessage]) any {
	switch msg.DataType {
	case DataTypeTssKeygenMsg, DataTypeTssReSharingMsg, DataTypeTssSignMsg:
		return types.UnmarshalJson[types.TssMessage](msg.Data)
	case GenKeySessionType, SignTaskSessionType, SignBatchTaskSessionType, ReShareGroupSessionType, TxSignatureSessionType:
		return types.UnmarshalJson[SessionMessage[ProposalID, Proposal]](msg.Data)
	}

	return types.UnmarshalJson[any](msg.Data)
}
