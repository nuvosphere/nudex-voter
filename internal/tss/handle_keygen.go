package tss

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func (t *TSSService) IsGenesis() bool {
	if t.localPartySaveData != nil && t.localPartySaveData.ECDSAPub != nil {
		return false
	}

	localPartySaveData, err := LoadTSSData()
	if err != nil {
		log.Errorf("Failed to load TSS data: %v", err)
	}

	if localPartySaveData == nil {
		return true
	}

	t.localPartySaveData = localPartySaveData

	return false
}

func (t *TSSService) Genesis(ctx context.Context) {
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		log.Fatalf("Failed to generate TSS preParams: %v", err)
	}

	log.Debugf("Generated TSS preParams: %+v", preParams)

	// todo online contact get address list
	t.partners = lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)
	params := tsslib.NewParameters(
		tsslib.S256(),
		peerCtx,
		partyIDs.FindByKey(new(big.Int).SetBytes(t.localAddress.Bytes())),
		len(partyIDs),
		config.AppConfig.TssThreshold,
	)

	party := keygen.NewLocalParty(params, t.keyOutCh, t.keyEndCh, *preParams).(*keygen.LocalParty)

	t.localParty = party
	t.partyIdMap = make(map[string]*tsslib.PartyID)

	for _, partyId := range partyIDs {
		t.partyIdMap[partyId.Id] = partyId
	}

	if err := party.Start(); err != nil {
		log.Errorf("TSS keygen process failed to start: %v", err)
	}
	// helper.RunKeyGen(ctx, preParams, params)

	log.Info("TSS keygen process started")
}

func (t *TSSService) Partners() []common.Address {
	// todo online contact get address list
	return lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}

//func (t *TSSService) Party() []common.Address {
//t.partners = t.Partners()
//partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
//peerCtx := tsslib.NewPeerContext(partyIDs)
//
//index := AddressIndex(t.partners, t.localAddress)
//params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)
//
//party := keygen.NewLocalParty(params, t.keyOutCh, t.keyEndCh, t.localPartySaveData.LocalPreParams).(*keygen.LocalParty)
//t.localParty = party
//return nil
//}

// init setup.
func (t *TSSService) setup() {
	t.localParty = nil
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
		t.localPartySaveData = localPartySaveData
		preParams = &localPartySaveData.LocalPreParams

		log.Infof("Loaded TSS data as prePrams")
	}

	t.partners = t.Partners()
	partyIDs := createPartyIDs(config.AppConfig.TssPublicKeys)
	peerCtx := tsslib.NewPeerContext(partyIDs)

	index := AddressIndex(t.partners, t.localAddress)
	params := tsslib.NewParameters(tsslib.S256(), peerCtx, partyIDs[index], len(partyIDs), config.AppConfig.TssThreshold)

	party := keygen.NewLocalParty(params, t.keyOutCh, t.keyEndCh, *preParams).(*keygen.LocalParty)

	t.setupTime = time.Now()
	t.localParty = party
	t.partyIdMap = make(map[string]*tsslib.PartyID)

	for _, partyId := range partyIDs {
		t.partyIdMap[partyId.Id] = partyId
	}

	if err := party.Start(); err != nil {
		log.Errorf("TSS keygen process failed to start: %v", err)
		return
	}
}

func (t *TSSService) handleTssKeyEnd(event *keygen.LocalPartySaveData) error {
	return saveTSSData(event)
}

var ErrLocalParty = errors.New("local party not initialized")
