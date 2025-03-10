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

// MemcoinMetaData contains all meta data concerning the Memcoin contract.
var MemcoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"routerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ArithmeticError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TaxTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAmountMustBePositive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eth\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaxCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEAD_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeFromTax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isExcludedFromTax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buyTax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sellTax\",\"type\":\"uint256\"}],\"name\":\"setTaxes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswaprouter\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MemcoinABI is the input ABI used to generate the binding from.
// Deprecated: Use MemcoinMetaData.ABI instead.
var MemcoinABI = MemcoinMetaData.ABI

// Memcoin is an auto generated Go binding around an Ethereum contract.
type Memcoin struct {
	MemcoinCaller     // Read-only binding to the contract
	MemcoinTransactor // Write-only binding to the contract
	MemcoinFilterer   // Log filterer for contract events
}

// MemcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemcoinSession struct {
	Contract     *Memcoin          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemcoinCallerSession struct {
	Contract *MemcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MemcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemcoinTransactorSession struct {
	Contract     *MemcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MemcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemcoinRaw struct {
	Contract *Memcoin // Generic contract binding to access the raw methods on
}

// MemcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemcoinCallerRaw struct {
	Contract *MemcoinCaller // Generic read-only contract binding to access the raw methods on
}

// MemcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemcoinTransactorRaw struct {
	Contract *MemcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemcoin creates a new instance of Memcoin, bound to a specific deployed contract.
func NewMemcoin(address common.Address, backend bind.ContractBackend) (*Memcoin, error) {
	contract, err := bindMemcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memcoin{MemcoinCaller: MemcoinCaller{contract: contract}, MemcoinTransactor: MemcoinTransactor{contract: contract}, MemcoinFilterer: MemcoinFilterer{contract: contract}}, nil
}

// NewMemcoinCaller creates a new read-only instance of Memcoin, bound to a specific deployed contract.
func NewMemcoinCaller(address common.Address, caller bind.ContractCaller) (*MemcoinCaller, error) {
	contract, err := bindMemcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemcoinCaller{contract: contract}, nil
}

// NewMemcoinTransactor creates a new write-only instance of Memcoin, bound to a specific deployed contract.
func NewMemcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*MemcoinTransactor, error) {
	contract, err := bindMemcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemcoinTransactor{contract: contract}, nil
}

// NewMemcoinFilterer creates a new log filterer instance of Memcoin, bound to a specific deployed contract.
func NewMemcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*MemcoinFilterer, error) {
	contract, err := bindMemcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemcoinFilterer{contract: contract}, nil
}

// bindMemcoin binds a generic wrapper to an already deployed contract.
func bindMemcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemcoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memcoin *MemcoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memcoin.Contract.MemcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memcoin *MemcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memcoin.Contract.MemcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memcoin *MemcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memcoin.Contract.MemcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memcoin *MemcoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memcoin *MemcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memcoin *MemcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memcoin.Contract.contract.Transact(opts, method, params...)
}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Memcoin *MemcoinCaller) DEADADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "DEAD_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Memcoin *MemcoinSession) DEADADDRESS() (common.Address, error) {
	return _Memcoin.Contract.DEADADDRESS(&_Memcoin.CallOpts)
}

// DEADADDRESS is a free data retrieval call binding the contract method 0x4e6fd6c4.
//
// Solidity: function DEAD_ADDRESS() view returns(address)
func (_Memcoin *MemcoinCallerSession) DEADADDRESS() (common.Address, error) {
	return _Memcoin.Contract.DEADADDRESS(&_Memcoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Memcoin *MemcoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Memcoin *MemcoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Memcoin.Contract.Allowance(&_Memcoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Memcoin *MemcoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Memcoin.Contract.Allowance(&_Memcoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Memcoin *MemcoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Memcoin *MemcoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Memcoin.Contract.BalanceOf(&_Memcoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Memcoin *MemcoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Memcoin.Contract.BalanceOf(&_Memcoin.CallOpts, account)
}

// BurnShare is a free data retrieval call binding the contract method 0x05bf0a54.
//
// Solidity: function burnShare() view returns(uint256)
func (_Memcoin *MemcoinCaller) BurnShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "burnShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnShare is a free data retrieval call binding the contract method 0x05bf0a54.
//
// Solidity: function burnShare() view returns(uint256)
func (_Memcoin *MemcoinSession) BurnShare() (*big.Int, error) {
	return _Memcoin.Contract.BurnShare(&_Memcoin.CallOpts)
}

// BurnShare is a free data retrieval call binding the contract method 0x05bf0a54.
//
// Solidity: function burnShare() view returns(uint256)
func (_Memcoin *MemcoinCallerSession) BurnShare() (*big.Int, error) {
	return _Memcoin.Contract.BurnShare(&_Memcoin.CallOpts)
}

// BuyTax is a free data retrieval call binding the contract method 0x4f7041a5.
//
// Solidity: function buyTax() view returns(uint256)
func (_Memcoin *MemcoinCaller) BuyTax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "buyTax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyTax is a free data retrieval call binding the contract method 0x4f7041a5.
//
// Solidity: function buyTax() view returns(uint256)
func (_Memcoin *MemcoinSession) BuyTax() (*big.Int, error) {
	return _Memcoin.Contract.BuyTax(&_Memcoin.CallOpts)
}

// BuyTax is a free data retrieval call binding the contract method 0x4f7041a5.
//
// Solidity: function buyTax() view returns(uint256)
func (_Memcoin *MemcoinCallerSession) BuyTax() (*big.Int, error) {
	return _Memcoin.Contract.BuyTax(&_Memcoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Memcoin *MemcoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Memcoin *MemcoinSession) Decimals() (uint8, error) {
	return _Memcoin.Contract.Decimals(&_Memcoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Memcoin *MemcoinCallerSession) Decimals() (uint8, error) {
	return _Memcoin.Contract.Decimals(&_Memcoin.CallOpts)
}

// IsExcludedFromTax is a free data retrieval call binding the contract method 0xcb4ca631.
//
// Solidity: function isExcludedFromTax(address ) view returns(bool)
func (_Memcoin *MemcoinCaller) IsExcludedFromTax(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "isExcludedFromTax", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromTax is a free data retrieval call binding the contract method 0xcb4ca631.
//
// Solidity: function isExcludedFromTax(address ) view returns(bool)
func (_Memcoin *MemcoinSession) IsExcludedFromTax(arg0 common.Address) (bool, error) {
	return _Memcoin.Contract.IsExcludedFromTax(&_Memcoin.CallOpts, arg0)
}

// IsExcludedFromTax is a free data retrieval call binding the contract method 0xcb4ca631.
//
// Solidity: function isExcludedFromTax(address ) view returns(bool)
func (_Memcoin *MemcoinCallerSession) IsExcludedFromTax(arg0 common.Address) (bool, error) {
	return _Memcoin.Contract.IsExcludedFromTax(&_Memcoin.CallOpts, arg0)
}

// LiquidityShare is a free data retrieval call binding the contract method 0x15291cd4.
//
// Solidity: function liquidityShare() view returns(uint256)
func (_Memcoin *MemcoinCaller) LiquidityShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "liquidityShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityShare is a free data retrieval call binding the contract method 0x15291cd4.
//
// Solidity: function liquidityShare() view returns(uint256)
func (_Memcoin *MemcoinSession) LiquidityShare() (*big.Int, error) {
	return _Memcoin.Contract.LiquidityShare(&_Memcoin.CallOpts)
}

// LiquidityShare is a free data retrieval call binding the contract method 0x15291cd4.
//
// Solidity: function liquidityShare() view returns(uint256)
func (_Memcoin *MemcoinCallerSession) LiquidityShare() (*big.Int, error) {
	return _Memcoin.Contract.LiquidityShare(&_Memcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Memcoin *MemcoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Memcoin *MemcoinSession) Name() (string, error) {
	return _Memcoin.Contract.Name(&_Memcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Memcoin *MemcoinCallerSession) Name() (string, error) {
	return _Memcoin.Contract.Name(&_Memcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Memcoin *MemcoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Memcoin *MemcoinSession) Owner() (common.Address, error) {
	return _Memcoin.Contract.Owner(&_Memcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Memcoin *MemcoinCallerSession) Owner() (common.Address, error) {
	return _Memcoin.Contract.Owner(&_Memcoin.CallOpts)
}

// RewardsShare is a free data retrieval call binding the contract method 0x33727077.
//
// Solidity: function rewardsShare() view returns(uint256)
func (_Memcoin *MemcoinCaller) RewardsShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "rewardsShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsShare is a free data retrieval call binding the contract method 0x33727077.
//
// Solidity: function rewardsShare() view returns(uint256)
func (_Memcoin *MemcoinSession) RewardsShare() (*big.Int, error) {
	return _Memcoin.Contract.RewardsShare(&_Memcoin.CallOpts)
}

// RewardsShare is a free data retrieval call binding the contract method 0x33727077.
//
// Solidity: function rewardsShare() view returns(uint256)
func (_Memcoin *MemcoinCallerSession) RewardsShare() (*big.Int, error) {
	return _Memcoin.Contract.RewardsShare(&_Memcoin.CallOpts)
}

// SellTax is a free data retrieval call binding the contract method 0xcc1776d3.
//
// Solidity: function sellTax() view returns(uint256)
func (_Memcoin *MemcoinCaller) SellTax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "sellTax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellTax is a free data retrieval call binding the contract method 0xcc1776d3.
//
// Solidity: function sellTax() view returns(uint256)
func (_Memcoin *MemcoinSession) SellTax() (*big.Int, error) {
	return _Memcoin.Contract.SellTax(&_Memcoin.CallOpts)
}

// SellTax is a free data retrieval call binding the contract method 0xcc1776d3.
//
// Solidity: function sellTax() view returns(uint256)
func (_Memcoin *MemcoinCallerSession) SellTax() (*big.Int, error) {
	return _Memcoin.Contract.SellTax(&_Memcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Memcoin *MemcoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Memcoin *MemcoinSession) Symbol() (string, error) {
	return _Memcoin.Contract.Symbol(&_Memcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Memcoin *MemcoinCallerSession) Symbol() (string, error) {
	return _Memcoin.Contract.Symbol(&_Memcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Memcoin *MemcoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Memcoin *MemcoinSession) TotalSupply() (*big.Int, error) {
	return _Memcoin.Contract.TotalSupply(&_Memcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Memcoin *MemcoinCallerSession) TotalSupply() (*big.Int, error) {
	return _Memcoin.Contract.TotalSupply(&_Memcoin.CallOpts)
}

// UniswapPair is a free data retrieval call binding the contract method 0xc816841b.
//
// Solidity: function uniswapPair() view returns(address)
func (_Memcoin *MemcoinCaller) UniswapPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "uniswapPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapPair is a free data retrieval call binding the contract method 0xc816841b.
//
// Solidity: function uniswapPair() view returns(address)
func (_Memcoin *MemcoinSession) UniswapPair() (common.Address, error) {
	return _Memcoin.Contract.UniswapPair(&_Memcoin.CallOpts)
}

// UniswapPair is a free data retrieval call binding the contract method 0xc816841b.
//
// Solidity: function uniswapPair() view returns(address)
func (_Memcoin *MemcoinCallerSession) UniswapPair() (common.Address, error) {
	return _Memcoin.Contract.UniswapPair(&_Memcoin.CallOpts)
}

// Uniswaprouter is a free data retrieval call binding the contract method 0x2431ba73.
//
// Solidity: function uniswaprouter() view returns(address)
func (_Memcoin *MemcoinCaller) Uniswaprouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Memcoin.contract.Call(opts, &out, "uniswaprouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswaprouter is a free data retrieval call binding the contract method 0x2431ba73.
//
// Solidity: function uniswaprouter() view returns(address)
func (_Memcoin *MemcoinSession) Uniswaprouter() (common.Address, error) {
	return _Memcoin.Contract.Uniswaprouter(&_Memcoin.CallOpts)
}

// Uniswaprouter is a free data retrieval call binding the contract method 0x2431ba73.
//
// Solidity: function uniswaprouter() view returns(address)
func (_Memcoin *MemcoinCallerSession) Uniswaprouter() (common.Address, error) {
	return _Memcoin.Contract.Uniswaprouter(&_Memcoin.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 tokenAmount) returns()
func (_Memcoin *MemcoinTransactor) AddLiquidity(opts *bind.TransactOpts, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "addLiquidity", tokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 tokenAmount) returns()
func (_Memcoin *MemcoinSession) AddLiquidity(tokenAmount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.AddLiquidity(&_Memcoin.TransactOpts, tokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 tokenAmount) returns()
func (_Memcoin *MemcoinTransactorSession) AddLiquidity(tokenAmount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.AddLiquidity(&_Memcoin.TransactOpts, tokenAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Memcoin *MemcoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Memcoin *MemcoinSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Approve(&_Memcoin.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Memcoin *MemcoinTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Approve(&_Memcoin.TransactOpts, spender, value)
}

// ExcludeFromTax is a paid mutator transaction binding the contract method 0xc6a30647.
//
// Solidity: function excludeFromTax(address account, bool excluded) returns()
func (_Memcoin *MemcoinTransactor) ExcludeFromTax(opts *bind.TransactOpts, account common.Address, excluded bool) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "excludeFromTax", account, excluded)
}

// ExcludeFromTax is a paid mutator transaction binding the contract method 0xc6a30647.
//
// Solidity: function excludeFromTax(address account, bool excluded) returns()
func (_Memcoin *MemcoinSession) ExcludeFromTax(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Memcoin.Contract.ExcludeFromTax(&_Memcoin.TransactOpts, account, excluded)
}

// ExcludeFromTax is a paid mutator transaction binding the contract method 0xc6a30647.
//
// Solidity: function excludeFromTax(address account, bool excluded) returns()
func (_Memcoin *MemcoinTransactorSession) ExcludeFromTax(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Memcoin.Contract.ExcludeFromTax(&_Memcoin.TransactOpts, account, excluded)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Memcoin *MemcoinTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Memcoin *MemcoinSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Mint(&_Memcoin.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Memcoin *MemcoinTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Mint(&_Memcoin.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Memcoin *MemcoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Memcoin *MemcoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _Memcoin.Contract.RenounceOwnership(&_Memcoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Memcoin *MemcoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Memcoin.Contract.RenounceOwnership(&_Memcoin.TransactOpts)
}

// SetTaxes is a paid mutator transaction binding the contract method 0xc647b20e.
//
// Solidity: function setTaxes(uint256 _buyTax, uint256 _sellTax) returns()
func (_Memcoin *MemcoinTransactor) SetTaxes(opts *bind.TransactOpts, _buyTax *big.Int, _sellTax *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "setTaxes", _buyTax, _sellTax)
}

// SetTaxes is a paid mutator transaction binding the contract method 0xc647b20e.
//
// Solidity: function setTaxes(uint256 _buyTax, uint256 _sellTax) returns()
func (_Memcoin *MemcoinSession) SetTaxes(_buyTax *big.Int, _sellTax *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.SetTaxes(&_Memcoin.TransactOpts, _buyTax, _sellTax)
}

// SetTaxes is a paid mutator transaction binding the contract method 0xc647b20e.
//
// Solidity: function setTaxes(uint256 _buyTax, uint256 _sellTax) returns()
func (_Memcoin *MemcoinTransactorSession) SetTaxes(_buyTax *big.Int, _sellTax *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.SetTaxes(&_Memcoin.TransactOpts, _buyTax, _sellTax)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Memcoin *MemcoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Memcoin *MemcoinSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Transfer(&_Memcoin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Memcoin *MemcoinTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Transfer(&_Memcoin.TransactOpts, to, value)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address sender, address recipient, uint256 amount) returns()
func (_Memcoin *MemcoinTransactor) Transfer0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "transfer0", sender, recipient, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address sender, address recipient, uint256 amount) returns()
func (_Memcoin *MemcoinSession) Transfer0(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Transfer0(&_Memcoin.TransactOpts, sender, recipient, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address sender, address recipient, uint256 amount) returns()
func (_Memcoin *MemcoinTransactorSession) Transfer0(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Transfer0(&_Memcoin.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Memcoin *MemcoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Memcoin *MemcoinSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.TransferFrom(&_Memcoin.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Memcoin *MemcoinTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.TransferFrom(&_Memcoin.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Memcoin *MemcoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Memcoin *MemcoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Memcoin.Contract.TransferOwnership(&_Memcoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Memcoin *MemcoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Memcoin.Contract.TransferOwnership(&_Memcoin.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Memcoin *MemcoinTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Memcoin *MemcoinSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Withdraw(&_Memcoin.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Memcoin *MemcoinTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Memcoin.Contract.Withdraw(&_Memcoin.TransactOpts, amount)
}

// MemcoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Memcoin contract.
type MemcoinApprovalIterator struct {
	Event *MemcoinApproval // Event containing the contract specifics and raw log

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
func (it *MemcoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinApproval)
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
		it.Event = new(MemcoinApproval)
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
func (it *MemcoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinApproval represents a Approval event raised by the Memcoin contract.
type MemcoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Memcoin *MemcoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MemcoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Memcoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MemcoinApprovalIterator{contract: _Memcoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Memcoin *MemcoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MemcoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Memcoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinApproval)
				if err := _Memcoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Memcoin *MemcoinFilterer) ParseApproval(log types.Log) (*MemcoinApproval, error) {
	event := new(MemcoinApproval)
	if err := _Memcoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemcoinLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Memcoin contract.
type MemcoinLiquidityAddedIterator struct {
	Event *MemcoinLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *MemcoinLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinLiquidityAdded)
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
		it.Event = new(MemcoinLiquidityAdded)
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
func (it *MemcoinLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinLiquidityAdded represents a LiquidityAdded event raised by the Memcoin contract.
type MemcoinLiquidityAdded struct {
	Tokens *big.Int
	Eth    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x38f8a0c92f4c5b0b6877f878cb4c0c8d348a47b76d716c8e78f425043df9515b.
//
// Solidity: event LiquidityAdded(uint256 tokens, uint256 eth)
func (_Memcoin *MemcoinFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*MemcoinLiquidityAddedIterator, error) {

	logs, sub, err := _Memcoin.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &MemcoinLiquidityAddedIterator{contract: _Memcoin.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x38f8a0c92f4c5b0b6877f878cb4c0c8d348a47b76d716c8e78f425043df9515b.
//
// Solidity: event LiquidityAdded(uint256 tokens, uint256 eth)
func (_Memcoin *MemcoinFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *MemcoinLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Memcoin.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinLiquidityAdded)
				if err := _Memcoin.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x38f8a0c92f4c5b0b6877f878cb4c0c8d348a47b76d716c8e78f425043df9515b.
//
// Solidity: event LiquidityAdded(uint256 tokens, uint256 eth)
func (_Memcoin *MemcoinFilterer) ParseLiquidityAdded(log types.Log) (*MemcoinLiquidityAdded, error) {
	event := new(MemcoinLiquidityAdded)
	if err := _Memcoin.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemcoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Memcoin contract.
type MemcoinOwnershipTransferredIterator struct {
	Event *MemcoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MemcoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinOwnershipTransferred)
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
		it.Event = new(MemcoinOwnershipTransferred)
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
func (it *MemcoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinOwnershipTransferred represents a OwnershipTransferred event raised by the Memcoin contract.
type MemcoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Memcoin *MemcoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MemcoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Memcoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MemcoinOwnershipTransferredIterator{contract: _Memcoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Memcoin *MemcoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MemcoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Memcoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinOwnershipTransferred)
				if err := _Memcoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Memcoin *MemcoinFilterer) ParseOwnershipTransferred(log types.Log) (*MemcoinOwnershipTransferred, error) {
	event := new(MemcoinOwnershipTransferred)
	if err := _Memcoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemcoinTaxCollectedIterator is returned from FilterTaxCollected and is used to iterate over the raw logs and unpacked data for TaxCollected events raised by the Memcoin contract.
type MemcoinTaxCollectedIterator struct {
	Event *MemcoinTaxCollected // Event containing the contract specifics and raw log

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
func (it *MemcoinTaxCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinTaxCollected)
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
		it.Event = new(MemcoinTaxCollected)
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
func (it *MemcoinTaxCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinTaxCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinTaxCollected represents a TaxCollected event raised by the Memcoin contract.
type MemcoinTaxCollected struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaxCollected is a free log retrieval operation binding the contract event 0x092c6768ba45dfdc937f94a27a0c5be8334952f1558bc7dd41b3e8707502e1b6.
//
// Solidity: event TaxCollected(uint256 amount)
func (_Memcoin *MemcoinFilterer) FilterTaxCollected(opts *bind.FilterOpts) (*MemcoinTaxCollectedIterator, error) {

	logs, sub, err := _Memcoin.contract.FilterLogs(opts, "TaxCollected")
	if err != nil {
		return nil, err
	}
	return &MemcoinTaxCollectedIterator{contract: _Memcoin.contract, event: "TaxCollected", logs: logs, sub: sub}, nil
}

// WatchTaxCollected is a free log subscription operation binding the contract event 0x092c6768ba45dfdc937f94a27a0c5be8334952f1558bc7dd41b3e8707502e1b6.
//
// Solidity: event TaxCollected(uint256 amount)
func (_Memcoin *MemcoinFilterer) WatchTaxCollected(opts *bind.WatchOpts, sink chan<- *MemcoinTaxCollected) (event.Subscription, error) {

	logs, sub, err := _Memcoin.contract.WatchLogs(opts, "TaxCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinTaxCollected)
				if err := _Memcoin.contract.UnpackLog(event, "TaxCollected", log); err != nil {
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

// ParseTaxCollected is a log parse operation binding the contract event 0x092c6768ba45dfdc937f94a27a0c5be8334952f1558bc7dd41b3e8707502e1b6.
//
// Solidity: event TaxCollected(uint256 amount)
func (_Memcoin *MemcoinFilterer) ParseTaxCollected(log types.Log) (*MemcoinTaxCollected, error) {
	event := new(MemcoinTaxCollected)
	if err := _Memcoin.contract.UnpackLog(event, "TaxCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemcoinTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the Memcoin contract.
type MemcoinTokensMintedIterator struct {
	Event *MemcoinTokensMinted // Event containing the contract specifics and raw log

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
func (it *MemcoinTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinTokensMinted)
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
		it.Event = new(MemcoinTokensMinted)
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
func (it *MemcoinTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinTokensMinted represents a TokensMinted event raised by the Memcoin contract.
type MemcoinTokensMinted struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x3f2c9d57c068687834f0de942a9babb9e5acab57d516d3480a3c16ee165a4273.
//
// Solidity: event TokensMinted(address indexed to, uint256 amount)
func (_Memcoin *MemcoinFilterer) FilterTokensMinted(opts *bind.FilterOpts, to []common.Address) (*MemcoinTokensMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Memcoin.contract.FilterLogs(opts, "TokensMinted", toRule)
	if err != nil {
		return nil, err
	}
	return &MemcoinTokensMintedIterator{contract: _Memcoin.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x3f2c9d57c068687834f0de942a9babb9e5acab57d516d3480a3c16ee165a4273.
//
// Solidity: event TokensMinted(address indexed to, uint256 amount)
func (_Memcoin *MemcoinFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *MemcoinTokensMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Memcoin.contract.WatchLogs(opts, "TokensMinted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinTokensMinted)
				if err := _Memcoin.contract.UnpackLog(event, "TokensMinted", log); err != nil {
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

// ParseTokensMinted is a log parse operation binding the contract event 0x3f2c9d57c068687834f0de942a9babb9e5acab57d516d3480a3c16ee165a4273.
//
// Solidity: event TokensMinted(address indexed to, uint256 amount)
func (_Memcoin *MemcoinFilterer) ParseTokensMinted(log types.Log) (*MemcoinTokensMinted, error) {
	event := new(MemcoinTokensMinted)
	if err := _Memcoin.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemcoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Memcoin contract.
type MemcoinTransferIterator struct {
	Event *MemcoinTransfer // Event containing the contract specifics and raw log

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
func (it *MemcoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MemcoinTransfer)
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
		it.Event = new(MemcoinTransfer)
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
func (it *MemcoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MemcoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MemcoinTransfer represents a Transfer event raised by the Memcoin contract.
type MemcoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Memcoin *MemcoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MemcoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Memcoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MemcoinTransferIterator{contract: _Memcoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Memcoin *MemcoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MemcoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Memcoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MemcoinTransfer)
				if err := _Memcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Memcoin *MemcoinFilterer) ParseTransfer(log types.Log) (*MemcoinTransfer, error) {
	event := new(MemcoinTransfer)
	if err := _Memcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
