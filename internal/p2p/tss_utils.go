package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	log "github.com/sirupsen/logrus"
)

func saveTSSData(data *keygen.LocalPartySaveData) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Unable to serialize TSS data: %v", err)
		return
	}

	dataDir := "tss_data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Errorf("Failed to create TSS data directory: %v", err)
		return
	}

	filePath := filepath.Join(dataDir, "tss_key_data.json")
	if err := os.WriteFile(filePath, dataBytes, 0644); err != nil {
		log.Errorf("Failed to save TSS data to file: %v", err)
		return
	}

	log.Infof("TSS data successfully saved to: %s", filePath)
}

func loadTSSData() (*keygen.LocalPartySaveData, error) {
	dataDir := "tss_data"
	filePath := filepath.Join(dataDir, "tss_key_data.json")

	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read TSS data file: %v", err)
	}

	var data keygen.LocalPartySaveData
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, fmt.Errorf("unable to deserialize TSS data: %v", err)
	}

	log.Infof("Successfully loaded TSS data from %s", filePath)
	return &data, nil
}

func publishTSSMessage(ctx context.Context, msg tss.Message, msgType MessageType) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("Unable to serialize TSS message to JSON: %v", err)
		return
	}

	message := Message{
		MessageType: msgType,
		Content:     string(jsonMsg),
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Errorf("Unable to serialize Message struct to JSON: %v", err)
		return
	}

	err = messageTopic.Publish(ctx, messageBytes)
	if err != nil {
		log.Errorf("Unable to publish TSS message: %v", err)
		return
	}

	log.Infof("TSS message successfully broadcasted")
}
