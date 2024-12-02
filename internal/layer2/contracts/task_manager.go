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

// ITaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type ITaskManagerTask struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	CreatedAt *big.Int
	UpdatedAt *big.Int
	Context   []byte
	Result    []byte
}

// TaskManagerContractMetaData contains all meta data concerning the TaskManagerContract contract.
var TaskManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getLatestTask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskManager.Task\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskState\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUncompletedTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structITaskManager.Task[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_taskSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingTaskIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingTasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmedTaskResults\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmedTasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTaskSubmitter\",\"inputs\":[{\"name\":\"_taskSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTask\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskRecords\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTask\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_state\",\"type\":\"uint8\",\"internalType\":\"enumState\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"context\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskUpdated\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"updateTime\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumState\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyExistTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EmptyTask\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTaskSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// TaskManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskManagerContractMetaData.ABI instead.
var TaskManagerContractABI = TaskManagerContractMetaData.ABI

// TaskManagerContract is an auto generated Go binding around an Ethereum contract.
type TaskManagerContract struct {
	TaskManagerContractCaller     // Read-only binding to the contract
	TaskManagerContractTransactor // Write-only binding to the contract
	TaskManagerContractFilterer   // Log filterer for contract events
}

// TaskManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskManagerContractSession struct {
	Contract     *TaskManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaskManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskManagerContractCallerSession struct {
	Contract *TaskManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TaskManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskManagerContractTransactorSession struct {
	Contract     *TaskManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TaskManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskManagerContractRaw struct {
	Contract *TaskManagerContract // Generic contract binding to access the raw methods on
}

// TaskManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskManagerContractCallerRaw struct {
	Contract *TaskManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// TaskManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskManagerContractTransactorRaw struct {
	Contract *TaskManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskManagerContract creates a new instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContract(address common.Address, backend bind.ContractBackend) (*TaskManagerContract, error) {
	contract, err := bindTaskManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContract{TaskManagerContractCaller: TaskManagerContractCaller{contract: contract}, TaskManagerContractTransactor: TaskManagerContractTransactor{contract: contract}, TaskManagerContractFilterer: TaskManagerContractFilterer{contract: contract}}, nil
}

// NewTaskManagerContractCaller creates a new read-only instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContractCaller(address common.Address, caller bind.ContractCaller) (*TaskManagerContractCaller, error) {
	contract, err := bindTaskManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractCaller{contract: contract}, nil
}

// NewTaskManagerContractTransactor creates a new write-only instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskManagerContractTransactor, error) {
	contract, err := bindTaskManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTransactor{contract: contract}, nil
}

// NewTaskManagerContractFilterer creates a new log filterer instance of TaskManagerContract, bound to a specific deployed contract.
func NewTaskManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskManagerContractFilterer, error) {
	contract, err := bindTaskManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractFilterer{contract: contract}, nil
}

// bindTaskManagerContract binds a generic wrapper to an already deployed contract.
func bindTaskManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskManagerContract *TaskManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskManagerContract.Contract.TaskManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskManagerContract *TaskManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.TaskManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskManagerContract *TaskManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.TaskManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskManagerContract *TaskManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskManagerContract *TaskManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskManagerContract *TaskManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.contract.Transact(opts, method, params...)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint64,uint8,address,uint256,uint256,bytes,bytes))
func (_TaskManagerContract *TaskManagerContractCaller) GetLatestTask(opts *bind.CallOpts) (ITaskManagerTask, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getLatestTask")

	if err != nil {
		return *new(ITaskManagerTask), err
	}

	out0 := *abi.ConvertType(out[0], new(ITaskManagerTask)).(*ITaskManagerTask)

	return out0, err

}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint64,uint8,address,uint256,uint256,bytes,bytes))
func (_TaskManagerContract *TaskManagerContractSession) GetLatestTask() (ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetLatestTask(&_TaskManagerContract.CallOpts)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint64,uint8,address,uint256,uint256,bytes,bytes))
func (_TaskManagerContract *TaskManagerContractCallerSession) GetLatestTask() (ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetLatestTask(&_TaskManagerContract.CallOpts)
}

// GetTaskState is a free data retrieval call binding the contract method 0xffcfc50a.
//
// Solidity: function getTaskState(uint64 _taskId) view returns(uint8)
func (_TaskManagerContract *TaskManagerContractCaller) GetTaskState(opts *bind.CallOpts, _taskId uint64) (uint8, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getTaskState", _taskId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTaskState is a free data retrieval call binding the contract method 0xffcfc50a.
//
// Solidity: function getTaskState(uint64 _taskId) view returns(uint8)
func (_TaskManagerContract *TaskManagerContractSession) GetTaskState(_taskId uint64) (uint8, error) {
	return _TaskManagerContract.Contract.GetTaskState(&_TaskManagerContract.CallOpts, _taskId)
}

// GetTaskState is a free data retrieval call binding the contract method 0xffcfc50a.
//
// Solidity: function getTaskState(uint64 _taskId) view returns(uint8)
func (_TaskManagerContract *TaskManagerContractCallerSession) GetTaskState(_taskId uint64) (uint8, error) {
	return _TaskManagerContract.Contract.GetTaskState(&_TaskManagerContract.CallOpts, _taskId)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint64,uint8,address,uint256,uint256,bytes,bytes)[])
func (_TaskManagerContract *TaskManagerContractCaller) GetUncompletedTasks(opts *bind.CallOpts) ([]ITaskManagerTask, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "getUncompletedTasks")

	if err != nil {
		return *new([]ITaskManagerTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITaskManagerTask)).(*[]ITaskManagerTask)

	return out0, err

}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint64,uint8,address,uint256,uint256,bytes,bytes)[])
func (_TaskManagerContract *TaskManagerContractSession) GetUncompletedTasks() ([]ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetUncompletedTasks(&_TaskManagerContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint64,uint8,address,uint256,uint256,bytes,bytes)[])
func (_TaskManagerContract *TaskManagerContractCallerSession) GetUncompletedTasks() ([]ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetUncompletedTasks(&_TaskManagerContract.CallOpts)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) NextTaskId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "nextTaskId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) NextTaskId() (uint64, error) {
	return _TaskManagerContract.Contract.NextTaskId(&_TaskManagerContract.CallOpts)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) NextTaskId() (uint64, error) {
	return _TaskManagerContract.Contract.NextTaskId(&_TaskManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskManagerContract *TaskManagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskManagerContract *TaskManagerContractSession) Owner() (common.Address, error) {
	return _TaskManagerContract.Contract.Owner(&_TaskManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskManagerContract *TaskManagerContractCallerSession) Owner() (common.Address, error) {
	return _TaskManagerContract.Contract.Owner(&_TaskManagerContract.CallOpts)
}

// PendingTaskIndex is a free data retrieval call binding the contract method 0xd9804297.
//
// Solidity: function pendingTaskIndex() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) PendingTaskIndex(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "pendingTaskIndex")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingTaskIndex is a free data retrieval call binding the contract method 0xd9804297.
//
// Solidity: function pendingTaskIndex() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) PendingTaskIndex() (uint64, error) {
	return _TaskManagerContract.Contract.PendingTaskIndex(&_TaskManagerContract.CallOpts)
}

// PendingTaskIndex is a free data retrieval call binding the contract method 0xd9804297.
//
// Solidity: function pendingTaskIndex() view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) PendingTaskIndex() (uint64, error) {
	return _TaskManagerContract.Contract.PendingTaskIndex(&_TaskManagerContract.CallOpts)
}

// PendingTasks is a free data retrieval call binding the contract method 0x42260541.
//
// Solidity: function pendingTasks(uint256 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) PendingTasks(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "pendingTasks", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingTasks is a free data retrieval call binding the contract method 0x42260541.
//
// Solidity: function pendingTasks(uint256 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) PendingTasks(arg0 *big.Int) (uint64, error) {
	return _TaskManagerContract.Contract.PendingTasks(&_TaskManagerContract.CallOpts, arg0)
}

// PendingTasks is a free data retrieval call binding the contract method 0x42260541.
//
// Solidity: function pendingTasks(uint256 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) PendingTasks(arg0 *big.Int) (uint64, error) {
	return _TaskManagerContract.Contract.PendingTasks(&_TaskManagerContract.CallOpts, arg0)
}

// PreconfirmedTaskResults is a free data retrieval call binding the contract method 0xc1e6dfeb.
//
// Solidity: function preconfirmedTaskResults(uint256 ) view returns(bytes)
func (_TaskManagerContract *TaskManagerContractCaller) PreconfirmedTaskResults(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "preconfirmedTaskResults", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PreconfirmedTaskResults is a free data retrieval call binding the contract method 0xc1e6dfeb.
//
// Solidity: function preconfirmedTaskResults(uint256 ) view returns(bytes)
func (_TaskManagerContract *TaskManagerContractSession) PreconfirmedTaskResults(arg0 *big.Int) ([]byte, error) {
	return _TaskManagerContract.Contract.PreconfirmedTaskResults(&_TaskManagerContract.CallOpts, arg0)
}

// PreconfirmedTaskResults is a free data retrieval call binding the contract method 0xc1e6dfeb.
//
// Solidity: function preconfirmedTaskResults(uint256 ) view returns(bytes)
func (_TaskManagerContract *TaskManagerContractCallerSession) PreconfirmedTaskResults(arg0 *big.Int) ([]byte, error) {
	return _TaskManagerContract.Contract.PreconfirmedTaskResults(&_TaskManagerContract.CallOpts, arg0)
}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) PreconfirmedTasks(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "preconfirmedTasks", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) PreconfirmedTasks(arg0 *big.Int) (uint64, error) {
	return _TaskManagerContract.Contract.PreconfirmedTasks(&_TaskManagerContract.CallOpts, arg0)
}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) PreconfirmedTasks(arg0 *big.Int) (uint64, error) {
	return _TaskManagerContract.Contract.PreconfirmedTasks(&_TaskManagerContract.CallOpts, arg0)
}

// TaskRecords is a free data retrieval call binding the contract method 0xdf9431ee.
//
// Solidity: function taskRecords(bytes32 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCaller) TaskRecords(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "taskRecords", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TaskRecords is a free data retrieval call binding the contract method 0xdf9431ee.
//
// Solidity: function taskRecords(bytes32 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) TaskRecords(arg0 [32]byte) (uint64, error) {
	return _TaskManagerContract.Contract.TaskRecords(&_TaskManagerContract.CallOpts, arg0)
}

// TaskRecords is a free data retrieval call binding the contract method 0xdf9431ee.
//
// Solidity: function taskRecords(bytes32 ) view returns(uint64)
func (_TaskManagerContract *TaskManagerContractCallerSession) TaskRecords(arg0 [32]byte) (uint64, error) {
	return _TaskManagerContract.Contract.TaskRecords(&_TaskManagerContract.CallOpts, arg0)
}

// TaskSubmitter is a free data retrieval call binding the contract method 0xcda20e13.
//
// Solidity: function taskSubmitter() view returns(address)
func (_TaskManagerContract *TaskManagerContractCaller) TaskSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "taskSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskSubmitter is a free data retrieval call binding the contract method 0xcda20e13.
//
// Solidity: function taskSubmitter() view returns(address)
func (_TaskManagerContract *TaskManagerContractSession) TaskSubmitter() (common.Address, error) {
	return _TaskManagerContract.Contract.TaskSubmitter(&_TaskManagerContract.CallOpts)
}

// TaskSubmitter is a free data retrieval call binding the contract method 0xcda20e13.
//
// Solidity: function taskSubmitter() view returns(address)
func (_TaskManagerContract *TaskManagerContractCallerSession) TaskSubmitter() (common.Address, error) {
	return _TaskManagerContract.Contract.TaskSubmitter(&_TaskManagerContract.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x6b895ab7.
//
// Solidity: function tasks(uint64 ) view returns(uint64 id, uint8 state, address submitter, uint256 createdAt, uint256 updatedAt, bytes context, bytes result)
func (_TaskManagerContract *TaskManagerContractCaller) Tasks(opts *bind.CallOpts, arg0 uint64) (struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	CreatedAt *big.Int
	UpdatedAt *big.Int
	Context   []byte
	Result    []byte
}, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id        uint64
		State     uint8
		Submitter common.Address
		CreatedAt *big.Int
		UpdatedAt *big.Int
		Context   []byte
		Result    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.State = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Submitter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Context = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.Result = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x6b895ab7.
//
// Solidity: function tasks(uint64 ) view returns(uint64 id, uint8 state, address submitter, uint256 createdAt, uint256 updatedAt, bytes context, bytes result)
func (_TaskManagerContract *TaskManagerContractSession) Tasks(arg0 uint64) (struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	CreatedAt *big.Int
	UpdatedAt *big.Int
	Context   []byte
	Result    []byte
}, error) {
	return _TaskManagerContract.Contract.Tasks(&_TaskManagerContract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x6b895ab7.
//
// Solidity: function tasks(uint64 ) view returns(uint64 id, uint8 state, address submitter, uint256 createdAt, uint256 updatedAt, bytes context, bytes result)
func (_TaskManagerContract *TaskManagerContractCallerSession) Tasks(arg0 uint64) (struct {
	Id        uint64
	State     uint8
	Submitter common.Address
	CreatedAt *big.Int
	UpdatedAt *big.Int
	Context   []byte
	Result    []byte
}, error) {
	return _TaskManagerContract.Contract.Tasks(&_TaskManagerContract.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _taskSubmitter, address _owner) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) Initialize(opts *bind.TransactOpts, _taskSubmitter common.Address, _owner common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "initialize", _taskSubmitter, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _taskSubmitter, address _owner) returns()
func (_TaskManagerContract *TaskManagerContractSession) Initialize(_taskSubmitter common.Address, _owner common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.Initialize(&_TaskManagerContract.TransactOpts, _taskSubmitter, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _taskSubmitter, address _owner) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) Initialize(_taskSubmitter common.Address, _owner common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.Initialize(&_TaskManagerContract.TransactOpts, _taskSubmitter, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskManagerContract *TaskManagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskManagerContract *TaskManagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskManagerContract.Contract.RenounceOwnership(&_TaskManagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskManagerContract.Contract.RenounceOwnership(&_TaskManagerContract.TransactOpts)
}

// SetTaskSubmitter is a paid mutator transaction binding the contract method 0xc72de714.
//
// Solidity: function setTaskSubmitter(address _taskSubmitter) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) SetTaskSubmitter(opts *bind.TransactOpts, _taskSubmitter common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "setTaskSubmitter", _taskSubmitter)
}

// SetTaskSubmitter is a paid mutator transaction binding the contract method 0xc72de714.
//
// Solidity: function setTaskSubmitter(address _taskSubmitter) returns()
func (_TaskManagerContract *TaskManagerContractSession) SetTaskSubmitter(_taskSubmitter common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SetTaskSubmitter(&_TaskManagerContract.TransactOpts, _taskSubmitter)
}

// SetTaskSubmitter is a paid mutator transaction binding the contract method 0xc72de714.
//
// Solidity: function setTaskSubmitter(address _taskSubmitter) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) SetTaskSubmitter(_taskSubmitter common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SetTaskSubmitter(&_TaskManagerContract.TransactOpts, _taskSubmitter)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _context) returns(uint64)
func (_TaskManagerContract *TaskManagerContractTransactor) SubmitTask(opts *bind.TransactOpts, _submitter common.Address, _context []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "submitTask", _submitter, _context)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _context) returns(uint64)
func (_TaskManagerContract *TaskManagerContractSession) SubmitTask(_submitter common.Address, _context []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SubmitTask(&_TaskManagerContract.TransactOpts, _submitter, _context)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _context) returns(uint64)
func (_TaskManagerContract *TaskManagerContractTransactorSession) SubmitTask(_submitter common.Address, _context []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SubmitTask(&_TaskManagerContract.TransactOpts, _submitter, _context)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskManagerContract *TaskManagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.TransferOwnership(&_TaskManagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.TransferOwnership(&_TaskManagerContract.TransactOpts, newOwner)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x875d4e7d.
//
// Solidity: function updateTask(uint64 _taskId, uint8 _state, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) UpdateTask(opts *bind.TransactOpts, _taskId uint64, _state uint8, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "updateTask", _taskId, _state, _result)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x875d4e7d.
//
// Solidity: function updateTask(uint64 _taskId, uint8 _state, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractSession) UpdateTask(_taskId uint64, _state uint8, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.UpdateTask(&_TaskManagerContract.TransactOpts, _taskId, _state, _result)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x875d4e7d.
//
// Solidity: function updateTask(uint64 _taskId, uint8 _state, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) UpdateTask(_taskId uint64, _state uint8, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.UpdateTask(&_TaskManagerContract.TransactOpts, _taskId, _state, _result)
}

// TaskManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaskManagerContract contract.
type TaskManagerContractInitializedIterator struct {
	Event *TaskManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractInitialized)
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
		it.Event = new(TaskManagerContractInitialized)
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
func (it *TaskManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractInitialized represents a Initialized event raised by the TaskManagerContract contract.
type TaskManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskManagerContractInitializedIterator, error) {

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractInitializedIterator{contract: _TaskManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractInitialized)
				if err := _TaskManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaskManagerContract *TaskManagerContractFilterer) ParseInitialized(log types.Log) (*TaskManagerContractInitialized, error) {
	event := new(TaskManagerContractInitialized)
	if err := _TaskManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaskManagerContract contract.
type TaskManagerContractOwnershipTransferredIterator struct {
	Event *TaskManagerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractOwnershipTransferred)
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
		it.Event = new(TaskManagerContractOwnershipTransferred)
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
func (it *TaskManagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the TaskManagerContract contract.
type TaskManagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskManagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractOwnershipTransferredIterator{contract: _TaskManagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskManagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractOwnershipTransferred)
				if err := _TaskManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaskManagerContract *TaskManagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*TaskManagerContractOwnershipTransferred, error) {
	event := new(TaskManagerContractOwnershipTransferred)
	if err := _TaskManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractTaskSubmittedIterator is returned from FilterTaskSubmitted and is used to iterate over the raw logs and unpacked data for TaskSubmitted events raised by the TaskManagerContract contract.
type TaskManagerContractTaskSubmittedIterator struct {
	Event *TaskManagerContractTaskSubmitted // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractTaskSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractTaskSubmitted)
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
		it.Event = new(TaskManagerContractTaskSubmitted)
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
func (it *TaskManagerContractTaskSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractTaskSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractTaskSubmitted represents a TaskSubmitted event raised by the TaskManagerContract contract.
type TaskManagerContractTaskSubmitted struct {
	TaskId    uint64
	Context   []byte
	Submitter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x96a9964cc016aebed6b4922209fc0404309ba528357eec22795575c934c6f0d7.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, bytes context, address indexed submitter)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, taskId []uint64, submitter []common.Address) (*TaskManagerContractTaskSubmittedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "TaskSubmitted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTaskSubmittedIterator{contract: _TaskManagerContract.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0x96a9964cc016aebed6b4922209fc0404309ba528357eec22795575c934c6f0d7.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, bytes context, address indexed submitter)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *TaskManagerContractTaskSubmitted, taskId []uint64, submitter []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "TaskSubmitted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractTaskSubmitted)
				if err := _TaskManagerContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0x96a9964cc016aebed6b4922209fc0404309ba528357eec22795575c934c6f0d7.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, bytes context, address indexed submitter)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseTaskSubmitted(log types.Log) (*TaskManagerContractTaskSubmitted, error) {
	event := new(TaskManagerContractTaskSubmitted)
	if err := _TaskManagerContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskManagerContractTaskUpdatedIterator is returned from FilterTaskUpdated and is used to iterate over the raw logs and unpacked data for TaskUpdated events raised by the TaskManagerContract contract.
type TaskManagerContractTaskUpdatedIterator struct {
	Event *TaskManagerContractTaskUpdated // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractTaskUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractTaskUpdated)
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
		it.Event = new(TaskManagerContractTaskUpdated)
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
func (it *TaskManagerContractTaskUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractTaskUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractTaskUpdated represents a TaskUpdated event raised by the TaskManagerContract contract.
type TaskManagerContractTaskUpdated struct {
	TaskId     uint64
	Submitter  common.Address
	UpdateTime *big.Int
	State      uint8
	Result     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskUpdated is a free log retrieval operation binding the contract event 0x5fdff117028da08af2dd3759035973e6e677707cdbd8b5c7f398eec3dd6a6d0f.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint256 indexed updateTime, uint8 state, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterTaskUpdated(opts *bind.FilterOpts, taskId []uint64, submitter []common.Address, updateTime []*big.Int) (*TaskManagerContractTaskUpdatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var updateTimeRule []interface{}
	for _, updateTimeItem := range updateTime {
		updateTimeRule = append(updateTimeRule, updateTimeItem)
	}

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "TaskUpdated", taskIdRule, submitterRule, updateTimeRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTaskUpdatedIterator{contract: _TaskManagerContract.contract, event: "TaskUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskUpdated is a free log subscription operation binding the contract event 0x5fdff117028da08af2dd3759035973e6e677707cdbd8b5c7f398eec3dd6a6d0f.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint256 indexed updateTime, uint8 state, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchTaskUpdated(opts *bind.WatchOpts, sink chan<- *TaskManagerContractTaskUpdated, taskId []uint64, submitter []common.Address, updateTime []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var updateTimeRule []interface{}
	for _, updateTimeItem := range updateTime {
		updateTimeRule = append(updateTimeRule, updateTimeItem)
	}

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "TaskUpdated", taskIdRule, submitterRule, updateTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractTaskUpdated)
				if err := _TaskManagerContract.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
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

// ParseTaskUpdated is a log parse operation binding the contract event 0x5fdff117028da08af2dd3759035973e6e677707cdbd8b5c7f398eec3dd6a6d0f.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint256 indexed updateTime, uint8 state, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseTaskUpdated(log types.Log) (*TaskManagerContractTaskUpdated, error) {
	event := new(TaskManagerContractTaskUpdated)
	if err := _TaskManagerContract.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
