package tss

import (
	"encoding/json"
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

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
	case DataTypeTssKeygenMsg, DataTypeTssReSharingMsg, DataTypeTssSignMsg:
		return unmarshal[types.TssMessage](msg.Data)
	case GenKeySessionType, SignTaskSessionType, ReShareGroupSessionType:
		return unmarshal[SessionMessage[ProposalID, Proposal]](msg.Data)
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
