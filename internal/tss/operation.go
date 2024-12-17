package tss

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
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

func (m *Scheduler) Operation(detailTask pool.Task[uint64]) *contracts.Operation {
	operation := &contracts.Operation{
		TaskId: detailTask.TaskID(),
	}

	switch task := detailTask.(type) {
	case *db.CreateWalletTask:
		coinType := types.GetCoinTypeByChain(task.Chain)
		ec := types.GetCurveTypeByCoinType(coinType)
		userAddress := address.GenerateAddressByPath(m.partyData.GetData(ec).ECPoint(), uint32(coinType), task.Account, task.Index)
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

func (m *Scheduler) saveOperations(nonce *big.Int, ops []contracts.Operation, dataHash, hash common.Hash) {
	operations := &Operations{
		Nonce:     nonce,
		Operation: ops,
		Hash:      hash,
		DataHash:  dataHash,
	}
	m.operationsQueue.Add(operations)
	m.currentVoterNonce.Store(nonce.Uint64())
}

func (m *Scheduler) JoinSignOperationSession(msg SessionMessage[ProposalID, Proposal]) error {
	log.Debugf("JoinSignBatchTaskSession: session id: %v, tss nonce(proposalID):%v", msg.SessionID, msg.ProposalID)

	batchData := &types.BatchData{}
	batchData.FromBytes(msg.Data)
	tasks := m.taskQueue.BatchGet(batchData.Ids)
	operations := lo.Map(tasks, func(item pool.Task[uint64], index int) contracts.Operation { return *m.Operation(item) })

	nonce, dataHash, unSignMsg, err := m.voterContract.GenerateVerifyTaskUnSignMsg(operations)
	if err != nil {
		return fmt.Errorf("batch task generate verify task unsign msg err:%v", err)
	}

	if nonce.Uint64() != msg.ProposalID {
		return fmt.Errorf("nonce error: %v", nonce.Uint64())
	}

	if msg.Proposal.Cmp(unSignMsg.Big()) != 0 {
		return fmt.Errorf("proposal error: %v", msg.Proposal.Text(16))
	}

	// only ecdsa batch
	m.NewSignOperationSession(
		msg.SessionID,
		msg.ProposalID,
		&msg.Proposal,
		msg.Data,
	)
	m.saveOperations(nonce, operations, dataHash, unSignMsg)

	return nil
}
