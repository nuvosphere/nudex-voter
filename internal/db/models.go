package db

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EVMSyncStatus struct {
	gorm.Model
	LastSyncBlock uint64 `gorm:"not null" json:"last_sync_block"`
}

// SubmitterChosen contains block number and current submitter.
type SubmitterChosen struct {
	gorm.Model
	BlockNumber uint64 `gorm:"index:block_number_submitter,unique;not null" json:"block_number"`
	Submitter   string `gorm:"index:block_number_submitter,unique;not null" json:"submitter"`
}

// Participant save all participants.
type Participant struct {
	gorm.Model
	Address string `gorm:"uniqueIndex;not null" json:"address"`
}

// Account save all accounts.
type Account struct {
	gorm.Model
	User    string `gorm:"not null"              json:"user"`
	Account uint64 `gorm:"not null"              json:"account"`
	ChainId uint8  `gorm:"not null"              json:"chain_id"`
	Index   uint64 `gorm:"not null"              json:"index"`
	Address string `gorm:"uniqueIndex; not null" json:"address"`
}

type DepositRecord struct {
	gorm.Model
	TargetAddress string `gorm:"not null" json:"target_address"`
	Amount        uint64 `gorm:"not null" json:"amount"`
	ChainId       uint64 `gorm:"not null" json:"chain_id"`
	TxInfo        []byte `gorm:"not null" json:"tx_info"`
	ExtraInfo     []byte `gorm:"not null" json:"extra_info"`
}

type WithdrawalRecord struct {
	gorm.Model
	TargetAddress string `gorm:"not null" json:"target_address"`
	Amount        uint64 `gorm:"not null" json:"amount"`
	ChainId       uint64 `gorm:"not null" json:"chain_id"`
	TxInfo        []byte `gorm:"not null" json:"tx_info"`
	ExtraInfo     []byte `gorm:"not null" json:"extra_info"`
}

const (
	New = iota
	Pending
	Completed
	Other
)

type Task struct {
	gorm.Model
	TaskId      uint32 `gorm:"unique;not null"    json:"task_id"`
	Context     []byte `gorm:"not null"           json:"Context"`
	Submitter   string `gorm:"not null"           json:"submitter"`
	BlockHeight uint64 `gorm:"not null"           json:"block_height"`
	Status      int    `gorm:"not null;default:0" json:"status"` // 0:new; 1:pending; 2:Completed; 3:other
}

type BTCTransaction struct {
	gorm.Model
	TxID        string    `gorm:"uniqueIndex;not null" json:"tx_id"`
	RawTxData   string    `gorm:"type:text;not null"   json:"raw_tx_data"`
	ReceivedAt  time.Time `gorm:"not null"             json:"received_at"`
	Processed   bool      `gorm:"default:false"        json:"processed"`
	ProcessedAt time.Time `json:"processed_at"`
}

type BtcBlock struct {
	gorm.Model
	Height uint64 `gorm:"not null;uniqueIndex" json:"height"`
	Hash   string `gorm:"not null"             json:"hash"`
	Status string `gorm:"not null"             json:"status"` // "unconfirm", "confirmed", "signing", "pending", "processed"
}

type BtcSyncStatus struct {
	gorm.Model
	UnconfirmHeight int64 `gorm:"not null" json:"unconfirm_height"`
	ConfirmedHeight int64 `gorm:"not null" json:"confirmed_height"`
}

type BtcBlockData struct {
	gorm.Model
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
	gorm.Model
	BlockID  uint   `json:"block_data_id"`
	TxHash   string `json:"tx_hash"`
	Value    uint64 `json:"value"`
	PkScript []byte `json:"pk_script"`
}

func (dm *DatabaseManager) autoMigrate() {
	if err := dm.relayerDb.AutoMigrate(
		&BTCTransaction{},
		&EVMSyncStatus{},
		&SubmitterChosen{},
		&Participant{},
		&Account{},
		&DepositRecord{},
		&WithdrawalRecord{},
		&Task{},
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
