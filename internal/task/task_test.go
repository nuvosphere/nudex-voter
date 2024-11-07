package task

import (
	"encoding/hex"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskEncode(t *testing.T) {
	createWalletTask := types.CreateWalletTask{
		BaseTask: types.BaseTask{
			TaskId: 1,
		},
		User:    "0xFa0c1810C5853348020e15a9C300c2363b5EBF41",
		Account: uint64(10001),
		Chain:   uint8(ETHEREUM),
		Index:   uint32(0),
	}
	bytes, err := encodeCreateWalletTask(createWalletTask)
	assert.NoError(t, err)
	assert.NotNil(t, bytes)

	assert.Equal(t, "0f1413e8d10cd1ec520cd20e110e7f744aadade9260edf12bea2ae80bf938c2e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000fa0c1810c5853348020e15a9c300c2363b5ebf41000000000000000000000000000000000000000000000000000000000000271100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", hex.EncodeToString(bytes))
}
