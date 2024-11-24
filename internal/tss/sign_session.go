package tss

import (
	"crypto/rand"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	log "github.com/sirupsen/logrus"
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
	innerSession.inToOut = m.sigInToOut
	innerSession.partyIdMap = partyIdMap
	// party, endCh, errCh := RunParty(innerSession.ctx, msg, params, key, innerSession, keyDerivationDelta) // todo
	party, endCh, outCh, errCh := createSignParty(msg, params, key, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	session := &SignSession[ProposalID, *Proposal, *tsscommon.SignatureData]{sessionTransport: innerSession}
	session.Run()
	m.AddSession(session)

	helper.RunParty(session.ctx, party, errCh, outCh, session, false)

	return session.SessionID()
}

func (s *SignSession[T, M, D]) Post(data *helper.ReceivedPartyState) {
	if data.IsBroadcast || s.Included(data.ToPartyIds) {
		s.sessionTransport.Post(data)
	}
}

func createSignParty(
	msg *big.Int,
	params *tss.Parameters,
	key helper.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) (tss.Party, chan *tsscommon.SignatureData, chan tss.Message, chan *tss.Error) {
	// outgoing messages to other peers - not one to not deadlock when a party
	// round is waiting for outgoing messages channel to clear
	outCh := make(chan tss.Message, 100000)
	// output signature when finished
	endCh := make(chan *tsscommon.SignatureData, 256)
	// error if signing fails, contains culprits to blame
	errCh := make(chan *tss.Error, 256)

	log.Debug("creating new local party")

	var party tss.Party

	switch key.CurveType() {
	case helper.ECDSA:
		if keyDerivationDelta != nil {
			party = signing.NewLocalPartyWithKDD(msg, params, *key.ECDSAData(), keyDerivationDelta, outCh, endCh)
		} else {
			party = signing.NewLocalParty(msg, params, *key.ECDSAData(), outCh, endCh)
		}

	default:
		panic("implement me")
	}

	log.Debug("local signing party created", "partyID", party.PartyID())

	// helper.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, outCh, errCh
}
