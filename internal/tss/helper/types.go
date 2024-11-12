package helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SessionID = common.Hash

type GroupID = common.Address

var (
	SenateSessionID = crypto.Keccak256Hash([]byte("The voter senate session，one and only one"))
	SenateGroupID   = common.BytesToAddress(SenateSessionID.Bytes())
	SenateTaskID    = SenateSessionID.Big().Int64()
)

type BaseMessage[T, M any] struct {
	GroupID   GroupID        `json:"groupID,omitempty"`
	SessionID SessionID      `json:"sessionID,omitempty"`
	Proposer  common.Address `json:"proposer,omitempty"` // current submitter
	TaskID    T              `json:"taskID,omitempty"`   // msg id
	Msg       M              `json:"msg"`
}

type Session[T, M any] struct {
	Group
	SessionID SessionID      `json:"sessionID,omitempty"`
	Proposer  common.Address `json:"proposer,omitempty"` // current submitter
	TaskID    T              `json:"taskID,omitempty"`   // msg id
	Msg       M              `json:"msg"`
	Threshold int            `json:"threshold"`
}

type Group struct {
	GroupID     GroupID          `json:"groupID,omitempty"`
	AllPartners []common.Address `json:"allPartners"` // all submitter
}
