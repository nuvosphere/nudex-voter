package task

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
)

func ParseTask(context []byte) (interface{}, error) {
	parsedABI, err := contracts.ParseABI(contracts.TaskPayloadContractMetaData.ABI)
	if err != nil {
		return nil, err
	}

	eventHash := common.BytesToHash(context[:32])
	switch eventHash {
	case layer2.WalletCreationRequestTopic:
		request := contracts.TaskPayloadContractWalletCreationRequest{}
		err = parsedABI.UnpackIntoInterface(&request, "WalletCreationRequest", context[32:])

		return request, err
	case layer2.DepositRequestTopic:
		request := contracts.TaskPayloadContractDepositRequest{}
		err = parsedABI.UnpackIntoInterface(&request, "DepositRequest", context[32:])

		return request, err
	case layer2.WithdrawalRequestTopic:
		request := contracts.TaskPayloadContractWithdrawalRequest{}
		err = parsedABI.UnpackIntoInterface(&request, "WithdrawalRequest", context[32:])

		return request, err
	}

	return nil, fmt.Errorf("unknown task type: %v", eventHash)
}

func encodeTask(taskType uint8, task interface{}) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task: %v", r)
			bytes = nil
		}
	}()

	switch taskType {
	case TaskTypeCreateWallet:
		t := task.(CreateWalletTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WalletCreationRequest",
			uint8(TaskVersionV1),
			taskType,
			common.HexToAddress(t.User),
			t.Account,
			t.Chain,
			t.Index,
		)

	case TaskTypeDeposit:
		t := task.(DepositTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"DepositRequest",
			uint8(TaskVersionV1),
			taskType,
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

	case TaskTypeWithdrawal:
		t := task.(WithdrawalTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WithdrawalRequest",
			uint8(TaskVersionV1),
			taskType,
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

func EncodeTaskResult(taskType uint8, result interface{}) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task result: %v", r)
			bytes = nil
		}
	}()

	switch taskType {
	case TaskTypeCreateWallet:
		t := result.(contracts.TaskPayloadContractWalletCreationResult)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WalletCreationResult",
			t.Version,
			t.Success,
			t.ErrorCode,
			t.WalletAddress,
		)
	case TaskTypeDeposit:
		t := result.(contracts.TaskPayloadContractDepositResult)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"DepositResult",
			t.Version,
			t.Success,
			t.ErrorCode,
		)
	case TaskTypeWithdrawal:
		t := result.(contracts.TaskPayloadContractWithdrawalResult)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WithdrawalResult",
			t.Version,
			t.Success,
			t.ErrorCode,
		)
	default:
		err = fmt.Errorf("unsupported task type: %v", taskType)
		bytes = nil
	}

	return bytes, err
}
