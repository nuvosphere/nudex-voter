package task

import (
	"encoding/hex"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeCreateWalletTask(t *testing.T) {
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

func TestEncodeDepositTask(t *testing.T) {
	depositTask := types.DepositTask{
		BaseTask: types.BaseTask{
			TaskId: 1,
		},
		TargetAddress:   "0xFa0c1810C5853348020e15a9C300c2363b5EBF41",
		Amount:          uint64(1000000000000000000),
		Chain:           uint8(ETHEREUM),
		ChainId:         uint32(1),
		BlockHeight:     uint64(21133979),
		TxHash:          "0x01cfa36f443bca6774be814ef667ead31be4493c6101e0093ab9a1d5142cb5a8",
		ContractAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		Ticker:          "USDC",
		AssetType:       uint8(ERC20),
		Decimal:         18,
	}
	bytes, err := encodeDepositTask(depositTask)
	assert.NoError(t, err)
	assert.NotNil(t, bytes)

	assert.Equal(t, "a16e6f3f5818b6d9cdc9da7bb4b22b721875ee019d05d2416544b3ec35fe7b8b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000", hex.EncodeToString(bytes))
}
