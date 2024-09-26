package db

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type DatabaseManager struct {
	relayerDb *gorm.DB
}

func NewDatabaseManager() *DatabaseManager {
	dm := &DatabaseManager{}
	dm.initDB()
	return dm
}

func (dm *DatabaseManager) initDB() {
	dbDir := config.AppConfig.DbDir
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	relayerPath := filepath.Join(dbDir, "relayer_data.db")
	relayerDb, err := gorm.Open(sqlite.Open(relayerPath), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database 1: %v", err)
	}
	dm.relayerDb = relayerDb
	log.Debugf("Database 1 connected successfully, path: %s", relayerPath)

	dm.autoMigrate()
	log.Debugf("Database migration completed successfully")
}

func (dm *DatabaseManager) GetRelayerDB() *gorm.DB {
	return dm.relayerDb
}
