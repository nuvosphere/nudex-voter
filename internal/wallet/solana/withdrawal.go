package solana

import (
	"math/big"

	soltypes "github.com/blocto/solana-go-sdk/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

type TxContext struct {
	c    *SolClient
	tx   *UnSignTx
	task pool.Task[uint64]
}

func (w *WalletClient) processWithdrawTxSign(task *db.WithdrawalTask) {
	var c *SolClient
	if w.IsProd() {
		c = NewSolClient()
	} else {
		c = NewDevSolClient()
	}
	hotAddress := address.HotAddressOfSui(w.tss.ECPoint(w.ChainType()))
	log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.TargetAddress, task.Amount)
	var (
		tx  *UnSignTx
		err error
	)
	switch task.AssetType {
	case types.AssetTypeMain:
		tx, err = c.BuildSolTransferWithAddress(hotAddress, task.TargetAddress, task.Amount)
	case types.AssetTypeErc20:
		// todo
	default:
		log.Errorf("unknown asset type: %v", task.AssetType)
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
	w.txContext.Store(task.TaskID(), &TxContext{
		c:    c,
		tx:   tx,
		task: task,
	})
	req := &suite.SignReq{
		SeqId:      task.TaskID(),
		Type:       types.SignTxSessionType,
		ChainType:  w.ChainType(),
		Signer:     hotAddress,
		DataDigest: proposal.String(),
		SignData:   proposal.Bytes(),
		ExtraData:  nil,
	}

	err = w.tss.Sign(req)
	if err != nil {
		log.Errorf("failed to sign tx: %v", err)
	}
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
	taskID := res.SeqId
	txCtx, ok := w.txContext.Load(taskID)
	defer w.txContext.Delete(taskID)
	if ok {
		switch ctx := txCtx.(type) {
		case *TxContext:
			unSignRawData, _ := ctx.tx.RawData()
			log.Debugf("SolTxContext: unSignRawData: %x, signature: %x", unSignRawData, res.Signature)
			sig, err := ctx.c.SyncSendTransaction(w.ctx, (*soltypes.Transaction)(ctx.tx.BuildSolTransaction(res.Signature)))
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}
			log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.task.TaskID(), sig)
		}
	}
}
