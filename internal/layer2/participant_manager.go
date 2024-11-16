package layer2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

func (l *Layer2Listener) Participants() []common.Address {
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
