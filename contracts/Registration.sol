pragma solidity >=0.5.0;

interface ChainValidator{
   function check_participant(uint vesting, address participant) external returns (bool);
}

interface ERC20{
   function totalSupply() external view returns (uint);
   function balanceOf(address tokenOwner) external view returns (uint balance);
   function allowance(address tokenOwner, address spender) external view returns (uint remaining);
   function transfer(address to, uint tokens) external returns (bool success);
   function approve(address spender, uint tokens) external returns (bool success);
   function transferFrom(address from, address to, uint tokens) external returns (bool success);
}

interface LitionSidechainRegistry{
   //prerequisite for the register functions is, that the executing account approves this contract to spend LIT tokens
   function register_chain( string calldata info, ChainValidator validator ) external returns (uint256 id);
   function vest_in_chain( uint id, uint vesting ) external;
   function deposit_in_chain( uint id, uint deposit ) external;
   function register_participant() external;
   function deregister_participant() external;
   function notary() external;

   event NewChain(uint id, string description);
   event NewChainEndpoint(uint id, string endpoint);
   event Deposit(uint indexed chain_id, uint deposit, address indexed depositer);
   event Vesting(uint indexed chain_id, uint deposit, address indexed depositer);

}

contract LitionRegistry is LitionSidechainRegistry{
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
   }

   mapping(uint256 => chain_info) public chains;
   uint256 public next_id = 0;
   
   constructor(ERC20 _token) public {
      token = _token;
   }

   function register_chain( string calldata info, ChainValidator validator, uint vesting, string calldata init_endpoint ) external returns ( uint256 id ){
      require( bytes(init_endpoint).length > 0 );
      id = next_id;
      chains[id].validator = validator;
      chains[id].active = true;
      chains[id].last_notary = 0;
      chains[id].users[msg.sender].active = true;
      chains[id].users[msg.sender].vesting = vesting;
      chains[id].users[msg.sender].endpoint = init_endpoint;
      emit NewChain( id, info );
      _vest_in_chain( id, vesting, msg.sender );
      emit NewChainEndpoint( id, init_endpoint );
      next_id++;
   }

   function vest_in_chain( uint id, uint vesting ) public {
      _vest_in_chain( id, vesting, msg.sender );
   }
   
   function _vest_in_chain( uint id, uint vesting, address user ) private {
      require( chains[id].active, "can't vest into non-existing chain" );
      require( chains[id].validator.check_participant( vesting, user ), "user does not meet chain criteria");
      if( chains[id].users[user].vesting > vesting ){
         uint to_withdraw = chains[id].users[user].vesting - vesting;
         token.transfer( user, to_withdraw);
      }else{
         uint to_deposit = vesting - chains[id].users[user].vesting;
         token.transferFrom( user, address(this), to_deposit);
      }
      chains[id].users[user].vesting = vesting;
      emit Vesting( id, vesting, user ); 
   }

   function deposit_in_chain( uint id, uint deposit ) public {
      _deposit_in_chain(id, deposit, msg.sender );
   }
   
   function _deposit_in_chain( uint id, uint deposit, address user ) public {
      require( chains[id].active, "can't deposit into non-existing chain" );
      if( chains[id].users[user].deposit > deposit ){
         uint to_withdraw = chains[id].users[user].deposit - deposit;
         token.transfer( user, to_withdraw);
      }else{
         uint to_deposit = deposit - chains[id].users[user].deposit;
         token.transferFrom(user, address(this), to_deposit);
      }
      chains[id].users[user].deposit = deposit;
      emit Deposit(id, deposit, user);
   }



}
