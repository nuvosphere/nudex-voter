package p2p

import (
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	mr "math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	tcp "github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"

	"github.com/nuvosphere/nudex-voter/internal/config"
)

const (
	handshakeProtocol  = "/nudex/voter/handshake/1.0.0"
	messageProtocol    = "/nudex/voter/message/1.0.0"
	expectedHandshake  = "nudexvoternudexbest"
	messageTopicName   = "gossip-topic"
	heartbeatTopicName = "heartbeat-topic"
	privKeyFile        = "node_private_key.pem"
)

var messageTopic *pubsub.Topic

func StartLibp2p() {
	// Start libp2p node
	ctx := context.Background()
	node, ps, err := createNodeWithPubSub(ctx)
	if err != nil {
		log.Fatalf("Failed to create libp2p node: %v", err)
	}

	// Print self boot node info
	printNodeAddrInfo(node)

	// Set handshake
	node.SetStreamHandler(protocol.ID(handshakeProtocol), func(s network.Stream) {
		log.Println("New handshake stream")
		handleHandshake(s)
		s.Close()
	})

	bootNodeAddrs := strings.Split(config.AppConfig.Libp2pBootNodes, ",")
	// Connect to bootnodes and handshake
	for _, addr := range bootNodeAddrs {
		if addr == "" {
			continue
		}
		connectToBootNode(ctx, node, addr)
	}

	messageTopic, err = ps.Join(messageTopicName)
	if err != nil {
		log.Fatalf("Failed to join message topic: %v", err)
	}

	sub, err := messageTopic.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to message topic: %v", err)
	}


	hbTopic, err := ps.Join(heartbeatTopicName)
	if err != nil {
		log.Fatalf("Failed to join heartbeat topic: %v", err)
	}

	hbSub, err := hbTopic.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to heartbeat topic: %v", err)
	}
	
	go handlePubSubMessages(ctx, sub, node)
	go handleHeartbeatMessages(ctx, hbSub, node)
	go startHeartbeat(ctx, node, hbTopic)

	go func() {
		time.Sleep(5 * time.Second)
		msg := Message{
			ID:      generateMessageID(),
			Content: "Hello, nudex voter libp2p PubSub network with handshake!",
		}
		publishMessage(ctx, msg)
	}()
}

func createNodeWithPubSub(ctx context.Context) (host.Host, *pubsub.PubSub, error) {
	privKey, err := loadOrCreatePrivateKey(privKeyFile)
	if err != nil {
		return nil, nil, err
	}

	listenAddr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", config.AppConfig.Libp2pPort)
	node, err := libp2p.New(
		libp2p.Identity(privKey),
		libp2p.Transport(tcp.NewTCPTransport), //TCP only
		libp2p.ListenAddrStrings(listenAddr),  // ipv4 only
	)
	if err != nil {
		return nil, nil, err
	}

	ps, err := pubsub.NewGossipSub(ctx, node)
	if err != nil {
		return nil, nil, err
	}

	return node, ps, nil
}

func generateMessageID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), mr.Int63())
}

func connectToBootNode(ctx context.Context, node host.Host, bootNodeAddr string) {
	multiAddr, err := multiaddr.NewMultiaddr(bootNodeAddr)
	if err != nil {
		log.Printf("Failed to parse bootnode address: %v", err)
		return
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(multiAddr)
	if err != nil {
		log.Printf("Failed to get peer info from address: %v", err)
		return
	}

	node.Peerstore().AddAddrs(peerInfo.ID, peerInfo.Addrs, peerstore.PermanentAddrTTL)
	if err := node.Connect(ctx, *peerInfo); err != nil {
		log.Errorf("Failed to connect to bootnode: %v", err)
	} else {
		log.Infof("Connected to bootnode: %s", peerInfo.ID.String())

		// Handshake after connect
		s, err := node.NewStream(ctx, peerInfo.ID, protocol.ID(handshakeProtocol))
		if err != nil {
			log.Errorf("Failed to create handshake stream to peer %s: %v", peerInfo.ID, err)
			return
		}

		_, err = s.Write([]byte(expectedHandshake))
		if err != nil {
			log.Errorf("Failed to send handshake to peer %s: %v", peerInfo.ID, err)
			s.Reset()
			return
		}

		s.Close()
	}
}

func loadOrCreatePrivateKey(fileName string) (crypto.PrivKey, error) {
	dbDir := config.AppConfig.DbDir
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	pemPath := filepath.Join(dbDir, fileName)
	if _, err := os.Stat(pemPath); err == nil {
		privKeyBytes, err := ioutil.ReadFile(pemPath)
		if err != nil {
			return nil, err
		}
		privKey, err := crypto.UnmarshalPrivateKey(privKeyBytes)
		if err != nil {
			return nil, err
		}
		return privKey, nil
	}

	privKey, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 2048, rand.Reader)
	if err != nil {
		return nil, err
	}

	privKeyBytes, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(pemPath, privKeyBytes, 0600); err != nil {
		return nil, err
	}

	return privKey, nil
}

func printNodeAddrInfo(node host.Host) {
	addrs := node.Addrs()
	peerID := node.ID().String()

	for _, addr := range addrs {
		fullAddr := fmt.Sprintf("%s/p2p/%s", addr, peerID)
		log.Infof("Bootnode address: %s", fullAddr)
	}
}
