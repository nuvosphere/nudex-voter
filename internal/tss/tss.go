package tss

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"time"
)

func NewTssService(libp2p *p2p.LibP2PService, state *state.State) *TSSService {
	return &TSSService{
		libp2p: libp2p,
		state:  state,
	}
}

func (tss *TSSService) Start(ctx context.Context) {
	keygenPrepareMessage := KeyGenPrepareMessage{
		PublicKeys: PublicKeysToHex(config.AppConfig.TssPublicKeys),
		Threshold:  config.AppConfig.TssThreshold,
		Timestamp:  time.Now().Unix(),
	}

	requestId := fmt.Sprintf("KEYGEN:%s", crypto.PubkeyToAddress(config.AppConfig.L2PrivateKey.PublicKey).Hex())
	p2pMsg := p2p.Message{
		MessageType: p2p.MessageTypeKeygen,
		RequestId:   requestId,
		DataType:    p2p.DataTypeKeygenPrepare,
		Data:        keygenPrepareMessage,
	}

	tss.libp2p.PublishMessage(ctx, p2pMsg)
}
