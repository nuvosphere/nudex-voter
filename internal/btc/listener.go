package btc

import (
	"context"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
)

type BTCListener struct {
	libp2p *p2p.Service
	dbm    *db.DatabaseManager
	state  *state.State

	notifier *BTCNotifier
}

func NewBTCListener(libp2p *p2p.Service, state *state.State, dbm *db.DatabaseManager) *BTCListener {
	db := dbm.GetBtcCacheDB()
	cache := NewBTCCache(db)
	poller := NewBTCPoller(state, db)

	connConfig := &rpcclient.ConnConfig{
		Host:         config.AppConfig.BtcRpc,
		User:         config.AppConfig.BtcRpcUser,
		Pass:         config.AppConfig.BtcRpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connConfig, nil)
	if err != nil {
		log.Fatalf("Failed to start bitcoin client: %v", err)
	}

	notifier := NewBTCNotifier(client, cache, poller)

	return &BTCListener{
		libp2p:   libp2p,
		dbm:      dbm,
		state:    state,
		notifier: notifier,
	}
}

func (bl *BTCListener) Start(ctx context.Context) {
	go bl.notifier.Start(ctx)
	log.Info("BTCListener started all modules")

	<-ctx.Done()
	log.Info("BTCListener is stopping...")
}
