package tss

import (
	"encoding/json"
	"fmt"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func (m *Scheduler) processOperationSignResult(operations *Operations) {
	// 1. save db
	// 2. update status
	if m.IsProposer() {
		log.Info("proposer submit signature")
		w := wallet.NewWallet()
		calldata := m.voterContract.EncodeVerifyAndCall(operations.Operation, operations.Signature)
		log.Infof("calldata: %x, signature: %x,nonce: %v,DataHash: %v, hash: %v", calldata, operations.Signature, operations.Nonce, operations.DataHash, operations.Hash)
		data, err := json.Marshal(operations)
		utils.Assert(err)
		tx, err := w.BuildUnsignTx(
			m.ctx,
			m.LocalSubmitter(),
			common.HexToAddress(config.AppConfig.VotingContract),
			big.NewInt(0),
			calldata,
			&db.Operations{
				Nonce: decimal.NewFromBigInt(operations.Nonce, 0),
				Data:  string(data),
			}, nil, nil,
		)
		if err != nil {
			log.Errorf("failed to build unsigned transaction: %v", err)
			return
		}

		chainId, err := w.ChainID(m.ctx)
		if err != nil {
			log.Errorf("failed to ChainID: %v", err)
			return
		}
		signedTx, err := ethtypes.SignTx(tx, ethtypes.LatestSignerForChainID(chainId), config.L2PrivateKey)
		if err != nil {
			log.Errorf("failed to sign transaction: %v", err)
			return
		}

		err = w.SendSingedTx(m.ctx, signedTx)
		if err != nil {
			log.Errorf("failed to send transaction: %v", err)
			return
		}
		// updated status to pending
		receipt, err := w.WaitTxSuccess(m.ctx, signedTx.Hash())
		if err != nil {
			log.Errorf("failed to wait transaction success: %v", err)
			return
		}

		if receipt.Status == 0 {
			// updated status to fail
			log.Errorf("failed to submit transaction for taskId: %d,txHash: %s", operations.TaskID(), signedTx.Hash().String())
		} else {
			// updated status to completed
			log.Infof("successfully submitted transaction for taskId: %d, txHash: %s", operations.TaskID(), signedTx.Hash().String())
		}
	}
}

type EvmTxContext struct {
	w    *wallet.Wallet
	tx   *ethtypes.Transaction
	task pool.Task[uint64]
}

func (m *Scheduler) processTxSign(msg *SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) {
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
			w := wallet.NewWallet()
			from := w.HotAddressOfCoin(types.CoinTypeEVM)
			to := common.HexToAddress(taskData.TargetAddress)
			var tx *ethtypes.Transaction
			var err error
			withdraw := &db.EvmWithdraw{
				TaskID: taskData.TaskId,
			}
			switch taskData.AssetType {
			case types.AssetTypeMain:
				tx, err = w.BuildUnsignTx(
					m.ctx,
					from,
					to,
					big.NewInt(int64(taskData.Amount)), nil, nil, withdraw, nil,
				)
			case types.AssetTypeErc20:
				tx, err = w.BuildUnsignTx(
					m.ctx,
					from,
					common.HexToAddress(taskData.ContractAddress),
					nil,
					contracts.EncodeTransferOfERC20(from, to, big.NewInt(int64(taskData.Amount))), nil, withdraw, nil,
				)
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
				return
			}
			if err != nil {
				log.Errorf("failed to build unsign tx: %v", err)
				return
			}
			hash := tx.Hash()
			sessionId := types.ZeroSessionID
			if msg != nil {
				if msg.Proposal.Cmp(hash.Big()) != 0 {
					log.Errorf("the proposal is incorrect")
					return
				}
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

func (m *Scheduler) processTxSignResult(taskID uint64, data *tsscommon.SignatureData) {
	txCtx, ok := m.txContext.Load(taskID)
	if ok {
		switch ctx := txCtx.(type) {
		case *EvmTxContext:
			hash := ctx.tx.Hash()
			err := ctx.w.SendTransactionWithSignature(m.ctx, ctx.tx, secp256k1Signature(data))
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}
			// updated status to pending
			receipt, err := ctx.w.WaitTxSuccess(m.ctx, hash)
			if err != nil {
				log.Errorf("failed to wait transaction success: %v", err)
				return
			}
			if receipt.Status == 0 {
				// updated status to fail
				log.Errorf("failed to submit transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), hash)
			} else {
				// updated status to completed
				log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), hash)
			}
		}
	}
}
