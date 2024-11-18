package tss

import (
	"context"
	"errors"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
)

var _ Session[any] = &ReShareGroupSession[any, any, any]{}

type ReShareGroupSession[T, M, D any] struct {
	oldSession *sessionTransport[T, M, D]
	newSession *sessionTransport[T, M, D]
}

func (m *Scheduler) NewReShareGroupSession(
	taskID ProposalID, // msg id
	msg *Proposal,
	oldPartners types.Participants,
	newPartners types.Participants,
) helper.SessionID {
	proposer := m.Proposer()
	reShareSession := &ReShareGroupSession[ProposalID, *Proposal, *keygen.LocalPartySaveData]{}

	oldPartyIDs := createOldPartyIDsByAddress(oldPartners)
	oldPeerCtx := tss.NewPeerContext(oldPartyIDs)
	oldPartyID := oldPartyIDs.FindByKey(m.LocalSubmitter().Big())
	oldPartyIdMap := lo.SliceToMap(oldPartyIDs, func(item *tss.PartyID) (string, *tss.PartyID) {
		return item.Id, item
	})

	newPartyIDs := createPartyIDsByAddress(newPartners)
	newPeerCtx := tss.NewPeerContext(newPartyIDs)
	newPartyID := newPartyIDs.FindByKey(m.LocalSubmitter().Big())
	newPartyIdMap := lo.SliceToMap(newPartyIDs, func(item *tss.PartyID) (string, *tss.PartyID) {
		return item.Id, item
	})

	oldParams := tss.NewReSharingParameters(
		tss.S256(),
		oldPeerCtx,
		newPeerCtx,
		oldPartyID,
		oldPartners.Len(),
		oldPartners.Threshold(),
		newPartners.Len(),
		newPartners.Threshold(),
	)

	oldInnerSession := newSession[ProposalID, *Proposal, *keygen.LocalPartySaveData](
		m.p2p,
		m,
		helper.SenateSessionID,
		m.MasterSigner(),
		proposer,
		taskID,
		msg,
		ReShareGroupSessionType,
		oldPartners,
	)
	reShareSession.oldSession = oldInnerSession

	party, endCh, errCh := helper.RunReshare(m.ctx, oldParams, m.masterLocalPartySaveData, reShareSession) // todo
	oldInnerSession.party = party
	oldInnerSession.partyIdMap = oldPartyIdMap
	oldInnerSession.endCh = endCh
	oldInnerSession.errCH = errCh
	oldInnerSession.inToOut = make(chan<- *SessionResult[ProposalID, *keygen.LocalPartySaveData], 1) // todo

	newParams := tss.NewReSharingParameters(
		tss.S256(),
		oldPeerCtx,
		newPeerCtx,
		newPartyID,
		oldPartners.Len(),
		oldPartners.Threshold(),
		newPartners.Len(),
		newPartners.Threshold(),
	)

	newInnerSession := newSession[ProposalID, *Proposal, *keygen.LocalPartySaveData](
		m.p2p,
		m,
		helper.SenateSessionID,
		m.MasterSigner(),
		proposer,
		taskID,
		msg,
		ReShareGroupSessionType,
		newPartners,
	)
	reShareSession.newSession = newInnerSession

	party, endCh, errCh = helper.RunReshare(m.ctx, newParams, m.masterLocalPartySaveData, reShareSession)
	newInnerSession.party = party
	newInnerSession.partyIdMap = newPartyIdMap
	newInnerSession.endCh = endCh
	newInnerSession.errCH = errCh
	newInnerSession.inToOut = m.senateInToOut

	reShareSession.Run()
	m.AddSession(reShareSession)

	return reShareSession.SessionID()
}

func From(bytes []byte, routing *tss.MessageRouting) *helper.ReceivedPartyState {
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
	if r.oldSession.Equal(routing.From.Id) {
		errs = append(errs, r.oldSession.Send(ctx, bytes, routing, b))

		if routing.IsBroadcast || (r.newSession.Included(lo.Map(routing.To, func(item *tss.PartyID, _ int) string { return item.Id }))) {
			r.newSession.Post(From(bytes, routing))
		}
	}

	if r.newSession.Equal(routing.From.Id) {
		errs = append(errs, r.newSession.Send(ctx, bytes, routing, b))

		if routing.IsBroadcast || (r.oldSession.Included(lo.Map(routing.To, func(item *tss.PartyID, _ int) string { return item.Id }))) {
			r.oldSession.Post(From(bytes, routing))
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
