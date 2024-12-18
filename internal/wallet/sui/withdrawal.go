package sui

import (
	"fmt"
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

type TxContext struct {
	signerPubKey []byte
	c            *TxClient
	tx           *UnSignTx
	task         pool.Task[uint64]
}

func (w *WalletClient) processWithdrawTxSign(task *db.WithdrawalTask) {
	var c *TxClient
	if w.IsProd() {
		c = NewClient(w.ctx)
	} else {
		c = NewDevClient()
	}
	hotAddress := address.HotAddressOfSui(w.tss.ECPoint(w.ChainType()))
	log.Infof("hotAddress: %v,targetAddress: %v, amount: %v", hotAddress, task.TargetAddress, task.Amount)
	unSignTx, err := c.BuildPaySuiTx(CoinType(task.ContractAddress, task.Ticker), hotAddress, []Recipient{
		{
			Recipient: task.TargetAddress,
			Amount:    fmt.Sprintf("%d", task.Amount),
		},
	})
	if err != nil {
		log.Errorf("failed to build unsign tx: %v", err)
		return
	}
	proposal := new(big.Int).SetBytes(unSignTx.Blake2bHash())
	w.txContext.Store(task.TaskID(), &TxContext{
		signerPubKey: w.tss.GetPublicKey(hotAddress).SerializeCompressed(),
		c:            c,
		tx:           unSignTx,
		task:         task,
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
			signTx := ctx.tx.SerializedSigWith(res.Signature, ctx.signerPubKey)
			log.Debugf("SuiTxContext: signTx: %v, signature: %x", signTx, res.Signature)
			digest, err := ctx.c.SendTx((*SignedTx)(signTx))
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
