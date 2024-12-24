package state

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EvmWalletState struct {
	pool *gorm.DB
}

func NewEvmWalletState(pool *gorm.DB) *EvmWalletState {
	return &EvmWalletState{pool: pool}
}

func (d *EvmWalletState) tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}

func (d *EvmWalletState) CreateTx(tx *gorm.DB,
	txHash common.Hash,
	txJsonData []byte,
	txNonce decimal.Decimal,
	buildHeight uint64,
	Operations *db.Operations,
	EvmWithdraw *db.EvmWithdraw,
	EvmConsolidation *db.EvmConsolidation,
) (*db.EvmTransaction, error) {
	tx = d.tx(tx)
	ty := 0
	switch {
	case Operations != nil:
		ty = db.TaskTypeOperations
	case EvmWithdraw != nil:
		ty = db.TaskTypeWithdrawal
	case EvmConsolidation != nil:
		ty = db.TaskTypeConsolidation
	}

	dbTx := &db.EvmTransaction{
		TxHash:           txHash,
		TxJsonData:       txJsonData,
		TxNonce:          txNonce,
		BuildHeight:      buildHeight,
		Status:           db.Created,
		Error:            "",
		Type:             ty,
		Operations:       Operations,
		EvmWithdraw:      EvmWithdraw,
		EvmConsolidation: EvmConsolidation,
	}

	return dbTx, tx.Create(dbTx).Error
}

func (d *EvmWalletState) PendingBlockchainTransaction(tx *gorm.DB, txHash common.Hash) (*db.EvmTransaction, error) {
	tx = d.tx(tx)
	bt := &db.EvmTransaction{}

	err := tx.
		Model(bt).
		Where("tx_hash = ? AND status = ?", txHash, db.Pending).
		Last(&bt).
		Error
	if err != nil {
		return nil, err
	}

	return bt, nil
}

func (d *EvmWalletState) PendingBlockchainTransactions(tx *gorm.DB) (txs []db.EvmTransaction, err error) {
	tx = d.tx(tx)

	err = tx.
		Model(&db.EvmTransaction{}).
		Where("status = ?", db.Pending).
		Find(&txs).
		Error
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (d *EvmWalletState) LatestNonce(tx *gorm.DB, account common.Address) (decimal.Decimal, error) {
	tx = d.tx(tx)
	bt := &db.EvmTransaction{}
	result := tx.
		Model(bt).
		Where("sender = ? AND status IN ?", account, []int{db.Created, db.Pending, db.Completed}).
		Last(&bt)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return decimal.Zero, nil
	}

	return bt.TxNonce, result.Error
}

func (d *EvmWalletState) UpdateTx(tx *gorm.DB, txHash common.Hash, status int, err error) error {
	db := d.tx(tx).
		Model(&db.EvmTransaction{}).
		Where("tx_hash = ? AND status != ?", txHash, db.Completed)
	if err != nil {
		db = db.Updates(&map[string]interface{}{
			"status": status,
			"error":  err.Error(),
		})
	} else {
		db = db.Updates(&map[string]interface{}{
			"status": status,
		})
	}

	return db.Error
}

func (d *EvmWalletState) UpdateFailTx(txHash common.Hash, err error) error {
	return d.UpdateTx(nil, txHash, db.Failed, err)
}

func (d *EvmWalletState) UpdatePendingTx(txHash common.Hash) error {
	return d.UpdateTx(nil, txHash, db.Pending, nil)
}

func (d *EvmWalletState) UpdateBookedTx(txHash common.Hash) error {
	return d.UpdateTx(
		nil,
		txHash,
		db.Completed,
		nil,
	)
}
