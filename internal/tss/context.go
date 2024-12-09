package tss

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
)

type EvmTxContext struct {
	w    *wallet.Wallet
	tx   *ethtypes.Transaction
	task pool.Task[uint64]
}
