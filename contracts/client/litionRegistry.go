// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package litionScClient

import (
	"math/big"
	"strings"

	ethereum "gitlab.com/lition/lition"
	"gitlab.com/lition/lition/accounts/abi"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/common"
	"gitlab.com/lition/lition/core/types"
	"gitlab.com/lition/lition/event"
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

// LitionScClientABI is the input ABI used to generate the binding from.
const LitionScClientABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLastNotary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notaryBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notaryTimestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"confirmVestInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"getAllowedToValidate\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"stopMining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"cancelVestInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"startMining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"getAllowedToTransact\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notaryStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notaryEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocksMined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"userGas\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"largestTx\",\"type\":\"uint32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getUserRequests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"vestingReqExists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqNotary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqControlState\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"depositReqExists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"depositReqTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositReqNotary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositReqValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositReqState\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getChainDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalVesting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastNotaryBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastNotaryTimestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"confirmDepositWithdrawalFromChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notaryBlockNo\",\"type\":\"uint256\"}],\"name\":\"testNotary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"hasDeposited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"requestDepositInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"cancelDepositInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getUserDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"contractChainValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"vesting\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"deposit\",\"type\":\"uint96\"},{\"internalType\":\"string\",\"name\":\"initEndpoint\",\"type\":\"string\"}],\"name\":\"registerChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"hasVested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"}],\"name\":\"requestVestInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"RequestDepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"ConfirmDepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"CancelDepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ForceWithdrawDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"RequestVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"ConfirmVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"CancelVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reqTimestamp\",\"type\":\"uint256\"}],\"name\":\"AcceptedVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ForceWithdrawVesting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"WhitelistAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"StartMining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"StopMining\",\"type\":\"event\"}]"

// LitionScClient is an auto generated Go binding around an Ethereum contract.
type LitionScClient struct {
	LitionScClientCaller     // Read-only binding to the contract
	LitionScClientTransactor // Write-only binding to the contract
	LitionScClientFilterer   // Log filterer for contract events
}

// LitionScClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type LitionScClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionScClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LitionScClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionScClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LitionScClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionScClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LitionScClientSession struct {
	Contract     *LitionScClient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LitionScClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LitionScClientCallerSession struct {
	Contract *LitionScClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LitionScClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LitionScClientTransactorSession struct {
	Contract     *LitionScClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LitionScClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type LitionScClientRaw struct {
	Contract *LitionScClient // Generic contract binding to access the raw methods on
}

// LitionScClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LitionScClientCallerRaw struct {
	Contract *LitionScClientCaller // Generic read-only contract binding to access the raw methods on
}

// LitionScClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LitionScClientTransactorRaw struct {
	Contract *LitionScClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLitionScClient creates a new instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClient(address common.Address, backend bind.ContractBackend) (*LitionScClient, error) {
	contract, err := bindLitionScClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LitionScClient{LitionScClientCaller: LitionScClientCaller{contract: contract}, LitionScClientTransactor: LitionScClientTransactor{contract: contract}, LitionScClientFilterer: LitionScClientFilterer{contract: contract}}, nil
}

// NewLitionScClientCaller creates a new read-only instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClientCaller(address common.Address, caller bind.ContractCaller) (*LitionScClientCaller, error) {
	contract, err := bindLitionScClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LitionScClientCaller{contract: contract}, nil
}

// NewLitionScClientTransactor creates a new write-only instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClientTransactor(address common.Address, transactor bind.ContractTransactor) (*LitionScClientTransactor, error) {
	contract, err := bindLitionScClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LitionScClientTransactor{contract: contract}, nil
}

// NewLitionScClientFilterer creates a new log filterer instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClientFilterer(address common.Address, filterer bind.ContractFilterer) (*LitionScClientFilterer, error) {
	contract, err := bindLitionScClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LitionScClientFilterer{contract: contract}, nil
}

// bindLitionScClient binds a generic wrapper to an already deployed contract.
func bindLitionScClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LitionScClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LitionScClient *LitionScClientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LitionScClient.Contract.LitionScClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LitionScClient *LitionScClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LitionScClient.Contract.LitionScClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LitionScClient *LitionScClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LitionScClient.Contract.LitionScClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LitionScClient *LitionScClientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LitionScClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LitionScClient *LitionScClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LitionScClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LitionScClient *LitionScClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LitionScClient.Contract.contract.Transact(opts, method, params...)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x7e233e26.
//
// Solidity: function getActiveValidators(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCaller) GetActiveValidators(opts *bind.CallOpts, chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	var (
		ret0 = new([100]common.Address)
		ret1 = new(*big.Int)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _LitionScClient.contract.Call(opts, out, "getActiveValidators", chainId, batch)
	return *ret0, *ret1, *ret2, err
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x7e233e26.
//
// Solidity: function getActiveValidators(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientSession) GetActiveValidators(chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetActiveValidators(&_LitionScClient.CallOpts, chainId, batch)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x7e233e26.
//
// Solidity: function getActiveValidators(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCallerSession) GetActiveValidators(chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetActiveValidators(&_LitionScClient.CallOpts, chainId, batch)
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0x529f31bf.
//
// Solidity: function getAllowedToTransact(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCaller) GetAllowedToTransact(opts *bind.CallOpts, chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	var (
		ret0 = new([100]common.Address)
		ret1 = new(*big.Int)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _LitionScClient.contract.Call(opts, out, "getAllowedToTransact", chainId, batch)
	return *ret0, *ret1, *ret2, err
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0x529f31bf.
//
// Solidity: function getAllowedToTransact(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientSession) GetAllowedToTransact(chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToTransact(&_LitionScClient.CallOpts, chainId, batch)
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0x529f31bf.
//
// Solidity: function getAllowedToTransact(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCallerSession) GetAllowedToTransact(chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToTransact(&_LitionScClient.CallOpts, chainId, batch)
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0x148dbe94.
//
// Solidity: function getAllowedToValidate(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCaller) GetAllowedToValidate(opts *bind.CallOpts, chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	var (
		ret0 = new([100]common.Address)
		ret1 = new(*big.Int)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _LitionScClient.contract.Call(opts, out, "getAllowedToValidate", chainId, batch)
	return *ret0, *ret1, *ret2, err
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0x148dbe94.
//
// Solidity: function getAllowedToValidate(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientSession) GetAllowedToValidate(chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToValidate(&_LitionScClient.CallOpts, chainId, batch)
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0x148dbe94.
//
// Solidity: function getAllowedToValidate(uint256 chainId, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCallerSession) GetAllowedToValidate(chainId *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToValidate(&_LitionScClient.CallOpts, chainId, batch)
}

// GetChainDetails is a free data retrieval call binding the contract method 0x79c767d6.
//
// Solidity: function getChainDetails(uint256 chainId) constant returns(bool registered, bool active, string endpoint, uint256 totalVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp)
func (_LitionScClient *LitionScClientCaller) GetChainDetails(opts *bind.CallOpts, chainId *big.Int) (struct {
	Registered          bool
	Active              bool
	Endpoint            string
	TotalVesting        *big.Int
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	ret := new(struct {
		Registered          bool
		Active              bool
		Endpoint            string
		TotalVesting        *big.Int
		LastNotaryBlock     *big.Int
		LastNotaryTimestamp *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getChainDetails", chainId)
	return *ret, err
}

// GetChainDetails is a free data retrieval call binding the contract method 0x79c767d6.
//
// Solidity: function getChainDetails(uint256 chainId) constant returns(bool registered, bool active, string endpoint, uint256 totalVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp)
func (_LitionScClient *LitionScClientSession) GetChainDetails(chainId *big.Int) (struct {
	Registered          bool
	Active              bool
	Endpoint            string
	TotalVesting        *big.Int
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetChainDetails(&_LitionScClient.CallOpts, chainId)
}

// GetChainDetails is a free data retrieval call binding the contract method 0x79c767d6.
//
// Solidity: function getChainDetails(uint256 chainId) constant returns(bool registered, bool active, string endpoint, uint256 totalVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp)
func (_LitionScClient *LitionScClientCallerSession) GetChainDetails(chainId *big.Int) (struct {
	Registered          bool
	Active              bool
	Endpoint            string
	TotalVesting        *big.Int
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetChainDetails(&_LitionScClient.CallOpts, chainId)
}

// GetLastNotary is a free data retrieval call binding the contract method 0x0ebd0ff7.
//
// Solidity: function getLastNotary(uint256 chainId) constant returns(uint256 notaryBlock, uint256 notaryTimestamp)
func (_LitionScClient *LitionScClientCaller) GetLastNotary(opts *bind.CallOpts, chainId *big.Int) (struct {
	NotaryBlock     *big.Int
	NotaryTimestamp *big.Int
}, error) {
	ret := new(struct {
		NotaryBlock     *big.Int
		NotaryTimestamp *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getLastNotary", chainId)
	return *ret, err
}

// GetLastNotary is a free data retrieval call binding the contract method 0x0ebd0ff7.
//
// Solidity: function getLastNotary(uint256 chainId) constant returns(uint256 notaryBlock, uint256 notaryTimestamp)
func (_LitionScClient *LitionScClientSession) GetLastNotary(chainId *big.Int) (struct {
	NotaryBlock     *big.Int
	NotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetLastNotary(&_LitionScClient.CallOpts, chainId)
}

// GetLastNotary is a free data retrieval call binding the contract method 0x0ebd0ff7.
//
// Solidity: function getLastNotary(uint256 chainId) constant returns(uint256 notaryBlock, uint256 notaryTimestamp)
func (_LitionScClient *LitionScClientCallerSession) GetLastNotary(chainId *big.Int) (struct {
	NotaryBlock     *big.Int
	NotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetLastNotary(&_LitionScClient.CallOpts, chainId)
}

// GetUserDetails is a free data retrieval call binding the contract method 0xc90902cb.
//
// Solidity: function getUserDetails(uint256 chainId, address acc) constant returns(bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining)
func (_LitionScClient *LitionScClientCaller) GetUserDetails(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (struct {
	Exists      bool
	Deposit     *big.Int
	Whitelisted bool
	Vesting     *big.Int
	Mining      bool
}, error) {
	ret := new(struct {
		Exists      bool
		Deposit     *big.Int
		Whitelisted bool
		Vesting     *big.Int
		Mining      bool
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getUserDetails", chainId, acc)
	return *ret, err
}

// GetUserDetails is a free data retrieval call binding the contract method 0xc90902cb.
//
// Solidity: function getUserDetails(uint256 chainId, address acc) constant returns(bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining)
func (_LitionScClient *LitionScClientSession) GetUserDetails(chainId *big.Int, acc common.Address) (struct {
	Exists      bool
	Deposit     *big.Int
	Whitelisted bool
	Vesting     *big.Int
	Mining      bool
}, error) {
	return _LitionScClient.Contract.GetUserDetails(&_LitionScClient.CallOpts, chainId, acc)
}

// GetUserDetails is a free data retrieval call binding the contract method 0xc90902cb.
//
// Solidity: function getUserDetails(uint256 chainId, address acc) constant returns(bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining)
func (_LitionScClient *LitionScClientCallerSession) GetUserDetails(chainId *big.Int, acc common.Address) (struct {
	Exists      bool
	Deposit     *big.Int
	Whitelisted bool
	Vesting     *big.Int
	Mining      bool
}, error) {
	return _LitionScClient.Contract.GetUserDetails(&_LitionScClient.CallOpts, chainId, acc)
}

// GetUserRequests is a free data retrieval call binding the contract method 0x75195b67.
//
// Solidity: function getUserRequests(uint256 chainId, address acc) constant returns(bool vestingReqExists, uint256 vestingReqTime, uint256 vestingReqNotary, uint256 vestingReqValue, uint256 vestingReqState, uint256 vestingReqControlState, bool depositReqExists, uint256 depositReqTime, uint256 depositReqNotary, uint256 depositReqValue, uint256 depositReqState)
func (_LitionScClient *LitionScClientCaller) GetUserRequests(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (struct {
	VestingReqExists       bool
	VestingReqTime         *big.Int
	VestingReqNotary       *big.Int
	VestingReqValue        *big.Int
	VestingReqState        *big.Int
	VestingReqControlState *big.Int
	DepositReqExists       bool
	DepositReqTime         *big.Int
	DepositReqNotary       *big.Int
	DepositReqValue        *big.Int
	DepositReqState        *big.Int
}, error) {
	ret := new(struct {
		VestingReqExists       bool
		VestingReqTime         *big.Int
		VestingReqNotary       *big.Int
		VestingReqValue        *big.Int
		VestingReqState        *big.Int
		VestingReqControlState *big.Int
		DepositReqExists       bool
		DepositReqTime         *big.Int
		DepositReqNotary       *big.Int
		DepositReqValue        *big.Int
		DepositReqState        *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getUserRequests", chainId, acc)
	return *ret, err
}

// GetUserRequests is a free data retrieval call binding the contract method 0x75195b67.
//
// Solidity: function getUserRequests(uint256 chainId, address acc) constant returns(bool vestingReqExists, uint256 vestingReqTime, uint256 vestingReqNotary, uint256 vestingReqValue, uint256 vestingReqState, uint256 vestingReqControlState, bool depositReqExists, uint256 depositReqTime, uint256 depositReqNotary, uint256 depositReqValue, uint256 depositReqState)
func (_LitionScClient *LitionScClientSession) GetUserRequests(chainId *big.Int, acc common.Address) (struct {
	VestingReqExists       bool
	VestingReqTime         *big.Int
	VestingReqNotary       *big.Int
	VestingReqValue        *big.Int
	VestingReqState        *big.Int
	VestingReqControlState *big.Int
	DepositReqExists       bool
	DepositReqTime         *big.Int
	DepositReqNotary       *big.Int
	DepositReqValue        *big.Int
	DepositReqState        *big.Int
}, error) {
	return _LitionScClient.Contract.GetUserRequests(&_LitionScClient.CallOpts, chainId, acc)
}

// GetUserRequests is a free data retrieval call binding the contract method 0x75195b67.
//
// Solidity: function getUserRequests(uint256 chainId, address acc) constant returns(bool vestingReqExists, uint256 vestingReqTime, uint256 vestingReqNotary, uint256 vestingReqValue, uint256 vestingReqState, uint256 vestingReqControlState, bool depositReqExists, uint256 depositReqTime, uint256 depositReqNotary, uint256 depositReqValue, uint256 depositReqState)
func (_LitionScClient *LitionScClientCallerSession) GetUserRequests(chainId *big.Int, acc common.Address) (struct {
	VestingReqExists       bool
	VestingReqTime         *big.Int
	VestingReqNotary       *big.Int
	VestingReqValue        *big.Int
	VestingReqState        *big.Int
	VestingReqControlState *big.Int
	DepositReqExists       bool
	DepositReqTime         *big.Int
	DepositReqNotary       *big.Int
	DepositReqValue        *big.Int
	DepositReqState        *big.Int
}, error) {
	return _LitionScClient.Contract.GetUserRequests(&_LitionScClient.CallOpts, chainId, acc)
}

// HasDeposited is a free data retrieval call binding the contract method 0x99b1eb2e.
//
// Solidity: function hasDeposited(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) HasDeposited(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "hasDeposited", chainId, acc)
	return *ret0, err
}

// HasDeposited is a free data retrieval call binding the contract method 0x99b1eb2e.
//
// Solidity: function hasDeposited(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) HasDeposited(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasDeposited(&_LitionScClient.CallOpts, chainId, acc)
}

// HasDeposited is a free data retrieval call binding the contract method 0x99b1eb2e.
//
// Solidity: function hasDeposited(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) HasDeposited(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasDeposited(&_LitionScClient.CallOpts, chainId, acc)
}

// HasVested is a free data retrieval call binding the contract method 0xda40da76.
//
// Solidity: function hasVested(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) HasVested(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "hasVested", chainId, acc)
	return *ret0, err
}

// HasVested is a free data retrieval call binding the contract method 0xda40da76.
//
// Solidity: function hasVested(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) HasVested(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasVested(&_LitionScClient.CallOpts, chainId, acc)
}

// HasVested is a free data retrieval call binding the contract method 0xda40da76.
//
// Solidity: function hasVested(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) HasVested(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasVested(&_LitionScClient.CallOpts, chainId, acc)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() constant returns(uint256)
func (_LitionScClient *LitionScClientCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "nextId")
	return *ret0, err
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() constant returns(uint256)
func (_LitionScClient *LitionScClientSession) NextId() (*big.Int, error) {
	return _LitionScClient.Contract.NextId(&_LitionScClient.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() constant returns(uint256)
func (_LitionScClient *LitionScClientCallerSession) NextId() (*big.Int, error) {
	return _LitionScClient.Contract.NextId(&_LitionScClient.CallOpts)
}

// CancelDepositInChain is a paid mutator transaction binding the contract method 0xa1a906e0.
//
// Solidity: function cancelDepositInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) CancelDepositInChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "cancelDepositInChain", chainId)
}

// CancelDepositInChain is a paid mutator transaction binding the contract method 0xa1a906e0.
//
// Solidity: function cancelDepositInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) CancelDepositInChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelDepositInChain(&_LitionScClient.TransactOpts, chainId)
}

// CancelDepositInChain is a paid mutator transaction binding the contract method 0xa1a906e0.
//
// Solidity: function cancelDepositInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) CancelDepositInChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelDepositInChain(&_LitionScClient.TransactOpts, chainId)
}

// CancelVestInChain is a paid mutator transaction binding the contract method 0x449c067f.
//
// Solidity: function cancelVestInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) CancelVestInChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "cancelVestInChain", chainId)
}

// CancelVestInChain is a paid mutator transaction binding the contract method 0x449c067f.
//
// Solidity: function cancelVestInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) CancelVestInChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelVestInChain(&_LitionScClient.TransactOpts, chainId)
}

// CancelVestInChain is a paid mutator transaction binding the contract method 0x449c067f.
//
// Solidity: function cancelVestInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) CancelVestInChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelVestInChain(&_LitionScClient.TransactOpts, chainId)
}

// ConfirmDepositWithdrawalFromChain is a paid mutator transaction binding the contract method 0x81577959.
//
// Solidity: function confirmDepositWithdrawalFromChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) ConfirmDepositWithdrawalFromChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "confirmDepositWithdrawalFromChain", chainId)
}

// ConfirmDepositWithdrawalFromChain is a paid mutator transaction binding the contract method 0x81577959.
//
// Solidity: function confirmDepositWithdrawalFromChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) ConfirmDepositWithdrawalFromChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmDepositWithdrawalFromChain(&_LitionScClient.TransactOpts, chainId)
}

// ConfirmDepositWithdrawalFromChain is a paid mutator transaction binding the contract method 0x81577959.
//
// Solidity: function confirmDepositWithdrawalFromChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) ConfirmDepositWithdrawalFromChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmDepositWithdrawalFromChain(&_LitionScClient.TransactOpts, chainId)
}

// ConfirmVestInChain is a paid mutator transaction binding the contract method 0x13456851.
//
// Solidity: function confirmVestInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) ConfirmVestInChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "confirmVestInChain", chainId)
}

// ConfirmVestInChain is a paid mutator transaction binding the contract method 0x13456851.
//
// Solidity: function confirmVestInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) ConfirmVestInChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmVestInChain(&_LitionScClient.TransactOpts, chainId)
}

// ConfirmVestInChain is a paid mutator transaction binding the contract method 0x13456851.
//
// Solidity: function confirmVestInChain(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) ConfirmVestInChain(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmVestInChain(&_LitionScClient.TransactOpts, chainId)
}

// Notary is a paid mutator transaction binding the contract method 0x6cd22ed8.
//
// Solidity: function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] miners, uint32[] blocksMined, address[] users, uint32[] userGas, uint32 largestTx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactor) Notary(opts *bind.TransactOpts, chainId *big.Int, notaryStartBlock *big.Int, notaryEndBlock *big.Int, miners []common.Address, blocksMined []uint32, users []common.Address, userGas []uint32, largestTx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "notary", chainId, notaryStartBlock, notaryEndBlock, miners, blocksMined, users, userGas, largestTx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0x6cd22ed8.
//
// Solidity: function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] miners, uint32[] blocksMined, address[] users, uint32[] userGas, uint32 largestTx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientSession) Notary(chainId *big.Int, notaryStartBlock *big.Int, notaryEndBlock *big.Int, miners []common.Address, blocksMined []uint32, users []common.Address, userGas []uint32, largestTx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, chainId, notaryStartBlock, notaryEndBlock, miners, blocksMined, users, userGas, largestTx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0x6cd22ed8.
//
// Solidity: function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] miners, uint32[] blocksMined, address[] users, uint32[] userGas, uint32 largestTx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactorSession) Notary(chainId *big.Int, notaryStartBlock *big.Int, notaryEndBlock *big.Int, miners []common.Address, blocksMined []uint32, users []common.Address, userGas []uint32, largestTx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, chainId, notaryStartBlock, notaryEndBlock, miners, blocksMined, users, userGas, largestTx, v, r, s)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xcf883380.
//
// Solidity: function registerChain(string description, address validator, uint96 vesting, uint96 deposit, string initEndpoint) returns(uint256 chainId)
func (_LitionScClient *LitionScClientTransactor) RegisterChain(opts *bind.TransactOpts, description string, validator common.Address, vesting *big.Int, deposit *big.Int, initEndpoint string) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "registerChain", description, validator, vesting, deposit, initEndpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xcf883380.
//
// Solidity: function registerChain(string description, address validator, uint96 vesting, uint96 deposit, string initEndpoint) returns(uint256 chainId)
func (_LitionScClient *LitionScClientSession) RegisterChain(description string, validator common.Address, vesting *big.Int, deposit *big.Int, initEndpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, description, validator, vesting, deposit, initEndpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xcf883380.
//
// Solidity: function registerChain(string description, address validator, uint96 vesting, uint96 deposit, string initEndpoint) returns(uint256 chainId)
func (_LitionScClient *LitionScClientTransactorSession) RegisterChain(description string, validator common.Address, vesting *big.Int, deposit *big.Int, initEndpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, description, validator, vesting, deposit, initEndpoint)
}

// RequestDepositInChain is a paid mutator transaction binding the contract method 0x9e9a4db9.
//
// Solidity: function requestDepositInChain(uint256 chainId, uint256 deposit) returns()
func (_LitionScClient *LitionScClientTransactor) RequestDepositInChain(opts *bind.TransactOpts, chainId *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "requestDepositInChain", chainId, deposit)
}

// RequestDepositInChain is a paid mutator transaction binding the contract method 0x9e9a4db9.
//
// Solidity: function requestDepositInChain(uint256 chainId, uint256 deposit) returns()
func (_LitionScClient *LitionScClientSession) RequestDepositInChain(chainId *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestDepositInChain(&_LitionScClient.TransactOpts, chainId, deposit)
}

// RequestDepositInChain is a paid mutator transaction binding the contract method 0x9e9a4db9.
//
// Solidity: function requestDepositInChain(uint256 chainId, uint256 deposit) returns()
func (_LitionScClient *LitionScClientTransactorSession) RequestDepositInChain(chainId *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestDepositInChain(&_LitionScClient.TransactOpts, chainId, deposit)
}

// RequestVestInChain is a paid mutator transaction binding the contract method 0xf2902fb0.
//
// Solidity: function requestVestInChain(uint256 chainId, uint256 vesting) returns()
func (_LitionScClient *LitionScClientTransactor) RequestVestInChain(opts *bind.TransactOpts, chainId *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "requestVestInChain", chainId, vesting)
}

// RequestVestInChain is a paid mutator transaction binding the contract method 0xf2902fb0.
//
// Solidity: function requestVestInChain(uint256 chainId, uint256 vesting) returns()
func (_LitionScClient *LitionScClientSession) RequestVestInChain(chainId *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestVestInChain(&_LitionScClient.TransactOpts, chainId, vesting)
}

// RequestVestInChain is a paid mutator transaction binding the contract method 0xf2902fb0.
//
// Solidity: function requestVestInChain(uint256 chainId, uint256 vesting) returns()
func (_LitionScClient *LitionScClientTransactorSession) RequestVestInChain(chainId *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestVestInChain(&_LitionScClient.TransactOpts, chainId, vesting)
}

// StartMining is a paid mutator transaction binding the contract method 0x47b272c0.
//
// Solidity: function startMining(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) StartMining(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "startMining", chainId)
}

// StartMining is a paid mutator transaction binding the contract method 0x47b272c0.
//
// Solidity: function startMining(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) StartMining(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StartMining(&_LitionScClient.TransactOpts, chainId)
}

// StartMining is a paid mutator transaction binding the contract method 0x47b272c0.
//
// Solidity: function startMining(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) StartMining(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StartMining(&_LitionScClient.TransactOpts, chainId)
}

// StopMining is a paid mutator transaction binding the contract method 0x3b714199.
//
// Solidity: function stopMining(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) StopMining(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "stopMining", chainId)
}

// StopMining is a paid mutator transaction binding the contract method 0x3b714199.
//
// Solidity: function stopMining(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) StopMining(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StopMining(&_LitionScClient.TransactOpts, chainId)
}

// StopMining is a paid mutator transaction binding the contract method 0x3b714199.
//
// Solidity: function stopMining(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) StopMining(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StopMining(&_LitionScClient.TransactOpts, chainId)
}

// TestNotary is a paid mutator transaction binding the contract method 0x83f2d9d2.
//
// Solidity: function testNotary(uint256 chainId, uint256 notaryBlockNo) returns()
func (_LitionScClient *LitionScClientTransactor) TestNotary(opts *bind.TransactOpts, chainId *big.Int, notaryBlockNo *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "testNotary", chainId, notaryBlockNo)
}

// TestNotary is a paid mutator transaction binding the contract method 0x83f2d9d2.
//
// Solidity: function testNotary(uint256 chainId, uint256 notaryBlockNo) returns()
func (_LitionScClient *LitionScClientSession) TestNotary(chainId *big.Int, notaryBlockNo *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.TestNotary(&_LitionScClient.TransactOpts, chainId, notaryBlockNo)
}

// TestNotary is a paid mutator transaction binding the contract method 0x83f2d9d2.
//
// Solidity: function testNotary(uint256 chainId, uint256 notaryBlockNo) returns()
func (_LitionScClient *LitionScClientTransactorSession) TestNotary(chainId *big.Int, notaryBlockNo *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.TestNotary(&_LitionScClient.TransactOpts, chainId, notaryBlockNo)
}

// LitionScClientAcceptedVestInChainIterator is returned from FilterAcceptedVestInChain and is used to iterate over the raw logs and unpacked data for AcceptedVestInChain events raised by the LitionScClient contract.
type LitionScClientAcceptedVestInChainIterator struct {
	Event *LitionScClientAcceptedVestInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientAcceptedVestInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientAcceptedVestInChain)
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
		it.Event = new(LitionScClientAcceptedVestInChain)
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
func (it *LitionScClientAcceptedVestInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientAcceptedVestInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientAcceptedVestInChain represents a AcceptedVestInChain event raised by the LitionScClient contract.
type LitionScClientAcceptedVestInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Vesting      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAcceptedVestInChain is a free log retrieval operation binding the contract event 0x00547b633d4da14c6f2e748a506f553393a70da4e5cdd6cb304865b140f92f6c.
//
// Solidity: event AcceptedVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterAcceptedVestInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientAcceptedVestInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "AcceptedVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientAcceptedVestInChainIterator{contract: _LitionScClient.contract, event: "AcceptedVestInChain", logs: logs, sub: sub}, nil
}

// WatchAcceptedVestInChain is a free log subscription operation binding the contract event 0x00547b633d4da14c6f2e748a506f553393a70da4e5cdd6cb304865b140f92f6c.
//
// Solidity: event AcceptedVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchAcceptedVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientAcceptedVestInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "AcceptedVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientAcceptedVestInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "AcceptedVestInChain", log); err != nil {
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

// ParseAcceptedVestInChain is a log parse operation binding the contract event 0x00547b633d4da14c6f2e748a506f553393a70da4e5cdd6cb304865b140f92f6c.
//
// Solidity: event AcceptedVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseAcceptedVestInChain(log types.Log) (*LitionScClientAcceptedVestInChain, error) {
	event := new(LitionScClientAcceptedVestInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "AcceptedVestInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientCancelDepositInChainIterator is returned from FilterCancelDepositInChain and is used to iterate over the raw logs and unpacked data for CancelDepositInChain events raised by the LitionScClient contract.
type LitionScClientCancelDepositInChainIterator struct {
	Event *LitionScClientCancelDepositInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientCancelDepositInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientCancelDepositInChain)
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
		it.Event = new(LitionScClientCancelDepositInChain)
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
func (it *LitionScClientCancelDepositInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientCancelDepositInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientCancelDepositInChain represents a CancelDepositInChain event raised by the LitionScClient contract.
type LitionScClientCancelDepositInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Deposit      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCancelDepositInChain is a free log retrieval operation binding the contract event 0x9e6b1e30d746be09c6f5cf9e8ec4c7bf584bfe1308d70b39a097363e118814e5.
//
// Solidity: event CancelDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterCancelDepositInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientCancelDepositInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "CancelDepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientCancelDepositInChainIterator{contract: _LitionScClient.contract, event: "CancelDepositInChain", logs: logs, sub: sub}, nil
}

// WatchCancelDepositInChain is a free log subscription operation binding the contract event 0x9e6b1e30d746be09c6f5cf9e8ec4c7bf584bfe1308d70b39a097363e118814e5.
//
// Solidity: event CancelDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchCancelDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientCancelDepositInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "CancelDepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientCancelDepositInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "CancelDepositInChain", log); err != nil {
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

// ParseCancelDepositInChain is a log parse operation binding the contract event 0x9e6b1e30d746be09c6f5cf9e8ec4c7bf584bfe1308d70b39a097363e118814e5.
//
// Solidity: event CancelDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseCancelDepositInChain(log types.Log) (*LitionScClientCancelDepositInChain, error) {
	event := new(LitionScClientCancelDepositInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "CancelDepositInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientCancelVestInChainIterator is returned from FilterCancelVestInChain and is used to iterate over the raw logs and unpacked data for CancelVestInChain events raised by the LitionScClient contract.
type LitionScClientCancelVestInChainIterator struct {
	Event *LitionScClientCancelVestInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientCancelVestInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientCancelVestInChain)
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
		it.Event = new(LitionScClientCancelVestInChain)
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
func (it *LitionScClientCancelVestInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientCancelVestInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientCancelVestInChain represents a CancelVestInChain event raised by the LitionScClient contract.
type LitionScClientCancelVestInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Vesting      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCancelVestInChain is a free log retrieval operation binding the contract event 0x698f91699fc921ec104a854863c577f1d32e8d48e41eac966924cbbdae212fea.
//
// Solidity: event CancelVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterCancelVestInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientCancelVestInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "CancelVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientCancelVestInChainIterator{contract: _LitionScClient.contract, event: "CancelVestInChain", logs: logs, sub: sub}, nil
}

// WatchCancelVestInChain is a free log subscription operation binding the contract event 0x698f91699fc921ec104a854863c577f1d32e8d48e41eac966924cbbdae212fea.
//
// Solidity: event CancelVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchCancelVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientCancelVestInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "CancelVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientCancelVestInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "CancelVestInChain", log); err != nil {
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

// ParseCancelVestInChain is a log parse operation binding the contract event 0x698f91699fc921ec104a854863c577f1d32e8d48e41eac966924cbbdae212fea.
//
// Solidity: event CancelVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseCancelVestInChain(log types.Log) (*LitionScClientCancelVestInChain, error) {
	event := new(LitionScClientCancelVestInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "CancelVestInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientConfirmDepositInChainIterator is returned from FilterConfirmDepositInChain and is used to iterate over the raw logs and unpacked data for ConfirmDepositInChain events raised by the LitionScClient contract.
type LitionScClientConfirmDepositInChainIterator struct {
	Event *LitionScClientConfirmDepositInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientConfirmDepositInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientConfirmDepositInChain)
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
		it.Event = new(LitionScClientConfirmDepositInChain)
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
func (it *LitionScClientConfirmDepositInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientConfirmDepositInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientConfirmDepositInChain represents a ConfirmDepositInChain event raised by the LitionScClient contract.
type LitionScClientConfirmDepositInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Deposit      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterConfirmDepositInChain is a free log retrieval operation binding the contract event 0x339a7256b6e5032769b6cee82c769df79aefc6215244052d987decc224251499.
//
// Solidity: event ConfirmDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterConfirmDepositInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientConfirmDepositInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ConfirmDepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientConfirmDepositInChainIterator{contract: _LitionScClient.contract, event: "ConfirmDepositInChain", logs: logs, sub: sub}, nil
}

// WatchConfirmDepositInChain is a free log subscription operation binding the contract event 0x339a7256b6e5032769b6cee82c769df79aefc6215244052d987decc224251499.
//
// Solidity: event ConfirmDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchConfirmDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientConfirmDepositInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ConfirmDepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientConfirmDepositInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "ConfirmDepositInChain", log); err != nil {
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

// ParseConfirmDepositInChain is a log parse operation binding the contract event 0x339a7256b6e5032769b6cee82c769df79aefc6215244052d987decc224251499.
//
// Solidity: event ConfirmDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseConfirmDepositInChain(log types.Log) (*LitionScClientConfirmDepositInChain, error) {
	event := new(LitionScClientConfirmDepositInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "ConfirmDepositInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientConfirmVestInChainIterator is returned from FilterConfirmVestInChain and is used to iterate over the raw logs and unpacked data for ConfirmVestInChain events raised by the LitionScClient contract.
type LitionScClientConfirmVestInChainIterator struct {
	Event *LitionScClientConfirmVestInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientConfirmVestInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientConfirmVestInChain)
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
		it.Event = new(LitionScClientConfirmVestInChain)
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
func (it *LitionScClientConfirmVestInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientConfirmVestInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientConfirmVestInChain represents a ConfirmVestInChain event raised by the LitionScClient contract.
type LitionScClientConfirmVestInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Vesting      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterConfirmVestInChain is a free log retrieval operation binding the contract event 0xb5aaacabfddd3428a7e8b351250df8f590b10b3eb0709a08223d2730aa110732.
//
// Solidity: event ConfirmVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterConfirmVestInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientConfirmVestInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ConfirmVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientConfirmVestInChainIterator{contract: _LitionScClient.contract, event: "ConfirmVestInChain", logs: logs, sub: sub}, nil
}

// WatchConfirmVestInChain is a free log subscription operation binding the contract event 0xb5aaacabfddd3428a7e8b351250df8f590b10b3eb0709a08223d2730aa110732.
//
// Solidity: event ConfirmVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchConfirmVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientConfirmVestInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ConfirmVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientConfirmVestInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "ConfirmVestInChain", log); err != nil {
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

// ParseConfirmVestInChain is a log parse operation binding the contract event 0xb5aaacabfddd3428a7e8b351250df8f590b10b3eb0709a08223d2730aa110732.
//
// Solidity: event ConfirmVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseConfirmVestInChain(log types.Log) (*LitionScClientConfirmVestInChain, error) {
	event := new(LitionScClientConfirmVestInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "ConfirmVestInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientForceWithdrawDepositIterator is returned from FilterForceWithdrawDeposit and is used to iterate over the raw logs and unpacked data for ForceWithdrawDeposit events raised by the LitionScClient contract.
type LitionScClientForceWithdrawDepositIterator struct {
	Event *LitionScClientForceWithdrawDeposit // Event containing the contract specifics and raw log

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
func (it *LitionScClientForceWithdrawDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientForceWithdrawDeposit)
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
		it.Event = new(LitionScClientForceWithdrawDeposit)
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
func (it *LitionScClientForceWithdrawDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientForceWithdrawDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientForceWithdrawDeposit represents a ForceWithdrawDeposit event raised by the LitionScClient contract.
type LitionScClientForceWithdrawDeposit struct {
	ChainId   *big.Int
	Account   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForceWithdrawDeposit is a free log retrieval operation binding the contract event 0x09ddb46db0b11a56fa7611442b214b389d5479513714d4f5c596cdd62b39a59c.
//
// Solidity: event ForceWithdrawDeposit(uint256 indexed chainId, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterForceWithdrawDeposit(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientForceWithdrawDepositIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ForceWithdrawDeposit", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientForceWithdrawDepositIterator{contract: _LitionScClient.contract, event: "ForceWithdrawDeposit", logs: logs, sub: sub}, nil
}

// WatchForceWithdrawDeposit is a free log subscription operation binding the contract event 0x09ddb46db0b11a56fa7611442b214b389d5479513714d4f5c596cdd62b39a59c.
//
// Solidity: event ForceWithdrawDeposit(uint256 indexed chainId, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchForceWithdrawDeposit(opts *bind.WatchOpts, sink chan<- *LitionScClientForceWithdrawDeposit, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ForceWithdrawDeposit", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientForceWithdrawDeposit)
				if err := _LitionScClient.contract.UnpackLog(event, "ForceWithdrawDeposit", log); err != nil {
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

// ParseForceWithdrawDeposit is a log parse operation binding the contract event 0x09ddb46db0b11a56fa7611442b214b389d5479513714d4f5c596cdd62b39a59c.
//
// Solidity: event ForceWithdrawDeposit(uint256 indexed chainId, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) ParseForceWithdrawDeposit(log types.Log) (*LitionScClientForceWithdrawDeposit, error) {
	event := new(LitionScClientForceWithdrawDeposit)
	if err := _LitionScClient.contract.UnpackLog(event, "ForceWithdrawDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientForceWithdrawVestingIterator is returned from FilterForceWithdrawVesting and is used to iterate over the raw logs and unpacked data for ForceWithdrawVesting events raised by the LitionScClient contract.
type LitionScClientForceWithdrawVestingIterator struct {
	Event *LitionScClientForceWithdrawVesting // Event containing the contract specifics and raw log

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
func (it *LitionScClientForceWithdrawVestingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientForceWithdrawVesting)
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
		it.Event = new(LitionScClientForceWithdrawVesting)
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
func (it *LitionScClientForceWithdrawVestingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientForceWithdrawVestingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientForceWithdrawVesting represents a ForceWithdrawVesting event raised by the LitionScClient contract.
type LitionScClientForceWithdrawVesting struct {
	ChainId   *big.Int
	Account   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForceWithdrawVesting is a free log retrieval operation binding the contract event 0xa3384ea61d5ab8755c01442bce838b3b1235b6f18e294620c2bc49de0073854f.
//
// Solidity: event ForceWithdrawVesting(uint256 indexed chainId, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterForceWithdrawVesting(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientForceWithdrawVestingIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ForceWithdrawVesting", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientForceWithdrawVestingIterator{contract: _LitionScClient.contract, event: "ForceWithdrawVesting", logs: logs, sub: sub}, nil
}

// WatchForceWithdrawVesting is a free log subscription operation binding the contract event 0xa3384ea61d5ab8755c01442bce838b3b1235b6f18e294620c2bc49de0073854f.
//
// Solidity: event ForceWithdrawVesting(uint256 indexed chainId, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchForceWithdrawVesting(opts *bind.WatchOpts, sink chan<- *LitionScClientForceWithdrawVesting, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ForceWithdrawVesting", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientForceWithdrawVesting)
				if err := _LitionScClient.contract.UnpackLog(event, "ForceWithdrawVesting", log); err != nil {
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

// ParseForceWithdrawVesting is a log parse operation binding the contract event 0xa3384ea61d5ab8755c01442bce838b3b1235b6f18e294620c2bc49de0073854f.
//
// Solidity: event ForceWithdrawVesting(uint256 indexed chainId, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) ParseForceWithdrawVesting(log types.Log) (*LitionScClientForceWithdrawVesting, error) {
	event := new(LitionScClientForceWithdrawVesting)
	if err := _LitionScClient.contract.UnpackLog(event, "ForceWithdrawVesting", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientNewChainIterator is returned from FilterNewChain and is used to iterate over the raw logs and unpacked data for NewChain events raised by the LitionScClient contract.
type LitionScClientNewChainIterator struct {
	Event *LitionScClientNewChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientNewChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientNewChain)
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
		it.Event = new(LitionScClientNewChain)
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
func (it *LitionScClientNewChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientNewChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientNewChain represents a NewChain event raised by the LitionScClient contract.
type LitionScClientNewChain struct {
	ChainId     *big.Int
	Description string
	Endpoint    string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewChain is a free log retrieval operation binding the contract event 0x86463e5a4c44c4d307742ef0abc183642f207e46cbaa6411b2bf7a118ab893e2.
//
// Solidity: event NewChain(uint256 chainId, string description, string endpoint)
func (_LitionScClient *LitionScClientFilterer) FilterNewChain(opts *bind.FilterOpts) (*LitionScClientNewChainIterator, error) {

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "NewChain")
	if err != nil {
		return nil, err
	}
	return &LitionScClientNewChainIterator{contract: _LitionScClient.contract, event: "NewChain", logs: logs, sub: sub}, nil
}

// WatchNewChain is a free log subscription operation binding the contract event 0x86463e5a4c44c4d307742ef0abc183642f207e46cbaa6411b2bf7a118ab893e2.
//
// Solidity: event NewChain(uint256 chainId, string description, string endpoint)
func (_LitionScClient *LitionScClientFilterer) WatchNewChain(opts *bind.WatchOpts, sink chan<- *LitionScClientNewChain) (event.Subscription, error) {

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "NewChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientNewChain)
				if err := _LitionScClient.contract.UnpackLog(event, "NewChain", log); err != nil {
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

// ParseNewChain is a log parse operation binding the contract event 0x86463e5a4c44c4d307742ef0abc183642f207e46cbaa6411b2bf7a118ab893e2.
//
// Solidity: event NewChain(uint256 chainId, string description, string endpoint)
func (_LitionScClient *LitionScClientFilterer) ParseNewChain(log types.Log) (*LitionScClientNewChain, error) {
	event := new(LitionScClientNewChain)
	if err := _LitionScClient.contract.UnpackLog(event, "NewChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientRequestDepositInChainIterator is returned from FilterRequestDepositInChain and is used to iterate over the raw logs and unpacked data for RequestDepositInChain events raised by the LitionScClient contract.
type LitionScClientRequestDepositInChainIterator struct {
	Event *LitionScClientRequestDepositInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientRequestDepositInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientRequestDepositInChain)
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
		it.Event = new(LitionScClientRequestDepositInChain)
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
func (it *LitionScClientRequestDepositInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientRequestDepositInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientRequestDepositInChain represents a RequestDepositInChain event raised by the LitionScClient contract.
type LitionScClientRequestDepositInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Deposit      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestDepositInChain is a free log retrieval operation binding the contract event 0x11b73dd0023b91d2be09f591166aeda0a3794d5c5d63607603e5d31c053c4804.
//
// Solidity: event RequestDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterRequestDepositInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientRequestDepositInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "RequestDepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientRequestDepositInChainIterator{contract: _LitionScClient.contract, event: "RequestDepositInChain", logs: logs, sub: sub}, nil
}

// WatchRequestDepositInChain is a free log subscription operation binding the contract event 0x11b73dd0023b91d2be09f591166aeda0a3794d5c5d63607603e5d31c053c4804.
//
// Solidity: event RequestDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchRequestDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientRequestDepositInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "RequestDepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientRequestDepositInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "RequestDepositInChain", log); err != nil {
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

// ParseRequestDepositInChain is a log parse operation binding the contract event 0x11b73dd0023b91d2be09f591166aeda0a3794d5c5d63607603e5d31c053c4804.
//
// Solidity: event RequestDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseRequestDepositInChain(log types.Log) (*LitionScClientRequestDepositInChain, error) {
	event := new(LitionScClientRequestDepositInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "RequestDepositInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientRequestVestInChainIterator is returned from FilterRequestVestInChain and is used to iterate over the raw logs and unpacked data for RequestVestInChain events raised by the LitionScClient contract.
type LitionScClientRequestVestInChainIterator struct {
	Event *LitionScClientRequestVestInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientRequestVestInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientRequestVestInChain)
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
		it.Event = new(LitionScClientRequestVestInChain)
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
func (it *LitionScClientRequestVestInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientRequestVestInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientRequestVestInChain represents a RequestVestInChain event raised by the LitionScClient contract.
type LitionScClientRequestVestInChain struct {
	ChainId      *big.Int
	Account      common.Address
	Vesting      *big.Int
	ReqTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestVestInChain is a free log retrieval operation binding the contract event 0x5be03bab1e7e8d384248dfcee1d12d03a399213ac09c422aac74b6e766de144d.
//
// Solidity: event RequestVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) FilterRequestVestInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientRequestVestInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "RequestVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientRequestVestInChainIterator{contract: _LitionScClient.contract, event: "RequestVestInChain", logs: logs, sub: sub}, nil
}

// WatchRequestVestInChain is a free log subscription operation binding the contract event 0x5be03bab1e7e8d384248dfcee1d12d03a399213ac09c422aac74b6e766de144d.
//
// Solidity: event RequestVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) WatchRequestVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientRequestVestInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "RequestVestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientRequestVestInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "RequestVestInChain", log); err != nil {
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

// ParseRequestVestInChain is a log parse operation binding the contract event 0x5be03bab1e7e8d384248dfcee1d12d03a399213ac09c422aac74b6e766de144d.
//
// Solidity: event RequestVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp)
func (_LitionScClient *LitionScClientFilterer) ParseRequestVestInChain(log types.Log) (*LitionScClientRequestVestInChain, error) {
	event := new(LitionScClientRequestVestInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "RequestVestInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientStartMiningIterator is returned from FilterStartMining and is used to iterate over the raw logs and unpacked data for StartMining events raised by the LitionScClient contract.
type LitionScClientStartMiningIterator struct {
	Event *LitionScClientStartMining // Event containing the contract specifics and raw log

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
func (it *LitionScClientStartMiningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientStartMining)
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
		it.Event = new(LitionScClientStartMining)
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
func (it *LitionScClientStartMiningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientStartMiningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientStartMining represents a StartMining event raised by the LitionScClient contract.
type LitionScClientStartMining struct {
	ChainId *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStartMining is a free log retrieval operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chainId, address indexed account)
func (_LitionScClient *LitionScClientFilterer) FilterStartMining(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientStartMiningIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "StartMining", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientStartMiningIterator{contract: _LitionScClient.contract, event: "StartMining", logs: logs, sub: sub}, nil
}

// WatchStartMining is a free log subscription operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chainId, address indexed account)
func (_LitionScClient *LitionScClientFilterer) WatchStartMining(opts *bind.WatchOpts, sink chan<- *LitionScClientStartMining, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "StartMining", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientStartMining)
				if err := _LitionScClient.contract.UnpackLog(event, "StartMining", log); err != nil {
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

// ParseStartMining is a log parse operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chainId, address indexed account)
func (_LitionScClient *LitionScClientFilterer) ParseStartMining(log types.Log) (*LitionScClientStartMining, error) {
	event := new(LitionScClientStartMining)
	if err := _LitionScClient.contract.UnpackLog(event, "StartMining", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientStopMiningIterator is returned from FilterStopMining and is used to iterate over the raw logs and unpacked data for StopMining events raised by the LitionScClient contract.
type LitionScClientStopMiningIterator struct {
	Event *LitionScClientStopMining // Event containing the contract specifics and raw log

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
func (it *LitionScClientStopMiningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientStopMining)
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
		it.Event = new(LitionScClientStopMining)
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
func (it *LitionScClientStopMiningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientStopMiningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientStopMining represents a StopMining event raised by the LitionScClient contract.
type LitionScClientStopMining struct {
	ChainId *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStopMining is a free log retrieval operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chainId, address indexed account)
func (_LitionScClient *LitionScClientFilterer) FilterStopMining(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientStopMiningIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "StopMining", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientStopMiningIterator{contract: _LitionScClient.contract, event: "StopMining", logs: logs, sub: sub}, nil
}

// WatchStopMining is a free log subscription operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chainId, address indexed account)
func (_LitionScClient *LitionScClientFilterer) WatchStopMining(opts *bind.WatchOpts, sink chan<- *LitionScClientStopMining, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "StopMining", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientStopMining)
				if err := _LitionScClient.contract.UnpackLog(event, "StopMining", log); err != nil {
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

// ParseStopMining is a log parse operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chainId, address indexed account)
func (_LitionScClient *LitionScClientFilterer) ParseStopMining(log types.Log) (*LitionScClientStopMining, error) {
	event := new(LitionScClientStopMining)
	if err := _LitionScClient.contract.UnpackLog(event, "StopMining", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientWhitelistAccountIterator is returned from FilterWhitelistAccount and is used to iterate over the raw logs and unpacked data for WhitelistAccount events raised by the LitionScClient contract.
type LitionScClientWhitelistAccountIterator struct {
	Event *LitionScClientWhitelistAccount // Event containing the contract specifics and raw log

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
func (it *LitionScClientWhitelistAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientWhitelistAccount)
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
		it.Event = new(LitionScClientWhitelistAccount)
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
func (it *LitionScClientWhitelistAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientWhitelistAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientWhitelistAccount represents a WhitelistAccount event raised by the LitionScClient contract.
type LitionScClientWhitelistAccount struct {
	ChainId   *big.Int
	Account   common.Address
	Whitelist bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAccount is a free log retrieval operation binding the contract event 0x2c8cbcff405aa83b0e160273c7f1159141d2bdc795ecddf1e3413dc2995f0d7f.
//
// Solidity: event WhitelistAccount(uint256 indexed chainId, address indexed account, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) FilterWhitelistAccount(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientWhitelistAccountIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "WhitelistAccount", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientWhitelistAccountIterator{contract: _LitionScClient.contract, event: "WhitelistAccount", logs: logs, sub: sub}, nil
}

// WatchWhitelistAccount is a free log subscription operation binding the contract event 0x2c8cbcff405aa83b0e160273c7f1159141d2bdc795ecddf1e3413dc2995f0d7f.
//
// Solidity: event WhitelistAccount(uint256 indexed chainId, address indexed account, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) WatchWhitelistAccount(opts *bind.WatchOpts, sink chan<- *LitionScClientWhitelistAccount, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "WhitelistAccount", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientWhitelistAccount)
				if err := _LitionScClient.contract.UnpackLog(event, "WhitelistAccount", log); err != nil {
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

// ParseWhitelistAccount is a log parse operation binding the contract event 0x2c8cbcff405aa83b0e160273c7f1159141d2bdc795ecddf1e3413dc2995f0d7f.
//
// Solidity: event WhitelistAccount(uint256 indexed chainId, address indexed account, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) ParseWhitelistAccount(log types.Log) (*LitionScClientWhitelistAccount, error) {
	event := new(LitionScClientWhitelistAccount)
	if err := _LitionScClient.contract.UnpackLog(event, "WhitelistAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}
