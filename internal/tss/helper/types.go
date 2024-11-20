package helper

import (
	"crypto/elliptic"

	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	SessionID = common.Hash
	GroupID   = common.Hash
)

var (
	ZeroSessionID           SessionID
	SenateSessionID         = crypto.Keccak256Hash([]byte("The voter senate session，one and only one"))
	SenateProposal          = SenateSessionID.Big()
	senateProposalID        = SenateSessionID.Big().Int64()
	SenateProposalIDOfECDSA = senateProposalID - 1
	SenateProposalIDOfEDDSA = senateProposalID - 2
	SenateSessionIDOfECDSA  = crypto.Keccak256Hash([]byte("ECDSA:The voter senate session，one and only one"))
	SenateSessionIDOfEDDSA  = crypto.Keccak256Hash([]byte("EDDSA:The voter senate session，one and only one"))
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
	EC          CurveType        `json:"ec,omitempty"`
	GroupID     GroupID          `json:"group_id,omitempty"`
	AllPartners []common.Address `json:"all_partners,omitempty"` // all submitter
}

type CurveType int

func (e *CurveType) EC() elliptic.Curve {
	switch *e {
	case EDDSA:
		return tss.Edwards()
	case ECDSA:
		return tss.S256()
	default:
		panic("implement me")
	}
}

func (e *CurveType) CurveName() string {
	switch *e {
	case EDDSA:
		return "ed25519"
	default:
		return "secp256k1"
	}
}

const (
	ECDSA CurveType = iota
	EDDSA
)

type LocalPartySaveData struct {
	ty   CurveType // 0:secp256k1; 1:ed25519
	data any
}

func BuildECDSALocalPartySaveData() *LocalPartySaveData {
	return &LocalPartySaveData{}
}

func BuildEDDSALocalPartySaveData() *LocalPartySaveData {
	return &LocalPartySaveData{ty: EDDSA}
}

func (d *LocalPartySaveData) SetData(data any) *LocalPartySaveData {
	d.data = data
	return d
}

func (d *LocalPartySaveData) GetData() any {
	return d.data
}

func (d *LocalPartySaveData) EC() elliptic.Curve {
	switch d.ty {
	case EDDSA:
		return tss.Edwards()
	default:
		return tss.S256()
	}
}

func (d *LocalPartySaveData) CurveType() CurveType {
	return d.ty
}

func (d *LocalPartySaveData) ECDSAData() *ecdsaKeygen.LocalPartySaveData {
	if d.ty == ECDSA && d.data != nil {
		return d.data.(*ecdsaKeygen.LocalPartySaveData)
	}

	return nil
}

func (d *LocalPartySaveData) EDDSAData() *eddsaKeygen.LocalPartySaveData {
	if d.ty == EDDSA && d.data != nil {
		return d.data.(*eddsaKeygen.LocalPartySaveData)
	}

	return nil
}

func (d *LocalPartySaveData) Address() string {
	switch d.ty {
	case ECDSA:
		return crypto.PubkeyToAddress(*d.EDDSAData().EDDSAPub.ToECDSAPubKey()).String()
	default:
		panic("implement me")
	}
}
