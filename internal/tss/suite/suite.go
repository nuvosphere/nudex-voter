package suite

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SignReq struct {
	SeqId      uint64
	ChainType  uint8
	Signer     string
	DataDigest string
	SignData   []byte
	ExtraData  []byte
}

type SignRes struct {
	SeqId      uint64
	DataDigest string
	Signature  []byte
}

type TssService interface {
	// PublicKey() crypto.PublicKey
	// ECPoint() *tssCrypto.ECPoint
	GetUserAddress(coinType, account uint32, index uint8) string
	TssSigner() common.Address
	IsMeeting(signDigest string) bool
	Sign(req *SignReq) error
	IsProposer() bool
	Proposer() common.Address
	LocalSubmitter() common.Address
	RegisterTssClient(client TssClient)
}

type TssClient interface {
	Verify(reqId *big.Int, signDigest string, ExtraData []byte) error
	ReceiveSignature(res *SignRes)
	ChainType() uint8
}
