package tss

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

type PartyData struct {
	basePath string
	rw       sync.RWMutex
	datas    map[helper.CurveType]*helper.LocalPartySaveData
}

func NewPartyData(basePath string) *PartyData {
	return &PartyData{
		basePath: basePath,
		rw:       sync.RWMutex{},
		datas:    make(map[helper.CurveType]*helper.LocalPartySaveData),
	}
}

func (p *PartyData) ECDSALocalData() *helper.LocalPartySaveData {
	return p.GetData(helper.ECDSA)
}

func (p *PartyData) EDDSALocalData() *helper.LocalPartySaveData {
	return p.GetData(helper.EDDSA)
}

func (p *PartyData) GetData(ec helper.CurveType) *helper.LocalPartySaveData {
	p.rw.RLock()
	data, ok := p.datas[ec]
	p.rw.RUnlock()

	if ok {
		return data
	}

	p.rw.Lock()
	data, _ = p.loadTSSData(ec)
	p.datas[ec] = data
	p.rw.Unlock()

	return data
}

func (p *PartyData) GenerateNewLocalPartySaveData(ec helper.CurveType, parties types.Participants) *helper.LocalPartySaveData {
	switch ec {
	case helper.ECDSA:
		save := ecdsaKeygen.NewLocalPartySaveData(parties.Len())
		// save.LocalPreParams = p.ECDSALocalData().ECDSAData().LocalPreParams //todo
		return helper.BuildECDSALocalPartySaveData().SetData(&save)
	case helper.EDDSA:
		save := eddsaKeygen.NewLocalPartySaveData(parties.Len())
		return helper.BuildEDDSALocalPartySaveData().SetData(&save)
	}

	return nil
}

func (p *PartyData) LoadData() bool {
	p.rw.Lock()
	defer p.rw.Unlock()

	data, err := p.loadTSSData(helper.ECDSA)
	if err != nil {
		return false
	}

	p.datas[data.CurveType()] = data

	data, err = p.loadTSSData(helper.EDDSA)
	if err != nil {
		return false
	}

	p.datas[data.CurveType()] = data

	return true
}

func (p *PartyData) SaveLocalData(data *helper.LocalPartySaveData) error {
	p.rw.Lock()
	defer p.rw.Unlock()
	p.datas[data.CurveType()] = data

	return p.saveTSSData(data)
}

func (p *PartyData) saveTSSData(data *helper.LocalPartySaveData) error {
	curveType := data.CurveType()

	dataDir := filepath.Join(p.basePath, "tss_data", curveType.CurveName())
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		log.Errorf("Failed to create TSS data directory: %v", err)
		return err
	}

	dataBytes, err := json.Marshal(data.GetData())
	if err != nil {
		log.Errorf("Unable to serialize TSS data: %v", err)
		return err
	}

	filePath := filepath.Join(dataDir, "tss_key_data.json")
	if err := os.WriteFile(filePath, dataBytes, 0o600); err != nil {
		log.Errorf("Failed to save TSS data to file: %v", err)
		return err
	}

	log.Infof("TSS data successfully saved to: %s", filePath)

	return nil
}

func (p *PartyData) loadTSSData(ec helper.CurveType) (*helper.LocalPartySaveData, error) {
	filePath := filepath.Join(p.basePath, "tss_data", ec.CurveName(), "tss_key_data.json")

	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read TSS data file: %v", err)
	}

	switch ec {
	case helper.ECDSA:
		var data ecdsaKeygen.LocalPartySaveData
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, fmt.Errorf("unable to deserialize TSS data: %v", err)
		}

		return helper.BuildECDSALocalPartySaveData().SetData(&data), nil
	case helper.EDDSA:
		var data eddsaKeygen.LocalPartySaveData
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, fmt.Errorf("unable to deserialize TSS data: %v", err)
		}

		return helper.BuildEDDSALocalPartySaveData().SetData(&data), nil
	}

	return nil, fmt.Errorf("unknown elliptic curve")
}
