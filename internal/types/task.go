package types

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdraw
)

type Task interface {
	GetTaskID() int32
}

type BaseTask struct {
	TaskId int32 `json:"task_id"`
}

func (t BaseTask) GetTaskID() int32 {
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
	Address   string `json:"address"`
	Amount    uint64 `json:"amount"`
	ChainId   uint64 `json:"chain_id"`
	Token     []byte `json:"token"`
	TxInfo    []byte `json:"tx_info"`
	ExtraInfo []byte `json:"extra_info"`
}

type WithdrawalTask struct {
	BaseTask
	Address   string `json:"address"`
	Amount    uint64 `json:"amount"`
	ChainId   uint64 `json:"chain_id"`
	Token     []byte `json:"token"`
	TxInfo    []byte `json:"tx_info"`
	ExtraInfo []byte `json:"extra_info"`
}

const (
	WalletTypeEVM = iota
	WalletTypeBTC
	WalletTypeSOL
	WalletTypeSUI
)
