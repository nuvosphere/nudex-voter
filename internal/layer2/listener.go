package layer2

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/abis"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"
)

var (
	votingManagerABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newParticipant","type":"address"}],"name":"ParticipantAdded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"participant","type":"address"}],"name":"ParticipantRemoved","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"requester","type":"address"},{"indexed":true,"internalType":"address","name":"currentSubmitter","type":"address"}],"name":"SubmitterRotationRequested","type":"event"},{"inputs":[{"internalType":"address","name":"targetAddress","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"txInfo","type":"bytes"},{"internalType":"uint256","name":"chainId","type":"uint256"},{"internalType":"bytes","name":"extraInfo","type":"bytes"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"submitDepositInfo","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
)

type Layer2Listener struct {
	libp2p    *p2p.LibP2PService
	db        *db.DatabaseManager
	state     *state.State
	ethClient *ethclient.Client

	contractVotingManager      *abis.VotingManagerContract
	contractAccountManager     *abis.AccountManagerContract
	contractOperations         *abis.NuDexOperationsContract
	contractParticipantManager *abis.ParticipantManagerContract
	contractDepositManager     *abis.DepositManagerContract

	sigFinishChan chan interface{}
}

func NewLayer2Listener(libp2p *p2p.LibP2PService, state *state.State, db *db.DatabaseManager) *Layer2Listener {
	ethClient, err := DialEthClient()
	if err != nil {
		log.Fatalf("Error creating Layer2 EVM RPC client: %v", err)
	}

	contractVotingManager, err := abis.NewVotingManagerContract(abis.VotingAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract VotingManager: %v", err)
	}

	contractAccountManager, err := abis.NewAccountManagerContract(abis.AccountAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract AccountManager: %v", err)
	}

	contractOperations, err := abis.NewNuDexOperationsContract(abis.OperationsAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract Operations: %v", err)
	}

	contractParticipantManager, err := abis.NewParticipantManagerContract(abis.ParticipantAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract ParticipantManager: %v", err)
	}

	contractDepositManager, err := abis.NewDepositManagerContract(abis.DepositAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate contract ParticipantManager: %v", err)
	}

	return &Layer2Listener{
		libp2p:    libp2p,
		db:        db,
		state:     state,
		ethClient: ethClient,

		contractVotingManager:      contractVotingManager,
		contractAccountManager:     contractAccountManager,
		contractOperations:         contractOperations,
		contractParticipantManager: contractParticipantManager,
		contractDepositManager:     contractDepositManager,

		sigFinishChan: make(chan interface{}, 256),
	}
}

// New an eth client
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

func (lis *Layer2Listener) Start(ctx context.Context) {
	// Get latest sync height
	var syncStatus db.EVMSyncStatus
	relayerDB := lis.db.GetRelayerDB()
	result := relayerDB.First(&syncStatus)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		syncStatus.LastSyncBlock = uint64(config.AppConfig.L2StartHeight)
		syncStatus.UpdatedAt = time.Now()
		relayerDB.Create(&syncStatus)
	} else if result.Error != nil {
		log.Fatalf("Error querying sync status: %v", result.Error)
	}

	for {
		if err := lis.scan(ctx, &syncStatus); err != nil {
			log.Errorf("scan : %v", err)
		}
		// Next loop
		time.Sleep(config.AppConfig.L2RequestInterval)
	}

}

func (lis *Layer2Listener) scan(ctx context.Context, syncStatus *db.EVMSyncStatus) error {
	latestBlock, err := lis.ethClient.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("error getting latest block number: %v", err)
	}
	targetBlock := latestBlock - uint64(config.AppConfig.L2Confirmations)
	if syncStatus.LastSyncBlock < targetBlock {
		fromBlock := syncStatus.LastSyncBlock + 1

		toBlock := fromBlock + uint64(config.AppConfig.L2MaxBlockRange) - 1
		if toBlock > targetBlock {
			toBlock = targetBlock
		}

		if toBlock <= targetBlock {
			log.WithFields(log.Fields{
				"fromBlock": fromBlock,
				"toBlock":   toBlock,
			}).Info("Syncing L2 nudex events")

			filterQuery := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(toBlock)),
				Addresses: []common.Address{
					abis.VotingAddress,
					abis.AccountAddress,
					abis.ParticipantAddress,
					abis.OperationsAddress,
				},
				Topics: [][]common.Hash{
					{SubmitterChosenTopic},
					{TaskSubmittedTopic},
					{AddressRegisteredTopic},
					{ParticipantRemovedTopic},
					{ParticipantAddedTopic},
				},
			}

			logs, err := lis.ethClient.FilterLogs(context.Background(), filterQuery)
			if err != nil {
				return fmt.Errorf("failed to filter logs: %w", err)
			}

			for _, vLog := range logs {
				lis.processLogs(vLog)
			}
			// Save sync status
			syncStatus.LastSyncBlock = toBlock
			syncStatus.UpdatedAt = time.Now()
			lis.db.GetRelayerDB().Save(syncStatus)

			err = lis.scan(ctx, syncStatus)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// stop ctx
func (lis *Layer2Listener) stop() {

}
