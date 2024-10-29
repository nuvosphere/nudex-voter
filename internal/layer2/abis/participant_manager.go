// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// ParticipantManagerContractMetaData contains all meta data concerning the ParticipantManagerContract contract.
var ParticipantManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addParticipant\",\"inputs\":[{\"name\":\"newParticipant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getParticipants\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomParticipant\",\"inputs\":[{\"name\":\"_salt\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"randParticipant\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_nuvoLock\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minLockAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minLockPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialParticipant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isEligible\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isParticipant\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minLockAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minLockPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nuvoLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINuvoLock\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participants\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeParticipant\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParticipantAdded\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParticipantRemoved\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ParticipantManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ParticipantManagerContractMetaData.ABI instead.
var ParticipantManagerContractABI = ParticipantManagerContractMetaData.ABI

// ParticipantManagerContract is an auto generated Go binding around an Ethereum contract.
type ParticipantManagerContract struct {
	ParticipantManagerContractCaller     // Read-only binding to the contract
	ParticipantManagerContractTransactor // Write-only binding to the contract
	ParticipantManagerContractFilterer   // Log filterer for contract events
}

// ParticipantManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParticipantManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParticipantManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParticipantManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParticipantManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParticipantManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParticipantManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParticipantManagerContractSession struct {
	Contract     *ParticipantManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ParticipantManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParticipantManagerContractCallerSession struct {
	Contract *ParticipantManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ParticipantManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParticipantManagerContractTransactorSession struct {
	Contract     *ParticipantManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ParticipantManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParticipantManagerContractRaw struct {
	Contract *ParticipantManagerContract // Generic contract binding to access the raw methods on
}

// ParticipantManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParticipantManagerContractCallerRaw struct {
	Contract *ParticipantManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// ParticipantManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParticipantManagerContractTransactorRaw struct {
	Contract *ParticipantManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParticipantManagerContract creates a new instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContract(address common.Address, backend bind.ContractBackend) (*ParticipantManagerContract, error) {
	contract, err := bindParticipantManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContract{ParticipantManagerContractCaller: ParticipantManagerContractCaller{contract: contract}, ParticipantManagerContractTransactor: ParticipantManagerContractTransactor{contract: contract}, ParticipantManagerContractFilterer: ParticipantManagerContractFilterer{contract: contract}}, nil
}

// NewParticipantManagerContractCaller creates a new read-only instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContractCaller(address common.Address, caller bind.ContractCaller) (*ParticipantManagerContractCaller, error) {
	contract, err := bindParticipantManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractCaller{contract: contract}, nil
}

// NewParticipantManagerContractTransactor creates a new write-only instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ParticipantManagerContractTransactor, error) {
	contract, err := bindParticipantManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractTransactor{contract: contract}, nil
}

// NewParticipantManagerContractFilterer creates a new log filterer instance of ParticipantManagerContract, bound to a specific deployed contract.
func NewParticipantManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ParticipantManagerContractFilterer, error) {
	contract, err := bindParticipantManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractFilterer{contract: contract}, nil
}

// bindParticipantManagerContract binds a generic wrapper to an already deployed contract.
func bindParticipantManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParticipantManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParticipantManagerContract *ParticipantManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParticipantManagerContract.Contract.ParticipantManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParticipantManagerContract *ParticipantManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.ParticipantManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParticipantManagerContract *ParticipantManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.ParticipantManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParticipantManagerContract *ParticipantManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParticipantManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParticipantManagerContract *ParticipantManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParticipantManagerContract *ParticipantManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.contract.Transact(opts, method, params...)
}

// GetParticipants is a free data retrieval call binding the contract method 0x5aa68ac0.
//
// Solidity: function getParticipants() view returns(address[])
func (_ParticipantManagerContract *ParticipantManagerContractCaller) GetParticipants(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "getParticipants")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetParticipants is a free data retrieval call binding the contract method 0x5aa68ac0.
//
// Solidity: function getParticipants() view returns(address[])
func (_ParticipantManagerContract *ParticipantManagerContractSession) GetParticipants() ([]common.Address, error) {
	return _ParticipantManagerContract.Contract.GetParticipants(&_ParticipantManagerContract.CallOpts)
}

// GetParticipants is a free data retrieval call binding the contract method 0x5aa68ac0.
//
// Solidity: function getParticipants() view returns(address[])
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) GetParticipants() ([]common.Address, error) {
	return _ParticipantManagerContract.Contract.GetParticipants(&_ParticipantManagerContract.CallOpts)
}

// GetRandomParticipant is a free data retrieval call binding the contract method 0x6235b415.
//
// Solidity: function getRandomParticipant(address _salt) view returns(address randParticipant)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) GetRandomParticipant(opts *bind.CallOpts, _salt common.Address) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "getRandomParticipant", _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRandomParticipant is a free data retrieval call binding the contract method 0x6235b415.
//
// Solidity: function getRandomParticipant(address _salt) view returns(address randParticipant)
func (_ParticipantManagerContract *ParticipantManagerContractSession) GetRandomParticipant(_salt common.Address) (common.Address, error) {
	return _ParticipantManagerContract.Contract.GetRandomParticipant(&_ParticipantManagerContract.CallOpts, _salt)
}

// GetRandomParticipant is a free data retrieval call binding the contract method 0x6235b415.
//
// Solidity: function getRandomParticipant(address _salt) view returns(address randParticipant)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) GetRandomParticipant(_salt common.Address) (common.Address, error) {
	return _ParticipantManagerContract.Contract.GetRandomParticipant(&_ParticipantManagerContract.CallOpts, _salt)
}

// IsEligible is a free data retrieval call binding the contract method 0x66e305fd.
//
// Solidity: function isEligible(address participant) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) IsEligible(opts *bind.CallOpts, participant common.Address) (bool, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "isEligible", participant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEligible is a free data retrieval call binding the contract method 0x66e305fd.
//
// Solidity: function isEligible(address participant) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractSession) IsEligible(participant common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.IsEligible(&_ParticipantManagerContract.CallOpts, participant)
}

// IsEligible is a free data retrieval call binding the contract method 0x66e305fd.
//
// Solidity: function isEligible(address participant) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) IsEligible(participant common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.IsEligible(&_ParticipantManagerContract.CallOpts, participant)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) IsParticipant(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "isParticipant", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.IsParticipant(&_ParticipantManagerContract.CallOpts, arg0)
}

// IsParticipant is a free data retrieval call binding the contract method 0x929066f5.
//
// Solidity: function isParticipant(address ) view returns(bool)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _ParticipantManagerContract.Contract.IsParticipant(&_ParticipantManagerContract.CallOpts, arg0)
}

// MinLockAmount is a free data retrieval call binding the contract method 0x08804275.
//
// Solidity: function minLockAmount() view returns(uint256)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) MinLockAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "minLockAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLockAmount is a free data retrieval call binding the contract method 0x08804275.
//
// Solidity: function minLockAmount() view returns(uint256)
func (_ParticipantManagerContract *ParticipantManagerContractSession) MinLockAmount() (*big.Int, error) {
	return _ParticipantManagerContract.Contract.MinLockAmount(&_ParticipantManagerContract.CallOpts)
}

// MinLockAmount is a free data retrieval call binding the contract method 0x08804275.
//
// Solidity: function minLockAmount() view returns(uint256)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) MinLockAmount() (*big.Int, error) {
	return _ParticipantManagerContract.Contract.MinLockAmount(&_ParticipantManagerContract.CallOpts)
}

// MinLockPeriod is a free data retrieval call binding the contract method 0x73ae54b5.
//
// Solidity: function minLockPeriod() view returns(uint256)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) MinLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "minLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLockPeriod is a free data retrieval call binding the contract method 0x73ae54b5.
//
// Solidity: function minLockPeriod() view returns(uint256)
func (_ParticipantManagerContract *ParticipantManagerContractSession) MinLockPeriod() (*big.Int, error) {
	return _ParticipantManagerContract.Contract.MinLockPeriod(&_ParticipantManagerContract.CallOpts)
}

// MinLockPeriod is a free data retrieval call binding the contract method 0x73ae54b5.
//
// Solidity: function minLockPeriod() view returns(uint256)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) MinLockPeriod() (*big.Int, error) {
	return _ParticipantManagerContract.Contract.MinLockPeriod(&_ParticipantManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) NuvoLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "nuvoLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractSession) NuvoLock() (common.Address, error) {
	return _ParticipantManagerContract.Contract.NuvoLock(&_ParticipantManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) NuvoLock() (common.Address, error) {
	return _ParticipantManagerContract.Contract.NuvoLock(&_ParticipantManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractSession) Owner() (common.Address, error) {
	return _ParticipantManagerContract.Contract.Owner(&_ParticipantManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) Owner() (common.Address, error) {
	return _ParticipantManagerContract.Contract.Owner(&_ParticipantManagerContract.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParticipantManagerContract.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _ParticipantManagerContract.Contract.Participants(&_ParticipantManagerContract.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_ParticipantManagerContract *ParticipantManagerContractCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _ParticipantManagerContract.Contract.Participants(&_ParticipantManagerContract.CallOpts, arg0)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address newParticipant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) AddParticipant(opts *bind.TransactOpts, newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "addParticipant", newParticipant)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address newParticipant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) AddParticipant(newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.AddParticipant(&_ParticipantManagerContract.TransactOpts, newParticipant)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address newParticipant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) AddParticipant(newParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.AddParticipant(&_ParticipantManagerContract.TransactOpts, newParticipant)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _nuvoLock, uint256 _minLockAmount, uint256 _minLockPeriod, address _owner, address _initialParticipant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) Initialize(opts *bind.TransactOpts, _nuvoLock common.Address, _minLockAmount *big.Int, _minLockPeriod *big.Int, _owner common.Address, _initialParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "initialize", _nuvoLock, _minLockAmount, _minLockPeriod, _owner, _initialParticipant)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _nuvoLock, uint256 _minLockAmount, uint256 _minLockPeriod, address _owner, address _initialParticipant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) Initialize(_nuvoLock common.Address, _minLockAmount *big.Int, _minLockPeriod *big.Int, _owner common.Address, _initialParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.Initialize(&_ParticipantManagerContract.TransactOpts, _nuvoLock, _minLockAmount, _minLockPeriod, _owner, _initialParticipant)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _nuvoLock, uint256 _minLockAmount, uint256 _minLockPeriod, address _owner, address _initialParticipant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) Initialize(_nuvoLock common.Address, _minLockAmount *big.Int, _minLockPeriod *big.Int, _owner common.Address, _initialParticipant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.Initialize(&_ParticipantManagerContract.TransactOpts, _nuvoLock, _minLockAmount, _minLockPeriod, _owner, _initialParticipant)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x668a2001.
//
// Solidity: function removeParticipant(address participant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) RemoveParticipant(opts *bind.TransactOpts, participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "removeParticipant", participant)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x668a2001.
//
// Solidity: function removeParticipant(address participant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) RemoveParticipant(participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RemoveParticipant(&_ParticipantManagerContract.TransactOpts, participant)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0x668a2001.
//
// Solidity: function removeParticipant(address participant) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) RemoveParticipant(participant common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RemoveParticipant(&_ParticipantManagerContract.TransactOpts, participant)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RenounceOwnership(&_ParticipantManagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.RenounceOwnership(&_ParticipantManagerContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParticipantManagerContract *ParticipantManagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.TransferOwnership(&_ParticipantManagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParticipantManagerContract *ParticipantManagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParticipantManagerContract.Contract.TransferOwnership(&_ParticipantManagerContract.TransactOpts, newOwner)
}

// ParticipantManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractInitializedIterator struct {
	Event *ParticipantManagerContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParticipantManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParticipantManagerContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParticipantManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractInitialized represents a Initialized event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ParticipantManagerContractInitializedIterator, error) {

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractInitializedIterator{contract: _ParticipantManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractInitialized)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseInitialized(log types.Log) (*ParticipantManagerContractInitialized, error) {
	event := new(ParticipantManagerContractInitialized)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractOwnershipTransferredIterator struct {
	Event *ParticipantManagerContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParticipantManagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParticipantManagerContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParticipantManagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParticipantManagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractOwnershipTransferredIterator{contract: _ParticipantManagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractOwnershipTransferred)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*ParticipantManagerContractOwnershipTransferred, error) {
	event := new(ParticipantManagerContractOwnershipTransferred)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractParticipantAddedIterator is returned from FilterParticipantAdded and is used to iterate over the raw logs and unpacked data for ParticipantAdded events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantAddedIterator struct {
	Event *ParticipantManagerContractParticipantAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParticipantManagerContractParticipantAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractParticipantAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParticipantManagerContractParticipantAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParticipantManagerContractParticipantAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractParticipantAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractParticipantAdded represents a ParticipantAdded event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantAdded struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantAdded is a free log retrieval operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterParticipantAdded(opts *bind.FilterOpts, participant []common.Address) (*ParticipantManagerContractParticipantAddedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "ParticipantAdded", participantRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractParticipantAddedIterator{contract: _ParticipantManagerContract.contract, event: "ParticipantAdded", logs: logs, sub: sub}, nil
}

// WatchParticipantAdded is a free log subscription operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchParticipantAdded(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractParticipantAdded, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "ParticipantAdded", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractParticipantAdded)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipantAdded is a log parse operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseParticipantAdded(log types.Log) (*ParticipantManagerContractParticipantAdded, error) {
	event := new(ParticipantManagerContractParticipantAdded)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParticipantManagerContractParticipantRemovedIterator is returned from FilterParticipantRemoved and is used to iterate over the raw logs and unpacked data for ParticipantRemoved events raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantRemovedIterator struct {
	Event *ParticipantManagerContractParticipantRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParticipantManagerContractParticipantRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParticipantManagerContractParticipantRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParticipantManagerContractParticipantRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParticipantManagerContractParticipantRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParticipantManagerContractParticipantRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParticipantManagerContractParticipantRemoved represents a ParticipantRemoved event raised by the ParticipantManagerContract contract.
type ParticipantManagerContractParticipantRemoved struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantRemoved is a free log retrieval operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) FilterParticipantRemoved(opts *bind.FilterOpts, participant []common.Address) (*ParticipantManagerContractParticipantRemovedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.FilterLogs(opts, "ParticipantRemoved", participantRule)
	if err != nil {
		return nil, err
	}
	return &ParticipantManagerContractParticipantRemovedIterator{contract: _ParticipantManagerContract.contract, event: "ParticipantRemoved", logs: logs, sub: sub}, nil
}

// WatchParticipantRemoved is a free log subscription operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) WatchParticipantRemoved(opts *bind.WatchOpts, sink chan<- *ParticipantManagerContractParticipantRemoved, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ParticipantManagerContract.contract.WatchLogs(opts, "ParticipantRemoved", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParticipantManagerContractParticipantRemoved)
				if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipantRemoved is a log parse operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_ParticipantManagerContract *ParticipantManagerContractFilterer) ParseParticipantRemoved(log types.Log) (*ParticipantManagerContractParticipantRemoved, error) {
	event := new(ParticipantManagerContractParticipantRemoved)
	if err := _ParticipantManagerContract.contract.UnpackLog(event, "ParticipantRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
