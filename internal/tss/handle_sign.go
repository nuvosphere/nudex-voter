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
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
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

func (t *TSSService) HandleSignCheck(ctx context.Context, dbTask db.Task, task interface{}) ([]byte, error) {
	return nil, nil
}

func (t *TSSService) HandleSignPrepare(ctx context.Context, dbTask db.Task, task interface{}) error {
	taskResult, err := t.HandleSignCheck(ctx, dbTask, task)
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

	nonce, err := t.layer2Listener.ContractVotingManager().TssNonce(nil)
	if err != nil {
		return err
	}

	messageToSign, err := serializeMessageToBeSigned(nonce.Uint64(), taskResult)
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
		Task: types.SignTask{TaskId: dbTask.TaskId, Data: taskResult},
	}

	p2pMsg := p2p.Message[types.SignMessage]{
		MessageType: p2p.MessageTypeSigReq,
		RequestId:   requestId,
		DataType:    DataTypeSignCreateWallet,
		Data:        reqMessage,
	}

	err = t.p2p.PublishMessage(ctx, p2pMsg)
	if err != nil {
		return err
	}

	log.Debugf("Publish p2p message SignReq: RequestId=%s, task=%v", requestId, task)

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(t.partners, t.localAddress)
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *t.LocalPartySaveData, t.sigOutCh, t.sigEndCh).(*signing.LocalParty)
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("Failed to start sign party: requestId=%s, error=%v", requestId, err)
			return
		} else {
			log.Infof("Sign party started: requestId=%s", requestId)
		}
	}()

	t.rw.Lock()
	t.sigMap[requestId] = make(map[uint32]*signing.LocalParty)
	t.sigMap[requestId][dbTask.TaskId] = party
	timeoutDuration := config.AppConfig.TssSigTimeout
	t.sigTimeoutMap[requestId] = time.Now().Add(timeoutDuration)
	t.rw.Unlock()

	return nil
}

func (t *TSSService) handleSignStart(ctx context.Context, e types.SignMessage) error {
	if t.localAddress.Hex() == e.BaseSignMsg.VoterAddress {
		log.Debugf("ignoring sign create wallet start, proposer self")
		return nil
	}

	// check map
	_, ok := t.sigExists(e.RequestId)
	if ok {
		return fmt.Errorf("sig exists: %s", e.RequestId)
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

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(t.partners, t.localAddress)
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)

	messageToSign, err := serializeMessageToBeSigned(e.Nonce, e.Task.Data)
	if err != nil {
		return err
	}

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *t.LocalPartySaveData, t.sigOutCh, t.sigEndCh).(*signing.LocalParty)
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("Failed to start sign party: requestId=%s, error=%v", e.RequestId, err)
			return
		} else {
			log.Infof("Sign party started: requestId=%s", e.RequestId)
		}
	}()

	t.rw.Lock()
	t.sigMap[e.RequestId] = make(map[uint32]*signing.LocalParty)
	t.sigMap[e.RequestId][e.Task.TaskId] = party
	t.sigTimeoutMap[e.RequestId] = time.Now().Add(config.AppConfig.TssSigTimeout)
	t.rw.Unlock()

	return nil
}

func (t *TSSService) handleSigStart(ctx context.Context, event interface{}) {
	if signMsg, ok := event.(types.SignMessage); ok {
		log.Debugf("Event handleSigStart is of type SignMessage, request id %s", signMsg.RequestId)

		if err := t.handleSignStart(ctx, signMsg); err != nil {
			log.Errorf("Error handleSigStart handleSignStart, %v", err)
			t.state.EventBus.Publish(state.EventSigFailed{}, event)
		}
	} else {
		log.Errorf("HandleSigStart error: event is not of type types.SignMessage")
	}
}

func (t *TSSService) handleSigFailed(ctx context.Context, event interface{}, reason string) {
	log.Infof("sig failed, taskId:%d, reason:%s", t.state.TssState.CurrentTask.TaskId, reason)
	t.cleanAllSigInfo()
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
				*(t.LocalPartySaveData.ECDSAPub.ToECDSAPubKey()),
				uint32(coinType),
				uint32(createWalletTask.Account),
				createWalletTask.Index,
			)
			log.Infof("user account address: %s", address)
		}
	}

	t.cleanAllSigInfo()

	t.rw.Unlock()
}
