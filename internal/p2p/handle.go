package p2p

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	log "github.com/sirupsen/logrus"
)

func handleHandshake(s network.Stream) {
	buffer := make([]byte, len(expectedHandshake))
	_, err := s.Read(buffer)
	if err != nil {
		log.Errorf("Failed to read handshake data: %v", err)
		return
	}

	if string(buffer) == expectedHandshake {
		log.Info("Handshake successful")
	} else {
		log.Warn("Handshake failed")
		s.Reset()
	}
}

func publishMessage(ctx context.Context, msg Message) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("Failed to marshal message: %v", err)
		return
	}

	if messageTopic == nil {
		log.Error("Message topic is nil, cannot publish message")
		return
	}

	if err := messageTopic.Publish(ctx, msgBytes); err != nil {
		log.Errorf("Failed to publish message: %v", err)
	}
}

func handlePubSubMessages(ctx context.Context, sub *pubsub.Subscription, node host.Host) {
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

		log.Infof("Received message via pubsub: ID=%d, Content=%s", receivedMsg.MessageType, receivedMsg.Content)

		// Implement key processes for keygen and signing in tsslib/v2
		switch receivedMsg.MessageType {
		case MessageTypeKeygen:
			handleKeygenMessage(ctx)
		case MessageTypeSigning:
			handleSigningMessage(ctx)
		default:
			log.Warnf("Unknown message type: %d", receivedMsg.MessageType)
		}
	}
}

func handleKeygenMessage(ctx context.Context) {
	outCh := make(chan tss.Message, 3)
	endCh := make(chan *keygen.LocalPartySaveData, 1)

	party := setupTSSParty(keygen.NewLocalParty, outCh, endCh)
	startTSSProcess(ctx, party, outCh, endCh, MessageTypeKeygen, saveTSSData)
}

func handleSigningMessage(ctx context.Context) {
	outCh := make(chan tss.Message, 3)
	endCh := make(chan *common.SignatureData, 1)

	msgToSign := big.NewInt(123456)
	savedData, err := loadTSSData()
	if err != nil {
		log.Errorf("Failed to load TSS data: %v", err)
		return
	}

	party := setupTSSParty(func(params *tss.Parameters, outCh chan<- tss.Message, endCh chan<- *common.SignatureData) tss.Party {
		return signing.NewLocalParty(msgToSign, params, *savedData, outCh, endCh)
	}, outCh, endCh)

	startTSSProcess(ctx, party, outCh, endCh, MessageTypeSigning, handleSignature)
}

func setupTSSParty(createParty interface{}, outCh, endCh interface{}) tss.Party {
	parties := 3
	threshold := 2
	partyIDs := createPartyIDs(parties)
	peerCtx := tss.NewPeerContext(partyIDs)
	params := tss.NewParameters(tss.S256(), peerCtx, partyIDs[0], parties, threshold)

	return createParty.(func(*tss.Parameters, interface{}, interface{}) tss.Party)(params, outCh, endCh)
}

func startTSSProcess(ctx context.Context, party tss.Party, outCh <-chan tss.Message, endCh interface{}, msgType MessageType, endHandler interface{}) {
	go func() {
		if err := party.Start(); err != nil {
			log.Errorf("TSS process failed to start: %v", err)
		}
	}()

	handleTSSProcess(ctx, outCh, endCh, msgType, endHandler)
}

func handleTSSProcess(ctx context.Context, outCh <-chan tss.Message, endCh interface{}, msgType MessageType, endHandler interface{}) {
	for {
		select {
		case msg := <-outCh:
			publishTSSMessage(ctx, msg, msgType)
		case data := <-endCh.(chan interface{}):
			log.Infof("TSS process completed, data: %v", data)
			endHandler.(func(interface{}))(data)
		case <-ctx.Done():
			return
		}
	}
}

func createPartyIDs(parties int) tss.SortedPartyIDs {
	partyIDs := make(tss.SortedPartyIDs, parties)
	for i := 0; i < parties; i++ {
		partyIDs[i] = tss.NewPartyID(strconv.Itoa(i), "", new(big.Int).SetInt64(int64(i)))
	}
	return partyIDs
}

func handleHeartbeatMessages(ctx context.Context, sub *pubsub.Subscription, node host.Host) {
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

func handleSignature(signatureData *common.SignatureData) {
	// Implement logic to handle generated signature data
	log.Infof("Received signature data: %+v", signatureData)

	// TODO: Send signature data to other systems or perform further processing
	// For example:
	// - Verify signature
	// - Store signature in database
	// - Broadcast signature to other nodes
	// - Trigger subsequent business logic
}
