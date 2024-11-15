package tss

import (
	"crypto/rand"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
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

func (m *Scheduler[T]) NewSignSession(
	groupID helper.GroupID,
	sessionID helper.SessionID,
	proposer, localSubmitter common.Address,
	taskID T,
	msg *Msg,
	allPartners types.Participants,
	key keygen.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) helper.SessionID {
	params, partyIdMap := NewParam(localSubmitter, allPartners.Threshold(), allPartners)
	innerSession := newSession[T, *Msg, *tsscommon.SignatureData](
		m.p2p,
		m,
		groupID,
		sessionID,
		proposer,
		taskID,
		msg,
		SignSessionType,
		allPartners,
	)
	party, endCh, errCh := helper.Run(m.ctx, msg, params, key, innerSession, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.partyIdMap = partyIdMap
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	innerSession.inToOut = m.sigInToOut
	session := &SignSession[T, *Msg, *tsscommon.SignatureData]{
		sessionTransport: innerSession,
	}
	session.Run()
	m.AddSession(session)

	return session.SessionID()
}
