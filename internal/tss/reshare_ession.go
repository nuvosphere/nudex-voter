package tss

import (
	"context"
	"errors"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

var _ Session[any] = &ReShareGroupSession[any, any, any]{}

type ReShareGroupSession[T, M, D any] struct {
	oldSession *sessionTransport[T, M, D]
	newSession *sessionTransport[T, M, D]
}

func (m *Scheduler) NewReShareGroupSession(
	runMode int64,
	ec helper.CurveType,
	taskID ProposalID, // msg id
	msg *Proposal,
	oldPartners types.Participants,
	newPartners types.Participants,
) helper.SessionID {
	localSubmitter := m.LocalSubmitter()

	log.Debugf("newPartners:%v, oldPartners: %v", newPartners, oldPartners)
	oldPartyIDs := createPartyIDsByGroup(ec, oldPartners)
	oldPeerCtx := tss.NewPeerContext(oldPartyIDs) // todo

	newPartyIDs := createPartyIDsByGroup(ec, newPartners)
	newPeerCtx := tss.NewPeerContext(newPartyIDs)
	newPartKey := PartyKey(ec, newPartners, localSubmitter)
	newPartyID := newPartyIDs.FindByKey(newPartKey)
	log.Debugf("newPartyID: %v", newPartyID)

	newPartyIdMap := lo.SliceToMap(newPartyIDs, func(item *tss.PartyID) (string, *tss.PartyID) { return item.Id, item })
	newParams := tss.NewReSharingParameters(
		ec.EC(),
		oldPeerCtx,
		newPeerCtx,
		newPartyID,
		oldPartners.Len(),
		oldPartners.Threshold(),
		newPartners.Len(),
		newPartners.Threshold(),
	)

	if runMode == JoinMode { // new node
		newInnerSession := newSession[ProposalID, *Proposal, *helper.LocalPartySaveData](
			ec,
			m.p2p,
			m,
			helper.SenateSessionID,
			common.Address{}, // todo
			localSubmitter,
			taskID,
			msg,
			ReShareGroupSessionType,
			newPartners,
		)

		localData := m.partyData.GenerateNewLocalPartySaveData(ec, newPartners)
		party, endCh, errCh := RunReshare(newInnerSession.ctx, newParams, *localData, newInnerSession) // new create node
		newInnerSession.party = party
		newInnerSession.partyIdMap = newPartyIdMap
		newInnerSession.endCh = endCh
		newInnerSession.errCH = errCh
		newInnerSession.inToOut = m.senateInToOut

		newInnerSession.Run()
		m.AddSession(newInnerSession)

		return newInnerSession.SessionID()
	}

	reShareSession := &ReShareGroupSession[ProposalID, *Proposal, *helper.LocalPartySaveData]{}

	oldPartKey := PartyKey(ec, oldPartners, localSubmitter)
	oldPartyID := oldPartyIDs.FindByKey(oldPartKey)
	oldParams := tss.NewReSharingParameters(
		ec.EC(),
		oldPeerCtx,
		newPeerCtx,
		oldPartyID,
		oldPartners.Len(),
		oldPartners.Threshold(),
		newPartners.Len(),
		newPartners.Threshold(),
	)
	oldPartyIdMap := lo.SliceToMap(oldPartyIDs, func(item *tss.PartyID) (string, *tss.PartyID) { return item.Id, item })
	oldInnerSession := newSession[ProposalID, *Proposal, *helper.LocalPartySaveData](
		ec,
		m.p2p,
		m,
		helper.SenateSessionID,
		common.Address{}, // todo
		localSubmitter,
		taskID,
		msg,
		ReShareGroupSessionType,
		newPartners, // todo
	)
	reShareSession.oldSession = oldInnerSession

	localData := m.partyData.GetData(ec)
	log.Debugf("localData: %v", localData.GetData())
	party, endCh, errCh := RunReshare(oldInnerSession.ctx, oldParams, *localData, reShareSession) // todo
	oldInnerSession.party = party
	oldInnerSession.partyIdMap = oldPartyIdMap
	oldInnerSession.endCh = endCh
	oldInnerSession.errCH = errCh
	oldInnerSession.inToOut = make(chan<- *SessionResult[ProposalID, *helper.LocalPartySaveData], 1) // todo

	newInnerSession := newSession[ProposalID, *Proposal, *helper.LocalPartySaveData](
		ec,
		m.p2p,
		m,
		helper.SenateSessionID,
		common.Address{}, // todo
		localSubmitter,
		taskID,
		msg,
		ReShareGroupSessionType,
		newPartners,
	)

	// todo
	localData = m.partyData.GenerateNewLocalPartySaveData(ec, newPartners)
	party, endCh, errCh = RunReshare(newInnerSession.ctx, newParams, *localData, reShareSession)
	newInnerSession.party = party
	newInnerSession.partyIdMap = newPartyIdMap
	newInnerSession.endCh = endCh
	newInnerSession.errCH = errCh
	newInnerSession.inToOut = m.senateInToOut

	reShareSession.newSession = newInnerSession
	reShareSession.Run()
	m.AddSession(reShareSession)

	return reShareSession.SessionID()
}

func FromRouteMsg(bytes []byte, routing *tss.MessageRouting) *helper.ReceivedPartyState {
	data := helper.ReceivedPartyState{
		WireBytes:               bytes,
		From:                    routing.From,
		IsBroadcast:             routing.IsBroadcast,
		IsToOldCommittee:        routing.IsToOldCommittee,
		IsToOldAndNewCommittees: routing.IsToOldAndNewCommittees,
	}

	return &data
}

func (r *ReShareGroupSession[T, M, D]) Send(ctx context.Context, bytes []byte, routing *tss.MessageRouting, b bool) error {
	var errs []error

	if r.oldSession.Equal(routing.From.Id) { // from
		log.Debugf("msg from oldSession: %v", routing.From.Id)
		errs = append(errs, r.oldSession.Send(ctx, bytes, routing, b))

		if routing.IsBroadcast || (r.newSession.Included(lo.Map(routing.To, func(item *tss.PartyID, _ int) string { return item.Id }))) {
			r.newSession.Post(FromRouteMsg(bytes, routing))
		}
	}

	if r.newSession.Equal(routing.From.Id) { // from
		log.Debugf("msg from newSession: %v", routing.From.Id)
		errs = append(errs, r.newSession.Send(ctx, bytes, routing, b))

		if routing.IsBroadcast || (r.oldSession.Included(lo.Map(routing.To, func(item *tss.PartyID, _ int) string { return item.Id }))) {
			r.oldSession.Post(FromRouteMsg(bytes, routing))
		}
	}

	return errors.Join(errs...)
}

func (r *ReShareGroupSession[T, M, D]) Receive(partyID string) chan *helper.ReceivedPartyState {
	if r.oldSession.Equal(partyID) {
		return r.oldSession.Receive(partyID)
	}

	return r.oldSession.Receive(partyID)
}

func (r *ReShareGroupSession[T, M, D]) Post(state *helper.ReceivedPartyState) {
	if state.IsBroadcast || state.IsToOldCommittee {
		r.oldSession.Post(state)
	}

	if state.IsBroadcast || state.IsToOldAndNewCommittees {
		r.newSession.Post(state)
	}
}

func (r *ReShareGroupSession[T, M, D]) Release() {
	r.oldSession.Release()
	r.newSession.Release()
}

func (r *ReShareGroupSession[T, M, D]) Name() string {
	return ReShareGroupSessionType
}

func (r *ReShareGroupSession[T, M, D]) Type() string {
	return ReShareGroupSessionType
}

func (r *ReShareGroupSession[T, M, D]) SessionID() helper.SessionID {
	return helper.SenateSessionID
}

func (r *ReShareGroupSession[T, M, D]) GroupID() helper.GroupID {
	return r.newSession.GroupID()
}

func (r *ReShareGroupSession[T, M, D]) ProposalID() T {
	return r.newSession.ProposalID()
}

func (r *ReShareGroupSession[T, M, D]) Proposer() common.Address {
	return r.newSession.Proposer()
}

func (r *ReShareGroupSession[T, M, D]) Threshold() int {
	return r.newSession.Threshold()
}

func (r *ReShareGroupSession[T, M, D]) PartyID(id string) *tss.PartyID {
	pid := r.newSession.PartyID(id)
	if pid != nil {
		return pid
	}

	return r.oldSession.PartyID(id)
}

func (r *ReShareGroupSession[T, M, D]) Equal(id string) bool {
	return r.newSession.Equal(id) || r.oldSession.Equal(id)
}

func (r *ReShareGroupSession[T, M, D]) Included(ids []string) bool {
	return r.newSession.Included(ids) || r.oldSession.Included(ids)
}

func (r *ReShareGroupSession[T, M, D]) Run() {
	r.oldSession.Run()
	r.newSession.Run()
}

func (r *ReShareGroupSession[T, M, D]) Participants() types.Participants {
	return r.newSession.Participants()
}

func (r *ReShareGroupSession[T, M, D]) Signer() common.Address {
	return r.newSession.Signer()
}
