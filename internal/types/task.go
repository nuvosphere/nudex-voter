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

type DepositTask struct {
	TaskId    int32  `json:"task_id"`
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
