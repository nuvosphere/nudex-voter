package rpc

import (
	"encoding/json"
	"net/http"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/btc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
)

type UTXOService struct {
	client *rpcclient.Client
}

func NewUTXOService() (*UTXOService, error) {
	connConfig := &rpcclient.ConnConfig{
		Host:         config.AppConfig.BTCRPC,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connConfig, nil)
	if err != nil {
		return nil, err
	}

	return &UTXOService{client: client}, nil
}

func (s *UTXOService) HandleSubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var tx wire.MsgTx
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid transaction format", http.StatusBadRequest)
		return
	}

	spvProof, err := btc.GenerateSPVProof(&tx)
	if err != nil {
		http.Error(w, "Failed to generate SPV proof", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"spv_proof": spvProof,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StartUTXOService() {
	service, err := NewUTXOService()
	if err != nil {
		log.Fatalf("Failed to create UTXO service: %v", err)
	}

	http.HandleFunc("/submit_transaction", service.HandleSubmitTransaction)
	// Use configuration port
	addr := ":" + config.AppConfig.HTTPPort
	log.Infof("RPC server is running on port %s", config.AppConfig.HTTPPort)
	log.Fatal(http.ListenAndServe(addr, nil))
}
