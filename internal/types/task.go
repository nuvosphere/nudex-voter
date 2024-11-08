package types

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
	TaskId uint32 `json:"task_id"`
}

func (t BaseTask) GetTaskID() uint32 {
	return t.TaskId
}

type CreateWalletTask struct {
	BaseTask
	User    string `json:"user"`
	Account uint64 `json:"account"`
	Chain   uint8  `json:"chain"` // evm_tss btc solana sui
	Index   uint32 `json:"index"`
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
	Decimal         uint32 `json:"decimal"`
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
	Decimal         uint32 `json:"decimal"`
	Fee             uint64 `json:"fee"`
}

const (
	WalletTypeEVM = iota
	WalletTypeBTC
	WalletTypeSOL
	WalletTypeSUI
)
