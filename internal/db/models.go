package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type LogIndex struct {
	gorm.Model
	ContractAddress common.Address `gorm:"index;size:160"                json:"contract_address"`
	EventName       string         `json:"eventName"`                                         // event name
	Log             *types.Log     `gorm:"serializer:json"               json:"log"`          // event content
	TxHash          common.Hash    `gorm:"index;size:256"                json:"tx_hash"`      // tx hash
	ChainId         vtypes.Byte32  `gorm:"index:log_index_unique,unique" json:"chain_id"`     // chainId
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
	TargetAddress string        `gorm:"not null"             json:"target_address"`
	Amount        uint64        `gorm:"not null"             json:"amount"`
	ChainId       vtypes.Byte32 `gorm:"not null"             json:"chain_id"`
	TxInfo        []byte        `gorm:"not null"             json:"tx_info"`
	ExtraInfo     []byte        `gorm:"not null"             json:"extra_info"`
	LogIndex      LogIndex      `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (DepositRecord) TableName() string {
	return "deposit_record"
}

type WithdrawalRecord struct {
	gorm.Model
	DepositAddress string        `gorm:"not null"             json:"target_address"`
	Amount         uint64        `gorm:"not null"             json:"amount"`
	ChainId        vtypes.Byte32 `gorm:"not null"             json:"chain_id"`
	LogIndex       LogIndex      `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}

func (WithdrawalRecord) TableName() string {
	return "withdrawal_record"
}

type AddressBalance struct {
	gorm.Model
	Address string          `gorm:"uniqueIndex; not null"             json:"address"`
	Token   string          `gorm:"uniqueIndex; not null"             json:"token"`
	Amount  decimal.Decimal `gorm:"not null"                          json:"amount"`
	ChainId vtypes.Byte32   `gorm:"not null"                          json:"chain_id"`
}

func (AddressBalance) TableName() string {
	return "address_balance"
}
