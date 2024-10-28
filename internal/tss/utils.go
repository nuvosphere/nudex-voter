package tss

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	log "github.com/sirupsen/logrus"
)

func createPartyIDs(publicKeys []*ecdsa.PublicKey) tss.SortedPartyIDs {
	var tssAllPartyIDs = make(tss.UnSortedPartyIDs, len(publicKeys))
	for i, publicKey := range publicKeys {
		key, _ := new(big.Int).SetString(ConvertPubKeyToHex(publicKey), 16)
		tssAllPartyIDs[i] = tss.NewPartyID(
			crypto.PubkeyToAddress(*publicKeys[i]).Hex(),
			crypto.PubkeyToAddress(*publicKeys[i]).Hex(),
			key,
		)
	}
	return tss.SortPartyIDs(tssAllPartyIDs)
}

func saveTSSData(data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Unable to serialize TSS data: %v", err)
		return err
	}

	dataDir := filepath.Join(config.AppConfig.DbDir, "tss_data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Errorf("Failed to create TSS data directory: %v", err)
		return err
	}

	filePath := filepath.Join(dataDir, "tss_key_data.json")
	if err := os.WriteFile(filePath, dataBytes, 0644); err != nil {
		log.Errorf("Failed to save TSS data to file: %v", err)
		return err
	}

	log.Infof("TSS data successfully saved to: %s", filePath)
	return nil
}

func LoadTSSData() (*keygen.LocalPartySaveData, error) {
	dataDir := filepath.Join(config.AppConfig.DbDir, "tss_data")
	filePath := filepath.Join(dataDir, "tss_key_data.json")

	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read TSS data file: %v", err)
	}

	var data keygen.LocalPartySaveData
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, fmt.Errorf("unable to deserialize TSS data: %v", err)
	}

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

func extractToIds(message tss.Message) []string {
	recipients := message.GetTo()

	ids := make([]string, len(recipients))

	for i, recipient := range recipients {
		ids[i] = recipient.GetId()
	}

	return ids
}

func serializeMsgSignCreateWalletMessageToBytes(task types.CreateWalletTask) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, task.TaskId); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, task.Account); err != nil {
		return nil, err
	}

	userBytes := []byte(task.User)
	chainBytes := []byte(task.Chain)

	if err := binary.Write(buf, binary.BigEndian, uint64(len(userBytes))); err != nil {
		return nil, err
	}
	if _, err := buf.Write(userBytes); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, uint64(len(chainBytes))); err != nil {
		return nil, err
	}
	if _, err := buf.Write(chainBytes); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getRequestId(task *db.Task) string {
	parts := strings.Split(task.Description, ":")
	taskType, _ := strconv.Atoi(parts[0])
	switch taskType {
	case types.TaskTypeUnknown:
		return ""
	case types.TaskTypeCreateWallet:
		return fmt.Sprintf("TSS_SIGN:CREATE_WALLET:%d", task.TaskId)
	case types.TaskTypeDeposit:
		return fmt.Sprintf("TSS_SIGN:DEPOSIT:%d", task.TaskId)
	case types.TaskTypeWithdraw:
		return fmt.Sprintf("TSS_SIGN:WITHDRAW:%d", task.TaskId)
	default:
	}
	return ""
}

func getCoinTypeByChain(chain string) int {
	if strings.HasPrefix(chain, "evm") {
		return 60
	} else if chain == "btc" {
		return 0
	} else if chain == "sol" {
		return 501
	} else if chain == "sui" {
		return 784
	}
	return -1
}
