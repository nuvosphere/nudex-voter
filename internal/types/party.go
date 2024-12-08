package types

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"

	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	SessionID = common.Hash
	GroupID   = common.Hash
)

var (
	ZeroSessionID           SessionID
	senateSessionID         = crypto.Keccak256Hash([]byte("The voter senate session，one and only one"))
	SenateProposal          = senateSessionID.Big()
	senateProposalID        = senateSessionID.Big().Uint64()
	SenateProposalIDOfECDSA = senateProposalID - 1
	SenateProposalIDOfEDDSA = senateProposalID - 2
	SenateSessionIDOfECDSA  = crypto.Keccak256Hash([]byte("ECDSA:The voter senate session，one and only one"))
	SenateSessionIDOfEDDSA  = crypto.Keccak256Hash([]byte("EDDSA:The voter senate session，one and only one"))
)

type Session[T, M any] struct {
	Group
	SessionID  SessionID      `json:"sessionID,omitempty"`
	Proposer   common.Address `json:"proposer,omitempty"`    // current submitter
	Signer     string         `json:"signer,omitempty"`      // current signer
	ProposalID T              `json:"proposal_id,omitempty"` // msg id
	Proposal   M              `json:"proposal,omitempty"`
	Data       []T            `json:"data,omitempty"`
}

type Group struct {
	EC          CurveType    `json:"ec,omitempty"`
	AllPartners Participants `json:"all_partners,omitempty"` // all submitter
}

func (g *Group) GroupID() GroupID {
	return g.AllPartners.GroupID()
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

func (e *CurveType) String() string {
	return e.CurveName()
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
	return &LocalPartySaveData{ty: ECDSA}
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

func (d *LocalPartySaveData) TssSigner() string {
	switch d.ty {
	case ECDSA:
		return crypto.PubkeyToAddress(*d.ECDSAData().ECDSAPub.ToECDSAPubKey()).String()
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) ToECDSAPubKey() *ecdsa.PublicKey {
	switch d.ty {
	case ECDSA:
		return d.ECDSAData().ECDSAPub.ToECDSAPubKey()
	case EDDSA:
		return d.EDDSAData().EDDSAPub.ToECDSAPubKey()
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) ECPoint() *tssCrypto.ECPoint {
	switch d.ty {
	case ECDSA:
		return d.ECDSAData().ECDSAPub
	case EDDSA:
		return d.EDDSAData().EDDSAPub
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) PublicKeyBase58() string {
	switch d.ty {
	case ECDSA:
		pubKey := d.ECDSAData().ECDSAPub.ToECDSAPubKey()
		pubKeyBytes := crypto.FromECDSAPub(pubKey)
		return base58.Encode(pubKeyBytes)
	case EDDSA:
		pubKey := d.EDDSAData().EDDSAPub.ToECDSAPubKey()
		pubKeyBytes := crypto.FromECDSAPub(pubKey)
		return base58.Encode(pubKeyBytes)
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) PublicKey() string {
	p := d.ECPoint()
	switch d.ty {
	case ECDSA:
		var (
			x = &btcec.FieldVal{}
			y = &btcec.FieldVal{}
		)
		x.SetByteSlice(p.X().Bytes())
		y.SetByteSlice(p.Y().Bytes())
		return hex.EncodeToString(btcec.NewPublicKey(x, y).SerializeCompressed())
	case EDDSA:
		p.ToECDSAPubKey()
		pubkey := edwards.NewPublicKey(p.X(), p.Y())
		return hex.EncodeToString(pubkey.SerializeCompressed())
	default:
		panic("implement me")
	}
}
