package tss

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/nuvosphere/nudex-voter/internal/wallet/solana"
)

type EvmTxContext struct {
	w    *wallet.Wallet
	tx   *ethtypes.Transaction
	task pool.Task[uint64]
}

type SolTxContext struct {
	c    *solana.SolClient
	tx   *solana.UnSignTx
	task pool.Task[uint64]
}

type BtcTxContext struct {
	c      types.TxClient
	sigCtx *SignContext
}
