package withdrawal

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func BuildTransaction(detailTask pool.Task[uint64]) (*ethereumTypes.Transaction, error) {
	switch task := detailTask.(type) {
	case *db.WithdrawalTask:
		targetAddress := common.HexToAddress(task.TargetAddress)
		toAddress := common.HexToAddress("0x00")

		var calldata []byte
		value := big.NewInt(0)
		if task.AssetType == types.AssetTypeMain {
			toAddress = targetAddress
			value = big.NewInt(int64(task.Amount))
		} else {
			// from = system address
			fromAddress := common.HexToAddress("0x00")
			calldata = contracts.EncodeTransferOfERC20(fromAddress, toAddress, big.NewInt(int64(task.Amount)))
		}
		baseTx := &ethereumTypes.DynamicFeeTx{
			To:    &toAddress,
			Value: value,
			Data:  calldata,
		}
		return ethereumTypes.NewTx(baseTx), nil
	default:
		return nil, fmt.Errorf("unknown transaction task:%d", detailTask.TaskID())
	}
}
