package layer2

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Layer2Listener struct {
	p2p             *p2p.Service
	db              *db.DatabaseManager
	state           *state.State
	ethClient       *ethclient.Client
	contractAddress []common.Address
	addressBind     map[common.Address]func(types.Log) error

	contractVotingManager *contracts.VotingManagerContract
	sigFinishChan         chan interface{}
}

func (l *Layer2Listener) ContractVotingManager() *contracts.VotingManagerContract {
	return l.contractVotingManager
}

func NewLayer2Listener(p *p2p.Service, state *state.State, db *db.DatabaseManager) *Layer2Listener {
	ethClient, err := DialEthClient()
	if err != nil {
		log.Fatalf("Error creating Layer2 EVM RPC client: %v", err)
	}

	self := &Layer2Listener{
		p2p:           p,
		db:            db,
		state:         state,
		ethClient:     ethClient,
		sigFinishChan: make(chan interface{}, 256),
	}

	self.addressBind = map[common.Address]func(types.Log) error{
		contracts.VotingAddress:      self.processVotingLog,
		contracts.AccountAddress:     self.processAccountLog,
		contracts.ParticipantAddress: self.processParticipantLog,
		contracts.TaskAddress:        self.processTaskLog,
		contracts.DepositAddress:     self.processDepositLog,
	}
	self.contractAddress = lo.MapToSlice(
		self.addressBind,
		func(item common.Address, _ func(log2 types.Log) error) common.Address { return item },
	)

	contractVotingManager, err := contracts.NewVotingManagerContract(contracts.VotingAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract VotingManager: %v", err)
	}

	self.contractVotingManager = contractVotingManager

	return self
}

// New an eth client.
func DialEthClient() (*ethclient.Client, error) {
	var opts []rpc.ClientOption

	if config.AppConfig.L2JwtSecret != "" {
		jwtSecret := common.FromHex(strings.TrimSpace(config.AppConfig.L2JwtSecret))
		if len(jwtSecret) != 32 {
			return nil, errors.New("jwt secret is not a 32 bytes hex string")
		}

		var jwtKey [32]byte

		copy(jwtKey[:], jwtSecret)
		opts = append(opts, rpc.WithHTTPAuth(node.NewJWTAuth(jwtKey)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Dial the Ethereum node with optional JWT authentication
	client, err := rpc.DialOptions(ctx, config.AppConfig.L2RPC, opts...)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(client), nil
}

func (l *Layer2Listener) Start(ctx context.Context) {
	// Get latest sync height
	var syncStatus db.EVMSyncStatus

	relayerDB := l.db.GetRelayerDB()

	result := relayerDB.First(&syncStatus)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		syncStatus.LastSyncBlock = uint64(config.AppConfig.L2StartHeight)
		syncStatus.UpdatedAt = time.Now()
		relayerDB.Create(&syncStatus)
	} else if result.Error != nil {
		log.Fatalf("Error querying sync status: %v", result.Error)
	}

L:
	for {
		isContinue, err := l.scan(ctx, &syncStatus)
		if err != nil {
			log.Errorf("scan : %v", err)
		}
		if isContinue {
			continue L
		}
		// Next loop
		time.Sleep(config.AppConfig.L2RequestInterval)
	}
}

func (l *Layer2Listener) scan(ctx context.Context, syncStatus *db.EVMSyncStatus) (isContinue bool, err error) {
	latestBlock, err := l.ethClient.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("error getting latest block number: %v", err)
	}

	targetBlock := latestBlock - uint64(config.AppConfig.L2Confirmations)
	if syncStatus.LastSyncBlock < targetBlock {
		fromBlock := syncStatus.LastSyncBlock + 1

		toBlock := fromBlock + uint64(config.AppConfig.L2MaxBlockRange) - 1
		if toBlock > targetBlock {
			toBlock = targetBlock
		}

		log.WithFields(log.Fields{"fromBlock": fromBlock, "toBlock": toBlock}).Info("Syncing L2 nudex events")

		filterQuery := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: l.contractAddress,
			Topics:    topics,
		}

		logs, err := l.ethClient.FilterLogs(context.Background(), filterQuery)
		if err != nil {
			return false, fmt.Errorf("failed to filter logs: %w", err)
		}

		for _, vLog := range logs {
			l.processLogs(vLog)
		}
		// Save sync status
		syncStatus.LastSyncBlock = toBlock
		syncStatus.UpdatedAt = time.Now()
		l.db.GetRelayerDB().Save(syncStatus)

		return true, nil
	}

	return false, nil
}

// stop ctx.
//
//lint:ignore U1000 Ignore unused function
func (l *Layer2Listener) stop() {}
