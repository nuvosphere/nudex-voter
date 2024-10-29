package types

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdraw
)

type CreateWalletTask struct {
	TaskId  int32 `json:"task_id"`
	User    int32 `json:"user"`
	Account int32 `json:"account"`
	Chain   int32 `json:"chain"` // evm_tss btc solana sui
}

const (
	WalletTypeEVM = iota
	WalletTypeBTC = iota
	WalletTypeSOL = iota
	WalletTypeSUI = iota
)
