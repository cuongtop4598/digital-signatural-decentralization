// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package document

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// DocumentMetaData contains all meta data concerning the Document contract.
var DocumentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numdoc\",\"type\":\"uint256\"}],\"name\":\"IndexDocument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Status\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"compareBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareString\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"name\":\"getHashUserInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSinger\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9201de55": "bytes32ToString(bytes32)",
		"ef1bade3": "compareBytes(bytes32,bytes32)",
		"75b83187": "compareString(string,string)",
		"fa540801": "getEthSignedMessageHash(bytes32)",
		"a23f1228": "getHashUserInfo(string)",
		"cfd6baa2": "getSignature(string,uint256)",
		"b26feeb9": "getSinger(string,bytes32,uint256)",
		"97aba7f9": "recoverSigner(bytes32,bytes)",
		"3a8fb351": "saveDoc(string,bytes)",
		"a7bb5803": "splitSignature(bytes)",
		"8e16afb2": "storeUser(string,string,string,string,string,string,address)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
		"a88c6e45": "verifyUser(string,string,string,string,string,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506112ee806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063a7bb58031161008c578063cfd6baa211610066578063cfd6baa2146101e7578063d2de39a814610208578063ef1bade31461021b578063fa5408011461022e57600080fd5b8063a7bb580314610190578063a88c6e45146101c1578063b26feeb9146101d457600080fd5b80633a8fb351146100d457806375b83187146100e95780638e16afb2146101115780639201de551461013257806397aba7f914610152578063a23f12281461017d575b600080fd5b6100e76100e2366004610e0e565b610241565b005b6100fc6100f7366004610e0e565b610405565b60405190151581526020015b60405180910390f35b61012461011f366004610f49565b610470565b604051908152602001610108565b610145610140366004610d01565b610517565b6040516101089190611164565b610165610160366004610d3c565b61064e565b6040516001600160a01b039091168152602001610108565b61012461018b366004610d83565b6106cd565b6101a361019e366004610d83565b610742565b60408051938452602084019290925260ff1690820152606001610108565b6100fc6101cf366004610e68565b6107b6565b6101456101e2366004610dc0565b610836565b6101fa6101f5366004611050565b6108a7565b604051610108929190611177565b6100fc610216366004610dc0565b610a95565b6100fc610229366004610d1a565b610adf565b61012461023c366004610d01565b610b19565b600061024c83610b6d565b905060005b60005481116103cc57610288600082815481106102705761027061128c565b90600052602060002090600602016002015483610adf565b156103ba5782600082815481106102a1576102a161128c565b906000526020600020906006020160050160008084815481106102c6576102c661128c565b906000526020600020906006020160040154815260200190815260200160002060000190805190602001906102fc929190610bbf565b50600081815481106103105761031061128c565b906000526020600020906006020160040154600161032e91906111a1565b600082815481106103415761034161128c565b9060005260206000209060060201600401819055507fca1bddf79d7a20af180f6d50bd9b30dd13435374d6c3e6d35e319f24eee11ba060016000838154811061038c5761038c61128c565b9060005260206000209060060201600401546103a891906111b9565b60405190815260200160405180910390a15b806103c48161123b565b915050610251565b60405162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b60448201526064015b60405180910390fd5b600081518351146104185750600061046a565b8160405160200161042991906110c1565b604051602081830303815290604052805190602001208360405160200161045091906110c1565b604051602081830303815290604052805190602001201490505b92915050565b600080610481888888888888610b80565b6000805460018101808355828052929350918190839081106104a5576104a561128c565b60009182526020918290208d5160069092020192506104c9918391908e0190610bbf565b50600181018390556003810180546001600160a01b0319166001600160a01b0387161790556104f787610b6d565b6002820181905560006004909201919091559a9950505050505050505050565b606060005b60208160ff161080156105505750828160ff166020811061053f5761053f61128c565b1a60f81b6001600160f81b03191615155b15610567578061055f81611256565b91505061051c565b60008160ff1667ffffffffffffffff811115610585576105856112a2565b6040519080825280601f01601f1916602001820160405280156105af576020820181803683370190505b509050600091505b60208260ff161080156105eb5750838260ff16602081106105da576105da61128c565b1a60f81b6001600160f81b03191615155b1561064757838260ff16602081106106055761060561128c565b1a60f81b818360ff168151811061061e5761061e61128c565b60200101906001600160f81b031916908160001a9053508161063f81611256565b9250506105b7565b9392505050565b60008060008061065d85610742565b6040805160008152602081018083528b905260ff8316918101919091526060810184905260808101839052929550909350915060019060a0016020604051602081039080840390855afa1580156106b8573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000806106d983610b6d565b905060005b60005481116103cc576106fd600082815481106102705761027061128c565b1561073057600081815481106107155761071561128c565b90600052602060002090600602016001015492505050919050565b8061073a8161123b565b9150506106de565b600080600083516041146107985760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e677468000000000000000060448201526064016103fc565b50505060208101516040820151606090920151909260009190911a90565b6000806107c7888888888888610b80565b905060005b600054811161082557610803600082815481106107eb576107eb61128c565b90600052602060002090600602016001015483610adf565b156108135760019250505061082c565b8061081d8161123b565b9150506107cc565b6000925050505b9695505050505050565b606080600061084586856108a7565b9092509050600061085586610b19565b90506000610863828561064e565b60405160200161088b919060609190911b6bffffffffffffffffffffffff1916815260140190565b60408051808303601f1901815291905298975050505050505050565b60606000806108db856040516020016108c091906110c1565b60405160208183030381529060405280519060200120610517565b90506000805b60005481101561097b57826040516020016108fc91906110c1565b60405160208183030381529060405280519060200120600082815481106109255761092561128c565b90600052602060002090600602016002015460405160200161094991815260200190565b604051602081830303815290604052805190602001201415610969578091505b806109738161123b565b9150506108e1565b5060008082815481106109905761099061128c565b906000526020600020906006020160050160008781526020019081526020016000206040518060200160405290816000820180546109cd90611200565b80601f01602080910402602001604051908101604052809291908181526020018280546109f990611200565b8015610a465780601f10610a1b57610100808354040283529160200191610a46565b820191906000526020600020905b815481529060010190602001808311610a2957829003601f168201915b5050505050815250509050806000015160008381548110610a6957610a6961128c565b60009182526020909120600360069092020101549095506001600160a01b031693505050509250929050565b600060606000610aa586856108a7565b90925090506000610ab586610b19565b90506000610ac3828561064e565b6001600160a01b03908116931692909214979650505050505050565b600081604051602001610af491815260200190565b60408051601f1981840301815282825280516020918201209083018690529101610450565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c015b604051602081830303815290604052805190602001209050919050565b600081604051602001610b5091906110c1565b6000868686868686604051602001610b9d969594939291906110dd565b6040516020818303038152906040528051906020012090509695505050505050565b828054610bcb90611200565b90600052602060002090601f016020900481019282610bed5760008555610c33565b82601f10610c0657805160ff1916838001178555610c33565b82800160010185558215610c33579182015b82811115610c33578251825591602001919060010190610c18565b50610c3f929150610c43565b5090565b5b80821115610c3f5760008155600101610c44565b80356001600160a01b0381168114610c6f57600080fd5b919050565b600082601f830112610c8557600080fd5b813567ffffffffffffffff80821115610ca057610ca06112a2565b604051601f8301601f19908116603f01168101908282118183101715610cc857610cc86112a2565b81604052838152866020858801011115610ce157600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215610d1357600080fd5b5035919050565b60008060408385031215610d2d57600080fd5b50508035926020909101359150565b60008060408385031215610d4f57600080fd5b82359150602083013567ffffffffffffffff811115610d6d57600080fd5b610d7985828601610c74565b9150509250929050565b600060208284031215610d9557600080fd5b813567ffffffffffffffff811115610dac57600080fd5b610db884828501610c74565b949350505050565b600080600060608486031215610dd557600080fd5b833567ffffffffffffffff811115610dec57600080fd5b610df886828701610c74565b9660208601359650604090950135949350505050565b60008060408385031215610e2157600080fd5b823567ffffffffffffffff80821115610e3957600080fd5b610e4586838701610c74565b93506020850135915080821115610e5b57600080fd5b50610d7985828601610c74565b60008060008060008060c08789031215610e8157600080fd5b863567ffffffffffffffff80821115610e9957600080fd5b610ea58a838b01610c74565b97506020890135915080821115610ebb57600080fd5b610ec78a838b01610c74565b96506040890135915080821115610edd57600080fd5b610ee98a838b01610c74565b95506060890135915080821115610eff57600080fd5b610f0b8a838b01610c74565b94506080890135915080821115610f2157600080fd5b50610f2e89828a01610c74565b925050610f3d60a08801610c58565b90509295509295509295565b600080600080600080600060e0888a031215610f6457600080fd5b873567ffffffffffffffff80821115610f7c57600080fd5b610f888b838c01610c74565b985060208a0135915080821115610f9e57600080fd5b610faa8b838c01610c74565b975060408a0135915080821115610fc057600080fd5b610fcc8b838c01610c74565b965060608a0135915080821115610fe257600080fd5b610fee8b838c01610c74565b955060808a013591508082111561100457600080fd5b6110108b838c01610c74565b945060a08a013591508082111561102657600080fd5b506110338a828b01610c74565b92505061104260c08901610c58565b905092959891949750929550565b6000806040838503121561106357600080fd5b823567ffffffffffffffff81111561107a57600080fd5b61108685828601610c74565b95602094909401359450505050565b600081518084526110ad8160208601602086016111d0565b601f01601f19169290920160200192915050565b600082516110d38184602087016111d0565b9190910192915050565b600087516110ef818460208c016111d0565b875190830190611103818360208c016111d0565b8751910190611116818360208b016111d0565b8651910190611129818360208a016111d0565b855191019061113c8183602089016111d0565b60609490941b6bffffffffffffffffffffffff19169301928352505060140195945050505050565b6020815260006106476020830184611095565b60408152600061118a6040830185611095565b905060018060a01b03831660208301529392505050565b600082198211156111b4576111b4611276565b500190565b6000828210156111cb576111cb611276565b500390565b60005b838110156111eb5781810151838201526020016111d3565b838111156111fa576000848401525b50505050565b600181811c9082168061121457607f821691505b6020821081141561123557634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561124f5761124f611276565b5060010190565b600060ff821660ff81141561126d5761126d611276565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122079c577a07207e47478d952d6db2576049748ba6e435e050c5338ece19f95a8ec64736f6c63430008070033",
}

// DocumentABI is the input ABI used to generate the binding from.
// Deprecated: Use DocumentMetaData.ABI instead.
var DocumentABI = DocumentMetaData.ABI

// Deprecated: Use DocumentMetaData.Sigs instead.
// DocumentFuncSigs maps the 4-byte function signature to its string representation.
var DocumentFuncSigs = DocumentMetaData.Sigs

// DocumentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DocumentMetaData.Bin instead.
var DocumentBin = DocumentMetaData.Bin

// DeployDocument deploys a new Ethereum contract, binding an instance of Document to it.
func DeployDocument(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Document, error) {
	parsed, err := DocumentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DocumentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Document{DocumentCaller: DocumentCaller{contract: contract}, DocumentTransactor: DocumentTransactor{contract: contract}, DocumentFilterer: DocumentFilterer{contract: contract}}, nil
}

// Document is an auto generated Go binding around an Ethereum contract.
type Document struct {
	DocumentCaller     // Read-only binding to the contract
	DocumentTransactor // Write-only binding to the contract
	DocumentFilterer   // Log filterer for contract events
}

// DocumentCaller is an auto generated read-only Go binding around an Ethereum contract.
type DocumentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DocumentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DocumentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DocumentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DocumentSession struct {
	Contract     *Document         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DocumentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DocumentCallerSession struct {
	Contract *DocumentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DocumentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DocumentTransactorSession struct {
	Contract     *DocumentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DocumentRaw is an auto generated low-level Go binding around an Ethereum contract.
type DocumentRaw struct {
	Contract *Document // Generic contract binding to access the raw methods on
}

// DocumentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DocumentCallerRaw struct {
	Contract *DocumentCaller // Generic read-only contract binding to access the raw methods on
}

// DocumentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DocumentTransactorRaw struct {
	Contract *DocumentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDocument creates a new instance of Document, bound to a specific deployed contract.
func NewDocument(address common.Address, backend bind.ContractBackend) (*Document, error) {
	contract, err := bindDocument(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Document{DocumentCaller: DocumentCaller{contract: contract}, DocumentTransactor: DocumentTransactor{contract: contract}, DocumentFilterer: DocumentFilterer{contract: contract}}, nil
}

// NewDocumentCaller creates a new read-only instance of Document, bound to a specific deployed contract.
func NewDocumentCaller(address common.Address, caller bind.ContractCaller) (*DocumentCaller, error) {
	contract, err := bindDocument(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DocumentCaller{contract: contract}, nil
}

// NewDocumentTransactor creates a new write-only instance of Document, bound to a specific deployed contract.
func NewDocumentTransactor(address common.Address, transactor bind.ContractTransactor) (*DocumentTransactor, error) {
	contract, err := bindDocument(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DocumentTransactor{contract: contract}, nil
}

// NewDocumentFilterer creates a new log filterer instance of Document, bound to a specific deployed contract.
func NewDocumentFilterer(address common.Address, filterer bind.ContractFilterer) (*DocumentFilterer, error) {
	contract, err := bindDocument(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DocumentFilterer{contract: contract}, nil
}

// bindDocument binds a generic wrapper to an already deployed contract.
func bindDocument(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DocumentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Document *DocumentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Document.Contract.DocumentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Document *DocumentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Document.Contract.DocumentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Document *DocumentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Document.Contract.DocumentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Document *DocumentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Document.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Document *DocumentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Document.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Document *DocumentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Document.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Document *DocumentCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Document *DocumentSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Document.Contract.Bytes32ToString(&_Document.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Document *DocumentCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Document.Contract.Bytes32ToString(&_Document.CallOpts, _bytes32)
}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Document *DocumentCaller) CompareBytes(opts *bind.CallOpts, a [32]byte, b [32]byte) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "compareBytes", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Document *DocumentSession) CompareBytes(a [32]byte, b [32]byte) (bool, error) {
	return _Document.Contract.CompareBytes(&_Document.CallOpts, a, b)
}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Document *DocumentCallerSession) CompareBytes(a [32]byte, b [32]byte) (bool, error) {
	return _Document.Contract.CompareBytes(&_Document.CallOpts, a, b)
}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Document *DocumentCaller) CompareString(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "compareString", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Document *DocumentSession) CompareString(a string, b string) (bool, error) {
	return _Document.Contract.CompareString(&_Document.CallOpts, a, b)
}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Document *DocumentCallerSession) CompareString(a string, b string) (bool, error) {
	return _Document.Contract.CompareString(&_Document.CallOpts, a, b)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Document *DocumentCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Document *DocumentSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Document.Contract.GetEthSignedMessageHash(&_Document.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Document *DocumentCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Document.Contract.GetEthSignedMessageHash(&_Document.CallOpts, _messageHash)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Document *DocumentCaller) GetHashUserInfo(opts *bind.CallOpts, phone string) ([32]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getHashUserInfo", phone)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Document *DocumentSession) GetHashUserInfo(phone string) ([32]byte, error) {
	return _Document.Contract.GetHashUserInfo(&_Document.CallOpts, phone)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Document *DocumentCallerSession) GetHashUserInfo(phone string) ([32]byte, error) {
	return _Document.Contract.GetHashUserInfo(&_Document.CallOpts, phone)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Document *DocumentCaller) GetSignature(opts *bind.CallOpts, phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getSignature", phone, indexDoc)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Document *DocumentSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	return _Document.Contract.GetSignature(&_Document.CallOpts, phone, indexDoc)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Document *DocumentCallerSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	return _Document.Contract.GetSignature(&_Document.CallOpts, phone, indexDoc)
}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentCaller) GetSinger(opts *bind.CallOpts, phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getSinger", phone, digest, indexDoc)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentSession) GetSinger(phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	return _Document.Contract.GetSinger(&_Document.CallOpts, phone, digest, indexDoc)
}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentCallerSession) GetSinger(phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	return _Document.Contract.GetSinger(&_Document.CallOpts, phone, digest, indexDoc)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Document *DocumentCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Document *DocumentSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Document.Contract.RecoverSigner(&_Document.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Document *DocumentCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Document.Contract.RecoverSigner(&_Document.CallOpts, _ethSignedMessageHash, _signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Document *DocumentCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Document *DocumentSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Document.Contract.SplitSignature(&_Document.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Document *DocumentCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Document.Contract.SplitSignature(&_Document.CallOpts, sig)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string phone, bytes32 digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentCaller) VerifyDoc(opts *bind.CallOpts, phone string, digest [32]byte, indexDoc *big.Int) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "verifyDoc", phone, digest, indexDoc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDoc is a free data retrieval call binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string phone, bytes32 digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentSession) VerifyDoc(phone string, digest [32]byte, indexDoc *big.Int) (bool, error) {
	return _Document.Contract.VerifyDoc(&_Document.CallOpts, phone, digest, indexDoc)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string phone, bytes32 digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentCallerSession) VerifyDoc(phone string, digest [32]byte, indexDoc *big.Int) (bool, error) {
	return _Document.Contract.VerifyDoc(&_Document.CallOpts, phone, digest, indexDoc)
}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Document *DocumentCaller) VerifyUser(opts *bind.CallOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "verifyUser", name, cmnd, dateOB, phone, gmail, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Document *DocumentSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	return _Document.Contract.VerifyUser(&_Document.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Document *DocumentCallerSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	return _Document.Contract.VerifyUser(&_Document.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string phone, bytes signature) returns()
func (_Document *DocumentTransactor) SaveDoc(opts *bind.TransactOpts, phone string, signature []byte) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "saveDoc", phone, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string phone, bytes signature) returns()
func (_Document *DocumentSession) SaveDoc(phone string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.SaveDoc(&_Document.TransactOpts, phone, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string phone, bytes signature) returns()
func (_Document *DocumentTransactorSession) SaveDoc(phone string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.SaveDoc(&_Document.TransactOpts, phone, signature)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Document *DocumentTransactor) StoreUser(opts *bind.TransactOpts, userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "storeUser", userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Document *DocumentSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Document *DocumentTransactorSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// DocumentIndexDocumentIterator is returned from FilterIndexDocument and is used to iterate over the raw logs and unpacked data for IndexDocument events raised by the Document contract.
type DocumentIndexDocumentIterator struct {
	Event *DocumentIndexDocument // Event containing the contract specifics and raw log

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
func (it *DocumentIndexDocumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DocumentIndexDocument)
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
		it.Event = new(DocumentIndexDocument)
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
func (it *DocumentIndexDocumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DocumentIndexDocumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DocumentIndexDocument represents a IndexDocument event raised by the Document contract.
type DocumentIndexDocument struct {
	Numdoc *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIndexDocument is a free log retrieval operation binding the contract event 0xca1bddf79d7a20af180f6d50bd9b30dd13435374d6c3e6d35e319f24eee11ba0.
//
// Solidity: event IndexDocument(uint256 numdoc)
func (_Document *DocumentFilterer) FilterIndexDocument(opts *bind.FilterOpts) (*DocumentIndexDocumentIterator, error) {

	logs, sub, err := _Document.contract.FilterLogs(opts, "IndexDocument")
	if err != nil {
		return nil, err
	}
	return &DocumentIndexDocumentIterator{contract: _Document.contract, event: "IndexDocument", logs: logs, sub: sub}, nil
}

// WatchIndexDocument is a free log subscription operation binding the contract event 0xca1bddf79d7a20af180f6d50bd9b30dd13435374d6c3e6d35e319f24eee11ba0.
//
// Solidity: event IndexDocument(uint256 numdoc)
func (_Document *DocumentFilterer) WatchIndexDocument(opts *bind.WatchOpts, sink chan<- *DocumentIndexDocument) (event.Subscription, error) {

	logs, sub, err := _Document.contract.WatchLogs(opts, "IndexDocument")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DocumentIndexDocument)
				if err := _Document.contract.UnpackLog(event, "IndexDocument", log); err != nil {
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

// ParseIndexDocument is a log parse operation binding the contract event 0xca1bddf79d7a20af180f6d50bd9b30dd13435374d6c3e6d35e319f24eee11ba0.
//
// Solidity: event IndexDocument(uint256 numdoc)
func (_Document *DocumentFilterer) ParseIndexDocument(log types.Log) (*DocumentIndexDocument, error) {
	event := new(DocumentIndexDocument)
	if err := _Document.contract.UnpackLog(event, "IndexDocument", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DocumentStatusIterator is returned from FilterStatus and is used to iterate over the raw logs and unpacked data for Status events raised by the Document contract.
type DocumentStatusIterator struct {
	Event *DocumentStatus // Event containing the contract specifics and raw log

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
func (it *DocumentStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DocumentStatus)
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
		it.Event = new(DocumentStatus)
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
func (it *DocumentStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DocumentStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DocumentStatus represents a Status event raised by the Document contract.
type DocumentStatus struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStatus is a free log retrieval operation binding the contract event 0x2844c95bf1b4598da931d527f903501abc60fe0199c65da52d5ce818c6c5e961.
//
// Solidity: event Status(string arg0)
func (_Document *DocumentFilterer) FilterStatus(opts *bind.FilterOpts) (*DocumentStatusIterator, error) {

	logs, sub, err := _Document.contract.FilterLogs(opts, "Status")
	if err != nil {
		return nil, err
	}
	return &DocumentStatusIterator{contract: _Document.contract, event: "Status", logs: logs, sub: sub}, nil
}

// WatchStatus is a free log subscription operation binding the contract event 0x2844c95bf1b4598da931d527f903501abc60fe0199c65da52d5ce818c6c5e961.
//
// Solidity: event Status(string arg0)
func (_Document *DocumentFilterer) WatchStatus(opts *bind.WatchOpts, sink chan<- *DocumentStatus) (event.Subscription, error) {

	logs, sub, err := _Document.contract.WatchLogs(opts, "Status")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DocumentStatus)
				if err := _Document.contract.UnpackLog(event, "Status", log); err != nil {
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

// ParseStatus is a log parse operation binding the contract event 0x2844c95bf1b4598da931d527f903501abc60fe0199c65da52d5ce818c6c5e961.
//
// Solidity: event Status(string arg0)
func (_Document *DocumentFilterer) ParseStatus(log types.Log) (*DocumentStatus, error) {
	event := new(DocumentStatus)
	if err := _Document.contract.UnpackLog(event, "Status", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDCMetaData contains all meta data concerning the IDC contract.
var IDCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3a8fb351": "saveDoc(string,bytes)",
		"8e16afb2": "storeUser(string,string,string,string,string,string,address)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
		"a88c6e45": "verifyUser(string,string,string,string,string,address)",
	},
}

// IDCABI is the input ABI used to generate the binding from.
// Deprecated: Use IDCMetaData.ABI instead.
var IDCABI = IDCMetaData.ABI

// Deprecated: Use IDCMetaData.Sigs instead.
// IDCFuncSigs maps the 4-byte function signature to its string representation.
var IDCFuncSigs = IDCMetaData.Sigs

// IDC is an auto generated Go binding around an Ethereum contract.
type IDC struct {
	IDCCaller     // Read-only binding to the contract
	IDCTransactor // Write-only binding to the contract
	IDCFilterer   // Log filterer for contract events
}

// IDCCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDCSession struct {
	Contract     *IDC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDCCallerSession struct {
	Contract *IDCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IDCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDCTransactorSession struct {
	Contract     *IDCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDCRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDCRaw struct {
	Contract *IDC // Generic contract binding to access the raw methods on
}

// IDCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDCCallerRaw struct {
	Contract *IDCCaller // Generic read-only contract binding to access the raw methods on
}

// IDCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDCTransactorRaw struct {
	Contract *IDCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDC creates a new instance of IDC, bound to a specific deployed contract.
func NewIDC(address common.Address, backend bind.ContractBackend) (*IDC, error) {
	contract, err := bindIDC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDC{IDCCaller: IDCCaller{contract: contract}, IDCTransactor: IDCTransactor{contract: contract}, IDCFilterer: IDCFilterer{contract: contract}}, nil
}

// NewIDCCaller creates a new read-only instance of IDC, bound to a specific deployed contract.
func NewIDCCaller(address common.Address, caller bind.ContractCaller) (*IDCCaller, error) {
	contract, err := bindIDC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDCCaller{contract: contract}, nil
}

// NewIDCTransactor creates a new write-only instance of IDC, bound to a specific deployed contract.
func NewIDCTransactor(address common.Address, transactor bind.ContractTransactor) (*IDCTransactor, error) {
	contract, err := bindIDC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDCTransactor{contract: contract}, nil
}

// NewIDCFilterer creates a new log filterer instance of IDC, bound to a specific deployed contract.
func NewIDCFilterer(address common.Address, filterer bind.ContractFilterer) (*IDCFilterer, error) {
	contract, err := bindIDC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDCFilterer{contract: contract}, nil
}

// bindIDC binds a generic wrapper to an already deployed contract.
func bindIDC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDC *IDCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDC.Contract.IDCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDC *IDCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDC.Contract.IDCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDC *IDCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDC.Contract.IDCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDC *IDCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDC *IDCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDC *IDCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDC.Contract.contract.Transact(opts, method, params...)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns()
func (_IDC *IDCTransactor) SaveDoc(opts *bind.TransactOpts, userID string, signature []byte) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "saveDoc", userID, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns()
func (_IDC *IDCSession) SaveDoc(userID string, signature []byte) (*types.Transaction, error) {
	return _IDC.Contract.SaveDoc(&_IDC.TransactOpts, userID, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns()
func (_IDC *IDCTransactorSession) SaveDoc(userID string, signature []byte) (*types.Transaction, error) {
	return _IDC.Contract.SaveDoc(&_IDC.TransactOpts, userID, signature)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_IDC *IDCTransactor) StoreUser(opts *bind.TransactOpts, userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "storeUser", userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_IDC *IDCSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.StoreUser(&_IDC.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_IDC *IDCTransactorSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.StoreUser(&_IDC.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyDoc is a paid mutator transaction binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string userID, bytes32 digest, uint256 indexDoc) returns(bool)
func (_IDC *IDCTransactor) VerifyDoc(opts *bind.TransactOpts, userID string, digest [32]byte, indexDoc *big.Int) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "verifyDoc", userID, digest, indexDoc)
}

// VerifyDoc is a paid mutator transaction binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string userID, bytes32 digest, uint256 indexDoc) returns(bool)
func (_IDC *IDCSession) VerifyDoc(userID string, digest [32]byte, indexDoc *big.Int) (*types.Transaction, error) {
	return _IDC.Contract.VerifyDoc(&_IDC.TransactOpts, userID, digest, indexDoc)
}

// VerifyDoc is a paid mutator transaction binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string userID, bytes32 digest, uint256 indexDoc) returns(bool)
func (_IDC *IDCTransactorSession) VerifyDoc(userID string, digest [32]byte, indexDoc *big.Int) (*types.Transaction, error) {
	return _IDC.Contract.VerifyDoc(&_IDC.TransactOpts, userID, digest, indexDoc)
}

// VerifyUser is a paid mutator transaction binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_IDC *IDCTransactor) VerifyUser(opts *bind.TransactOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "verifyUser", name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a paid mutator transaction binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_IDC *IDCSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.VerifyUser(&_IDC.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a paid mutator transaction binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_IDC *IDCTransactorSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.VerifyUser(&_IDC.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StringUtilsMetaData contains all meta data concerning the StringUtils contract.
var StringUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Sigs: map[string]string{
		"3a96fdd7": "compare(string,string)",
		"46bdca9a": "equal(string,string)",
		"8a0807b7": "indexOf(string,string)",
	},
	Bin: "0x61051d61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80633a96fdd71461005057806346bdca9a146100835780638a0807b7146100b3575b600080fd5b81801561005c57600080fd5b5061007061006b36600461040e565b6100d3565b6040519081526020015b60405180910390f35b81801561008f57600080fd5b506100a361009e36600461040e565b6101fa565b604051901515815260200161007a565b8180156100bf57600080fd5b506100706100ce36600461040e565b61020e565b8151815160009184918491908111156100ea575080515b60005b818110156101be57828181518110610107576101076104bb565b602001015160f81c60f81b6001600160f81b03191684828151811061012e5761012e6104bb565b01602001516001600160f81b0319161015610151576000199450505050506101f4565b828181518110610163576101636104bb565b602001015160f81c60f81b6001600160f81b03191684828151811061018a5761018a6104bb565b01602001516001600160f81b03191611156101ac5760019450505050506101f4565b806101b68161048a565b9150506100ed565b508151835110156101d65760001993505050506101f4565b8151835111156101ec57600193505050506101f4565b600093505050505b92915050565b600061020683836100d3565b159392505050565b81516000908390839060011180610226575060018151105b80610232575081518151115b1561024357600019925050506101f4565b6fffffffffffffffffffffffffffffffff8251111561026857600019925050506101f4565b6000805b83518110156103735782600081518110610288576102886104bb565b602001015160f81c60f81b6001600160f81b0319168482815181106102af576102af6104bb565b01602001516001600160f81b031916141561036157600191505b8251821080156102e2575083516102e08383610472565b105b801561033657508282815181106102fb576102fb6104bb565b01602001516001600160f81b031916846103158484610472565b81518110610325576103256104bb565b01602001516001600160f81b031916145b1561034d57816103458161048a565b9250506102c9565b82518214156103615793506101f492505050565b8061036b8161048a565b91505061026c565b5060001993505050506101f4565b600082601f83011261039257600080fd5b813567ffffffffffffffff808211156103ad576103ad6104d1565b604051601f8301601f19908116603f011681019082821181831017156103d5576103d56104d1565b816040528381528660208588010111156103ee57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561042157600080fd5b823567ffffffffffffffff8082111561043957600080fd5b61044586838701610381565b9350602085013591508082111561045b57600080fd5b5061046885828601610381565b9150509250929050565b60008219821115610485576104856104a5565b500190565b600060001982141561049e5761049e6104a5565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220c44f61b9adc7ccfb844e5bd4c738e1dcda38bb3f022fc659fec415b4c6ffcd2564736f6c63430008070033",
}

// StringUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringUtilsMetaData.ABI instead.
var StringUtilsABI = StringUtilsMetaData.ABI

// Deprecated: Use StringUtilsMetaData.Sigs instead.
// StringUtilsFuncSigs maps the 4-byte function signature to its string representation.
var StringUtilsFuncSigs = StringUtilsMetaData.Sigs

// StringUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringUtilsMetaData.Bin instead.
var StringUtilsBin = StringUtilsMetaData.Bin

// DeployStringUtils deploys a new Ethereum contract, binding an instance of StringUtils to it.
func DeployStringUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StringUtils, error) {
	parsed, err := StringUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StringUtils{StringUtilsCaller: StringUtilsCaller{contract: contract}, StringUtilsTransactor: StringUtilsTransactor{contract: contract}, StringUtilsFilterer: StringUtilsFilterer{contract: contract}}, nil
}

// StringUtils is an auto generated Go binding around an Ethereum contract.
type StringUtils struct {
	StringUtilsCaller     // Read-only binding to the contract
	StringUtilsTransactor // Write-only binding to the contract
	StringUtilsFilterer   // Log filterer for contract events
}

// StringUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringUtilsSession struct {
	Contract     *StringUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringUtilsCallerSession struct {
	Contract *StringUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StringUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringUtilsTransactorSession struct {
	Contract     *StringUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StringUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringUtilsRaw struct {
	Contract *StringUtils // Generic contract binding to access the raw methods on
}

// StringUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringUtilsCallerRaw struct {
	Contract *StringUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// StringUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringUtilsTransactorRaw struct {
	Contract *StringUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStringUtils creates a new instance of StringUtils, bound to a specific deployed contract.
func NewStringUtils(address common.Address, backend bind.ContractBackend) (*StringUtils, error) {
	contract, err := bindStringUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StringUtils{StringUtilsCaller: StringUtilsCaller{contract: contract}, StringUtilsTransactor: StringUtilsTransactor{contract: contract}, StringUtilsFilterer: StringUtilsFilterer{contract: contract}}, nil
}

// NewStringUtilsCaller creates a new read-only instance of StringUtils, bound to a specific deployed contract.
func NewStringUtilsCaller(address common.Address, caller bind.ContractCaller) (*StringUtilsCaller, error) {
	contract, err := bindStringUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringUtilsCaller{contract: contract}, nil
}

// NewStringUtilsTransactor creates a new write-only instance of StringUtils, bound to a specific deployed contract.
func NewStringUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringUtilsTransactor, error) {
	contract, err := bindStringUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringUtilsTransactor{contract: contract}, nil
}

// NewStringUtilsFilterer creates a new log filterer instance of StringUtils, bound to a specific deployed contract.
func NewStringUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringUtilsFilterer, error) {
	contract, err := bindStringUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringUtilsFilterer{contract: contract}, nil
}

// bindStringUtils binds a generic wrapper to an already deployed contract.
func bindStringUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StringUtils *StringUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StringUtils.Contract.StringUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StringUtils *StringUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StringUtils.Contract.StringUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StringUtils *StringUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StringUtils.Contract.StringUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StringUtils *StringUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StringUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StringUtils *StringUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StringUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StringUtils *StringUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StringUtils.Contract.contract.Transact(opts, method, params...)
}
