package tss

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	soltypes "github.com/blocto/solana-go-sdk/types"
	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/crypto/ckd"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaSigning "github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	eddsaSigning "github.com/bnb-chain/tss-lib/v2/eddsa/signing"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/nuvosphere/nudex-voter/internal/wallet/btc"
	"github.com/nuvosphere/nudex-voter/internal/wallet/solana"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func (m *Scheduler) GenerateDerivationWalletProposal(coinType, account uint32, index uint8) (LocalPartySaveData, *big.Int) {
	// coinType := types.GetCoinTypeByChain(coinType)
	path := wallet.Bip44DerivationPath(coinType, account, index)
	param, err := path.ToParams()
	utils.Assert(err)
	ec := types.GetCurveTypeByCoinType(int(coinType))
	localPartySaveData := m.partyData.GetData(ec)
	l := *localPartySaveData

	chainCode := big.NewInt(int64(coinType)).Bytes() // todo
	keyDerivationDelta, extendedChildPk, err := ckd.DerivingPubkeyFromPath(l.ECPoint(), chainCode, param.Indexes(), ec.EC())
	utils.Assert(err)
	switch ec {
	case crypto.ECDSA:
		data := []ecdsaKeygen.LocalPartySaveData{*l.ECDSAData()}
		err = ecdsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			data,
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)
		l.SetData(&data[0])
		return l, keyDerivationDelta
	case crypto.EDDSA:
		data := []eddsaKeygen.LocalPartySaveData{*l.EDDSAData()}
		err = eddsaSigning.UpdatePublicKeyAndAdjustBigXj(
			keyDerivationDelta,
			data,
			extendedChildPk.PublicKey,
			ec.EC(),
		)
		utils.Assert(err)
		l.SetData(&data[0])
		return l, keyDerivationDelta

	default:
		panic(fmt.Errorf("unknown EC type: %v", ec))
	}
}

func (m *Scheduler) loopSigInToOut() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				log.Info("tss signature read result loop stopped")
			case result := <-m.sigInToOut:
				log.Infof("finish consensus success, sessionID:%s", result.SessionID)
				info := fmt.Sprintf("tss signature sessionID=%v, groupID=%v, ProposalID=%v", result.SessionID, result.GroupID, result.ProposalID)

				if result.Err != nil {
					log.Errorf("result error:%v, error: %v", result.Err, info)
				} else {
					switch result.Type {
					case SignBatchTaskSessionType:
						ops := m.operationsQueue.Get(result.ProposalID).(*Operations)
						ops.Signature = secp256k1Signature(result.Data)
						log.Infof("result.Data.Signature: len: %d, result.Data.Signature: %x", len(result.Data.Signature), result.Data.Signature)
						log.Infof("result.Data.SignatureRecovery: len: %d, result.Data.SignatureRecovery: %x", len(result.Data.SignatureRecovery), result.Data.SignatureRecovery)
						log.Infof("ops.Signature: len: %d, ops.Signature: %x, Hash: %v,dataHash: %v", len(ops.Signature), ops.Signature, ops.Hash, ops.DataHash)
						m.processOperationSignResult(ops)
						lo.ForEach(ops.Operation, func(item contracts.Operation, _ int) { m.AddDiscussedTask(item.TaskId) })
						m.operationsQueue.RemoveTopN(ops.TaskID() - 1)

					case TxSignatureSessionType:
						m.processTxSignResult(result.ProposalID, result.Data)
					default:
						log.Infof("tss signature result: %v", result)
					}
				}
			}
		}
	}()
}

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

func (m *Scheduler) processTxSign(msg *SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) {
	log.Debugf("processTxSign taskId: %v", task.TaskID())
	switch taskData := task.(type) {
	case *db.WithdrawalTask:
		log.Debugf("processTxSign task: %v", taskData)
		switch taskData.Chain {
		case types.ChainBitcoin: // todo
			switch taskData.AssetType {
			case types.AssetTypeMain:
				// coinType := types.GetCoinTypeByChain(taskData.Chain)
				localData, keyDerivationDelta := m.GenerateDerivationWalletProposal(types.CoinTypeBTC, 0, 0)
				c := btc.NewTxClient(m.ctx, time.Second*60, &chaincfg.MainNetParams, localData.PublicKey())
				sigCtx := &SignContext{
					chainType:          taskData.Chain,
					localData:          localData,
					keyDerivationDelta: keyDerivationDelta,
				}
				from := localData.Address(taskData.Chain)
				m.sigContext.Store(from, sigCtx)

				sessionId := ZeroSessionID
				var proposal *big.Int
				go func() {
					err := c.BuildTx(taskData.TargetAddress, int64(taskData.Amount))
					if err != nil {
						log.Errorf("failed to send transaction: %v", err)
						return
					}
				}()
				hash := c.NextSignTask()
				if hash == nil {
					log.Errorf("failed to next sign ")
					return
				}
				proposal.SetBytes(hash)

				if msg != nil {
					sessionId = msg.SessionID
					if proposal.Cmp(&msg.Proposal) != 0 {
						log.Errorf("the proposal is incorrect of btc")
					}
				}

				m.NewTxSignSession(
					sessionId,
					taskData.TaskId,
					proposal,
					localData,
					keyDerivationDelta,
					from,
				)

			case types.AssetTypeErc20:
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainEthereum:
			w := wallet.NewWallet()
			hotAddress := w.HotAddressOfCoin(types.CoinTypeEVM)
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
					hotAddress,
					to,
					big.NewInt(int64(taskData.Amount)), nil, nil, withdraw, nil,
				)
			case types.AssetTypeErc20:
				tx, err = w.BuildUnsignTx(
					m.ctx,
					hotAddress,
					common.HexToAddress(taskData.ContractAddress),
					nil,
					contracts.EncodeTransferOfERC20(hotAddress, to, big.NewInt(int64(taskData.Amount))), nil, withdraw, nil,
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
			sessionId := ZeroSessionID
			if msg != nil {
				if msg.Proposal.Cmp(hash.Big()) != 0 {
					log.Errorf("the proposal is incorrect of evm")
					return
				}
				sessionId = msg.SessionID
			} else {
				m.txContext.Store(task.TaskID(), &EvmTxContext{
					w:    w,
					tx:   tx,
					task: taskData,
				})
			}
			// coinType := types.GetCoinTypeByChain(coinType)
			localData, keyDerivationDelta := m.GenerateDerivationWalletProposal(types.CoinTypeEVM, 0, 0)
			m.NewTxSignSession(
				sessionId,
				taskData.TaskId,
				hash.Big(),
				localData,
				keyDerivationDelta,
				hotAddress.String(),
			)

		case types.ChainSolana:
			var c *solana.SolClient
			if m.isProd {
				c = solana.NewSolClient()
			} else {
				c = solana.NewDevSolClient()
			}
			hotAddress := wallet.HotAddressOfSolana(m.partyData.GetData(crypto.EDDSA).ECPoint())
			log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, taskData.TargetAddress, taskData.Amount)
			var (
				tx  *solana.UnSignTx
				err error
			)
			switch taskData.AssetType {
			case types.AssetTypeMain:
				tx, err = c.BuildSolTransferWithAddress(hotAddress, taskData.TargetAddress, taskData.Amount)
			case types.AssetTypeErc20:
				// todo
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
				return
			}
			if err != nil {
				log.Errorf("failed to build unsign tx: %v", err)
				return
			}
			raw, err := tx.RawData()
			if err != nil {
				log.Errorf("failed to build unsign tx: %v", err)
				return
			}
			log.Infof("raw: %x", raw)
			proposal := new(big.Int).SetBytes(raw)
			sessionId := ZeroSessionID
			if msg != nil {
				log.Infof("msg.proposal: %v, proposal: %v", msg.Proposal.String(), proposal.String())
				// todo
				//if msg.proposal.Cmp(proposal) != 0 {
				//	log.Errorf("the proposal is incorrect of solana")
				//	return
				//}
				proposal = &msg.Proposal // todo
				sessionId = msg.SessionID
			} else {
				m.txContext.Store(task.TaskID(), &SolTxContext{
					c:    c,
					tx:   tx,
					task: taskData,
				})
			}
			// coinType := types.GetCoinTypeByChain(coinType)
			localData, keyDerivationDelta := m.GenerateDerivationWalletProposal(types.CoinTypeSOL, 0, 0)
			m.NewTxSignSession(
				sessionId,
				taskData.TaskId,
				proposal,
				localData,
				keyDerivationDelta,
				hotAddress,
			)

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
	defer m.txContext.Delete(taskID)
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
			receipt, err := ctx.w.WaitTxSuccess(m.ctx, hash) // todo track error: re-sign、request
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

		case *SolTxContext:
			unSignRawData, _ := ctx.tx.RawData()
			log.Debugf("SolTxContext: unSignRawData: %x, signature: %x", unSignRawData, data.Signature)
			sig, err := ctx.c.SyncSendTransaction(m.ctx, (*soltypes.Transaction)(ctx.tx.BuildSolTransaction(data.Signature)))
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}
			log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), sig)
		}
	}
}
