package testutil

import (
	"encoding/json"
	"fmt"
	"path"
	"testing"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/test"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/require"
)

func KeyPath(index int) string {
	return path.Join("testutil", "test-fixtures", fmt.Sprintf("localparty-savedata-%02d.json", index))
}

func ReadTestKey(index int) keygen.LocalPartySaveData {
	path := KeyPath(index)

	bytes := utils.MustReadFile(path)

	var key keygen.LocalPartySaveData
	if err := json.Unmarshal(bytes, &key); err != nil {
		panic(err)
	}

	return key
}

func GetTestTssKeys(count int) []keygen.LocalPartySaveData {
	var keys []keygen.LocalPartySaveData

	for i := 0; i < count; i++ {
		key := ReadTestKey(i)
		keys = append(keys, key)
	}

	return keys
}

func WriteTestKey(index int, key keygen.LocalPartySaveData) {
	bz, err := json.MarshalIndent(&key, "", "  ")
	if err != nil {
		panic(err)
	}

	utils.MustWriteFile(KeyPath(index), bz, 0o600)
}

func TestLoadKey(t *testing.T) {
	for i := 0; i < test.TestParticipants; i++ {
		key := ReadTestKey(i)
		require.True(t, key.Validate(), "test-fixture keys should be valid")
	}
}
