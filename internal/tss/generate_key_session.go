package tss

import (
	"math/big"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

var _ Session[any] = &GenerateKeySession[any, any, any]{}

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (m *Scheduler[T]) NewGenerateKeySession(
	proposer common.Address, // current submitter
	taskID T, // msg id
	msg *big.Int,
	threshold int,
	allPartners []common.Address,
) helper.SessionID {
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	params, partyIdMap := NewParam(m.localSubmitter, threshold, allPartners)
	s := newSession[T, *big.Int, *keygen.LocalPartySaveData](
		m.p2p,
		m,
		helper.SenateGroupID,
		helper.SenateSessionID,
		proposer,
		taskID,
		msg,
		threshold,
		GenKeySessionType,
		allPartners,
	)
	party, endCh, errCh := helper.RunKeyGen(m.ctx, preParams, params, s) // todo
	s.party = party
	s.partyIdMap = partyIdMap
	s.inToOut = m.senateInToOut
	s.errCH = errCh
	s.endCh = endCh
	s.Run()
	session := &GenerateKeySession[T, *big.Int, *keygen.LocalPartySaveData]{sessionTransport: s}
	m.AddSession(session)

	return session.SessionID()
}

func (m *GenerateKeySession[T, M, D]) Release() {
	m.sessionTransport.Release()
}
