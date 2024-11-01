package tss

import (
	"context"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
	"slices"
)

func (tss *TSSService) startReSharing(newAddressList []common.Address, threshold int) error {
	if tss.LocalParty == nil || tss.LocalPartySaveData == nil || tss.LocalPartySaveData.ECDSAPub == nil {
		return fmt.Errorf("local party not initialized")
	}

	currentPartyIDs := createPartyIDsByAddress(tss.addressList)
	currentPeerCtx := tsslib.NewPeerContext(currentPartyIDs)

	newPartyIDs := createPartyIDsByAddress(newAddressList)
	newPeerCtx := tsslib.NewPeerContext(newPartyIDs)

	currentIndex := slices.Index(tss.addressList, tss.Address)
	currentParams := tsslib.NewReSharingParameters(
		tsslib.S256(),
		currentPeerCtx,
		newPeerCtx,
		currentPartyIDs[currentIndex],
		len(tss.addressList),
		config.AppConfig.TssThreshold,
		len(newAddressList),
		threshold,
	)
	currentParty := resharing.NewLocalParty(currentParams, *tss.LocalPartySaveData, tss.reSharingOutCh, tss.reSharingEndCh).(*keygen.LocalParty)

	go func() {
		if err := currentParty.Start(); err != nil {
			log.Errorf("Failed to start resharing old party, error=%v", err)
			return
		} else {
			log.Infof("Resharing old party started")
		}
	}()

	newIndex := slices.Index(newAddressList, tss.Address)
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
	newParty := resharing.NewLocalParty(newParams, *tss.LocalPartySaveData, tss.reSharingOutCh, tss.reSharingEndCh).(*keygen.LocalParty)

	go func() {
		if err := newParty.Start(); err != nil {
			log.Errorf("Failed to start resharing new party, error=%v", err)
			return
		} else {
			log.Infof("Resharing new party started")
		}
	}()
	return nil
}

func (tss *TSSService) handleTssReSharingOut(ctx context.Context, msg tsslib.Message) (err error) {
	dest := msg.GetTo()
	if dest == nil {
		return fmt.Errorf("did not expect a msg to have a nil destination during resharing")
	}

	_, err = tss.sendTssMsg(ctx, DataTypeTssReSharingMsg, msg)
	return err
}
