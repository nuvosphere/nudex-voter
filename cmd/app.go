package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nuvosphere/nudex-voter/internal/btc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/http"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/task"
	"github.com/nuvosphere/nudex-voter/internal/tss"
	"github.com/samber/lo/parallel"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	DatabaseManager *db.DatabaseManager
	State           *state.State
	LibP2PService   *p2p.LibP2PService
	Layer2Listener  *layer2.Layer2Listener
	BTCListener     *btc.BTCListener
	TssService      *tss.TSSService
	TaskService     *task.TaskService
	HTTPServer      *http.HTTPServer
}

func NewApplication() *Application {
	config.InitConfig()

	dbm := db.NewDatabaseManager()
	state := state.InitializeState(dbm)
	libP2PService := p2p.NewLibP2PService(state)
	layer2Listener := layer2.NewLayer2Listener(libP2PService, state, dbm)
	btcListener := btc.NewBTCListener(libP2PService, state, dbm)
	tssService := tss.NewTssService(libP2PService, dbm, state)
	taskService := task.NewTaskService(state, dbm, tssService)
	httpServer := http.NewHTTPServer(libP2PService, state, dbm)

	return &Application{
		DatabaseManager: dbm,
		State:           state,
		LibP2PService:   libP2PService,
		Layer2Listener:  layer2Listener,
		BTCListener:     btcListener,
		TssService:      tssService,
		TaskService:     taskService,
		HTTPServer:      httpServer,
	}
}

func (app *Application) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	var moules []Module
	if config.AppConfig.EnableRelayer {
		moules = append(moules, app.Layer2Listener)
	}

	otherModules := []Module{
		app.LibP2PService,
		app.BTCListener,
		app.TssService,
		app.TaskService,
		app.HTTPServer,
	}
	moules = append(moules, otherModules...)
	parallel.ForEach(moules, func(module Module, _ int) { module.Start(ctx) })

	<-stop
	log.Info("Receiving exit signal...")

	cancel()

	app.State.Bus().Close()
	log.Info("Server stopped")
}

func main() {
	app := NewApplication()
	app.Run()
}

type Module interface {
	Start(ctx context.Context)
}
