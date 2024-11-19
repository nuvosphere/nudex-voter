package tss

import (
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

type PartyData struct {
	rw    sync.RWMutex
	datas map[helper.CurveType]*helper.LocalPartySaveData
}

func NewPartyData() *PartyData {
	return &PartyData{
		rw:    sync.RWMutex{},
		datas: make(map[helper.CurveType]*helper.LocalPartySaveData),
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
	data, _ = LoadTSSData(ec)
	p.datas[ec] = data
	p.rw.Unlock()

	return data
}

func (p *PartyData) LoadData() bool {
	p.rw.Lock()
	defer p.rw.Unlock()

	data, err := LoadTSSData(helper.ECDSA)
	if err != nil {
		return false
	}

	p.datas[data.CurveType()] = data

	data, err = LoadTSSData(helper.EDDSA)
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

	return saveTSSData(data)
}
