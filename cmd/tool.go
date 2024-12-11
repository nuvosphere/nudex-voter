package main

import (
	"encoding/hex"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "generate config info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate config info")
	},
}

var printPublicKeyCmd = &cobra.Command{
	Use:     "publicKey",
	Short:   "generate public key from secp256k1 private key",
	Example: `nudex-voter tool publicKey 76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := hex.DecodeString(args[0])
		utils.Assert(err)
		pk := ethCrypto.ToECDSAUnsafe(data)
		publicKey := ethCrypto.CompressPubkey(&pk.PublicKey)
		fmt.Println(hex.EncodeToString(publicKey[:]))
	},
}

var printPeerIDCmd = &cobra.Command{
	Use:     "peerID",
	Short:   "generate p2p peerID from secp256k1 public key",
	Example: `nudex-voter tool peerID 020b537f46c6da81f84824ce1409bab1f9825fb58b57dcafbf4f4b074e90a0c040`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := hex.DecodeString(args[0])
		utils.Assert(err)
		pubkey, err := crypto.UnmarshalSecp256k1PublicKey(data)
		utils.Assert(err)
		id, err := peer.IDFromPublicKey(pubkey)
		utils.Assert(err)
		fmt.Println(id)
	},
}

var printAddressCmd = &cobra.Command{
	Use:     "address",
	Short:   "generate eth address from secp256k1 public key",
	Example: `nudex-voter tool address 020b537f46c6da81f84824ce1409bab1f9825fb58b57dcafbf4f4b074e90a0c040`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := hex.DecodeString(args[0])
		utils.Assert(err)
		pubkey, err := ethCrypto.DecompressPubkey(data)
		utils.Assert(err)
		address := ethCrypto.PubkeyToAddress(*pubkey)
		fmt.Println(address)
	},
}

var printTssAddressCmd = &cobra.Command{
	Use:     "tssAddress",
	Short:   "print master tss eth address from config data",
	Example: `nudex-voter tool tssAddress`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(configPath)
		partyData := tss.NewPartyData(config.AppConfig.DbDir)
		fmt.Println(partyData.ECDSALocalData().TssSigner())
	},
}

var printEDDSAPublicKeyCmd = &cobra.Command{
	Use:     "eddsaPublicKey",
	Short:   "print EDDSA PublicKey from config data",
	Example: `nudex-voter tool public key`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(configPath)
		partyData := tss.NewPartyData(config.AppConfig.DbDir)
		fmt.Println(partyData.EDDSALocalData().CompressedPublicKey())
	},
}

var printECDSAPublicKeyCmd = &cobra.Command{
	Use:     "ecdsaPublicKey",
	Short:   "print ECDSA PublicKey from config data",
	Example: `nudex-voter tool public key`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(configPath)
		partyData := tss.NewPartyData(config.AppConfig.DbDir)
		fmt.Println(partyData.ECDSALocalData().CompressedPublicKey())
	},
}

var printP2pFullAddressCmd = &cobra.Command{
	Use:     "p2pFullAddr",
	Short:   "print p2p full address from config",
	Example: `nudex-voter tool p2pFullAddr`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(configPath)
		secp256k1PrivateKey, err := crypto.UnmarshalSecp256k1PrivateKey(ethCrypto.FromECDSA(config.L2PrivateKey))
		utils.Assert(err)
		listenAddr := p2p.ListenAddr()
		node, err := libp2p.New(
			libp2p.Identity(secp256k1PrivateKey),
			libp2p.Transport(tcp.NewTCPTransport), // TCP only
			libp2p.ListenAddrStrings(listenAddr),  // ipv4 only
		)
		utils.Assert(err)
		p2p.PrintNodeAddrInfo(node)
	},
}
