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

// AssetParam is an auto generated low-level Go binding around an user-defined struct.
type AssetParam struct {
	AssetType         uint8
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	WithdrawFee       *big.Int
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
	AssetLogo         string
}

// NudexAsset is an auto generated low-level Go binding around an user-defined struct.
type NudexAsset struct {
	ListIndex         uint32
	AssetType         uint8
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	WithdrawFee       *big.Int
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
	AssetLogo         string
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	ChainId         *big.Int
	IsActive        bool
	AssetType       uint8
	Decimals        uint8
	ContractAddress common.Address
	Symbol          string
	Balance         *big.Int
}

// AssetHandlerContractMetaData contains all meta data concerning the AssetHandlerContract contract.
var AssetHandlerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"assetTickerList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consolidate\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delistAsset\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetDetails\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNudexAsset\",\"components\":[{\"name\":\"listIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isListed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetLogo\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAssetListed\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkToken\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_tokenInfos\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"linkedTokenList\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkedTokens\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listNewAsset\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_assetParam\",\"type\":\"tuple\",\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetLogo\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nudexAssets\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"listIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isListed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetLogo\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetlinkedToken\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenSwitch\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAsset\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_assetParam\",\"type\":\"tuple\",\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetLogo\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetDelisted\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetListed\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"assetParam\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetLogo\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetUpdated\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"assetParam\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumAssetType\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetLogo\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"assetIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"assetIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AssetNotListed\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assetIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// AssetHandlerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetHandlerContractMetaData.ABI instead.
var AssetHandlerContractABI = AssetHandlerContractMetaData.ABI

// AssetHandlerContract is an auto generated Go binding around an Ethereum contract.
type AssetHandlerContract struct {
	AssetHandlerContractCaller     // Read-only binding to the contract
	AssetHandlerContractTransactor // Write-only binding to the contract
	AssetHandlerContractFilterer   // Log filterer for contract events
}

// AssetHandlerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetHandlerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetHandlerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetHandlerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetHandlerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetHandlerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetHandlerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetHandlerContractSession struct {
	Contract     *AssetHandlerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AssetHandlerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetHandlerContractCallerSession struct {
	Contract *AssetHandlerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AssetHandlerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetHandlerContractTransactorSession struct {
	Contract     *AssetHandlerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AssetHandlerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetHandlerContractRaw struct {
	Contract *AssetHandlerContract // Generic contract binding to access the raw methods on
}

// AssetHandlerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetHandlerContractCallerRaw struct {
	Contract *AssetHandlerContractCaller // Generic read-only contract binding to access the raw methods on
}

// AssetHandlerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetHandlerContractTransactorRaw struct {
	Contract *AssetHandlerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetHandlerContract creates a new instance of AssetHandlerContract, bound to a specific deployed contract.
func NewAssetHandlerContract(address common.Address, backend bind.ContractBackend) (*AssetHandlerContract, error) {
	contract, err := bindAssetHandlerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContract{AssetHandlerContractCaller: AssetHandlerContractCaller{contract: contract}, AssetHandlerContractTransactor: AssetHandlerContractTransactor{contract: contract}, AssetHandlerContractFilterer: AssetHandlerContractFilterer{contract: contract}}, nil
}

// NewAssetHandlerContractCaller creates a new read-only instance of AssetHandlerContract, bound to a specific deployed contract.
func NewAssetHandlerContractCaller(address common.Address, caller bind.ContractCaller) (*AssetHandlerContractCaller, error) {
	contract, err := bindAssetHandlerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractCaller{contract: contract}, nil
}

// NewAssetHandlerContractTransactor creates a new write-only instance of AssetHandlerContract, bound to a specific deployed contract.
func NewAssetHandlerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetHandlerContractTransactor, error) {
	contract, err := bindAssetHandlerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractTransactor{contract: contract}, nil
}

// NewAssetHandlerContractFilterer creates a new log filterer instance of AssetHandlerContract, bound to a specific deployed contract.
func NewAssetHandlerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetHandlerContractFilterer, error) {
	contract, err := bindAssetHandlerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractFilterer{contract: contract}, nil
}

// bindAssetHandlerContract binds a generic wrapper to an already deployed contract.
func bindAssetHandlerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetHandlerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetHandlerContract *AssetHandlerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetHandlerContract.Contract.AssetHandlerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetHandlerContract *AssetHandlerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.AssetHandlerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetHandlerContract *AssetHandlerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.AssetHandlerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetHandlerContract *AssetHandlerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetHandlerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetHandlerContract *AssetHandlerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetHandlerContract *AssetHandlerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.contract.Transact(opts, method, params...)
}

// AssetTickerList is a free data retrieval call binding the contract method 0x75e4f746.
//
// Solidity: function assetTickerList(uint256 ) view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCaller) AssetTickerList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "assetTickerList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetTickerList is a free data retrieval call binding the contract method 0x75e4f746.
//
// Solidity: function assetTickerList(uint256 ) view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractSession) AssetTickerList(arg0 *big.Int) ([32]byte, error) {
	return _AssetHandlerContract.Contract.AssetTickerList(&_AssetHandlerContract.CallOpts, arg0)
}

// AssetTickerList is a free data retrieval call binding the contract method 0x75e4f746.
//
// Solidity: function assetTickerList(uint256 ) view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) AssetTickerList(arg0 *big.Int) ([32]byte, error) {
	return _AssetHandlerContract.Contract.AssetTickerList(&_AssetHandlerContract.CallOpts, arg0)
}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns(bytes32[])
func (_AssetHandlerContract *AssetHandlerContractCaller) GetAllAssets(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "getAllAssets")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns(bytes32[])
func (_AssetHandlerContract *AssetHandlerContractSession) GetAllAssets() ([][32]byte, error) {
	return _AssetHandlerContract.Contract.GetAllAssets(&_AssetHandlerContract.CallOpts)
}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns(bytes32[])
func (_AssetHandlerContract *AssetHandlerContractCallerSession) GetAllAssets() ([][32]byte, error) {
	return _AssetHandlerContract.Contract.GetAllAssets(&_AssetHandlerContract.CallOpts)
}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 _ticker) view returns((uint32,uint8,uint8,bool,bool,bool,uint32,uint32,uint256,uint256,uint256,string,string))
func (_AssetHandlerContract *AssetHandlerContractCaller) GetAssetDetails(opts *bind.CallOpts, _ticker [32]byte) (NudexAsset, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "getAssetDetails", _ticker)

	if err != nil {
		return *new(NudexAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(NudexAsset)).(*NudexAsset)

	return out0, err

}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 _ticker) view returns((uint32,uint8,uint8,bool,bool,bool,uint32,uint32,uint256,uint256,uint256,string,string))
func (_AssetHandlerContract *AssetHandlerContractSession) GetAssetDetails(_ticker [32]byte) (NudexAsset, error) {
	return _AssetHandlerContract.Contract.GetAssetDetails(&_AssetHandlerContract.CallOpts, _ticker)
}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 _ticker) view returns((uint32,uint8,uint8,bool,bool,bool,uint32,uint32,uint256,uint256,uint256,string,string))
func (_AssetHandlerContract *AssetHandlerContractCallerSession) GetAssetDetails(_ticker [32]byte) (NudexAsset, error) {
	return _AssetHandlerContract.Contract.GetAssetDetails(&_AssetHandlerContract.CallOpts, _ticker)
}

// IsAssetListed is a free data retrieval call binding the contract method 0x7581b94c.
//
// Solidity: function isAssetListed(bytes32 _ticker) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractCaller) IsAssetListed(opts *bind.CallOpts, _ticker [32]byte) (bool, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "isAssetListed", _ticker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetListed is a free data retrieval call binding the contract method 0x7581b94c.
//
// Solidity: function isAssetListed(bytes32 _ticker) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractSession) IsAssetListed(_ticker [32]byte) (bool, error) {
	return _AssetHandlerContract.Contract.IsAssetListed(&_AssetHandlerContract.CallOpts, _ticker)
}

// IsAssetListed is a free data retrieval call binding the contract method 0x7581b94c.
//
// Solidity: function isAssetListed(bytes32 _ticker) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) IsAssetListed(_ticker [32]byte) (bool, error) {
	return _AssetHandlerContract.Contract.IsAssetListed(&_AssetHandlerContract.CallOpts, _ticker)
}

// LinkedTokenList is a free data retrieval call binding the contract method 0xae109286.
//
// Solidity: function linkedTokenList(bytes32 ticker, uint256 ) view returns(uint256 chainIds)
func (_AssetHandlerContract *AssetHandlerContractCaller) LinkedTokenList(opts *bind.CallOpts, ticker [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "linkedTokenList", ticker, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LinkedTokenList is a free data retrieval call binding the contract method 0xae109286.
//
// Solidity: function linkedTokenList(bytes32 ticker, uint256 ) view returns(uint256 chainIds)
func (_AssetHandlerContract *AssetHandlerContractSession) LinkedTokenList(ticker [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _AssetHandlerContract.Contract.LinkedTokenList(&_AssetHandlerContract.CallOpts, ticker, arg1)
}

// LinkedTokenList is a free data retrieval call binding the contract method 0xae109286.
//
// Solidity: function linkedTokenList(bytes32 ticker, uint256 ) view returns(uint256 chainIds)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) LinkedTokenList(ticker [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _AssetHandlerContract.Contract.LinkedTokenList(&_AssetHandlerContract.CallOpts, ticker, arg1)
}

// LinkedTokens is a free data retrieval call binding the contract method 0xb5a32b98.
//
// Solidity: function linkedTokens(bytes32 ticker, uint256 chainId) view returns(uint256 chainId, bool isActive, uint8 assetType, uint8 decimals, address contractAddress, string symbol, uint256 balance)
func (_AssetHandlerContract *AssetHandlerContractCaller) LinkedTokens(opts *bind.CallOpts, ticker [32]byte, chainId *big.Int) (struct {
	ChainId         *big.Int
	IsActive        bool
	AssetType       uint8
	Decimals        uint8
	ContractAddress common.Address
	Symbol          string
	Balance         *big.Int
}, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "linkedTokens", ticker, chainId)

	outstruct := new(struct {
		ChainId         *big.Int
		IsActive        bool
		AssetType       uint8
		Decimals        uint8
		ContractAddress common.Address
		Symbol          string
		Balance         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.AssetType = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.ContractAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Symbol = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Balance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LinkedTokens is a free data retrieval call binding the contract method 0xb5a32b98.
//
// Solidity: function linkedTokens(bytes32 ticker, uint256 chainId) view returns(uint256 chainId, bool isActive, uint8 assetType, uint8 decimals, address contractAddress, string symbol, uint256 balance)
func (_AssetHandlerContract *AssetHandlerContractSession) LinkedTokens(ticker [32]byte, chainId *big.Int) (struct {
	ChainId         *big.Int
	IsActive        bool
	AssetType       uint8
	Decimals        uint8
	ContractAddress common.Address
	Symbol          string
	Balance         *big.Int
}, error) {
	return _AssetHandlerContract.Contract.LinkedTokens(&_AssetHandlerContract.CallOpts, ticker, chainId)
}

// LinkedTokens is a free data retrieval call binding the contract method 0xb5a32b98.
//
// Solidity: function linkedTokens(bytes32 ticker, uint256 chainId) view returns(uint256 chainId, bool isActive, uint8 assetType, uint8 decimals, address contractAddress, string symbol, uint256 balance)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) LinkedTokens(ticker [32]byte, chainId *big.Int) (struct {
	ChainId         *big.Int
	IsActive        bool
	AssetType       uint8
	Decimals        uint8
	ContractAddress common.Address
	Symbol          string
	Balance         *big.Int
}, error) {
	return _AssetHandlerContract.Contract.LinkedTokens(&_AssetHandlerContract.CallOpts, ticker, chainId)
}

// NudexAssets is a free data retrieval call binding the contract method 0x79979120.
//
// Solidity: function nudexAssets(bytes32 ticker) view returns(uint32 listIndex, uint8 assetType, uint8 decimals, bool depositEnabled, bool withdrawalEnabled, bool isListed, uint32 createdTime, uint32 updatedTime, uint256 withdrawFee, uint256 minDepositAmount, uint256 minWithdrawAmount, string assetAlias, string assetLogo)
func (_AssetHandlerContract *AssetHandlerContractCaller) NudexAssets(opts *bind.CallOpts, ticker [32]byte) (struct {
	ListIndex         uint32
	AssetType         uint8
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	WithdrawFee       *big.Int
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
	AssetLogo         string
}, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "nudexAssets", ticker)

	outstruct := new(struct {
		ListIndex         uint32
		AssetType         uint8
		Decimals          uint8
		DepositEnabled    bool
		WithdrawalEnabled bool
		IsListed          bool
		CreatedTime       uint32
		UpdatedTime       uint32
		WithdrawFee       *big.Int
		MinDepositAmount  *big.Int
		MinWithdrawAmount *big.Int
		AssetAlias        string
		AssetLogo         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListIndex = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.AssetType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Decimals = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.DepositEnabled = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.WithdrawalEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.IsListed = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.CreatedTime = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.UpdatedTime = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.WithdrawFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MinDepositAmount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.MinWithdrawAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.AssetAlias = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.AssetLogo = *abi.ConvertType(out[12], new(string)).(*string)

	return *outstruct, err

}

// NudexAssets is a free data retrieval call binding the contract method 0x79979120.
//
// Solidity: function nudexAssets(bytes32 ticker) view returns(uint32 listIndex, uint8 assetType, uint8 decimals, bool depositEnabled, bool withdrawalEnabled, bool isListed, uint32 createdTime, uint32 updatedTime, uint256 withdrawFee, uint256 minDepositAmount, uint256 minWithdrawAmount, string assetAlias, string assetLogo)
func (_AssetHandlerContract *AssetHandlerContractSession) NudexAssets(ticker [32]byte) (struct {
	ListIndex         uint32
	AssetType         uint8
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	WithdrawFee       *big.Int
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
	AssetLogo         string
}, error) {
	return _AssetHandlerContract.Contract.NudexAssets(&_AssetHandlerContract.CallOpts, ticker)
}

// NudexAssets is a free data retrieval call binding the contract method 0x79979120.
//
// Solidity: function nudexAssets(bytes32 ticker) view returns(uint32 listIndex, uint8 assetType, uint8 decimals, bool depositEnabled, bool withdrawalEnabled, bool isListed, uint32 createdTime, uint32 updatedTime, uint256 withdrawFee, uint256 minDepositAmount, uint256 minWithdrawAmount, string assetAlias, string assetLogo)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) NudexAssets(ticker [32]byte) (struct {
	ListIndex         uint32
	AssetType         uint8
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	WithdrawFee       *big.Int
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
	AssetLogo         string
}, error) {
	return _AssetHandlerContract.Contract.NudexAssets(&_AssetHandlerContract.CallOpts, ticker)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetHandlerContract *AssetHandlerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetHandlerContract *AssetHandlerContractSession) Owner() (common.Address, error) {
	return _AssetHandlerContract.Contract.Owner(&_AssetHandlerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) Owner() (common.Address, error) {
	return _AssetHandlerContract.Contract.Owner(&_AssetHandlerContract.CallOpts)
}

// Consolidate is a paid mutator transaction binding the contract method 0xb68a5e3e.
//
// Solidity: function consolidate(bytes32 _ticker, uint256 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) Consolidate(opts *bind.TransactOpts, _ticker [32]byte, _chainId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "consolidate", _ticker, _chainId, _amount)
}

// Consolidate is a paid mutator transaction binding the contract method 0xb68a5e3e.
//
// Solidity: function consolidate(bytes32 _ticker, uint256 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) Consolidate(_ticker [32]byte, _chainId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Consolidate(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _amount)
}

// Consolidate is a paid mutator transaction binding the contract method 0xb68a5e3e.
//
// Solidity: function consolidate(bytes32 _ticker, uint256 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) Consolidate(_ticker [32]byte, _chainId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Consolidate(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _amount)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe06e60a5.
//
// Solidity: function delistAsset(bytes32 _ticker) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) DelistAsset(opts *bind.TransactOpts, _ticker [32]byte) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "delistAsset", _ticker)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe06e60a5.
//
// Solidity: function delistAsset(bytes32 _ticker) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) DelistAsset(_ticker [32]byte) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.DelistAsset(&_AssetHandlerContract.TransactOpts, _ticker)
}

// DelistAsset is a paid mutator transaction binding the contract method 0xe06e60a5.
//
// Solidity: function delistAsset(bytes32 _ticker) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) DelistAsset(_ticker [32]byte) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.DelistAsset(&_AssetHandlerContract.TransactOpts, _ticker)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Initialize(&_AssetHandlerContract.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Initialize(&_AssetHandlerContract.TransactOpts, _owner)
}

// LinkToken is a paid mutator transaction binding the contract method 0x63875efc.
//
// Solidity: function linkToken(bytes32 _ticker, (uint256,bool,uint8,uint8,address,string,uint256)[] _tokenInfos) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) LinkToken(opts *bind.TransactOpts, _ticker [32]byte, _tokenInfos []TokenInfo) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "linkToken", _ticker, _tokenInfos)
}

// LinkToken is a paid mutator transaction binding the contract method 0x63875efc.
//
// Solidity: function linkToken(bytes32 _ticker, (uint256,bool,uint8,uint8,address,string,uint256)[] _tokenInfos) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) LinkToken(_ticker [32]byte, _tokenInfos []TokenInfo) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.LinkToken(&_AssetHandlerContract.TransactOpts, _ticker, _tokenInfos)
}

// LinkToken is a paid mutator transaction binding the contract method 0x63875efc.
//
// Solidity: function linkToken(bytes32 _ticker, (uint256,bool,uint8,uint8,address,string,uint256)[] _tokenInfos) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) LinkToken(_ticker [32]byte, _tokenInfos []TokenInfo) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.LinkToken(&_AssetHandlerContract.TransactOpts, _ticker, _tokenInfos)
}

// ListNewAsset is a paid mutator transaction binding the contract method 0x0c834ce0.
//
// Solidity: function listNewAsset(bytes32 _ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) ListNewAsset(opts *bind.TransactOpts, _ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "listNewAsset", _ticker, _assetParam)
}

// ListNewAsset is a paid mutator transaction binding the contract method 0x0c834ce0.
//
// Solidity: function listNewAsset(bytes32 _ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) ListNewAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.ListNewAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// ListNewAsset is a paid mutator transaction binding the contract method 0x0c834ce0.
//
// Solidity: function listNewAsset(bytes32 _ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) ListNewAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.ListNewAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetHandlerContract *AssetHandlerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.RenounceOwnership(&_AssetHandlerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.RenounceOwnership(&_AssetHandlerContract.TransactOpts)
}

// ResetlinkedToken is a paid mutator transaction binding the contract method 0x9bded0ec.
//
// Solidity: function resetlinkedToken(bytes32 _ticker) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) ResetlinkedToken(opts *bind.TransactOpts, _ticker [32]byte) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "resetlinkedToken", _ticker)
}

// ResetlinkedToken is a paid mutator transaction binding the contract method 0x9bded0ec.
//
// Solidity: function resetlinkedToken(bytes32 _ticker) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) ResetlinkedToken(_ticker [32]byte) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.ResetlinkedToken(&_AssetHandlerContract.TransactOpts, _ticker)
}

// ResetlinkedToken is a paid mutator transaction binding the contract method 0x9bded0ec.
//
// Solidity: function resetlinkedToken(bytes32 _ticker) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) ResetlinkedToken(_ticker [32]byte) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.ResetlinkedToken(&_AssetHandlerContract.TransactOpts, _ticker)
}

// TokenSwitch is a paid mutator transaction binding the contract method 0x6965feef.
//
// Solidity: function tokenSwitch(bytes32 _ticker, uint256 _chainId, bool _isActive) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) TokenSwitch(opts *bind.TransactOpts, _ticker [32]byte, _chainId *big.Int, _isActive bool) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "tokenSwitch", _ticker, _chainId, _isActive)
}

// TokenSwitch is a paid mutator transaction binding the contract method 0x6965feef.
//
// Solidity: function tokenSwitch(bytes32 _ticker, uint256 _chainId, bool _isActive) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) TokenSwitch(_ticker [32]byte, _chainId *big.Int, _isActive bool) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.TokenSwitch(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _isActive)
}

// TokenSwitch is a paid mutator transaction binding the contract method 0x6965feef.
//
// Solidity: function tokenSwitch(bytes32 _ticker, uint256 _chainId, bool _isActive) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) TokenSwitch(_ticker [32]byte, _chainId *big.Int, _isActive bool) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.TokenSwitch(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _isActive)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.TransferOwnership(&_AssetHandlerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.TransferOwnership(&_AssetHandlerContract.TransactOpts, newOwner)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x874ed715.
//
// Solidity: function updateAsset(bytes32 _ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) UpdateAsset(opts *bind.TransactOpts, _ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "updateAsset", _ticker, _assetParam)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x874ed715.
//
// Solidity: function updateAsset(bytes32 _ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) UpdateAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.UpdateAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x874ed715.
//
// Solidity: function updateAsset(bytes32 _ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) UpdateAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.UpdateAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5c334e8f.
//
// Solidity: function withdraw(bytes32 _ticker, uint256 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) Withdraw(opts *bind.TransactOpts, _ticker [32]byte, _chainId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "withdraw", _ticker, _chainId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5c334e8f.
//
// Solidity: function withdraw(bytes32 _ticker, uint256 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) Withdraw(_ticker [32]byte, _chainId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Withdraw(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5c334e8f.
//
// Solidity: function withdraw(bytes32 _ticker, uint256 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) Withdraw(_ticker [32]byte, _chainId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Withdraw(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _amount)
}

// AssetHandlerContractAssetDelistedIterator is returned from FilterAssetDelisted and is used to iterate over the raw logs and unpacked data for AssetDelisted events raised by the AssetHandlerContract contract.
type AssetHandlerContractAssetDelistedIterator struct {
	Event *AssetHandlerContractAssetDelisted // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractAssetDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractAssetDelisted)
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
		it.Event = new(AssetHandlerContractAssetDelisted)
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
func (it *AssetHandlerContractAssetDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractAssetDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractAssetDelisted represents a AssetDelisted event raised by the AssetHandlerContract contract.
type AssetHandlerContractAssetDelisted struct {
	AssetId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetDelisted is a free log retrieval operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterAssetDelisted(opts *bind.FilterOpts, assetId [][32]byte) (*AssetHandlerContractAssetDelistedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "AssetDelisted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractAssetDelistedIterator{contract: _AssetHandlerContract.contract, event: "AssetDelisted", logs: logs, sub: sub}, nil
}

// WatchAssetDelisted is a free log subscription operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchAssetDelisted(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractAssetDelisted, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "AssetDelisted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractAssetDelisted)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
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
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseAssetDelisted(log types.Log) (*AssetHandlerContractAssetDelisted, error) {
	event := new(AssetHandlerContractAssetDelisted)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractAssetListedIterator is returned from FilterAssetListed and is used to iterate over the raw logs and unpacked data for AssetListed events raised by the AssetHandlerContract contract.
type AssetHandlerContractAssetListedIterator struct {
	Event *AssetHandlerContractAssetListed // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractAssetListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractAssetListed)
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
		it.Event = new(AssetHandlerContractAssetListed)
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
func (it *AssetHandlerContractAssetListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractAssetListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractAssetListed represents a AssetListed event raised by the AssetHandlerContract contract.
type AssetHandlerContractAssetListed struct {
	Ticker     [32]byte
	AssetParam AssetParam
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetListed is a free log retrieval operation binding the contract event 0x47a32dea00882b59c5ae47fa19013c8c2fa5c7183864cfe26810382eb223406c.
//
// Solidity: event AssetListed(bytes32 indexed ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterAssetListed(opts *bind.FilterOpts, ticker [][32]byte) (*AssetHandlerContractAssetListedIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "AssetListed", tickerRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractAssetListedIterator{contract: _AssetHandlerContract.contract, event: "AssetListed", logs: logs, sub: sub}, nil
}

// WatchAssetListed is a free log subscription operation binding the contract event 0x47a32dea00882b59c5ae47fa19013c8c2fa5c7183864cfe26810382eb223406c.
//
// Solidity: event AssetListed(bytes32 indexed ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchAssetListed(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractAssetListed, ticker [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "AssetListed", tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractAssetListed)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetListed", log); err != nil {
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

// ParseAssetListed is a log parse operation binding the contract event 0x47a32dea00882b59c5ae47fa19013c8c2fa5c7183864cfe26810382eb223406c.
//
// Solidity: event AssetListed(bytes32 indexed ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseAssetListed(log types.Log) (*AssetHandlerContractAssetListed, error) {
	event := new(AssetHandlerContractAssetListed)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractAssetUpdatedIterator is returned from FilterAssetUpdated and is used to iterate over the raw logs and unpacked data for AssetUpdated events raised by the AssetHandlerContract contract.
type AssetHandlerContractAssetUpdatedIterator struct {
	Event *AssetHandlerContractAssetUpdated // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractAssetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractAssetUpdated)
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
		it.Event = new(AssetHandlerContractAssetUpdated)
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
func (it *AssetHandlerContractAssetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractAssetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractAssetUpdated represents a AssetUpdated event raised by the AssetHandlerContract contract.
type AssetHandlerContractAssetUpdated struct {
	Ticker     [32]byte
	AssetParam AssetParam
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetUpdated is a free log retrieval operation binding the contract event 0xd40d00347e52d79ab828acdc3f68593b7286254edbdd8553defcc2e3c6bf74f1.
//
// Solidity: event AssetUpdated(bytes32 indexed ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterAssetUpdated(opts *bind.FilterOpts, ticker [][32]byte) (*AssetHandlerContractAssetUpdatedIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "AssetUpdated", tickerRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractAssetUpdatedIterator{contract: _AssetHandlerContract.contract, event: "AssetUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetUpdated is a free log subscription operation binding the contract event 0xd40d00347e52d79ab828acdc3f68593b7286254edbdd8553defcc2e3c6bf74f1.
//
// Solidity: event AssetUpdated(bytes32 indexed ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchAssetUpdated(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractAssetUpdated, ticker [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "AssetUpdated", tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractAssetUpdated)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
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

// ParseAssetUpdated is a log parse operation binding the contract event 0xd40d00347e52d79ab828acdc3f68593b7286254edbdd8553defcc2e3c6bf74f1.
//
// Solidity: event AssetUpdated(bytes32 indexed ticker, (uint8,uint8,bool,bool,uint256,uint256,uint256,string,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseAssetUpdated(log types.Log) (*AssetHandlerContractAssetUpdated, error) {
	event := new(AssetHandlerContractAssetUpdated)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AssetHandlerContract contract.
type AssetHandlerContractDepositIterator struct {
	Event *AssetHandlerContractDeposit // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractDeposit)
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
		it.Event = new(AssetHandlerContractDeposit)
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
func (it *AssetHandlerContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractDeposit represents a Deposit event raised by the AssetHandlerContract contract.
type AssetHandlerContractDeposit struct {
	AssetId    [32]byte
	AssetIndex *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1f1096fd8bc7d572fb7ad7e4102736b6615500975c0252ea91ef1b765c49897.
//
// Solidity: event Deposit(bytes32 indexed assetId, uint256 indexed assetIndex, uint256 indexed amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterDeposit(opts *bind.FilterOpts, assetId [][32]byte, assetIndex []*big.Int, amount []*big.Int) (*AssetHandlerContractDepositIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetIndexRule []interface{}
	for _, assetIndexItem := range assetIndex {
		assetIndexRule = append(assetIndexRule, assetIndexItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "Deposit", assetIdRule, assetIndexRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractDepositIterator{contract: _AssetHandlerContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1f1096fd8bc7d572fb7ad7e4102736b6615500975c0252ea91ef1b765c49897.
//
// Solidity: event Deposit(bytes32 indexed assetId, uint256 indexed assetIndex, uint256 indexed amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractDeposit, assetId [][32]byte, assetIndex []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetIndexRule []interface{}
	for _, assetIndexItem := range assetIndex {
		assetIndexRule = append(assetIndexRule, assetIndexItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "Deposit", assetIdRule, assetIndexRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractDeposit)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1f1096fd8bc7d572fb7ad7e4102736b6615500975c0252ea91ef1b765c49897.
//
// Solidity: event Deposit(bytes32 indexed assetId, uint256 indexed assetIndex, uint256 indexed amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseDeposit(log types.Log) (*AssetHandlerContractDeposit, error) {
	event := new(AssetHandlerContractDeposit)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AssetHandlerContract contract.
type AssetHandlerContractInitializedIterator struct {
	Event *AssetHandlerContractInitialized // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractInitialized)
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
		it.Event = new(AssetHandlerContractInitialized)
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
func (it *AssetHandlerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractInitialized represents a Initialized event raised by the AssetHandlerContract contract.
type AssetHandlerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*AssetHandlerContractInitializedIterator, error) {

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractInitializedIterator{contract: _AssetHandlerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractInitialized)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseInitialized(log types.Log) (*AssetHandlerContractInitialized, error) {
	event := new(AssetHandlerContractInitialized)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AssetHandlerContract contract.
type AssetHandlerContractOwnershipTransferredIterator struct {
	Event *AssetHandlerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractOwnershipTransferred)
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
		it.Event = new(AssetHandlerContractOwnershipTransferred)
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
func (it *AssetHandlerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractOwnershipTransferred represents a OwnershipTransferred event raised by the AssetHandlerContract contract.
type AssetHandlerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AssetHandlerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractOwnershipTransferredIterator{contract: _AssetHandlerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractOwnershipTransferred)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseOwnershipTransferred(log types.Log) (*AssetHandlerContractOwnershipTransferred, error) {
	event := new(AssetHandlerContractOwnershipTransferred)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AssetHandlerContract contract.
type AssetHandlerContractWithdrawIterator struct {
	Event *AssetHandlerContractWithdraw // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractWithdraw)
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
		it.Event = new(AssetHandlerContractWithdraw)
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
func (it *AssetHandlerContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractWithdraw represents a Withdraw event raised by the AssetHandlerContract contract.
type AssetHandlerContractWithdraw struct {
	AssetId    [32]byte
	AssetIndex *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xcaa03f03cdee1c04d0a939a3664af76dafeb30903f32620a78e7c7cafa15202b.
//
// Solidity: event Withdraw(bytes32 indexed assetId, uint256 indexed assetIndex, uint256 indexed amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterWithdraw(opts *bind.FilterOpts, assetId [][32]byte, assetIndex []*big.Int, amount []*big.Int) (*AssetHandlerContractWithdrawIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetIndexRule []interface{}
	for _, assetIndexItem := range assetIndex {
		assetIndexRule = append(assetIndexRule, assetIndexItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "Withdraw", assetIdRule, assetIndexRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractWithdrawIterator{contract: _AssetHandlerContract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xcaa03f03cdee1c04d0a939a3664af76dafeb30903f32620a78e7c7cafa15202b.
//
// Solidity: event Withdraw(bytes32 indexed assetId, uint256 indexed assetIndex, uint256 indexed amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractWithdraw, assetId [][32]byte, assetIndex []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetIndexRule []interface{}
	for _, assetIndexItem := range assetIndex {
		assetIndexRule = append(assetIndexRule, assetIndexItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "Withdraw", assetIdRule, assetIndexRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractWithdraw)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xcaa03f03cdee1c04d0a939a3664af76dafeb30903f32620a78e7c7cafa15202b.
//
// Solidity: event Withdraw(bytes32 indexed assetId, uint256 indexed assetIndex, uint256 indexed amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseWithdraw(log types.Log) (*AssetHandlerContractWithdraw, error) {
	event := new(AssetHandlerContractWithdraw)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
