package tss

import (
	"context"
	"fmt"
	"math/big"
	"reflect"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	tasks "github.com/nuvosphere/nudex-voter/internal/task"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

var taskPrefix = map[reflect.Type]string{
	reflect.TypeOf(&contracts.TaskPayloadContractWalletCreationRequest{}): "CREATE_WALLET",
	reflect.TypeOf(&contracts.TaskPayloadContractDepositRequest{}):        "DEPOSIT",
	reflect.TypeOf(&contracts.TaskPayloadContractWithdrawalRequest{}):     "WITHDRAWAL",
}

func (t *Service) HandleSignCheck(ctx context.Context, dbTask db.Task) (interface{}, *big.Int, []byte, error) {
	task, err := tasks.ParseTask(dbTask.Context)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("parse task %x error: %v", dbTask.Context, err)
	}

	nonce, err := t.layer2Listener.ContractVotingManager().TssNonce(nil)
	if err != nil {
		return task, nonce, nil, fmt.Errorf("get nonce error for task %x, error: %v", dbTask.Context, err)
	}

	switch taskRequest := task.(type) {
	case *contracts.TaskPayloadContractWalletCreationRequest:
		coinType := getCoinTypeByChain(taskRequest.Chain)

		var result contracts.TaskPayloadContractWalletCreationResult

		if coinType == -1 {
			result = contracts.TaskPayloadContractWalletCreationResult{
				Version:       tasks.TaskVersionV1,
				Success:       false,
				ErrorCode:     tasks.TaskErrorCodeChainNotSupported,
				WalletAddress: "",
			}
		} else {
			address := wallet.GenerateAddressByPath(
				*(t.scheduler.MasterPublicKey()),
				uint32(coinType),
				taskRequest.Account,
				taskRequest.Index,
			)
			log.Infof("Generated address: %s for task: %d", address, dbTask.TaskId)
			result = contracts.TaskPayloadContractWalletCreationResult{
				Version:       tasks.TaskVersionV1,
				Success:       true,
				ErrorCode:     tasks.TaskErrorCodeSuccess,
				WalletAddress: address.Hex(),
			}
		}

		resultBytes, err := tasks.EncodeTaskResult(tasks.TaskTypeCreateWallet, result)
		if err != nil {
			return taskRequest, nonce, resultBytes, err
		}

		///todo
		serialized, err := serializeMessageToBeSigned(nonce.Uint64(), resultBytes)

		return taskRequest, nonce, serialized, err
	}

	return task, nonce, nil, err
}

func (t *Service) handleSigFinish(ctx context.Context, signatureData *common.SignatureData) {
	//t.rw.Lock()
	//
	//log.Infof("sig finish, taskId:%d, R:%x, S:%x, V:%x", t.state.TssState.CurrentTask.TaskId, signatureData.R, signatureData.S, signatureData.SignatureRecovery)
	//
	//if t.state.TssState.CurrentTask.Submitter == t.localAddress.Hex() {
	//	buf := bytes.NewReader(t.state.TssState.CurrentTask.Context)
	//
	//		// @todo
	//		// generate wallet and send to chain
	//		address := wallet.GenerateAddressByPath(
	//			*(t.scheduler.MasterPublicKey()),
	//			uint32(coinType),
	//			createWalletTask.Account,
	//			createWalletTask.Index,
	//		)
	//		log.Infof("user account address: %s", address)
	//	}
	//}
	//
	//t.rw.Unlock()
}
