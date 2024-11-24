package tss

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

var _ Session[any] = &sessionTransport[any, any, any]{}

type (
	ProposalID = int64
	Proposal   = big.Int
)

type SessionMessage[T, M any] struct {
	Type                    string           `json:"type"`
	GroupID                 helper.GroupID   `json:"group_id,omitempty"`
	SessionID               helper.SessionID `json:"session_id,omitempty"`
	Signer                  common.Address   `json:"signer,omitempty"`
	Proposer                common.Address   `json:"proposer,omitempty"`    // current submitter
	ProposalID              T                `json:"proposal_id,omitempty"` // msg id
	Proposal                M                `json:"proposal,omitempty"`
	FromPartyId             string           `json:"from_party_id"`
	ToPartyIds              []string         `json:"to_party_ids"`
	IsBroadcast             bool             `json:"is_broadcast"`
	IsToOldCommittee        bool             `json:"is_to_old_committee"`          // whether the message should be sent to old committee participants rather than the new committee
	IsToOldAndNewCommittees bool             `json:"is_to_old_and_new_committees"` // whether the message should be sent to both old and new committee participants
	MsgWireBytes            []byte           `json:"msg_wire_bytes"`
}

func (s *SessionMessage[T, M]) State(from *tss.PartyID) *helper.ReceivedPartyState {
	return &helper.ReceivedPartyState{
		WireBytes:               s.MsgWireBytes,
		From:                    from,
		ToPartyIds:              s.ToPartyIds,
		IsBroadcast:             s.IsBroadcast,
		IsToOldCommittee:        s.IsToOldCommittee,
		IsToOldAndNewCommittees: s.IsToOldAndNewCommittees,
	}
}

// sessionTransport is a transport for a specific session.
type sessionTransport[T, M, D any] struct {
	broadcaster    p2p.P2PService                  // send data
	recvChan       chan *helper.ReceivedPartyState // receive data
	session        helper.Session[T, M]
	sessionRelease SessionReleaser
	isReleased     atomic.Bool
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
	SignTaskSessionType     = "SignTaskSession"
	TxSignatureSessionType  = "TxSignatureSession"
)

func NewParam(
	ec helper.CurveType,
	localSubmitter common.Address,
	allPartners types.Participants,
) (*tss.Parameters, map[string]*tss.PartyID) {
	partyIDs := createPartyIDsByGroup(ec, allPartners)

	part := PartyKey(ec, allPartners, localSubmitter)
	partyID := partyIDs.FindByKey(part)

	peerCtx := tss.NewPeerContext(partyIDs)
	params := tss.NewParameters(ec.EC(), peerCtx, partyID, partyIDs.Len(), allPartners.Threshold())
	partyIdMap := lo.SliceToMap(partyIDs, func(item *tss.PartyID) (string, *tss.PartyID) { return strings.ToLower(item.Id), item })

	return params, partyIdMap
}

func newSession[T comparable, M, D any](
	ec helper.CurveType,
	p p2p.P2PService,
	m *Scheduler,
	sessionID helper.SessionID,
	signer common.Address, // current signer
	proposer common.Address, // current submitter
	ProposalId T, // msg id
	proposal M,
	ty string,
	allPartners types.Participants,
) *sessionTransport[T, M, D] {
	if sessionID == helper.ZeroSessionID {
		sessionID = RandSessionID()
	}

	recvChan := make(chan *helper.ReceivedPartyState, 1)
	ctx, cancel := context.WithCancel(context.Background())

	return &sessionTransport[T, M, D]{
		broadcaster: p,
		recvChan:    recvChan,
		session: helper.Session[T, M]{
			Group: helper.Group{
				EC:          ec,
				AllPartners: allPartners,
				GroupID:     allPartners.GroupID(),
			},
			SessionID:  sessionID,
			Signer:     signer,
			Proposer:   proposer,
			ProposalID: ProposalId,
			Proposal:   proposal,
			Threshold:  allPartners.Threshold(),
		},
		sessionRelease: m,
		ty:             ty,
		partyIdMap:     make(map[string]*tss.PartyID),
		rw:             sync.RWMutex{},
		ctx:            ctx,
		cancel:         cancel,
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
		s.ProposalID(),
		s.session.Proposal,
		s.Threshold(),
	)
}

func (s *sessionTransport[T, M, D]) PartyID(id string) *tss.PartyID {
	return s.partyIdMap[strings.ToLower(id)]
}

func (s *sessionTransport[T, M, D]) Equal(id string) bool {
	return strings.EqualFold(s.party.PartyID().Id, id)
}

func (s *sessionTransport[T, M, D]) Included(ids []string) bool {
	return slices.Contains(ids, strings.ToLower(s.party.PartyID().Id))
}

func (s *sessionTransport[T, M, D]) SessionID() helper.SessionID {
	return s.session.SessionID
}

func (s *sessionTransport[T, M, D]) GroupID() helper.GroupID {
	return s.session.GroupID
}

func (s *sessionTransport[T, M, D]) ProposalID() T {
	return s.session.ProposalID
}

func (s *sessionTransport[T, M, D]) Proposer() common.Address {
	return s.session.Proposer
}

func (s *sessionTransport[T, M, D]) Threshold() int {
	return s.session.Threshold
}

func (s *sessionTransport[T, M, D]) Release() {
	if !s.isReleased.Swap(true) {
		log.Infof("release session : %v, party id:%v", s.Name(), s.party.PartyID())
		s.sessionRelease.SessionRelease(s.SessionID())
		s.cancel()
		close(s.recvChan)
		close(s.endCh)
		close(s.errCH)
	}
}

func (s *sessionTransport[T, M, D]) Send(ctx context.Context, bytes []byte, routing *tss.MessageRouting, b bool) error {
	msg := p2p.Message[any]{
		MessageType: p2p.MessageTypeTssMsg,
		RequestId:   fmt.Sprintf("%v", s.ProposalID()), // todo taskID
		DataType:    s.Type(),                          // todo
		Data: SessionMessage[T, M]{
			Type:                    s.Type(),
			GroupID:                 s.GroupID(),
			SessionID:               s.SessionID(),
			Signer:                  s.Signer(),
			Proposer:                s.Proposer(),
			ProposalID:              s.ProposalID(),
			Proposal:                s.session.Proposal,
			FromPartyId:             strings.ToLower(routing.From.Id),
			ToPartyIds:              lo.Map(routing.To, func(to *tss.PartyID, _ int) string { return strings.ToLower(to.Id) }),
			IsBroadcast:             routing.IsBroadcast,
			IsToOldCommittee:        routing.IsToOldCommittee,
			IsToOldAndNewCommittees: routing.IsToOldAndNewCommittees,
			MsgWireBytes:            bytes,
		},
	}

	return s.broadcaster.PublishMessage(ctx, msg)
}

func (s *sessionTransport[T, M, D]) GetReceiveChannel(_ string) chan *helper.ReceivedPartyState {
	return s.recvChan
}

func (s *sessionTransport[T, M, D]) Post(data *helper.ReceivedPartyState) {
	if !s.Equal(data.From.Id) {
		s.recvChan <- data
	}
}

func (s *sessionTransport[T, M, D]) Run() {
	go func() {
		defer s.Release()
		select {
		case <-s.ctx.Done():
		case data := <-s.endCh:
			log.Infof("recived data: session id: %v", s.SessionID())
			s.inToOut <- s.newDataResult(data)
		case err := <-s.errCH:
			log.Infof("recived err: session id: %v", s.SessionID())
			s.inToOut <- s.newErrResult(err)
		}
	}()
}

func (s *sessionTransport[T, M, D]) Participants() types.Participants {
	return s.session.AllPartners
}

func (s *sessionTransport[T, M, D]) Signer() common.Address {
	return s.session.Signer
}

func (s *sessionTransport[T, M, D]) newDataResult(data D) *SessionResult[T, D] {
	return &SessionResult[T, D]{
		ProposalID: s.ProposalID(),
		SessionID:  s.SessionID(),
		GroupID:    s.GroupID(),
		Data:       data,
		Err:        nil,
	}
}

func (s *sessionTransport[T, M, D]) newErrResult(err error) *SessionResult[T, D] {
	return &SessionResult[T, D]{
		ProposalID: s.ProposalID(),
		SessionID:  s.SessionID(),
		GroupID:    s.GroupID(),
		Err:        err,
	}
}

type SessionResult[T, D any] struct {
	ProposalID T
	SessionID  helper.SessionID
	GroupID    helper.GroupID
	Data       D
	Err        error
}
