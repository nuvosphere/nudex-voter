package wallet

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

const (
	ecdsaPublicKey = "0x026ae06bb6b7a4779ef7d2fbcb5da36bec729c54e8b9c235aa75b09a5e22dd427b"
	eddsaPublicKey = "44a3e1108c206006fbcc5d3a5e33dfba38b0f3bca00fe0ccdfc2267e712271a1"
)

func TestMasterPublicKey(t *testing.T) {
	pubKey, err := ethCrypto.DecompressPubkey(hexutil.MustDecode(ecdsaPublicKey))
	assert.Nil(t, err)

	tssAddress := ethCrypto.PubkeyToAddress(*pubKey)
	t.Log(tssAddress.String())
	assert.Equal(t, tssAddress, common.HexToAddress("0xB43EB0e9Ec8040737FFcc144073C72Cf68bC4bab")) // 🙆
}
