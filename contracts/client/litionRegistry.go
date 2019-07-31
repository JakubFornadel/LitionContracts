// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package litionContractClient

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LitionABI is the input ABI used to generate the binding from.
const LitionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"start_mining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"notary_block\",\"type\":\"uint32\"},{\"name\":\"miners\",\"type\":\"address[]\"},{\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"name\":\"users\",\"type\":\"address[]\"},{\"name\":\"user_gas\",\"type\":\"uint32[]\"},{\"name\":\"largest_tx\",\"type\":\"uint32\"},{\"name\":\"notary_data\",\"type\":\"bytes\"}],\"name\":\"notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chains\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"},{\"name\":\"last_notary\",\"type\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"total_vesting\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"stop_mining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"vesting\",\"type\":\"uint256\"}],\"name\":\"vest_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"has_deposited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"info\",\"type\":\"string\"},{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"vesting\",\"type\":\"uint256\"},{\"name\":\"init_endpoint\",\"type\":\"string\"}],\"name\":\"register_chain\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"has_vested\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"deposit_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"NewChainEndpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"datetime\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"datetime\",\"type\":\"uint256\"}],\"name\":\"Vesting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"StartMining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"StopMining\",\"type\":\"event\"}]"

// Lition is an auto generated Go binding around an Ethereum contract.
type Lition struct {
	LitionCaller     // Read-only binding to the contract
	LitionTransactor // Write-only binding to the contract
	LitionFilterer   // Log filterer for contract events
}

// LitionCaller is an auto generated read-only Go binding around an Ethereum contract.
type LitionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LitionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LitionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LitionSession struct {
	Contract     *Lition           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LitionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LitionCallerSession struct {
	Contract *LitionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LitionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LitionTransactorSession struct {
	Contract     *LitionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LitionRaw is an auto generated low-level Go binding around an Ethereum contract.
type LitionRaw struct {
	Contract *Lition // Generic contract binding to access the raw methods on
}

// LitionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LitionCallerRaw struct {
	Contract *LitionCaller // Generic read-only contract binding to access the raw methods on
}

// LitionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LitionTransactorRaw struct {
	Contract *LitionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLition creates a new instance of Lition, bound to a specific deployed contract.
func NewLition(address common.Address, backend bind.ContractBackend) (*Lition, error) {
	contract, err := bindLition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lition{LitionCaller: LitionCaller{contract: contract}, LitionTransactor: LitionTransactor{contract: contract}, LitionFilterer: LitionFilterer{contract: contract}}, nil
}

// NewLitionCaller creates a new read-only instance of Lition, bound to a specific deployed contract.
func NewLitionCaller(address common.Address, caller bind.ContractCaller) (*LitionCaller, error) {
	contract, err := bindLition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LitionCaller{contract: contract}, nil
}

// NewLitionTransactor creates a new write-only instance of Lition, bound to a specific deployed contract.
func NewLitionTransactor(address common.Address, transactor bind.ContractTransactor) (*LitionTransactor, error) {
	contract, err := bindLition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LitionTransactor{contract: contract}, nil
}

// NewLitionFilterer creates a new log filterer instance of Lition, bound to a specific deployed contract.
func NewLitionFilterer(address common.Address, filterer bind.ContractFilterer) (*LitionFilterer, error) {
	contract, err := bindLition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LitionFilterer{contract: contract}, nil
}

// bindLition binds a generic wrapper to an already deployed contract.
func bindLition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LitionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lition *LitionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lition.Contract.LitionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lition *LitionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lition.Contract.LitionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lition *LitionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lition.Contract.LitionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lition *LitionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lition *LitionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lition *LitionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lition.Contract.contract.Transact(opts, method, params...)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) constant returns(bool active, uint256 last_notary, address validator, uint256 total_vesting)
func (_Lition *LitionCaller) Chains(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Lition.contract.Call(opts, out, "chains", arg0)
	return *ret, err
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) constant returns(bool active, uint256 last_notary, address validator, uint256 total_vesting)
func (_Lition *LitionSession) Chains(arg0 *big.Int) (struct {
	Active       bool
	LastNotary   *big.Int
	Validator    common.Address
	TotalVesting *big.Int
}, error) {
	return _Lition.Contract.Chains(&_Lition.CallOpts, arg0)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) constant returns(bool active, uint256 last_notary, address validator, uint256 total_vesting)
func (_Lition *LitionCallerSession) Chains(arg0 *big.Int) (struct {
	Active       bool
	LastNotary   *big.Int
	Validator    common.Address
	TotalVesting *big.Int
}, error) {
	return _Lition.Contract.Chains(&_Lition.CallOpts, arg0)
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 id, address user) constant returns(bool)
func (_Lition *LitionCaller) HasDeposited(opts *bind.CallOpts, id *big.Int, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lition.contract.Call(opts, out, "has_deposited", id, user)
	return *ret0, err
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 id, address user) constant returns(bool)
func (_Lition *LitionSession) HasDeposited(id *big.Int, user common.Address) (bool, error) {
	return _Lition.Contract.HasDeposited(&_Lition.CallOpts, id, user)
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 id, address user) constant returns(bool)
func (_Lition *LitionCallerSession) HasDeposited(id *big.Int, user common.Address) (bool, error) {
	return _Lition.Contract.HasDeposited(&_Lition.CallOpts, id, user)
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 id, address user) constant returns(bool)
func (_Lition *LitionCaller) HasVested(opts *bind.CallOpts, id *big.Int, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lition.contract.Call(opts, out, "has_vested", id, user)
	return *ret0, err
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 id, address user) constant returns(bool)
func (_Lition *LitionSession) HasVested(id *big.Int, user common.Address) (bool, error) {
	return _Lition.Contract.HasVested(&_Lition.CallOpts, id, user)
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 id, address user) constant returns(bool)
func (_Lition *LitionCallerSession) HasVested(id *big.Int, user common.Address) (bool, error) {
	return _Lition.Contract.HasVested(&_Lition.CallOpts, id, user)
}

// NextId is a free data retrieval call binding the contract method 0xe31bfa00.
//
// Solidity: function next_id() constant returns(uint256)
func (_Lition *LitionCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lition.contract.Call(opts, out, "next_id")
	return *ret0, err
}

// NextId is a free data retrieval call binding the contract method 0xe31bfa00.
//
// Solidity: function next_id() constant returns(uint256)
func (_Lition *LitionSession) NextId() (*big.Int, error) {
	return _Lition.Contract.NextId(&_Lition.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0xe31bfa00.
//
// Solidity: function next_id() constant returns(uint256)
func (_Lition *LitionCallerSession) NextId() (*big.Int, error) {
	return _Lition.Contract.NextId(&_Lition.CallOpts)
}

// DepositInChain is a paid mutator transaction binding the contract method 0xd647820e.
//
// Solidity: function deposit_in_chain(uint256 id, uint256 deposit) returns()
func (_Lition *LitionTransactor) DepositInChain(opts *bind.TransactOpts, id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "deposit_in_chain", id, deposit)
}

// DepositInChain is a paid mutator transaction binding the contract method 0xd647820e.
//
// Solidity: function deposit_in_chain(uint256 id, uint256 deposit) returns()
func (_Lition *LitionSession) DepositInChain(id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.DepositInChain(&_Lition.TransactOpts, id, deposit)
}

// DepositInChain is a paid mutator transaction binding the contract method 0xd647820e.
//
// Solidity: function deposit_in_chain(uint256 id, uint256 deposit) returns()
func (_Lition *LitionTransactorSession) DepositInChain(id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.DepositInChain(&_Lition.TransactOpts, id, deposit)
}

// Notary is a paid mutator transaction binding the contract method 0x3a68156f.
//
// Solidity: function notary(uint256 id, uint32 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, bytes notary_data) returns()
func (_Lition *LitionTransactor) Notary(opts *bind.TransactOpts, id *big.Int, notary_block uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, notary_data []byte) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "notary", id, notary_block, miners, blocks_mined, users, user_gas, largest_tx, notary_data)
}

// Notary is a paid mutator transaction binding the contract method 0x3a68156f.
//
// Solidity: function notary(uint256 id, uint32 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, bytes notary_data) returns()
func (_Lition *LitionSession) Notary(id *big.Int, notary_block uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, notary_data []byte) (*types.Transaction, error) {
	return _Lition.Contract.Notary(&_Lition.TransactOpts, id, notary_block, miners, blocks_mined, users, user_gas, largest_tx, notary_data)
}

// Notary is a paid mutator transaction binding the contract method 0x3a68156f.
//
// Solidity: function notary(uint256 id, uint32 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, bytes notary_data) returns()
func (_Lition *LitionTransactorSession) Notary(id *big.Int, notary_block uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, notary_data []byte) (*types.Transaction, error) {
	return _Lition.Contract.Notary(&_Lition.TransactOpts, id, notary_block, miners, blocks_mined, users, user_gas, largest_tx, notary_data)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb7f7968f.
//
// Solidity: function register_chain(string info, address validator, uint256 vesting, string init_endpoint) returns(uint256 id)
func (_Lition *LitionTransactor) RegisterChain(opts *bind.TransactOpts, info string, validator common.Address, vesting *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "register_chain", info, validator, vesting, init_endpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb7f7968f.
//
// Solidity: function register_chain(string info, address validator, uint256 vesting, string init_endpoint) returns(uint256 id)
func (_Lition *LitionSession) RegisterChain(info string, validator common.Address, vesting *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _Lition.Contract.RegisterChain(&_Lition.TransactOpts, info, validator, vesting, init_endpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb7f7968f.
//
// Solidity: function register_chain(string info, address validator, uint256 vesting, string init_endpoint) returns(uint256 id)
func (_Lition *LitionTransactorSession) RegisterChain(info string, validator common.Address, vesting *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _Lition.Contract.RegisterChain(&_Lition.TransactOpts, info, validator, vesting, init_endpoint)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 id) returns()
func (_Lition *LitionTransactor) StartMining(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "start_mining", id)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 id) returns()
func (_Lition *LitionSession) StartMining(id *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.StartMining(&_Lition.TransactOpts, id)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 id) returns()
func (_Lition *LitionTransactorSession) StartMining(id *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.StartMining(&_Lition.TransactOpts, id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 id) returns()
func (_Lition *LitionTransactor) StopMining(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "stop_mining", id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 id) returns()
func (_Lition *LitionSession) StopMining(id *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.StopMining(&_Lition.TransactOpts, id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 id) returns()
func (_Lition *LitionTransactorSession) StopMining(id *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.StopMining(&_Lition.TransactOpts, id)
}

// VestInChain is a paid mutator transaction binding the contract method 0x6d443d19.
//
// Solidity: function vest_in_chain(uint256 id, uint256 vesting) returns()
func (_Lition *LitionTransactor) VestInChain(opts *bind.TransactOpts, id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "vest_in_chain", id, vesting)
}

// VestInChain is a paid mutator transaction binding the contract method 0x6d443d19.
//
// Solidity: function vest_in_chain(uint256 id, uint256 vesting) returns()
func (_Lition *LitionSession) VestInChain(id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.VestInChain(&_Lition.TransactOpts, id, vesting)
}

// VestInChain is a paid mutator transaction binding the contract method 0x6d443d19.
//
// Solidity: function vest_in_chain(uint256 id, uint256 vesting) returns()
func (_Lition *LitionTransactorSession) VestInChain(id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _Lition.Contract.VestInChain(&_Lition.TransactOpts, id, vesting)
}

// LitionDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Lition contract.
type LitionDepositIterator struct {
	Event *LitionDeposit // Event containing the contract specifics and raw log

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
func (it *LitionDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionDeposit)
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
		it.Event = new(LitionDeposit)
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
func (it *LitionDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionDeposit represents a Deposit event raised by the Lition contract.
type LitionDeposit struct {
	ChainId   *big.Int
	Deposit   *big.Int
	Depositer common.Address
	Datetime  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x9a2a1e97e6d641080089aafc36750cfdef4c79f8b3ace6fa4c384fa2f0476959.
//
// Solidity: event Deposit(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_Lition *LitionFilterer) FilterDeposit(opts *bind.FilterOpts, chain_id []*big.Int, depositer []common.Address) (*LitionDepositIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _Lition.contract.FilterLogs(opts, "Deposit", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return &LitionDepositIterator{contract: _Lition.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x9a2a1e97e6d641080089aafc36750cfdef4c79f8b3ace6fa4c384fa2f0476959.
//
// Solidity: event Deposit(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_Lition *LitionFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LitionDeposit, chain_id []*big.Int, depositer []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _Lition.contract.WatchLogs(opts, "Deposit", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionDeposit)
				if err := _Lition.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Lition *LitionFilterer) ParseDeposit(log types.Log) (*LitionDeposit, error) {
	event := new(LitionDeposit)
	if err := _Lition.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionNewChainIterator is returned from FilterNewChain and is used to iterate over the raw logs and unpacked data for NewChain events raised by the Lition contract.
type LitionNewChainIterator struct {
	Event *LitionNewChain // Event containing the contract specifics and raw log

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
func (it *LitionNewChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionNewChain)
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
		it.Event = new(LitionNewChain)
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
func (it *LitionNewChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionNewChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionNewChain represents a NewChain event raised by the Lition contract.
type LitionNewChain struct {
	Id          *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewChain is a free log retrieval operation binding the contract event 0xd8a08a6069bc34d54779d98bc9eaf69c8398a605b914243d592ae8e9bd94a728.
//
// Solidity: event NewChain(uint256 id, string description)
func (_Lition *LitionFilterer) FilterNewChain(opts *bind.FilterOpts) (*LitionNewChainIterator, error) {

	logs, sub, err := _Lition.contract.FilterLogs(opts, "NewChain")
	if err != nil {
		return nil, err
	}
	return &LitionNewChainIterator{contract: _Lition.contract, event: "NewChain", logs: logs, sub: sub}, nil
}

// WatchNewChain is a free log subscription operation binding the contract event 0xd8a08a6069bc34d54779d98bc9eaf69c8398a605b914243d592ae8e9bd94a728.
//
// Solidity: event NewChain(uint256 id, string description)
func (_Lition *LitionFilterer) WatchNewChain(opts *bind.WatchOpts, sink chan<- *LitionNewChain) (event.Subscription, error) {

	logs, sub, err := _Lition.contract.WatchLogs(opts, "NewChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionNewChain)
				if err := _Lition.contract.UnpackLog(event, "NewChain", log); err != nil {
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
func (_Lition *LitionFilterer) ParseNewChain(log types.Log) (*LitionNewChain, error) {
	event := new(LitionNewChain)
	if err := _Lition.contract.UnpackLog(event, "NewChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionNewChainEndpointIterator is returned from FilterNewChainEndpoint and is used to iterate over the raw logs and unpacked data for NewChainEndpoint events raised by the Lition contract.
type LitionNewChainEndpointIterator struct {
	Event *LitionNewChainEndpoint // Event containing the contract specifics and raw log

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
func (it *LitionNewChainEndpointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionNewChainEndpoint)
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
		it.Event = new(LitionNewChainEndpoint)
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
func (it *LitionNewChainEndpointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionNewChainEndpointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionNewChainEndpoint represents a NewChainEndpoint event raised by the Lition contract.
type LitionNewChainEndpoint struct {
	Id       *big.Int
	Endpoint string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewChainEndpoint is a free log retrieval operation binding the contract event 0x0e51cb8a86e5362fecae09c9dd8b8f25ad4e35cc650cdc3fe4ec2da11bcd5004.
//
// Solidity: event NewChainEndpoint(uint256 id, string endpoint)
func (_Lition *LitionFilterer) FilterNewChainEndpoint(opts *bind.FilterOpts) (*LitionNewChainEndpointIterator, error) {

	logs, sub, err := _Lition.contract.FilterLogs(opts, "NewChainEndpoint")
	if err != nil {
		return nil, err
	}
	return &LitionNewChainEndpointIterator{contract: _Lition.contract, event: "NewChainEndpoint", logs: logs, sub: sub}, nil
}

// WatchNewChainEndpoint is a free log subscription operation binding the contract event 0x0e51cb8a86e5362fecae09c9dd8b8f25ad4e35cc650cdc3fe4ec2da11bcd5004.
//
// Solidity: event NewChainEndpoint(uint256 id, string endpoint)
func (_Lition *LitionFilterer) WatchNewChainEndpoint(opts *bind.WatchOpts, sink chan<- *LitionNewChainEndpoint) (event.Subscription, error) {

	logs, sub, err := _Lition.contract.WatchLogs(opts, "NewChainEndpoint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionNewChainEndpoint)
				if err := _Lition.contract.UnpackLog(event, "NewChainEndpoint", log); err != nil {
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
func (_Lition *LitionFilterer) ParseNewChainEndpoint(log types.Log) (*LitionNewChainEndpoint, error) {
	event := new(LitionNewChainEndpoint)
	if err := _Lition.contract.UnpackLog(event, "NewChainEndpoint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionStartMiningIterator is returned from FilterStartMining and is used to iterate over the raw logs and unpacked data for StartMining events raised by the Lition contract.
type LitionStartMiningIterator struct {
	Event *LitionStartMining // Event containing the contract specifics and raw log

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
func (it *LitionStartMiningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionStartMining)
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
		it.Event = new(LitionStartMining)
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
func (it *LitionStartMiningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionStartMiningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionStartMining represents a StartMining event raised by the Lition contract.
type LitionStartMining struct {
	ChainId *big.Int
	Miner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStartMining is a free log retrieval operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chain_id, address miner)
func (_Lition *LitionFilterer) FilterStartMining(opts *bind.FilterOpts, chain_id []*big.Int) (*LitionStartMiningIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _Lition.contract.FilterLogs(opts, "StartMining", chain_idRule)
	if err != nil {
		return nil, err
	}
	return &LitionStartMiningIterator{contract: _Lition.contract, event: "StartMining", logs: logs, sub: sub}, nil
}

// WatchStartMining is a free log subscription operation binding the contract event 0x1090e92433a132c15edc16a996682566af9d40d581e34d732ac2b39991847892.
//
// Solidity: event StartMining(uint256 indexed chain_id, address miner)
func (_Lition *LitionFilterer) WatchStartMining(opts *bind.WatchOpts, sink chan<- *LitionStartMining, chain_id []*big.Int) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _Lition.contract.WatchLogs(opts, "StartMining", chain_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionStartMining)
				if err := _Lition.contract.UnpackLog(event, "StartMining", log); err != nil {
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
func (_Lition *LitionFilterer) ParseStartMining(log types.Log) (*LitionStartMining, error) {
	event := new(LitionStartMining)
	if err := _Lition.contract.UnpackLog(event, "StartMining", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionStopMiningIterator is returned from FilterStopMining and is used to iterate over the raw logs and unpacked data for StopMining events raised by the Lition contract.
type LitionStopMiningIterator struct {
	Event *LitionStopMining // Event containing the contract specifics and raw log

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
func (it *LitionStopMiningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionStopMining)
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
		it.Event = new(LitionStopMining)
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
func (it *LitionStopMiningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionStopMiningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionStopMining represents a StopMining event raised by the Lition contract.
type LitionStopMining struct {
	ChainId *big.Int
	Miner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStopMining is a free log retrieval operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chain_id, address miner)
func (_Lition *LitionFilterer) FilterStopMining(opts *bind.FilterOpts, chain_id []*big.Int) (*LitionStopMiningIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _Lition.contract.FilterLogs(opts, "StopMining", chain_idRule)
	if err != nil {
		return nil, err
	}
	return &LitionStopMiningIterator{contract: _Lition.contract, event: "StopMining", logs: logs, sub: sub}, nil
}

// WatchStopMining is a free log subscription operation binding the contract event 0x921c933fe5d237f13ecce36e8ce6e7370d68826ed08698f6d2dd81caf298aaa3.
//
// Solidity: event StopMining(uint256 indexed chain_id, address miner)
func (_Lition *LitionFilterer) WatchStopMining(opts *bind.WatchOpts, sink chan<- *LitionStopMining, chain_id []*big.Int) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _Lition.contract.WatchLogs(opts, "StopMining", chain_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionStopMining)
				if err := _Lition.contract.UnpackLog(event, "StopMining", log); err != nil {
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
func (_Lition *LitionFilterer) ParseStopMining(log types.Log) (*LitionStopMining, error) {
	event := new(LitionStopMining)
	if err := _Lition.contract.UnpackLog(event, "StopMining", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionVestingIterator is returned from FilterVesting and is used to iterate over the raw logs and unpacked data for Vesting events raised by the Lition contract.
type LitionVestingIterator struct {
	Event *LitionVesting // Event containing the contract specifics and raw log

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
func (it *LitionVestingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionVesting)
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
		it.Event = new(LitionVesting)
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
func (it *LitionVestingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionVestingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionVesting represents a Vesting event raised by the Lition contract.
type LitionVesting struct {
	ChainId   *big.Int
	Deposit   *big.Int
	Depositer common.Address
	Datetime  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVesting is a free log retrieval operation binding the contract event 0x538c8cde66e8496c8eccec720ba64db094e4bd703ad2b39ab1b76fff5b799854.
//
// Solidity: event Vesting(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_Lition *LitionFilterer) FilterVesting(opts *bind.FilterOpts, chain_id []*big.Int, depositer []common.Address) (*LitionVestingIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _Lition.contract.FilterLogs(opts, "Vesting", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return &LitionVestingIterator{contract: _Lition.contract, event: "Vesting", logs: logs, sub: sub}, nil
}

// WatchVesting is a free log subscription operation binding the contract event 0x538c8cde66e8496c8eccec720ba64db094e4bd703ad2b39ab1b76fff5b799854.
//
// Solidity: event Vesting(uint256 indexed chain_id, uint256 deposit, address indexed depositer, uint256 datetime)
func (_Lition *LitionFilterer) WatchVesting(opts *bind.WatchOpts, sink chan<- *LitionVesting, chain_id []*big.Int, depositer []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}

	logs, sub, err := _Lition.contract.WatchLogs(opts, "Vesting", chain_idRule, depositerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionVesting)
				if err := _Lition.contract.UnpackLog(event, "Vesting", log); err != nil {
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
func (_Lition *LitionFilterer) ParseVesting(log types.Log) (*LitionVesting, error) {
	event := new(LitionVesting)
	if err := _Lition.contract.UnpackLog(event, "Vesting", log); err != nil {
		return nil, err
	}
	return event, nil
}
