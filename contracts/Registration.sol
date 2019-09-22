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
    
    /**** Events related to the deposit requests ****/
    // Deposit request created
    event RequestDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp /* now */);
    // Deposit request confirmed - tokens were transferred and account's deposit balance was updated
    event ConfirmDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp /* request creation time */);
    // Deposit request cancelled
    event CancelDepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 reqTimestamp /* request creation time */);
    // Whole deposit withdrawned - this is special withdrawal, which as applied after the chain validators are not able to reach consensus for 1 month 
    event ForceWithdrawDeposit(uint256 indexed chainId, address indexed account, uint256 timestamp /* now */);
    
    /**** Events related to the vesting requests ****/
    // Vesting request created
    event RequestVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp /* now */);
    // Vesting request confirmed - tokens were transferred
    event ConfirmVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp /* request creation time */);
    // Vesting request cancelled
    event CancelVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp /* request creation time */);
    // Vesting request accepted - account's vesting balance was updated 
    event AcceptedVestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 reqTimestamp /* request creation time */);
    // Whole vesting withdrawned - this is special withdrawal, which as applied after the chain validators are not able to reach consensus for 1 month 
    event ForceWithdrawVesting(uint256 indexed chainId, address indexed account, uint256 timestamp /* now */);
    
    // if whitelist == true  - allow user to transact
    // if whitelist == false - do not allow user to transact
    event WhitelistAccount(uint256 indexed chainId, address miner, bool whitelist);
    
    // Validator start/stop mining
    event StartMining(uint256 indexed chainId, address miner);
    event StopMining(uint256 indexed chainId, address miner);


    /**************************************************************************************************************************/
    /***************************************** Structs related to the list of users *******************************************/
    /**************************************************************************************************************************/
    struct Validator {
        // Actual user's vesting
        uint96  vesting;
        // Flag if user is mining -> set in start/stopMining
        bool    mining;
    }
    
    struct Transactor {
        // Actual user's deposit
        uint96  deposit;
        // Flag if user is whitelisted (allowed to transact) -> actual deposit must be greater than min. required deposit condition 
        bool    whitelisted;
    }
    
    // Optimized "UserEntry" so it packs together with User_details to 256 Bits (32 Bytes) chunk of memory
    struct UserEntry {
        // Validator's data
        Validator    validator;
        // Transactor's data
        Transactor   transactor;
        // index to the usersList, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing user
        uint48       index;
    }
    
    struct Users {
        mapping(address => UserEntry) accounts;
        address[]                     list;        
    }
    
    
    /**************************************************************************************************************************/
    /*********************************** Structs related to the vesting/deposit requests **************************************/
    /**************************************************************************************************************************/
    enum RequestState {
        REQUEST_CREATED,
        REQUEST_CONFIRMED
    }
    
    enum VestingRequestControlState {
        // Wait for vestingRequest_confirm to be called. It is used when newVesting > actual validator's vesting
        WAIT_FOR_CONFIRMATION,
        // Validator's actual vesting will be replaced by the newVesting from VestingRequest after the next notary
        REPLACE_VESTING,
        // Validator's actual vesting was replaced by the newVesting from VestingRequest after the previous notary 
        VESTING_REPLACED
    }
    
    // 512 Bits
    struct VestingRequestData {
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
        // Timestamp(now) when the request was accepted/created
        uint48                          timestamp;
        // In oldVesting is stored current vesting that validator had when new VestingRequest was accepted
        uint96                          oldVesting;
        // New value of vesting to be set
        uint96                          newVesting;
        // Actual state of the request
        RequestState                   state;
        // Actual control state of the request
        VestingRequestControlState      controlState;
    }
    
    // Only full deposit withdrawals are saved as deposit requests - other types of deposits do not need to be confirmed
    // 312 Bits
    struct DepositWithdrawRequestData {
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
        // Timestamp(now) when the request was accepted/created
        uint48                          timestamp;
        // Actual state of the request
        RequestState                   state;
    }
    
    // 872 Bits
    struct RequestsEntry {
        // index to the requestsList, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing request
        uint48                          index;
        // Deposit withdrawal request details
        DepositWithdrawRequestData     depositWithdrawRequest;
        // Vesting request details
        VestingRequestData             vestingRequest;   
    }
    
    struct Requests {
        mapping(address => RequestsEntry)  accounts;
        address[]                           list;        
    }
    
    
    /**************************************************************************************************************************/
    /***************************************************** Other structs ******************************************************/
    /**************************************************************************************************************************/
    ERC20 token;
   
    // struct Signature {
    //     uint8 v; bytes32 r; bytes32 s;
    // }
    
    struct LastNotary {
        uint256 timestamp;
        // Actual block number, when the last notary was accepted
        uint256 block;
    }
    
    struct ChainInfo {
        bool              registered;
        bool              active;
        uint256           totalVesting;
        LastNotary        lastNotary;
        ChainValidator    chainValidator;
        Users             users;
        Requests          requests;
        string            endpoint;
    }
    
    mapping(uint256 => ChainInfo) private chains;
    uint256 public nextId = 0;

    
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
          require(validatorExists(chainId, msg.sender) == true, "Trying to withdraw vesting from non-existing validator account");
          
          // If last notary is older than 30 days, it means that validators cannot reach consensus and side-chain is basically stuck.
          // In such case ignore multi-step vesting process and allow users to withdraw all vested tokens
          if (chain.lastNotary.timestamp + 30 days < now) {
              forceWithdrawVestFromChain(chainId, msg.sender);
              return;
          }
          
          require(chain.users.accounts[msg.sender].validator.mining == false, "Can't withdraw any tokens, stop_minig must be called first.");  
      }
      // Vest in chain or withdraw just part of vesting
      else {
         require(vesting <= ~uint96(0), "vesting is greater than uint96_max_value");
         require(chain.users.accounts[msg.sender].validator.vesting != vesting, "Cannot vest the same amount of tokens as you already has vested.");
         require(checkLitionMinVesting(vesting), "user does not meet Lition's min.required vesting condition");
         require(chain.chainValidator.checkVesting(vesting, msg.sender), "user does not meet chain validator's min.required vesting condition");
      }
      
      require(vestingRequestExists(chainId, msg.sender) == false, "Cannot vest in chain. There is already ongoing request being processed for this acc.");
        
      _requestVestInChain(chainId, vesting, msg.sender);
    }
    
    // Confirms vest request, token transfer is processed during confirmation
    function confirmVestInChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
        
        require(vestingRequestExists(chainId, msg.sender) == true, "Cannot confirm non-existing vesting request.");
        require(chain.lastNotary.block > chains[chainId].requests.accounts[msg.sender].vestingRequest.notaryBlock, "Request confirmation can be called in the next notary window after request was accepted.");
        require(chain.requests.accounts[msg.sender].vestingRequest.state == RequestState.REQUEST_CREATED, "Cannot confirm already confirmed request.");
        
        _confirmVestInChain(chainId, msg.sender);
    }
    
    // Cancels the existing vest request. Such request can be cancelled only if it was not already confirmed
    function cancelVestInChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
      
        require(vestingRequestExists(chainId, msg.sender) == true, "Cannot cancel non-existing vesting request.");
        require(chain.requests.accounts[msg.sender].vestingRequest.state == RequestState.REQUEST_CREATED, "Cannot cancel already confirmed request." );
        
        _cancelVestInChain(chainId, msg.sender);
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
          require(transactorExists(chainId, msg.sender) == true, "Non-existing transactor account");
          
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
        
        require(depositWithdrawRequestExists(chainId, msg.sender) == false, "Cannot deposit in chain. There is ongoing withdrawal request being processed for this acc.");
                
        _requestDepositInChain(chainId, deposit, msg.sender);
    }
    
    // Confirms deposit withdrawal request, token transfer is processed during confirmation
    function confirmDepositWithdrawalFromChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
        
        require(depositWithdrawRequestExists(chainId, msg.sender) == true, "Cannot confirm non-existing deposit withdrawal request.");
        require(chain.lastNotary.block > chains[chainId].requests.accounts[msg.sender].depositWithdrawRequest.notaryBlock, "Request confirmation can be called in the next notary window after request was accepted.");
        require(chain.requests.accounts[msg.sender].depositWithdrawRequest.state == RequestState.REQUEST_CREATED, "Cannot confirm already confirmed request.");
        
        _confirmDepositWithdrawalFromChain(chainId, msg.sender);
    }
    
    // Cancels the existing deposit request. Such request can be cancelled only if it was not already confirmed
    function cancelDepositInChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.active == true, "Non-active chain");
        
        require(depositWithdrawRequestExists(chainId, msg.sender) == true, "Cannot cancel non-existing deposit withdrawal request.");
        require(chain.requests.accounts[msg.sender].depositWithdrawRequest.state == RequestState.REQUEST_CREATED, "Cannot cancel already confirmed request." );
        
        cancelDepositWithdrawalFromChain(chainId, msg.sender);
    }
    
    // Internally creates/registers new side-chain. Creator must be also validator at least from the beginning as joining process take multiple steps
    // and these steps cannot be done in the same notary window
    function registerChain(string calldata info, ChainValidator validator, uint96 vesting, uint96 deposit, string calldata initEndpoint) external returns (uint256 chainId) {
        require(bytes(initEndpoint).length > 0);
        require(deposit <= ~uint96(0), "deposit is greater than uint96_max_value");
        require(vesting <= ~uint96(0), "vesting is greater than uint96_max_value");
        
        address creator         = msg.sender;
        uint256 timestamp       = now;
        
        // Inits chain data
        chainId                = nextId;
        ChainInfo storage chain = chains[chainId];
        
        chain.chainValidator   = validator;
        
        // Validates vesting
        require(checkLitionMinVesting(vesting), "chain creator does not meet Lition's min.required vesting condition");
        require(chain.chainValidator.checkVesting(vesting, creator), "chain creator does not meet chain validator's min.required vesting condition");
        
        // Validates deposit
        require(checkLitionMinDeposit(deposit), "chain creator does not meet Lition's min.required deposit condition");
        require(chain.chainValidator.checkDeposit(deposit, creator), "chain creator does not meet chain validator's min.required deposit condition");
        
        // Internally creates new user
        userCreate(chainId, creator);
        
        // Transfers vesting tokens
        token.transferFrom(creator, address(this), vesting);
        chain.users.accounts[creator].validator.vesting = vesting;      
        
        // Transfers deposit tokens
        token.transferFrom(creator, address(this), deposit);
        chain.users.accounts[creator].transactor.deposit = deposit;      
        chain.users.accounts[creator].transactor.whitelisted = true;
        
        chain.registered        = true;
        chain.endpoint          = initEndpoint;
        
        emit NewChain(chainId, info, initEndpoint);
        
        emit RequestVestInChain(chainId, creator, vesting, timestamp);
        emit ConfirmVestInChain(chainId, creator, vesting, timestamp);
        emit AcceptedVestInChain(chainId, creator, vesting, timestamp);
        
        emit RequestDepositInChain(chainId, creator, deposit, timestamp);
        emit ConfirmDepositInChain(chainId, creator, deposit, timestamp);
        
        nextId++;
    }
    
    // Reurns true, if acc has vested enough to become validator, othervise false
    function hasVested(uint256 chainId, address acc) view external returns (bool) {
        // No need to check vesting balance as it cannot be lover than min. required
        return validatorExists(chainId, acc);
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
    
    function testNotary(uint256 chainId, uint256 notaryBlockNo) external {
        ChainInfo storage chain = chains[chainId];
        require(chain.registered, "Non-registered chain");
        
        // Process existing vesting/deposit withdrawals requests as these cannot be 
        // processed automatically - math for user's usage and miner's rewards calculations would be invalid
        processExistingRequests(chainId);
        
        // TODO: remove validators(call stopMining) who signed no block during this notary window and have mining flag == true
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryBlockNo;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
    }
    

    function getChainDetails(uint256 chainId) external view returns (bool registered, bool active, string memory endpoint, uint256 totalVesting,
                                                                     uint256 lastNotaryBlock, uint256 lastNotaryTimestamp) {
        ChainInfo storage chain = chains[chainId];
        
        registered          = chain.registered;
        active              = chain.active;
        endpoint            = chain.endpoint;
        totalVesting        = chain.totalVesting;
        lastNotaryBlock     = chain.lastNotary.block;
        lastNotaryTimestamp = chain.lastNotary.timestamp;
    }
    
    function getUserDetails(uint256 chainId, address acc) external view returns (bool exists, uint256 deposit, bool whitelisted, uint256 vesting, bool mining) {
        ChainInfo storage chain = chains[chainId];
        
        exists = chain.users.accounts[acc].index > 0;
         
        deposit = chain.users.accounts[acc].transactor.deposit;
        whitelisted = chain.users.accounts[acc].transactor.whitelisted;
        
        vesting = chain.users.accounts[acc].validator.vesting;
        mining = chain.users.accounts[acc].validator.mining;  
    }
    
    function getUserRequests(uint256 chainId, address acc) external view returns (bool vestingReqExists, uint256 vestingReqTime, uint256 vestingReqNotary, uint256 vestingReqValue, uint256 vestingReqState, uint256 vestingReqControlState,
                                                                                     bool depositReqExists, uint256 depositReqTime, uint256 depositReqNotary, uint256 depositReqValue, uint256 depositReqState) {
        if (vestingRequestExists(chainId, acc)) {
            VestingRequestData storage request = chains[chainId].requests.accounts[acc].vestingRequest;
            
            vestingReqExists        = true;
            vestingReqTime          = request.timestamp;
            vestingReqNotary        = request.notaryBlock;
            vestingReqValue         = request.newVesting;
            vestingReqState         = uint256(request.state);
            vestingReqControlState  = uint256(request.controlState);
        }
        
        if (depositWithdrawRequestExists(chainId, acc)) {
            DepositWithdrawRequestData storage request = chains[chainId].requests.accounts[acc].depositWithdrawRequest;
            
            depositReqExists = true;
            depositReqTime   = request.timestamp;
            depositReqNotary = request.notaryBlock;
            depositReqValue  = 0;
            depositReqState  = uint256(request.state);
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
        require(notaryStartBlock    >  chain.lastNotary.block, "Invalid data: notaryBlock_start from statistics must be greater than the last known notary block");
        require(notaryEndBlock      >  notaryStartBlock,      "Invalid data: notaryEndBlock must be greater than notaryStartBlock");
        require(largestTx           >  0,                       "Invalid data: Largest tx must be greater than zero");
        require(miners.length       == blocksMined.length,     "Invalid data: num of miners != num of block mined");
        require(users.length        == userGas.length,         "Invalid data: num of users != num of users gas");
        
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
        
        // Process existing vesting/deposit withdrawals requests as these cannot be 
        // processed automatically - math for user's usage and miner's rewards calculations would be invalid
        processExistingRequests(chainId);
        
        // TODO: remove validators(call stopMining) who signed no block during this notary window and have mining flag == true
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryEndBlock;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
    }
    
    
    // Returns list of user's addresses that are allowed to transact - their deposit >= min. required deposit
    function getAllowedToTransact(uint256 chainId, uint256 batch) view external returns (address[100] memory, uint256, bool) {
        return getUsers(chainId, batch, false, false);
    }
    
    // Returns list of validator's addresses that are actively participating in mining
    function getAllowedToValidate(uint256 chainId, uint256 batch) view external returns (address[100] memory, uint256, bool) {
        return getUsers(chainId, batch, true, false);
    }
    
    // Returns list of validator's addresses that are allowed to be participating in mining based on their vesting, but not yet mining
    function getActiveValidators(uint256 chainId, uint256 batch) view external returns (address[100] memory, uint256, bool) {
        return getUsers(chainId, batch, true, true);
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
    
    // Creates new user - does not set it's data yet as it is done after vesting/deposit_withdraw request is confirmed
    function userCreate(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        require(chain.users.list.length < ~uint48(0), "count(users) is equal to max_count");
        
        chain.users.list.push(acc);
        chain.users.accounts[acc].index = uint48(chains[chainId].users.list.length); // indexes are stored + 1
    }
    
    // Deletes existing user from the internal list of users
    // This method should never be called directly, validatorDelete & transactorDelete should be called instead
    function userDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        UserEntry storage entry = chain.users.accounts[acc];
        // userExists(chainId, acc) could be used instead
        require(entry.index != 0, "User does not exist");
        
        address[] storage usersList = chain.users.list;
        require(entry.index <= usersList.length, "Invalid index value");
    
        // Move an last element of array into the vacated key slot.
        uint48 foundIndex = uint48(entry.index - 1);
        uint48 lastIndex = uint48(usersList.length - 1);
    
        chain.users.accounts[usersList[lastIndex]].index = foundIndex + 1;
        usersList[foundIndex] = usersList[lastIndex];
        usersList.length--;
    
        delete chains[chainId].users.accounts[acc];
    }
    
    // Deletes validator
    function validatorDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no existing transactor for this account - delete whole requests struct 
        if (chain.users.accounts[acc].transactor.deposit == 0) {
            userDelete(chainId, acc);
            return;
        } 
        
        // There is exiting transactor for this account - only reset validator
        Validator storage validator = chain.users.accounts[acc].validator;
        validator.vesting   = 0;
        validator.mining    = false;
    }
    
    // Deletes transactor
    function transactorDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no existing validator for this account - delete whole requests struct 
        if (chain.users.accounts[acc].validator.vesting == 0) {
            userDelete(chainId, acc);
            return;
        } 
        
        // There is existing validator for this account - only reset transactor
        Transactor storage transactor = chain.users.accounts[acc].transactor;
        transactor.deposit      = 0;
        transactor.whitelisted  = false;
    }
    
    function userExists(uint256 chainId, address acc) internal view returns (bool) {
      return chains[chainId].users.accounts[acc].index != 0;
    }
    
    function validatorExists(uint256 chainId, address acc) internal view returns (bool) {
      return chains[chainId].users.accounts[acc].validator.vesting > 0;
    }
    
    function transactorExists(uint256 chainId, address acc) internal view returns (bool) {
      return chains[chainId].users.accounts[acc].transactor.deposit > 0;
    }
    
    
    
    /**************************************************************************************************************************/
    /*********************************** Functions related to the vesting/deposit requests ************************************/
    /**************************************************************************************************************************/
    
    // Creates new vesting request
    function vestingRequestCreate(uint256 chainId, address acc, uint256 vesting) internal {
        ChainInfo storage chain = chains[chainId];
        
        require(chain.requests.list.length < ~uint48(0), "count(requests) is equal to max_count");
        
        RequestsEntry storage entry = chain.requests.accounts[acc];
        
        entry.vestingRequest.oldVesting = chain.users.accounts[acc].validator.vesting;
        entry.vestingRequest.newVesting = uint96(vesting);
        if (entry.vestingRequest.newVesting >= entry.vestingRequest.oldVesting) { // == case should never happen as it is handled in the caller's function
            entry.vestingRequest.controlState = VestingRequestControlState.WAIT_FOR_CONFIRMATION;
        } else {
            entry.vestingRequest.controlState = VestingRequestControlState.REPLACE_VESTING;
        }
        
        entry.vestingRequest.state = RequestState.REQUEST_CREATED;
        entry.vestingRequest.timestamp = uint48(now);
        entry.vestingRequest.notaryBlock = chain.lastNotary.block; 
        
        
        // There is no deposit or vesting ongoing request - create new RequestsEntry structure
        if (entry.index == 0) { // anyRequestExists(chainId, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chain.requests.list.push(acc);    
            entry.index = uint48(chain.requests.list.length); // indexes are stored + 1
        }
    }

    // Creates new deposit withdrawal request
    function depositWithdrawRequestCreate(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        require(chain.requests.list.length < ~uint48(0), "count(requests) is equal to max_count");
        
        RequestsEntry storage entry = chain.requests.accounts[acc];
        
        entry.depositWithdrawRequest.state = RequestState.REQUEST_CREATED;
        entry.depositWithdrawRequest.timestamp = uint48(now);
        entry.depositWithdrawRequest.notaryBlock = chain.lastNotary.block; 
        
        // There is no deposit or vesting ongoing request - create new RequestsEntry structure
        if (entry.index == 0) { // anyRequestExists(chainId, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chain.requests.list.push(acc);    
            entry.index = uint48(chain.requests.list.length); // indexes are stored + 1
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
        uint48 foundIndex = uint48(entry.index - 1);
        uint48 lastIndex = uint48(requestsList.length - 1);
    
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
        VestingRequestData storage request = chain.requests.accounts[acc].vestingRequest;
        request.notaryBlock    = 0;
        request.timestamp      = 0;
        request.oldVesting     = 0;
        request.newVesting     = 0;
        // First enum value is default
        request.state          = RequestState.REQUEST_CREATED; 
        request.controlState   = VestingRequestControlState.WAIT_FOR_CONFIRMATION;
    }
    
    function depositWithdrawRequestDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no ongoing vesting request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].vestingRequest.timestamp == 0) {
            requestsPairDelete(chainId, acc);
            return;
        } 
        
        // There is ongoing vesting request for this account - only reset vesting request
        DepositWithdrawRequestData storage request = chain.requests.accounts[acc].depositWithdrawRequest;
        request.notaryBlock    = 0;
        request.timestamp      = 0;
        // First enum value is default
        request.state          = RequestState.REQUEST_CREATED; 
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
      // Internally creates new user
      if (vesting != 0 && userExists(chainId, acc) == false) {
          userCreate(chainId, acc);
      }
      
      vestingRequestCreate(chainId, acc, vesting);
      emit RequestVestInChain(chainId, acc, vesting, now);
    }
    
    // Basically just transfers the tokens, validator's vesting balance update is always done at the of notary atomatically
    function _confirmVestInChain(uint256 chainId, address acc) internal {
        VestingRequestData storage request = chains[chainId].requests.accounts[acc].vestingRequest;
        
        request.state = RequestState.REQUEST_CONFIRMED;
        
        // Decreases account's vesting in chain
        if(request.newVesting < request.oldVesting) {
            // This should never happen during normal conditions as vesting balance state is updated during notary
            require(request.controlState == VestingRequestControlState.VESTING_REPLACED, "Cannot withdraw vesting tokens, internal balance was not updated yet");
            
            uint96 toWithdraw = request.oldVesting - request.newVesting;
            
            // If it was request to withdraw whole vesting balance, delete validator
            if (request.newVesting == 0) {
                validatorDelete(chainId, acc);
            }
            
            emit ConfirmVestInChain(chainId, acc, request.newVesting, request.timestamp);
            vestingRequestDelete(chainId, acc);
            
            token.transfer(acc, toWithdraw);
            
            return;
        }
        
        // Increases account's vesting in chain
        uint96 toVest = request.newVesting - request.oldVesting;
        token.transferFrom(acc, address(this), toVest);
        
        request.controlState = VestingRequestControlState.REPLACE_VESTING;
        
        emit ConfirmVestInChain(chainId, acc, request.newVesting, request.timestamp);
    }
    
    function _cancelVestInChain(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        VestingRequestData storage request = chain.requests.accounts[acc].vestingRequest;
        Validator storage validator = chain.users.accounts[acc].validator;
        
        // Replace back the original validator's vesting
        if (request.controlState == VestingRequestControlState.VESTING_REPLACED) {
            validator.vesting = request.oldVesting;
            
            // If validator is actively mining, updates chain totalVesting
            if (validator.mining == true) {
                chain.totalVesting = chain.totalVesting.sub(request.newVesting - request.oldVesting);
            }
        }
        
        emit CancelVestInChain(chainId, acc, request.newVesting, request.timestamp);
        vestingRequestDelete(chainId, acc);
    }
    
    // Forcefully withdraw whole vesting from chain.
    // Because vesting is processed during 2(new_vest < act_vest) or even 3(new_vest > act_vest) notary windows,
    // user might end up with locked tokens in SC in case validators never reach consesnsus. In such case these tokens stay locked in
    // SC for 1 month and after that can be withdrawned. Any existing vest requests are deleted after this withdraw.
    function forceWithdrawVestFromChain(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        uint96 toWithdraw = 0;
        bool requestExists = vestingRequestExists(chainId, acc);
        
        // No ongoing vesting request is present
        if (requestExists == false) {
            toWithdraw = chain.users.accounts[acc].validator.vesting;
            chain.totalVesting = chain.totalVesting.sub(toWithdraw);
        }
        // There is ongoing vesting request
        else { 
            VestingRequestData storage request = chain.requests.accounts[acc].vestingRequest;
            // Token transfer was already processed -> use new vesting balance as actual user's vesting balance to withdraw
            if (request.state == RequestState.REQUEST_CONFIRMED) {
                toWithdraw = request.newVesting;
            }
            // Token transfer was not yet processed -> use saved old vesting balance as actual user's vesting balance to withdraw
            else {
                toWithdraw = request.oldVesting;
            }
            
            // If validor is actively mining update chain's totalVesting
            if (chain.users.accounts[acc].validator.mining == true) {
                // Vesting balance and chain's totalVesting were already internally updated
                if (request.controlState == VestingRequestControlState.VESTING_REPLACED) {
                    chain.totalVesting = chain.totalVesting.sub(request.newVesting);
                }
                // Vesting balance and chain's totalVesting were not yet internally updated
                else {
                    chain.totalVesting = chain.totalVesting.sub(request.oldVesting);
                }
            }
            
            vestingRequestDelete(chainId, acc);
        }
        
        validatorDelete(chainId, acc);
        
        // Transfers all remaining tokens to the user account
        token.transfer(acc, toWithdraw);
        
        // Stops mining
        emit StopMining(chainId, acc);
        
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
    
    function cancelDepositWithdrawalFromChain(uint256 chainId, address acc) internal {
        DepositWithdrawRequestData storage request = chains[chainId].requests.accounts[acc].depositWithdrawRequest;
        Transactor storage transactor = chains[chainId].users.accounts[acc].transactor;
        
        // If withdrawal was cancelled and transactor has >= min. required deposit, re-enable user to transact
        if (transactor.whitelisted == false && 
            checkLitionMinDeposit(transactor.deposit) == true &&
            chains[chainId].chainValidator.checkDeposit(transactor.deposit, acc) == true) {
                
          transactor.whitelisted = true;
          emit WhitelistAccount(chainId, acc, true);
        }
        
        depositWithdrawRequestDelete(chainId, acc);
        
        emit CancelDepositInChain(chainId, acc, 0, request.timestamp);
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

    // Process existing vesting/deposit withdrawals requests as these cannot be 
    // processed automatically - math for user's usage and miner's rewards calculations would be invalid
    function processExistingRequests(uint256 chainId) internal {
        ChainInfo storage chain = chains[chainId];
        Requests storage requests = chain.requests;
        
        // Runs through all existing requests
        for (uint256 i = 0; i < requests.list.length; i++) {
            address acc = requests.list[i];
            RequestsEntry storage entry = requests.accounts[acc];
            UserEntry storage user = chain.users.accounts[acc];
            
            // There is existing vesting request - process it 
            if (entry.vestingRequest.timestamp != 0) {
                // Current validator's vesting balance should be replaced during this notary
                if (entry.vestingRequest.controlState == VestingRequestControlState.REPLACE_VESTING) {
                    user.validator.vesting = entry.vestingRequest.newVesting;
                    entry.vestingRequest.controlState = VestingRequestControlState.VESTING_REPLACED;
                    
                    // If validator is actively mining, updates also chain's total vesting
                    if (user.validator.mining == true) {
                        chains[chainId].totalVesting = chains[chainId].totalVesting.add(entry.vestingRequest.newVesting - entry.vestingRequest.oldVesting);
                    }
                    
                    emit AcceptedVestInChain(chainId, acc, entry.vestingRequest.newVesting, entry.vestingRequest.timestamp);
                    
                    // If it was request to increase validator's vesting balance and we got here, it means we can delete this request.
                    // Requests to decrease vesting balance are deleted in confirmation
                    if (entry.vestingRequest.newVesting > entry.vestingRequest.oldVesting) {
                        vestingRequestDelete(chainId, acc);
                    }
                }
            }
            
            // There is existing deposit withdrawal request - process it
            // If there is such request and user is still allowed to transact(might be not because he runs out of tokens during usage calculations), 
            // do not allow him to transact anymore
            if (entry.depositWithdrawRequest.timestamp != 0 && user.transactor.whitelisted == true) {
                // User wants to withdraw whole deposit - do not allow him to transact anymore, token transfer is processed when he confirms withdrawal
                user.transactor.whitelisted = false;
                emit WhitelistAccount(chainId, acc, false);
            }
        }
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
