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
const LitionScClientABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"confirmVestInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"getAllowedToValidate\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"validators\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"getTransactors\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"transactors\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"isAllowedToTransact\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"stopMining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"startMining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notaryStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notaryEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocksMined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"userGas\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"largestTx\",\"type\":\"uint64\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"confirmDepositWithdrawalFromChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getChainStaticDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxNumOfValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumOfTransactors\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"involvedVestingNotaryCond\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"participationNotaryCond\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"requestDepositInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"initEndpoint\",\"type\":\"string\"},{\"internalType\":\"contractChainValidator\",\"name\":\"chainValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumOfValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumOfTransactors\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"involvedVestingNotaryCond\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"participationNotaryCond\",\"type\":\"bool\"}],\"name\":\"registerChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"isAllowedToValidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"validators\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getChainDynamicDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalVesting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastValidatorVesting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastNotaryBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastNotaryTimestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getUserDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"prevNotaryMined\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"secondPrevNotaryMined\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"vestingReqExist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqNotary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingReqValue\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"depositFullWithdrawalReqExist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"depositReqNotary\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"testNotary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"setChainStaticDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"}],\"name\":\"requestVestInChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastNotaryBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"name\":\"DepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastNotaryBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"name\":\"VestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"AccountWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"}],\"name\":\"AccountMining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"notaryBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"name\":\"Notary\",\"type\":\"event\"}]"

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

// GetAllowedToValidate is a free data retrieval call binding the contract method 0x148dbe94.
//
// Solidity: function getAllowedToValidate(uint256 chainId, uint256 batch) constant returns(address[100] validators, uint256 count, bool end)
func (_LitionScClient *LitionScClientCaller) GetAllowedToValidate(opts *bind.CallOpts, chainId *big.Int, batch *big.Int) (struct {
	Validators [100]common.Address
	Count      *big.Int
	End        bool
}, error) {
	ret := new(struct {
		Validators [100]common.Address
		Count      *big.Int
		End        bool
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getAllowedToValidate", chainId, batch)
	return *ret, err
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0x148dbe94.
//
// Solidity: function getAllowedToValidate(uint256 chainId, uint256 batch) constant returns(address[100] validators, uint256 count, bool end)
func (_LitionScClient *LitionScClientSession) GetAllowedToValidate(chainId *big.Int, batch *big.Int) (struct {
	Validators [100]common.Address
	Count      *big.Int
	End        bool
}, error) {
	return _LitionScClient.Contract.GetAllowedToValidate(&_LitionScClient.CallOpts, chainId, batch)
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0x148dbe94.
//
// Solidity: function getAllowedToValidate(uint256 chainId, uint256 batch) constant returns(address[100] validators, uint256 count, bool end)
func (_LitionScClient *LitionScClientCallerSession) GetAllowedToValidate(chainId *big.Int, batch *big.Int) (struct {
	Validators [100]common.Address
	Count      *big.Int
	End        bool
}, error) {
	return _LitionScClient.Contract.GetAllowedToValidate(&_LitionScClient.CallOpts, chainId, batch)
}

// GetChainDynamicDetails is a free data retrieval call binding the contract method 0xc660c93e.
//
// Solidity: function getChainDynamicDetails(uint256 chainId) constant returns(bool active, uint256 totalVesting, uint256 validatorsCount, uint256 transactorsCount, uint256 lastValidatorVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp)
func (_LitionScClient *LitionScClientCaller) GetChainDynamicDetails(opts *bind.CallOpts, chainId *big.Int) (struct {
	Active               bool
	TotalVesting         *big.Int
	ValidatorsCount      *big.Int
	TransactorsCount     *big.Int
	LastValidatorVesting *big.Int
	LastNotaryBlock      *big.Int
	LastNotaryTimestamp  *big.Int
}, error) {
	ret := new(struct {
		Active               bool
		TotalVesting         *big.Int
		ValidatorsCount      *big.Int
		TransactorsCount     *big.Int
		LastValidatorVesting *big.Int
		LastNotaryBlock      *big.Int
		LastNotaryTimestamp  *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getChainDynamicDetails", chainId)
	return *ret, err
}

// GetChainDynamicDetails is a free data retrieval call binding the contract method 0xc660c93e.
//
// Solidity: function getChainDynamicDetails(uint256 chainId) constant returns(bool active, uint256 totalVesting, uint256 validatorsCount, uint256 transactorsCount, uint256 lastValidatorVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp)
func (_LitionScClient *LitionScClientSession) GetChainDynamicDetails(chainId *big.Int) (struct {
	Active               bool
	TotalVesting         *big.Int
	ValidatorsCount      *big.Int
	TransactorsCount     *big.Int
	LastValidatorVesting *big.Int
	LastNotaryBlock      *big.Int
	LastNotaryTimestamp  *big.Int
}, error) {
	return _LitionScClient.Contract.GetChainDynamicDetails(&_LitionScClient.CallOpts, chainId)
}

// GetChainDynamicDetails is a free data retrieval call binding the contract method 0xc660c93e.
//
// Solidity: function getChainDynamicDetails(uint256 chainId) constant returns(bool active, uint256 totalVesting, uint256 validatorsCount, uint256 transactorsCount, uint256 lastValidatorVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp)
func (_LitionScClient *LitionScClientCallerSession) GetChainDynamicDetails(chainId *big.Int) (struct {
	Active               bool
	TotalVesting         *big.Int
	ValidatorsCount      *big.Int
	TransactorsCount     *big.Int
	LastValidatorVesting *big.Int
	LastNotaryBlock      *big.Int
	LastNotaryTimestamp  *big.Int
}, error) {
	return _LitionScClient.Contract.GetChainDynamicDetails(&_LitionScClient.CallOpts, chainId)
}

// GetChainStaticDetails is a free data retrieval call binding the contract method 0x8ea1e8c3.
//
// Solidity: function getChainStaticDetails(uint256 chainId) constant returns(string description, string endpoint, bool registered, uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond)
func (_LitionScClient *LitionScClientCaller) GetChainStaticDetails(opts *bind.CallOpts, chainId *big.Int) (struct {
	Description               string
	Endpoint                  string
	Registered                bool
	MaxNumOfValidators        *big.Int
	MaxNumOfTransactors       *big.Int
	InvolvedVestingNotaryCond bool
	ParticipationNotaryCond   bool
}, error) {
	ret := new(struct {
		Description               string
		Endpoint                  string
		Registered                bool
		MaxNumOfValidators        *big.Int
		MaxNumOfTransactors       *big.Int
		InvolvedVestingNotaryCond bool
		ParticipationNotaryCond   bool
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getChainStaticDetails", chainId)
	return *ret, err
}

// GetChainStaticDetails is a free data retrieval call binding the contract method 0x8ea1e8c3.
//
// Solidity: function getChainStaticDetails(uint256 chainId) constant returns(string description, string endpoint, bool registered, uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond)
func (_LitionScClient *LitionScClientSession) GetChainStaticDetails(chainId *big.Int) (struct {
	Description               string
	Endpoint                  string
	Registered                bool
	MaxNumOfValidators        *big.Int
	MaxNumOfTransactors       *big.Int
	InvolvedVestingNotaryCond bool
	ParticipationNotaryCond   bool
}, error) {
	return _LitionScClient.Contract.GetChainStaticDetails(&_LitionScClient.CallOpts, chainId)
}

// GetChainStaticDetails is a free data retrieval call binding the contract method 0x8ea1e8c3.
//
// Solidity: function getChainStaticDetails(uint256 chainId) constant returns(string description, string endpoint, bool registered, uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond)
func (_LitionScClient *LitionScClientCallerSession) GetChainStaticDetails(chainId *big.Int) (struct {
	Description               string
	Endpoint                  string
	Registered                bool
	MaxNumOfValidators        *big.Int
	MaxNumOfTransactors       *big.Int
	InvolvedVestingNotaryCond bool
	ParticipationNotaryCond   bool
}, error) {
	return _LitionScClient.Contract.GetChainStaticDetails(&_LitionScClient.CallOpts, chainId)
}

// GetTransactors is a free data retrieval call binding the contract method 0x2d94c932.
//
// Solidity: function getTransactors(uint256 chainId, uint256 batch) constant returns(address[100] transactors, uint256 count, bool end)
func (_LitionScClient *LitionScClientCaller) GetTransactors(opts *bind.CallOpts, chainId *big.Int, batch *big.Int) (struct {
	Transactors [100]common.Address
	Count       *big.Int
	End         bool
}, error) {
	ret := new(struct {
		Transactors [100]common.Address
		Count       *big.Int
		End         bool
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getTransactors", chainId, batch)
	return *ret, err
}

// GetTransactors is a free data retrieval call binding the contract method 0x2d94c932.
//
// Solidity: function getTransactors(uint256 chainId, uint256 batch) constant returns(address[100] transactors, uint256 count, bool end)
func (_LitionScClient *LitionScClientSession) GetTransactors(chainId *big.Int, batch *big.Int) (struct {
	Transactors [100]common.Address
	Count       *big.Int
	End         bool
}, error) {
	return _LitionScClient.Contract.GetTransactors(&_LitionScClient.CallOpts, chainId, batch)
}

// GetTransactors is a free data retrieval call binding the contract method 0x2d94c932.
//
// Solidity: function getTransactors(uint256 chainId, uint256 batch) constant returns(address[100] transactors, uint256 count, bool end)
func (_LitionScClient *LitionScClientCallerSession) GetTransactors(chainId *big.Int, batch *big.Int) (struct {
	Transactors [100]common.Address
	Count       *big.Int
	End         bool
}, error) {
	return _LitionScClient.Contract.GetTransactors(&_LitionScClient.CallOpts, chainId, batch)
}

// GetUserDetails is a free data retrieval call binding the contract method 0xc90902cb.
//
// Solidity: function getUserDetails(uint256 chainId, address acc) constant returns(uint256 deposit, bool whitelisted, uint256 vesting, bool mining, bool prevNotaryMined, bool secondPrevNotaryMined, bool vestingReqExist, uint256 vestingReqNotary, uint256 vestingReqValue, bool depositFullWithdrawalReqExist, uint256 depositReqNotary)
func (_LitionScClient *LitionScClientCaller) GetUserDetails(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (struct {
	Deposit                       *big.Int
	Whitelisted                   bool
	Vesting                       *big.Int
	Mining                        bool
	PrevNotaryMined               bool
	SecondPrevNotaryMined         bool
	VestingReqExist               bool
	VestingReqNotary              *big.Int
	VestingReqValue               *big.Int
	DepositFullWithdrawalReqExist bool
	DepositReqNotary              *big.Int
}, error) {
	ret := new(struct {
		Deposit                       *big.Int
		Whitelisted                   bool
		Vesting                       *big.Int
		Mining                        bool
		PrevNotaryMined               bool
		SecondPrevNotaryMined         bool
		VestingReqExist               bool
		VestingReqNotary              *big.Int
		VestingReqValue               *big.Int
		DepositFullWithdrawalReqExist bool
		DepositReqNotary              *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getUserDetails", chainId, acc)
	return *ret, err
}

// GetUserDetails is a free data retrieval call binding the contract method 0xc90902cb.
//
// Solidity: function getUserDetails(uint256 chainId, address acc) constant returns(uint256 deposit, bool whitelisted, uint256 vesting, bool mining, bool prevNotaryMined, bool secondPrevNotaryMined, bool vestingReqExist, uint256 vestingReqNotary, uint256 vestingReqValue, bool depositFullWithdrawalReqExist, uint256 depositReqNotary)
func (_LitionScClient *LitionScClientSession) GetUserDetails(chainId *big.Int, acc common.Address) (struct {
	Deposit                       *big.Int
	Whitelisted                   bool
	Vesting                       *big.Int
	Mining                        bool
	PrevNotaryMined               bool
	SecondPrevNotaryMined         bool
	VestingReqExist               bool
	VestingReqNotary              *big.Int
	VestingReqValue               *big.Int
	DepositFullWithdrawalReqExist bool
	DepositReqNotary              *big.Int
}, error) {
	return _LitionScClient.Contract.GetUserDetails(&_LitionScClient.CallOpts, chainId, acc)
}

// GetUserDetails is a free data retrieval call binding the contract method 0xc90902cb.
//
// Solidity: function getUserDetails(uint256 chainId, address acc) constant returns(uint256 deposit, bool whitelisted, uint256 vesting, bool mining, bool prevNotaryMined, bool secondPrevNotaryMined, bool vestingReqExist, uint256 vestingReqNotary, uint256 vestingReqValue, bool depositFullWithdrawalReqExist, uint256 depositReqNotary)
func (_LitionScClient *LitionScClientCallerSession) GetUserDetails(chainId *big.Int, acc common.Address) (struct {
	Deposit                       *big.Int
	Whitelisted                   bool
	Vesting                       *big.Int
	Mining                        bool
	PrevNotaryMined               bool
	SecondPrevNotaryMined         bool
	VestingReqExist               bool
	VestingReqNotary              *big.Int
	VestingReqValue               *big.Int
	DepositFullWithdrawalReqExist bool
	DepositReqNotary              *big.Int
}, error) {
	return _LitionScClient.Contract.GetUserDetails(&_LitionScClient.CallOpts, chainId, acc)
}

// GetValidators is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 chainId, uint256 batch) constant returns(address[100] validators, uint256 count, bool end)
func (_LitionScClient *LitionScClientCaller) GetValidators(opts *bind.CallOpts, chainId *big.Int, batch *big.Int) (struct {
	Validators [100]common.Address
	Count      *big.Int
	End        bool
}, error) {
	ret := new(struct {
		Validators [100]common.Address
		Count      *big.Int
		End        bool
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "getValidators", chainId, batch)
	return *ret, err
}

// GetValidators is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 chainId, uint256 batch) constant returns(address[100] validators, uint256 count, bool end)
func (_LitionScClient *LitionScClientSession) GetValidators(chainId *big.Int, batch *big.Int) (struct {
	Validators [100]common.Address
	Count      *big.Int
	End        bool
}, error) {
	return _LitionScClient.Contract.GetValidators(&_LitionScClient.CallOpts, chainId, batch)
}

// GetValidators is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 chainId, uint256 batch) constant returns(address[100] validators, uint256 count, bool end)
func (_LitionScClient *LitionScClientCallerSession) GetValidators(chainId *big.Int, batch *big.Int) (struct {
	Validators [100]common.Address
	Count      *big.Int
	End        bool
}, error) {
	return _LitionScClient.Contract.GetValidators(&_LitionScClient.CallOpts, chainId, batch)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0xff347ae8.
//
// Solidity: function isActiveValidator(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) IsActiveValidator(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "isActiveValidator", chainId, acc)
	return *ret0, err
}

// IsActiveValidator is a free data retrieval call binding the contract method 0xff347ae8.
//
// Solidity: function isActiveValidator(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) IsActiveValidator(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.IsActiveValidator(&_LitionScClient.CallOpts, chainId, acc)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0xff347ae8.
//
// Solidity: function isActiveValidator(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) IsActiveValidator(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.IsActiveValidator(&_LitionScClient.CallOpts, chainId, acc)
}

// IsAllowedToTransact is a free data retrieval call binding the contract method 0x36f3eecb.
//
// Solidity: function isAllowedToTransact(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) IsAllowedToTransact(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "isAllowedToTransact", chainId, acc)
	return *ret0, err
}

// IsAllowedToTransact is a free data retrieval call binding the contract method 0x36f3eecb.
//
// Solidity: function isAllowedToTransact(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) IsAllowedToTransact(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.IsAllowedToTransact(&_LitionScClient.CallOpts, chainId, acc)
}

// IsAllowedToTransact is a free data retrieval call binding the contract method 0x36f3eecb.
//
// Solidity: function isAllowedToTransact(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) IsAllowedToTransact(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.IsAllowedToTransact(&_LitionScClient.CallOpts, chainId, acc)
}

// IsAllowedToValidate is a free data retrieval call binding the contract method 0xb5fc7ef7.
//
// Solidity: function isAllowedToValidate(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) IsAllowedToValidate(opts *bind.CallOpts, chainId *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "isAllowedToValidate", chainId, acc)
	return *ret0, err
}

// IsAllowedToValidate is a free data retrieval call binding the contract method 0xb5fc7ef7.
//
// Solidity: function isAllowedToValidate(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) IsAllowedToValidate(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.IsAllowedToValidate(&_LitionScClient.CallOpts, chainId, acc)
}

// IsAllowedToValidate is a free data retrieval call binding the contract method 0xb5fc7ef7.
//
// Solidity: function isAllowedToValidate(uint256 chainId, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) IsAllowedToValidate(chainId *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.IsAllowedToValidate(&_LitionScClient.CallOpts, chainId, acc)
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

// Notary is a paid mutator transaction binding the contract method 0x4d8ab6aa.
//
// Solidity: function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] validators, uint32[] blocksMined, address[] users, uint64[] userGas, uint64 largestTx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactor) Notary(opts *bind.TransactOpts, chainId *big.Int, notaryStartBlock *big.Int, notaryEndBlock *big.Int, validators []common.Address, blocksMined []uint32, users []common.Address, userGas []uint64, largestTx uint64, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "notary", chainId, notaryStartBlock, notaryEndBlock, validators, blocksMined, users, userGas, largestTx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0x4d8ab6aa.
//
// Solidity: function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] validators, uint32[] blocksMined, address[] users, uint64[] userGas, uint64 largestTx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientSession) Notary(chainId *big.Int, notaryStartBlock *big.Int, notaryEndBlock *big.Int, validators []common.Address, blocksMined []uint32, users []common.Address, userGas []uint64, largestTx uint64, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, chainId, notaryStartBlock, notaryEndBlock, validators, blocksMined, users, userGas, largestTx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0x4d8ab6aa.
//
// Solidity: function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] validators, uint32[] blocksMined, address[] users, uint64[] userGas, uint64 largestTx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactorSession) Notary(chainId *big.Int, notaryStartBlock *big.Int, notaryEndBlock *big.Int, validators []common.Address, blocksMined []uint32, users []common.Address, userGas []uint64, largestTx uint64, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, chainId, notaryStartBlock, notaryEndBlock, validators, blocksMined, users, userGas, largestTx, v, r, s)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xab1e2f74.
//
// Solidity: function registerChain(string description, string initEndpoint, address chainValidator, uint256 vesting, uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond) returns(uint256 chainId)
func (_LitionScClient *LitionScClientTransactor) RegisterChain(opts *bind.TransactOpts, description string, initEndpoint string, chainValidator common.Address, vesting *big.Int, maxNumOfValidators *big.Int, maxNumOfTransactors *big.Int, involvedVestingNotaryCond bool, participationNotaryCond bool) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "registerChain", description, initEndpoint, chainValidator, vesting, maxNumOfValidators, maxNumOfTransactors, involvedVestingNotaryCond, participationNotaryCond)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xab1e2f74.
//
// Solidity: function registerChain(string description, string initEndpoint, address chainValidator, uint256 vesting, uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond) returns(uint256 chainId)
func (_LitionScClient *LitionScClientSession) RegisterChain(description string, initEndpoint string, chainValidator common.Address, vesting *big.Int, maxNumOfValidators *big.Int, maxNumOfTransactors *big.Int, involvedVestingNotaryCond bool, participationNotaryCond bool) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, description, initEndpoint, chainValidator, vesting, maxNumOfValidators, maxNumOfTransactors, involvedVestingNotaryCond, participationNotaryCond)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xab1e2f74.
//
// Solidity: function registerChain(string description, string initEndpoint, address chainValidator, uint256 vesting, uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond) returns(uint256 chainId)
func (_LitionScClient *LitionScClientTransactorSession) RegisterChain(description string, initEndpoint string, chainValidator common.Address, vesting *big.Int, maxNumOfValidators *big.Int, maxNumOfTransactors *big.Int, involvedVestingNotaryCond bool, participationNotaryCond bool) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, description, initEndpoint, chainValidator, vesting, maxNumOfValidators, maxNumOfTransactors, involvedVestingNotaryCond, participationNotaryCond)
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

// SetChainStaticDetails is a paid mutator transaction binding the contract method 0xe2b39ae3.
//
// Solidity: function setChainStaticDetails(uint256 chainId, string description, string endpoint) returns()
func (_LitionScClient *LitionScClientTransactor) SetChainStaticDetails(opts *bind.TransactOpts, chainId *big.Int, description string, endpoint string) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "setChainStaticDetails", chainId, description, endpoint)
}

// SetChainStaticDetails is a paid mutator transaction binding the contract method 0xe2b39ae3.
//
// Solidity: function setChainStaticDetails(uint256 chainId, string description, string endpoint) returns()
func (_LitionScClient *LitionScClientSession) SetChainStaticDetails(chainId *big.Int, description string, endpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.SetChainStaticDetails(&_LitionScClient.TransactOpts, chainId, description, endpoint)
}

// SetChainStaticDetails is a paid mutator transaction binding the contract method 0xe2b39ae3.
//
// Solidity: function setChainStaticDetails(uint256 chainId, string description, string endpoint) returns()
func (_LitionScClient *LitionScClientTransactorSession) SetChainStaticDetails(chainId *big.Int, description string, endpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.SetChainStaticDetails(&_LitionScClient.TransactOpts, chainId, description, endpoint)
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

// TestNotary is a paid mutator transaction binding the contract method 0xdaada67b.
//
// Solidity: function testNotary(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactor) TestNotary(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "testNotary", chainId)
}

// TestNotary is a paid mutator transaction binding the contract method 0xdaada67b.
//
// Solidity: function testNotary(uint256 chainId) returns()
func (_LitionScClient *LitionScClientSession) TestNotary(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.TestNotary(&_LitionScClient.TransactOpts, chainId)
}

// TestNotary is a paid mutator transaction binding the contract method 0xdaada67b.
//
// Solidity: function testNotary(uint256 chainId) returns()
func (_LitionScClient *LitionScClientTransactorSession) TestNotary(chainId *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.TestNotary(&_LitionScClient.TransactOpts, chainId)
}

// LitionScClientAccountMiningIterator is returned from FilterAccountMining and is used to iterate over the raw logs and unpacked data for AccountMining events raised by the LitionScClient contract.
type LitionScClientAccountMiningIterator struct {
	Event *LitionScClientAccountMining // Event containing the contract specifics and raw log

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
func (it *LitionScClientAccountMiningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientAccountMining)
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
		it.Event = new(LitionScClientAccountMining)
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
func (it *LitionScClientAccountMiningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientAccountMiningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientAccountMining represents a AccountMining event raised by the LitionScClient contract.
type LitionScClientAccountMining struct {
	ChainId *big.Int
	Account common.Address
	Mining  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountMining is a free log retrieval operation binding the contract event 0xbdcecfacc87dc76e8dd0ce9ec36aafac83c67468e83445dac5e316e8a60824cb.
//
// Solidity: event AccountMining(uint256 indexed chainId, address indexed account, bool mining)
func (_LitionScClient *LitionScClientFilterer) FilterAccountMining(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientAccountMiningIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "AccountMining", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientAccountMiningIterator{contract: _LitionScClient.contract, event: "AccountMining", logs: logs, sub: sub}, nil
}

// WatchAccountMining is a free log subscription operation binding the contract event 0xbdcecfacc87dc76e8dd0ce9ec36aafac83c67468e83445dac5e316e8a60824cb.
//
// Solidity: event AccountMining(uint256 indexed chainId, address indexed account, bool mining)
func (_LitionScClient *LitionScClientFilterer) WatchAccountMining(opts *bind.WatchOpts, sink chan<- *LitionScClientAccountMining, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "AccountMining", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientAccountMining)
				if err := _LitionScClient.contract.UnpackLog(event, "AccountMining", log); err != nil {
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

// ParseAccountMining is a log parse operation binding the contract event 0xbdcecfacc87dc76e8dd0ce9ec36aafac83c67468e83445dac5e316e8a60824cb.
//
// Solidity: event AccountMining(uint256 indexed chainId, address indexed account, bool mining)
func (_LitionScClient *LitionScClientFilterer) ParseAccountMining(log types.Log) (*LitionScClientAccountMining, error) {
	event := new(LitionScClientAccountMining)
	if err := _LitionScClient.contract.UnpackLog(event, "AccountMining", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientAccountWhitelistIterator is returned from FilterAccountWhitelist and is used to iterate over the raw logs and unpacked data for AccountWhitelist events raised by the LitionScClient contract.
type LitionScClientAccountWhitelistIterator struct {
	Event *LitionScClientAccountWhitelist // Event containing the contract specifics and raw log

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
func (it *LitionScClientAccountWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientAccountWhitelist)
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
		it.Event = new(LitionScClientAccountWhitelist)
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
func (it *LitionScClientAccountWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientAccountWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientAccountWhitelist represents a AccountWhitelist event raised by the LitionScClient contract.
type LitionScClientAccountWhitelist struct {
	ChainId   *big.Int
	Account   common.Address
	Whitelist bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountWhitelist is a free log retrieval operation binding the contract event 0xc732c367e58816f41ca8b732ffca6bc97e995866c236984c146082ce8b1f49d0.
//
// Solidity: event AccountWhitelist(uint256 indexed chainId, address indexed account, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) FilterAccountWhitelist(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientAccountWhitelistIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "AccountWhitelist", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientAccountWhitelistIterator{contract: _LitionScClient.contract, event: "AccountWhitelist", logs: logs, sub: sub}, nil
}

// WatchAccountWhitelist is a free log subscription operation binding the contract event 0xc732c367e58816f41ca8b732ffca6bc97e995866c236984c146082ce8b1f49d0.
//
// Solidity: event AccountWhitelist(uint256 indexed chainId, address indexed account, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) WatchAccountWhitelist(opts *bind.WatchOpts, sink chan<- *LitionScClientAccountWhitelist, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "AccountWhitelist", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientAccountWhitelist)
				if err := _LitionScClient.contract.UnpackLog(event, "AccountWhitelist", log); err != nil {
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

// ParseAccountWhitelist is a log parse operation binding the contract event 0xc732c367e58816f41ca8b732ffca6bc97e995866c236984c146082ce8b1f49d0.
//
// Solidity: event AccountWhitelist(uint256 indexed chainId, address indexed account, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) ParseAccountWhitelist(log types.Log) (*LitionScClientAccountWhitelist, error) {
	event := new(LitionScClientAccountWhitelist)
	if err := _LitionScClient.contract.UnpackLog(event, "AccountWhitelist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientDepositInChainIterator is returned from FilterDepositInChain and is used to iterate over the raw logs and unpacked data for DepositInChain events raised by the LitionScClient contract.
type LitionScClientDepositInChainIterator struct {
	Event *LitionScClientDepositInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientDepositInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientDepositInChain)
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
		it.Event = new(LitionScClientDepositInChain)
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
func (it *LitionScClientDepositInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientDepositInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientDepositInChain represents a DepositInChain event raised by the LitionScClient contract.
type LitionScClientDepositInChain struct {
	ChainId         *big.Int
	Account         common.Address
	Deposit         *big.Int
	LastNotaryBlock *big.Int
	Confirmed       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositInChain is a free log retrieval operation binding the contract event 0xe56ded5a257157adb0d590ebbb8ca76f1e27d8df7dc2d7c42ea40e2269b634b4.
//
// Solidity: event DepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 lastNotaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) FilterDepositInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientDepositInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "DepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientDepositInChainIterator{contract: _LitionScClient.contract, event: "DepositInChain", logs: logs, sub: sub}, nil
}

// WatchDepositInChain is a free log subscription operation binding the contract event 0xe56ded5a257157adb0d590ebbb8ca76f1e27d8df7dc2d7c42ea40e2269b634b4.
//
// Solidity: event DepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 lastNotaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) WatchDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientDepositInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "DepositInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientDepositInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "DepositInChain", log); err != nil {
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

// ParseDepositInChain is a log parse operation binding the contract event 0xe56ded5a257157adb0d590ebbb8ca76f1e27d8df7dc2d7c42ea40e2269b634b4.
//
// Solidity: event DepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 lastNotaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) ParseDepositInChain(log types.Log) (*LitionScClientDepositInChain, error) {
	event := new(LitionScClientDepositInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "DepositInChain", log); err != nil {
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

// LitionScClientNotaryIterator is returned from FilterNotary and is used to iterate over the raw logs and unpacked data for Notary events raised by the LitionScClient contract.
type LitionScClientNotaryIterator struct {
	Event *LitionScClientNotary // Event containing the contract specifics and raw log

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
func (it *LitionScClientNotaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientNotary)
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
		it.Event = new(LitionScClientNotary)
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
func (it *LitionScClientNotaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientNotaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientNotary represents a Notary event raised by the LitionScClient contract.
type LitionScClientNotary struct {
	ChainId     *big.Int
	NotaryBlock *big.Int
	Confirmed   bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNotary is a free log retrieval operation binding the contract event 0x2bf6913b394ab8c4cb188e7c518c7f79fa5ca7a96701d295071195c41bcd6457.
//
// Solidity: event Notary(uint256 indexed chainId, uint256 notaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) FilterNotary(opts *bind.FilterOpts, chainId []*big.Int) (*LitionScClientNotaryIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "Notary", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientNotaryIterator{contract: _LitionScClient.contract, event: "Notary", logs: logs, sub: sub}, nil
}

// WatchNotary is a free log subscription operation binding the contract event 0x2bf6913b394ab8c4cb188e7c518c7f79fa5ca7a96701d295071195c41bcd6457.
//
// Solidity: event Notary(uint256 indexed chainId, uint256 notaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) WatchNotary(opts *bind.WatchOpts, sink chan<- *LitionScClientNotary, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "Notary", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientNotary)
				if err := _LitionScClient.contract.UnpackLog(event, "Notary", log); err != nil {
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

// ParseNotary is a log parse operation binding the contract event 0x2bf6913b394ab8c4cb188e7c518c7f79fa5ca7a96701d295071195c41bcd6457.
//
// Solidity: event Notary(uint256 indexed chainId, uint256 notaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) ParseNotary(log types.Log) (*LitionScClientNotary, error) {
	event := new(LitionScClientNotary)
	if err := _LitionScClient.contract.UnpackLog(event, "Notary", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientVestInChainIterator is returned from FilterVestInChain and is used to iterate over the raw logs and unpacked data for VestInChain events raised by the LitionScClient contract.
type LitionScClientVestInChainIterator struct {
	Event *LitionScClientVestInChain // Event containing the contract specifics and raw log

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
func (it *LitionScClientVestInChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientVestInChain)
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
		it.Event = new(LitionScClientVestInChain)
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
func (it *LitionScClientVestInChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientVestInChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientVestInChain represents a VestInChain event raised by the LitionScClient contract.
type LitionScClientVestInChain struct {
	ChainId         *big.Int
	Account         common.Address
	Vesting         *big.Int
	LastNotaryBlock *big.Int
	Confirmed       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVestInChain is a free log retrieval operation binding the contract event 0xcdf71206e0f94bf5111dc95b5392b83fd7f390a69ec0afe1181cc0edfaee834c.
//
// Solidity: event VestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 lastNotaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) FilterVestInChain(opts *bind.FilterOpts, chainId []*big.Int, account []common.Address) (*LitionScClientVestInChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "VestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientVestInChainIterator{contract: _LitionScClient.contract, event: "VestInChain", logs: logs, sub: sub}, nil
}

// WatchVestInChain is a free log subscription operation binding the contract event 0xcdf71206e0f94bf5111dc95b5392b83fd7f390a69ec0afe1181cc0edfaee834c.
//
// Solidity: event VestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 lastNotaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) WatchVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientVestInChain, chainId []*big.Int, account []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "VestInChain", chainIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientVestInChain)
				if err := _LitionScClient.contract.UnpackLog(event, "VestInChain", log); err != nil {
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

// ParseVestInChain is a log parse operation binding the contract event 0xcdf71206e0f94bf5111dc95b5392b83fd7f390a69ec0afe1181cc0edfaee834c.
//
// Solidity: event VestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 lastNotaryBlock, bool confirmed)
func (_LitionScClient *LitionScClientFilterer) ParseVestInChain(log types.Log) (*LitionScClientVestInChain, error) {
	event := new(LitionScClientVestInChain)
	if err := _LitionScClient.contract.UnpackLog(event, "VestInChain", log); err != nil {
		return nil, err
	}
	return event, nil
}
