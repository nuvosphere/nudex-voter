package btc

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type RawSignTx struct{}

type UnSignTx struct{}

type SignedTx struct{}

type BtcClient struct {
	privateKey *btcec.PrivateKey
	publicKey  *btcec.PublicKey
	client     *rpcclient.Client
	params     *chaincfg.Params
}

func NewBtcClient() *BtcClient {
	connConfig := &rpcclient.ConnConfig{
		Host:         config.AppConfig.BtcRpc,
		User:         config.AppConfig.BtcRpcUser,
		Pass:         config.AppConfig.BtcRpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connConfig, nil)
	utils.Assert(err)

	privateKeyBytes, err := hex.DecodeString(config.AppConfig.BtcPrivKey)
	utils.Assert(err)

	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	return &BtcClient{
		privateKey: privateKey,
		publicKey:  publicKey,
		client:     client,
	}
}

// BuildTxOut TxOut https://www.mengbin.top/2024-07-24-btcd_raw_tx/
func (c *BtcClient) BuildTxOut(addr string, amount int64, params *chaincfg.Params) (*wire.TxOut, []byte, error) {
	destinationAddress, err := btcutil.DecodeAddress(addr, params)
	if err != nil {
		return nil, nil, err
	}

	// todo
	pkScript, err := txscript.PayToAddrScript(destinationAddress)
	if err != nil {
		return nil, nil, err
	}

	// satoshis
	return wire.NewTxOut(amount, pkScript), pkScript, nil
}

func (c *BtcClient) GetUTXOs(address btcutil.Address) ([]btcjson.ListUnspentResult, error) {
	// todo
	return c.client.ListUnspentMinMaxAddresses(0, 9999999, []btcutil.Address{address})
}

func (c *BtcClient) BuildTxIn(amount int64, txOut *wire.TxOut) (*wire.MsgTx, error) {
	fromAddr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(c.publicKey.SerializeCompressed()), c.params)
	if err != nil {
		return nil, err
	}

	utxos, err := c.GetUTXOs(fromAddr)
	if err != nil {
		return nil, err
	}

	msgTx := wire.NewMsgTx(wire.TxVersion)
	// satoshis
	totalInput := int64(0)
	for _, utxo := range utxos {
		if totalInput > amount {
			break
		}
		txHash, err := chainhash.NewHashFromStr(utxo.TxID)
		if err != nil {
			return nil, err
		}

		txIn := wire.NewTxIn(&wire.OutPoint{Hash: *txHash, Index: utxo.Vout}, nil, nil)
		msgTx.AddTxIn(txIn)
		totalInput += int64(utxo.Amount * 1e8)
	}
	msgTx.AddTxOut(txOut)

	fee := int64(msgTx.SerializeSize())
	change := totalInput - amount
	if change > fee {
		changePkScript, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return nil, err
		}
		txOut := wire.NewTxOut(change-fee, changePkScript)
		msgTx.AddTxOut(txOut)
	}

	for i, txIn := range msgTx.TxIn {
		prevOutputScript, err := hex.DecodeString(utxos[i].ScriptPubKey)
		if err != nil {
			return nil, err
		}
		txHash, err := chainhash.NewHashFromStr(utxos[i].TxID)
		if err != nil {
			return nil, err
		}
		outPoint := wire.OutPoint{Hash: *txHash, Index: utxos[i].Vout}
		prevOutputFetcher := txscript.NewMultiPrevOutFetcher(map[wire.OutPoint]*wire.TxOut{
			outPoint: {Value: int64(utxos[i].Amount * 1e8), PkScript: prevOutputScript},
		})
		sigHashes := txscript.NewTxSigHashes(msgTx, prevOutputFetcher)
		// todo
		sigScript, err := txscript.WitnessSignature(
			msgTx,
			sigHashes,
			int(utxos[i].Vout),
			int64(utxos[i].Amount*1e8),
			prevOutputScript,
			txscript.SigHashAll,
			c.privateKey,
			true,
		)
		if err != nil {
			return nil, err
		}
		txIn.Witness = sigScript
	}
	return msgTx, nil
}
