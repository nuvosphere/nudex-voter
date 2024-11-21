package layer2

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Layer2Listener struct {
	p2p                   *p2p.Service
	db                    *db.DatabaseManager
	state                 *state.State
	ethClient             *ethclient.Client
	chainID               atomic.Int64
	contractAddress       []common.Address
	addressBind           map[common.Address]func(types.Log) error
	contractVotingManager *contracts.VotingManagerContract
	participantManager    *contracts.ParticipantManagerContract
	taskManager           *contracts.TaskManagerContract
	accountManager        *contracts.AccountManagerContract
}

func (l *Layer2Listener) postTask(task any) {
	l.state.Bus().Publish(eventbus.EventTask{}, task)
}

func NewLayer2Listener(p *p2p.Service, state *state.State, db *db.DatabaseManager) *Layer2Listener {
	ethClient, err := DialEthClient()
	if err != nil {
		log.Fatalf("Error creating Layer2 EVM RPC client: %v", err)
	}

	self := &Layer2Listener{
		p2p:       p,
		db:        db,
		state:     state,
		ethClient: ethClient,
		chainID:   atomic.Int64{},
	}

	var (
		VotingAddress      = common.HexToAddress(config.AppConfig.VotingContract)
		AccountAddress     = common.HexToAddress(config.AppConfig.AccountContract)
		TaskAddress        = common.HexToAddress(config.AppConfig.TaskManagerContract)
		ParticipantAddress = common.HexToAddress(config.AppConfig.ParticipantContract)
		DepositAddress     = common.HexToAddress(config.AppConfig.DepositContract)
	)

	self.addressBind = map[common.Address]func(types.Log) error{
		VotingAddress:      self.processVotingLog,
		AccountAddress:     self.processAccountLog,
		ParticipantAddress: self.processParticipantLog,
		TaskAddress:        self.processTaskLog,
		DepositAddress:     self.processDepositLog,
	}
	self.contractAddress = lo.MapToSlice(
		self.addressBind,
		func(item common.Address, _ func(log2 types.Log) error) common.Address { return item },
	)

	var errs []error

	chainId, err := self.ChainID(context.Background())
	self.chainID.Store(chainId.Int64())
	if chainId.Int64() != config.AppConfig.L2ChainId.Int64() {
		err = fmt.Errorf("ChainId mismatch: expected %d, got %d", config.AppConfig.L2ChainId.Int64(), chainId.Int64())
		errs = append(errs, err)
	}
	errs = append(errs, err)
	contractVotingManager, err := contracts.NewVotingManagerContract(VotingAddress, ethClient)
	errs = append(errs, err)
	participantManager, err := contracts.NewParticipantManagerContract(ParticipantAddress, ethClient)
	errs = append(errs, err)
	taskManager, err := contracts.NewTaskManagerContract(TaskAddress, ethClient)
	errs = append(errs, err)
	accountManager, err := contracts.NewAccountManagerContract(TaskAddress, ethClient)
	errs = append(errs, err)

	utils.Assert(errors.Join(errs...))

	self.taskManager = taskManager
	self.contractVotingManager = contractVotingManager
	self.participantManager = participantManager
	self.accountManager = accountManager

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
			// Topics:    batch,
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

func (l *Layer2Listener) ChainID(ctx context.Context) (*big.Int, error) {
	if l.chainID.Load() == 0 {
		chainID, err := l.ethClient.ChainID(ctx)
		if err != nil {
			return nil, fmt.Errorf("chainID error: %w", err)
		}

		l.chainID.Store(chainID.Int64())
	}

	return big.NewInt(l.chainID.Load()), nil
}
