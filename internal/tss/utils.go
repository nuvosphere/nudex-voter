package tss

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sort"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db/task"
	log "github.com/sirupsen/logrus"
)

func createPartyIDs(publicKeys []*ecdsa.PublicKey) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(publicKeys))

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

func createPartyIDsByAddress(addressList []common.Address) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	for i, address := range addressList {
		key := new(big.Int).SetBytes(address.Bytes())
		tssAllPartyIDs[i] = tss.NewPartyID(
			address.Hex(),
			address.Hex(),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func createOldPartyIDsByAddress(addressList []common.Address) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	for i, address := range addressList {
		key := new(big.Int).SetBytes(address.Bytes())
		key = new(big.Int).Add(key, big.NewInt(1)) // key + 1
		tssAllPartyIDs[i] = tss.NewPartyID(
			key.Text(16),
			key.Text(16),
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
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		log.Errorf("Failed to create TSS data directory: %v", err)
		return err
	}

	filePath := filepath.Join(dataDir, "tss_key_data.json")
	if err := os.WriteFile(filePath, dataBytes, 0o600); err != nil {
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

func AddressIndex(addressList []common.Address, tssAddress common.Address) int {
	for i, address := range addressList {
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

func serializeMessageToBeSigned(nonce uint64, data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	bufBytes := buf.Bytes()
	length := len(bufBytes)
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)

	return append(append(nonceBytes, lengthBytes...), data...), nil
}

func getCoinTypeByChain(chain uint8) int {
	switch chain {
	case task.WalletTypeEVM:
		return 60
	case task.WalletTypeBTC:
		return 0
	case task.WalletTypeSOL:
		return 501
	case task.WalletTypeSUI:
		return 784
	default:
		return -1
	}
}
