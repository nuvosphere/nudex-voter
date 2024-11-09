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

type Chain uint8

const (
	ETHEREUM Chain = iota
	BITCOIN
	SOLANA
	SUI
)

type AssetType uint32

const (
	MAIN AssetType = iota
	ERC20
)

type Version uint8

const (
	INITIAL Version = iota
	V1
)

type ErrorCode uint8

const (
	SUCCESS ErrorCode = iota
)
