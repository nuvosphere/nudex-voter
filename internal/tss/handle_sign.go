package tss

import (
	"context"
	"errors"
	"fmt"
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
		DataType:    p2p.DataTypeSignCreateWallet,
		Data:        reqMessage,
	}

	err := tss.libp2p.PublishMessage(ctx, p2pMsg)
	if err != nil {
		return err
	}
	log.Debugf("Publish p2p message SignReq: RequestId=%s, task=%v", requestId, task)

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(config.AppConfig.TssPublicKeys, tss.Address.Hex())
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)
	messageToSign, err := serializeMsgSignCreateWalletMessageToBytes(task)
	if err != nil {
		return err
	}

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *tss.LocalPartySaveData, tss.keyOutCh, tss.signEndCh).(*signing.LocalParty)
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("Failed to start sign party: requestId=%s, error=%v", requestId, err)
			return
		} else {
			log.Infof("Sign party started: requestId=%s", requestId)
		}
	}()

	tss.sigMu.Lock()
	tss.sigMap[requestId] = make(map[uint64]*signing.LocalParty)
	tss.sigMap[requestId][task.TaskId] = party
	timeoutDuration := config.AppConfig.TssSigTimeout
	tss.sigTimeoutMap[requestId] = time.Now().Add(timeoutDuration)
	tss.sigMu.Unlock()
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
		if tss.state.TssState.CurrentTask.TaskId > e.Task.TaskId {
			var existingTask db.Task
			result := tss.dbm.GetRelayerDB().Where("task_id = ?", e.Task.TaskId).First(&existingTask)
			if result.Error == nil {
				tss.state.TssState.CurrentTask = &existingTask
			} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return fmt.Errorf("find no task from db for taskId:%d", e.Task.TaskId)
			}
		}
	}

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(config.AppConfig.TssPublicKeys, tss.Address.Hex())
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)
	messageToSign, err := serializeMsgSignCreateWalletMessageToBytes(e.Task)
	if err != nil {
		return err
	}

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *tss.LocalPartySaveData, tss.keyOutCh, tss.signEndCh).(*signing.LocalParty)
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("Failed to start sign party: requestId=%s, error=%v", e.RequestId, err)
			return
		} else {
			log.Infof("Sign party started: requestId=%s", e.RequestId)
		}
	}()

	tss.sigMu.Lock()
	tss.sigMap[e.RequestId] = make(map[uint64]*signing.LocalParty)
	tss.sigMap[e.RequestId][e.Task.TaskId] = party
	timeoutDuration := config.AppConfig.TssSigTimeout
	tss.sigTimeoutMap[e.RequestId] = time.Now().Add(timeoutDuration)
	tss.sigMu.Unlock()
	return nil
}

func (tss *TSSService) handleSigStart(ctx context.Context, event interface{}) {
	switch e := event.(type) {
	case types.MsgSignCreateWalletMessage:
		log.Debugf("Event handleSigStart is of type MsgSignCreateWalletMessage, request id %s", e.RequestId)
		if err := tss.handleSignCreateWalletStart(ctx, e); err != nil {
			log.Errorf("Error handleSigStart MsgSignCreateWalletMessage, %v", err)
			tss.state.EventBus.Publish(state.SigFailed, e)
		}
	default:
		log.Debug("Unknown event handleSigStart type")
	}
}

func (tss *TSSService) handleSigReceive(ctx context.Context, event interface{}) {
}

func (tss *TSSService) handleSigFailed(ctx context.Context, event interface{}, reason string) {
	if e, ok := event.(map[uint64]*signing.LocalParty); ok {
		taskId := tss.state.TssState.CurrentTask.TaskId
		if _, exists := e[taskId]; exists {
			tss.state.TssState.CurrentTask = nil
		}
		for key := range e {
			log.Infof("handle sig failed, taskId:%d, reason:%s", key, reason)
			break
		}
	} else {
		log.Warnf("event is not sign type, actual type: %T", event)
	}
}

func (tss *TSSService) handleSigFinish(ctx context.Context, event interface{}) {
}
