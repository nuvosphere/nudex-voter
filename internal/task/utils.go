package task

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
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

func parseTask(context []byte) (interface{}, error) {
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
