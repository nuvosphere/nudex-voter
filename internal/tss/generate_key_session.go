package tss

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

var _ Session[any] = &GenerateKeySession[any, any, any]{}

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (m *Scheduler) NewGenerateKeySession(
	ec helper.CurveType,
	proposalID ProposalID, // msg id
	sessionID helper.SessionID,
	signer common.Address,
	msg *Proposal,
) helper.SessionID {
	allPartners := m.Participants()
	s := newSession[ProposalID, *Proposal, *helper.LocalPartySaveData](
		ec,
		m.p2p,
		m,
		sessionID,
		signer,
		m.Proposer(),
		proposalID,
		msg,
		GenKeySessionType,
		allPartners,
	)
	party, partyIdMap, endCh, errCh := RunKeyGen(m.ctx, ec, m.localSubmitter, allPartners, s) // todo
	s.party = party
	s.partyIdMap = partyIdMap
	s.errCH = errCh
	s.endCh = endCh
	s.inToOut = m.senateInToOut
	s.Run()
	session := &GenerateKeySession[ProposalID, *Proposal, *helper.LocalPartySaveData]{sessionTransport: s}
	m.AddSession(session)

	return session.SessionID()
}

func (m *GenerateKeySession[T, M, D]) Release() {
	m.sessionTransport.Release()
}
