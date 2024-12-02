package tss

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
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

func (m *Scheduler) Operation(detailTask pool.Task[uint64]) *contracts.Operation {
	operation := &contracts.Operation{
		TaskId: detailTask.TaskID(),
	}

	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := getCoinTypeByChain(task.Chain)

		userAddress := wallet.GenerateAddressByPath(*m.partyData.ECDSALocalData().ECDSAData().ECDSAPub.ToECDSAPubKey(), uint32(coinType), task.Account, task.Index)
		data := m.voterContract.EncodeRegisterNewAddress(big.NewInt(int64(task.Account)), task.Chain, big.NewInt(int64(task.Index)), userAddress.Hex())
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.AccountContract)
		operation.State = db.Completed
	case *db.DepositTask:
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
		data := m.voterContract.EncodeRecordWithdrawal(
			common.HexToAddress(task.TargetAddress),
			big.NewInt(int64(task.Amount)),
			uint64(task.ChainId),
			common.HexToHash(task.TxHash).Bytes(), // todo
			nil,
		)
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = db.Pending
	default:
		log.Errorf("unhandled default case")
		operation.State = db.Completed
		operation.OptData = nil // todo
	}

	return operation
}
