package task

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func encodeCreateWalletTask(task types.CreateWalletTask) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task: %v", r)
			bytes = nil
		}
	}()

	bytes = contracts.PackEvent(
		contracts.TaskPayloadContractMetaData,
		"WalletCreationRequest",
		uint32(V1),
		uint32(types.TaskTypeCreateWallet),
		common.HexToAddress(task.User),
		task.Account,
		task.Chain,
		task.Index,
	)
	return bytes, nil
}

func encodeDepositTask(task types.DepositTask) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to encode task: %v", r)
			bytes = nil
		}
	}()

	bytes = contracts.PackEvent(
		contracts.TaskPayloadContractMetaData,
		"DepositRequest",
		uint32(V1),
		uint32(types.TaskTypeDeposit),
		task.TargetAddress,
		task.Amount,
		task.Chain,
		task.ChainId,
		task.BlockHeight,
		task.TxHash,
		task.ContractAddress,
		task.Ticker,
		task.AssetType,
		task.Decimal,
	)
	return bytes, nil
}
