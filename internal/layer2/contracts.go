package layer2

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

type VoterContract interface {
	IsSyncing() bool
	ContractVotingManager
	ParticipantManager
	TaskManager
	AccountManager
	DepositManager
}

type Operation interface {
	Operation(detail db.DetailTask) *contracts.TaskOperation
}
