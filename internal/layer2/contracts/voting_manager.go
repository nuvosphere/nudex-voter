// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VotingManagerContractMetaData contains all meta data concerning the VotingManagerContract contract.
var VotingManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"chooseNewSubmitter\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forcedRotationWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_tssSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_participantManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nuvoLock\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastSubmissionTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nuvoLock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINuvoLock\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operationHash\",\"inputs\":[{\"name\":\"_opts\",\"type\":\"tuple[]\",\"internalType\":\"structOperation[]\",\"components\":[{\"name\":\"managerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"optData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"participantManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIParticipantManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setSignerAddress\",\"inputs\":[{\"name\":\"_newSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskCompletionThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndCall\",\"inputs\":[{\"name\":\"_opts\",\"type\":\"tuple[]\",\"internalType\":\"structOperation[]\",\"components\":[{\"name\":\"managerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"optData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyOperation\",\"inputs\":[{\"name\":\"_opts\",\"type\":\"tuple[]\",\"internalType\":\"structOperation[]\",\"components\":[{\"name\":\"managerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"optData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperationFailed\",\"inputs\":[{\"name\":\"errInfo\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubmitterChosen\",\"inputs\":[{\"name\":\"newSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubmitterRotationRequested\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"currentSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyOperationsArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectSubmitter\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RotationWindowNotPassed\",\"inputs\":[{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"window\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TaskAlreadyCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// VotingManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingManagerContractMetaData.ABI instead.
var VotingManagerContractABI = VotingManagerContractMetaData.ABI

// VotingManagerContract is an auto generated Go binding around an Ethereum contract.
type VotingManagerContract struct {
	VotingManagerContractCaller     // Read-only binding to the contract
	VotingManagerContractTransactor // Write-only binding to the contract
	VotingManagerContractFilterer   // Log filterer for contract events
}

// VotingManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingManagerContractSession struct {
	Contract     *VotingManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VotingManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingManagerContractCallerSession struct {
	Contract *VotingManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// VotingManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingManagerContractTransactorSession struct {
	Contract     *VotingManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// VotingManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingManagerContractRaw struct {
	Contract *VotingManagerContract // Generic contract binding to access the raw methods on
}

// VotingManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingManagerContractCallerRaw struct {
	Contract *VotingManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// VotingManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingManagerContractTransactorRaw struct {
	Contract *VotingManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingManagerContract creates a new instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContract(address common.Address, backend bind.ContractBackend) (*VotingManagerContract, error) {
	contract, err := bindVotingManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContract{VotingManagerContractCaller: VotingManagerContractCaller{contract: contract}, VotingManagerContractTransactor: VotingManagerContractTransactor{contract: contract}, VotingManagerContractFilterer: VotingManagerContractFilterer{contract: contract}}, nil
}

// NewVotingManagerContractCaller creates a new read-only instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContractCaller(address common.Address, caller bind.ContractCaller) (*VotingManagerContractCaller, error) {
	contract, err := bindVotingManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractCaller{contract: contract}, nil
}

// NewVotingManagerContractTransactor creates a new write-only instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingManagerContractTransactor, error) {
	contract, err := bindVotingManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractTransactor{contract: contract}, nil
}

// NewVotingManagerContractFilterer creates a new log filterer instance of VotingManagerContract, bound to a specific deployed contract.
func NewVotingManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingManagerContractFilterer, error) {
	contract, err := bindVotingManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractFilterer{contract: contract}, nil
}

// bindVotingManagerContract binds a generic wrapper to an already deployed contract.
func bindVotingManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingManagerContract *VotingManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingManagerContract.Contract.VotingManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingManagerContract *VotingManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.VotingManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingManagerContract *VotingManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.VotingManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingManagerContract *VotingManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingManagerContract *VotingManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingManagerContract *VotingManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.contract.Transact(opts, method, params...)
}

// ForcedRotationWindow is a free data retrieval call binding the contract method 0xc858379d.
//
// Solidity: function forcedRotationWindow() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) ForcedRotationWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "forcedRotationWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForcedRotationWindow is a free data retrieval call binding the contract method 0xc858379d.
//
// Solidity: function forcedRotationWindow() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) ForcedRotationWindow() (*big.Int, error) {
	return _VotingManagerContract.Contract.ForcedRotationWindow(&_VotingManagerContract.CallOpts)
}

// ForcedRotationWindow is a free data retrieval call binding the contract method 0xc858379d.
//
// Solidity: function forcedRotationWindow() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) ForcedRotationWindow() (*big.Int, error) {
	return _VotingManagerContract.Contract.ForcedRotationWindow(&_VotingManagerContract.CallOpts)
}

// LastSubmissionTime is a free data retrieval call binding the contract method 0x4f70104e.
//
// Solidity: function lastSubmissionTime() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) LastSubmissionTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "lastSubmissionTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSubmissionTime is a free data retrieval call binding the contract method 0x4f70104e.
//
// Solidity: function lastSubmissionTime() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) LastSubmissionTime() (*big.Int, error) {
	return _VotingManagerContract.Contract.LastSubmissionTime(&_VotingManagerContract.CallOpts)
}

// LastSubmissionTime is a free data retrieval call binding the contract method 0x4f70104e.
//
// Solidity: function lastSubmissionTime() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) LastSubmissionTime() (*big.Int, error) {
	return _VotingManagerContract.Contract.LastSubmissionTime(&_VotingManagerContract.CallOpts)
}

// NextSubmitter is a free data retrieval call binding the contract method 0x249c94e8.
//
// Solidity: function nextSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) NextSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "nextSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextSubmitter is a free data retrieval call binding the contract method 0x249c94e8.
//
// Solidity: function nextSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) NextSubmitter() (common.Address, error) {
	return _VotingManagerContract.Contract.NextSubmitter(&_VotingManagerContract.CallOpts)
}

// NextSubmitter is a free data retrieval call binding the contract method 0x249c94e8.
//
// Solidity: function nextSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) NextSubmitter() (common.Address, error) {
	return _VotingManagerContract.Contract.NextSubmitter(&_VotingManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) NuvoLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "nuvoLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) NuvoLock() (common.Address, error) {
	return _VotingManagerContract.Contract.NuvoLock(&_VotingManagerContract.CallOpts)
}

// NuvoLock is a free data retrieval call binding the contract method 0xe1bf9f1a.
//
// Solidity: function nuvoLock() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) NuvoLock() (common.Address, error) {
	return _VotingManagerContract.Contract.NuvoLock(&_VotingManagerContract.CallOpts)
}

// OperationHash is a free data retrieval call binding the contract method 0x80f12264.
//
// Solidity: function operationHash((address,uint8,uint64,bytes)[] _opts, uint256 _nonce) pure returns(bytes32 hash, bytes32 messageHash)
func (_VotingManagerContract *VotingManagerContractCaller) OperationHash(opts *bind.CallOpts, _opts []Operation, _nonce *big.Int) (struct {
	Hash        [32]byte
	MessageHash [32]byte
}, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "operationHash", _opts, _nonce)

	outstruct := new(struct {
		Hash        [32]byte
		MessageHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.MessageHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// OperationHash is a free data retrieval call binding the contract method 0x80f12264.
//
// Solidity: function operationHash((address,uint8,uint64,bytes)[] _opts, uint256 _nonce) pure returns(bytes32 hash, bytes32 messageHash)
func (_VotingManagerContract *VotingManagerContractSession) OperationHash(_opts []Operation, _nonce *big.Int) (struct {
	Hash        [32]byte
	MessageHash [32]byte
}, error) {
	return _VotingManagerContract.Contract.OperationHash(&_VotingManagerContract.CallOpts, _opts, _nonce)
}

// OperationHash is a free data retrieval call binding the contract method 0x80f12264.
//
// Solidity: function operationHash((address,uint8,uint64,bytes)[] _opts, uint256 _nonce) pure returns(bytes32 hash, bytes32 messageHash)
func (_VotingManagerContract *VotingManagerContractCallerSession) OperationHash(_opts []Operation, _nonce *big.Int) (struct {
	Hash        [32]byte
	MessageHash [32]byte
}, error) {
	return _VotingManagerContract.Contract.OperationHash(&_VotingManagerContract.CallOpts, _opts, _nonce)
}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) ParticipantManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "participantManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) ParticipantManager() (common.Address, error) {
	return _VotingManagerContract.Contract.ParticipantManager(&_VotingManagerContract.CallOpts)
}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) ParticipantManager() (common.Address, error) {
	return _VotingManagerContract.Contract.ParticipantManager(&_VotingManagerContract.CallOpts)
}

// TaskCompletionThreshold is a free data retrieval call binding the contract method 0xc6a1f28b.
//
// Solidity: function taskCompletionThreshold() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) TaskCompletionThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "taskCompletionThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCompletionThreshold is a free data retrieval call binding the contract method 0xc6a1f28b.
//
// Solidity: function taskCompletionThreshold() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) TaskCompletionThreshold() (*big.Int, error) {
	return _VotingManagerContract.Contract.TaskCompletionThreshold(&_VotingManagerContract.CallOpts)
}

// TaskCompletionThreshold is a free data retrieval call binding the contract method 0xc6a1f28b.
//
// Solidity: function taskCompletionThreshold() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) TaskCompletionThreshold() (*big.Int, error) {
	return _VotingManagerContract.Contract.TaskCompletionThreshold(&_VotingManagerContract.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) TaskManager() (common.Address, error) {
	return _VotingManagerContract.Contract.TaskManager(&_VotingManagerContract.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) TaskManager() (common.Address, error) {
	return _VotingManagerContract.Contract.TaskManager(&_VotingManagerContract.CallOpts)
}

// TssNonce is a free data retrieval call binding the contract method 0x3596d7eb.
//
// Solidity: function tssNonce() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) TssNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "tssNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TssNonce is a free data retrieval call binding the contract method 0x3596d7eb.
//
// Solidity: function tssNonce() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) TssNonce() (*big.Int, error) {
	return _VotingManagerContract.Contract.TssNonce(&_VotingManagerContract.CallOpts)
}

// TssNonce is a free data retrieval call binding the contract method 0x3596d7eb.
//
// Solidity: function tssNonce() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) TssNonce() (*big.Int, error) {
	return _VotingManagerContract.Contract.TssNonce(&_VotingManagerContract.CallOpts)
}

// TssSigner is a free data retrieval call binding the contract method 0xd5e2cb37.
//
// Solidity: function tssSigner() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) TssSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "tssSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssSigner is a free data retrieval call binding the contract method 0xd5e2cb37.
//
// Solidity: function tssSigner() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) TssSigner() (common.Address, error) {
	return _VotingManagerContract.Contract.TssSigner(&_VotingManagerContract.CallOpts)
}

// TssSigner is a free data retrieval call binding the contract method 0xd5e2cb37.
//
// Solidity: function tssSigner() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) TssSigner() (common.Address, error) {
	return _VotingManagerContract.Contract.TssSigner(&_VotingManagerContract.CallOpts)
}

// VerifyOperation is a free data retrieval call binding the contract method 0xf4706ec2.
//
// Solidity: function verifyOperation((address,uint8,uint64,bytes)[] _opts, bytes _signature, uint256 _nonce) view returns()
func (_VotingManagerContract *VotingManagerContractCaller) VerifyOperation(opts *bind.CallOpts, _opts []Operation, _signature []byte, _nonce *big.Int) error {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "verifyOperation", _opts, _signature, _nonce)

	if err != nil {
		return err
	}

	return err

}

// VerifyOperation is a free data retrieval call binding the contract method 0xf4706ec2.
//
// Solidity: function verifyOperation((address,uint8,uint64,bytes)[] _opts, bytes _signature, uint256 _nonce) view returns()
func (_VotingManagerContract *VotingManagerContractSession) VerifyOperation(_opts []Operation, _signature []byte, _nonce *big.Int) error {
	return _VotingManagerContract.Contract.VerifyOperation(&_VotingManagerContract.CallOpts, _opts, _signature, _nonce)
}

// VerifyOperation is a free data retrieval call binding the contract method 0xf4706ec2.
//
// Solidity: function verifyOperation((address,uint8,uint64,bytes)[] _opts, bytes _signature, uint256 _nonce) view returns()
func (_VotingManagerContract *VotingManagerContractCallerSession) VerifyOperation(_opts []Operation, _signature []byte, _nonce *big.Int) error {
	return _VotingManagerContract.Contract.VerifyOperation(&_VotingManagerContract.CallOpts, _opts, _signature, _nonce)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0xfe22e244.
//
// Solidity: function chooseNewSubmitter(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) ChooseNewSubmitter(opts *bind.TransactOpts, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "chooseNewSubmitter", _signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0xfe22e244.
//
// Solidity: function chooseNewSubmitter(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) ChooseNewSubmitter(_signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ChooseNewSubmitter(&_VotingManagerContract.TransactOpts, _signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0xfe22e244.
//
// Solidity: function chooseNewSubmitter(bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) ChooseNewSubmitter(_signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ChooseNewSubmitter(&_VotingManagerContract.TransactOpts, _signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _tssSigner, address _participantManager, address _taskManager, address _nuvoLock) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) Initialize(opts *bind.TransactOpts, _tssSigner common.Address, _participantManager common.Address, _taskManager common.Address, _nuvoLock common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "initialize", _tssSigner, _participantManager, _taskManager, _nuvoLock)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _tssSigner, address _participantManager, address _taskManager, address _nuvoLock) returns()
func (_VotingManagerContract *VotingManagerContractSession) Initialize(_tssSigner common.Address, _participantManager common.Address, _taskManager common.Address, _nuvoLock common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.Initialize(&_VotingManagerContract.TransactOpts, _tssSigner, _participantManager, _taskManager, _nuvoLock)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _tssSigner, address _participantManager, address _taskManager, address _nuvoLock) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) Initialize(_tssSigner common.Address, _participantManager common.Address, _taskManager common.Address, _nuvoLock common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.Initialize(&_VotingManagerContract.TransactOpts, _tssSigner, _participantManager, _taskManager, _nuvoLock)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x1b004ebb.
//
// Solidity: function setSignerAddress(address _newSigner, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SetSignerAddress(opts *bind.TransactOpts, _newSigner common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "setSignerAddress", _newSigner, _signature)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x1b004ebb.
//
// Solidity: function setSignerAddress(address _newSigner, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SetSignerAddress(_newSigner common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SetSignerAddress(&_VotingManagerContract.TransactOpts, _newSigner, _signature)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x1b004ebb.
//
// Solidity: function setSignerAddress(address _newSigner, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SetSignerAddress(_newSigner common.Address, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SetSignerAddress(&_VotingManagerContract.TransactOpts, _newSigner, _signature)
}

// VerifyAndCall is a paid mutator transaction binding the contract method 0x2da438c6.
//
// Solidity: function verifyAndCall((address,uint8,uint64,bytes)[] _opts, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) VerifyAndCall(opts *bind.TransactOpts, _opts []Operation, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "verifyAndCall", _opts, _signature)
}

// VerifyAndCall is a paid mutator transaction binding the contract method 0x2da438c6.
//
// Solidity: function verifyAndCall((address,uint8,uint64,bytes)[] _opts, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) VerifyAndCall(_opts []Operation, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.VerifyAndCall(&_VotingManagerContract.TransactOpts, _opts, _signature)
}

// VerifyAndCall is a paid mutator transaction binding the contract method 0x2da438c6.
//
// Solidity: function verifyAndCall((address,uint8,uint64,bytes)[] _opts, bytes _signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) VerifyAndCall(_opts []Operation, _signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.VerifyAndCall(&_VotingManagerContract.TransactOpts, _opts, _signature)
}

// VotingManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VotingManagerContract contract.
type VotingManagerContractInitializedIterator struct {
	Event *VotingManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractInitialized)
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
		it.Event = new(VotingManagerContractInitialized)
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
func (it *VotingManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractInitialized represents a Initialized event raised by the VotingManagerContract contract.
type VotingManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotingManagerContractInitializedIterator, error) {

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractInitializedIterator{contract: _VotingManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotingManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractInitialized)
				if err := _VotingManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VotingManagerContract *VotingManagerContractFilterer) ParseInitialized(log types.Log) (*VotingManagerContractInitialized, error) {
	event := new(VotingManagerContractInitialized)
	if err := _VotingManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractOperationFailedIterator is returned from FilterOperationFailed and is used to iterate over the raw logs and unpacked data for OperationFailed events raised by the VotingManagerContract contract.
type VotingManagerContractOperationFailedIterator struct {
	Event *VotingManagerContractOperationFailed // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractOperationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractOperationFailed)
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
		it.Event = new(VotingManagerContractOperationFailed)
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
func (it *VotingManagerContractOperationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractOperationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractOperationFailed represents a OperationFailed event raised by the VotingManagerContract contract.
type VotingManagerContractOperationFailed struct {
	ErrInfo common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperationFailed is a free log retrieval operation binding the contract event 0xd8653653db9841582c39f8003b7268dec862ced73491acb8db13ca0681a8ecc2.
//
// Solidity: event OperationFailed(bytes indexed errInfo)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterOperationFailed(opts *bind.FilterOpts, errInfo [][]byte) (*VotingManagerContractOperationFailedIterator, error) {

	var errInfoRule []interface{}
	for _, errInfoItem := range errInfo {
		errInfoRule = append(errInfoRule, errInfoItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "OperationFailed", errInfoRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractOperationFailedIterator{contract: _VotingManagerContract.contract, event: "OperationFailed", logs: logs, sub: sub}, nil
}

// WatchOperationFailed is a free log subscription operation binding the contract event 0xd8653653db9841582c39f8003b7268dec862ced73491acb8db13ca0681a8ecc2.
//
// Solidity: event OperationFailed(bytes indexed errInfo)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchOperationFailed(opts *bind.WatchOpts, sink chan<- *VotingManagerContractOperationFailed, errInfo [][]byte) (event.Subscription, error) {

	var errInfoRule []interface{}
	for _, errInfoItem := range errInfo {
		errInfoRule = append(errInfoRule, errInfoItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "OperationFailed", errInfoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractOperationFailed)
				if err := _VotingManagerContract.contract.UnpackLog(event, "OperationFailed", log); err != nil {
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

// ParseOperationFailed is a log parse operation binding the contract event 0xd8653653db9841582c39f8003b7268dec862ced73491acb8db13ca0681a8ecc2.
//
// Solidity: event OperationFailed(bytes indexed errInfo)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseOperationFailed(log types.Log) (*VotingManagerContractOperationFailed, error) {
	event := new(VotingManagerContractOperationFailed)
	if err := _VotingManagerContract.contract.UnpackLog(event, "OperationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractSubmitterChosenIterator is returned from FilterSubmitterChosen and is used to iterate over the raw logs and unpacked data for SubmitterChosen events raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterChosenIterator struct {
	Event *VotingManagerContractSubmitterChosen // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractSubmitterChosenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractSubmitterChosen)
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
		it.Event = new(VotingManagerContractSubmitterChosen)
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
func (it *VotingManagerContractSubmitterChosenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractSubmitterChosenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractSubmitterChosen represents a SubmitterChosen event raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterChosen struct {
	NewSubmitter common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubmitterChosen is a free log retrieval operation binding the contract event 0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d.
//
// Solidity: event SubmitterChosen(address indexed newSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterSubmitterChosen(opts *bind.FilterOpts, newSubmitter []common.Address) (*VotingManagerContractSubmitterChosenIterator, error) {

	var newSubmitterRule []interface{}
	for _, newSubmitterItem := range newSubmitter {
		newSubmitterRule = append(newSubmitterRule, newSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "SubmitterChosen", newSubmitterRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractSubmitterChosenIterator{contract: _VotingManagerContract.contract, event: "SubmitterChosen", logs: logs, sub: sub}, nil
}

// WatchSubmitterChosen is a free log subscription operation binding the contract event 0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d.
//
// Solidity: event SubmitterChosen(address indexed newSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchSubmitterChosen(opts *bind.WatchOpts, sink chan<- *VotingManagerContractSubmitterChosen, newSubmitter []common.Address) (event.Subscription, error) {

	var newSubmitterRule []interface{}
	for _, newSubmitterItem := range newSubmitter {
		newSubmitterRule = append(newSubmitterRule, newSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "SubmitterChosen", newSubmitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractSubmitterChosen)
				if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterChosen", log); err != nil {
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

// ParseSubmitterChosen is a log parse operation binding the contract event 0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d.
//
// Solidity: event SubmitterChosen(address indexed newSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseSubmitterChosen(log types.Log) (*VotingManagerContractSubmitterChosen, error) {
	event := new(VotingManagerContractSubmitterChosen)
	if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterChosen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractSubmitterRotationRequestedIterator is returned from FilterSubmitterRotationRequested and is used to iterate over the raw logs and unpacked data for SubmitterRotationRequested events raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterRotationRequestedIterator struct {
	Event *VotingManagerContractSubmitterRotationRequested // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractSubmitterRotationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractSubmitterRotationRequested)
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
		it.Event = new(VotingManagerContractSubmitterRotationRequested)
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
func (it *VotingManagerContractSubmitterRotationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractSubmitterRotationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractSubmitterRotationRequested represents a SubmitterRotationRequested event raised by the VotingManagerContract contract.
type VotingManagerContractSubmitterRotationRequested struct {
	Requester        common.Address
	CurrentSubmitter common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSubmitterRotationRequested is a free log retrieval operation binding the contract event 0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584.
//
// Solidity: event SubmitterRotationRequested(address indexed requester, address indexed currentSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterSubmitterRotationRequested(opts *bind.FilterOpts, requester []common.Address, currentSubmitter []common.Address) (*VotingManagerContractSubmitterRotationRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var currentSubmitterRule []interface{}
	for _, currentSubmitterItem := range currentSubmitter {
		currentSubmitterRule = append(currentSubmitterRule, currentSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "SubmitterRotationRequested", requesterRule, currentSubmitterRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractSubmitterRotationRequestedIterator{contract: _VotingManagerContract.contract, event: "SubmitterRotationRequested", logs: logs, sub: sub}, nil
}

// WatchSubmitterRotationRequested is a free log subscription operation binding the contract event 0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584.
//
// Solidity: event SubmitterRotationRequested(address indexed requester, address indexed currentSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchSubmitterRotationRequested(opts *bind.WatchOpts, sink chan<- *VotingManagerContractSubmitterRotationRequested, requester []common.Address, currentSubmitter []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var currentSubmitterRule []interface{}
	for _, currentSubmitterItem := range currentSubmitter {
		currentSubmitterRule = append(currentSubmitterRule, currentSubmitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "SubmitterRotationRequested", requesterRule, currentSubmitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractSubmitterRotationRequested)
				if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterRotationRequested", log); err != nil {
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

// ParseSubmitterRotationRequested is a log parse operation binding the contract event 0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584.
//
// Solidity: event SubmitterRotationRequested(address indexed requester, address indexed currentSubmitter)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseSubmitterRotationRequested(log types.Log) (*VotingManagerContractSubmitterRotationRequested, error) {
	event := new(VotingManagerContractSubmitterRotationRequested)
	if err := _VotingManagerContract.contract.UnpackLog(event, "SubmitterRotationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
