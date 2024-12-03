// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package codec

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Operation is an auto generated low-level Go binding around an user-defined struct.
type Operation struct {
	ManagerAddr common.Address
	State       uint8
	TaskId      uint64
	OptData     []byte
}

// VoterCodecMetaData contains all meta data concerning the VoterCodec contract.
var VoterCodecMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddr\",\"type\":\"address\"},{\"internalType\":\"enumState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"taskId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structOperation[]\",\"name\":\"opts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// VoterCodecABI is the input ABI used to generate the binding from.
// Deprecated: Use VoterCodecMetaData.ABI instead.
var VoterCodecABI = VoterCodecMetaData.ABI

// VoterCodec is an auto generated Go binding around an Ethereum contract.
type VoterCodec struct {
	VoterCodecCaller     // Read-only binding to the contract
	VoterCodecTransactor // Write-only binding to the contract
	VoterCodecFilterer   // Log filterer for contract events
}

// VoterCodecCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterCodecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterCodecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterCodecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterCodecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterCodecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterCodecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterCodecSession struct {
	Contract     *VoterCodec       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterCodecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterCodecCallerSession struct {
	Contract *VoterCodecCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VoterCodecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterCodecTransactorSession struct {
	Contract     *VoterCodecTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VoterCodecRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterCodecRaw struct {
	Contract *VoterCodec // Generic contract binding to access the raw methods on
}

// VoterCodecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterCodecCallerRaw struct {
	Contract *VoterCodecCaller // Generic read-only contract binding to access the raw methods on
}

// VoterCodecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterCodecTransactorRaw struct {
	Contract *VoterCodecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterCodec creates a new instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodec(address common.Address, backend bind.ContractBackend) (*VoterCodec, error) {
	contract, err := bindVoterCodec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterCodec{VoterCodecCaller: VoterCodecCaller{contract: contract}, VoterCodecTransactor: VoterCodecTransactor{contract: contract}, VoterCodecFilterer: VoterCodecFilterer{contract: contract}}, nil
}

// NewVoterCodecCaller creates a new read-only instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodecCaller(address common.Address, caller bind.ContractCaller) (*VoterCodecCaller, error) {
	contract, err := bindVoterCodec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterCodecCaller{contract: contract}, nil
}

// NewVoterCodecTransactor creates a new write-only instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodecTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterCodecTransactor, error) {
	contract, err := bindVoterCodec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterCodecTransactor{contract: contract}, nil
}

// NewVoterCodecFilterer creates a new log filterer instance of VoterCodec, bound to a specific deployed contract.
func NewVoterCodecFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterCodecFilterer, error) {
	contract, err := bindVoterCodec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterCodecFilterer{contract: contract}, nil
}

// bindVoterCodec binds a generic wrapper to an already deployed contract.
func bindVoterCodec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoterCodecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterCodec *VoterCodecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterCodec.Contract.VoterCodecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterCodec *VoterCodecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterCodec.Contract.VoterCodecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterCodec *VoterCodecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterCodec.Contract.VoterCodecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterCodec *VoterCodecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterCodec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterCodec *VoterCodecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterCodec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterCodec *VoterCodecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterCodec.Contract.contract.Transact(opts, method, params...)
}
