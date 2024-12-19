package db

import (
	"fmt"
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
	l2SyncDb   *gorm.DB
	l2InfoDb   *gorm.DB
	btcLightDb *gorm.DB
	btcCacheDb *gorm.DB
	walletDb   *gorm.DB
}

func NewDatabaseManager() *DatabaseManager {
	dm := &DatabaseManager{}
	dm.initDB()

	return dm
}

func SetConnParam(db *gorm.DB) {
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

	databaseConfigs := []struct {
		dbPath string
		dbRef  **gorm.DB
		dbName string
	}{
		{filepath.Join(dbDir, "l2_sync.db"), &dm.l2SyncDb, "Database l2_sync"},
		{filepath.Join(dbDir, "l2_info.db"), &dm.l2InfoDb, "Database l2_info"},
		{filepath.Join(dbDir, "btc_light.db"), &dm.btcLightDb, "Database btc_light"},
		{filepath.Join(dbDir, "wallet.db"), &dm.walletDb, "Database wallet"},
		{filepath.Join(dbDir, "btc_cache.db"), &dm.btcCacheDb, "Database btc_cache"},
	}

	for _, dbConfig := range databaseConfigs {
		if err := dm.connectDatabase(dbConfig.dbPath, dbConfig.dbRef, dbConfig.dbName); err != nil {
			log.Fatalf("Failed to connect to %s: %v", dbConfig.dbName, err)
		}
	}

	dm.autoMigrate()
	log.Debugf("Database migration completed successfully")
}

func (dm *DatabaseManager) GetL2InfoDB() *gorm.DB {
	return dm.l2InfoDb
}

func (dm *DatabaseManager) GetBtcLightDB() *gorm.DB {
	return dm.btcLightDb
}

func (dm *DatabaseManager) GetBtcCacheDB() *gorm.DB {
	return dm.btcCacheDb
}

func (dm *DatabaseManager) GetWalletDB() *gorm.DB {
	return dm.walletDb
}

func (dm *DatabaseManager) GetL2SyncDB() *gorm.DB { return dm.l2SyncDb }

func (dm *DatabaseManager) connectDatabase(dbPath string, dbRef **gorm.DB, dbName string) error {
	// open database and set WAL mode
	db, err := gorm.Open(sqlite.Open(dbPath+"?_journal_mode=WAL"), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Warn),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", dbName, err)
	}
	SetConnParam(db)
	*dbRef = db
	log.Debugf("%s connected successfully in WAL mode, path: %s", dbName, dbPath)
	return nil
}

func (dm *DatabaseManager) autoMigrate() {
	if err := dm.l2SyncDb.AutoMigrate(
		&LogIndex{},
		&BTCTransaction{},
		&EVMSyncStatus{},
		&SubmitterChosen{},
		&Participant{},
		&ParticipantEvent{},
		&Account{},
		&DepositRecord{},
		&WithdrawalRecord{},
		&Task{},
		&CreateWalletTask{},
		&DepositTask{},
		&WithdrawalTask{},
		&AddressBalance{},
		&TaskUpdatedEvent{},
		&Operations{},
	); err != nil {
		log.Fatalf("Failed to migrate database 1: %v", err)
	}

	if err := dm.btcLightDb.AutoMigrate(&BtcBlock{}); err != nil {
		log.Fatalf("Failed to migrate database 3: %v", err)
	}

	if err := dm.walletDb.AutoMigrate(&Utxo{}, &Withdraw{}, &SendOrder{}, &Vin{}, &Vout{}, &DepositResult{}); err != nil {
		log.Fatalf("Failed to migrate database 4: %v", err)
	}
	if err := dm.btcCacheDb.AutoMigrate(&BtcSyncStatus{}, &BtcBlockData{}, &BtcTXOutput{}, &Deposit{}); err != nil {
		log.Fatalf("Failed to migrate database 5: %v", err)
	}
}
