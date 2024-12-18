package helper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/test"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper/testutil"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReshare(t *testing.T) {
	// t.Skip("intensive test")
	utils.SkipCI(t)

	log.SetLevel(log.DebugLevel)

	// newTotalPartyCount := 10
	// Number of participants in resharing -- t+1 + num new peers
	// newThreshold := 9
	newThreshold := test.TestThreshold

	// err := logging.SetLogLevel("*", "debug")
	// require.NoError(t, err)

	// 1. Get t+1 current keys
	oldPartyIDs := testutil.GetTestPartyIDs(testutil.TestPartyCount - 1)

	oldKeys := testutil.GetTestTssKeys(testutil.TestThreshold + 1)

	t.Log("len(oldKeys)", len(oldKeys))
	t.Log("len(oldPartyIDs)", len(oldPartyIDs))
	// require.Equal(t, keygen.TestThreshold+1, len(oldKeys))
	// require.Equal(t, keygen.TestThreshold+1, len(oldPartyIDs))

	// 2. Create new party IDs to add.. or replace ? confused
	newPartyIDs := tss.GenerateTestPartyIDs(test.TestParticipants)
	require.Len(t, newPartyIDs, test.TestParticipants)

	for _, vla := range oldPartyIDs {
		t.Log("vla", vla, "key", new(big.Int).SetBytes(vla.Key))
	}

	t.Logf("oldPartyIDs: %v;", oldPartyIDs)

	t.Logf("sort old partyIDs: %v", tss.SortPartyIDs(oldPartyIDs))
	t.Logf("new partyIDs: %v", newPartyIDs)

	// 3. Create and connect transport between peers
	oldTransports, newTransports := CreateAndConnectReSharingTransports(t, oldPartyIDs, newPartyIDs)

	// 4. Start resharing party for each peer
	outputAgg := make(chan *keygen.LocalPartySaveData)
	errAgg := make(chan *tss.Error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start old parties
	for i, partyID := range oldPartyIDs {
		params := CreateReShareParams(
			oldPartyIDs,
			newPartyIDs.ToUnSorted(),
			partyID,
			testutil.TestThreshold,
			newThreshold,
		)
		// t.Log(params.PartyID(), hex.EncodeToString(partyID.Key))
		t.Log(partyID, hex.EncodeToString(partyID.Key))

		_, outputCh, errCh := RunReshare(ctx, params, oldKeys[i], oldTransports[i])

		go func(outputCh chan *keygen.LocalPartySaveData, errCh chan *tss.Error) {
			for {
				select {
				case output := <-outputCh:
					outputAgg <- output
				case err := <-errCh:
					errAgg <- err
				}
			}
		}(outputCh, errCh)
	}

	// Start new parties
	for i, partyID := range newPartyIDs {
		params := CreateReShareParams(
			oldPartyIDs,
			newPartyIDs.ToUnSorted(),
			partyID,
			testutil.TestThreshold,
			newThreshold,
		)
		// t.Log(params.PartyID(), hex.EncodeToString(partyID.Key))
		t.Log(partyID, hex.EncodeToString(partyID.Key))

		save := keygen.NewLocalPartySaveData(len(newPartyIDs))
		// Reuse fixture pre-generated preparams
		// save.LocalPreParams = testutil.ReadTestKey(i).LocalPreParams

		_, outputCh, errCh := RunReshare(ctx, params, save, newTransports[i])

		go func(outputCh chan *keygen.LocalPartySaveData, errCh chan *tss.Error) {
			for {
				select {
				case output := <-outputCh:
					outputAgg <- output
				case err := <-errCh:
					errAgg <- err
				}
			}
		}(outputCh, errCh)
	}

	t.Logf("started key reshare")

	newKeys := make([]*keygen.LocalPartySaveData, len(newPartyIDs))

	// Wait for parties to finish
	for i := 0; i < len(oldPartyIDs)+len(newPartyIDs); i++ {
		select {
		case output := <-outputAgg:
			bz, err := json.Marshal(&output)
			require.NoError(t, err)
			t.Log(string(bz))

			// Old committee parties have Xi zeroed, ignore those
			if output.Xi == nil {
				continue
			}

			// new committee -- must use original index in slice
			index, err := output.OriginalIndex()
			assert.NoErrorf(t, err, "should not be an error getting a party's index from save data")

			newKeys[index] = output
		case err := <-errAgg:
			t.Fatal(err)
		}
	}

	require.Equal(t, len(newPartyIDs), len(newKeys), "each party should get a key")

	// xj tests: BigXj == xj*G
	for j, key := range newKeys {
		// xj test: BigXj == xj*G
		xj := key.Xi
		gXj := crypto.ScalarBaseMult(Curve, xj)

		// Uses index here so it must use OriginalIndex(), not append() in arbitrary order
		BigXj := key.BigXj[j]

		assert.True(t, BigXj.Equals(gXj), "ensure BigX_j == g^x_j")
	}

	// New reshared pubkey should match old pubkey
	assert.Truef(t, oldKeys[0].ECDSAPub.Equals(newKeys[0].ECDSAPub), "reshared pubkey should match old pubkey")

	// make sure everyone has the same ECDSA public key
	for i, key := range newKeys {
		for j, key2 := range newKeys {
			// Skip self and previous keys
			if j <= i {
				continue
			}

			assert.Truef(t, key.ECDSAPub.Equals(key2.ECDSAPub), "key %v != %v", i, j)
		}
	}
}
