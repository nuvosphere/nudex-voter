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

// INuDexOperationsTask is an auto generated low-level Go binding around an user-defined struct.
type INuDexOperationsTask struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}

// NuDexOperationsContractMetaData contains all meta data concerning the NuDexOperationsContract contract.
var NuDexOperationsContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"confirmAllTasks\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestTask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structINuDexOperations.Task\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUncompletedTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINuDexOperations.Task[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_participantManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTaskCompleted\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markTaskCompleted\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markTaskCompleted_Batch\",\"inputs\":[{\"name\":\"_taskIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participantManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIParticipantManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmTask\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preconfirmedTaskResults\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmedTasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTask\",\"inputs\":[{\"name\":\"_context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyTask\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// NuDexOperationsContractABI is the input ABI used to generate the binding from.
// Deprecated: Use NuDexOperationsContractMetaData.ABI instead.
var NuDexOperationsContractABI = NuDexOperationsContractMetaData.ABI

// NuDexOperationsContract is an auto generated Go binding around an Ethereum contract.
type NuDexOperationsContract struct {
	NuDexOperationsContractCaller     // Read-only binding to the contract
	NuDexOperationsContractTransactor // Write-only binding to the contract
	NuDexOperationsContractFilterer   // Log filterer for contract events
}

// NuDexOperationsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NuDexOperationsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NuDexOperationsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NuDexOperationsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NuDexOperationsContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NuDexOperationsContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NuDexOperationsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NuDexOperationsContractSession struct {
	Contract     *NuDexOperationsContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NuDexOperationsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NuDexOperationsContractCallerSession struct {
	Contract *NuDexOperationsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// NuDexOperationsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NuDexOperationsContractTransactorSession struct {
	Contract     *NuDexOperationsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// NuDexOperationsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NuDexOperationsContractRaw struct {
	Contract *NuDexOperationsContract // Generic contract binding to access the raw methods on
}

// NuDexOperationsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NuDexOperationsContractCallerRaw struct {
	Contract *NuDexOperationsContractCaller // Generic read-only contract binding to access the raw methods on
}

// NuDexOperationsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NuDexOperationsContractTransactorRaw struct {
	Contract *NuDexOperationsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNuDexOperationsContract creates a new instance of NuDexOperationsContract, bound to a specific deployed contract.
func NewNuDexOperationsContract(address common.Address, backend bind.ContractBackend) (*NuDexOperationsContract, error) {
	contract, err := bindNuDexOperationsContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContract{NuDexOperationsContractCaller: NuDexOperationsContractCaller{contract: contract}, NuDexOperationsContractTransactor: NuDexOperationsContractTransactor{contract: contract}, NuDexOperationsContractFilterer: NuDexOperationsContractFilterer{contract: contract}}, nil
}

// NewNuDexOperationsContractCaller creates a new read-only instance of NuDexOperationsContract, bound to a specific deployed contract.
func NewNuDexOperationsContractCaller(address common.Address, caller bind.ContractCaller) (*NuDexOperationsContractCaller, error) {
	contract, err := bindNuDexOperationsContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractCaller{contract: contract}, nil
}

// NewNuDexOperationsContractTransactor creates a new write-only instance of NuDexOperationsContract, bound to a specific deployed contract.
func NewNuDexOperationsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NuDexOperationsContractTransactor, error) {
	contract, err := bindNuDexOperationsContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractTransactor{contract: contract}, nil
}

// NewNuDexOperationsContractFilterer creates a new log filterer instance of NuDexOperationsContract, bound to a specific deployed contract.
func NewNuDexOperationsContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NuDexOperationsContractFilterer, error) {
	contract, err := bindNuDexOperationsContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractFilterer{contract: contract}, nil
}

// bindNuDexOperationsContract binds a generic wrapper to an already deployed contract.
func bindNuDexOperationsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NuDexOperationsContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NuDexOperationsContract *NuDexOperationsContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NuDexOperationsContract.Contract.NuDexOperationsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NuDexOperationsContract *NuDexOperationsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.NuDexOperationsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NuDexOperationsContract *NuDexOperationsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.NuDexOperationsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NuDexOperationsContract *NuDexOperationsContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NuDexOperationsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NuDexOperationsContract *NuDexOperationsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NuDexOperationsContract *NuDexOperationsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.contract.Transact(opts, method, params...)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint256,bytes,address,bool,uint256,uint256,bytes))
func (_NuDexOperationsContract *NuDexOperationsContractCaller) GetLatestTask(opts *bind.CallOpts) (INuDexOperationsTask, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "getLatestTask")

	if err != nil {
		return *new(INuDexOperationsTask), err
	}

	out0 := *abi.ConvertType(out[0], new(INuDexOperationsTask)).(*INuDexOperationsTask)

	return out0, err

}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint256,bytes,address,bool,uint256,uint256,bytes))
func (_NuDexOperationsContract *NuDexOperationsContractSession) GetLatestTask() (INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetLatestTask(&_NuDexOperationsContract.CallOpts)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint256,bytes,address,bool,uint256,uint256,bytes))
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) GetLatestTask() (INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetLatestTask(&_NuDexOperationsContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,bytes,address,bool,uint256,uint256,bytes)[])
func (_NuDexOperationsContract *NuDexOperationsContractCaller) GetUncompletedTasks(opts *bind.CallOpts) ([]INuDexOperationsTask, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "getUncompletedTasks")

	if err != nil {
		return *new([]INuDexOperationsTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]INuDexOperationsTask)).(*[]INuDexOperationsTask)

	return out0, err

}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,bytes,address,bool,uint256,uint256,bytes)[])
func (_NuDexOperationsContract *NuDexOperationsContractSession) GetUncompletedTasks() ([]INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetUncompletedTasks(&_NuDexOperationsContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,bytes,address,bool,uint256,uint256,bytes)[])
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) GetUncompletedTasks() ([]INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetUncompletedTasks(&_NuDexOperationsContract.CallOpts)
}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 _taskId) view returns(bool)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) IsTaskCompleted(opts *bind.CallOpts, _taskId *big.Int) (bool, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "isTaskCompleted", _taskId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 _taskId) view returns(bool)
func (_NuDexOperationsContract *NuDexOperationsContractSession) IsTaskCompleted(_taskId *big.Int) (bool, error) {
	return _NuDexOperationsContract.Contract.IsTaskCompleted(&_NuDexOperationsContract.CallOpts, _taskId)
}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 _taskId) view returns(bool)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) IsTaskCompleted(_taskId *big.Int) (bool, error) {
	return _NuDexOperationsContract.Contract.IsTaskCompleted(&_NuDexOperationsContract.CallOpts, _taskId)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint256)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) NextTaskId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "nextTaskId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint256)
func (_NuDexOperationsContract *NuDexOperationsContractSession) NextTaskId() (*big.Int, error) {
	return _NuDexOperationsContract.Contract.NextTaskId(&_NuDexOperationsContract.CallOpts)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint256)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) NextTaskId() (*big.Int, error) {
	return _NuDexOperationsContract.Contract.NextTaskId(&_NuDexOperationsContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NuDexOperationsContract *NuDexOperationsContractSession) Owner() (common.Address, error) {
	return _NuDexOperationsContract.Contract.Owner(&_NuDexOperationsContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) Owner() (common.Address, error) {
	return _NuDexOperationsContract.Contract.Owner(&_NuDexOperationsContract.CallOpts)
}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) ParticipantManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "participantManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_NuDexOperationsContract *NuDexOperationsContractSession) ParticipantManager() (common.Address, error) {
	return _NuDexOperationsContract.Contract.ParticipantManager(&_NuDexOperationsContract.CallOpts)
}

// ParticipantManager is a free data retrieval call binding the contract method 0x464dbe6e.
//
// Solidity: function participantManager() view returns(address)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) ParticipantManager() (common.Address, error) {
	return _NuDexOperationsContract.Contract.ParticipantManager(&_NuDexOperationsContract.CallOpts)
}

// PreconfirmedTaskResults is a free data retrieval call binding the contract method 0xc1e6dfeb.
//
// Solidity: function preconfirmedTaskResults(uint256 ) view returns(bytes)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) PreconfirmedTaskResults(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "preconfirmedTaskResults", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PreconfirmedTaskResults is a free data retrieval call binding the contract method 0xc1e6dfeb.
//
// Solidity: function preconfirmedTaskResults(uint256 ) view returns(bytes)
func (_NuDexOperationsContract *NuDexOperationsContractSession) PreconfirmedTaskResults(arg0 *big.Int) ([]byte, error) {
	return _NuDexOperationsContract.Contract.PreconfirmedTaskResults(&_NuDexOperationsContract.CallOpts, arg0)
}

// PreconfirmedTaskResults is a free data retrieval call binding the contract method 0xc1e6dfeb.
//
// Solidity: function preconfirmedTaskResults(uint256 ) view returns(bytes)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) PreconfirmedTaskResults(arg0 *big.Int) ([]byte, error) {
	return _NuDexOperationsContract.Contract.PreconfirmedTaskResults(&_NuDexOperationsContract.CallOpts, arg0)
}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint256)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) PreconfirmedTasks(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "preconfirmedTasks", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint256)
func (_NuDexOperationsContract *NuDexOperationsContractSession) PreconfirmedTasks(arg0 *big.Int) (*big.Int, error) {
	return _NuDexOperationsContract.Contract.PreconfirmedTasks(&_NuDexOperationsContract.CallOpts, arg0)
}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint256)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) PreconfirmedTasks(arg0 *big.Int) (*big.Int, error) {
	return _NuDexOperationsContract.Contract.PreconfirmedTasks(&_NuDexOperationsContract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, bytes context, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Context     []byte
		Submitter   common.Address
		IsCompleted bool
		CreatedAt   *big.Int
		CompletedAt *big.Int
		Result      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Context = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Submitter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsCompleted = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.CreatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CompletedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, bytes context, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	return _NuDexOperationsContract.Contract.Tasks(&_NuDexOperationsContract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, bytes context, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	return _NuDexOperationsContract.Contract.Tasks(&_NuDexOperationsContract.CallOpts, arg0)
}

// ConfirmAllTasks is a paid mutator transaction binding the contract method 0xa7ec744e.
//
// Solidity: function confirmAllTasks() returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) ConfirmAllTasks(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "confirmAllTasks")
}

// ConfirmAllTasks is a paid mutator transaction binding the contract method 0xa7ec744e.
//
// Solidity: function confirmAllTasks() returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) ConfirmAllTasks() (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.ConfirmAllTasks(&_NuDexOperationsContract.TransactOpts)
}

// ConfirmAllTasks is a paid mutator transaction binding the contract method 0xa7ec744e.
//
// Solidity: function confirmAllTasks() returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) ConfirmAllTasks() (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.ConfirmAllTasks(&_NuDexOperationsContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _participantManager, address _owner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) Initialize(opts *bind.TransactOpts, _participantManager common.Address, _owner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "initialize", _participantManager, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _participantManager, address _owner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) Initialize(_participantManager common.Address, _owner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.Initialize(&_NuDexOperationsContract.TransactOpts, _participantManager, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _participantManager, address _owner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) Initialize(_participantManager common.Address, _owner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.Initialize(&_NuDexOperationsContract.TransactOpts, _participantManager, _owner)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 _taskId, bytes _result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) MarkTaskCompleted(opts *bind.TransactOpts, _taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "markTaskCompleted", _taskId, _result)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 _taskId, bytes _result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) MarkTaskCompleted(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.MarkTaskCompleted(&_NuDexOperationsContract.TransactOpts, _taskId, _result)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 _taskId, bytes _result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) MarkTaskCompleted(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.MarkTaskCompleted(&_NuDexOperationsContract.TransactOpts, _taskId, _result)
}

// MarkTaskCompletedBatch is a paid mutator transaction binding the contract method 0xdf2f7649.
//
// Solidity: function markTaskCompleted_Batch(uint256[] _taskIds, bytes[] _results) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) MarkTaskCompletedBatch(opts *bind.TransactOpts, _taskIds []*big.Int, _results [][]byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "markTaskCompleted_Batch", _taskIds, _results)
}

// MarkTaskCompletedBatch is a paid mutator transaction binding the contract method 0xdf2f7649.
//
// Solidity: function markTaskCompleted_Batch(uint256[] _taskIds, bytes[] _results) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) MarkTaskCompletedBatch(_taskIds []*big.Int, _results [][]byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.MarkTaskCompletedBatch(&_NuDexOperationsContract.TransactOpts, _taskIds, _results)
}

// MarkTaskCompletedBatch is a paid mutator transaction binding the contract method 0xdf2f7649.
//
// Solidity: function markTaskCompleted_Batch(uint256[] _taskIds, bytes[] _results) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) MarkTaskCompletedBatch(_taskIds []*big.Int, _results [][]byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.MarkTaskCompletedBatch(&_NuDexOperationsContract.TransactOpts, _taskIds, _results)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x84302728.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) PreconfirmTask(opts *bind.TransactOpts, _taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "preconfirmTask", _taskId, _result)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x84302728.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) PreconfirmTask(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.PreconfirmTask(&_NuDexOperationsContract.TransactOpts, _taskId, _result)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x84302728.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) PreconfirmTask(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.PreconfirmTask(&_NuDexOperationsContract.TransactOpts, _taskId, _result)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.RenounceOwnership(&_NuDexOperationsContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.RenounceOwnership(&_NuDexOperationsContract.TransactOpts)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x5e0c6100.
//
// Solidity: function submitTask(bytes _context) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) SubmitTask(opts *bind.TransactOpts, _context []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "submitTask", _context)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x5e0c6100.
//
// Solidity: function submitTask(bytes _context) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) SubmitTask(_context []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.SubmitTask(&_NuDexOperationsContract.TransactOpts, _context)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x5e0c6100.
//
// Solidity: function submitTask(bytes _context) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) SubmitTask(_context []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.SubmitTask(&_NuDexOperationsContract.TransactOpts, _context)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.TransferOwnership(&_NuDexOperationsContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.TransferOwnership(&_NuDexOperationsContract.TransactOpts, newOwner)
}

// NuDexOperationsContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NuDexOperationsContract contract.
type NuDexOperationsContractInitializedIterator struct {
	Event *NuDexOperationsContractInitialized // Event containing the contract specifics and raw log

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
func (it *NuDexOperationsContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NuDexOperationsContractInitialized)
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
		it.Event = new(NuDexOperationsContractInitialized)
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
func (it *NuDexOperationsContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NuDexOperationsContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NuDexOperationsContractInitialized represents a Initialized event raised by the NuDexOperationsContract contract.
type NuDexOperationsContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*NuDexOperationsContractInitializedIterator, error) {

	logs, sub, err := _NuDexOperationsContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractInitializedIterator{contract: _NuDexOperationsContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NuDexOperationsContractInitialized) (event.Subscription, error) {

	logs, sub, err := _NuDexOperationsContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NuDexOperationsContractInitialized)
				if err := _NuDexOperationsContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) ParseInitialized(log types.Log) (*NuDexOperationsContractInitialized, error) {
	event := new(NuDexOperationsContractInitialized)
	if err := _NuDexOperationsContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NuDexOperationsContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NuDexOperationsContract contract.
type NuDexOperationsContractOwnershipTransferredIterator struct {
	Event *NuDexOperationsContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NuDexOperationsContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NuDexOperationsContractOwnershipTransferred)
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
		it.Event = new(NuDexOperationsContractOwnershipTransferred)
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
func (it *NuDexOperationsContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NuDexOperationsContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NuDexOperationsContractOwnershipTransferred represents a OwnershipTransferred event raised by the NuDexOperationsContract contract.
type NuDexOperationsContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NuDexOperationsContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractOwnershipTransferredIterator{contract: _NuDexOperationsContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NuDexOperationsContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NuDexOperationsContractOwnershipTransferred)
				if err := _NuDexOperationsContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) ParseOwnershipTransferred(log types.Log) (*NuDexOperationsContractOwnershipTransferred, error) {
	event := new(NuDexOperationsContractOwnershipTransferred)
	if err := _NuDexOperationsContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NuDexOperationsContractTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the NuDexOperationsContract contract.
type NuDexOperationsContractTaskCompletedIterator struct {
	Event *NuDexOperationsContractTaskCompleted // Event containing the contract specifics and raw log

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
func (it *NuDexOperationsContractTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NuDexOperationsContractTaskCompleted)
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
		it.Event = new(NuDexOperationsContractTaskCompleted)
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
func (it *NuDexOperationsContractTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NuDexOperationsContractTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NuDexOperationsContractTaskCompleted represents a TaskCompleted event raised by the NuDexOperationsContract contract.
type NuDexOperationsContractTaskCompleted struct {
	TaskId      *big.Int
	Submitter   common.Address
	CompletedAt *big.Int
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 indexed completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int, submitter []common.Address, completedAt []*big.Int) (*NuDexOperationsContractTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var completedAtRule []interface{}
	for _, completedAtItem := range completedAt {
		completedAtRule = append(completedAtRule, completedAtItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.FilterLogs(opts, "TaskCompleted", taskIdRule, submitterRule, completedAtRule)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractTaskCompletedIterator{contract: _NuDexOperationsContract.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 indexed completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *NuDexOperationsContractTaskCompleted, taskId []*big.Int, submitter []common.Address, completedAt []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var completedAtRule []interface{}
	for _, completedAtItem := range completedAt {
		completedAtRule = append(completedAtRule, completedAtItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.WatchLogs(opts, "TaskCompleted", taskIdRule, submitterRule, completedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NuDexOperationsContractTaskCompleted)
				if err := _NuDexOperationsContract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 indexed completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) ParseTaskCompleted(log types.Log) (*NuDexOperationsContractTaskCompleted, error) {
	event := new(NuDexOperationsContractTaskCompleted)
	if err := _NuDexOperationsContract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NuDexOperationsContractTaskSubmittedIterator is returned from FilterTaskSubmitted and is used to iterate over the raw logs and unpacked data for TaskSubmitted events raised by the NuDexOperationsContract contract.
type NuDexOperationsContractTaskSubmittedIterator struct {
	Event *NuDexOperationsContractTaskSubmitted // Event containing the contract specifics and raw log

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
func (it *NuDexOperationsContractTaskSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NuDexOperationsContractTaskSubmitted)
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
		it.Event = new(NuDexOperationsContractTaskSubmitted)
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
func (it *NuDexOperationsContractTaskSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NuDexOperationsContractTaskSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NuDexOperationsContractTaskSubmitted represents a TaskSubmitted event raised by the NuDexOperationsContract contract.
type NuDexOperationsContractTaskSubmitted struct {
	TaskId    *big.Int
	Context   []byte
	Submitter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, bytes context, address indexed submitter)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, taskId []*big.Int, submitter []common.Address) (*NuDexOperationsContractTaskSubmittedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.FilterLogs(opts, "TaskSubmitted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractTaskSubmittedIterator{contract: _NuDexOperationsContract.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, bytes context, address indexed submitter)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *NuDexOperationsContractTaskSubmitted, taskId []*big.Int, submitter []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.WatchLogs(opts, "TaskSubmitted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NuDexOperationsContractTaskSubmitted)
				if err := _NuDexOperationsContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, bytes context, address indexed submitter)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) ParseTaskSubmitted(log types.Log) (*NuDexOperationsContractTaskSubmitted, error) {
	event := new(NuDexOperationsContractTaskSubmitted)
	if err := _NuDexOperationsContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
