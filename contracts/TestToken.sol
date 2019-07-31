pragma solidity >=0.5.4;

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

contract TestToken is ERC20{
   mapping(address => uint) _holding;
   mapping(address => mapping(address => uint)) _allowances;
   uint _totalSupply;
   string public constant symbol = "LIT";
   uint8 public constant decimals = 18;

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
      emit Transfer(msg.sender, to, tokens);
      return true;
   }

   function approve(address spender, uint tokens) public returns (bool success){
      _allowances[msg.sender][spender] = tokens;
      emit Approval(msg.sender, spender, tokens);
      return true;
   }

   function transferFrom(address from, address to, uint tokens) public returns (bool success){
      require( _allowances[from][msg.sender] >= tokens );
      _holding[from] -= tokens;
      _holding[to] += tokens;
      emit Transfer(from, to, tokens);
      return true;
   }

   function mint(address to, uint tokens) public{
      _holding[to] += tokens;
      _totalSupply += tokens;
   }

}
