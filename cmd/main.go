package main

import (
	"github.com/nuvosphere/nudex-voter/internal/btc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/http"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
)

func main() {
	// Init configuration
	config.InitConfig()

	// Init database: sqlite
	db.InitDB()

	// Start services
	if config.AppConfig.EnableRelayer {
		go layer2.StartLayer2Monitor()
	}
	go http.StartHTTPServer()
	go p2p.StartLibp2p()
	go btc.StartBTCListener()

	// Blocking the main thread
	select {}
}
