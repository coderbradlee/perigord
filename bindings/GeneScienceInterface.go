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

// GeneScienceInterfaceABI is the input ABI used to generate the binding from.
const GeneScienceInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"genes1\",\"type\":\"uint256\"},{\"name\":\"genes2\",\"type\":\"uint256\"},{\"name\":\"targetBlock\",\"type\":\"uint256\"}],\"name\":\"mixGenes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGeneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// GeneScienceInterfaceBin is the compiled bytecode used for deploying new contracts.
const GeneScienceInterfaceBin = `608060405234801561001057600080fd5b5060d68061001f6000396000f30060806040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630d9f5aed8114604d57806354c15b8214607a575b600080fd5b348015605857600080fd5b50606860043560243560443560a0565b60408051918252519081900360200190f35b348015608557600080fd5b50608c60a5565b604080519115158252519081900360200190f35b501690565b6001905600a165627a7a723058209d83d545dbad3670d6aaa1d90c7766db7febe6e1f96c4475454b9bb4bf0344940029`

// DeployGeneScienceInterface deploys a new Ethereum contract, binding an instance of GeneScienceInterface to it.
func DeployGeneScienceInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GeneScienceInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneScienceInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GeneScienceInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GeneScienceInterface{GeneScienceInterfaceCaller: GeneScienceInterfaceCaller{contract: contract}, GeneScienceInterfaceTransactor: GeneScienceInterfaceTransactor{contract: contract}, GeneScienceInterfaceFilterer: GeneScienceInterfaceFilterer{contract: contract}}, nil
}

// GeneScienceInterface is an auto generated Go binding around an Ethereum contract.
type GeneScienceInterface struct {
	GeneScienceInterfaceCaller     // Read-only binding to the contract
	GeneScienceInterfaceTransactor // Write-only binding to the contract
	GeneScienceInterfaceFilterer   // Log filterer for contract events
}

// GeneScienceInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneScienceInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneScienceInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneScienceInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneScienceInterfaceSession struct {
	Contract     *GeneScienceInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GeneScienceInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneScienceInterfaceCallerSession struct {
	Contract *GeneScienceInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// GeneScienceInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneScienceInterfaceTransactorSession struct {
	Contract     *GeneScienceInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// GeneScienceInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneScienceInterfaceRaw struct {
	Contract *GeneScienceInterface // Generic contract binding to access the raw methods on
}

// GeneScienceInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneScienceInterfaceCallerRaw struct {
	Contract *GeneScienceInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// GeneScienceInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneScienceInterfaceTransactorRaw struct {
	Contract *GeneScienceInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneScienceInterface creates a new instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterface(address common.Address, backend bind.ContractBackend) (*GeneScienceInterface, error) {
	contract, err := bindGeneScienceInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterface{GeneScienceInterfaceCaller: GeneScienceInterfaceCaller{contract: contract}, GeneScienceInterfaceTransactor: GeneScienceInterfaceTransactor{contract: contract}, GeneScienceInterfaceFilterer: GeneScienceInterfaceFilterer{contract: contract}}, nil
}

// NewGeneScienceInterfaceCaller creates a new read-only instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceCaller(address common.Address, caller bind.ContractCaller) (*GeneScienceInterfaceCaller, error) {
	contract, err := bindGeneScienceInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceCaller{contract: contract}, nil
}

// NewGeneScienceInterfaceTransactor creates a new write-only instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneScienceInterfaceTransactor, error) {
	contract, err := bindGeneScienceInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceTransactor{contract: contract}, nil
}

// NewGeneScienceInterfaceFilterer creates a new log filterer instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneScienceInterfaceFilterer, error) {
	contract, err := bindGeneScienceInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceFilterer{contract: contract}, nil
}

// bindGeneScienceInterface binds a generic wrapper to an already deployed contract.
func bindGeneScienceInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneScienceInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneScienceInterface *GeneScienceInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GeneScienceInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneScienceInterface *GeneScienceInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneScienceInterface *GeneScienceInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.contract.Transact(opts, method, params...)
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() constant returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceCaller) IsGeneScience(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GeneScienceInterface.contract.Call(opts, out, "isGeneScience")
	return *ret0, err
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() constant returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceSession) IsGeneScience() (bool, error) {
	return _GeneScienceInterface.Contract.IsGeneScience(&_GeneScienceInterface.CallOpts)
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() constant returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceCallerSession) IsGeneScience() (bool, error) {
	return _GeneScienceInterface.Contract.IsGeneScience(&_GeneScienceInterface.CallOpts)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(genes1 uint256, genes2 uint256, targetBlock uint256) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceTransactor) MixGenes(opts *bind.TransactOpts, genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.contract.Transact(opts, "mixGenes", genes1, genes2, targetBlock)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(genes1 uint256, genes2 uint256, targetBlock uint256) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceSession) MixGenes(genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.MixGenes(&_GeneScienceInterface.TransactOpts, genes1, genes2, targetBlock)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(genes1 uint256, genes2 uint256, targetBlock uint256) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceTransactorSession) MixGenes(genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.MixGenes(&_GeneScienceInterface.TransactOpts, genes1, genes2, targetBlock)
}
