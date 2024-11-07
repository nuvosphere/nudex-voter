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
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/types"
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

func createPartyIDsByAddress(publicKeys []common.Address) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(publicKeys))

	for i, publicKey := range publicKeys {
		key := new(big.Int).SetBytes(publicKey.Bytes())
		tssAllPartyIDs[i] = tss.NewPartyID(
			publicKey.Hex(),
			publicKey.Hex(),
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

func serializeTaskMessageToBytes(nonce uint64, baseTask types.Task) ([]byte, error) {
	buf := new(bytes.Buffer)
	writeField := func(data interface{}) error {
		return binary.Write(buf, binary.BigEndian, data)
	}

	switch task := baseTask.(type) {
	case *types.CreateWalletTask:
		if err := writeField(task.GetTaskID()); err != nil {
			return nil, err
		}

		if err := writeField(task.Account); err != nil {
			return nil, err
		}

		if err := writeField(task.User); err != nil {
			return nil, err
		}

		if err := writeField(task.Chain); err != nil {
			return nil, err
		}

	case *types.DepositTask:
		if err := writeField(task.GetTaskID()); err != nil {
			return nil, err
		}

		if err := writeField(task.TargetAddress); err != nil {
			return nil, err
		}

		if err := writeField(task.Amount); err != nil {
			return nil, err
		}

		if err := writeField(task.ChainId); err != nil {
			return nil, err
		}

		if err := writeField(task.Ticker); err != nil {
			return nil, err
		}

		if err := writeField(task.BlockHeight); err != nil {
			return nil, err
		}

		if err := writeField(task.TxHash); err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported task type: %T", task)
	}

	bufBytes := buf.Bytes()
	length := len(bufBytes)
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)

	return append(append(nonceBytes, lengthBytes...), bufBytes...), nil
}

func getRequestId(task *db.Task) string {
	buf := bytes.NewReader(task.Context)

	var taskType int32
	_ = binary.Read(buf, binary.LittleEndian, &taskType)

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

func getCoinTypeByChain(chain uint8) int {
	switch chain {
	case types.WalletTypeEVM:
		return 60
	case types.WalletTypeBTC:
		return 0
	case types.WalletTypeSOL:
		return 501
	case types.WalletTypeSUI:
		return 784
	default:
		return -1
	}
}
