package db

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DbDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	task := CreateWalletTask{
		BaseTask: BaseTask{
			TaskType: 0,
			TaskId:   1011,
			Task: Task{
				TaskId:      1011,
				TaskType:    0,
				Context:     []byte("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
				Submitter:   "1111",
				BlockHeight: 0,
				Status:      0,
				LogIndexID:  0,
				LogIndex: LogIndex{
					ContractAddress: common.HexToAddress("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
					EventName:       "121221",
					TxHash:          common.Hash{},
					ChainId:         0,
					BlockNumber:     0,
					LogIndex:        0,
				},
			},
		},
		User:    "12",
		Account: 0,
		Chain:   0,
		Index:   0,
	}
	db := dbm.GetRelayerDB().Debug()
	db.DryRun = true
	err := db.Save(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))
	err = db.Model(&CreateWalletTask{}).Where("task_id", 10).Last(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))

	task = CreateWalletTask{}
	// db.DryRun = false
	err = db.Model(&task).Preload("Task").Where("task_id", 10).Last(&task).Error
	assert.Nil(t, err)
	t.Log(utils.FormatJSON(task))
}
