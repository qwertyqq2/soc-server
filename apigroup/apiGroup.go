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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lotSnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"BuyLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_roundAddress\",\"type\":\"address\"}],\"name\":\"CreateRoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"}],\"name\":\"CreatedLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"EnterRoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeFirst\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lotSnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"NewLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"Price1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"name\":\"Price2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_psnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"}],\"name\":\"ReceiveLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_receiveTokens\",\"type\":\"uint256\"}],\"name\":\"SendLotEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bsnap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_psnap\",\"type\":\"uint256\"}],\"name\":\"StartRoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nwin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spos\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sneg\",\"type\":\"uint256\"}],\"name\":\"UpdatePlayerParams\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_Hres\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_Hd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevBalance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_prevPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prevSnap\",\"type\":\"uint256\"}],\"name\":\"BuyLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CreateLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"}],\"name\":\"CreateRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"CurrentPrice1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"CurrentPrice2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Enter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetDepositRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetInitSnapRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetParamsSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRound\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSnap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSnapRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timeFirst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_Hres\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"NewLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initParamsData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofResData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"playerParamsData\",\"type\":\"bytes\"}],\"name\":\"ReceiveLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lotAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initParamsData\",\"type\":\"bytes\"}],\"name\":\"SendLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StartRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"SwapDaiToEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"SwapEthToDai\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"playerParamsData\",\"type\":\"bytes\"}],\"name\":\"VerifyParamsPlayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_H\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"VerifyProofRes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetDepositRound is a free data retrieval call binding the contract method 0x571b268b.
//
// Solidity: function GetDepositRound() view returns(uint256, uint256)
func (_Apigroup *ApigroupCaller) GetDepositRound(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "GetDepositRound")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDepositRound is a free data retrieval call binding the contract method 0x571b268b.
//
// Solidity: function GetDepositRound() view returns(uint256, uint256)
func (_Apigroup *ApigroupSession) GetDepositRound() (*big.Int, *big.Int, error) {
	return _Apigroup.Contract.GetDepositRound(&_Apigroup.CallOpts)
}

// GetDepositRound is a free data retrieval call binding the contract method 0x571b268b.
//
// Solidity: function GetDepositRound() view returns(uint256, uint256)
func (_Apigroup *ApigroupCallerSession) GetDepositRound() (*big.Int, *big.Int, error) {
	return _Apigroup.Contract.GetDepositRound(&_Apigroup.CallOpts)
}

// GetInitSnapRound is a free data retrieval call binding the contract method 0xa06a032d.
//
// Solidity: function GetInitSnapRound() view returns(uint256)
func (_Apigroup *ApigroupCaller) GetInitSnapRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "GetInitSnapRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitSnapRound is a free data retrieval call binding the contract method 0xa06a032d.
//
// Solidity: function GetInitSnapRound() view returns(uint256)
func (_Apigroup *ApigroupSession) GetInitSnapRound() (*big.Int, error) {
	return _Apigroup.Contract.GetInitSnapRound(&_Apigroup.CallOpts)
}

// GetInitSnapRound is a free data retrieval call binding the contract method 0xa06a032d.
//
// Solidity: function GetInitSnapRound() view returns(uint256)
func (_Apigroup *ApigroupCallerSession) GetInitSnapRound() (*big.Int, error) {
	return _Apigroup.Contract.GetInitSnapRound(&_Apigroup.CallOpts)
}

// GetParamsSnapshot is a free data retrieval call binding the contract method 0xa88f19ad.
//
// Solidity: function GetParamsSnapshot() view returns(uint256)
func (_Apigroup *ApigroupCaller) GetParamsSnapshot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "GetParamsSnapshot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParamsSnapshot is a free data retrieval call binding the contract method 0xa88f19ad.
//
// Solidity: function GetParamsSnapshot() view returns(uint256)
func (_Apigroup *ApigroupSession) GetParamsSnapshot() (*big.Int, error) {
	return _Apigroup.Contract.GetParamsSnapshot(&_Apigroup.CallOpts)
}

// GetParamsSnapshot is a free data retrieval call binding the contract method 0xa88f19ad.
//
// Solidity: function GetParamsSnapshot() view returns(uint256)
func (_Apigroup *ApigroupCallerSession) GetParamsSnapshot() (*big.Int, error) {
	return _Apigroup.Contract.GetParamsSnapshot(&_Apigroup.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0xfdf2c5b5.
//
// Solidity: function GetRound() view returns(address)
func (_Apigroup *ApigroupCaller) GetRound(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "GetRound")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0xfdf2c5b5.
//
// Solidity: function GetRound() view returns(address)
func (_Apigroup *ApigroupSession) GetRound() (common.Address, error) {
	return _Apigroup.Contract.GetRound(&_Apigroup.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0xfdf2c5b5.
//
// Solidity: function GetRound() view returns(address)
func (_Apigroup *ApigroupCallerSession) GetRound() (common.Address, error) {
	return _Apigroup.Contract.GetRound(&_Apigroup.CallOpts)
}

// GetSnap is a free data retrieval call binding the contract method 0x259fd68e.
//
// Solidity: function GetSnap() view returns(uint256)
func (_Apigroup *ApigroupCaller) GetSnap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "GetSnap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSnap is a free data retrieval call binding the contract method 0x259fd68e.
//
// Solidity: function GetSnap() view returns(uint256)
func (_Apigroup *ApigroupSession) GetSnap() (*big.Int, error) {
	return _Apigroup.Contract.GetSnap(&_Apigroup.CallOpts)
}

// GetSnap is a free data retrieval call binding the contract method 0x259fd68e.
//
// Solidity: function GetSnap() view returns(uint256)
func (_Apigroup *ApigroupCallerSession) GetSnap() (*big.Int, error) {
	return _Apigroup.Contract.GetSnap(&_Apigroup.CallOpts)
}

// GetSnapRound is a free data retrieval call binding the contract method 0xf3ea6846.
//
// Solidity: function GetSnapRound() view returns(uint256)
func (_Apigroup *ApigroupCaller) GetSnapRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "GetSnapRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSnapRound is a free data retrieval call binding the contract method 0xf3ea6846.
//
// Solidity: function GetSnapRound() view returns(uint256)
func (_Apigroup *ApigroupSession) GetSnapRound() (*big.Int, error) {
	return _Apigroup.Contract.GetSnapRound(&_Apigroup.CallOpts)
}

// GetSnapRound is a free data retrieval call binding the contract method 0xf3ea6846.
//
// Solidity: function GetSnapRound() view returns(uint256)
func (_Apigroup *ApigroupCallerSession) GetSnapRound() (*big.Int, error) {
	return _Apigroup.Contract.GetSnapRound(&_Apigroup.CallOpts)
}

// VerifyParamsPlayer is a free data retrieval call binding the contract method 0x085fe405.
//
// Solidity: function VerifyParamsPlayer(bytes playerParamsData) view returns(bool)
func (_Apigroup *ApigroupCaller) VerifyParamsPlayer(opts *bind.CallOpts, playerParamsData []byte) (bool, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "VerifyParamsPlayer", playerParamsData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyParamsPlayer is a free data retrieval call binding the contract method 0x085fe405.
//
// Solidity: function VerifyParamsPlayer(bytes playerParamsData) view returns(bool)
func (_Apigroup *ApigroupSession) VerifyParamsPlayer(playerParamsData []byte) (bool, error) {
	return _Apigroup.Contract.VerifyParamsPlayer(&_Apigroup.CallOpts, playerParamsData)
}

// VerifyParamsPlayer is a free data retrieval call binding the contract method 0x085fe405.
//
// Solidity: function VerifyParamsPlayer(bytes playerParamsData) view returns(bool)
func (_Apigroup *ApigroupCallerSession) VerifyParamsPlayer(playerParamsData []byte) (bool, error) {
	return _Apigroup.Contract.VerifyParamsPlayer(&_Apigroup.CallOpts, playerParamsData)
}

// VerifyProofRes is a free data retrieval call binding the contract method 0x31435bbd.
//
// Solidity: function VerifyProofRes(address _addr, uint256 _H, uint256 _balance) view returns(bool)
func (_Apigroup *ApigroupCaller) VerifyProofRes(opts *bind.CallOpts, _addr common.Address, _H *big.Int, _balance *big.Int) (bool, error) {
	var out []interface{}
	err := _Apigroup.contract.Call(opts, &out, "VerifyProofRes", _addr, _H, _balance)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofRes is a free data retrieval call binding the contract method 0x31435bbd.
//
// Solidity: function VerifyProofRes(address _addr, uint256 _H, uint256 _balance) view returns(bool)
func (_Apigroup *ApigroupSession) VerifyProofRes(_addr common.Address, _H *big.Int, _balance *big.Int) (bool, error) {
	return _Apigroup.Contract.VerifyProofRes(&_Apigroup.CallOpts, _addr, _H, _balance)
}

// VerifyProofRes is a free data retrieval call binding the contract method 0x31435bbd.
//
// Solidity: function VerifyProofRes(address _addr, uint256 _H, uint256 _balance) view returns(bool)
func (_Apigroup *ApigroupCallerSession) VerifyProofRes(_addr common.Address, _H *big.Int, _balance *big.Int) (bool, error) {
	return _Apigroup.Contract.VerifyProofRes(&_Apigroup.CallOpts, _addr, _H, _balance)
}

// BuyLot is a paid mutator transaction binding the contract method 0xcd0c467e.
//
// Solidity: function BuyLot(address _lotAddr, uint256 _price, uint256 _Hres, uint256 _Hd, uint256 _balance, uint256 _prevBalance, address _prevOwner, uint256 _prevPrice, uint256 _prevSnap) returns()
func (_Apigroup *ApigroupTransactor) BuyLot(opts *bind.TransactOpts, _lotAddr common.Address, _price *big.Int, _Hres *big.Int, _Hd *big.Int, _balance *big.Int, _prevBalance *big.Int, _prevOwner common.Address, _prevPrice *big.Int, _prevSnap *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "BuyLot", _lotAddr, _price, _Hres, _Hd, _balance, _prevBalance, _prevOwner, _prevPrice, _prevSnap)
}

// BuyLot is a paid mutator transaction binding the contract method 0xcd0c467e.
//
// Solidity: function BuyLot(address _lotAddr, uint256 _price, uint256 _Hres, uint256 _Hd, uint256 _balance, uint256 _prevBalance, address _prevOwner, uint256 _prevPrice, uint256 _prevSnap) returns()
func (_Apigroup *ApigroupSession) BuyLot(_lotAddr common.Address, _price *big.Int, _Hres *big.Int, _Hd *big.Int, _balance *big.Int, _prevBalance *big.Int, _prevOwner common.Address, _prevPrice *big.Int, _prevSnap *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.BuyLot(&_Apigroup.TransactOpts, _lotAddr, _price, _Hres, _Hd, _balance, _prevBalance, _prevOwner, _prevPrice, _prevSnap)
}

// BuyLot is a paid mutator transaction binding the contract method 0xcd0c467e.
//
// Solidity: function BuyLot(address _lotAddr, uint256 _price, uint256 _Hres, uint256 _Hd, uint256 _balance, uint256 _prevBalance, address _prevOwner, uint256 _prevPrice, uint256 _prevSnap) returns()
func (_Apigroup *ApigroupTransactorSession) BuyLot(_lotAddr common.Address, _price *big.Int, _Hres *big.Int, _Hd *big.Int, _balance *big.Int, _prevBalance *big.Int, _prevOwner common.Address, _prevPrice *big.Int, _prevSnap *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.BuyLot(&_Apigroup.TransactOpts, _lotAddr, _price, _Hres, _Hd, _balance, _prevBalance, _prevOwner, _prevPrice, _prevSnap)
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

// CurrentPrice1 is a paid mutator transaction binding the contract method 0x2888e6f9.
//
// Solidity: function CurrentPrice1(uint256 _amountIn) returns()
func (_Apigroup *ApigroupTransactor) CurrentPrice1(opts *bind.TransactOpts, _amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "CurrentPrice1", _amountIn)
}

// CurrentPrice1 is a paid mutator transaction binding the contract method 0x2888e6f9.
//
// Solidity: function CurrentPrice1(uint256 _amountIn) returns()
func (_Apigroup *ApigroupSession) CurrentPrice1(_amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.CurrentPrice1(&_Apigroup.TransactOpts, _amountIn)
}

// CurrentPrice1 is a paid mutator transaction binding the contract method 0x2888e6f9.
//
// Solidity: function CurrentPrice1(uint256 _amountIn) returns()
func (_Apigroup *ApigroupTransactorSession) CurrentPrice1(_amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.CurrentPrice1(&_Apigroup.TransactOpts, _amountIn)
}

// CurrentPrice2 is a paid mutator transaction binding the contract method 0xcda283c7.
//
// Solidity: function CurrentPrice2(uint256 _amountIn) returns()
func (_Apigroup *ApigroupTransactor) CurrentPrice2(opts *bind.TransactOpts, _amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "CurrentPrice2", _amountIn)
}

// CurrentPrice2 is a paid mutator transaction binding the contract method 0xcda283c7.
//
// Solidity: function CurrentPrice2(uint256 _amountIn) returns()
func (_Apigroup *ApigroupSession) CurrentPrice2(_amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.CurrentPrice2(&_Apigroup.TransactOpts, _amountIn)
}

// CurrentPrice2 is a paid mutator transaction binding the contract method 0xcda283c7.
//
// Solidity: function CurrentPrice2(uint256 _amountIn) returns()
func (_Apigroup *ApigroupTransactorSession) CurrentPrice2(_amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.CurrentPrice2(&_Apigroup.TransactOpts, _amountIn)
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

// NewLot is a paid mutator transaction binding the contract method 0x4ef625b8.
//
// Solidity: function NewLot(address _lotAddr, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _Hres, uint256 _balance) returns()
func (_Apigroup *ApigroupTransactor) NewLot(opts *bind.TransactOpts, _lotAddr common.Address, _timeFirst *big.Int, _timeSecond *big.Int, _price *big.Int, _val *big.Int, _Hres *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "NewLot", _lotAddr, _timeFirst, _timeSecond, _price, _val, _Hres, _balance)
}

// NewLot is a paid mutator transaction binding the contract method 0x4ef625b8.
//
// Solidity: function NewLot(address _lotAddr, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _Hres, uint256 _balance) returns()
func (_Apigroup *ApigroupSession) NewLot(_lotAddr common.Address, _timeFirst *big.Int, _timeSecond *big.Int, _price *big.Int, _val *big.Int, _Hres *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.NewLot(&_Apigroup.TransactOpts, _lotAddr, _timeFirst, _timeSecond, _price, _val, _Hres, _balance)
}

// NewLot is a paid mutator transaction binding the contract method 0x4ef625b8.
//
// Solidity: function NewLot(address _lotAddr, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _Hres, uint256 _balance) returns()
func (_Apigroup *ApigroupTransactorSession) NewLot(_lotAddr common.Address, _timeFirst *big.Int, _timeSecond *big.Int, _price *big.Int, _val *big.Int, _Hres *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.NewLot(&_Apigroup.TransactOpts, _lotAddr, _timeFirst, _timeSecond, _price, _val, _Hres, _balance)
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

// SwapDaiToEth is a paid mutator transaction binding the contract method 0xc70312ce.
//
// Solidity: function SwapDaiToEth(uint256 amountIn) returns()
func (_Apigroup *ApigroupTransactor) SwapDaiToEth(opts *bind.TransactOpts, amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "SwapDaiToEth", amountIn)
}

// SwapDaiToEth is a paid mutator transaction binding the contract method 0xc70312ce.
//
// Solidity: function SwapDaiToEth(uint256 amountIn) returns()
func (_Apigroup *ApigroupSession) SwapDaiToEth(amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.SwapDaiToEth(&_Apigroup.TransactOpts, amountIn)
}

// SwapDaiToEth is a paid mutator transaction binding the contract method 0xc70312ce.
//
// Solidity: function SwapDaiToEth(uint256 amountIn) returns()
func (_Apigroup *ApigroupTransactorSession) SwapDaiToEth(amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.SwapDaiToEth(&_Apigroup.TransactOpts, amountIn)
}

// SwapEthToDai is a paid mutator transaction binding the contract method 0x79facaee.
//
// Solidity: function SwapEthToDai(uint256 amountIn) returns()
func (_Apigroup *ApigroupTransactor) SwapEthToDai(opts *bind.TransactOpts, amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.contract.Transact(opts, "SwapEthToDai", amountIn)
}

// SwapEthToDai is a paid mutator transaction binding the contract method 0x79facaee.
//
// Solidity: function SwapEthToDai(uint256 amountIn) returns()
func (_Apigroup *ApigroupSession) SwapEthToDai(amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.SwapEthToDai(&_Apigroup.TransactOpts, amountIn)
}

// SwapEthToDai is a paid mutator transaction binding the contract method 0x79facaee.
//
// Solidity: function SwapEthToDai(uint256 amountIn) returns()
func (_Apigroup *ApigroupTransactorSession) SwapEthToDai(amountIn *big.Int) (*types.Transaction, error) {
	return _Apigroup.Contract.SwapEthToDai(&_Apigroup.TransactOpts, amountIn)
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
	LotAddr common.Address
	Sender  common.Address
	Price   *big.Int
	LotSnap *big.Int
	Bsnap   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyLotEvent is a free log retrieval operation binding the contract event 0xd7cf0720ef557d2e4e0050f02ea656eba1644e7acc538778f2eee03fa2820559.
//
// Solidity: event BuyLotEvent(address _lotAddr, address _sender, uint256 _price, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterBuyLotEvent(opts *bind.FilterOpts) (*ApigroupBuyLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "BuyLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupBuyLotEventIterator{contract: _Apigroup.contract, event: "BuyLotEvent", logs: logs, sub: sub}, nil
}

// WatchBuyLotEvent is a free log subscription operation binding the contract event 0xd7cf0720ef557d2e4e0050f02ea656eba1644e7acc538778f2eee03fa2820559.
//
// Solidity: event BuyLotEvent(address _lotAddr, address _sender, uint256 _price, uint256 _lotSnap, uint256 _bsnap)
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

// ParseBuyLotEvent is a log parse operation binding the contract event 0xd7cf0720ef557d2e4e0050f02ea656eba1644e7acc538778f2eee03fa2820559.
//
// Solidity: event BuyLotEvent(address _lotAddr, address _sender, uint256 _price, uint256 _lotSnap, uint256 _bsnap)
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
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateRoundEvent is a free log retrieval operation binding the contract event 0x207317a9a3c9d67474683b98d588e6a3d6a34b201fdd24ab9ac44e6b7ee6b321.
//
// Solidity: event CreateRoundEvent(address _roundAddress)
func (_Apigroup *ApigroupFilterer) FilterCreateRoundEvent(opts *bind.FilterOpts) (*ApigroupCreateRoundEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "CreateRoundEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupCreateRoundEventIterator{contract: _Apigroup.contract, event: "CreateRoundEvent", logs: logs, sub: sub}, nil
}

// WatchCreateRoundEvent is a free log subscription operation binding the contract event 0x207317a9a3c9d67474683b98d588e6a3d6a34b201fdd24ab9ac44e6b7ee6b321.
//
// Solidity: event CreateRoundEvent(address _roundAddress)
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

// ParseCreateRoundEvent is a log parse operation binding the contract event 0x207317a9a3c9d67474683b98d588e6a3d6a34b201fdd24ab9ac44e6b7ee6b321.
//
// Solidity: event CreateRoundEvent(address _roundAddress)
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
	LotAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatedLotEvent is a free log retrieval operation binding the contract event 0x30615d585e73fd9de5159f10134e8b21a24bcbc07b2659223a82e87b4bad21c8.
//
// Solidity: event CreatedLotEvent(address _lotAddr)
func (_Apigroup *ApigroupFilterer) FilterCreatedLotEvent(opts *bind.FilterOpts) (*ApigroupCreatedLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "CreatedLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupCreatedLotEventIterator{contract: _Apigroup.contract, event: "CreatedLotEvent", logs: logs, sub: sub}, nil
}

// WatchCreatedLotEvent is a free log subscription operation binding the contract event 0x30615d585e73fd9de5159f10134e8b21a24bcbc07b2659223a82e87b4bad21c8.
//
// Solidity: event CreatedLotEvent(address _lotAddr)
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

// ParseCreatedLotEvent is a log parse operation binding the contract event 0x30615d585e73fd9de5159f10134e8b21a24bcbc07b2659223a82e87b4bad21c8.
//
// Solidity: event CreatedLotEvent(address _lotAddr)
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
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterRoundEvent is a free log retrieval operation binding the contract event 0x68dc3ce5e812fe17247634e9cb61ecf3e69ce8801ab6e08d501c5a38eff9ca15.
//
// Solidity: event EnterRoundEvent(address _sender)
func (_Apigroup *ApigroupFilterer) FilterEnterRoundEvent(opts *bind.FilterOpts) (*ApigroupEnterRoundEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "EnterRoundEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupEnterRoundEventIterator{contract: _Apigroup.contract, event: "EnterRoundEvent", logs: logs, sub: sub}, nil
}

// WatchEnterRoundEvent is a free log subscription operation binding the contract event 0x68dc3ce5e812fe17247634e9cb61ecf3e69ce8801ab6e08d501c5a38eff9ca15.
//
// Solidity: event EnterRoundEvent(address _sender)
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

// ParseEnterRoundEvent is a log parse operation binding the contract event 0x68dc3ce5e812fe17247634e9cb61ecf3e69ce8801ab6e08d501c5a38eff9ca15.
//
// Solidity: event EnterRoundEvent(address _sender)
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
	LotAddr    common.Address
	TimeFirst  *big.Int
	TimeSecond *big.Int
	Price      *big.Int
	Val        *big.Int
	LotSnap    *big.Int
	Bsnap      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewLotEvent is a free log retrieval operation binding the contract event 0x23fd0571e24a340c152acab101a9ae991b90a575893eb90bf89ebfddbe895478.
//
// Solidity: event NewLotEvent(address _lotAddr, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterNewLotEvent(opts *bind.FilterOpts) (*ApigroupNewLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "NewLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupNewLotEventIterator{contract: _Apigroup.contract, event: "NewLotEvent", logs: logs, sub: sub}, nil
}

// WatchNewLotEvent is a free log subscription operation binding the contract event 0x23fd0571e24a340c152acab101a9ae991b90a575893eb90bf89ebfddbe895478.
//
// Solidity: event NewLotEvent(address _lotAddr, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _lotSnap, uint256 _bsnap)
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

// ParseNewLotEvent is a log parse operation binding the contract event 0x23fd0571e24a340c152acab101a9ae991b90a575893eb90bf89ebfddbe895478.
//
// Solidity: event NewLotEvent(address _lotAddr, uint256 _timeFirst, uint256 _timeSecond, uint256 _price, uint256 _val, uint256 _lotSnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) ParseNewLotEvent(log types.Log) (*ApigroupNewLotEvent, error) {
	event := new(ApigroupNewLotEvent)
	if err := _Apigroup.contract.UnpackLog(event, "NewLotEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupPrice1Iterator is returned from FilterPrice1 and is used to iterate over the raw logs and unpacked data for Price1 events raised by the Apigroup contract.
type ApigroupPrice1Iterator struct {
	Event *ApigroupPrice1 // Event containing the contract specifics and raw log

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
func (it *ApigroupPrice1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupPrice1)
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
		it.Event = new(ApigroupPrice1)
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
func (it *ApigroupPrice1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupPrice1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupPrice1 represents a Price1 event raised by the Apigroup contract.
type ApigroupPrice1 struct {
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrice1 is a free log retrieval operation binding the contract event 0x0737c274963b06ee8d1f3269a1e91e7d372d7b2d31ddd86fa11e87158ef61f16.
//
// Solidity: event Price1(uint256 _amountOut)
func (_Apigroup *ApigroupFilterer) FilterPrice1(opts *bind.FilterOpts) (*ApigroupPrice1Iterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "Price1")
	if err != nil {
		return nil, err
	}
	return &ApigroupPrice1Iterator{contract: _Apigroup.contract, event: "Price1", logs: logs, sub: sub}, nil
}

// WatchPrice1 is a free log subscription operation binding the contract event 0x0737c274963b06ee8d1f3269a1e91e7d372d7b2d31ddd86fa11e87158ef61f16.
//
// Solidity: event Price1(uint256 _amountOut)
func (_Apigroup *ApigroupFilterer) WatchPrice1(opts *bind.WatchOpts, sink chan<- *ApigroupPrice1) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "Price1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupPrice1)
				if err := _Apigroup.contract.UnpackLog(event, "Price1", log); err != nil {
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

// ParsePrice1 is a log parse operation binding the contract event 0x0737c274963b06ee8d1f3269a1e91e7d372d7b2d31ddd86fa11e87158ef61f16.
//
// Solidity: event Price1(uint256 _amountOut)
func (_Apigroup *ApigroupFilterer) ParsePrice1(log types.Log) (*ApigroupPrice1, error) {
	event := new(ApigroupPrice1)
	if err := _Apigroup.contract.UnpackLog(event, "Price1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApigroupPrice2Iterator is returned from FilterPrice2 and is used to iterate over the raw logs and unpacked data for Price2 events raised by the Apigroup contract.
type ApigroupPrice2Iterator struct {
	Event *ApigroupPrice2 // Event containing the contract specifics and raw log

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
func (it *ApigroupPrice2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApigroupPrice2)
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
		it.Event = new(ApigroupPrice2)
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
func (it *ApigroupPrice2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApigroupPrice2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApigroupPrice2 represents a Price2 event raised by the Apigroup contract.
type ApigroupPrice2 struct {
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrice2 is a free log retrieval operation binding the contract event 0x87d8dce4cc6b399964f8c8bbcfee5777d877362b3fa1dfa767ecad511a12b917.
//
// Solidity: event Price2(uint256 _amountOut)
func (_Apigroup *ApigroupFilterer) FilterPrice2(opts *bind.FilterOpts) (*ApigroupPrice2Iterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "Price2")
	if err != nil {
		return nil, err
	}
	return &ApigroupPrice2Iterator{contract: _Apigroup.contract, event: "Price2", logs: logs, sub: sub}, nil
}

// WatchPrice2 is a free log subscription operation binding the contract event 0x87d8dce4cc6b399964f8c8bbcfee5777d877362b3fa1dfa767ecad511a12b917.
//
// Solidity: event Price2(uint256 _amountOut)
func (_Apigroup *ApigroupFilterer) WatchPrice2(opts *bind.WatchOpts, sink chan<- *ApigroupPrice2) (event.Subscription, error) {

	logs, sub, err := _Apigroup.contract.WatchLogs(opts, "Price2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApigroupPrice2)
				if err := _Apigroup.contract.UnpackLog(event, "Price2", log); err != nil {
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

// ParsePrice2 is a log parse operation binding the contract event 0x87d8dce4cc6b399964f8c8bbcfee5777d877362b3fa1dfa767ecad511a12b917.
//
// Solidity: event Price2(uint256 _amountOut)
func (_Apigroup *ApigroupFilterer) ParsePrice2(log types.Log) (*ApigroupPrice2, error) {
	event := new(ApigroupPrice2)
	if err := _Apigroup.contract.UnpackLog(event, "Price2", log); err != nil {
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
	LotAddr common.Address
	Owner   common.Address
	Balance *big.Int
	Psnap   *big.Int
	Bsnap   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceiveLotEvent is a free log retrieval operation binding the contract event 0x06ff557b072e8bbcb5472acedec9b50d725df241103eb1d9a79456c79ba3ce89.
//
// Solidity: event ReceiveLotEvent(address _lotAddr, address _owner, uint256 _balance, uint256 _psnap, uint256 _bsnap)
func (_Apigroup *ApigroupFilterer) FilterReceiveLotEvent(opts *bind.FilterOpts) (*ApigroupReceiveLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "ReceiveLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupReceiveLotEventIterator{contract: _Apigroup.contract, event: "ReceiveLotEvent", logs: logs, sub: sub}, nil
}

// WatchReceiveLotEvent is a free log subscription operation binding the contract event 0x06ff557b072e8bbcb5472acedec9b50d725df241103eb1d9a79456c79ba3ce89.
//
// Solidity: event ReceiveLotEvent(address _lotAddr, address _owner, uint256 _balance, uint256 _psnap, uint256 _bsnap)
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

// ParseReceiveLotEvent is a log parse operation binding the contract event 0x06ff557b072e8bbcb5472acedec9b50d725df241103eb1d9a79456c79ba3ce89.
//
// Solidity: event ReceiveLotEvent(address _lotAddr, address _owner, uint256 _balance, uint256 _psnap, uint256 _bsnap)
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
	LotAddr       common.Address
	ReceiveTokens *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendLotEvent is a free log retrieval operation binding the contract event 0x0251ddb2e5f8f355396a07600859eb047c13bd51e3a69c7abb836b504d848fad.
//
// Solidity: event SendLotEvent(address _lotAddr, uint256 _receiveTokens)
func (_Apigroup *ApigroupFilterer) FilterSendLotEvent(opts *bind.FilterOpts) (*ApigroupSendLotEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "SendLotEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupSendLotEventIterator{contract: _Apigroup.contract, event: "SendLotEvent", logs: logs, sub: sub}, nil
}

// WatchSendLotEvent is a free log subscription operation binding the contract event 0x0251ddb2e5f8f355396a07600859eb047c13bd51e3a69c7abb836b504d848fad.
//
// Solidity: event SendLotEvent(address _lotAddr, uint256 _receiveTokens)
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

// ParseSendLotEvent is a log parse operation binding the contract event 0x0251ddb2e5f8f355396a07600859eb047c13bd51e3a69c7abb836b504d848fad.
//
// Solidity: event SendLotEvent(address _lotAddr, uint256 _receiveTokens)
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
	Bsnap *big.Int
	Psnap *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartRoundEvent is a free log retrieval operation binding the contract event 0x116830b501908852222abf8f74b7e51a1f1619bdf216ccd719458025b3307e1e.
//
// Solidity: event StartRoundEvent(uint256 _bsnap, uint256 _psnap)
func (_Apigroup *ApigroupFilterer) FilterStartRoundEvent(opts *bind.FilterOpts) (*ApigroupStartRoundEventIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "StartRoundEvent")
	if err != nil {
		return nil, err
	}
	return &ApigroupStartRoundEventIterator{contract: _Apigroup.contract, event: "StartRoundEvent", logs: logs, sub: sub}, nil
}

// WatchStartRoundEvent is a free log subscription operation binding the contract event 0x116830b501908852222abf8f74b7e51a1f1619bdf216ccd719458025b3307e1e.
//
// Solidity: event StartRoundEvent(uint256 _bsnap, uint256 _psnap)
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

// ParseStartRoundEvent is a log parse operation binding the contract event 0x116830b501908852222abf8f74b7e51a1f1619bdf216ccd719458025b3307e1e.
//
// Solidity: event StartRoundEvent(uint256 _bsnap, uint256 _psnap)
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
	Owner common.Address
	Nwin  *big.Int
	N     *big.Int
	Spos  *big.Int
	Sneg  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlayerParams is a free log retrieval operation binding the contract event 0x66641debe84b19c57d6a387cbfd4810ec4ac03fd84aac0d5605c3b164ef99f7e.
//
// Solidity: event UpdatePlayerParams(address _owner, uint256 _nwin, uint256 _n, uint256 _spos, uint256 _sneg)
func (_Apigroup *ApigroupFilterer) FilterUpdatePlayerParams(opts *bind.FilterOpts) (*ApigroupUpdatePlayerParamsIterator, error) {

	logs, sub, err := _Apigroup.contract.FilterLogs(opts, "UpdatePlayerParams")
	if err != nil {
		return nil, err
	}
	return &ApigroupUpdatePlayerParamsIterator{contract: _Apigroup.contract, event: "UpdatePlayerParams", logs: logs, sub: sub}, nil
}

// WatchUpdatePlayerParams is a free log subscription operation binding the contract event 0x66641debe84b19c57d6a387cbfd4810ec4ac03fd84aac0d5605c3b164ef99f7e.
//
// Solidity: event UpdatePlayerParams(address _owner, uint256 _nwin, uint256 _n, uint256 _spos, uint256 _sneg)
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

// ParseUpdatePlayerParams is a log parse operation binding the contract event 0x66641debe84b19c57d6a387cbfd4810ec4ac03fd84aac0d5605c3b164ef99f7e.
//
// Solidity: event UpdatePlayerParams(address _owner, uint256 _nwin, uint256 _n, uint256 _spos, uint256 _sneg)
func (_Apigroup *ApigroupFilterer) ParseUpdatePlayerParams(log types.Log) (*ApigroupUpdatePlayerParams, error) {
	event := new(ApigroupUpdatePlayerParams)
	if err := _Apigroup.contract.UnpackLog(event, "UpdatePlayerParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
