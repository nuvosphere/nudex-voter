package tss

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestOldPartyID(t *testing.T) {
	addressList := []common.Address{
		common.HexToAddress("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
		common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD"),
		common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		common.HexToAddress("0x38015D05f4fEC8AFe15D7cc0386a126574e8077B"),
		common.HexToAddress("0x41ff9AA7e16B8B1a8a8dc4f0eFacd93D02d071c9"),
	}

	oldPartyIDs := createOldPartyIDsByAddress(addressList)
	t.Log(oldPartyIDs[0].Id)
	t.Log(oldPartyIDs[0].Moniker)
	t.Logf("%x", oldPartyIDs[0].Key)

	assert.Equal(t, oldPartyIDs[0].Id, oldPartyIDs[0].Moniker)
	assert.Equal(t, oldPartyIDs[0].Id, "33128a8fc17869897dce68ed026d694621f6fdfe")
}