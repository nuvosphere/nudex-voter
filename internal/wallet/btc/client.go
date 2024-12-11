package btc

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
)

const precision = 1e8

type RawSignTx struct{}

type UnSignTx struct{}

type SignedTx struct{}

type Client struct {
	ctx       context.Context
	cancel    context.CancelFunc
	publicKey crypto.PublicKey // SerializeCompressed p2wpkh address
	client    *rpcclient.Client
	params    *chaincfg.Params
	tx        *wire.MsgTx
	signer    types.Signer
	signChan  chan []byte
}

func NewBtcClient(
	ctx context.Context,
	timeout time.Duration,
	singer types.Signer,
	params *chaincfg.Params,
	publicKey crypto.PublicKey,
) types.Requester {
	connConfig := &rpcclient.ConnConfig{
		Host:         config.AppConfig.BtcRpc,
		User:         config.AppConfig.BtcRpcUser,
		Pass:         config.AppConfig.BtcRpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connConfig, nil)
	utils.Assert(err)

	ctx, cancel := context.WithTimeout(ctx, timeout)

	return &Client{
		publicKey: publicKey,
		client:    client,
		tx:        wire.NewMsgTx(wire.TxVersion),
		ctx:       ctx,
		cancel:    cancel,
		signChan:  make(chan []byte, 1),
		signer:    singer,
		params:    params,
	}
}

// buildTxOut TxOut https://www.mengbin.top/2024-07-24-btcd_raw_tx/
func (c *Client) buildTxOut(addr string, amount int64) error {
	destinationAddress, err := btcutil.DecodeAddress(addr, c.params)
	if err != nil {
		return err
	}

	pkScript, err := txscript.PayToAddrScript(destinationAddress)
	if err != nil {
		return err
	}

	// satoshis
	c.tx.AddTxOut(wire.NewTxOut(amount, pkScript))
	return nil
}

func (c *Client) getUTXOs(from btcutil.Address) ([]btcjson.ListUnspentResult, error) {
	// todo
	return c.client.ListUnspentMinMaxAddresses(0, 9999999, []btcutil.Address{from})
}

func (c *Client) buildTxIn(amount int64) error {
	// p2wpkh address
	fromAddr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(c.publicKey.SerializeCompressed()), c.params)
	if err != nil {
		return err
	}

	utxos, err := c.getUTXOs(fromAddr)
	if err != nil {
		return err
	}

	// satoshis
	totalInput := int64(0)
	for _, utxo := range utxos {
		if totalInput > amount {
			break
		}
		txHash, err := chainhash.NewHashFromStr(utxo.TxID)
		if err != nil {
			return err
		}

		txIn := wire.NewTxIn(&wire.OutPoint{Hash: *txHash, Index: utxo.Vout}, nil, nil)
		c.tx.AddTxIn(txIn)
		totalInput += int64(utxo.Amount * precision)
	}

	// tx fee
	fee := int64(c.tx.SerializeSize())
	change := totalInput - amount
	if change > fee {
		changePkScript, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return err
		}
		txOut := wire.NewTxOut(change-fee, changePkScript)
		c.tx.AddTxOut(txOut)
	}

	// group := errgroup.WithContext(c.ctx)
	// group.Go(func() error {})
	for i, txIn := range c.tx.TxIn {
		prevOutputScript, err := hex.DecodeString(utxos[i].ScriptPubKey)
		if err != nil {
			return err
		}
		txHash, err := chainhash.NewHashFromStr(utxos[i].TxID)
		if err != nil {
			return err
		}
		outPoint := wire.OutPoint{Hash: *txHash, Index: utxos[i].Vout}
		prevOutputFetcher := txscript.NewMultiPrevOutFetcher(map[wire.OutPoint]*wire.TxOut{
			outPoint: {Value: int64(utxos[i].Amount * precision), PkScript: prevOutputScript},
		})
		sigHashes := txscript.NewTxSigHashes(c.tx, prevOutputFetcher)
		sigScript, err := c.witnessSignature(
			sigHashes,
			int(utxos[i].Vout),
			int64(utxos[i].Amount*precision), // todo
			prevOutputScript,
		)
		if err != nil {
			return err
		}
		txIn.Witness = sigScript
	}
	return nil
}

// rawTxInWitnessSignature returns the serialized ECDA signature for the input
// idx of the given transaction, with the hashType appended to it. This
// function is identical to RawTxInSignature, however the signature generated
// signs a new sighash digest defined in BIP0143.
func (c *Client) rawTxInWitnessSignature(sigHashes *txscript.TxSigHashes, idx int, amt int64, subScript []byte) ([]byte, error) {
	hash, err := txscript.CalcWitnessSigHash(subScript, sigHashes, txscript.SigHashAll, c.tx, idx, amt)
	if err != nil {
		return nil, err
	}
	signature, err := c.sign(hash)
	if err != nil {
		return nil, err
	}
	// txscript.SigHashAll, // https://www.btcstudy.org/2021/11/09/bitcoin-signature-types-sighash/
	return append(signature, byte(txscript.SigHashAll)), nil
}

// witnessSignature creates an input witness stack for tx to spend BTC sent
// from a previous output to the owner of privKey using the p2wkh script
// template. The passed transaction must contain all the inputs and outputs as
// dictated by the passed hashType. The signature generated observes the new
// transaction digest algorithm defined within BIP0143.
func (c *Client) witnessSignature(sigHashes *txscript.TxSigHashes, idx int, amt int64, subscript []byte) (wire.TxWitness, error) {
	sig, err := c.rawTxInWitnessSignature(sigHashes, idx, amt, subscript)
	if err != nil {
		return nil, err
	}

	// A witness script is actually a stack, so we return an array of byte
	// slices here, rather than a single byte slice.
	return wire.TxWitness{sig, c.publicKey.SerializeCompressed()}, nil
}

func (c *Client) sign(hash []byte) ([]byte, error) {
	err := c.signer.Sign(c, hash)
	if err != nil {
		return nil, err
	}
	select {
	case signature := <-c.signChan:
		return signature, nil
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	}
}

func (c *Client) sendTx(allowHighFees bool) (*chainhash.Hash, error) {
	hash, err := c.client.SendRawTransaction(c.tx, allowHighFees)
	if err != nil {
		return nil, fmt.Errorf("send raw transaction: %w", err)
	}
	return hash, nil
}

func (c *Client) SendTransaction(to string, amount int64, allowHighFees bool) ([]byte, error) {
	err := c.buildTxOut(to, amount)
	if err != nil {
		return nil, err
	}
	err = c.buildTxIn(amount)
	if err != nil {
		return nil, err
	}

	hash, err := c.sendTx(allowHighFees)
	if err != nil {
		return nil, err
	}
	return hash.CloneBytes(), nil
}

func (c *Client) WaitTxSuccess(hash []byte) error {
	h, err := chainhash.NewHash(hash)
	if err != nil {
		return err
	}

	begin := time.Now()
	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		res, err := c.client.GetRawTransactionVerbose(h)
		if err != nil {
			return fmt.Errorf("get raw transaction: %w", err)
		}
		if res == nil {
			count--
			time.Sleep(time.Second)
		} else {
			return nil
		}
	}
	return fmt.Errorf("get raw transaction fail")
}

func (c *Client) Close() error {
	c.cancel()
	return nil
}

func (c *Client) Post(hash, signature []byte) {
	c.signChan <- hash
}

func (c *Client) ChainType() int {
	return types.ChainBitcoin
}
