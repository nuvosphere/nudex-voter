package tss

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

// init setup.
func (t *TSSService) setup() {
	t.LocalParty = nil
	t.setupTime = time.Time{}

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

	// todo online contact get address list
	t.partners = lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(t.partners, t.localAddress)
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)

	party := keygen.NewLocalParty(params, t.keyOutCh, t.keyEndCh, *preParams).(*keygen.LocalParty)

	t.setupTime = time.Now()
	t.LocalParty = party
	t.partyIdMap = make(map[string]*tsslib.PartyID)

	for _, partyId := range partyIDs {
		t.partyIdMap[partyId.Id] = partyId
	}

	t.LocalPartySaveData = localPartySaveData

	if localPartySaveData == nil || localPartySaveData.ECDSAPub == nil {
		if err := party.Start(); err != nil {
			log.Errorf("TSS keygen process failed to start: %v", err)
			return
		}
	}
}

func (t *TSSService) handleTssKeyEnd(event *keygen.LocalPartySaveData) error {
	t.assertLocalParty(event)
	return saveTSSData(event)
}

var ErrLocalParty = errors.New("local party not initialized")

func (t *TSSService) assertLocalParty(extra any) {
	if t.LocalParty == nil {
		panic(fmt.Errorf("%w, extra:%v", ErrLocalParty, extra))
	}
}
