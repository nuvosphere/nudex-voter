package db

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Operations struct {
	gorm.Model
	Nonce  decimal.Decimal `gorm:"index"`
	Data   string
	Status int
}

func (*Operations) TableName() string {
	return "operations"
}
