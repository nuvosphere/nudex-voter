package db

import (
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"gorm.io/gorm"
)

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdrawal
)

const (
	Created = iota
	Pending
	Completed
	Failed
)

type DetailTask interface {
	pool.Task[uint64]
	SetBaseTask(task Task)
}

type Task struct {
	gorm.Model
	TaskId           uint64            `gorm:"unique;not null"                     json:"task_id"`
	TaskType         int               `gorm:"not null;default:0"                  json:"task_type"`
	Context          []byte            `gorm:"not null"                            json:"context"`
	Submitter        string            `gorm:"not null"                            json:"submitter"`
	Status           int               `gorm:"not null;default:0"                  json:"status"` // 0:Created; 1:pending; 2:Completed; 3:Failed
	LogIndex         LogIndex          `gorm:"foreignKey:ForeignID"`                              // has one https://gorm.io/zh_CN/docs/has_one.html
	CreateWalletTask *CreateWalletTask `gorm:"foreignKey:TaskId;references:TaskId"`
	DepositTask      *DepositTask      `gorm:"foreignKey:TaskId;references:TaskId"`
	WithdrawalTask   *WithdrawalTask   `gorm:"foreignKey:TaskId;references:TaskId"`
}

func (*Task) TableName() string {
	return "task"
}

func (t *Task) Type() int {
	return t.TaskType
}

func (t *Task) TaskID() uint64 {
	return t.TaskId
}

func (t *Task) DetailTask() DetailTask {
	var c DetailTask

	switch t.TaskType {
	case TaskTypeCreateWallet:
		c = t.CreateWalletTask
	case TaskTypeDeposit:
		c = t.DepositTask
	case TaskTypeWithdrawal:
		c = t.WithdrawalTask
	default:
		panic("DetailTask: unhandled default case")
	}

	c.SetBaseTask(*t)

	return c
}

type BaseTask struct {
	gorm.Model
	TaskType int    `gorm:"not null;default:0"                  json:"task_type"`
	TaskId   uint64 `gorm:"unique;not null"                     json:"task_id"`
	Task     Task   `gorm:"foreignKey:TaskId;references:TaskId"`
}

func (t *BaseTask) Type() int {
	return t.TaskType
}

func (t *BaseTask) TaskID() uint64 {
	return t.TaskId
}

func (t *BaseTask) SetBaseTask(task Task) {
	t.Task = task
}

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

func NewCreateWalletTask(taskId uint64, req *contracts.TaskPayloadContractWalletCreationRequest) *CreateWalletTask {
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

func NewDepositTask(taskId uint64, req *contracts.TaskPayloadContractDepositRequest) *DepositTask {
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

func NewWithdrawalTask(taskId uint64, req *contracts.TaskPayloadContractWithdrawalRequest) *WithdrawalTask {
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
	TaskVersionInitial = iota
	TaskVersionV1
)

const (
	TaskErrorCodeSuccess = iota
	TaskErrorCodePending
	TaskErrorCodeChainNotSupported
	TaskErrorCodeAssetNotSupported
	TaskErrorCodeCheckWithdrawalTxFailed
	TaskErrorCodeCheckWithdrawalInscriptionFailed
	TaskErrorCodeCheckWithdrawalBalanceFailed
)

type TaskUpdatedEvent struct {
	gorm.Model
	TaskId     uint64   `gorm:"unique;not null"                     json:"task_id"`
	Submitter  string   `gorm:"not null"                            json:"submitter"`
	UpdateTime int64    `gorm:"unique;not null"                     json:"completed_at"`
	State      uint8    `gorm:"state"                               json:"state"`
	Result     []byte   `json:"result"`
	Task       Task     `gorm:"foreignKey:TaskId;references:TaskId"`
	LogIndex   LogIndex `gorm:"foreignKey:ForeignID"` // has one https://gorm.io/zh_CN/docs/has_one.html
}
