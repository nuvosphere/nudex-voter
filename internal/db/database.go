package db

import (
	"os"
	"path/filepath"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type DatabaseManager struct {
	relayerDb  *gorm.DB
	btcLightDb *gorm.DB
	btcCacheDb *gorm.DB
}

func NewDatabaseManager() *DatabaseManager {
	dm := &DatabaseManager{}
	dm.initDB()

	return dm
}

func setConnParam(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func (dm *DatabaseManager) initDB() {
	dbDir := config.AppConfig.DbDir
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	relayerPath := filepath.Join(dbDir, "relayer_data.db")

	relayerDb, err := gorm.Open(sqlite.Open(relayerPath), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Warn),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	if err != nil {
		log.Fatalf("Failed to connect to database 1: %v", err)
	}

	setConnParam(relayerDb)
	dm.relayerDb = relayerDb

	log.Debugf("Database 1 connected successfully, path: %s", relayerPath)

	btcLightPath := filepath.Join(dbDir, "btc_light.db")

	btcLightDb, err := gorm.Open(sqlite.Open(btcLightPath), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Warn),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	if err != nil {
		log.Fatalf("Failed to connect to database 2: %v", err)
	}

	setConnParam(btcLightDb)
	dm.btcLightDb = btcLightDb

	log.Debugf("Database 2 connected successfully, path: %s", btcLightPath)

	btcCachePath := filepath.Join(dbDir, "btc_cache.db")

	btcCacheDb, err := gorm.Open(sqlite.Open(btcCachePath), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Warn),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	if err != nil {
		log.Fatalf("Failed to connect to database 3: %v", err)
	}

	setConnParam(btcCacheDb)
	dm.btcCacheDb = btcCacheDb

	log.Debugf("Database 3 connected successfully, path: %s", btcCachePath)

	dm.autoMigrate()
	log.Debugf("Database migration completed successfully")
}

func (dm *DatabaseManager) GetRelayerDB() *gorm.DB {
	return dm.relayerDb
}

func (dm *DatabaseManager) GetBtcLightDB() *gorm.DB {
	return dm.btcLightDb
}

func (dm *DatabaseManager) GetBtcCacheDB() *gorm.DB {
	return dm.btcCacheDb
}

func (dm *DatabaseManager) autoMigrate() {
	if err := dm.relayerDb.AutoMigrate(
		&LogIndex{},
		&BTCTransaction{},
		&EVMSyncStatus{},
		&SubmitterChosen{},
		&Participant{},
		&Account{},
		&DepositRecord{},
		&WithdrawalRecord{},
		&Task{},
		&CreateWalletTask{},
		&DepositTask{},
		&WithdrawalTask{},
	); err != nil {
		log.Fatalf("Failed to migrate database 1: %v", err)
	}

	if err := dm.btcLightDb.AutoMigrate(&BtcBlock{}); err != nil {
		log.Fatalf("Failed to migrate database 3: %v", err)
	}

	if err := dm.btcCacheDb.AutoMigrate(&BtcSyncStatus{}, &BtcBlockData{}, &BtcTXOutput{}); err != nil {
		log.Fatalf("Failed to migrate database 2: %v", err)
	}
}
