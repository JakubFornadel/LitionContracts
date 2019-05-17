pragma solidity >=0.5.0;
pragma experimental ABIEncoderV2;


interface ChainValidator{
   function check_participant(uint vesting, address participant) external returns (bool);
}

contract DummyChainValidator{
   function check_participant(uint vesting, address participant) pure public returns (bool){
      if(vesting > 100*10^18 )
         return true;
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
}

contract TestToken is ERC20{
   mapping(address => uint) _holding;
   mapping(address => mapping(address => uint)) _allowances;
   uint _totalSupply;

   function totalSupply() public view returns (uint){
      return _totalSupply;
   }

   function balanceOf(address tokenOwner) public view returns (uint balance){
      return _holding[tokenOwner];
   }

   function allowance(address tokenOwner, address spender) public view returns (uint remaining){
      return _allowances[tokenOwner][spender];
   }

   function transfer(address to, uint tokens) public returns (bool success){
      require( _holding[msg.sender] >= tokens );
      _holding[msg.sender] -= tokens;
      _holding[to] += tokens;
      return true;
   }
   
   function approve(address spender, uint tokens) public returns (bool success){
      _allowances[msg.sender][spender] = tokens;
      return true;
   }

   function transferFrom(address from, address to, uint tokens) public returns (bool success){
      require( _allowances[from][msg.sender] >= tokens );
      _holding[from] -= tokens;
      _holding[to] += tokens;
      return true;
   }
   
   function mint(address to, uint tokens) public{
      _holding[to] += tokens;
      _totalSupply += tokens;
   }

}

contract LitionRegistry{
   event NewChain(uint id, string description);
   event NewChainEndpoint(uint id, string endpoint);
   event Deposit(uint indexed chain_id, uint deposit, address indexed depositer, uint256 datetime);
   event Vesting(uint indexed chain_id, uint deposit, address indexed depositer, uint256 datetime);

   ERC20 token;
   struct user_details{
      bool active;
      uint vesting;
      uint deposit;
      string endpoint;
   }

   struct chain_info{
      bool active;
      mapping(address => user_details) users;
      uint last_notary;
      ChainValidator validator;
      uint total_vesting;
   }

   mapping(uint256 => chain_info) public chains;
   uint256 public next_id = 0;
   
   constructor(ERC20 _token) public {
      token = _token;
   }

   struct signature {
      uint8 v; bytes32 r; bytes32 s;
   }

   function register_chain( string calldata info, ChainValidator validator, uint vesting, string calldata init_endpoint ) external returns ( uint256 id ){
      require( bytes(init_endpoint).length > 0 );
      id = next_id;
      chains[id].validator = validator;
      chains[id].active = true;
      chains[id].last_notary = 0;
      chains[id].users[msg.sender].active = true;
      chains[id].users[msg.sender].endpoint = init_endpoint;
      emit NewChain( id, info );
      _vest_in_chain( id, vesting, msg.sender );
      emit NewChainEndpoint( id, init_endpoint );
      next_id++;
   }

   function vest_in_chain( uint id, uint vesting ) public {
      _vest_in_chain( id, vesting, msg.sender );
   }

  function has_vested( uint id, address user) view external returns (bool){
      return chains[id].users[user].vesting > 0; 
   }

   function has_deposited( uint id, address user) view external returns (bool){
      return chains[id].users[user].deposit > 0; 
   }

   function notary(uint id, uint notary_block, address[] memory miners, uint[] memory blocks_mined, address[] memory users, uint[] memory user_gas, uint largest_tx, signature[] memory signatures) public{
      //first, calculate hash from miners, block_mined, users and user_gas 
      //then, do ec_recover of the signatures to determine signers
      //check if there is enough signers (total vesting of signers > 50% of all vestings)
      //then, calculate reward
      
      chain_info storage chain = chains[id];
      require(chain.active, "Trying to report about non-existing chain");
      chain.last_notary = notary_block;

      uint total_mined = 0;
      uint involved_vesting = 0;
      for(uint i = 0; i < miners.length; i++){
         total_mined += blocks_mined[i];
         involved_vesting += chain.users[miners[i]].vesting;
      }

      require(involved_vesting * 3/2 > chain.total_vesting);

      uint total_gas = 0; 
      uint total_cost = 0;
      //largest tx fixed at 0.1 LIT - rework that to work with current price
      uint largest_reward = 10**17;
      
      for(uint i = 0; i < users.length; i++){
         total_gas +=user_gas[i];
         uint user_cost = (user_gas[i] / largest_tx) * largest_reward;
         if( user_cost > chain.users[users[i]].deposit ) 
            user_cost = chain.users[users[i]].deposit;
         chain.users[users[i]].deposit -= user_cost;   
         total_cost += user_cost;
      }
      
      for( uint i = 0; i < miners.length - 1; i++ ){
         uint miner_reward = blocks_mined[i] * total_cost / total_mined;
         token.transfer( miners[i], miner_reward );
         total_cost -= miner_reward;
         total_mined -= blocks_mined[i];
      }

      token.transfer( miners[miners.length - 1], total_cost );

   }
   
   //TODO - rework so withdrawals are not processed immediatelly but after notary window
   function _vest_in_chain( uint id, uint vesting, address user ) private {
      if(vesting > 0 ){
         require( chains[id].active, "can't vest into non-existing chain" );
         require( chains[id].validator.check_participant( vesting, user ), "user does not meet chain criteria");
      }
      if( chains[id].users[user].vesting > vesting ){
         uint to_withdraw = chains[id].users[user].vesting - vesting;
         chains[id].total_vesting -= to_withdraw; //TODO - safe math here;
         token.transfer( user, to_withdraw);
      }else{
         uint to_deposit = vesting - chains[id].users[user].vesting;
         chains[id].total_vesting += to_deposit;
         token.transferFrom( user, address(this), to_deposit);
      }
      chains[id].users[user].vesting = vesting;
      emit Vesting( id, vesting, user, now ); 
   }

   function deposit_in_chain( uint id, uint deposit ) public {
      _deposit_in_chain(id, deposit, msg.sender );
   }
   
   //TODO - rework so withdrawals are not processed immediatelly but after notary window
   function _deposit_in_chain( uint id, uint deposit, address user ) public {
      if(deposit > 0){
         require( chains[id].active, "can't deposit into non-existing chain" );
      }
      if( chains[id].users[user].deposit > deposit ){
         uint to_withdraw = chains[id].users[user].deposit - deposit;
         token.transfer( user, to_withdraw);
      }else{
         uint to_deposit = deposit - chains[id].users[user].deposit;
         token.transferFrom(user, address(this), to_deposit);
      }
      chains[id].users[user].deposit = deposit;
      emit Deposit(id, deposit, user, now);
   }



}
