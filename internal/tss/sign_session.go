package tss

import (
	"crypto/rand"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaSigning "github.com/bnb-chain/tss-lib/v2/eddsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

var _ Session[any] = &SignSession[any, any, any]{}

type SignSession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func RandSessionID() types.SessionID {
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	return common.BytesToHash(b[:])
}

func (m *Scheduler) NewMasterSignBatchSession(
	sessionID types.SessionID,
	proposalID ProposalID,
	msg *Proposal,
	data []ProposalID,
) types.SessionID {
	return m.NewSignSessionWitKey(sessionID, proposalID, msg, *m.partyData.ECDSALocalData(), nil, SignBatchTaskSessionType, data, m.partyData.ECDSALocalData().TssSigner())
}

func (m *Scheduler) NewSignSession(
	sessionID types.SessionID,
	proposalID ProposalID,
	msg *Proposal,
	key types.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) types.SessionID {
	return m.NewSignSessionWitKey(sessionID, proposalID, msg, key, keyDerivationDelta, SignTaskSessionType, nil, key.TssSigner())
}

func (m *Scheduler) NewTxSignSession(
	sessionID types.SessionID,
	proposalID ProposalID,
	msg *Proposal,
	key types.LocalPartySaveData,
	keyDerivationDelta *big.Int,
	signer string, // current signer
) types.SessionID {
	return m.NewSignSessionWitKey(sessionID, proposalID, msg, key, keyDerivationDelta, TxSignatureSessionType, nil, signer)
}

func (m *Scheduler) NewSignSessionWitKey(
	sessionID types.SessionID,
	proposalID ProposalID,
	msg *Proposal,
	key types.LocalPartySaveData,
	keyDerivationDelta *big.Int,
	ty string,
	data []ProposalID,
	signer string, // current signer
) types.SessionID {
	allPartners := m.Participants()
	params, partyIdMap := NewParam(key.CurveType(), m.LocalSubmitter(), allPartners)
	innerSession := newSession[ProposalID, *Proposal, *tsscommon.SignatureData](
		key.CurveType(),
		m.p2p,
		m,
		sessionID,
		signer,
		m.Proposer(),
		proposalID,
		msg,
		ty,
		allPartners,
	)
	innerSession.session.Data = data
	innerSession.inToOut = m.sigInToOut
	innerSession.partyIdMap = partyIdMap
	party, endCh, outCh, errCh := createSignParty(msg, params, key, keyDerivationDelta)
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
	key types.LocalPartySaveData,
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
	case types.ECDSA:
		if keyDerivationDelta != nil {
			party = ecdsaSigning.NewLocalPartyWithKDD(msg, params, *key.ECDSAData(), keyDerivationDelta, outCh, endCh)
		} else {
			party = ecdsaSigning.NewLocalParty(msg, params, *key.ECDSAData(), outCh, endCh)
		}
	case types.EDDSA:
		if keyDerivationDelta != nil {
			party = eddsaSigning.NewLocalPartyWithKDD(msg, params, *key.EDDSAData(), keyDerivationDelta, outCh, endCh)
		} else {
			party = eddsaSigning.NewLocalParty(msg, params, *key.EDDSAData(), outCh, endCh)
		}

	default:
		panic("implement me")
	}

	log.Debug("local signing party created", "partyID", party.PartyID())

	// types.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, outCh, errCh
}
