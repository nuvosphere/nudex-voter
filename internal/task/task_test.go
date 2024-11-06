package task

import (
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
}
