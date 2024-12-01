package db

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type LogIndex struct {
	gorm.Model
	ContractAddress common.Address `gorm:"index;size:160"                json:"contract_address"`
	EventName       string         `json:"eventName"`                                         // event name
	Log             *types.Log     `gorm:"serializer:json"               json:"log"`          // event content
	TxHash          common.Hash    `gorm:"index;size:256"                json:"tx_hash"`      // tx hash
	ChainId         uint64         `gorm:"index:log_index_unique,unique" json:"chain_id"`     // chainId
	BlockNumber     uint64         `gorm:"index:log_index_unique,unique" json:"block_number"` // block number of the tx
	LogIndex        uint64         `gorm:"index:log_index_unique,unique" json:"log_index"`    // block log index
	ForeignID       uint           `gorm:"index"                         json:"foreign_id"`   // task table ID;submitter table ID;participant_event table ID;...
}

func (LogIndex) TableName() string {
	return "log_index"
}

type EVMSyncStatus struct {
	gorm.Model
	LastSyncBlock uint64 `gorm:"not null" json:"last_sync_block"`
}

func (EVMSyncStatus) TableName() string {
	return "evm_sync_status"
}

// SubmitterChosen contains block number and current submitter.
type SubmitterChosen struct {
	gorm.Model
	Submitter   string   `gorm:"index:submitter_block_number_unique,unique" json:"submitter"`
	BlockNumber uint64   `gorm:"index:submitter_block_number_unique,unique" json:"block_number"`
	LogIndex    LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (SubmitterChosen) TableName() string {
	return "submitter"
}

// Participant save all participants.
type Participant struct {
	gorm.Model
	Address string `gorm:"uniqueIndex;not null" json:"address"`
}

func (Participant) TableName() string {
	return "participant"
}

// ParticipantEvent save all participants.
type ParticipantEvent struct {
	gorm.Model
	EventName   string   `json:"eventName"` // event name
	Address     string   `gorm:"index;not null"       json:"address"`
	BlockNumber uint64   `gorm:"index;not null"       json:"block_number"`
	LogIndex    LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (ParticipantEvent) TableName() string {
	return "participant_event"
}

// Account save all accounts.
type Account struct {
	gorm.Model
	User     string   `gorm:"not null"              json:"user"`
	Account  uint64   `gorm:"not null"              json:"account"`
	Chain    uint8    `gorm:"not null"              json:"chain"`
	Index    uint64   `gorm:"not null"              json:"index"`
	Address  string   `gorm:"uniqueIndex; not null" json:"address"`
	LogIndex LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (Account) TableName() string {
	return "account"
}

type DepositRecord struct {
	gorm.Model
	TargetAddress string   `gorm:"not null"             json:"target_address"`
	Amount        uint64   `gorm:"not null"             json:"amount"`
	ChainId       uint64   `gorm:"not null"             json:"chain_id"`
	TxInfo        []byte   `gorm:"not null"             json:"tx_info"`
	ExtraInfo     []byte   `gorm:"not null"             json:"extra_info"`
	LogIndex      LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (DepositRecord) TableName() string {
	return "deposit_record"
}

type WithdrawalRecord struct {
	gorm.Model
	TargetAddress string   `gorm:"not null"             json:"target_address"`
	Amount        uint64   `gorm:"not null"             json:"amount"`
	ChainId       uint64   `gorm:"not null"             json:"chain_id"`
	TxInfo        []byte   `gorm:"not null"             json:"tx_info"`
	ExtraInfo     []byte   `gorm:"not null"             json:"extra_info"`
	LogIndex      LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (WithdrawalRecord) TableName() string {
	return "withdrawal_record"
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
