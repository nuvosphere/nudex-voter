package tss

import (
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
)

func (tss *TSSService) handleTssKeyEnd(event *keygen.LocalPartySaveData) error {
	if tss.Party == nil {
		return fmt.Errorf("handleTssEnd error, event %v, self not init", event)
	}
	return saveTSSData(event)
}
