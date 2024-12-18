package state

import "gorm.io/gorm"

type BtcWalletState struct {
	pool *gorm.DB
}

func NewBtcWalletState(pool *gorm.DB) *BtcWalletState {
	return &BtcWalletState{pool: pool}
}

func (d *BtcWalletState) tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}
