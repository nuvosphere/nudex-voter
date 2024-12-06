package wallet

import (
	"encoding/base64"
	"testing"

	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/stretchr/testify/assert"
)

func TestSuiAddress(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString("ALoFvhYh7S9eDM+hxB9cx6O1UHkRSfEdOAL9geHj0DME")
	t.Logf("%x", data)
	assert.Nil(t, err)
	pubkey, err := edwards.ParsePubKey(data[1:])
	assert.Nil(t, err)

	address := Ed25519PublicKeyToSuiAddress(pubkey.Serialize())
	t.Log(address)
	assert.Equal(t, "0x9cbec822bd17762d757bf741ef9dcea763cc414d970baffb52c191978fabe266", address)
}
