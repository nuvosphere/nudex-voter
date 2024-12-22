package tss

import (
	"fmt"
	"math/big"
	"time"

	soltypes "github.com/blocto/solana-go-sdk/types"
	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/nuvosphere/nudex-voter/internal/wallet/btc"
	"github.com/nuvosphere/nudex-voter/internal/wallet/solana"
	"github.com/nuvosphere/nudex-voter/internal/wallet/sui"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

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
					case types.SignTestOperationSessionType:
						ops := m.operationsQueue.Get(result.SeqId).(*Operations)
						ops.Signature = secp256k1Signature(result.Data)
						log.Infof("result.Data.Signature: len: %d, result.Data.Signature: %x", len(result.Data.Signature), result.Data.Signature)
						log.Infof("result.Data.SignatureRecovery: len: %d, result.Data.SignatureRecovery: %x", len(result.Data.SignatureRecovery), result.Data.SignatureRecovery)
						log.Infof("ops.Signature: len: %d, ops.Signature: %x, Hash: %v,dataHash: %v", len(ops.Signature), ops.Signature, ops.Hash, ops.DataHash)
						m.processOperationSignResult(ops)
						lo.ForEach(ops.Operation, func(item contracts.Operation, _ int) { m.AddDiscussedTask(item.TaskId) })
						m.operationsQueue.RemoveTopN(ops.TaskID() - 1)

					case types.SignTestTxSessionType:
						m.processTxSignResult(result.SeqId, result.Data)
					case types.SignTxSessionType, types.SignOperationSessionType:
						signature := result.Data.Signature
						if result.ChainType == types.ChainEthereum {
							signature = secp256k1Signature(result.Data)
						}
						m.PostClient(result.ChainType, result.SeqId, result.ProposalID, signature)
					default:
						log.Infof("tss signature result: %v", result)
					}
				}
			}
		}
	}()
}

func (m *Scheduler) PostClient(chainType uint8, SeqId uint64, signDigest string, signature []byte) {
	defer m.crw.RUnlock()
	m.crw.RLock()
	c, ok := m.tssClients[chainType]
	if ok {
		c.ReceiveSignature(&suite.SignRes{
			SeqId:      SeqId,
			DataDigest: signDigest,
			Signature:  signature,
		})
	}
}

func (m *Scheduler) processTxSign(msg *SessionMessage[ProposalID, Proposal]) error {
	defer m.crw.RUnlock()
	m.crw.RLock()
	c, ok := m.tssClients[msg.ChainType]
	if ok {
		return c.Verify(msg.SeqId, msg.ProposalID, msg.Data)
	}
	return nil
}

// only used test
func (m *Scheduler) processTxSignForTest(msg *SessionMessage[ProposalID, Proposal], task pool.Task[uint64]) {
	log.Debugf("processTxSign taskId: %v", task.TaskID())
	sessionId := ZeroSessionID
	seqId := task.TaskID()
	var signer *SignerContext
	proposalID := ""
	var proposal *Proposal

	switch taskData := task.(type) {
	case *db.WithdrawalTask:
		log.Debugf("processTxSign task: %v", taskData)
		switch taskData.Chain {
		case types.ChainBitcoin: // todo
			switch taskData.AssetType {
			case types.AssetTypeMain:
				// coinType := types.GetCoinTypeByChain(taskData.AddressType)
				localData, _ := m.GenerateDerivationWalletProposal(types.CoinTypeBTC, 0, 0)
				c := btc.NewTxClient(m.ctx, time.Second*60, &chaincfg.MainNetParams, localData.PublicKey())
				//sigCtx := &SignerContext{
				//	chainType:          taskData.AddressType,
				//	localData:          localData,
				//	keyDerivationDelta: keyDerivationDelta,
				//}
				hotAddress := localData.Address(taskData.Chain)
				// m.sigContext.Store(hotAddress, sigCtx)
				signer = m.GetSigner(hotAddress)
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

			case types.AssetTypeErc20:
			default:
				log.Errorf("unknown asset type: %v", taskData.AssetType)
			}
		case types.ChainEthereum:
			w := wallet.NewWallet()
			hotAddress := w.HotAddressOfCoin(types.CoinTypeEVM)
			signer = m.GetSigner(hotAddress.String())
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

		case types.ChainSolana:
			var c *solana.SolClient
			if m.isProd {
				c = solana.NewSolClient()
			} else {
				c = solana.NewDevSolClient()
			}
			hotAddress := address.HotAddressOfSolana(m.partyData.GetData(crypto.EDDSA).ECPoint())
			log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, taskData.TargetAddress, taskData.Amount)
			signer = m.GetSigner(hotAddress)
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
			proposal = new(big.Int).SetBytes(raw)
			if msg != nil {
				log.Infof("msg.proposal: %v, proposal: %v", msg.Proposal.String(), proposal.String())
				// todo
				//if msg.proposal.Cmp(proposal) != 0 {
				//	log.Errorf("the proposal is incorrect of solana")
				//	return
				//}
				proposalID = msg.ProposalID
				proposal = &msg.Proposal // todo
				sessionId = msg.SessionID
			} else {
				m.txContext.Store(task.TaskID(), &SolTxContext{
					c:    c,
					tx:   tx,
					task: taskData,
				})
			}

		case types.ChainSui:
			var c *sui.TxClient
			if m.isProd {
				c = sui.NewClient(m.ctx)
			} else {
				c = sui.NewDevClient()
			}
			hotAddress := address.HotAddressOfSui(m.partyData.GetData(crypto.EDDSA).ECPoint())
			log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, taskData.TargetAddress, taskData.Amount)
			signer = m.GetSigner(hotAddress)
			unSignTx, err := c.BuildPaySuiTx(sui.CoinType(taskData.ContractAddress, taskData.Ticker.String()), hotAddress, []sui.Recipient{
				{
					Recipient: taskData.TargetAddress,
					Amount:    fmt.Sprintf("%d", taskData.Amount),
				},
			})
			if err != nil {
				log.Errorf("failed to build unsign tx: %v", err)
				return
			}
			localData, _ := m.GenerateDerivationWalletProposal(types.CoinTypeSUI, 0, 0)
			proposal = new(big.Int).SetBytes(unSignTx.Blake2bHash())
			if msg != nil {
				log.Infof("msg.proposal: %v, proposal: %v", msg.Proposal.String(), proposal.String())
				//if msg.Proposal.Cmp(proposal) != 0 {//todo
				//	log.Errorf("the proposal is incorrect of sui")
				//	return
				//}
				proposalID = msg.ProposalID
				proposal = &msg.Proposal
				sessionId = msg.SessionID
			} else {
				m.txContext.Store(task.TaskID(), &SuiTxContext{
					signerPubKey: localData.PublicKey().SerializeCompressed(),
					c:            c,
					tx:           unSignTx,
					task:         taskData,
				})
			}
		default:
			panic(fmt.Errorf("unknown AddressType type: %v", taskData.Chain))
		}
	default:
		log.Errorf("error pending task id: %v", task.TaskID())
		return
	}
	if proposalID == "" {
		proposalID = proposal.String()
	}
	m.NewTxSignSession(
		sessionId,
		seqId,
		proposalID,
		proposal,
		signer,
	)
}

// only used test
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

		case *SuiTxContext:
			signTx := ctx.tx.SerializedSigWith(data.Signature, ctx.signerPubKey)
			log.Debugf("SuiTxContext: signTx: %v, signature: %x", signTx, data.Signature)
			digest, err := ctx.c.SendTx((*sui.SignedTx)(signTx))
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}

			err = ctx.c.WaitSuccess(digest)
			if err != nil {
				log.Errorf("check transaction success err: %v", err)
				return
			}

			log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), digest)
		}
	}
}

func (m *Scheduler) Sign(req *suite.SignReq) error {
	signerCtx := m.GetSigner(req.Signer)
	if signerCtx == nil {
		return fmt.Errorf("signer not found in context")
	}
	if m.isCanProposal() {
		m.NewSignSessionWitKey(
			ZeroSessionID,
			req.SeqId,
			req.Type,
			new(big.Int).SetBytes(req.SignData),
			req.DataDigest,
			req.ExtraData,
			signerCtx,
		)
	}
	return nil
}
