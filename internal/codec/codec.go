package codec

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type WalletCreationRequest struct {
	Account     uint32
	AddressType uint8
	Index       uint8
}

func DecodeTask(taskId uint64, context []byte) db.DetailTask {
	parsedABI, err := contracts.ParseABI(contracts.TaskPayloadContractMetaData.ABI)
	utils.Assert(err)

	eventHash := common.BytesToHash(context[:32])
	switch eventHash {
	case contracts.WalletCreationRequestTopic:
		request := &contracts.TaskPayloadContractWalletCreationRequest{}
		err = parsedABI.UnpackIntoInterface(request, "WalletCreationRequest", context[32:])
		utils.Assert(err)

		return db.NewCreateWalletTask(taskId, request)
	case contracts.DepositRequestTopic:
		request := &contracts.TaskPayloadContractDepositRequest{}
		err = parsedABI.UnpackIntoInterface(request, "DepositRequest", context[32:])
		utils.Assert(err)

		return db.NewDepositTask(taskId, request)
	case contracts.WithdrawalRequestTopic:
		request := &contracts.TaskPayloadContractWithdrawalRequest{}
		err = parsedABI.UnpackIntoInterface(request, "WithdrawalRequest", context[32:])
		utils.Assert(err)

		return db.NewWithdrawalTask(taskId, request)
	}

	return nil
}

func EncodeTask(taskType uint8, task any) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task: %v", r)
			bytes = nil
		}
	}()

	switch taskType {
	case db.TaskTypeCreateWallet:
		t := task.(db.CreateWalletTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WalletCreationRequest",
			uint8(db.TaskVersionV1),
			taskType,
			t.Account,
			t.ChainType(),
			t.Index,
		)

	case db.TaskTypeDeposit:
		t := task.(db.DepositTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"DepositRequest",
			uint8(db.TaskVersionV1),
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

	case db.TaskTypeWithdrawal:
		t := task.(db.WithdrawalTask)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WithdrawalRequest",
			uint8(db.TaskVersionV1),
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
	case db.TaskTypeCreateWallet:
		t := result.(contracts.TaskPayloadContractWalletCreationResult)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"WalletCreationResult",
			t.Version,
			t.Success,
			t.ErrorCode,
			t.WalletAddress,
		)
	case db.TaskTypeDeposit:
		t := result.(contracts.TaskPayloadContractDepositResult)
		bytes = contracts.PackEvent(
			contracts.TaskPayloadContractMetaData,
			"DepositResult",
			t.Version,
			t.Success,
			t.ErrorCode,
		)
	case db.TaskTypeWithdrawal:
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
