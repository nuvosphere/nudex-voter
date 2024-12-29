package db

import (
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"gorm.io/gorm"
)

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdrawal
	TaskTypeConsolidation
	TaskTypeOperations
)

const (
	Created = iota
	Pending
	Completed
	Failed
)

type DetailTask interface {
	types.ChainType
	pool.Task[uint64]
	SetBaseTask(task Task)
	Status() int
}

type Task struct {
	gorm.Model
	TaskId           uint64            `gorm:"unique;not null"                     json:"task_id"`
	TaskType         int               `gorm:"not null;default:0"                  json:"task_type"`
	Context          []byte            `gorm:"not null"                            json:"context"`
	Submitter        string            `gorm:"not null"                            json:"submitter"`
	State            int               `gorm:"not null;default:0"                  json:"status"` // 0:Created; 1:pending; 2:Completed; 3:Failed
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

func (t *Task) ChainType() uint8 {
	return t.DetailTask().ChainType()
}

func (t *Task) Status() int {
	return t.State
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

func (t *BaseTask) Status() int {
	return t.Task.Status()
}

type CreateWalletTask struct {
	BaseTask
	Account uint32 `json:"account"`
	Chain   uint8  `json:"chain"` // evm_tss btc solana sui
	Index   uint32 `json:"index"`
	Address string `json:"address"` // new create bip44 address
}

func (*CreateWalletTask) TableName() string {
	return "create_wallet_task"
}

func (c *CreateWalletTask) ChainType() uint8 {
	return c.Chain
}

func NewCreateWalletTask(taskId uint64, req *contracts.TaskPayloadContractWalletCreationRequest) *CreateWalletTask {
	return &CreateWalletTask{
		BaseTask: BaseTask{
			TaskId:   taskId,
			TaskType: TaskTypeCreateWallet,
		},
		Account: req.Account,
		Chain:   req.AddressType,
		Index:   uint32(req.Index),
	}
}

type DepositTask struct {
	BaseTask
	TargetAddress   string       `json:"target_address"`
	Amount          uint64       `json:"amount"`
	Chain           uint8        `json:"chain"`
	ChainId         types.Byte32 `json:"chain_id"`
	BlockHeight     uint64       `json:"block_height"`
	TxHash          string       `json:"tx_hash"`
	ContractAddress string       `json:"contract_address"`
	Ticker          string       `json:"ticker"`
	AssetType       uint8        `json:"asset_type"`
	Decimal         uint8        `json:"decimal"`
}

func (c *DepositTask) Status() int {
	// TODO implement me
	panic("implement me")
}

func (*DepositTask) TableName() string {
	return "deposit_task"
}

func (c *DepositTask) ChainType() uint8 {
	return c.Chain
}

func NewDepositTask(taskId uint64, req *contracts.TaskPayloadContractDepositRequest) *DepositTask {
	return &DepositTask{
		BaseTask: BaseTask{
			TaskId:   taskId,
			TaskType: TaskTypeDeposit,
		},
		TargetAddress:   req.UserTssAddress,
		Amount:          req.Amount,
		ChainId:         req.ChainId,
		TxHash:          req.TxHash,
		ContractAddress: req.ContractAddress,
		Ticker:          req.Ticker,
		AssetType:       req.AssetType,
		Decimal:         req.Decimal,
	}
}

type WithdrawalTask struct {
	BaseTask
	TargetAddress   string       `json:"target_address"`
	Amount          uint64       `json:"amount"`
	Chain           uint8        `json:"chain"`
	ChainId         types.Byte32 `json:"chain_id"`
	BlockHeight     uint64       `json:"block_height"`
	TxHash          string       `json:"tx_hash"`
	ContractAddress string       `json:"contract_address"`
	Ticker          types.Byte32 `json:"ticker"`
	AssetType       uint8        `json:"asset_type"`
	Decimal         uint8        `json:"decimal"`
	Fee             uint64       `json:"fee"`
}

func (*WithdrawalTask) TableName() string {
	return "withdrawal_task"
}

func (c *WithdrawalTask) ChainType() uint8 {
	return c.Chain
}

func NewWithdrawalTask(taskId uint64, req *contracts.TaskPayloadContractWithdrawalRequest) *WithdrawalTask {
	return &WithdrawalTask{
		BaseTask: BaseTask{
			TaskId:   taskId,
			TaskType: TaskTypeWithdrawal,
		},
		TargetAddress:   req.UserTssAddress,
		Amount:          req.Amount,
		ChainId:         req.ChainId,
		TxHash:          req.TxHash,
		ContractAddress: req.ContractAddress,
		Ticker:          req.Ticker,
		AssetType:       req.AssetType,
		Decimal:         req.Decimal,
		Fee:             req.Fee,
	}
}

type ConsolidationTask struct {
	BaseTask
	TargetAddress   string       `json:"target_address"`
	Amount          uint64       `json:"amount"`
	Chain           uint8        `json:"chain"`
	ChainId         types.Byte32 `json:"chain_id"`
	BlockHeight     uint64       `json:"block_height"`
	TxHash          string       `json:"tx_hash"`
	ContractAddress string       `json:"contract_address"`
	Ticker          string       `json:"ticker"`
	AssetType       uint8        `json:"asset_type"`
	Decimal         uint8        `json:"decimal"`
	Fee             uint64       `json:"fee"`
}

func (*ConsolidationTask) TableName() string {
	return "consolidation_task"
}

func (c *ConsolidationTask) ChainType() uint8 {
	return c.Chain
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
	TaskErrorCodeCheckTxFailed
	TaskErrorCodeCheckInscriptionFailed
	TaskErrorCodeCheckAmountFailed
	TaskErrorCodeCheckAssetFailed
	TaskErrorCodeDepositAssetNotEnabled
	TaskErrorCodeDepositAmountTooLow
	TaskErrorCodeDepositTokenNotSupported
	TaskErrorCodeDepositTokenNotActive
	TaskErrorCodeWithdrawalAssetNotEnabled
	TaskErrorCodeWithdrawalAmountTooLow
	TaskErrorCodeWithdrawalTokenNotSupported
	TaskErrorCodeWithdrawalTokenNotActive
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

func (t *TaskUpdatedEvent) TaskID() uint64 {
	return t.TaskId
}

func (t *TaskUpdatedEvent) Type() int {
	return t.Task.Type()
}

func (t *TaskUpdatedEvent) ChainType() uint8 {
	return t.Task.ChainType()
}

func (t *TaskUpdatedEvent) Status() int {
	return int(t.State)
}

func (t *TaskUpdatedEvent) SetBaseTask(task Task) {
	t.Task = task
}
