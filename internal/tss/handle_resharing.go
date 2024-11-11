package tss

import (
	"context"
	"fmt"
	"math/big"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
)

func (t *TSSService) startReSharing(newAddressList []common.Address, threshold int) error {
	oldPartyIDs := createPartyIDsByAddress(t.oldPartners())
	oldPeerCtx := tsslib.NewPeerContext(oldPartyIDs)
	oldPartyID := oldPartyIDs.FindByKey(new(big.Int).SetBytes(t.localAddress.Bytes()))

	newPartyIDs := createPartyIDsByAddress(newAddressList)
	newPeerCtx := tsslib.NewPeerContext(newPartyIDs)
	newPartyID := newPartyIDs.FindByKey(new(big.Int).SetBytes(t.localAddress.Bytes()))

	oldParams := tsslib.NewReSharingParameters(
		tsslib.S256(),
		oldPeerCtx,
		newPeerCtx,
		oldPartyID,
		len(t.oldPartners()),
		config.AppConfig.TssThreshold,
		len(newAddressList),
		threshold,
	)
	oldParty := resharing.NewLocalParty(oldParams, *t.localPartySaveData, t.reSharingOutCh, t.reSharingEndCh).(*resharing.LocalParty)

	go func() {
		if err := oldParty.Start(); err != nil {
			log.Errorf("Failed to start resharing old party, error=%v", err)
			return
		} else {
			log.Infof("Resharing old party started")
		}
	}()

	newParams := tsslib.NewReSharingParameters(
		tsslib.S256(),
		oldPeerCtx,
		newPeerCtx,
		newPartyID,
		len(t.oldPartners()),
		config.AppConfig.TssThreshold,
		len(newAddressList),
		threshold,
	)
	t.reLocalParty = resharing.NewLocalParty(newParams, *t.localPartySaveData, t.reSharingOutCh, t.reSharingEndCh).(*resharing.LocalParty)

	go func() {
		if err := t.reLocalParty.Start(); err != nil {
			log.Errorf("Failed to start resharing new party, error=%v", err)
			return
		} else {
			log.Infof("Resharing new party started")
		}
	}()

	return nil
}

func (t *TSSService) handleTssReSharingOut(ctx context.Context, msg tsslib.Message) (err error) {
	dest := msg.GetTo()
	if dest == nil && !msg.IsBroadcast() {
		return fmt.Errorf("did not expect a msg to have a nil destination and not broadcast during resharing")
	}

	_, err = t.sendTssMsg(ctx, DataTypeTssReSharingMsg, msg)

	return err
}
