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

// TokenERC20ABI is the input ABI used to generate the binding from.
const TokenERC20ABI = "[]"

// TokenERC20Bin is the compiled bytecode used for deploying new contracts.
const TokenERC20Bin = `6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582036a906f9ecd17a17a92bbd312ec3317cc4bf552bff0f66fe14174b59dacfcf9f0029`

// DeployTokenERC20 deploys a new Ethereum contract, binding an instance of TokenERC20 to it.
func DeployTokenERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// TokenERC20 is an auto generated Go binding around an Ethereum contract.
type TokenERC20 struct {
	TokenERC20Caller     // Read-only binding to the contract
	TokenERC20Transactor // Write-only binding to the contract
	TokenERC20Filterer   // Log filterer for contract events
}

// TokenERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenERC20Session struct {
	Contract     *TokenERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenERC20CallerSession struct {
	Contract *TokenERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenERC20TransactorSession struct {
	Contract     *TokenERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenERC20Raw struct {
	Contract *TokenERC20 // Generic contract binding to access the raw methods on
}

// TokenERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenERC20CallerRaw struct {
	Contract *TokenERC20Caller // Generic read-only contract binding to access the raw methods on
}

// TokenERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenERC20TransactorRaw struct {
	Contract *TokenERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenERC20 creates a new instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20(address common.Address, backend bind.ContractBackend) (*TokenERC20, error) {
	contract, err := bindTokenERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// NewTokenERC20Caller creates a new read-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Caller(address common.Address, caller bind.ContractCaller) (*TokenERC20Caller, error) {
	contract, err := bindTokenERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Caller{contract: contract}, nil
}

// NewTokenERC20Transactor creates a new write-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenERC20Transactor, error) {
	contract, err := bindTokenERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Transactor{contract: contract}, nil
}

// NewTokenERC20Filterer creates a new log filterer instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenERC20Filterer, error) {
	contract, err := bindTokenERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Filterer{contract: contract}, nil
}

// bindTokenERC20 binds a generic wrapper to an already deployed contract.
func bindTokenERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.TokenERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transact(opts, method, params...)
}
