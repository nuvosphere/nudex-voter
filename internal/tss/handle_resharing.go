package tss

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
)

func (tss *TSSService) startResharing(publicKeys []*ecdsa.PublicKey, threshold int) error {
	if tss.Party == nil {
		return fmt.Errorf("startResharing error, self not init")
	}

	localPartySaveData, err := LoadTSSData()
	if err != nil {
		return err
	}
	if localPartySaveData == nil || localPartySaveData.ECDSAPub == nil {
		return fmt.Errorf("startResharing error, self not init")
	}
	currentPartyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	currentPeerCtx := tsslib.NewPeerContext(currentPartyIDs)

	newPartyIDs := createPartyIDs(publicKeys)
	newPeerCtx := tsslib.NewPeerContext(newPartyIDs)

	currentIndex := AddressIndex(config.AppConfig.TssPublicKeys, tss.Address.Hex())
	currentParams := tsslib.NewReSharingParameters(tsslib.S256(), currentPeerCtx, newPeerCtx, currentPartyIDs[currentIndex], len(config.AppConfig.TssPublicKeys), config.AppConfig.TssThreshold, len(publicKeys), threshold)
	currentParty := resharing.NewLocalParty(currentParams, *localPartySaveData, tss.keyOutCh, tss.keygenEndCh).(*keygen.LocalParty)

	go func() {
		if err := currentParty.Start(); err != nil {
			log.Errorf("Failed to start resharing old party, error=%v", err)
			return
		} else {
			log.Infof("Resharing old party started")
		}
	}()

	newIndex := AddressIndex(publicKeys, tss.Address.Hex())
	newParams := tsslib.NewReSharingParameters(tsslib.S256(), currentPeerCtx, newPeerCtx, newPartyIDs[newIndex], len(config.AppConfig.TssPublicKeys), threshold, len(publicKeys), threshold)
	newParty := resharing.NewLocalParty(newParams, *localPartySaveData, tss.keyOutCh, tss.keygenEndCh).(*keygen.LocalParty)

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
