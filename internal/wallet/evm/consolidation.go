package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	log "github.com/sirupsen/logrus"
)

func (w *WalletClient) processConsolidation(task *db.ConsolidationTask) {
	balance, err := w.ContractState().GetAddressBalance(task.TargetAddress, task.ContractAddress)
	if err != nil {
		log.Errorf("Failed to get address balance for %v: %v", task.TargetAddress, err)
		return
	}

	if balance.Cmp(task.Amount) == 1 {
		to := common.HexToAddress(address.HotAddressOfEth(w.tss.ECPoint(types.ChainEthereum)))
		_, err := w.signTask(common.HexToAddress(task.TargetAddress), to, task.Amount.BigInt(), task.TaskID(), task.Ticker, db.TaskTypeConsolidation)
		if err != nil {
			log.Errorf("Failed to sign task for %v: %v", task.TargetAddress, err)
		}
	}
}
