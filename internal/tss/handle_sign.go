package tss

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var taskPrefix = map[reflect.Type]string{
	reflect.TypeOf(&contracts.TaskPayloadContractWalletCreationRequest{}): "CREATE_WALLET",
	reflect.TypeOf(&contracts.TaskPayloadContractDepositRequest{}):        "DEPOSIT",
	reflect.TypeOf(&contracts.TaskPayloadContractWithdrawalRequest{}):     "WITHDRAWAL",
}

func (t *TSSService) HandleSignCheck(ctx context.Context, dbTask db.Task) (interface{}, *big.Int, []byte, error) {
	task, err := ParseTask(dbTask.Context)
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
				Version:       types.TaskVersionV1,
				Success:       false,
				ErrorCode:     types.TaskErrorCodeChainNotSupported,
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
				Version:       types.TaskVersionV1,
				Success:       true,
				ErrorCode:     types.TaskErrorCodeSuccess,
				WalletAddress: address.Hex(),
			}
		}

		resultBytes, err := utils.EncodeTaskResult(types.TaskTypeCreateWallet, result)
		if err != nil {
			return taskRequest, nonce, resultBytes, err
		}

		serialized, err := serializeMessageToBeSigned(nonce.Uint64(), resultBytes)

		return taskRequest, nonce, serialized, err
	}

	return task, nonce, nil, err
}

func (t *TSSService) HandleSignPrepare(ctx context.Context, dbTask db.Task) error {
	task, nonce, taskResult, err := t.HandleSignCheck(ctx, dbTask)
	if err != nil {
		return err
	}

	var requestId string

	taskType := reflect.TypeOf(task)
	if prefix, found := taskPrefix[taskType]; found {
		requestId = fmt.Sprintf("TSS_SIGN:%s:%d", prefix, dbTask.TaskId)
	} else {
		return fmt.Errorf("task type %T not found in task prefix", task)
	}

	reqMessage := types.SignMessage{
		BaseSignMsg: types.BaseSignMsg{
			RequestId:    requestId,
			IsProposer:   true,
			Nonce:        nonce.Uint64(),
			VoterAddress: t.localAddress.Hex(),
			CreateTime:   time.Now().Unix(),
		},
		Task: types.SignTask{TaskId: dbTask.TaskId, Data: taskResult},
	}

	_ = p2p.Message[types.SignMessage]{
		MessageType: p2p.MessageTypeSigReq,
		RequestId:   requestId,
		DataType:    DataTypeSignCreateWallet,
		Data:        reqMessage,
	}

	return nil
}

func (t *TSSService) handleSignStart(ctx context.Context, e types.SignMessage) error {
	if t.localAddress.Hex() == e.BaseSignMsg.VoterAddress {
		log.Debugf("ignoring sign create wallet start, proposer self")
		return nil
	}

	if t.state.TssState.CurrentTask == nil {
		var existingTask db.Task
		result := t.dbm.GetRelayerDB().Where("task_id = ?", e.Task.TaskId).First(&existingTask)

		if result.Error == nil {
			t.state.TssState.CurrentTask = &existingTask
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("find no task from db for taskId:%d", e.Task.TaskId)
		}
	} else {
		if t.state.TssState.CurrentTask.TaskId > e.Task.TaskId {
			var existingTask db.Task
			result := t.dbm.GetRelayerDB().Where("task_id = ?", e.Task.TaskId).First(&existingTask)

			if result.Error == nil {
				t.state.TssState.CurrentTask = &existingTask
			} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("find no task from db for taskId:%d", e.Task.TaskId)
			}
		} else if t.state.TssState.CurrentTask.TaskId < e.Task.TaskId {
			return fmt.Errorf("new task from p2p: %d is greater than current: %d", e.Task.TaskId, t.state.TssState.CurrentTask.TaskId)
		}
	}

	return nil
}

func (t *TSSService) handleSigFinish(ctx context.Context, signatureData *common.SignatureData) {
	t.rw.Lock()

	log.Infof("sig finish, taskId:%d, R:%x, S:%x, V:%x", t.state.TssState.CurrentTask.TaskId, signatureData.R, signatureData.S, signatureData.SignatureRecovery)

	if t.state.TssState.CurrentTask.Submitter == t.localAddress.Hex() {
		buf := bytes.NewReader(t.state.TssState.CurrentTask.Context)

		var taskType int32
		_ = binary.Read(buf, binary.LittleEndian, &taskType)

		if taskType == types.TaskTypeCreateWallet {
			createWalletTask := types.CreateWalletTask{
				BaseTask: types.BaseTask{
					TaskId: t.state.TssState.CurrentTask.TaskId,
				},
			}

			_ = binary.Read(buf, binary.LittleEndian, &createWalletTask.User)
			_ = binary.Read(buf, binary.LittleEndian, &createWalletTask.Account)
			_ = binary.Read(buf, binary.LittleEndian, &createWalletTask.Chain)

			coinType := getCoinTypeByChain(createWalletTask.Chain)
			if coinType == -1 {
				log.Errorf("chain %d not supported", createWalletTask.Chain)
			}

			// @todo
			// generate wallet and send to chain
			address := wallet.GenerateAddressByPath(
				*(t.scheduler.MasterPublicKey()),
				uint32(coinType),
				createWalletTask.Account,
				createWalletTask.Index,
			)
			log.Infof("user account address: %s", address)
		}
	}

	t.rw.Unlock()
}
