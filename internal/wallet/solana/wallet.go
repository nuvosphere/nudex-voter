package solana

import (
	"context"
	"math/big"
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

type WalletClient struct {
	ctx           context.Context
	cancel        context.CancelFunc
	event         eventbus.Bus
	state         *state.WalletEvmState
	tss           suite.TssService
	voterContract layer2.VoterContract
	client        *SolClient
	taskQueue     *pool.Pool[uint64]
	pendingTx     sync.Map // txHash: bool
}

func NewWallet(event eventbus.Bus, tss suite.TssService, voterContract layer2.VoterContract) *WalletClient {
	client := NewSolClient()
	ctx, cancel := context.WithCancel(context.Background())
	return &WalletClient{
		ctx:           ctx,
		cancel:        cancel,
		event:         event,
		client:        client,
		pendingTx:     sync.Map{},
		state:         nil,
		tss:           tss,
		voterContract: voterContract,
		taskQueue:     pool.NewTaskPool[uint64](),
	}
}

func (s *WalletClient) Start(context.Context) {
	log.Info("solana wallet client is stopping...")
	// s.loopApproveProposal()
}

func (s *WalletClient) Stop(context.Context) {
	s.cancel()
}

func (s *WalletClient) Verify(reqId *big.Int, signDigest string, ExtraData []byte) error {
	// TODO implement me
	panic("implement me")
}

func (s *WalletClient) PostSignature(res suite.SignRes) error {
	// TODO implement me
	panic("implement me")
}

func (s *WalletClient) ChainType() uint8 {
	return types.ChainSolana
}
