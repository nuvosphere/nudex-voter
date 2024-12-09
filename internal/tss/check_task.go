package tss

import (
	"fmt"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

func (m *Scheduler) checkTask(task pool.Task[uint64]) (int, error) {
	switch taskData := task.(type) {
	// @todo check tx
	case *db.WithdrawalTask:
		switch taskData.Chain {
		case types.CoinTypeBTC:
			switch taskData.AssetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				return db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainEthereum:
			switch taskData.AssetType {
			case types.AssetTypeMain:
				// @todo check main asset
			case types.AssetTypeErc20:
				// @todo check erc20 asset
			default:
				return db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainSolana:
			switch taskData.AssetType {
			case types.AssetTypeMain:
			case types.AssetTypeErc20:
			default:
				return db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainSui:
			switch taskData.AssetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				return db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		default:
			return db.TaskErrorCodeChainNotSupported, fmt.Errorf("unknown Chain type: %v", taskData.Chain)
		}
	default:
		panic(fmt.Errorf("error pending task id: %v", task.TaskID()))
	}
	return db.TaskErrorCodePending, nil
}
