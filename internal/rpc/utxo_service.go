package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"net"

	"github.com/btcsuite/btcd/wire"
	bitcointypes "github.com/goatnetwork/goat/x/bitcoin/types"
	relayertypes "github.com/goatnetwork/goat/x/relayer/types"
	"github.com/nuvosphere/nudex-voter/internal/btc"
	"github.com/nuvosphere/nudex-voter/internal/config"
	pb "github.com/nuvosphere/nudex-voter/internal/proto"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

type UtxoServer struct {
	pb.UnimplementedBitcoinLightWalletServer
}

func NewUtxoServer() *UtxoServer {
	return &UtxoServer{}
}

func (s *UtxoServer) NewTransaction(ctx context.Context, req *pb.NewTransactionRequest) (*pb.NewTransactionResponse, error) {
	var tx wire.MsgTx
	if err := json.NewDecoder(bytes.NewReader(req.RawTransaction)).Decode(&tx); err != nil {
		log.Errorf("Failed to decode transaction: %v", err)
		return nil, err
	}

	_, err := btc.GenerateSPVProof(&tx)
	if err != nil {
		log.Errorf("Failed to generate SPV proof: %v", err)
		return nil, err
	}

	deposit := &bitcointypes.Deposit{
		EvmAddress:    req.EvmAddress,
		NoWitnessTx:   req.RawTransaction,
		RelayerPubkey: &relayertypes.PublicKey{}, // Need to fill in the correct RelayerPubkey
		OutputIndex:   uint32(0),                 // Assuming output index is 0, adjust according to actual situation
		Version:       uint32(0),                 // Assuming version is 0, adjust according to actual situation
	}

	btc.SendDepositData(deposit)

	return &pb.NewTransactionResponse{
		TransactionId: "txhash",
		ErrorMessage:  "",
	}, nil
}

func StartUTXOService() {
	addr := ":" + config.AppConfig.HTTPPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBitcoinLightWalletServer(s, &UtxoServer{})
	reflection.Register(s)

	log.Infof("gRPC server is running on port %s", config.AppConfig.HTTPPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
