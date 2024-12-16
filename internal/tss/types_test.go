package tss

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestJsonParse(t *testing.T) {
	msg := p2p.Message[types.TssMessage]{
		MessageType: 12,
		RequestId:   "test",
		DataType:    "test",
		Data: types.TssMessage{
			FromPartyId:  "rewr",
			ToPartyIds:   []string{"1", "2"},
			IsBroadcast:  false,
			MsgWireBytes: []byte("1234567890"),
		},
	}

	data, err := json.Marshal(msg)
	assert.Nil(t, err)
	t.Log(string(data))

	msgByte := p2p.Message[json.RawMessage]{}
	err = json.Unmarshal(data, &msgByte)
	assert.Nil(t, err)
	t.Log(msg)
	t.Log(msg.Data)
	t.Log(reflect.ValueOf(msg.Data).String())

	dataMsg := types.TssMessage{}
	err = json.Unmarshal(msgByte.Data, &dataMsg)
	assert.Nil(t, err)
	t.Log(utils.FormatJSON(dataMsg))
}

func TestConvertMsgData(t *testing.T) {
	originMsg := p2p.Message[types.TssMessage]{
		MessageType: 10,
		RequestId:   "test",
		DataType:    DataTypeTssKeygenMsg,
		Data: types.TssMessage{
			FromPartyId:  "test",
			ToPartyIds:   []string{"1", "2"},
			IsBroadcast:  false,
			MsgWireBytes: []byte("1234567890"),
		},
	}

	data, err := json.Marshal(originMsg)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)

	rawMsg := p2p.Message[json.RawMessage]{}
	err = json.Unmarshal(data, &rawMsg)
	assert.Nil(t, err)

	event := ConvertP2PMsgData(rawMsg)
	orderMag, ok := event.(types.TssMessage)
	assert.True(t, ok)

	msg := p2p.Message[types.TssMessage]{
		MessageType: rawMsg.MessageType,
		RequestId:   rawMsg.RequestId,
		DataType:    rawMsg.DataType,
		Data:        orderMag,
	}
	t.Log("msg", msg)
	assert.Equal(t, originMsg, msg)
}
