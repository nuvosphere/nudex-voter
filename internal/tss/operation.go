package tss

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/codec"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

const TypeOperations = 100

type Operations struct {
	Nonce     *big.Int
	Operation []contracts.Operation
	Hash      common.Hash
	DataHash  common.Hash
	Signature []byte
}

func (o *Operations) TaskID() uint64 {
	return o.Nonce.Uint64()
}

func (o *Operations) Type() int {
	return TypeOperations
}

func (m *Scheduler) Operation(detailTask pool.Task[uint64]) (*contracts.Operation, error) {
	operation := &contracts.Operation{
		TaskId: detailTask.TaskID(),
	}

	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		ec := types.GetCurveTypeByCoinType(coinType)
		userAddress := wallet.GenerateAddressByPath(m.partyData.GetData(ec).ECPoint(), uint32(coinType), task.Account, task.Index)
		data := m.voterContract.EncodeRegisterNewAddress(big.NewInt(int64(task.Account)), task.Chain, big.NewInt(int64(task.Index)), strings.ToLower(userAddress))
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.AccountContract)
		operation.State = db.Completed
	case *db.DepositTask:
		needConfirm, checkCode, err := m.checkTask(task)
		if !needConfirm {
			return nil, fmt.Errorf("task %d: hash:%s check failed, %w", task.TaskId, task.TxHash, err)
		}
		if err != nil || checkCode != db.TaskErrorCodeSuccess {
			taskResult := contracts.TaskPayloadContractWithdrawalResult{
				Version:   uint8(db.TaskVersionV1),
				Success:   false,
				ErrorCode: uint8(checkCode),
			}
			taskBytes, err := codec.EncodeTaskResult(db.TaskTypeDeposit, taskResult)
			if err != nil {
				panic(fmt.Errorf("encode result failed for task %d: %w", task.TaskId, err))
			}

			data := m.voterContract.EncodeMarkTaskCompleted(new(big.Int).SetUint64(task.TaskId), taskBytes)
			operation.OptData = data
			operation.State = db.Failed
			return operation, err
		}

		data := m.voterContract.EncodeRecordDeposit(
			common.HexToAddress(task.TargetAddress),
			big.NewInt(int64(task.Amount)),
			uint64(task.ChainId),
			common.HexToHash(task.TxHash).Bytes(), // todo
			nil,
		)
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = db.Completed
	case *db.WithdrawalTask:
		needConfirm, checkCode, err := m.checkTask(task)
		if !needConfirm {
			return nil, fmt.Errorf("task %d: hash:%s check failed, %w", task.TaskId, task.TxHash, err)
		}

		if err != nil || checkCode != db.TaskErrorCodeSuccess {
			taskResult := contracts.TaskPayloadContractWithdrawalResult{
				Version:   uint8(db.TaskVersionV1),
				Success:   false,
				ErrorCode: uint8(checkCode),
			}
			taskBytes, err := codec.EncodeTaskResult(db.TaskTypeWithdrawal, taskResult)
			if err != nil {
				panic(fmt.Errorf("encode result failed for task %d: %w", task.TaskId, err))
			}

			data := m.voterContract.EncodeMarkTaskCompleted(new(big.Int).SetUint64(task.TaskId), taskBytes)
			operation.OptData = data
			operation.State = db.Failed
			return operation, err
		}

		taskResult := contracts.TaskPayloadContractWithdrawalResult{
			Version:   uint8(db.TaskVersionV1),
			Success:   false,
			ErrorCode: uint8(db.TaskErrorCodePending),
		}
		taskBytes, err := codec.EncodeTaskResult(db.TaskTypeWithdrawal, taskResult)
		if err != nil {
			return nil, fmt.Errorf("encode result failed for task %d: %w", task.TaskId, err)
		}

		data := m.voterContract.EncodeMarkTaskCompleted(new(big.Int).SetUint64(task.TaskId), taskBytes)
		operation.OptData = data
		operation.State = db.Pending
	default:
		log.Errorf("unhandled default case")
		operation.State = db.Completed
		operation.OptData = nil // todo
	}

	return operation, nil
}
