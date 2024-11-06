package task

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func encodeCreateWalletTask(task types.CreateWalletTask) ([]byte, error) {
	bytes := contracts.PackEvent(contracts.TaskPayloadContractMetaData, "WalletCreationRequest", uint32(V1),
		uint32(types.TaskTypeCreateWallet), common.HexToAddress(task.User), task.Account, task.Chain, task.Index)
	return bytes, nil
}
