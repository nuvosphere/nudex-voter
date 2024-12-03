package main

import (
	"encoding/hex"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
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
