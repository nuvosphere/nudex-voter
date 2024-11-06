package contracts

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/config"
)

var (
	VotingAddress      = common.HexToAddress(config.AppConfig.VotingContract)
	AccountAddress     = common.HexToAddress(config.AppConfig.AccountContract)
	OperationsAddress  = common.HexToAddress(config.AppConfig.OperationsContract)
	ParticipantAddress = common.HexToAddress(config.AppConfig.ParticipantContract)
	DepositAddress     = common.HexToAddress(config.AppConfig.DepositContract)
)

func Pack(meta *bind.MetaData, method string, params ...interface{}) []byte {
	a, err := meta.GetAbi()
	if err != nil {
		panic(err)
	}

	// Otherwise pack up the parameters and invoke the contract
	input, err := a.Pack(method, params...)
	if err != nil {
		panic(err)
	}

	return input
}

func EncodeFun(abiStr string, method string, params ...interface{}) []byte {
	return Pack(
		&bind.MetaData{
			ABI: abiStr,
		},
		method,
		params...,
	)
}

func ParseABI(abiJSON string) (abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unable to parse abi json: %w", err)
	}

	return parsed, nil
}

func MustParseABI(abiJSON string) abi.ABI {
	parsed, err := ParseABI(abiJSON)
	if err != nil {
		panic(err)
	}

	return parsed
}

type Contract struct {
	Address common.Address
	ABI     abi.ABI
}

func NewContract(address common.Address, abi abi.ABI) *Contract {
	return &Contract{
		Address: address,
		ABI:     abi,
	}
}

func (c *Contract) Encode(method string, args ...any) ([]byte, error) {
	// Otherwise pack up the parameters and invoke the contract
	return c.ABI.Pack(method, args...)
}

func (c *Contract) Decode(method string, data []byte) ([]any, error) {
	return c.ABI.Unpack(method, data)
}

func (c *Contract) EventTopicHash(eventName string) (common.Hash, error) {
	ev, ok := c.ABI.Events[eventName]
	if !ok {
		return common.Hash{}, fmt.Errorf("ethcontract: event '%s' not found in contract abi", eventName)
	}

	h := crypto.Keccak256Hash([]byte(ev.Sig))

	return h, nil
}

func EncodeTransferOfERC20(from, to common.Address, amount *big.Int) []byte {
	return EncodeFun(ERC20ABI, "transferFrom", from, to, amount)
}

func EncodeVerifyAndCall(_target common.Address, _data []byte, _signature []byte) []byte {
	return EncodeFun(VotingManagerContractABI, "verifyAndCall", _target, _data, _signature)
}

func EncodeSubmitTaskReceipt(taskId *big.Int, result []byte, signature []byte) []byte {
	return EncodeFun(VotingManagerContractABI, "submitTaskReceipt", taskId, result, signature)
}
