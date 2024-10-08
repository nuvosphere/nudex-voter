package tss

import (
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/state"
)

type TSSService struct {
	libp2p *p2p.LibP2PService
	state  *state.State
}

type KeyGenPrepareMessage struct {
	PublicKeys []string `json:"public_keys"`
	Threshold  int      `json:"threshold"`
	Timestamp  int64    `json:"ts"`
}

type KeygenMessage struct {
	Content string
}

type SigningMessage struct {
	Content string
}
