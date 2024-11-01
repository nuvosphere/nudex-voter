package state

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	t.Log("test eventbus begin")

	testLen := 1000
	exist := make(chan struct{}, testLen)
	wg := sync.WaitGroup{}
	count := atomic.Uint64{}
	for i := 0; i < testLen; i++ {
		tssUpdated := make(chan interface{})
		bus.Subscribe(TssMsg, tssUpdated)
		wg.Add(1)
		go func() {
			exist <- struct{}{}
			result := <-tssUpdated
			t.Logf("subtest:index = %d, result = %v", i, result)
			count.Add(1)

			wg.Done()
		}()
	}
	<-exist
	bus.Publish(TssMsg, "OK")
	t.Log("eventbus publish end")
	wg.Wait()
	assert.Equal(t, count.Load(), uint64(len(bus.subscribers[TssMsg.String()])))
	t.Log("test eventbus end")
}
