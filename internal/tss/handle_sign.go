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
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var taskPrefix = map[reflect.Type]string{
	reflect.TypeOf(&types.CreateWalletTask{}): "CREATE_WALLET",
	reflect.TypeOf(&types.DepositTask{}):      "DEPOSIT",
	reflect.TypeOf(&types.WithdrawalTask{}):   "WITHDRAWAL",
}

func (t *TSSService) HandleSignPrepare(ctx context.Context, task types.Task) error {
	var requestId string

	taskType := reflect.TypeOf(task)
	if prefix, found := taskPrefix[taskType]; found {
		requestId = fmt.Sprintf("TSS_SIGN:%s:%d", prefix, task.GetTaskID())
	} else {
		return fmt.Errorf("task type %T not found in task prefix", task)
	}

	nonce, err := t.layer2Listener.ContractVotingManager().TssNonce(nil)
	if err != nil {
		return err
	}

	reqMessage := types.SignMessage{
		BaseSignMsg: types.BaseSignMsg{
			RequestId:    requestId,
			IsProposer:   true,
			Nonce:        nonce.Uint64(),
			VoterAddress: t.localAddress.Hex(),
			CreateTime:   time.Now().Unix(),
		},
		Task: task,
	}

	_ = p2p.Message[types.SignMessage]{
		MessageType: p2p.MessageTypeSigReq,
		RequestId:   requestId,
		DataType:    DataTypeSignCreateWallet,
		Data:        reqMessage,
	}

	messageToSign, err := serializeTaskMessageToBytes(nonce.Uint64(), task)
	if err != nil {
		return err
	}

	msg := new(big.Int).SetBytes(messageToSign)

	log.Debug(msg)
	// todo
	return nil
}

func (t *TSSService) handleSignStart(ctx context.Context, e types.SignMessage) error {
	if t.localAddress.Hex() == e.BaseSignMsg.VoterAddress {
		log.Debugf("ignoring sign create wallet start, proposer self")
		return nil
	}

	if t.state.TssState.CurrentTask == nil {
		var existingTask db.Task
		result := t.dbm.GetRelayerDB().Where("task_id = ?", e.Task.GetTaskID()).First(&existingTask)

		if result.Error == nil {
			t.state.TssState.CurrentTask = &existingTask
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("find no task from db for taskId:%d", e.Task.GetTaskID())
		}
	} else {
		if t.state.TssState.CurrentTask.TaskId > uint64(e.Task.GetTaskID()) {
			var existingTask db.Task
			result := t.dbm.GetRelayerDB().Where("task_id = ?", e.Task.GetTaskID()).First(&existingTask)

			if result.Error == nil {
				t.state.TssState.CurrentTask = &existingTask
			} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("find no task from db for taskId:%d", e.Task.GetTaskID())
			}
		} else if t.state.TssState.CurrentTask.TaskId < uint64(e.Task.GetTaskID()) {
			return fmt.Errorf("new task from p2p: %d is greater than current: %d", e.Task.GetTaskID(), t.state.TssState.CurrentTask.TaskId)
		}
	}

	messageToSign, err := serializeTaskMessageToBytes(e.Nonce, e.Task)
	if err != nil {
		return err
	}

	msg := new(big.Int).SetBytes(messageToSign)
	log.Debug(msg)
	// todo

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
					TaskId: int32(t.state.TssState.CurrentTask.TaskId),
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
				*(t.localPartySaveData.ECDSAPub.ToECDSAPubKey()),
				uint32(coinType),
				uint32(createWalletTask.Account),
				createWalletTask.Index,
			)
			log.Infof("user account address: %s", address)
		}
	}

	t.rw.Unlock()
}
