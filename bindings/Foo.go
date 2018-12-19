// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FooABI is the input ABI used to generate the binding from.
const FooABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"bar\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FooBin is the compiled bytecode used for deploying new contracts.
const FooBin = `6080604052348015600f57600080fd5b5060998061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663febb0f7e81146043575b600080fd5b348015604e57600080fd5b5060556067565b60408051918252519081900360200190f35b610539905600a165627a7a723058206cceea502b3a5e6c425e3c3686f77036ec670825da7497362d50f52a87c6dad00029`

// DeployFoo deploys a new Ethereum contract, binding an instance of Foo to it.
func DeployFoo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Foo, error) {
	parsed, err := abi.JSON(strings.NewReader(FooABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FooBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Foo{FooCaller: FooCaller{contract: contract}, FooTransactor: FooTransactor{contract: contract}, FooFilterer: FooFilterer{contract: contract}}, nil
}

// Foo is an auto generated Go binding around an Ethereum contract.
type Foo struct {
	FooCaller     // Read-only binding to the contract
	FooTransactor // Write-only binding to the contract
	FooFilterer   // Log filterer for contract events
}

// FooCaller is an auto generated read-only Go binding around an Ethereum contract.
type FooCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FooTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FooFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FooSession struct {
	Contract     *Foo              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FooCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FooCallerSession struct {
	Contract *FooCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FooTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FooTransactorSession struct {
	Contract     *FooTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FooRaw is an auto generated low-level Go binding around an Ethereum contract.
type FooRaw struct {
	Contract *Foo // Generic contract binding to access the raw methods on
}

// FooCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FooCallerRaw struct {
	Contract *FooCaller // Generic read-only contract binding to access the raw methods on
}

// FooTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FooTransactorRaw struct {
	Contract *FooTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoo creates a new instance of Foo, bound to a specific deployed contract.
func NewFoo(address common.Address, backend bind.ContractBackend) (*Foo, error) {
	contract, err := bindFoo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foo{FooCaller: FooCaller{contract: contract}, FooTransactor: FooTransactor{contract: contract}, FooFilterer: FooFilterer{contract: contract}}, nil
}

// NewFooCaller creates a new read-only instance of Foo, bound to a specific deployed contract.
func NewFooCaller(address common.Address, caller bind.ContractCaller) (*FooCaller, error) {
	contract, err := bindFoo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FooCaller{contract: contract}, nil
}

// NewFooTransactor creates a new write-only instance of Foo, bound to a specific deployed contract.
func NewFooTransactor(address common.Address, transactor bind.ContractTransactor) (*FooTransactor, error) {
	contract, err := bindFoo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FooTransactor{contract: contract}, nil
}

// NewFooFilterer creates a new log filterer instance of Foo, bound to a specific deployed contract.
func NewFooFilterer(address common.Address, filterer bind.ContractFilterer) (*FooFilterer, error) {
	contract, err := bindFoo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FooFilterer{contract: contract}, nil
}

// bindFoo binds a generic wrapper to an already deployed contract.
func bindFoo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FooABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foo *FooRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Foo.Contract.FooCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foo *FooRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foo.Contract.FooTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foo *FooRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foo.Contract.FooTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foo *FooCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Foo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foo *FooTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foo *FooTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foo.Contract.contract.Transact(opts, method, params...)
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() constant returns(int256)
func (_Foo *FooCaller) Bar(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Foo.contract.Call(opts, out, "bar")
	return *ret0, err
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() constant returns(int256)
func (_Foo *FooSession) Bar() (*big.Int, error) {
	return _Foo.Contract.Bar(&_Foo.CallOpts)
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() constant returns(int256)
func (_Foo *FooCallerSession) Bar() (*big.Int, error) {
	return _Foo.Contract.Bar(&_Foo.CallOpts)
}
