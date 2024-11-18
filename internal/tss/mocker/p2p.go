package mocker

import (
	"context"

	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
)

type P2PMocker struct{}

func (p P2PMocker) Bind(msgType p2p.MessageType, event eventbus.Event) {
	// TODO implement me
	panic("implement me")
}

func (p P2PMocker) PublishMessage(ctx context.Context, msg any) error {
	// TODO implement me
	panic("implement me")
}

func (p P2PMocker) OnlinePeerCount() int {
	// TODO implement me
	panic("implement me")
}

func (p P2PMocker) IsOnline(partyID string) bool {
	// TODO implement me
	panic("implement me")
}
