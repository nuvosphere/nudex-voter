package tss

import (
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

var _ Session[any] = &GenerateKeySession[any, any, any]{}

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (m *Scheduler) NewGenerateKeySession(
	proposer common.Address, // current submitter
	taskID TaskId, // msg id
	msg *Msg,
	allPartners types.Participants,
) helper.SessionID {
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	params, partyIdMap := NewParam(m.LocalSubmitter(), allPartners.Threshold(), allPartners)
	s := newSession[TaskId, *Msg, *keygen.LocalPartySaveData](
		m.p2p,
		m,
		helper.SenateGroupID,
		helper.SenateSessionID,
		proposer,
		taskID,
		msg,
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
	session := &GenerateKeySession[TaskId, *Msg, *keygen.LocalPartySaveData]{sessionTransport: s}
	m.AddSession(session)

	return session.SessionID()
}

func (m *GenerateKeySession[T, M, D]) Release() {
	m.sessionTransport.Release()
}
