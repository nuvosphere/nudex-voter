package tss

import (
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
)

var _ Session[any] = &GenerateKeySession[any, any, any]{}

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (m *Scheduler) NewGenerateKeySession(
	ec crypto.CurveType,
	proposalID ProposalID, // msg id
	sessionID party.SessionID,
	signer string,
	msg *Proposal,
) party.SessionID {
	allPartners := m.Participants()
	s := newSession[ProposalID, *Proposal, *LocalPartySaveData](
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
	party, partyIdMap, endCh, errCh := RunKeyGen(s.ctx, m.isProd, ec, m.LocalSubmitter(), allPartners, s) // todo
	s.party = party
	s.partyIdMap = partyIdMap
	s.errCH = errCh
	s.endCh = endCh
	s.inToOut = m.senateInToOut
	s.Run()
	session := &GenerateKeySession[ProposalID, *Proposal, *LocalPartySaveData]{sessionTransport: s}
	m.AddSession(session)

	return session.SessionID()
}

func (g *GenerateKeySession[T, M, D]) Post(data *helper.ReceivedPartyState) {
	if data.IsBroadcast || g.Included(data.ToPartyIds) {
		g.sessionTransport.Post(data)
	}
}
