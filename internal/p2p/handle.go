package p2p

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	log "github.com/sirupsen/logrus"
)

func handleHandshake(s network.Stream, node host.Host) {
	buf := make([]byte, 1024)
	n, err := s.Read(buf)
	if err != nil {
		log.Errorf("Error reading handshake message: %v", err)
		return
	}

	handshakeMsg := buf[:n]
	log.Infof("Received handshake message: %s", string(handshakeMsg))

	expectedMsg := []byte(expectedHandshake)
	if !bytes.Equal(handshakeMsg, expectedMsg) {
		log.Warn("Invalid handshake message received, closing connection")
		s.Reset()

		// disconnect peer
		peerID := s.Conn().RemotePeer()
		// s.Conn().Close()
		node.Network().ClosePeer(peerID)
		return
	}

	_, err = s.Write(handshakeMsg)
	if err != nil {
		log.Errorf("Error writing handshake response: %v", err)
		return
	}

	log.Info("Handshake successful")
}
func (libp2p *LibP2PService) PublishMessage(ctx context.Context, msg Message) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("Failed to marshal message: %v", err)
		return
	}

	if libp2p.messageTopic == nil {
		startTime := time.Now()
		if time.Since(startTime) >= 10*time.Second {
			log.Error("Message topic is nil, cannot publish message")
			return
		}
		if libp2p.messageTopic == nil {
			time.Sleep(1 * time.Second)
		}
	}

	if err := libp2p.messageTopic.Publish(ctx, msgBytes); err != nil {
		log.Errorf("Failed to publish message: %v", err)
	}
}

func (libp2p *LibP2PService) handlePubSubMessages(ctx context.Context, sub *pubsub.Subscription, node host.Host) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			log.Errorf("Error reading message from pubsub: %v", err)
			continue
		}

		if msg.ReceivedFrom == node.ID() {
			log.Debug("Received message from self, ignore")
			continue
		}

		var receivedMsg Message
		if err := json.Unmarshal(msg.Data, &receivedMsg); err != nil {
			log.Errorf("Error unmarshaling pubsub message: %v", err)
			continue
		}

		log.Debugf("Received message via pubsub: ID=%d, RequestId=%s, Data=%v", receivedMsg.MessageType, receivedMsg.RequestId, receivedMsg.Data)

		switch receivedMsg.MessageType {
		case MessageTypeSigReq:
			libp2p.state.EventBus.Publish(state.SigReceive, convertMsgData(receivedMsg))
		case MessageTypeSigResp:
			libp2p.state.EventBus.Publish(state.SigReceive, convertMsgData(receivedMsg))
		case MessageTypeDepositReceive:
			libp2p.state.EventBus.Publish(state.DepositReceive, convertMsgData(receivedMsg))
		default:
			log.Warnf("Unknown message type: %d", receivedMsg.MessageType)
		}
	}
}

// convertMsgData converts the message data to the corresponding struct
// TODO: use reflector to optimize this function
func convertMsgData(msg Message) interface{} {
	if msg.DataType == DataTypeKeygenPrepare {
		jsonBytes, _ := json.Marshal(msg.Data)
		var rawData types.KeyGenPrepareMessage
		_ = json.Unmarshal(jsonBytes, &rawData)
		return rawData
	}
	return msg.Data
}

func (libp2p *LibP2PService) handleHeartbeatMessages(ctx context.Context, sub *pubsub.Subscription, node host.Host) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			log.Errorf("Error reading heartbeat message from pubsub: %v", err)
			continue
		}

		if msg.ReceivedFrom == node.ID() {
			log.Debug("Received heartbeat from self, ignore")
			continue
		}

		var hbMsg HeartbeatMessage
		if err := json.Unmarshal(msg.Data, &hbMsg); err != nil {
			log.Errorf("Error unmarshaling heartbeat message: %v", err)
			continue
		}

		log.Infof("Received heartbeat from %d-%s: %s", hbMsg.Timestamp, hbMsg.PeerID, hbMsg.Message)
	}
}

func startHeartbeat(ctx context.Context, node host.Host, topic *pubsub.Topic) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hbMsg := HeartbeatMessage{
				PeerID:    node.ID().String(),
				Message:   "heartbeat",
				Timestamp: time.Now().Unix(),
			}

			msgBytes, err := json.Marshal(hbMsg)
			if err != nil {
				log.Errorf("Failed to marshal heartbeat message: %v", err)
				continue
			}

			if err := topic.Publish(ctx, msgBytes); err != nil {
				log.Errorf("Failed to publish heartbeat message: %v", err)
			} else {
				log.Infof("Heartbeat message sent by %s", hbMsg.PeerID)
			}

		case <-ctx.Done():
			return
		}
	}
}
