package helper

import (
	"encoding/json"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper/testutil"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/require"
)

func PreParamPath(index int) string {
	return path.Join("testutil", "test-fixtures", fmt.Sprintf("pre-params%d.json", index))
}

func LoadTestPreParam(index int) *keygen.LocalPreParams {
	path := PreParamPath(index)
	bytes := utils.MustReadFile(path)

	var preParams *keygen.LocalPreParams
	if err := json.Unmarshal(bytes, &preParams); err != nil {
		panic(err)
	}

	return preParams
}

func TestPreParam(t *testing.T) {
	utils.SkipCI(t)
	for i := 0; i < 2; i++ {
		preParams := LoadTestPreParam(i)
		require.True(t, preParams.Validate(), "test-fixture preparams should be valid")
	}
}

func TestCreatePreParamFixtures(t *testing.T) {
	utils.SkipCI(t)
	// Takes a while to run -- only run when generating new fixtures or adding
	// more parties
	t.Skip("skip create pre-param fixtures")

	for i := 0; i < testutil.TestPartyCount; i++ {
		t.Logf("creating pre-params for party %d", i)

		preParams, err := keygen.GeneratePreParams(1 * time.Minute)
		require.NoError(t, err)

		b, err := json.MarshalIndent(preParams, "", "  ")
		require.NoError(t, err)

		utils.MustWriteFile(PreParamPath(i), b, 0o644)
	}
}
