package tss

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func createPartyIDs(publicKeys []*ecdsa.PublicKey) tss.SortedPartyIDs {
	var tssAllPartyIDs = make(tss.UnSortedPartyIDs, len(publicKeys))
	for i, publicKey := range publicKeys {
		key, _ := new(big.Int).SetString(ConvertPubKeyToHex(publicKey), 16)
		tssAllPartyIDs[i] = tss.NewPartyID(
			strconv.Itoa(i),
			crypto.PubkeyToAddress(*publicKeys[i]).Hex(),
			key,
		)
	}
	return tss.SortPartyIDs(tssAllPartyIDs)
}

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

func PublicKeysToHex(pubKeys []*ecdsa.PublicKey) []string {
	hexStrings := make([]string, len(pubKeys))
	for i, pubKey := range pubKeys {
		hexStrings[i] = ConvertPubKeyToHex(pubKey)
	}
	return hexStrings
}

func ConvertPubKeyToHex(pubKey *ecdsa.PublicKey) string {
	if pubKey == nil {
		return ""
	}

	var prefix byte
	if pubKey.Y.Bit(0) == 0 {
		prefix = 0x02 // Even Y-coordinate
	} else {
		prefix = 0x03 // Odd Y-coordinate
	}

	// Create the public key byte slice with the prefix
	pubKeyBytes := append([]byte{prefix}, pubKey.X.Bytes()...)

	// Convert to hex string
	return hex.EncodeToString(pubKeyBytes)
}

func CompareStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func AddressIndex(publicKeys []*ecdsa.PublicKey, tssAddress string) int {
	for i, pubKey := range publicKeys {
		address := crypto.PubkeyToAddress(*pubKey).Hex()
		if address == tssAddress {
			return i // Return the index if a match is found
		}
	}
	return -1 // Return -1 if not found
}
