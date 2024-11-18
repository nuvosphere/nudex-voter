package helper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	SessionID = common.Hash
	GroupID   = common.Hash
)

var (
	ZeroSessionID    SessionID
	SenateSessionID  = crypto.Keccak256Hash([]byte("The voter senate session，one and only one"))
	SenateProposal   = SenateSessionID.Big()
	SenateProposalID = SenateSessionID.Big().Int64()
)

type BaseMessage[T, M any] struct {
	GroupID    GroupID        `json:"group_id,omitempty"`
	SessionID  SessionID      `json:"session_id,omitempty"`
	Proposer   common.Address `json:"proposer,omitempty"`    // current submitter
	ProposalID T              `json:"proposal_id,omitempty"` // msg id
	Proposal   M              `json:"proposal"`
}

type Session[T, M any] struct {
	Group
	SessionID  SessionID      `json:"sessionID,omitempty"`
	Proposer   common.Address `json:"proposer,omitempty"`    // current submitter
	Signer     common.Address `json:"signer,omitempty"`      // current signer
	ProposalID T              `json:"proposal_id,omitempty"` // msg id
	Proposal   M              `json:"proposal,omitempty"`
	Threshold  int            `json:"threshold"`
}

type Group struct {
	GroupID     GroupID          `json:"groupID,omitempty"`
	AllPartners []common.Address `json:"allPartners"` // all submitter
}
