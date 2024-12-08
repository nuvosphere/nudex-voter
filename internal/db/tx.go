package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EvmTransaction struct {
	gorm.Model
	CalldataHash common.Hash     `gorm:"index;size:256"                         json:"calldataHash"` // ta data hash
	Calldata     []byte          `json:"calldata"`                                                   // tx data
	TxNonce      decimal.Decimal `gorm:"index:sender_nonce"                     json:"txNonce"`      // tx nonce
	TxHash       common.Hash     `gorm:"uniqueIndex;size:256"                   json:"txHash"`       // tx hash
	TxJsonData   []byte          `json:"tx"`                                                         // blockchain origin tx of json format
	Sender       common.Address  `gorm:"index:sender_nonce; size:160"           json:"sender"`       // tx sender
	BuildHeight  uint64          `json:"buildHeight"`
	Status       int             `json:"status"` // 0: new，1:booked
	Error        string          `json:"error"`
}

func (EvmTransaction) TableName() string {
	return "evm_transactions"
}
