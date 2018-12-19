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

// ClockAuctionABI is the input ABI used to generate the binding from.
const ClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddress\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ClockAuctionBin is the compiled bytecode used for deploying new contracts.
const ClockAuctionBin = `60806040526000805460a060020a60ff021916905534801561002057600080fd5b50604051604080610e8483398101604052805160209091015160008054600160a060020a0319163317815561271082111561005a57600080fd5b506002819055604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f9a20483d00000000000000000000000000000000000000000000000000000000600482015290518391600160a060020a038316916301ffc9a7916024808201926020929091908290030181600087803b1580156100e457600080fd5b505af11580156100f8573d6000803e3d6000fd5b505050506040513d602081101561010e57600080fd5b5051151561011b57600080fd5b60018054600160a060020a03909216600160a060020a03199092169190911790555050610d378061014d6000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327ebe40a81146100d45780633f4ba83a14610103578063454a2ab31461012c5780635c975abb146101375780635fd8c7101461014c57806378bd79351461016157806383b5ff8b146101ae5780638456cb59146101d5578063878eb368146101ea5780638da5cb5b1461020257806396b5a75514610233578063c55d0f561461024b578063dd1b7a0f14610263578063f2fde38b14610278575b600080fd5b3480156100e057600080fd5b50610101600435602435604435606435600160a060020a0360843516610299565b005b34801561010f57600080fd5b506101186103a6565b604080519115158252519081900360200190f35b610101600435610421565b34801561014357600080fd5b50610118610450565b34801561015857600080fd5b50610101610460565b34801561016d57600080fd5b506101796004356104bd565b60408051600160a060020a03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b3480156101ba57600080fd5b506101c3610553565b60408051918252519081900360200190f35b3480156101e157600080fd5b50610118610559565b3480156101f657600080fd5b506101016004356105d9565b34801561020e57600080fd5b50610217610646565b60408051600160a060020a039092168252519081900360200190f35b34801561023f57600080fd5b50610101600435610655565b34801561025757600080fd5b506101c360043561069f565b34801561026f57600080fd5b506102176106d1565b34801561028457600080fd5b50610101600160a060020a03600435166106e0565b6102a1610cdd565b60005460a060020a900460ff16156102b857600080fd5b6fffffffffffffffffffffffffffffffff851685146102d657600080fd5b6fffffffffffffffffffffffffffffffff841684146102f457600080fd5b67ffffffffffffffff8316831461030a57600080fd5b6103143387610733565b151561031f57600080fd5b61032933876107dc565b60a06040519081016040528083600160a060020a03168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061039e8682610864565b505050505050565b60008054600160a060020a031633146103be57600080fd5b60005460a060020a900460ff1615156103d657600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a150600190565b60005460a060020a900460ff161561043857600080fd5b61044281346109b8565b5061044d3382610ade565b50565b60005460a060020a900460ff1681565b60015460008054600160a060020a039283169216331480610489575033600160a060020a038316145b151561049457600080fd5b604051600160a060020a03831690303180156108fc02916000818181858888f150505050505050565b600081815260036020526040812081908190819081906104dc81610b4c565b15156104e757600080fd5b80546001820154600290920154600160a060020a03909116986fffffffffffffffffffffffffffffffff8084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b60008054600160a060020a0316331461057157600080fd5b60005460a060020a900460ff161561058857600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a150600190565b6000805460a060020a900460ff1615156105f257600080fd5b600054600160a060020a0316331461060957600080fd5b50600081815260036020526040902061062181610b4c565b151561062c57600080fd5b8054610642908390600160a060020a0316610b6d565b5050565b600054600160a060020a031681565b60008181526003602052604081209061066d82610b4c565b151561067857600080fd5b508054600160a060020a031633811461069057600080fd5b61069a8382610b6d565b505050565b60008181526003602052604081206106b681610b4c565b15156106c157600080fd5b6106ca81610bb7565b9392505050565b600154600160a060020a031681565b600054600160a060020a031633146106f757600080fd5b600160a060020a0381161561044d5760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b600154604080517f6352211e000000000000000000000000000000000000000000000000000000008152600481018490529051600092600160a060020a0380871693911691636352211e9160248082019260209290919082900301818887803b15801561079f57600080fd5b505af11580156107b3573d6000803e3d6000fd5b505050506040513d60208110156107c957600080fd5b5051600160a060020a0316149392505050565b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd91606480830192600092919082900301818387803b15801561085057600080fd5b505af115801561039e573d6000803e3d6000fd5b603c816060015167ffffffffffffffff161015151561088257600080fd5b60008281526003602090815260409182902083518154600160a060020a0390911673ffffffffffffffffffffffffffffffffffffffff1990911617815581840151600182018054858701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000081029482166fffffffffffffffffffffffffffffffff19909316831790911693909317909155606080870151600290940180546080808a015167ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190981667ffffffffffffffff1990931683171696909617909155865189815295860192909252848601929092529083015291517fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7929181900390910190a15050565b600082815260036020526040812081808080806109d486610b4c565b15156109df57600080fd5b6109e886610bb7565b9450848810156109f757600080fd5b8554600160a060020a03169350610a0d89610c47565b6000851115610a5f57610a1f85610c94565b6040519093508386039250600160a060020a0385169083156108fc029084906000818181858888f19350505050158015610a5d573d6000803e3d6000fd5b505b5060405184880390339082156108fc029083906000818181858888f19350505050158015610a91573d6000803e3d6000fd5b50604080518a815260208101879052338183015290517f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd29181900360600190a15092979650505050505050565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb91604480830192600092919082900301818387803b15801561085057600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610b7682610c47565b610b808183610ade565b6040805183815290517f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df9181900360200190a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610bfd5750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106ca916fffffffffffffffffffffffffffffffff80821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610ca0565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610cb457869350610cd2565b878703925085858402811515610cc657fe5b05915081880190508093505b505050949350505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152905600a165627a7a72305820a1acc42a175b8a9744a4b7a26479be770e68f105b045d9ee78492e6c84801c750029`

// DeployClockAuction deploys a new Ethereum contract, binding an instance of ClockAuction to it.
func DeployClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address, _cut *big.Int) (common.Address, *types.Transaction, *ClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClockAuctionBin), backend, _nftAddress, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClockAuction{ClockAuctionCaller: ClockAuctionCaller{contract: contract}, ClockAuctionTransactor: ClockAuctionTransactor{contract: contract}, ClockAuctionFilterer: ClockAuctionFilterer{contract: contract}}, nil
}

// ClockAuction is an auto generated Go binding around an Ethereum contract.
type ClockAuction struct {
	ClockAuctionCaller     // Read-only binding to the contract
	ClockAuctionTransactor // Write-only binding to the contract
	ClockAuctionFilterer   // Log filterer for contract events
}

// ClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockAuctionSession struct {
	Contract     *ClockAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockAuctionCallerSession struct {
	Contract *ClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockAuctionTransactorSession struct {
	Contract     *ClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockAuctionRaw struct {
	Contract *ClockAuction // Generic contract binding to access the raw methods on
}

// ClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockAuctionCallerRaw struct {
	Contract *ClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// ClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockAuctionTransactorRaw struct {
	Contract *ClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockAuction creates a new instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuction(address common.Address, backend bind.ContractBackend) (*ClockAuction, error) {
	contract, err := bindClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClockAuction{ClockAuctionCaller: ClockAuctionCaller{contract: contract}, ClockAuctionTransactor: ClockAuctionTransactor{contract: contract}, ClockAuctionFilterer: ClockAuctionFilterer{contract: contract}}, nil
}

// NewClockAuctionCaller creates a new read-only instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*ClockAuctionCaller, error) {
	contract, err := bindClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionCaller{contract: contract}, nil
}

// NewClockAuctionTransactor creates a new write-only instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockAuctionTransactor, error) {
	contract, err := bindClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionTransactor{contract: contract}, nil
}

// NewClockAuctionFilterer creates a new log filterer instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockAuctionFilterer, error) {
	contract, err := bindClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionFilterer{contract: contract}, nil
}

// bindClockAuction binds a generic wrapper to an already deployed contract.
func bindClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuction *ClockAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuction.Contract.ClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuction *ClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.Contract.ClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuction *ClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuction.Contract.ClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuction *ClockAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuction *ClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuction *ClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_ClockAuction *ClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
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
	err := _ClockAuction.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_ClockAuction *ClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _ClockAuction.Contract.GetAuction(&_ClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_ClockAuction *ClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _ClockAuction.Contract.GetAuction(&_ClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_ClockAuction *ClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_ClockAuction *ClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _ClockAuction.Contract.GetCurrentPrice(&_ClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_ClockAuction *ClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _ClockAuction.Contract.GetCurrentPrice(&_ClockAuction.CallOpts, _tokenId)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuction *ClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuction *ClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuction.Contract.NonFungibleContract(&_ClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_ClockAuction *ClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuction.Contract.NonFungibleContract(&_ClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClockAuction *ClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClockAuction *ClockAuctionSession) Owner() (common.Address, error) {
	return _ClockAuction.Contract.Owner(&_ClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClockAuction *ClockAuctionCallerSession) Owner() (common.Address, error) {
	return _ClockAuction.Contract.Owner(&_ClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuction *ClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuction *ClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _ClockAuction.Contract.OwnerCut(&_ClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_ClockAuction *ClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _ClockAuction.Contract.OwnerCut(&_ClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClockAuction *ClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClockAuction.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClockAuction *ClockAuctionSession) Paused() (bool, error) {
	return _ClockAuction.Contract.Paused(&_ClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClockAuction *ClockAuctionCallerSession) Paused() (bool, error) {
	return _ClockAuction.Contract.Paused(&_ClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.Bid(&_ClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.Bid(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuction(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuction(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuctionWhenPaused(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuctionWhenPaused(&_ClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_ClockAuction *ClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_ClockAuction *ClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.CreateAuction(&_ClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.CreateAuction(&_ClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionSession) Pause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Pause(&_ClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Pause(&_ClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClockAuction *ClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClockAuction *ClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.TransferOwnership(&_ClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClockAuction *ClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.TransferOwnership(&_ClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Unpause(&_ClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Unpause(&_ClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _ClockAuction.Contract.WithdrawBalance(&_ClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _ClockAuction.Contract.WithdrawBalance(&_ClockAuction.TransactOpts)
}

// ClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the ClockAuction contract.
type ClockAuctionAuctionCancelledIterator struct {
	Event *ClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionCancelled)
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
		it.Event = new(ClockAuctionAuctionCancelled)
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
func (it *ClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the ClockAuction contract.
type ClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionCancelledIterator{contract: _ClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: e AuctionCancelled(tokenId uint256)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionCancelled)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the ClockAuction contract.
type ClockAuctionAuctionCreatedIterator struct {
	Event *ClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionCreated)
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
		it.Event = new(ClockAuctionAuctionCreated)
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
func (it *ClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionCreated represents a AuctionCreated event raised by the ClockAuction contract.
type ClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionCreatedIterator{contract: _ClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: e AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionCreated)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the ClockAuction contract.
type ClockAuctionAuctionSuccessfulIterator struct {
	Event *ClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionSuccessful)
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
		it.Event = new(ClockAuctionAuctionSuccessful)
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
func (it *ClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the ClockAuction contract.
type ClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionSuccessfulIterator{contract: _ClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: e AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionSuccessful)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ClockAuction contract.
type ClockAuctionPauseIterator struct {
	Event *ClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *ClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionPause)
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
		it.Event = new(ClockAuctionPause)
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
func (it *ClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionPause represents a Pause event raised by the ClockAuction contract.
type ClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_ClockAuction *ClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*ClockAuctionPauseIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionPauseIterator{contract: _ClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_ClockAuction *ClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionPause)
				if err := _ClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ClockAuction contract.
type ClockAuctionUnpauseIterator struct {
	Event *ClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *ClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionUnpause)
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
		it.Event = new(ClockAuctionUnpause)
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
func (it *ClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionUnpause represents a Unpause event raised by the ClockAuction contract.
type ClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_ClockAuction *ClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*ClockAuctionUnpauseIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionUnpauseIterator{contract: _ClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_ClockAuction *ClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionUnpause)
				if err := _ClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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
