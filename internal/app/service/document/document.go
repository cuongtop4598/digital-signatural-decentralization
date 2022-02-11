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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"GetHashUserInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getUserID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d203f63c": "GetHashUserInfo(address)",
		"bff90821": "getPublicKey(string)",
		"0ef12aa9": "getUserID(address)",
		"3a8fb351": "saveDoc(string,bytes)",
		"8e16afb2": "storeUser(string,string,string,string,string,string,address)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610bd0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630ef12aa9146100675780633a8fb351146100905780638e16afb2146100b1578063bff90821146100d4578063d203f63c146100ff578063d2de39a81461012b575b600080fd5b61007a610075366004610811565b61013e565b6040516100879190610abd565b60405180910390f35b6100a361009e3660046108b7565b6101ea565b604051908152602001610087565b6100c46100bf36600461092f565b61029e565b6040519015158152602001610087565b6100e76100e236600461082c565b610321565b6040516001600160a01b039091168152602001610087565b6100a361010d366004610811565b6001600160a01b031660009081526020819052604090206001015490565b6100c4610139366004610869565b610357565b6001600160a01b038116600090815260208190526040902080546060919061016590610b20565b80601f016020809104026020016040519081016040528092919081815260200182805461019190610b20565b80156101de5780601f106101b3576101008083540402835291602001916101de565b820191906000526020600020905b8154815290600101906020018083116101c157829003601f168201915b50505050509050919050565b6000806000600101846040516102009190610a36565b9081526020016040518091039020600301549050808061021f90610b5b565b915050600080600101856040516102369190610a36565b9081526040805160209281900383019020600085815260049091018352208651909250610268918391908801906106bf565b50835161027e90600183019060208701906106bf565b50835161029490600283019060208701906106bf565b5090949350505050565b6000806102ae888888888861058c565b6001600160a01b0384166000908152602081815260409091208b51929350916102dc918391908d01906106bf565b5060018181018390556002820180546001600160a01b0319166001600160a01b038716179055604051610310908c90610a36565b525060019998505050505050505050565b600080600101826040516103359190610a36565b908152604051908190036020019020600201546001600160a01b031692915050565b60008060006001018560405161036d9190610a36565b9081526020016040518091039020905060008160040160008581526020019081526020016000206040518060600160405290816000820180546103af90610b20565b80601f01602080910402602001604051908101604052809291908181526020018280546103db90610b20565b80156104285780601f106103fd57610100808354040283529160200191610428565b820191906000526020600020905b81548152906001019060200180831161040b57829003601f168201915b5050505050815260200160018201805461044190610b20565b80601f016020809104026020016040519081016040528092919081815260200182805461046d90610b20565b80156104ba5780601f1061048f576101008083540402835291602001916104ba565b820191906000526020600020905b81548152906001019060200180831161049d57829003601f168201915b505050505081526020016002820180546104d390610b20565b80601f01602080910402602001604051908101604052809291908181526020018280546104ff90610b20565b801561054c5780601f106105215761010080835404028352916020019161054c565b820191906000526020600020905b81548152906001019060200180831161052f57829003601f168201915b50505091909252505050600283015460208201519192506001600160a01b0316906105789087906105c8565b6001600160a01b0316149695505050505050565b600085858585856040516020016105a7959493929190610a52565b60405160208183030381529060405280519060200120905095945050505050565b6000806000806105d785610647565b6040805160008152602081018083528b905260ff8316918101919091526060810184905260808101839052929550909350915060019060a0016020604051602081039080840390855afa158015610632573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b600080600083516041146106a15760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e6774680000000000000000604482015260640160405180910390fd5b50505060208101516040820151606090920151909260009190911a90565b8280546106cb90610b20565b90600052602060002090601f0160209004810192826106ed5760008555610733565b82601f1061070657805160ff1916838001178555610733565b82800160010185558215610733579182015b82811115610733578251825591602001919060010190610718565b5061073f929150610743565b5090565b5b8082111561073f5760008155600101610744565b600067ffffffffffffffff8084111561077357610773610b84565b604051601f8501601f19908116603f0116810190828211818310171561079b5761079b610b84565b816040528093508581528686860111156107b457600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b03811681146107e557600080fd5b919050565b600082601f8301126107fb57600080fd5b61080a83833560208501610758565b9392505050565b60006020828403121561082357600080fd5b61080a826107ce565b60006020828403121561083e57600080fd5b813567ffffffffffffffff81111561085557600080fd5b610861848285016107ea565b949350505050565b60008060006060848603121561087e57600080fd5b833567ffffffffffffffff81111561089557600080fd5b6108a1868287016107ea565b9660208601359650604090950135949350505050565b600080604083850312156108ca57600080fd5b823567ffffffffffffffff808211156108e257600080fd5b6108ee868387016107ea565b9350602085013591508082111561090457600080fd5b508301601f8101851361091657600080fd5b61092585823560208401610758565b9150509250929050565b600080600080600080600060e0888a03121561094a57600080fd5b873567ffffffffffffffff8082111561096257600080fd5b61096e8b838c016107ea565b985060208a013591508082111561098457600080fd5b6109908b838c016107ea565b975060408a01359150808211156109a657600080fd5b6109b28b838c016107ea565b965060608a01359150808211156109c857600080fd5b6109d48b838c016107ea565b955060808a01359150808211156109ea57600080fd5b6109f68b838c016107ea565b945060a08a0135915080821115610a0c57600080fd5b50610a198a828b016107ea565b925050610a2860c089016107ce565b905092959891949750929550565b60008251610a48818460208701610af0565b9190910192915050565b60008651610a64818460208b01610af0565b865190830190610a78818360208b01610af0565b8651910190610a8b818360208a01610af0565b8551910190610a9e818360208901610af0565b8451910190610ab1818360208801610af0565b01979650505050505050565b6020815260008251806020840152610adc816040850160208701610af0565b601f01601f19169190910160400192915050565b60005b83811015610b0b578181015183820152602001610af3565b83811115610b1a576000848401525b50505050565b600181811c90821680610b3457607f821691505b60208210811415610b5557634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610b7d57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052604160045260246000fdfea264697066735822122031868b000766c18e8c0c378a3809b58fa487038437c4a3621608a29fd6adc96a64736f6c63430008070033",
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

// GetHashUserInfo is a free data retrieval call binding the contract method 0xd203f63c.
//
// Solidity: function GetHashUserInfo(address userAddress) view returns(bytes32)
func (_Document *DocumentCaller) GetHashUserInfo(opts *bind.CallOpts, userAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "GetHashUserInfo", userAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xd203f63c.
//
// Solidity: function GetHashUserInfo(address userAddress) view returns(bytes32)
func (_Document *DocumentSession) GetHashUserInfo(userAddress common.Address) ([32]byte, error) {
	return _Document.Contract.GetHashUserInfo(&_Document.CallOpts, userAddress)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xd203f63c.
//
// Solidity: function GetHashUserInfo(address userAddress) view returns(bytes32)
func (_Document *DocumentCallerSession) GetHashUserInfo(userAddress common.Address) ([32]byte, error) {
	return _Document.Contract.GetHashUserInfo(&_Document.CallOpts, userAddress)
}

// GetPublicKey is a free data retrieval call binding the contract method 0xbff90821.
//
// Solidity: function getPublicKey(string userID) view returns(address)
func (_Document *DocumentCaller) GetPublicKey(opts *bind.CallOpts, userID string) (common.Address, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getPublicKey", userID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0xbff90821.
//
// Solidity: function getPublicKey(string userID) view returns(address)
func (_Document *DocumentSession) GetPublicKey(userID string) (common.Address, error) {
	return _Document.Contract.GetPublicKey(&_Document.CallOpts, userID)
}

// GetPublicKey is a free data retrieval call binding the contract method 0xbff90821.
//
// Solidity: function getPublicKey(string userID) view returns(address)
func (_Document *DocumentCallerSession) GetPublicKey(userID string) (common.Address, error) {
	return _Document.Contract.GetPublicKey(&_Document.CallOpts, userID)
}

// GetUserID is a free data retrieval call binding the contract method 0x0ef12aa9.
//
// Solidity: function getUserID(address publicKey) view returns(string)
func (_Document *DocumentCaller) GetUserID(opts *bind.CallOpts, publicKey common.Address) (string, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "getUserID", publicKey)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserID is a free data retrieval call binding the contract method 0x0ef12aa9.
//
// Solidity: function getUserID(address publicKey) view returns(string)
func (_Document *DocumentSession) GetUserID(publicKey common.Address) (string, error) {
	return _Document.Contract.GetUserID(&_Document.CallOpts, publicKey)
}

// GetUserID is a free data retrieval call binding the contract method 0x0ef12aa9.
//
// Solidity: function getUserID(address publicKey) view returns(string)
func (_Document *DocumentCallerSession) GetUserID(publicKey common.Address) (string, error) {
	return _Document.Contract.GetUserID(&_Document.CallOpts, publicKey)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string userID, bytes32 digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentCaller) VerifyDoc(opts *bind.CallOpts, userID string, digest [32]byte, indexDoc *big.Int) (bool, error) {
	var out []interface{}
	err := _Document.contract.Call(opts, &out, "verifyDoc", userID, digest, indexDoc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDoc is a free data retrieval call binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string userID, bytes32 digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentSession) VerifyDoc(userID string, digest [32]byte, indexDoc *big.Int) (bool, error) {
	return _Document.Contract.VerifyDoc(&_Document.CallOpts, userID, digest, indexDoc)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xd2de39a8.
//
// Solidity: function verifyDoc(string userID, bytes32 digest, uint256 indexDoc) view returns(bool)
func (_Document *DocumentCallerSession) VerifyDoc(userID string, digest [32]byte, indexDoc *big.Int) (bool, error) {
	return _Document.Contract.VerifyDoc(&_Document.CallOpts, userID, digest, indexDoc)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns(uint256)
func (_Document *DocumentTransactor) SaveDoc(opts *bind.TransactOpts, userID string, signature []byte) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "saveDoc", userID, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns(uint256)
func (_Document *DocumentSession) SaveDoc(userID string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.SaveDoc(&_Document.TransactOpts, userID, signature)
}

// SaveDoc is a paid mutator transaction binding the contract method 0x3a8fb351.
//
// Solidity: function saveDoc(string userID, bytes signature) returns(uint256)
func (_Document *DocumentTransactorSession) SaveDoc(userID string, signature []byte) (*types.Transaction, error) {
	return _Document.Contract.SaveDoc(&_Document.TransactOpts, userID, signature)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_Document *DocumentTransactor) StoreUser(opts *bind.TransactOpts, userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.contract.Transact(opts, "storeUser", userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_Document *DocumentSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_Document *DocumentTransactorSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Document.Contract.StoreUser(&_Document.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// IDCMetaData contains all meta data concerning the IDC contract.
var IDCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getUserID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0ef12aa9": "getUserID(address)",
		"3a8fb351": "saveDoc(string,bytes)",
		"8e16afb2": "storeUser(string,string,string,string,string,string,address)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
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

// GetUserID is a paid mutator transaction binding the contract method 0x0ef12aa9.
//
// Solidity: function getUserID(address publicKey) returns(string)
func (_IDC *IDCTransactor) GetUserID(opts *bind.TransactOpts, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "getUserID", publicKey)
}

// GetUserID is a paid mutator transaction binding the contract method 0x0ef12aa9.
//
// Solidity: function getUserID(address publicKey) returns(string)
func (_IDC *IDCSession) GetUserID(publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.GetUserID(&_IDC.TransactOpts, publicKey)
}

// GetUserID is a paid mutator transaction binding the contract method 0x0ef12aa9.
//
// Solidity: function getUserID(address publicKey) returns(string)
func (_IDC *IDCTransactorSession) GetUserID(publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.GetUserID(&_IDC.TransactOpts, publicKey)
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

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_IDC *IDCTransactor) StoreUser(opts *bind.TransactOpts, userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.contract.Transact(opts, "storeUser", userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
func (_IDC *IDCSession) StoreUser(userID string, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _IDC.Contract.StoreUser(&_IDC.TransactOpts, userID, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x8e16afb2.
//
// Solidity: function storeUser(string userID, string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bool)
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

// IDCMetadataMetaData contains all meta data concerning the IDCMetadata contract.
var IDCMetadataMetaData = &bind.MetaData{
	ABI: "[]",
}

// IDCMetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IDCMetadataMetaData.ABI instead.
var IDCMetadataABI = IDCMetadataMetaData.ABI

// IDCMetadata is an auto generated Go binding around an Ethereum contract.
type IDCMetadata struct {
	IDCMetadataCaller     // Read-only binding to the contract
	IDCMetadataTransactor // Write-only binding to the contract
	IDCMetadataFilterer   // Log filterer for contract events
}

// IDCMetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDCMetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCMetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDCMetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCMetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDCMetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCMetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDCMetadataSession struct {
	Contract     *IDCMetadata      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDCMetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDCMetadataCallerSession struct {
	Contract *IDCMetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDCMetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDCMetadataTransactorSession struct {
	Contract     *IDCMetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDCMetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDCMetadataRaw struct {
	Contract *IDCMetadata // Generic contract binding to access the raw methods on
}

// IDCMetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDCMetadataCallerRaw struct {
	Contract *IDCMetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IDCMetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDCMetadataTransactorRaw struct {
	Contract *IDCMetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDCMetadata creates a new instance of IDCMetadata, bound to a specific deployed contract.
func NewIDCMetadata(address common.Address, backend bind.ContractBackend) (*IDCMetadata, error) {
	contract, err := bindIDCMetadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDCMetadata{IDCMetadataCaller: IDCMetadataCaller{contract: contract}, IDCMetadataTransactor: IDCMetadataTransactor{contract: contract}, IDCMetadataFilterer: IDCMetadataFilterer{contract: contract}}, nil
}

// NewIDCMetadataCaller creates a new read-only instance of IDCMetadata, bound to a specific deployed contract.
func NewIDCMetadataCaller(address common.Address, caller bind.ContractCaller) (*IDCMetadataCaller, error) {
	contract, err := bindIDCMetadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDCMetadataCaller{contract: contract}, nil
}

// NewIDCMetadataTransactor creates a new write-only instance of IDCMetadata, bound to a specific deployed contract.
func NewIDCMetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IDCMetadataTransactor, error) {
	contract, err := bindIDCMetadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDCMetadataTransactor{contract: contract}, nil
}

// NewIDCMetadataFilterer creates a new log filterer instance of IDCMetadata, bound to a specific deployed contract.
func NewIDCMetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IDCMetadataFilterer, error) {
	contract, err := bindIDCMetadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDCMetadataFilterer{contract: contract}, nil
}

// bindIDCMetadata binds a generic wrapper to an already deployed contract.
func bindIDCMetadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDCMetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDCMetadata *IDCMetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDCMetadata.Contract.IDCMetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDCMetadata *IDCMetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDCMetadata.Contract.IDCMetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDCMetadata *IDCMetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDCMetadata.Contract.IDCMetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDCMetadata *IDCMetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDCMetadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDCMetadata *IDCMetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDCMetadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDCMetadata *IDCMetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDCMetadata.Contract.contract.Transact(opts, method, params...)
}
