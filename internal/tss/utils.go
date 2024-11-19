package tss

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"time"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	ecdsaResharing "github.com/bnb-chain/tss-lib/v2/ecdsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	eddsaResharing "github.com/bnb-chain/tss-lib/v2/eddsa/resharing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

func Partners() types.Participants {
	// todo online contact get address list
	return lo.Map(
		config.AppConfig.TssPublicKeys,
		func(pubKey *ecdsa.PublicKey, _ int) common.Address { return crypto.PubkeyToAddress(*pubKey) },
	)
}

func createPartyIDs(publicKeys []*ecdsa.PublicKey) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(publicKeys))

	for i, publicKey := range publicKeys {
		key, _ := new(big.Int).SetString(ConvertPubKeyToHex(publicKey), 16)
		tssAllPartyIDs[i] = tss.NewPartyID(
			crypto.PubkeyToAddress(*publicKeys[i]).Hex(),
			crypto.PubkeyToAddress(*publicKeys[i]).Hex(),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func createPartyIDsByAddress(addressList types.Participants) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	for i, address := range addressList {
		key := address.Big()
		tssAllPartyIDs[i] = tss.NewPartyID(
			address.Hex(),
			address.Hex(),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func createOldPartyIDsByAddress(addressList types.Participants) tss.SortedPartyIDs {
	tssAllPartyIDs := make(tss.UnSortedPartyIDs, len(addressList))

	for i, address := range addressList {
		key := address.Big()
		key = new(big.Int).Add(key, big.NewInt(1)) // key + 1
		tssAllPartyIDs[i] = tss.NewPartyID(
			key.Text(16),
			key.Text(16),
			key,
		)
	}

	return tss.SortPartyIDs(tssAllPartyIDs)
}

func SaveTSSData(data *helper.LocalPartySaveData) error {
	curveType := data.CurveType()

	dataDir := filepath.Join(config.AppConfig.DbDir, "tss_data", curveType.CurveName())
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		log.Errorf("Failed to create TSS data directory: %v", err)
		return err
	}

	dataBytes, err := json.Marshal(data.GetData())
	if err != nil {
		log.Errorf("Unable to serialize TSS data: %v", err)
		return err
	}

	filePath := filepath.Join(dataDir, "tss_key_data.json")
	if err := os.WriteFile(filePath, dataBytes, 0o600); err != nil {
		log.Errorf("Failed to save TSS data to file: %v", err)
		return err
	}

	log.Infof("TSS data successfully saved to: %s", filePath)

	return nil
}

func LoadTSSData(ec helper.CurveType) (*helper.LocalPartySaveData, error) {
	filePath := filepath.Join(config.AppConfig.DbDir, "tss_data", ec.CurveName(), "tss_key_data.json")

	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read TSS data file: %v", err)
	}

	switch ec {
	case helper.ECDSA:
		var data ecdsaKeygen.LocalPartySaveData
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, fmt.Errorf("unable to deserialize TSS data: %v", err)
		}

		return helper.BuildECDSALocalPartySaveData().SetData(&data), nil
	case helper.EDDSA:
		var data keygen.LocalPartySaveData
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, fmt.Errorf("unable to deserialize TSS data: %v", err)
		}

		return helper.BuildEDDSALocalPartySaveData().SetData(&data), nil
	}

	return nil, fmt.Errorf("unknown elliptic curve")
}

func PublicKeysToHex(pubKeys []*ecdsa.PublicKey) []string {
	hexStrings := make([]string, len(pubKeys))
	for i, pubKey := range pubKeys {
		hexStrings[i] = ConvertPubKeyToHex(pubKey)
	}

	return hexStrings
}

func ConvertPubKeyToHex(pubKey *ecdsa.PublicKey) string {
	if pubKey == nil {
		return ""
	}

	var prefix byte
	if pubKey.Y.Bit(0) == 0 {
		prefix = 0x02 // Even Y-coordinate
	} else {
		prefix = 0x03 // Odd Y-coordinate
	}

	// Create the public key byte slice with the prefix
	pubKeyBytes := append([]byte{prefix}, pubKey.X.Bytes()...)

	// Convert to hex string
	return hex.EncodeToString(pubKeyBytes)
}

func CompareStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func AddressIndex(addressList types.Participants, tssAddress common.Address) int {
	for i, address := range addressList {
		if address == tssAddress {
			return i // Return the index if a match is found
		}
	}

	return -1 // Return -1 if not found
}

func extractToIds(message tss.Message) []string {
	recipients := message.GetTo()

	ids := make([]string, len(recipients))

	for i, recipient := range recipients {
		ids[i] = recipient.GetId()
	}

	return ids
}

func serializeMessageToBeSigned(nonce uint64, data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	bufBytes := buf.Bytes()
	length := len(bufBytes)
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)

	return append(append(nonceBytes, lengthBytes...), data...), nil
}

var ErrCoinType = fmt.Errorf("error coin type")

func getCoinTypeByChain(chain uint8) int {
	switch chain {
	case db.WalletTypeEVM:
		return 60
	case db.WalletTypeBTC:
		return 0
	case db.WalletTypeSOL:
		return 501
	case db.WalletTypeSUI:
		return 784
	default:
		panic(ErrCoinType)
	}
}

// RunKeyGen starts the local keygen party and handles incoming and outgoing
// messages to other parties.
func RunKeyGen(
	ctx context.Context,
	ty helper.CurveType,
	localSubmitter common.Address, // current submitter
	allPartners types.Participants,
	transport helper.Transporter,
) (tss.Party, map[string]*tss.PartyID, chan *helper.LocalPartySaveData, chan *tss.Error) {
	// outgoing messages to other peers
	outCh := make(chan tss.Message, 10)
	// error if keygen fails, contains culprits to blame
	errCh := make(chan *tss.Error, 10)

	var party tss.Party

	log.Debug("creating new local party")

	outEndCh := make(chan *helper.LocalPartySaveData)
	// output data when keygen finished
	ecdsaEndCh := make(chan *ecdsaKeygen.LocalPartySaveData)
	eddsaEndCh := make(chan *eddsaKeygen.LocalPartySaveData)

	params, partyIdMap := NewParam(ty.EC(), localSubmitter, allPartners)

	switch ty {
	case helper.ECDSA:
		preParams, err := ecdsaKeygen.GeneratePreParams(1 * time.Minute)
		if err != nil {
			panic(err)
		}

		party = ecdsaKeygen.NewLocalParty(params, outCh, ecdsaEndCh, *preParams)
	case helper.EDDSA:
		party = eddsaKeygen.NewLocalParty(params, outCh, eddsaEndCh)
	default:
		panic("implement me")
	}

	go func() {
		defer close(eddsaEndCh)
		defer close(ecdsaEndCh)
		select {
		case <-ctx.Done():
			return
		case data := <-ecdsaEndCh:
			outEndCh <- helper.BuildECDSALocalPartySaveData().SetData(data)

			close(ecdsaEndCh)
		case data := <-eddsaEndCh:
			outEndCh <- helper.BuildEDDSALocalPartySaveData().SetData(data)

			close(eddsaEndCh)
		}
	}()

	log.Debug("local party created", "partyID", party.PartyID())

	helper.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, partyIdMap, outEndCh, errCh
}

// RunReshare starts the local reshare party and handles incoming and outgoing
// messages to other parties.
func RunReshare(
	ctx context.Context,
	params *tss.ReSharingParameters,
	key helper.LocalPartySaveData,
	transport helper.Transporter,
) (tss.Party, chan *helper.LocalPartySaveData, chan *tss.Error) {
	// outgoing messages to other peers
	outCh := make(chan tss.Message, 1)
	// error if reshare fails, contains culprits to blame
	errCh := make(chan *tss.Error, 1)

	log.Debug("creating new local party")

	outEndCh := make(chan *helper.LocalPartySaveData)
	// output data when keygen finished
	ecdsaEndCh := make(chan *ecdsaKeygen.LocalPartySaveData)
	eddsaEndCh := make(chan *eddsaKeygen.LocalPartySaveData)

	var party tss.Party

	switch key.CurveType() {
	case helper.ECDSA:
		party = ecdsaResharing.NewLocalParty(params, *key.ECDSAData(), outCh, ecdsaEndCh)

	case helper.EDDSA:
		party = eddsaResharing.NewLocalParty(params, *key.EDDSAData(), outCh, eddsaEndCh)

	default:
		panic("implement me")
	}

	go func() {
		defer close(eddsaEndCh)
		defer close(ecdsaEndCh)
		select {
		case <-ctx.Done():
			return
		case data := <-ecdsaEndCh:
			outEndCh <- helper.BuildECDSALocalPartySaveData().SetData(data)

			close(ecdsaEndCh)
		case data := <-eddsaEndCh:
			outEndCh <- helper.BuildEDDSALocalPartySaveData().SetData(data)

			close(eddsaEndCh)
		}
	}()

	log.Debug("local resharing party created", "partyID", party.PartyID())

	helper.RunParty(ctx, party, errCh, outCh, transport, true)

	return party, outEndCh, errCh
}

func RunParty(
	ctx context.Context,
	msg *big.Int,
	params *tss.Parameters,
	key helper.LocalPartySaveData,
	transport helper.Transporter,
	keyDerivationDelta *big.Int,
) (tss.Party, chan *tsscommon.SignatureData, chan *tss.Error) {
	// outgoing messages to other peers - not one to not deadlock when a party
	// round is waiting for outgoing messages channel to clear
	outCh := make(chan tss.Message, params.PartyCount())
	// output signature when finished
	endCh := make(chan *tsscommon.SignatureData, 1)
	// error if signing fails, contains culprits to blame
	errCh := make(chan *tss.Error, 1)

	log.Debug("creating new local party")

	var party tss.Party

	switch key.CurveType() {
	case helper.ECDSA:
		if keyDerivationDelta != nil {
			party = signing.NewLocalPartyWithKDD(msg, params, *key.ECDSAData(), keyDerivationDelta, outCh, endCh)
		} else {
			party = signing.NewLocalParty(msg, params, *key.ECDSAData(), outCh, endCh)
		}

	default:
		panic("implement me")
	}

	log.Debug("local signing party created", "partyID", party.PartyID())

	helper.RunParty(ctx, party, errCh, outCh, transport, false)

	return party, endCh, errCh
}
