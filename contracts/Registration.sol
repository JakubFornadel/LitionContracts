
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
   event NewChainEndpoint(uint id, string endpoint);
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
    /********************************* Structs & Functions related to the list of validators **********************************/
    /**************************************************************************************************************************/
    struct Validator_details {
        // Actual validator vesting
        uint96  vesting;
        // Flag that is set in start/stop_mining
        bool    mining;
    }
    
    struct Validator_entry {
        // index to the validators_list, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing validator
        uint                index;
        // Validator details
        Validator_details   info;   
    }
    
    struct Validators {
        mapping(address => Validator_entry) accounts;
        address[]                           list;        
    }
    
    function validator_create(uint chain_id, address acc, uint96 vesting) internal {
        require(validator_exists(chain_id, acc) == false, 'validator_create: Validator already exists.');
        chains[chain_id].validators.list.push(acc);
    
        Validator_entry storage entry = chains[chain_id].validators.accounts[acc];
        entry.info.vesting = vesting;
        entry.info.mining = false;
        entry.index = chains[chain_id].users_list.length; // indexes are stored + 1
    }
    
    function validator_delete(uint chain_id, address acc) internal {
        require(validator_exists(chain_id, acc) == true, 'validator_delete: Validator does not exist.');
        address[] storage validators_list = chains[chain_id].validators.list;
    
        Validator_entry storage entry = chains[chain_id].validators.accounts[acc];
        require(entry.index <= validators_list.length, 'validator_delete: Invalid index value.');
    
        // Move an last element of array into the vacated key slot.
        uint found_index = entry.index - 1;
        uint last_index = validators_list.length - 1;
    
        chains[chain_id].validators.accounts[validators_list[last_index]].index = found_index + 1;
        validators_list[found_index] = validators_list[last_index];
        validators_list.length--;
    
        delete chains[chain_id].validators.accounts[acc];
    }
    
    function validator_exists(uint chain_id, address acc) internal view returns (bool) {
      return chains[chain_id].validators.accounts[acc].index != 0;
    }
    /**************************************************************************************************************************/
    /***************************** End of Structs & Functions related to the list of validators *******************************/
    /**************************************************************************************************************************/
   

    /**************************************************************************************************************************/
    /********************************** Structs & Functions related to the vesting requests ***********************************/
    /**************************************************************************************************************************/
    enum VestingRequest_state {
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
        // Timestamp when the request was accepted
        uint                           timestamp;
        // In old_vesting is stored actual vesting that validator had when new VestingRequest was accepted
        uint96                          old_vesting;
        // New value of vesting to be set
        uint96                          new_vesting;
        // Actual state of the request
        VestingRequest_state            state;
        // Actual control state of the request
        VestingRequestControl_state     control_state;
    }
    
    struct VestingRequest_entry {
        // index to the validators_list, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing validator
        uint                  index;
        // Validator details
        VestingRequest_data   data;   
    }
    
    struct VestingRequests {
        mapping(address => VestingRequest_entry) accounts;
        address[]                                list;        
    }
    
    // Creates new vesting request and stores it in the requests list
    // TODO: handle all requires, etc... in the vest_in_chain function !!!
    function vesting_request_create(uint chain_id, address acc, uint96 vesting) internal {
        require(vesting_request_exists(chain_id, acc) == false, 'vesting_request_create: vesting_request already exists.');
        chains[chain_id].vesting_requests.list.push(acc);
    
        VestingRequest_entry storage entry = chains[chain_id].vesting_requests.accounts[acc];
        entry.data.timestamp = now;
        entry.data.old_vesting = chains[chain_id].validators.accounts[acc].info.vesting;
        entry.data.new_vesting = vesting;
        if (entry.data.new_vesting >= entry.data.old_vesting) { // == case should never happen as it is handled in the caller's function
            entry.data.control_state = VestingRequestControl_state.WAIT_FOR_CONFIRMATION;
        } else {
            entry.data.control_state = VestingRequestControl_state.REPLACE_VESTING;
        }
        
        entry.index = chains[chain_id].users_list.length; // indexes are stored + 1
    }
    
    // Deletes existing vesting request from the internal list of requests
    function vesting_request_delete(uint chain_id, address acc) internal {
        require(vesting_request_exists(chain_id, acc) == true, 'vesting_request_delete: vesting_request does not exist.');
        address[] storage vesting_requests_list = chains[chain_id].vesting_requests.list;
    
        VestingRequest_entry storage entry = chains[chain_id].vesting_requests.accounts[acc];
        require(entry.index <= vesting_requests_list.length, 'vesting_request_delete: Invalid index value.');
    
        // Move an last element of array into the vacated key slot.
        uint found_index = entry.index - 1;
        uint last_index = vesting_requests_list.length - 1;
    
        chains[chain_id].vesting_requests.accounts[vesting_requests_list[last_index]].index = found_index + 1;
        vesting_requests_list[found_index] = vesting_requests_list[last_index];
        vesting_requests_list.length--;
    
        delete chains[chain_id].vesting_requests.accounts[acc];
    }
    
    // Checks if acc has any ongoing vesting request
    function vesting_request_exists(uint chain_id, address acc) internal view returns (bool) {
      return chains[chain_id].vesting_requests.accounts[acc].index != 0;
    }
    
    /********************** Public functions visible to the callers **********************************/
    
    // Requests vest in chain. It will be processed and applied to the actual user state after next:
    //      * 2 notary windows - in case new vesting < actual vesting
    //      * 3 notary windows - in case new vesting > actual vesting
    function request_vest_in_chain(uint chain_id, uint96 vesting) public {
      _request_vest_in_chain( chain_id, vesting, msg.sender );
    }
    
    // Confirms vest request, token transfer is processed during confirmation
    function confirm_vest_in_chain(uint chain_id) public {
      _confirm_vest_in_chain(chain_id, msg.sender);
    }
    
    // Cancels the existing vest request. Such request can be cancelled only if it was not already confirmed
    function cancel_vest_in_chain(uint chain_id) public {
      _cancel_vest_in_chain(chain_id, msg.sender);
    }
    
    // Forcefully withdraw all vesting from chain.
    // Because vesting is processed during 2(new_vest < act_vest) or even 3(new_vest > act_vest) notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing vest requests are deleted after this withdraw.
    function force_withdraw_vest_from_chain( uint id ) public {
      //_confirm_vest_in_chain( id, vesting, msg.sender );
      // TODO: important require that request.timestamp - last_notary.timestamp >= 1 month 
    }
    
    /****************** End of Public functions visible to the callers *******************************/
    
    
    function _request_vest_in_chain( uint chain_id, uint96 vesting, address acc ) private {
      require( vesting_request_exists(chain_id, acc) == false, "Cannot vest in chain. There is already ongoing request being processed for this acc." );
      require( chains[chain_id].validators.accounts[acc].info.vesting != vesting, "Cannot vest the same amount of tokens as you already has vested." );
      
      bool existing_validator = validator_exists(chain_id, acc);
       
      // Withdraw all vesting and delete validator account
      if (vesting == 0) {
        require( existing_validator == true, "Trying to withdraw vesting from non-existing validator account" );
        require( chains[chain_id].validators.accounts[acc].info.mining == false, "Can't withdraw any tokens, stop_minig must be called first." );  
      }
      // Vest in chain or withdraw just part of vesting
      else {
         require( chains[chain_id].active, "can't vest into non-existing chain" );
         require( check_lition_min_vesting( vesting ), "user does not meet min. required chain criteria");
         require( chains[chain_id].validator.check_vesting( vesting, acc ), "user does not meet chain criteria");
         
         if (existing_validator == false) {
             validator_create(chain_id, acc, vesting);
         }
      }
      
      vesting_request_create(chain_id, acc, vesting);
      chains[chain_id].vesting_requests.accounts[acc].data.state = VestingRequest_state.REQUEST_CREATED;
      emit RequestVestInChain(chain_id, msg.sender, vesting, now);
    }
    
    function _confirm_vest_in_chain(uint chain_id, address acc) private {
        require(vesting_request_exists(chain_id, acc) == true, "Cannot confirm non-existing vesting request.");
        VestingRequest_entry storage request = chains[chain_id].vesting_requests.accounts[acc];
        require(request.data.state == VestingRequest_state.REQUEST_CREATED, "Cannot confirm already confirmed request." );
        // TODO: require that confirm was called in the next notary window after request
        
        
        // Decreases account's vesting in chain
        if(request.data.new_vesting < chains[chain_id].validators.accounts[acc].info.vesting) {
            uint96 to_withdraw = chains[chain_id].validators.accounts[acc].info.vesting - request.data.new_vesting;
            token.transfer(acc, to_withdraw);
            
            if (chains[chain_id].validators.accounts[acc].info.mining == true) {
                chains[chain_id].total_vesting -= to_withdraw; //TODO -= safe math here;
            }
            
            emit ConfirmVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].data.new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].data.timestamp, now);
            emit FinishedVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].data.new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].data.timestamp, now);
            vesting_request_delete(chain_id, acc);
            return;
        }
        
        // Increases account's vesting in chain
        uint96 to_vest = request.data.new_vesting - chains[chain_id].validators.accounts[acc].info.vesting;
        token.transferFrom( acc, address(this), to_vest);
        
        if (chains[chain_id].validators.accounts[acc].info.mining == true) {
            chains[chain_id].total_vesting += to_vest;
        }
        
        request.data.control_state = VestingRequestControl_state.REPLACE_VESTING;
        request.data.state = VestingRequest_state.REQUEST_CONFIRMED;
        emit ConfirmVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].data.new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].data.timestamp, now);
    }
    
    function _cancel_vest_in_chain(uint chain_id, address acc) private {
        require(vesting_request_exists(chain_id, acc) == true, "Cannot cancel non-existing vesting request.");
        VestingRequest_entry storage request = chains[chain_id].vesting_requests.accounts[acc];
        require(request.data.state == VestingRequest_state.REQUEST_CREATED, "Cannot cancel already confirmed request." );
        
        // Replace back the original validator's vesting
        if (request.data.control_state == VestingRequestControl_state.VESTING_REPLACED) {
            chains[chain_id].validators.accounts[acc].info.vesting = request.data.old_vesting;
        }
        
        emit CancelVestInChain(chain_id, msg.sender, chains[chain_id].vesting_requests.accounts[msg.sender].data.new_vesting, chains[chain_id].vesting_requests.accounts[msg.sender].data.timestamp, now);
        vesting_request_delete(chain_id, acc);
    }

    /**************************************************************************************************************************/
    /****************************** End of Structs & Functions related to the vesting requests ********************************/
    /**************************************************************************************************************************/

   ERC20 token;
   struct user_details{
      bool mining;
      uint vesting;
      uint deposit;
      string endpoint;
   }
   struct user_entry {
     uint index; // index start 1 to users_list.length
     user_details info;
   }

   struct chain_info{
      bool active;
      mapping(address => user_entry) users;
      address[]                      users_list;
      uint last_notary;
      uint last_notary_timestamp;
      ChainValidator validator;
      uint96 total_vesting;
      Validators validators;
      VestingRequests vesting_requests;
   }
   
   struct signature {
      uint8 v; bytes32 r; bytes32 s;
   }

   mapping(uint256 => chain_info) private chains;
   uint256 public next_id = 0;

   constructor(ERC20 _token) public {
      token = _token;
   }
   
//   function users_list_add(uint chain_id, address user) internal {
//      // User is already in list, do nothing
//      if (chains[chain_id].users[user].index != 0) {
//         return;
//      }
    
//      chains[chain_id].users_list.push(user);
//      chains[chain_id].users[user].index = chains[chain_id].users_list.length; // indexes are stored + 1
//   }
   
  // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
  function check_lition_min_vesting(uint vesting) internal pure returns (bool) {
      if(vesting >= 10*(uint256(10)**uint256(18))) {
        return true;   
      }
      return false;
  }
   
  // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
  function check_lition_min_deposit(uint deposit) internal pure returns (bool) {
      if(deposit >= 1*(uint256(10)**uint256(18))) {
        return true;   
      }
      return false; 
  }

//   function register_chain( string calldata info, ChainValidator validator, uint vesting, string calldata init_endpoint ) external returns ( uint256 id ){
//       require( bytes(init_endpoint).length > 0 );
//       id = next_id;
//       chains[id].validator = validator;
//       chains[id].active = true;
//       chains[id].last_notary = 0;
//       chains[id].users[msg.sender].info.endpoint = init_endpoint;
//       emit NewChain( id, info );
//       _vest_in_chain( id, vesting, msg.sender );
//       emit NewChainEndpoint( id, init_endpoint );
//       next_id++;
//   }

//   function vest_in_chain( uint id, uint vesting ) public {
//       _vest_in_chain( id, vesting, msg.sender );
//   }
   
// //   function withdraw_vest_from_chain( uint id, uint vesting ) public {
// //       _withdraw_vest_from_chain( id, vesting, msg.sender );
// //   }
   
//   function deposit_in_chain( uint id, uint deposit ) public {
//       _deposit_in_chain(id, deposit, msg.sender );
//   }

//   function has_vested( uint id, address user) view external returns (bool){
//       return chains[id].users[user].info.vesting > 0;
//   }

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
//   function _vest_in_chain( uint chain_id, uint96 vesting, address acc ) private {
//       bool existing_validator = validator_exists(chain_id, acc);
       
//       // Withdraw all vesting and delete validator account
//       if (vesting == 0) {
//         require( existing_validator == true, "Trying to withdraw vesting from non-existing validator account" );
//         require( chains[chain_id].validators.accounts[acc].info.next_vesting > 0, "Zero vesting balance. Can't withdraw any tokens" );
//         require( chains[chain_id].validators.accounts[acc].info.mining == false, "Can't withdraw any tokens, stop_minig must be called first." );  
//       }
//       // Vest in chain or withdraw just part of vesting
//       else {
//          require( chains[chain_id].active, "can't vest into non-existing chain" );
//          require( check_lition_min_vesting( vesting ), "user does not meet min. required chain criteria");
//          require( chains[chain_id].validator.check_vesting( vesting, acc ), "user does not meet chain criteria");
         
//          if (existing_validator == false) {
//              validator_create(chain_id, acc, 0, vesting, false);
//          }
//       }
      
//       // Withdraw all or just part of vesting
//       if( chains[chain_id].validators.accounts[acc].info.next_vesting > vesting ){
//          uint96 to_withdraw = chains[chain_id].validators.accounts[acc].info.vesting - vesting;
//          token.transfer( acc, chains[chain_id].validators.accounts[acc], to_withdraw);
         
//          if (chains[chain_id].validators.accounts[acc].info.mining == true) {
//             chains[chain_id].validators.accounts[acc].total_vesting -= to_withdraw; //TODO -= safe math here;
//          }
//       // Vest in chain
//       } else{
//          uint to_vest = vesting - chains[chain_id].validators.accounts[acc].info.vesting;
//          token.transferFrom( acc, address(this), to_vest);
         
//          if (chains[chain_id].validators.accounts[acc].info.mining == true) {
//             chains[chain_id].total_vesting += to_vest;
//          }
//       }
      
//       chains[chain_id].validators.accounts[acc].info.vesting = vesting;
//       //users_list_add(id, user);
      
//       // Event vesting can be either emitted here with flag it will applied in the next notary, 
//       // or it can be emitted automatically during the next notary when next_vesting replaces act_vesting 
//       // emit Vesting( id, vesting, user, now );
//   }
   
//     // function _withdraw_vest_from_chain( uint chain_id, uint vesting, address acc ) private {
//     //     require( chains[id].active, "can't withdraw vesting from non-existing chain" );
//     //     require (validator_exists(chain_id, acc) == true, "Trying to withdraw vesting from non-existing validator account" );
//     //     // require( chains[id].validators.accounts[acc].info.vesting > 0, "Zero vesting balance. Can't withdraw any tokens" ); // this shoudl never happen so no need to test it
//     //     require( chains[id].validators.accounts[acc].info.mining == false, "Can't withdraw any tokens, stop_minig must be called first." );  
        
//     //     // Withdraw all vesting and delete validator account
//     //     if (vesting == 0) {
//     //         chains[id].validators.accounts[acc].next_vesting = 
//     //     }    
        
//     //     // Withdraw just part of vesting
//     //     require( check_lition_min_vesting( vesting ), "user does not meet min. required chain criteria");
//     //     require( chains[id].validator.check_vesting( vesting, user ), "user does not meet chain criteria");
        
//     //     uint96 to_withdraw = chains[id].users[user].info.vesting - vesting;
//     //      token.transfer( user, to_withdraw);
         
//     //      if (chains[id].users[user].info.mining == true) {
//     //         chains[id].total_vesting -= to_withdraw; //TODO -= safe math here;
//     //      }
//     // }

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

//   function get_allowed_to_validate( uint id, uint batch ) view external returns (address[100] memory users, uint count) {
//      count = 0;
//      uint j = batch * 100;
//      while( j < (batch + 1)*100 && j < chains[id].users_list.length ) {
//       address user = chains[id].users_list[j];
//       if(chains[id].users[user].info.vesting > 0) {
//          users[count] = user;
//          count++;
//       }
//       j++;
//      }
//   }

//   function get_active_validators( uint id, uint batch ) view external returns (address[100] memory users, uint count) {
//      count = 0;
//      uint j = batch * 100;
//      while( j < (batch + 1)*100 && j < chains[id].users_list.length ) {
//       address user = chains[id].users_list[j];
//       if(chains[id].users[user].info.vesting > 0 && chains[id].users[user].info.mining == true) {
//          users[count] = user;
//          count++;
//       }
//       j++;
//      }
//   }

//   function start_mining(uint id) public {
//       require(chains[id].active == true, "Can't start mining on non-existing chain");
//       require(check_lition_min_vesting( chains[id].users[msg.sender].info.vesting) == true, "user does not meet min. required chain criteria");
//       require(chains[id].validator.check_vesting(chains[id].users[msg.sender].info.vesting, msg.sender) == true, "User does not meet chain criteria");
      
//       if (chains[id].users[msg.sender].info.mining == false) {
//           chains[id].total_vesting += chains[id].users[msg.sender].info.vesting;
//       }
//       chains[id].users[msg.sender].info.mining = true;
      
//       emit StartMining(id, msg.sender);
//   }

//   function stop_mining(uint id) public {
//       require(chains[id].active == true, "Can't start mining on non-existing chain");
//       require(check_lition_min_vesting( chains[id].users[msg.sender].info.vesting) == true, "user does not meet min. required chain criteria");
      
//       if (chains[id].users[msg.sender].info.mining == true) {
//           chains[id].total_vesting -= chains[id].users[msg.sender].info.vesting;
//       }
//       chains[id].users[msg.sender].info.mining = false;
      
//       emit StopMining(id, msg.sender);
//   }

}