package p2p

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMsgDecode(t *testing.T) {
	originMsg := Message[any]{
		MessageType: 10,
		RequestId:   "test",
		DataType:    "test",
		Data: HeartbeatMessage{
			PeerID:    "100",
			Message:   "test",
			Timestamp: time.Now().Unix(),
		},
	}

	data, err := json.Marshal(originMsg)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)

	rawMsg := Message[json.RawMessage]{}
	err = json.Unmarshal(data, &rawMsg)
	assert.Nil(t, err)

	heartMsg := HeartbeatMessage{}
	err = json.Unmarshal(rawMsg.Data, &heartMsg)
	assert.Nil(t, err)

	msg := Message[any]{
		MessageType: rawMsg.MessageType,
		RequestId:   rawMsg.RequestId,
		DataType:    rawMsg.DataType,
		Data:        heartMsg,
	}
	assert.Equal(t, originMsg, msg)
}
