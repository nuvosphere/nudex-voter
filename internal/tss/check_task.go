package tss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

type TxStatusResponse struct {
	Status string `json:"status"`
}

func (m *Scheduler) checkTask(task pool.Task[uint64]) (bool, int, error) {
	switch taskData := task.(type) {
	case *db.WithdrawalTask:
		hashCheckStatus, err := checkTxStatus(taskData.TxHash)
		if err != nil {
			return false, -1, err
		}
		if hashCheckStatus != "success" {
			return false, -1, err
		}

		// @todo check inscription info

		switch taskData.Chain {
		case types.CoinTypeBTC:
			switch taskData.AssetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainEthereum:
			switch taskData.AssetType {
			case types.AssetTypeMain:
				// @todo check main asset
			case types.AssetTypeErc20:
				// @todo check erc20 asset
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainSolana:
			switch taskData.AssetType {
			case types.AssetTypeMain:
			case types.AssetTypeErc20:
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainSui:
			switch taskData.AssetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				return true, db.TaskErrorCodeAssetNotSupported, fmt.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		default:
			return true, db.TaskErrorCodeChainNotSupported, fmt.Errorf("unknown Chain type: %v", taskData.Chain)
		}
	default:
		panic(fmt.Errorf("error pending task id: %v", task.TaskID()))
	}
	return false, db.TaskErrorCodePending, nil
}
func checkTxStatus(txHash string) (string, error) {
	url := fmt.Sprintf("https://nip-api.testnet.nudex.io/%s/status", txHash)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var txStatus TxStatusResponse
	if err := json.Unmarshal(body, &txStatus); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	return txStatus.Status, nil
}
