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

// SiringClockAuctionABI is the input ABI used to generate the binding from.
const SiringClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSiringClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// SiringClockAuctionBin is the compiled bytecode used for deploying new contracts.
const SiringClockAuctionBin = `60806040526000805460a060020a60ff02191690556004805460ff1916600117905534801561002d57600080fd5b50604051604080610e1a83398101604052805160209091015160008054600160a060020a031916331781558290829061271082111561006b57600080fd5b506002819055604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f9a20483d00000000000000000000000000000000000000000000000000000000600482015290518391600160a060020a038316916301ffc9a7916024808201926020929091908290030181600087803b1580156100f557600080fd5b505af1158015610109573d6000803e3d6000fd5b505050506040513d602081101561011f57600080fd5b5051151561012c57600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610cba806101606000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327ebe40a81146100df5780633f4ba83a1461010e578063454a2ab3146101375780635c975abb146101425780635fd8c7101461015757806376190f8f1461016c57806378bd79351461018157806383b5ff8b146101ce5780638456cb59146101f5578063878eb3681461020a5780638da5cb5b1461022257806396b5a75514610253578063c55d0f561461026b578063dd1b7a0f14610283578063f2fde38b14610298575b600080fd5b3480156100eb57600080fd5b5061010c600435602435604435606435600160a060020a03608435166102b9565b005b34801561011a57600080fd5b506101236103b1565b604080519115158252519081900360200190f35b61010c60043561042c565b34801561014e57600080fd5b50610123610478565b34801561016357600080fd5b5061010c610488565b34801561017857600080fd5b506101236104e5565b34801561018d57600080fd5b506101996004356104ee565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b3480156101da57600080fd5b506101e3610584565b60408051918252519081900360200190f35b34801561020157600080fd5b5061012361058a565b34801561021657600080fd5b5061010c60043561060a565b34801561022e57600080fd5b50610237610673565b60408051600160a060020a039092168252519081900360200190f35b34801561025f57600080fd5b5061010c600435610682565b34801561027757600080fd5b506101e36004356106cc565b34801561028f57600080fd5b506102376106fe565b3480156102a457600080fd5b5061010c600160a060020a036004351661070d565b6102c1610c60565b6fffffffffffffffffffffffffffffffff851685146102df57600080fd5b6fffffffffffffffffffffffffffffffff841684146102fd57600080fd5b67ffffffffffffffff8316831461031357600080fd5b600154600160a060020a0316331461032a57600080fd5b610334828761075f565b60a06040519081016040528083600160a060020a03168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681525090506103a986826107e7565b505050505050565b60008054600160a060020a031633146103c957600080fd5b60005460a060020a900460ff1615156103e157600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b600154600090600160a060020a0316331461044657600080fd5b50600081815260036020526040902054600160a060020a0316610469823461093b565b506104748183610a61565b5050565b60005460a060020a900460ff1681565b60015460008054600160a060020a0392831692163314806104b1575033600160a060020a038316145b15156104bc57600080fd5b604051600160a060020a03831690303180156108fc02916000818181858888f150505050505050565b60045460ff1681565b6000818152600360205260408120819081908190819061050d81610acf565b151561051857600080fd5b80546001820154600290920154600160a060020a03909116986fffffffffffffffffffffffffffffffff8084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b60008054600160a060020a031633146105a257600080fd5b60005460a060020a900460ff16156105b957600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b6000805460a060020a900460ff16151561062357600080fd5b600054600160a060020a0316331461063a57600080fd5b50600081815260036020526040902061065281610acf565b151561065d57600080fd5b8054610474908390600160a060020a0316610af0565b600054600160a060020a031681565b60008181526003602052604081209061069a82610acf565b15156106a557600080fd5b508054600160a060020a03163381146106bd57600080fd5b6106c78382610af0565b505050565b60008181526003602052604081206106e381610acf565b15156106ee57600080fd5b6106f781610b3a565b9392505050565b600154600160a060020a031681565b600054600160a060020a0316331461072457600080fd5b600160a060020a0381161561075c576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd91606480830192600092919082900301818387803b1580156107d357600080fd5b505af11580156103a9573d6000803e3d6000fd5b603c816060015167ffffffffffffffff161015151561080557600080fd5b60008281526003602090815260409182902083518154600160a060020a0390911673ffffffffffffffffffffffffffffffffffffffff1990911617815581840151600182018054858701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000081029482166fffffffffffffffffffffffffffffffff19909316831790911693909317909155606080870151600290940180546080808a015167ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190981667ffffffffffffffff1990931683171696909617909155865189815295860192909252848601929092529083015291517fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7929181900390910190a15050565b6000828152600360205260408120818080808061095786610acf565b151561096257600080fd5b61096b86610b3a565b94508488101561097a57600080fd5b8554600160a060020a0316935061099089610bca565b60008511156109e2576109a285610c17565b6040519093508386039250600160a060020a0385169083156108fc029084906000818181858888f193505050501580156109e0573d6000803e3d6000fd5b505b5060405184880390339082156108fc029083906000818181858888f19350505050158015610a14573d6000803e3d6000fd5b50604080518a815260208101879052338183015290517f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd29181900360600190a15092979650505050505050565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb91604480830192600092919082900301818387803b1580156107d357600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610af982610bca565b610b038183610a61565b6040805183815290517f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df9181900360200190a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610b805750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106f7916fffffffffffffffffffffffffffffffff80821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610c23565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610c3757869350610c55565b878703925085858402811515610c4957fe5b05915081880190508093505b505050949350505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a7230582064020b7b04e63db2c1ec70d90e7c35e82e5540d89a79fd96c64806526d8850f30029`

// DeploySiringClockAuction deploys a new Ethereum contract, binding an instance of SiringClockAuction to it.
func DeploySiringClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddr common.Address, _cut *big.Int) (common.Address, *types.Transaction, *SiringClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SiringClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SiringClockAuctionBin), backend, _nftAddr, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiringClockAuction{SiringClockAuctionCaller: SiringClockAuctionCaller{contract: contract}, SiringClockAuctionTransactor: SiringClockAuctionTransactor{contract: contract}, SiringClockAuctionFilterer: SiringClockAuctionFilterer{contract: contract}}, nil
}

// SiringClockAuction is an auto generated Go binding around an Ethereum contract.
type SiringClockAuction struct {
	SiringClockAuctionCaller     // Read-only binding to the contract
	SiringClockAuctionTransactor // Write-only binding to the contract
	SiringClockAuctionFilterer   // Log filterer for contract events
}

// SiringClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiringClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiringClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiringClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiringClockAuctionSession struct {
	Contract     *SiringClockAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SiringClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiringClockAuctionCallerSession struct {
	Contract *SiringClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SiringClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiringClockAuctionTransactorSession struct {
	Contract     *SiringClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SiringClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiringClockAuctionRaw struct {
	Contract *SiringClockAuction // Generic contract binding to access the raw methods on
}

// SiringClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiringClockAuctionCallerRaw struct {
	Contract *SiringClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SiringClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiringClockAuctionTransactorRaw struct {
	Contract *SiringClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiringClockAuction creates a new instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuction(address common.Address, backend bind.ContractBackend) (*SiringClockAuction, error) {
	contract, err := bindSiringClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuction{SiringClockAuctionCaller: SiringClockAuctionCaller{contract: contract}, SiringClockAuctionTransactor: SiringClockAuctionTransactor{contract: contract}, SiringClockAuctionFilterer: SiringClockAuctionFilterer{contract: contract}}, nil
}

// NewSiringClockAuctionCaller creates a new read-only instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*SiringClockAuctionCaller, error) {
	contract, err := bindSiringClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionCaller{contract: contract}, nil
}

// NewSiringClockAuctionTransactor creates a new write-only instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SiringClockAuctionTransactor, error) {
	contract, err := bindSiringClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionTransactor{contract: contract}, nil
}

// NewSiringClockAuctionFilterer creates a new log filterer instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SiringClockAuctionFilterer, error) {
	contract, err := bindSiringClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionFilterer{contract: contract}, nil
}

// bindSiringClockAuction binds a generic wrapper to an already deployed contract.
func bindSiringClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SiringClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiringClockAuction *SiringClockAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SiringClockAuction.Contract.SiringClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiringClockAuction *SiringClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.SiringClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiringClockAuction *SiringClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.SiringClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiringClockAuction *SiringClockAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SiringClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiringClockAuction *SiringClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiringClockAuction *SiringClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	ret := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	out := ret
	err := _SiringClockAuction.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SiringClockAuction *SiringClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SiringClockAuction.Contract.GetAuction(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SiringClockAuction.Contract.GetAuction(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SiringClockAuction.Contract.GetCurrentPrice(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SiringClockAuction.Contract.GetCurrentPrice(&_SiringClockAuction.CallOpts, _tokenId)
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCaller) IsSiringClockAuction(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "isSiringClockAuction")
	return *ret0, err
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) IsSiringClockAuction() (bool, error) {
	return _SiringClockAuction.Contract.IsSiringClockAuction(&_SiringClockAuction.CallOpts)
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCallerSession) IsSiringClockAuction() (bool, error) {
	return _SiringClockAuction.Contract.IsSiringClockAuction(&_SiringClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _SiringClockAuction.Contract.NonFungibleContract(&_SiringClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _SiringClockAuction.Contract.NonFungibleContract(&_SiringClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionSession) Owner() (common.Address, error) {
	return _SiringClockAuction.Contract.Owner(&_SiringClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SiringClockAuction *SiringClockAuctionCallerSession) Owner() (common.Address, error) {
	return _SiringClockAuction.Contract.Owner(&_SiringClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _SiringClockAuction.Contract.OwnerCut(&_SiringClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _SiringClockAuction.Contract.OwnerCut(&_SiringClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SiringClockAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Paused() (bool, error) {
	return _SiringClockAuction.Contract.Paused(&_SiringClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SiringClockAuction *SiringClockAuctionCallerSession) Paused() (bool, error) {
	return _SiringClockAuction.Contract.Paused(&_SiringClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Bid(&_SiringClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Bid(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuction(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuction(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuctionWhenPaused(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuctionWhenPaused(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CreateAuction(&_SiringClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CreateAuction(&_SiringClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Pause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Pause(&_SiringClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Pause(&_SiringClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SiringClockAuction *SiringClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.TransferOwnership(&_SiringClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.TransferOwnership(&_SiringClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Unpause(&_SiringClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Unpause(&_SiringClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.WithdrawBalance(&_SiringClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.WithdrawBalance(&_SiringClockAuction.TransactOpts)
}

// SiringClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCancelledIterator struct {
	Event *SiringClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionCancelled)
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
		it.Event = new(SiringClockAuctionAuctionCancelled)
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
func (it *SiringClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*SiringClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionCancelledIterator{contract: _SiringClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionCancelled)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// SiringClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCreatedIterator struct {
	Event *SiringClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionCreated)
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
		it.Event = new(SiringClockAuctionAuctionCreated)
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
func (it *SiringClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionCreated represents a AuctionCreated event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*SiringClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionCreatedIterator{contract: _SiringClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionCreated)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// SiringClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionSuccessfulIterator struct {
	Event *SiringClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionSuccessful)
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
		it.Event = new(SiringClockAuctionAuctionSuccessful)
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
func (it *SiringClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*SiringClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionSuccessfulIterator{contract: _SiringClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionSuccessful)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// SiringClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the SiringClockAuction contract.
type SiringClockAuctionPauseIterator struct {
	Event *SiringClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionPause)
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
		it.Event = new(SiringClockAuctionPause)
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
func (it *SiringClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionPause represents a Pause event raised by the SiringClockAuction contract.
type SiringClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*SiringClockAuctionPauseIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionPauseIterator{contract: _SiringClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionPause)
				if err := _SiringClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// SiringClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the SiringClockAuction contract.
type SiringClockAuctionUnpauseIterator struct {
	Event *SiringClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionUnpause)
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
		it.Event = new(SiringClockAuctionUnpause)
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
func (it *SiringClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionUnpause represents a Unpause event raised by the SiringClockAuction contract.
type SiringClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*SiringClockAuctionUnpauseIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionUnpauseIterator{contract: _SiringClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionUnpause)
				if err := _SiringClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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
