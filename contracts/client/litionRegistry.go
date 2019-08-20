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
const LitionScClientABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notary_block\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"user_gas\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"largest_tx\",\"type\":\"uint32\"}],\"name\":\"get_signature_hash_from_notary\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"get_validators\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"users\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"start_mining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"last_notary\",\"type\":\"uint256\"},{\"internalType\":\"contractChainValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total_vesting\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"stop_mining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"}],\"name\":\"vest_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"notary_block_no\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"user_gas\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"largest_tx\",\"type\":\"uint32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"get_allowed_to_transact\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"users\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"has_deposited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"contractChainValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"init_endpoint\",\"type\":\"string\"}],\"name\":\"register_chain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"has_vested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"deposit_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"get_last_notary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"NewChainEndpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"datetime\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"datetime\",\"type\":\"uint256\"}],\"name\":\"Vesting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"StartMining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"StopMining\",\"type\":\"event\"}]"

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

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) constant returns(bool active, uint256 last_notary, address validator, uint256 total_vesting)
func (_LitionScClient *LitionScClientCaller) Chains(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Active       bool
	LastNotary   *big.Int
	Validator    common.Address
	TotalVesting *big.Int
}, error) {
	ret := new(struct {
		Active       bool
		LastNotary   *big.Int
		Validator    common.Address
		TotalVesting *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "chains", arg0)
	return *ret, err
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) constant returns(bool active, uint256 last_notary, address validator, uint256 total_vesting)
func (_LitionScClient *LitionScClientSession) Chains(arg0 *big.Int) (struct {
	Active       bool
	LastNotary   *big.Int
	Validator    common.Address
	TotalVesting *big.Int
}, error) {
	return _LitionScClient.Contract.Chains(&_LitionScClient.CallOpts, arg0)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) constant returns(bool active, uint256 last_notary, address validator, uint256 total_vesting)
func (_LitionScClient *LitionScClientCallerSession) Chains(arg0 *big.Int) (struct {
	Active       bool
	LastNotary   *big.Int
	Validator    common.Address
	TotalVesting *big.Int
}, error) {
	return _LitionScClient.Contract.Chains(&_LitionScClient.CallOpts, arg0)
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0xb190042c.
//
// Solidity: function get_allowed_to_transact(uint256 id, uint256 batch) constant returns(address[100] users, uint256 count)
func (_LitionScClient *LitionScClientCaller) GetAllowedToTransact(opts *bind.CallOpts, id *big.Int, batch *big.Int) (struct {
	Users [100]common.Address
	Count *big.Int
}, error) {
	ret := new(struct {
		Users [100]common.Address
		Count *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "get_allowed_to_transact", id, batch)
	return *ret, err
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0xb190042c.
//
// Solidity: function get_allowed_to_transact(uint256 id, uint256 batch) constant returns(address[100] users, uint256 count)
func (_LitionScClient *LitionScClientSession) GetAllowedToTransact(id *big.Int, batch *big.Int) (struct {
	Users [100]common.Address
	Count *big.Int
}, error) {
	return _LitionScClient.Contract.GetAllowedToTransact(&_LitionScClient.CallOpts, id, batch)
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0xb190042c.
//
// Solidity: function get_allowed_to_transact(uint256 id, uint256 batch) constant returns(address[100] users, uint256 count)
func (_LitionScClient *LitionScClientCallerSession) GetAllowedToTransact(id *big.Int, batch *big.Int) (struct {
	Users [100]common.Address
	Count *big.Int
}, error) {
	return _LitionScClient.Contract.GetAllowedToTransact(&_LitionScClient.CallOpts, id, batch)
}

// GetLastNotary is a free data retrieval call binding the contract method 0xda26501d.
//
// Solidity: function get_last_notary(uint256 id) constant returns(uint256)
func (_LitionScClient *LitionScClientCaller) GetLastNotary(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "get_last_notary", id)
	return *ret0, err
}

// GetLastNotary is a free data retrieval call binding the contract method 0xda26501d.
//
// Solidity: function get_last_notary(uint256 id) constant returns(uint256)
func (_LitionScClient *LitionScClientSession) GetLastNotary(id *big.Int) (*big.Int, error) {
	return _LitionScClient.Contract.GetLastNotary(&_LitionScClient.CallOpts, id)
}

// GetLastNotary is a free data retrieval call binding the contract method 0xda26501d.
//
// Solidity: function get_last_notary(uint256 id) constant returns(uint256)
func (_LitionScClient *LitionScClientCallerSession) GetLastNotary(id *big.Int) (*big.Int, error) {
	return _LitionScClient.Contract.GetLastNotary(&_LitionScClient.CallOpts, id)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x040c4c08.
//
// Solidity: function get_signature_hash_from_notary(uint32 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientCaller) GetSignatureHashFromNotary(opts *bind.CallOpts, notary_block uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "get_signature_hash_from_notary", notary_block, miners, blocks_mined, users, user_gas, largest_tx)
	return *ret0, err
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x040c4c08.
//
// Solidity: function get_signature_hash_from_notary(uint32 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientSession) GetSignatureHashFromNotary(notary_block uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	return _LitionScClient.Contract.GetSignatureHashFromNotary(&_LitionScClient.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x040c4c08.
//
// Solidity: function get_signature_hash_from_notary(uint32 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientCallerSession) GetSignatureHashFromNotary(notary_block uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	return _LitionScClient.Contract.GetSignatureHashFromNotary(&_LitionScClient.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetValidators is a free data retrieval call binding the contract method 0x1d9307a5.
//
// Solidity: function get_validators(uint256 id, uint256 batch) constant returns(address[100] users, uint256 count)
func (_LitionScClient *LitionScClientCaller) GetValidators(opts *bind.CallOpts, id *big.Int, batch *big.Int) (struct {
	Users [100]common.Address
	Count *big.Int
}, error) {
	ret := new(struct {
		Users [100]common.Address
		Count *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "get_validators", id, batch)
	return *ret, err
}

// GetValidators is a free data retrieval call binding the contract method 0x1d9307a5.
//
// Solidity: function get_validators(uint256 id, uint256 batch) constant returns(address[100] users, uint256 count)
func (_LitionScClient *LitionScClientSession) GetValidators(id *big.Int, batch *big.Int) (struct {
	Users [100]common.Address
	Count *big.Int
}, error) {
	return _LitionScClient.Contract.GetValidators(&_LitionScClient.CallOpts, id, batch)
}

// GetValidators is a free data retrieval call binding the contract method 0x1d9307a5.
//
// Solidity: function get_validators(uint256 id, uint256 batch) constant returns(address[100] users, uint256 count)
func (_LitionScClient *LitionScClientCallerSession) GetValidators(id *big.Int, batch *big.Int) (struct {
	Users [100]common.Address
	Count *big.Int
}, error) {
	return _LitionScClient.Contract.GetValidators(&_LitionScClient.CallOpts, id, batch)
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 id, address user) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) HasDeposited(opts *bind.CallOpts, id *big.Int, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "has_deposited", id, user)
	return *ret0, err
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 id, address user) constant returns(bool)
func (_LitionScClient *LitionScClientSession) HasDeposited(id *big.Int, user common.Address) (bool, error) {
	return _LitionScClient.Contract.HasDeposited(&_LitionScClient.CallOpts, id, user)
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 id, address user) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) HasDeposited(id *big.Int, user common.Address) (bool, error) {
	return _LitionScClient.Contract.HasDeposited(&_LitionScClient.CallOpts, id, user)
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 id, address user) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) HasVested(opts *bind.CallOpts, id *big.Int, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "has_vested", id, user)
	return *ret0, err
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 id, address user) constant returns(bool)
func (_LitionScClient *LitionScClientSession) HasVested(id *big.Int, user common.Address) (bool, error) {
	return _LitionScClient.Contract.HasVested(&_LitionScClient.CallOpts, id, user)
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 id, address user) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) HasVested(id *big.Int, user common.Address) (bool, error) {
	return _LitionScClient.Contract.HasVested(&_LitionScClient.CallOpts, id, user)
}

// NextId is a free data retrieval call binding the contract method 0xe31bfa00.
//
// Solidity: function next_id() constant returns(uint256)
func (_LitionScClient *LitionScClientCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "next_id")
	return *ret0, err
}

// NextId is a free data retrieval call binding the contract method 0xe31bfa00.
//
// Solidity: function next_id() constant returns(uint256)
func (_LitionScClient *LitionScClientSession) NextId() (*big.Int, error) {
	return _LitionScClient.Contract.NextId(&_LitionScClient.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0xe31bfa00.
//
// Solidity: function next_id() constant returns(uint256)
func (_LitionScClient *LitionScClientCallerSession) NextId() (*big.Int, error) {
	return _LitionScClient.Contract.NextId(&_LitionScClient.CallOpts)
}

// DepositInChain is a paid mutator transaction binding the contract method 0xd647820e.
//
// Solidity: function deposit_in_chain(uint256 id, uint256 deposit) returns()
func (_LitionScClient *LitionScClientTransactor) DepositInChain(opts *bind.TransactOpts, id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "deposit_in_chain", id, deposit)
}

// DepositInChain is a paid mutator transaction binding the contract method 0xd647820e.
//
// Solidity: function deposit_in_chain(uint256 id, uint256 deposit) returns()
func (_LitionScClient *LitionScClientSession) DepositInChain(id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.DepositInChain(&_LitionScClient.TransactOpts, id, deposit)
}

// DepositInChain is a paid mutator transaction binding the contract method 0xd647820e.
//
// Solidity: function deposit_in_chain(uint256 id, uint256 deposit) returns()
func (_LitionScClient *LitionScClientTransactorSession) DepositInChain(id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.DepositInChain(&_LitionScClient.TransactOpts, id, deposit)
}

// Notary is a paid mutator transaction binding the contract method 0xade5d22f.
//
// Solidity: function notary(uint256 id, uint32 notary_block_no, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactor) Notary(opts *bind.TransactOpts, id *big.Int, notary_block_no uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "notary", id, notary_block_no, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0xade5d22f.
//
// Solidity: function notary(uint256 id, uint32 notary_block_no, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientSession) Notary(id *big.Int, notary_block_no uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, id, notary_block_no, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0xade5d22f.
//
// Solidity: function notary(uint256 id, uint32 notary_block_no, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactorSession) Notary(id *big.Int, notary_block_no uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, id, notary_block_no, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb7f7968f.
//
// Solidity: function register_chain(string info, address validator, uint256 vesting, string init_endpoint) returns(uint256 id)
func (_LitionScClient *LitionScClientTransactor) RegisterChain(opts *bind.TransactOpts, info string, validator common.Address, vesting *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "register_chain", info, validator, vesting, init_endpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb7f7968f.
//
// Solidity: function register_chain(string info, address validator, uint256 vesting, string init_endpoint) returns(uint256 id)
func (_LitionScClient *LitionScClientSession) RegisterChain(info string, validator common.Address, vesting *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, info, validator, vesting, init_endpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb7f7968f.
//
// Solidity: function register_chain(string info, address validator, uint256 vesting, string init_endpoint) returns(uint256 id)
func (_LitionScClient *LitionScClientTransactorSession) RegisterChain(info string, validator common.Address, vesting *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, info, validator, vesting, init_endpoint)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 id) returns()
func (_LitionScClient *LitionScClientTransactor) StartMining(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "start_mining", id)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 id) returns()
func (_LitionScClient *LitionScClientSession) StartMining(id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StartMining(&_LitionScClient.TransactOpts, id)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 id) returns()
func (_LitionScClient *LitionScClientTransactorSession) StartMining(id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StartMining(&_LitionScClient.TransactOpts, id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 id) returns()
func (_LitionScClient *LitionScClientTransactor) StopMining(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "stop_mining", id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 id) returns()
func (_LitionScClient *LitionScClientSession) StopMining(id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StopMining(&_LitionScClient.TransactOpts, id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 id) returns()
func (_LitionScClient *LitionScClientTransactorSession) StopMining(id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StopMining(&_LitionScClient.TransactOpts, id)
}

// VestInChain is a paid mutator transaction binding the contract method 0x6d443d19.
//
// Solidity: function vest_in_chain(uint256 id, uint256 vesting) returns()
func (_LitionScClient *LitionScClientTransactor) VestInChain(opts *bind.TransactOpts, id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "vest_in_chain", id, vesting)
}

// VestInChain is a paid mutator transaction binding the contract method 0x6d443d19.
//
// Solidity: function vest_in_chain(uint256 id, uint256 vesting) returns()
func (_LitionScClient *LitionScClientSession) VestInChain(id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.VestInChain(&_LitionScClient.TransactOpts, id, vesting)
}

// VestInChain is a paid mutator transaction binding the contract method 0x6d443d19.
//
// Solidity: function vest_in_chain(uint256 id, uint256 vesting) returns()
func (_LitionScClient *LitionScClientTransactorSession) VestInChain(id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.VestInChain(&_LitionScClient.TransactOpts, id, vesting)
}

// LitionScClientDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LitionScClient contract.
type LitionScClientDepositIterator struct {
	Event *LitionScClientDeposit // Event containing the contract specifics and raw log

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
func (it *LitionScClientDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientDeposit)
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
		it.Event = new(LitionScClientDeposit)
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
func (it *LitionScClientDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientDeposit represents a Deposit event raised by the LitionScClient contract.
type LitionScClientDeposit struct {
	ChainId   *big.Int
	Deposit   *big.Int
	Depositer common.Address
	Datetime  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x9a2a1e97e6d641080089aafc36750cfdef4c79f8b3ace6fa4c384fa2f0476959.
//
// Solidity: event Deposit(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_LitionScClient *LitionScClientFilterer) FilterDeposit(opts *bind.FilterOpts, chain_id []*big.Int, depositer []common.Address) (*LitionScClientDepositIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "Deposit", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientDepositIterator{contract: _LitionScClient.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x9a2a1e97e6d641080089aafc36750cfdef4c79f8b3ace6fa4c384fa2f0476959.
//
// Solidity: event Deposit(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_LitionScClient *LitionScClientFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LitionScClientDeposit, chain_id []*big.Int, depositer []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "Deposit", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientDeposit)
				if err := _LitionScClient.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x9a2a1e97e6d641080089aafc36750cfdef4c79f8b3ace6fa4c384fa2f0476959.
//
// Solidity: event Deposit(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_LitionScClient *LitionScClientFilterer) ParseDeposit(log types.Log) (*LitionScClientDeposit, error) {
	event := new(LitionScClientDeposit)
	if err := _LitionScClient.contract.UnpackLog(event, "Deposit", log); err != nil {
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
	Id          *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewChain is a free log retrieval operation binding the contract event 0xd8a08a6069bc34d54779d98bc9eaf69c8398a605b914243d592ae8e9bd94a728.
//
// Solidity: event NewChain(uint256 id, string description)
func (_LitionScClient *LitionScClientFilterer) FilterNewChain(opts *bind.FilterOpts) (*LitionScClientNewChainIterator, error) {

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "NewChain")
	if err != nil {
		return nil, err
	}
	return &LitionScClientNewChainIterator{contract: _LitionScClient.contract, event: "NewChain", logs: logs, sub: sub}, nil
}

// WatchNewChain is a free log subscription operation binding the contract event 0xd8a08a6069bc34d54779d98bc9eaf69c8398a605b914243d592ae8e9bd94a728.
//
// Solidity: event NewChain(uint256 id, string description)
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

// ParseNewChain is a log parse operation binding the contract event 0xd8a08a6069bc34d54779d98bc9eaf69c8398a605b914243d592ae8e9bd94a728.
//
// Solidity: event NewChain(uint256 id, string description)
func (_LitionScClient *LitionScClientFilterer) ParseNewChain(log types.Log) (*LitionScClientNewChain, error) {
	event := new(LitionScClientNewChain)
	if err := _LitionScClient.contract.UnpackLog(event, "NewChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientNewChainEndpointIterator is returned from FilterNewChainEndpoint and is used to iterate over the raw logs and unpacked data for NewChainEndpoint events raised by the LitionScClient contract.
type LitionScClientNewChainEndpointIterator struct {
	Event *LitionScClientNewChainEndpoint // Event containing the contract specifics and raw log

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
func (it *LitionScClientNewChainEndpointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientNewChainEndpoint)
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
		it.Event = new(LitionScClientNewChainEndpoint)
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
func (it *LitionScClientNewChainEndpointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientNewChainEndpointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientNewChainEndpoint represents a NewChainEndpoint event raised by the LitionScClient contract.
type LitionScClientNewChainEndpoint struct {
	Id       *big.Int
	Endpoint string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewChainEndpoint is a free log retrieval operation binding the contract event 0x0e51cb8a86e5362fecae09c9dd8b8f25ad4e35cc650cdc3fe4ec2da11bcd5004.
//
// Solidity: event NewChainEndpoint(uint256 id, string endpoint)
func (_LitionScClient *LitionScClientFilterer) FilterNewChainEndpoint(opts *bind.FilterOpts) (*LitionScClientNewChainEndpointIterator, error) {

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "NewChainEndpoint")
	if err != nil {
		return nil, err
	}
	return &LitionScClientNewChainEndpointIterator{contract: _LitionScClient.contract, event: "NewChainEndpoint", logs: logs, sub: sub}, nil
}

// WatchNewChainEndpoint is a free log subscription operation binding the contract event 0x0e51cb8a86e5362fecae09c9dd8b8f25ad4e35cc650cdc3fe4ec2da11bcd5004.
//
// Solidity: event NewChainEndpoint(uint256 id, string endpoint)
func (_LitionScClient *LitionScClientFilterer) WatchNewChainEndpoint(opts *bind.WatchOpts, sink chan<- *LitionScClientNewChainEndpoint) (event.Subscription, error) {

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "NewChainEndpoint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientNewChainEndpoint)
				if err := _LitionScClient.contract.UnpackLog(event, "NewChainEndpoint", log); err != nil {
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

// ParseNewChainEndpoint is a log parse operation binding the contract event 0x0e51cb8a86e5362fecae09c9dd8b8f25ad4e35cc650cdc3fe4ec2da11bcd5004.
//
// Solidity: event NewChainEndpoint(uint256 id, string endpoint)
func (_LitionScClient *LitionScClientFilterer) ParseNewChainEndpoint(log types.Log) (*LitionScClientNewChainEndpoint, error) {
	event := new(LitionScClientNewChainEndpoint)
	if err := _LitionScClient.contract.UnpackLog(event, "NewChainEndpoint", log); err != nil {
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
	Miner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStartMining is a free log retrieval operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chain_id, address miner)
func (_LitionScClient *LitionScClientFilterer) FilterStartMining(opts *bind.FilterOpts, chain_id []*big.Int) (*LitionScClientStartMiningIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "StartMining", chain_idRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientStartMiningIterator{contract: _LitionScClient.contract, event: "StartMining", logs: logs, sub: sub}, nil
}

// WatchStartMining is a free log subscription operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chain_id, address miner)
func (_LitionScClient *LitionScClientFilterer) WatchStartMining(opts *bind.WatchOpts, sink chan<- *LitionScClientStartMining, chain_id []*big.Int) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "StartMining", chain_idRule)
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
// Solidity: event StartMining(uint256 indexed chain_id, address miner)
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
	Miner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStopMining is a free log retrieval operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chain_id, address miner)
func (_LitionScClient *LitionScClientFilterer) FilterStopMining(opts *bind.FilterOpts, chain_id []*big.Int) (*LitionScClientStopMiningIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "StopMining", chain_idRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientStopMiningIterator{contract: _LitionScClient.contract, event: "StopMining", logs: logs, sub: sub}, nil
}

// WatchStopMining is a free log subscription operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chain_id, address miner)
func (_LitionScClient *LitionScClientFilterer) WatchStopMining(opts *bind.WatchOpts, sink chan<- *LitionScClientStopMining, chain_id []*big.Int) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "StopMining", chain_idRule)
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
// Solidity: event StopMining(uint256 indexed chain_id, address miner)
func (_LitionScClient *LitionScClientFilterer) ParseStopMining(log types.Log) (*LitionScClientStopMining, error) {
	event := new(LitionScClientStopMining)
	if err := _LitionScClient.contract.UnpackLog(event, "StopMining", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientVestingIterator is returned from FilterVesting and is used to iterate over the raw logs and unpacked data for Vesting events raised by the LitionScClient contract.
type LitionScClientVestingIterator struct {
	Event *LitionScClientVesting // Event containing the contract specifics and raw log

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
func (it *LitionScClientVestingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientVesting)
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
		it.Event = new(LitionScClientVesting)
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
func (it *LitionScClientVestingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientVestingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientVesting represents a Vesting event raised by the LitionScClient contract.
type LitionScClientVesting struct {
	ChainId   *big.Int
	Deposit   *big.Int
	Depositer common.Address
	Datetime  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVesting is a free log retrieval operation binding the contract event 0x538c8cde66e8496c8eccec720ba64db094e4bd703ad2b39ab1b76fff5b799854.
//
// Solidity: event Vesting(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_LitionScClient *LitionScClientFilterer) FilterVesting(opts *bind.FilterOpts, chain_id []*big.Int, depositer []common.Address) (*LitionScClientVestingIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "Vesting", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientVestingIterator{contract: _LitionScClient.contract, event: "Vesting", logs: logs, sub: sub}, nil
}

// WatchVesting is a free log subscription operation binding the contract event 0x538c8cde66e8496c8eccec720ba64db094e4bd703ad2b39ab1b76fff5b799854.
//
// Solidity: event Vesting(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_LitionScClient *LitionScClientFilterer) WatchVesting(opts *bind.WatchOpts, sink chan<- *LitionScClientVesting, chain_id []*big.Int, depositer []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "Vesting", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientVesting)
				if err := _LitionScClient.contract.UnpackLog(event, "Vesting", log); err != nil {
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

// ParseVesting is a log parse operation binding the contract event 0x538c8cde66e8496c8eccec720ba64db094e4bd703ad2b39ab1b76fff5b799854.
//
// Solidity: event Vesting(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_LitionScClient *LitionScClientFilterer) ParseVesting(log types.Log) (*LitionScClientVesting, error) {
	event := new(LitionScClientVesting)
	if err := _LitionScClient.contract.UnpackLog(event, "Vesting", log); err != nil {
		return nil, err
	}
	return event, nil
}
