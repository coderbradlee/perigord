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

// SaleClockAuctionABI is the input ABI used to generate the binding from.
const SaleClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastGen0SalePrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSaleClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0SaleCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"averageGen0SalePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// SaleClockAuctionBin is the compiled bytecode used for deploying new contracts.
const SaleClockAuctionBin = `60806040526000805460a060020a60ff02191690556004805460ff1916600117905534801561002d57600080fd5b50604051604080610ee883398101604052805160209091015160008054600160a060020a031916331781558290829061271082111561006b57600080fd5b506002819055604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f9a20483d00000000000000000000000000000000000000000000000000000000600482015290518391600160a060020a038316916301ffc9a7916024808201926020929091908290030181600087803b1580156100f557600080fd5b505af1158015610109573d6000803e3d6000fd5b505050506040513d602081101561011f57600080fd5b5051151561012c57600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610d88806101606000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327ebe40a81146101005780633f4ba83a1461012f578063454a2ab314610158578063484eccb4146101635780635c975abb1461018d5780635fd8c710146101a257806378bd7935146101b757806383b5ff8b146102045780638456cb591461021957806385b861881461022e578063878eb368146102435780638a98a9cc1461025b5780638da5cb5b1461027057806396b5a755146102a1578063c55d0f56146102b9578063dd1b7a0f146102d1578063eac9d94c146102e6578063f2fde38b146102fb575b600080fd5b34801561010c57600080fd5b5061012d600435602435604435606435600160a060020a036084351661031c565b005b34801561013b57600080fd5b50610144610414565b604080519115158252519081900360200190f35b61012d60043561048f565b34801561016f57600080fd5b5061017b6004356104f9565b60408051918252519081900360200190f35b34801561019957600080fd5b5061014461050d565b3480156101ae57600080fd5b5061012d61051d565b3480156101c357600080fd5b506101cf60043561057a565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561021057600080fd5b5061017b610610565b34801561022557600080fd5b50610144610616565b34801561023a57600080fd5b50610144610696565b34801561024f57600080fd5b5061012d60043561069f565b34801561026757600080fd5b5061017b61070c565b34801561027c57600080fd5b50610285610712565b60408051600160a060020a039092168252519081900360200190f35b3480156102ad57600080fd5b5061012d600435610721565b3480156102c557600080fd5b5061017b600435610766565b3480156102dd57600080fd5b50610285610798565b3480156102f257600080fd5b5061017b6107a7565b34801561030757600080fd5b5061012d600160a060020a03600435166107db565b610324610d2e565b6fffffffffffffffffffffffffffffffff8516851461034257600080fd5b6fffffffffffffffffffffffffffffffff8416841461036057600080fd5b67ffffffffffffffff8316831461037657600080fd5b600154600160a060020a0316331461038d57600080fd5b610397828761082d565b60a06040519081016040528083600160a060020a03168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061040c86826108b5565b505050505050565b60008054600160a060020a0316331461042c57600080fd5b60005460a060020a900460ff16151561044457600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b600081815260036020526040812054600160a060020a0316906104b28334610a09565b90506104be3384610b2f565b600154600160a060020a03838116911614156104f45760058054829160069106600581106104e857fe5b01556005805460010190555b505050565b6006816005811061050657fe5b0154905081565b60005460a060020a900460ff1681565b60015460008054600160a060020a039283169216331480610546575033600160a060020a038316145b151561055157600080fd5b604051600160a060020a03831690303180156108fc02916000818181858888f150505050505050565b6000818152600360205260408120819081908190819061059981610b9d565b15156105a457600080fd5b80546001820154600290920154600160a060020a03909116986fffffffffffffffffffffffffffffffff8084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b60008054600160a060020a0316331461062e57600080fd5b60005460a060020a900460ff161561064557600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b60045460ff1681565b6000805460a060020a900460ff1615156106b857600080fd5b600054600160a060020a031633146106cf57600080fd5b5060008181526003602052604090206106e781610b9d565b15156106f257600080fd5b8054610708908390600160a060020a0316610bbe565b5050565b60055481565b600054600160a060020a031681565b60008181526003602052604081209061073982610b9d565b151561074457600080fd5b508054600160a060020a031633811461075c57600080fd5b6104f48382610bbe565b600081815260036020526040812061077d81610b9d565b151561078857600080fd5b61079181610c08565b9392505050565b600154600160a060020a031681565b600080805b60058110156107d157600681600581106107c257fe5b015491909101906001016107ac565b5060059004919050565b600054600160a060020a031633146107f257600080fd5b600160a060020a0381161561082a576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd91606480830192600092919082900301818387803b1580156108a157600080fd5b505af115801561040c573d6000803e3d6000fd5b603c816060015167ffffffffffffffff16101515156108d357600080fd5b60008281526003602090815260409182902083518154600160a060020a0390911673ffffffffffffffffffffffffffffffffffffffff1990911617815581840151600182018054858701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000081029482166fffffffffffffffffffffffffffffffff19909316831790911693909317909155606080870151600290940180546080808a015167ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190981667ffffffffffffffff1990931683171696909617909155865189815295860192909252848601929092529083015291517fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7929181900390910190a15050565b60008281526003602052604081208180808080610a2586610b9d565b1515610a3057600080fd5b610a3986610c08565b945084881015610a4857600080fd5b8554600160a060020a03169350610a5e89610c98565b6000851115610ab057610a7085610ce5565b6040519093508386039250600160a060020a0385169083156108fc029084906000818181858888f19350505050158015610aae573d6000803e3d6000fd5b505b5060405184880390339082156108fc029083906000818181858888f19350505050158015610ae2573d6000803e3d6000fd5b50604080518a815260208101879052338183015290517f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd29181900360600190a15092979650505050505050565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb91604480830192600092919082900301818387803b1580156108a157600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610bc782610c98565b610bd18183610b2f565b6040805183815290517f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df9181900360200190a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610c4e5750600282015468010000000000000000900467ffffffffffffffff1642035b60018301546002840154610791916fffffffffffffffffffffffffffffffff80821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610cf1565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610d0557869350610d23565b878703925085858402811515610d1757fe5b05915081880190508093505b505050949350505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a723058209ff93b60ab4b1b61bb9e14d9d4dce4565fa620ddd0404b40f9245cc52db6f8b80029`

// DeploySaleClockAuction deploys a new Ethereum contract, binding an instance of SaleClockAuction to it.
func DeploySaleClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddr common.Address, _cut *big.Int) (common.Address, *types.Transaction, *SaleClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaleClockAuctionBin), backend, _nftAddr, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaleClockAuction{SaleClockAuctionCaller: SaleClockAuctionCaller{contract: contract}, SaleClockAuctionTransactor: SaleClockAuctionTransactor{contract: contract}, SaleClockAuctionFilterer: SaleClockAuctionFilterer{contract: contract}}, nil
}

// SaleClockAuction is an auto generated Go binding around an Ethereum contract.
type SaleClockAuction struct {
	SaleClockAuctionCaller     // Read-only binding to the contract
	SaleClockAuctionTransactor // Write-only binding to the contract
	SaleClockAuctionFilterer   // Log filterer for contract events
}

// SaleClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleClockAuctionSession struct {
	Contract     *SaleClockAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleClockAuctionCallerSession struct {
	Contract *SaleClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SaleClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleClockAuctionTransactorSession struct {
	Contract     *SaleClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SaleClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleClockAuctionRaw struct {
	Contract *SaleClockAuction // Generic contract binding to access the raw methods on
}

// SaleClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleClockAuctionCallerRaw struct {
	Contract *SaleClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SaleClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleClockAuctionTransactorRaw struct {
	Contract *SaleClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaleClockAuction creates a new instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuction(address common.Address, backend bind.ContractBackend) (*SaleClockAuction, error) {
	contract, err := bindSaleClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuction{SaleClockAuctionCaller: SaleClockAuctionCaller{contract: contract}, SaleClockAuctionTransactor: SaleClockAuctionTransactor{contract: contract}, SaleClockAuctionFilterer: SaleClockAuctionFilterer{contract: contract}}, nil
}

// NewSaleClockAuctionCaller creates a new read-only instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*SaleClockAuctionCaller, error) {
	contract, err := bindSaleClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionCaller{contract: contract}, nil
}

// NewSaleClockAuctionTransactor creates a new write-only instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleClockAuctionTransactor, error) {
	contract, err := bindSaleClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionTransactor{contract: contract}, nil
}

// NewSaleClockAuctionFilterer creates a new log filterer instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleClockAuctionFilterer, error) {
	contract, err := bindSaleClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionFilterer{contract: contract}, nil
}

// bindSaleClockAuction binds a generic wrapper to an already deployed contract.
func bindSaleClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleClockAuction *SaleClockAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleClockAuction.Contract.SaleClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleClockAuction *SaleClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.SaleClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleClockAuction *SaleClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.SaleClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleClockAuction *SaleClockAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleClockAuction *SaleClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleClockAuction *SaleClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.contract.Transact(opts, method, params...)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) AverageGen0SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "averageGen0SalePrice")
	return *ret0, err
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) AverageGen0SalePrice() (*big.Int, error) {
	return _SaleClockAuction.Contract.AverageGen0SalePrice(&_SaleClockAuction.CallOpts)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) AverageGen0SalePrice() (*big.Int, error) {
	return _SaleClockAuction.Contract.AverageGen0SalePrice(&_SaleClockAuction.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) Gen0SaleCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "gen0SaleCount")
	return *ret0, err
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) Gen0SaleCount() (*big.Int, error) {
	return _SaleClockAuction.Contract.Gen0SaleCount(&_SaleClockAuction.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Gen0SaleCount() (*big.Int, error) {
	return _SaleClockAuction.Contract.Gen0SaleCount(&_SaleClockAuction.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
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
	err := _SaleClockAuction.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SaleClockAuction *SaleClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SaleClockAuction.Contract.GetAuction(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SaleClockAuction.Contract.GetAuction(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.GetCurrentPrice(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.GetCurrentPrice(&_SaleClockAuction.CallOpts, _tokenId)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCaller) IsSaleClockAuction(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "isSaleClockAuction")
	return *ret0, err
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) IsSaleClockAuction() (bool, error) {
	return _SaleClockAuction.Contract.IsSaleClockAuction(&_SaleClockAuction.CallOpts)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCallerSession) IsSaleClockAuction() (bool, error) {
	return _SaleClockAuction.Contract.IsSaleClockAuction(&_SaleClockAuction.CallOpts)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) LastGen0SalePrices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "lastGen0SalePrices", arg0)
	return *ret0, err
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.LastGen0SalePrices(&_SaleClockAuction.CallOpts, arg0)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.LastGen0SalePrices(&_SaleClockAuction.CallOpts, arg0)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _SaleClockAuction.Contract.NonFungibleContract(&_SaleClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _SaleClockAuction.Contract.NonFungibleContract(&_SaleClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionSession) Owner() (common.Address, error) {
	return _SaleClockAuction.Contract.Owner(&_SaleClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Owner() (common.Address, error) {
	return _SaleClockAuction.Contract.Owner(&_SaleClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _SaleClockAuction.Contract.OwnerCut(&_SaleClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _SaleClockAuction.Contract.OwnerCut(&_SaleClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SaleClockAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Paused() (bool, error) {
	return _SaleClockAuction.Contract.Paused(&_SaleClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Paused() (bool, error) {
	return _SaleClockAuction.Contract.Paused(&_SaleClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Bid(&_SaleClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Bid(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuction(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuction(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuctionWhenPaused(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuctionWhenPaused(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CreateAuction(&_SaleClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CreateAuction(&_SaleClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Pause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Pause(&_SaleClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Pause(&_SaleClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleClockAuction *SaleClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.TransferOwnership(&_SaleClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.TransferOwnership(&_SaleClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Unpause(&_SaleClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Unpause(&_SaleClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.WithdrawBalance(&_SaleClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.WithdrawBalance(&_SaleClockAuction.TransactOpts)
}

// SaleClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCancelledIterator struct {
	Event *SaleClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionCancelled)
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
		it.Event = new(SaleClockAuctionAuctionCancelled)
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
func (it *SaleClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*SaleClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionCancelledIterator{contract: _SaleClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionCancelled)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// SaleClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCreatedIterator struct {
	Event *SaleClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionCreated)
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
		it.Event = new(SaleClockAuctionAuctionCreated)
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
func (it *SaleClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionCreated represents a AuctionCreated event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*SaleClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionCreatedIterator{contract: _SaleClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionCreated)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// SaleClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionSuccessfulIterator struct {
	Event *SaleClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionSuccessful)
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
		it.Event = new(SaleClockAuctionAuctionSuccessful)
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
func (it *SaleClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*SaleClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionSuccessfulIterator{contract: _SaleClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionSuccessful)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// SaleClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the SaleClockAuction contract.
type SaleClockAuctionPauseIterator struct {
	Event *SaleClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionPause)
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
		it.Event = new(SaleClockAuctionPause)
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
func (it *SaleClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionPause represents a Pause event raised by the SaleClockAuction contract.
type SaleClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*SaleClockAuctionPauseIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionPauseIterator{contract: _SaleClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionPause)
				if err := _SaleClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// SaleClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the SaleClockAuction contract.
type SaleClockAuctionUnpauseIterator struct {
	Event *SaleClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionUnpause)
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
		it.Event = new(SaleClockAuctionUnpause)
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
func (it *SaleClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionUnpause represents a Unpause event raised by the SaleClockAuction contract.
type SaleClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*SaleClockAuctionUnpauseIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionUnpauseIterator{contract: _SaleClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionUnpause)
				if err := _SaleClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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
