package db

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DbDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	var taskId uint32 = 13
	task := CreateWalletTask{
		BaseTask: BaseTask{
			TaskType: 0,
			TaskId:   taskId,
			Task: Task{
				TaskId:    taskId,
				TaskType:  0,
				Context:   []byte("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
				Submitter: "1111",
				Status:    0,
				LogIndex: LogIndex{
					ContractAddress: common.HexToAddress("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
					EventName:       "121221",
					TxHash:          common.Hash{},
					ChainId:         60,
					BlockNumber:     10,
					LogIndex:        20,
				},
			},
		},
		Account: 0,
		Chain:   0,
		Index:   0,
	}
	db := dbm.GetRelayerDB().Debug()
	db.DryRun = true
	// err := db.Create(&task).Error
	// err := db.Save(&task).Error
	// assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))
	err := db.Model(&CreateWalletTask{}).Where("task_id", taskId).Last(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))

	task = CreateWalletTask{}
	// db.DryRun = tr
	err = db.Model(&task).Preload(clause.Associations).Where("task_id", taskId).Last(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))

	baseTask := Task{}
	err = db.Model(&Task{}).Preload(clause.Associations).Where("task_id", taskId).Last(&baseTask).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(baseTask))
	// err = db.Preload(clause.Associations).Where("task_id", taskId).Last(&baseTask).Error
	// assert.Nil(t, err)
	// t.Log(utils.FormatJSON(baseTask))
	t.Log("end")
}

func TestUniqueTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DbDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	var taskId uint32 = 26
	task := CreateWalletTask{
		BaseTask: BaseTask{
			TaskType: 0,
			TaskId:   taskId,
			Task: Task{
				TaskId:    taskId,
				TaskType:  0,
				Context:   []byte("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
				Submitter: "1111",
				Status:    0,
				LogIndex: LogIndex{
					ContractAddress: common.HexToAddress("0x5e7BB104d84c7CB9B682AaC2F3d509f5F406809A"),
					EventName:       "121221",
					TxHash:          common.Hash{},
					ChainId:         60,
					BlockNumber:     10,
					LogIndex:        20,
				},
			},
		},
		Account: 0,
		Chain:   0,
		Index:   0,
	}
	db := dbm.GetRelayerDB().Debug()
	db.DryRun = true
	err := db.Save(&task).Error
	// err := db.Save(&task).Error
	assert.Nil(t, err)
	// t.Log(utils.FormatJSON(task))
	task = CreateWalletTask{}
	err = db.Model(&CreateWalletTask{}).Where("task_id", 25).Last(&task).Error
	assert.Nil(t, err)
	t.Log(utils.FormatJSON(task))
	t.Log("end")
}

func TestSubmitterChosenUniqueTask(t *testing.T) {
	utils.SkipCI(t)

	config.AppConfig.DbDir = "./"
	dbm := NewDatabaseManager()
	dbm.initDB()

	s := SubmitterChosen{
		Submitter: "122",
	}
	db := dbm.GetRelayerDB().Debug()
	db.DryRun = true
	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Save(&s).Error
	assert.Nil(t, err)
	t.Log("end")
}
