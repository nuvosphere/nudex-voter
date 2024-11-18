package tss

import (
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

var _ Session[any] = &GenerateKeySession[any, any, any]{}

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (m *Scheduler) NewGenerateKeySession(
	taskID TaskId, // msg id
	msg *Msg,
) helper.SessionID {
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	allPartners := m.Participants()
	params, partyIdMap := NewParam(m.LocalSubmitter(), allPartners.Threshold(), allPartners)
	s := newSession[TaskId, *Msg, *keygen.LocalPartySaveData](
		m.p2p,
		m,
		helper.SenateSessionID,
		m.MasterSigner(),
		m.Proposer(),
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

func (m *Scheduler) MasterSigner() common.Address {
	return crypto.PubkeyToAddress(m.MasterPublicKey())
}

func (m *GenerateKeySession[T, M, D]) Release() {
	m.sessionTransport.Release()
}
