package state

import "github.com/nuvosphere/nudex-voter/internal/db"

// BtcHeadState to manage BTC head
type BtcHeadState struct {
	Latest         db.BtcBlock
	UnconfirmQueue []*db.BtcBlock // status in 'unconfirm', 'confirmed'
	SigQueue       []*db.BtcBlock // status in 'signing', 'pending'
}
