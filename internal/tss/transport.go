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
	_ Session[any] = &GenerateKeySession[any, any]{}
	_ Session[any] = &ReShareGroupSession[any, any]{}
	_ Session[any] = &SignSession[any]{}
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
type sessionTransport[T, M any] struct {
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
}

const (
	GenKeySessionType       = "GenerateKeySession"
	ReShareGroupSessionType = "ReShareGroupSession"
	SignSessionType         = "SignSession"
)

type GenerateKeySession[T, M any] struct {
	*sessionTransport[T, M]
	endCh chan *keygen.LocalPartySaveData
	errCH chan *tss.Error
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

func NewGenerateKeySession[T comparable, M any](
	p p2p.P2PService,
	m *Manager[T],
	sponsor common.Address, // current submitter
	taskID T, // msg id
	msg M,
	threshold int,
	allPartners []common.Address,
) Session[T] {
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	params, partyIdMap := NewParam(sponsor, threshold, allPartners)
	s := newSession(p, m, helper.SenateGroupID, helper.SenateSessionID, sponsor, taskID, msg, threshold, GenKeySessionType, allPartners)
	party, endCh, errCh := helper.RunKeyGen(context.Background(), preParams, params, s) // todo
	s.party = party
	s.partyIdMap = partyIdMap

	return &GenerateKeySession[T, M]{
		sessionTransport: s,
		endCh:            endCh,
		errCH:            errCh,
	}
}

func (m *GenerateKeySession[T, M]) Release() {
	m.sessionTransport.Release()
	close(m.endCh)
	close(m.errCH)
}

func (m *GenerateKeySession[T, M]) Run() {
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

type ReShareGroupSession[T, M any] struct {
	*GenerateKeySession[T, M]
}

func NewReShareGroupSession[T comparable, M any](
	p p2p.P2PService,
	m *Manager[T],
	sponsor common.Address, // current submitter
	taskID T, // msg id
	msg M,
	threshold int,
	allPartners []common.Address,
) Session[T] {
	// return newSession(p, m, groupID, helper.SenateSessionID, sponsor, taskID, msg, threshold, ReShareGroupSessionType, allPartners)
	return nil
}

type SignSession[T any] struct {
	*sessionTransport[T, *big.Int]
	inToOut chan<- *SignResult[T]
	endCh   chan *tsscommon.SignatureData
	errCH   chan *tss.Error
}

func (m *SignSession[T]) Release() {
	m.sessionTransport.Release()
	close(m.endCh)
	close(m.errCH)
}

func (m *SignSession[T]) Run() {
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

type SignResult[T any] struct {
	TaskID    T
	SessionID helper.SessionID
	GroupID   helper.GroupID
	Data      *tsscommon.SignatureData
	Err       error
}

func (m *SignSession[T]) newDataResult(data *tsscommon.SignatureData) *SignResult[T] {
	return &SignResult[T]{
		TaskID:    m.TaskID(),
		SessionID: m.SessionID(),
		GroupID:   m.GroupID(),
		Data:      data,
		Err:       nil,
	}
}

func (m *SignSession[T]) newErrResult(err error) *SignResult[T] {
	return &SignResult[T]{
		TaskID:    m.TaskID(),
		SessionID: m.SessionID(),
		GroupID:   m.GroupID(),
		Data:      nil,
		Err:       err,
	}
}

func RandSessionID() helper.SessionID {
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	return common.BytesToHash(b[:])
}

func NewSignSession[T comparable](
	p p2p.P2PService,
	m *Manager[T],
	groupID helper.GroupID,
	sponsor common.Address,
	taskID T,
	msg *big.Int,
	threshold int,
	allPartners []common.Address,
	key keygen.LocalPartySaveData,
	keyDerivationDelta *big.Int,
	inToOut chan<- *SignResult[T],
) *SignSession[T] {
	params, partyIdMap := NewParam(sponsor, threshold, allPartners)
	innerSession := newSession(p, m, groupID, RandSessionID(), sponsor, taskID, msg, threshold, GenKeySessionType, allPartners)
	party, endCh, errCh := helper.Run(context.Background(), msg, params, key, innerSession, keyDerivationDelta) // todo
	innerSession.party = party
	innerSession.partyIdMap = partyIdMap

	return &SignSession[T]{
		sessionTransport: innerSession,
		inToOut:          inToOut,
		endCh:            endCh,
		errCH:            errCh,
	}
}

func newSession[T comparable, M any](
	p p2p.P2PService,
	m *Manager[T],
	groupID helper.GroupID,
	sessionID helper.SessionID,
	sponsor common.Address, // current submitter
	taskID T, // msg id
	msg M,
	threshold int,
	ty string,
	allPartners []common.Address,
) *sessionTransport[T, M] {
	recvChan := make(chan *helper.ReceivedPartyState, 1)

	return &sessionTransport[T, M]{
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

func (s *sessionTransport[T, M]) Type() string {
	return s.ty
}

func (s *sessionTransport[T, M]) Name() string {
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

func (s *sessionTransport[T, M]) PartyID(id string) *tss.PartyID {
	return s.party.PartyID()
}

func (s *sessionTransport[T, M]) Party() tss.Party {
	return s.party
}

func (s *sessionTransport[T, M]) SessionID() helper.SessionID {
	return s.session.SessionID
}

func (s *sessionTransport[T, M]) GroupID() helper.GroupID {
	return s.session.GroupID
}

func (s *sessionTransport[T, M]) TaskID() T {
	return s.session.TaskID
}

func (s *sessionTransport[T, M]) Sponsor() common.Address {
	return s.session.Sponsor
}

func (s *sessionTransport[T, M]) Threshold() int {
	return s.session.Threshold
}

func (s *sessionTransport[T, M]) Release() {
	s.sessionRelease.SessionRelease(s.SessionID())
	s.cancel()
	close(s.recvChan)
}

func (s *sessionTransport[T, M]) Send(ctx context.Context, bytes []byte, routing *tss.MessageRouting, b bool) error {
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

func (s *sessionTransport[T, M]) Receive() chan *helper.ReceivedPartyState {
	return s.recvChan
}

func (s *sessionTransport[T, M]) Post(data *helper.ReceivedPartyState) {
	s.recvChan <- data
}
