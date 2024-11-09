package task

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func encodeTask(taskType int, task interface{}) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task: %v", r)
			bytes = nil
		}
	}()

	switch taskType {
	case types.TaskTypeCreateWallet:
		t := task.(types.CreateWalletTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WalletCreationRequest",
			uint32(V1),
			uint32(taskType),
			common.HexToAddress(t.User),
			t.Account,
			t.Chain,
			t.Index,
		)

	case types.TaskTypeDeposit:
		t := task.(types.DepositTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"DepositRequest",
			uint32(V1),
			uint32(taskType),
			t.TargetAddress,
			t.Amount,
			t.Chain,
			t.ChainId,
			t.BlockHeight,
			t.TxHash,
			t.ContractAddress,
			t.Ticker,
			t.AssetType,
			t.Decimal,
		)

	case types.TaskTypeWithdrawal:
		t := task.(types.WithdrawalTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WithdrawalRequest",
			uint32(V1),
			uint32(taskType),
			t.TargetAddress,
			t.Amount,
			t.Chain,
			t.ChainId,
			t.BlockHeight,
			t.TxHash,
			t.ContractAddress,
			t.Ticker,
			t.AssetType,
			t.Decimal,
			t.Fee,
		)

	default:
		err = fmt.Errorf("unsupported task type: %v", taskType)
		bytes = nil
	}

	return bytes, err
}

func encodeTaskResult(taskType int, result interface{}) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task result: %v", r)
			bytes = nil
		}
	}()

	switch taskType {
	case types.TaskTypeCreateWallet:
		t := result.(contracts.TaskPayloadContractWalletCreationResult)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WalletCreationResult",
			t.Version,
			t.Success,
			t.ErrorCode,
			t.ErrorMsg,
			t.WalletAddress,
		)
	default:
		err = fmt.Errorf("unsupported task type: %v", taskType)
		bytes = nil
	}

	return bytes, err
}
