package tss

import (
	"context"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
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
	tssErr := party.Start()
	if tssErr != nil && tssErr.Cause() != nil {
		return tssErr.Cause()
	}

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

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(config.AppConfig.TssPublicKeys, tss.Address.Hex())
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)
	messageToSign, err := serializeMsgSignCreateWalletMessageToBytes(e.Task)
	if err != nil {
		return err
	}

	party := signing.NewLocalParty(new(big.Int).SetBytes(messageToSign), params, *tss.LocalPartySaveData, tss.keyOutCh, tss.signEndCh).(*signing.LocalParty)
	tssErr := party.Start()
	if tssErr != nil && tssErr.Cause() != nil {
		return tssErr.Cause()
	}

	tss.sigMu.Lock()
	tss.sigMap[e.RequestId] = make(map[uint64]*signing.LocalParty)
	tss.sigMap[e.RequestId][e.Task.TaskId] = party
	timeoutDuration := config.AppConfig.TssSigTimeout
	tss.sigTimeoutMap[e.RequestId] = time.Now().Add(timeoutDuration)
	tss.sigMu.Unlock()
	return nil
}
