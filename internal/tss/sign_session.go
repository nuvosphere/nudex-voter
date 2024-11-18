package tss

import (
	"crypto/rand"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	sessionID helper.SessionID,
	taskID TaskId,
	msg *Msg,
	key keygen.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) helper.SessionID {
	allPartners := m.Participants()
	params, partyIdMap := NewParam(m.LocalSubmitter(), allPartners.Threshold(), allPartners)
	innerSession := newSession[TaskId, *Msg, *tsscommon.SignatureData](
		m.p2p,
		m,
		sessionID,
		crypto.PubkeyToAddress(*key.ECDSAPub.ToECDSAPubKey()), // todo
		m.Proposer(),
		taskID,
		msg,
		SignTaskSessionType,
		allPartners,
	)
	party, endCh, errCh := helper.Run(m.ctx, msg, params, key, innerSession, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.partyIdMap = partyIdMap
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	innerSession.inToOut = m.sigInToOut
	session := &SignSession[TaskId, *Msg, *tsscommon.SignatureData]{
		sessionTransport: innerSession,
	}
	session.Run()
	m.AddSession(session)

	return session.SessionID()
}
