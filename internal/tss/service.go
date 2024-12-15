package tss

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	scheduler *Scheduler
}

func NewTssService(p p2p.P2PService, dbm *db.DatabaseManager, bus eventbus.Bus, voterContract layer2.VoterContract) *Service {
	scheduler := NewScheduler(
		true,
		p,
		bus,
		state.NewContractState(dbm.GetL2InfoDB()),
		voterContract,
		crypto.PubkeyToAddress(config.L2PrivateKey.PublicKey),
	)

	return &Service{
		scheduler: scheduler,
	}
}

func (t *Service) Start(ctx context.Context) {
	t.scheduler.Start()

	<-ctx.Done()
	log.Info("TSSService is stopping...")
}

func (t *Service) Stop(ctx context.Context) {
	t.scheduler.Stop()
}
