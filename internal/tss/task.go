package tss

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
