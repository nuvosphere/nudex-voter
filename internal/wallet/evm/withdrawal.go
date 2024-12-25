package evm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) processWithdrawTxSign(task *db.WithdrawalTask) {
	log.Debugf("processTxSign taskId: %v", task.TaskID())

	hotAddress := common.HexToAddress(address.HotAddressOfEth(w.tss.ECPoint(w.ChainType())))
	to := common.HexToAddress(task.TargetAddress)
	var tx *db.EvmTransaction
	var err error
	withdraw := &db.EvmWithdraw{
		TaskID: task.TaskId,
	}
	switch task.AssetType {
	case types.AssetTypeMain:
		tx, err = w.BuildUnsignTx(
			hotAddress,
			to,
			big.NewInt(int64(task.Amount)), nil, nil, withdraw, nil,
		)
	case types.AssetTypeErc20:
		tx, err = w.BuildUnsignTx(

			hotAddress,
			common.HexToAddress(task.ContractAddress),
			nil,
			contracts.EncodeTransferOfERC20(hotAddress, to, big.NewInt(int64(task.Amount))), nil, withdraw, nil,
		)
	default:
		log.Errorf("unknown asset type: %v", task.AssetType)
		return
	}
	if err != nil {
		log.Errorf("failed to build unsign tx: %v", err)
		return
	}
	hash := tx.TxHash

	req := &suite.SignReq{
		SeqId:      task.TaskID(),
		Type:       types.SignTxSessionType,
		ChainType:  w.ChainType(),
		Signer:     hotAddress.String(),
		DataDigest: hash.String(),
		SignData:   hash.Bytes(),
		ExtraData:  nil,
	}
	w.pendingTx.Store(hash, &TxContext{dbTX: tx})

	err = w.tss.Sign(req)
	if err != nil {
		log.Errorf("failed to sign tx: %v", err)
	}
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
	txCtx, ok := w.pendingTx.Load(res.DataDigest)
	defer w.pendingTx.Delete(res.DataDigest)
	if ok {
		switch ctx := txCtx.(type) {
		case *TxContext:
			ctx.sig = res.Signature
			err := w.SendSingedTx(ctx)
			if err != nil {
				log.Errorf("send transaction err: %v", err)
				return
			}
			hash := ctx.TxHash()
			// updated status to pending
			receipt, err := w.WaitTxSuccess(hash)
			if err != nil {
				log.Errorf("failed to wait transaction success: %v", err)
				return
			}
			if receipt.Status == 0 {
				// updated status to fail
				log.Errorf("failed to submit transaction for taskId: %d,txHash: %v", res.SeqId, hash)
			} else {
				// updated status to completed
				log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", res.SeqId, hash)
			}
		}
	}
}
