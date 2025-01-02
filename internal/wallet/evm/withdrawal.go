package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) signTask(from, to common.Address, amount *big.Int, taskId uint64, ticker types.Byte32, ty int) (common.Hash, error) {
	var tx *db.EvmTransaction

	var err error

	tokenInfo, err := w.ContractState().GetTokenInfo(ticker)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get token info: %w", err)
	}

	switch tokenInfo.AssetType {
	case types.AssetTypeMain:
		tx, err = w.BuildUnSignTx(
			from,
			to,
			amount,
			nil,
			ty,
			taskId,
		)
	case types.AssetTypeErc20:
		tx, err = w.BuildUnSignTx(

			from,
			common.HexToAddress(tokenInfo.ContractAddress),
			nil,
			contracts.EncodeTransferOfERC20(from, to, amount),
			db.TaskTypeWithdrawal,
			taskId,
		)
	default:
		return common.Hash{}, fmt.Errorf("unknown asset type: %v", tokenInfo.AssetType)
	}

	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to build unsign tx: %w", err)
	}

	ctx := w.NewTxContext(tx)
	w.pendingTx.Store(ctx.TxHash(), ctx)

	defer w.pendingTx.Delete(ctx.TxHash())

	err = w.signTx(ctx)
	if err != nil {
		return ctx.TxHash(), fmt.Errorf("failed to sign tx: %w", err)
	}

	err = w.SendSingedTx(ctx)
	if err != nil {
		return ctx.TxHash(), fmt.Errorf("failed to send transaction: %w", err)
	}
	// updated status to pending
	receipt, err := w.WaitTxSuccess(ctx.TxHash())
	if err != nil {
		return ctx.TxHash(), fmt.Errorf("failed to wait transaction success: %w", err)
	}

	if receipt.Status == 0 {
		// updated status to fail
		log.Errorf("failed to submit transaction for taskId: %d,txHash: %v", ctx.SeqID(), ctx.TxHash())
		return ctx.TxHash(), fmt.Errorf("failed to submit transaction for taskId: %d", ctx.SeqID())
	} else {
		// updated status to completed
		log.Infof("successfully submitted transaction for taskId: %d,txHash: %v", ctx.SeqID(), ctx.TxHash())
		return ctx.TxHash(), nil
	}
}

func (w *WalletClient) processWithdrawTxSign(task *db.WithdrawalTask) {
	log.Debugf("processTxSign taskId: %v", task.TaskID())

	hotAddress := common.HexToAddress(address.HotAddressOfEth(w.tss.ECPoint(w.ChainType())))
	to := common.HexToAddress(task.TargetAddress)

	_, err := w.signTask(hotAddress, to, task.Amount.BigInt(), task.TaskID(), task.Ticker, db.TaskTypeWithdrawal)
	if err != nil {
		log.Errorf("failed to sign task %d: %v", task.TaskID(), err)
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
