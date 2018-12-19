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

// MonsterBreedingABI is the input ABI used to generate the binding from.
const MonsterBreedingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_monsterId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsterIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsterIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantMonsters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_monsterId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"MonsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// MonsterBreedingBin is the compiled bytecode used for deploying new contracts.
const MonsterBreedingBin = `6002805460a060020a60ff0219169055610240604052603c6080908152607860a05261012c60c05261025860e05261070861010052610e1061012052611c2061014052613840610160526170806101805261e1006101a052620151806101c0526202a3006101e052620546006102005262093a80610220526200008790600390600e620000ac565b50600f60055566071afd498d0000600e55348015620000a557600080fd5b5062000176565b6002830191839082156200013d5791602002820160005b838211156200010957835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000c3565b80156200013b5782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000109565b505b506200014b9291506200014f565b5090565b6200017391905b808211156200014b57805463ffffffff1916815560010162000156565b90565b6122a380620001866000396000f3006080604052600436106101df5763ffffffff60e060020a60003504166301ffc9a781146101e45780630519ce791461022f5780630560ff441461026057806306fdde03146102f9578063095ea7b31461030e5780630a0f81681461033457806318160ddd146103495780631940a9361461037057806321717ebf1461038857806323b872dd1461039d57806324e7a38a146103c757806327d7874c146103e85780632ba73c15146104095780633f4ba83a1461042a57806346116e6f1461043f57806346d22c70146104575780634b85fd55146104725780634dfff04f1461048a5780634e0a3379146104ae5780635663896e146104cf5780635c975abb146104e75780636352211e146104fc57806370a08231146105145780637a7d4937146105355780637d55aeea1461054a5780638456cb59146105625780638462151c1461057757806388c2a0bf146105e857806395d89b41146106005780639d6fac6f14610615578063a9059cbb14610646578063ad4b558c1461066a578063b047fb5014610682578063b0c35c0514610697578063bc4006f5146106ac578063c7e3ff4b146106c1578063d3e6f49f146106d6578063e17b25af146106ee578063e6cbe3511461070f578063f2b47d5214610724578063f7d8c88314610739575b600080fd5b3480156101f057600080fd5b5061021b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1960043516610747565b604080519115158252519081900360200190f35b34801561023b57600080fd5b506102446109da565b60408051600160a060020a039092168252519081900360200190f35b34801561026c57600080fd5b506102846004803590602480359081019101356109e9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102be5781810151838201526020016102a6565b50505050905090810190601f1680156102eb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561030557600080fd5b50610284610aec565b34801561031a57600080fd5b50610332600160a060020a0360043516602435610b23565b005b34801561034057600080fd5b50610244610ba5565b34801561035557600080fd5b5061035e610bb4565b60408051918252519081900360200190f35b34801561037c57600080fd5b5061021b600435610bbe565b34801561039457600080fd5b50610244610c03565b3480156103a957600080fd5b50610332600160a060020a0360043581169060243516604435610c12565b3480156103d357600080fd5b50610332600160a060020a0360043516610c8e565b3480156103f457600080fd5b50610332600160a060020a0360043516610d41565b34801561041557600080fd5b50610332600160a060020a0360043516610d8f565b34801561043657600080fd5b50610332610ddd565b34801561044b57600080fd5b50610244600435610e2c565b34801561046357600080fd5b5061021b600435602435610e47565b34801561047e57600080fd5b50610332600435610ec7565b34801561049657600080fd5b50610332600160a060020a0360043516602435610ee3565b3480156104ba57600080fd5b50610332600160a060020a0360043516610f3d565b3480156104db57600080fd5b50610332600435610f8b565b3480156104f357600080fd5b5061021b610fe7565b34801561050857600080fd5b50610244600435610ff7565b34801561052057600080fd5b5061035e600160a060020a036004351661101b565b34801561054157600080fd5b5061035e611036565b34801561055657600080fd5b5061024460043561103c565b34801561056e57600080fd5b50610332611057565b34801561058357600080fd5b50610598600160a060020a03600435166110d7565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105d45781810151838201526020016105bc565b505050509050019250505060405180910390f35b3480156105f457600080fd5b5061035e6004356111a9565b34801561060c57600080fd5b50610284611485565b34801561062157600080fd5b5061062d6004356114bc565b6040805163ffffffff9092168252519081900360200190f35b34801561065257600080fd5b50610332600160a060020a03600435166024356114e9565b34801561067657600080fd5b50610244600435611585565b34801561068e57600080fd5b506102446115a0565b3480156106a357600080fd5b5061035e6115af565b3480156106b857600080fd5b506102446115b5565b3480156106cd57600080fd5b5061035e6115c4565b3480156106e257600080fd5b5061021b6004356115ca565b3480156106fa57600080fd5b50610332600160a060020a0360043516611699565b34801561071b57600080fd5b506102446116d2565b34801561073057600080fd5b506102446116e1565b6103326004356024356116f0565b604080517f737570706f727473496e74657266616365286279746573342900000000000000815290519081900360190190206000907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19838116911614806109d25750604080517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152815190819003601d0181207f746f6b656e734f664f776e657228616464726573732900000000000000000000825282519182900360160182207f7472616e7366657246726f6d28616464726573732c616464726573732c75696e83527f7432353629000000000000000000000000000000000000000000000000000000602084015283519283900360250183207f7472616e7366657228616464726573732c75696e743235362900000000000000845284519384900360190184207f617070726f766528616464726573732c75696e74323536290000000000000000855285519485900360180185207f6f776e65724f662875696e743235362900000000000000000000000000000000865286519586900360100186207f62616c616e63654f662861646472657373290000000000000000000000000000875287519687900360120187207f746f74616c537570706c792829000000000000000000000000000000000000008852885197889003600d0188207f73796d626f6c2829000000000000000000000000000000000000000000000000895289519889900360080189207f6e616d65282900000000000000000000000000000000000000000000000000008a529951988990036006019098207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff198c811691909a189098181818181818181891909116145b90505b919050565b600154600160a060020a031681565b60606109f3612214565b600d54600090600160a060020a03161515610a0d57600080fd5b600d54604080517fcb4799f2000000000000000000000000000000000000000000000000000000008152600481018981526024820192835260448201889052600160a060020a039093169263cb4799f2928a928a928a929091606401848480828437820191505094505050505060a060405180830381600087803b158015610a9457600080fd5b505af1158015610aa8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060a0811015610acd57600080fd5b5060808101519092509050610ae282826118f4565b9695505050505050565b60408051808201909152600781527f6d6f6e7374657200000000000000000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610b3a57600080fd5b610b443382611948565b1515610b4f57600080fd5b610b598183611968565b60408051338152600160a060020a038416602082015280820183905290517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360600190a15050565b600054600160a060020a031681565b6006546000190190565b6000808211610bcc57600080fd5b6006805483908110610bda57fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b600c54600160a060020a031681565b60025460a060020a900460ff1615610c2957600080fd5b600160a060020a0382161515610c3e57600080fd5b600160a060020a038216301415610c5457600080fd5b610c5e3382611996565b1515610c6957600080fd5b610c738382611948565b1515610c7e57600080fd5b610c898383836119b6565b505050565b60008054600160a060020a03163314610ca657600080fd5b81905080600160a060020a03166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610ce757600080fd5b505af1158015610cfb573d6000803e3d6000fd5b505050506040513d6020811015610d1157600080fd5b50511515610d1e57600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b600054600160a060020a03163314610d5857600080fd5b600160a060020a0381161515610d6d57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610da657600080fd5b600160a060020a0381161515610dbb57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610df457600080fd5b60025460a060020a900460ff161515610e0c57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b60008080808511610e5757600080fd5b60008411610e6457600080fd5b6006805486908110610e7257fe5b90600052602060002090600202019150600684815481101515610e9157fe5b90600052602060002090600202019050610ead82868387611a98565b8015610ebe5750610ebe8486611c18565b95945050505050565b600254600160a060020a03163314610ede57600080fd5b600e55565b60025460a060020a900460ff1615610efa57600080fd5b610f043382611948565b1515610f0f57600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a03163314610f5457600080fd5b600160a060020a0381161515610f6957600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a0316331480610fae5750600054600160a060020a031633145b80610fc35750600154600160a060020a031633145b1515610fce57600080fd5b60035463ffffffff168110610fe257600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a03168015156109d557600080fd5b600160a060020a031660009081526008602052604090205490565b60055481565b600960205260009081526040902054600160a060020a031681565b600254600160a060020a031633148061107a5750600054600160a060020a031633145b8061108f5750600154600160a060020a031633145b151561109a57600080fd5b60025460a060020a900460ff16156110b157600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b60606000606060008060006110eb8761101b565b945084151561110a57604080516000815260208101909152955061119f565b84604051908082528060200260200182016040528015611134578160200160208202803883390190505b50935061113f610bb4565b925060009150600190505b82811161119b57600081815260076020526040902054600160a060020a03888116911614156111935780848381518110151561118257fe5b602090810290910101526001909101905b60010161114a565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff161515156111d157600080fd5b600680548a9081106111df57fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff16151561120c57600080fd5b604080516101008101825288548152600189015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e082015261129d90611c6d565b15156112a857600080fd5b60018701546006805460c060020a90920463ffffffff16975090879081106112cc57fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a92839004811696509190041684101561131957600185015460f060020a900461ffff1693505b6010548754865460018a0154604080517f0d9f5aed0000000000000000000000000000000000000000000000000000000081526004810194909452602484019290925260001967ffffffffffffffff6801000000000000000090920482160116604483015251600160a060020a0390921691630d9f5aed916064808201926020929091908290030181600087803b1580156113b357600080fd5b505af11580156113c7573d6000803e3d6000fd5b505050506040513d60208110156113dd57600080fd5b505160008a815260076020526040902054600189810154929550600160a060020a039091169350611426918b9160c060020a90910463ffffffff1690870161ffff168686611c9d565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54604051919250339181156108fc0291906000818181858888f150939c9b505050505050505050505050565b60408051808201909152600281527f4d52000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e81106114c957fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b60025460a060020a900460ff161561150057600080fd5b600160a060020a038216151561151557600080fd5b600160a060020a03821630141561152b57600080fd5b600b54600160a060020a038381169116141561154657600080fd5b600c54600160a060020a038381169116141561156157600080fd5b61156b3382611948565b151561157657600080fd5b6115813383836119b6565b5050565b600760205260009081526040902054600160a060020a031681565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b600f5481565b6000808083116115d957600080fd5b60068054849081106115e757fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e082015290915061169290611f7a565b9392505050565b600054600160a060020a031633146116b057600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b601054600160a060020a031681565b600254600090819060a060020a900460ff161561170c57600080fd5b600e5434101561171b57600080fd5b6117253385611948565b151561173057600080fd5b61173a8385611c18565b151561174557600080fd5b600680548590811061175357fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e08201529092506117fe90611f7a565b151561180957600080fd5b600680548490811061181757fe5b60009182526020918290206040805161010081018252600290930290910180548352600181015467ffffffffffffffff808216958501959095526801000000000000000081049094169183019190915263ffffffff608060020a84048116606084015260a060020a84048116608084015260c060020a84041660a083015261ffff60e060020a8404811660c084015260f060020a90930490921660e08201529091506118c290611f7a565b15156118cd57600080fd5b6118d982858386611a98565b15156118e457600080fd5b6118ee8484611fa9565b50505050565b606080600080846040519080825280601f01601f191660200182016040528015611928578160200160208202803883390190505b509250506020820190508461193e8282876120e7565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a0319169091179055831615611a4957600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b60408051600160a060020a0380861682528416602082015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a1505050565b600081841415611aaa57506000611c10565b6001850154608060020a900463ffffffff16821480611ad95750600185015460a060020a900463ffffffff1682145b15611ae657506000611c10565b6001830154608060020a900463ffffffff16841480611b155750600183015460a060020a900463ffffffff1684145b15611b2257506000611c10565b6001830154608060020a900463ffffffff161580611b4f57506001850154608060020a900463ffffffff16155b15611b5c57506001611c10565b60018581015490840154608060020a9182900463ffffffff90811692909104161480611ba7575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b15611bb457506000611c10565b6001808601549084015460a060020a900463ffffffff908116608060020a909204161480611bff57506001858101549084015460a060020a9182900463ffffffff9081169290910416145b15611c0c57506000611c10565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a03918216911680821480610ebe57506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b60008160a0015163ffffffff166000141580156109d25750506040015167ffffffffffffffff4381169116111590565b600080611ca8612233565b600063ffffffff89168914611cbc57600080fd5b63ffffffff88168814611cce57600080fd5b61ffff87168714611cde57600080fd5b600287049250600d8361ffff161115611cf657600d92505b505060408051610100810182528581524267ffffffffffffffff90811660208301908152600093830184815263ffffffff8c8116606086019081528c82166080870190815260a0870188815261ffff8a811660c08a019081528f821660e08b01908152600680546001810182559c528a517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60028e029081019190915598517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909901805498519651955194519251915167ffffffffffffffff19909916998b16999099176fffffffffffffffff000000000000000019166801000000000000000096909a16959095029890981773ffffffff000000000000000000000000000000001916608060020a938616939093029290921777ffffffff0000000000000000000000000000000000000000191660a060020a91851691909102177bffffffff000000000000000000000000000000000000000000000000191660c060020a96841696909602959095177fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e060020a91861691909102177dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f060020a929094169190910292909217905590919081168114611ef457600080fd5b606080830151608080850151855160408051600160a060020a038c1681526020810188905263ffffffff95861681830152929094169482019490945290810192909252517f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad59181900360a00190a1611f6e600086836119b6565b98975050505050505050565b60008160a0015163ffffffff1660001480156109d25750506040015167ffffffffffffffff4381169116111590565b600080600683815481101515611fbb57fe5b90600052602060002090600202019150600684815481101515611fda57fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff871602179055905061202a8261212b565b6120338161212b565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f80546001908101909155878452600783529281902054928401548151600160a060020a0390941684529183018790528281018690526801000000000000000090910467ffffffffffffffff166060830152517f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80916080908290030190a150505050565b60005b6020821061210c578251845260209384019390920191601f19909101906120ea565b50905182516020929092036101000a6000190180199091169116179052565b600554600182015443919060039060e060020a900461ffff16600e811061214e57fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff1681151561217957fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff161015612211576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b6080604051908101604052806004906020820280388339509192915050565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152905600a165627a7a723058208d5008cee9d540f7c644cc8db3d37d7d96c84e2b7dba124aee67eb994ec724290029`

// DeployMonsterBreeding deploys a new Ethereum contract, binding an instance of MonsterBreeding to it.
func DeployMonsterBreeding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonsterBreeding, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterBreedingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MonsterBreedingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonsterBreeding{MonsterBreedingCaller: MonsterBreedingCaller{contract: contract}, MonsterBreedingTransactor: MonsterBreedingTransactor{contract: contract}, MonsterBreedingFilterer: MonsterBreedingFilterer{contract: contract}}, nil
}

// MonsterBreeding is an auto generated Go binding around an Ethereum contract.
type MonsterBreeding struct {
	MonsterBreedingCaller     // Read-only binding to the contract
	MonsterBreedingTransactor // Write-only binding to the contract
	MonsterBreedingFilterer   // Log filterer for contract events
}

// MonsterBreedingCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonsterBreedingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterBreedingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonsterBreedingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterBreedingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonsterBreedingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterBreedingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonsterBreedingSession struct {
	Contract     *MonsterBreeding  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MonsterBreedingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonsterBreedingCallerSession struct {
	Contract *MonsterBreedingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MonsterBreedingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonsterBreedingTransactorSession struct {
	Contract     *MonsterBreedingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MonsterBreedingRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonsterBreedingRaw struct {
	Contract *MonsterBreeding // Generic contract binding to access the raw methods on
}

// MonsterBreedingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonsterBreedingCallerRaw struct {
	Contract *MonsterBreedingCaller // Generic read-only contract binding to access the raw methods on
}

// MonsterBreedingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonsterBreedingTransactorRaw struct {
	Contract *MonsterBreedingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonsterBreeding creates a new instance of MonsterBreeding, bound to a specific deployed contract.
func NewMonsterBreeding(address common.Address, backend bind.ContractBackend) (*MonsterBreeding, error) {
	contract, err := bindMonsterBreeding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonsterBreeding{MonsterBreedingCaller: MonsterBreedingCaller{contract: contract}, MonsterBreedingTransactor: MonsterBreedingTransactor{contract: contract}, MonsterBreedingFilterer: MonsterBreedingFilterer{contract: contract}}, nil
}

// NewMonsterBreedingCaller creates a new read-only instance of MonsterBreeding, bound to a specific deployed contract.
func NewMonsterBreedingCaller(address common.Address, caller bind.ContractCaller) (*MonsterBreedingCaller, error) {
	contract, err := bindMonsterBreeding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingCaller{contract: contract}, nil
}

// NewMonsterBreedingTransactor creates a new write-only instance of MonsterBreeding, bound to a specific deployed contract.
func NewMonsterBreedingTransactor(address common.Address, transactor bind.ContractTransactor) (*MonsterBreedingTransactor, error) {
	contract, err := bindMonsterBreeding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingTransactor{contract: contract}, nil
}

// NewMonsterBreedingFilterer creates a new log filterer instance of MonsterBreeding, bound to a specific deployed contract.
func NewMonsterBreedingFilterer(address common.Address, filterer bind.ContractFilterer) (*MonsterBreedingFilterer, error) {
	contract, err := bindMonsterBreeding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingFilterer{contract: contract}, nil
}

// bindMonsterBreeding binds a generic wrapper to an already deployed contract.
func bindMonsterBreeding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MonsterBreedingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterBreeding *MonsterBreedingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterBreeding.Contract.MonsterBreedingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterBreeding *MonsterBreedingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.MonsterBreedingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterBreeding *MonsterBreedingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.MonsterBreedingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterBreeding *MonsterBreedingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MonsterBreeding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterBreeding *MonsterBreedingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterBreeding *MonsterBreedingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.contract.Transact(opts, method, params...)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "autoBirthFee")
	return *ret0, err
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingSession) AutoBirthFee() (*big.Int, error) {
	return _MonsterBreeding.Contract.AutoBirthFee(&_MonsterBreeding.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCallerSession) AutoBirthFee() (*big.Int, error) {
	return _MonsterBreeding.Contract.AutoBirthFee(&_MonsterBreeding.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_MonsterBreeding *MonsterBreedingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_MonsterBreeding *MonsterBreedingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MonsterBreeding.Contract.BalanceOf(&_MonsterBreeding.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_MonsterBreeding *MonsterBreedingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MonsterBreeding.Contract.BalanceOf(&_MonsterBreeding.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "canBreedWith", _matronId, _sireId)
	return *ret0, err
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _MonsterBreeding.Contract.CanBreedWith(&_MonsterBreeding.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _MonsterBreeding.Contract.CanBreedWith(&_MonsterBreeding.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) CeoAddress() (common.Address, error) {
	return _MonsterBreeding.Contract.CeoAddress(&_MonsterBreeding.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) CeoAddress() (common.Address, error) {
	return _MonsterBreeding.Contract.CeoAddress(&_MonsterBreeding.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) CfoAddress() (common.Address, error) {
	return _MonsterBreeding.Contract.CfoAddress(&_MonsterBreeding.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) CfoAddress() (common.Address, error) {
	return _MonsterBreeding.Contract.CfoAddress(&_MonsterBreeding.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) CooAddress() (common.Address, error) {
	return _MonsterBreeding.Contract.CooAddress(&_MonsterBreeding.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) CooAddress() (common.Address, error) {
	return _MonsterBreeding.Contract.CooAddress(&_MonsterBreeding.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterBreeding *MonsterBreedingCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterBreeding *MonsterBreedingSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _MonsterBreeding.Contract.Cooldowns(&_MonsterBreeding.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_MonsterBreeding *MonsterBreedingCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _MonsterBreeding.Contract.Cooldowns(&_MonsterBreeding.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) Erc721Metadata() (common.Address, error) {
	return _MonsterBreeding.Contract.Erc721Metadata(&_MonsterBreeding.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) Erc721Metadata() (common.Address, error) {
	return _MonsterBreeding.Contract.Erc721Metadata(&_MonsterBreeding.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "geneScience")
	return *ret0, err
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) GeneScience() (common.Address, error) {
	return _MonsterBreeding.Contract.GeneScience(&_MonsterBreeding.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) GeneScience() (common.Address, error) {
	return _MonsterBreeding.Contract.GeneScience(&_MonsterBreeding.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_monsterId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCaller) IsPregnant(opts *bind.CallOpts, _monsterId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "isPregnant", _monsterId)
	return *ret0, err
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_monsterId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingSession) IsPregnant(_monsterId *big.Int) (bool, error) {
	return _MonsterBreeding.Contract.IsPregnant(&_MonsterBreeding.CallOpts, _monsterId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_monsterId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCallerSession) IsPregnant(_monsterId *big.Int) (bool, error) {
	return _MonsterBreeding.Contract.IsPregnant(&_MonsterBreeding.CallOpts, _monsterId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_monsterId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCaller) IsReadyToBreed(opts *bind.CallOpts, _monsterId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "isReadyToBreed", _monsterId)
	return *ret0, err
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_monsterId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingSession) IsReadyToBreed(_monsterId *big.Int) (bool, error) {
	return _MonsterBreeding.Contract.IsReadyToBreed(&_MonsterBreeding.CallOpts, _monsterId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_monsterId uint256) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCallerSession) IsReadyToBreed(_monsterId *big.Int) (bool, error) {
	return _MonsterBreeding.Contract.IsReadyToBreed(&_MonsterBreeding.CallOpts, _monsterId)
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) MonsterIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "monsterIndexToApproved", arg0)
	return *ret0, err
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) MonsterIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.MonsterIndexToApproved(&_MonsterBreeding.CallOpts, arg0)
}

// MonsterIndexToApproved is a free data retrieval call binding the contract method 0x7d55aeea.
//
// Solidity: function monsterIndexToApproved( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) MonsterIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.MonsterIndexToApproved(&_MonsterBreeding.CallOpts, arg0)
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) MonsterIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "monsterIndexToOwner", arg0)
	return *ret0, err
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) MonsterIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.MonsterIndexToOwner(&_MonsterBreeding.CallOpts, arg0)
}

// MonsterIndexToOwner is a free data retrieval call binding the contract method 0xad4b558c.
//
// Solidity: function monsterIndexToOwner( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) MonsterIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.MonsterIndexToOwner(&_MonsterBreeding.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonsterBreeding *MonsterBreedingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonsterBreeding *MonsterBreedingSession) Name() (string, error) {
	return _MonsterBreeding.Contract.Name(&_MonsterBreeding.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MonsterBreeding *MonsterBreedingCallerSession) Name() (string, error) {
	return _MonsterBreeding.Contract.Name(&_MonsterBreeding.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_MonsterBreeding *MonsterBreedingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_MonsterBreeding *MonsterBreedingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.OwnerOf(&_MonsterBreeding.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_MonsterBreeding *MonsterBreedingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.OwnerOf(&_MonsterBreeding.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterBreeding *MonsterBreedingSession) Paused() (bool, error) {
	return _MonsterBreeding.Contract.Paused(&_MonsterBreeding.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCallerSession) Paused() (bool, error) {
	return _MonsterBreeding.Contract.Paused(&_MonsterBreeding.CallOpts)
}

// PregnantMonsters is a free data retrieval call binding the contract method 0xc7e3ff4b.
//
// Solidity: function pregnantMonsters() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCaller) PregnantMonsters(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "pregnantMonsters")
	return *ret0, err
}

// PregnantMonsters is a free data retrieval call binding the contract method 0xc7e3ff4b.
//
// Solidity: function pregnantMonsters() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingSession) PregnantMonsters() (*big.Int, error) {
	return _MonsterBreeding.Contract.PregnantMonsters(&_MonsterBreeding.CallOpts)
}

// PregnantMonsters is a free data retrieval call binding the contract method 0xc7e3ff4b.
//
// Solidity: function pregnantMonsters() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCallerSession) PregnantMonsters() (*big.Int, error) {
	return _MonsterBreeding.Contract.PregnantMonsters(&_MonsterBreeding.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) SaleAuction() (common.Address, error) {
	return _MonsterBreeding.Contract.SaleAuction(&_MonsterBreeding.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) SaleAuction() (common.Address, error) {
	return _MonsterBreeding.Contract.SaleAuction(&_MonsterBreeding.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingSession) SecondsPerBlock() (*big.Int, error) {
	return _MonsterBreeding.Contract.SecondsPerBlock(&_MonsterBreeding.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _MonsterBreeding.Contract.SecondsPerBlock(&_MonsterBreeding.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.SireAllowedToAddress(&_MonsterBreeding.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _MonsterBreeding.Contract.SireAllowedToAddress(&_MonsterBreeding.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterBreeding *MonsterBreedingSession) SiringAuction() (common.Address, error) {
	return _MonsterBreeding.Contract.SiringAuction(&_MonsterBreeding.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_MonsterBreeding *MonsterBreedingCallerSession) SiringAuction() (common.Address, error) {
	return _MonsterBreeding.Contract.SiringAuction(&_MonsterBreeding.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _MonsterBreeding.Contract.SupportsInterface(&_MonsterBreeding.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_MonsterBreeding *MonsterBreedingCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _MonsterBreeding.Contract.SupportsInterface(&_MonsterBreeding.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonsterBreeding *MonsterBreedingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonsterBreeding *MonsterBreedingSession) Symbol() (string, error) {
	return _MonsterBreeding.Contract.Symbol(&_MonsterBreeding.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MonsterBreeding *MonsterBreedingCallerSession) Symbol() (string, error) {
	return _MonsterBreeding.Contract.Symbol(&_MonsterBreeding.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_MonsterBreeding *MonsterBreedingCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_MonsterBreeding *MonsterBreedingSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _MonsterBreeding.Contract.TokenMetadata(&_MonsterBreeding.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_MonsterBreeding *MonsterBreedingCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _MonsterBreeding.Contract.TokenMetadata(&_MonsterBreeding.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_MonsterBreeding *MonsterBreedingCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_MonsterBreeding *MonsterBreedingSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _MonsterBreeding.Contract.TokensOfOwner(&_MonsterBreeding.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_MonsterBreeding *MonsterBreedingCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _MonsterBreeding.Contract.TokensOfOwner(&_MonsterBreeding.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MonsterBreeding.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingSession) TotalSupply() (*big.Int, error) {
	return _MonsterBreeding.Contract.TotalSupply(&_MonsterBreeding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MonsterBreeding *MonsterBreedingCallerSession) TotalSupply() (*big.Int, error) {
	return _MonsterBreeding.Contract.TotalSupply(&_MonsterBreeding.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Approve(&_MonsterBreeding.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Approve(&_MonsterBreeding.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.ApproveSiring(&_MonsterBreeding.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.ApproveSiring(&_MonsterBreeding.TransactOpts, _addr, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.BreedWithAuto(&_MonsterBreeding.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.BreedWithAuto(&_MonsterBreeding.TransactOpts, _matronId, _sireId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_MonsterBreeding *MonsterBreedingTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_MonsterBreeding *MonsterBreedingSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.GiveBirth(&_MonsterBreeding.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_MonsterBreeding *MonsterBreedingTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.GiveBirth(&_MonsterBreeding.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterBreeding *MonsterBreedingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterBreeding *MonsterBreedingSession) Pause() (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Pause(&_MonsterBreeding.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) Pause() (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Pause(&_MonsterBreeding.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetAutoBirthFee(&_MonsterBreeding.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetAutoBirthFee(&_MonsterBreeding.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetCEO(&_MonsterBreeding.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetCEO(&_MonsterBreeding.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetCFO(&_MonsterBreeding.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetCFO(&_MonsterBreeding.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetCOO(&_MonsterBreeding.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetCOO(&_MonsterBreeding.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetGeneScienceAddress(&_MonsterBreeding.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetGeneScienceAddress(&_MonsterBreeding.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetMetadataAddress(&_MonsterBreeding.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetMetadataAddress(&_MonsterBreeding.TransactOpts, _contractAddress)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetSecondsPerBlock(&_MonsterBreeding.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.SetSecondsPerBlock(&_MonsterBreeding.TransactOpts, secs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Transfer(&_MonsterBreeding.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Transfer(&_MonsterBreeding.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.TransferFrom(&_MonsterBreeding.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MonsterBreeding.Contract.TransferFrom(&_MonsterBreeding.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterBreeding *MonsterBreedingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterBreeding.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterBreeding *MonsterBreedingSession) Unpause() (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Unpause(&_MonsterBreeding.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterBreeding *MonsterBreedingTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonsterBreeding.Contract.Unpause(&_MonsterBreeding.TransactOpts)
}

// MonsterBreedingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MonsterBreeding contract.
type MonsterBreedingApprovalIterator struct {
	Event *MonsterBreedingApproval // Event containing the contract specifics and raw log

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
func (it *MonsterBreedingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBreedingApproval)
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
		it.Event = new(MonsterBreedingApproval)
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
func (it *MonsterBreedingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBreedingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBreedingApproval represents a Approval event raised by the MonsterBreeding contract.
type MonsterBreedingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) FilterApproval(opts *bind.FilterOpts) (*MonsterBreedingApprovalIterator, error) {

	logs, sub, err := _MonsterBreeding.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingApprovalIterator{contract: _MonsterBreeding.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner address, approved address, tokenId uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MonsterBreedingApproval) (event.Subscription, error) {

	logs, sub, err := _MonsterBreeding.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBreedingApproval)
				if err := _MonsterBreeding.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MonsterBreedingBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the MonsterBreeding contract.
type MonsterBreedingBirthIterator struct {
	Event *MonsterBreedingBirth // Event containing the contract specifics and raw log

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
func (it *MonsterBreedingBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBreedingBirth)
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
		it.Event = new(MonsterBreedingBirth)
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
func (it *MonsterBreedingBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBreedingBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBreedingBirth represents a Birth event raised by the MonsterBreeding contract.
type MonsterBreedingBirth struct {
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
func (_MonsterBreeding *MonsterBreedingFilterer) FilterBirth(opts *bind.FilterOpts) (*MonsterBreedingBirthIterator, error) {

	logs, sub, err := _MonsterBreeding.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingBirthIterator{contract: _MonsterBreeding.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: e Birth(owner address, MonsterId uint256, matronId uint256, sireId uint256, genes uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *MonsterBreedingBirth) (event.Subscription, error) {

	logs, sub, err := _MonsterBreeding.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBreedingBirth)
				if err := _MonsterBreeding.contract.UnpackLog(event, "Birth", log); err != nil {
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

// MonsterBreedingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the MonsterBreeding contract.
type MonsterBreedingContractUpgradeIterator struct {
	Event *MonsterBreedingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *MonsterBreedingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBreedingContractUpgrade)
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
		it.Event = new(MonsterBreedingContractUpgrade)
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
func (it *MonsterBreedingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBreedingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBreedingContractUpgrade represents a ContractUpgrade event raised by the MonsterBreeding contract.
type MonsterBreedingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterBreeding *MonsterBreedingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*MonsterBreedingContractUpgradeIterator, error) {

	logs, sub, err := _MonsterBreeding.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingContractUpgradeIterator{contract: _MonsterBreeding.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: e ContractUpgrade(newContract address)
func (_MonsterBreeding *MonsterBreedingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *MonsterBreedingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _MonsterBreeding.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBreedingContractUpgrade)
				if err := _MonsterBreeding.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// MonsterBreedingPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the MonsterBreeding contract.
type MonsterBreedingPregnantIterator struct {
	Event *MonsterBreedingPregnant // Event containing the contract specifics and raw log

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
func (it *MonsterBreedingPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBreedingPregnant)
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
		it.Event = new(MonsterBreedingPregnant)
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
func (it *MonsterBreedingPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBreedingPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBreedingPregnant represents a Pregnant event raised by the MonsterBreeding contract.
type MonsterBreedingPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) FilterPregnant(opts *bind.FilterOpts) (*MonsterBreedingPregnantIterator, error) {

	logs, sub, err := _MonsterBreeding.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingPregnantIterator{contract: _MonsterBreeding.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: e Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *MonsterBreedingPregnant) (event.Subscription, error) {

	logs, sub, err := _MonsterBreeding.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBreedingPregnant)
				if err := _MonsterBreeding.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// MonsterBreedingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MonsterBreeding contract.
type MonsterBreedingTransferIterator struct {
	Event *MonsterBreedingTransfer // Event containing the contract specifics and raw log

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
func (it *MonsterBreedingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterBreedingTransfer)
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
		it.Event = new(MonsterBreedingTransfer)
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
func (it *MonsterBreedingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterBreedingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterBreedingTransfer represents a Transfer event raised by the MonsterBreeding contract.
type MonsterBreedingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) FilterTransfer(opts *bind.FilterOpts) (*MonsterBreedingTransferIterator, error) {

	logs, sub, err := _MonsterBreeding.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &MonsterBreedingTransferIterator{contract: _MonsterBreeding.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from address, to address, tokenId uint256)
func (_MonsterBreeding *MonsterBreedingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MonsterBreedingTransfer) (event.Subscription, error) {

	logs, sub, err := _MonsterBreeding.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterBreedingTransfer)
				if err := _MonsterBreeding.contract.UnpackLog(event, "Transfer", log); err != nil {
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
