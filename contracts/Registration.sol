
pragma solidity >= 0.5.11;

interface ChainValidator {
   function check_vesting(uint vesting, address participant) external returns (bool);
   function check_deposit(uint vesting, address participant) external returns (bool);
}

interface ERC20 {
   function totalSupply() external view returns (uint);
   function balanceOf(address tokenOwner) external view returns (uint balance);
   function allowance(address tokenOwner, address spender) external view returns (uint remaining);
   function transfer(address to, uint tokens) external returns (bool success);
   function approve(address spender, uint tokens) external returns (bool success);
   function transferFrom(address from, address to, uint tokens) external returns (bool success);
   event Transfer(address indexed from, address indexed to, uint tokens);
   event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract LitionChainValidator is ChainValidator {
   function check_vesting(uint vesting, address participant) public returns (bool) {
      if(vesting >= 1000*(10**18) && vesting <= 500000*(10**18)) {
        return true;   
      }
      return false;
   }

   function check_deposit(uint deposit, address participant) public returns (bool) {
      if(deposit >= 1000*(10**18)) {
         return true;
      }
      return false;
   }
}

contract LitionRegistry{
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function check_lition_min_vesting(uint vesting) private pure returns (bool) {
        if(vesting >= 1000*(10**18)) {
            return true;   
        }
        return false;
    }
    
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function check_lition_min_deposit(uint deposit) private pure returns (bool) {
        if(deposit >= 1000*(10**18)) {
            return true;   
        }
        return false; 
    }
    
    // New chain was registered
    event NewChain(uint256 chain_id, string description, string endpoint);
    
    /**** Events related to the deposit requests ****/
    // Deposit request created
    event RequestDepositInChain(uint indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp /* now */);
    // Deposit request confirmed - tokens were transferred and account's deposit balance was updated
    event ConfirmDepositInChain(uint indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp /* request creation time */);
    // Deposit request cancelled
    event CancelDepositInChain(uint indexed chain_id, address indexed account, uint256 deposit, uint256 req_timestamp /* request creation time */);
    // Whole deposit withdrawned - this is special withdrawal, which as applied after the chain validators are not able to reach consensus for 1 month 
    event ForceWithdrawDeposit(uint indexed chain_id, address indexed account, uint timestamp /* now */);
    
    /**** Events related to the vesting requests ****/
    // Vesting request created
    event RequestVestInChain(uint indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp /* now */);
    // Vesting request confirmed - tokens were transferred
    event ConfirmVestInChain(uint indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp /* request creation time */);
    // Vesting request cancelled
    event CancelVestInChain(uint indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp /* request creation time */);
    // Vesting request accepted - account's vesting balance was updated 
    event AcceptedVestInChain(uint indexed chain_id, address indexed account, uint256 vesting, uint256 req_timestamp /* request creation time */);
    // Whole vesting withdrawned - this is special withdrawal, which as applied after the chain validators are not able to reach consensus for 1 month 
    event ForceWithdrawVesting(uint indexed chain_id, address indexed account, uint timestamp /* now */);
    
    // if whitelist == true  - allow user to transact
    // if whitelist == false - do not allow user to transact
    event WhitelistAccount(uint indexed chain_id, address miner, bool whitelist);
    
    // Validator start/stop mining
    event StartMining(uint indexed chain_id, address miner);
    event StopMining(uint indexed chain_id, address miner);


    /**************************************************************************************************************************/
    /***************************************** Structs related to the list of users *******************************************/
    /**************************************************************************************************************************/
    struct Validator {
        // Actual user's vesting
        uint96  vesting;
        // Flag if user is mining -> set in start/stop_mining
        bool    mining;
    }
    
    struct Transactor {
        // Actual user's deposit
        uint96  deposit;
        // Flag if user is whitelisted (allowed to transact) -> actual deposit must be greater than min. required deposit condition 
        bool    whitelisted;
    }
    
    // Optimized "User_entry" so it packs together with User_details to 256 Bits (32 Bytes) chunk of memory
    struct User_entry {
        // Validator's data
        Validator    validator;
        // Transactor's data
        Transactor   transactor;
        // index to the users_list, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing user
        uint48          index;
    }
    
    struct Users {
        mapping(address => User_entry) accounts;
        address[]                      list;        
    }
    
    
    /**************************************************************************************************************************/
    /*********************************** Structs related to the vesting/deposit requests **************************************/
    /**************************************************************************************************************************/
    enum Request_state {
        REQUEST_CREATED,
        REQUEST_CONFIRMED
    }
    
    enum VestingRequestControl_state {
        // Wait for vesting_request_confirm to be called. It is used when new_vesting > actual validator's vesting
        WAIT_FOR_CONFIRMATION,
        // Validator's actual vesting will be replaced by the new_vesting from VestingRequest after the next notary
        REPLACE_VESTING,
        // Validator's actual vesting was replaced by the new_vesting from VestingRequest after the previous notary 
        VESTING_REPLACED
    }
    
    // 512 Bits
    struct VestingRequest_data {
        // Last notary block number when the request was accepted 
        uint256                         notary_block;
        // Timestamp(now) when the request was accepted/created
        uint48                          timestamp;
        // In old_vesting is stored current vesting that validator had when new VestingRequest was accepted
        uint96                          old_vesting;
        // New value of vesting to be set
        uint96                          new_vesting;
        // Actual state of the request
        Request_state                   state;
        // Actual control state of the request
        VestingRequestControl_state     control_state;
    }
    
    // Only full deposit withdrawals are saved as deposit requests - other types of deposits do not need to be confirmed
    // 312 Bits
    struct DepositWithdrawRequest_data {
        // Last notary block number when the request was accepted 
        uint256                         notary_block;
        // Timestamp(now) when the request was accepted/created
        uint48                          timestamp;
        // Actual state of the request
        Request_state                   state;
    }
    
    // 872 Bits
    struct Requests_entry {
        // index to the requests_list, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing request
        uint48                          index;
        // Deposit withdrawal request details
        DepositWithdrawRequest_data     deposit_withdraw_request;
        // Vesting request details
        VestingRequest_data             vesting_request;   
    }
    
    struct Requests {
        mapping(address => Requests_entry)  accounts;
        address[]                           list;        
    }
    
    
    /**************************************************************************************************************************/
    /***************************************************** Other structs ******************************************************/
    /**************************************************************************************************************************/
    ERC20 token;
   
    struct signature {
        uint8 v; bytes32 r; bytes32 s;
    }
    
    struct LastNotary {
        // Timestamp, when the last notary was accepted
        uint timestamp;
        // Actual block number, when the last notary was accepted
        uint block;
    }
    
    struct ChainInfo {
        bool              active;
        uint96            total_vesting;
        LastNotary        last_notary;
        ChainValidator    chain_validator;
        Users             users;
        Requests          requests;
        string            endpoint;
    }
    
    mapping(uint256 => ChainInfo) private chains;
    uint256 public next_id = 0;

    
    /**************************************************************************************************************************/
    /************************************ Contract interface - external/public functions **************************************/
    /**************************************************************************************************************************/
    
    // Requests vest in chain. It will be processed and applied to the actual user state after next:
    //      * 2 notary windows - in case new vesting < actual vesting
    //      * 3 notary windows - in case new vesting > actual vesting
    function request_vest_in_chain(uint chain_id, uint256 vesting) external {
      // Withdraw all vesting
      if (vesting == 0) {
          // If last notary is older than 30 days, it means that validators cannot reach consensus and side-chain is basically stuck.
          // In such case ignore multi-step vesting process and allow users to withdraw all vested tokens
          if (chains[chain_id].last_notary.timestamp + 30 days < now) {
              _force_withdraw_vest_from_chain(chain_id, msg.sender);
              return;
          }
          
          require(validator_exists(chain_id, msg.sender) == true, "Trying to withdraw vesting from non-existing validator account");
          require(chains[chain_id].users.accounts[msg.sender].validator.mining == false, "Can't withdraw any tokens, stop_minig must be called first.");  
      }
      // Vest in chain or withdraw just part of vesting
      else {
         require(chains[chain_id].active, "can't vest into non-existing chain");
         require(check_lition_min_vesting(vesting), "user does not meet Lition's min.required vesting condition");
         require(chains[chain_id].chain_validator.check_vesting(vesting, msg.sender), "user does not meet chain validator's min.required vesting condition");
         require(vesting <= ~uint96(0), "vesting is greater than uint96_max_value");
      }

      require(vesting_request_exists(chain_id, msg.sender) == false, "Cannot vest in chain. There is already ongoing request being processed for this acc.");
      require(chains[chain_id].users.accounts[msg.sender].validator.vesting != vesting, "Cannot vest the same amount of tokens as you already has vested.");
      
      _request_vest_in_chain(chain_id, vesting, msg.sender);
    }
    
    // Confirms vest request, token transfer is processed during confirmation
    function confirm_vest_in_chain(uint chain_id) external {
        require(vesting_request_exists(chain_id, msg.sender) == true, "Cannot confirm non-existing vesting request.");
        require(chains[chain_id].last_notary.block > chains[chain_id].requests.accounts[msg.sender].vesting_request.notary_block, "Request confirmation can be called in the next notary window after request was accepted.");
        require(chains[chain_id].requests.accounts[msg.sender].vesting_request.state == Request_state.REQUEST_CREATED, "Cannot confirm already confirmed request.");
        
        _confirm_vest_in_chain(chain_id, msg.sender);
    }
    
    // Cancels the existing vest request. Such request can be cancelled only if it was not already confirmed
    function cancel_vest_in_chain(uint chain_id) external {
        require(vesting_request_exists(chain_id, msg.sender) == true, "Cannot cancel non-existing vesting request.");
        require(chains[chain_id].requests.accounts[msg.sender].vesting_request.state == Request_state.REQUEST_CREATED, "Cannot cancel already confirmed request." );
        
        _cancel_vest_in_chain(chain_id, msg.sender);
    }
    
    function request_deposit_in_chain(uint chain_id, uint256 deposit) external {
        // Withdraw whole deposit
        if (deposit == 0) {
          // If last notary is older than 30 days, it means that validators cannot reach consensus and side-chain is basically stuck.
          // In such case ignore multi-step deposit withdrawal process and allow users to withdraw all deposited tokens
          if (chains[chain_id].last_notary.timestamp + 30 days < now) {
              _force_withdraw_deposit_from_chain(chain_id, msg.sender);
              return;
          }
          
          require(transactor_exists(chain_id, msg.sender) == true, "Trying to withdraw deposit from non-existing transactor account");
        }
        // Deposit in chain or withdraw just part of vesting
        else {
         require(chains[chain_id].active, "can't deposit into non-existing chain");
         require(check_lition_min_deposit(deposit), "user does not meet Lition's min.required deposit condition");
         require(chains[chain_id].chain_validator.check_deposit(deposit, msg.sender), "user does not meet chain validator's min.required deposit condition");
         require(deposit <= ~uint96(0), "deposit is greater than uint96_max_value");
        }
        
        require(deposit_withdraw_request_exists(chain_id, msg.sender) == false, "Cannot deposit in chain. There is ongoing withdrawal request being processed for this acc.");
        require(chains[chain_id].users.accounts[msg.sender].transactor.deposit != deposit, "Cannot deposit the same amount of tokens as you already has vested.");
        
        _request_deposit_in_chain(chain_id, deposit, msg.sender);
    }
    
    // Confirms deposit withdrawal request, token transfer is processed during confirmation
    function confirm_deposit_withdrawal_from_chain(uint chain_id) external {
        require(deposit_withdraw_request_exists(chain_id, msg.sender) == true, "Cannot confirm non-existing deposit withdrawal request.");
        require(chains[chain_id].last_notary.block > chains[chain_id].requests.accounts[msg.sender].deposit_withdraw_request.notary_block, "Request confirmation can be called in the next notary window after request was accepted.");
        require(chains[chain_id].requests.accounts[msg.sender].deposit_withdraw_request.state == Request_state.REQUEST_CREATED, "Cannot confirm already confirmed request.");
        
        _confirm_deposit_withdrawal_from_chain(chain_id, msg.sender);
    }
    
    // Cancels the existing deposit request. Such request can be cancelled only if it was not already confirmed
    function cancel_deposit_in_chain(uint chain_id) external {
        require(deposit_withdraw_request_exists(chain_id, msg.sender) == true, "Cannot cancel non-existing deposit withdrawal request.");
        require(chains[chain_id].requests.accounts[msg.sender].deposit_withdraw_request.state == Request_state.REQUEST_CREATED, "Cannot cancel already confirmed request." );
        
        _cancel_deposit_withdrawal_from_chain(chain_id, msg.sender);
    }
    
    // Internally creates/registers new side-chain. Creator must be also validator at least from the beginning as joining process take multiple steps
    // and these steps cannot be done in the same notary window
    function register_chain(string calldata info, ChainValidator validator, uint96 vesting, uint96 deposit, string calldata init_endpoint) external returns (uint256 chain_id) {
        require(bytes(init_endpoint).length > 0);
        
        address creator         = msg.sender;
        uint256 timestamp       = now;
        
        // Inits chain data
        chain_id                = next_id;
        ChainInfo storage chain = chains[chain_id];
        
        
        chain.chain_validator   = validator;
        chain.active            = true;
        chain.endpoint          = init_endpoint;
        
        // Validates vesting
        require(check_lition_min_vesting(vesting), "chain creator does not meet Lition's min.required vesting condition");
        require(chain.chain_validator.check_vesting(vesting, creator), "chain creator does not meet chain validator's min.required vesting condition");
        
        // Validates deposit
        require(check_lition_min_deposit(deposit), "chain creator does not meet Lition's min.required deposit condition");
        require(chain.chain_validator.check_deposit(deposit, creator), "chain creator does not meet chain validator's min.required deposit condition");
        
        // Internally creates new user
        user_create(chain_id, creator);
        
        // Transfers vesting tokens
        token.transferFrom(creator, address(this), vesting);
        chain.users.accounts[creator].validator.vesting = vesting;      
        
        // Transfers deposit tokens
        token.transferFrom(creator, address(this), deposit);
        chain.users.accounts[creator].transactor.deposit = deposit;      
        chain.users.accounts[creator].transactor.whitelisted = true;
        
        emit NewChain(chain_id, info, init_endpoint);
        
        emit RequestVestInChain(chain_id, creator, vesting, timestamp);
        emit ConfirmVestInChain(chain_id, creator, vesting, timestamp);
        emit AcceptedVestInChain(chain_id, creator, vesting, timestamp);
        
        emit RequestDepositInChain(chain_id, creator, deposit, timestamp);
        emit ConfirmDepositInChain(chain_id, creator, deposit, timestamp);
        
        next_id++;
    }
    
    // Reurns true, if acc has vested enough to become validator, othervise false
    function has_vested(uint chain_id, address acc) view external returns (bool) {
        // No need to check vesting balance as it cannot be lover than min. required
        return validator_exists(chain_id, acc);
    }
    
    // Returns true if user's remaining deposit balance >= min. required deposit and is allowed to transact
    function has_deposited(uint chain_id, address acc) view external returns (bool) {
        // No need to check deposit balance as whitelisted flag should be alwyas set accordingly
        return chains[chain_id].users.accounts[acc].transactor.whitelisted;
    }
    
    function get_signature_hash_from_notary(uint256 notary_block, address[] memory miners,
                                          uint32[] memory blocks_mined, address[] memory users,
                                          uint32[] memory user_gas, uint32 largest_tx)
                                            public pure returns (bytes32) {
        return keccak256(abi.encodePacked(notary_block, miners, blocks_mined, users, user_gas, largest_tx));
    }
    
    function get_last_notary(uint256 chain_id) external view returns (uint256 last_notary_block, uint256 last_notary_timestamp) {
        last_notary_block = chains[chain_id].last_notary.block;
        last_notary_timestamp = chains[chain_id].last_notary.timestamp;
    }
    
    function test_notary(uint256 chain_id, uint256 notary_block_no) external {
        ChainInfo storage chain = chains[chain_id];
        
        // Process existing vesting/deposit withdrawals requests as these cannot be 
        // processed automatically - math for user's usage and miner's rewards calculations would be invalid
        process_existing_requests(chain_id);
        
        // TODO: remove validators(call stop_mining) who signed no block during this notary window and have mining flag == true
        
        // Updates info when the last notary was processed 
        chain.last_notary.block = notary_block_no;
        chain.last_notary.timestamp = now;
    }
    
    function get_user_details(uint256 chain_id, address acc) external view returns (uint256 deposit, bool whitelisted, uint256 vesting, bool mining) {
        deposit = chains[chain_id].users.accounts[acc].transactor.deposit;
        whitelisted = chains[chain_id].users.accounts[acc].transactor.whitelisted;
        
        vesting = chains[chain_id].users.accounts[acc].validator.vesting;
        mining = chains[chain_id].users.accounts[acc].validator.mining;  
    }
    
    function get_user_requests(uint256 chain_id, address acc) external view returns (bool vesting_req_exists, uint256 vesting_req_time, uint256 vesting_req_notary, uint256 vesting_req_value, uint256 vesting_req_state, uint256 vesting_req_control_state,
                                                                                     bool deposit_req_exists, uint256 deposit_req_time, uint256 deposit_req_notary, uint256 deposit_req_value, uint256 deposit_req_state) {
        if (vesting_request_exists(chain_id, acc)) {
            VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
            
            vesting_req_exists = true;
            vesting_req_time = request.timestamp;
            vesting_req_notary = request.notary_block;
            vesting_req_value = request.new_vesting;
            vesting_req_state = uint256(request.state);
            vesting_req_control_state = uint256(request.control_state);
        }
        
        if (deposit_withdraw_request_exists(chain_id, acc)) {
            DepositWithdrawRequest_data storage request = chains[chain_id].requests.accounts[acc].deposit_withdraw_request;
            
            deposit_req_exists = true;
            deposit_req_time = request.timestamp;
            deposit_req_notary = request.notary_block;
            deposit_req_value = 0;
            deposit_req_state = uint256(request.state);
        }
    }
    
    function notary(uint256 chain_id, uint256 notary_block_no, address[] memory miners, uint32[] memory blocks_mined,
              address[] memory users, uint32[] memory user_gas, uint32 largest_tx,
              uint8[] memory v, bytes32[] memory r, bytes32[] memory s) public {
    
        // First, calculate hash from miners, block_mined, users and user_gas
        // then, do ec_recover of the signatures to determine signers
        // check if there is enough signers (total vesting of signers > 50% of all vestings)
        // then, calculate reward
        require(v.length == r.length);
        require(v.length == s.length);
        
        bytes32 signature_hash = get_signature_hash_from_notary(notary_block_no, miners, blocks_mined, users, user_gas, largest_tx);
        
        ChainInfo storage chain = chains[chain_id];
        require(chain.active, "Trying to report about non-existing chain");
        // TODO: check all other required stuff
        
        // Involved vesting based on validator's, who signed statistics for this notary window. 
        // These statistics are used for calculating usage cost and miner rewards are calculated
        uint256 involved_vesting = 0;
        for(uint256 i = 0; i < v.length; i++) {
            involved_vesting += chain.users.accounts[ecrecover(signature_hash, v[i], r[i], s[i])].validator.vesting;
        }
        
        // There must be at least 50% out of total possible vesting involved
        require(involved_vesting * 2 >= chain.total_vesting);
        
        // Calculates total cost based on user's usage durint current notary window
        uint256 total_cost = process_users_consumptions(chain_id, users, user_gas, largest_tx);
        
        // Calculates and process validator's rewards based on their participation rate and vesting balance
        process_miners_rewards(chain_id, notary_block_no, miners, blocks_mined, total_cost);
        
        // Process existing vesting/deposit withdrawals requests as these cannot be 
        // processed automatically - math for user's usage and miner's rewards calculations would be invalid
        process_existing_requests(chain_id);
        
        // TODO: remove validators(call stop_mining) who signed no block during this notary window and have mining flag == true
        
        // Updates info when the last notary was processed 
        chain.last_notary.block = notary_block_no;
        chain.last_notary.timestamp = now;
    }
    
    
    // Returns list of user's addresses that are allowed to transact - their deposit >= min. required deposit
    function get_allowed_to_transact(uint chain_id, uint batch) view external returns (address[100] memory, uint256, bool) {
        return _get_users(chain_id, batch, false, false);
    }
    
    // Returns list of validator's addresses that are actively participating in mining
    function get_allowed_to_validate(uint chain_id, uint batch) view external returns (address[100] memory, uint256, bool) {
        return _get_users(chain_id, batch, true, false);
    }
    
    // Returns list of validator's addresses that are allowed to be participating in mining based on their vesting, but not yet mining
    function get_active_validators(uint chain_id, uint batch) view external returns (address[100] memory, uint256, bool) {
        return _get_users(chain_id, batch, true, true);
    }
    
    function start_mining(uint chain_id) external {
        require(chains[chain_id].active == true, "Can't start mining on non-existing chain");
        require(check_lition_min_vesting(chains[chain_id].users.accounts[msg.sender].validator.vesting) == true, "user does not meet Lition's min.required vesting condition");
        require(chains[chain_id].chain_validator.check_vesting(chains[chain_id].users.accounts[msg.sender].validator.vesting, msg.sender) == true, "User does not meet chain validator's min.required vesting condition");
        
        _start_mining(chain_id, msg.sender);
    }
  
    function stop_mining(uint chain_id) external {
        require(chains[chain_id].active == true, "Can't start mining on non-existing chain");
        require(check_lition_min_vesting(chains[chain_id].users.accounts[msg.sender].validator.vesting) == true, "user does not meet Lition's min.required vesting condition");
        
        _stop_mining(chain_id, msg.sender);
    }
    

    /**************************************************************************************************************************/
    /**************************************** Functions related to the list of users ******************************************/
    /**************************************************************************************************************************/
    
    // Creates new user - does not set it's data yet as it is done after vesting/deposit_withdraw request is confirmed
    function user_create(uint256 chain_id, address acc) private {
        require(chains[chain_id].users.list.length < ~uint48(0), "count(users) is equal to max_count");
        require(user_exists(chain_id, acc) == false, "Creating already-created user");
        
        chains[chain_id].users.list.push(acc);
        chains[chain_id].users.accounts[acc].index = uint48(chains[chain_id].users.list.length); // indexes are stored + 1
    }
    
    // Deletes existing user from the internal list of users
    // This method should never be called directly, validator_delete & transactor_delete should be called instead
    function user_delete(uint256 chain_id, address acc) private {
        User_entry storage entry = chains[chain_id].users.accounts[acc];
        // user_exists(chain_id, acc) could be used instead
        require(entry.index != 0, "User does not exist");
        
        address[] storage users_list = chains[chain_id].users.list;
        require(entry.index <= users_list.length, "Invalid index value");
    
        // Move an last element of array into the vacated key slot.
        uint48 found_index = uint48(entry.index - 1);
        uint48 last_index = uint48(users_list.length - 1);
    
        chains[chain_id].users.accounts[users_list[last_index]].index = found_index + 1;
        users_list[found_index] = users_list[last_index];
        users_list.length--;
    
        delete chains[chain_id].users.accounts[acc];
    }
    
    // Deletes validator
    function validator_delete(uint chain_id, address acc) private {
        // There is no existing transactor for this account - delete whole requests struct 
        if (chains[chain_id].users.accounts[acc].transactor.deposit == 0) {
            user_delete(chain_id, acc);
            return;
        } 
        
        // There is exiting transactor for this account - only reset validator
        Validator storage validator = chains[chain_id].users.accounts[acc].validator;
        validator.vesting   = 0;
        validator.mining    = false;
    }
    
    // Deletes transactor
    function transactor_delete(uint chain_id, address acc) private {
        // There is no existing validator for this account - delete whole requests struct 
        if (chains[chain_id].users.accounts[acc].validator.vesting == 0) {
            user_delete(chain_id, acc);
            return;
        } 
        
        // There is existing validator for this account - only reset transactor
        Transactor storage transactor = chains[chain_id].users.accounts[acc].transactor;
        transactor.deposit      = 0;
        transactor.whitelisted  = false;
    }
    
    function user_exists(uint256 chain_id, address acc) internal view returns (bool) {
      return chains[chain_id].users.accounts[acc].index != 0;
    }
    
    function validator_exists(uint256 chain_id, address acc) internal view returns (bool) {
      return chains[chain_id].users.accounts[acc].validator.vesting > 0;
    }
    
    function transactor_exists(uint256 chain_id, address acc) internal view returns (bool) {
      return chains[chain_id].users.accounts[acc].transactor.deposit > 0;
    }
    
    
    
    /**************************************************************************************************************************/
    /*********************************** Functions related to the vesting/deposit requests ************************************/
    /**************************************************************************************************************************/
    
    // Creates new vesting request
    function vesting_request_create(uint256 chain_id, address acc, uint256 vesting) private {
        require(chains[chain_id].requests.list.length < ~uint48(0), "count(requests) is equal to max_count");
        
        Requests_entry storage requests_entry = chains[chain_id].requests.accounts[acc];
        
        requests_entry.vesting_request.old_vesting = chains[chain_id].users.accounts[acc].validator.vesting;
        requests_entry.vesting_request.new_vesting = uint96(vesting);
        if (requests_entry.vesting_request.new_vesting >= requests_entry.vesting_request.old_vesting) { // == case should never happen as it is handled in the caller's function
            requests_entry.vesting_request.control_state = VestingRequestControl_state.WAIT_FOR_CONFIRMATION;
        } else {
            requests_entry.vesting_request.control_state = VestingRequestControl_state.REPLACE_VESTING;
        }
        
        requests_entry.vesting_request.state = Request_state.REQUEST_CREATED;
        requests_entry.vesting_request.timestamp = uint48(now);
        requests_entry.vesting_request.notary_block = chains[chain_id].last_notary.block; 
        
        
        // There is no deposit or vesting ongoing request - create new requests_entry structure
        if (requests_entry.index == 0) { // any_request_exists(chain_id, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chains[chain_id].requests.list.push(acc);    
            requests_entry.index = uint48(chains[chain_id].requests.list.length); // indexes are stored + 1
        }
    }

    // Creates new deposit withdrawal request
    function deposit_withdraw_request_create(uint256 chain_id, address acc) private {
        require(chains[chain_id].requests.list.length < ~uint48(0), "count(requests) is equal to max_count");
        
        Requests_entry storage requests_entry = chains[chain_id].requests.accounts[acc];
        
        requests_entry.deposit_withdraw_request.state = Request_state.REQUEST_CREATED;
        requests_entry.deposit_withdraw_request.timestamp = uint48(now);
        requests_entry.deposit_withdraw_request.notary_block = chains[chain_id].last_notary.block; 
        
        // There is no deposit or vesting ongoing request - create new requests_entry structure
        if (requests_entry.index == 0) { // any_request_exists(chain_id, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chains[chain_id].requests.list.push(acc);    
            requests_entry.index = uint48(chains[chain_id].requests.list.length); // indexes are stored + 1
        }
    }

        
    // Deletes existing requests pair(vesting & deposit) from the internal list of requests
    // This method should never be called directly, vesting_request_delete & deposit_withdraw_request_delete should be called instead
    function requests_pair_delete(uint chain_id, address acc) private {
        Requests_entry storage entry = chains[chain_id].requests.accounts[acc];
        
        // request_exists(chain_id, acc), vesting_request_exists(chain_id, acc) and deposoti_withdraw_exists(chain_id, acc) could be used instead
        require(entry.index != 0, "Request does not exist");
        
        address[] storage requests_list = chains[chain_id].requests.list;
    
        require(entry.index <= requests_list.length, "Invalid index value");
    
        // Move an last element of array into the vacated key slot.
        uint48 found_index = uint48(entry.index - 1);
        uint48 last_index = uint48(requests_list.length - 1);
    
        chains[chain_id].requests.accounts[requests_list[last_index]].index = found_index + 1;
        requests_list[found_index] = requests_list[last_index];
        requests_list.length--;
    
        delete chains[chain_id].requests.accounts[acc];
    }
    
    function vesting_request_delete(uint chain_id, address acc) private {
        // There is no ongoing deposit request for this account - delete whole requests struct 
        if (chains[chain_id].requests.accounts[acc].deposit_withdraw_request.timestamp == 0) {
            requests_pair_delete(chain_id, acc);
            return;
        } 
        
        // There is ongoing deposit request for this account - only reset vesting request
        VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
        request.notary_block    = 0;
        request.timestamp       = 0;
        request.old_vesting     = 0;
        request.new_vesting     = 0;
        // First enum value is default
        request.state           = Request_state.REQUEST_CREATED; 
        request.control_state   = VestingRequestControl_state.WAIT_FOR_CONFIRMATION;
    }
    
    function deposit_withdraw_request_delete(uint chain_id, address acc) private {
        // There is no ongoing vesting request for this account - delete whole requests struct 
        if (chains[chain_id].requests.accounts[acc].vesting_request.timestamp == 0) {
            requests_pair_delete(chain_id, acc);
            return;
        } 
        
        // There is ongoing vesting request for this account - only reset vesting request
        DepositWithdrawRequest_data storage request = chains[chain_id].requests.accounts[acc].deposit_withdraw_request;
        request.notary_block    = 0;
        request.timestamp       = 0;
        // First enum value is default
        request.state           = Request_state.REQUEST_CREATED; 
    }
    
    // Checks if acc has any ongoing vesting or deposit request
    function any_request_exists(uint chain_id, address acc) private view returns (bool) {
      return chains[chain_id].requests.accounts[acc].index != 0;
    }
    
    // Checks if acc has any ongoing vesting request
    function vesting_request_exists(uint chain_id, address acc) private view returns (bool) {
      return chains[chain_id].requests.accounts[acc].vesting_request.timestamp != 0;
    }
    
    // Checks if acc has any ongoing DEPOSIT WITHDRAWAL request
    function deposit_withdraw_request_exists(uint chain_id, address acc) private view returns (bool) {
      return chains[chain_id].requests.accounts[acc].deposit_withdraw_request.timestamp != 0;
    }
    
    function _request_vest_in_chain(uint chain_id, uint256 vesting, address acc) private {
      // Internally creates new user
      if (vesting != 0 && user_exists(chain_id, acc) == false) {
          user_create(chain_id, acc);
      }
      
      vesting_request_create(chain_id, acc, vesting);
      emit RequestVestInChain(chain_id, acc, vesting, now);
    }
    
    // Basically just transfers the tokens, validator's vesting balance update is always done at the of notary atomatically
    function _confirm_vest_in_chain(uint chain_id, address acc) private {
        VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
        Validator storage validator = chains[chain_id].users.accounts[acc].validator;
        
        request.state = Request_state.REQUEST_CONFIRMED;
        
        // Decreases account's vesting in chain
        if(request.new_vesting < request.old_vesting) {
            // This should never happen during normal conditions as vesting balance state is updated during notary
            require(request.control_state == VestingRequestControl_state.VESTING_REPLACED, "Cannot withdraw vesting tokens, internal balance was not updated yet");
            
            uint96 to_withdraw = request.old_vesting - request.new_vesting;
            token.transfer(acc, to_withdraw);
            
            if (validator.mining == true) {
                chains[chain_id].total_vesting -= to_withdraw;
            }
            
            // If it was request to withdraw whole vesting balance, delete validator
            if (request.new_vesting == 0) {
                validator_delete(chain_id, acc);
            }
            
            emit ConfirmVestInChain(chain_id, acc, request.new_vesting, request.timestamp);
            vesting_request_delete(chain_id, acc);
            
            return;
        }
        
        // Increases account's vesting in chain
        uint96 to_vest = request.new_vesting - request.old_vesting;
        token.transferFrom(acc, address(this), to_vest);
        
        if (validator.mining == true) {
            chains[chain_id].total_vesting += to_vest;
        }
        
        request.control_state = VestingRequestControl_state.REPLACE_VESTING;
        
        emit ConfirmVestInChain(chain_id, acc, request.new_vesting, request.timestamp);
    }
    
    function _cancel_vest_in_chain(uint chain_id, address acc) private {
        VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
        
        // Replace back the original validator's vesting
        if (request.control_state == VestingRequestControl_state.VESTING_REPLACED) {
            chains[chain_id].users.accounts[acc].validator.vesting = request.old_vesting;
        }
        
        emit CancelVestInChain(chain_id, acc, request.new_vesting, request.timestamp);
        vesting_request_delete(chain_id, acc);
    }
    
    // Forcefully withdraw whole vesting from chain.
    // Because vesting is processed during 2(new_vest < act_vest) or even 3(new_vest > act_vest) notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing vest requests are deleted after this withdraw.
    function _force_withdraw_vest_from_chain(uint chain_id, address acc) private {
        uint96 to_withdraw = 0;
        bool requestExists = vesting_request_exists(chain_id, acc);
        
        // No ongoing vesting request is present
        if (requestExists == false) {
            to_withdraw = chains[chain_id].users.accounts[acc].validator.vesting;
        }
        // There is ongoing vesting request
        else { 
            VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
            // Token transfer was not yet processed -> use saved old vesting balance as actual user's vesting balance to withdraw
            if (request.state == Request_state.REQUEST_CREATED) {
                to_withdraw = request.old_vesting;
            }
            // Token transfer was already processed -> use new vesting balance as actual user's vesting balance to withdraw
            else {
                to_withdraw = request.new_vesting;
            }
            
            vesting_request_delete(chain_id, acc);
        }
        
        // Stops mining
        _stop_mining(chain_id, acc);
        
        validator_delete(chain_id, acc);
        
        // Transfers all remaining tokens to the user account
        token.transfer(acc, to_withdraw);
        
        emit ForceWithdrawVesting(chain_id, acc, now);
    }
    
    function _request_deposit_in_chain(uint chain_id, uint256 deposit, address acc) private {
      uint256 timestamp = now;
      
      // If user wants to withdraw whole deposit - create withdrawal request
      if (deposit == 0) {
        deposit_withdraw_request_create(chain_id, acc);
        emit RequestDepositInChain(chain_id, acc, deposit, timestamp);  
        
        return;
      }
      
      // Internally creates new user
      if (user_exists(chain_id, acc) == false) {
          user_create(chain_id, acc);
      }
      
      // If user wants to deposit in chain, process it immediately
      Transactor storage transactor = chains[chain_id].users.accounts[acc].transactor;
      uint256 act_transactor_deposit = transactor.deposit;
      
      if(act_transactor_deposit > deposit) {
         transactor.deposit = uint96(deposit);
         
         uint256 to_withdraw = act_transactor_deposit - deposit;
         token.transfer(acc, to_withdraw);
      } else {
         uint to_deposit = deposit - act_transactor_deposit;
         token.transferFrom(acc, address(this), to_deposit);
         
         transactor.deposit = uint96(deposit);
      }
      
      emit RequestDepositInChain(chain_id, acc, deposit, timestamp);
      emit ConfirmDepositInChain(chain_id, acc, deposit, timestamp);
      
      if (transactor.whitelisted == false) {
        transactor.whitelisted = true;
        emit WhitelistAccount(chain_id, acc, true);
      }
    }
    
    function _confirm_deposit_withdrawal_from_chain(uint chain_id, address acc) private {
        DepositWithdrawRequest_data storage request = chains[chain_id].requests.accounts[acc].deposit_withdraw_request;
        Transactor storage transactor = chains[chain_id].users.accounts[acc].transactor;
        
        uint256 to_withdraw = transactor.deposit;
        
        deposit_withdraw_request_delete(chain_id, acc);
        transactor_delete(chain_id, acc);
        
        // Withdraw whole deposit
        token.transfer(acc, to_withdraw);
        
        emit ConfirmDepositInChain(chain_id, acc, 0, request.timestamp);
    }
    
    function _cancel_deposit_withdrawal_from_chain(uint chain_id, address acc) private {
        DepositWithdrawRequest_data storage request = chains[chain_id].requests.accounts[acc].deposit_withdraw_request;
        Transactor storage transactor = chains[chain_id].users.accounts[acc].transactor;
        
        // If withdrawal was cancelled and transactor has >= min. required deposit, re-enable user to transact
        if (transactor.whitelisted == false && 
            check_lition_min_deposit(transactor.deposit) == true &&
            chains[chain_id].chain_validator.check_deposit(transactor.deposit, acc) == true) {
                
          transactor.whitelisted = true;
          emit WhitelistAccount(chain_id, acc, true);
        }
        
        deposit_withdraw_request_delete(chain_id, acc);
        
        emit CancelDepositInChain(chain_id, acc, 0, request.timestamp);
    }
    
    // Forcefully withdraw whole deposit.
    // Because deposit withdrawal is processed during 2 notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing deposit requests are deleted after this withdraw.
    function _force_withdraw_deposit_from_chain(uint chain_id, address acc) private {
        Transactor storage transactor = chains[chain_id].users.accounts[acc].transactor; 
        
        uint256 to_withdraw = transactor.deposit;
        transactor_delete(chain_id, acc);
        
        // If deposit withdrawal request exists, delete it
        if (deposit_withdraw_request_exists(chain_id, acc) == true) {
            deposit_withdraw_request_delete(chain_id, acc);    
        }
        
        // Transfers all remaining tokens to the user account
        token.transfer(acc, to_withdraw);
        
        emit ForceWithdrawDeposit(chain_id, acc, now);
    }

    /**************************************************************************************************************************/
    /*************************************************** Other functions ******************************************************/
    /**************************************************************************************************************************/

   constructor(ERC20 _token) public {
      token = _token;
   }
  
  // Process users consumption based on their usage
  function process_users_consumptions(uint256 chain_id, address[] memory users, uint32[] memory user_gas, uint32 largest_tx) internal returns (uint256 total_cost) {
     // Total usage cost in LIT tokens
     total_cost = 0;
     
     // largest tx fee fixed at 0.1 LIT
     uint256 largest_reward = 10**17;
     
     // Individual user's usage cost in LIT tokens
     uint256 user_cost;
     for(uint256 i = 0; i < users.length; i++) {
        address acc = users[i];
        Transactor storage transactor = chains[chain_id].users.accounts[acc].transactor;
        
        user_cost = (user_gas[i] / largest_tx) * largest_reward;
        
        // This should never happen as it is handled by 2-step deposit withdrawal system and
        // by checking user's deposit balance is >= min. required deposit conditions
        if(user_cost > transactor.deposit ) {
           user_cost = transactor.deposit;
           
           transactor.whitelisted = false;
           emit WhitelistAccount(chain_id, users[i], false);
        }
        
        // Updates user's current deposit balance based on his usage
        transactor.deposit -= uint96(user_cost);
        
        // Check if user's deposit balance is >= min. required deposit conditions
        if (check_lition_min_deposit(transactor.deposit) == false || chains[chain_id].chain_validator.check_deposit(transactor.deposit, acc) == false) {
            // If not, do not allow him to transact anymore
            transactor.whitelisted = false;
            emit WhitelistAccount(chain_id, acc, false);
        }
        
        // Adds user's cost to the total cost
        total_cost += user_cost;
     }
   }

   // Process miners rewards based on their participation rate(how many blocks they signed) and their vesting balance
   function process_miners_rewards(uint256 chain_id, uint256 notary_block_no, address[] memory miners, uint32[] memory blocks_mined, uint256 lit_to_distribute) internal {
     ChainInfo storage chain = chains[chain_id];
     
     // Min. vesting balance to be a trust node. Trust Nodes haved doubled(virtually) vesting
     uint256 min_trust_node_vesting = 50000*(10**18); 
     
     // How many block could validator mined since the last notary in case he did sign every possible block 
     uint256 max_blocks_mined = notary_block_no - chain.last_notary.block;

     // Total involved vesting 
     uint256 total_involved_vesting = 0;
     
     // Selected validator's vesting balance
     uint256 validator_vesting;
     
     // Runs through all miners and calculates total involved vesting based on:
     for(uint256 i = 0; i < miners.length; i++) {
        validator_vesting = chain.users.accounts[miners[i]].validator.vesting;
        
        // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
        if (validator_vesting >= min_trust_node_vesting) {
            validator_vesting *= 2;
        }

        total_involved_vesting += (max_blocks_mined / blocks_mined[i]) * validator_vesting;
     }

     
     // Runs through all miners and calculates their reward based on:
     //     1. How many blocks out of max_blocks_moned each miner signed
     //     2. How many token each miner vested
     for(uint256 i = 0; i < miners.length; i++) {
        validator_vesting = chain.users.accounts[miners[i]].validator.vesting;
        
        // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
        if (validator_vesting >= min_trust_node_vesting) {
            validator_vesting *= 2;
        }
        
        uint256 miner_reward = (max_blocks_mined / blocks_mined[i]) * (validator_vesting / total_involved_vesting) * lit_to_distribute;
        token.transfer(miners[i], miner_reward);
        
        lit_to_distribute -= miner_reward;
     }

     // Sends the rest(math rounding) to the miner, who called notary function
     token.transfer(msg.sender, lit_to_distribute);
   }

    // Process existing vesting/deposit withdrawals requests as these cannot be 
    // processed automatically - math for user's usage and miner's rewards calculations would be invalid
    function process_existing_requests(uint256 chain_id) internal {
        ChainInfo storage chain = chains[chain_id];
        Requests storage requests = chain.requests;
        
        // Runs through all existing requests
        for (uint256 i = 0; i < requests.list.length; i++) {
            address acc = requests.list[i];
            Requests_entry storage entry = requests.accounts[acc];
            User_entry storage user = chain.users.accounts[acc];
            
            // There is existing vesting request - process it 
            if (entry.vesting_request.timestamp != 0) {
                // Current validator's vesting balance should be replaced during this notary
                if (entry.vesting_request.control_state == VestingRequestControl_state.REPLACE_VESTING) {
                    user.validator.vesting = entry.vesting_request.new_vesting;
                    entry.vesting_request.control_state = VestingRequestControl_state.VESTING_REPLACED;
                    
                    emit AcceptedVestInChain(chain_id, acc, entry.vesting_request.new_vesting, entry.vesting_request.timestamp);
                    
                    // If it was request to increase validator's vesting balance and we got here, it means we can delete this request
                    // request to decrease vestin balance are deleted in confirmation
                    if (entry.vesting_request.new_vesting > entry.vesting_request.old_vesting) {
                        vesting_request_delete(chain_id, acc);
                    }
                }
            }
            
            // There is existing deposit withdrawal request - process it
            // If there is such request and user is still allowed to transact(might be not because he runs out of tokens during usage calculations), 
            // do not allow him to transact anymore
            if (entry.deposit_withdraw_request.timestamp != 0 && user.transactor.whitelisted == true) {
                // User wants to withdraw whole deposit - do not allow him to transact anymore, token transfer is processed when he confirms withdrawal
                user.transactor.whitelisted = false;
                emit WhitelistAccount(chain_id, acc, false);
            }
        }
    }

  function _get_users(uint chain_id, uint batch, bool validators, bool active) private view returns (address[100] memory users, uint count, bool end) {
     count = 0;
     uint256 j = batch * 100;
     uint256 users_total_count = chains[chain_id].users.list.length;
     
     while(j < (batch + 1)*100 && j < users_total_count) {
      address acc = chains[chain_id].users.list[j];
      // Get validators
      if(validators == true) {
        // if active == true, get only those validators who are also mining
        // if active == false, get those who are allowed to mine based on their vesting
        if (chains[chain_id].users.accounts[acc].validator.mining == active) {
          users[count] = acc;
          count++;
        }
      }
      // Get transactors
      else {
        if (chains[chain_id].users.accounts[acc].transactor.whitelisted == true) {
          users[count] = acc;
          count++;
        } 
      }
      j++;
     }
     
     if (j+1 == users_total_count) {
         end = true;
     }
     else {
         end = false;
     }
  }
      
  function _start_mining(uint chain_id, address acc) private {      
      Validator storage validator = chains[chain_id].users.accounts[acc].validator;
      if (validator.mining == false) {
          chains[chain_id].total_vesting += validator.vesting;
      }
      validator.mining = true;
      
      emit StartMining(chain_id, acc);
  }
      
  function _stop_mining(uint chain_id, address acc) private {      
      Validator storage validator = chains[chain_id].users.accounts[acc].validator;
      if (validator.mining == true) {
          chains[chain_id].total_vesting -= validator.vesting;
      }
      validator.mining = false;
      
      emit StopMining(chain_id, acc);
  }
}