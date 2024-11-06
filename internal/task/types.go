package task

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss"
)

type TaskService struct {
	state *state.State
	dbm   *db.DatabaseManager
	Tss   *tss.TSSService
}

type Chain int

const (
	ETHEREUM Chain = iota
	BITCOIN
	SOLANA
	SUI
)

type AssetType int

const (
	MAIN AssetType = iota
	ERC20
)
