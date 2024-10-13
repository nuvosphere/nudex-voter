package tss

import (
	"context"
	"fmt"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (tss *TSSService) handleKeygenReq(ctx context.Context, event interface{}) error {
	message, ok := event.(types.KeygenReqMessage)
	if !ok {
		return fmt.Errorf("handleKeygenReq error, event %v, is not keygen req", event)
	}

	if tss.address.Hex() == message.VoterAddress {
		log.Debugf("Ignore handleKeygenReq request id %s, is proposer: %s", message.RequestId, tss.address.Hex())
		return fmt.Errorf("cannot handleKeygenReq %s, is proposer: %s", message.RequestId, tss.address.Hex())
	}

	tss.sigMu.Lock()
	defer tss.sigMu.Unlock()

	keygenReqMessage := types.KeygenReceiveMessage{
		RequestId:         message.RequestId,
		VoterAddress:      tss.address.Hex(),
		CreateTime:        time.Now().Unix(),
		PublicKeys:        PublicKeysToHex(config.AppConfig.TssPublicKeys),
		Threshold:         config.AppConfig.TssThreshold,
		PublicKeysMatched: CompareStrings(PublicKeysToHex(config.AppConfig.TssPublicKeys), message.PublicKeys),
		ThresholdMatched:  config.AppConfig.TssThreshold == message.Threshold,
	}

	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeKeygenResp,
		RequestId:   message.RequestId,
		DataType:    p2p.DataTypeKeygenResponse,
		Data:        keygenReqMessage,
	}

	err := tss.libp2p.PublishMessage(ctx, p2pMsg)
	if err != nil {
		return err
	}
	log.Debugf("Publish p2p message keygenResponseMessage: RequestId=%s, Key Length=%d, Threshold=%d, Keys=%v",
		message.RequestId, len(keygenReqMessage.PublicKeys), keygenReqMessage.Threshold,
		keygenReqMessage.PublicKeys)
	return nil
}

func (tss *TSSService) handleKeygenReceive(ctx context.Context, event interface{}) error {
	message, ok := event.(types.KeygenReceiveMessage)
	if !ok {
		return fmt.Errorf("handleKeygenReceive error, event %v, is not keygen resp", event)
	}

	tss.sigMu.Lock()
	voteMap, ok := tss.sigMap[message.RequestId]
	if !ok {
		tss.sigMu.Unlock()
		return fmt.Errorf("keygen receive proposer process no sig found, request id: %s", message.RequestId)
	}
	_, ok = voteMap[message.VoterAddress]
	if ok {
		tss.sigMu.Unlock()
		log.Debugf("Keygen proposer process voter multi receive, request id: %s, voter address: %s",
			message.RequestId, message.VoterAddress)
		return nil
	}
	voteMap[message.VoterAddress] = message.ThresholdMatched && message.PublicKeysMatched

	trueCount := countTrueValues(voteMap)
	if trueCount < config.AppConfig.TssThreshold {
		tss.sigMu.Unlock()
		return fmt.Errorf("keygen threshold not reach, request id: %s, true values: %d, threshold: %d",
			message.RequestId, trueCount, config.AppConfig.TssThreshold)
	}

	tss.removeSigMap(message.RequestId, false)

	tss.setup()
	tss.sigMu.Unlock()

	return nil
}

func countTrueValues(voteMap map[string]interface{}) int {
	count := 0

	// Iterate through the map
	for _, value := range voteMap {
		// Assert if the value is of type bool and check if it's true
		if boolValue, ok := value.(bool); ok && boolValue {
			count++
		}
	}

	return count
}
