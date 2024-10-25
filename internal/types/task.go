package types

const (
	TaskTypeUnknown = iota
	TaskTypeCreateWallet
	TaskTypeDeposit
	TaskTypeWithdraw
)

type CreateWalletTask struct {
	TaskId  uint64 `json:"task_id"`
	User    string `json:"user"`
	Account uint64 `json:"account"`
	Chain   string `json:"chain"` // evm_tss btc solana sui
}
