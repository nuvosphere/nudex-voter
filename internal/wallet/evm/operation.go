package evm

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type Operations struct {
	Nonce     *big.Int
	Operation []contracts.Operation
	Hash      common.Hash
	DataHash  common.Hash
	Signature []byte
}

func (o *Operations) TaskID() uint64 {
	return o.Nonce.Uint64()
}

func (o *Operations) Type() int {
	return db.TypeOperations
}

func (m *WalletClient) Operation(detailTask pool.Task[uint64]) *contracts.Operation {
	operation := &contracts.Operation{
		TaskId: detailTask.TaskID(),
	}

	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		userAddress := m.tss.GetUserAddress(uint32(coinType), task.Account, task.Index)
		data := m.voterContract.EncodeRegisterNewAddress(big.NewInt(int64(task.Account)), task.Chain, big.NewInt(int64(task.Index)), strings.ToLower(userAddress))
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.AccountContract)
		operation.State = db.Completed
	case *db.DepositTask:
		data := m.voterContract.EncodeRecordDeposit(
			common.HexToAddress(task.TargetAddress),
			big.NewInt(int64(task.Amount)),
			uint64(task.ChainId),
			common.HexToHash(task.TxHash).Bytes(), // todo
			nil,
		)
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = db.Completed
	case *db.WithdrawalTask:
		data := m.voterContract.EncodeRecordWithdrawal(
			common.HexToAddress(task.TargetAddress),
			big.NewInt(int64(task.Amount)),
			uint64(task.ChainId),
			common.HexToHash(task.TxHash).Bytes(), // todo
			nil,
		)
		operation.OptData = data
		operation.ManagerAddr = common.HexToAddress(config.AppConfig.DepositContract)
		operation.State = db.Pending
	default:
		log.Errorf("unhandled default case")
		operation.State = db.Completed
		operation.OptData = nil // todo
	}

	return operation
}

func (m *WalletClient) loopApproveProposal() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		select {
		case <-m.ctx.Done():
			log.Info("approve proposal done")

		case <-ticker.C:
			m.BatchTask()

		case <-m.notify:
			m.BatchTask()
		}
	}()
}

const TopN = 20

func (m *WalletClient) BatchTask() {
	log.Info("batch proposal")
	tasks := m.taskQueue.GetTopN(TopN)
	operations := lo.Map(tasks, func(item pool.Task[uint64], index int) contracts.Operation { return *m.Operation(item) })
	if len(operations) == 0 {
		log.Warnf("operationsQueue is empty")
		return
	}
	nonce, dataHash, msg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(operations)
	if err != nil {
		log.Errorf("batch task generate verify task unsign msg err:%v", err)
		return
	}
	log.Infof("nonce: %v, dataHash: %v, msg: %v", nonce, dataHash, msg)

	data := lo.Map(tasks, func(item pool.Task[uint64], index int) uint64 { return item.TaskID() })
	// batchData := types.BatchData{Ids: data}
	_ = types.BatchData{Ids: data}

	// only ecdsa batch
	//m.NewMasterSignBatchSession(
	//	ZeroSessionID,
	//	nonce.Uint64(), // ProposalID
	//	msg.Big(),
	//	batchData.Bytes(),
	//)
	//m.saveOperations(nonce, operations, dataHash, msg)
}

func (m *WalletClient) saveOperations(nonce *big.Int, ops []contracts.Operation, dataHash, hash common.Hash) {
	operations := &Operations{
		Nonce:     nonce,
		Operation: ops,
		Hash:      hash,
		DataHash:  dataHash,
	}
	m.operationsQueue.Add(operations)
	m.currentVoterNonce.Store(nonce.Uint64())
}
