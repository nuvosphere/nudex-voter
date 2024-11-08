package tss

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
)

var (
	_ Session[any] = &GenerateKeySession[any, any, any]{}
	_ Session[any] = &ReShareGroupSession[any, any, any]{}
	_ Session[any] = &SignSession[any, any, any]{}
)

type SessionMessage[T any] struct {
	Type                    string           `json:"type"`
	GroupID                 helper.GroupID   `json:"groupID,omitempty"`
	SessionID               helper.SessionID `json:"sessionID,omitempty"`
	Sponsor                 common.Address   `json:"sponsor,omitempty"` // current submitter
	TaskID                  T                `json:"taskID,omitempty"`  // msg id
	FromPartyId             string           `json:"from_party_id"`
	ToPartyIds              []string         `json:"to_party_ids"`
	IsBroadcast             bool             `json:"is_broadcast"`
	IsToOldCommittee        bool             `json:"is_to_old_committee"`          // whether the message should be sent to old committee participants rather than the new committee
	IsToOldAndNewCommittees bool             `json:"is_to_old_and_new_committees"` // whether the message should be sent to both old and new committee participants
	MsgWireBytes            []byte           `json:"msg_wire_bytes"`
}

func (s *SessionMessage[T]) State(from *tss.PartyID) *helper.ReceivedPartyState {
	return &helper.ReceivedPartyState{
		WireBytes:               s.MsgWireBytes,
		From:                    from,
		IsBroadcast:             s.IsBroadcast,
		IsToOldCommittee:        s.IsToOldCommittee,
		IsToOldAndNewCommittees: s.IsToOldAndNewCommittees,
	}
}

// sessionTransport is a transport for a specific session.
type sessionTransport[T, M, D any] struct {
	broadcaster    p2p.P2PService                  // send data
	recvChan       chan *helper.ReceivedPartyState // receive data
	partyIDStore   *helper.PartyIDStore
	session        helper.Session[T, M]
	sessionRelease SessionReleaser
	ty             string
	party          tss.Party
	partyIdMap     map[string]*tss.PartyID
	rw             sync.RWMutex
	ctx            context.Context
	cancel         context.CancelFunc
	endCh          chan D
	errCH          chan *tss.Error
	inToOut        chan<- *SessionResult[T, D]
}

const (
	GenKeySessionType       = "GenerateKeySession"
	ReShareGroupSessionType = "ReShareGroupSession"
	SignSessionType         = "SignSession"
)

type GenerateKeySession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func NewParam(
	sponsor common.Address, // current submitter
	threshold int,
	allPartners []common.Address,
) (*tss.Parameters, map[string]*tss.PartyID) {
	partyIDs := createPartyIDsByAddress(allPartners)
	partyID := partyIDs.FindByKey(new(big.Int).SetBytes(sponsor.Bytes()))
	peerCtx := tss.NewPeerContext(partyIDs)
	params := tss.NewParameters(tss.S256(), peerCtx, partyID, len(partyIDs), threshold)
	partyIdMap := lo.SliceToMap(partyIDs, func(item *tss.PartyID) (string, *tss.PartyID) {
		return item.Id, item
	})

	return params, partyIdMap
}

func (m *Scheduler[T]) NewGenerateKeySession(
	p p2p.P2PService,
	sponsor common.Address, // current submitter
	taskID T, // msg id
	msg *big.Int,
	threshold int,
	allPartners []common.Address,
) helper.SessionID {
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	params, partyIdMap := NewParam(sponsor, threshold, allPartners)
	s := newSession[T, *big.Int, *keygen.LocalPartySaveData](p, m, helper.SenateGroupID, helper.SenateSessionID, sponsor, taskID, msg, threshold, GenKeySessionType, allPartners)
	party, endCh, errCh := helper.RunKeyGen(context.Background(), preParams, params, s) // todo
	s.party = party
	s.partyIdMap = partyIdMap
	s.inToOut = m.genKeyInToOut
	s.errCH = errCh
	s.endCh = endCh
	session := &GenerateKeySession[T, *big.Int, *keygen.LocalPartySaveData]{sessionTransport: s}
	m.AddSession(session)
	session.Run()

	return session.SessionID()
}

func (m *GenerateKeySession[T, M, D]) Release() {
	m.sessionTransport.Release()
	close(m.endCh)
	close(m.errCH)
}

func (m *GenerateKeySession[T, M, D]) Run() {
	go func() {
		defer m.Release()
		select {
		case <-m.ctx.Done():
		case data := <-m.endCh:
			err := saveTSSData(data)
			utils.Assert(err)
		case err := <-m.errCH:
			panic(err)
		}
	}()
}

type ReShareGroupSession[T, M, D any] struct {
	*GenerateKeySession[T, M, D]
}

func (m *Scheduler[T]) NewReShareGroupSession(
	p p2p.P2PService,
	sponsor common.Address, // current submitter
	taskID T, // msg id
	msg *big.Int,
	threshold int,
	allPartners []common.Address,
) helper.SessionID {
	// return newSession(p, m, groupID, helper.SenateSessionID, sponsor, taskID, msg, threshold, ReShareGroupSessionType, allPartners)
	return helper.SessionID{}
}

type SignSession[T, M, D any] struct {
	*sessionTransport[T, M, D]
}

func (m *SignSession[T, M, D]) Release() {
	m.sessionTransport.Release()
	close(m.endCh)
	close(m.errCH)
}

func (m *SignSession[T, M, D]) Run() {
	go func() {
		defer m.Release()
		select {
		case <-m.ctx.Done():
		case data := <-m.endCh:
			m.inToOut <- m.newDataResult(data)
		case err := <-m.errCH:
			m.inToOut <- m.newErrResult(err)
		}
	}()
}

func RandSessionID() helper.SessionID {
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	return common.BytesToHash(b[:])
}

func (m *Scheduler[T]) NewSignSession(
	p p2p.P2PService,
	groupID helper.GroupID,
	sponsor common.Address,
	taskID T,
	msg *big.Int,
	threshold int,
	allPartners []common.Address,
	key keygen.LocalPartySaveData,
	keyDerivationDelta *big.Int,
) helper.SessionID {
	params, partyIdMap := NewParam(sponsor, threshold, allPartners)
	innerSession := newSession[T, *big.Int, *tsscommon.SignatureData](p, m, groupID, RandSessionID(), sponsor, taskID, msg, threshold, GenKeySessionType, allPartners)
	party, endCh, errCh := helper.Run(context.Background(), msg, params, key, innerSession, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.partyIdMap = partyIdMap
	innerSession.endCh = endCh
	innerSession.errCH = errCh
	innerSession.inToOut = m.sigInToOut
	session := &SignSession[T, *big.Int, *tsscommon.SignatureData]{
		sessionTransport: innerSession,
	}
	m.AddSession(session)
	session.Run()

	return session.SessionID()
}

func newSession[T comparable, M, D any](
	p p2p.P2PService,
	m *Scheduler[T],
	groupID helper.GroupID,
	sessionID helper.SessionID,
	sponsor common.Address, // current submitter
	taskID T, // msg id
	msg M,
	threshold int,
	ty string,
	allPartners []common.Address,
) *sessionTransport[T, M, D] {
	recvChan := make(chan *helper.ReceivedPartyState, 1)

	return &sessionTransport[T, M, D]{
		broadcaster: p,
		recvChan:    recvChan,
		session: helper.Session[T, M]{
			GroupID:   groupID,
			SessionID: sessionID,
			Sponsor:   sponsor,
			TaskID:    taskID,
			Msg:       msg,
			Threshold: threshold,
		},
		sessionRelease: m,
		ty:             ty,
		partyIDStore:   helper.NewPartyIDStore(),
		partyIdMap:     make(map[string]*tss.PartyID),
		rw:             sync.RWMutex{},
	}
}

func (s *sessionTransport[T, M, D]) Type() string {
	return s.ty
}

func (s *sessionTransport[T, M, D]) Name() string {
	return fmt.Sprintf(
		"%s: seesionID=%v, groupID=%v,taskID=%v,msg=%v,threshold=%d",
		s.Type(),
		s.SessionID(),
		s.GroupID(),
		s.TaskID(),
		s.session.Msg,
		s.Threshold(),
	)
}

func (s *sessionTransport[T, M, D]) PartyID(id string) *tss.PartyID {
	return s.party.PartyID()
}

func (s *sessionTransport[T, M, D]) Party() tss.Party {
	return s.party
}

func (s *sessionTransport[T, M, D]) SessionID() helper.SessionID {
	return s.session.SessionID
}

func (s *sessionTransport[T, M, D]) GroupID() helper.GroupID {
	return s.session.GroupID
}

func (s *sessionTransport[T, M, D]) TaskID() T {
	return s.session.TaskID
}

func (s *sessionTransport[T, M, D]) Sponsor() common.Address {
	return s.session.Sponsor
}

func (s *sessionTransport[T, M, D]) Threshold() int {
	return s.session.Threshold
}

func (s *sessionTransport[T, M, D]) Release() {
	s.sessionRelease.SessionRelease(s.SessionID())
	s.cancel()
	close(s.recvChan)
}

func (s *sessionTransport[T, M, D]) Send(ctx context.Context, bytes []byte, routing *tss.MessageRouting, b bool) error {
	msg := p2p.Message[any]{
		MessageType: p2p.MessageTypeTssMsg,
		RequestId:   fmt.Sprintf("%v", s.TaskID()), // todo taskID
		DataType:    s.Type(),                      // todo
		Data: SessionMessage[T]{
			Type:                    s.Type(),
			GroupID:                 s.GroupID(),
			SessionID:               s.SessionID(),
			Sponsor:                 s.Sponsor(),
			TaskID:                  s.TaskID(),
			FromPartyId:             routing.From.Id,
			ToPartyIds:              lo.Map(routing.To, func(to *tss.PartyID, _ int) string { return to.Id }),
			IsBroadcast:             routing.IsBroadcast,
			IsToOldCommittee:        routing.IsToOldCommittee,
			IsToOldAndNewCommittees: routing.IsToOldAndNewCommittees,
			MsgWireBytes:            bytes,
		},
	}

	return s.broadcaster.PublishMessage(ctx, msg)
}

func (s *sessionTransport[T, M, D]) Receive() chan *helper.ReceivedPartyState {
	return s.recvChan
}

func (s *sessionTransport[T, M, D]) Post(data *helper.ReceivedPartyState) {
	s.recvChan <- data
}

func (s *sessionTransport[T, M, D]) newDataResult(data D) *SessionResult[T, D] {
	return &SessionResult[T, D]{
		TaskID:    s.TaskID(),
		SessionID: s.SessionID(),
		GroupID:   s.GroupID(),
		Data:      data,
		Err:       nil,
	}
}

func (s *sessionTransport[T, M, D]) newErrResult(err error) *SessionResult[T, D] {
	return &SessionResult[T, D]{
		TaskID:    s.TaskID(),
		SessionID: s.SessionID(),
		GroupID:   s.GroupID(),
		Err:       err,
	}
}

type SessionResult[T, D any] struct {
	TaskID    T
	SessionID helper.SessionID
	GroupID   helper.GroupID
	Data      D
	Err       error
}
