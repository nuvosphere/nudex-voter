package db

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type BTCTransaction struct {
	ID          uint      `gorm:"primaryKey"`
	TxID        string    `gorm:"uniqueIndex;not null"`
	RawTxData   string    `gorm:"type:text;not null"`
	ReceivedAt  time.Time `gorm:"not null"`
	Processed   bool      `gorm:"default:false"`
	ProcessedAt time.Time
}

type EVMSyncStatus struct {
	ID            uint      `gorm:"primaryKey"`
	LastSyncBlock uint64    `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}

type WithdrawalRecord struct {
	ID           uint      `gorm:"primaryKey"`
	WithdrawalID string    `gorm:"uniqueIndex;not null"`
	UserAddress  string    `gorm:"not null"`
	Amount       string    `gorm:"not null"`
	DetectedAt   time.Time `gorm:"not null"`
	OnChain      bool      `gorm:"default:false"`
	OnChainTxID  string
	Processed    bool `gorm:"default:false"`
	ProcessedAt  time.Time
}

// SubmitterRotation contains block number and current submitter
type SubmitterRotation struct {
	ID               uint64 `gorm:"primaryKey"`
	BlockNumber      uint64 `gorm:"not null"`
	CurrentSubmitter string `gorm:"not null"`
}

// Participant save all participants
type Participant struct {
	ID      uint64 `gorm:"primaryKey"`
	Address string `gorm:"uniqueIndex;not null"`
}

type BtcBlock struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Height    uint64    `gorm:"not null;uniqueIndex" json:"height"`
	Hash      string    `gorm:"not null" json:"hash"`
	Status    string    `gorm:"not null" json:"status"` // "unconfirm", "confirmed", "signing", "pending", "processed"
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type BtcSyncStatus struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UnconfirmHeight int64     `gorm:"not null" json:"unconfirm_height"`
	ConfirmedHeight int64     `gorm:"not null" json:"confirmed_height"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`
}

type BtcBlockData struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	BlockHeight  uint64 `gorm:"unique;not null" json:"block_height"`
	BlockHash    string `gorm:"unique;not null" json:"block_hash"`
	Header       []byte `json:"header"`
	Difficulty   uint32 `json:"difficulty"`
	RandomNumber uint32 `json:"random_number"`
	MerkleRoot   string `json:"merkle_root"`
	BlockTime    int64  `json:"block_time"`
	TxHashes     string `json:"tx_hashes"`
}

type BtcTXOutput struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	BlockID  uint   `json:"block_data_id"`
	TxHash   string `json:"tx_hash"`
	Value    uint64 `json:"value"`
	PkScript []byte `json:"pk_script"`
}

func (dm *DatabaseManager) autoMigrate() {
	if err := dm.relayerDb.AutoMigrate(&BTCTransaction{}, &EVMSyncStatus{}, &WithdrawalRecord{}, &SubmitterRotation{}, &Participant{}); err != nil {
		log.Fatalf("Failed to migrate database 1: %v", err)
	}
	if err := dm.btcLightDb.AutoMigrate(&BtcBlock{}); err != nil {
		log.Fatalf("Failed to migrate database 3: %v", err)
	}
	if err := dm.btcCacheDb.AutoMigrate(&BtcSyncStatus{}, &BtcBlockData{}, &BtcTXOutput{}); err != nil {
		log.Fatalf("Failed to migrate database 2: %v", err)
	}
}
