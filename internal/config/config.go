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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Env                 string        `yaml:"env"` // dev、test、prod
	HttpPort            string        `yaml:"httpPort"`
	P2pPort             int           `yaml:"p2PPort"`
	P2pBootNodes        string        `yaml:"p2PBootNodes"`
	BtcRpc              string        `yaml:"btcRpc"`
	BtcRpcUser          string        `yaml:"btcRpcUser"`
	BtcRpcPass          string        `yaml:"btcRpcPass"`
	BtcStartHeight      int           `yaml:"btcStartHeight"`
	BtcPrivKey          string        `yaml:"btcPrivKey"`
	SolRPC              string        `yaml:"solRPC"`
	BolPrivKey          string        `yaml:"solPrivKey"`
	SuiRPC              string        `yaml:"suiRPC"`
	SuiPrivKey          string        `yaml:"suiPrivKey"`
	L2Rpc               string        `yaml:"l2Rpc"`
	L2JwtSecret         string        `yaml:"l2JwtSecret"`
	L2StartHeight       int           `yaml:"l2StartHeight"`
	L2Confirmations     int           `yaml:"l2Confirmations"`
	L2MaxBlockRange     int           `yaml:"l2MaxBlockRange"`
	L2RequestInterval   time.Duration `yaml:"l2RequestInterval"`
	FireblocksPubKey    string        `yaml:"fireblocksPubKey"`
	FireblocksPrivKey   string        `yaml:"fireblocksPrivKey"`
	TssThreshold        int           `yaml:"tssThreshold"`
	TssSigTimeout       time.Duration `yaml:"tssSigTimeout"`
	EnableWebhook       bool          `yaml:"enableWebhook"`
	EnableRelayer       bool          `yaml:"enableRelayer"`
	DbDir               string        `yaml:"dbDir"`
	VotingContract      string        `yaml:"votingContract"`
	AccountContract     string        `yaml:"accountContract"`
	WithdrawContract    string        `yaml:"withdrawContract"`
	TaskManagerContract string        `yaml:"taskManagerContract"`
	ParticipantContract string        `yaml:"participantContract"`
	DepositContract     string        `yaml:"depositContract"`
	L2ChainID           string        `yaml:"l2ChainID"`
	LogLevel            string        `yaml:"logLevel"`
	L2PrivateKey        string        `yaml:"l2PrivateKey"`
	TssPublicKeys       []string      `yaml:"tssPublicKeys"`

	// TssPublicKeys []*ecdsa.PublicKey
	// L2PrivateKey  *ecdsa.PrivateKey
	// L2ChainId     *big.Int
}

func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

var (
	AppConfig     Config
	TssPublicKeys []*ecdsa.PublicKey
	L2PrivateKey  *ecdsa.PrivateKey
	L2ChainId     *big.Int
)

func InitConfig(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/etc")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	notFound := viper.ConfigFileNotFoundError{}
	if !errors.As(err, &notFound) {
		logrus.Info("load yaml config")
		err = viper.Unmarshal(&AppConfig)
		if err != nil {
			panic(err)
		}

		setL2PrivateKey(AppConfig.L2PrivateKey)
		setL2ChainId(AppConfig.L2ChainID)
		setTssPublicKeys(strings.Join(AppConfig.TssPublicKeys, ","))
		setLogLevel()
		return
	}

	logrus.Info("load .env config")
	viper.AutomaticEnv()
	// viper.Debug()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

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

	setL2PrivateKey(viper.GetString("L2_PRIVATE_KEY"))
	setL2ChainId(viper.GetString("L2_CHAIN_ID"))
	setTssPublicKeys(viper.GetString("TSS_PUBLIC_KEYS"))

	AppConfig = Config{
		HttpPort:            viper.GetString("HTTP_PORT"),
		P2pPort:             viper.GetInt("LIBP2P_PORT"),
		P2pBootNodes:        viper.GetString("LIBP2P_BOOT_NODES"),
		BtcRpc:              viper.GetString("BTC_RPC"),
		BtcRpcUser:          viper.GetString("BTC_RPC_USER"),
		BtcRpcPass:          viper.GetString("BTC_RPC_PASS"),
		BtcStartHeight:      viper.GetInt("BTC_START_HEIGHT"),
		L2Rpc:               viper.GetString("L2_RPC"),
		L2JwtSecret:         viper.GetString("L2_JWT_SECRET"),
		L2StartHeight:       viper.GetInt("L2_START_HEIGHT"),
		L2Confirmations:     viper.GetInt("L2_CONFIRMATIONS"),
		L2MaxBlockRange:     viper.GetInt("L2_MAX_BLOCK_RANGE"),
		L2RequestInterval:   viper.GetDuration("L2_REQUEST_INTERVAL"),
		FireblocksPubKey:    viper.GetString("FIREBLOCKS_PUBKEY"),
		FireblocksPrivKey:   viper.GetString("FIREBLOCKS_PRIVKEY"),
		TssThreshold:        viper.GetInt("TSS_THRESHOLD"),
		TssSigTimeout:       viper.GetDuration("TSS_SIG_TIMEOUT"),
		EnableWebhook:       viper.GetBool("ENABLE_WEBHOOK"),
		EnableRelayer:       viper.GetBool("ENABLE_RELAYER"),
		DbDir:               viper.GetString("DB_DIR"),
		LogLevel:            viper.GetString("LOG_LEVEL"),
		VotingContract:      viper.GetString("VOTING_CONTRACT"),
		AccountContract:     viper.GetString("ACCOUNT_CONTRACT"),
		ParticipantContract: viper.GetString("PARTICIPANT_CONTRACT"),
		TaskManagerContract: viper.GetString("TASK_CONTRACT"),
		DepositContract:     viper.GetString("DEPOSIT_CONTRACT"),
	}
	setLogLevel()
}

func setTssPublicKeys(tssList string) {
	tssPublicKeys, err := ParseECDSAPublicKeys(tssList)
	if err != nil {
		logrus.Fatalf("Failed to parse tss public keys: %v", err)
	}

	TssPublicKeys = tssPublicKeys
}

func setL2ChainId(chainId string) {
	l2ChainId, err := strconv.ParseInt(chainId, 10, 64)
	if err != nil {
		logrus.Fatalf("Failed to parse l2 chain id: %v", err)
	}

	L2ChainId = big.NewInt(l2ChainId)
}

func setL2PrivateKey(pk string) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		logrus.Fatalf("Failed to load l2 private key: %v, given length %d", err, len(pk))
	}
	L2PrivateKey = privateKey
}

func setLogLevel() {
	logrus.SetOutput(os.Stdout)
	logLvl, err := logrus.ParseLevel(AppConfig.LogLevel)
	if err != nil {
		logLvl = logrus.WarnLevel
	}

	logrus.SetLevel(logLvl)

	if !AppConfig.IsProd() {
		logrus.SetReportCaller(true)
	}
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
