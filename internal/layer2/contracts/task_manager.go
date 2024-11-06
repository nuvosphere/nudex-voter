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
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}

// TaskManagerContractMetaData contains all meta data concerning the TaskManagerContract contract.
var TaskManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"confirmAllTasks\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestTask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITaskManager.Task\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUncompletedTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structITaskManager.Task[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_taskSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTaskCompleted\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markTaskCompleted\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markTaskCompleted_Batch\",\"inputs\":[{\"name\":\"_taskIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextTaskId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmTask\",\"inputs\":[{\"name\":\"_taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preconfirmedTaskResults\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preconfirmedTasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTaskSubmitter\",\"inputs\":[{\"name\":\"_taskSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTask\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isCompleted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"completedAt\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyTask\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTaskSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
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
// Solidity: function getLatestTask() view returns((uint256,bytes,address,bool,uint256,uint256,bytes))
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
// Solidity: function getLatestTask() view returns((uint256,bytes,address,bool,uint256,uint256,bytes))
func (_TaskManagerContract *TaskManagerContractSession) GetLatestTask() (ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetLatestTask(&_TaskManagerContract.CallOpts)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint256,bytes,address,bool,uint256,uint256,bytes))
func (_TaskManagerContract *TaskManagerContractCallerSession) GetLatestTask() (ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetLatestTask(&_TaskManagerContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,bytes,address,bool,uint256,uint256,bytes)[])
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
// Solidity: function getUncompletedTasks() view returns((uint256,bytes,address,bool,uint256,uint256,bytes)[])
func (_TaskManagerContract *TaskManagerContractSession) GetUncompletedTasks() ([]ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetUncompletedTasks(&_TaskManagerContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,bytes,address,bool,uint256,uint256,bytes)[])
func (_TaskManagerContract *TaskManagerContractCallerSession) GetUncompletedTasks() ([]ITaskManagerTask, error) {
	return _TaskManagerContract.Contract.GetUncompletedTasks(&_TaskManagerContract.CallOpts)
}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 _taskId) view returns(bool)
func (_TaskManagerContract *TaskManagerContractCaller) IsTaskCompleted(opts *bind.CallOpts, _taskId *big.Int) (bool, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "isTaskCompleted", _taskId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 _taskId) view returns(bool)
func (_TaskManagerContract *TaskManagerContractSession) IsTaskCompleted(_taskId *big.Int) (bool, error) {
	return _TaskManagerContract.Contract.IsTaskCompleted(&_TaskManagerContract.CallOpts, _taskId)
}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 _taskId) view returns(bool)
func (_TaskManagerContract *TaskManagerContractCallerSession) IsTaskCompleted(_taskId *big.Int) (bool, error) {
	return _TaskManagerContract.Contract.IsTaskCompleted(&_TaskManagerContract.CallOpts, _taskId)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint256)
func (_TaskManagerContract *TaskManagerContractCaller) NextTaskId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "nextTaskId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint256)
func (_TaskManagerContract *TaskManagerContractSession) NextTaskId() (*big.Int, error) {
	return _TaskManagerContract.Contract.NextTaskId(&_TaskManagerContract.CallOpts)
}

// NextTaskId is a free data retrieval call binding the contract method 0xfdc3d8d7.
//
// Solidity: function nextTaskId() view returns(uint256)
func (_TaskManagerContract *TaskManagerContractCallerSession) NextTaskId() (*big.Int, error) {
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
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint256)
func (_TaskManagerContract *TaskManagerContractCaller) PreconfirmedTasks(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "preconfirmedTasks", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint256)
func (_TaskManagerContract *TaskManagerContractSession) PreconfirmedTasks(arg0 *big.Int) (*big.Int, error) {
	return _TaskManagerContract.Contract.PreconfirmedTasks(&_TaskManagerContract.CallOpts, arg0)
}

// PreconfirmedTasks is a free data retrieval call binding the contract method 0x1f07caa8.
//
// Solidity: function preconfirmedTasks(uint256 ) view returns(uint256)
func (_TaskManagerContract *TaskManagerContractCallerSession) PreconfirmedTasks(arg0 *big.Int) (*big.Int, error) {
	return _TaskManagerContract.Contract.PreconfirmedTasks(&_TaskManagerContract.CallOpts, arg0)
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

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, bytes context, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	var out []interface{}
	err := _TaskManagerContract.contract.Call(opts, &out, "tasks", arg0)

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
func (_TaskManagerContract *TaskManagerContractSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	return _TaskManagerContract.Contract.Tasks(&_TaskManagerContract.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, bytes context, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractCallerSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Context     []byte
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	return _TaskManagerContract.Contract.Tasks(&_TaskManagerContract.CallOpts, arg0)
}

// ConfirmAllTasks is a paid mutator transaction binding the contract method 0xa7ec744e.
//
// Solidity: function confirmAllTasks() returns()
func (_TaskManagerContract *TaskManagerContractTransactor) ConfirmAllTasks(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "confirmAllTasks")
}

// ConfirmAllTasks is a paid mutator transaction binding the contract method 0xa7ec744e.
//
// Solidity: function confirmAllTasks() returns()
func (_TaskManagerContract *TaskManagerContractSession) ConfirmAllTasks() (*types.Transaction, error) {
	return _TaskManagerContract.Contract.ConfirmAllTasks(&_TaskManagerContract.TransactOpts)
}

// ConfirmAllTasks is a paid mutator transaction binding the contract method 0xa7ec744e.
//
// Solidity: function confirmAllTasks() returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) ConfirmAllTasks() (*types.Transaction, error) {
	return _TaskManagerContract.Contract.ConfirmAllTasks(&_TaskManagerContract.TransactOpts)
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

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 _taskId, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) MarkTaskCompleted(opts *bind.TransactOpts, _taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "markTaskCompleted", _taskId, _result)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 _taskId, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractSession) MarkTaskCompleted(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.MarkTaskCompleted(&_TaskManagerContract.TransactOpts, _taskId, _result)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 _taskId, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) MarkTaskCompleted(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.MarkTaskCompleted(&_TaskManagerContract.TransactOpts, _taskId, _result)
}

// MarkTaskCompletedBatch is a paid mutator transaction binding the contract method 0xdf2f7649.
//
// Solidity: function markTaskCompleted_Batch(uint256[] _taskIds, bytes[] _results) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) MarkTaskCompletedBatch(opts *bind.TransactOpts, _taskIds []*big.Int, _results [][]byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "markTaskCompleted_Batch", _taskIds, _results)
}

// MarkTaskCompletedBatch is a paid mutator transaction binding the contract method 0xdf2f7649.
//
// Solidity: function markTaskCompleted_Batch(uint256[] _taskIds, bytes[] _results) returns()
func (_TaskManagerContract *TaskManagerContractSession) MarkTaskCompletedBatch(_taskIds []*big.Int, _results [][]byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.MarkTaskCompletedBatch(&_TaskManagerContract.TransactOpts, _taskIds, _results)
}

// MarkTaskCompletedBatch is a paid mutator transaction binding the contract method 0xdf2f7649.
//
// Solidity: function markTaskCompleted_Batch(uint256[] _taskIds, bytes[] _results) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) MarkTaskCompletedBatch(_taskIds []*big.Int, _results [][]byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.MarkTaskCompletedBatch(&_TaskManagerContract.TransactOpts, _taskIds, _results)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x84302728.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) PreconfirmTask(opts *bind.TransactOpts, _taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "preconfirmTask", _taskId, _result)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x84302728.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractSession) PreconfirmTask(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.PreconfirmTask(&_TaskManagerContract.TransactOpts, _taskId, _result)
}

// PreconfirmTask is a paid mutator transaction binding the contract method 0x84302728.
//
// Solidity: function preconfirmTask(uint256 _taskId, bytes _result) returns()
func (_TaskManagerContract *TaskManagerContractTransactorSession) PreconfirmTask(_taskId *big.Int, _result []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.PreconfirmTask(&_TaskManagerContract.TransactOpts, _taskId, _result)
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
// Solidity: function submitTask(address _submitter, bytes _context) returns()
func (_TaskManagerContract *TaskManagerContractTransactor) SubmitTask(opts *bind.TransactOpts, _submitter common.Address, _context []byte) (*types.Transaction, error) {
	return _TaskManagerContract.contract.Transact(opts, "submitTask", _submitter, _context)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _context) returns()
func (_TaskManagerContract *TaskManagerContractSession) SubmitTask(_submitter common.Address, _context []byte) (*types.Transaction, error) {
	return _TaskManagerContract.Contract.SubmitTask(&_TaskManagerContract.TransactOpts, _submitter, _context)
}

// SubmitTask is a paid mutator transaction binding the contract method 0xe2a944e6.
//
// Solidity: function submitTask(address _submitter, bytes _context) returns()
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

// TaskManagerContractTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the TaskManagerContract contract.
type TaskManagerContractTaskCompletedIterator struct {
	Event *TaskManagerContractTaskCompleted // Event containing the contract specifics and raw log

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
func (it *TaskManagerContractTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskManagerContractTaskCompleted)
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
		it.Event = new(TaskManagerContractTaskCompleted)
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
func (it *TaskManagerContractTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskManagerContractTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskManagerContractTaskCompleted represents a TaskCompleted event raised by the TaskManagerContract contract.
type TaskManagerContractTaskCompleted struct {
	TaskId      *big.Int
	Submitter   common.Address
	CompletedAt *big.Int
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 indexed completedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int, submitter []common.Address, completedAt []*big.Int) (*TaskManagerContractTaskCompletedIterator, error) {

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

	logs, sub, err := _TaskManagerContract.contract.FilterLogs(opts, "TaskCompleted", taskIdRule, submitterRule, completedAtRule)
	if err != nil {
		return nil, err
	}
	return &TaskManagerContractTaskCompletedIterator{contract: _TaskManagerContract.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 indexed completedAt, bytes result)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *TaskManagerContractTaskCompleted, taskId []*big.Int, submitter []common.Address, completedAt []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _TaskManagerContract.contract.WatchLogs(opts, "TaskCompleted", taskIdRule, submitterRule, completedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskManagerContractTaskCompleted)
				if err := _TaskManagerContract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_TaskManagerContract *TaskManagerContractFilterer) ParseTaskCompleted(log types.Log) (*TaskManagerContractTaskCompleted, error) {
	event := new(TaskManagerContractTaskCompleted)
	if err := _TaskManagerContract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
	TaskId    *big.Int
	Context   []byte
	Submitter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, bytes context, address indexed submitter)
func (_TaskManagerContract *TaskManagerContractFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, taskId []*big.Int, submitter []common.Address) (*TaskManagerContractTaskSubmittedIterator, error) {

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

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, bytes context, address indexed submitter)
func (_TaskManagerContract *TaskManagerContractFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *TaskManagerContractTaskSubmitted, taskId []*big.Int, submitter []common.Address) (event.Subscription, error) {

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

// ParseTaskSubmitted is a log parse operation binding the contract event 0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, bytes context, address indexed submitter)
func (_TaskManagerContract *TaskManagerContractFilterer) ParseTaskSubmitted(log types.Log) (*TaskManagerContractTaskSubmitted, error) {
	event := new(TaskManagerContractTaskSubmitted)
	if err := _TaskManagerContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
