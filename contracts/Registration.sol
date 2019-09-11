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
   event StartMining(uint indexed chain_id, address miner);
   event StopMining(uint indexed chain_id, address miner);

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
      uint256 last_notary;
      ChainValidator validator;
      uint total_vesting;
   }
   
   struct signature {
      uint8 v; bytes32 r; bytes32 s;
   }

   mapping(uint256 => chain_info) public chains;
   uint256 public next_id = 0;

   constructor(ERC20 _token) public {
      token = _token;
   }
   
   function users_list_add(uint chain_id, address user) internal {
     // User is already in list, do nothing
     if (chains[chain_id].users[user].index != 0) {
        return;
     }
    
     chains[chain_id].users_list.push(user);
     chains[chain_id].users[user].index = chains[chain_id].users_list.length; // indexes are stored + 1
   }
   
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

   function register_chain( string calldata info, ChainValidator validator, uint vesting, string calldata init_endpoint ) external returns ( uint256 id ){
      require( bytes(init_endpoint).length > 0 );
      id = next_id;
      chains[id].validator = validator;
      chains[id].active = true;
      chains[id].last_notary = 0;
      chains[id].users[msg.sender].info.endpoint = init_endpoint;
      emit NewChain( id, info );
      _vest_in_chain( id, vesting, msg.sender );
      emit NewChainEndpoint( id, init_endpoint );
      next_id++;
   }

   function vest_in_chain( uint id, uint vesting ) public {
      _vest_in_chain( id, vesting, msg.sender );
   }
   
   function deposit_in_chain( uint id, uint deposit ) public {
      _deposit_in_chain(id, deposit, msg.sender );
   }

   function has_vested( uint id, address user) view external returns (bool){
      return chains[id].users[user].info.vesting > 0;
   }

   function has_deposited(uint id, address user) view external returns (bool) {
      return chains[id].users[user].info.deposit > 0;
   }

   function get_signature_hash_from_notary(uint256 notary_block, address[] memory miners,
                                 uint32[] memory blocks_mined, address[] memory users,
                                 uint32[] memory user_gas, uint32 largest_tx)
                                     public pure returns (bytes32) {
      return keccak256(abi.encodePacked(notary_block, miners, blocks_mined, users, user_gas, largest_tx));
   }

   function get_last_notary(uint id) external view returns (uint256) {
     return chains[id].last_notary;
   }

   function process_users_consumptions(uint id, address[] memory users, uint32[] memory user_gas, uint32 largest_tx) internal returns (uint256 total_cost) {
     uint total_gas = 0;
     total_cost = 0;
     //largest tx fixed at 0.1 LIT - rework that to work with current price
     uint largest_reward = 10**17;

     for(uint i = 0; i < users.length; i++) {
        total_gas +=user_gas[i];
        uint user_cost = (user_gas[i] / largest_tx) * largest_reward;
        if( user_cost > chains[id].users[users[i]].info.deposit ) {
           user_cost = chains[id].users[users[i]].info.deposit;
           emit Deposit(id, 0, users[i], now);
        }
        chains[id].users[users[i]].info.deposit -= user_cost;
        total_cost += user_cost;
     }
   }

   function process_miners_rewards(uint id, uint256 notary_block_no, address[] memory miners, uint32[] memory blocks_mined, uint lit_to_distribute) internal {
     uint total_signatures = 0;
     chain_info storage chain = chains[id];
     uint max_blocks = notary_block_no - chain.last_notary ;
     for(uint i = 0; i < miners.length - 1; i++) {
         //TODO multiplier here for white paper specifics
        total_signatures += max_blocks / blocks_mined[i] * chain.users[miners[i]].info.vesting;
     }

     for(uint i = 0; i < miners.length - 1; i++) {
        uint miner_reward = max_blocks / blocks_mined[i] * chain.users[miners[i]].info.vesting * lit_to_distribute / total_signatures;
        token.transfer( miners[i], miner_reward );
        lit_to_distribute -= miner_reward;
     }

     token.transfer( miners[miners.length - 1], lit_to_distribute );
   }

   function notary(uint id, uint256 notary_block_no, address[] memory miners, uint32[] memory blocks_mined,
                                 address[] memory users, uint32[] memory user_gas, uint32 largest_tx,
                                 uint8[] memory v, bytes32[] memory r, bytes32[] memory s) public {
      //first, calculate hash from miners, block_mined, users and user_gas
      //then, do ec_recover of the signatures to determine signers
      //check if there is enough signers (total vesting of signers > 50% of all vestings)
      //then, calculate reward
      require(v.length == r.length);
      require(v.length == s.length);
      bytes32 signature_hash = get_signature_hash_from_notary(notary_block_no, miners, blocks_mined, users, user_gas, largest_tx);
      chain_info storage chain = chains[id];
      require(chain.active, "Trying to report about non-existing chain");

      uint involved_vesting = 0;

      for(uint i =0; i<v.length; i++) {
         involved_vesting += chain.users[ecrecover(signature_hash, v[i], r[i], s[i])].info.vesting;
      }

      require(involved_vesting * 2 >= chain.total_vesting);

      uint256 total_cost = process_users_consumptions(id, users, user_gas, largest_tx);
      process_miners_rewards(id, notary_block_no, miners, blocks_mined, total_cost);
      chain.last_notary = notary_block_no;
   }

   //TODO - rework so withdrawals are not processed immediatelly but after notary window
   function _vest_in_chain( uint id, uint vesting, address user ) private {
      //Validate value of vesting
      if (vesting == 0) {
        require( chains[id].users[user].info.vesting > 0, "Zero vesting balance. Can't withdraw any tokens" );
        require( chains[id].users[user].info.mining == false, "Can't withdraw any tokens, stop_minig must be called first." );  
      }
      else {
         require( chains[id].active, "can't vest into non-existing chain" );
         require( check_lition_min_vesting( vesting ), "user does not meet min. required chain criteria");
         require( chains[id].validator.check_vesting( vesting, user ), "user does not meet chain criteria");
      }
      
      if( chains[id].users[user].info.vesting > vesting ){
         uint to_withdraw = chains[id].users[user].info.vesting - vesting;
         token.transfer( user, to_withdraw);
         
         if (chains[id].users[user].info.mining == true) {
            chains[id].total_vesting -= to_withdraw; //TODO -= safe math here;
         }
      } else{
         uint to_vest = vesting - chains[id].users[user].info.vesting;
         token.transferFrom( user, address(this), to_vest);
         
         if (chains[id].users[user].info.mining == true) {
            chains[id].total_vesting += to_vest;
         }
      }
      
      chains[id].users[user].info.vesting = vesting;
      users_list_add(id, user);
      emit Vesting( id, vesting, user, now );
   }

   //TODO - rework so withdrawals are not processed immediatelly but after notary window
   function _deposit_in_chain( uint id, uint deposit, address user ) private {
      //Validate value of deposit
      if (deposit == 0) {
        require( chains[id].users[user].info.deposit > 0, "Zero deposit balance. Can't withdraw any tokens" );  
      }
      else {
        require( chains[id].active, "can't deposit into non-existing chain" );
        require( check_lition_min_deposit( deposit), "user does not meet min. required chain criteria");
        require( chains[id].validator.check_deposit( deposit, user ), "user does not meet chain criteria");
      }
      
      if( chains[id].users[user].info.deposit > deposit ){
         uint to_withdraw = chains[id].users[user].info.deposit - deposit;
         token.transfer( user, to_withdraw);
      } else{
         uint to_deposit = deposit - chains[id].users[user].info.deposit;
         token.transferFrom(user, address(this), to_deposit);
      }
      
      chains[id].users[user].info.deposit = deposit;
      users_list_add(id, user);
      emit Deposit(id, deposit, user, now);
   }

   function get_allowed_to_transact( uint id, uint batch ) view external returns (address[100] memory users, uint count) {
     count = 0;
     uint j = batch * 100;
     while( j < (batch + 1)*100 && j < chains[id].users_list.length ) {
       address user = chains[id].users_list[j];
       if(chains[id].users[user].info.deposit > 0) {
         users[count] = user;
         count++;
       }
       j++;
     }
   }

   function get_allowed_to_validate( uint id, uint batch ) view external returns (address[100] memory users, uint count) {
     count = 0;
     uint j = batch * 100;
     while( j < (batch + 1)*100 && j < chains[id].users_list.length ) {
       address user = chains[id].users_list[j];
       if(chains[id].users[user].info.vesting > 0) {
         users[count] = user;
         count++;
       }
       j++;
     }
   }

   function get_active_validators( uint id, uint batch ) view external returns (address[100] memory users, uint count) {
     count = 0;
     uint j = batch * 100;
     while( j < (batch + 1)*100 && j < chains[id].users_list.length ) {
       address user = chains[id].users_list[j];
       if(chains[id].users[user].info.vesting > 0 && chains[id].users[user].info.mining == true) {
         users[count] = user;
         count++;
       }
       j++;
     }
   }

   function start_mining(uint id) public {
      require(chains[id].active == true, "Can't start mining on non-existing chain");
      require(check_lition_min_vesting( chains[id].users[msg.sender].info.vesting) == true, "user does not meet min. required chain criteria");
      require(chains[id].validator.check_vesting(chains[id].users[msg.sender].info.vesting, msg.sender) == true, "User does not meet chain criteria");
      
      if (chains[id].users[msg.sender].info.mining == false) {
          chains[id].total_vesting += chains[id].users[msg.sender].info.vesting;
      }
      chains[id].users[msg.sender].info.mining = true;
      
      emit StartMining(id, msg.sender);
   }

   function stop_mining(uint id) public {
      require(chains[id].active == true, "Can't start mining on non-existing chain");
      require(check_lition_min_vesting( chains[id].users[msg.sender].info.vesting) == true, "user does not meet min. required chain criteria");
      
      if (chains[id].users[msg.sender].info.mining == true) {
          chains[id].total_vesting -= chains[id].users[msg.sender].info.vesting;
      }
      chains[id].users[msg.sender].info.mining = false;
      
      emit StopMining(id, msg.sender);
   }

}