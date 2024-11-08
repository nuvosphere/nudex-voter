package helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type TaskID = common.Hash

type SessionID = common.Hash

type GroupID = common.Address

var SenateSessionID = crypto.Keccak256Hash([]byte("The voter senate session，one and only one"))

type BaseMessage struct {
	GroupID   GroupID        `json:"groupID,omitempty"`
	SessionID SessionID      `json:"sessionID,omitempty"`
	Sponsor   common.Address `json:"sponsor,omitempty"` // current submitter
	TaskID    TaskID         `json:"taskID,omitempty"`  // msg id
	Msg       []byte         `json:"msg"`
}

type Session struct {
	BaseMessage
	Group
	threshold int
	// allPartners    []common.Address // all submitter
	// actualPartners []common.Address // todo
}

type Group struct {
	GroupID     GroupID
	AllPartners []common.Address // all submitter
}
