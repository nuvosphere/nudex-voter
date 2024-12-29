package btc

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"sync/atomic"
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

type SignatureCtx struct {
	Hash, Signature []byte
}

type txClient struct {
	ctx          context.Context
	cancel       context.CancelFunc
	publicKey    crypto.PublicKey // SerializeCompressed p2wpkh address
	client       *rpcclient.Client
	params       *chaincfg.Params
	tx           *wire.MsgTx
	txInHash     [][]byte
	hashCounter  atomic.Int64
	signChan     chan *SignatureCtx
	nextSignChan chan []byte
}

func NewTxClient(
	ctx context.Context,
	timeout time.Duration,
	params *chaincfg.Params,
	publicKey crypto.PublicKey,
) types.TxClient {
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

	return &txClient{
		publicKey:    publicKey,
		client:       client,
		tx:           wire.NewMsgTx(wire.TxVersion),
		ctx:          ctx,
		cancel:       cancel,
		signChan:     make(chan *SignatureCtx, 1),
		nextSignChan: make(chan []byte, 1),
		params:       params,
	}
}

// buildTxOut TxOut https://www.mengbin.top/2024-07-24-btcd_raw_tx/
func (c *txClient) buildTxOut(addr string, amount int64) error {
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

func (c *txClient) getUTXOs(from btcutil.Address) ([]btcjson.ListUnspentResult, error) {
	// todo
	return c.client.ListUnspentMinMaxAddresses(0, 50, []btcutil.Address{from})
}

func (c *txClient) buildTxIn(amount int64) error {
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

	for i := range c.tx.TxIn {
		prevOutputScript, err := hex.DecodeString(utxos[i].ScriptPubKey)
		if err != nil {
			return err
		}
		txHash, err := chainhash.NewHashFromStr(utxos[i].TxID)
		if err != nil {
			return err
		}
		outPoint := wire.OutPoint{Hash: *txHash, Index: utxos[i].Vout}
		prevOutputFetcher := txscript.NewMultiPrevOutFetcher(
			map[wire.OutPoint]*wire.TxOut{outPoint: {Value: int64(utxos[i].Amount * precision), PkScript: prevOutputScript}},
		)
		sigHashes := txscript.NewTxSigHashes(c.tx, prevOutputFetcher)

		hash, err := txscript.CalcWitnessSigHash(prevOutputScript, sigHashes, txscript.SigHashAll, c.tx, int(utxos[i].Vout), int64(utxos[i].Amount*precision))
		if err != nil {
			return err
		}
		c.txInHash = append(c.txInHash, hash)
		c.hashCounter.Add(1)

		//signature, err := c.sign(hash) //todo
		//if err != nil {
		//	return err
		//}
		//txscript.SigHashAll, // https://www.btcstudy.org/2021/11/09/bitcoin-signature-types-sighash/
		//signature = append(signature, byte(txscript.SigHashAll)) //todo
		//
		//txIn.Witness = wire.TxWitness{signature, c.publicKey.SerializeCompressed()}
	}
	return nil
}

func (c *txClient) AddWitnessSignature(hash, signature []byte) bool {
	for i, inHash := range c.txInHash {
		if bytes.Equal(inHash, hash) {
			// txscript.SigHashAll, // https://www.btcstudy.org/2021/11/09/bitcoin-signature-types-sighash/
			signature = append(signature, byte(txscript.SigHashAll)) // todo
			c.tx.TxIn[i].Witness = wire.TxWitness{signature, c.publicKey.SerializeCompressed()}
			c.hashCounter.Add(-1)
		}
	}
	return c.hashCounter.Load() == 0
}

func (c *txClient) IsHaveWitnessSignature(hash []byte) bool {
	for _, inHash := range c.txInHash {
		if bytes.Equal(inHash, hash) {
			return true
		}
	}
	return false
}

// rawTxInWitnessSignature returns the serialized ECDA signature for the input
// idx of the given transaction, with the hashType appended to it. This
// function is identical to RawTxInSignature, however the signature generated
// signs a new sighash digest defined in BIP0143.
func (c *txClient) rawTxInWitnessSignature(sigHashes *txscript.TxSigHashes, idx int, amt int64, subScript []byte) ([]byte, error) {
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
func (c *txClient) witnessSignature(sigHashes *txscript.TxSigHashes, idx int, amt int64, subScript []byte) (wire.TxWitness, error) {
	hash, err := txscript.CalcWitnessSigHash(subScript, sigHashes, txscript.SigHashAll, c.tx, idx, amt)
	if err != nil {
		return nil, err
	}
	signature, err := c.sign(hash) ///todo
	if err != nil {
		return nil, err
	}
	// txscript.SigHashAll, // https://www.btcstudy.org/2021/11/09/bitcoin-signature-types-sighash/
	signature = append(signature, byte(txscript.SigHashAll))

	// A witness script is actually a stack, so we return an array of byte
	// slices here, rather than a single byte slice.
	return wire.TxWitness{signature, c.publicKey.SerializeCompressed()}, nil
}

func (c *txClient) sign(hash []byte) ([]byte, error) {
	c.nextSignChan <- hash
	select {
	case signature := <-c.signChan:
		return signature.Signature, nil
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	}
}

func (c *txClient) sendTx(allowHighFees bool) (*chainhash.Hash, error) {
	hash, err := c.client.SendRawTransaction(c.tx, allowHighFees)
	if err != nil {
		return nil, fmt.Errorf("send raw transaction: %w", err)
	}
	return hash, nil
}

func (c *txClient) BuildTx(to string, amount int64) error {
	return errors.Join(c.buildTxOut(to, amount), c.buildTxIn(amount))
}

type Receipt struct {
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

func (c *txClient) BuildMultiTransfer(receipt []Receipt) error {
	errs := make([]error, 0)
	var amount uint64
	for _, r := range receipt {
		errs = append(errs, c.buildTxOut(r.To, int64(r.Amount)))
		amount += r.Amount
	}

	if amount > 0 {
		errs = append(errs, c.buildTxIn(int64(amount)))
	}

	return errors.Join(errs...)
}

func (c *txClient) SendTx() error {
	_, err := c.sendTx(false)
	return err
}

func (c *txClient) TxHash() []byte {
	hash := c.tx.TxHash()
	return hash.CloneBytes()
}

func (c *txClient) NextSignTask() []byte {
	ctx, cancel := context.WithTimeout(c.ctx, time.Second*10)
	defer cancel()
	select {
	case msg := <-c.nextSignChan:
		return msg
	case <-ctx.Done():
		return nil
	}
}

func (c *txClient) WaitTxSuccess() error {
	h := c.tx.TxHash()
	begin := time.Now()
	defer func() {
		log.Infof("waitTxSuccess, duration_ms: %v", time.Since(begin).Milliseconds())
	}()

	count := 60
	for count > 0 {
		res, err := c.client.GetRawTransactionVerbose(&h)
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

func (c *txClient) Close() error {
	c.cancel()
	return nil
}

func (c *txClient) Post(hash, signature []byte) {
	c.signChan <- &SignatureCtx{
		Hash:      hash,
		Signature: signature,
	}
}

func (c *txClient) ChainType() int {
	return types.ChainBitcoin
}
