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

// IDepositManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IDepositManagerDepositInfo struct {
	TargetAddress common.Address
	ChainId       *big.Int
	Amount        *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}

// IDepositManagerWithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type IDepositManagerWithdrawalInfo struct {
	TargetAddress common.Address
	Amount        *big.Int
	ChainId       *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}

// DepositManagerContractMetaData contains all meta data concerning the DepositManagerContract contract.
var DepositManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposit\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.DepositInfo\",\"components\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDepositManager.DepositInfo[]\",\"components\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawal\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDepositManager.WithdrawalInfo\",\"components\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawals\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDepositManager.WithdrawalInfo[]\",\"components\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordDeposit\",\"inputs\":[{\"name\":\"_targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordWithdrawal\",\"inputs\":[{\"name\":\"_targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositRecorded\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRecorded\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"txInfo\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"extraInfo\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// DepositManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositManagerContractMetaData.ABI instead.
var DepositManagerContractABI = DepositManagerContractMetaData.ABI

// DepositManagerContract is an auto generated Go binding around an Ethereum contract.
type DepositManagerContract struct {
	DepositManagerContractCaller     // Read-only binding to the contract
	DepositManagerContractTransactor // Write-only binding to the contract
	DepositManagerContractFilterer   // Log filterer for contract events
}

// DepositManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositManagerContractSession struct {
	Contract     *DepositManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DepositManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositManagerContractCallerSession struct {
	Contract *DepositManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DepositManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositManagerContractTransactorSession struct {
	Contract     *DepositManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DepositManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositManagerContractRaw struct {
	Contract *DepositManagerContract // Generic contract binding to access the raw methods on
}

// DepositManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositManagerContractCallerRaw struct {
	Contract *DepositManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// DepositManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositManagerContractTransactorRaw struct {
	Contract *DepositManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositManagerContract creates a new instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContract(address common.Address, backend bind.ContractBackend) (*DepositManagerContract, error) {
	contract, err := bindDepositManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContract{DepositManagerContractCaller: DepositManagerContractCaller{contract: contract}, DepositManagerContractTransactor: DepositManagerContractTransactor{contract: contract}, DepositManagerContractFilterer: DepositManagerContractFilterer{contract: contract}}, nil
}

// NewDepositManagerContractCaller creates a new read-only instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContractCaller(address common.Address, caller bind.ContractCaller) (*DepositManagerContractCaller, error) {
	contract, err := bindDepositManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractCaller{contract: contract}, nil
}

// NewDepositManagerContractTransactor creates a new write-only instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositManagerContractTransactor, error) {
	contract, err := bindDepositManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractTransactor{contract: contract}, nil
}

// NewDepositManagerContractFilterer creates a new log filterer instance of DepositManagerContract, bound to a specific deployed contract.
func NewDepositManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositManagerContractFilterer, error) {
	contract, err := bindDepositManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractFilterer{contract: contract}, nil
}

// bindDepositManagerContract binds a generic wrapper to an already deployed contract.
func bindDepositManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManagerContract *DepositManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositManagerContract.Contract.DepositManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManagerContract *DepositManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.DepositManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManagerContract *DepositManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.DepositManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManagerContract *DepositManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManagerContract *DepositManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManagerContract *DepositManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0xd6d68177.
//
// Solidity: function deposits(address , uint256 ) view returns(address targetAddress, uint256 chainId, uint256 amount, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TargetAddress common.Address
	ChainId       *big.Int
	Amount        *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "deposits", arg0, arg1)

	outstruct := new(struct {
		TargetAddress common.Address
		ChainId       *big.Int
		Amount        *big.Int
		TxInfo        []byte
		ExtraInfo     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TargetAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TxInfo = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ExtraInfo = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xd6d68177.
//
// Solidity: function deposits(address , uint256 ) view returns(address targetAddress, uint256 chainId, uint256 amount, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractSession) Deposits(arg0 common.Address, arg1 *big.Int) (struct {
	TargetAddress common.Address
	ChainId       *big.Int
	Amount        *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}, error) {
	return _DepositManagerContract.Contract.Deposits(&_DepositManagerContract.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0xd6d68177.
//
// Solidity: function deposits(address , uint256 ) view returns(address targetAddress, uint256 chainId, uint256 amount, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractCallerSession) Deposits(arg0 common.Address, arg1 *big.Int) (struct {
	TargetAddress common.Address
	ChainId       *big.Int
	Amount        *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}, error) {
	return _DepositManagerContract.Contract.Deposits(&_DepositManagerContract.CallOpts, arg0, arg1)
}

// GetDeposit is a free data retrieval call binding the contract method 0x2726b506.
//
// Solidity: function getDeposit(address targetAddress, uint256 index) view returns((address,uint256,uint256,bytes,bytes))
func (_DepositManagerContract *DepositManagerContractCaller) GetDeposit(opts *bind.CallOpts, targetAddress common.Address, index *big.Int) (IDepositManagerDepositInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getDeposit", targetAddress, index)

	if err != nil {
		return *new(IDepositManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDepositManagerDepositInfo)).(*IDepositManagerDepositInfo)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0x2726b506.
//
// Solidity: function getDeposit(address targetAddress, uint256 index) view returns((address,uint256,uint256,bytes,bytes))
func (_DepositManagerContract *DepositManagerContractSession) GetDeposit(targetAddress common.Address, index *big.Int) (IDepositManagerDepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposit(&_DepositManagerContract.CallOpts, targetAddress, index)
}

// GetDeposit is a free data retrieval call binding the contract method 0x2726b506.
//
// Solidity: function getDeposit(address targetAddress, uint256 index) view returns((address,uint256,uint256,bytes,bytes))
func (_DepositManagerContract *DepositManagerContractCallerSession) GetDeposit(targetAddress common.Address, index *big.Int) (IDepositManagerDepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposit(&_DepositManagerContract.CallOpts, targetAddress, index)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address targetAddress) view returns((address,uint256,uint256,bytes,bytes)[])
func (_DepositManagerContract *DepositManagerContractCaller) GetDeposits(opts *bind.CallOpts, targetAddress common.Address) ([]IDepositManagerDepositInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getDeposits", targetAddress)

	if err != nil {
		return *new([]IDepositManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDepositManagerDepositInfo)).(*[]IDepositManagerDepositInfo)

	return out0, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address targetAddress) view returns((address,uint256,uint256,bytes,bytes)[])
func (_DepositManagerContract *DepositManagerContractSession) GetDeposits(targetAddress common.Address) ([]IDepositManagerDepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposits(&_DepositManagerContract.CallOpts, targetAddress)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address targetAddress) view returns((address,uint256,uint256,bytes,bytes)[])
func (_DepositManagerContract *DepositManagerContractCallerSession) GetDeposits(targetAddress common.Address) ([]IDepositManagerDepositInfo, error) {
	return _DepositManagerContract.Contract.GetDeposits(&_DepositManagerContract.CallOpts, targetAddress)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address targetAddress, uint256 index) view returns((address,uint256,uint256,bytes,bytes))
func (_DepositManagerContract *DepositManagerContractCaller) GetWithdrawal(opts *bind.CallOpts, targetAddress common.Address, index *big.Int) (IDepositManagerWithdrawalInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getWithdrawal", targetAddress, index)

	if err != nil {
		return *new(IDepositManagerWithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDepositManagerWithdrawalInfo)).(*IDepositManagerWithdrawalInfo)

	return out0, err

}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address targetAddress, uint256 index) view returns((address,uint256,uint256,bytes,bytes))
func (_DepositManagerContract *DepositManagerContractSession) GetWithdrawal(targetAddress common.Address, index *big.Int) (IDepositManagerWithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawal(&_DepositManagerContract.CallOpts, targetAddress, index)
}

// GetWithdrawal is a free data retrieval call binding the contract method 0x5be4f765.
//
// Solidity: function getWithdrawal(address targetAddress, uint256 index) view returns((address,uint256,uint256,bytes,bytes))
func (_DepositManagerContract *DepositManagerContractCallerSession) GetWithdrawal(targetAddress common.Address, index *big.Int) (IDepositManagerWithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawal(&_DepositManagerContract.CallOpts, targetAddress, index)
}

// GetWithdrawals is a free data retrieval call binding the contract method 0x3a2b643a.
//
// Solidity: function getWithdrawals(address targetAddress) view returns((address,uint256,uint256,bytes,bytes)[])
func (_DepositManagerContract *DepositManagerContractCaller) GetWithdrawals(opts *bind.CallOpts, targetAddress common.Address) ([]IDepositManagerWithdrawalInfo, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "getWithdrawals", targetAddress)

	if err != nil {
		return *new([]IDepositManagerWithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDepositManagerWithdrawalInfo)).(*[]IDepositManagerWithdrawalInfo)

	return out0, err

}

// GetWithdrawals is a free data retrieval call binding the contract method 0x3a2b643a.
//
// Solidity: function getWithdrawals(address targetAddress) view returns((address,uint256,uint256,bytes,bytes)[])
func (_DepositManagerContract *DepositManagerContractSession) GetWithdrawals(targetAddress common.Address) ([]IDepositManagerWithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawals(&_DepositManagerContract.CallOpts, targetAddress)
}

// GetWithdrawals is a free data retrieval call binding the contract method 0x3a2b643a.
//
// Solidity: function getWithdrawals(address targetAddress) view returns((address,uint256,uint256,bytes,bytes)[])
func (_DepositManagerContract *DepositManagerContractCallerSession) GetWithdrawals(targetAddress common.Address) ([]IDepositManagerWithdrawalInfo, error) {
	return _DepositManagerContract.Contract.GetWithdrawals(&_DepositManagerContract.CallOpts, targetAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DepositManagerContract *DepositManagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DepositManagerContract *DepositManagerContractSession) Owner() (common.Address, error) {
	return _DepositManagerContract.Contract.Owner(&_DepositManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DepositManagerContract *DepositManagerContractCallerSession) Owner() (common.Address, error) {
	return _DepositManagerContract.Contract.Owner(&_DepositManagerContract.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x422b1077.
//
// Solidity: function withdrawals(address , uint256 ) view returns(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractCaller) Withdrawals(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TargetAddress common.Address
	Amount        *big.Int
	ChainId       *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}, error) {
	var out []interface{}
	err := _DepositManagerContract.contract.Call(opts, &out, "withdrawals", arg0, arg1)

	outstruct := new(struct {
		TargetAddress common.Address
		Amount        *big.Int
		ChainId       *big.Int
		TxInfo        []byte
		ExtraInfo     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TargetAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TxInfo = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ExtraInfo = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x422b1077.
//
// Solidity: function withdrawals(address , uint256 ) view returns(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractSession) Withdrawals(arg0 common.Address, arg1 *big.Int) (struct {
	TargetAddress common.Address
	Amount        *big.Int
	ChainId       *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}, error) {
	return _DepositManagerContract.Contract.Withdrawals(&_DepositManagerContract.CallOpts, arg0, arg1)
}

// Withdrawals is a free data retrieval call binding the contract method 0x422b1077.
//
// Solidity: function withdrawals(address , uint256 ) view returns(address targetAddress, uint256 amount, uint256 chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractCallerSession) Withdrawals(arg0 common.Address, arg1 *big.Int) (struct {
	TargetAddress common.Address
	Amount        *big.Int
	ChainId       *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
}, error) {
	return _DepositManagerContract.Contract.Withdrawals(&_DepositManagerContract.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_DepositManagerContract *DepositManagerContractSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.Initialize(&_DepositManagerContract.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.Initialize(&_DepositManagerContract.TransactOpts, _owner)
}

// RecordDeposit is a paid mutator transaction binding the contract method 0xbac89758.
//
// Solidity: function recordDeposit(address _targetAddress, uint256 _amount, uint256 _chainId, bytes _txInfo, bytes _extraInfo) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactor) RecordDeposit(opts *bind.TransactOpts, _targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "recordDeposit", _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

// RecordDeposit is a paid mutator transaction binding the contract method 0xbac89758.
//
// Solidity: function recordDeposit(address _targetAddress, uint256 _amount, uint256 _chainId, bytes _txInfo, bytes _extraInfo) returns(bytes)
func (_DepositManagerContract *DepositManagerContractSession) RecordDeposit(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordDeposit(&_DepositManagerContract.TransactOpts, _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

// RecordDeposit is a paid mutator transaction binding the contract method 0xbac89758.
//
// Solidity: function recordDeposit(address _targetAddress, uint256 _amount, uint256 _chainId, bytes _txInfo, bytes _extraInfo) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactorSession) RecordDeposit(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordDeposit(&_DepositManagerContract.TransactOpts, _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

// RecordWithdrawal is a paid mutator transaction binding the contract method 0x57673f43.
//
// Solidity: function recordWithdrawal(address _targetAddress, uint256 _amount, uint256 _chainId, bytes _txInfo, bytes _extraInfo) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactor) RecordWithdrawal(opts *bind.TransactOpts, _targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "recordWithdrawal", _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

// RecordWithdrawal is a paid mutator transaction binding the contract method 0x57673f43.
//
// Solidity: function recordWithdrawal(address _targetAddress, uint256 _amount, uint256 _chainId, bytes _txInfo, bytes _extraInfo) returns(bytes)
func (_DepositManagerContract *DepositManagerContractSession) RecordWithdrawal(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordWithdrawal(&_DepositManagerContract.TransactOpts, _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

// RecordWithdrawal is a paid mutator transaction binding the contract method 0x57673f43.
//
// Solidity: function recordWithdrawal(address _targetAddress, uint256 _amount, uint256 _chainId, bytes _txInfo, bytes _extraInfo) returns(bytes)
func (_DepositManagerContract *DepositManagerContractTransactorSession) RecordWithdrawal(_targetAddress common.Address, _amount *big.Int, _chainId *big.Int, _txInfo []byte, _extraInfo []byte) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RecordWithdrawal(&_DepositManagerContract.TransactOpts, _targetAddress, _amount, _chainId, _txInfo, _extraInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositManagerContract *DepositManagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositManagerContract *DepositManagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RenounceOwnership(&_DepositManagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DepositManagerContract.Contract.RenounceOwnership(&_DepositManagerContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositManagerContract *DepositManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositManagerContract *DepositManagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.TransferOwnership(&_DepositManagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositManagerContract *DepositManagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositManagerContract.Contract.TransferOwnership(&_DepositManagerContract.TransactOpts, newOwner)
}

// DepositManagerContractDepositRecordedIterator is returned from FilterDepositRecorded and is used to iterate over the raw logs and unpacked data for DepositRecorded events raised by the DepositManagerContract contract.
type DepositManagerContractDepositRecordedIterator struct {
	Event *DepositManagerContractDepositRecorded // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractDepositRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractDepositRecorded)
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
		it.Event = new(DepositManagerContractDepositRecorded)
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
func (it *DepositManagerContractDepositRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractDepositRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractDepositRecorded represents a DepositRecorded event raised by the DepositManagerContract contract.
type DepositManagerContractDepositRecorded struct {
	TargetAddress common.Address
	Amount        *big.Int
	ChainId       *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositRecorded is a free log retrieval operation binding the contract event 0xda0e7c971690dbb1d8118c31cf27f8303b471719eb78a6200d35300175974100.
//
// Solidity: event DepositRecorded(address indexed targetAddress, uint256 indexed amount, uint256 indexed chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterDepositRecorded(opts *bind.FilterOpts, targetAddress []common.Address, amount []*big.Int, chainId []*big.Int) (*DepositManagerContractDepositRecordedIterator, error) {

	var targetAddressRule []interface{}
	for _, targetAddressItem := range targetAddress {
		targetAddressRule = append(targetAddressRule, targetAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "DepositRecorded", targetAddressRule, amountRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractDepositRecordedIterator{contract: _DepositManagerContract.contract, event: "DepositRecorded", logs: logs, sub: sub}, nil
}

// WatchDepositRecorded is a free log subscription operation binding the contract event 0xda0e7c971690dbb1d8118c31cf27f8303b471719eb78a6200d35300175974100.
//
// Solidity: event DepositRecorded(address indexed targetAddress, uint256 indexed amount, uint256 indexed chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchDepositRecorded(opts *bind.WatchOpts, sink chan<- *DepositManagerContractDepositRecorded, targetAddress []common.Address, amount []*big.Int, chainId []*big.Int) (event.Subscription, error) {

	var targetAddressRule []interface{}
	for _, targetAddressItem := range targetAddress {
		targetAddressRule = append(targetAddressRule, targetAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "DepositRecorded", targetAddressRule, amountRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractDepositRecorded)
				if err := _DepositManagerContract.contract.UnpackLog(event, "DepositRecorded", log); err != nil {
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

// ParseDepositRecorded is a log parse operation binding the contract event 0xda0e7c971690dbb1d8118c31cf27f8303b471719eb78a6200d35300175974100.
//
// Solidity: event DepositRecorded(address indexed targetAddress, uint256 indexed amount, uint256 indexed chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseDepositRecorded(log types.Log) (*DepositManagerContractDepositRecorded, error) {
	event := new(DepositManagerContractDepositRecorded)
	if err := _DepositManagerContract.contract.UnpackLog(event, "DepositRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DepositManagerContract contract.
type DepositManagerContractInitializedIterator struct {
	Event *DepositManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractInitialized)
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
		it.Event = new(DepositManagerContractInitialized)
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
func (it *DepositManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractInitialized represents a Initialized event raised by the DepositManagerContract contract.
type DepositManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*DepositManagerContractInitializedIterator, error) {

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractInitializedIterator{contract: _DepositManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DepositManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractInitialized)
				if err := _DepositManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DepositManagerContract *DepositManagerContractFilterer) ParseInitialized(log types.Log) (*DepositManagerContractInitialized, error) {
	event := new(DepositManagerContractInitialized)
	if err := _DepositManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DepositManagerContract contract.
type DepositManagerContractOwnershipTransferredIterator struct {
	Event *DepositManagerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractOwnershipTransferred)
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
		it.Event = new(DepositManagerContractOwnershipTransferred)
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
func (it *DepositManagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the DepositManagerContract contract.
type DepositManagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositManagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractOwnershipTransferredIterator{contract: _DepositManagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositManagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractOwnershipTransferred)
				if err := _DepositManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DepositManagerContract *DepositManagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*DepositManagerContractOwnershipTransferred, error) {
	event := new(DepositManagerContractOwnershipTransferred)
	if err := _DepositManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerContractWithdrawalRecordedIterator is returned from FilterWithdrawalRecorded and is used to iterate over the raw logs and unpacked data for WithdrawalRecorded events raised by the DepositManagerContract contract.
type DepositManagerContractWithdrawalRecordedIterator struct {
	Event *DepositManagerContractWithdrawalRecorded // Event containing the contract specifics and raw log

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
func (it *DepositManagerContractWithdrawalRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerContractWithdrawalRecorded)
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
		it.Event = new(DepositManagerContractWithdrawalRecorded)
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
func (it *DepositManagerContractWithdrawalRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerContractWithdrawalRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerContractWithdrawalRecorded represents a WithdrawalRecorded event raised by the DepositManagerContract contract.
type DepositManagerContractWithdrawalRecorded struct {
	TargetAddress common.Address
	Amount        *big.Int
	ChainId       *big.Int
	TxInfo        []byte
	ExtraInfo     []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRecorded is a free log retrieval operation binding the contract event 0x07c8f2d211076c7cba51f2504af48025acdaf410e993e6c7f62b066a51d9b068.
//
// Solidity: event WithdrawalRecorded(address indexed targetAddress, uint256 indexed amount, uint256 indexed chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) FilterWithdrawalRecorded(opts *bind.FilterOpts, targetAddress []common.Address, amount []*big.Int, chainId []*big.Int) (*DepositManagerContractWithdrawalRecordedIterator, error) {

	var targetAddressRule []interface{}
	for _, targetAddressItem := range targetAddress {
		targetAddressRule = append(targetAddressRule, targetAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.FilterLogs(opts, "WithdrawalRecorded", targetAddressRule, amountRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerContractWithdrawalRecordedIterator{contract: _DepositManagerContract.contract, event: "WithdrawalRecorded", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRecorded is a free log subscription operation binding the contract event 0x07c8f2d211076c7cba51f2504af48025acdaf410e993e6c7f62b066a51d9b068.
//
// Solidity: event WithdrawalRecorded(address indexed targetAddress, uint256 indexed amount, uint256 indexed chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) WatchWithdrawalRecorded(opts *bind.WatchOpts, sink chan<- *DepositManagerContractWithdrawalRecorded, targetAddress []common.Address, amount []*big.Int, chainId []*big.Int) (event.Subscription, error) {

	var targetAddressRule []interface{}
	for _, targetAddressItem := range targetAddress {
		targetAddressRule = append(targetAddressRule, targetAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _DepositManagerContract.contract.WatchLogs(opts, "WithdrawalRecorded", targetAddressRule, amountRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerContractWithdrawalRecorded)
				if err := _DepositManagerContract.contract.UnpackLog(event, "WithdrawalRecorded", log); err != nil {
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

// ParseWithdrawalRecorded is a log parse operation binding the contract event 0x07c8f2d211076c7cba51f2504af48025acdaf410e993e6c7f62b066a51d9b068.
//
// Solidity: event WithdrawalRecorded(address indexed targetAddress, uint256 indexed amount, uint256 indexed chainId, bytes txInfo, bytes extraInfo)
func (_DepositManagerContract *DepositManagerContractFilterer) ParseWithdrawalRecorded(log types.Log) (*DepositManagerContractWithdrawalRecorded, error) {
	event := new(DepositManagerContractWithdrawalRecorded)
	if err := _DepositManagerContract.contract.UnpackLog(event, "WithdrawalRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
