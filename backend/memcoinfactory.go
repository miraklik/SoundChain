// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package memcoin
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

// MemcoinfactoryMetaData contains all meta data concerning the Memcoinfactory contract.
var MemcoinfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_routerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"artist\",\"type\":\"address\"}],\"name\":\"MemCoinCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"createMemCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MemcoinfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MemcoinfactoryMetaData.ABI instead.
var MemcoinfactoryABI = MemcoinfactoryMetaData.ABI

// Memcoinfactory is an auto generated Go binding around an Ethereum contract.
type Memcoinfactory struct {
	MemcoinfactoryCaller     // Read-only binding to the contract
	MemcoinfactoryTransactor // Write-only binding to the contract
	MemcoinfactoryFilterer   // Log filterer for contract events
}

// MemcoinfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemcoinfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemcoinfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemcoinfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemcoinfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemcoinfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemcoinfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemcoinfactorySession struct {
	Contract     *Memcoinfactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemcoinfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemcoinfactoryCallerSession struct {
	Contract *MemcoinfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MemcoinfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemcoinfactoryTransactorSession struct {
	Contract     *MemcoinfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MemcoinfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemcoinfactoryRaw struct {
	Contract *Memcoinfactory // Generic contract binding to access the raw methods on
}

// MemcoinfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemcoinfactoryCallerRaw struct {
	Contract *MemcoinfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemcoinfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemcoinfactoryTransactorRaw struct {
	Contract *MemcoinfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemcoinfactory creates a new instance of Memcoinfactory, bound to a specific deployed contract.
func NewMemcoinfactory(address common.Address, backend bind.ContractBackend) (*Memcoinfactory, error) {
	contract, err := bindMemcoinfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memcoinfactory{MemcoinfactoryCaller: MemcoinfactoryCaller{contract: contract}, MemcoinfactoryTransactor: MemcoinfactoryTransactor{contract: contract}, MemcoinfactoryFilterer: MemcoinfactoryFilterer{contract: contract}}, nil
}

// NewMemcoinfactoryCaller creates a new read-only instance of Memcoinfactory, bound to a specific deployed contract.
func NewMemcoinfactoryCaller(address common.Address, caller bind.ContractCaller) (*MemcoinfactoryCaller, error) {
	contract, err := bindMemcoinfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemcoinfactoryCaller{contract: contract}, nil
}

// NewMemcoinfactoryTransactor creates a new write-only instance of Memcoinfactory, bound to a specific deployed contract.
func NewMemcoinfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemcoinfactoryTransactor, error) {
	contract, err := bindMemcoinfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemcoinfactoryTransactor{contract: contract}, nil
}

// NewMemcoinfactoryFilterer creates a new log filterer instance of Memcoinfactory, bound to a specific deployed contract.
func NewMemcoinfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemcoinfactoryFilterer, error) {
	contract, err := bindMemcoinfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemcoinfactoryFilterer{contract: contract}, nil
}

// bindMemcoinfactory binds a generic wrapper to an already deployed contract.
func bindMemcoinfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemcoinfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memcoinfactory *MemcoinfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memcoinfactory.Contract.MemcoinfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memcoinfactory *MemcoinfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.MemcoinfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memcoinfactory *MemcoinfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.MemcoinfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memcoinfactory *MemcoinfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memcoinfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memcoinfactory *MemcoinfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memcoinfactory *MemcoinfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.contract.Transact(opts, method, params...)
}

// AllTokens is a free data retrieval call binding the contract method 0x634282af.
//
// Solidity: function allTokens(uint256 ) view returns(address)
func (_Memcoinfactory *MemcoinfactoryCaller) AllTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Memcoinfactory.contract.Call(opts, &out, "allTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x634282af.
//
// Solidity: function allTokens(uint256 ) view returns(address)
func (_Memcoinfactory *MemcoinfactorySession) AllTokens(arg0 *big.Int) (common.Address, error) {
	return _Memcoinfactory.Contract.AllTokens(&_Memcoinfactory.CallOpts, arg0)
}

// AllTokens is a free data retrieval call binding the contract method 0x634282af.
//
// Solidity: function allTokens(uint256 ) view returns(address)
func (_Memcoinfactory *MemcoinfactoryCallerSession) AllTokens(arg0 *big.Int) (common.Address, error) {
	return _Memcoinfactory.Contract.AllTokens(&_Memcoinfactory.CallOpts, arg0)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns(address[])
func (_Memcoinfactory *MemcoinfactoryCaller) GetAllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Memcoinfactory.contract.Call(opts, &out, "getAllTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns(address[])
func (_Memcoinfactory *MemcoinfactorySession) GetAllTokens() ([]common.Address, error) {
	return _Memcoinfactory.Contract.GetAllTokens(&_Memcoinfactory.CallOpts)
}

// GetAllTokens is a free data retrieval call binding the contract method 0x2a5c792a.
//
// Solidity: function getAllTokens() view returns(address[])
func (_Memcoinfactory *MemcoinfactoryCallerSession) GetAllTokens() ([]common.Address, error) {
	return _Memcoinfactory.Contract.GetAllTokens(&_Memcoinfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Memcoinfactory *MemcoinfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Memcoinfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Memcoinfactory *MemcoinfactorySession) Owner() (common.Address, error) {
	return _Memcoinfactory.Contract.Owner(&_Memcoinfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Memcoinfactory *MemcoinfactoryCallerSession) Owner() (common.Address, error) {
	return _Memcoinfactory.Contract.Owner(&_Memcoinfactory.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_Memcoinfactory *MemcoinfactoryCaller) RouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Memcoinfactory.contract.Call(opts, &out, "routerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_Memcoinfactory *MemcoinfactorySession) RouterAddress() (common.Address, error) {
	return _Memcoinfactory.Contract.RouterAddress(&_Memcoinfactory.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_Memcoinfactory *MemcoinfactoryCallerSession) RouterAddress() (common.Address, error) {
	return _Memcoinfactory.Contract.RouterAddress(&_Memcoinfactory.CallOpts)
}

// CreateMemCoin is a paid mutator transaction binding the contract method 0x02beb2d7.
//
// Solidity: function createMemCoin(string name, string symbol, uint256 initialSupply) returns(address)
func (_Memcoinfactory *MemcoinfactoryTransactor) CreateMemCoin(opts *bind.TransactOpts, name string, symbol string, initialSupply *big.Int) (*types.Transaction, error) {
	return _Memcoinfactory.contract.Transact(opts, "createMemCoin", name, symbol, initialSupply)
}

// CreateMemCoin is a paid mutator transaction binding the contract method 0x02beb2d7.
//
// Solidity: function createMemCoin(string name, string symbol, uint256 initialSupply) returns(address)
func (_Memcoinfactory *MemcoinfactorySession) CreateMemCoin(name string, symbol string, initialSupply *big.Int) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.CreateMemCoin(&_Memcoinfactory.TransactOpts, name, symbol, initialSupply)
}

// CreateMemCoin is a paid mutator transaction binding the contract method 0x02beb2d7.
//
// Solidity: function createMemCoin(string name, string symbol, uint256 initialSupply) returns(address)
func (_Memcoinfactory *MemcoinfactoryTransactorSession) CreateMemCoin(name string, symbol string, initialSupply *big.Int) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.CreateMemCoin(&_Memcoinfactory.TransactOpts, name, symbol, initialSupply)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Memcoinfactory *MemcoinfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memcoinfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Memcoinfactory *MemcoinfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Memcoinfactory.Contract.RenounceOwnership(&_Memcoinfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Memcoinfactory *MemcoinfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Memcoinfactory.Contract.RenounceOwnership(&_Memcoinfactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Memcoinfactory *MemcoinfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Memcoinfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Memcoinfactory *MemcoinfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.TransferOwnership(&_Memcoinfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Memcoinfactory *MemcoinfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Memcoinfactory.Contract.TransferOwnership(&_Memcoinfactory.TransactOpts, newOwner)
}

// MemcoinfactoryMemCoinCreatedIterator is returned from FilterMemCoinCreated and is used to iterate over the raw logs and unpacked data for MemCoinCreated events raised by the Memcoinfactory contract.
type MemcoinfactoryMemCoinCreatedIterator struct {
	Event *MemcoinfactoryMemCoinCreated // Event containing the contract specifics and raw log

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
func (it *MemcoinfactoryMemCoinCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinfactoryMemCoinCreated)
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
		it.Event = new(MemcoinfactoryMemCoinCreated)
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
func (it *MemcoinfactoryMemCoinCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinfactoryMemCoinCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinfactoryMemCoinCreated represents a MemCoinCreated event raised by the Memcoinfactory contract.
type MemcoinfactoryMemCoinCreated struct {
	TokenAddress common.Address
	Artist       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMemCoinCreated is a free log retrieval operation binding the contract event 0xd6607fa9d7ab84d8e91463d6ff59add1a8532e4bdf6ac8d7e44d598d9d2cac13.
//
// Solidity: event MemCoinCreated(address indexed tokenAddress, address indexed artist)
func (_Memcoinfactory *MemcoinfactoryFilterer) FilterMemCoinCreated(opts *bind.FilterOpts, tokenAddress []common.Address, artist []common.Address) (*MemcoinfactoryMemCoinCreatedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var artistRule []interface{}
	for _, artistItem := range artist {
		artistRule = append(artistRule, artistItem)
	}

	logs, sub, err := _Memcoinfactory.contract.FilterLogs(opts, "MemCoinCreated", tokenAddressRule, artistRule)
	if err != nil {
		return nil, err
	}
	return &MemcoinfactoryMemCoinCreatedIterator{contract: _Memcoinfactory.contract, event: "MemCoinCreated", logs: logs, sub: sub}, nil
}

// WatchMemCoinCreated is a free log subscription operation binding the contract event 0xd6607fa9d7ab84d8e91463d6ff59add1a8532e4bdf6ac8d7e44d598d9d2cac13.
//
// Solidity: event MemCoinCreated(address indexed tokenAddress, address indexed artist)
func (_Memcoinfactory *MemcoinfactoryFilterer) WatchMemCoinCreated(opts *bind.WatchOpts, sink chan<- *MemcoinfactoryMemCoinCreated, tokenAddress []common.Address, artist []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var artistRule []interface{}
	for _, artistItem := range artist {
		artistRule = append(artistRule, artistItem)
	}

	logs, sub, err := _Memcoinfactory.contract.WatchLogs(opts, "MemCoinCreated", tokenAddressRule, artistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinfactoryMemCoinCreated)
				if err := _Memcoinfactory.contract.UnpackLog(event, "MemCoinCreated", log); err != nil {
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

// ParseMemCoinCreated is a log parse operation binding the contract event 0xd6607fa9d7ab84d8e91463d6ff59add1a8532e4bdf6ac8d7e44d598d9d2cac13.
//
// Solidity: event MemCoinCreated(address indexed tokenAddress, address indexed artist)
func (_Memcoinfactory *MemcoinfactoryFilterer) ParseMemCoinCreated(log types.Log) (*MemcoinfactoryMemCoinCreated, error) {
	event := new(MemcoinfactoryMemCoinCreated)
	if err := _Memcoinfactory.contract.UnpackLog(event, "MemCoinCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemcoinfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Memcoinfactory contract.
type MemcoinfactoryOwnershipTransferredIterator struct {
	Event *MemcoinfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MemcoinfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinfactoryOwnershipTransferred)
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
		it.Event = new(MemcoinfactoryOwnershipTransferred)
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
func (it *MemcoinfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Memcoinfactory contract.
type MemcoinfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Memcoinfactory *MemcoinfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MemcoinfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Memcoinfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MemcoinfactoryOwnershipTransferredIterator{contract: _Memcoinfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Memcoinfactory *MemcoinfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MemcoinfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Memcoinfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinfactoryOwnershipTransferred)
				if err := _Memcoinfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Memcoinfactory *MemcoinfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*MemcoinfactoryOwnershipTransferred, error) {
	event := new(MemcoinfactoryOwnershipTransferred)
	if err := _Memcoinfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
