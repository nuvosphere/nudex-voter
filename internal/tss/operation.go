package tss

import (
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
)

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
	case *db.DepositTask:
	case *db.WithdrawalTask:
	default:
		panic("unhandled default case")
	}

	return operation
}
