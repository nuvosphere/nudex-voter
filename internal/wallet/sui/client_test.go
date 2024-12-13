package sui

import (
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	utils.SkipCI(t)

	c := NewDevClient()

	data := base58.Decode("5ZnCSBuoktAiv1titQWUzHd9iqvy9sD8vQNMrHxZMR8KMzjwkM3GQyX7qfoZJ6cYU1HLEX6bT25B2rtRhKiM8MVc")
	pk, pubKey := edwards.PrivKeyFromBytes(data)

	fromAddress := Ed25519PublicKeyToSuiAddress(pubKey.Serialize())
	t.Logf("fromAddress: %s", fromAddress) // 0x9099b85cce1e63a584f981390bf3457611df3f7778c0d77de3f16cb57951bcf9

	// from := "0x9cbec822bd17762d757bf741ef9dcea763cc414d970baffb52c191978fabe266"
	to := "0x5283816ef0fe030955141418c61ac7e362101eb251ca6e9e9d812ca2e803320c"

	// curl --location --request POST 'https://faucet.devnet.sui.io/v1/gas' --header 'Content-Type: application/json' --data-raw '{"FixedAmountRequest": {"recipient": "0x9099b85cce1e63a584f981390bf3457611df3f7778c0d77de3f16cb57951bcf9"}}'
	amount, err := c.GetBalance("0x9099b85cce1e63a584f981390bf3457611df3f7778c0d77de3f16cb57951bcf9", "0x02", "sui")
	assert.Nil(t, err)
	t.Logf("amount: %s", amount)

	unSignTx, err := c.BuildTransferTx(fromAddress, to, "0x09ef901eb531f155fd58e5c52f6047d9d0b358d2c6ef974f57038be4b258a7cc", 111)
	assert.NoError(t, err)

	_, err = c.TryExecuteTx(unSignTx)
	assert.NoError(t, err)

	signature, err := pk.Sign(unSignTx.TxHash())
	assert.NoError(t, err)

	signTx := unSignTx.SerializedSigWith(signature.Serialize(), pubKey.Serialize())

	hash, err := c.SendTx((*SignedTx)(signTx))
	assert.NoError(t, err)
	t.Log("hash", string(hash))
}
