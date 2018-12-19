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

// MonsterBaseABI is the input ABI used to generate the binding from.
const MonsterBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsterIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsterIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"MonsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// MonsterBaseBin is the compiled bytecode used for deploying new contracts.
const MonsterBaseBin = `6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a806102205261008590600390600e61009d565b50600f60055534801561009757600080fd5b5061015d565b6002830191839082156101295791602002820160005b838211156100f757835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026100b3565b80156101275782816101000a81549063ffffffff02191690556004016020816003010492830192600103026100f7565b505b50610135929150610139565b5090565b61015a91905b8082111561013557805463ffffffff1916815560010161013f565b90565b6106758061016c6000396000f3006080604052600436106100f05763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce7981146100f55780630a0f81681461012657806321717ebf1461013b57806327d7874c146101505780632ba73c15146101735780633f4ba83a1461019457806346116e6f146101a95780634e0a3379146101c15780635663896e146101e25780635c975abb146101fa5780637a7d4937146102235780637d55aeea1461024a5780638456cb59146102625780639d6fac6f14610277578063ad4b558c146102a8578063b047fb50146102c0578063e6cbe351146102d5575b600080fd5b34801561010157600080fd5b5061010a6102ea565b60408051600160a060020a039092168252519081900360200190f35b34801561013257600080fd5b5061010a6102f9565b34801561014757600080fd5b5061010a610308565b34801561015c57600080fd5b50610171600160a060020a0360043516610317565b005b34801561017f57600080fd5b50610171600160a060020a0360043516610372565b3480156101a057600080fd5b506101716103cd565b3480156101b557600080fd5b5061010a60043561042d565b3480156101cd57600080fd5b50610171600160a060020a0360043516610448565b3480156101ee57600080fd5b506101716004356104a3565b34801561020657600080fd5b5061020f6104ff565b604080519115158252519081900360200190f35b34801561022f57600080fd5b50610238610520565b60408051918252519081900360200190f35b34801561025657600080fd5b5061010a600435610526565b34801561026e57600080fd5b50610171610541565b34801561028357600080fd5b5061028f6004356105e3565b6040805163ffffffff9092168252519081900360200190f35b3480156102b457600080fd5b5061010a600435610610565b3480156102cc57600080fd5b5061010a61062b565b3480156102e157600080fd5b5061010a61063a565b600154600160a060020a031681565b600054600160a060020a031681565b600c54600160a060020a031681565b600054600160a060020a0316331461032e57600080fd5b600160a060020a038116151561034357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461038957600080fd5b600160a060020a038116151561039e57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146103e457600080fd5b60025474010000000000000000000000000000000000000000900460ff16151561040d57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600054600160a060020a0316331461045f57600080fd5b600160a060020a038116151561047457600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163314806104c65750600054600160a060020a031633145b806104db5750600154600160a060020a031633145b15156104e657600080fd5b60035463ffffffff1681106104fa57600080fd5b600555565b60025474010000000000000000000000000000000000000000900460ff1681565b60055481565b600960205260009081526040902054600160a060020a031681565b600254600160a060020a03163314806105645750600054600160a060020a031633145b806105795750600154600160a060020a031633145b151561058457600080fd5b60025474010000000000000000000000000000000000000000900460ff16156105ac57600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600381600e81106105f057fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b600254600160a060020a031681565b600b54600160a060020a0316815600a165627a7a72305820f5d57957246c721659fbeb52502d8324ee397ac82ca8ac4e9566032701056dc30029`

// DeployMonsterBase deploys a new Ethereum contract, binding an instance of MonsterBase to it.
func DeployMonsterBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonsterBase, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonsterBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonsterBase{MonsterBaseCaller: MonsterBaseCaller{contract: contract}, MonsterBaseTransactor: MonsterBaseTransactor{contract: contract}, MonsterBaseFilterer: MonsterBaseFilterer{contract: contract}}, nil
}

// MonsterBase is an auto generated Go binding around an Ethereum contract.
type MonsterBase struct {
	MonsterBaseCaller     // Read-only binding to the contract
	MonsterBaseTransactor // Write-only binding to the contract
	MonsterBaseFilterer   // Log filterer for contract events
}

// MonsterBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonsterBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonsterBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonsterBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonsterBaseSession struct {
	Contract     *MonsterBase      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MonsterBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonsterBaseCallerSession struct {
	Contract *MonsterBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MonsterBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonsterBaseTransactorSession struct {
	Contract     *MonsterBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MonsterBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonsterBaseRaw struct {
	Contract *MonsterBase // Generic contract binding to access the raw methods on
}

// MonsterBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonsterBaseCallerRaw struct {
	Contract *MonsterBaseCaller // Generic read-only contract binding to access the raw methods on
}

// MonsterBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonsterBaseTransactorRaw struct {
	Contract *MonsterBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonsterBase creates a new instance of MonsterBase, bound to a specific deployed contract.
func NewMonsterBase(address common.Address, backend bind.ContractBackend) (*MonsterBase, error) {
	contract, err := bindMonsterBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonsterBase{MonsterBaseCaller: MonsterBaseCaller{contract: contract}, MonsterBaseTransactor: MonsterBaseTransactor{contract: contract}, MonsterBaseFilterer: MonsterBaseFilterer{contract: contract}}, nil
}

// NewMonsterBaseCaller creates a new read-only instance of MonsterBase, bound to a specific deployed contract.
func NewMonsterBaseCaller(address common.Address, caller bind.ContractCaller) (*MonsterBaseCaller, error) {
	contract, err := bindMonsterBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterBaseCaller{contract: contract}, nil
}

// NewMonsterBaseTransactor creates a new write-only instance of MonsterBase, bound to a specific deployed contract.
func NewMonsterBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*MonsterBaseTransactor, error) {
	contract, err := bindMonsterBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterBaseTransactor{contract: contract}, nil
}

// NewMonsterBaseFilterer creates a new log filterer instance of MonsterBase, bound to a specific deployed contract.
func NewMonsterBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*MonsterBaseFilterer, error) {
	contract, err := bindMonsterBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonsterBaseFilterer{contract: contract}, nil
}

// bindMonsterBase binds a generic wrapper to an already deployed contract.
func bindMonsterBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterBase *MonsterBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterBase.Contract.MonsterBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterBase *MonsterBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBase.Contract.MonsterBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterBase *MonsterBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterBase.Contract.MonsterBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterBase *MonsterBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterBase *MonsterBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterBase *MonsterBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterBase.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterBase *MonsterBaseCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterBase *MonsterBaseSession) CeoAddress() (common.Address, error) {
	return _MonsterBase.Contract.CeoAddress(&_MonsterBase.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) CeoAddress() (common.Address, error) {
	return _MonsterBase.Contract.CeoAddress(&_MonsterBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterBase *MonsterBaseCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterBase *MonsterBaseSession) CfoAddress() (common.Address, error) {
	return _MonsterBase.Contract.CfoAddress(&_MonsterBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) CfoAddress() (common.Address, error) {
	return _MonsterBase.Contract.CfoAddress(&_MonsterBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterBase *MonsterBaseCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterBase *MonsterBaseSession) CooAddress() (common.Address, error) {
	return _MonsterBase.Contract.CooAddress(&_MonsterBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) CooAddress() (common.Address, error) {
	return _MonsterBase.Contract.CooAddress(&_MonsterBase.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterBase *MonsterBaseCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterBase *MonsterBaseSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _MonsterBase.Contract.Cooldowns(&_MonsterBase.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterBase *MonsterBaseCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _MonsterBase.Contract.Cooldowns(&_MonsterBase.CallOpts, arg0)
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseCaller) MonsterIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "monsterIndexToApproved", arg0)
	return *ret0, err
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseSession) MonsterIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _MonsterBase.Contract.MonsterIndexToApproved(&_MonsterBase.CallOpts, arg0)
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) MonsterIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _MonsterBase.Contract.MonsterIndexToApproved(&_MonsterBase.CallOpts, arg0)
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseCaller) MonsterIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "monsterIndexToOwner", arg0)
	return *ret0, err
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseSession) MonsterIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _MonsterBase.Contract.MonsterIndexToOwner(&_MonsterBase.CallOpts, arg0)
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) MonsterIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _MonsterBase.Contract.MonsterIndexToOwner(&_MonsterBase.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterBase *MonsterBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterBase *MonsterBaseSession) Paused() (bool, error) {
	return _MonsterBase.Contract.Paused(&_MonsterBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterBase *MonsterBaseCallerSession) Paused() (bool, error) {
	return _MonsterBase.Contract.Paused(&_MonsterBase.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterBase *MonsterBaseCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterBase *MonsterBaseSession) SaleAuction() (common.Address, error) {
	return _MonsterBase.Contract.SaleAuction(&_MonsterBase.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) SaleAuction() (common.Address, error) {
	return _MonsterBase.Contract.SaleAuction(&_MonsterBase.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterBase *MonsterBaseCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterBase *MonsterBaseSession) SecondsPerBlock() (*big.Int, error) {
	return _MonsterBase.Contract.SecondsPerBlock(&_MonsterBase.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterBase *MonsterBaseCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _MonsterBase.Contract.SecondsPerBlock(&_MonsterBase.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _MonsterBase.Contract.SireAllowedToAddress(&_MonsterBase.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _MonsterBase.Contract.SireAllowedToAddress(&_MonsterBase.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterBase *MonsterBaseCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBase.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterBase *MonsterBaseSession) SiringAuction() (common.Address, error) {
	return _MonsterBase.Contract.SiringAuction(&_MonsterBase.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterBase *MonsterBaseCallerSession) SiringAuction() (common.Address, error) {
	return _MonsterBase.Contract.SiringAuction(&_MonsterBase.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterBase *MonsterBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterBase *MonsterBaseSession) Pause() (*types.Transaction, error) {
	return _MonsterBase.Contract.Pause(&_MonsterBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterBase *MonsterBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _MonsterBase.Contract.Pause(&_MonsterBase.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterBase *MonsterBaseTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _MonsterBase.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterBase *MonsterBaseSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetCEO(&_MonsterBase.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterBase *MonsterBaseTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetCEO(&_MonsterBase.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterBase *MonsterBaseTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _MonsterBase.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterBase *MonsterBaseSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetCFO(&_MonsterBase.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterBase *MonsterBaseTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetCFO(&_MonsterBase.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterBase *MonsterBaseTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _MonsterBase.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterBase *MonsterBaseSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetCOO(&_MonsterBase.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterBase *MonsterBaseTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetCOO(&_MonsterBase.TransactOpts, _newCOO)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterBase *MonsterBaseTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _MonsterBase.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterBase *MonsterBaseSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetSecondsPerBlock(&_MonsterBase.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterBase *MonsterBaseTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _MonsterBase.Contract.SetSecondsPerBlock(&_MonsterBase.TransactOpts, secs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterBase *MonsterBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterBase *MonsterBaseSession) Unpause() (*types.Transaction, error) {
	return _MonsterBase.Contract.Unpause(&_MonsterBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterBase *MonsterBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonsterBase.Contract.Unpause(&_MonsterBase.TransactOpts)
}

// MonsterBaseBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the MonsterBase contract.
type MonsterBaseBirthIterator struct {
	Event *MonsterBaseBirth // Event containing the contract specifics and raw log

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
func (it *MonsterBaseBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBaseBirth)
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
		it.Event = new(MonsterBaseBirth)
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
func (it *MonsterBaseBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBaseBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBaseBirth represents a Birth event raised by the MonsterBase contract.
type MonsterBaseBirth struct {
	Owner     common.Address
	MonsterId *big.Int
	MatronId  *big.Int
	SireId    *big.Int
	Genes     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, MonsterId uint256, matronId uint256, sireId uint256, genes uint256)
func (_MonsterBase *MonsterBaseFilterer) FilterBirth(opts *bind.FilterOpts) (*MonsterBaseBirthIterator, error) {

	logs, sub, err := _MonsterBase.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &MonsterBaseBirthIterator{contract: _MonsterBase.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, MonsterId uint256, matronId uint256, sireId uint256, genes uint256)
func (_MonsterBase *MonsterBaseFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *MonsterBaseBirth) (event.Subscription, error) {

	logs, sub, err := _MonsterBase.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBaseBirth)
				if err := _MonsterBase.contract.UnpackLog(event, "Birth", log); err != nil {
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

// MonsterBaseContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the MonsterBase contract.
type MonsterBaseContractUpgradeIterator struct {
	Event *MonsterBaseContractUpgrade // Event containing the contract specifics and raw log

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
func (it *MonsterBaseContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBaseContractUpgrade)
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
		it.Event = new(MonsterBaseContractUpgrade)
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
func (it *MonsterBaseContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBaseContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBaseContractUpgrade represents a ContractUpgrade event raised by the MonsterBase contract.
type MonsterBaseContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterBase *MonsterBaseFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*MonsterBaseContractUpgradeIterator, error) {

	logs, sub, err := _MonsterBase.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &MonsterBaseContractUpgradeIterator{contract: _MonsterBase.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterBase *MonsterBaseFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *MonsterBaseContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _MonsterBase.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBaseContractUpgrade)
				if err := _MonsterBase.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// MonsterBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MonsterBase contract.
type MonsterBaseTransferIterator struct {
	Event *MonsterBaseTransfer // Event containing the contract specifics and raw log

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
func (it *MonsterBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBaseTransfer)
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
		it.Event = new(MonsterBaseTransfer)
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
func (it *MonsterBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBaseTransfer represents a Transfer event raised by the MonsterBase contract.
type MonsterBaseTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_MonsterBase *MonsterBaseFilterer) FilterTransfer(opts *bind.FilterOpts) (*MonsterBaseTransferIterator, error) {

	logs, sub, err := _MonsterBase.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &MonsterBaseTransferIterator{contract: _MonsterBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_MonsterBase *MonsterBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MonsterBaseTransfer) (event.Subscription, error) {

	logs, sub, err := _MonsterBase.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBaseTransfer)
				if err := _MonsterBase.contract.UnpackLog(event, "Transfer", log); err != nil {
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
