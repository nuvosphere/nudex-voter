package state

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	account common.Address,
	txNonce decimal.Decimal,
	txJsonData, calldata []byte,
	txHash common.Hash,
	buildHeight uint64,
	Operations *db.Operations,
	EvmWithdraw *db.EvmWithdraw,
	EvmConsolidation *db.EvmConsolidation,
) error {
	tx = d.tx(tx)
	return tx.Create(&db.EvmTransaction{
		CalldataHash:     crypto.Keccak256Hash(calldata),
		Calldata:         calldata,
		TxNonce:          txNonce,
		TxHash:           txHash,
		TxJsonData:       txJsonData,
		Sender:           account,
		BuildHeight:      buildHeight,
		Status:           db.Created,
		Error:            "",
		Type:             0, // todo
		Operations:       Operations,
		EvmWithdraw:      EvmWithdraw,
		EvmConsolidation: EvmConsolidation,
	}).Error
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

func (d *EvmWalletState) LatestNonce(tx *gorm.DB, account common.Address) (decimal.Decimal, error) {
	tx = d.tx(tx)
	bt := &db.EvmTransaction{}
	result := tx.
		Model(bt).
		Where("sender = ? AND status IN ?", account, []int{db.Created, db.Pending}).
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
