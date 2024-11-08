package helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type TaskID = common.Hash

type SessionID = common.Hash

type GroupID = common.Address

var (
	SenateSessionID = crypto.Keccak256Hash([]byte("The voter senate session，one and only one"))
	SenateGroupID   = common.BytesToAddress(SenateSessionID.Bytes())
)

type BaseMessage[T, M any] struct {
	GroupID   GroupID        `json:"groupID,omitempty"`
	SessionID SessionID      `json:"sessionID,omitempty"`
	Sponsor   common.Address `json:"sponsor,omitempty"` // current submitter
	TaskID    T              `json:"taskID,omitempty"`  // msg id
	Msg       M              `json:"msg"`
}

type Session[T, M any] struct {
	GroupID   GroupID        `json:"groupID,omitempty"`
	SessionID SessionID      `json:"sessionID,omitempty"`
	Sponsor   common.Address `json:"sponsor,omitempty"` // current submitter
	TaskID    T              `json:"taskID,omitempty"`  // msg id
	Msg       M              `json:"msg"`
	Threshold int            `json:"threshold"`
	// Group
	allPartners []common.Address // all submitter
	// actualPartners []common.Address // todo
}

type Group struct {
	GroupID     GroupID
	AllPartners []common.Address // all submitter
}
