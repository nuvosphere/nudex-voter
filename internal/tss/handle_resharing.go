package tss

import (
	"context"
	"fmt"
	"slices"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
)

func (t *TSSService) startReSharing(newAddressList []common.Address, threshold int) error {
	if t.LocalParty == nil || t.LocalPartySaveData == nil || t.LocalPartySaveData.ECDSAPub == nil {
		return fmt.Errorf("local party not initialized")
	}

	currentPartyIDs := createPartyIDsByAddress(t.partners)
	currentPeerCtx := tsslib.NewPeerContext(currentPartyIDs)

	newPartyIDs := createPartyIDsByAddress(newAddressList)
	newPeerCtx := tsslib.NewPeerContext(newPartyIDs)

	currentIndex := slices.Index(t.partners, t.localAddress)
	currentParams := tsslib.NewReSharingParameters(
		tsslib.S256(),
		currentPeerCtx,
		newPeerCtx,
		currentPartyIDs[currentIndex],
		len(t.partners),
		config.AppConfig.TssThreshold,
		len(newAddressList),
		threshold,
	)
	currentParty := resharing.NewLocalParty(currentParams, *t.LocalPartySaveData, t.reSharingOutCh, t.reSharingEndCh).(*resharing.LocalParty)

	go func() {
		if err := currentParty.Start(); err != nil {
			log.Errorf("Failed to start resharing old party, error=%v", err)
			return
		} else {
			log.Infof("Resharing old party started")
		}
	}()

	newIndex := slices.Index(newAddressList, t.localAddress)
	newParams := tsslib.NewReSharingParameters(
		tsslib.S256(),
		currentPeerCtx,
		newPeerCtx,
		newPartyIDs[newIndex],
		len(config.AppConfig.TssPublicKeys),
		threshold,
		len(newAddressList),
		threshold,
	)
	t.reLocalParty = resharing.NewLocalParty(newParams, *t.LocalPartySaveData, t.reSharingOutCh, t.reSharingEndCh).(*resharing.LocalParty)

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
