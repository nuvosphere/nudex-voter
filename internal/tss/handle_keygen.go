package tss

import (
	"errors"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
	"time"
)

func (tss *TSSService) setup() {
	tss.LocalParty = nil
	tss.setupTime = time.Time{}

	var preParams *keygen.LocalPreParams
	localPartySaveData, err := LoadTSSData()
	if err != nil {
		log.Errorf("Failed to load TSS data: %v", err)
		preParams, err = keygen.GeneratePreParams(1 * time.Minute)
		if err != nil {
			log.Fatalf("Failed to generate TSS preParams: %v", err)
		}
		log.Debugf("Generated TSS preParams: %+v", preParams)
		err = saveTSSData(preParams)
		if err != nil {
			log.Fatalf("Failed to save TSS data: %v", err)
		}
	} else {
		preParams = &localPartySaveData.LocalPreParams
		log.Infof("Loaded TSS data as prePrams")
	}

	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(config.AppConfig.TssPublicKeys, tss.Address.Hex())
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)

	party := keygen.NewLocalParty(params, tss.keyOutCh, tss.keyEndCh, *preParams).(*keygen.LocalParty)

	tss.setupTime = time.Now()
	tss.LocalParty = party
	tss.partyIdMap = make(map[string]*tsslib.PartyID)
	for _, partyId := range partyIDs {
		tss.partyIdMap[partyId.Id] = partyId
	}
	tss.LocalPartySaveData = localPartySaveData

	if localPartySaveData == nil || localPartySaveData.ECDSAPub == nil {
		if err := party.Start(); err != nil {
			log.Errorf("TSS keygen process failed to start: %v", err)
			return
		}
	}
}

func (tss *TSSService) handleTssKeyEnd(event *keygen.LocalPartySaveData) error {
	tss.assertLocalParty(event)
	return saveTSSData(event)
}

var ErrLocalParty = errors.New("local party not initialized")

func (tss *TSSService) assertLocalParty(extra any) {
	if tss.LocalParty == nil {
		panic(fmt.Errorf("%w, extra:%v", ErrLocalParty, extra))
	}
}
