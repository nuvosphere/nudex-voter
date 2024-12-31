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
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
}

// ConsolidateTaskParam is an auto generated low-level Go binding around an user-defined struct.
type ConsolidateTaskParam struct {
	FromAddr []string
	Ticker   [32]byte
	ChainId  [32]byte
	Amount   *big.Int
}

// NudexAsset is an auto generated low-level Go binding around an user-defined struct.
type NudexAsset struct {
	ListIndex         uint32
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	ChainId         [32]byte
	IsActive        bool
	Decimals        uint8
	ContractAddress string
	Symbol          string
	WithdrawFee     *big.Int
}

// AssetHandlerContractMetaData contains all meta data concerning the AssetHandlerContract contract.
var AssetHandlerContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTRYPOINT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNDS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SUBMITTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assetTickerList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consolidate\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structConsolidateTaskParam\",\"components\":[{\"name\":\"fromAddr\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delistAsset\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllLinkedTokens\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetDetails\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNudexAsset\",\"components\":[{\"name\":\"listIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isListed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedToken\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAssetListed\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkToken\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_tokenInfos\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"linkedTokenList\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainIds\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkedTokens\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listNewAsset\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_assetParam\",\"type\":\"tuple\",\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nudexAssets\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"listIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isListed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"createdTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetlinkedToken\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitAssetTask\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitConsolidateTask\",\"inputs\":[{\"name\":\"_params\",\"type\":\"tuple[]\",\"internalType\":\"structConsolidateTaskParam[]\",\"components\":[{\"name\":\"fromAddr\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitListAssetTask\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_assetParam\",\"type\":\"tuple\",\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenSwitch\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAsset\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_assetParam\",\"type\":\"tuple\",\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetDelisted\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetListed\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"assetParam\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetUpdated\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"assetParam\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAssetParam\",\"components\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depositEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"withdrawalEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetAlias\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Consolidate\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fromAddr\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LinkToken\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"contractAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResetLinkedToken\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSwitch\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AssetNotListed\",\"inputs\":[{\"name\":\"ticker\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.DEFAULTADMINROLE(&_AssetHandlerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.DEFAULTADMINROLE(&_AssetHandlerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCaller) ENTRYPOINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "ENTRYPOINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.ENTRYPOINTROLE(&_AssetHandlerContract.CallOpts)
}

// ENTRYPOINTROLE is a free data retrieval call binding the contract method 0x5445bd5d.
//
// Solidity: function ENTRYPOINT_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) ENTRYPOINTROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.ENTRYPOINTROLE(&_AssetHandlerContract.CallOpts)
}

// FUNDSROLE is a free data retrieval call binding the contract method 0x9b87ab6d.
//
// Solidity: function FUNDS_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCaller) FUNDSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "FUNDS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FUNDSROLE is a free data retrieval call binding the contract method 0x9b87ab6d.
//
// Solidity: function FUNDS_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractSession) FUNDSROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.FUNDSROLE(&_AssetHandlerContract.CallOpts)
}

// FUNDSROLE is a free data retrieval call binding the contract method 0x9b87ab6d.
//
// Solidity: function FUNDS_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) FUNDSROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.FUNDSROLE(&_AssetHandlerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCaller) SUBMITTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "SUBMITTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractSession) SUBMITTERROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.SUBMITTERROLE(&_AssetHandlerContract.CallOpts)
}

// SUBMITTERROLE is a free data retrieval call binding the contract method 0x91712a0b.
//
// Solidity: function SUBMITTER_ROLE() view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) SUBMITTERROLE() ([32]byte, error) {
	return _AssetHandlerContract.Contract.SUBMITTERROLE(&_AssetHandlerContract.CallOpts)
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

// GetAllLinkedTokens is a free data retrieval call binding the contract method 0xa0966e3e.
//
// Solidity: function getAllLinkedTokens(bytes32 _ticker) view returns(bytes32[])
func (_AssetHandlerContract *AssetHandlerContractCaller) GetAllLinkedTokens(opts *bind.CallOpts, _ticker [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "getAllLinkedTokens", _ticker)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllLinkedTokens is a free data retrieval call binding the contract method 0xa0966e3e.
//
// Solidity: function getAllLinkedTokens(bytes32 _ticker) view returns(bytes32[])
func (_AssetHandlerContract *AssetHandlerContractSession) GetAllLinkedTokens(_ticker [32]byte) ([][32]byte, error) {
	return _AssetHandlerContract.Contract.GetAllLinkedTokens(&_AssetHandlerContract.CallOpts, _ticker)
}

// GetAllLinkedTokens is a free data retrieval call binding the contract method 0xa0966e3e.
//
// Solidity: function getAllLinkedTokens(bytes32 _ticker) view returns(bytes32[])
func (_AssetHandlerContract *AssetHandlerContractCallerSession) GetAllLinkedTokens(_ticker [32]byte) ([][32]byte, error) {
	return _AssetHandlerContract.Contract.GetAllLinkedTokens(&_AssetHandlerContract.CallOpts, _ticker)
}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 _ticker) view returns((uint32,uint8,bool,bool,bool,uint32,uint32,uint256,uint256,string))
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
// Solidity: function getAssetDetails(bytes32 _ticker) view returns((uint32,uint8,bool,bool,bool,uint32,uint32,uint256,uint256,string))
func (_AssetHandlerContract *AssetHandlerContractSession) GetAssetDetails(_ticker [32]byte) (NudexAsset, error) {
	return _AssetHandlerContract.Contract.GetAssetDetails(&_AssetHandlerContract.CallOpts, _ticker)
}

// GetAssetDetails is a free data retrieval call binding the contract method 0xc4381674.
//
// Solidity: function getAssetDetails(bytes32 _ticker) view returns((uint32,uint8,bool,bool,bool,uint32,uint32,uint256,uint256,string))
func (_AssetHandlerContract *AssetHandlerContractCallerSession) GetAssetDetails(_ticker [32]byte) (NudexAsset, error) {
	return _AssetHandlerContract.Contract.GetAssetDetails(&_AssetHandlerContract.CallOpts, _ticker)
}

// GetLinkedToken is a free data retrieval call binding the contract method 0x2180b8f3.
//
// Solidity: function getLinkedToken(bytes32 _ticker, bytes32 _chainId) view returns((bytes32,bool,uint8,string,string,uint256))
func (_AssetHandlerContract *AssetHandlerContractCaller) GetLinkedToken(opts *bind.CallOpts, _ticker [32]byte, _chainId [32]byte) (TokenInfo, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "getLinkedToken", _ticker, _chainId)

	if err != nil {
		return *new(TokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenInfo)).(*TokenInfo)

	return out0, err

}

// GetLinkedToken is a free data retrieval call binding the contract method 0x2180b8f3.
//
// Solidity: function getLinkedToken(bytes32 _ticker, bytes32 _chainId) view returns((bytes32,bool,uint8,string,string,uint256))
func (_AssetHandlerContract *AssetHandlerContractSession) GetLinkedToken(_ticker [32]byte, _chainId [32]byte) (TokenInfo, error) {
	return _AssetHandlerContract.Contract.GetLinkedToken(&_AssetHandlerContract.CallOpts, _ticker, _chainId)
}

// GetLinkedToken is a free data retrieval call binding the contract method 0x2180b8f3.
//
// Solidity: function getLinkedToken(bytes32 _ticker, bytes32 _chainId) view returns((bytes32,bool,uint8,string,string,uint256))
func (_AssetHandlerContract *AssetHandlerContractCallerSession) GetLinkedToken(_ticker [32]byte, _chainId [32]byte) (TokenInfo, error) {
	return _AssetHandlerContract.Contract.GetLinkedToken(&_AssetHandlerContract.CallOpts, _ticker, _chainId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetHandlerContract.Contract.GetRoleAdmin(&_AssetHandlerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetHandlerContract.Contract.GetRoleAdmin(&_AssetHandlerContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetHandlerContract.Contract.HasRole(&_AssetHandlerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetHandlerContract.Contract.HasRole(&_AssetHandlerContract.CallOpts, role, account)
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
// Solidity: function linkedTokenList(bytes32 ticker, uint256 ) view returns(bytes32 chainIds)
func (_AssetHandlerContract *AssetHandlerContractCaller) LinkedTokenList(opts *bind.CallOpts, ticker [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "linkedTokenList", ticker, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LinkedTokenList is a free data retrieval call binding the contract method 0xae109286.
//
// Solidity: function linkedTokenList(bytes32 ticker, uint256 ) view returns(bytes32 chainIds)
func (_AssetHandlerContract *AssetHandlerContractSession) LinkedTokenList(ticker [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _AssetHandlerContract.Contract.LinkedTokenList(&_AssetHandlerContract.CallOpts, ticker, arg1)
}

// LinkedTokenList is a free data retrieval call binding the contract method 0xae109286.
//
// Solidity: function linkedTokenList(bytes32 ticker, uint256 ) view returns(bytes32 chainIds)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) LinkedTokenList(ticker [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _AssetHandlerContract.Contract.LinkedTokenList(&_AssetHandlerContract.CallOpts, ticker, arg1)
}

// LinkedTokens is a free data retrieval call binding the contract method 0x4dcb66fb.
//
// Solidity: function linkedTokens(bytes32 ticker, bytes32 chainId) view returns(bytes32 chainId, bool isActive, uint8 decimals, string contractAddress, string symbol, uint256 withdrawFee)
func (_AssetHandlerContract *AssetHandlerContractCaller) LinkedTokens(opts *bind.CallOpts, ticker [32]byte, chainId [32]byte) (struct {
	ChainId         [32]byte
	IsActive        bool
	Decimals        uint8
	ContractAddress string
	Symbol          string
	WithdrawFee     *big.Int
}, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "linkedTokens", ticker, chainId)

	outstruct := new(struct {
		ChainId         [32]byte
		IsActive        bool
		Decimals        uint8
		ContractAddress string
		Symbol          string
		WithdrawFee     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Decimals = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ContractAddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.WithdrawFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LinkedTokens is a free data retrieval call binding the contract method 0x4dcb66fb.
//
// Solidity: function linkedTokens(bytes32 ticker, bytes32 chainId) view returns(bytes32 chainId, bool isActive, uint8 decimals, string contractAddress, string symbol, uint256 withdrawFee)
func (_AssetHandlerContract *AssetHandlerContractSession) LinkedTokens(ticker [32]byte, chainId [32]byte) (struct {
	ChainId         [32]byte
	IsActive        bool
	Decimals        uint8
	ContractAddress string
	Symbol          string
	WithdrawFee     *big.Int
}, error) {
	return _AssetHandlerContract.Contract.LinkedTokens(&_AssetHandlerContract.CallOpts, ticker, chainId)
}

// LinkedTokens is a free data retrieval call binding the contract method 0x4dcb66fb.
//
// Solidity: function linkedTokens(bytes32 ticker, bytes32 chainId) view returns(bytes32 chainId, bool isActive, uint8 decimals, string contractAddress, string symbol, uint256 withdrawFee)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) LinkedTokens(ticker [32]byte, chainId [32]byte) (struct {
	ChainId         [32]byte
	IsActive        bool
	Decimals        uint8
	ContractAddress string
	Symbol          string
	WithdrawFee     *big.Int
}, error) {
	return _AssetHandlerContract.Contract.LinkedTokens(&_AssetHandlerContract.CallOpts, ticker, chainId)
}

// NudexAssets is a free data retrieval call binding the contract method 0x79979120.
//
// Solidity: function nudexAssets(bytes32 ticker) view returns(uint32 listIndex, uint8 decimals, bool depositEnabled, bool withdrawalEnabled, bool isListed, uint32 createdTime, uint32 updatedTime, uint256 minDepositAmount, uint256 minWithdrawAmount, string assetAlias)
func (_AssetHandlerContract *AssetHandlerContractCaller) NudexAssets(opts *bind.CallOpts, ticker [32]byte) (struct {
	ListIndex         uint32
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
}, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "nudexAssets", ticker)

	outstruct := new(struct {
		ListIndex         uint32
		Decimals          uint8
		DepositEnabled    bool
		WithdrawalEnabled bool
		IsListed          bool
		CreatedTime       uint32
		UpdatedTime       uint32
		MinDepositAmount  *big.Int
		MinWithdrawAmount *big.Int
		AssetAlias        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListIndex = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.DepositEnabled = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.WithdrawalEnabled = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsListed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.CreatedTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.UpdatedTime = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.MinDepositAmount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MinWithdrawAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.AssetAlias = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// NudexAssets is a free data retrieval call binding the contract method 0x79979120.
//
// Solidity: function nudexAssets(bytes32 ticker) view returns(uint32 listIndex, uint8 decimals, bool depositEnabled, bool withdrawalEnabled, bool isListed, uint32 createdTime, uint32 updatedTime, uint256 minDepositAmount, uint256 minWithdrawAmount, string assetAlias)
func (_AssetHandlerContract *AssetHandlerContractSession) NudexAssets(ticker [32]byte) (struct {
	ListIndex         uint32
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
}, error) {
	return _AssetHandlerContract.Contract.NudexAssets(&_AssetHandlerContract.CallOpts, ticker)
}

// NudexAssets is a free data retrieval call binding the contract method 0x79979120.
//
// Solidity: function nudexAssets(bytes32 ticker) view returns(uint32 listIndex, uint8 decimals, bool depositEnabled, bool withdrawalEnabled, bool isListed, uint32 createdTime, uint32 updatedTime, uint256 minDepositAmount, uint256 minWithdrawAmount, string assetAlias)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) NudexAssets(ticker [32]byte) (struct {
	ListIndex         uint32
	Decimals          uint8
	DepositEnabled    bool
	WithdrawalEnabled bool
	IsListed          bool
	CreatedTime       uint32
	UpdatedTime       uint32
	MinDepositAmount  *big.Int
	MinWithdrawAmount *big.Int
	AssetAlias        string
}, error) {
	return _AssetHandlerContract.Contract.NudexAssets(&_AssetHandlerContract.CallOpts, ticker)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetHandlerContract.Contract.SupportsInterface(&_AssetHandlerContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetHandlerContract.Contract.SupportsInterface(&_AssetHandlerContract.CallOpts, interfaceId)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_AssetHandlerContract *AssetHandlerContractCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetHandlerContract.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_AssetHandlerContract *AssetHandlerContractSession) TaskManager() (common.Address, error) {
	return _AssetHandlerContract.Contract.TaskManager(&_AssetHandlerContract.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_AssetHandlerContract *AssetHandlerContractCallerSession) TaskManager() (common.Address, error) {
	return _AssetHandlerContract.Contract.TaskManager(&_AssetHandlerContract.CallOpts)
}

// Consolidate is a paid mutator transaction binding the contract method 0x1b391425.
//
// Solidity: function consolidate((string[],bytes32,bytes32,uint256) _param) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) Consolidate(opts *bind.TransactOpts, _param ConsolidateTaskParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "consolidate", _param)
}

// Consolidate is a paid mutator transaction binding the contract method 0x1b391425.
//
// Solidity: function consolidate((string[],bytes32,bytes32,uint256) _param) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) Consolidate(_param ConsolidateTaskParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Consolidate(&_AssetHandlerContract.TransactOpts, _param)
}

// Consolidate is a paid mutator transaction binding the contract method 0x1b391425.
//
// Solidity: function consolidate((string[],bytes32,bytes32,uint256) _param) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) Consolidate(_param ConsolidateTaskParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Consolidate(&_AssetHandlerContract.TransactOpts, _param)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.GrantRole(&_AssetHandlerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.GrantRole(&_AssetHandlerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "initialize", _owner, _entryPoint, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Initialize(&_AssetHandlerContract.TransactOpts, _owner, _entryPoint, _submitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _entryPoint, address _submitter) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) Initialize(_owner common.Address, _entryPoint common.Address, _submitter common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Initialize(&_AssetHandlerContract.TransactOpts, _owner, _entryPoint, _submitter)
}

// LinkToken is a paid mutator transaction binding the contract method 0x457f39f8.
//
// Solidity: function linkToken(bytes32 _ticker, (bytes32,bool,uint8,string,string,uint256)[] _tokenInfos) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) LinkToken(opts *bind.TransactOpts, _ticker [32]byte, _tokenInfos []TokenInfo) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "linkToken", _ticker, _tokenInfos)
}

// LinkToken is a paid mutator transaction binding the contract method 0x457f39f8.
//
// Solidity: function linkToken(bytes32 _ticker, (bytes32,bool,uint8,string,string,uint256)[] _tokenInfos) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) LinkToken(_ticker [32]byte, _tokenInfos []TokenInfo) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.LinkToken(&_AssetHandlerContract.TransactOpts, _ticker, _tokenInfos)
}

// LinkToken is a paid mutator transaction binding the contract method 0x457f39f8.
//
// Solidity: function linkToken(bytes32 _ticker, (bytes32,bool,uint8,string,string,uint256)[] _tokenInfos) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) LinkToken(_ticker [32]byte, _tokenInfos []TokenInfo) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.LinkToken(&_AssetHandlerContract.TransactOpts, _ticker, _tokenInfos)
}

// ListNewAsset is a paid mutator transaction binding the contract method 0xf131ea5b.
//
// Solidity: function listNewAsset(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) ListNewAsset(opts *bind.TransactOpts, _ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "listNewAsset", _ticker, _assetParam)
}

// ListNewAsset is a paid mutator transaction binding the contract method 0xf131ea5b.
//
// Solidity: function listNewAsset(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) ListNewAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.ListNewAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// ListNewAsset is a paid mutator transaction binding the contract method 0xf131ea5b.
//
// Solidity: function listNewAsset(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) ListNewAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.ListNewAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.RenounceRole(&_AssetHandlerContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.RenounceRole(&_AssetHandlerContract.TransactOpts, role, callerConfirmation)
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

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.RevokeRole(&_AssetHandlerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.RevokeRole(&_AssetHandlerContract.TransactOpts, role, account)
}

// SubmitAssetTask is a paid mutator transaction binding the contract method 0x32b0708b.
//
// Solidity: function submitAssetTask(bytes32 _ticker, bytes _callData) returns(uint64)
func (_AssetHandlerContract *AssetHandlerContractTransactor) SubmitAssetTask(opts *bind.TransactOpts, _ticker [32]byte, _callData []byte) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "submitAssetTask", _ticker, _callData)
}

// SubmitAssetTask is a paid mutator transaction binding the contract method 0x32b0708b.
//
// Solidity: function submitAssetTask(bytes32 _ticker, bytes _callData) returns(uint64)
func (_AssetHandlerContract *AssetHandlerContractSession) SubmitAssetTask(_ticker [32]byte, _callData []byte) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.SubmitAssetTask(&_AssetHandlerContract.TransactOpts, _ticker, _callData)
}

// SubmitAssetTask is a paid mutator transaction binding the contract method 0x32b0708b.
//
// Solidity: function submitAssetTask(bytes32 _ticker, bytes _callData) returns(uint64)
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) SubmitAssetTask(_ticker [32]byte, _callData []byte) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.SubmitAssetTask(&_AssetHandlerContract.TransactOpts, _ticker, _callData)
}

// SubmitConsolidateTask is a paid mutator transaction binding the contract method 0x4700ca07.
//
// Solidity: function submitConsolidateTask((string[],bytes32,bytes32,uint256)[] _params) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) SubmitConsolidateTask(opts *bind.TransactOpts, _params []ConsolidateTaskParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "submitConsolidateTask", _params)
}

// SubmitConsolidateTask is a paid mutator transaction binding the contract method 0x4700ca07.
//
// Solidity: function submitConsolidateTask((string[],bytes32,bytes32,uint256)[] _params) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) SubmitConsolidateTask(_params []ConsolidateTaskParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.SubmitConsolidateTask(&_AssetHandlerContract.TransactOpts, _params)
}

// SubmitConsolidateTask is a paid mutator transaction binding the contract method 0x4700ca07.
//
// Solidity: function submitConsolidateTask((string[],bytes32,bytes32,uint256)[] _params) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) SubmitConsolidateTask(_params []ConsolidateTaskParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.SubmitConsolidateTask(&_AssetHandlerContract.TransactOpts, _params)
}

// SubmitListAssetTask is a paid mutator transaction binding the contract method 0xe1550b8c.
//
// Solidity: function submitListAssetTask(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns(uint64)
func (_AssetHandlerContract *AssetHandlerContractTransactor) SubmitListAssetTask(opts *bind.TransactOpts, _ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "submitListAssetTask", _ticker, _assetParam)
}

// SubmitListAssetTask is a paid mutator transaction binding the contract method 0xe1550b8c.
//
// Solidity: function submitListAssetTask(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns(uint64)
func (_AssetHandlerContract *AssetHandlerContractSession) SubmitListAssetTask(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.SubmitListAssetTask(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// SubmitListAssetTask is a paid mutator transaction binding the contract method 0xe1550b8c.
//
// Solidity: function submitListAssetTask(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns(uint64)
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) SubmitListAssetTask(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.SubmitListAssetTask(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// TokenSwitch is a paid mutator transaction binding the contract method 0x18ea92aa.
//
// Solidity: function tokenSwitch(bytes32 _ticker, bytes32 _chainId, bool _isActive) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) TokenSwitch(opts *bind.TransactOpts, _ticker [32]byte, _chainId [32]byte, _isActive bool) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "tokenSwitch", _ticker, _chainId, _isActive)
}

// TokenSwitch is a paid mutator transaction binding the contract method 0x18ea92aa.
//
// Solidity: function tokenSwitch(bytes32 _ticker, bytes32 _chainId, bool _isActive) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) TokenSwitch(_ticker [32]byte, _chainId [32]byte, _isActive bool) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.TokenSwitch(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _isActive)
}

// TokenSwitch is a paid mutator transaction binding the contract method 0x18ea92aa.
//
// Solidity: function tokenSwitch(bytes32 _ticker, bytes32 _chainId, bool _isActive) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) TokenSwitch(_ticker [32]byte, _chainId [32]byte, _isActive bool) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.TokenSwitch(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _isActive)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0xa15a1e77.
//
// Solidity: function updateAsset(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) UpdateAsset(opts *bind.TransactOpts, _ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "updateAsset", _ticker, _assetParam)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0xa15a1e77.
//
// Solidity: function updateAsset(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) UpdateAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.UpdateAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0xa15a1e77.
//
// Solidity: function updateAsset(bytes32 _ticker, (uint8,bool,bool,uint256,uint256,string) _assetParam) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) UpdateAsset(_ticker [32]byte, _assetParam AssetParam) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.UpdateAsset(&_AssetHandlerContract.TransactOpts, _ticker, _assetParam)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcf0d9098.
//
// Solidity: function withdraw(bytes32 _ticker, bytes32 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactor) Withdraw(opts *bind.TransactOpts, _ticker [32]byte, _chainId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.contract.Transact(opts, "withdraw", _ticker, _chainId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcf0d9098.
//
// Solidity: function withdraw(bytes32 _ticker, bytes32 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractSession) Withdraw(_ticker [32]byte, _chainId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _AssetHandlerContract.Contract.Withdraw(&_AssetHandlerContract.TransactOpts, _ticker, _chainId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcf0d9098.
//
// Solidity: function withdraw(bytes32 _ticker, bytes32 _chainId, uint256 _amount) returns()
func (_AssetHandlerContract *AssetHandlerContractTransactorSession) Withdraw(_ticker [32]byte, _chainId [32]byte, _amount *big.Int) (*types.Transaction, error) {
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
	Ticker [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetDelisted is a free log retrieval operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed ticker)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterAssetDelisted(opts *bind.FilterOpts, ticker [][32]byte) (*AssetHandlerContractAssetDelistedIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "AssetDelisted", tickerRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractAssetDelistedIterator{contract: _AssetHandlerContract.contract, event: "AssetDelisted", logs: logs, sub: sub}, nil
}

// WatchAssetDelisted is a free log subscription operation binding the contract event 0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6.
//
// Solidity: event AssetDelisted(bytes32 indexed ticker)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchAssetDelisted(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractAssetDelisted, ticker [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "AssetDelisted", tickerRule)
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
// Solidity: event AssetDelisted(bytes32 indexed ticker)
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

// FilterAssetListed is a free log retrieval operation binding the contract event 0x9a321153d23feb212bd45898d567f7472a2ecc6c752d8e6757ea0914ab2b7009.
//
// Solidity: event AssetListed(bytes32 indexed ticker, (uint8,bool,bool,uint256,uint256,string) assetParam)
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

// WatchAssetListed is a free log subscription operation binding the contract event 0x9a321153d23feb212bd45898d567f7472a2ecc6c752d8e6757ea0914ab2b7009.
//
// Solidity: event AssetListed(bytes32 indexed ticker, (uint8,bool,bool,uint256,uint256,string) assetParam)
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

// ParseAssetListed is a log parse operation binding the contract event 0x9a321153d23feb212bd45898d567f7472a2ecc6c752d8e6757ea0914ab2b7009.
//
// Solidity: event AssetListed(bytes32 indexed ticker, (uint8,bool,bool,uint256,uint256,string) assetParam)
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

// FilterAssetUpdated is a free log retrieval operation binding the contract event 0xc91c349180cb8d82f404c5b3cb776276676bba19867657b1e45e84e3103d0b36.
//
// Solidity: event AssetUpdated(bytes32 indexed ticker, (uint8,bool,bool,uint256,uint256,string) assetParam)
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

// WatchAssetUpdated is a free log subscription operation binding the contract event 0xc91c349180cb8d82f404c5b3cb776276676bba19867657b1e45e84e3103d0b36.
//
// Solidity: event AssetUpdated(bytes32 indexed ticker, (uint8,bool,bool,uint256,uint256,string) assetParam)
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

// ParseAssetUpdated is a log parse operation binding the contract event 0xc91c349180cb8d82f404c5b3cb776276676bba19867657b1e45e84e3103d0b36.
//
// Solidity: event AssetUpdated(bytes32 indexed ticker, (uint8,bool,bool,uint256,uint256,string) assetParam)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseAssetUpdated(log types.Log) (*AssetHandlerContractAssetUpdated, error) {
	event := new(AssetHandlerContractAssetUpdated)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractConsolidateIterator is returned from FilterConsolidate and is used to iterate over the raw logs and unpacked data for Consolidate events raised by the AssetHandlerContract contract.
type AssetHandlerContractConsolidateIterator struct {
	Event *AssetHandlerContractConsolidate // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractConsolidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractConsolidate)
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
		it.Event = new(AssetHandlerContractConsolidate)
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
func (it *AssetHandlerContractConsolidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractConsolidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractConsolidate represents a Consolidate event raised by the AssetHandlerContract contract.
type AssetHandlerContractConsolidate struct {
	Ticker   [32]byte
	ChainId  [32]byte
	Amount   *big.Int
	FromAddr []string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsolidate is a free log retrieval operation binding the contract event 0x1a8dfb1b33ec086667484356f91a2c6ba31efa82f32375459055863764f3d927.
//
// Solidity: event Consolidate(bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount, string[] fromAddr)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterConsolidate(opts *bind.FilterOpts, ticker [][32]byte, chainId [][32]byte) (*AssetHandlerContractConsolidateIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "Consolidate", tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractConsolidateIterator{contract: _AssetHandlerContract.contract, event: "Consolidate", logs: logs, sub: sub}, nil
}

// WatchConsolidate is a free log subscription operation binding the contract event 0x1a8dfb1b33ec086667484356f91a2c6ba31efa82f32375459055863764f3d927.
//
// Solidity: event Consolidate(bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount, string[] fromAddr)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchConsolidate(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractConsolidate, ticker [][32]byte, chainId [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "Consolidate", tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractConsolidate)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "Consolidate", log); err != nil {
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

// ParseConsolidate is a log parse operation binding the contract event 0x1a8dfb1b33ec086667484356f91a2c6ba31efa82f32375459055863764f3d927.
//
// Solidity: event Consolidate(bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount, string[] fromAddr)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseConsolidate(log types.Log) (*AssetHandlerContractConsolidate, error) {
	event := new(AssetHandlerContractConsolidate)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "Consolidate", log); err != nil {
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

// AssetHandlerContractLinkTokenIterator is returned from FilterLinkToken and is used to iterate over the raw logs and unpacked data for LinkToken events raised by the AssetHandlerContract contract.
type AssetHandlerContractLinkTokenIterator struct {
	Event *AssetHandlerContractLinkToken // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractLinkTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractLinkToken)
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
		it.Event = new(AssetHandlerContractLinkToken)
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
func (it *AssetHandlerContractLinkTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractLinkTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractLinkToken represents a LinkToken event raised by the AssetHandlerContract contract.
type AssetHandlerContractLinkToken struct {
	Ticker [32]byte
	Tokens []TokenInfo
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLinkToken is a free log retrieval operation binding the contract event 0x8ebde2f255b6646fbdd0aa775c6c067d015c7537705c818c486d48deee03d69f.
//
// Solidity: event LinkToken(bytes32 indexed ticker, (bytes32,bool,uint8,string,string,uint256)[] tokens)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterLinkToken(opts *bind.FilterOpts, ticker [][32]byte) (*AssetHandlerContractLinkTokenIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "LinkToken", tickerRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractLinkTokenIterator{contract: _AssetHandlerContract.contract, event: "LinkToken", logs: logs, sub: sub}, nil
}

// WatchLinkToken is a free log subscription operation binding the contract event 0x8ebde2f255b6646fbdd0aa775c6c067d015c7537705c818c486d48deee03d69f.
//
// Solidity: event LinkToken(bytes32 indexed ticker, (bytes32,bool,uint8,string,string,uint256)[] tokens)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchLinkToken(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractLinkToken, ticker [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "LinkToken", tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractLinkToken)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "LinkToken", log); err != nil {
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

// ParseLinkToken is a log parse operation binding the contract event 0x8ebde2f255b6646fbdd0aa775c6c067d015c7537705c818c486d48deee03d69f.
//
// Solidity: event LinkToken(bytes32 indexed ticker, (bytes32,bool,uint8,string,string,uint256)[] tokens)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseLinkToken(log types.Log) (*AssetHandlerContractLinkToken, error) {
	event := new(AssetHandlerContractLinkToken)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "LinkToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractResetLinkedTokenIterator is returned from FilterResetLinkedToken and is used to iterate over the raw logs and unpacked data for ResetLinkedToken events raised by the AssetHandlerContract contract.
type AssetHandlerContractResetLinkedTokenIterator struct {
	Event *AssetHandlerContractResetLinkedToken // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractResetLinkedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractResetLinkedToken)
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
		it.Event = new(AssetHandlerContractResetLinkedToken)
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
func (it *AssetHandlerContractResetLinkedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractResetLinkedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractResetLinkedToken represents a ResetLinkedToken event raised by the AssetHandlerContract contract.
type AssetHandlerContractResetLinkedToken struct {
	Ticker [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResetLinkedToken is a free log retrieval operation binding the contract event 0x1f036e69048c80897a67509c6ce12379bd2f893606404ec226c2db215f15eb32.
//
// Solidity: event ResetLinkedToken(bytes32 indexed ticker)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterResetLinkedToken(opts *bind.FilterOpts, ticker [][32]byte) (*AssetHandlerContractResetLinkedTokenIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "ResetLinkedToken", tickerRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractResetLinkedTokenIterator{contract: _AssetHandlerContract.contract, event: "ResetLinkedToken", logs: logs, sub: sub}, nil
}

// WatchResetLinkedToken is a free log subscription operation binding the contract event 0x1f036e69048c80897a67509c6ce12379bd2f893606404ec226c2db215f15eb32.
//
// Solidity: event ResetLinkedToken(bytes32 indexed ticker)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchResetLinkedToken(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractResetLinkedToken, ticker [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "ResetLinkedToken", tickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractResetLinkedToken)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "ResetLinkedToken", log); err != nil {
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

// ParseResetLinkedToken is a log parse operation binding the contract event 0x1f036e69048c80897a67509c6ce12379bd2f893606404ec226c2db215f15eb32.
//
// Solidity: event ResetLinkedToken(bytes32 indexed ticker)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseResetLinkedToken(log types.Log) (*AssetHandlerContractResetLinkedToken, error) {
	event := new(AssetHandlerContractResetLinkedToken)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "ResetLinkedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AssetHandlerContract contract.
type AssetHandlerContractRoleAdminChangedIterator struct {
	Event *AssetHandlerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractRoleAdminChanged)
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
		it.Event = new(AssetHandlerContractRoleAdminChanged)
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
func (it *AssetHandlerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractRoleAdminChanged represents a RoleAdminChanged event raised by the AssetHandlerContract contract.
type AssetHandlerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AssetHandlerContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractRoleAdminChangedIterator{contract: _AssetHandlerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractRoleAdminChanged)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseRoleAdminChanged(log types.Log) (*AssetHandlerContractRoleAdminChanged, error) {
	event := new(AssetHandlerContractRoleAdminChanged)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AssetHandlerContract contract.
type AssetHandlerContractRoleGrantedIterator struct {
	Event *AssetHandlerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractRoleGranted)
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
		it.Event = new(AssetHandlerContractRoleGranted)
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
func (it *AssetHandlerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractRoleGranted represents a RoleGranted event raised by the AssetHandlerContract contract.
type AssetHandlerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetHandlerContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractRoleGrantedIterator{contract: _AssetHandlerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractRoleGranted)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseRoleGranted(log types.Log) (*AssetHandlerContractRoleGranted, error) {
	event := new(AssetHandlerContractRoleGranted)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AssetHandlerContract contract.
type AssetHandlerContractRoleRevokedIterator struct {
	Event *AssetHandlerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractRoleRevoked)
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
		it.Event = new(AssetHandlerContractRoleRevoked)
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
func (it *AssetHandlerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractRoleRevoked represents a RoleRevoked event raised by the AssetHandlerContract contract.
type AssetHandlerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetHandlerContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractRoleRevokedIterator{contract: _AssetHandlerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractRoleRevoked)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseRoleRevoked(log types.Log) (*AssetHandlerContractRoleRevoked, error) {
	event := new(AssetHandlerContractRoleRevoked)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetHandlerContractTokenSwitchIterator is returned from FilterTokenSwitch and is used to iterate over the raw logs and unpacked data for TokenSwitch events raised by the AssetHandlerContract contract.
type AssetHandlerContractTokenSwitchIterator struct {
	Event *AssetHandlerContractTokenSwitch // Event containing the contract specifics and raw log

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
func (it *AssetHandlerContractTokenSwitchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetHandlerContractTokenSwitch)
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
		it.Event = new(AssetHandlerContractTokenSwitch)
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
func (it *AssetHandlerContractTokenSwitchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetHandlerContractTokenSwitchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetHandlerContractTokenSwitch represents a TokenSwitch event raised by the AssetHandlerContract contract.
type AssetHandlerContractTokenSwitch struct {
	Ticker   [32]byte
	ChainId  [32]byte
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenSwitch is a free log retrieval operation binding the contract event 0x75e7e58cf4ee88ac2cce397e121b3463db42e4b190a214537949ce3466348396.
//
// Solidity: event TokenSwitch(bytes32 indexed ticker, bytes32 indexed chainId, bool isActive)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterTokenSwitch(opts *bind.FilterOpts, ticker [][32]byte, chainId [][32]byte) (*AssetHandlerContractTokenSwitchIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "TokenSwitch", tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractTokenSwitchIterator{contract: _AssetHandlerContract.contract, event: "TokenSwitch", logs: logs, sub: sub}, nil
}

// WatchTokenSwitch is a free log subscription operation binding the contract event 0x75e7e58cf4ee88ac2cce397e121b3463db42e4b190a214537949ce3466348396.
//
// Solidity: event TokenSwitch(bytes32 indexed ticker, bytes32 indexed chainId, bool isActive)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchTokenSwitch(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractTokenSwitch, ticker [][32]byte, chainId [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "TokenSwitch", tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetHandlerContractTokenSwitch)
				if err := _AssetHandlerContract.contract.UnpackLog(event, "TokenSwitch", log); err != nil {
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

// ParseTokenSwitch is a log parse operation binding the contract event 0x75e7e58cf4ee88ac2cce397e121b3463db42e4b190a214537949ce3466348396.
//
// Solidity: event TokenSwitch(bytes32 indexed ticker, bytes32 indexed chainId, bool isActive)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseTokenSwitch(log types.Log) (*AssetHandlerContractTokenSwitch, error) {
	event := new(AssetHandlerContractTokenSwitch)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "TokenSwitch", log); err != nil {
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
	Ticker  [32]byte
	ChainId [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3cbda7e3dafe8e1d194e7a9bbff9a8a52c3104f5fd0677b5c1ef2bc0a7892970.
//
// Solidity: event Withdraw(bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) FilterWithdraw(opts *bind.FilterOpts, ticker [][32]byte, chainId [][32]byte) (*AssetHandlerContractWithdrawIterator, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.FilterLogs(opts, "Withdraw", tickerRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetHandlerContractWithdrawIterator{contract: _AssetHandlerContract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3cbda7e3dafe8e1d194e7a9bbff9a8a52c3104f5fd0677b5c1ef2bc0a7892970.
//
// Solidity: event Withdraw(bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AssetHandlerContractWithdraw, ticker [][32]byte, chainId [][32]byte) (event.Subscription, error) {

	var tickerRule []interface{}
	for _, tickerItem := range ticker {
		tickerRule = append(tickerRule, tickerItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _AssetHandlerContract.contract.WatchLogs(opts, "Withdraw", tickerRule, chainIdRule)
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

// ParseWithdraw is a log parse operation binding the contract event 0x3cbda7e3dafe8e1d194e7a9bbff9a8a52c3104f5fd0677b5c1ef2bc0a7892970.
//
// Solidity: event Withdraw(bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount)
func (_AssetHandlerContract *AssetHandlerContractFilterer) ParseWithdraw(log types.Log) (*AssetHandlerContractWithdraw, error) {
	event := new(AssetHandlerContractWithdraw)
	if err := _AssetHandlerContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
