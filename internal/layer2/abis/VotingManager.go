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

// VotingManagerContractMetaData contains all meta data concerning the VotingManagerContract contract.
var VotingManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txInfo\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraInfo\",\"type\":\"bytes\"}],\"name\":\"DepositInfoSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newParticipant\",\"type\":\"address\"}],\"name\":\"ParticipantAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"ParticipantRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRewardPerPeriod\",\"type\":\"uint256\"}],\"name\":\"RewardPerPeriodVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSubmitter\",\"type\":\"address\"}],\"name\":\"SubmitterChosen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentSubmitter\",\"type\":\"address\"}],\"name\":\"SubmitterRotationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"taskResult\",\"type\":\"bytes\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newParticipant\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetManager\",\"outputs\":[{\"internalType\":\"contractAssetManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentSubmitter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"chooseNewSubmitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIAssetManager.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"delistAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositManager\",\"outputs\":[{\"internalType\":\"contractDepositManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forcedRotationWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSubmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participantManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nuvoLock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assetManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nuDexOperations\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSubmissionTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSubmitterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nuDexName\",\"type\":\"string\"},{\"internalType\":\"enumIAssetManager.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"listAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nuDexOperations\",\"outputs\":[{\"internalType\":\"contractNuDexOperations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nuvoLock\",\"outputs\":[{\"internalType\":\"contractNuvoLockUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"participantManager\",\"outputs\":[{\"internalType\":\"contractParticipantManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRewardPerPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setRewardPerPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txInfo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraInfo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"submitDepositInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"submitTaskReceipt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskCompletionThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AssetManager is a free data retrieval call binding the contract method 0x94217ad1.
//
// Solidity: function assetManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) AssetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "assetManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetManager is a free data retrieval call binding the contract method 0x94217ad1.
//
// Solidity: function assetManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) AssetManager() (common.Address, error) {
	return _VotingManagerContract.Contract.AssetManager(&_VotingManagerContract.CallOpts)
}

// AssetManager is a free data retrieval call binding the contract method 0x94217ad1.
//
// Solidity: function assetManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) AssetManager() (common.Address, error) {
	return _VotingManagerContract.Contract.AssetManager(&_VotingManagerContract.CallOpts)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) DepositManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "depositManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) DepositManager() (common.Address, error) {
	return _VotingManagerContract.Contract.DepositManager(&_VotingManagerContract.CallOpts)
}

// DepositManager is a free data retrieval call binding the contract method 0x6c7ac9d8.
//
// Solidity: function depositManager() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) DepositManager() (common.Address, error) {
	return _VotingManagerContract.Contract.DepositManager(&_VotingManagerContract.CallOpts)
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

// GetCurrentSubmitter is a free data retrieval call binding the contract method 0x6b511e82.
//
// Solidity: function getCurrentSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) GetCurrentSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "getCurrentSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentSubmitter is a free data retrieval call binding the contract method 0x6b511e82.
//
// Solidity: function getCurrentSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) GetCurrentSubmitter() (common.Address, error) {
	return _VotingManagerContract.Contract.GetCurrentSubmitter(&_VotingManagerContract.CallOpts)
}

// GetCurrentSubmitter is a free data retrieval call binding the contract method 0x6b511e82.
//
// Solidity: function getCurrentSubmitter() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) GetCurrentSubmitter() (common.Address, error) {
	return _VotingManagerContract.Contract.GetCurrentSubmitter(&_VotingManagerContract.CallOpts)
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

// LastSubmitterIndex is a free data retrieval call binding the contract method 0x886150da.
//
// Solidity: function lastSubmitterIndex() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCaller) LastSubmitterIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "lastSubmitterIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSubmitterIndex is a free data retrieval call binding the contract method 0x886150da.
//
// Solidity: function lastSubmitterIndex() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractSession) LastSubmitterIndex() (*big.Int, error) {
	return _VotingManagerContract.Contract.LastSubmitterIndex(&_VotingManagerContract.CallOpts)
}

// LastSubmitterIndex is a free data retrieval call binding the contract method 0x886150da.
//
// Solidity: function lastSubmitterIndex() view returns(uint256)
func (_VotingManagerContract *VotingManagerContractCallerSession) LastSubmitterIndex() (*big.Int, error) {
	return _VotingManagerContract.Contract.LastSubmitterIndex(&_VotingManagerContract.CallOpts)
}

// NuDexOperations is a free data retrieval call binding the contract method 0xb02ae352.
//
// Solidity: function nuDexOperations() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) NuDexOperations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "nuDexOperations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NuDexOperations is a free data retrieval call binding the contract method 0xb02ae352.
//
// Solidity: function nuDexOperations() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) NuDexOperations() (common.Address, error) {
	return _VotingManagerContract.Contract.NuDexOperations(&_VotingManagerContract.CallOpts)
}

// NuDexOperations is a free data retrieval call binding the contract method 0xb02ae352.
//
// Solidity: function nuDexOperations() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) NuDexOperations() (common.Address, error) {
	return _VotingManagerContract.Contract.NuDexOperations(&_VotingManagerContract.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotingManagerContract *VotingManagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingManagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotingManagerContract *VotingManagerContractSession) Owner() (common.Address, error) {
	return _VotingManagerContract.Contract.Owner(&_VotingManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotingManagerContract *VotingManagerContractCallerSession) Owner() (common.Address, error) {
	return _VotingManagerContract.Contract.Owner(&_VotingManagerContract.CallOpts)
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

// AddParticipant is a paid mutator transaction binding the contract method 0x1b8d2365.
//
// Solidity: function addParticipant(address newParticipant, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) AddParticipant(opts *bind.TransactOpts, newParticipant common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "addParticipant", newParticipant, params, signature)
}

// AddParticipant is a paid mutator transaction binding the contract method 0x1b8d2365.
//
// Solidity: function addParticipant(address newParticipant, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) AddParticipant(newParticipant common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.AddParticipant(&_VotingManagerContract.TransactOpts, newParticipant, params, signature)
}

// AddParticipant is a paid mutator transaction binding the contract method 0x1b8d2365.
//
// Solidity: function addParticipant(address newParticipant, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) AddParticipant(newParticipant common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.AddParticipant(&_VotingManagerContract.TransactOpts, newParticipant, params, signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0x3ac785bb.
//
// Solidity: function chooseNewSubmitter(address currentSubmitter, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) ChooseNewSubmitter(opts *bind.TransactOpts, currentSubmitter common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "chooseNewSubmitter", currentSubmitter, params, signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0x3ac785bb.
//
// Solidity: function chooseNewSubmitter(address currentSubmitter, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) ChooseNewSubmitter(currentSubmitter common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ChooseNewSubmitter(&_VotingManagerContract.TransactOpts, currentSubmitter, params, signature)
}

// ChooseNewSubmitter is a paid mutator transaction binding the contract method 0x3ac785bb.
//
// Solidity: function chooseNewSubmitter(address currentSubmitter, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) ChooseNewSubmitter(currentSubmitter common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ChooseNewSubmitter(&_VotingManagerContract.TransactOpts, currentSubmitter, params, signature)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe7a36191.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) DelistAsset(opts *bind.TransactOpts, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "delistAsset", assetType, contractAddress, chainId, signature)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe7a36191.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) DelistAsset(assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.DelistAsset(&_VotingManagerContract.TransactOpts, assetType, contractAddress, chainId, signature)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe7a36191.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) DelistAsset(assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.DelistAsset(&_VotingManagerContract.TransactOpts, assetType, contractAddress, chainId, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _participantManager, address _nuvoLock, address _assetManager, address _depositManager, address _nuDexOperations, address _initialOwner) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) Initialize(opts *bind.TransactOpts, _participantManager common.Address, _nuvoLock common.Address, _assetManager common.Address, _depositManager common.Address, _nuDexOperations common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "initialize", _participantManager, _nuvoLock, _assetManager, _depositManager, _nuDexOperations, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _participantManager, address _nuvoLock, address _assetManager, address _depositManager, address _nuDexOperations, address _initialOwner) returns()
func (_VotingManagerContract *VotingManagerContractSession) Initialize(_participantManager common.Address, _nuvoLock common.Address, _assetManager common.Address, _depositManager common.Address, _nuDexOperations common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.Initialize(&_VotingManagerContract.TransactOpts, _participantManager, _nuvoLock, _assetManager, _depositManager, _nuDexOperations, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _participantManager, address _nuvoLock, address _assetManager, address _depositManager, address _nuDexOperations, address _initialOwner) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) Initialize(_participantManager common.Address, _nuvoLock common.Address, _assetManager common.Address, _depositManager common.Address, _nuDexOperations common.Address, _initialOwner common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.Initialize(&_VotingManagerContract.TransactOpts, _participantManager, _nuvoLock, _assetManager, _depositManager, _nuDexOperations, _initialOwner)
}

// ListAsset is a paid mutator transaction binding the contract method 0xd1c0a28a.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) ListAsset(opts *bind.TransactOpts, name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "listAsset", name, nuDexName, assetType, contractAddress, chainId, signature)
}

// ListAsset is a paid mutator transaction binding the contract method 0xd1c0a28a.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) ListAsset(name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ListAsset(&_VotingManagerContract.TransactOpts, name, nuDexName, assetType, contractAddress, chainId, signature)
}

// ListAsset is a paid mutator transaction binding the contract method 0xd1c0a28a.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) ListAsset(name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.ListAsset(&_VotingManagerContract.TransactOpts, name, nuDexName, assetType, contractAddress, chainId, signature)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0xdd1ee03e.
//
// Solidity: function removeParticipant(address participant, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) RemoveParticipant(opts *bind.TransactOpts, participant common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "removeParticipant", participant, params, signature)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0xdd1ee03e.
//
// Solidity: function removeParticipant(address participant, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) RemoveParticipant(participant common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RemoveParticipant(&_VotingManagerContract.TransactOpts, participant, params, signature)
}

// RemoveParticipant is a paid mutator transaction binding the contract method 0xdd1ee03e.
//
// Solidity: function removeParticipant(address participant, bytes params, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) RemoveParticipant(participant common.Address, params []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RemoveParticipant(&_VotingManagerContract.TransactOpts, participant, params, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotingManagerContract *VotingManagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotingManagerContract *VotingManagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RenounceOwnership(&_VotingManagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VotingManagerContract.Contract.RenounceOwnership(&_VotingManagerContract.TransactOpts)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0x77833f18.
//
// Solidity: function setRewardPerPeriod(uint256 newRewardPerPeriod, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SetRewardPerPeriod(opts *bind.TransactOpts, newRewardPerPeriod *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "setRewardPerPeriod", newRewardPerPeriod, signature)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0x77833f18.
//
// Solidity: function setRewardPerPeriod(uint256 newRewardPerPeriod, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SetRewardPerPeriod(newRewardPerPeriod *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SetRewardPerPeriod(&_VotingManagerContract.TransactOpts, newRewardPerPeriod, signature)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0x77833f18.
//
// Solidity: function setRewardPerPeriod(uint256 newRewardPerPeriod, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SetRewardPerPeriod(newRewardPerPeriod *big.Int, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SetRewardPerPeriod(&_VotingManagerContract.TransactOpts, newRewardPerPeriod, signature)
}

// SubmitDepositInfo is a paid mutator transaction binding the contract method 0xbfe07e7a.
//
// Solidity: function submitDepositInfo(address targetAddress, uint256 amount, bytes txInfo, uint256 chainId, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SubmitDepositInfo(opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int, txInfo []byte, chainId *big.Int, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "submitDepositInfo", targetAddress, amount, txInfo, chainId, extraInfo, signature)
}

// SubmitDepositInfo is a paid mutator transaction binding the contract method 0xbfe07e7a.
//
// Solidity: function submitDepositInfo(address targetAddress, uint256 amount, bytes txInfo, uint256 chainId, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SubmitDepositInfo(targetAddress common.Address, amount *big.Int, txInfo []byte, chainId *big.Int, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitDepositInfo(&_VotingManagerContract.TransactOpts, targetAddress, amount, txInfo, chainId, extraInfo, signature)
}

// SubmitDepositInfo is a paid mutator transaction binding the contract method 0xbfe07e7a.
//
// Solidity: function submitDepositInfo(address targetAddress, uint256 amount, bytes txInfo, uint256 chainId, bytes extraInfo, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SubmitDepositInfo(targetAddress common.Address, amount *big.Int, txInfo []byte, chainId *big.Int, extraInfo []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitDepositInfo(&_VotingManagerContract.TransactOpts, targetAddress, amount, txInfo, chainId, extraInfo, signature)
}

// SubmitTaskReceipt is a paid mutator transaction binding the contract method 0x980d16c9.
//
// Solidity: function submitTaskReceipt(uint256 taskId, bytes result, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) SubmitTaskReceipt(opts *bind.TransactOpts, taskId *big.Int, result []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "submitTaskReceipt", taskId, result, signature)
}

// SubmitTaskReceipt is a paid mutator transaction binding the contract method 0x980d16c9.
//
// Solidity: function submitTaskReceipt(uint256 taskId, bytes result, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractSession) SubmitTaskReceipt(taskId *big.Int, result []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitTaskReceipt(&_VotingManagerContract.TransactOpts, taskId, result, signature)
}

// SubmitTaskReceipt is a paid mutator transaction binding the contract method 0x980d16c9.
//
// Solidity: function submitTaskReceipt(uint256 taskId, bytes result, bytes signature) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) SubmitTaskReceipt(taskId *big.Int, result []byte, signature []byte) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.SubmitTaskReceipt(&_VotingManagerContract.TransactOpts, taskId, result, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotingManagerContract *VotingManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotingManagerContract *VotingManagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.TransferOwnership(&_VotingManagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotingManagerContract *VotingManagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotingManagerContract.Contract.TransferOwnership(&_VotingManagerContract.TransactOpts, newOwner)
}

// VotingManagerContractAssetDelistedIterator is returned from FilterAssetDelisted and is used to iterate over the raw logs and unpacked data for AssetDelisted events raised by the VotingManagerContract contract.
type VotingManagerContractAssetDelistedIterator struct {
	Event *VotingManagerContractAssetDelisted // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractAssetDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractAssetDelisted)
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
		it.Event = new(VotingManagerContractAssetDelisted)
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
func (it *VotingManagerContractAssetDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractAssetDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractAssetDelisted represents a AssetDelisted event raised by the VotingManagerContract contract.
type VotingManagerContractAssetDelisted struct {
	AssetId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetDelisted is a free log retrieval operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterAssetDelisted(opts *bind.FilterOpts, assetId [][32]byte) (*VotingManagerContractAssetDelistedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "AssetDelisted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractAssetDelistedIterator{contract: _VotingManagerContract.contract, event: "AssetDelisted", logs: logs, sub: sub}, nil
}

// WatchAssetDelisted is a free log subscription operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchAssetDelisted(opts *bind.WatchOpts, sink chan<- *VotingManagerContractAssetDelisted, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "AssetDelisted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractAssetDelisted)
				if err := _VotingManagerContract.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
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

// ParseAssetDelisted is a log parse operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseAssetDelisted(log types.Log) (*VotingManagerContractAssetDelisted, error) {
	event := new(VotingManagerContractAssetDelisted)
	if err := _VotingManagerContract.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractAssetListedIterator is returned from FilterAssetListed and is used to iterate over the raw logs and unpacked data for AssetListed events raised by the VotingManagerContract contract.
type VotingManagerContractAssetListedIterator struct {
	Event *VotingManagerContractAssetListed // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractAssetListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractAssetListed)
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
		it.Event = new(VotingManagerContractAssetListed)
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
func (it *VotingManagerContractAssetListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractAssetListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractAssetListed represents a AssetListed event raised by the VotingManagerContract contract.
type VotingManagerContractAssetListed struct {
	AssetId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetListed is a free log retrieval operation binding the contract event 0x731530c0188d99aa7001e25e5fbeb8dbfce5cdedfe3d0a1c252b1ab4e1c3bd45.
//
// Solidity: event AssetListed(bytes32 indexed assetId)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterAssetListed(opts *bind.FilterOpts, assetId [][32]byte) (*VotingManagerContractAssetListedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "AssetListed", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractAssetListedIterator{contract: _VotingManagerContract.contract, event: "AssetListed", logs: logs, sub: sub}, nil
}

// WatchAssetListed is a free log subscription operation binding the contract event 0x731530c0188d99aa7001e25e5fbeb8dbfce5cdedfe3d0a1c252b1ab4e1c3bd45.
//
// Solidity: event AssetListed(bytes32 indexed assetId)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchAssetListed(opts *bind.WatchOpts, sink chan<- *VotingManagerContractAssetListed, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "AssetListed", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractAssetListed)
				if err := _VotingManagerContract.contract.UnpackLog(event, "AssetListed", log); err != nil {
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

// ParseAssetListed is a log parse operation binding the contract event 0x731530c0188d99aa7001e25e5fbeb8dbfce5cdedfe3d0a1c252b1ab4e1c3bd45.
//
// Solidity: event AssetListed(bytes32 indexed assetId)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseAssetListed(log types.Log) (*VotingManagerContractAssetListed, error) {
	event := new(VotingManagerContractAssetListed)
	if err := _VotingManagerContract.contract.UnpackLog(event, "AssetListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractDepositInfoSubmittedIterator is returned from FilterDepositInfoSubmitted and is used to iterate over the raw logs and unpacked data for DepositInfoSubmitted events raised by the VotingManagerContract contract.
type VotingManagerContractDepositInfoSubmittedIterator struct {
	Event *VotingManagerContractDepositInfoSubmitted // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractDepositInfoSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractDepositInfoSubmitted)
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
		it.Event = new(VotingManagerContractDepositInfoSubmitted)
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
func (it *VotingManagerContractDepositInfoSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractDepositInfoSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractDepositInfoSubmitted represents a DepositInfoSubmitted event raised by the VotingManagerContract contract.
type VotingManagerContractDepositInfoSubmitted struct {
	TargetAddress common.Address
	Amount        *big.Int
	TxInfo        []byte
	ChainId       *big.Int
	ExtraInfo     []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositInfoSubmitted is a free log retrieval operation binding the contract event 0x7bbd2c16415fef4cd9dcb330e9c3e9124c0d053f6c85075b29b5e3798fd4fea6.
//
// Solidity: event DepositInfoSubmitted(address indexed targetAddress, uint256 amount, bytes txInfo, uint256 chainId, bytes extraInfo)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterDepositInfoSubmitted(opts *bind.FilterOpts, targetAddress []common.Address) (*VotingManagerContractDepositInfoSubmittedIterator, error) {

	var targetAddressRule []interface{}
	for _, targetAddressItem := range targetAddress {
		targetAddressRule = append(targetAddressRule, targetAddressItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "DepositInfoSubmitted", targetAddressRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractDepositInfoSubmittedIterator{contract: _VotingManagerContract.contract, event: "DepositInfoSubmitted", logs: logs, sub: sub}, nil
}

// WatchDepositInfoSubmitted is a free log subscription operation binding the contract event 0x7bbd2c16415fef4cd9dcb330e9c3e9124c0d053f6c85075b29b5e3798fd4fea6.
//
// Solidity: event DepositInfoSubmitted(address indexed targetAddress, uint256 amount, bytes txInfo, uint256 chainId, bytes extraInfo)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchDepositInfoSubmitted(opts *bind.WatchOpts, sink chan<- *VotingManagerContractDepositInfoSubmitted, targetAddress []common.Address) (event.Subscription, error) {

	var targetAddressRule []interface{}
	for _, targetAddressItem := range targetAddress {
		targetAddressRule = append(targetAddressRule, targetAddressItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "DepositInfoSubmitted", targetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractDepositInfoSubmitted)
				if err := _VotingManagerContract.contract.UnpackLog(event, "DepositInfoSubmitted", log); err != nil {
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

// ParseDepositInfoSubmitted is a log parse operation binding the contract event 0x7bbd2c16415fef4cd9dcb330e9c3e9124c0d053f6c85075b29b5e3798fd4fea6.
//
// Solidity: event DepositInfoSubmitted(address indexed targetAddress, uint256 amount, bytes txInfo, uint256 chainId, bytes extraInfo)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseDepositInfoSubmitted(log types.Log) (*VotingManagerContractDepositInfoSubmitted, error) {
	event := new(VotingManagerContractDepositInfoSubmitted)
	if err := _VotingManagerContract.contract.UnpackLog(event, "DepositInfoSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// VotingManagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotingManagerContract contract.
type VotingManagerContractOwnershipTransferredIterator struct {
	Event *VotingManagerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractOwnershipTransferred)
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
		it.Event = new(VotingManagerContractOwnershipTransferred)
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
func (it *VotingManagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the VotingManagerContract contract.
type VotingManagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotingManagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractOwnershipTransferredIterator{contract: _VotingManagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotingManagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractOwnershipTransferred)
				if err := _VotingManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VotingManagerContract *VotingManagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*VotingManagerContractOwnershipTransferred, error) {
	event := new(VotingManagerContractOwnershipTransferred)
	if err := _VotingManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractParticipantAddedIterator is returned from FilterParticipantAdded and is used to iterate over the raw logs and unpacked data for ParticipantAdded events raised by the VotingManagerContract contract.
type VotingManagerContractParticipantAddedIterator struct {
	Event *VotingManagerContractParticipantAdded // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractParticipantAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractParticipantAdded)
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
		it.Event = new(VotingManagerContractParticipantAdded)
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
func (it *VotingManagerContractParticipantAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractParticipantAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractParticipantAdded represents a ParticipantAdded event raised by the VotingManagerContract contract.
type VotingManagerContractParticipantAdded struct {
	NewParticipant common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterParticipantAdded is a free log retrieval operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed newParticipant)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterParticipantAdded(opts *bind.FilterOpts, newParticipant []common.Address) (*VotingManagerContractParticipantAddedIterator, error) {

	var newParticipantRule []interface{}
	for _, newParticipantItem := range newParticipant {
		newParticipantRule = append(newParticipantRule, newParticipantItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "ParticipantAdded", newParticipantRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractParticipantAddedIterator{contract: _VotingManagerContract.contract, event: "ParticipantAdded", logs: logs, sub: sub}, nil
}

// WatchParticipantAdded is a free log subscription operation binding the contract event 0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b.
//
// Solidity: event ParticipantAdded(address indexed newParticipant)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchParticipantAdded(opts *bind.WatchOpts, sink chan<- *VotingManagerContractParticipantAdded, newParticipant []common.Address) (event.Subscription, error) {

	var newParticipantRule []interface{}
	for _, newParticipantItem := range newParticipant {
		newParticipantRule = append(newParticipantRule, newParticipantItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "ParticipantAdded", newParticipantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractParticipantAdded)
				if err := _VotingManagerContract.contract.UnpackLog(event, "ParticipantAdded", log); err != nil {
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
// Solidity: event ParticipantAdded(address indexed newParticipant)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseParticipantAdded(log types.Log) (*VotingManagerContractParticipantAdded, error) {
	event := new(VotingManagerContractParticipantAdded)
	if err := _VotingManagerContract.contract.UnpackLog(event, "ParticipantAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractParticipantRemovedIterator is returned from FilterParticipantRemoved and is used to iterate over the raw logs and unpacked data for ParticipantRemoved events raised by the VotingManagerContract contract.
type VotingManagerContractParticipantRemovedIterator struct {
	Event *VotingManagerContractParticipantRemoved // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractParticipantRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractParticipantRemoved)
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
		it.Event = new(VotingManagerContractParticipantRemoved)
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
func (it *VotingManagerContractParticipantRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractParticipantRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractParticipantRemoved represents a ParticipantRemoved event raised by the VotingManagerContract contract.
type VotingManagerContractParticipantRemoved struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantRemoved is a free log retrieval operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterParticipantRemoved(opts *bind.FilterOpts, participant []common.Address) (*VotingManagerContractParticipantRemovedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "ParticipantRemoved", participantRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractParticipantRemovedIterator{contract: _VotingManagerContract.contract, event: "ParticipantRemoved", logs: logs, sub: sub}, nil
}

// WatchParticipantRemoved is a free log subscription operation binding the contract event 0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc.
//
// Solidity: event ParticipantRemoved(address indexed participant)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchParticipantRemoved(opts *bind.WatchOpts, sink chan<- *VotingManagerContractParticipantRemoved, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "ParticipantRemoved", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractParticipantRemoved)
				if err := _VotingManagerContract.contract.UnpackLog(event, "ParticipantRemoved", log); err != nil {
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
func (_VotingManagerContract *VotingManagerContractFilterer) ParseParticipantRemoved(log types.Log) (*VotingManagerContractParticipantRemoved, error) {
	event := new(VotingManagerContractParticipantRemoved)
	if err := _VotingManagerContract.contract.UnpackLog(event, "ParticipantRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingManagerContractRewardPerPeriodVotedIterator is returned from FilterRewardPerPeriodVoted and is used to iterate over the raw logs and unpacked data for RewardPerPeriodVoted events raised by the VotingManagerContract contract.
type VotingManagerContractRewardPerPeriodVotedIterator struct {
	Event *VotingManagerContractRewardPerPeriodVoted // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractRewardPerPeriodVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractRewardPerPeriodVoted)
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
		it.Event = new(VotingManagerContractRewardPerPeriodVoted)
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
func (it *VotingManagerContractRewardPerPeriodVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractRewardPerPeriodVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractRewardPerPeriodVoted represents a RewardPerPeriodVoted event raised by the VotingManagerContract contract.
type VotingManagerContractRewardPerPeriodVoted struct {
	NewRewardPerPeriod *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardPerPeriodVoted is a free log retrieval operation binding the contract event 0x49c1ed88b0c8d97e2b2c5ed61e83a948d5ceb786bc26227eaab0cd748878cc66.
//
// Solidity: event RewardPerPeriodVoted(uint256 newRewardPerPeriod)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterRewardPerPeriodVoted(opts *bind.FilterOpts) (*VotingManagerContractRewardPerPeriodVotedIterator, error) {

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "RewardPerPeriodVoted")
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractRewardPerPeriodVotedIterator{contract: _VotingManagerContract.contract, event: "RewardPerPeriodVoted", logs: logs, sub: sub}, nil
}

// WatchRewardPerPeriodVoted is a free log subscription operation binding the contract event 0x49c1ed88b0c8d97e2b2c5ed61e83a948d5ceb786bc26227eaab0cd748878cc66.
//
// Solidity: event RewardPerPeriodVoted(uint256 newRewardPerPeriod)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchRewardPerPeriodVoted(opts *bind.WatchOpts, sink chan<- *VotingManagerContractRewardPerPeriodVoted) (event.Subscription, error) {

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "RewardPerPeriodVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractRewardPerPeriodVoted)
				if err := _VotingManagerContract.contract.UnpackLog(event, "RewardPerPeriodVoted", log); err != nil {
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

// ParseRewardPerPeriodVoted is a log parse operation binding the contract event 0x49c1ed88b0c8d97e2b2c5ed61e83a948d5ceb786bc26227eaab0cd748878cc66.
//
// Solidity: event RewardPerPeriodVoted(uint256 newRewardPerPeriod)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseRewardPerPeriodVoted(log types.Log) (*VotingManagerContractRewardPerPeriodVoted, error) {
	event := new(VotingManagerContractRewardPerPeriodVoted)
	if err := _VotingManagerContract.contract.UnpackLog(event, "RewardPerPeriodVoted", log); err != nil {
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

// VotingManagerContractTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the VotingManagerContract contract.
type VotingManagerContractTaskCompletedIterator struct {
	Event *VotingManagerContractTaskCompleted // Event containing the contract specifics and raw log

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
func (it *VotingManagerContractTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingManagerContractTaskCompleted)
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
		it.Event = new(VotingManagerContractTaskCompleted)
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
func (it *VotingManagerContractTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingManagerContractTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingManagerContractTaskCompleted represents a TaskCompleted event raised by the VotingManagerContract contract.
type VotingManagerContractTaskCompleted struct {
	TaskId      *big.Int
	Submitter   common.Address
	CompletedAt *big.Int
	TaskResult  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 completedAt, bytes taskResult)
func (_VotingManagerContract *VotingManagerContractFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int, submitter []common.Address) (*VotingManagerContractTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.FilterLogs(opts, "TaskCompleted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &VotingManagerContractTaskCompletedIterator{contract: _VotingManagerContract.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee.
//
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 completedAt, bytes taskResult)
func (_VotingManagerContract *VotingManagerContractFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *VotingManagerContractTaskCompleted, taskId []*big.Int, submitter []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _VotingManagerContract.contract.WatchLogs(opts, "TaskCompleted", taskIdRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingManagerContractTaskCompleted)
				if err := _VotingManagerContract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
// Solidity: event TaskCompleted(uint256 indexed taskId, address indexed submitter, uint256 completedAt, bytes taskResult)
func (_VotingManagerContract *VotingManagerContractFilterer) ParseTaskCompleted(log types.Log) (*VotingManagerContractTaskCompleted, error) {
	event := new(VotingManagerContractTaskCompleted)
	if err := _VotingManagerContract.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
