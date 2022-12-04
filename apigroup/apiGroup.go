// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package apigroup

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

// ApigroupMetaData contains all meta data concerning the Apigroup contract.
var ApigroupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lotSnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"BuyLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"}],\"name\":\"CreateRoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"}],\"name\":\"CreatedLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"EnterRoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeFirst\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lotSnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"NewLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_SposDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_SnegDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_psnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"ReceiveLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_receiveTokens\",\"type\":\"uint256\"}],\"name\":\"SendLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_psnap\",\"type\":\"uint256\"}],\"name\":\"StartRoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nwin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spos\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sneg\",\"type\":\"uint256\"}],\"name\":\"UpdatePlayerParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_psnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proofResData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofEPData\",\"type\":\"bytes\"}],\"name\":\"BuyLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CreateLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"}],\"name\":\"CreateRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Enter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initParamsData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofResData\",\"type\":\"bytes\"}],\"name\":\"NewLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initParamsData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofResData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"playerParamsData\",\"type\":\"bytes\"}],\"name\":\"ReceiveLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initParamsData\",\"type\":\"bytes\"}],\"name\":\"SendLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StartRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proofResData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"playerParamsData\",\"type\":\"bytes\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ApigroupABI is the input ABI used to generate the binding from.
// Deprecated: Use ApigroupMetaData.ABI instead.
var ApigroupABI = ApigroupMetaData.ABI

// Apigroup is an auto generated Go binding around an Ethereum contract.
type Apigroup struct {
	ApigroupCaller     // Read-only binding to the contract
	ApigroupTransactor // Write-only binding to the contract
	ApigroupFilterer   // Log filterer for contract events
}

// ApigroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApigroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApigroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApigroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApigroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApigroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApigroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApigroupSession struct {
	Contract     *Apigroup         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApigroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApigroupCallerSession struct {
	Contract *ApigroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ApigroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApigroupTransactorSession struct {
	Contract     *ApigroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ApigroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApigroupRaw struct {
	Contract *Apigroup // Generic contract binding to access the raw methods on
}

// ApigroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApigroupCallerRaw struct {
	Contract *ApigroupCaller // Generic read-only contract binding to access the raw methods on
}

// ApigroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApigroupTransactorRaw struct {
	Contract *ApigroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApigroup creates a new instance of Apigroup, bound to a specific deployed contract.
func NewApigroup(address common.Address, backend bind.ContractBackend) (*Apigroup, error) {
	contract, err := bindApigroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apigroup{ApigroupCaller: ApigroupCaller{contract: contract}, ApigroupTransactor: ApigroupTransactor{contract: contract}, ApigroupFilterer: ApigroupFilterer{contract: contract}}, nil
}

// NewApigroupCaller creates a new read-only instance of Apigroup, bound to a specific deployed contract.
func NewApigroupCaller(address common.Address, caller bind.ContractCaller) (*ApigroupCaller, error) {
	contract, err := bindApigroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApigroupCaller{contract: contract}, nil
}

// NewApigroupTransactor creates a new write-only instance of Apigroup, bound to a specific deployed contract.
func NewApigroupTransactor(address common.Address, transactor bind.ContractTransactor) (*ApigroupTransactor, error) {
	contract, err := bindApigroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApigroupTransactor{contract: contract}, nil
}

// NewApigroupFilterer creates a new log filterer instance of Apigroup, bound to a specific deployed contract.
func NewApigroupFilterer(address common.Address, filterer bind.ContractFilterer) (*ApigroupFilterer, error) {
	contract, err := bindApigroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApigroupFilterer{contract: contract}, nil
}

// bindApigroup binds a generic wrapper to an already deployed contract.
func bindApigroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApigroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apigroup *ApigroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apigroup.Contract.ApigroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apigroup *ApigroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apigroup.Contract.ApigroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apigroup *ApigroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apigroup.Contract.ApigroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apigroup *ApigroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apigroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apigroup *ApigroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apigroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apigroup *ApigroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apigroup.Contract.contract.Transact(opts, method, params...)
}

// BuyLot is a paid mutator transaction binding the contract method 0x5be4a96c.
//
// Solidity: function BuyLot(address _lotAddr, bytes proofResData, bytes proofEPData) returns()
func (_Apigroup *ApigroupTransactor) BuyLot(opts *bind.TransactOpts, _lotAddr common.Address, proofResData []byte, proofEPData []byte) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "BuyLot", _lotAddr, proofResData, proofEPData)
}

// BuyLot is a paid mutator transaction binding the contract method 0x5be4a96c.
//
// Solidity: function BuyLot(address _lotAddr, bytes proofResData, bytes proofEPData) returns()
func (_Apigroup *ApigroupSession) BuyLot(_lotAddr common.Address, proofResData []byte, proofEPData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.BuyLot(&_Apigroup.TransactOpts, _lotAddr, proofResData, proofEPData)
}

// BuyLot is a paid mutator transaction binding the contract method 0x5be4a96c.
//
// Solidity: function BuyLot(address _lotAddr, bytes proofResData, bytes proofEPData) returns()
func (_Apigroup *ApigroupTransactorSession) BuyLot(_lotAddr common.Address, proofResData []byte, proofEPData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.BuyLot(&_Apigroup.TransactOpts, _lotAddr, proofResData, proofEPData)
}

// CreateLot is a paid mutator transaction binding the contract method 0xd53e743e.
//
// Solidity: function CreateLot() returns()
func (_Apigroup *ApigroupTransactor) CreateLot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "CreateLot")
}

// CreateLot is a paid mutator transaction binding the contract method 0xd53e743e.
//
// Solidity: function CreateLot() returns()
func (_Apigroup *ApigroupSession) CreateLot() (*types.Transaction, error) {
	return _Apigroup.Contract.CreateLot(&_Apigroup.TransactOpts)
}

// CreateLot is a paid mutator transaction binding the contract method 0xd53e743e.
//
// Solidity: function CreateLot() returns()
func (_Apigroup *ApigroupTransactorSession) CreateLot() (*types.Transaction, error) {
	return _Apigroup.Contract.CreateLot(&_Apigroup.TransactOpts)
}

// CreateRound is a paid mutator transaction binding the contract method 0x257b6e0c.
//
// Solidity: function CreateRound(uint256 _deposit) returns()
func (_Apigroup *ApigroupTransactor) CreateRound(opts *bind.TransactOpts, _deposit *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "CreateRound", _deposit)
}

// CreateRound is a paid mutator transaction binding the contract method 0x257b6e0c.
//
// Solidity: function CreateRound(uint256 _deposit) returns()
func (_Apigroup *ApigroupSession) CreateRound(_deposit *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.CreateRound(&_Apigroup.TransactOpts, _deposit)
}

// CreateRound is a paid mutator transaction binding the contract method 0x257b6e0c.
//
// Solidity: function CreateRound(uint256 _deposit) returns()
func (_Apigroup *ApigroupTransactorSession) CreateRound(_deposit *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.CreateRound(&_Apigroup.TransactOpts, _deposit)
}

// Enter is a paid mutator transaction binding the contract method 0x1097e579.
//
// Solidity: function Enter() payable returns()
func (_Apigroup *ApigroupTransactor) Enter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "Enter")
}

// Enter is a paid mutator transaction binding the contract method 0x1097e579.
//
// Solidity: function Enter() payable returns()
func (_Apigroup *ApigroupSession) Enter() (*types.Transaction, error) {
	return _Apigroup.Contract.Enter(&_Apigroup.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0x1097e579.
//
// Solidity: function Enter() payable returns()
func (_Apigroup *ApigroupTransactorSession) Enter() (*types.Transaction, error) {
	return _Apigroup.Contract.Enter(&_Apigroup.TransactOpts)
}

// NewLot is a paid mutator transaction binding the contract method 0x70ebfc2d.
//
// Solidity: function NewLot(address _lotAddr, bytes initParamsData, bytes proofResData) returns()
func (_Apigroup *ApigroupTransactor) NewLot(opts *bind.TransactOpts, _lotAddr common.Address, initParamsData []byte, proofResData []byte) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "NewLot", _lotAddr, initParamsData, proofResData)
}

// NewLot is a paid mutator transaction binding the contract method 0x70ebfc2d.
//
// Solidity: function NewLot(address _lotAddr, bytes initParamsData, bytes proofResData) returns()
func (_Apigroup *ApigroupSession) NewLot(_lotAddr common.Address, initParamsData []byte, proofResData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.NewLot(&_Apigroup.TransactOpts, _lotAddr, initParamsData, proofResData)
}

// NewLot is a paid mutator transaction binding the contract method 0x70ebfc2d.
//
// Solidity: function NewLot(address _lotAddr, bytes initParamsData, bytes proofResData) returns()
func (_Apigroup *ApigroupTransactorSession) NewLot(_lotAddr common.Address, initParamsData []byte, proofResData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.NewLot(&_Apigroup.TransactOpts, _lotAddr, initParamsData, proofResData)
}

// ReceiveLot is a paid mutator transaction binding the contract method 0xc9001884.
//
// Solidity: function ReceiveLot(address _lotAddr, address _owner, bytes initParamsData, bytes proofResData, bytes playerParamsData) returns()
func (_Apigroup *ApigroupTransactor) ReceiveLot(opts *bind.TransactOpts, _lotAddr common.Address, _owner common.Address, initParamsData []byte, proofResData []byte, playerParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "ReceiveLot", _lotAddr, _owner, initParamsData, proofResData, playerParamsData)
}

// ReceiveLot is a paid mutator transaction binding the contract method 0xc9001884.
//
// Solidity: function ReceiveLot(address _lotAddr, address _owner, bytes initParamsData, bytes proofResData, bytes playerParamsData) returns()
func (_Apigroup *ApigroupSession) ReceiveLot(_lotAddr common.Address, _owner common.Address, initParamsData []byte, proofResData []byte, playerParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.ReceiveLot(&_Apigroup.TransactOpts, _lotAddr, _owner, initParamsData, proofResData, playerParamsData)
}

// ReceiveLot is a paid mutator transaction binding the contract method 0xc9001884.
//
// Solidity: function ReceiveLot(address _lotAddr, address _owner, bytes initParamsData, bytes proofResData, bytes playerParamsData) returns()
func (_Apigroup *ApigroupTransactorSession) ReceiveLot(_lotAddr common.Address, _owner common.Address, initParamsData []byte, proofResData []byte, playerParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.ReceiveLot(&_Apigroup.TransactOpts, _lotAddr, _owner, initParamsData, proofResData, playerParamsData)
}

// SendLot is a paid mutator transaction binding the contract method 0x5ebb86c0.
//
// Solidity: function SendLot(address _lotAddr, bytes initParamsData) returns()
func (_Apigroup *ApigroupTransactor) SendLot(opts *bind.TransactOpts, _lotAddr common.Address, initParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "SendLot", _lotAddr, initParamsData)
}

// SendLot is a paid mutator transaction binding the contract method 0x5ebb86c0.
//
// Solidity: function SendLot(address _lotAddr, bytes initParamsData) returns()
func (_Apigroup *ApigroupSession) SendLot(_lotAddr common.Address, initParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.SendLot(&_Apigroup.TransactOpts, _lotAddr, initParamsData)
}

// SendLot is a paid mutator transaction binding the contract method 0x5ebb86c0.
//
// Solidity: function SendLot(address _lotAddr, bytes initParamsData) returns()
func (_Apigroup *ApigroupTransactorSession) SendLot(_lotAddr common.Address, initParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.SendLot(&_Apigroup.TransactOpts, _lotAddr, initParamsData)
}

// StartRound is a paid mutator transaction binding the contract method 0xaa7d7568.
//
// Solidity: function StartRound() returns()
func (_Apigroup *ApigroupTransactor) StartRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "StartRound")
}

// StartRound is a paid mutator transaction binding the contract method 0xaa7d7568.
//
// Solidity: function StartRound() returns()
func (_Apigroup *ApigroupSession) StartRound() (*types.Transaction, error) {
	return _Apigroup.Contract.StartRound(&_Apigroup.TransactOpts)
}

// StartRound is a paid mutator transaction binding the contract method 0xaa7d7568.
//
// Solidity: function StartRound() returns()
func (_Apigroup *ApigroupTransactorSession) StartRound() (*types.Transaction, error) {
	return _Apigroup.Contract.StartRound(&_Apigroup.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd224efc8.
//
// Solidity: function Withdraw(bytes proofResData, bytes playerParamsData) returns()
func (_Apigroup *ApigroupTransactor) Withdraw(opts *bind.TransactOpts, proofResData []byte, playerParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "Withdraw", proofResData, playerParamsData)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd224efc8.
//
// Solidity: function Withdraw(bytes proofResData, bytes playerParamsData) returns()
func (_Apigroup *ApigroupSession) Withdraw(proofResData []byte, playerParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.Withdraw(&_Apigroup.TransactOpts, proofResData, playerParamsData)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd224efc8.
//
// Solidity: function Withdraw(bytes proofResData, bytes playerParamsData) returns()
func (_Apigroup *ApigroupTransactorSession) Withdraw(proofResData []byte, playerParamsData []byte) (*types.Transaction, error) {
	return _Apigroup.Contract.Withdraw(&_Apigroup.TransactOpts, proofResData, playerParamsData)
}

// ApigroupBuyLotEventIterator is returned from FilterBuyLotEvent and is used to iterate over the raw logs and unpacked data for BuyLotEvent events raised by the Apigroup contract.
type ApigroupBuyLotEventIterator struct {
	Event *ApigroupBuyLotEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupBuyLotEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupBuyLotEvent)
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
		it.Event = new(ApigroupBuyLotEvent)
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
func (it *ApigroupBuyLotEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupBuyLotEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupBuyLotEvent represents a BuyLotEvent event raised by the Apigroup contract.
type ApigroupBuyLotEvent struct {
	RoundAddress common.Address
	LotAddr      common.Address
	Sender       common.Address
	Price        *big.Int
	LotSnap      *big.Int
	Bsnap        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBuyLotEvent is a free log retrieval operation binding the contract event 0x4fc33a0e6a259fb2ce75e5cbd9fdde7d4eefdaac773afbd1a02c1f2fdb2c66f9.
//
// Solidity: event BuyLotEvent(address _roundAddress, address _lotAddr, address _sender, uint256 _price, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterBuyLotEvent(opts *bind.FilterOpts) (*ApigroupBuyLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "BuyLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupBuyLotEventIterator{contract: _Apigroup.contract, event: "BuyLotEvent", logs: logs, sub: sub}, nil
}

// WatchBuyLotEvent is a free log subscription operation binding the contract event 0x4fc33a0e6a259fb2ce75e5cbd9fdde7d4eefdaac773afbd1a02c1f2fdb2c66f9.
//
// Solidity: event BuyLotEvent(address _roundAddress, address _lotAddr, address _sender, uint256 _price, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) WatchBuyLotEvent(opts *bind.WatchOpts, sink chan<- *ApigroupBuyLotEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "BuyLotEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupBuyLotEvent)
				if err := _Apigroup.contract.UnpackLog(event, "BuyLotEvent", log); err != nil {
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

// ParseBuyLotEvent is a log parse operation binding the contract event 0x4fc33a0e6a259fb2ce75e5cbd9fdde7d4eefdaac773afbd1a02c1f2fdb2c66f9.
//
// Solidity: event BuyLotEvent(address _roundAddress, address _lotAddr, address _sender, uint256 _price, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) ParseBuyLotEvent(log types.Log) (*ApigroupBuyLotEvent, error) {
	event := new(ApigroupBuyLotEvent)
	if err := _Apigroup.contract.UnpackLog(event, "BuyLotEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupCreateRoundEventIterator is returned from FilterCreateRoundEvent and is used to iterate over the raw logs and unpacked data for CreateRoundEvent events raised by the Apigroup contract.
type ApigroupCreateRoundEventIterator struct {
	Event *ApigroupCreateRoundEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupCreateRoundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupCreateRoundEvent)
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
		it.Event = new(ApigroupCreateRoundEvent)
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
func (it *ApigroupCreateRoundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupCreateRoundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupCreateRoundEvent represents a CreateRoundEvent event raised by the Apigroup contract.
type ApigroupCreateRoundEvent struct {
	RoundAddress common.Address
	Deposit      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateRoundEvent is a free log retrieval operation binding the contract event 0xf901c2f537cd421ec196e0b11577d59e041a570ece9cdca8a901ddd7f20b4254.
//
// Solidity: event CreateRoundEvent(address _roundAddress, uint256 _deposit)
func (_Apigroup *ApigroupFilterer) FilterCreateRoundEvent(opts *bind.FilterOpts) (*ApigroupCreateRoundEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "CreateRoundEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupCreateRoundEventIterator{contract: _Apigroup.contract, event: "CreateRoundEvent", logs: logs, sub: sub}, nil
}

// WatchCreateRoundEvent is a free log subscription operation binding the contract event 0xf901c2f537cd421ec196e0b11577d59e041a570ece9cdca8a901ddd7f20b4254.
//
// Solidity: event CreateRoundEvent(address _roundAddress, uint256 _deposit)
func (_Apigroup *ApigroupFilterer) WatchCreateRoundEvent(opts *bind.WatchOpts, sink chan<- *ApigroupCreateRoundEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "CreateRoundEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupCreateRoundEvent)
				if err := _Apigroup.contract.UnpackLog(event, "CreateRoundEvent", log); err != nil {
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

// ParseCreateRoundEvent is a log parse operation binding the contract event 0xf901c2f537cd421ec196e0b11577d59e041a570ece9cdca8a901ddd7f20b4254.
//
// Solidity: event CreateRoundEvent(address _roundAddress, uint256 _deposit)
func (_Apigroup *ApigroupFilterer) ParseCreateRoundEvent(log types.Log) (*ApigroupCreateRoundEvent, error) {
	event := new(ApigroupCreateRoundEvent)
	if err := _Apigroup.contract.UnpackLog(event, "CreateRoundEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupCreatedLotEventIterator is returned from FilterCreatedLotEvent and is used to iterate over the raw logs and unpacked data for CreatedLotEvent events raised by the Apigroup contract.
type ApigroupCreatedLotEventIterator struct {
	Event *ApigroupCreatedLotEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupCreatedLotEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupCreatedLotEvent)
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
		it.Event = new(ApigroupCreatedLotEvent)
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
func (it *ApigroupCreatedLotEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupCreatedLotEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupCreatedLotEvent represents a CreatedLotEvent event raised by the Apigroup contract.
type ApigroupCreatedLotEvent struct {
	RoundAddress common.Address
	LotAddr      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatedLotEvent is a free log retrieval operation binding the contract event 0x918ba15a8dc4906fa0a8ac698245a6a565c4724cf9a27b8be97bf3af6e315035.
//
// Solidity: event CreatedLotEvent(address _roundAddress, address _lotAddr)
func (_Apigroup *ApigroupFilterer) FilterCreatedLotEvent(opts *bind.FilterOpts) (*ApigroupCreatedLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "CreatedLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupCreatedLotEventIterator{contract: _Apigroup.contract, event: "CreatedLotEvent", logs: logs, sub: sub}, nil
}

// WatchCreatedLotEvent is a free log subscription operation binding the contract event 0x918ba15a8dc4906fa0a8ac698245a6a565c4724cf9a27b8be97bf3af6e315035.
//
// Solidity: event CreatedLotEvent(address _roundAddress, address _lotAddr)
func (_Apigroup *ApigroupFilterer) WatchCreatedLotEvent(opts *bind.WatchOpts, sink chan<- *ApigroupCreatedLotEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "CreatedLotEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupCreatedLotEvent)
				if err := _Apigroup.contract.UnpackLog(event, "CreatedLotEvent", log); err != nil {
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

// ParseCreatedLotEvent is a log parse operation binding the contract event 0x918ba15a8dc4906fa0a8ac698245a6a565c4724cf9a27b8be97bf3af6e315035.
//
// Solidity: event CreatedLotEvent(address _roundAddress, address _lotAddr)
func (_Apigroup *ApigroupFilterer) ParseCreatedLotEvent(log types.Log) (*ApigroupCreatedLotEvent, error) {
	event := new(ApigroupCreatedLotEvent)
	if err := _Apigroup.contract.UnpackLog(event, "CreatedLotEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupEnterRoundEventIterator is returned from FilterEnterRoundEvent and is used to iterate over the raw logs and unpacked data for EnterRoundEvent events raised by the Apigroup contract.
type ApigroupEnterRoundEventIterator struct {
	Event *ApigroupEnterRoundEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupEnterRoundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupEnterRoundEvent)
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
		it.Event = new(ApigroupEnterRoundEvent)
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
func (it *ApigroupEnterRoundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupEnterRoundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupEnterRoundEvent represents a EnterRoundEvent event raised by the Apigroup contract.
type ApigroupEnterRoundEvent struct {
	RoundAddress common.Address
	Sender       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEnterRoundEvent is a free log retrieval operation binding the contract event 0xc27ed72edbb571182c3c5aa21b035674fc49c06957440f22d55f7e2bf00f3180.
//
// Solidity: event EnterRoundEvent(address _roundAddress, address _sender)
func (_Apigroup *ApigroupFilterer) FilterEnterRoundEvent(opts *bind.FilterOpts) (*ApigroupEnterRoundEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "EnterRoundEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupEnterRoundEventIterator{contract: _Apigroup.contract, event: "EnterRoundEvent", logs: logs, sub: sub}, nil
}

// WatchEnterRoundEvent is a free log subscription operation binding the contract event 0xc27ed72edbb571182c3c5aa21b035674fc49c06957440f22d55f7e2bf00f3180.
//
// Solidity: event EnterRoundEvent(address _roundAddress, address _sender)
func (_Apigroup *ApigroupFilterer) WatchEnterRoundEvent(opts *bind.WatchOpts, sink chan<- *ApigroupEnterRoundEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "EnterRoundEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupEnterRoundEvent)
				if err := _Apigroup.contract.UnpackLog(event, "EnterRoundEvent", log); err != nil {
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

// ParseEnterRoundEvent is a log parse operation binding the contract event 0xc27ed72edbb571182c3c5aa21b035674fc49c06957440f22d55f7e2bf00f3180.
//
// Solidity: event EnterRoundEvent(address _roundAddress, address _sender)
func (_Apigroup *ApigroupFilterer) ParseEnterRoundEvent(log types.Log) (*ApigroupEnterRoundEvent, error) {
	event := new(ApigroupEnterRoundEvent)
	if err := _Apigroup.contract.UnpackLog(event, "EnterRoundEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupNewLotEventIterator is returned from FilterNewLotEvent and is used to iterate over the raw logs and unpacked data for NewLotEvent events raised by the Apigroup contract.
type ApigroupNewLotEventIterator struct {
	Event *ApigroupNewLotEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupNewLotEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupNewLotEvent)
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
		it.Event = new(ApigroupNewLotEvent)
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
func (it *ApigroupNewLotEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupNewLotEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupNewLotEvent represents a NewLotEvent event raised by the Apigroup contract.
type ApigroupNewLotEvent struct {
	RoundAddress common.Address
	LotAddr      common.Address
	Owner        common.Address
	TimeFirst    *big.Int
	TimeSecond   *big.Int
	Price        *big.Int
	Val          *big.Int
	LotSnap      *big.Int
	Bsnap        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewLotEvent is a free log retrieval operation binding the contract event 0x0ec6421e5db42f49d5bc9b51a59d5babed4e72573d1b56303b98a9496e3be021.
//
// Solidity: event NewLotEvent(address _roundAddress, address _lotAddr, address _owner, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterNewLotEvent(opts *bind.FilterOpts) (*ApigroupNewLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "NewLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupNewLotEventIterator{contract: _Apigroup.contract, event: "NewLotEvent", logs: logs, sub: sub}, nil
}

// WatchNewLotEvent is a free log subscription operation binding the contract event 0x0ec6421e5db42f49d5bc9b51a59d5babed4e72573d1b56303b98a9496e3be021.
//
// Solidity: event NewLotEvent(address _roundAddress, address _lotAddr, address _owner, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) WatchNewLotEvent(opts *bind.WatchOpts, sink chan<- *ApigroupNewLotEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "NewLotEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupNewLotEvent)
				if err := _Apigroup.contract.UnpackLog(event, "NewLotEvent", log); err != nil {
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

// ParseNewLotEvent is a log parse operation binding the contract event 0x0ec6421e5db42f49d5bc9b51a59d5babed4e72573d1b56303b98a9496e3be021.
//
// Solidity: event NewLotEvent(address _roundAddress, address _lotAddr, address _owner, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) ParseNewLotEvent(log types.Log) (*ApigroupNewLotEvent, error) {
	event := new(ApigroupNewLotEvent)
	if err := _Apigroup.contract.UnpackLog(event, "NewLotEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupReceiveLotEventIterator is returned from FilterReceiveLotEvent and is used to iterate over the raw logs and unpacked data for ReceiveLotEvent events raised by the Apigroup contract.
type ApigroupReceiveLotEventIterator struct {
	Event *ApigroupReceiveLotEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupReceiveLotEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupReceiveLotEvent)
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
		it.Event = new(ApigroupReceiveLotEvent)
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
func (it *ApigroupReceiveLotEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupReceiveLotEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupReceiveLotEvent represents a ReceiveLotEvent event raised by the Apigroup contract.
type ApigroupReceiveLotEvent struct {
	RoundAddress common.Address
	LotAddr      common.Address
	Owner        common.Address
	Balance      *big.Int
	SposDelta    *big.Int
	SnegDelta    *big.Int
	Psnap        *big.Int
	Bsnap        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReceiveLotEvent is a free log retrieval operation binding the contract event 0xf961f8b2e1723623c5ae736a74446f0a5b8b0d68f8e42035a217e011b9543fb9.
//
// Solidity: event ReceiveLotEvent(address _roundAddress, address _lotAddr, address _owner, uint256 _balance, uint256 _SposDelta, uint256 _SnegDelta, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterReceiveLotEvent(opts *bind.FilterOpts) (*ApigroupReceiveLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "ReceiveLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupReceiveLotEventIterator{contract: _Apigroup.contract, event: "ReceiveLotEvent", logs: logs, sub: sub}, nil
}

// WatchReceiveLotEvent is a free log subscription operation binding the contract event 0xf961f8b2e1723623c5ae736a74446f0a5b8b0d68f8e42035a217e011b9543fb9.
//
// Solidity: event ReceiveLotEvent(address _roundAddress, address _lotAddr, address _owner, uint256 _balance, uint256 _SposDelta, uint256 _SnegDelta, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) WatchReceiveLotEvent(opts *bind.WatchOpts, sink chan<- *ApigroupReceiveLotEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "ReceiveLotEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupReceiveLotEvent)
				if err := _Apigroup.contract.UnpackLog(event, "ReceiveLotEvent", log); err != nil {
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

// ParseReceiveLotEvent is a log parse operation binding the contract event 0xf961f8b2e1723623c5ae736a74446f0a5b8b0d68f8e42035a217e011b9543fb9.
//
// Solidity: event ReceiveLotEvent(address _roundAddress, address _lotAddr, address _owner, uint256 _balance, uint256 _SposDelta, uint256 _SnegDelta, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) ParseReceiveLotEvent(log types.Log) (*ApigroupReceiveLotEvent, error) {
	event := new(ApigroupReceiveLotEvent)
	if err := _Apigroup.contract.UnpackLog(event, "ReceiveLotEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupSendLotEventIterator is returned from FilterSendLotEvent and is used to iterate over the raw logs and unpacked data for SendLotEvent events raised by the Apigroup contract.
type ApigroupSendLotEventIterator struct {
	Event *ApigroupSendLotEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupSendLotEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupSendLotEvent)
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
		it.Event = new(ApigroupSendLotEvent)
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
func (it *ApigroupSendLotEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupSendLotEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupSendLotEvent represents a SendLotEvent event raised by the Apigroup contract.
type ApigroupSendLotEvent struct {
	RoundAddress  common.Address
	LotAddr       common.Address
	ReceiveTokens *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendLotEvent is a free log retrieval operation binding the contract event 0x7c63f4ca5f3947a628e97b9822a09acf9382a1c3916ccbf7493097dc2f081ecf.
//
// Solidity: event SendLotEvent(address _roundAddress, address _lotAddr, uint256 _receiveTokens)
func (_Apigroup *ApigroupFilterer) FilterSendLotEvent(opts *bind.FilterOpts) (*ApigroupSendLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "SendLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupSendLotEventIterator{contract: _Apigroup.contract, event: "SendLotEvent", logs: logs, sub: sub}, nil
}

// WatchSendLotEvent is a free log subscription operation binding the contract event 0x7c63f4ca5f3947a628e97b9822a09acf9382a1c3916ccbf7493097dc2f081ecf.
//
// Solidity: event SendLotEvent(address _roundAddress, address _lotAddr, uint256 _receiveTokens)
func (_Apigroup *ApigroupFilterer) WatchSendLotEvent(opts *bind.WatchOpts, sink chan<- *ApigroupSendLotEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "SendLotEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupSendLotEvent)
				if err := _Apigroup.contract.UnpackLog(event, "SendLotEvent", log); err != nil {
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

// ParseSendLotEvent is a log parse operation binding the contract event 0x7c63f4ca5f3947a628e97b9822a09acf9382a1c3916ccbf7493097dc2f081ecf.
//
// Solidity: event SendLotEvent(address _roundAddress, address _lotAddr, uint256 _receiveTokens)
func (_Apigroup *ApigroupFilterer) ParseSendLotEvent(log types.Log) (*ApigroupSendLotEvent, error) {
	event := new(ApigroupSendLotEvent)
	if err := _Apigroup.contract.UnpackLog(event, "SendLotEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupStartRoundEventIterator is returned from FilterStartRoundEvent and is used to iterate over the raw logs and unpacked data for StartRoundEvent events raised by the Apigroup contract.
type ApigroupStartRoundEventIterator struct {
	Event *ApigroupStartRoundEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupStartRoundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupStartRoundEvent)
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
		it.Event = new(ApigroupStartRoundEvent)
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
func (it *ApigroupStartRoundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupStartRoundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupStartRoundEvent represents a StartRoundEvent event raised by the Apigroup contract.
type ApigroupStartRoundEvent struct {
	RoundAddress common.Address
	Bsnap        *big.Int
	Psnap        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStartRoundEvent is a free log retrieval operation binding the contract event 0x87d0d34a259a5bf664a28dc7fdadc47f8abc8fe0f896b1356c050f6d5773418c.
//
// Solidity: event StartRoundEvent(address _roundAddress, uint256 _bsnap, uint256 _psnap)
func (_Apigroup *ApigroupFilterer) FilterStartRoundEvent(opts *bind.FilterOpts) (*ApigroupStartRoundEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "StartRoundEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupStartRoundEventIterator{contract: _Apigroup.contract, event: "StartRoundEvent", logs: logs, sub: sub}, nil
}

// WatchStartRoundEvent is a free log subscription operation binding the contract event 0x87d0d34a259a5bf664a28dc7fdadc47f8abc8fe0f896b1356c050f6d5773418c.
//
// Solidity: event StartRoundEvent(address _roundAddress, uint256 _bsnap, uint256 _psnap)
func (_Apigroup *ApigroupFilterer) WatchStartRoundEvent(opts *bind.WatchOpts, sink chan<- *ApigroupStartRoundEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "StartRoundEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupStartRoundEvent)
				if err := _Apigroup.contract.UnpackLog(event, "StartRoundEvent", log); err != nil {
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

// ParseStartRoundEvent is a log parse operation binding the contract event 0x87d0d34a259a5bf664a28dc7fdadc47f8abc8fe0f896b1356c050f6d5773418c.
//
// Solidity: event StartRoundEvent(address _roundAddress, uint256 _bsnap, uint256 _psnap)
func (_Apigroup *ApigroupFilterer) ParseStartRoundEvent(log types.Log) (*ApigroupStartRoundEvent, error) {
	event := new(ApigroupStartRoundEvent)
	if err := _Apigroup.contract.UnpackLog(event, "StartRoundEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupUpdatePlayerParamsIterator is returned from FilterUpdatePlayerParams and is used to iterate over the raw logs and unpacked data for UpdatePlayerParams events raised by the Apigroup contract.
type ApigroupUpdatePlayerParamsIterator struct {
	Event *ApigroupUpdatePlayerParams // Event containing the contract specifics and raw log

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
func (it *ApigroupUpdatePlayerParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupUpdatePlayerParams)
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
		it.Event = new(ApigroupUpdatePlayerParams)
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
func (it *ApigroupUpdatePlayerParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupUpdatePlayerParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupUpdatePlayerParams represents a UpdatePlayerParams event raised by the Apigroup contract.
type ApigroupUpdatePlayerParams struct {
	RoundAddress common.Address
	Owner        common.Address
	Nwin         *big.Int
	N            *big.Int
	Spos         *big.Int
	Sneg         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlayerParams is a free log retrieval operation binding the contract event 0x549bf9b2e230d5c7984eceaafcffe0407f05426c6cc61641bad8e2cd415ff461.
//
// Solidity: event UpdatePlayerParams(address _roundAddress, address _owner, uint256 _nwin, uint256 _n, uint256 _spos, uint256 _sneg)
func (_Apigroup *ApigroupFilterer) FilterUpdatePlayerParams(opts *bind.FilterOpts) (*ApigroupUpdatePlayerParamsIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "UpdatePlayerParams")
	if err != nil {
		return nil, err
	}
	return &ApigroupUpdatePlayerParamsIterator{contract: _Apigroup.contract, event: "UpdatePlayerParams", logs: logs, sub: sub}, nil
}

// WatchUpdatePlayerParams is a free log subscription operation binding the contract event 0x549bf9b2e230d5c7984eceaafcffe0407f05426c6cc61641bad8e2cd415ff461.
//
// Solidity: event UpdatePlayerParams(address _roundAddress, address _owner, uint256 _nwin, uint256 _n, uint256 _spos, uint256 _sneg)
func (_Apigroup *ApigroupFilterer) WatchUpdatePlayerParams(opts *bind.WatchOpts, sink chan<- *ApigroupUpdatePlayerParams) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "UpdatePlayerParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupUpdatePlayerParams)
				if err := _Apigroup.contract.UnpackLog(event, "UpdatePlayerParams", log); err != nil {
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

// ParseUpdatePlayerParams is a log parse operation binding the contract event 0x549bf9b2e230d5c7984eceaafcffe0407f05426c6cc61641bad8e2cd415ff461.
//
// Solidity: event UpdatePlayerParams(address _roundAddress, address _owner, uint256 _nwin, uint256 _n, uint256 _spos, uint256 _sneg)
func (_Apigroup *ApigroupFilterer) ParseUpdatePlayerParams(log types.Log) (*ApigroupUpdatePlayerParams, error) {
	event := new(ApigroupUpdatePlayerParams)
	if err := _Apigroup.contract.UnpackLog(event, "UpdatePlayerParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Apigroup contract.
type ApigroupWithdrawEventIterator struct {
	Event *ApigroupWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *ApigroupWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupWithdrawEvent)
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
		it.Event = new(ApigroupWithdrawEvent)
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
func (it *ApigroupWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupWithdrawEvent represents a WithdrawEvent event raised by the Apigroup contract.
type ApigroupWithdrawEvent struct {
	Sender common.Address
	Psnap  *big.Int
	Bsnap  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address _sender, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterWithdrawEvent(opts *bind.FilterOpts) (*ApigroupWithdrawEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupWithdrawEventIterator{contract: _Apigroup.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address _sender, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *ApigroupWithdrawEvent) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupWithdrawEvent)
				if err := _Apigroup.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// ParseWithdrawEvent is a log parse operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address _sender, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) ParseWithdrawEvent(log types.Log) (*ApigroupWithdrawEvent, error) {
	event := new(ApigroupWithdrawEvent)
	if err := _Apigroup.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
