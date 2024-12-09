package tss

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

type EvmTxContext struct {
	w    *wallet.Wallet
	tx   *ethtypes.Transaction
	task pool.Task[uint64]
}

func (m *Scheduler) processPendingTaskSign(msg *SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) {
	switch taskData := task.(type) {
	case *db.WithdrawalTask:
		switch taskData.Chain {
		case types.CoinTypeBTC:
			switch taskData.AssetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainEthereum:
			switch taskData.AssetType {
			case types.AssetTypeMain:
				w := wallet.NewWallet()
				tx, err := w.BuildUnsignTx(
					m.ctx,
					w.HotAddressOfCoin(types.CoinTypeEVM),
					common.HexToAddress(taskData.ContractAddress),
					big.NewInt(int64(taskData.Amount)),
					nil, nil, &db.EvmWithdraw{
						TaskID: taskData.TaskId,
					}, nil,
				)
				if err != nil {
					log.Errorf("failed to build unsign tx: %v", err)
				} else {
					hash := tx.Hash()
					sessionId := types.ZeroSessionID
					if msg != nil && msg.Proposal.Cmp(hash.Big()) != 0 {
						log.Errorf("the proposal is incorrect")
					} else {
						sessionId = msg.SessionID
					}
					// coinType := types.GetCoinTypeByChain(coinType)
					localData, keyDerivationDelta := m.GenerateDerivationWalletProposal(types.CoinTypeEVM, 0, 0)
					m.NewTxSignSession(
						sessionId,
						taskData.TaskId,
						hash.Big(),
						localData,
						keyDerivationDelta,
					)
					m.txContext.Store(task.TaskID(), &EvmTxContext{
						w:    w,
						tx:   tx,
						task: taskData,
					})
				}
			case types.AssetTypeErc20:

			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainSolana:
			switch taskData.AssetType {
			case types.AssetTypeMain:
			case types.AssetTypeErc20:
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainSui:
			switch taskData.AssetType {
			case types.AssetTypeMain:

			case types.AssetTypeErc20:
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		default:
			panic(fmt.Errorf("unknown Chain type: %v", taskData.Chain))
		}
	default:
		log.Errorf("error pending task id: %v", task.TaskID())
	}
}
