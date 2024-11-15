package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// CalculateThreshold  2/3.
func CalculateThreshold(length int) int {
	threshold := decimal.NewFromUint64(uint64(length)).
		Mul(decimal.NewFromInt(2)).
		Div(decimal.NewFromInt(3)).
		Ceil().
		IntPart()

	return int(threshold)
}

type Participants []common.Address

func (p Participants) Threshold() int {
	return CalculateThreshold(len(p))
}

func (p Participants) Len() int {
	return len(p)
}
