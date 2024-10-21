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
	Description string
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}

// NuDexOperationsContractMetaData contains all meta data concerning the NuDexOperationsContract contract.
var NuDexOperationsContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"TaskSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLatestTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structINuDexOperations.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUncompletedTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structINuDexOperations.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participantManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"isTaskCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"markTaskCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTaskId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"participantManager\",\"outputs\":[{\"internalType\":\"contractIParticipantManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"submitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function getLatestTask() view returns((uint256,string,address,bool,uint256,uint256,bytes))
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
// Solidity: function getLatestTask() view returns((uint256,string,address,bool,uint256,uint256,bytes))
func (_NuDexOperationsContract *NuDexOperationsContractSession) GetLatestTask() (INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetLatestTask(&_NuDexOperationsContract.CallOpts)
}

// GetLatestTask is a free data retrieval call binding the contract method 0xaccf9a36.
//
// Solidity: function getLatestTask() view returns((uint256,string,address,bool,uint256,uint256,bytes))
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) GetLatestTask() (INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetLatestTask(&_NuDexOperationsContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,string,address,bool,uint256,uint256,bytes)[])
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
// Solidity: function getUncompletedTasks() view returns((uint256,string,address,bool,uint256,uint256,bytes)[])
func (_NuDexOperationsContract *NuDexOperationsContractSession) GetUncompletedTasks() ([]INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetUncompletedTasks(&_NuDexOperationsContract.CallOpts)
}

// GetUncompletedTasks is a free data retrieval call binding the contract method 0x52ee8b0f.
//
// Solidity: function getUncompletedTasks() view returns((uint256,string,address,bool,uint256,uint256,bytes)[])
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) GetUncompletedTasks() ([]INuDexOperationsTask, error) {
	return _NuDexOperationsContract.Contract.GetUncompletedTasks(&_NuDexOperationsContract.CallOpts)
}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 taskId) view returns(bool)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) IsTaskCompleted(opts *bind.CallOpts, taskId *big.Int) (bool, error) {
	var out []interface{}
	err := _NuDexOperationsContract.contract.Call(opts, &out, "isTaskCompleted", taskId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 taskId) view returns(bool)
func (_NuDexOperationsContract *NuDexOperationsContractSession) IsTaskCompleted(taskId *big.Int) (bool, error) {
	return _NuDexOperationsContract.Contract.IsTaskCompleted(&_NuDexOperationsContract.CallOpts, taskId)
}

// IsTaskCompleted is a free data retrieval call binding the contract method 0xb1dc11cb.
//
// Solidity: function isTaskCompleted(uint256 taskId) view returns(bool)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) IsTaskCompleted(taskId *big.Int) (bool, error) {
	return _NuDexOperationsContract.Contract.IsTaskCompleted(&_NuDexOperationsContract.CallOpts, taskId)
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

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string description, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Description string
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
		Description string
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
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Submitter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsCompleted = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.CreatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CompletedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string description, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Description string
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
// Solidity: function tasks(uint256 ) view returns(uint256 id, string description, address submitter, bool isCompleted, uint256 createdAt, uint256 completedAt, bytes result)
func (_NuDexOperationsContract *NuDexOperationsContractCallerSession) Tasks(arg0 *big.Int) (struct {
	Id          *big.Int
	Description string
	Submitter   common.Address
	IsCompleted bool
	CreatedAt   *big.Int
	CompletedAt *big.Int
	Result      []byte
}, error) {
	return _NuDexOperationsContract.Contract.Tasks(&_NuDexOperationsContract.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _participantManager, address _initialOwner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) Initialize(opts *bind.TransactOpts, _participantManager common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "initialize", _participantManager, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _participantManager, address _initialOwner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) Initialize(_participantManager common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.Initialize(&_NuDexOperationsContract.TransactOpts, _participantManager, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _participantManager, address _initialOwner) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) Initialize(_participantManager common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.Initialize(&_NuDexOperationsContract.TransactOpts, _participantManager, _initialOwner)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 taskId, bytes result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) MarkTaskCompleted(opts *bind.TransactOpts, taskId *big.Int, result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "markTaskCompleted", taskId, result)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 taskId, bytes result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) MarkTaskCompleted(taskId *big.Int, result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.MarkTaskCompleted(&_NuDexOperationsContract.TransactOpts, taskId, result)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x4d703d9e.
//
// Solidity: function markTaskCompleted(uint256 taskId, bytes result) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) MarkTaskCompleted(taskId *big.Int, result []byte) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.MarkTaskCompleted(&_NuDexOperationsContract.TransactOpts, taskId, result)
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

// SubmitTask is a paid mutator transaction binding the contract method 0x9a7b22b9.
//
// Solidity: function submitTask(string description) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactor) SubmitTask(opts *bind.TransactOpts, description string) (*types.Transaction, error) {
	return _NuDexOperationsContract.contract.Transact(opts, "submitTask", description)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x9a7b22b9.
//
// Solidity: function submitTask(string description) returns()
func (_NuDexOperationsContract *NuDexOperationsContractSession) SubmitTask(description string) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.SubmitTask(&_NuDexOperationsContract.TransactOpts, description)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x9a7b22b9.
//
// Solidity: function submitTask(string description) returns()
func (_NuDexOperationsContract *NuDexOperationsContractTransactorSession) SubmitTask(description string) (*types.Transaction, error) {
	return _NuDexOperationsContract.Contract.SubmitTask(&_NuDexOperationsContract.TransactOpts, description)
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
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x843af93d40addceac6932508439844b897d4df9e971db326d557e3cdaa9f3ebf.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 completedAt)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int, submitter []common.Address) (*NuDexOperationsContractTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.FilterLogs(opts, "TaskCompleted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &NuDexOperationsContractTaskCompletedIterator{contract: _NuDexOperationsContract.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x843af93d40addceac6932508439844b897d4df9e971db326d557e3cdaa9f3ebf.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 completedAt)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *NuDexOperationsContractTaskCompleted, taskId []*big.Int, submitter []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _NuDexOperationsContract.contract.WatchLogs(opts, "TaskCompleted", taskIdRule, submitterRule)
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x843af93d40addceac6932508439844b897d4df9e971db326d557e3cdaa9f3ebf.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 completedAt)
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
	TaskId      *big.Int
	Description string
	Submitter   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x475a19c94ceb828d9d7c5bd38d863932bba45ea65488dcb5d1c386b4ec2bb76c.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, string description, address indexed submitter)
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

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0x475a19c94ceb828d9d7c5bd38d863932bba45ea65488dcb5d1c386b4ec2bb76c.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, string description, address indexed submitter)
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0x475a19c94ceb828d9d7c5bd38d863932bba45ea65488dcb5d1c386b4ec2bb76c.
//
// Solidity: event TaskSubmitted(uint256 indexed taskId, string description, address indexed submitter)
func (_NuDexOperationsContract *NuDexOperationsContractFilterer) ParseTaskSubmitted(log types.Log) (*NuDexOperationsContractTaskSubmitted, error) {
	event := new(NuDexOperationsContractTaskSubmitted)
	if err := _NuDexOperationsContract.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
