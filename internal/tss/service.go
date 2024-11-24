package tss

import (
	"context"
	"fmt"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	scheduler *Scheduler
}

func NewTssService(p p2p.P2PService, stateDB *gorm.DB, bus eventbus.Bus, voterContract layer2.VoterContract) *Service {
	return &Service{
		scheduler: NewScheduler(
			true,
			p,
			bus,
			stateDB,
			voterContract,
			crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey),
		),
	}
}

func (t *Service) Start(ctx context.Context) {
	t.scheduler.Start()
	t.loop(ctx)
	log.Info("TSSService is stopping...")
}

func (t *Service) Stop() {
	t.scheduler.Stop()
}

func (t *Service) loop(ctx context.Context) {
	out := t.scheduler.sigInToOut

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("tss signature read result loop stopped")
			case result := <-out:
				log.Infof("finish consensus sessionID:%s", result.SessionID)
				info := fmt.Sprintf("tss signature sessionID=%v, groupID=%v, taskID=%v", result.SessionID, result.GroupID, result.ProposalID)
				t.scheduler.AddDiscussedTask(result.ProposalID) // todo

				if result.Err != nil {
					log.Errorf("%s, result error:%v", info, result.Err)
				} else {
					t.handleSigFinish(ctx, uint64(result.ProposalID), result.Data)
				}
			}
		}
	}()
}

func (t *Service) handleSigFinish(ctx context.Context, taskID uint64, signatureData *tsscommon.SignatureData) {
	// 1. save db
	// 2. update status
	if t.scheduler.IsProposer() {
		log.Info("proposer submit signature")

		task, err := t.scheduler.GetTask(taskID)
		if err != nil {
			log.Errorf("get task err:%v", err)
			return
		}

		if task == nil {
			log.Errorf("find no task by taskId: %d", taskID)
			return
		}
		// createWalletTask := task.
		// localPartySaveData, _, unSignMsg, operations := t.scheduler.CreateWalletProposal(createWalletTask)
		//
		// calldata := t.scheduler.voterContract.EncodeVerifyAndCall(operations, signatureData.Signature)
		// publicKey := *localPartySaveData.ECDSAData().ECDSAPub.ToECDSAPubKey()
		// verifyOk := ecdsa.Verify(&publicKey, unSignMsg.Bytes(), new(big.Int).SetBytes(signatureData.R), new(big.Int).SetBytes(signatureData.S))
		//
		//	if !verifyOk {
		//		log.Errorf("verify signature fail for taskId: %d", taskID)
		//		return
		//	}
		//
		// submitterWallet := wallet.NewWallet(config.AppConfig.L2RPC, publicKey, *config.AppConfig.L2PrivateKey)
		// tx, err := submitterWallet.BuildUnsignTx(context.Background(), common.HexToAddress(config.AppConfig.AccountContract), big.NewInt(0), calldata)
		//
		//	if err != nil {
		//		log.Fatalf("failed to build unsigned transaction: %v", err)
		//	}
		//
		// signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(config.AppConfig.L2ChainId), config.AppConfig.L2PrivateKey)
		//
		//	if err != nil {
		//		log.Fatalf("failed to sign transaction: %v", err)
		//	}
		//
		// err = submitterWallet.SendSingedTx(context.Background(), signedTx)
		//
		//	if err != nil {
		//		log.Fatalf("failed to send transaction: %v", err)
		//	}
		//
		// receipt, err := submitterWallet.WaitTxSuccess(ctx, signedTx.Hash())
		//
		//	if err != nil {
		//		log.Fatalf("failed to wait transaction success: %v", err)
		//	}
		//
		//	if receipt.Status == 0 {
		//		log.Errorf("failed to submit transaction for taskId: %d,txHash: %s", taskID, signedTx.Hash().String())
		//	} else {
		//
		//		log.Infof("successfully submitted transaction for taskId: %d, txHash: %s", taskID, signedTx.Hash().String())
		//	}
	}
}
