package tss

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
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

// only used test
func (m *Scheduler) operation(detailTask pool.Task[uint64]) (*contracts.TaskOperation, error) {
	operation := &contracts.TaskOperation{
		TaskId: detailTask.TaskID(),
	}

	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		ec := types.GetCurveTypeByCoinType(coinType)
		userAddress := address.GenerateAddressByPath(m.partyData.GetData(ec).ECPoint(), uint32(coinType), task.Account, task.Index)
		//_ = address.GenerateAddressByPath(m.partyData.GetData(ec).ECPoint(), uint32(coinType), task.Account, task.Index)
		// data := m.voterContract.EncodeRegisterNewAddress(big.NewInt(int64(task.Account)), task.AddressType, big.NewInt(int64(task.Index)), strings.ToLower(userAddress))
		// operation.OptData = data
		// operation.ManagerAddr = common.HexToAddress(config.AppConfig.AccountContract)
		operation.State = db.Completed
		operation.ExtraData = []byte(userAddress) // todo
	case *db.DepositTask:
		needConfirm, checkCode, err := m.checkTask(task)
		if !needConfirm {
			return nil, fmt.Errorf("task %d: hash:%s check failed, %w", task.TaskId, task.TxHash, err)
		}
		if err != nil || checkCode != db.TaskErrorCodeSuccess {
			//taskResult := contracts.TaskPayloadContractWithdrawalResult{
			//	Version:   uint8(db.TaskVersionV1),
			//	Success:   false,
			//	ErrorCode: uint8(checkCode),
			//}
			//taskBytes, err := codec.EncodeTaskResult(db.TaskTypeDeposit, taskResult)
			//if err != nil {
			//	panic(fmt.Errorf("encode result failed for task %d: %w", task.TaskId, err))
			//}

			// data := m.voterContract.EncodeMarkTaskCompleted(new(big.Int).SetUint64(task.TaskId), taskBytes)
			// operation.OptData = data
			// operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
			operation.State = db.Failed
			return operation, err
		}

		//data := m.voterContract.EncodeRecordDeposit(
		//	common.HexToAddress(task.TargetAddress),
		//	big.NewInt(int64(task.Amount)),
		//	task.ChainId.Big(),
		//	common.HexToHash(task.TxHash).Bytes(), // todo
		//	nil,
		//)
		//operation.OptData = data
		//operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = db.Completed
	case *db.WithdrawalTask:
		//data := m.voterContract.EncodeRecordWithdrawal(
		//	common.HexToAddress(task.TargetAddress),
		//	big.NewInt(int64(task.Amount)),
		//	task.ChainId.Big(),
		//	common.HexToHash(task.TxHash).Bytes(), // todo
		//	nil,
		//)
		//operation.OptData = data
		//operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = db.Pending
	default:
		log.Errorf("unhandled default case")
		operation.State = db.Completed
		// operation.OptData = nil // todo
	}

	return operation, nil
}

// only used test
func (m *Scheduler) saveOperations(nonce *big.Int, ops []contracts.TaskOperation, dataHash, hash common.Hash) {
	operations := &Operations{
		Nonce:     nonce,
		Operation: ops,
		Hash:      hash,
		DataHash:  dataHash,
	}
	m.operationsQueue.Add(operations)
	m.currentVoterNonce.Store(nonce.Uint64())
}

// only used test
func (m *Scheduler) joinSignOperationSession(msg SessionMessage[ProposalID, Proposal]) error {
	log.Debugf("joinSignOperationSession: session id: %v, tss nonce(proposalID):%v", msg.SessionID, msg.ProposalID)

	batchData := &types.BatchData{}
	batchData.FromBytes(msg.Data)
	tasks := m.taskQueue.BatchGet(batchData.Ids)
	operations := make([]contracts.TaskOperation, 0, len(tasks))
	for _, item := range tasks {
		op, err := m.operation(item)
		if err != nil {
			return fmt.Errorf("failed to process task: %d, %w", item.TaskID(), err)
		}
		operations = append(operations, *op)
	}

	if len(operations) == 0 {
		return fmt.Errorf("operations queue is empty")
	}

	nonce, dataHash, unSignMsg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(operations)
	if err != nil {
		return fmt.Errorf("batch task generate verify task unsign msg err:%v", err)
	}

	if nonce.Uint64() != msg.SeqId {
		return fmt.Errorf("nonce error: %v", nonce.Uint64())
	}

	if msg.Proposal.Cmp(unSignMsg.Big()) != 0 {
		return fmt.Errorf("proposal error: %v", msg.Proposal.Text(16))
	}

	// only ecdsa batch
	m.NewSignOperationSession(
		msg.SessionID,
		msg.SeqId,
		msg.ProposalID,
		&msg.Proposal,
		msg.Data,
	)
	m.saveOperations(nonce, operations, dataHash, unSignMsg)

	return nil
}

// only used test
func (m *Scheduler) processOperationSignResult(operations *Operations) {
	// 1. save db
	// 2. update status
	if m.IsProposer() {
		log.Info("proposer submit signature")
		w := wallet.NewWallet()
		calldata := m.voterContract.EncodeVerifyAndCall(operations.Operation, operations.Signature)
		log.Infof("calldata: %x, signature: %x,nonce: %v,DataHash: %v, hash: %v", calldata, operations.Signature, operations.Nonce, operations.DataHash, operations.Hash)
		tx, err := w.BuildUnsignTx(
			m.ctx,
			m.LocalSubmitter(),
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
