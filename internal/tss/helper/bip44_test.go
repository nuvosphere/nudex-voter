package helper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper/testutil"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBip44GenerateAddress(t *testing.T) {
	localData := testutil.ReadTestKey(1)
	masterPublicKey := *(localData.ECDSAPub.ToECDSAPubKey())
	t.Log("master address: ", crypto.PubkeyToAddress(masterPublicKey))
	address := wallet.GenerateAddressByPath(masterPublicKey, 60, 0, 0)
	t.Log("address: ", address)
}

func TestHDSign(t *testing.T) {
	// err := logging.SetLogLevel("*", "debug")
	// require.NoError(t, err)
	_, _, keys, signPIDs := testutil.GetTestKeys(t, testutil.TestThreshold+1)

	// 2. Create and connect transport between peers
	transports := CreateAndConnectTransports(t, signPIDs)
	require.Len(t, transports, testutil.TestThreshold+1)

	// 3. Start signing party for each peer
	outputAgg := make(chan *common.SignatureData, testutil.TestThreshold)
	errAgg := make(chan *tss.Error, testutil.TestThreshold)

	msgHash := big.NewInt(1234)

	keyDerivationDelta, extendedChildPk, errorDerivation := wallet.DerivingPubKeyFromPath(*(keys[0].ECDSAPub.ToECDSAPubKey()), []uint32{44, 60, 0, 0, 0})
	assert.NoErrorf(t, errorDerivation, "there should not be an error deriving the child public key")

	err := signing.UpdatePublicKeyAndAdjustBigXj(keyDerivationDelta, keys, &extendedChildPk.PublicKey, Curve)
	assert.NoErrorf(t, err, "there should not be an error setting the derived keys")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range signPIDs {
		params := CreateParams(signPIDs, signPIDs[i], testutil.TestThreshold)
		t.Log(params.PartyID())

		// big.Int message, would be message hash converted to big int
		outputCh, errCh := RunSignWithHD(ctx, msgHash, params, keys[i], transports[i], keyDerivationDelta)

		go func(outputCh chan *common.SignatureData, errCh chan *tss.Error) {
			for {
				select {
				//nolint:govet // https://github.com/bnb-chain/tss-lib/pull/167
				case output := <-outputCh:
					outputAgg <- output
				case err := <-errCh:
					errAgg <- err
				}
			}
		}(outputCh, errCh)
	}

	t.Logf("started signing")

	var signatures []*common.SignatureData

	for range signPIDs {
		select {
		//nolint:govet
		case output := <-outputAgg:
			bz, err := json.Marshal(&output)
			require.NoError(t, err)
			t.Log(string(bz))

			//nolint:govet
			signatures = append(signatures, output)

		case err := <-errAgg:
			t.Logf("err: %v", err)
		}
	}

	require.Len(t, signatures, testutil.TestThreshold+1, "each party should get a signature")

	//nolint:govet
	for i, sig := range signatures {
		//nolint:govet
		for j, sig2 := range signatures {
			// Skip self and previous keys
			if j <= i {
				continue
			}

			// t.Log(utils.VerifySig(ethcommon.BytesToHash(msgHash.Bytes()), sig.Signature, ethcommon.HexToAddress("0x4326E0BcE74b754A29627D3A15C2CADD6F7b5DaA")))
			// make sure everyone has the same signature
			assert.True(t, bytes.Equal(sig.Signature, sig2.Signature))
		}
	}

	// Verify signature
	pkX, pkY := extendedChildPk.X, extendedChildPk.Y
	pk := ecdsa.PublicKey{
		Curve: Curve,
		X:     pkX,
		Y:     pkY,
	}

	ok := ecdsa.Verify(
		&pk,                                    // pubkey
		msgHash.Bytes(),                        // message hash
		new(big.Int).SetBytes(signatures[0].R), // R
		new(big.Int).SetBytes(signatures[0].S), // S
	)
	assert.True(t, ok, "ecdsa verify must pass")

	t.Log("signature verified")
}
