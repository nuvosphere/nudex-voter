package config

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	golog "github.com/ipfs/go-log/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort            string
	Libp2pPort          int
	Libp2pBootNodes     string
	BTCRPC              string
	BTCRPC_USER         string
	BTCRPC_PASS         string
	BTCStartHeight      int
	L2RPC               string
	L2JwtSecret         string
	L2StartHeight       int
	L2Confirmations     int
	L2MaxBlockRange     int
	L2RequestInterval   time.Duration
	FireblocksPubKey    string
	FireblocksPrivKey   string
	TssPublicKeys       []*ecdsa.PublicKey
	TssThreshold        int
	TssSigTimeout       time.Duration
	EnableWebhook       bool
	EnableRelayer       bool
	DbDir               string
	LogLevel            logrus.Level
	VotingContract      string
	AccountContract     string
	WithdrawContract    string
	TaskManagerContract string
	ParticipantContract string
	DepositContract     string
	L2PrivateKey        *ecdsa.PrivateKey
	L2ChainId           *big.Int
}

var AppConfig Config

func InitConfig() {
	viper.AutomaticEnv()

	// Default config
	viper.SetDefault("HTTP_PORT", "8080")
	viper.SetDefault("RPC_PORT", "50051")
	viper.SetDefault("LIBP2P_PORT", 4001)
	viper.SetDefault("LIBP2P_BOOT_NODES", "")
	viper.SetDefault("BTC_RPC", "http://localhost:8332")
	viper.SetDefault("BTC_RPC_USER", "")
	viper.SetDefault("BTC_RPC_PASS", "")
	viper.SetDefault("BTC_START_HEIGHT", 0)
	viper.SetDefault("L2_RPC", "http://localhost:8545")
	viper.SetDefault("L2_JWT_SECRET", "")
	viper.SetDefault("L2_START_HEIGHT", 0)
	viper.SetDefault("L2_CONFIRMATIONS", 3)
	viper.SetDefault("L2_MAX_BLOCK_RANGE", 500)
	viper.SetDefault("L2_REQUEST_INTERVAL", "10s")
	viper.SetDefault("L2_PRIVATE_KEY", "")
	viper.SetDefault("L2_CHAIN_ID", "2345")
	viper.SetDefault("ENABLE_WEBHOOK", true)
	viper.SetDefault("ENABLE_RELAYER", true)
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("DB_DIR", "/app/db")
	viper.SetDefault("VOTING_CONTRACT", "")
	viper.SetDefault("ACCOUNT_CONTRACT", "")
	viper.SetDefault("OPERATIONS_CONTRACT", "")
	viper.SetDefault("PARTICIPANT_CONTRACT", "")
	viper.SetDefault("DEPOSIT_CONTRACT", "")
	viper.SetDefault("FIREBLOCKS_PUBKEY", "")
	viper.SetDefault("FIREBLOCKS_PRIVKEY", "")
	viper.SetDefault("TSS_PUBLIC_KEYS", "")
	viper.SetDefault("TSS_THRESHOLD", "1")
	viper.SetDefault("TSS_SIG_TIMEOUT", "300s")

	logLevel, err := logrus.ParseLevel(strings.ToLower(viper.GetString("LOG_LEVEL")))
	if err != nil {
		logrus.Fatalf("Invalid log level: %v", err)
	}

	l2PrivateKey, err := crypto.HexToECDSA(viper.GetString("L2_PRIVATE_KEY"))
	if err != nil {
		logrus.Fatalf("Failed to load l2 private key: %v, given length %d", err, len(viper.GetString("L2_PRIVATE_KEY")))
	}

	l2ChainId, err := strconv.ParseInt(viper.GetString("L2_CHAIN_ID"), 10, 64)
	if err != nil {
		logrus.Fatalf("Failed to parse l2 chain id: %v", err)
	}

	tssPublicKeys, err := ParseECDSAPublicKeys(viper.GetString("TSS_PUBLIC_KEYS"))
	if err != nil {
		logrus.Fatalf("Failed to parse tss public keys: %v", err)
	}

	AppConfig = Config{
		HTTPPort:            viper.GetString("HTTP_PORT"),
		Libp2pPort:          viper.GetInt("LIBP2P_PORT"),
		Libp2pBootNodes:     viper.GetString("LIBP2P_BOOT_NODES"),
		BTCRPC:              viper.GetString("BTC_RPC"),
		BTCRPC_USER:         viper.GetString("BTC_RPC_USER"),
		BTCRPC_PASS:         viper.GetString("BTC_RPC_PASS"),
		BTCStartHeight:      viper.GetInt("BTC_START_HEIGHT"),
		L2RPC:               viper.GetString("L2_RPC"),
		L2JwtSecret:         viper.GetString("L2_JWT_SECRET"),
		L2StartHeight:       viper.GetInt("L2_START_HEIGHT"),
		L2Confirmations:     viper.GetInt("L2_CONFIRMATIONS"),
		L2MaxBlockRange:     viper.GetInt("L2_MAX_BLOCK_RANGE"),
		L2RequestInterval:   viper.GetDuration("L2_REQUEST_INTERVAL"),
		FireblocksPubKey:    viper.GetString("FIREBLOCKS_PUBKEY"),
		FireblocksPrivKey:   viper.GetString("FIREBLOCKS_PRIVKEY"),
		TssPublicKeys:       tssPublicKeys,
		TssThreshold:        viper.GetInt("TSS_THRESHOLD"),
		TssSigTimeout:       viper.GetDuration("TSS_SIG_TIMEOUT"),
		EnableWebhook:       viper.GetBool("ENABLE_WEBHOOK"),
		EnableRelayer:       viper.GetBool("ENABLE_RELAYER"),
		DbDir:               viper.GetString("DB_DIR"),
		LogLevel:            logLevel,
		VotingContract:      viper.GetString("VOTING_CONTRACT"),
		AccountContract:     viper.GetString("ACCOUNT_CONTRACT"),
		ParticipantContract: viper.GetString("PARTICIPANT_CONTRACT"),
		TaskManagerContract: viper.GetString("TASK_CONTRACT"),
		DepositContract:     viper.GetString("DEPOSIT_CONTRACT"),
		L2PrivateKey:        l2PrivateKey,
		L2ChainId:           big.NewInt(l2ChainId),
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(AppConfig.LogLevel)

	logLvl, err := golog.LevelFromString(AppConfig.LogLevel.String())
	if err != nil {
		logrus.Fatalf("LevelFromString: %v", err)
	}

	golog.SetAllLoggers(logLvl)
}

// ParseECDSAPublicKeys parses a comma-separated string of 132-character public keys with '0x' prefix
// into an array of *ecdsa.PublicKey. It uses the secp256k1 elliptic curve.
func ParseECDSAPublicKeys(publicKeyStr string) ([]*ecdsa.PublicKey, error) {
	// Remove '0x' prefix and split the string by commas
	publicKeyHexArray := strings.Split(publicKeyStr, ",")
	for i := range publicKeyHexArray {
		publicKeyHexArray[i] = strings.TrimPrefix(publicKeyHexArray[i], "0x")
	}

	sort.Strings(publicKeyHexArray)

	publicKeys := make([]*ecdsa.PublicKey, len(publicKeyHexArray))

	for i, keyHex := range publicKeyHexArray {
		if len(keyHex) != 66 {
			return nil, errors.New("invalid compressed public key length, expected 33 bytes")
		}

		pubBytes, err := hex.DecodeString(keyHex)
		if err != nil {
			return nil, errors.New("failed to decode public key hex: " + err.Error())
		}

		pubKey, err := crypto.DecompressPubkey(pubBytes)
		if err != nil {
			return nil, errors.New("failed to decompress public key: " + err.Error())
		}

		publicKeys[i] = pubKey
	}

	return publicKeys, nil
}
