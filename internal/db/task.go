package db

import (
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"gorm.io/gorm"
)

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdrawal
)

const (
	New = iota
	Pending
	Completed
	Other
)

type ITask interface {
	Type() int
	TaskID() uint32
	SetBaseTask(task Task)
}

type Task struct {
	gorm.Model
	TaskId           uint32            `gorm:"unique;not null"                     json:"task_id"`
	TaskType         int               `gorm:"not null;default:0"                  json:"task_type"`
	Context          []byte            `gorm:"not null"                            json:"context"`
	Submitter        string            `gorm:"not null"                            json:"submitter"`
	Status           int               `gorm:"not null;default:0"                  json:"status"` // 0:new; 1:pending; 2:Completed; 3:other
	LogIndex         LogIndex          `gorm:"foreignKey:ForeignID"`                              // has one https://gorm.io/zh_CN/docs/has_one.html
	CreateWalletTask *CreateWalletTask `gorm:"foreignKey:TaskId;references:TaskId"`
	DepositTask      *DepositTask      `gorm:"foreignKey:TaskId;references:TaskId"`
	WithdrawalTask   *WithdrawalTask   `gorm:"foreignKey:TaskId;references:TaskId"`
}

func (Task) TableName() string {
	return "task"
}

type BaseTask struct {
	gorm.Model
	TaskType int    `gorm:"not null;default:0"                  json:"task_type"`
	TaskId   uint32 `gorm:"unique;not null"                     json:"task_id"`
	Task     Task   `gorm:"foreignKey:TaskId;references:TaskId"`
}

func (t *BaseTask) Type() int {
	return t.TaskType
}

func (t *BaseTask) TaskID() uint32 {
	return t.TaskId
}

func (t *BaseTask) SetBaseTask(task Task) {
	t.Task = task
}

//func (t *BaseTask) BaseTask() Task {
//	return t.Task
//}

type CreateWalletTask struct {
	BaseTask
	Account uint32 `json:"account"`
	Chain   uint8  `json:"chain"` // evm_tss btc solana sui
	Index   uint8  `json:"index"`
	Address string `json:"address"` // new create bip44 address
}

func (CreateWalletTask) TableName() string {
	return "create_wallet_task"
}

func NewCreateWalletTask(taskId uint32, req *contracts.TaskPayloadContractWalletCreationRequest) *CreateWalletTask {
	return &CreateWalletTask{
		BaseTask: BaseTask{
			TaskId:   taskId,
			TaskType: TaskTypeCreateWallet,
		},
		Account: req.Account,
		Chain:   req.Chain,
		Index:   req.Index,
	}
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

func (DepositTask) TableName() string {
	return "deposit_task"
}

func NewDepositTask(taskId uint32, req *contracts.TaskPayloadContractDepositRequest) *DepositTask {
	return &DepositTask{
		BaseTask: BaseTask{
			TaskId:   taskId,
			TaskType: TaskTypeDeposit,
		},
		TargetAddress:   req.TargetAddress,
		Amount:          req.Amount,
		Chain:           req.Chain,
		ChainId:         req.ChainId,
		BlockHeight:     req.BlockHeight,
		TxHash:          req.TxHash,
		ContractAddress: req.ContractAddress,
		Ticker:          req.Ticker,
		AssetType:       req.AssetType,
		Decimal:         req.Decimal,
	}
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

func (WithdrawalTask) TableName() string {
	return "withdrawal_task"
}

func NewWithdrawalTask(taskId uint32, req *contracts.TaskPayloadContractWithdrawalRequest) *WithdrawalTask {
	return &WithdrawalTask{
		BaseTask: BaseTask{
			TaskId:   taskId,
			TaskType: TaskTypeWithdrawal,
		},
		TargetAddress:   req.TargetAddress,
		Amount:          req.Amount,
		Chain:           req.Chain,
		ChainId:         req.ChainId,
		BlockHeight:     req.BlockHeight,
		TxHash:          req.TxHash,
		ContractAddress: req.ContractAddress,
		Ticker:          req.Ticker,
		AssetType:       req.AssetType,
		Decimal:         req.Decimal,
		Fee:             req.Fee,
	}
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

type TaskCompletedEvent struct {
	gorm.Model
	TaskId      uint32   `gorm:"unique;not null"                     json:"task_id"`
	Submitter   string   `gorm:"not null"                            json:"submitter"`
	CompletedAt int64    `gorm:"unique;not null"                     json:"completed_at"`
	Result      []byte   `json:"result"`
	Task        Task     `gorm:"foreignKey:TaskId;references:TaskId"`
	LogIndex    LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}
