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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getUserID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bff90821": "getPublicKey(string)",
		"0ef12aa9": "getUserID(address)",
		"3a8fb351": "saveDoc(string,bytes)",
		"8e16afb2": "storeUser(string,string,string,string,string,string,address)",
		"d2de39a8": "verifyDoc(string,bytes32,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610b99806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630ef12aa91461005c5780633a8fb351146100855780638e16afb2146100a6578063bff90821146100c9578063d2de39a8146100f4575b600080fd5b61006f61006a3660046107da565b610107565b60405161007c9190610a86565b60405180910390f35b610098610093366004610880565b6101b3565b60405190815260200161007c565b6100b96100b43660046108f8565b610267565b604051901515815260200161007c565b6100dc6100d73660046107f5565b6102ea565b6040516001600160a01b03909116815260200161007c565b6100b9610102366004610832565b610320565b6001600160a01b038116600090815260208190526040902080546060919061012e90610ae9565b80601f016020809104026020016040519081016040528092919081815260200182805461015a90610ae9565b80156101a75780601f1061017c576101008083540402835291602001916101a7565b820191906000526020600020905b81548152906001019060200180831161018a57829003601f168201915b50505050509050919050565b6000806000600101846040516101c991906109ff565b908152602001604051809103902060030154905080806101e890610b24565b915050600080600101856040516101ff91906109ff565b908152604080516020928190038301902060008581526004909101835220865190925061023191839190880190610688565b5083516102479060018301906020870190610688565b50835161025d9060028301906020870190610688565b5090949350505050565b6000806102778888888888610555565b6001600160a01b0384166000908152602081815260409091208b51929350916102a5918391908d0190610688565b5060018181018390556002820180546001600160a01b0319166001600160a01b0387161790556040516102d9908c906109ff565b525060019998505050505050505050565b600080600101826040516102fe91906109ff565b908152604051908190036020019020600201546001600160a01b031692915050565b60008060006001018560405161033691906109ff565b90815260200160405180910390209050600081600401600085815260200190815260200160002060405180606001604052908160008201805461037890610ae9565b80601f01602080910402602001604051908101604052809291908181526020018280546103a490610ae9565b80156103f15780601f106103c6576101008083540402835291602001916103f1565b820191906000526020600020905b8154815290600101906020018083116103d457829003601f168201915b5050505050815260200160018201805461040a90610ae9565b80601f016020809104026020016040519081016040528092919081815260200182805461043690610ae9565b80156104835780601f1061045857610100808354040283529160200191610483565b820191906000526020600020905b81548152906001019060200180831161046657829003601f168201915b5050505050815260200160028201805461049c90610ae9565b80601f01602080910402602001604051908101604052809291908181526020018280546104c890610ae9565b80156105155780601f106104ea57610100808354040283529160200191610515565b820191906000526020600020905b8154815290600101906020018083116104f857829003601f168201915b50505091909252505050600283015460208201519192506001600160a01b031690610541908790610591565b6001600160a01b0316149695505050505050565b60008585858585604051602001610570959493929190610a1b565b60405160208183030381529060405280519060200120905095945050505050565b6000806000806105a085610610565b6040805160008152602081018083528b905260ff8316918101919091526060810184905260808101839052929550909350915060019060a0016020604051602081039080840390855afa1580156105fb573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000806000835160411461066a5760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e6774680000000000000000604482015260640160405180910390fd5b50505060208101516040820151606090920151909260009190911a90565b82805461069490610ae9565b90600052602060002090601f0160209004810192826106b657600085556106fc565b82601f106106cf57805160ff19168380011785556106fc565b828001600101855582156106fc579182015b828111156106fc5782518255916020019190600101906106e1565b5061070892915061070c565b5090565b5b80821115610708576000815560010161070d565b600067ffffffffffffffff8084111561073c5761073c610b4d565b604051601f8501601f19908116603f0116810190828211818310171561076457610764610b4d565b8160405280935085815286868601111561077d57600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b03811681146107ae57600080fd5b919050565b600082601f8301126107c457600080fd5b6107d383833560208501610721565b9392505050565b6000602082840312156107ec57600080fd5b6107d382610797565b60006020828403121561080757600080fd5b813567ffffffffffffffff81111561081e57600080fd5b61082a848285016107b3565b949350505050565b60008060006060848603121561084757600080fd5b833567ffffffffffffffff81111561085e57600080fd5b61086a868287016107b3565b9660208601359650604090950135949350505050565b6000806040838503121561089357600080fd5b823567ffffffffffffffff808211156108ab57600080fd5b6108b7868387016107b3565b935060208501359150808211156108cd57600080fd5b508301601f810185136108df57600080fd5b6108ee85823560208401610721565b9150509250929050565b600080600080600080600060e0888a03121561091357600080fd5b873567ffffffffffffffff8082111561092b57600080fd5b6109378b838c016107b3565b985060208a013591508082111561094d57600080fd5b6109598b838c016107b3565b975060408a013591508082111561096f57600080fd5b61097b8b838c016107b3565b965060608a013591508082111561099157600080fd5b61099d8b838c016107b3565b955060808a01359150808211156109b357600080fd5b6109bf8b838c016107b3565b945060a08a01359150808211156109d557600080fd5b506109e28a828b016107b3565b9250506109f160c08901610797565b905092959891949750929550565b60008251610a11818460208701610ab9565b9190910192915050565b60008651610a2d818460208b01610ab9565b865190830190610a41818360208b01610ab9565b8651910190610a54818360208a01610ab9565b8551910190610a67818360208901610ab9565b8451910190610a7a818360208801610ab9565b01979650505050505050565b6020815260008251806020840152610aa5816040850160208701610ab9565b601f01601f19169190910160400192915050565b60005b83811015610ad4578181015183820152602001610abc565b83811115610ae3576000848401525b50505050565b600181811c90821680610afd57607f821691505b60208210811415610b1e57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610b4657634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052604160045260246000fdfea26469706673582212207faa2ebaa4a2bdca8656e0c20ff9bee5602dd6b3af028cc0b975b728c7a144de64736f6c63430008070033",
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
