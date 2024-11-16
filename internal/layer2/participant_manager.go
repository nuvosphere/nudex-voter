package layer2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type ParticipantManager interface {
	Participants() types.Participants
	IsParticipant(participant common.Address) bool
	GetRandomParticipant(participant common.Address) common.Address
}

func (l *Layer2Listener) Participants() types.Participants {
	participants, err := l.participantManager.GetParticipants(nil)
	utils.Assert(err)

	return participants
}

func (l *Layer2Listener) IsParticipant(participant common.Address) bool {
	is, err := l.participantManager.IsParticipant(nil, participant)
	utils.Assert(err)

	return is
}

func (l *Layer2Listener) GetRandomParticipant(participant common.Address) common.Address {
	nextSubmitter, err := l.participantManager.GetRandomParticipant(nil, participant)
	utils.Assert(err)

	return nextSubmitter
}
