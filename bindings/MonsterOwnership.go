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

// MonsterOwnershipABI is the input ABI used to generate the binding from.
const MonsterOwnershipABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsterIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsterIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"MonsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// MonsterOwnershipBin is the compiled bytecode used for deploying new contracts.
const MonsterOwnershipBin = `6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620000a1565b50600f6005553480156200009a57600080fd5b506200016b565b600283019183908215620001325791602002820160005b83821115620000fe57835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000b8565b8015620001305782816101000a81549063ffffffff0219169055600401602081600301049283019260010302620000fe565b505b506200014092915062000144565b5090565b6200016891905b808211156200014057805463ffffffff191681556001016200014b565b90565b611202806200017b6000396000f30060806040526004361061017f5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146101845780630519ce79146101cf5780630560ff441461020057806306fdde0314610299578063095ea7b3146102ae5780630a0f8168146102d457806318160ddd146102e957806321717ebf1461031057806323b872dd1461032557806327d7874c1461034f5780632ba73c15146103705780633f4ba83a1461039157806346116e6f146103a65780634e0a3379146103be5780635663896e146103df5780635c975abb146103f75780636352211e1461040c57806370a08231146104245780637a7d4937146104455780637d55aeea1461045a5780638456cb59146104725780638462151c1461048757806395d89b41146104f85780639d6fac6f1461050d578063a9059cbb1461053e578063ad4b558c14610562578063b047fb501461057a578063bc4006f51461058f578063e17b25af146105a4578063e6cbe351146105c5575b600080fd5b34801561019057600080fd5b506101bb7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166105da565b604080519115158252519081900360200190f35b3480156101db57600080fd5b506101e461086d565b60408051600160a060020a039092168252519081900360200190f35b34801561020c57600080fd5b5061022460048035906024803590810191013561087c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561025e578181015183820152602001610246565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102a557600080fd5b5061022461097f565b3480156102ba57600080fd5b506102d2600160a060020a03600435166024356109b6565b005b3480156102e057600080fd5b506101e4610a38565b3480156102f557600080fd5b506102fe610a47565b60408051918252519081900360200190f35b34801561031c57600080fd5b506101e4610a51565b34801561033157600080fd5b506102d2600160a060020a0360043581169060243516604435610a60565b34801561035b57600080fd5b506102d2600160a060020a0360043516610adc565b34801561037c57600080fd5b506102d2600160a060020a0360043516610b2a565b34801561039d57600080fd5b506102d2610b78565b3480156103b257600080fd5b506101e4600435610bc7565b3480156103ca57600080fd5b506102d2600160a060020a0360043516610be2565b3480156103eb57600080fd5b506102d2600435610c30565b34801561040357600080fd5b506101bb610c8c565b34801561041857600080fd5b506101e4600435610c9c565b34801561043057600080fd5b506102fe600160a060020a0360043516610cc0565b34801561045157600080fd5b506102fe610cdb565b34801561046657600080fd5b506101e4600435610ce1565b34801561047e57600080fd5b506102d2610cfc565b34801561049357600080fd5b506104a8600160a060020a0360043516610d7c565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156104e45781810151838201526020016104cc565b505050509050019250505060405180910390f35b34801561050457600080fd5b50610224610e4e565b34801561051957600080fd5b50610525600435610e85565b6040805163ffffffff9092168252519081900360200190f35b34801561054a57600080fd5b506102d2600160a060020a0360043516602435610eb2565b34801561056e57600080fd5b506101e4600435610f4e565b34801561058657600080fd5b506101e4610f69565b34801561059b57600080fd5b506101e4610f78565b3480156105b057600080fd5b506102d2600160a060020a0360043516610f87565b3480156105d157600080fd5b506101e4610fc0565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19838116911614806108655750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b60606108866111b7565b600d54600090600160a060020a031615156108a057600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b15801561092757600080fd5b505af115801561093b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a081101561096057600080fd5b50608081015190925090506109758282610fcf565b9695505050505050565b60408051808201909152600781527f6d6f6e7374657200000000000000000000000000000000000000000000000000602082015281565b60025460a060020a900460ff16156109cd57600080fd5b6109d73382611023565b15156109e257600080fd5b6109ec8183611043565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b6006546000190190565b600c54600160a060020a031681565b60025460a060020a900460ff1615610a7757600080fd5b600160a060020a0382161515610a8c57600080fd5b600160a060020a038216301415610aa257600080fd5b610aac3382611071565b1515610ab757600080fd5b610ac18382611023565b1515610acc57600080fd5b610ad7838383611091565b505050565b600054600160a060020a03163314610af357600080fd5b600160a060020a0381161515610b0857600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610b4157600080fd5b600160a060020a0381161515610b5657600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610b8f57600080fd5b60025460a060020a900460ff161515610ba757600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600054600160a060020a03163314610bf957600080fd5b600160a060020a0381161515610c0e57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a0316331480610c535750600054600160a060020a031633145b80610c685750600154600160a060020a031633145b1515610c7357600080fd5b60035463ffffffff168110610c8757600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a031680151561086857600080fd5b600160a060020a031660009081526008602052604090205490565b60055481565b600960205260009081526040902054600160a060020a031681565b600254600160a060020a0316331480610d1f5750600054600160a060020a031633145b80610d345750600154600160a060020a031633145b1515610d3f57600080fd5b60025460a060020a900460ff1615610d5657600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b6060600060606000806000610d9087610cc0565b9450841515610daf576040805160008152602081019091529550610e44565b84604051908082528060200260200182016040528015610dd9578160200160208202803883390190505b509350610de4610a47565b925060009150600190505b828111610e4057600081815260076020526040902054600160a060020a0388811691161415610e3857808483815181101515610e2757fe5b602090810290910101526001909101905b600101610def565b8395505b5050505050919050565b60408051808201909152600281527f4d52000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110610e9257fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b60025460a060020a900460ff1615610ec957600080fd5b600160a060020a0382161515610ede57600080fd5b600160a060020a038216301415610ef457600080fd5b600b54600160a060020a0383811691161415610f0f57600080fd5b600c54600160a060020a0383811691161415610f2a57600080fd5b610f343382611023565b1515610f3f57600080fd5b610f4a338383611091565b5050565b600760205260009081526040902054600160a060020a031681565b600254600160a060020a031681565b600d54600160a060020a031681565b600054600160a060020a03163314610f9e57600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b606080600080846040519080825280601f01601f191660200182016040528015611003578160200160208202803883390190505b5092505060208201905084611019828287611173565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561112457600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b60005b60208210611198578251845260209384019390920191601f1990910190611176565b50905182516020929092036101000a6000190180199091169116179052565b60806040519081016040528060049060208202803883395091929150505600a165627a7a72305820fb1222c117964c614b82db0512c59d7978b3a3f60047ec32f43f05d9a75931cf0029`

// DeployMonsterOwnership deploys a new Ethereum contract, binding an instance of MonsterOwnership to it.
func DeployMonsterOwnership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonsterOwnership, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterOwnershipABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonsterOwnershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonsterOwnership{MonsterOwnershipCaller: MonsterOwnershipCaller{contract: contract}, MonsterOwnershipTransactor: MonsterOwnershipTransactor{contract: contract}, MonsterOwnershipFilterer: MonsterOwnershipFilterer{contract: contract}}, nil
}

// MonsterOwnership is an auto generated Go binding around an Ethereum contract.
type MonsterOwnership struct {
	MonsterOwnershipCaller     // Read-only binding to the contract
	MonsterOwnershipTransactor // Write-only binding to the contract
	MonsterOwnershipFilterer   // Log filterer for contract events
}

// MonsterOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonsterOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonsterOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonsterOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonsterOwnershipSession struct {
	Contract     *MonsterOwnership // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MonsterOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonsterOwnershipCallerSession struct {
	Contract *MonsterOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MonsterOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonsterOwnershipTransactorSession struct {
	Contract     *MonsterOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MonsterOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonsterOwnershipRaw struct {
	Contract *MonsterOwnership // Generic contract binding to access the raw methods on
}

// MonsterOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonsterOwnershipCallerRaw struct {
	Contract *MonsterOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// MonsterOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonsterOwnershipTransactorRaw struct {
	Contract *MonsterOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonsterOwnership creates a new instance of MonsterOwnership, bound to a specific deployed contract.
func NewMonsterOwnership(address common.Address, backend bind.ContractBackend) (*MonsterOwnership, error) {
	contract, err := bindMonsterOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonsterOwnership{MonsterOwnershipCaller: MonsterOwnershipCaller{contract: contract}, MonsterOwnershipTransactor: MonsterOwnershipTransactor{contract: contract}, MonsterOwnershipFilterer: MonsterOwnershipFilterer{contract: contract}}, nil
}

// NewMonsterOwnershipCaller creates a new read-only instance of MonsterOwnership, bound to a specific deployed contract.
func NewMonsterOwnershipCaller(address common.Address, caller bind.ContractCaller) (*MonsterOwnershipCaller, error) {
	contract, err := bindMonsterOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipCaller{contract: contract}, nil
}

// NewMonsterOwnershipTransactor creates a new write-only instance of MonsterOwnership, bound to a specific deployed contract.
func NewMonsterOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*MonsterOwnershipTransactor, error) {
	contract, err := bindMonsterOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipTransactor{contract: contract}, nil
}

// NewMonsterOwnershipFilterer creates a new log filterer instance of MonsterOwnership, bound to a specific deployed contract.
func NewMonsterOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*MonsterOwnershipFilterer, error) {
	contract, err := bindMonsterOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipFilterer{contract: contract}, nil
}

// bindMonsterOwnership binds a generic wrapper to an already deployed contract.
func bindMonsterOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterOwnership *MonsterOwnershipRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterOwnership.Contract.MonsterOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterOwnership *MonsterOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.MonsterOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterOwnership *MonsterOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.MonsterOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterOwnership *MonsterOwnershipCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterOwnership *MonsterOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterOwnership *MonsterOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_MonsterOwnership *MonsterOwnershipCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_MonsterOwnership *MonsterOwnershipSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MonsterOwnership.Contract.BalanceOf(&_MonsterOwnership.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_MonsterOwnership *MonsterOwnershipCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MonsterOwnership.Contract.BalanceOf(&_MonsterOwnership.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) CeoAddress() (common.Address, error) {
	return _MonsterOwnership.Contract.CeoAddress(&_MonsterOwnership.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) CeoAddress() (common.Address, error) {
	return _MonsterOwnership.Contract.CeoAddress(&_MonsterOwnership.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) CfoAddress() (common.Address, error) {
	return _MonsterOwnership.Contract.CfoAddress(&_MonsterOwnership.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) CfoAddress() (common.Address, error) {
	return _MonsterOwnership.Contract.CfoAddress(&_MonsterOwnership.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) CooAddress() (common.Address, error) {
	return _MonsterOwnership.Contract.CooAddress(&_MonsterOwnership.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) CooAddress() (common.Address, error) {
	return _MonsterOwnership.Contract.CooAddress(&_MonsterOwnership.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterOwnership *MonsterOwnershipCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterOwnership *MonsterOwnershipSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _MonsterOwnership.Contract.Cooldowns(&_MonsterOwnership.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterOwnership *MonsterOwnershipCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _MonsterOwnership.Contract.Cooldowns(&_MonsterOwnership.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) Erc721Metadata() (common.Address, error) {
	return _MonsterOwnership.Contract.Erc721Metadata(&_MonsterOwnership.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) Erc721Metadata() (common.Address, error) {
	return _MonsterOwnership.Contract.Erc721Metadata(&_MonsterOwnership.CallOpts)
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) MonsterIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "monsterIndexToApproved", arg0)
	return *ret0, err
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) MonsterIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.MonsterIndexToApproved(&_MonsterOwnership.CallOpts, arg0)
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) MonsterIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.MonsterIndexToApproved(&_MonsterOwnership.CallOpts, arg0)
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) MonsterIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "monsterIndexToOwner", arg0)
	return *ret0, err
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) MonsterIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.MonsterIndexToOwner(&_MonsterOwnership.CallOpts, arg0)
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) MonsterIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.MonsterIndexToOwner(&_MonsterOwnership.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonsterOwnership *MonsterOwnershipCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonsterOwnership *MonsterOwnershipSession) Name() (string, error) {
	return _MonsterOwnership.Contract.Name(&_MonsterOwnership.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonsterOwnership *MonsterOwnershipCallerSession) Name() (string, error) {
	return _MonsterOwnership.Contract.Name(&_MonsterOwnership.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_MonsterOwnership *MonsterOwnershipCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_MonsterOwnership *MonsterOwnershipSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.OwnerOf(&_MonsterOwnership.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.OwnerOf(&_MonsterOwnership.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterOwnership *MonsterOwnershipCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterOwnership *MonsterOwnershipSession) Paused() (bool, error) {
	return _MonsterOwnership.Contract.Paused(&_MonsterOwnership.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterOwnership *MonsterOwnershipCallerSession) Paused() (bool, error) {
	return _MonsterOwnership.Contract.Paused(&_MonsterOwnership.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) SaleAuction() (common.Address, error) {
	return _MonsterOwnership.Contract.SaleAuction(&_MonsterOwnership.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) SaleAuction() (common.Address, error) {
	return _MonsterOwnership.Contract.SaleAuction(&_MonsterOwnership.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterOwnership *MonsterOwnershipCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterOwnership *MonsterOwnershipSession) SecondsPerBlock() (*big.Int, error) {
	return _MonsterOwnership.Contract.SecondsPerBlock(&_MonsterOwnership.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterOwnership *MonsterOwnershipCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _MonsterOwnership.Contract.SecondsPerBlock(&_MonsterOwnership.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.SireAllowedToAddress(&_MonsterOwnership.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _MonsterOwnership.Contract.SireAllowedToAddress(&_MonsterOwnership.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipSession) SiringAuction() (common.Address, error) {
	return _MonsterOwnership.Contract.SiringAuction(&_MonsterOwnership.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterOwnership *MonsterOwnershipCallerSession) SiringAuction() (common.Address, error) {
	return _MonsterOwnership.Contract.SiringAuction(&_MonsterOwnership.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_MonsterOwnership *MonsterOwnershipCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_MonsterOwnership *MonsterOwnershipSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _MonsterOwnership.Contract.SupportsInterface(&_MonsterOwnership.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_MonsterOwnership *MonsterOwnershipCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _MonsterOwnership.Contract.SupportsInterface(&_MonsterOwnership.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonsterOwnership *MonsterOwnershipCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonsterOwnership *MonsterOwnershipSession) Symbol() (string, error) {
	return _MonsterOwnership.Contract.Symbol(&_MonsterOwnership.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonsterOwnership *MonsterOwnershipCallerSession) Symbol() (string, error) {
	return _MonsterOwnership.Contract.Symbol(&_MonsterOwnership.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_MonsterOwnership *MonsterOwnershipCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_MonsterOwnership *MonsterOwnershipSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _MonsterOwnership.Contract.TokenMetadata(&_MonsterOwnership.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_MonsterOwnership *MonsterOwnershipCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _MonsterOwnership.Contract.TokenMetadata(&_MonsterOwnership.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_MonsterOwnership *MonsterOwnershipCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_MonsterOwnership *MonsterOwnershipSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _MonsterOwnership.Contract.TokensOfOwner(&_MonsterOwnership.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_MonsterOwnership *MonsterOwnershipCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _MonsterOwnership.Contract.TokensOfOwner(&_MonsterOwnership.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonsterOwnership *MonsterOwnershipCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterOwnership.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonsterOwnership *MonsterOwnershipSession) TotalSupply() (*big.Int, error) {
	return _MonsterOwnership.Contract.TotalSupply(&_MonsterOwnership.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonsterOwnership *MonsterOwnershipCallerSession) TotalSupply() (*big.Int, error) {
	return _MonsterOwnership.Contract.TotalSupply(&_MonsterOwnership.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Approve(&_MonsterOwnership.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Approve(&_MonsterOwnership.TransactOpts, _to, _tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterOwnership *MonsterOwnershipSession) Pause() (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Pause(&_MonsterOwnership.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) Pause() (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Pause(&_MonsterOwnership.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterOwnership *MonsterOwnershipSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetCEO(&_MonsterOwnership.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetCEO(&_MonsterOwnership.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterOwnership *MonsterOwnershipSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetCFO(&_MonsterOwnership.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetCFO(&_MonsterOwnership.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterOwnership *MonsterOwnershipSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetCOO(&_MonsterOwnership.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetCOO(&_MonsterOwnership.TransactOpts, _newCOO)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_MonsterOwnership *MonsterOwnershipSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetMetadataAddress(&_MonsterOwnership.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetMetadataAddress(&_MonsterOwnership.TransactOpts, _contractAddress)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterOwnership *MonsterOwnershipSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetSecondsPerBlock(&_MonsterOwnership.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.SetSecondsPerBlock(&_MonsterOwnership.TransactOpts, secs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Transfer(&_MonsterOwnership.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Transfer(&_MonsterOwnership.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.TransferFrom(&_MonsterOwnership.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterOwnership.Contract.TransferFrom(&_MonsterOwnership.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterOwnership *MonsterOwnershipTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterOwnership.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterOwnership *MonsterOwnershipSession) Unpause() (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Unpause(&_MonsterOwnership.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterOwnership *MonsterOwnershipTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonsterOwnership.Contract.Unpause(&_MonsterOwnership.TransactOpts)
}

// MonsterOwnershipApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MonsterOwnership contract.
type MonsterOwnershipApprovalIterator struct {
	Event *MonsterOwnershipApproval // Event containing the contract specifics and raw log

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
func (it *MonsterOwnershipApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterOwnershipApproval)
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
		it.Event = new(MonsterOwnershipApproval)
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
func (it *MonsterOwnershipApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterOwnershipApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterOwnershipApproval represents a Approval event raised by the MonsterOwnership contract.
type MonsterOwnershipApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_MonsterOwnership *MonsterOwnershipFilterer) FilterApproval(opts *bind.FilterOpts) (*MonsterOwnershipApprovalIterator, error) {

	logs, sub, err := _MonsterOwnership.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipApprovalIterator{contract: _MonsterOwnership.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_MonsterOwnership *MonsterOwnershipFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MonsterOwnershipApproval) (event.Subscription, error) {

	logs, sub, err := _MonsterOwnership.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterOwnershipApproval)
				if err := _MonsterOwnership.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MonsterOwnershipBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the MonsterOwnership contract.
type MonsterOwnershipBirthIterator struct {
	Event *MonsterOwnershipBirth // Event containing the contract specifics and raw log

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
func (it *MonsterOwnershipBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterOwnershipBirth)
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
		it.Event = new(MonsterOwnershipBirth)
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
func (it *MonsterOwnershipBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterOwnershipBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterOwnershipBirth represents a Birth event raised by the MonsterOwnership contract.
type MonsterOwnershipBirth struct {
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
func (_MonsterOwnership *MonsterOwnershipFilterer) FilterBirth(opts *bind.FilterOpts) (*MonsterOwnershipBirthIterator, error) {

	logs, sub, err := _MonsterOwnership.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipBirthIterator{contract: _MonsterOwnership.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, MonsterId uint256, matronId uint256, sireId uint256, genes uint256)
func (_MonsterOwnership *MonsterOwnershipFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *MonsterOwnershipBirth) (event.Subscription, error) {

	logs, sub, err := _MonsterOwnership.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterOwnershipBirth)
				if err := _MonsterOwnership.contract.UnpackLog(event, "Birth", log); err != nil {
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

// MonsterOwnershipContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the MonsterOwnership contract.
type MonsterOwnershipContractUpgradeIterator struct {
	Event *MonsterOwnershipContractUpgrade // Event containing the contract specifics and raw log

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
func (it *MonsterOwnershipContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterOwnershipContractUpgrade)
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
		it.Event = new(MonsterOwnershipContractUpgrade)
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
func (it *MonsterOwnershipContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterOwnershipContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterOwnershipContractUpgrade represents a ContractUpgrade event raised by the MonsterOwnership contract.
type MonsterOwnershipContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterOwnership *MonsterOwnershipFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*MonsterOwnershipContractUpgradeIterator, error) {

	logs, sub, err := _MonsterOwnership.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipContractUpgradeIterator{contract: _MonsterOwnership.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterOwnership *MonsterOwnershipFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *MonsterOwnershipContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _MonsterOwnership.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterOwnershipContractUpgrade)
				if err := _MonsterOwnership.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// MonsterOwnershipTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MonsterOwnership contract.
type MonsterOwnershipTransferIterator struct {
	Event *MonsterOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *MonsterOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterOwnershipTransfer)
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
		it.Event = new(MonsterOwnershipTransfer)
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
func (it *MonsterOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterOwnershipTransfer represents a Transfer event raised by the MonsterOwnership contract.
type MonsterOwnershipTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_MonsterOwnership *MonsterOwnershipFilterer) FilterTransfer(opts *bind.FilterOpts) (*MonsterOwnershipTransferIterator, error) {

	logs, sub, err := _MonsterOwnership.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &MonsterOwnershipTransferIterator{contract: _MonsterOwnership.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_MonsterOwnership *MonsterOwnershipFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MonsterOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _MonsterOwnership.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterOwnershipTransfer)
				if err := _MonsterOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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
