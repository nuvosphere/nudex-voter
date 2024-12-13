package sui

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/block-vision/sui-go-sdk/common/keypair"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/blake2b"
)

type RawSignTx struct{}

type UnSignTx models.TxnMetaData

type SignedTx models.SignedTransactionSerializedSig

func (t *UnSignTx) TxHash() []byte {
	return TxHash(t.TxBytes)
}

func (t *UnSignTx) SerializedSigWith(signature, publicKey []byte) *models.SignedTransactionSerializedSig {
	return &models.SignedTransactionSerializedSig{
		TxBytes:   t.TxBytes,
		Signature: models.ToSerializedSignature(signature, publicKey),
	}
}

// TxClient https://docs.sui.io/sui-api-ref#unsafe_paysui
type TxClient struct {
	client sui.ISuiAPI
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *TxClient) SetUrl(url string) *TxClient {
	c.client = sui.NewSuiClient(url)
	return c
}

func NewClient(ctx context.Context) *TxClient {
	ctx, cancel := context.WithCancel(ctx)
	return &TxClient{
		client: sui.NewSuiClient(constant.SuiMainnetEndpoint),
		ctx:    ctx,
		cancel: cancel,
	}
}

func NewDevClient() *TxClient {
	c := NewClient(context.Background())
	c.SetUrl("https://sui-devnet-endpoint.blockvision.org")
	return c
}

func (c *TxClient) GetBalance(address string, coin string, coinName string) (decimal.Decimal, error) {
	res, err := c.client.SuiXGetBalance(c.ctx, models.SuiXGetBalanceRequest{
		Owner:    address,
		CoinType: fmt.Sprintf("%s::%s::%s", coin, strings.ToLower(coinName), strings.ToUpper(coinName)), //"0x2::sui::SUI",
	})
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.RequireFromString(res.TotalBalance), nil
}

// pay: src20 token
// paySui: sui token

func (c *TxClient) BuildTransferTx(from, to, token string, amount uint64) (*UnSignTx, error) {
	//txIn, err := c.client.PaySui(c.ctx, models.PaySuiRequest{
	//	Signer:      from,
	//	SuiObjectId: []string{token},
	//	Recipient:   []string{to},
	//	Amount:      []string{fmt.Sprintf("%d", amount)},
	//	GasBudget:   "1000000", // todo
	//})
	txIn, err := c.client.TransferSui(c.ctx, models.TransferSuiRequest{
		Signer:      from,
		SuiObjectId: token,
		Recipient:   to,
		Amount:      fmt.Sprintf("%d", amount),
		GasBudget:   "1000000", // todo https://docs.sui.io/concepts/tokenomics/gas-in-sui
	})
	if err != nil {
		return nil, fmt.Errorf("transfer sui: %w", err)
	}

	return (*UnSignTx)(&txIn), nil
}

func (c *TxClient) SendTx(tx *SignedTx) ([]byte, error) {
	res, err := c.client.SuiExecuteTransactionBlock(c.ctx, models.SuiExecuteTransactionBlockRequest{
		TxBytes:   tx.TxBytes,
		Signature: []string{tx.Signature},
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
		RequestType: "WaitForLocalExecution",
	})
	if err != nil {
		return nil, fmt.Errorf("send transaction block: %w", err)
	}

	return []byte(res.Digest), nil
}

// TryExecuteTx todo: bug
func (c *TxClient) TryExecuteTx(tx *UnSignTx) (string, error) {
	res, err := c.client.SuiDryRunTransactionBlock(c.ctx, models.SuiDryRunTransactionBlockRequest{
		TxBytes: tx.TxBytes,
	})
	if err != nil {
		return "", fmt.Errorf("try execute transaction block: %w", err)
	}

	fmt.Println(utils.FormatJSON(res))
	return "", nil
}

// Ed25519PublicKeyToSuiAddress https://github.com/MystenLabs/sui/blob/main/sdk/typescript/src/cryptography/publickey.ts#L112
func Ed25519PublicKeyToSuiAddress(pubKey []byte) string {
	newPubkey := []byte{byte(keypair.Ed25519Flag)}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)
	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}

func TxHash(data string) []byte {
	txBytes, _ := base64.StdEncoding.DecodeString(data)
	message := models.MessageWithIntent(txBytes)
	digest := blake2b.Sum256(message)
	return digest[:]
}
