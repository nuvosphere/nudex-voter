package tss

import (
	"encoding/json"
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
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

// convertMsgData converts the message data to the corresponding struct.
func convertMsgData(msg p2p.Message[json.RawMessage]) any {
	switch msg.DataType {
	case DataTypeTssKeygenMsg, DataTypeTssReSharingMsg, DataTypeTssSignMsg:
		return unmarshal[types.TssMessage](msg.Data)
	case GenKeySessionType, SignTaskSessionType, SignBatchTaskSessionType, ReShareGroupSessionType, TxSignatureSessionType:
		return unmarshal[SessionMessage[ProposalID, Proposal]](msg.Data)
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

type BatchData struct {
	Ids []uint64 `json:"ids"`
}

func (b *BatchData) Bytes() []byte {
	data, err := json.Marshal(b)
	utils.Assert(err)
	return data
}

func (b *BatchData) FromBytes(data []byte) {
	*b = unmarshal[BatchData](data)
}
