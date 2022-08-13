// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userInfo\",\"type\":\"string\"}],\"name\":\"CreateAccountSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"NumberOfOwnerDocument\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"compareBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareString\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatureA\",\"type\":\"bytes\"}],\"name\":\"createMultipleSignerDocument\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createSingleSignerDocument\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"name\":\"getHashUserInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"getSinger\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"DocID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatureB\",\"type\":\"bytes\"}],\"name\":\"signMultipleSignerDocument\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"storeUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"digest\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"indexDoc\",\"type\":\"uint256\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cmnd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateOB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gmail\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Abi *AbiCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Abi *AbiSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Abi.Contract.Bytes32ToString(&_Abi.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Abi *AbiCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Abi.Contract.Bytes32ToString(&_Abi.CallOpts, _bytes32)
}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Abi *AbiCaller) CompareBytes(opts *bind.CallOpts, a [32]byte, b [32]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "compareBytes", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Abi *AbiSession) CompareBytes(a [32]byte, b [32]byte) (bool, error) {
	return _Abi.Contract.CompareBytes(&_Abi.CallOpts, a, b)
}

// CompareBytes is a free data retrieval call binding the contract method 0xef1bade3.
//
// Solidity: function compareBytes(bytes32 a, bytes32 b) pure returns(bool)
func (_Abi *AbiCallerSession) CompareBytes(a [32]byte, b [32]byte) (bool, error) {
	return _Abi.Contract.CompareBytes(&_Abi.CallOpts, a, b)
}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Abi *AbiCaller) CompareString(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "compareString", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Abi *AbiSession) CompareString(a string, b string) (bool, error) {
	return _Abi.Contract.CompareString(&_Abi.CallOpts, a, b)
}

// CompareString is a free data retrieval call binding the contract method 0x75b83187.
//
// Solidity: function compareString(string a, string b) pure returns(bool)
func (_Abi *AbiCallerSession) CompareString(a string, b string) (bool, error) {
	return _Abi.Contract.CompareString(&_Abi.CallOpts, a, b)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Abi *AbiCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Abi *AbiSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Abi.Contract.GetEthSignedMessageHash(&_Abi.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Abi *AbiCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Abi.Contract.GetEthSignedMessageHash(&_Abi.CallOpts, _messageHash)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Abi *AbiCaller) GetHashUserInfo(opts *bind.CallOpts, phone string) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getHashUserInfo", phone)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Abi *AbiSession) GetHashUserInfo(phone string) ([32]byte, error) {
	return _Abi.Contract.GetHashUserInfo(&_Abi.CallOpts, phone)
}

// GetHashUserInfo is a free data retrieval call binding the contract method 0xa23f1228.
//
// Solidity: function getHashUserInfo(string phone) view returns(bytes32)
func (_Abi *AbiCallerSession) GetHashUserInfo(phone string) ([32]byte, error) {
	return _Abi.Contract.GetHashUserInfo(&_Abi.CallOpts, phone)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Abi *AbiCaller) GetMessageHash(opts *bind.CallOpts, _message string) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getMessageHash", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Abi *AbiSession) GetMessageHash(_message string) ([32]byte, error) {
	return _Abi.Contract.GetMessageHash(&_Abi.CallOpts, _message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Abi *AbiCallerSession) GetMessageHash(_message string) ([32]byte, error) {
	return _Abi.Contract.GetMessageHash(&_Abi.CallOpts, _message)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Abi *AbiCaller) GetSignature(opts *bind.CallOpts, phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getSignature", phone, indexDoc)

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
func (_Abi *AbiSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	return _Abi.Contract.GetSignature(&_Abi.CallOpts, phone, indexDoc)
}

// GetSignature is a free data retrieval call binding the contract method 0xcfd6baa2.
//
// Solidity: function getSignature(string phone, uint256 indexDoc) view returns(bytes, address)
func (_Abi *AbiCallerSession) GetSignature(phone string, indexDoc *big.Int) ([]byte, common.Address, error) {
	return _Abi.Contract.GetSignature(&_Abi.CallOpts, phone, indexDoc)
}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Abi *AbiCaller) GetSinger(opts *bind.CallOpts, phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getSinger", phone, digest, indexDoc)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Abi *AbiSession) GetSinger(phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	return _Abi.Contract.GetSinger(&_Abi.CallOpts, phone, digest, indexDoc)
}

// GetSinger is a free data retrieval call binding the contract method 0xb26feeb9.
//
// Solidity: function getSinger(string phone, bytes32 digest, uint256 indexDoc) view returns(bytes)
func (_Abi *AbiCallerSession) GetSinger(phone string, digest [32]byte, indexDoc *big.Int) ([]byte, error) {
	return _Abi.Contract.GetSinger(&_Abi.CallOpts, phone, digest, indexDoc)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Abi *AbiCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Abi *AbiSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Abi.Contract.RecoverSigner(&_Abi.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Abi *AbiCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Abi.Contract.RecoverSigner(&_Abi.CallOpts, _ethSignedMessageHash, _signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Abi *AbiCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "splitSignature", sig)

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
func (_Abi *AbiSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Abi.Contract.SplitSignature(&_Abi.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Abi *AbiCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Abi.Contract.SplitSignature(&_Abi.CallOpts, sig)
}

// VerifyDoc is a free data retrieval call binding the contract method 0x9ae5c2ee.
//
// Solidity: function verifyDoc(string phone, string digest, uint256 indexDoc) view returns(bool)
func (_Abi *AbiCaller) VerifyDoc(opts *bind.CallOpts, phone string, digest string, indexDoc *big.Int) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "verifyDoc", phone, digest, indexDoc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDoc is a free data retrieval call binding the contract method 0x9ae5c2ee.
//
// Solidity: function verifyDoc(string phone, string digest, uint256 indexDoc) view returns(bool)
func (_Abi *AbiSession) VerifyDoc(phone string, digest string, indexDoc *big.Int) (bool, error) {
	return _Abi.Contract.VerifyDoc(&_Abi.CallOpts, phone, digest, indexDoc)
}

// VerifyDoc is a free data retrieval call binding the contract method 0x9ae5c2ee.
//
// Solidity: function verifyDoc(string phone, string digest, uint256 indexDoc) view returns(bool)
func (_Abi *AbiCallerSession) VerifyDoc(phone string, digest string, indexDoc *big.Int) (bool, error) {
	return _Abi.Contract.VerifyDoc(&_Abi.CallOpts, phone, digest, indexDoc)
}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Abi *AbiCaller) VerifyUser(opts *bind.CallOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "verifyUser", name, cmnd, dateOB, phone, gmail, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Abi *AbiSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	return _Abi.Contract.VerifyUser(&_Abi.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// VerifyUser is a free data retrieval call binding the contract method 0xa88c6e45.
//
// Solidity: function verifyUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) view returns(bool)
func (_Abi *AbiCallerSession) VerifyUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (bool, error) {
	return _Abi.Contract.VerifyUser(&_Abi.CallOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// CreateMultipleSignerDocument is a paid mutator transaction binding the contract method 0x69fc026b.
//
// Solidity: function createMultipleSignerDocument(address partner, bytes signatureA) returns(uint256, uint256, bool)
func (_Abi *AbiTransactor) CreateMultipleSignerDocument(opts *bind.TransactOpts, partner common.Address, signatureA []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "createMultipleSignerDocument", partner, signatureA)
}

// CreateMultipleSignerDocument is a paid mutator transaction binding the contract method 0x69fc026b.
//
// Solidity: function createMultipleSignerDocument(address partner, bytes signatureA) returns(uint256, uint256, bool)
func (_Abi *AbiSession) CreateMultipleSignerDocument(partner common.Address, signatureA []byte) (*types.Transaction, error) {
	return _Abi.Contract.CreateMultipleSignerDocument(&_Abi.TransactOpts, partner, signatureA)
}

// CreateMultipleSignerDocument is a paid mutator transaction binding the contract method 0x69fc026b.
//
// Solidity: function createMultipleSignerDocument(address partner, bytes signatureA) returns(uint256, uint256, bool)
func (_Abi *AbiTransactorSession) CreateMultipleSignerDocument(partner common.Address, signatureA []byte) (*types.Transaction, error) {
	return _Abi.Contract.CreateMultipleSignerDocument(&_Abi.TransactOpts, partner, signatureA)
}

// CreateSingleSignerDocument is a paid mutator transaction binding the contract method 0xc552e61a.
//
// Solidity: function createSingleSignerDocument(string phone, bytes signature) returns(bytes32)
func (_Abi *AbiTransactor) CreateSingleSignerDocument(opts *bind.TransactOpts, phone string, signature []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "createSingleSignerDocument", phone, signature)
}

// CreateSingleSignerDocument is a paid mutator transaction binding the contract method 0xc552e61a.
//
// Solidity: function createSingleSignerDocument(string phone, bytes signature) returns(bytes32)
func (_Abi *AbiSession) CreateSingleSignerDocument(phone string, signature []byte) (*types.Transaction, error) {
	return _Abi.Contract.CreateSingleSignerDocument(&_Abi.TransactOpts, phone, signature)
}

// CreateSingleSignerDocument is a paid mutator transaction binding the contract method 0xc552e61a.
//
// Solidity: function createSingleSignerDocument(string phone, bytes signature) returns(bytes32)
func (_Abi *AbiTransactorSession) CreateSingleSignerDocument(phone string, signature []byte) (*types.Transaction, error) {
	return _Abi.Contract.CreateSingleSignerDocument(&_Abi.TransactOpts, phone, signature)
}

// SignMultipleSignerDocument is a paid mutator transaction binding the contract method 0x1c3c6d2b.
//
// Solidity: function signMultipleSignerDocument(uint256 DocID, bytes signatureB) returns(uint256, bool)
func (_Abi *AbiTransactor) SignMultipleSignerDocument(opts *bind.TransactOpts, DocID *big.Int, signatureB []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "signMultipleSignerDocument", DocID, signatureB)
}

// SignMultipleSignerDocument is a paid mutator transaction binding the contract method 0x1c3c6d2b.
//
// Solidity: function signMultipleSignerDocument(uint256 DocID, bytes signatureB) returns(uint256, bool)
func (_Abi *AbiSession) SignMultipleSignerDocument(DocID *big.Int, signatureB []byte) (*types.Transaction, error) {
	return _Abi.Contract.SignMultipleSignerDocument(&_Abi.TransactOpts, DocID, signatureB)
}

// SignMultipleSignerDocument is a paid mutator transaction binding the contract method 0x1c3c6d2b.
//
// Solidity: function signMultipleSignerDocument(uint256 DocID, bytes signatureB) returns(uint256, bool)
func (_Abi *AbiTransactorSession) SignMultipleSignerDocument(DocID *big.Int, signatureB []byte) (*types.Transaction, error) {
	return _Abi.Contract.SignMultipleSignerDocument(&_Abi.TransactOpts, DocID, signatureB)
}

// StoreUser is a paid mutator transaction binding the contract method 0x9a77e1c1.
//
// Solidity: function storeUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Abi *AbiTransactor) StoreUser(opts *bind.TransactOpts, name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "storeUser", name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x9a77e1c1.
//
// Solidity: function storeUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Abi *AbiSession) StoreUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Abi.Contract.StoreUser(&_Abi.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// StoreUser is a paid mutator transaction binding the contract method 0x9a77e1c1.
//
// Solidity: function storeUser(string name, string cmnd, string dateOB, string phone, string gmail, address publicKey) returns(bytes32)
func (_Abi *AbiTransactorSession) StoreUser(name string, cmnd string, dateOB string, phone string, gmail string, publicKey common.Address) (*types.Transaction, error) {
	return _Abi.Contract.StoreUser(&_Abi.TransactOpts, name, cmnd, dateOB, phone, gmail, publicKey)
}

// AbiCreateAccountSuccessIterator is returned from FilterCreateAccountSuccess and is used to iterate over the raw logs and unpacked data for CreateAccountSuccess events raised by the Abi contract.
type AbiCreateAccountSuccessIterator struct {
	Event *AbiCreateAccountSuccess // Event containing the contract specifics and raw log

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
func (it *AbiCreateAccountSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiCreateAccountSuccess)
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
		it.Event = new(AbiCreateAccountSuccess)
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
func (it *AbiCreateAccountSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiCreateAccountSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiCreateAccountSuccess represents a CreateAccountSuccess event raised by the Abi contract.
type AbiCreateAccountSuccess struct {
	UserInfo string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreateAccountSuccess is a free log retrieval operation binding the contract event 0x9905c6c3f7ce69816a7cdf2185623a1e3ef1fa5a03010d86987f9d1035be0300.
//
// Solidity: event CreateAccountSuccess(string userInfo)
func (_Abi *AbiFilterer) FilterCreateAccountSuccess(opts *bind.FilterOpts) (*AbiCreateAccountSuccessIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "CreateAccountSuccess")
	if err != nil {
		return nil, err
	}
	return &AbiCreateAccountSuccessIterator{contract: _Abi.contract, event: "CreateAccountSuccess", logs: logs, sub: sub}, nil
}

// WatchCreateAccountSuccess is a free log subscription operation binding the contract event 0x9905c6c3f7ce69816a7cdf2185623a1e3ef1fa5a03010d86987f9d1035be0300.
//
// Solidity: event CreateAccountSuccess(string userInfo)
func (_Abi *AbiFilterer) WatchCreateAccountSuccess(opts *bind.WatchOpts, sink chan<- *AbiCreateAccountSuccess) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "CreateAccountSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiCreateAccountSuccess)
				if err := _Abi.contract.UnpackLog(event, "CreateAccountSuccess", log); err != nil {
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

// ParseCreateAccountSuccess is a log parse operation binding the contract event 0x9905c6c3f7ce69816a7cdf2185623a1e3ef1fa5a03010d86987f9d1035be0300.
//
// Solidity: event CreateAccountSuccess(string userInfo)
func (_Abi *AbiFilterer) ParseCreateAccountSuccess(log types.Log) (*AbiCreateAccountSuccess, error) {
	event := new(AbiCreateAccountSuccess)
	if err := _Abi.contract.UnpackLog(event, "CreateAccountSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiNumberOfOwnerDocumentIterator is returned from FilterNumberOfOwnerDocument and is used to iterate over the raw logs and unpacked data for NumberOfOwnerDocument events raised by the Abi contract.
type AbiNumberOfOwnerDocumentIterator struct {
	Event *AbiNumberOfOwnerDocument // Event containing the contract specifics and raw log

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
func (it *AbiNumberOfOwnerDocumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiNumberOfOwnerDocument)
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
		it.Event = new(AbiNumberOfOwnerDocument)
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
func (it *AbiNumberOfOwnerDocumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiNumberOfOwnerDocumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiNumberOfOwnerDocument represents a NumberOfOwnerDocument event raised by the Abi contract.
type AbiNumberOfOwnerDocument struct {
	Num *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNumberOfOwnerDocument is a free log retrieval operation binding the contract event 0x6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db904.
//
// Solidity: event NumberOfOwnerDocument(uint256 num)
func (_Abi *AbiFilterer) FilterNumberOfOwnerDocument(opts *bind.FilterOpts) (*AbiNumberOfOwnerDocumentIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "NumberOfOwnerDocument")
	if err != nil {
		return nil, err
	}
	return &AbiNumberOfOwnerDocumentIterator{contract: _Abi.contract, event: "NumberOfOwnerDocument", logs: logs, sub: sub}, nil
}

// WatchNumberOfOwnerDocument is a free log subscription operation binding the contract event 0x6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db904.
//
// Solidity: event NumberOfOwnerDocument(uint256 num)
func (_Abi *AbiFilterer) WatchNumberOfOwnerDocument(opts *bind.WatchOpts, sink chan<- *AbiNumberOfOwnerDocument) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "NumberOfOwnerDocument")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiNumberOfOwnerDocument)
				if err := _Abi.contract.UnpackLog(event, "NumberOfOwnerDocument", log); err != nil {
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

// ParseNumberOfOwnerDocument is a log parse operation binding the contract event 0x6fc5b69e508af84c5302ddb2ca1e6965cac21ddcad708dce220155101a5db904.
//
// Solidity: event NumberOfOwnerDocument(uint256 num)
func (_Abi *AbiFilterer) ParseNumberOfOwnerDocument(log types.Log) (*AbiNumberOfOwnerDocument, error) {
	event := new(AbiNumberOfOwnerDocument)
	if err := _Abi.contract.UnpackLog(event, "NumberOfOwnerDocument", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
