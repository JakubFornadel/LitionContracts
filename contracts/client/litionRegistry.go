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
const LitionScClientABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"cancel_vest_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"get_active_validators\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"start_mining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"notary_block\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"user_gas\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"largest_tx\",\"type\":\"uint32\"}],\"name\":\"get_signature_hash_from_notary\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"request_deposit_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"cancel_deposit_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"confirm_deposit_withdrawal_from_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"stop_mining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"get_user_details\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notary_start_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notary_end_block\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"user_gas\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"largest_tx\",\"type\":\"uint32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"contractChainValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"vesting\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"deposit\",\"type\":\"uint96\"},{\"internalType\":\"string\",\"name\":\"init_endpoint\",\"type\":\"string\"}],\"name\":\"register_chain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"notary_block_no\",\"type\":\"uint256\"}],\"name\":\"test_notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"get_allowed_to_transact\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"has_deposited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"get_user_requests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"vesting_req_exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vesting_req_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting_req_notary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting_req_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting_req_state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting_req_control_state\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deposit_req_exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposit_req_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit_req_notary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit_req_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit_req_state\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"has_vested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"}],\"name\":\"request_vest_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"get_last_notary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"last_notary_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_notary_timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"get_chain_details\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"total_vesting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_notary_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_notary_timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"}],\"name\":\"get_allowed_to_validate\",\"outputs\":[{\"internalType\":\"address[100]\",\"name\":\"\",\"type\":\"address[100]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"}],\"name\":\"confirm_vest_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"RequestDepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"ConfirmDepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"CancelDepositInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ForceWithdrawDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"RequestVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"ConfirmVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"CancelVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vesting\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"req_timestamp\",\"type\":\"uint256\"}],\"name\":\"AcceptedVestInChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ForceWithdrawVesting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"WhitelistAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"StartMining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"StopMining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"new_contract\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"}]"

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

// GetActiveValidators is a free data retrieval call binding the contract method 0x1ff7c270.
//
// Solidity: function get_active_validators(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCaller) GetActiveValidators(opts *bind.CallOpts, chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
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
	err := _LitionScClient.contract.Call(opts, out, "get_active_validators", chain_id, batch)
	return *ret0, *ret1, *ret2, err
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x1ff7c270.
//
// Solidity: function get_active_validators(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientSession) GetActiveValidators(chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetActiveValidators(&_LitionScClient.CallOpts, chain_id, batch)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x1ff7c270.
//
// Solidity: function get_active_validators(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCallerSession) GetActiveValidators(chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetActiveValidators(&_LitionScClient.CallOpts, chain_id, batch)
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0xb190042c.
//
// Solidity: function get_allowed_to_transact(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCaller) GetAllowedToTransact(opts *bind.CallOpts, chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
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
	err := _LitionScClient.contract.Call(opts, out, "get_allowed_to_transact", chain_id, batch)
	return *ret0, *ret1, *ret2, err
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0xb190042c.
//
// Solidity: function get_allowed_to_transact(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientSession) GetAllowedToTransact(chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToTransact(&_LitionScClient.CallOpts, chain_id, batch)
}

// GetAllowedToTransact is a free data retrieval call binding the contract method 0xb190042c.
//
// Solidity: function get_allowed_to_transact(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCallerSession) GetAllowedToTransact(chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToTransact(&_LitionScClient.CallOpts, chain_id, batch)
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0xe34c31b1.
//
// Solidity: function get_allowed_to_validate(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCaller) GetAllowedToValidate(opts *bind.CallOpts, chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
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
	err := _LitionScClient.contract.Call(opts, out, "get_allowed_to_validate", chain_id, batch)
	return *ret0, *ret1, *ret2, err
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0xe34c31b1.
//
// Solidity: function get_allowed_to_validate(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientSession) GetAllowedToValidate(chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToValidate(&_LitionScClient.CallOpts, chain_id, batch)
}

// GetAllowedToValidate is a free data retrieval call binding the contract method 0xe34c31b1.
//
// Solidity: function get_allowed_to_validate(uint256 chain_id, uint256 batch) constant returns(address[100], uint256, bool)
func (_LitionScClient *LitionScClientCallerSession) GetAllowedToValidate(chain_id *big.Int, batch *big.Int) ([100]common.Address, *big.Int, bool, error) {
	return _LitionScClient.Contract.GetAllowedToValidate(&_LitionScClient.CallOpts, chain_id, batch)
}

// GetChainDetails is a free data retrieval call binding the contract method 0xddaad348.
//
// Solidity: function get_chain_details(uint256 chain_id) constant returns(bool registered, bool active, string endpoint, uint256 total_vesting, uint256 last_notary_block, uint256 last_notary_timestamp)
func (_LitionScClient *LitionScClientCaller) GetChainDetails(opts *bind.CallOpts, chain_id *big.Int) (struct {
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
	err := _LitionScClient.contract.Call(opts, out, "get_chain_details", chain_id)
	return *ret, err
}

// GetChainDetails is a free data retrieval call binding the contract method 0xddaad348.
//
// Solidity: function get_chain_details(uint256 chain_id) constant returns(bool registered, bool active, string endpoint, uint256 total_vesting, uint256 last_notary_block, uint256 last_notary_timestamp)
func (_LitionScClient *LitionScClientSession) GetChainDetails(chain_id *big.Int) (struct {
	Registered          bool
	Active              bool
	Endpoint            string
	TotalVesting        *big.Int
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetChainDetails(&_LitionScClient.CallOpts, chain_id)
}

// GetChainDetails is a free data retrieval call binding the contract method 0xddaad348.
//
// Solidity: function get_chain_details(uint256 chain_id) constant returns(bool registered, bool active, string endpoint, uint256 total_vesting, uint256 last_notary_block, uint256 last_notary_timestamp)
func (_LitionScClient *LitionScClientCallerSession) GetChainDetails(chain_id *big.Int) (struct {
	Registered          bool
	Active              bool
	Endpoint            string
	TotalVesting        *big.Int
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetChainDetails(&_LitionScClient.CallOpts, chain_id)
}

// GetLastNotary is a free data retrieval call binding the contract method 0xda26501d.
//
// Solidity: function get_last_notary(uint256 chain_id) constant returns(uint256 last_notary_block, uint256 last_notary_timestamp)
func (_LitionScClient *LitionScClientCaller) GetLastNotary(opts *bind.CallOpts, chain_id *big.Int) (struct {
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	ret := new(struct {
		LastNotaryBlock     *big.Int
		LastNotaryTimestamp *big.Int
	})
	out := ret
	err := _LitionScClient.contract.Call(opts, out, "get_last_notary", chain_id)
	return *ret, err
}

// GetLastNotary is a free data retrieval call binding the contract method 0xda26501d.
//
// Solidity: function get_last_notary(uint256 chain_id) constant returns(uint256 last_notary_block, uint256 last_notary_timestamp)
func (_LitionScClient *LitionScClientSession) GetLastNotary(chain_id *big.Int) (struct {
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetLastNotary(&_LitionScClient.CallOpts, chain_id)
}

// GetLastNotary is a free data retrieval call binding the contract method 0xda26501d.
//
// Solidity: function get_last_notary(uint256 chain_id) constant returns(uint256 last_notary_block, uint256 last_notary_timestamp)
func (_LitionScClient *LitionScClientCallerSession) GetLastNotary(chain_id *big.Int) (struct {
	LastNotaryBlock     *big.Int
	LastNotaryTimestamp *big.Int
}, error) {
	return _LitionScClient.Contract.GetLastNotary(&_LitionScClient.CallOpts, chain_id)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x2aad5f6e.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientCaller) GetSignatureHashFromNotary(opts *bind.CallOpts, notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "get_signature_hash_from_notary", notary_block, miners, blocks_mined, users, user_gas, largest_tx)
	return *ret0, err
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x2aad5f6e.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientSession) GetSignatureHashFromNotary(notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	return _LitionScClient.Contract.GetSignatureHashFromNotary(&_LitionScClient.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x2aad5f6e.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientCallerSession) GetSignatureHashFromNotary(notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	return _LitionScClient.Contract.GetSignatureHashFromNotary(&_LitionScClient.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetUserDetails is a free data retrieval call binding the contract method 0x677f7a3d.
//
// Solidity: function get_user_details(uint256 chain_id, address acc) constant returns(bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining)
func (_LitionScClient *LitionScClientCaller) GetUserDetails(opts *bind.CallOpts, chain_id *big.Int, acc common.Address) (struct {
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
	err := _LitionScClient.contract.Call(opts, out, "get_user_details", chain_id, acc)
	return *ret, err
}

// GetUserDetails is a free data retrieval call binding the contract method 0x677f7a3d.
//
// Solidity: function get_user_details(uint256 chain_id, address acc) constant returns(bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining)
func (_LitionScClient *LitionScClientSession) GetUserDetails(chain_id *big.Int, acc common.Address) (struct {
	Exists      bool
	Deposit     *big.Int
	Whitelisted bool
	Vesting     *big.Int
	Mining      bool
}, error) {
	return _LitionScClient.Contract.GetUserDetails(&_LitionScClient.CallOpts, chain_id, acc)
}

// GetUserDetails is a free data retrieval call binding the contract method 0x677f7a3d.
//
// Solidity: function get_user_details(uint256 chain_id, address acc) constant returns(bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining)
func (_LitionScClient *LitionScClientCallerSession) GetUserDetails(chain_id *big.Int, acc common.Address) (struct {
	Exists      bool
	Deposit     *big.Int
	Whitelisted bool
	Vesting     *big.Int
	Mining      bool
}, error) {
	return _LitionScClient.Contract.GetUserDetails(&_LitionScClient.CallOpts, chain_id, acc)
}

// GetUserRequests is a free data retrieval call binding the contract method 0xc6ae0f58.
//
// Solidity: function get_user_requests(uint256 chain_id, address acc) constant returns(bool vesting_req_exists, uint256 vesting_req_time, uint256 vesting_req_notary, uint256 vesting_req_value, uint256 vesting_req_state, uint256 vesting_req_control_state, bool deposit_req_exists, uint256 deposit_req_time, uint256 deposit_req_notary, uint256 deposit_req_value, uint256 deposit_req_state)
func (_LitionScClient *LitionScClientCaller) GetUserRequests(opts *bind.CallOpts, chain_id *big.Int, acc common.Address) (struct {
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
	err := _LitionScClient.contract.Call(opts, out, "get_user_requests", chain_id, acc)
	return *ret, err
}

// GetUserRequests is a free data retrieval call binding the contract method 0xc6ae0f58.
//
// Solidity: function get_user_requests(uint256 chain_id, address acc) constant returns(bool vesting_req_exists, uint256 vesting_req_time, uint256 vesting_req_notary, uint256 vesting_req_value, uint256 vesting_req_state, uint256 vesting_req_control_state, bool deposit_req_exists, uint256 deposit_req_time, uint256 deposit_req_notary, uint256 deposit_req_value, uint256 deposit_req_state)
func (_LitionScClient *LitionScClientSession) GetUserRequests(chain_id *big.Int, acc common.Address) (struct {
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
	return _LitionScClient.Contract.GetUserRequests(&_LitionScClient.CallOpts, chain_id, acc)
}

// GetUserRequests is a free data retrieval call binding the contract method 0xc6ae0f58.
//
// Solidity: function get_user_requests(uint256 chain_id, address acc) constant returns(bool vesting_req_exists, uint256 vesting_req_time, uint256 vesting_req_notary, uint256 vesting_req_value, uint256 vesting_req_state, uint256 vesting_req_control_state, bool deposit_req_exists, uint256 deposit_req_time, uint256 deposit_req_notary, uint256 deposit_req_value, uint256 deposit_req_state)
func (_LitionScClient *LitionScClientCallerSession) GetUserRequests(chain_id *big.Int, acc common.Address) (struct {
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
	return _LitionScClient.Contract.GetUserRequests(&_LitionScClient.CallOpts, chain_id, acc)
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 chain_id, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) HasDeposited(opts *bind.CallOpts, chain_id *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "has_deposited", chain_id, acc)
	return *ret0, err
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 chain_id, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) HasDeposited(chain_id *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasDeposited(&_LitionScClient.CallOpts, chain_id, acc)
}

// HasDeposited is a free data retrieval call binding the contract method 0xb747bb9b.
//
// Solidity: function has_deposited(uint256 chain_id, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) HasDeposited(chain_id *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasDeposited(&_LitionScClient.CallOpts, chain_id, acc)
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 chain_id, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCaller) HasVested(opts *bind.CallOpts, chain_id *big.Int, acc common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "has_vested", chain_id, acc)
	return *ret0, err
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 chain_id, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientSession) HasVested(chain_id *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasVested(&_LitionScClient.CallOpts, chain_id, acc)
}

// HasVested is a free data retrieval call binding the contract method 0xc87ef2fb.
//
// Solidity: function has_vested(uint256 chain_id, address acc) constant returns(bool)
func (_LitionScClient *LitionScClientCallerSession) HasVested(chain_id *big.Int, acc common.Address) (bool, error) {
	return _LitionScClient.Contract.HasVested(&_LitionScClient.CallOpts, chain_id, acc)
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

// CancelDepositInChain is a paid mutator transaction binding the contract method 0x3271dc81.
//
// Solidity: function cancel_deposit_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactor) CancelDepositInChain(opts *bind.TransactOpts, chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "cancel_deposit_in_chain", chain_id)
}

// CancelDepositInChain is a paid mutator transaction binding the contract method 0x3271dc81.
//
// Solidity: function cancel_deposit_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientSession) CancelDepositInChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelDepositInChain(&_LitionScClient.TransactOpts, chain_id)
}

// CancelDepositInChain is a paid mutator transaction binding the contract method 0x3271dc81.
//
// Solidity: function cancel_deposit_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactorSession) CancelDepositInChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelDepositInChain(&_LitionScClient.TransactOpts, chain_id)
}

// CancelVestInChain is a paid mutator transaction binding the contract method 0x17010a1a.
//
// Solidity: function cancel_vest_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactor) CancelVestInChain(opts *bind.TransactOpts, chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "cancel_vest_in_chain", chain_id)
}

// CancelVestInChain is a paid mutator transaction binding the contract method 0x17010a1a.
//
// Solidity: function cancel_vest_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientSession) CancelVestInChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelVestInChain(&_LitionScClient.TransactOpts, chain_id)
}

// CancelVestInChain is a paid mutator transaction binding the contract method 0x17010a1a.
//
// Solidity: function cancel_vest_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactorSession) CancelVestInChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.CancelVestInChain(&_LitionScClient.TransactOpts, chain_id)
}

// ConfirmDepositWithdrawalFromChain is a paid mutator transaction binding the contract method 0x3ea8c1e5.
//
// Solidity: function confirm_deposit_withdrawal_from_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactor) ConfirmDepositWithdrawalFromChain(opts *bind.TransactOpts, chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "confirm_deposit_withdrawal_from_chain", chain_id)
}

// ConfirmDepositWithdrawalFromChain is a paid mutator transaction binding the contract method 0x3ea8c1e5.
//
// Solidity: function confirm_deposit_withdrawal_from_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientSession) ConfirmDepositWithdrawalFromChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmDepositWithdrawalFromChain(&_LitionScClient.TransactOpts, chain_id)
}

// ConfirmDepositWithdrawalFromChain is a paid mutator transaction binding the contract method 0x3ea8c1e5.
//
// Solidity: function confirm_deposit_withdrawal_from_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactorSession) ConfirmDepositWithdrawalFromChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmDepositWithdrawalFromChain(&_LitionScClient.TransactOpts, chain_id)
}

// ConfirmVestInChain is a paid mutator transaction binding the contract method 0xf517cd96.
//
// Solidity: function confirm_vest_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactor) ConfirmVestInChain(opts *bind.TransactOpts, chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "confirm_vest_in_chain", chain_id)
}

// ConfirmVestInChain is a paid mutator transaction binding the contract method 0xf517cd96.
//
// Solidity: function confirm_vest_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientSession) ConfirmVestInChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmVestInChain(&_LitionScClient.TransactOpts, chain_id)
}

// ConfirmVestInChain is a paid mutator transaction binding the contract method 0xf517cd96.
//
// Solidity: function confirm_vest_in_chain(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactorSession) ConfirmVestInChain(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.ConfirmVestInChain(&_LitionScClient.TransactOpts, chain_id)
}

// Notary is a paid mutator transaction binding the contract method 0x6cd22ed8.
//
// Solidity: function notary(uint256 chain_id, uint256 notary_start_block, uint256 notary_end_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactor) Notary(opts *bind.TransactOpts, chain_id *big.Int, notary_start_block *big.Int, notary_end_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "notary", chain_id, notary_start_block, notary_end_block, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0x6cd22ed8.
//
// Solidity: function notary(uint256 chain_id, uint256 notary_start_block, uint256 notary_end_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientSession) Notary(chain_id *big.Int, notary_start_block *big.Int, notary_end_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, chain_id, notary_start_block, notary_end_block, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

// Notary is a paid mutator transaction binding the contract method 0x6cd22ed8.
//
// Solidity: function notary(uint256 chain_id, uint256 notary_start_block, uint256 notary_end_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_LitionScClient *LitionScClientTransactorSession) Notary(chain_id *big.Int, notary_start_block *big.Int, notary_end_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.Notary(&_LitionScClient.TransactOpts, chain_id, notary_start_block, notary_end_block, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x8e0808ed.
//
// Solidity: function register_chain(string info, address validator, uint96 vesting, uint96 deposit, string init_endpoint) returns(uint256 chain_id)
func (_LitionScClient *LitionScClientTransactor) RegisterChain(opts *bind.TransactOpts, info string, validator common.Address, vesting *big.Int, deposit *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "register_chain", info, validator, vesting, deposit, init_endpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x8e0808ed.
//
// Solidity: function register_chain(string info, address validator, uint96 vesting, uint96 deposit, string init_endpoint) returns(uint256 chain_id)
func (_LitionScClient *LitionScClientSession) RegisterChain(info string, validator common.Address, vesting *big.Int, deposit *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, info, validator, vesting, deposit, init_endpoint)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x8e0808ed.
//
// Solidity: function register_chain(string info, address validator, uint96 vesting, uint96 deposit, string init_endpoint) returns(uint256 chain_id)
func (_LitionScClient *LitionScClientTransactorSession) RegisterChain(info string, validator common.Address, vesting *big.Int, deposit *big.Int, init_endpoint string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterChain(&_LitionScClient.TransactOpts, info, validator, vesting, deposit, init_endpoint)
}

// RequestDepositInChain is a paid mutator transaction binding the contract method 0x2f0ae693.
//
// Solidity: function request_deposit_in_chain(uint256 chain_id, uint256 deposit) returns()
func (_LitionScClient *LitionScClientTransactor) RequestDepositInChain(opts *bind.TransactOpts, chain_id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "request_deposit_in_chain", chain_id, deposit)
}

// RequestDepositInChain is a paid mutator transaction binding the contract method 0x2f0ae693.
//
// Solidity: function request_deposit_in_chain(uint256 chain_id, uint256 deposit) returns()
func (_LitionScClient *LitionScClientSession) RequestDepositInChain(chain_id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestDepositInChain(&_LitionScClient.TransactOpts, chain_id, deposit)
}

// RequestDepositInChain is a paid mutator transaction binding the contract method 0x2f0ae693.
//
// Solidity: function request_deposit_in_chain(uint256 chain_id, uint256 deposit) returns()
func (_LitionScClient *LitionScClientTransactorSession) RequestDepositInChain(chain_id *big.Int, deposit *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestDepositInChain(&_LitionScClient.TransactOpts, chain_id, deposit)
}

// RequestVestInChain is a paid mutator transaction binding the contract method 0xd4a37ef9.
//
// Solidity: function request_vest_in_chain(uint256 chain_id, uint256 vesting) returns()
func (_LitionScClient *LitionScClientTransactor) RequestVestInChain(opts *bind.TransactOpts, chain_id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "request_vest_in_chain", chain_id, vesting)
}

// RequestVestInChain is a paid mutator transaction binding the contract method 0xd4a37ef9.
//
// Solidity: function request_vest_in_chain(uint256 chain_id, uint256 vesting) returns()
func (_LitionScClient *LitionScClientSession) RequestVestInChain(chain_id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestVestInChain(&_LitionScClient.TransactOpts, chain_id, vesting)
}

// RequestVestInChain is a paid mutator transaction binding the contract method 0xd4a37ef9.
//
// Solidity: function request_vest_in_chain(uint256 chain_id, uint256 vesting) returns()
func (_LitionScClient *LitionScClientTransactorSession) RequestVestInChain(chain_id *big.Int, vesting *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.RequestVestInChain(&_LitionScClient.TransactOpts, chain_id, vesting)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactor) StartMining(opts *bind.TransactOpts, chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "start_mining", chain_id)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientSession) StartMining(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StartMining(&_LitionScClient.TransactOpts, chain_id)
}

// StartMining is a paid mutator transaction binding the contract method 0x2a42c457.
//
// Solidity: function start_mining(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactorSession) StartMining(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StartMining(&_LitionScClient.TransactOpts, chain_id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactor) StopMining(opts *bind.TransactOpts, chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "stop_mining", chain_id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientSession) StopMining(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StopMining(&_LitionScClient.TransactOpts, chain_id)
}

// StopMining is a paid mutator transaction binding the contract method 0x67245c1d.
//
// Solidity: function stop_mining(uint256 chain_id) returns()
func (_LitionScClient *LitionScClientTransactorSession) StopMining(chain_id *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.StopMining(&_LitionScClient.TransactOpts, chain_id)
}

// TestNotary is a paid mutator transaction binding the contract method 0xa0b1eaa3.
//
// Solidity: function test_notary(uint256 chain_id, uint256 notary_block_no) returns()
func (_LitionScClient *LitionScClientTransactor) TestNotary(opts *bind.TransactOpts, chain_id *big.Int, notary_block_no *big.Int) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "test_notary", chain_id, notary_block_no)
}

// TestNotary is a paid mutator transaction binding the contract method 0xa0b1eaa3.
//
// Solidity: function test_notary(uint256 chain_id, uint256 notary_block_no) returns()
func (_LitionScClient *LitionScClientSession) TestNotary(chain_id *big.Int, notary_block_no *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.TestNotary(&_LitionScClient.TransactOpts, chain_id, notary_block_no)
}

// TestNotary is a paid mutator transaction binding the contract method 0xa0b1eaa3.
//
// Solidity: function test_notary(uint256 chain_id, uint256 notary_block_no) returns()
func (_LitionScClient *LitionScClientTransactorSession) TestNotary(chain_id *big.Int, notary_block_no *big.Int) (*types.Transaction, error) {
	return _LitionScClient.Contract.TestNotary(&_LitionScClient.TransactOpts, chain_id, notary_block_no)
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
// Solidity: event AcceptedVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterAcceptedVestInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientAcceptedVestInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "AcceptedVestInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientAcceptedVestInChainIterator{contract: _LitionScClient.contract, event: "AcceptedVestInChain", logs: logs, sub: sub}, nil
}

// WatchAcceptedVestInChain is a free log subscription operation binding the contract event 0x00547b633d4da14c6f2e748a506f553393a70da4e5cdd6cb304865b140f92f6c.
//
// Solidity: event AcceptedVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchAcceptedVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientAcceptedVestInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "AcceptedVestInChain", chain_idRule, accountRule)
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
// Solidity: event AcceptedVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
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
// Solidity: event CancelDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterCancelDepositInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientCancelDepositInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "CancelDepositInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientCancelDepositInChainIterator{contract: _LitionScClient.contract, event: "CancelDepositInChain", logs: logs, sub: sub}, nil
}

// WatchCancelDepositInChain is a free log subscription operation binding the contract event 0x9e6b1e30d746be09c6f5cf9e8ec4c7bf584bfe1308d70b39a097363e118814e5.
//
// Solidity: event CancelDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchCancelDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientCancelDepositInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "CancelDepositInChain", chain_idRule, accountRule)
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
// Solidity: event CancelDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
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
// Solidity: event CancelVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterCancelVestInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientCancelVestInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "CancelVestInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientCancelVestInChainIterator{contract: _LitionScClient.contract, event: "CancelVestInChain", logs: logs, sub: sub}, nil
}

// WatchCancelVestInChain is a free log subscription operation binding the contract event 0x698f91699fc921ec104a854863c577f1d32e8d48e41eac966924cbbdae212fea.
//
// Solidity: event CancelVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchCancelVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientCancelVestInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "CancelVestInChain", chain_idRule, accountRule)
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
// Solidity: event CancelVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
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
// Solidity: event ConfirmDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterConfirmDepositInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientConfirmDepositInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ConfirmDepositInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientConfirmDepositInChainIterator{contract: _LitionScClient.contract, event: "ConfirmDepositInChain", logs: logs, sub: sub}, nil
}

// WatchConfirmDepositInChain is a free log subscription operation binding the contract event 0x339a7256b6e5032769b6cee82c769df79aefc6215244052d987decc224251499.
//
// Solidity: event ConfirmDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchConfirmDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientConfirmDepositInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ConfirmDepositInChain", chain_idRule, accountRule)
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
// Solidity: event ConfirmDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
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
// Solidity: event ConfirmVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterConfirmVestInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientConfirmVestInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ConfirmVestInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientConfirmVestInChainIterator{contract: _LitionScClient.contract, event: "ConfirmVestInChain", logs: logs, sub: sub}, nil
}

// WatchConfirmVestInChain is a free log subscription operation binding the contract event 0xb5aaacabfddd3428a7e8b351250df8f590b10b3eb0709a08223d2730aa110732.
//
// Solidity: event ConfirmVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchConfirmVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientConfirmVestInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ConfirmVestInChain", chain_idRule, accountRule)
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
// Solidity: event ConfirmVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
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
// Solidity: event ForceWithdrawDeposit(uint256 indexed chain_id, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterForceWithdrawDeposit(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientForceWithdrawDepositIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ForceWithdrawDeposit", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientForceWithdrawDepositIterator{contract: _LitionScClient.contract, event: "ForceWithdrawDeposit", logs: logs, sub: sub}, nil
}

// WatchForceWithdrawDeposit is a free log subscription operation binding the contract event 0x09ddb46db0b11a56fa7611442b214b389d5479513714d4f5c596cdd62b39a59c.
//
// Solidity: event ForceWithdrawDeposit(uint256 indexed chain_id, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchForceWithdrawDeposit(opts *bind.WatchOpts, sink chan<- *LitionScClientForceWithdrawDeposit, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ForceWithdrawDeposit", chain_idRule, accountRule)
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
// Solidity: event ForceWithdrawDeposit(uint256 indexed chain_id, address indexed account, uint256 timestamp)
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
// Solidity: event ForceWithdrawVesting(uint256 indexed chain_id, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterForceWithdrawVesting(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientForceWithdrawVestingIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "ForceWithdrawVesting", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientForceWithdrawVestingIterator{contract: _LitionScClient.contract, event: "ForceWithdrawVesting", logs: logs, sub: sub}, nil
}

// WatchForceWithdrawVesting is a free log subscription operation binding the contract event 0xa3384ea61d5ab8755c01442bce838b3b1235b6f18e294620c2bc49de0073854f.
//
// Solidity: event ForceWithdrawVesting(uint256 indexed chain_id, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchForceWithdrawVesting(opts *bind.WatchOpts, sink chan<- *LitionScClientForceWithdrawVesting, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "ForceWithdrawVesting", chain_idRule, accountRule)
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
// Solidity: event ForceWithdrawVesting(uint256 indexed chain_id, address indexed account, uint256 timestamp)
func (_LitionScClient *LitionScClientFilterer) ParseForceWithdrawVesting(log types.Log) (*LitionScClientForceWithdrawVesting, error) {
	event := new(LitionScClientForceWithdrawVesting)
	if err := _LitionScClient.contract.UnpackLog(event, "ForceWithdrawVesting", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LitionScClientMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the LitionScClient contract.
type LitionScClientMigrateIterator struct {
	Event *LitionScClientMigrate // Event containing the contract specifics and raw log

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
func (it *LitionScClientMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientMigrate)
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
		it.Event = new(LitionScClientMigrate)
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
func (it *LitionScClientMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientMigrate represents a Migrate event raised by the LitionScClient contract.
type LitionScClientMigrate struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address new_contract)
func (_LitionScClient *LitionScClientFilterer) FilterMigrate(opts *bind.FilterOpts) (*LitionScClientMigrateIterator, error) {

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return &LitionScClientMigrateIterator{contract: _LitionScClient.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address new_contract)
func (_LitionScClient *LitionScClientFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *LitionScClientMigrate) (event.Subscription, error) {

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientMigrate)
				if err := _LitionScClient.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// ParseMigrate is a log parse operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address new_contract)
func (_LitionScClient *LitionScClientFilterer) ParseMigrate(log types.Log) (*LitionScClientMigrate, error) {
	event := new(LitionScClientMigrate)
	if err := _LitionScClient.contract.UnpackLog(event, "Migrate", log); err != nil {
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
// Solidity: event NewChain(uint256 chain_id, string description, string endpoint)
func (_LitionScClient *LitionScClientFilterer) FilterNewChain(opts *bind.FilterOpts) (*LitionScClientNewChainIterator, error) {

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "NewChain")
	if err != nil {
		return nil, err
	}
	return &LitionScClientNewChainIterator{contract: _LitionScClient.contract, event: "NewChain", logs: logs, sub: sub}, nil
}

// WatchNewChain is a free log subscription operation binding the contract event 0x86463e5a4c44c4d307742ef0abc183642f207e46cbaa6411b2bf7a118ab893e2.
//
// Solidity: event NewChain(uint256 chain_id, string description, string endpoint)
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
// Solidity: event NewChain(uint256 chain_id, string description, string endpoint)
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
// Solidity: event RequestDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterRequestDepositInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientRequestDepositInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "RequestDepositInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientRequestDepositInChainIterator{contract: _LitionScClient.contract, event: "RequestDepositInChain", logs: logs, sub: sub}, nil
}

// WatchRequestDepositInChain is a free log subscription operation binding the contract event 0x11b73dd0023b91d2be09f591166aeda0a3794d5c5d63607603e5d31c053c4804.
//
// Solidity: event RequestDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchRequestDepositInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientRequestDepositInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "RequestDepositInChain", chain_idRule, accountRule)
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
// Solidity: event RequestDepositInChain(uint256 indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp)
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
// Solidity: event RequestVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) FilterRequestVestInChain(opts *bind.FilterOpts, chain_id []*big.Int, account []common.Address) (*LitionScClientRequestVestInChainIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "RequestVestInChain", chain_idRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientRequestVestInChainIterator{contract: _LitionScClient.contract, event: "RequestVestInChain", logs: logs, sub: sub}, nil
}

// WatchRequestVestInChain is a free log subscription operation binding the contract event 0x5be03bab1e7e8d384248dfcee1d12d03a399213ac09c422aac74b6e766de144d.
//
// Solidity: event RequestVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
func (_LitionScClient *LitionScClientFilterer) WatchRequestVestInChain(opts *bind.WatchOpts, sink chan<- *LitionScClientRequestVestInChain, chain_id []*big.Int, account []common.Address) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "RequestVestInChain", chain_idRule, accountRule)
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
// Solidity: event RequestVestInChain(uint256 indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp)
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
	Miner     common.Address
	Whitelist bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAccount is a free log retrieval operation binding the contract event 0x2c8cbcff405aa83b0e160273c7f1159141d2bdc795ecddf1e3413dc2995f0d7f.
//
// Solidity: event WhitelistAccount(uint256 indexed chain_id, address miner, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) FilterWhitelistAccount(opts *bind.FilterOpts, chain_id []*big.Int) (*LitionScClientWhitelistAccountIterator, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "WhitelistAccount", chain_idRule)
	if err != nil {
		return nil, err
	}
	return &LitionScClientWhitelistAccountIterator{contract: _LitionScClient.contract, event: "WhitelistAccount", logs: logs, sub: sub}, nil
}

// WatchWhitelistAccount is a free log subscription operation binding the contract event 0x2c8cbcff405aa83b0e160273c7f1159141d2bdc795ecddf1e3413dc2995f0d7f.
//
// Solidity: event WhitelistAccount(uint256 indexed chain_id, address miner, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) WatchWhitelistAccount(opts *bind.WatchOpts, sink chan<- *LitionScClientWhitelistAccount, chain_id []*big.Int) (event.Subscription, error) {

	var chain_idRule []interface{}
	for _, chain_idItem := range chain_id {
		chain_idRule = append(chain_idRule, chain_idItem)
	}

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "WhitelistAccount", chain_idRule)
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
// Solidity: event WhitelistAccount(uint256 indexed chain_id, address miner, bool whitelist)
func (_LitionScClient *LitionScClientFilterer) ParseWhitelistAccount(log types.Log) (*LitionScClientWhitelistAccount, error) {
	event := new(LitionScClientWhitelistAccount)
	if err := _LitionScClient.contract.UnpackLog(event, "WhitelistAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}
