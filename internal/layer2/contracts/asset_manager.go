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

// IAssetManagerAsset is an auto generated low-level Go binding around an user-defined struct.
type IAssetManagerAsset struct {
	Name            string
	NuDexName       string
	AssetType       uint8
	ContractAddress common.Address
	ChainId         *big.Int
	IsListed        bool
}

// AssetManagerContractMetaData contains all meta data concerning the AssetManagerContract contract.
var AssetManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"assetList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assets\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nuDexName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isListed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delistAsset\",\"inputs\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetDetails\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAssetManager.Asset\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nuDexName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isListed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetIdentifier\",\"inputs\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAssetListed\",\"inputs\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listAsset\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nuDexName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetDelisted\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetListed\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"nuDexName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAssetManager.AssetType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// AssetManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetManagerContractMetaData.ABI instead.
var AssetManagerContractABI = AssetManagerContractMetaData.ABI

// AssetManagerContract is an auto generated Go binding around an Ethereum contract.
type AssetManagerContract struct {
	AssetManagerContractCaller     // Read-only binding to the contract
	AssetManagerContractTransactor // Write-only binding to the contract
	AssetManagerContractFilterer   // Log filterer for contract events
}

// AssetManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetManagerContractSession struct {
	Contract     *AssetManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AssetManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetManagerContractCallerSession struct {
	Contract *AssetManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AssetManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetManagerContractTransactorSession struct {
	Contract     *AssetManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AssetManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetManagerContractRaw struct {
	Contract *AssetManagerContract // Generic contract binding to access the raw methods on
}

// AssetManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetManagerContractCallerRaw struct {
	Contract *AssetManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// AssetManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetManagerContractTransactorRaw struct {
	Contract *AssetManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetManagerContract creates a new instance of AssetManagerContract, bound to a specific deployed contract.
func NewAssetManagerContract(address common.Address, backend bind.ContractBackend) (*AssetManagerContract, error) {
	contract, err := bindAssetManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContract{AssetManagerContractCaller: AssetManagerContractCaller{contract: contract}, AssetManagerContractTransactor: AssetManagerContractTransactor{contract: contract}, AssetManagerContractFilterer: AssetManagerContractFilterer{contract: contract}}, nil
}

// NewAssetManagerContractCaller creates a new read-only instance of AssetManagerContract, bound to a specific deployed contract.
func NewAssetManagerContractCaller(address common.Address, caller bind.ContractCaller) (*AssetManagerContractCaller, error) {
	contract, err := bindAssetManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractCaller{contract: contract}, nil
}

// NewAssetManagerContractTransactor creates a new write-only instance of AssetManagerContract, bound to a specific deployed contract.
func NewAssetManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetManagerContractTransactor, error) {
	contract, err := bindAssetManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractTransactor{contract: contract}, nil
}

// NewAssetManagerContractFilterer creates a new log filterer instance of AssetManagerContract, bound to a specific deployed contract.
func NewAssetManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetManagerContractFilterer, error) {
	contract, err := bindAssetManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractFilterer{contract: contract}, nil
}

// bindAssetManagerContract binds a generic wrapper to an already deployed contract.
func bindAssetManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManagerContract *AssetManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManagerContract.Contract.AssetManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManagerContract *AssetManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.AssetManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManagerContract *AssetManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.AssetManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManagerContract *AssetManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManagerContract *AssetManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManagerContract *AssetManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.contract.Transact(opts, method, params...)
}

// AssetList is a free data retrieval call binding the contract method 0xa0b4b301.
//
// Solidity: function assetList(uint256 ) view returns(bytes32)
func (_AssetManagerContract *AssetManagerContractCaller) AssetList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "assetList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetList is a free data retrieval call binding the contract method 0xa0b4b301.
//
// Solidity: function assetList(uint256 ) view returns(bytes32)
func (_AssetManagerContract *AssetManagerContractSession) AssetList(arg0 *big.Int) ([32]byte, error) {
	return _AssetManagerContract.Contract.AssetList(&_AssetManagerContract.CallOpts, arg0)
}

// AssetList is a free data retrieval call binding the contract method 0xa0b4b301.
//
// Solidity: function assetList(uint256 ) view returns(bytes32)
func (_AssetManagerContract *AssetManagerContractCallerSession) AssetList(arg0 *big.Int) ([32]byte, error) {
	return _AssetManagerContract.Contract.AssetList(&_AssetManagerContract.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) view returns(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bool isListed)
func (_AssetManagerContract *AssetManagerContractCaller) Assets(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Name            string
	NuDexName       string
	AssetType       uint8
	ContractAddress common.Address
	ChainId         *big.Int
	IsListed        bool
}, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		Name            string
		NuDexName       string
		AssetType       uint8
		ContractAddress common.Address
		ChainId         *big.Int
		IsListed        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.NuDexName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AssetType = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ContractAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsListed = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) view returns(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bool isListed)
func (_AssetManagerContract *AssetManagerContractSession) Assets(arg0 [32]byte) (struct {
	Name            string
	NuDexName       string
	AssetType       uint8
	ContractAddress common.Address
	ChainId         *big.Int
	IsListed        bool
}, error) {
	return _AssetManagerContract.Contract.Assets(&_AssetManagerContract.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) view returns(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId, bool isListed)
func (_AssetManagerContract *AssetManagerContractCallerSession) Assets(arg0 [32]byte) (struct {
	Name            string
	NuDexName       string
	AssetType       uint8
	ContractAddress common.Address
	ChainId         *big.Int
	IsListed        bool
}, error) {
	return _AssetManagerContract.Contract.Assets(&_AssetManagerContract.CallOpts, arg0)
}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns(bytes32[])
func (_AssetManagerContract *AssetManagerContractCaller) GetAllAssets(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "getAllAssets")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns(bytes32[])
func (_AssetManagerContract *AssetManagerContractSession) GetAllAssets() ([][32]byte, error) {
	return _AssetManagerContract.Contract.GetAllAssets(&_AssetManagerContract.CallOpts)
}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns(bytes32[])
func (_AssetManagerContract *AssetManagerContractCallerSession) GetAllAssets() ([][32]byte, error) {
	return _AssetManagerContract.Contract.GetAllAssets(&_AssetManagerContract.CallOpts)
}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 assetId) view returns((string,string,uint8,address,uint256,bool))
func (_AssetManagerContract *AssetManagerContractCaller) GetAssetDetails(opts *bind.CallOpts, assetId [32]byte) (IAssetManagerAsset, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "getAssetDetails", assetId)

	if err != nil {
		return *new(IAssetManagerAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(IAssetManagerAsset)).(*IAssetManagerAsset)

	return out0, err

}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 assetId) view returns((string,string,uint8,address,uint256,bool))
func (_AssetManagerContract *AssetManagerContractSession) GetAssetDetails(assetId [32]byte) (IAssetManagerAsset, error) {
	return _AssetManagerContract.Contract.GetAssetDetails(&_AssetManagerContract.CallOpts, assetId)
}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 assetId) view returns((string,string,uint8,address,uint256,bool))
func (_AssetManagerContract *AssetManagerContractCallerSession) GetAssetDetails(assetId [32]byte) (IAssetManagerAsset, error) {
	return _AssetManagerContract.Contract.GetAssetDetails(&_AssetManagerContract.CallOpts, assetId)
}

// GetAssetIdentifier is a free data retrieval call binding the contract method 0xcf008e17.
//
// Solidity: function getAssetIdentifier(uint8 assetType, address contractAddress, uint256 chainId) pure returns(bytes32)
func (_AssetManagerContract *AssetManagerContractCaller) GetAssetIdentifier(opts *bind.CallOpts, assetType uint8, contractAddress common.Address, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "getAssetIdentifier", assetType, contractAddress, chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAssetIdentifier is a free data retrieval call binding the contract method 0xcf008e17.
//
// Solidity: function getAssetIdentifier(uint8 assetType, address contractAddress, uint256 chainId) pure returns(bytes32)
func (_AssetManagerContract *AssetManagerContractSession) GetAssetIdentifier(assetType uint8, contractAddress common.Address, chainId *big.Int) ([32]byte, error) {
	return _AssetManagerContract.Contract.GetAssetIdentifier(&_AssetManagerContract.CallOpts, assetType, contractAddress, chainId)
}

// GetAssetIdentifier is a free data retrieval call binding the contract method 0xcf008e17.
//
// Solidity: function getAssetIdentifier(uint8 assetType, address contractAddress, uint256 chainId) pure returns(bytes32)
func (_AssetManagerContract *AssetManagerContractCallerSession) GetAssetIdentifier(assetType uint8, contractAddress common.Address, chainId *big.Int) ([32]byte, error) {
	return _AssetManagerContract.Contract.GetAssetIdentifier(&_AssetManagerContract.CallOpts, assetType, contractAddress, chainId)
}

// IsAssetListed is a free data retrieval call binding the contract method 0x615c548e.
//
// Solidity: function isAssetListed(uint8 assetType, address contractAddress, uint256 chainId) view returns(bool)
func (_AssetManagerContract *AssetManagerContractCaller) IsAssetListed(opts *bind.CallOpts, assetType uint8, contractAddress common.Address, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "isAssetListed", assetType, contractAddress, chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetListed is a free data retrieval call binding the contract method 0x615c548e.
//
// Solidity: function isAssetListed(uint8 assetType, address contractAddress, uint256 chainId) view returns(bool)
func (_AssetManagerContract *AssetManagerContractSession) IsAssetListed(assetType uint8, contractAddress common.Address, chainId *big.Int) (bool, error) {
	return _AssetManagerContract.Contract.IsAssetListed(&_AssetManagerContract.CallOpts, assetType, contractAddress, chainId)
}

// IsAssetListed is a free data retrieval call binding the contract method 0x615c548e.
//
// Solidity: function isAssetListed(uint8 assetType, address contractAddress, uint256 chainId) view returns(bool)
func (_AssetManagerContract *AssetManagerContractCallerSession) IsAssetListed(assetType uint8, contractAddress common.Address, chainId *big.Int) (bool, error) {
	return _AssetManagerContract.Contract.IsAssetListed(&_AssetManagerContract.CallOpts, assetType, contractAddress, chainId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetManagerContract *AssetManagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetManagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetManagerContract *AssetManagerContractSession) Owner() (common.Address, error) {
	return _AssetManagerContract.Contract.Owner(&_AssetManagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetManagerContract *AssetManagerContractCallerSession) Owner() (common.Address, error) {
	return _AssetManagerContract.Contract.Owner(&_AssetManagerContract.CallOpts)
}

// DelistAsset is a paid mutator transaction binding the contract method 0x4dfa1608.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId) returns()
func (_AssetManagerContract *AssetManagerContractTransactor) DelistAsset(opts *bind.TransactOpts, assetType uint8, contractAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _AssetManagerContract.contract.Transact(opts, "delistAsset", assetType, contractAddress, chainId)
}

// DelistAsset is a paid mutator transaction binding the contract method 0x4dfa1608.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId) returns()
func (_AssetManagerContract *AssetManagerContractSession) DelistAsset(assetType uint8, contractAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.DelistAsset(&_AssetManagerContract.TransactOpts, assetType, contractAddress, chainId)
}

// DelistAsset is a paid mutator transaction binding the contract method 0x4dfa1608.
//
// Solidity: function delistAsset(uint8 assetType, address contractAddress, uint256 chainId) returns()
func (_AssetManagerContract *AssetManagerContractTransactorSession) DelistAsset(assetType uint8, contractAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.DelistAsset(&_AssetManagerContract.TransactOpts, assetType, contractAddress, chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AssetManagerContract *AssetManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AssetManagerContract.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AssetManagerContract *AssetManagerContractSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.Initialize(&_AssetManagerContract.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_AssetManagerContract *AssetManagerContractTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.Initialize(&_AssetManagerContract.TransactOpts, _owner)
}

// ListAsset is a paid mutator transaction binding the contract method 0x4499b9bd.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId) returns()
func (_AssetManagerContract *AssetManagerContractTransactor) ListAsset(opts *bind.TransactOpts, name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _AssetManagerContract.contract.Transact(opts, "listAsset", name, nuDexName, assetType, contractAddress, chainId)
}

// ListAsset is a paid mutator transaction binding the contract method 0x4499b9bd.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId) returns()
func (_AssetManagerContract *AssetManagerContractSession) ListAsset(name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.ListAsset(&_AssetManagerContract.TransactOpts, name, nuDexName, assetType, contractAddress, chainId)
}

// ListAsset is a paid mutator transaction binding the contract method 0x4499b9bd.
//
// Solidity: function listAsset(string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId) returns()
func (_AssetManagerContract *AssetManagerContractTransactorSession) ListAsset(name string, nuDexName string, assetType uint8, contractAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.ListAsset(&_AssetManagerContract.TransactOpts, name, nuDexName, assetType, contractAddress, chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetManagerContract *AssetManagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetManagerContract *AssetManagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetManagerContract.Contract.RenounceOwnership(&_AssetManagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AssetManagerContract *AssetManagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AssetManagerContract.Contract.RenounceOwnership(&_AssetManagerContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetManagerContract *AssetManagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AssetManagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetManagerContract *AssetManagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.TransferOwnership(&_AssetManagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AssetManagerContract *AssetManagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AssetManagerContract.Contract.TransferOwnership(&_AssetManagerContract.TransactOpts, newOwner)
}

// AssetManagerContractAssetDelistedIterator is returned from FilterAssetDelisted and is used to iterate over the raw logs and unpacked data for AssetDelisted events raised by the AssetManagerContract contract.
type AssetManagerContractAssetDelistedIterator struct {
	Event *AssetManagerContractAssetDelisted // Event containing the contract specifics and raw log

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
func (it *AssetManagerContractAssetDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerContractAssetDelisted)
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
		it.Event = new(AssetManagerContractAssetDelisted)
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
func (it *AssetManagerContractAssetDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerContractAssetDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerContractAssetDelisted represents a AssetDelisted event raised by the AssetManagerContract contract.
type AssetManagerContractAssetDelisted struct {
	AssetId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetDelisted is a free log retrieval operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_AssetManagerContract *AssetManagerContractFilterer) FilterAssetDelisted(opts *bind.FilterOpts, assetId [][32]byte) (*AssetManagerContractAssetDelistedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetManagerContract.contract.FilterLogs(opts, "AssetDelisted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractAssetDelistedIterator{contract: _AssetManagerContract.contract, event: "AssetDelisted", logs: logs, sub: sub}, nil
}

// WatchAssetDelisted is a free log subscription operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed assetId)
func (_AssetManagerContract *AssetManagerContractFilterer) WatchAssetDelisted(opts *bind.WatchOpts, sink chan<- *AssetManagerContractAssetDelisted, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetManagerContract.contract.WatchLogs(opts, "AssetDelisted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerContractAssetDelisted)
				if err := _AssetManagerContract.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
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
func (_AssetManagerContract *AssetManagerContractFilterer) ParseAssetDelisted(log types.Log) (*AssetManagerContractAssetDelisted, error) {
	event := new(AssetManagerContractAssetDelisted)
	if err := _AssetManagerContract.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerContractAssetListedIterator is returned from FilterAssetListed and is used to iterate over the raw logs and unpacked data for AssetListed events raised by the AssetManagerContract contract.
type AssetManagerContractAssetListedIterator struct {
	Event *AssetManagerContractAssetListed // Event containing the contract specifics and raw log

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
func (it *AssetManagerContractAssetListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerContractAssetListed)
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
		it.Event = new(AssetManagerContractAssetListed)
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
func (it *AssetManagerContractAssetListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerContractAssetListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerContractAssetListed represents a AssetListed event raised by the AssetManagerContract contract.
type AssetManagerContractAssetListed struct {
	AssetId         [32]byte
	Name            string
	NuDexName       string
	AssetType       uint8
	ContractAddress common.Address
	ChainId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAssetListed is a free log retrieval operation binding the contract event 0x11efbdbc8e7ee6ca3aa4e2edb68a457967796d196fc76715c6609a387f12b0be.
//
// Solidity: event AssetListed(bytes32 indexed assetId, string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId)
func (_AssetManagerContract *AssetManagerContractFilterer) FilterAssetListed(opts *bind.FilterOpts, assetId [][32]byte) (*AssetManagerContractAssetListedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetManagerContract.contract.FilterLogs(opts, "AssetListed", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractAssetListedIterator{contract: _AssetManagerContract.contract, event: "AssetListed", logs: logs, sub: sub}, nil
}

// WatchAssetListed is a free log subscription operation binding the contract event 0x11efbdbc8e7ee6ca3aa4e2edb68a457967796d196fc76715c6609a387f12b0be.
//
// Solidity: event AssetListed(bytes32 indexed assetId, string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId)
func (_AssetManagerContract *AssetManagerContractFilterer) WatchAssetListed(opts *bind.WatchOpts, sink chan<- *AssetManagerContractAssetListed, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetManagerContract.contract.WatchLogs(opts, "AssetListed", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerContractAssetListed)
				if err := _AssetManagerContract.contract.UnpackLog(event, "AssetListed", log); err != nil {
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

// ParseAssetListed is a log parse operation binding the contract event 0x11efbdbc8e7ee6ca3aa4e2edb68a457967796d196fc76715c6609a387f12b0be.
//
// Solidity: event AssetListed(bytes32 indexed assetId, string name, string nuDexName, uint8 assetType, address contractAddress, uint256 chainId)
func (_AssetManagerContract *AssetManagerContractFilterer) ParseAssetListed(log types.Log) (*AssetManagerContractAssetListed, error) {
	event := new(AssetManagerContractAssetListed)
	if err := _AssetManagerContract.contract.UnpackLog(event, "AssetListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AssetManagerContract contract.
type AssetManagerContractInitializedIterator struct {
	Event *AssetManagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *AssetManagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerContractInitialized)
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
		it.Event = new(AssetManagerContractInitialized)
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
func (it *AssetManagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerContractInitialized represents a Initialized event raised by the AssetManagerContract contract.
type AssetManagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AssetManagerContract *AssetManagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*AssetManagerContractInitializedIterator, error) {

	logs, sub, err := _AssetManagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractInitializedIterator{contract: _AssetManagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AssetManagerContract *AssetManagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AssetManagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _AssetManagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerContractInitialized)
				if err := _AssetManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AssetManagerContract *AssetManagerContractFilterer) ParseInitialized(log types.Log) (*AssetManagerContractInitialized, error) {
	event := new(AssetManagerContractInitialized)
	if err := _AssetManagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AssetManagerContract contract.
type AssetManagerContractOwnershipTransferredIterator struct {
	Event *AssetManagerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AssetManagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerContractOwnershipTransferred)
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
		it.Event = new(AssetManagerContractOwnershipTransferred)
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
func (it *AssetManagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the AssetManagerContract contract.
type AssetManagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetManagerContract *AssetManagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AssetManagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetManagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerContractOwnershipTransferredIterator{contract: _AssetManagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AssetManagerContract *AssetManagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AssetManagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AssetManagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerContractOwnershipTransferred)
				if err := _AssetManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AssetManagerContract *AssetManagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*AssetManagerContractOwnershipTransferred, error) {
	event := new(AssetManagerContractOwnershipTransferred)
	if err := _AssetManagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
