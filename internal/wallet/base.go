package wallet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nuvosphere/nudex-voter/internal/codec"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/patrickmn/go-cache"
)

type BaseWallet struct {
	stateDB            *state.ContractState
	voterContract      layer2.VoterContract
	discussedTaskCache *cache.Cache
	taskQueue          *pool.Pool[uint64] // receive l2 task
}

func NewBaseWallet(stateDB *state.ContractState, voterContract layer2.VoterContract) *BaseWallet {
	return &BaseWallet{
		stateDB:            stateDB,
		voterContract:      voterContract,
		discussedTaskCache: cache.New(5*time.Minute, 10*time.Minute),
		taskQueue:          pool.NewTaskPool[uint64](),
	}
}

func (w *BaseWallet) IsProd() bool {
	return false
}

func (w *BaseWallet) VoterContract() layer2.VoterContract {
	return w.voterContract
}

func (w *BaseWallet) AddTask(task pool.Task[uint64]) {
	w.taskQueue.Add(task)
}

func (w *BaseWallet) RemoveTask(taskId uint64) {
	w.taskQueue.Remove(taskId)
}

func (w *BaseWallet) GetTask(taskID uint64) (pool.Task[uint64], error) {
	t := w.taskQueue.Get(taskID)
	if t != nil {
		return t, nil
	}

	task, err := w.stateDB.GetUnCompletedTask(taskID)
	//todo
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return w.GetOnlineTask(taskID)
	//}
	if err != nil {
		return nil, fmt.Errorf("taskID:%v, %w", taskID, err)
	}

	return task.DetailTask(), err
}

func (w *BaseWallet) GetOnlineTask(taskId uint64) (pool.Task[uint64], error) {
	t, err := w.voterContract.Tasks(taskId)
	if err != nil {
		return nil, err
	}

	detailTask := codec.DecodeTask(t.Id, t.Context)

	baseTask := db.Task{
		TaskId:    t.Id,
		TaskType:  detailTask.Type(),
		Context:   t.Context,
		Submitter: t.Submitter.Hex(),
		State:     int(t.State),
	}
	detailTask.SetBaseTask(baseTask)

	return detailTask, nil
}

func (w *BaseWallet) IsDiscussed(taskID uint64) bool {
	_, ok := w.discussedTaskCache.Get(fmt.Sprintf("%d", taskID))
	if !ok {
		ok, _ = w.voterContract.IsTaskCompleted(taskID)
	}

	return ok
}

func (w *BaseWallet) AddDiscussedTask(taskID uint64) {
	w.discussedTaskCache.SetDefault(fmt.Sprintf("%d", taskID), struct{}{})
}

func (w *BaseWallet) CheckTask(task pool.Task[uint64]) (bool, int, error) {
	switch taskData := task.(type) {
	case *db.DepositTask:
		hashCheckStatus, err := checkTxStatus(taskData.TxHash)
		if err != nil {
			return false, -1, err
		}
		if hashCheckStatus != "success" {
			return true, db.TaskErrorCodeCheckTxFailed, err
		}
		inscriptionMintb, err := w.stateDB.GetInscriptionMintb(taskData.TxHash)
		if err != nil || inscriptionMintb == nil {
			return true, db.TaskErrorCodeCheckInscriptionFailed, err
		}
		if inscriptionMintb.Amount != taskData.Amount {
			return true, db.TaskErrorCodeCheckAmountFailed, err
		}

		asset, err := w.stateDB.GetAsset(inscriptionMintb.Ticker)
		if err != nil || asset == nil {
			return true, db.TaskErrorCodeCheckAssetFailed, err
		}
		if !asset.DepositEnabled {
			return true, db.TaskErrorCodeDepositAssetNotEnabled, err
		}
		if taskData.Amount < asset.MinDepositAmount {
			return true, db.TaskErrorCodeDepositAmountTooLow, err
		}

		tokenInfo, err := w.stateDB.GetTokenInfo(inscriptionMintb.Ticker, uint64(taskData.ChainId))
		if err != nil || tokenInfo == nil {
			return true, db.TaskErrorCodeDepositTokenNotSupported, err
		}
		if !tokenInfo.IsActive {
			return true, db.TaskErrorCodeDepositTokenNotActive, err
		}
		return true, db.TaskErrorCodeSuccess, nil
	case *db.WithdrawalTask:
		hashCheckStatus, err := checkTxStatus(taskData.TxHash)
		if err != nil {
			return false, -1, err
		}
		if hashCheckStatus != "success" {
			return true, db.TaskErrorCodeCheckTxFailed, err
		}
		inscriptionBurnb, err := w.stateDB.GetInscriptionBurnb(taskData.TxHash)
		if err != nil || inscriptionBurnb == nil {
			return true, db.TaskErrorCodeCheckInscriptionFailed, err
		}

		if inscriptionBurnb.Amount != taskData.Amount {
			return true, db.TaskErrorCodeCheckAmountFailed, err
		}

		asset, err := w.stateDB.GetAsset(inscriptionBurnb.Ticker)
		if err != nil || asset == nil {
			return true, db.TaskErrorCodeCheckAssetFailed, err
		}
		if !asset.WithdrawalEnabled {
			return true, db.TaskErrorCodeWithdrawalAssetNotEnabled, err
		}
		if taskData.Amount < asset.MinWithdrawAmount {
			return true, db.TaskErrorCodeWithdrawalAmountTooLow, err
		}

		tokenInfo, err := w.stateDB.GetTokenInfo(inscriptionBurnb.Ticker, uint64(taskData.ChainId))
		if err != nil || tokenInfo == nil {
			return true, db.TaskErrorCodeWithdrawalTokenNotSupported, err
		}
		if !tokenInfo.IsActive {
			return true, db.TaskErrorCodeWithdrawalTokenNotActive, err
		}

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

	type TxStatusResponse struct {
		Status string `json:"status"`
	}

	var txStatus TxStatusResponse
	if err := json.Unmarshal(body, &txStatus); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	return txStatus.Status, nil
}
