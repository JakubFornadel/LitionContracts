pragma solidity >= 0.5.11;

interface ChainValidator {
   function checkVesting(uint256 vesting, address participant) external returns (bool);
   function checkDeposit(uint256 vesting, address participant) external returns (bool);
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
   function checkVesting(uint256 vesting, address participant) public returns (bool) {
      if(vesting >= 1000*(10**18) && vesting <= 500000*(10**18)) {
        return true;   
      }
      return false;
   }

   function checkDeposit(uint256 deposit, address participant) public returns (bool) {
      if(deposit >= 1000*(10**18)) {
         return true;
      }
      return false;
   }
}

contract LitionRegistry {
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function checkLitionMinVesting(uint256 vesting) private pure returns (bool) {
        if(vesting >= 1000*(10**18)) {
            return true;   
        }
        return false;
    }
    
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function checkLitionMinDeposit(uint256 deposit) private pure returns (bool) {
        if(deposit >= 1000*(10**18)) {
            return true;   
        }
        return false; 
    }
    
    using SafeMath for uint256;
    
    // New chain was registered
    event NewChain(uint256 chainId, string description, string endpoint);
    
    // New notary was processed
    event Notary(uint256 indexed chainId, uint256 notaryBlock);
    
    // Deposit request created
    event RequestDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 lastNotaryBlock);
    // Deposit request confirmed - tokens were transferred and account's deposit balance was updated
    event ConfirmDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqNotaryBlock);
    // Whole deposit withdrawned - this is special withdrawal, which as applied after the chain validators are not able to reach consensus for 1 month 
    event ForceWithdrawDeposit(uint256 indexed chainId, address indexed account);
    
    // Vesting request created
    event RequestVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 lastNotaryBlock);
    // Vesting request confirmed - tokens were transferred
    event ConfirmVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqNotaryBlock);
    // Whole vesting withdrawned - this is special withdrawal, which as applied after the chain validators are not able to reach consensus for 1 month 
    event ForceWithdrawVesting(uint256 indexed chainId, address indexed account);
    
    // if whitelist == true  - allow user to transact
    // if whitelist == false - do not allow user to transact
    event WhitelistAccount(uint256 indexed chainId, address indexed account, bool whitelist);
    
    // Validator start/stop mining
    event StartMining(uint256 indexed chainId, address indexed account);
    event StopMining(uint256 indexed chainId, address indexed account);


    /**************************************************************************************************************************/
    /***************************************** Structs related to the list of users *******************************************/
    /**************************************************************************************************************************/
    
    // Iterable map that is used only together with the Users mapping as data holder
    struct IterableMap {
        // map of indexes to the list array
        // indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing element
        mapping(address => uint256) listIndex;
        // list of addresses 
        address[]                   list;        
    }
    
    struct Validator {
        // Actual user's vesting
        uint96  vesting;
        // Flag if validator mined at least 1 block in the previous notary window
        bool    prevNotaryMined;
        // Flag if validator mined at least 1 block in second the previous notary window
        bool    secondPrevNotaryMined;
    }
    
    struct Transactor {
        // Actual user's deposit
        uint96  deposit;
        // Flag if user is whitelisted (allowed to transact) -> actual deposit must be greater than min. required deposit condition 
        bool    whitelisted;
    }
    
    struct User {
        // Validator's data
        Validator    validator;
        // Transactor's data
        Transactor   transactor;
    }
    
    
    /**************************************************************************************************************************/
    /*********************************** Structs related to the vesting/deposit requests **************************************/
    /**************************************************************************************************************************/
    
    struct VestingRequest {
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
        // In oldVesting is stored current vesting that validator had when new VestingRequest was accepted
        uint96                          oldVesting;
        // New value of vesting to be set
        uint96                          newVesting;
    }
    
    // Only full deposit withdrawals are saved as deposit requests - other types of deposits do not need to be confirmed
    struct DepositWithdrawRequest {
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
    }
    
    struct RequestsEntry {
        // index to the requestsList, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing request
        uint256                         index;
        // Deposit withdrawal request details
        DepositWithdrawRequest          depositWithdrawRequest;
        // Vesting request details
        VestingRequest                  vestingRequest;   
    }
    
    struct Requests {
        mapping(address => RequestsEntry)   accounts;
        address[]                           list;        
    }
    
    
    /**************************************************************************************************************************/
    /***************************************************** Other structs ******************************************************/
    /**************************************************************************************************************************/
    ERC20 token;
    
    struct LastNotary {
        // Timestamp(eth block time), when last notary was accepted
        uint256 timestamp;
        // Actual Lition block number, when the last notary was accepted
        uint256 block;
    }
    
    struct ChainInfo {
        // Side-chain description
        string                          description;
        
        // Side-chain endpoint
        string                          endpoint;
        
        // Flag that says if the side-chain was already(& sucessfully) registered
        bool                            registered;
        
        // Flag that says if the side-chain is active - creator already mined first block
        bool                            active;
        
        // How much vesting in total is vested in the side-chain by the active validators
        uint256                         totalVesting;
        
        // When was the last notary function called (block and time)
        LastNotary                      lastNotary;
        
        // Smart-contract for validating min.required deposit and vesting
        ChainValidator                  chainValidator;
        
        // Mapping data holder for users (validators & transactors) data
        mapping(address => User)        users;
        
        // List of existing validators - only those who has vested enough and are actively mining are here
        IterableMap                     validators;
        
        // Index of the validator with the smallest vesting balance among all existing validators
        // In case someone vests more tokens, he will replace the one with smallest vesting balance
        // index is shifted +1 compared to the real index of the list
        uint256                         lastValidatorIndex;
        
        // List of existing transactors - all user, who have != deposit balance are here even if they cannot transact anymore
        IterableMap                     transactors;
        
        // List of users requests for changing their vesting/deposit balance
        Requests                        requests;
    }
    
    // Registered chains
    mapping(uint256 => ChainInfo)   private chains;
    
    // Next chain id
    uint256                         public  nextId = 0;

    
    /**************************************************************************************************************************/
    /************************************ Contract interface - external/public functions **************************************/
    /**************************************************************************************************************************/
    
    // Requests vest in chain. It will be processed and applied to the actual user state after next:
    //      * 2 notary windows - in case new vesting < actual vesting
    //      * 3 notary windows - in case new vesting > actual vesting
    //
    // 
    // In case new vesting < actual vesting, user first creates request, his balance is internally updated, he should confirm this request as soon as possible(next window)
    // ann tokens are transferred back to his account. It takes 2 notary windows to finish this process. 
    //
    // In case new vesting > actual vesting, user first creates request, than he must confirm this request, tokens are transferred to the sc and his internal balance is updated 
    // in the next notary window after the one, in which confirm was called. It takes 3 notary windows to finish this process. 
    function requestVestInChain(uint256 chainId, uint256 vesting) external {
      ChainInfo storage chain = chains[chainId];
      require(chain.active == true, "Non-active chain");
      
      // Withdraw all vesting
      if (vesting == 0) {
          require(validatorExist(chainId, msg.sender) == true, "Non-existing validator account");
          
          // If last notary is older than 30 days, it means that validators cannot reach consensus and side-chain is basically stuck.
          // In such case ignore multi-step vesting process and allow users to withdraw all vested tokens
          if (chain.lastNotary.timestamp + 30 days < now) {
              forceWithdrawVestFromChain(chainId, msg.sender);
              return;
          }
          
          require(activeValidatorExist(chainId, msg.sender) == false, "Stop_minig must be called first.");  
      }
      // Vest in chain or withdraw just part of vesting
      else {
         require(vesting <= ~uint96(0), "vesting is greater than uint96_max_value");
         require(chain.users.accounts[msg.sender].validator.vesting != vesting, "Cannot vest the same amount of tokens as you already has vested.");
         require(checkLitionMinVesting(vesting), "user does not meet Lition's min.required vesting condition");
         require(chain.chainValidator.checkVesting(vesting, msg.sender), "user does not meet chain validator's min.required vesting condition");
      }
      
      require(vestingRequestExists(chainId, msg.sender) == false, "There is already existing request being processed for this acc.");
        
      _requestVestInChain(chainId, vesting, msg.sender);
    }
    
    // Confirms vest request, token transfer is processed during confirmation
    function confirmVestInChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
        
        require(vestingRequestExists(chainId, msg.sender) == true, "Non-existing vesting request.");
        require(chain.lastNotary.block > chains[chainId].requests.accounts[msg.sender].vestingRequest.notaryBlock, "Confirm can be called in the next notary window after request was accepted.");
        
        _confirmVestInChain(chainId, msg.sender);
    }
    
    // Requests deposit in chain. It will be processed and applied to the actual user state after next:
    //      * 1 notary window - in case new deposit != 0
    //      * emmidiately     - in case new deposit == 0
    // We need to handle only whole deposit withdrawals as it would allow users to send unlimited amount of txs to the sidechain and 
    // withdraw whole deposit right before notary function, in which user's comsumption is calculated and tokens are transferred. He would pay nothing...
    function requestDepositInChain(uint256 chainId, uint256 deposit) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
        
        // Withdraw whole deposit
        if (deposit == 0) {
          require(transactorExist(chainId, msg.sender) == true, "Non-existing transactor account");
          
          // If last notary is older than 30 days, it means that validators cannot reach consensus and side-chain is basically stuck.
          // In such case ignore multi-step deposit withdrawal process and allow users to withdraw all deposited tokens
          if (chain.lastNotary.timestamp + 30 days < now) {
              forceWithdrawDepositFromChain(chainId, msg.sender);
              return;
          }
          
        }
        // Deposit in chain or withdraw just part of deposit
        else {
         require(chain.users.accounts[msg.sender].transactor.deposit != deposit, "Cannot deposit the same amount of tokens as you already has deposited.");
         require(checkLitionMinDeposit(deposit), "user does not meet Lition's min.required deposit condition");
         require(chain.chainValidator.checkDeposit(deposit, msg.sender), "user does not meet chain validator's min.required deposit condition");
         require(deposit <= ~uint96(0), "deposit is greater than uint96_max_value");
        }
        
        require(depositWithdrawRequestExists(chainId, msg.sender) == false, "There is already existing withdrawal request being processed for this acc.");
                
        _requestDepositInChain(chainId, deposit, msg.sender);
    }
    
    // Confirms deposit withdrawal request, token transfer is processed during confirmation
    function confirmDepositWithdrawalFromChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
        
        require(depositWithdrawRequestExists(chainId, msg.sender) == true, "Cannot confirm non-existing deposit withdrawal request.");
        require(chain.lastNotary.block > chains[chainId].requests.accounts[msg.sender].depositWithdrawRequest.notaryBlock, "Confirm can be called in the next notary window after request was accepted.");
        
        _confirmDepositWithdrawalFromChain(chainId, msg.sender);
    }
    
    // Internally creates/registers new side-chain. Creator must be also validator at least from the beginning as joining process take multiple steps
    // and these steps cannot be done in the same notary window
    function registerChain(string calldata description, ChainValidator validator, uint96 vesting, uint96 deposit, string calldata initEndpoint) external returns (uint256 chainId) {
        require(bytes(description).length   > 0,            "Chain description cannot be empty");
        require(bytes(initEndpoint).length  > 0,            "Chain endpoint cannot be empty");
        require(deposit                     <= ~uint96(0),  "deposit is greater than uint96_max_value");
        require(vesting                     <= ~uint96(0),  "vesting is greater than uint96_max_value");
        
        address creator         = msg.sender;
        uint256 timestamp       = now;
        
        // Inits chain data
        chainId                 = nextId;
        ChainInfo storage chain = chains[chainId];
        
        chain.chainValidator    = validator;
        
        // Validates vesting
        require(checkLitionMinVesting(vesting), "chain creator does not meet Lition's min.required vesting condition");
        require(chain.chainValidator.checkVesting(vesting, creator), "chain creator does not meet chain validator's min.required vesting condition");
        
        // Validates deposit
        require(checkLitionMinDeposit(deposit), "chain creator does not meet Lition's min.required deposit condition");
        require(chain.chainValidator.checkDeposit(deposit, creator), "chain creator does not meet chain validator's min.required deposit condition");
        
        // Transfers vesting tokens
        token.transferFrom(creator, address(this), vesting);
        // Inits validator's data
        validatorCreate(chainId, creator, vesting);
        
        // Transfers deposit tokens
        token.transferFrom(creator, address(this), deposit);
        // Inits transactor's data and inserts it into the list of validators
        transactorCreate(chainId, creator, deposit);
        
        chain.description       = description;
        chain.registered        = true;
        chain.endpoint          = initEndpoint;
        
        emit NewChain(chainId, description, initEndpoint);
        
        emit RequestVestInChain(chainId, creator, vesting, timestamp);
        emit ConfirmVestInChain(chainId, creator, vesting, timestamp);
        
        emit RequestDepositInChain(chainId, creator, deposit, timestamp);
        emit ConfirmDepositInChain(chainId, creator, deposit, timestamp);
        
        nextId++;
    }
    
    // Reurns true, if acc has vested enough to become validator, othervise false
    function hasVested(uint256 chainId, address acc) view external returns (bool) {
        // No need to check vesting balance as it cannot be lover than min. required
        return validatorExist(chainId, acc);
    }
    
    // Returns true if user's remaining deposit balance >= min. required deposit and is allowed to transact
    function hasDeposited(uint256 chainId, address acc) view external returns (bool) {
        // No need to check deposit balance as whitelisted flag should be alwyas set accordingly
        return chains[chainId].users.accounts[acc].transactor.whitelisted;
    }
    
    function getLastNotary(uint256 chainId) external view returns (uint256 notaryBlock, uint256 notaryTimestamp) {
        notaryBlock = chains[chainId].lastNotary.block;
        notaryTimestamp = chains[chainId].lastNotary.timestamp;
    }
    
    function testNotary(uint256 chainId, uint256 notaryBlock) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.registered, "Non-registered chain");
        
        // TODO: remove validators(call stopMining) who signed no block during this notary window and have mining flag == true
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryBlock;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
        
        emit Notary(chainId, notaryBlock);
    }
    

    function getChainDetails(uint256 chainId) external view returns (string memory description, string memory endpoint, bool registered, bool active, 
                                                                     uint256 totalVesting, address validator, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp) {
        ChainInfo storage chain = chains[chainId];
        
        description         = chain.description;
        endpoint            = chain.endpoint;
        registered          = chain.registered;
        active              = chain.active;
        totalVesting        = chain.totalVesting;
        validator           = chain.chainValidator;
        lastNotaryBlock     = chain.lastNotary.block;
        lastNotaryTimestamp = chain.lastNotary.timestamp;
    }
    
    function getUserDetails(uint256 chainId, address acc) external view returns (uint256 deposit, bool whitelisted, uint256 vesting, 
                                                                                 bool mining, bool prevNotaryMined, bool secondPrevNotaryMined) {
        ChainInfo storage chain = chains[chainId];
         
        deposit                 = chain.users[acc].transactor.deposit;
        whitelisted             = chain.users[acc].transactor.whitelisted;
        vesting                 = chain.users[acc].validator.vesting;
        mining                  = activeValidatorExist(chainId, acc);
        prevNotaryMined         = chain.users[acc].validator.prevNotaryMined;  
        secondPrevNotaryMined   = chain.users[acc].validator.secondPrevNotaryMined;  
    }
    
    function getUserRequests(uint256 chainId, address acc) external view returns (bool vestingReqExists, uint256 vestingReqNotary, uint256 vestingReqValue, 
                                                                                  bool depositReqExists, uint256 depositReqNotary) {
        if (vestingRequestExists(chainId, acc)) {
            VestingRequest storage request = chains[chainId].requests.accounts[acc].vestingRequest;
            
            vestingReqExists        = true;
            vestingReqNotary        = request.notaryBlock;
            vestingReqValue         = request.newVesting;
        }
        
        if (depositWithdrawRequestExists(chainId, acc)) {
            DepositWithdrawRequest storage request = chains[chainId].requests.accounts[acc].depositWithdrawRequest;
            
            depositReqExists = true;
            depositReqNotary = request.notaryBlock;
        }
    }
    
    
    // Notarization function - calculates user consumption as well as miner rewards
    // First, calculate hash from miners, block_mined, users and userGas
    // then, do ec_recover of the signatures to determine signers
    // check if there is enough signers (total vesting of signers > 50% of all vestings)
    // then, calculate reward
    function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] memory miners, uint32[] memory blocksMined,
                    address[] memory users, uint32[] memory userGas, uint32 largestTx,
                    uint8[] memory v, bytes32[] memory r, bytes32[] memory s) public {
                  
        ChainInfo storage chain = chains[chainId];
        require(chain.registered, "Non-registered chain");
    
        // Validates statistics data
        require(v.length            == r.length,                "Invalid data: v.length != r.length");
        require(v.length            == s.length,                "Invalid data: v.length != s.length");
        require(notaryStartBlock    >  chain.lastNotary.block,  "Invalid data: notaryBlock_start from statistics must be greater than the last known notary block");
        require(notaryEndBlock      >  notaryStartBlock,        "Invalid data: notaryEndBlock must be greater than notaryStartBlock");
        require(largestTx           >  0,                       "Invalid data: Largest tx must be greater than zero");
        require(miners.length       == blocksMined.length,      "Invalid data: num of miners != num of block mined");
        require(users.length        == userGas.length,          "Invalid data: num of users != num of users gas");
        
        bytes32 signatureHash = keccak256(abi.encodePacked(notaryEndBlock, miners, blocksMined, users, userGas, largestTx));
        
        // Involved vesting based on validator's, who signed statistics for this notary window. 
        // These statistics are used for calculating usage cost and miner rewards are calculated
        uint256 involvedVesting = 0;
        for(uint256 i = 0; i < v.length; i++) {
            involvedVesting += chain.users.accounts[ecrecover(signatureHash, v[i], r[i], s[i])].validator.vesting;
        }
        
        // There must be at least 50% out of total possible vesting involved
        involvedVesting = involvedVesting.mul(2);
        require(involvedVesting >= chain.totalVesting);
        
        // Calculates total cost based on user's usage durint current notary window
        uint256 totalCost = processUsersConsumptions(chainId, users, userGas, largestTx);
        
        // Calculates and process validator's rewards based on their participation rate and vesting balance
        processMinersRewards(chainId, notaryEndBlock, miners, blocksMined, totalCost);
        
        // TODO: remove validators(call stopMining) who signed no block during this notary window and have mining flag == true
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryEndBlock;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
        
        emit Notary(chainId, notaryEndBlock);
    }
    
    
    // Returns list of user's addresses that are allowed to transact - their deposit >= min. required deposit
    function getAllowedToTransact(uint256 chainId, uint256 batch) view external returns (address[100] memory, uint256, bool) {
        
    }
    
    // Returns list of addresses of active validators
    function getActiveValidators(uint256 chainId, uint256 batch) view external returns (address[100] memory, uint256, bool) {
        
    }
    
    // Sets mining validator's mining flag to true and emit event so other nodes vote him
    function startMining(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.registered == true, "Non-registered chain");
        require(validatorExists(chainId, msg.sender) == true, "Non-existing validator");
        require(vestingRequestExists(chainId, msg.sender) == false, "Cannot start mining - there is ongoing vesting request.");
        require(chains[chainId].chainValidator.checkVesting(chain.users.accounts[msg.sender].validator.vesting, msg.sender) == true, "User does not meet chain validator's min.required vesting condition");
        
        _startMining(chainId, msg.sender);
    }
  
    // Sets mining validator's mining flag to false and emit event so other nodes unvote
    function stopMining(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.registered == true, "Non-registered chain");
        require(validatorExists(chainId, msg.sender) == true, "Non-existing validator");
        
        _stopMining(chainId, msg.sender);
    }
    

    /**************************************************************************************************************************/
    /**************************************** Functions related to the list of users ******************************************/
    /**************************************************************************************************************************/
    
    // Adds acc from the map
    function insertAcc(IterableMap storage map, address acc) internal {
        map.list.push(acc);
        // indexes are stored + 1   
        map.listIndex[acc] = map.list.length;
    }
    
    // Removes acc from the map
    function removeAcc(IterableMap storage map, address acc) internal {
        require(map.listIndex[acc] <= map.list.length, "RemoveAcc invalid index");
    
        // Move an last element of array into the vacated key slot.
        uint256 foundIndex = map.listIndex[acc] - 1;
        uint256 lastIndex  = map.list.length - 1;
    
        map.listIndex[map.list[lastIndex]] = foundIndex + 1;
        map.list[foundIndex] = map.list[lastIndex];
        map.list.length--;
    
        // Deletes element
        map.listIndex[acc] = 0;
    }
    
    // Returns true, if acc exists in the iterable map, otherwise false
    function existAcc(IterableMap storage map, address acc) internal returns (bool) {
        return map.listIndex[acc] != 0;
    }
    
    // Inits validator data holder in the users mapping 
    function validatorCreate(uint256 chainId, address acc, uint256 vesting) internal {
        Validator storage validator = chains[chainId].users[acc].validator;
        validator.vesting               = vesting;
        // Inits previously notary windows as mined so validator does not get removed from the list of actively mining validators right after the creation
        validator.prevNotaryMined       = 1;
        validator.secondPrevNotaryMined = 1;
    }
    
    // Deinits validator data holder in the users mapping 
    function validatorDelete(uint256 chainId, address acc) internal {
        Validator storage validator = chains[chainId].users[acc].validator;
        validator.vesting               = 0;
        validator.prevNotaryMined       = false;
        validator.secondPrevNotaryMined = false;
        
        if (validatorExist(chainId, acc) == true) {
            validatorRemove(chainId, acc);
        }
    }
    
    // Inserts validator into the list of actively mining validators
    function validatorInsert(uint256 chainId, address acc) internal {
        insertAcc(chains[chainId].validators, acc);   
    }
    
    // Removes validator from the list of actively mining validators
    function validatorRemove(uint256 chainId, address acc, uint) internal {
        removeAcc(chains[chainId].validators, acc);   
    }
    
    // Returns true, if acc is in the list of actively mining validators, otherwise false
    function activeValidatorExist(uint256 chainId, address acc) internal returns (bool) {
        return existAcc(chains[chainId].validators, acc);
    }
    
    // Returns true, if acc hase vesting > 0, otherwise false
    function validatorExist(uint256 chainId, address acc) internal returns (bool) {
        return chains[chainId].users[acc].validator.vesting > 0;
    }
    
    // Inits transactor data holder in the users mapping and inserts it into the list of transactors
    function transactorCreate(uint256 chainId, address acc, uint256 deposit) internal {
        Transactor storage transactor = chains[chainId].users[acc].transactor;
        transactor.deposit            = deposit;
        transactor.whitelisted        = true;
        
        insertAcc(chains[chainId].transactors, acc);
    }
    
    // Deinits transactor data holder in the users mapping and removes it from the list of transactors
    function transactorDelete(uint256 chainId, address acc) internal {
        Transactor storage transactor = chains[chainId].users[acc].transactor;
        transactor.deposit            = 0;
        transactor.whitelisted        = false;
        
        removeAcc(chains[chainId].transactors, acc);
    }
    
    // Returns true, if acc is in the list of transactors
    function transactorExist(uint256 chainId, address acc) internal returns (bool) {
        return existAcc(chains[chainId].transactors, acc);
    }
    
    /**************************************************************************************************************************/
    /*********************************** Functions related to the vesting/deposit requests ************************************/
    /**************************************************************************************************************************/
    
    // Creates new vesting request
    function vestingRequestCreate(uint256 chainId, address acc, uint256 vesting) internal {
        ChainInfo storage chain          = chains[chainId];
        RequestsEntry storage entry      = chain.requests.accounts[acc];
        
        entry.vestingRequest.oldVesting  = chain.users.accounts[acc].validator.vesting;
        entry.vestingRequest.newVesting  = uint96(vesting);
        entry.vestingRequest.notaryBlock = chain.lastNotary.block; 
        
        // There is no deposit or vesting ongoing request - create new RequestsEntry structure
        if (entry.index == 0) { // anyRequestExists(chainId, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chain.requests.list.push(acc);    
            entry.index = chain.requests.list.length; // indexes are stored + 1
        }
    }

    // Creates new deposit withdrawal request
    function depositWithdrawRequestCreate(uint256 chainId, address acc) internal {
        ChainInfo storage chain                  = chains[chainId];
        RequestsEntry storage entry              = chain.requests.accounts[acc];
        
        entry.depositWithdrawRequest.notaryBlock = chain.lastNotary.block; 
        
        // There is no deposit or vesting ongoing request - create new RequestsEntry structure
        if (entry.index == 0) { // anyRequestExists(chainId, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chain.requests.list.push(acc);    
            entry.index = chain.requests.list.length; // indexes are stored + 1
        }
    }

        
    // Deletes existing requests pair(vesting & deposit) from the internal list of requests
    // This method should never be called directly, vestingRequestDelete & depositWithdrawRequestDelete should be called instead
    function requestsPairDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        RequestsEntry storage entry = chain.requests.accounts[acc];
        
        // request_exists(chainId, acc), vestingRequestExists(chainId, acc) and deposoti_withdraw_exists(chainId, acc) could be used instead
        require(entry.index != 0, "Request does not exist");
        
        address[] storage requestsList = chain.requests.list;
    
        require(entry.index <= requestsList.length, "Invalid index value");
    
        // Move an last element of array into the vacated key slot.
        uint256 foundIndex = entry.index - 1;
        uint256 lastIndex = requestsList.length - 1;
    
        chain.requests.accounts[requestsList[lastIndex]].index = foundIndex + 1;
        requestsList[foundIndex] = requestsList[lastIndex];
        requestsList.length--;
    
        delete chain.requests.accounts[acc];
    }
    
    function vestingRequestDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no ongoing deposit request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].depositWithdrawRequest.timestamp == 0) {
            requestsPairDelete(chainId, acc);
            return;
        } 
        
        // There is ongoing deposit request for this account - only reset vesting request
        VestingRequest storage request = chain.requests.accounts[acc].vestingRequest;
        request.notaryBlock    = 0;
        request.oldVesting     = 0;
        request.newVesting     = 0;
    }
    
    function depositWithdrawRequestDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no ongoing vesting request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].vestingRequest.timestamp == 0) {
            requestsPairDelete(chainId, acc);
            return;
        } 
        
        // There is ongoing vesting request for this account - only reset vesting request
        DepositWithdrawRequest storage request = chain.requests.accounts[acc].depositWithdrawRequest;
        request.notaryBlock = 0;
    }
    
    // Checks if acc has any ongoing vesting or deposit request
    function anyRequestExists(uint256 chainId, address acc) internal view returns (bool) {
      return chains[chainId].requests.accounts[acc].index != 0;
    }
    
    // Checks if acc has any ongoing vesting request
    function vestingRequestExists(uint256 chainId, address acc) internal view returns (bool) {
      return chains[chainId].requests.accounts[acc].vestingRequest.timestamp != 0;
    }
    
    // Checks if acc has any ongoing DEPOSIT WITHDRAWAL request
    function depositWithdrawRequestExists(uint256 chainId, address acc) internal view returns (bool) {
      return chains[chainId].requests.accounts[acc].depositWithdrawRequest.timestamp != 0;
    }
    
    function _requestVestInChain(uint256 chainId, uint256 vesting, address acc) internal {
      ChainInfo storage chain     = chains[chainId];
      Validator storage validator = chain.users[acc].validator;
      
      // if newVesting > oldVesting, send tokens first and update internal vesting balance in cofirm method
      if (vesting > validator.vesting) {
        // Internally creates new user
        if (validator.vesting == 0) {
          validatorCreate(chainId, acc);
        }
        
        uint256 toVest = vesting - validator.vesting;
        token.transferFrom(acc, address(this), toVest);
      }
      // if newVesting < oldVesting, update internal vesting balance first and send tokens in cofirm method
      else {
        uint256 toDecreaseChainVesting;
         
        // If it was request to withdraw whole vesting balance, delete validator
        if (vesting == 0) {
          toDecreaseChainVesting = validator.vesting;
          validatorDelete(chainId, acc);
        }
        else {
          toDecreaseChainVesting = validator.vesting - vesting;
          validator.vesting = vesting;    
        } 
        
        // If validator is actively mining, decrease chain's total vesting
        if (activeValidatorExist(chainId, acc) == true) {
            chain.totalVesting = chain.totalVesting.sub(toDecreaseChainVesting);
        }
      }
      
      vestingRequestCreate(chainId, acc, vesting);
      emit RequestVestInChain(chainId, acc, vesting, chain.lastNotary.block);
    }
    
    // Basically just transfers the tokens, validator's vesting balance update is always done at the of notary atomatically
    function _confirmVestInChain(uint256 chainId, address acc) internal {
        ChainInfo storage chain        = chains[chainId];
        VestingRequest storage request = chain.requests.accounts[acc].vestingRequest;
        
        uint256 oldVesting = request.oldVesting;
        uint256 newVesting = request.newVesting;
        vestingRequestDelete(chainId, acc);
        
        if(newVesting > oldVesting) {
            Validator storage validator = chains[chainId].users[acc].validator;
            validator.vesting = request.newVesting;
            
            if (activeValidatorExist(chainId, acc) == true) {
                chain.totalVesting = chain.totalVesting.add(newVesting - oldVesting);
            }
        }
        else {
            uint256 toWithdraw = oldVesting - newVesting;
            token.transfer(acc, toWithdraw);
        }
        
        emit ConfirmVestInChain(chainId, acc, newVesting, chain.lastNotary.block);
    }
    
    // Forcefully withdraw whole vesting from chain.
    // Because vesting is processed during 2(new_vest < act_vest) or even 3(new_vest > act_vest) notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing vest requests are deleted after this withdraw.
    function forceWithdrawVestFromChain(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        uint256 toWithdraw;
        uint256 toDecreaseChainVesting;
        
        // No ongoing vesting request is present
        if (vestingRequestExists(chainId, acc) == false) {
            toWithdraw = chain.users.accounts[acc].validator.vesting;
            toDecreaseChainVesting = toWithdraw;
        }
        // There is ongoing vesting request
        else { 
            VestingRequest storage request = chain.requests.accounts[acc].vestingRequest;
            // Token transfer was already processed -> use new vesting balance as actual user's vesting balance to withdraw
            if (request.newVesting > request.oldVesting) {
                toWithdraw              = request.newVesting;
                toDecreaseChainVesting  = request.oldVesting;
            }
            // Token transfer was not yet processed -> use saved old vesting balance as actual user's vesting balance to withdraw
            else {
                toWithdraw              = request.oldVesting;
                toDecreaseChainVesting  = request.newVesting;
            }
            
            vestingRequestDelete(chainId, acc);
        }
        
        // If validor is actively mining update chain's totalVesting
        if (activeValidatorExist(acc) == true) {
            chain.totalVesting = chain.totalVesting.sub(toDecreaseChainVesting);
            emit StopMining(chainId, acc);
        }
        
        validatorDelete(chainId, acc);
        
        // Transfers all remaining tokens to the user account
        token.transfer(acc, toWithdraw);
        
        emit ForceWithdrawVesting(chainId, acc, now);
    }
    
    function _requestDepositInChain(uint256 chainId, uint256 deposit, address acc) internal {
      uint256 timestamp = now;
      
      // If user wants to withdraw whole deposit - create withdrawal request
      if (deposit == 0) {
        depositWithdrawRequestCreate(chainId, acc);
        emit RequestDepositInChain(chainId, acc, deposit, timestamp);  
        
        return;
      }
      
      // Internally creates new user
      if (userExists(chainId, acc) == false) {
          userCreate(chainId, acc);
      }
      
      // If user wants to deposit in chain, process it immediately
      Transactor storage transactor = chains[chainId].users.accounts[acc].transactor;
      uint256 actTransactorDeposit = transactor.deposit;
      
      if(actTransactorDeposit > deposit) {
         transactor.deposit = uint96(deposit);
         
         uint256 toWithdraw = actTransactorDeposit - deposit;
         token.transfer(acc, toWithdraw);
      } else {
         uint256 toDeposit = deposit - actTransactorDeposit;
         token.transferFrom(acc, address(this), toDeposit);
         
         transactor.deposit = uint96(deposit);
      }
      
      emit RequestDepositInChain(chainId, acc, deposit, timestamp);
      emit ConfirmDepositInChain(chainId, acc, deposit, timestamp);
      
      if (transactor.whitelisted == false) {
        transactor.whitelisted = true;
        emit WhitelistAccount(chainId, acc, true);
      }
    }
    
    function _confirmDepositWithdrawalFromChain(uint256 chainId, address acc) internal {
        DepositWithdrawRequestData storage request = chains[chainId].requests.accounts[acc].depositWithdrawRequest;
        Transactor storage transactor = chains[chainId].users.accounts[acc].transactor;
        
        uint256 toWithdraw = transactor.deposit;
        
        depositWithdrawRequestDelete(chainId, acc);
        transactorDelete(chainId, acc);
        
        // Withdraw whole deposit
        token.transfer(acc, toWithdraw);
        
        emit ConfirmDepositInChain(chainId, acc, 0, request.timestamp);
    }
    
    // Forcefully withdraw whole deposit.
    // Because deposit withdrawal is processed during 2 notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing deposit requests are deleted after this withdraw.
    function forceWithdrawDepositFromChain(uint256 chainId, address acc) internal {
        Transactor storage transactor = chains[chainId].users.accounts[acc].transactor; 
        
        uint256 toWithdraw = transactor.deposit;
        transactorDelete(chainId, acc);
        
        // If deposit withdrawal request exists, delete it
        if (depositWithdrawRequestExists(chainId, acc) == true) {
            depositWithdrawRequestDelete(chainId, acc);    
        }
        
        // Transfers all remaining tokens to the user account
        token.transfer(acc, toWithdraw);
        
        emit ForceWithdrawDeposit(chainId, acc, now);
    }

    /**************************************************************************************************************************/
    /*************************************************** Other functions ******************************************************/
    /**************************************************************************************************************************/

   constructor(ERC20 _token) public {
      token = _token;
   }
  
  // Process users consumption based on their usage
  function processUsersConsumptions(uint256 chainId, address[] memory users, uint32[] memory userGas, uint32 largestTx) internal returns (uint256 totalCost) {
     ChainInfo storage chain = chains[chainId];
     
     // Total usage cost in LIT tokens
     totalCost = 0;
     
     // largest tx fee fixed at 0.1 LIT
     uint256 largestReward = 10**17;
     
     // Individual user's usage cost in LIT tokens
     uint256 userCost;
     
     // Use uint256 transactorDeposit instead of stored uint96 transactor.deposit because of simplified math  
     uint256 transactorDeposit;
     address acc;
     for(uint256 i = 0; i < users.length; i++) {
        acc = users[i];
        Transactor storage transactor = chain.users.accounts[acc].transactor;
        transactorDeposit = transactor.deposit;
        
        // This can happen only if there is non-registered user in statistics, which means there is probably ongoing coordinated attack
        // This if should ideally never evaluate to true
        if (transactorDeposit == 0) {
            // Ignores non-registred user and let nodes know he is not allowed to transact
            emit WhitelistAccount(chainId, users[i], false);
            continue;
        }
        
        userCost = (userGas[i] / largestTx) * largestReward;
        
        // This should never happen as it is handled by 2-step deposit withdrawal system and
        // by checking user's deposit balance is >= min. required deposit conditions
        if(userCost > transactorDeposit ) {
           userCost = transactorDeposit;
           
           transactor.whitelisted = false;
           emit WhitelistAccount(chainId, users[i], false);
        }
        
        transactorDeposit -= userCost;
        
        // Updates user's stored deposit balance based on his usage
        transactor.deposit = uint96(transactorDeposit);
        
        // Check if user's deposit balance is >= min. required deposit conditions
        if (checkLitionMinDeposit(transactorDeposit) == false || chain.chainValidator.checkDeposit(transactorDeposit, acc) == false) {
            // If not, do not allow him to transact anymore
            transactor.whitelisted = false;
            emit WhitelistAccount(chainId, acc, false);
        }
        
        // Adds user's cost to the total cost
        // No need for safe math as we internally allow max 10^48 users(even if there is theoretically more users in statistics, they are ignored here)
        // max possible userCost is 10^32 * 10^17, so max possible totalCost is 10^48 * 10^32 * 10^17 = 10^97, which will never overfloww uint256
        totalCost += userCost;  
     }
   }

   // Process miners rewards based on their participation rate(how many blocks they signed) and their vesting balance
   function processMinersRewards(uint256 chainId, uint256 notaryBlockNo, address[] memory miners, uint32[] memory blocksMined, uint256 litToDistribute) internal {
     ChainInfo storage chain = chains[chainId];
     
     // Min. vesting balance to be a trust node. Trust Nodes haved doubled(virtually) vesting
     uint256 minTrustNodeVesting = 50000*(10**18); 
     
     // How many block could validator mined since the last notary in case he did sign every possible block 
     uint256 maxBlocksMined = notaryBlockNo - chain.lastNotary.block;

     // Total involved vesting 
     uint256 totalInvolvedVesting = 0;
     
     // Selected validator's vesting balance
     uint256 validatorVesting;
     
     // Max number of miners for which we do not need to worry about uin256 overflow in math
     uint256 overflowMaxMiners = 10**20;
     
     // Runs through all miners and calculates total involved vesting based on:
     for(uint256 i = 0; i < miners.length; i++) {
        validatorVesting = chain.users.accounts[miners[i]].validator.vesting;
        
        // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
        if (validatorVesting >= minTrustNodeVesting) {
            // Validator's stored vesting is max uint96
            validatorVesting *= 2;
        }
        // This can happen only if there is non-registered validator in statistics, which means there is probably ongoing coordinated attack
        // This if should ideally never evaluate to true
        else if (validatorVesting == 0) {
            // Ignores non-registred miner
            continue;
        }

        // No need for safe math as we internally allow max 10^48 users(even if there is theoretically more validators in statistics, they are ignored here)
        // max possible (maxBlocksMined / blocksMined[i]) valuse is 10^32, max possible validatorVesting value is 10^96, when virtually doubled it is 10^192, 
        // so max possible totalInvolvedVesting value is 10^48 * 10^32 * 10^192 = 10^272, which can possibly overfloww uint256 only if there is more than 2^20 miners
        if (i < overflowMaxMiners) { 
            totalInvolvedVesting += (maxBlocksMined / blocksMined[i]) * validatorVesting;
        }
        else {
            totalInvolvedVesting = totalInvolvedVesting.add((maxBlocksMined / blocksMined[i]) * validatorVesting);
        }
     }

     
     // Runs through all miners and calculates their reward based on:
     //     1. How many blocks out of max_blocks_moned each miner signed
     //     2. How many token each miner vested
     for(uint256 i = 0; i < miners.length; i++) {
        validatorVesting = chain.users.accounts[miners[i]].validator.vesting;
        
        // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
        if (validatorVesting >= minTrustNodeVesting) {
            // Validator's stored vesting is max uint96
            validatorVesting *= 2;
        }
        // This can happen only if there is non-registered validator in statistics, which means there is probably ongoing coordinated attack
        // This if should ideally never evaluate to true
        else if (validatorVesting == 0) {
            // Ignores non-registred miner
            continue;
        }
        
        // No need for safe math as max value of (maxBlocksMined / blocksMined[i]) is 10^32, max value of (validatorVesting / totalInvolvedVesting) is 1 and 
        // max value of litToDistribute(calculated in processUsersConsumptions) is 10^97, so max possible miner reward is 10^32 * 1 * 10^97 = 10^129
        uint256 minerReward = (maxBlocksMined / blocksMined[i]) * (validatorVesting / totalInvolvedVesting) * litToDistribute;
        token.transfer(miners[i], minerReward);
        
        // No need for safe math as miner reward is calculated as fraction of total litToDistribute and all miners rewards must always be <= litToDistribute
        litToDistribute -= minerReward;
     }

     // Sends the rest(math rounding) to the miner, who called notary function
     token.transfer(msg.sender, litToDistribute);
   }
   
  function getUsers(uint256 chainId, uint256 batch, bool validators, bool active) internal view returns (address[100] memory users, uint256 count, bool end) {
     ChainInfo storage chain = chains[chainId];
     
     count = 0;
     uint256 j = batch * 100;
     uint256 usersTotalCount = chain.users.list.length;
     
     while(j < (batch + 1)*100 && j < usersTotalCount) {
      address acc = chain.users.list[j];
      // Get validators
      if(validators == true) {
        // if active == true, get only those validators who are also mining
        // if active == false, get those who are allowed to mine based on their vesting
        if (chain.users.accounts[acc].validator.mining == active) {
          users[count] = acc;
          count++;
        }
      }
      // Get transactors
      else {
        if (chain.users.accounts[acc].transactor.whitelisted == true) {
          users[count] = acc;
          count++;
        } 
      }
      j++;
     }
     
     if (j == usersTotalCount) {
         end = true;
     }
     else {
         end = false;
     }
  }
      
  function _startMining(uint256 chainId, address acc) internal {      
      ChainInfo storage chain = chains[chainId];
      Validator storage validator = chain.users.accounts[acc].validator;
      
      if (validator.mining == false) {
          chain.totalVesting = chain.totalVesting.add(validator.vesting);
      }
      validator.mining = true;
      
      emit StartMining(chainId, acc);
  }
      
  function _stopMining(uint256 chainId, address acc) internal {      
      ChainInfo storage chain = chains[chainId];
      Validator storage validator = chain.users.accounts[acc].validator;
      
      if (validator.mining == true) {
          chain.totalVesting = chain.totalVesting.sub(validator.vesting);
      }
      validator.mining = false;
      
      emit StopMining(chainId, acc);
  }
}

// SafeMath library from: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/math/SafeMath.sol
library SafeMath {
    /**
     * @dev Returns the addition of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `+` operator.
     *
     * Requirements:
     * - Addition cannot overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return sub(a, b, "SafeMath: subtraction overflow");
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting with custom message on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     * - Subtraction cannot overflow.
     *
     * NOTE: This is a feature of the next version of OpenZeppelin Contracts.
     * @dev Get it via `npm install @openzeppelin/contracts@next`.
     */
    function sub(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b <= a, errorMessage);
        uint256 c = a - b;

        return c;
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `*` operator.
     *
     * Requirements:
     * - Multiplication cannot overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");

        return c;
    }

    /**
     * @dev Returns the integer division of two unsigned integers. Reverts on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return div(a, b, "SafeMath: division by zero");
    }

    /**
     * @dev Returns the integer division of two unsigned integers. Reverts with custom message on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     * - The divisor cannot be zero.
     * NOTE: This is a feature of the next version of OpenZeppelin Contracts.
     * @dev Get it via `npm install @openzeppelin/contracts@next`.
     */
    function div(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0, errorMessage);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }
}
