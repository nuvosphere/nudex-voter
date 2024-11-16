package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// CalculateThreshold (CalculateThreshold + 1) / length >= 2/3
func CalculateThreshold(length int) int {
	threshold := decimal.NewFromUint64(uint64(length)).
		Mul(decimal.NewFromInt(2)).
		Div(decimal.NewFromInt(3)).
		Ceil().
		IntPart() - 1

	for float64(threshold+1)/float64(length) < 2.0/3.0 {
		threshold++
	}

	return int(threshold)
}

type Participants []common.Address

func (p Participants) Threshold() int {
	return CalculateThreshold(len(p))
}

func (p Participants) Len() int {
	return len(p)
}
