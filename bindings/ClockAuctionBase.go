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

// ClockAuctionBaseABI is the input ABI used to generate the binding from.
const ClockAuctionBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"}]"

// ClockAuctionBaseBin is the compiled bytecode used for deploying new contracts.
const ClockAuctionBaseBin = `608060405234801561001057600080fd5b5060fa8061001f6000396000f30060806040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166383b5ff8b8114604d578063dd1b7a0f146071575b600080fd5b348015605857600080fd5b50605f60ac565b60408051918252519081900360200190f35b348015607c57600080fd5b50608360b2565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60015481565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058202819e50f53cd8bd2db0c80e6972f3d46c0b02cf37b88c0ff4e0156b0ddb46afd0029`

// DeployClockAuctionBase deploys a new Ethereum contract, binding an instance of ClockAuctionBase to it.
func DeployClockAuctionBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClockAuctionBase, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClockAuctionBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClockAuctionBase{ClockAuctionBaseCaller: ClockAuctionBaseCaller{contract: contract}, ClockAuctionBaseTransactor: ClockAuctionBaseTransactor{contract: contract}, ClockAuctionBaseFilterer: ClockAuctionBaseFilterer{contract: contract}}, nil
}

// ClockAuctionBase is an auto generated Go binding around an Ethereum contract.
type ClockAuctionBase struct {
	ClockAuctionBaseCaller     // Read-only binding to the contract
	ClockAuctionBaseTransactor // Write-only binding to the contract
	ClockAuctionBaseFilterer   // Log filterer for contract events
}

// ClockAuctionBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockAuctionBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockAuctionBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockAuctionBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockAuctionBaseSession struct {
	Contract     *ClockAuctionBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockAuctionBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockAuctionBaseCallerSession struct {
	Contract *ClockAuctionBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ClockAuctionBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockAuctionBaseTransactorSession struct {
	Contract     *ClockAuctionBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ClockAuctionBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockAuctionBaseRaw struct {
	Contract *ClockAuctionBase // Generic contract binding to access the raw methods on
}

// ClockAuctionBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockAuctionBaseCallerRaw struct {
	Contract *ClockAuctionBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ClockAuctionBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockAuctionBaseTransactorRaw struct {
	Contract *ClockAuctionBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockAuctionBase creates a new instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBase(address common.Address, backend bind.ContractBackend) (*ClockAuctionBase, error) {
	contract, err := bindClockAuctionBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBase{ClockAuctionBaseCaller: ClockAuctionBaseCaller{contract: contract}, ClockAuctionBaseTransactor: ClockAuctionBaseTransactor{contract: contract}, ClockAuctionBaseFilterer: ClockAuctionBaseFilterer{contract: contract}}, nil
}

// NewClockAuctionBaseCaller creates a new read-only instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseCaller(address common.Address, caller bind.ContractCaller) (*ClockAuctionBaseCaller, error) {
	contract, err := bindClockAuctionBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseCaller{contract: contract}, nil
}

// NewClockAuctionBaseTransactor creates a new write-only instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockAuctionBaseTransactor, error) {
	contract, err := bindClockAuctionBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseTransactor{contract: contract}, nil
}

// NewClockAuctionBaseFilterer creates a new log filterer instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockAuctionBaseFilterer, error) {
	contract, err := bindClockAuctionBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseFilterer{contract: contract}, nil
}

// bindClockAuctionBase binds a generic wrapper to an already deployed contract.
func bindClockAuctionBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuctionBase.Contract.ClockAuctionBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.ClockAuctionBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.ClockAuctionBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuctionBase *ClockAuctionBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuctionBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuctionBase *ClockAuctionBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuctionBase *ClockAuctionBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.contract.Transact(opts, method, params...)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuctionBase *ClockAuctionBaseCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClockAuctionBase.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuctionBase *ClockAuctionBaseSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuctionBase.Contract.NonFungibleContract(&_ClockAuctionBase.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuctionBase *ClockAuctionBaseCallerSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuctionBase.Contract.NonFungibleContract(&_ClockAuctionBase.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClockAuctionBase.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseSession) OwnerCut() (*big.Int, error) {
	return _ClockAuctionBase.Contract.OwnerCut(&_ClockAuctionBase.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseCallerSession) OwnerCut() (*big.Int, error) {
	return _ClockAuctionBase.Contract.OwnerCut(&_ClockAuctionBase.CallOpts)
}

// ClockAuctionBaseAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCancelledIterator struct {
	Event *ClockAuctionBaseAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionCancelled)
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
		it.Event = new(ClockAuctionBaseAuctionCancelled)
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
func (it *ClockAuctionBaseAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionCancelled represents a AuctionCancelled event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionCancelledIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionCancelledIterator{contract: _ClockAuctionBase.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionCancelled)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ClockAuctionBaseAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCreatedIterator struct {
	Event *ClockAuctionBaseAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionCreated)
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
		it.Event = new(ClockAuctionBaseAuctionCreated)
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
func (it *ClockAuctionBaseAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionCreated represents a AuctionCreated event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionCreatedIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionCreatedIterator{contract: _ClockAuctionBase.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionCreated)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ClockAuctionBaseAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionSuccessfulIterator struct {
	Event *ClockAuctionBaseAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionSuccessful)
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
		it.Event = new(ClockAuctionBaseAuctionSuccessful)
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
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionSuccessful represents a AuctionSuccessful event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionSuccessfulIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionSuccessfulIterator{contract: _ClockAuctionBase.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionSuccessful)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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
