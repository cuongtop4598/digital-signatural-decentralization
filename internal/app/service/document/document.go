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

// DocumentMetaData contains all meta data concerning the Document contract.
var DocumentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getUserID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userID\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061110b806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630ef12aa9146100515780633a8fb351146100815780638e16afb2146100b1578063d2de39a8146100e1575b600080fd5b61006b60048036038101906100669190610967565b610111565b6040516100789190610d5d565b60405180910390f35b61009b60048036038101906100969190610a03565b6101e7565b6040516100a89190610d9f565b60405180910390f35b6100cb60048036038101906100c69190610a7b565b6102ad565b6040516100d89190610cfd565b60405180910390f35b6100fb60048036038101906100f69190610994565b61039e565b6040516101089190610cfd565b60405180910390f35b60606000800160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001805461016290610f09565b80601f016020809104026020016040519081016040528092919081815260200182805461018e90610f09565b80156101db5780601f106101b0576101008083540402835291602001916101db565b820191906000526020600020905b8154815290600101906020018083116101be57829003601f168201915b50505050509050919050565b6000806000600101846040516101fd9190610c8e565b9081526020016040518091039020600301549050808061021c90610f6c565b915050600080600101856040516102339190610c8e565b9081526020016040518091039020600401600083815260200190815260200160002090508481600001908051906020019061026f92919061071f565b50838160010190805190602001906102889291906107a5565b50838160020190805190602001906102a19291906107a5565b50819250505092915050565b6000806102be898989898989610609565b905060008060000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508981600001908051906020019061031d92919061071f565b50818160010181905550838160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006001018a60405161037d9190610c8e565b90815260200160405180910390209050600192505050979650505050505050565b6000806000600101856040516103b49190610c8e565b9081526020016040518091039020905060008160040160008581526020019081526020016000206040518060600160405290816000820180546103f690610f09565b80601f016020809104026020016040519081016040528092919081815260200182805461042290610f09565b801561046f5780601f106104445761010080835404028352916020019161046f565b820191906000526020600020905b81548152906001019060200180831161045257829003601f168201915b5050505050815260200160018201805461048890610f09565b80601f01602080910402602001604051908101604052809291908181526020018280546104b490610f09565b80156105015780601f106104d657610100808354040283529160200191610501565b820191906000526020600020905b8154815290600101906020018083116104e457829003601f168201915b5050505050815260200160028201805461051a90610f09565b80601f016020809104026020016040519081016040528092919081815260200182805461054690610f09565b80156105935780601f1061056857610100808354040283529160200191610593565b820191906000526020600020905b81548152906001019060200180831161057657829003601f168201915b50505050508152505090508160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166105e7868360200151610648565b73ffffffffffffffffffffffffffffffffffffffff1614925050509392505050565b600086868686868660405160200161062696959493929190610ca5565b6040516020818303038152906040528051906020012090509695505050505050565b600080600080610657856106b7565b925092509250600186828585604051600081526020016040526040516106809493929190610d18565b6020604051602081039080840390855afa1580156106a2573d6000803e3d6000fd5b50505060206040510351935050505092915050565b60008060006041845114610700576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f790610d7f565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b82805461072b90610f09565b90600052602060002090601f01602090048101928261074d5760008555610794565b82601f1061076657805160ff1916838001178555610794565b82800160010185558215610794579182015b82811115610793578251825591602001919060010190610778565b5b5090506107a1919061082b565b5090565b8280546107b190610f09565b90600052602060002090601f0160209004810192826107d3576000855561081a565b82601f106107ec57805160ff191683800117855561081a565b8280016001018555821561081a579182015b828111156108195782518255916020019190600101906107fe565b5b509050610827919061082b565b5090565b5b8082111561084457600081600090555060010161082c565b5090565b600061085b61085684610ddf565b610dba565b90508281526020810184848401111561087757610876611047565b5b610882848285610ec7565b509392505050565b600061089d61089884610e10565b610dba565b9050828152602081018484840111156108b9576108b8611047565b5b6108c4848285610ec7565b509392505050565b6000813590506108db81611090565b92915050565b6000813590506108f0816110a7565b92915050565b600082601f83011261090b5761090a611042565b5b813561091b848260208601610848565b91505092915050565b600082601f83011261093957610938611042565b5b813561094984826020860161088a565b91505092915050565b600081359050610961816110be565b92915050565b60006020828403121561097d5761097c611051565b5b600061098b848285016108cc565b91505092915050565b6000806000606084860312156109ad576109ac611051565b5b600084013567ffffffffffffffff8111156109cb576109ca61104c565b5b6109d786828701610924565b93505060206109e8868287016108e1565b92505060406109f986828701610952565b9150509250925092565b60008060408385031215610a1a57610a19611051565b5b600083013567ffffffffffffffff811115610a3857610a3761104c565b5b610a4485828601610924565b925050602083013567ffffffffffffffff811115610a6557610a6461104c565b5b610a71858286016108f6565b9150509250929050565b600080600080600080600060e0888a031215610a9a57610a99611051565b5b600088013567ffffffffffffffff811115610ab857610ab761104c565b5b610ac48a828b01610924565b975050602088013567ffffffffffffffff811115610ae557610ae461104c565b5b610af18a828b01610924565b965050604088013567ffffffffffffffff811115610b1257610b1161104c565b5b610b1e8a828b01610924565b955050606088013567ffffffffffffffff811115610b3f57610b3e61104c565b5b610b4b8a828b01610924565b945050608088013567ffffffffffffffff811115610b6c57610b6b61104c565b5b610b788a828b01610924565b93505060a088013567ffffffffffffffff811115610b9957610b9861104c565b5b610ba58a828b01610924565b92505060c0610bb68a828b016108cc565b91505092959891949750929550565b610bce81610e7a565b82525050565b610bdd81610e86565b82525050565b6000610bee82610e41565b610bf88185610e4c565b9350610c08818560208601610ed6565b610c1181611056565b840191505092915050565b6000610c2782610e41565b610c318185610e5d565b9350610c41818560208601610ed6565b80840191505092915050565b6000610c5a601883610e4c565b9150610c6582611067565b602082019050919050565b610c7981610eb0565b82525050565b610c8881610eba565b82525050565b6000610c9a8284610c1c565b915081905092915050565b6000610cb18289610c1c565b9150610cbd8288610c1c565b9150610cc98287610c1c565b9150610cd58286610c1c565b9150610ce18285610c1c565b9150610ced8284610c1c565b9150819050979650505050505050565b6000602082019050610d126000830184610bc5565b92915050565b6000608082019050610d2d6000830187610bd4565b610d3a6020830186610c7f565b610d476040830185610bd4565b610d546060830184610bd4565b95945050505050565b60006020820190508181036000830152610d778184610be3565b905092915050565b60006020820190508181036000830152610d9881610c4d565b9050919050565b6000602082019050610db46000830184610c70565b92915050565b6000610dc4610dd5565b9050610dd08282610f3b565b919050565b6000604051905090565b600067ffffffffffffffff821115610dfa57610df9611013565b5b610e0382611056565b9050602081019050919050565b600067ffffffffffffffff821115610e2b57610e2a611013565b5b610e3482611056565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000610e7382610e90565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610ef4578082015181840152602081019050610ed9565b83811115610f03576000848401525b50505050565b60006002820490506001821680610f2157607f821691505b60208210811415610f3557610f34610fe4565b5b50919050565b610f4482611056565b810181811067ffffffffffffffff82111715610f6357610f62611013565b5b80604052505050565b6000610f7782610eb0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610faa57610fa9610fb5565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b61109981610e68565b81146110a457600080fd5b50565b6110b081610e86565b81146110bb57600080fd5b50565b6110c781610eb0565b81146110d257600080fd5b5056fea2646970667358221220b0577cae86791204e2084a9be99129d287474c3b70e670f276e4015fcaa5c0b764736f6c63430008070033",
}

// DocumentABI is the input ABI used to generate the binding from.
// Deprecated: Use DocumentMetaData.ABI instead.
var DocumentABI = DocumentMetaData.ABI

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