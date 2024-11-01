package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
)

// BtcHeadState to manage BTC head
type BtcHeadState struct {
	Latest         db.BtcBlock
	UnconfirmQueue []*db.BtcBlock // status in 'unconfirm', 'confirmed'
	SigQueue       []*db.BtcBlock // status in 'signing', 'pending'
}

type TssState struct {
	BlockNumber      uint64
	CurrentSubmitter common.Address
	Participants     []common.Address

	CurrentTask *db.Task
}
