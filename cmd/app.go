package main

import (
	"context"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/btc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/http"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/rpc"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Application struct {
	DatabaseManager *db.DatabaseManager
	State           *state.State
	LibP2PService   *p2p.LibP2PService
	Layer2Listener  *layer2.Layer2Listener
	BTCListener     *btc.BTCListener
	TssService      *tss.TSSService
	HTTPServer      *http.HTTPServer
	TssKeyInCh      chan types.KeygenMessage
	TssKeyOutCh     chan tsslib.Message
	TssKeyEndCh     chan *keygen.LocalPartySaveData
}

func NewApplication() *Application {
	config.InitConfig()

	dbm := db.NewDatabaseManager()
	state := state.InitializeState(dbm)
	libP2PService := p2p.NewLibP2PService(state)
	layer2Listener := layer2.NewLayer2Listener(libP2PService, state, dbm)
	btcListener := btc.NewBTCListener(libP2PService, state, dbm)
	tssService := tss.NewTssService(libP2PService, state)
	httpServer := http.NewHTTPServer(libP2PService, state, dbm)

	return &Application{
		DatabaseManager: dbm,
		State:           state,
		LibP2PService:   libP2PService,
		Layer2Listener:  layer2Listener,
		BTCListener:     btcListener,
		TssService:      tssService,
		HTTPServer:      httpServer,
		TssKeyInCh:      make(chan types.KeygenMessage),
	}
}

func (app *Application) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if config.AppConfig.EnableRelayer {
			app.Layer2Listener.Start(ctx)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		app.LibP2PService.Start(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		app.BTCListener.Start(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		go app.TssService.Start(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		go rpc.StartUTXOService()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		app.HTTPServer.Start(ctx)
	}()

	<-stop
	log.Info("Receiving exit signal...")

	cancel()

	wg.Wait()
	log.Info("Server stopped")
}

func main() {
	app := NewApplication()
	app.Run()
}
