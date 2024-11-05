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
	*bind.BoundContract
	Address common.Address
	ABI     abi.ABI
}

func NewContractCaller(address common.Address, abi abi.ABI, caller bind.ContractCaller) *Contract {
	return NewContract(address, abi, caller, nil, nil)
}

func NewContractTransactor(address common.Address, abi abi.ABI, caller bind.ContractCaller, transactor bind.ContractTransactor) *Contract {
	return NewContract(address, abi, caller, transactor, nil)
}

func NewContractFilterer(address common.Address, abi abi.ABI, filterer bind.ContractFilterer) *Contract {
	return NewContract(address, abi, nil, nil, filterer)
}

func NewContract(address common.Address, abi abi.ABI, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) *Contract {
	contract := &Contract{
		BoundContract: bind.NewBoundContract(address, abi, caller, transactor, filterer),
		Address:       address,
		ABI:           abi,
	}

	return contract
}

func (c *Contract) Encode(method string, args ...interface{}) ([]byte, error) {
	m, ok := c.ABI.Methods[method]
	if !ok {
		return nil, fmt.Errorf("contract method %s not found", method)
	}

	input, err := m.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}

	input = append(m.ID, input...)

	return input, nil
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
