package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"gorm.io/gorm"
)

type Operations struct {
	gorm.Model
	TssNonce  uint64 `gorm:"index"`
	Signature []byte
	Hash      common.Hash
	DataHash  common.Hash
	Data      string
}

func (*Operations) TableName() string {
	return "operations"
}

func (o *Operations) TaskID() uint64 {
	return o.TssNonce
}

func (o *Operations) Type() int {
	return TaskTypeOperations
}

type TaskOperations struct {
	Operation []contracts.TaskOperation
}
