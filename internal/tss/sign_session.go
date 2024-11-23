package tss

import (
	"crypto/rand"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

var _ Session[any] = &SignSession[any, any, any]{}

type SignSession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func RandSessionID() helper.SessionID {
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	return common.BytesToHash(b[:])
}

func (m *Scheduler) NewSignSession(
	ec helper.CurveType,
	sessionID helper.SessionID,
	proposalID ProposalID,
	msg *Proposal,
	key helper.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) helper.SessionID {
	allPartners := m.Participants()
	params, partyIdMap := NewParam(ec, m.LocalSubmitter(), allPartners)
	innerSession := newSession[ProposalID, *Proposal, *tsscommon.SignatureData](
		ec,
		m.p2p,
		m,
		sessionID,
		common.HexToAddress(key.Address()), // todo
		m.Proposer(),
		proposalID,
		msg,
		SignTaskSessionType,
		allPartners,
	)
	party, endCh, errCh := RunParty(innerSession.ctx, msg, params, key, innerSession, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.partyIdMap = partyIdMap
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	innerSession.inToOut = m.sigInToOut
	session := &SignSession[ProposalID, *Proposal, *tsscommon.SignatureData]{
		sessionTransport: innerSession,
	}
	session.Run()
	m.AddSession(session)

	return session.SessionID()
}

func (s *SignSession[T, M, D]) Post(data *helper.ReceivedPartyState) {
	if data.IsBroadcast || s.Included(data.ToPartyIds) {
		s.sessionTransport.Post(data)
	}
}
