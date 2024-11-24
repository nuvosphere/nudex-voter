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
					log.Infof("finish consensus:%s", info)
					t.handleSigFinish(ctx, result.Data)
				}
			}
		}
	}()
}

func (t *Service) handleSigFinish(ctx context.Context, signatureData *tsscommon.SignatureData) {
	// 1. save db
	// 2. update status
	if t.scheduler.IsProposer() {
		// 2. send signature data to blockchain
		log.Info("proposer submit signature")
	}
}
