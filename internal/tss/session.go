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
	Sponsor
	Threshold
	PartyID
	Party
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

type Sponsor interface {
	Sponsor() common.Address
}

type PartyID interface {
	PartyID(id string) *tss.PartyID
}

type Party interface {
	Party() tss.Party
}

type Threshold interface {
	Threshold() int
}
