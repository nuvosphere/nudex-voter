package tss

import (
	"context"
	"fmt"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"time"
)

func (tss *TSSService) handleSigStartKeyPrepare(ctx context.Context, e types.MsgSignKeyPrepareMessage) error {
	if tss.address.Hex() != e.MsgSign.VoterAddress {
		log.Debugf("Ignore SigStart request id %s, not proposer: %v", e.RequestId)
		return fmt.Errorf("cannot start sig %s, not proposer: %v", e.RequestId)
	}

	// check map
	_, ok := tss.sigExists(e.RequestId)
	if ok {
		return fmt.Errorf("sig exists: %s", e.RequestId)
	}

	// build sign
	newSign := &types.MsgSignKeyPrepareMessage{
		MsgSign: types.MsgSign{
			RequestId:    e.RequestId,
			IsProposer:   true,
			VoterAddress: tss.address.Hex(),
			SigData:      s.makeSigNewBlock(e.PublicKeys, e.Threshold),
			CreateTime:   time.Now().Unix(),
		},
		PublicKeys: e.PublicKeys,
		Threshold:  e.Threshold,
	}

	// p2p broadcast
	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeSigReq,
		RequestId:   e.RequestId,
		DataType:    "MsgSignNewBlock",
		Data:        *newSign,
	}
	if err := tss.libp2p.PublishMessage(ctx, p2pMsg); err != nil {
		log.Errorf("SigStart key prepare to p2p error, request id: %s, err: %v", e.RequestId, err)
		return err
	}

	return nil
}
