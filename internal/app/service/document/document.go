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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publickey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numdoc\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"IndexDocument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Status\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"infoHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publickey\",\"type\":\"string\"}],\"name\":\"StoreUserStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"compareBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"name\":\"getHashUserInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9201de55": "bytes32ToString(bytes32)",
		"ef1bade3": "compareBytes(bytes32,bytes32)",
		"a23f1228": "getHashUserInfo(string)",
		"cfd6baa2": "getSignature(string,uint256)",
		"3a8fb351": "saveDoc(string,bytes)",
		"a01ef4b9": "storeUser(string,string,string,string,string,string,string)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
		"7ecde66f": "verifyUser(string,string,string,string,string,string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611489806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a23f12281161005b578063a23f122814610109578063cfd6baa21461011c578063d2de39a81461012f578063ef1bade31461014257600080fd5b80633a8fb3511461008d5780637ecde66f146100b35780639201de55146100d6578063a01ef4b9146100f6575b600080fd5b6100a061009b366004610e16565b610155565b6040519081526020015b60405180910390f35b6100c66100c1366004610e8e565b6103ec565b60405190151581526020016100aa565b6100e96100e4366004610d50565b61046c565b6040516100aa9190611293565b6100a0610104366004610f83565b6105a3565b6100a0610117366004610d8b565b61063a565b6100e961012a36600461109e565b6106e3565b6100c661013d366004610dc8565b610897565b6100c6610150366004610d69565b610a6b565b60008061016184610abf565b905060005b60005481116103ac5761019d6000828154811061018557610185611427565b90600052602060002090600602016002015483610a6b565b1561039a5783600082815481106101b6576101b6611427565b906000526020600020906006020160050160008084815481106101db576101db611427565b90600052602060002090600602016004015481526020019081526020016000206000019080519060200190610211929190610c21565b506000818154811061022557610225611427565b9060005260206000209060060201600401546001610243919061133c565b6000828154811061025657610256611427565b9060005260206000209060060201600401819055507f4b6f2d7d86556d3bffa2a6dfd4d4752578724e9907eb7344f072896ca1fe6ce56000828154811061029f5761029f611427565b90600052602060002090600602016003016001600084815481106102c5576102c5611427565b9060005260206000209060060201600401546102e19190611354565b600084815481106102f4576102f4611427565b9060005260206000209060060201600501600060016000878154811061031c5761031c611427565b9060005260206000209060060201600401546103389190611354565b8152602001908152602001600020600001604051610358939291906112a6565b60405180910390a160016000828154811061037557610375611427565b9060005260206000209060060201600401546103919190611354565b925050506103e6565b806103a4816113d6565b915050610166565b5060405162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b60448201526064015b60405180910390fd5b92915050565b6000806103fd888888888888610aef565b905060005b600054811161045b576104396000828154811061042157610421611427565b90600052602060002090600602016001015483610a6b565b1561044957600192505050610462565b80610453816113d6565b915050610402565b6000925050505b9695505050505050565b606060005b60208160ff161080156104a55750828160ff166020811061049457610494611427565b1a60f81b6001600160f81b03191615155b156104bc57806104b4816113f1565b915050610471565b60008160ff1667ffffffffffffffff8111156104da576104da61143d565b6040519080825280601f01601f191660200182016040528015610504576020820181803683370190505b509050600091505b60208260ff161080156105405750838260ff166020811061052f5761052f611427565b1a60f81b6001600160f81b03191615155b1561059c57838260ff166020811061055a5761055a611427565b1a60f81b818360ff168151811061057357610573611427565b60200101906001600160f81b031916908160001a90535081610594816113f1565b92505061050c565b9392505050565b6000806105b4888888888888610aef565b6000805460018101808355828052929350918190839081106105d8576105d8611427565b60009182526020918290208d5160069092020192506105fc918391908e0190610c21565b506001810183905584516106199060038301906020880190610c21565b5061062387610abf565b600290910181905592505050979650505050505050565b60008061064683610abf565b905060005b60005481116106af5761066a6000828154811061018557610185611427565b1561069d576000818154811061068257610682611427565b90600052602060002090600602016001015492505050919050565b806106a7816113d6565b91505061064b565b60405162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b60448201526064016103dd565b60606000610716846040516020016106fb9190611186565b6040516020818303038152906040528051906020012061046c565b90506000805b6000548110156107b657826040516020016107379190611186565b604051602081830303815290604052805190602001206000828154811061076057610760611427565b90600052602060002090600602016002015460405160200161078491815260200190565b6040516020818303038152906040528051906020012014156107a4578091505b806107ae816113d6565b91505061071c565b5060008082815481106107cb576107cb611427565b906000526020600020906006020160050160008681526020019081526020016000206040518060200160405290816000820180546108089061139b565b80601f01602080910402602001604051908101604052809291908181526020018280546108349061139b565b80156108815780601f1061085657610100808354040283529160200191610881565b820191906000526020600020905b81548152906001019060200180831161086457829003601f168201915b5050509190925250509051979650505050505050565b6000806108a385610abf565b905064174876e80060005b6000548111610900576108e5600082815481106108cd576108cd611427565b90600052602060002090600602016002015484610a6b565b156108ee578091505b806108f8816113d6565b9150506108ae565b508064174876e800141561091357600080fd5b6000818154811061092657610926611427565b90600052602060002090600602016003016040516020016109479190611221565b60405160208183030381529060405280519060200120610a21866000848154811061097457610974611427565b600091825260208083208a8452600560069093020191909101905260409020805461099e9061139b565b80601f01602080910402602001604051908101604052809291908181526020018280546109ca9061139b565b8015610a175780601f106109ec57610100808354040283529160200191610a17565b820191906000526020600020905b8154815290600101906020018083116109fa57829003601f168201915b5050505050610b2e565b604051602001610a49919060609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528051906020012014925050509392505050565b600081604051602001610a8091815260200190565b60408051601f19818403018152828252805160209182012090830186905291016040516020818303038152906040528051906020012014905092915050565b600081604051602001610ad29190611186565b604051602081830303815290604052805190602001209050919050565b6000868686868686604051602001610b0c969594939291906111a2565b6040516020818303038152906040528051906020012090509695505050505050565b600080600080610b3d85610bad565b6040805160008152602081018083528b905260ff8316918101919091526060810184905260808101839052929550909350915060019060a0016020604051602081039080840390855afa158015610b98573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008060008351604114610c035760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e677468000000000000000060448201526064016103dd565b50505060208101516040820151606090920151909260009190911a90565b828054610c2d9061139b565b90600052602060002090601f016020900481019282610c4f5760008555610c95565b82601f10610c6857805160ff1916838001178555610c95565b82800160010185558215610c95579182015b82811115610c95578251825591602001919060010190610c7a565b50610ca1929150610ca5565b5090565b5b80821115610ca15760008155600101610ca6565b600067ffffffffffffffff80841115610cd557610cd561143d565b604051601f8501601f19908116603f01168101908282118183101715610cfd57610cfd61143d565b81604052809350858152868686011115610d1657600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112610d4157600080fd5b61059c83833560208501610cba565b600060208284031215610d6257600080fd5b5035919050565b60008060408385031215610d7c57600080fd5b50508035926020909101359150565b600060208284031215610d9d57600080fd5b813567ffffffffffffffff811115610db457600080fd5b610dc084828501610d30565b949350505050565b600080600060608486031215610ddd57600080fd5b833567ffffffffffffffff811115610df457600080fd5b610e0086828701610d30565b9660208601359650604090950135949350505050565b60008060408385031215610e2957600080fd5b823567ffffffffffffffff80821115610e4157600080fd5b610e4d86838701610d30565b93506020850135915080821115610e6357600080fd5b508301601f81018513610e7557600080fd5b610e8485823560208401610cba565b9150509250929050565b60008060008060008060c08789031215610ea757600080fd5b863567ffffffffffffffff80821115610ebf57600080fd5b610ecb8a838b01610d30565b97506020890135915080821115610ee157600080fd5b610eed8a838b01610d30565b96506040890135915080821115610f0357600080fd5b610f0f8a838b01610d30565b95506060890135915080821115610f2557600080fd5b610f318a838b01610d30565b94506080890135915080821115610f4757600080fd5b610f538a838b01610d30565b935060a0890135915080821115610f6957600080fd5b50610f7689828a01610d30565b9150509295509295509295565b600080600080600080600060e0888a031215610f9e57600080fd5b873567ffffffffffffffff80821115610fb657600080fd5b610fc28b838c01610d30565b985060208a0135915080821115610fd857600080fd5b610fe48b838c01610d30565b975060408a0135915080821115610ffa57600080fd5b6110068b838c01610d30565b965060608a013591508082111561101c57600080fd5b6110288b838c01610d30565b955060808a013591508082111561103e57600080fd5b61104a8b838c01610d30565b945060a08a013591508082111561106057600080fd5b61106c8b838c01610d30565b935060c08a013591508082111561108257600080fd5b5061108f8a828b01610d30565b91505092959891949750929550565b600080604083850312156110b157600080fd5b823567ffffffffffffffff8111156110c857600080fd5b6110d485828601610d30565b95602094909401359450505050565b600081518084526110fb81602086016020860161136b565b601f01601f19169290920160200192915050565b6000815461111c8161139b565b808552602060018381168015611139576001811461114d5761117b565b60ff1985168884015260408801955061117b565b866000528260002060005b858110156111735781548a8201860152908301908401611158565b890184019650505b505050505092915050565b6000825161119881846020870161136b565b9190910192915050565b6000875160206111b58285838d0161136b565b8851918401916111c88184848d0161136b565b88519201916111da8184848c0161136b565b87519201916111ec8184848b0161136b565b86519201916111fe8184848a0161136b565b8551920191611210818484890161136b565b919091019998505050505050505050565b600080835461122f8161139b565b60018281168015611247576001811461125857611287565b60ff19841687528287019450611287565b8760005260208060002060005b8581101561127e5781548a820152908401908201611265565b50505082870194505b50929695505050505050565b60208152600061059c60208301846110e3565b6060815260008085546112b88161139b565b80606086015260806001808416600081146112da57600181146112ee5761131f565b60ff1985168884015260a08801955061131f565b8a60005260208060002060005b868110156113165781548b82018701529084019082016112fb565b8a018501975050505b50505050508460208401528281036040840152610462818561110f565b6000821982111561134f5761134f611411565b500190565b60008282101561136657611366611411565b500390565b60005b8381101561138657818101518382015260200161136e565b83811115611395576000848401525b50505050565b600181811c908216806113af57607f821691505b602082108114156113d057634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156113ea576113ea611411565b5060010190565b600060ff821660ff81141561140857611408611411565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220cd296020a943bbd96f97c27c364f02f4d74d6979d6ce567a1f0fbeac35c700ad64736f6c63430008070033",
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
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentCaller) GetSignature(opts *bind.CallOpts, phone string, indexDoc *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getSignature", phone, indexDoc)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, error) {
	return _Document.Contract.GetSignature(&_Document.CallOpts, phone, indexDoc)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes)
func (_Document *DocumentCallerSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, error) {
	return _Document.Contract.GetSignature(&_Document.CallOpts, phone, indexDoc)
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

// VerifyUser is a free data retrieval call binding the contract method 0x7ecde66f.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) view returns(bool)
func (_Document *DocumentCaller) VerifyUser(opts *bind.CallOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "verifyUser", name, cmnd, dateOB, phone, gmail, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUser is a free data retrieval call binding the contract method 0x7ecde66f.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) view returns(bool)
func (_Document *DocumentSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (bool, error) {
	return _Document.Contract.VerifyUser(&_Document.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a free data retrieval call binding the contract method 0x7ecde66f.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) view returns(bool)
func (_Document *DocumentCallerSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (bool, error) {
	return _Document.Contract.VerifyUser(&_Document.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string phone, bytes signature) returns(uint256)
func (_Document *DocumentTransactor) SaveDoc(opts *bind.TransactOpts, phone string, signature []byte) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "saveDoc", phone, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string phone, bytes signature) returns(uint256)
func (_Document *DocumentSession) SaveDoc(phone string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.SaveDoc(&_Document.TransactOpts, phone, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string phone, bytes signature) returns(uint256)
func (_Document *DocumentTransactorSession) SaveDoc(phone string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.SaveDoc(&_Document.TransactOpts, phone, signature)
}

// StoreUser is a paid mutator transaction binding the contract method 0xa01ef4b9.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bytes32)
func (_Document *DocumentTransactor) StoreUser(opts *bind.TransactOpts, userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "storeUser", userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0xa01ef4b9.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bytes32)
func (_Document *DocumentSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0xa01ef4b9.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bytes32)
func (_Document *DocumentTransactorSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
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
	Publickey string
	Numdoc    *big.Int
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIndexDocument is a free log retrieval operation binding the contract event 0x4b6f2d7d86556d3bffa2a6dfd4d4752578724e9907eb7344f072896ca1fe6ce5.
//
// Solidity: event IndexDocument(string publickey, uint256 numdoc, bytes signature)
func (_Document *DocumentFilterer) FilterIndexDocument(opts *bind.FilterOpts) (*DocumentIndexDocumentIterator, error) {

	logs, sub, err := _Document.contract.FilterLogs(opts, "IndexDocument")
	if err != nil {
		return nil, err
	}
	return &DocumentIndexDocumentIterator{contract: _Document.contract, event: "IndexDocument", logs: logs, sub: sub}, nil
}

// WatchIndexDocument is a free log subscription operation binding the contract event 0x4b6f2d7d86556d3bffa2a6dfd4d4752578724e9907eb7344f072896ca1fe6ce5.
//
// Solidity: event IndexDocument(string publickey, uint256 numdoc, bytes signature)
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

// ParseIndexDocument is a log parse operation binding the contract event 0x4b6f2d7d86556d3bffa2a6dfd4d4752578724e9907eb7344f072896ca1fe6ce5.
//
// Solidity: event IndexDocument(string publickey, uint256 numdoc, bytes signature)
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

// DocumentStoreUserStatusIterator is returned from FilterStoreUserStatus and is used to iterate over the raw logs and unpacked data for StoreUserStatus events raised by the Document contract.
type DocumentStoreUserStatusIterator struct {
	Event *DocumentStoreUserStatus // Event containing the contract specifics and raw log

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
func (it *DocumentStoreUserStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DocumentStoreUserStatus)
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
		it.Event = new(DocumentStoreUserStatus)
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
func (it *DocumentStoreUserStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DocumentStoreUserStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DocumentStoreUserStatus represents a StoreUserStatus event raised by the Document contract.
type DocumentStoreUserStatus struct {
	InfoHash  [32]byte
	Publickey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStoreUserStatus is a free log retrieval operation binding the contract event 0x5a8eb3f4ce99516f1c41b07acada760ba5bcec300c07c1293f0137eb7f4be192.
//
// Solidity: event StoreUserStatus(bytes32 infoHash, string publickey)
func (_Document *DocumentFilterer) FilterStoreUserStatus(opts *bind.FilterOpts) (*DocumentStoreUserStatusIterator, error) {

	logs, sub, err := _Document.contract.FilterLogs(opts, "StoreUserStatus")
	if err != nil {
		return nil, err
	}
	return &DocumentStoreUserStatusIterator{contract: _Document.contract, event: "StoreUserStatus", logs: logs, sub: sub}, nil
}

// WatchStoreUserStatus is a free log subscription operation binding the contract event 0x5a8eb3f4ce99516f1c41b07acada760ba5bcec300c07c1293f0137eb7f4be192.
//
// Solidity: event StoreUserStatus(bytes32 infoHash, string publickey)
func (_Document *DocumentFilterer) WatchStoreUserStatus(opts *bind.WatchOpts, sink chan<- *DocumentStoreUserStatus) (event.Subscription, error) {

	logs, sub, err := _Document.contract.WatchLogs(opts, "StoreUserStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DocumentStoreUserStatus)
				if err := _Document.contract.UnpackLog(event, "StoreUserStatus", log); err != nil {
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

// ParseStoreUserStatus is a log parse operation binding the contract event 0x5a8eb3f4ce99516f1c41b07acada760ba5bcec300c07c1293f0137eb7f4be192.
//
// Solidity: event StoreUserStatus(bytes32 infoHash, string publickey)
func (_Document *DocumentFilterer) ParseStoreUserStatus(log types.Log) (*DocumentStoreUserStatus, error) {
	event := new(DocumentStoreUserStatus)
	if err := _Document.contract.UnpackLog(event, "StoreUserStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDCMetaData contains all meta data concerning the IDC contract.
var IDCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3a8fb351": "saveDoc(string,bytes)",
		"a01ef4b9": "storeUser(string,string,string,string,string,string,string)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
		"7ecde66f": "verifyUser(string,string,string,string,string,string)",
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
// Solidity: function saveDoc(string userID, bytes signature) returns(uint256)
func (_IDC *IDCTransactor) SaveDoc(opts *bind.TransactOpts, userID string, signature []byte) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "saveDoc", userID, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns(uint256)
func (_IDC *IDCSession) SaveDoc(userID string, signature []byte) (*types.Transaction, error) {
	return _IDC.Contract.SaveDoc(&_IDC.TransactOpts, userID, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns(uint256)
func (_IDC *IDCTransactorSession) SaveDoc(userID string, signature []byte) (*types.Transaction, error) {
	return _IDC.Contract.SaveDoc(&_IDC.TransactOpts, userID, signature)
}

// StoreUser is a paid mutator transaction binding the contract method 0xa01ef4b9.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bytes32)
func (_IDC *IDCTransactor) StoreUser(opts *bind.TransactOpts, userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "storeUser", userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0xa01ef4b9.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bytes32)
func (_IDC *IDCSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
	return _IDC.Contract.StoreUser(&_IDC.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0xa01ef4b9.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bytes32)
func (_IDC *IDCTransactorSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
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

// VerifyUser is a paid mutator transaction binding the contract method 0x7ecde66f.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bool)
func (_IDC *IDCTransactor) VerifyUser(opts *bind.TransactOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "verifyUser", name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x7ecde66f.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bool)
func (_IDC *IDCSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
	return _IDC.Contract.VerifyUser(&_IDC.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x7ecde66f.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, string publicKey) returns(bool)
func (_IDC *IDCTransactorSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey string) (*types.Transaction, error) {
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
