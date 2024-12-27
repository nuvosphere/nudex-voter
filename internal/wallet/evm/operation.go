package evm

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Operations struct {
	Nonce     *big.Int
	Operation []contracts.TaskOperation
	Hash      common.Hash
	DataHash  common.Hash
	Signature []byte
}

func (o *Operations) TaskID() uint64 {
	return o.Nonce.Uint64()
}

func (o *Operations) Type() int {
	return db.TaskTypeOperations
}

func (w *WalletClient) Operation(detailTask pool.Task[uint64]) *contracts.TaskOperation {
	operation := &contracts.TaskOperation{
		TaskId: detailTask.TaskID(),
	}

	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		userAddress := w.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
		operation.State = uint8(task.Task.State)
		operation.ExtraData = []byte(userAddress) // todo
	case *db.DepositTask:
		//data := w.VoterContract().EncodeRecordDeposit(
		//	common.HexToAddress(task.TargetAddress),
		//	big.NewInt(int64(task.Amount)),
		//	task.ChainId.Big(),
		//	common.HexToHash(task.TxHash).Bytes(), // todo
		//	nil,
		//)
		//operation.OptData = data
		//operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = uint8(task.Task.State)
	case *db.WithdrawalTask:
		//data := w.VoterContract().EncodeRecordWithdrawal(
		//	common.HexToAddress(task.TargetAddress),
		//	big.NewInt(int64(task.Amount)),
		//	task.ChainId.Big(),
		//	common.HexToHash(task.TxHash).Bytes(), // todo
		//	nil,
		//)
		//operation.OptData = data
		//operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = uint8(task.Task.State)
	default:
		log.Errorf("unhandled default case")
		operation.State = db.Completed
		// operation.OptData = nil // todo
	}

	return operation
}

func (w *WalletClient) loopProcessOperation() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("approve proposal done")

			case <-ticker.C:
				w.processOperation()

			case <-w.notify:
				w.processOperation()
			}
		}
	}()
}

const TopN = 20

func (w *WalletClient) processOperation() {
	log.Info("batch proposal")
	tasks := w.submitTaskQueue.GetTopN(TopN)
	operations := lo.Map(tasks, func(item pool.Task[uint64], index int) contracts.TaskOperation { return *w.Operation(item) })
	if len(operations) == 0 {
		log.Warnf("operationsQueue is empty")
		return
	}
	nonce, dataHash, msg, err := w.VoterContract().GenerateVerifyTaskUnSignMsg(operations)
	if err != nil {
		log.Errorf("batch task generate verify task unsign msg err:%v", err)
		return
	}
	log.Infof("nonce: %v, dataHash: %v, msg: %v", nonce, dataHash, msg)

	data := lo.Map(tasks, func(item pool.Task[uint64], index int) uint64 { return item.TaskID() })
	batchData := types.BatchData{Ids: data}

	signReq := &suite.SignReq{
		SeqId:      nonce.Uint64(),
		Type:       types.SignOperationSessionType,
		ChainType:  w.ChainType(),
		Signer:     w.tss.TssSigner().String(),
		DataDigest: msg.String(),
		SignData:   msg.Bytes(),
		ExtraData:  batchData.Bytes(),
	}

	err = w.tss.Sign(signReq)
	if err != nil {
		log.Errorf("batch task signTx err:%v", err)
		return
	}
	w.saveOperations(nonce, operations, dataHash, msg)
}

func (w *WalletClient) saveOperations(nonce *big.Int, ops []contracts.TaskOperation, dataHash, hash common.Hash) {
	operations := &Operations{
		Nonce:     nonce,
		Operation: ops,
		Hash:      hash,
		DataHash:  dataHash,
	}
	w.operationsQueue.Add(operations)
	w.currentVoterNonce.Store(nonce.Uint64())
}

func (w *WalletClient) processOperationSignResult(operations *Operations) {
	// 1. save db
	// 2. update status
	if w.tss.IsProposer() {
		log.Info("proposer submit signature")
		calldata := w.VoterContract().EncodeVerifyAndCall(operations.Operation, operations.Signature)
		log.Infof("calldata: %x, signature: %x,nonce: %v,DataHash: %v, hash: %v", calldata, operations.Signature, operations.Nonce, operations.DataHash, operations.Hash)
		tx, err := w.BuildUnSignTx(
			w.tss.LocalSubmitter(),
			common.HexToAddress(config.AppConfig.VotingContract),
			big.NewInt(0),
			calldata,
			operations.Type(),
			operations.TaskID(),
		)
		if err != nil {
			log.Errorf("failed to build unsigned transaction: %v", err)
			return
		}

		ctx := w.NewTxContext(tx)
		w.pendingTx.Store(ctx.TxHash(), ctx)
		defer w.pendingTx.Delete(ctx.TxHash())
		err = w.signTx(ctx)
		if err != nil {
			log.Errorf("failed to signTx transaction: %v", err)
			return
		}

		err = w.SendSingedTx(ctx)
		if err != nil {
			log.Errorf("failed to send transaction: %v", err)
			return
		}
		// updated status to pending
		receipt, err := w.WaitTxSuccess(ctx.TxHash())
		if err != nil {
			log.Errorf("failed to wait transaction success: %v", err)
			return
		}

		if receipt.Status == 0 {
			// updated status to fail
			log.Errorf("failed to submit transaction for taskId: %d,txHash: %s", operations.TaskID(), ctx.TxHash().String())
		} else {
			// updated status to completed
			log.Infof("successfully submitted transaction for taskId: %d, txHash: %s", operations.TaskID(), ctx.TxHash().String())
		}
	}
}

func (w *WalletClient) receiveSubmitTaskLoop() {
	taskEvent := w.event.Subscribe(eventbus.EventSubmitTask{})

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				log.Info("evm wallet receive task event done")

			case detailTask := <-taskEvent:
				val, ok := detailTask.(db.DetailTask)
				if ok {
					w.submitTaskQueue.Add(val)
					if w.submitTaskQueue.Len() >= TopN {
						w.notify <- struct{}{}
					}
				}
			}
		}
	}()
}
