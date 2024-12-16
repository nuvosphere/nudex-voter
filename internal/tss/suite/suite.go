package suite

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//type SignReq struct {
//	ReqId     string
//	chainType uint8
//	Signer    string
//	Data      any
//}
//
//type SignRes struct {
//	ReqId string
//	Data  any
//}

type SignReq interface {
	ReqId() *big.Int
	ChainType() uint8
	Signer() string
	SignDigest() string
	SignData() []byte
	ExtraData() []byte
}

type SignRes interface {
	ReqId() *big.Int
	Signature() []byte
}

type TssService interface {
	// PublicKey() crypto.PublicKey
	// ECPoint() *tssCrypto.ECPoint
	TssSigner() common.Address
	IsMeeting(signDigest string) bool
	Sign(req SignReq) error
	RegisterTssClient(client TssClient)
}

type TssClient interface {
	Verify(reqId *big.Int, signDigest string, ExtraData []byte) error
	PostSignature(res SignRes) error
	ChainType() uint8
}
