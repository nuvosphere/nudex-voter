package tss

import (
	"crypto/rand"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
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

func (m *Scheduler[T]) NewSignSession(
	groupID helper.GroupID,
	sessionID helper.SessionID,
	proposer, localSubmitter common.Address,
	taskID T,
	msg *big.Int,
	threshold int,
	allPartners []common.Address,
	key keygen.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) helper.SessionID {
	params, partyIdMap := NewParam(localSubmitter, threshold, allPartners)
	innerSession := newSession[T, *big.Int, *tsscommon.SignatureData](
		m.p2p,
		m,
		groupID,
		sessionID,
		proposer,
		taskID,
		msg,
		threshold,
		SignSessionType,
		allPartners,
	)
	party, endCh, errCh := helper.Run(m.ctx, msg, params, key, innerSession, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.partyIdMap = partyIdMap
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	innerSession.inToOut = m.sigInToOut
	session := &SignSession[T, *big.Int, *tsscommon.SignatureData]{
		sessionTransport: innerSession,
	}
	session.Run()
	m.AddSession(session)

	return session.SessionID()
}
