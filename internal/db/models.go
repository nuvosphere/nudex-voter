package db

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type BTCTransaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TxID        string    `gorm:"uniqueIndex;not null" json:"tx_id"`
	RawTxData   string    `gorm:"type:text;not null" json:"raw_tx_data"`
	ReceivedAt  time.Time `gorm:"not null" json:"received_at"`
	Processed   bool      `gorm:"default:false" json:"processed"`
	ProcessedAt time.Time `json:"processed_at"`
}

type EVMSyncStatus struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	LastSyncBlock uint64    `gorm:"not null" json:"last_sync_block"`
	UpdatedAt     time.Time `gorm:"not null" json:"updated_at"`
}

type WithdrawalRecord struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	WithdrawalID string    `gorm:"uniqueIndex;not null" json:"withdrawal_id"`
	UserAddress  string    `gorm:"not null" json:"user_address"`
	Amount       string    `gorm:"not null" json:"amount"`
	DetectedAt   time.Time `gorm:"not null" json:"detected_at"`
	OnChain      bool      `gorm:"default:false" json:"on_chain"`
	OnChainTxID  string    `json:"on_chain_tx_id"`
	Processed    bool      `gorm:"default:false" json:"processed"`
	ProcessedAt  time.Time `json:"processed_at"`
}

// SubmitterChosen contains block number and current submitter
type SubmitterChosen struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	BlockNumber uint64 `gorm:"not null" json:"block_number"`
	Submitter   string `gorm:"not null" json:"submitter"`
}

// Participant save all participants
type Participant struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Address string `gorm:"uniqueIndex;not null" json:"address"`
}

// Account save all accounts
type Account struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	User    string `gorm:"not null" json:"user"`
	Account uint64 `gorm:"not null" json:"account"`
	ChainId uint8  `gorm:"not null" json:"chain_id"`
	Index   uint64 `gorm:"not null" json:"index"`
	Address string `gorm:"not null" json:"address"`
}

type DepositRecord struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	TargetAddress string `gorm:"not null" json:"target_address"`
	Amount        uint64 `gorm:"not null" json:"amount"`
	ChainId       uint64 `gorm:"not null" json:"chain_id"`
	TxInfo        []byte `gorm:"not null" json:"tx_info"`
	ExtraInfo     []byte `gorm:"not null" json:"extra_info"`
}

type Task struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	TaskId      uint64    `gorm:"unique;not null" json:"task_id"`
	Context     []byte    `gorm:"not null" json:"Context"`
	Submitter   string    `gorm:"not null" json:"submitter"`
	BlockHeight uint64    `gorm:"not null" json:"block_height"`
	IsCompleted bool      `gorm:"not null" json:"is_completed"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
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
	if err := dm.relayerDb.AutoMigrate(&BTCTransaction{}, &EVMSyncStatus{}, &WithdrawalRecord{}, &SubmitterChosen{},
		&Participant{}, &Account{}, &DepositRecord{}, &Task{}); err != nil {
		log.Fatalf("Failed to migrate database 1: %v", err)
	}
	if err := dm.btcLightDb.AutoMigrate(&BtcBlock{}); err != nil {
		log.Fatalf("Failed to migrate database 3: %v", err)
	}
	if err := dm.btcCacheDb.AutoMigrate(&BtcSyncStatus{}, &BtcBlockData{}, &BtcTXOutput{}); err != nil {
		log.Fatalf("Failed to migrate database 2: %v", err)
	}
}
