package state

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletEvmState struct {
	pool *gorm.DB
}

func (d *WalletEvmState) tx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		tx = d.pool
	}

	return tx
}

func (d *WalletEvmState) CreateTx(tx *gorm.DB,
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

func (d *WalletEvmState) PendingBlockchainTransaction(tx *gorm.DB, txHash common.Hash) (*db.EvmTransaction, error) {
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

func (d *WalletEvmState) LatestNonce(tx *gorm.DB, account common.Address) (decimal.Decimal, error) {
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

func (d *WalletEvmState) UpdateTx(tx *gorm.DB, txHash common.Hash, status int, err error) error {
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

func (d *WalletEvmState) UpdateFailTx(txHash common.Hash, err error) error {
	return d.UpdateTx(nil, txHash, db.Failed, err)
}

func (d *WalletEvmState) UpdatePendingTx(txHash common.Hash) error {
	return d.UpdateTx(nil, txHash, db.Pending, nil)
}

func (d *WalletEvmState) UpdateBookedTx(txHash common.Hash) error {
	return d.UpdateTx(
		nil,
		txHash,
		db.Completed,
		nil,
	)
}

type ContractState struct {
	l2InfoDb *gorm.DB
}

func NewContractState(l2InfoDb *gorm.DB) *ContractState {
	return &ContractState{
		l2InfoDb: l2InfoDb,
	}
}

func (s *ContractState) Account(address string) (*db.Account, error) {
	account := &db.Account{}
	err := s.l2InfoDb.
		Preload(clause.Associations).
		Where("address = ?", address).
		Last(account).
		Error
	return account, err
}

func (s *ContractState) Task(taskID uint64) (*db.Task, error) {
	task := &db.Task{}
	err := s.l2InfoDb.
		Preload(clause.Associations).
		Where("task_id", taskID).
		Last(task).
		Error
	return task, err
}

func (s *ContractState) GetPendingTask() (tasks []db.Task, err error) {
	err = s.l2InfoDb.
		Preload(clause.Associations).
		Where("status = ?", db.Pending).
		First(tasks).
		Error
	return tasks, err
}
