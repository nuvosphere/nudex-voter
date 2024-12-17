package solana

import (
	"context"
	"math/big"

	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

type WalletClient struct {
	*wallet.BaseWallet
	ctx    context.Context
	cancel context.CancelFunc
	event  eventbus.Bus
	state  *state.EvmWalletState
	tss    suite.TssService
	client *SolClient
}

func NewWallet(
	event eventbus.Bus,
	tss suite.TssService,
	stateDB *state.ContractState,
	state *state.EvmWalletState,
	voterContract layer2.VoterContract,
) *WalletClient {
	client := NewSolClient()
	ctx, cancel := context.WithCancel(context.Background())
	return &WalletClient{
		BaseWallet: wallet.NewBaseWallet(stateDB, voterContract),
		ctx:        ctx,
		cancel:     cancel,
		event:      event,
		state:      state,
		tss:        tss,
		client:     client,
	}
}

func (w *WalletClient) Start(context.Context) {
	log.Info("solana wallet client is stopping...")
	// w.loopApproveProposal()
}

func (w *WalletClient) Stop(context.Context) {
	w.cancel()
}

func (w *WalletClient) Verify(reqId *big.Int, signDigest string, ExtraData []byte) error {
	// TODO implement me
	panic("implement me")
}

func (w *WalletClient) PostSignature(res suite.SignRes) error {
	// TODO implement me
	panic("implement me")
}

func (w *WalletClient) ChainType() uint8 {
	return types.ChainSolana
}
