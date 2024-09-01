package http

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// UNCHECKED this is an example
type FireblocksWebhookRequest struct {
	Event     string `json:"event"`
	TxID      string `json:"txId"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

// handleFireblocksWebhook process webhook callback of Fireblocks, should validate request data first
func handleFireblocksWebhook(c *gin.Context) {
	var req FireblocksWebhookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Infof("Received webhook event: %s, TxID: %s, Status: %s, Timestamp: %d",
		req.Event, req.TxID, req.Status, req.Timestamp)

	// TODO: webhook parameters check, signature and others..
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// handleFireblocksCosigner process cosigner callback of Fireblocks, it will check tx and sign
func handleFireblocksCosigner(c *gin.Context) {
	// TODO: cosigner tx validation
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
