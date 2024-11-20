package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParticipantsThreshold(t *testing.T) {
	assert.Equal(t, CalculateThreshold(3), 1)   // (1 + 1) / 3 = 2/3
	assert.Equal(t, CalculateThreshold(4), 2)   // (2 + 1) / 4 = 3/4 > 2/3
	assert.Equal(t, CalculateThreshold(5), 3)   // (3 + 1) / 5 = 4/5 > 2/3
	assert.Equal(t, CalculateThreshold(6), 3)   // (3 + 1) / 6 = 4/6 = 2/3
	assert.Equal(t, CalculateThreshold(7), 4)   // (4 + 1) / 7 = 5/7 > 2/3
	assert.Equal(t, CalculateThreshold(8), 5)   // (5 + 1) / 8 = 6/8 > 2/3
	assert.Equal(t, CalculateThreshold(9), 5)   // (5 + 1) / 9 = 6/9 = 2/3
	assert.Equal(t, CalculateThreshold(10), 6)  // (6 + 1) / 10 = 7/10 > 2/3
	assert.Equal(t, CalculateThreshold(11), 7)  // (7 + 1) / 11 = 8/11 > 2/3
	assert.Equal(t, CalculateThreshold(12), 7)  // (7 + 1) / 12 = 8/12 = 2/3
	assert.Equal(t, CalculateThreshold(13), 8)  // (8 + 1) / 13 = 9/13 > 2/3
	assert.Equal(t, CalculateThreshold(15), 9)  // (9 + 1) / 15 = 10/15 = 2/3
	assert.Equal(t, CalculateThreshold(20), 13) // (13 + 1) / 20 = 14/20 = 7/10 > 2/3
}
