package db

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type EvmTransaction struct {
	gorm.Model
	TxHash     common.Hash     `gorm:"uniqueIndex;size:256"                   json:"txHash"`  // tx hash
	TxJsonData []byte          `json:"tx"`                                                    // blockchain origin tx of json format
	TxNonce    decimal.Decimal `gorm:"index:sender_nonce"                     json:"txNonce"` // tx nonce
	Sender     common.Address  `json:"sender"`
	Status     int             `json:"status"` // 0: new，1:booked
	Error      string          `json:"error"`
	Type       int             // operation、withdraw、consolidation
	SeqID      uint64
}

func (*EvmTransaction) TableName() string {
	return "evm_transactions"
}

func (e *EvmTransaction) Tx() *types.Transaction {
	tx := new(types.Transaction)
	err := json.Unmarshal(e.TxJsonData, tx)
	utils.Assert(err)
	return tx
}
