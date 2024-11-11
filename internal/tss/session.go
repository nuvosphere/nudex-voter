package tss

import (
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

type Session[T any] interface {
	helper.Transporter
	Releaser
	Name
	Type
	SessionID
	GroupID
	TaskID[T]
	Proposer
	Threshold
	PartyID
	Equal
	Included
	Run
}

type Run interface {
	Run()
}

type Name interface {
	Name() string
}

type Type interface {
	Type() string
}

type SessionID interface {
	SessionID() helper.SessionID
}

type GroupID interface {
	GroupID() helper.GroupID
}

type TaskID[T any] interface {
	TaskID() T
}

type Proposer interface {
	Proposer() common.Address
}

type PartyID interface {
	PartyID(id string) *tss.PartyID
}

type Included interface {
	Included(ids []string) bool
}

type Equal interface {
	Equal(id string) bool
}

type Threshold interface {
	Threshold() int
}
