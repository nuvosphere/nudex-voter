package tss

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (tss *TSSService) HandleSignCreateAccount(ctx context.Context, task types.CreateWalletTask) error {
	requestId := fmt.Sprintf("TSS_SIGN:CREATE_WALLET:%d", task.TaskId)
	reqMessage := types.MsgSignCreateWalletMessage{
		MsgSign: types.MsgSign{
			RequestId:    requestId,
			IsProposer:   true,
			VoterAddress: tss.Address.Hex(),
			CreateTime:   time.Now().Unix(),
		},
		Task: task,
	}

	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeSigReq,
		RequestId:   requestId,
		DataType:    DataTypeSignCreateWallet,
		Data:        reqMessage,
	}

	err := tss.p2p.PublishMessage(ctx, p2pMsg)
	if err != nil {
		return err
	}
	log.Debugf("Publish p2p message SignReq: RequestId=%s, task=%v", requestId, task)

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(tss.addressList, tss.Address)
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)
	messageToSign, err := serializeMsgSignCreateWalletMessageToBytes(task)
	if err != nil {
		return err
	}

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *tss.LocalPartySaveData, tss.sigOutCh, tss.sigEndCh).(*signing.LocalParty)
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("Failed to start sign party: requestId=%s, error=%v", requestId, err)
			return
		} else {
			log.Infof("Sign party started: requestId=%s", requestId)
		}
	}()

	tss.rw.Lock()
	tss.sigMap[requestId] = make(map[int32]*signing.LocalParty)
	tss.sigMap[requestId][task.TaskId] = party
	timeoutDuration := config.AppConfig.TssSigTimeout
	tss.sigTimeoutMap[requestId] = time.Now().Add(timeoutDuration)
	tss.rw.Unlock()
	return nil
}

func (tss *TSSService) handleSignCreateWalletStart(ctx context.Context, e types.MsgSignCreateWalletMessage) error {
	if tss.Address.Hex() == e.MsgSign.VoterAddress {
		log.Debugf("ignoring sign create wallet start, proposer self")
		return nil
	}

	// check map
	_, ok := tss.sigExists(e.RequestId)
	if ok {
		return fmt.Errorf("sig exists: %s", e.RequestId)
	}

	if tss.state.TssState.CurrentTask == nil {
		var existingTask db.Task
		result := tss.dbm.GetRelayerDB().Where("task_id = ?", e.Task.TaskId).First(&existingTask)

		if result.Error == nil {
			tss.state.TssState.CurrentTask = &existingTask
		} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("find no task from db for taskId:%d", e.Task.TaskId)
		}
	} else {
		if tss.state.TssState.CurrentTask.TaskId > uint64(e.Task.TaskId) {
			var existingTask db.Task
			result := tss.dbm.GetRelayerDB().Where("task_id = ?", e.Task.TaskId).First(&existingTask)
			if result.Error == nil {
				tss.state.TssState.CurrentTask = &existingTask
			} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("find no task from db for taskId:%d", e.Task.TaskId)
			}
		} else if tss.state.TssState.CurrentTask.TaskId < uint64(e.Task.TaskId) {
			return fmt.Errorf("new task from p2p: %d is greater than current: %d", e.Task.TaskId, tss.state.TssState.CurrentTask.TaskId)
		}
	}

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(tss.addressList, tss.Address)
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)
	messageToSign, err := serializeMsgSignCreateWalletMessageToBytes(e.Task)
	if err != nil {
		return err
	}

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *tss.LocalPartySaveData, tss.sigOutCh, tss.sigEndCh).(*signing.LocalParty)
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("Failed to start sign party: requestId=%s, error=%v", e.RequestId, err)
			return
		} else {
			log.Infof("Sign party started: requestId=%s", e.RequestId)
		}
	}()

	tss.rw.Lock()
	tss.sigMap[e.RequestId] = make(map[int32]*signing.LocalParty)
	tss.sigMap[e.RequestId][e.Task.TaskId] = party
	tss.sigTimeoutMap[e.RequestId] = time.Now().Add(config.AppConfig.TssSigTimeout)
	tss.rw.Unlock()
	return nil
}

func (tss *TSSService) handleSigStart(ctx context.Context, event interface{}) {
	switch e := event.(type) {
	case types.MsgSignCreateWalletMessage:
		log.Debugf("Event handleSigStart is of type MsgSignCreateWalletMessage, request id %s", e.RequestId)
		if err := tss.handleSignCreateWalletStart(ctx, e); err != nil {
			log.Errorf("Error handleSigStart MsgSignCreateWalletMessage, %v", err)
			tss.state.EventBus.Publish(state.EventSigFailed{}, e)
		}
	default:
		log.Debug("Unknown event handleSigStart type")
	}
}

func (tss *TSSService) handleSigReceive(ctx context.Context, event interface{}) {
}

func (tss *TSSService) handleSigFailed(ctx context.Context, event interface{}, reason string) {
	log.Infof("sig failed, taskId:%d, reason:%s", tss.state.TssState.CurrentTask.TaskId, reason)
	tss.CleanAllSigInfo()
}

func (tss *TSSService) handleSigFinish(ctx context.Context, signatureData *common.SignatureData) {
	tss.rw.Lock()

	log.Infof("sig finish, taskId:%d, R:%x, S:%x, V:%x", tss.state.TssState.CurrentTask.TaskId, signatureData.R, signatureData.S, signatureData.SignatureRecovery)
	if tss.state.TssState.CurrentTask.Submitter == tss.Address.Hex() {
		buf := bytes.NewReader(tss.state.TssState.CurrentTask.Context)
		var taskType int32
		_ = binary.Read(buf, binary.LittleEndian, &taskType)

		if taskType == types.TaskTypeCreateWallet {
			createWalletTask := types.CreateWalletTask{TaskId: int32(tss.state.TssState.CurrentTask.TaskId)}

			_ = binary.Read(buf, binary.LittleEndian, &createWalletTask.User)
			_ = binary.Read(buf, binary.LittleEndian, &createWalletTask.Account)
			_ = binary.Read(buf, binary.LittleEndian, &createWalletTask.Chain)

			coinType := getCoinTypeByChain(createWalletTask.Chain)
			if coinType == -1 {
				log.Errorf("chain %d not supported", createWalletTask.Chain)
			}

			bip44Path := fmt.Sprintf("m/44'/%d'/%d'/0/%d", coinType, createWalletTask.User, createWalletTask.Account)
			log.Infof("bip44Path: %s", bip44Path)
			// @todo
			// generate wallet and send to chain
		}

	}
	tss.CleanAllSigInfo()

	tss.rw.Unlock()
}
