package db

import (
	"time"

	"gorm.io/gorm"
)

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
