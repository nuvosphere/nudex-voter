package btc

import (
	"context"
	"sync"

	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/tss/suite"
	"github.com/nuvosphere/nudex-voter/internal/wallet"
	log "github.com/sirupsen/logrus"
)

type WalletClient struct {
	*wallet.BaseWallet
	ctx       context.Context
	cancel    context.CancelFunc
	event     eventbus.Bus
	state     *state.BtcWalletState
	tss       suite.TssService
	txContext sync.Map // taskID:TxContext
	// client *txClient todo
}

func NewWallet(
	event eventbus.Bus,
	tss suite.TssService,
	stateDB *state.ContractState,
	state *state.BtcWalletState,
	voterContract layer2.VoterContract,
) *WalletClient {
	ctx, cancel := context.WithCancel(context.Background())
	return &WalletClient{
		BaseWallet: wallet.NewBaseWallet(stateDB, voterContract),
		ctx:        ctx,
		cancel:     cancel,
		event:      event,
		state:      state,
		tss:        tss,
		txContext:  sync.Map{},
		// client:     client,
	}
}

func (w *WalletClient) Start(context.Context) {
	log.Info("btc wallet client is starting...")
	w.tss.RegisterTssClient(w)
	w.receiveL2TaskLoop()
}

func (w *WalletClient) Stop(context.Context) {
	log.Info("btc wallet client is stopping...")
	w.cancel()
}
