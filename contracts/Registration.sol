
pragma solidity >=0.5.4;
//pragma experimental ABIEncoderV2;

interface ChainValidator {
   function check_vesting(uint vesting, address participant) external returns (bool);
   function check_deposit(uint vesting, address participant) external returns (bool);
}

contract LitionChainValidator is ChainValidator {
   function check_vesting(uint vesting, address participant) public returns (bool) {
      if(vesting >= 1000*(uint256(10)**uint256(18)) && vesting <= 500000*(uint256(10)**uint256(18))) {
        return true;   
      }
      return false;
   }

   function check_deposit(uint deposit, address participant) public returns (bool) {
      if(deposit >= 10*(uint256(10)**uint256(18))) {
         return true;
      }
      return false;
   }
}

interface ERC20{
   function totalSupply() external view returns (uint);
   function balanceOf(address tokenOwner) external view returns (uint balance);
   function allowance(address tokenOwner, address spender) external view returns (uint remaining);
   function transfer(address to, uint tokens) external returns (bool success);
   function approve(address spender, uint tokens) external returns (bool success);
   function transferFrom(address from, address to, uint tokens) external returns (bool success);
   event Transfer(address indexed from, address indexed to, uint tokens);
   event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract LitionRegistry{
   event NewChain(uint id, string description);
   //event NewChainEndpoint(uint id, string endpoint);
   event Deposit(uint indexed chain_id, uint deposit, address indexed depositer, uint256 datetime);
   event Vesting(uint indexed chain_id, uint deposit, address indexed depositer, uint256 datetime);
   
   
   event RequestVestInChain(uint indexed chain_id, address indexed account, uint96 vesting, uint req_timestamp /* now */);
   //TODO: All events connected to the vesting_request(except RequestVestInChain), might trhow just req_timestamp as id of the original request and timestamp as actual time
   event ConfirmVestInChain(uint indexed chain_id, address indexed account, uint96 vesting, uint req_timestamp /* when was the request created/it's id */, uint timestamp /* now */);
   event CancelVestInChain(uint indexed chain_id, address indexed account, uint96 vesting, uint req_timestamp /* when was the request created/it's id */, uint timestamp /* now */);
   event FinishedVestInChain(uint indexed chain_id, address indexed account, uint96 vesting, uint req_timestamp /* when was the request created/it's id */, uint timestamp /* now */);
   event ForceWithdrawVesting(uint indexed chain_id, address indexed account, uint timestamp /* now */);
  
   
   event StartMining(uint indexed chain_id, address miner);
   event StopMining(uint indexed chain_id, address miner);


    /**************************************************************************************************************************/
    /*********************************** Structs & Functions related to the list of users *************************************/
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
    
    // Creates new user - does not set it's data yet as it is done after vesting/deposit_withdraw request is confirmed
    function user_create(uint256 chain_id, address acc) private {
        require(chains[chain_id].users.list.length < ~uint48(0), "count(users) is equal to max_count");
        require(user_exists(chain_id, acc) == false, "Creating already-created user");
        
        chains[chain_id].users.list.push(acc);
        chains[chain_id].users.accounts[acc].index = uint48(chains[chain_id].users.list.length); // indexes are stored + 1
        
        // TODO: move this functionality to the requests confirmation
        // // TODO: check if ~uint96(0) is cheaper to have as constant
        // require(vesting <= ~uint96(0), "vesting is greater than uint96_max_value");
        // require(deposit <= ~uint96(0), "deposit is greater than uint96_max_value");
        
        // address[] storage users_list = chains[chain_id].users.list; 
        // require(users_list.length < ~uint48(0), "count(users) is equal to max_count");
        
        // // Push account to the list of accounts
        // users_list.push(acc);
        
        // User_entry storage new_user = chains[chain_id].users.accounts[acc];
        
        // // Default vesting values are vesting = 0, mining = false (it has to be set in separate call to start/stop mining); no need to expliciyl set them
        // if (vesting > 0) {
        //     new_user.info.vesting = uint96(vesting);
        // }

        // // Default deposit values are vesting = 0, whitelisted = false; no need to expliciyl set them
        // if (deposit > 0) {
        //     new_user.info.deposit = uint96(deposit);
        //     // In case deposit > 0, min. required deposit conditions have already been checked - allow user to transact
        //     new_user.info.whitelisted = true;    
        // }
        
        // new_user.index = uint48(users_list.length); // indexes are stored + 1
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
    /******************************* End of Structs & Functions related to the list of users **********************************/
    /**************************************************************************************************************************/


    /**************************************************************************************************************************/
    /****************************** Structs & Functions related to the vesting/deposit requests *******************************/
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
    
    struct VestingRequest_data {
        // Last notary block number when the request was accepted 
        uint256                         notary_block;
        // Timestamp(now) when the request was accepted/created
        uint48                          timestamp;
        // In old_vesting is stored actual vesting that validator had when new VestingRequest was accepted
        uint96                          old_vesting;
        // New value of vesting to be set
        uint96                          new_vesting;
        // Actual state of the request
        Request_state                   state;
        // Actual control state of the request
        VestingRequestControl_state     control_state;
    }
    
    // Only full deposit withdrawals are saved as deposit requests - other types of deposits do not need to be confirmed
    struct DeposiWithdrawtRequest_data {
        // Last notary block number when the request was accepted 
        uint256                         notary_block;
        // Timestamp(now) when the request was accepted/created
        uint48                          timestamp;
        // Actual state of the request
        Request_state                   state;
    }
    
    struct Requests_entry {
        // index to the requests_list, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing request
        uint48                          index;
        // Deposit withdrawal request details
        DeposiWithdrawtRequest_data     deposit_withdraw_request;
        // Vesting request details
        VestingRequest_data             vesting_request;   
    }
    
    struct Requests {
        mapping(address => Requests_entry)  accounts;
        address[]                           list;        
    }
    
    // Creates new vesting request
    function vesting_request_create(uint256 chain_id, address acc, uint256 vesting) private {
        require(chains[chain_id].requests.list.length < ~uint48(0), "count(requests) is equal to max_count");
        require(vesting <= ~uint96(0), "request: vesting is greater than uint96_max_value");
        
        // TODO: should already-existing vesting request be checked here or in caller's function ???
        require(vesting_request_exists(chain_id, acc) == false, "There is already ongoing vesting request");
        
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
        
        // TODO: should already-existing deposit withdrawal request be checked here or in caller's function ???
        require(deposit_withdraw_request_exists(chain_id, acc) == false, "There is already ongoing deposit withdrawal request");
        
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
        DeposiWithdrawtRequest_data storage request = chains[chain_id].requests.accounts[acc].deposit_withdraw_request;
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
    
    /********************** Public functions visible to the callers **********************************/
    
    // Requests vest in chain. It will be processed and applied to the actual user state after next:
    //      * 2 notary windows - in case new vesting < actual vesting
    //      * 3 notary windows - in case new vesting > actual vesting
    function request_vest_in_chain(uint chain_id, uint96 vesting) external {
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
    
    /****************** End of Public functions visible to the callers *******************************/
    
    
    function _request_vest_in_chain(uint chain_id, uint96 vesting, address acc) private {
      // Internally creates new user
      if (vesting != 0 && user_exists(chain_id, acc) == false) {
          user_create(chain_id, acc);
      }
      
      vesting_request_create(chain_id, acc, vesting);
      emit RequestVestInChain(chain_id, msg.sender, vesting, now);
    }
    
    function _confirm_vest_in_chain(uint chain_id, address acc) private {
        VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
        Validator storage validator = chains[chain_id].users.accounts[acc].validator;
        
        // Decreases account's vesting in chain
        if(request.new_vesting < validator.vesting) {
            uint96 to_withdraw = validator.vesting - request.new_vesting;
            token.transfer(acc, to_withdraw);
            
            if (validator.mining == true) {
                chains[chain_id].total_vesting -= to_withdraw; //TODO -= safe math here;
            }
            
            emit ConfirmVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].timestamp, now);
            emit FinishedVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].timestamp, now);
            vesting_request_delete(chain_id, acc);
            return;
        }
        
        // Increases account's vesting in chain
        uint96 to_vest = request.new_vesting - validator.info.vesting;
        token.transferFrom(acc, address(this), to_vest);
        
        if (validator.info.mining == true) {
            chains[chain_id].total_vesting += to_vest;
        }
        
        request.control_state  = VestingRequestControl_state.REPLACE_VESTING;
        request.state          = Request_state.REQUEST_CONFIRMED;
        emit ConfirmVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].timestamp, now);
    }
    
    function _cancel_vest_in_chain(uint chain_id, address acc) private {
        VestingRequest_data storage request = chains[chain_id].requests.accounts[acc].vesting_request;
        
        // Replace back the original validator's vesting
        if (request.control_state == VestingRequestControl_state.VESTING_REPLACED) {
            chains[chain_id].users.accounts[acc].validator.info.vesting = request.old_vesting;
        }
        
        emit CancelVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].timestamp, now);
        vesting_request_delete(chain_id, acc);
    }
    
    // Forcefully withdraw all vesting from chain.
    // Because vesting is processed during 2(new_vest < act_vest) or even 3(new_vest > act_vest) notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing vest requests are deleted after this withdraw.
    function _force_withdraw_vest_from_chain(uint chain_id, address acc) private {
        uint96 to_withdraw = 0;
        bool requestExists = vesting_request_exists(chain_id, acc);
        
        // No ongoing vesting request is present
        if (requestExists == false) {
            to_withdraw = chains[chain_id].users.accounts[acc].validator.info.vesting;
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
        }
        
        // Stops mining
        _stop_mining(chain_id, acc);
        
        // Transfers all remaining tokens to the user account
        token.transfer(acc, to_withdraw);
        chains[chain_id].users.accounts[acc].validator.info.vesting = 0;
        
        // If vesting request exists, delete it
        if (requestExists == true) {
            vesting_request_delete(chain_id, acc);    
        }
    }

    /**************************************************************************************************************************/
    /************************** End of Structs & Functions related to the vesting/deppsit requests ****************************/
    /**************************************************************************************************************************/

   ERC20 token;
   
   struct LastNotary {
       // Timestamp, when the last notary was accepted
       uint timestamp;
       // Actual block number, when the last notary was accepted
       uint block;
   }

   struct chain_info{
      bool              active;
      uint96            total_vesting;
      LastNotary        last_notary;
      ChainValidator    chain_validator;
      Users             users;
      Requests          requests;
      string            endpoint;
   }
   
   struct signature {
      uint8 v; bytes32 r; bytes32 s;
   }

   mapping(uint256 => chain_info) private chains;
   uint256 public next_id = 0;

   constructor(ERC20 _token) public {
      token = _token;
   }
   
  // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
  function check_lition_min_vesting(uint vesting) private pure returns (bool) {
      if(vesting >= 1000*(uint256(10)**uint256(18))) {
        return true;   
      }
      return false;
  }
   
  // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
  function check_lition_min_deposit(uint deposit) private pure returns (bool) {
      if(deposit >= 1000*(uint256(10)**uint256(18))) {
        return true;   
      }
      return false; 
  }

  // Internally creates/registers new side-chain 
  function register_chain(string calldata info, ChainValidator validator, uint96 vesting, uint96 deposit, string calldata init_endpoint) external returns (uint256 chain_id) {
      require(bytes(init_endpoint).length > 0);
      
      // Validates vesting
      require(check_lition_min_vesting(vesting), "chain creator does not meet Lition's min.required vesting condition");
      require(chains[chain_id].chain_validator.check_vesting(vesting, msg.sender), "chain creator does not meet chain validator's min.required vesting condition");
      
      // Validates deposit
      require(check_lition_min_deposit(deposit), "chain creator does not meet Lition's min.required deposit condition");
      require(chains[chain_id].chain_validator.check_deposit(deposit, msg.sender), "chain creator does not meet chain validator's min.required deposit condition");
      
      // Inits chain data
      chains[chain_id].chain_validator          = validator;
      chains[chain_id].active                   = true;
      chains[chain_id].endpoint                 = init_endpoint;
      // Use default value (0) for timestamp and block
      
      chain_id                                  = next_id;
      
      // Internally creates new user
      user_create(chain_id, msg.sender);
      
      // Transfers vesting tokens
      token.transferFrom(msg.sender, address(this), vesting);
      chains[chain_id].users.accounts[msg.sender].validator.vesting = vesting;      
      // TODO: emit events about vesting request/confirmation/finish          
      
      // Transfers deposit tokens
      token.transferFrom(msg.sender, address(this), deposit);
      chains[chain_id].users.accounts[msg.sender].transactor.deposit = deposit;      
      // TODO: emit events about deposit request/confirmation        
      
      emit NewChain(chain_id, info);
      //emit NewChainEndpoint( id, init_endpoint );
      
      next_id++;
  }
   
//   function deposit_in_chain( uint id, uint deposit ) public {
//       _deposit_in_chain(id, deposit, msg.sender );
//   }

  function has_vested(uint chain_id, address acc) view external returns (bool) {
      return validator_exists(chain_id, acc);
  }

//   function has_deposited(uint id, address user) view external returns (bool) {
//       return chains[id].users[user].info.deposit > 0;
//   }

//   function get_signature_hash_from_notary(uint32 notary_block, address[] memory miners,
//                                  uint32[] memory blocks_mined, address[] memory users,
//                                  uint32[] memory user_gas, uint32 largest_tx)
//                                      public pure returns (bytes32) {
//       return keccak256(abi.encodePacked(notary_block, miners, blocks_mined, users, user_gas, largest_tx));
//   }

//   function get_last_notary(uint id) external view returns (uint256) {
//      return chains[id].last_notary;
//   }

//   function process_users_consumptions(uint id, address[] memory users, uint32[] memory user_gas, uint32 largest_tx) internal returns (uint256 total_cost) {
//      uint total_gas = 0;
//      total_cost = 0;
//      //largest tx fixed at 0.1 LIT - rework that to work with current price
//      uint largest_reward = 10**17;

//      for(uint i = 0; i < users.length; i++) {
//         total_gas +=user_gas[i];
//         uint user_cost = (user_gas[i] / largest_tx) * largest_reward;
//         if( user_cost > chains[id].users[users[i]].info.deposit ) {
//           user_cost = chains[id].users[users[i]].info.deposit;
//           emit Deposit(id, 0, users[i], now);
//         }
//         chains[id].users[users[i]].info.deposit -= user_cost;
//         total_cost += user_cost;
//      }
//   }

//   function process_miners_rewards(uint id, address[] memory miners, uint32[] memory blocks_mined, uint lit_to_distribute) internal {
//      uint total_signatures = 0;
//      for(uint i = 0; i < miners.length - 1; i++) {
//         total_signatures += blocks_mined[i];
//      }

//      for(uint i = 0; i < miners.length - 1; i++) {
//         uint miner_reward = blocks_mined[i] * lit_to_distribute / total_signatures;
//         token.transfer( miners[i], miner_reward );
//         lit_to_distribute -= miner_reward;
//      }

//      token.transfer( miners[miners.length - 1], lit_to_distribute );
//   }

//   function notary(uint id, uint32 notary_block_no, address[] memory miners, uint32[] memory blocks_mined,
//                                  address[] memory users, uint32[] memory user_gas, uint32 largest_tx,
//                                  uint8[] memory v, bytes32[] memory r, bytes32[] memory s) public {
//       //first, calculate hash from miners, block_mined, users and user_gas
//       //then, do ec_recover of the signatures to determine signers
//       //check if there is enough signers (total vesting of signers > 50% of all vestings)
//       //then, calculate reward
//       require(v.length == r.length);
//       require(v.length == s.length);
//       bytes32 signature_hash = get_signature_hash_from_notary(notary_block_no, miners, blocks_mined, users, user_gas, largest_tx);
//       chain_info storage chain = chains[id];
//       require(chain.active, "Trying to report about non-existing chain");

//       uint involved_vesting = 0;

//       for(uint i =0; i<v.length; i++) {
//          address signer = ecrecover(signature_hash, v[i], r[i], s[i]);
//          involved_vesting += chain.users[signer].info.vesting;
//       }

//       require(involved_vesting * 2 >= chain.total_vesting);

//       uint256 total_cost = process_users_consumptions(id, users, user_gas, largest_tx);
//       process_miners_rewards(id, miners, blocks_mined, total_cost);
//       chain.last_notary = notary_block_no;
//   }

//   //TODO - rework so withdrawals are not processed immediatelly but after notary window
//   function _deposit_in_chain( uint id, uint deposit, address user ) private {
//       //Validate value of deposit
//       if (deposit == 0) {
//         require( chains[id].users[user].info.deposit > 0, "Zero deposit balance. Can't withdraw any tokens" );  
//       }
//       else {
//         require( chains[id].active, "can't deposit into non-existing chain" );
//         require( check_lition_min_deposit( deposit), "user does not meet min. required chain criteria");
//         require( chains[id].validator.check_deposit( deposit, user ), "user does not meet chain criteria");
//       }
      
//       if( chains[id].users[user].info.deposit > deposit ){
//          uint to_withdraw = chains[id].users[user].info.deposit - deposit;
//          token.transfer( user, to_withdraw);
//       } else{
//          uint to_deposit = deposit - chains[id].users[user].info.deposit;
//          token.transferFrom(user, address(this), to_deposit);
//       }
      
//       chains[id].users[user].info.deposit = deposit;
//       users_list_add(id, user);
//       emit Deposit(id, deposit, user, now);
//   }

//   function get_allowed_to_transact( uint id, uint batch ) view external returns (address[100] memory users, uint count) {
//      count = 0;
//      uint j = batch * 100;
//      while( j < (batch + 1)*100 && j < chains[id].users_list.length ) {
//       address user = chains[id].users_list[j];
//       if(chains[id].users[user].info.deposit > 0) {
//          users[count] = user;
//          count++;
//       }
//       j++;
//      }
//   }

  // Returns list of validator's addresses that are actively participating in mining
  function get_allowed_to_validate(uint chain_id, uint batch) view external returns (address[100] memory, uint) {
     return _get_validators(chain_id, batch, false);
  }

  // Returns list of validator's addresses that are allowed to be participating in mining based on their vesting, but not yet mining
  function get_active_validators(uint chain_id, uint batch) view external returns (address[100] memory, uint) {
     return _get_validators(chain_id, batch, true);
  }
  
  function _get_validators(uint chain_id, uint batch, bool active) private view returns (address[100] memory validators, uint count) {
     count = 0;
     uint j = batch * 100;
     while( j < (batch + 1)*100 && j < chains[chain_id].validators.list.length ) {
      address acc = chains[chain_id].validators.list[j];
      if(chains[chain_id].validators.accounts[acc].info.mining == active) {
         validators[count] = acc;
         count++;
      }
      j++;
     }
  }
  
  function start_mining(uint chain_id) external {
      require(chains[chain_id].active == true, "Can't start mining on non-existing chain");
      require(check_lition_min_vesting(chains[chain_id].validators.accounts[msg.sender].info.vesting) == true, "user does not meet Lition's min.required vesting condition");
      require(chains[chain_id].chain_validator.check_vesting(chains[chain_id].validators.accounts[msg.sender].info.vesting, msg.sender) == true, "User does not meet chain validator's min.required vesting condition");
      
      _start_mining(chain_id, msg.sender);
  }
      
  function _start_mining(uint chain_id, address acc) private {      
      if (chains[chain_id].validators.accounts[acc].info.mining == false) {
          chains[chain_id].total_vesting += chains[chain_id].validators.accounts[acc].info.vesting;
      }
      chains[chain_id].validators.accounts[acc].info.mining = true;
      
      emit StartMining(chain_id, acc);
  }

  function stop_mining(uint chain_id) external {
      require(chains[chain_id].active == true, "Can't start mining on non-existing chain");
      require(check_lition_min_vesting( chains[chain_id].validators.accounts[msg.sender].info.vesting) == true, "user does not meet Lition's min.required vesting condition");
      
      _stop_mining(chain_id, msg.sender);
  }
      
  function _stop_mining(uint chain_id, address acc) private {      
      if (chains[chain_id].validators.accounts[acc].info.mining == true) {
          chains[chain_id].total_vesting -= chains[chain_id].validators.accounts[acc].info.vesting;
      }
      chains[chain_id].validators.accounts[acc].info.mining = false;
      
      emit StopMining(chain_id, acc);
  }

}