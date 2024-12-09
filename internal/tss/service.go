package tss

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	scheduler *Scheduler
	wallet    *wallet.Wallet
}

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, bus eventbus.Bus, voterContract layer2.VoterContract) *Service {
	scheduler := NewScheduler(
		true,
		p,
		bus,
		state.NewContractState(dbm.GetL2SyncDB()),
		voterContract,
		crypto.PubkeyToAddress(config.L2PrivateKey.PublicKey),
	)

	return &Service{
		scheduler: scheduler,
		wallet:    wallet.NewWallet(config.AppConfig.L2Rpc, config.L2PrivateKey),
	}
}

func (t *Service) Start(ctx context.Context) {
	t.scheduler.handleSigFinish = t.handleSigFinish
	t.scheduler.Start()

	<-ctx.Done()
	log.Info("TSSService is stopping...")
}

func (t *Service) Stop(ctx context.Context) {
	t.scheduler.Stop()
}

func (t *Service) handleSigFinish(operations *Operations) {
	// 1. save db
	// 2. update status
	if t.scheduler.IsProposer() {
		log.Info("proposer submit signature")

		calldata := t.scheduler.voterContract.EncodeVerifyAndCall(operations.Operation, operations.Signature)

		log.Infof("calldata: %x, signature: %x,nonce: %v,DataHash: %v, hash: %v", calldata, operations.Signature, operations.Nonce, operations.DataHash, operations.Hash)

		data, _ := json.Marshal(operations)
		tx, err := t.wallet.BuildUnsignTx(
			context.Background(),
			t.scheduler.LocalSubmitter(),
			common.HexToAddress(config.AppConfig.VotingContract),
			big.NewInt(0), // todo
			calldata,
			&db.Operations{
				Nonce: decimal.NewFromBigInt(operations.Nonce, 0),
				Data:  string(data),
			}, nil, nil,
		)
		if err != nil {
			log.Fatalf("failed to build unsigned transaction: %v", err)
		}

		chainId, err := t.wallet.ChainID(context.Background())
		if err != nil {
			log.Fatalf("failed to ChainID: %v", err)
		}
		signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainId), config.L2PrivateKey)
		if err != nil {
			log.Fatalf("failed to sign transaction: %v", err)
		}

		err = t.wallet.SendSingedTx(context.Background(), signedTx)
		if err != nil {
			log.Fatalf("failed to send transaction: %v", err)
		}
		// updated status to pending

		receipt, err := t.wallet.WaitTxSuccess(context.Background(), signedTx.Hash())
		if err != nil {
			log.Fatalf("failed to wait transaction success: %v", err)
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
