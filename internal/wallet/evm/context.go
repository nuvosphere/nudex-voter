package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
)

type TxContext struct {
	dbTX   *db.EvmTransaction
	sig    []byte
	notify chan error
	ctx    context.Context
	cancel context.CancelFunc
}

func (t *TxContext) TxHash() common.Hash {
	return t.dbTX.TxHash
}

func (t *TxContext) SeqID() uint64 {
	switch t.dbTX.Type {
	case db.TaskTypeConsolidation:
		return t.dbTX.EvmConsolidation.TaskID
	case db.TaskTypeWithdrawal:
		return t.dbTX.EvmWithdraw.TaskID
	case db.TaskTypeOperations:
		return t.dbTX.Operations.Nonce.BigInt().Uint64()
	default:
		panic("unknown task type")
	}
}

func (t *TxContext) IsSig() bool {
	return t.sig != nil
}

func (t *TxContext) UnSignTx() *types.Transaction {
	return t.dbTX.Tx()
}

func (t *TxContext) SignType() string {
	switch t.dbTX.Type {
	case db.TaskTypeConsolidation, db.TaskTypeWithdrawal:
		return vtypes.SignTxSessionType
	case db.TaskTypeOperations:
		return vtypes.SignOperationSessionType
	default:
		panic("unknown task type")
	}
}
