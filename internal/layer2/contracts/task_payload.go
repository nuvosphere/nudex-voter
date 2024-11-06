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

// TaskPayloadContractMetaData contains all meta data concerning the TaskPayloadContract contract.
var TaskPayloadContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"DepositRequest\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"taskType\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"targetAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"txHash\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumTaskPayload.AssetType\"},{\"name\":\"decimal\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositResult\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"success\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorCode\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorMsg\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResult\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"success\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorCode\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorMsg\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WalletCreationRequest\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"taskType\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"chain\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumTaskPayload.Chain\"},{\"name\":\"index\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WalletCreationResult\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"success\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorCode\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorMsg\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"walletAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRequest\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"taskType\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"targetAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"txHash\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ticker\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"assetType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumTaskPayload.AssetType\"},{\"name\":\"decimal\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalResult\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"success\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorCode\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"errorMsg\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// TaskPayloadContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskPayloadContractMetaData.ABI instead.
var TaskPayloadContractABI = TaskPayloadContractMetaData.ABI

// TaskPayloadContract is an auto generated Go binding around an Ethereum contract.
type TaskPayloadContract struct {
	TaskPayloadContractCaller     // Read-only binding to the contract
	TaskPayloadContractTransactor // Write-only binding to the contract
	TaskPayloadContractFilterer   // Log filterer for contract events
}

// TaskPayloadContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskPayloadContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskPayloadContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskPayloadContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskPayloadContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskPayloadContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskPayloadContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskPayloadContractSession struct {
	Contract     *TaskPayloadContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaskPayloadContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskPayloadContractCallerSession struct {
	Contract *TaskPayloadContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TaskPayloadContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskPayloadContractTransactorSession struct {
	Contract     *TaskPayloadContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TaskPayloadContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskPayloadContractRaw struct {
	Contract *TaskPayloadContract // Generic contract binding to access the raw methods on
}

// TaskPayloadContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskPayloadContractCallerRaw struct {
	Contract *TaskPayloadContractCaller // Generic read-only contract binding to access the raw methods on
}

// TaskPayloadContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskPayloadContractTransactorRaw struct {
	Contract *TaskPayloadContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskPayloadContract creates a new instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContract(address common.Address, backend bind.ContractBackend) (*TaskPayloadContract, error) {
	contract, err := bindTaskPayloadContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContract{TaskPayloadContractCaller: TaskPayloadContractCaller{contract: contract}, TaskPayloadContractTransactor: TaskPayloadContractTransactor{contract: contract}, TaskPayloadContractFilterer: TaskPayloadContractFilterer{contract: contract}}, nil
}

// NewTaskPayloadContractCaller creates a new read-only instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContractCaller(address common.Address, caller bind.ContractCaller) (*TaskPayloadContractCaller, error) {
	contract, err := bindTaskPayloadContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractCaller{contract: contract}, nil
}

// NewTaskPayloadContractTransactor creates a new write-only instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskPayloadContractTransactor, error) {
	contract, err := bindTaskPayloadContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractTransactor{contract: contract}, nil
}

// NewTaskPayloadContractFilterer creates a new log filterer instance of TaskPayloadContract, bound to a specific deployed contract.
func NewTaskPayloadContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskPayloadContractFilterer, error) {
	contract, err := bindTaskPayloadContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractFilterer{contract: contract}, nil
}

// bindTaskPayloadContract binds a generic wrapper to an already deployed contract.
func bindTaskPayloadContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskPayloadContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskPayloadContract *TaskPayloadContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskPayloadContract.Contract.TaskPayloadContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskPayloadContract *TaskPayloadContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.TaskPayloadContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskPayloadContract *TaskPayloadContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.TaskPayloadContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskPayloadContract *TaskPayloadContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskPayloadContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskPayloadContract *TaskPayloadContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskPayloadContract *TaskPayloadContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskPayloadContract.Contract.contract.Transact(opts, method, params...)
}

// TaskPayloadContractDepositRequestIterator is returned from FilterDepositRequest and is used to iterate over the raw logs and unpacked data for DepositRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositRequestIterator struct {
	Event *TaskPayloadContractDepositRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractDepositRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractDepositRequest)
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
		it.Event = new(TaskPayloadContractDepositRequest)
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
func (it *TaskPayloadContractDepositRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractDepositRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractDepositRequest represents a DepositRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositRequest struct {
	Version         uint32
	TaskType        uint32
	TargetAddress   common.Address
	Amount          uint64
	ChainId         uint32
	BlockHeight     uint64
	TxHash          []byte
	ContractAddress common.Address
	Ticker          string
	AssetType       uint8
	Decimal         uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositRequest is a free log retrieval operation binding the contract event 0xea93332254f87d488690c41de539c76f81103e770e84d209304f8ce20f417c54.
//
// Solidity: event DepositRequest(uint32 version, uint32 taskType, address targetAddress, uint64 amount, uint32 chainId, uint64 blockHeight, bytes txHash, address contractAddress, string ticker, uint8 assetType, uint32 decimal)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterDepositRequest(opts *bind.FilterOpts) (*TaskPayloadContractDepositRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "DepositRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractDepositRequestIterator{contract: _TaskPayloadContract.contract, event: "DepositRequest", logs: logs, sub: sub}, nil
}

// WatchDepositRequest is a free log subscription operation binding the contract event 0xea93332254f87d488690c41de539c76f81103e770e84d209304f8ce20f417c54.
//
// Solidity: event DepositRequest(uint32 version, uint32 taskType, address targetAddress, uint64 amount, uint32 chainId, uint64 blockHeight, bytes txHash, address contractAddress, string ticker, uint8 assetType, uint32 decimal)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchDepositRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractDepositRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "DepositRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractDepositRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositRequest", log); err != nil {
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

// ParseDepositRequest is a log parse operation binding the contract event 0xea93332254f87d488690c41de539c76f81103e770e84d209304f8ce20f417c54.
//
// Solidity: event DepositRequest(uint32 version, uint32 taskType, address targetAddress, uint64 amount, uint32 chainId, uint64 blockHeight, bytes txHash, address contractAddress, string ticker, uint8 assetType, uint32 decimal)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseDepositRequest(log types.Log) (*TaskPayloadContractDepositRequest, error) {
	event := new(TaskPayloadContractDepositRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractDepositResultIterator is returned from FilterDepositResult and is used to iterate over the raw logs and unpacked data for DepositResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositResultIterator struct {
	Event *TaskPayloadContractDepositResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractDepositResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractDepositResult)
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
		it.Event = new(TaskPayloadContractDepositResult)
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
func (it *TaskPayloadContractDepositResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractDepositResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractDepositResult represents a DepositResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractDepositResult struct {
	Version   uint32
	Success   uint32
	ErrorCode uint32
	ErrorMsg  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositResult is a free log retrieval operation binding the contract event 0x4b9d422b81766e21945af43bf87aded9d3c761f290a88bf9e694f99dc1f584d3.
//
// Solidity: event DepositResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterDepositResult(opts *bind.FilterOpts) (*TaskPayloadContractDepositResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "DepositResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractDepositResultIterator{contract: _TaskPayloadContract.contract, event: "DepositResult", logs: logs, sub: sub}, nil
}

// WatchDepositResult is a free log subscription operation binding the contract event 0x4b9d422b81766e21945af43bf87aded9d3c761f290a88bf9e694f99dc1f584d3.
//
// Solidity: event DepositResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchDepositResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractDepositResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "DepositResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractDepositResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositResult", log); err != nil {
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

// ParseDepositResult is a log parse operation binding the contract event 0x4b9d422b81766e21945af43bf87aded9d3c761f290a88bf9e694f99dc1f584d3.
//
// Solidity: event DepositResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseDepositResult(log types.Log) (*TaskPayloadContractDepositResult, error) {
	event := new(TaskPayloadContractDepositResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "DepositResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractTaskResultIterator is returned from FilterTaskResult and is used to iterate over the raw logs and unpacked data for TaskResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractTaskResultIterator struct {
	Event *TaskPayloadContractTaskResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractTaskResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractTaskResult)
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
		it.Event = new(TaskPayloadContractTaskResult)
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
func (it *TaskPayloadContractTaskResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractTaskResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractTaskResult represents a TaskResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractTaskResult struct {
	Version   uint32
	Success   uint32
	ErrorCode uint32
	ErrorMsg  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskResult is a free log retrieval operation binding the contract event 0xe13ccd793fa68300bc498250a8e759a348adcbbf2eff9d6361f8de8b791f6316.
//
// Solidity: event TaskResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterTaskResult(opts *bind.FilterOpts) (*TaskPayloadContractTaskResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "TaskResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractTaskResultIterator{contract: _TaskPayloadContract.contract, event: "TaskResult", logs: logs, sub: sub}, nil
}

// WatchTaskResult is a free log subscription operation binding the contract event 0xe13ccd793fa68300bc498250a8e759a348adcbbf2eff9d6361f8de8b791f6316.
//
// Solidity: event TaskResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchTaskResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractTaskResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "TaskResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractTaskResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "TaskResult", log); err != nil {
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

// ParseTaskResult is a log parse operation binding the contract event 0xe13ccd793fa68300bc498250a8e759a348adcbbf2eff9d6361f8de8b791f6316.
//
// Solidity: event TaskResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseTaskResult(log types.Log) (*TaskPayloadContractTaskResult, error) {
	event := new(TaskPayloadContractTaskResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "TaskResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWalletCreationRequestIterator is returned from FilterWalletCreationRequest and is used to iterate over the raw logs and unpacked data for WalletCreationRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationRequestIterator struct {
	Event *TaskPayloadContractWalletCreationRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWalletCreationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWalletCreationRequest)
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
		it.Event = new(TaskPayloadContractWalletCreationRequest)
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
func (it *TaskPayloadContractWalletCreationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWalletCreationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWalletCreationRequest represents a WalletCreationRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationRequest struct {
	Version  uint32
	TaskType uint32
	User     common.Address
	Account  uint64
	Chain    uint8
	Index    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletCreationRequest is a free log retrieval operation binding the contract event 0x0f1413e8d10cd1ec520cd20e110e7f744aadade9260edf12bea2ae80bf938c2e.
//
// Solidity: event WalletCreationRequest(uint32 version, uint32 taskType, address user, uint64 account, uint8 chain, uint32 index)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWalletCreationRequest(opts *bind.FilterOpts) (*TaskPayloadContractWalletCreationRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WalletCreationRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWalletCreationRequestIterator{contract: _TaskPayloadContract.contract, event: "WalletCreationRequest", logs: logs, sub: sub}, nil
}

// WatchWalletCreationRequest is a free log subscription operation binding the contract event 0x0f1413e8d10cd1ec520cd20e110e7f744aadade9260edf12bea2ae80bf938c2e.
//
// Solidity: event WalletCreationRequest(uint32 version, uint32 taskType, address user, uint64 account, uint8 chain, uint32 index)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWalletCreationRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWalletCreationRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WalletCreationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWalletCreationRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationRequest", log); err != nil {
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

// ParseWalletCreationRequest is a log parse operation binding the contract event 0x0f1413e8d10cd1ec520cd20e110e7f744aadade9260edf12bea2ae80bf938c2e.
//
// Solidity: event WalletCreationRequest(uint32 version, uint32 taskType, address user, uint64 account, uint8 chain, uint32 index)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWalletCreationRequest(log types.Log) (*TaskPayloadContractWalletCreationRequest, error) {
	event := new(TaskPayloadContractWalletCreationRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWalletCreationResultIterator is returned from FilterWalletCreationResult and is used to iterate over the raw logs and unpacked data for WalletCreationResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationResultIterator struct {
	Event *TaskPayloadContractWalletCreationResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWalletCreationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWalletCreationResult)
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
		it.Event = new(TaskPayloadContractWalletCreationResult)
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
func (it *TaskPayloadContractWalletCreationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWalletCreationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWalletCreationResult represents a WalletCreationResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractWalletCreationResult struct {
	Version       uint32
	Success       uint32
	ErrorCode     uint32
	ErrorMsg      string
	WalletAddress string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWalletCreationResult is a free log retrieval operation binding the contract event 0x99f6ad63176d2f97471cb4d5385e8fa14767b05685160ef30ba50f9415008e8e.
//
// Solidity: event WalletCreationResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg, string walletAddress)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWalletCreationResult(opts *bind.FilterOpts) (*TaskPayloadContractWalletCreationResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WalletCreationResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWalletCreationResultIterator{contract: _TaskPayloadContract.contract, event: "WalletCreationResult", logs: logs, sub: sub}, nil
}

// WatchWalletCreationResult is a free log subscription operation binding the contract event 0x99f6ad63176d2f97471cb4d5385e8fa14767b05685160ef30ba50f9415008e8e.
//
// Solidity: event WalletCreationResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg, string walletAddress)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWalletCreationResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWalletCreationResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WalletCreationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWalletCreationResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationResult", log); err != nil {
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

// ParseWalletCreationResult is a log parse operation binding the contract event 0x99f6ad63176d2f97471cb4d5385e8fa14767b05685160ef30ba50f9415008e8e.
//
// Solidity: event WalletCreationResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg, string walletAddress)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWalletCreationResult(log types.Log) (*TaskPayloadContractWalletCreationResult, error) {
	event := new(TaskPayloadContractWalletCreationResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WalletCreationResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWithdrawalRequestIterator is returned from FilterWithdrawalRequest and is used to iterate over the raw logs and unpacked data for WithdrawalRequest events raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalRequestIterator struct {
	Event *TaskPayloadContractWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWithdrawalRequest)
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
		it.Event = new(TaskPayloadContractWithdrawalRequest)
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
func (it *TaskPayloadContractWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWithdrawalRequest represents a WithdrawalRequest event raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalRequest struct {
	Version         uint32
	TaskType        uint32
	TargetAddress   common.Address
	Amount          uint64
	ChainId         uint32
	BlockHeight     uint64
	TxHash          []byte
	ContractAddress common.Address
	Ticker          string
	AssetType       uint8
	Decimal         uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequest is a free log retrieval operation binding the contract event 0x6d27c4501566f5d35d6ba3fd06c4f9fd407374318f22de6f9895d24d56a0e6dc.
//
// Solidity: event WithdrawalRequest(uint32 version, uint32 taskType, address targetAddress, uint64 amount, uint32 chainId, uint64 blockHeight, bytes txHash, address contractAddress, string ticker, uint8 assetType, uint32 decimal)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWithdrawalRequest(opts *bind.FilterOpts) (*TaskPayloadContractWithdrawalRequestIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWithdrawalRequestIterator{contract: _TaskPayloadContract.contract, event: "WithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequest is a free log subscription operation binding the contract event 0x6d27c4501566f5d35d6ba3fd06c4f9fd407374318f22de6f9895d24d56a0e6dc.
//
// Solidity: event WithdrawalRequest(uint32 version, uint32 taskType, address targetAddress, uint64 amount, uint32 chainId, uint64 blockHeight, bytes txHash, address contractAddress, string ticker, uint8 assetType, uint32 decimal)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWithdrawalRequest)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalRequest", log); err != nil {
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

// ParseWithdrawalRequest is a log parse operation binding the contract event 0x6d27c4501566f5d35d6ba3fd06c4f9fd407374318f22de6f9895d24d56a0e6dc.
//
// Solidity: event WithdrawalRequest(uint32 version, uint32 taskType, address targetAddress, uint64 amount, uint32 chainId, uint64 blockHeight, bytes txHash, address contractAddress, string ticker, uint8 assetType, uint32 decimal)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWithdrawalRequest(log types.Log) (*TaskPayloadContractWithdrawalRequest, error) {
	event := new(TaskPayloadContractWithdrawalRequest)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskPayloadContractWithdrawalResultIterator is returned from FilterWithdrawalResult and is used to iterate over the raw logs and unpacked data for WithdrawalResult events raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalResultIterator struct {
	Event *TaskPayloadContractWithdrawalResult // Event containing the contract specifics and raw log

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
func (it *TaskPayloadContractWithdrawalResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskPayloadContractWithdrawalResult)
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
		it.Event = new(TaskPayloadContractWithdrawalResult)
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
func (it *TaskPayloadContractWithdrawalResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskPayloadContractWithdrawalResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskPayloadContractWithdrawalResult represents a WithdrawalResult event raised by the TaskPayloadContract contract.
type TaskPayloadContractWithdrawalResult struct {
	Version   uint32
	Success   uint32
	ErrorCode uint32
	ErrorMsg  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalResult is a free log retrieval operation binding the contract event 0xf011021fc9bcde3609110858e94203136e99da4f02bd0804d05021317d58f728.
//
// Solidity: event WithdrawalResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) FilterWithdrawalResult(opts *bind.FilterOpts) (*TaskPayloadContractWithdrawalResultIterator, error) {

	logs, sub, err := _TaskPayloadContract.contract.FilterLogs(opts, "WithdrawalResult")
	if err != nil {
		return nil, err
	}
	return &TaskPayloadContractWithdrawalResultIterator{contract: _TaskPayloadContract.contract, event: "WithdrawalResult", logs: logs, sub: sub}, nil
}

// WatchWithdrawalResult is a free log subscription operation binding the contract event 0xf011021fc9bcde3609110858e94203136e99da4f02bd0804d05021317d58f728.
//
// Solidity: event WithdrawalResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) WatchWithdrawalResult(opts *bind.WatchOpts, sink chan<- *TaskPayloadContractWithdrawalResult) (event.Subscription, error) {

	logs, sub, err := _TaskPayloadContract.contract.WatchLogs(opts, "WithdrawalResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskPayloadContractWithdrawalResult)
				if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalResult", log); err != nil {
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

// ParseWithdrawalResult is a log parse operation binding the contract event 0xf011021fc9bcde3609110858e94203136e99da4f02bd0804d05021317d58f728.
//
// Solidity: event WithdrawalResult(uint32 version, uint32 success, uint32 errorCode, string errorMsg)
func (_TaskPayloadContract *TaskPayloadContractFilterer) ParseWithdrawalResult(log types.Log) (*TaskPayloadContractWithdrawalResult, error) {
	event := new(TaskPayloadContractWithdrawalResult)
	if err := _TaskPayloadContract.contract.UnpackLog(event, "WithdrawalResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
