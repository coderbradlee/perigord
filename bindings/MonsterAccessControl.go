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

// MonsterAccessControlABI is the input ABI used to generate the binding from.
const MonsterAccessControlABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// MonsterAccessControlBin is the compiled bytecode used for deploying new contracts.
const MonsterAccessControlBin = `60806040526002805460a060020a60ff021916905534801561002057600080fd5b5061043d806100306000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461009d5780630a0f8168146100ce57806327d7874c146100e35780632ba73c15146101065780633f4ba83a146101275780634e0a33791461013c5780635c975abb1461015d5780638456cb5914610186578063b047fb501461019b575b600080fd5b3480156100a957600080fd5b506100b26101b0565b60408051600160a060020a039092168252519081900360200190f35b3480156100da57600080fd5b506100b26101bf565b3480156100ef57600080fd5b50610104600160a060020a03600435166101ce565b005b34801561011257600080fd5b50610104600160a060020a0360043516610229565b34801561013357600080fd5b50610104610284565b34801561014857600080fd5b50610104600160a060020a03600435166102e4565b34801561016957600080fd5b5061017261033f565b604080519115158252519081900360200190f35b34801561019257600080fd5b50610104610360565b3480156101a757600080fd5b506100b2610402565b600154600160a060020a031681565b600054600160a060020a031681565b600054600160a060020a031633146101e557600080fd5b600160a060020a03811615156101fa57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461024057600080fd5b600160a060020a038116151561025557600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461029b57600080fd5b60025474010000000000000000000000000000000000000000900460ff1615156102c457600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600054600160a060020a031633146102fb57600080fd5b600160a060020a038116151561031057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025474010000000000000000000000000000000000000000900460ff1681565b600254600160a060020a03163314806103835750600054600160a060020a031633145b806103985750600154600160a060020a031633145b15156103a357600080fd5b60025474010000000000000000000000000000000000000000900460ff16156103cb57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600254600160a060020a0316815600a165627a7a723058203a980c5d0bb7b84efb0d0c7b698f280cd0f7dce74029a98e4e2cdddd9b87fe2c0029`

// DeployMonsterAccessControl deploys a new Ethereum contract, binding an instance of MonsterAccessControl to it.
func DeployMonsterAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonsterAccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterAccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonsterAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonsterAccessControl{MonsterAccessControlCaller: MonsterAccessControlCaller{contract: contract}, MonsterAccessControlTransactor: MonsterAccessControlTransactor{contract: contract}, MonsterAccessControlFilterer: MonsterAccessControlFilterer{contract: contract}}, nil
}

// MonsterAccessControl is an auto generated Go binding around an Ethereum contract.
type MonsterAccessControl struct {
	MonsterAccessControlCaller     // Read-only binding to the contract
	MonsterAccessControlTransactor // Write-only binding to the contract
	MonsterAccessControlFilterer   // Log filterer for contract events
}

// MonsterAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonsterAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonsterAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonsterAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonsterAccessControlSession struct {
	Contract     *MonsterAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MonsterAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonsterAccessControlCallerSession struct {
	Contract *MonsterAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MonsterAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonsterAccessControlTransactorSession struct {
	Contract     *MonsterAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MonsterAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonsterAccessControlRaw struct {
	Contract *MonsterAccessControl // Generic contract binding to access the raw methods on
}

// MonsterAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonsterAccessControlCallerRaw struct {
	Contract *MonsterAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// MonsterAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonsterAccessControlTransactorRaw struct {
	Contract *MonsterAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonsterAccessControl creates a new instance of MonsterAccessControl, bound to a specific deployed contract.
func NewMonsterAccessControl(address common.Address, backend bind.ContractBackend) (*MonsterAccessControl, error) {
	contract, err := bindMonsterAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonsterAccessControl{MonsterAccessControlCaller: MonsterAccessControlCaller{contract: contract}, MonsterAccessControlTransactor: MonsterAccessControlTransactor{contract: contract}, MonsterAccessControlFilterer: MonsterAccessControlFilterer{contract: contract}}, nil
}

// NewMonsterAccessControlCaller creates a new read-only instance of MonsterAccessControl, bound to a specific deployed contract.
func NewMonsterAccessControlCaller(address common.Address, caller bind.ContractCaller) (*MonsterAccessControlCaller, error) {
	contract, err := bindMonsterAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterAccessControlCaller{contract: contract}, nil
}

// NewMonsterAccessControlTransactor creates a new write-only instance of MonsterAccessControl, bound to a specific deployed contract.
func NewMonsterAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*MonsterAccessControlTransactor, error) {
	contract, err := bindMonsterAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterAccessControlTransactor{contract: contract}, nil
}

// NewMonsterAccessControlFilterer creates a new log filterer instance of MonsterAccessControl, bound to a specific deployed contract.
func NewMonsterAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*MonsterAccessControlFilterer, error) {
	contract, err := bindMonsterAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonsterAccessControlFilterer{contract: contract}, nil
}

// bindMonsterAccessControl binds a generic wrapper to an already deployed contract.
func bindMonsterAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterAccessControl *MonsterAccessControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterAccessControl.Contract.MonsterAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterAccessControl *MonsterAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.MonsterAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterAccessControl *MonsterAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.MonsterAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterAccessControl *MonsterAccessControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterAccessControl *MonsterAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterAccessControl *MonsterAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterAccessControl.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlSession) CeoAddress() (common.Address, error) {
	return _MonsterAccessControl.Contract.CeoAddress(&_MonsterAccessControl.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlCallerSession) CeoAddress() (common.Address, error) {
	return _MonsterAccessControl.Contract.CeoAddress(&_MonsterAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterAccessControl.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlSession) CfoAddress() (common.Address, error) {
	return _MonsterAccessControl.Contract.CfoAddress(&_MonsterAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlCallerSession) CfoAddress() (common.Address, error) {
	return _MonsterAccessControl.Contract.CfoAddress(&_MonsterAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterAccessControl.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlSession) CooAddress() (common.Address, error) {
	return _MonsterAccessControl.Contract.CooAddress(&_MonsterAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterAccessControl *MonsterAccessControlCallerSession) CooAddress() (common.Address, error) {
	return _MonsterAccessControl.Contract.CooAddress(&_MonsterAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterAccessControl *MonsterAccessControlCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterAccessControl.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterAccessControl *MonsterAccessControlSession) Paused() (bool, error) {
	return _MonsterAccessControl.Contract.Paused(&_MonsterAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterAccessControl *MonsterAccessControlCallerSession) Paused() (bool, error) {
	return _MonsterAccessControl.Contract.Paused(&_MonsterAccessControl.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterAccessControl *MonsterAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterAccessControl *MonsterAccessControlSession) Pause() (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.Pause(&_MonsterAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterAccessControl *MonsterAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.Pause(&_MonsterAccessControl.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterAccessControl *MonsterAccessControlTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterAccessControl *MonsterAccessControlSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.SetCEO(&_MonsterAccessControl.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterAccessControl *MonsterAccessControlTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.SetCEO(&_MonsterAccessControl.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterAccessControl *MonsterAccessControlTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterAccessControl *MonsterAccessControlSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.SetCFO(&_MonsterAccessControl.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterAccessControl *MonsterAccessControlTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.SetCFO(&_MonsterAccessControl.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterAccessControl *MonsterAccessControlTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterAccessControl *MonsterAccessControlSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.SetCOO(&_MonsterAccessControl.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterAccessControl *MonsterAccessControlTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.SetCOO(&_MonsterAccessControl.TransactOpts, _newCOO)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterAccessControl *MonsterAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterAccessControl *MonsterAccessControlSession) Unpause() (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.Unpause(&_MonsterAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterAccessControl *MonsterAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonsterAccessControl.Contract.Unpause(&_MonsterAccessControl.TransactOpts)
}

// MonsterAccessControlContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the MonsterAccessControl contract.
type MonsterAccessControlContractUpgradeIterator struct {
	Event *MonsterAccessControlContractUpgrade // Event containing the contract specifics and raw log

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
func (it *MonsterAccessControlContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterAccessControlContractUpgrade)
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
		it.Event = new(MonsterAccessControlContractUpgrade)
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
func (it *MonsterAccessControlContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterAccessControlContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterAccessControlContractUpgrade represents a ContractUpgrade event raised by the MonsterAccessControl contract.
type MonsterAccessControlContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterAccessControl *MonsterAccessControlFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*MonsterAccessControlContractUpgradeIterator, error) {

	logs, sub, err := _MonsterAccessControl.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &MonsterAccessControlContractUpgradeIterator{contract: _MonsterAccessControl.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterAccessControl *MonsterAccessControlFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *MonsterAccessControlContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _MonsterAccessControl.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterAccessControlContractUpgrade)
				if err := _MonsterAccessControl.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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
