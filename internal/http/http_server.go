package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
)

type HTTPServer struct {
	libp2p *p2p.Service
	db     *db.DatabaseManager
	state  *state.State
}

func NewHTTPServer(libp2p *p2p.Service, state *state.State, db *db.DatabaseManager) *HTTPServer {
	return &HTTPServer{
		libp2p: libp2p,
		db:     db,
		state:  state,
	}
}

// Start http server.
func (s *HTTPServer) Start(ctx context.Context) {
	r := gin.Default()

	if gin.IsDebugging() {
		r.GET("/api/v1/helloworld", s.handleHelloWorld)
	}

	if config.AppConfig.EnableWebhook {
		r.POST("/api/fireblocks/webhook", s.handleFireblocksWebhook)
	}

	if config.AppConfig.EnableRelayer {
		r.POST("/api/fireblocks/cosigner/v2/tx_sign_request", s.handleFireblocksCosignerTxSign)
	}

	// Use configuration port
	addr := ":" + config.AppConfig.HttpPort
	log.Infof("HTTP server is running on port %s", config.AppConfig.HttpPort)

	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: time.Second * 30,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	<-ctx.Done()

	log.Info("HTTP server is stopping...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("HTTP server forced to shutdown: %v", err)
	}

	log.Info("HTTP server has stopped.")
}

func (s *HTTPServer) Stop(ctx context.Context) {
	log.Info("HTTP server is shutted down...")
}

// a demo handler.
func (s *HTTPServer) handleHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": "hello world."})
}
