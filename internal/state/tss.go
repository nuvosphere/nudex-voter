package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
)

type TssState struct {
	rw sync.RWMutex

	BlockNumber      uint64         // charge submitter height
	CurrentSubmitter common.Address // charge submitter

	Participants []common.Address // re-share
	CurrentTask  *db.Task         // task
}
