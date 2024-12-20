package state

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContractState struct {
	l2InfoDb *gorm.DB
}

func NewContractState(l2InfoDb *gorm.DB) *ContractState {
	return &ContractState{
		l2InfoDb: l2InfoDb,
	}
}

func (s *ContractState) AccountByChainAndAddress(chain uint8, address string) (*db.Account, error) {
	account := &db.Account{}
	err := s.l2InfoDb.
		Preload(clause.Associations).
		Where("chain = ? AND address = ?", chain, address).
		Last(account).
		Error
	return account, err
}

func (s *ContractState) AccountByChain(chain uint8) ([]db.Account, error) {
	var accounts []db.Account
	err := s.l2InfoDb.
		Preload(clause.Associations).
		Where("chain = ?", chain).
		Last(accounts).
		Error
	return accounts, err
}

func (s *ContractState) GetAddressBalance(address string) (decimal.Decimal, error) {
	var balance decimal.Decimal
	err := s.l2InfoDb.
		Model(&db.AddressBalance{}).
		Where("address = ?", address).
		Select("SUM(amount)").
		Scan(&balance).
		Error
	return balance, err
}

func (s *ContractState) GetAddressBalanceByCondition(chainId uint64, minAmount uint64) ([]db.AddressBalance, error) {
	var balances []db.AddressBalance
	err := s.l2InfoDb.
		Model(&db.AddressBalance{}).
		Where("chain_id = ? AND amount >= ?", chainId, decimal.NewFromUint64(minAmount)).
		Find(&balances).
		Error
	return balances, err
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

func (s *ContractState) GetUnCompletedTask(taskID uint64) (*db.Task, error) {
	task := &db.Task{}
	err := s.l2InfoDb.
		Preload(clause.Associations).
		Where("task_id = ? and state in ?", taskID, []int{db.Created, db.Pending}).
		Last(task).
		Error
	return task, err
}

func (s *ContractState) GetCreatedTask() (tasks []db.Task, err error) {
	return s.GetTaskByStatus(db.Created)
}

func (s *ContractState) GetPendingTask() (tasks []db.Task, err error) {
	return s.GetTaskByStatus(db.Pending)
}

func (s *ContractState) GetTaskByStatus(status int) (tasks []db.Task, err error) {
	err = s.l2InfoDb.
		Preload(clause.Associations).
		Where("status = ?", status).
		First(tasks).
		Error
	return tasks, err
}
