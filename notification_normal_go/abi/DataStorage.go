// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"strings"
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_modelParamsHash\",\"type\":\"bytes32\"}],\"name\":\"setModelParamsHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"encryptedAesKeyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_finalDistance\",\"type\":\"bytes32\"}],\"name\":\"setFinalDistance\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modelHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"trainedResultHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalDistance\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_encryptedAesKeyHash\",\"type\":\"bytes32\"}],\"name\":\"setEncryptedAesKeyHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fheAddressHash\",\"type\":\"bytes32\"}],\"name\":\"setFheAddressHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fheAddressHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trainedResultHash\",\"type\":\"bytes32\"}],\"name\":\"setTrainedResultHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rsaPublicKeyHash\",\"type\":\"bytes32\"}],\"name\":\"setRsaPublicKeyHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_modelHash\",\"type\":\"bytes32\"}],\"name\":\"setModelHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"name\":\"setDataHash\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"modelParamsHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rsaPublicKeyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_fheAddressHash\",\"type\":\"bytes32\"}],\"name\":\"LogSetFheAddressHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_modelHash\",\"type\":\"bytes32\"}],\"name\":\"LogSetModelHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_rsaPublicKeyHash\",\"type\":\"bytes32\"}],\"name\":\"LogSetRsaPublicKeyHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"name\":\"LogSetDataHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_encryptedAesKeyHash\",\"type\":\"bytes32\"}],\"name\":\"LogEncryptedAesKeyHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_trainedResultHash\",\"type\":\"bytes32\"}],\"name\":\"LogTrainedResultHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_finalDistance\",\"type\":\"bytes32\"}],\"name\":\"LogSetFinalDistance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_modelParamsHash\",\"type\":\"bytes32\"}],\"name\":\"LogSetModelParamsHash\",\"type\":\"event\"}]"

// MainBin is the compiled bytecode used for deploying new contracts.
const MainBin = `608060405234801561001057600080fd5b50610676806100206000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630483cb37146100eb57806312ddf9ae1461010f5780631a8fdab9146101425780631b2ccc09146101665780631b3012a3146101995780634b4ece52146101cc5780634c1c8789146101ff5780635ee13861146102325780637e0f9ddf146102565780638b9ccebb1461027a578063994d762e146102ad578063a1ad03c2146102d1578063a86fe55d146102f5578063b6a8f70314610319578063b6b726741461033d578063c138127e14610370575b600080fd5b61010d60048036038101908080356000191690602001909291905050506103a3565b005b34801561011b57600080fd5b506101246103f2565b60405180826000191660001916815260200191505060405180910390f35b61016460048036038101908080356000191690602001909291905050506103f8565b005b34801561017257600080fd5b5061017b610447565b60405180826000191660001916815260200191505060405180910390f35b3480156101a557600080fd5b506101ae61044d565b60405180826000191660001916815260200191505060405180910390f35b3480156101d857600080fd5b506101e1610453565b60405180826000191660001916815260200191505060405180910390f35b34801561020b57600080fd5b50610214610459565b60405180826000191660001916815260200191505060405180910390f35b610254600480360381019080803560001916906020019092919050505061045f565b005b61027860048036038101908080356000191690602001909291905050506104ae565b005b34801561028657600080fd5b5061028f6104fd565b60405180826000191660001916815260200191505060405180910390f35b6102cf6004803603810190808035600019169060200190929190505050610503565b005b6102f36004803603810190808035600019169060200190929190505050610552565b005b61031760048036038101908080356000191690602001909291905050506105a0565b005b61033b60048036038101908080356000191690602001909291905050506105ef565b005b34801561034957600080fd5b5061035261063e565b60405180826000191660001916815260200191505060405180910390f35b34801561037c57600080fd5b50610385610644565b60405180826000191660001916815260200191505060405180910390f35b80600781600019169055507f7034869c2aa7e7265ca374e8476f25d1256880850a27a346a0d3d85d0af4477260075460405180826000191660001916815260200191505060405180910390a150565b60045481565b80600681600019169055507f9c6f5a9ba93907c83634f959c44c55107e452ca00a217f4d7a4ff515e66c0cd060065460405180826000191660001916815260200191505060405180910390a150565b60035481565b60025481565b60055481565b60065481565b80600481600019169055507f2a09c1917aa0c55675d6fd9d1354c315ad731dd3b2e1ec6520d0f8307e10714860045460405180826000191660001916815260200191505060405180910390a150565b80600081600019169055507f84cd278fd309da11907c3a07e62f8fb36f0f4a4a1ad30851d0b4c91e7cc8af8260005460405180826000191660001916815260200191505060405180910390a150565b60005481565b80600581600019169055507f906e27cc089c73cd1d44fcd513d956b221515cee6e280c87ea18fa3a027c16f960055460405180826000191660001916815260200191505060405180910390a150565b80600181600019169055507e09665638a273ab94b50f85ed606cdb9b6b7118c15fd7c373d7d64e09cefe8760015460405180826000191660001916815260200191505060405180910390a150565b80600381600019169055507fe6d3727ea0a293c75dcb4456e3041ff031283e8155fef30f9a350a003859e8c360035460405180826000191660001916815260200191505060405180910390a150565b80600281600019169055507fd3bb8ce1bbf26702369bf6da605267726e9ffd668acd1106e8868499bef21dc660025460405180826000191660001916815260200191505060405180910390a150565b60075481565b600154815600a165627a7a72305820898e811db361c8a868763fcfe44ca1f2a1ad8c76025d9d38550bd930ad9342030029`

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(bytes32)
func (_Main *MainCaller) DataHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "dataHash")
	return *ret0, err
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(bytes32)
func (_Main *MainSession) DataHash() ([32]byte, error) {
	return _Main.Contract.DataHash(&_Main.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(bytes32)
func (_Main *MainCallerSession) DataHash() ([32]byte, error) {
	return _Main.Contract.DataHash(&_Main.CallOpts)
}

// EncryptedAesKeyHash is a free data retrieval call binding the contract method 0x12ddf9ae.
//
// Solidity: function encryptedAesKeyHash() constant returns(bytes32)
func (_Main *MainCaller) EncryptedAesKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "encryptedAesKeyHash")
	return *ret0, err
}

// EncryptedAesKeyHash is a free data retrieval call binding the contract method 0x12ddf9ae.
//
// Solidity: function encryptedAesKeyHash() constant returns(bytes32)
func (_Main *MainSession) EncryptedAesKeyHash() ([32]byte, error) {
	return _Main.Contract.EncryptedAesKeyHash(&_Main.CallOpts)
}

// EncryptedAesKeyHash is a free data retrieval call binding the contract method 0x12ddf9ae.
//
// Solidity: function encryptedAesKeyHash() constant returns(bytes32)
func (_Main *MainCallerSession) EncryptedAesKeyHash() ([32]byte, error) {
	return _Main.Contract.EncryptedAesKeyHash(&_Main.CallOpts)
}

// FheAddressHash is a free data retrieval call binding the contract method 0x8b9ccebb.
//
// Solidity: function fheAddressHash() constant returns(bytes32)
func (_Main *MainCaller) FheAddressHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "fheAddressHash")
	return *ret0, err
}

// FheAddressHash is a free data retrieval call binding the contract method 0x8b9ccebb.
//
// Solidity: function fheAddressHash() constant returns(bytes32)
func (_Main *MainSession) FheAddressHash() ([32]byte, error) {
	return _Main.Contract.FheAddressHash(&_Main.CallOpts)
}

// FheAddressHash is a free data retrieval call binding the contract method 0x8b9ccebb.
//
// Solidity: function fheAddressHash() constant returns(bytes32)
func (_Main *MainCallerSession) FheAddressHash() ([32]byte, error) {
	return _Main.Contract.FheAddressHash(&_Main.CallOpts)
}

// FinalDistance is a free data retrieval call binding the contract method 0x4c1c8789.
//
// Solidity: function finalDistance() constant returns(bytes32)
func (_Main *MainCaller) FinalDistance(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "finalDistance")
	return *ret0, err
}

// FinalDistance is a free data retrieval call binding the contract method 0x4c1c8789.
//
// Solidity: function finalDistance() constant returns(bytes32)
func (_Main *MainSession) FinalDistance() ([32]byte, error) {
	return _Main.Contract.FinalDistance(&_Main.CallOpts)
}

// FinalDistance is a free data retrieval call binding the contract method 0x4c1c8789.
//
// Solidity: function finalDistance() constant returns(bytes32)
func (_Main *MainCallerSession) FinalDistance() ([32]byte, error) {
	return _Main.Contract.FinalDistance(&_Main.CallOpts)
}

// ModelHash is a free data retrieval call binding the contract method 0x1b2ccc09.
//
// Solidity: function modelHash() constant returns(bytes32)
func (_Main *MainCaller) ModelHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "modelHash")
	return *ret0, err
}

// ModelHash is a free data retrieval call binding the contract method 0x1b2ccc09.
//
// Solidity: function modelHash() constant returns(bytes32)
func (_Main *MainSession) ModelHash() ([32]byte, error) {
	return _Main.Contract.ModelHash(&_Main.CallOpts)
}

// ModelHash is a free data retrieval call binding the contract method 0x1b2ccc09.
//
// Solidity: function modelHash() constant returns(bytes32)
func (_Main *MainCallerSession) ModelHash() ([32]byte, error) {
	return _Main.Contract.ModelHash(&_Main.CallOpts)
}

// ModelParamsHash is a free data retrieval call binding the contract method 0xb6b72674.
//
// Solidity: function modelParamsHash() constant returns(bytes32)
func (_Main *MainCaller) ModelParamsHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "modelParamsHash")
	return *ret0, err
}

// ModelParamsHash is a free data retrieval call binding the contract method 0xb6b72674.
//
// Solidity: function modelParamsHash() constant returns(bytes32)
func (_Main *MainSession) ModelParamsHash() ([32]byte, error) {
	return _Main.Contract.ModelParamsHash(&_Main.CallOpts)
}

// ModelParamsHash is a free data retrieval call binding the contract method 0xb6b72674.
//
// Solidity: function modelParamsHash() constant returns(bytes32)
func (_Main *MainCallerSession) ModelParamsHash() ([32]byte, error) {
	return _Main.Contract.ModelParamsHash(&_Main.CallOpts)
}

// RsaPublicKeyHash is a free data retrieval call binding the contract method 0xc138127e.
//
// Solidity: function rsaPublicKeyHash() constant returns(bytes32)
func (_Main *MainCaller) RsaPublicKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "rsaPublicKeyHash")
	return *ret0, err
}

// RsaPublicKeyHash is a free data retrieval call binding the contract method 0xc138127e.
//
// Solidity: function rsaPublicKeyHash() constant returns(bytes32)
func (_Main *MainSession) RsaPublicKeyHash() ([32]byte, error) {
	return _Main.Contract.RsaPublicKeyHash(&_Main.CallOpts)
}

// RsaPublicKeyHash is a free data retrieval call binding the contract method 0xc138127e.
//
// Solidity: function rsaPublicKeyHash() constant returns(bytes32)
func (_Main *MainCallerSession) RsaPublicKeyHash() ([32]byte, error) {
	return _Main.Contract.RsaPublicKeyHash(&_Main.CallOpts)
}

// TrainedResultHash is a free data retrieval call binding the contract method 0x4b4ece52.
//
// Solidity: function trainedResultHash() constant returns(bytes32)
func (_Main *MainCaller) TrainedResultHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "trainedResultHash")
	return *ret0, err
}

// TrainedResultHash is a free data retrieval call binding the contract method 0x4b4ece52.
//
// Solidity: function trainedResultHash() constant returns(bytes32)
func (_Main *MainSession) TrainedResultHash() ([32]byte, error) {
	return _Main.Contract.TrainedResultHash(&_Main.CallOpts)
}

// TrainedResultHash is a free data retrieval call binding the contract method 0x4b4ece52.
//
// Solidity: function trainedResultHash() constant returns(bytes32)
func (_Main *MainCallerSession) TrainedResultHash() ([32]byte, error) {
	return _Main.Contract.TrainedResultHash(&_Main.CallOpts)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb6a8f703.
//
// Solidity: function setDataHash(_dataHash bytes32) returns()
func (_Main *MainTransactor) SetDataHash(opts *bind.TransactOpts, _dataHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setDataHash", _dataHash)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb6a8f703.
//
// Solidity: function setDataHash(_dataHash bytes32) returns()
func (_Main *MainSession) SetDataHash(_dataHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetDataHash(&_Main.TransactOpts, _dataHash)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb6a8f703.
//
// Solidity: function setDataHash(_dataHash bytes32) returns()
func (_Main *MainTransactorSession) SetDataHash(_dataHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetDataHash(&_Main.TransactOpts, _dataHash)
}

// SetEncryptedAesKeyHash is a paid mutator transaction binding the contract method 0x5ee13861.
//
// Solidity: function setEncryptedAesKeyHash(_encryptedAesKeyHash bytes32) returns()
func (_Main *MainTransactor) SetEncryptedAesKeyHash(opts *bind.TransactOpts, _encryptedAesKeyHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setEncryptedAesKeyHash", _encryptedAesKeyHash)
}

// SetEncryptedAesKeyHash is a paid mutator transaction binding the contract method 0x5ee13861.
//
// Solidity: function setEncryptedAesKeyHash(_encryptedAesKeyHash bytes32) returns()
func (_Main *MainSession) SetEncryptedAesKeyHash(_encryptedAesKeyHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetEncryptedAesKeyHash(&_Main.TransactOpts, _encryptedAesKeyHash)
}

// SetEncryptedAesKeyHash is a paid mutator transaction binding the contract method 0x5ee13861.
//
// Solidity: function setEncryptedAesKeyHash(_encryptedAesKeyHash bytes32) returns()
func (_Main *MainTransactorSession) SetEncryptedAesKeyHash(_encryptedAesKeyHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetEncryptedAesKeyHash(&_Main.TransactOpts, _encryptedAesKeyHash)
}

// SetFheAddressHash is a paid mutator transaction binding the contract method 0x7e0f9ddf.
//
// Solidity: function setFheAddressHash(_fheAddressHash bytes32) returns()
func (_Main *MainTransactor) SetFheAddressHash(opts *bind.TransactOpts, _fheAddressHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFheAddressHash", _fheAddressHash)
}

// SetFheAddressHash is a paid mutator transaction binding the contract method 0x7e0f9ddf.
//
// Solidity: function setFheAddressHash(_fheAddressHash bytes32) returns()
func (_Main *MainSession) SetFheAddressHash(_fheAddressHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetFheAddressHash(&_Main.TransactOpts, _fheAddressHash)
}

// SetFheAddressHash is a paid mutator transaction binding the contract method 0x7e0f9ddf.
//
// Solidity: function setFheAddressHash(_fheAddressHash bytes32) returns()
func (_Main *MainTransactorSession) SetFheAddressHash(_fheAddressHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetFheAddressHash(&_Main.TransactOpts, _fheAddressHash)
}

// SetFinalDistance is a paid mutator transaction binding the contract method 0x1a8fdab9.
//
// Solidity: function setFinalDistance(_finalDistance bytes32) returns()
func (_Main *MainTransactor) SetFinalDistance(opts *bind.TransactOpts, _finalDistance [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFinalDistance", _finalDistance)
}

// SetFinalDistance is a paid mutator transaction binding the contract method 0x1a8fdab9.
//
// Solidity: function setFinalDistance(_finalDistance bytes32) returns()
func (_Main *MainSession) SetFinalDistance(_finalDistance [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetFinalDistance(&_Main.TransactOpts, _finalDistance)
}

// SetFinalDistance is a paid mutator transaction binding the contract method 0x1a8fdab9.
//
// Solidity: function setFinalDistance(_finalDistance bytes32) returns()
func (_Main *MainTransactorSession) SetFinalDistance(_finalDistance [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetFinalDistance(&_Main.TransactOpts, _finalDistance)
}

// SetModelHash is a paid mutator transaction binding the contract method 0xa86fe55d.
//
// Solidity: function setModelHash(_modelHash bytes32) returns()
func (_Main *MainTransactor) SetModelHash(opts *bind.TransactOpts, _modelHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setModelHash", _modelHash)
}

// SetModelHash is a paid mutator transaction binding the contract method 0xa86fe55d.
//
// Solidity: function setModelHash(_modelHash bytes32) returns()
func (_Main *MainSession) SetModelHash(_modelHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetModelHash(&_Main.TransactOpts, _modelHash)
}

// SetModelHash is a paid mutator transaction binding the contract method 0xa86fe55d.
//
// Solidity: function setModelHash(_modelHash bytes32) returns()
func (_Main *MainTransactorSession) SetModelHash(_modelHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetModelHash(&_Main.TransactOpts, _modelHash)
}

// SetModelParamsHash is a paid mutator transaction binding the contract method 0x0483cb37.
//
// Solidity: function setModelParamsHash(_modelParamsHash bytes32) returns()
func (_Main *MainTransactor) SetModelParamsHash(opts *bind.TransactOpts, _modelParamsHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setModelParamsHash", _modelParamsHash)
}

// SetModelParamsHash is a paid mutator transaction binding the contract method 0x0483cb37.
//
// Solidity: function setModelParamsHash(_modelParamsHash bytes32) returns()
func (_Main *MainSession) SetModelParamsHash(_modelParamsHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetModelParamsHash(&_Main.TransactOpts, _modelParamsHash)
}

// SetModelParamsHash is a paid mutator transaction binding the contract method 0x0483cb37.
//
// Solidity: function setModelParamsHash(_modelParamsHash bytes32) returns()
func (_Main *MainTransactorSession) SetModelParamsHash(_modelParamsHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetModelParamsHash(&_Main.TransactOpts, _modelParamsHash)
}

// SetRsaPublicKeyHash is a paid mutator transaction binding the contract method 0xa1ad03c2.
//
// Solidity: function setRsaPublicKeyHash(_rsaPublicKeyHash bytes32) returns()
func (_Main *MainTransactor) SetRsaPublicKeyHash(opts *bind.TransactOpts, _rsaPublicKeyHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setRsaPublicKeyHash", _rsaPublicKeyHash)
}

// SetRsaPublicKeyHash is a paid mutator transaction binding the contract method 0xa1ad03c2.
//
// Solidity: function setRsaPublicKeyHash(_rsaPublicKeyHash bytes32) returns()
func (_Main *MainSession) SetRsaPublicKeyHash(_rsaPublicKeyHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetRsaPublicKeyHash(&_Main.TransactOpts, _rsaPublicKeyHash)
}

// SetRsaPublicKeyHash is a paid mutator transaction binding the contract method 0xa1ad03c2.
//
// Solidity: function setRsaPublicKeyHash(_rsaPublicKeyHash bytes32) returns()
func (_Main *MainTransactorSession) SetRsaPublicKeyHash(_rsaPublicKeyHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetRsaPublicKeyHash(&_Main.TransactOpts, _rsaPublicKeyHash)
}

// SetTrainedResultHash is a paid mutator transaction binding the contract method 0x994d762e.
//
// Solidity: function setTrainedResultHash(_trainedResultHash bytes32) returns()
func (_Main *MainTransactor) SetTrainedResultHash(opts *bind.TransactOpts, _trainedResultHash [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setTrainedResultHash", _trainedResultHash)
}

// SetTrainedResultHash is a paid mutator transaction binding the contract method 0x994d762e.
//
// Solidity: function setTrainedResultHash(_trainedResultHash bytes32) returns()
func (_Main *MainSession) SetTrainedResultHash(_trainedResultHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetTrainedResultHash(&_Main.TransactOpts, _trainedResultHash)
}

// SetTrainedResultHash is a paid mutator transaction binding the contract method 0x994d762e.
//
// Solidity: function setTrainedResultHash(_trainedResultHash bytes32) returns()
func (_Main *MainTransactorSession) SetTrainedResultHash(_trainedResultHash [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetTrainedResultHash(&_Main.TransactOpts, _trainedResultHash)
}

// MainLogEncryptedAesKeyHashIterator is returned from FilterLogEncryptedAesKeyHash and is used to iterate over the raw logs and unpacked data for LogEncryptedAesKeyHash events raised by the Main contract.
type MainLogEncryptedAesKeyHashIterator struct {
	Event *MainLogEncryptedAesKeyHash // Event containing the contract specifics and raw log

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
func (it *MainLogEncryptedAesKeyHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogEncryptedAesKeyHash)
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
		it.Event = new(MainLogEncryptedAesKeyHash)
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
func (it *MainLogEncryptedAesKeyHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogEncryptedAesKeyHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogEncryptedAesKeyHash represents a LogEncryptedAesKeyHash event raised by the Main contract.
type MainLogEncryptedAesKeyHash struct {
	EncryptedAesKeyHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogEncryptedAesKeyHash is a free log retrieval operation binding the contract event 0x2a09c1917aa0c55675d6fd9d1354c315ad731dd3b2e1ec6520d0f8307e107148.
//
// Solidity: e LogEncryptedAesKeyHash(_encryptedAesKeyHash bytes32)
func (_Main *MainFilterer) FilterLogEncryptedAesKeyHash(opts *bind.FilterOpts) (*MainLogEncryptedAesKeyHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogEncryptedAesKeyHash")
	if err != nil {
		return nil, err
	}
	return &MainLogEncryptedAesKeyHashIterator{contract: _Main.contract, event: "LogEncryptedAesKeyHash", logs: logs, sub: sub}, nil
}

// WatchLogEncryptedAesKeyHash is a free log subscription operation binding the contract event 0x2a09c1917aa0c55675d6fd9d1354c315ad731dd3b2e1ec6520d0f8307e107148.
//
// Solidity: e LogEncryptedAesKeyHash(_encryptedAesKeyHash bytes32)
func (_Main *MainFilterer) WatchLogEncryptedAesKeyHash(opts *bind.WatchOpts, sink chan<- *MainLogEncryptedAesKeyHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogEncryptedAesKeyHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogEncryptedAesKeyHash)
				if err := _Main.contract.UnpackLog(event, "LogEncryptedAesKeyHash", log); err != nil {
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

// MainLogSetDataHashIterator is returned from FilterLogSetDataHash and is used to iterate over the raw logs and unpacked data for LogSetDataHash events raised by the Main contract.
type MainLogSetDataHashIterator struct {
	Event *MainLogSetDataHash // Event containing the contract specifics and raw log

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
func (it *MainLogSetDataHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogSetDataHash)
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
		it.Event = new(MainLogSetDataHash)
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
func (it *MainLogSetDataHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogSetDataHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogSetDataHash represents a LogSetDataHash event raised by the Main contract.
type MainLogSetDataHash struct {
	DataHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSetDataHash is a free log retrieval operation binding the contract event 0xd3bb8ce1bbf26702369bf6da605267726e9ffd668acd1106e8868499bef21dc6.
//
// Solidity: e LogSetDataHash(_dataHash bytes32)
func (_Main *MainFilterer) FilterLogSetDataHash(opts *bind.FilterOpts) (*MainLogSetDataHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogSetDataHash")
	if err != nil {
		return nil, err
	}
	return &MainLogSetDataHashIterator{contract: _Main.contract, event: "LogSetDataHash", logs: logs, sub: sub}, nil
}

// WatchLogSetDataHash is a free log subscription operation binding the contract event 0xd3bb8ce1bbf26702369bf6da605267726e9ffd668acd1106e8868499bef21dc6.
//
// Solidity: e LogSetDataHash(_dataHash bytes32)
func (_Main *MainFilterer) WatchLogSetDataHash(opts *bind.WatchOpts, sink chan<- *MainLogSetDataHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogSetDataHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogSetDataHash)
				if err := _Main.contract.UnpackLog(event, "LogSetDataHash", log); err != nil {
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

// MainLogSetFheAddressHashIterator is returned from FilterLogSetFheAddressHash and is used to iterate over the raw logs and unpacked data for LogSetFheAddressHash events raised by the Main contract.
type MainLogSetFheAddressHashIterator struct {
	Event *MainLogSetFheAddressHash // Event containing the contract specifics and raw log

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
func (it *MainLogSetFheAddressHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogSetFheAddressHash)
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
		it.Event = new(MainLogSetFheAddressHash)
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
func (it *MainLogSetFheAddressHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogSetFheAddressHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogSetFheAddressHash represents a LogSetFheAddressHash event raised by the Main contract.
type MainLogSetFheAddressHash struct {
	FheAddressHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogSetFheAddressHash is a free log retrieval operation binding the contract event 0x84cd278fd309da11907c3a07e62f8fb36f0f4a4a1ad30851d0b4c91e7cc8af82.
//
// Solidity: e LogSetFheAddressHash(_fheAddressHash bytes32)
func (_Main *MainFilterer) FilterLogSetFheAddressHash(opts *bind.FilterOpts) (*MainLogSetFheAddressHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogSetFheAddressHash")
	if err != nil {
		return nil, err
	}
	return &MainLogSetFheAddressHashIterator{contract: _Main.contract, event: "LogSetFheAddressHash", logs: logs, sub: sub}, nil
}

// WatchLogSetFheAddressHash is a free log subscription operation binding the contract event 0x84cd278fd309da11907c3a07e62f8fb36f0f4a4a1ad30851d0b4c91e7cc8af82.
//
// Solidity: e LogSetFheAddressHash(_fheAddressHash bytes32)
func (_Main *MainFilterer) WatchLogSetFheAddressHash(opts *bind.WatchOpts, sink chan<- *MainLogSetFheAddressHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogSetFheAddressHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogSetFheAddressHash)
				if err := _Main.contract.UnpackLog(event, "LogSetFheAddressHash", log); err != nil {
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

// MainLogSetFinalDistanceIterator is returned from FilterLogSetFinalDistance and is used to iterate over the raw logs and unpacked data for LogSetFinalDistance events raised by the Main contract.
type MainLogSetFinalDistanceIterator struct {
	Event *MainLogSetFinalDistance // Event containing the contract specifics and raw log

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
func (it *MainLogSetFinalDistanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogSetFinalDistance)
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
		it.Event = new(MainLogSetFinalDistance)
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
func (it *MainLogSetFinalDistanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogSetFinalDistanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogSetFinalDistance represents a LogSetFinalDistance event raised by the Main contract.
type MainLogSetFinalDistance struct {
	FinalDistance [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogSetFinalDistance is a free log retrieval operation binding the contract event 0x9c6f5a9ba93907c83634f959c44c55107e452ca00a217f4d7a4ff515e66c0cd0.
//
// Solidity: e LogSetFinalDistance(_finalDistance bytes32)
func (_Main *MainFilterer) FilterLogSetFinalDistance(opts *bind.FilterOpts) (*MainLogSetFinalDistanceIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogSetFinalDistance")
	if err != nil {
		return nil, err
	}
	return &MainLogSetFinalDistanceIterator{contract: _Main.contract, event: "LogSetFinalDistance", logs: logs, sub: sub}, nil
}

// WatchLogSetFinalDistance is a free log subscription operation binding the contract event 0x9c6f5a9ba93907c83634f959c44c55107e452ca00a217f4d7a4ff515e66c0cd0.
//
// Solidity: e LogSetFinalDistance(_finalDistance bytes32)
func (_Main *MainFilterer) WatchLogSetFinalDistance(opts *bind.WatchOpts, sink chan<- *MainLogSetFinalDistance) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogSetFinalDistance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogSetFinalDistance)
				if err := _Main.contract.UnpackLog(event, "LogSetFinalDistance", log); err != nil {
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

// MainLogSetModelHashIterator is returned from FilterLogSetModelHash and is used to iterate over the raw logs and unpacked data for LogSetModelHash events raised by the Main contract.
type MainLogSetModelHashIterator struct {
	Event *MainLogSetModelHash // Event containing the contract specifics and raw log

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
func (it *MainLogSetModelHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogSetModelHash)
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
		it.Event = new(MainLogSetModelHash)
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
func (it *MainLogSetModelHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogSetModelHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogSetModelHash represents a LogSetModelHash event raised by the Main contract.
type MainLogSetModelHash struct {
	ModelHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetModelHash is a free log retrieval operation binding the contract event 0xe6d3727ea0a293c75dcb4456e3041ff031283e8155fef30f9a350a003859e8c3.
//
// Solidity: e LogSetModelHash(_modelHash bytes32)
func (_Main *MainFilterer) FilterLogSetModelHash(opts *bind.FilterOpts) (*MainLogSetModelHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogSetModelHash")
	if err != nil {
		return nil, err
	}
	return &MainLogSetModelHashIterator{contract: _Main.contract, event: "LogSetModelHash", logs: logs, sub: sub}, nil
}

// WatchLogSetModelHash is a free log subscription operation binding the contract event 0xe6d3727ea0a293c75dcb4456e3041ff031283e8155fef30f9a350a003859e8c3.
//
// Solidity: e LogSetModelHash(_modelHash bytes32)
func (_Main *MainFilterer) WatchLogSetModelHash(opts *bind.WatchOpts, sink chan<- *MainLogSetModelHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogSetModelHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogSetModelHash)
				if err := _Main.contract.UnpackLog(event, "LogSetModelHash", log); err != nil {
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

// MainLogSetModelParamsHashIterator is returned from FilterLogSetModelParamsHash and is used to iterate over the raw logs and unpacked data for LogSetModelParamsHash events raised by the Main contract.
type MainLogSetModelParamsHashIterator struct {
	Event *MainLogSetModelParamsHash // Event containing the contract specifics and raw log

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
func (it *MainLogSetModelParamsHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogSetModelParamsHash)
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
		it.Event = new(MainLogSetModelParamsHash)
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
func (it *MainLogSetModelParamsHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogSetModelParamsHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogSetModelParamsHash represents a LogSetModelParamsHash event raised by the Main contract.
type MainLogSetModelParamsHash struct {
	ModelParamsHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSetModelParamsHash is a free log retrieval operation binding the contract event 0x7034869c2aa7e7265ca374e8476f25d1256880850a27a346a0d3d85d0af44772.
//
// Solidity: e LogSetModelParamsHash(_modelParamsHash bytes32)
func (_Main *MainFilterer) FilterLogSetModelParamsHash(opts *bind.FilterOpts) (*MainLogSetModelParamsHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogSetModelParamsHash")
	if err != nil {
		return nil, err
	}
	return &MainLogSetModelParamsHashIterator{contract: _Main.contract, event: "LogSetModelParamsHash", logs: logs, sub: sub}, nil
}

// WatchLogSetModelParamsHash is a free log subscription operation binding the contract event 0x7034869c2aa7e7265ca374e8476f25d1256880850a27a346a0d3d85d0af44772.
//
// Solidity: e LogSetModelParamsHash(_modelParamsHash bytes32)
func (_Main *MainFilterer) WatchLogSetModelParamsHash(opts *bind.WatchOpts, sink chan<- *MainLogSetModelParamsHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogSetModelParamsHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogSetModelParamsHash)
				if err := _Main.contract.UnpackLog(event, "LogSetModelParamsHash", log); err != nil {
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

// MainLogSetRsaPublicKeyHashIterator is returned from FilterLogSetRsaPublicKeyHash and is used to iterate over the raw logs and unpacked data for LogSetRsaPublicKeyHash events raised by the Main contract.
type MainLogSetRsaPublicKeyHashIterator struct {
	Event *MainLogSetRsaPublicKeyHash // Event containing the contract specifics and raw log

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
func (it *MainLogSetRsaPublicKeyHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogSetRsaPublicKeyHash)
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
		it.Event = new(MainLogSetRsaPublicKeyHash)
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
func (it *MainLogSetRsaPublicKeyHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogSetRsaPublicKeyHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogSetRsaPublicKeyHash represents a LogSetRsaPublicKeyHash event raised by the Main contract.
type MainLogSetRsaPublicKeyHash struct {
	RsaPublicKeyHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogSetRsaPublicKeyHash is a free log retrieval operation binding the contract event 0x0009665638a273ab94b50f85ed606cdb9b6b7118c15fd7c373d7d64e09cefe87.
//
// Solidity: e LogSetRsaPublicKeyHash(_rsaPublicKeyHash bytes32)
func (_Main *MainFilterer) FilterLogSetRsaPublicKeyHash(opts *bind.FilterOpts) (*MainLogSetRsaPublicKeyHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogSetRsaPublicKeyHash")
	if err != nil {
		return nil, err
	}
	return &MainLogSetRsaPublicKeyHashIterator{contract: _Main.contract, event: "LogSetRsaPublicKeyHash", logs: logs, sub: sub}, nil
}

// WatchLogSetRsaPublicKeyHash is a free log subscription operation binding the contract event 0x0009665638a273ab94b50f85ed606cdb9b6b7118c15fd7c373d7d64e09cefe87.
//
// Solidity: e LogSetRsaPublicKeyHash(_rsaPublicKeyHash bytes32)
func (_Main *MainFilterer) WatchLogSetRsaPublicKeyHash(opts *bind.WatchOpts, sink chan<- *MainLogSetRsaPublicKeyHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogSetRsaPublicKeyHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogSetRsaPublicKeyHash)
				if err := _Main.contract.UnpackLog(event, "LogSetRsaPublicKeyHash", log); err != nil {
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

// MainLogTrainedResultHashIterator is returned from FilterLogTrainedResultHash and is used to iterate over the raw logs and unpacked data for LogTrainedResultHash events raised by the Main contract.
type MainLogTrainedResultHashIterator struct {
	Event *MainLogTrainedResultHash // Event containing the contract specifics and raw log

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
func (it *MainLogTrainedResultHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogTrainedResultHash)
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
		it.Event = new(MainLogTrainedResultHash)
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
func (it *MainLogTrainedResultHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogTrainedResultHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogTrainedResultHash represents a LogTrainedResultHash event raised by the Main contract.
type MainLogTrainedResultHash struct {
	TrainedResultHash [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogTrainedResultHash is a free log retrieval operation binding the contract event 0x906e27cc089c73cd1d44fcd513d956b221515cee6e280c87ea18fa3a027c16f9.
//
// Solidity: e LogTrainedResultHash(_trainedResultHash bytes32)
func (_Main *MainFilterer) FilterLogTrainedResultHash(opts *bind.FilterOpts) (*MainLogTrainedResultHashIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogTrainedResultHash")
	if err != nil {
		return nil, err
	}
	return &MainLogTrainedResultHashIterator{contract: _Main.contract, event: "LogTrainedResultHash", logs: logs, sub: sub}, nil
}

// WatchLogTrainedResultHash is a free log subscription operation binding the contract event 0x906e27cc089c73cd1d44fcd513d956b221515cee6e280c87ea18fa3a027c16f9.
//
// Solidity: e LogTrainedResultHash(_trainedResultHash bytes32)
func (_Main *MainFilterer) WatchLogTrainedResultHash(opts *bind.WatchOpts, sink chan<- *MainLogTrainedResultHash) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogTrainedResultHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogTrainedResultHash)
				if err := _Main.contract.UnpackLog(event, "LogTrainedResultHash", log); err != nil {
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
