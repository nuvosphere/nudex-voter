package state

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
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

func (s *ContractState) GetAddressBalance(address, token string) (decimal.Decimal, error) {
	var balance decimal.Decimal
	err := s.l2InfoDb.
		Model(&db.AddressBalance{}).
		Where("address = ? AND token = ?", address, token).
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

func (s *ContractState) GetUnCompletedTasks() (tasks []db.Task, err error) {
	err = s.l2InfoDb.
		Preload(clause.Associations).
		Where("state in ?", []int{db.Created, db.Pending}).
		Last(&db.Task{}).
		Error
	return tasks, err
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

func (s *ContractState) GetInscriptionMintb(txHash string) (*db.InscriptionMintb, error) {
	var inscriptionMintB db.InscriptionMintb

	txHashBytes := common.HexToHash(txHash)
	result := s.l2InfoDb.
		Preload("LogIndex", "tx_hash = ?", txHashBytes).
		First(&inscriptionMintB)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("inscriptionMintB not found for TxHash: %s", txHash)
		}
		return nil, fmt.Errorf("failed to query inscriptionMintB: %w", result.Error)
	}

	return &inscriptionMintB, nil
}

func (s *ContractState) GetInscriptionBurnb(txHash string) (*db.InscriptionBurnb, error) {
	var inscriptionBurnb db.InscriptionBurnb

	txHashBytes := common.HexToHash(txHash)
	result := s.l2InfoDb.
		Preload("LogIndex", "tx_hash = ?", txHashBytes).
		First(&inscriptionBurnb)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("inscriptionBurnb not found for TxHash: %s", txHash)
		}
		return nil, fmt.Errorf("failed to query inscriptionBurnb: %w", result.Error)
	}

	return &inscriptionBurnb, nil
}

func (s *ContractState) GetAsset(ticker string) (*db.Asset, error) {
	var asset db.Asset

	result := s.l2InfoDb.
		Preload(clause.Associations).
		Where("ticker = ?", ticker).
		First(asset)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("asset not found for ticker: %s", ticker)
		}
		return nil, fmt.Errorf("failed to query asset for ticker: %s,  %w", ticker, result.Error)
	}

	return &asset, nil
}

func (s *ContractState) GetTokenInfo(ticker string, chainId vtypes.Byte32) (*db.TokenInfo, error) {
	var tokenInfo db.TokenInfo

	result := s.l2InfoDb.
		Preload(clause.Associations).
		Where("ticker = ? and chain_id = ?", ticker, chainId).
		First(tokenInfo)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("token info not found for ticker: %s, chainId: %d", ticker, chainId)
		}
		return nil, fmt.Errorf("failed to query asset for ticker: %s, chainId: %d, %w", ticker, chainId, result.Error)
	}

	return &tokenInfo, nil
}
