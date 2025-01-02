package layer2

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

var (
	WalletCreationReq = contracts.MethodID(contracts.AccountManagerContractMetaData, "registerNewAddress")
	DepositReq        = contracts.MethodID(contracts.DepositManagerContractMetaData, "recordDeposit")
	WithdrawalReq     = contracts.MethodID(contracts.DepositManagerContractMetaData, "recordWithdrawal")
)

func DecodeTask(taskId uint64, context []byte) db.DetailTask {
	context = context[4:]

	contextHex := hexutil.Encode(context)
	switch contextHex {
	case WalletCreationReq:
		request := &db.WalletCreationRequest{}
		err := contracts.NewContract(contracts.AccountManagerContractMetaData).UnPackInput("registerNewAddress", request, context)
		utils.Assert(err)

		return db.NewCreateWalletTask(taskId, request)
	case DepositReq:
		request := &db.DepositRequest{}
		err := contracts.NewContract(contracts.DepositManagerContractMetaData).UnPackInput("recordDeposit", request, context)
		utils.Assert(err)

		return db.NewDepositTask(taskId, request)
	case WithdrawalReq:
		request := &db.WithdrawalRequest{}
		err := contracts.NewContract(contracts.DepositManagerContractMetaData).UnPackInput("recordWithdrawal", request, context)
		utils.Assert(err)

		return db.NewWithdrawalTask(taskId, request)
	}

	return nil
}
