package task

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"gorm.io/gorm"
)

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdrawal
)

type Task interface {
	GetTaskID() uint32
}

type BaseTask struct {
	gorm.Model
	TaskId uint32 `gorm:"unique;not null" json:"task_id"`
}

func (t BaseTask) GetTaskID() uint32 {
	return t.TaskId
}

type CreateWalletTask struct {
	BaseTask
	User    string `json:"user"`
	Account uint32 `json:"account"`
	Chain   uint8  `json:"chain"` // evm_tss btc solana sui
	Index   uint8  `json:"index"`
}

type DepositTask struct {
	BaseTask
	TargetAddress   string `json:"target_address"`
	Amount          uint64 `json:"amount"`
	Chain           uint8  `json:"chain"`
	ChainId         uint32 `json:"chain_id"`
	BlockHeight     uint64 `json:"block_height"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Ticker          string `json:"ticker"`
	AssetType       uint8  `json:"asset_type"`
	Decimal         uint8  `json:"decimal"`
}

type WithdrawalTask struct {
	BaseTask
	TargetAddress   string `json:"target_address"`
	Amount          uint64 `json:"amount"`
	Chain           uint8  `json:"chain"`
	ChainId         uint32 `json:"chain_id"`
	BlockHeight     uint64 `json:"block_height"`
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Ticker          string `json:"ticker"`
	AssetType       uint8  `json:"asset_type"`
	Decimal         uint8  `json:"decimal"`
	Fee             uint64 `json:"fee"`
}

const (
	WalletTypeEVM = iota
	WalletTypeBTC
	WalletTypeSOL
	WalletTypeSUI
)

const (
	ChainEthereum = iota
	ChainBitcoin
	ChainSolana
	ChainSui
)

const (
	AssetTypeMain = iota
	AssetTypeErc20
)

const (
	TaskVersionInitial = iota
	TaskVersionV1
)

const (
	TaskErrorCodeSuccess = iota
	TaskErrorCodeChainNotSupported
)

type ParticipantPair struct {
	Old []common.Address
	New []common.Address
}

type SubmitterChosenPair struct {
	Old db.SubmitterChosen
	New db.SubmitterChosen
}
