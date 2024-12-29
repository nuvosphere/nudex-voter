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

func (w *WalletClient) signTask(from, to, contract common.Address, amount *big.Int) error {
	panic("todo")
}

func (w *WalletClient) processWithdrawTxSign(task *db.WithdrawalTask) {
	log.Debugf("processTxSign taskId: %v", task.TaskID())

	hotAddress := common.HexToAddress(address.HotAddressOfEth(w.tss.ECPoint(w.ChainType())))
	to := common.HexToAddress(task.TargetAddress)
	var tx *db.EvmTransaction
	var err error
	tokenInfo, err := w.ContractState().GetTokenInfo(task.Ticker)
	if err != nil {
		log.Errorf("Failed to get token info: %v", err)
		return
	}

	switch tokenInfo.AssetType {
	case types.AssetTypeMain:
		tx, err = w.BuildUnSignTx(
			hotAddress,
			to,
			task.Amount.BigInt(),
			nil,
			db.TaskTypeWithdrawal,
			task.TaskId,
		)
	case types.AssetTypeErc20:
		tx, err = w.BuildUnSignTx(

			hotAddress,
			common.HexToAddress(tokenInfo.ContractAddress),
			nil,
			contracts.EncodeTransferOfERC20(hotAddress, to, task.Amount.BigInt()),
			db.TaskTypeWithdrawal,
			task.TaskId,
		)
	default:
		log.Errorf("unknown asset type: %v", tokenInfo.AssetType)
		return
	}
	if err != nil {
		log.Errorf("failed to build unsign tx: %v", err)
		return
	}
	ctx := w.NewTxContext(tx)
	w.pendingTx.Store(ctx.TxHash(), ctx)
	defer w.pendingTx.Delete(ctx.TxHash())

	err = w.signTx(ctx)
	if err != nil {
		log.Errorf("failed to signTx tx: %v", err)
	}
	err = w.SendSingedTx(ctx)
	if err != nil {
		log.Errorf("send transaction err: %v", err)
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
		log.Errorf("failed to submit transaction for taskId: %d,txHash: %v", ctx.SeqID(), ctx.TxHash())
	} else {
		// updated status to completed
		log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.SeqID(), ctx.TxHash())
	}
}

func (w *WalletClient) processTxSignResult(res *suite.SignRes) {
	txCtx, ok := w.pendingTx.Load(res.DataDigest)
	if ok {
		switch ctx := txCtx.(type) {
		case *TxContext:
			ctx.sig = res.Signature
			ctx.notify <- res.Err
		}
	}
}
