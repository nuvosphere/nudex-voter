package tss

import (
	"context"
	"fmt"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
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
			p,
			bus,
			stateDB,
			voterContract,
		),
	}
}

func (t *Service) Start(ctx context.Context) {
	t.scheduler.Start()

	<-ctx.Done()
	log.Info("TSSService is stopping...")
	t.Stop()
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
				info := fmt.Sprintf("tss signature sessionID=%v, groupID=%v, taskID=%v", result.SessionID, result.GroupID, result.TaskID)
				t.scheduler.AddDiscussedTask(result.TaskID) // todo

				if result.Err != nil {
					log.Errorf("%s, result error:%v", info, result.Err)
				} else {
					t.handleSigFinish(ctx, result.Data)
				}
			}
		}
	}()
}

func (t *Service) handleSigFinish(ctx context.Context, signatureData *tsscommon.SignatureData) {
	//t.rw.Lock()
	//
	//log.Infof("sig finish, taskId:%d, R:%x, S:%x, V:%x", t.stateDB.TssState.CurrentTask.TaskId, signatureData.R, signatureData.S, signatureData.SignatureRecovery)
	//
	//if t.stateDB.TssState.CurrentTask.Submitter == t.localAddress.Hex() {
	//	buf := bytes.NewReader(t.stateDB.TssState.CurrentTask.Context)
	//
	//		// @todo
	//		// generate wallet and send to chain
	//		address := wallet.GenerateAddressByPath(
	//			*(t.scheduler.MasterPublicKey()),
	//			uint32(coinType),
	//			createWalletTask.Account,
	//			createWalletTask.Index,
	//		)
	//		log.Infof("user account address: %s", address)
	//	}
	//}
	//
	//t.rw.Unlock()
}
