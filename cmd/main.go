package main

import (
	"context"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	tsslib "github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/btc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/http"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/rpc"
	"github.com/nuvosphere/nudex-voter/internal/tss"
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

	ctx := context.Background()
	tssKeyInCh := make(chan tss.KeygenMessage)
	tssKeyOutCh := make(chan tsslib.Message)
	tssKeyEndCh := make(chan *keygen.LocalPartySaveData)
	tssSignInCh := make(chan tss.SigningMessage)
	tssSignOutCh := make(chan tsslib.Message)
	tssSignEndCh := make(chan *common.SignatureData)

	go p2p.StartLibp2p(tssKeyInCh, tssSignInCh)
	go tss.HandleKeygenMessages(ctx, tssKeyInCh, tssKeyOutCh, tssKeyEndCh)
	go tss.HandleSigningMessages(ctx, tssSignInCh, tssSignOutCh, tssSignEndCh)
	go btc.StartBTCListener()
	go rpc.StartUTXOService()

	// Blocking the main thread
	select {}
}
