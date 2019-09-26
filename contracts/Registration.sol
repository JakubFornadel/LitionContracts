pragma solidity >= 0.5.11;

interface ChainValidator {
   function checkVesting(uint256 vesting, address acc, uint256 actNumOfValidators) external returns (bool);
   function checkDeposit(uint256 vesting, address acc, uint256 actNumOfTransactors) external returns (bool);
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
   function checkVesting(uint256 vesting, address acc, uint256 actNumOfValidators) public returns (bool) {
      if(vesting >= 1000*(10**18) && vesting <= 500000*(10**18)) {
        return true;   
      }
      return false;
   }

   function checkDeposit(uint256 deposit, address acc, uint256 actNumOfTransactors) public returns (bool) {
      if(deposit >= 1000*(10**18)) {
         return true;
      }
      return false;
   }
}

contract LitionRegistry {
    using SafeMath for uint256;
    
    /**************************************************************************************************************************/
    /************************************************** Constants *************************************************************/
    /**************************************************************************************************************************/
    
    // Token precision. 1 LIT token = 1*10^18
    uint256 constant LIT_PRECISION               = 10**18;
    
    // Max deposit value
    uint256 constant MAX_DEPOSIT                 = ~uint96(0);
    
    // Max vesting value
    uint256 constant MAX_VESTING                 = ~uint96(0);
    
    // Max number of active validators allowed to be registered in one chain is set by chain creator
    // This is upper limit of what creator can specify
    uint256 constant MAX_NUM_OF_VALIDATORS_LIMIT = 101;
    
    // Max number of unique users allowed to be processed during notary
    uint256 constant MAX_NUM_OF_PROCESSED_USERS  = 500;
    
    // Time after which chain becomes inactive in case there was no successfull notary processed
    // Users can then increase their vesting instantly and bypass 2-step process. For instant full deposit withdrawal users must wait *CHAIN_INACTIVITY_TIMEOUT
    uint256 constant CHAIN_INACTIVITY_TIMEOUT    = 7 days;
    
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function checkLitionMinVesting(uint256 vesting) private pure returns (bool) {
        if(vesting >= 1000*LIT_PRECISION) {
            return true;   
        }
        return false;
    }
    
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function checkLitionMinDeposit(uint256 deposit) private pure returns (bool) {
        if(deposit >= 1000*LIT_PRECISION) {
            return true;   
        }
        return false; 
    }
    
    
    /**************************************************************************************************************************/
    /**************************************************** Events **************************************************************/
    /**************************************************************************************************************************/
    
    // New chain was registered
    event NewChain(uint256 chainId, string description, string endpoint);
    
    // Deposit request created
    // in case confirmed == true, tokens were transferred and account's deposit balance was updated
    // in case confirmed == false, listener needs to wait for another DepositInChain event with confirmed flag == true.
    // It can paired up with the first event if first 4 parameters are the same
    event DepositInChain(uint256 indexed chainId, address indexed account, uint256 deposit, uint256 lastNotaryBlock, bool confirmed);
    
    // Vesting request created
    // in case confirmed == true, tokens were transferred and account's vesting balance was updated
    // in case confirmed == false, listener needs to wait for another VestInChain event with confirmed flag == true.
    // It can paired up with the first event if first 4 parameters are the same
    event VestInChain(uint256 indexed chainId, address indexed account, uint256 vesting, uint256 lastNotaryBlock, bool confirmed);
    
    // if whitelist == true  - allow user to transact
    // if whitelist == false - do not allow user to transact
    event AccountWhitelist(uint256 indexed chainId, address indexed account, bool whitelist);
    
    // Validator start/stop mining
    event AccountMining(uint256 indexed chainId, address indexed account, bool mining);

    // New notary was processed
    // in case confirmed == false, notary processing just stared and listener needs to wait for another event with confirmed flag == true.
    // in case confirmed == true, notary was processed successfully
    // It can paired up with the first event if first 4 parameters are the same
    // If the second event with confirmed flag set to true is not emmited in reasonable time, it means notary failed
    event Notary(uint256 indexed chainId, uint256 notaryBlock, bool confirmed);

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
        // Flag if validator mined at least 1 block in current notary window
        bool    currentNotaryMined;
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
    
    struct VestingIncreaseRequest {
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
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
        VestingIncreaseRequest          vestingIncreaseRequest;   
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
        // Chain description
        string                          description;
        
        // Chain endpoint
        string                          endpoint;
        
        // Flag that says if the chain was already(& sucessfully) registered
        bool                            registered;
        
        // Flag that says if chain is active - num of active validators > 0 && first block was already mined 
        bool                            active;
        
        // How much vesting in total is vested in the chain by the active validators
        uint256                         totalVesting;
        
        // When was the last notary function called (block and time)
        LastNotary                      lastNotary;
        
        // Smart-contract for validating min.required deposit and vesting
        ChainValidator                  chainValidator;
        
        // Mapping data holder for users (validators & transactors) data
        mapping(address => User)        users;
        
        // List of existing validators - only those who vested enough and are actively mining are here
        IterableMap                     validators;
        
        // Validator with the smallest vesting balance among all existing validators
        // In case someone vests more tokens, he will replace the one with smallest vesting balance
        address                         lastValidator;
        
        // List of existing transactors - all users, who have deposited (even if they cannot transact anymore)
        IterableMap                     transactors;
        
        // Actual number of whitelisted transactors (their current depost > min.required deposit)
        uint256                         actNumOfTransactors;
        
        // List of users requests for changing their vesting/deposit balance
        Requests                        requests;
        
        // Max number of active validators for chain. There is no limit to how many users can vest.
        // Tthis is limit for active validators (startMining function) 
        uint256                         maxNumOfValidators;
        
        // Max number of users(transactors) for chain.
        // Tthis is limit for whitelisted transactors (their current depost > min.required deposit)
        uint256                         maxNumOfTransactors;
        
        // Flag for condition to be checked during notary:
        // InvolvedVesting of validators who signed notary statistics must be greater than 50% of chain total vesting(sum of all active validator's vesting)
        // to notary statistics to be accepted is accepted
        bool                            involvedVestingNotaryCond;
        
        // Flag for condition to be checked during notary:
        // Number of validators who signed notary statistics must be greater or equal than 2/3+1 of all active validators
        // to notary statistics to be accepted is accepted
        bool                            participationNotaryCond;
    }
    
    // Registered chains
    mapping(uint256 => ChainInfo)   private chains;
    
    // Next chain id
    uint256                         public  nextId = 0;

    
    /**************************************************************************************************************************/
    /************************************ Contract interface - external/public functions **************************************/
    /**************************************************************************************************************************/
    
    // Requests vest in chain. It will be processed and applied to the actual user state after next:
    //      * 1 notary windows after the confirm method is called - in case new vesting > actual vesting
    //      * emmidiately - in case new vesting < actual vesting
    //
    // In case new vesting > actual vesting, user first creates request, tokens are transferred to the sc, then he must confirm this request in the next notary window 
    // and his internal vesting balance is updated 
    function requestVestInChain(uint256 chainId, uint256 vesting) external {
        ChainInfo storage chain = chains[chainId];
        
        // Enable users to vest into registered (but non-active) chain and start minig so it becomes active
        require(chain.registered == true, "Non-registered chain");
        require(vestingIncreaseRequestExist(chainId, msg.sender) == false, "There is already existing request being processed for this acc.");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chainId);
            
        // Withdraw all vesting
        if (vesting == 0) {
            require(validatorExist(chainId, msg.sender) == true, "Non-existing validator account (0 vesting balance)");
            require(activeValidatorExist(chainId, msg.sender) == false, "StopMinig must be called first.");  
        }
        // Vest in chain or withdraw just part of vesting
        else {
            require(vesting <= MAX_VESTING, "vesting is greater than uint96_max_value");
            require(chain.users[msg.sender].validator.vesting != vesting, "Cannot vest the same amount of tokens as you already has vested.");
            require(checkLitionMinVesting(vesting), "user does not meet Lition's min.required vesting condition");
            require(chain.chainValidator.checkVesting(vesting, msg.sender, chain.validators.list.length), "user does not meet chain validator's min.required vesting condition");
        }
        
        _requestVestInChain(chainId, vesting, msg.sender);
    }
    
    // Confirms vesting increase request
    function confirmVestIncreaseInChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        
        // Enable users to confirm vesting request into registered (but non-active) chain and start minig so it becomes active
        require(chain.registered == true, "Non-registered chain");
        require(vestingIncreaseRequestExist(chainId, msg.sender) == true, "Non-existing vesting request.");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chainId);
        
        // Chain is active
        if (chain.active == true) {
            require(chain.lastNotary.block > chain.requests.accounts[msg.sender].vestingIncreaseRequest.notaryBlock, "Confirm can be called in the next notary window after request was accepted.");    
        }
        
        _confirmVestIncreaseInChain(chainId, msg.sender);
    }
    
    // Requests deposit in chain. It will be processed and applied to the actual user state after next:
    //      * 1 notary windows after the confirm method is called - in case deposit == 0
    //      * emmidiately                                         - in case deposit != 0
    // Only full deposit withdrawals needs to be handled in 2 steps as it would allow users to send unlimited amount of txs to the sidechain and 
    // withdraw whole deposit right before notary function, in which case he would pay nothing...
    function requestDepositInChain(uint256 chainId, uint256 deposit) external {
        ChainInfo storage chain = chains[chainId];
        
        require(chain.registered == true, "Non-registered chain");
        require(depositWithdrawRequestExist(chainId, msg.sender) == false, "There is already existing withdrawal request being processed for this acc.");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chainId);
        
        // Withdraw whole deposit
        if (deposit == 0) {
            require(transactorExist(chainId, msg.sender) == true, "Non-existing transactor account (0 deposit balance)");
        }
        // Deposit in chain or withdraw just part of deposit
        else {
            require(chain.users[msg.sender].transactor.deposit != deposit, "Cannot deposit the same amount of tokens as you already has deposited.");
            require(checkLitionMinDeposit(deposit), "user does not meet Lition's min.required deposit condition");
            require(chain.chainValidator.checkDeposit(deposit, msg.sender, chain.actNumOfTransactors), "user does not meet chain validator's min.required deposit condition");
            require(deposit <= MAX_DEPOSIT, "deposit is greater than uint96_max_value");
        }
                
        _requestDepositInChain(chainId, deposit, msg.sender);
    }
    
    // Confirms deposit withdrawal request, token transfer is processed during confirmation
    function confirmDepositWithdrawalFromChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];

        require(chain.registered == true, "Non-registered chain");
        require(depositWithdrawRequestExist(chainId, msg.sender) == true, "Non-existing deposit withdrawal request");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chainId);
        
        // Chain is active or it's been inactive for less than 2*CHAIN_INACTIVITY_TIMEOUT
        if (chain.active == true || chain.lastNotary.timestamp + 2*CHAIN_INACTIVITY_TIMEOUT > now) {
            require(chain.lastNotary.block > chains[chainId].requests.accounts[msg.sender].depositWithdrawRequest.notaryBlock, "Confirm can be called in the next notary window after request was accepted.");
        }
        
        _confirmDepositWithdrawalFromChain(chainId, msg.sender);
    }
    
    // Internally creates/registers new chain.
    function registerChain(string calldata description, string calldata initEndpoint, ChainValidator validator, uint256 vesting, uint256 deposit, 
                           uint256 maxNumOfValidators, uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond) external returns (uint256 chainId) {
        require(bytes(description).length > 0 && bytes(description).length <= 200,           "Chain description cannot be empty");
        require(bytes(initEndpoint).length > 0 && bytes(initEndpoint).length <= 100,         "Chain endpoint cannot be empty");
        require(maxNumOfValidators > 0 && maxNumOfValidators <= MAX_NUM_OF_VALIDATORS_LIMIT, "Specified max num of validators out of limits");
        require(maxNumOfTransactors > 0,                                                     "Specified max num of transactors out of limits");
        require(involvedVestingNotaryCond == true || participationNotaryCond == true,        "At least on notary condition must be specified");
        require(deposit <= MAX_DEPOSIT,                                                      "deposit is greater than uint96_max_value");
        require(vesting <= MAX_VESTING,                                                      "vesting is greater than uint96_max_value");
        
        address creator         = msg.sender;
        
        // Inits chain data
        chainId                 = nextId;
        ChainInfo storage chain = chains[chainId];
        
        chain.chainValidator    = validator;
        
        // Validates vesting
        require(checkLitionMinVesting(vesting), "chain creator does not meet Lition's min.required vesting condition");
        require(chain.chainValidator.checkVesting(vesting, creator, chain.validators.list.length), "chain creator does not meet chain validator's min.required vesting condition");
        
        // Validates deposit
        require(checkLitionMinDeposit(deposit), "chain creator does not meet Lition's min.required deposit condition");
        require(chain.chainValidator.checkDeposit(deposit, creator, chain.actNumOfTransactors), "chain creator does not meet chain validator's min.required deposit condition");
        
        // Transfers vesting tokens
        token.transferFrom(creator, address(this), vesting);
        
        // Transfers deposit tokens
        token.transferFrom(creator, address(this), deposit);
        
        chain.description               = description;
        chain.endpoint                  = initEndpoint;
        chain.registered                = true;
        chain.maxNumOfValidators        = maxNumOfValidators;
        chain.maxNumOfTransactors       = maxNumOfTransactors;
        chain.involvedVestingNotaryCond = involvedVestingNotaryCond;
        chain.participationNotaryCond   = participationNotaryCond;
        
        emit NewChain(chainId, description, initEndpoint);
        
        emit VestInChain(chainId, creator, vesting, 0 /* zero block */, true);
        // Inits validator's data
        validatorCreate(chainId, creator, vesting);
        
        emit DepositInChain(chainId, creator, deposit, 0 /* zero block */, true);
        // Inits transactor's data
        transactorCreate(chainId, creator, deposit);
        
        nextId++;
    }
    
    // Reurns true, if user has vested enough tokensto become validator, othervise false
    function hasVested(uint256 chainId, address acc) view external returns (bool) {
        // No need to check vesting balance as it cannot be lover than min. required
        return validatorExist(chainId, acc);
    }
    
    // Returns true if user's remaining deposit balance >= min. required deposit and is allowed to transact
    function isAllowedToTransact(uint256 chainId, address acc) view external returns (bool) {
        // No need to check deposit balance as whitelisted flag should be alwyas set accordingly
        return chains[chainId].users[acc].transactor.whitelisted;
    }
    
    // Returns info about when was the last notary processed
    function getLastNotary(uint256 chainId) external view returns (uint256 notaryBlock, uint256 notaryTimestamp) {
        notaryBlock = chains[chainId].lastNotary.block;
        notaryTimestamp = chains[chainId].lastNotary.timestamp;
    }
    
    // TODO: remove on production version
    function testNotary(uint256 chainId, uint256 notaryBlock) external {
        emit Notary(chainId, notaryBlock, true);
        
        ChainInfo storage chain = chains[chainId];
        require(chain.registered, "Non-registered chain");
        
        // Remove validators who signed no block during this notary window and have mining flag == true
        removeInactiveValidators(chainId);
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryBlock;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
        
        emit Notary(chainId, notaryBlock, true);
    }
    
    // Returns static chain details
    function getStaticChainDetails(uint256 chainId) external view returns (string memory description, string memory endpoint, bool registered, uint256 maxNumOfValidators, uint256 maxNumOfTransactors,
                                                                           bool involvedVestingNotaryCond, bool participationNotaryCond) {
        ChainInfo storage chain = chains[chainId];
        
        description                 = chain.description;
        endpoint                    = chain.endpoint;
        registered                  = chain.registered;
        maxNumOfValidators          = chain.maxNumOfValidators;
        maxNumOfTransactors         = chain.maxNumOfTransactors;
        involvedVestingNotaryCond   = chain.involvedVestingNotaryCond;
        participationNotaryCond     = chain.participationNotaryCond;
    }
    
    // Returns dynamic chain details 
    function getDynamicChainDetails(uint256 chainId) public view returns (bool active, uint256 totalVesting, uint256 validatorsCount,
                                                                          uint256 lastValidatorVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp) {
        ChainInfo storage chain = chains[chainId];
        
        active               = chain.active;
        totalVesting         = chain.totalVesting;
        validatorsCount      = chain.validators.list.length;
        lastValidatorVesting = chain.users[chain.lastValidator].validator.vesting;   
        lastNotaryBlock      = chain.lastNotary.block;
        lastNotaryTimestamp  = chain.lastNotary.timestamp;
    }
    
    // Returns user details
    function getUserDetails(uint256 chainId, address acc) external view returns (uint256 deposit, bool whitelisted, uint256 vesting, 
                                                                                 bool mining, bool prevNotaryMined, bool secondPrevNotaryMined, bool thirdPrevNotaryMined) {
        ChainInfo storage chain = chains[chainId];
         
        deposit                 = chain.users[acc].transactor.deposit;
        whitelisted             = chain.users[acc].transactor.whitelisted;
        vesting                 = chain.users[acc].validator.vesting;
        mining                  = activeValidatorExist(chainId, acc);
        prevNotaryMined         = chain.users[acc].validator.currentNotaryMined;  
        secondPrevNotaryMined   = chain.users[acc].validator.prevNotaryMined;  
        thirdPrevNotaryMined    = chain.users[acc].validator.secondPrevNotaryMined;  
    }
    
    // Returns user requests details
    function getUserRequests(uint256 chainId, address acc) external view returns (bool vestingIncreaceReqExists, uint256 vestingReqNotary, uint256 vestingReqValue, 
                                                                                  bool depositFullWithdrawalReqExists, uint256 depositReqNotary) {
        if (vestingIncreaseRequestExist(chainId, acc)) {
            VestingIncreaseRequest storage request = chains[chainId].requests.accounts[acc].vestingIncreaseRequest;
            
            vestingIncreaceReqExists    = true;
            vestingReqNotary            = request.notaryBlock;
            vestingReqValue             = request.newVesting;
        }
        
        if (depositWithdrawRequestExist(chainId, acc)) {
            DepositWithdrawRequest storage request = chains[chainId].requests.accounts[acc].depositWithdrawRequest;
            
            depositFullWithdrawalReqExists  = true;
            depositReqNotary                = request.notaryBlock;
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
        
        emit Notary(chainId, notaryEndBlock, false);
                  
        ChainInfo storage chain = chains[chainId];
        require(chain.registered    == true,    "Invalid chain data: Non-registered chain");
        require(chain.totalVesting  > 0,        "Invalid chain data: current chain total_vesting == 0");
        
        require(miners.length       > 0,                            "Invalid statistics data: miners.length == 0");
        // There is upper limit of max num of validators == chain.maxNumOfValidators, but
        // loop can run up to chain.maxNumOfValidators*2 times as there is an edge case where during one window: validator mines a few blocks, then withdraws whole vesting, 
        // he is removed from validators list and new validator joins sidechain as there is empty space now. As a result there can be theoretically up to chain.maxNumOfValidators*2 validators 
        // in the statistics
        require(miners.length       < chain.maxNumOfValidators*2,   "Invalid statistics data: miners.length >= maxNumOfValidators*2");
        require(miners.length       == blocksMined.length,          "Invalid statistics data: miners.length != num of block mined");
        
        require(users.length        > 0,                            "Invalid statistics data: users.length == 0");
        require(users.length        < MAX_NUM_OF_PROCESSED_USERS,   "Invalid statistics data: users.length >= MAX_NUM_OF_PROCESSED_USERS");
        require(users.length        == userGas.length,              "Invalid statistics data: users.length != usersGas.length");
        
        require(v.length            < chain.maxNumOfValidators*2,      "Invalid statistics data: v.length >= maxNumOfValidators*2");
        require(v.length            == r.length,                    "Invalid statistics data: v.length != r.length");
        require(v.length            == s.length,                    "Invalid statistics data: v.length != s.length");
        
        require(notaryStartBlock    >  chain.lastNotary.block,      "Invalid statistics data: notaryBlock_start <= last known notary block");
        require(notaryEndBlock      >  notaryStartBlock,            "Invalid statistics data: notaryEndBlock <= notaryStartBlock");
        require(largestTx           >  0,                           "Invalid statistics data: Largest tx <= 0");
        
        bytes32 signatureHash = keccak256(abi.encodePacked(notaryEndBlock, miners, blocksMined, users, userGas, largestTx));
        
        // Validates notary conditions(involvedVesting && participation) to statistics to be accepted
        validateNotaryConditions(chainId, signatureHash, v, r, s);
        
        // Calculates total cost based on user's usage durint current notary window
        uint256 totalCost = processUsersConsumptions(chainId, users, userGas, largestTx);
        
        // In case totalCost == 0, something is wrong and there is no need for notary to continue as there is no tokens to be distributed to the miners.
        // There is probably ongoing coordinated attack based on invalid statistics sent to the notary
        require(totalCost > 0, "Invalid statistics data: users totalUsageCost == 0");
        
        // Calculates and process validator's rewards based on their participation rate and vesting balance
        processMinersRewards(chainId, notaryStartBlock, notaryEndBlock, miners, blocksMined, totalCost);
        
        // Remove validators who signed no block during this notary window and have mining flag == true
        removeInactiveValidators(chainId);
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryEndBlock;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
        
        emit Notary(chainId, notaryEndBlock, true);
    }
    
    // Validates notary conditions(involvedVesting && participation) to statistics to be accepted
    function validateNotaryConditions(uint256 chainId, bytes32 signatureHash, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) internal view {
        ChainInfo storage chain = chains[chainId];
        
        uint256 involvedVestingSum = 0;
        uint256 involvedSignaturesCount = 0;
        
        address signerAcc;
        for(uint256 i = 0; i < v.length; i++) {
            signerAcc = ecrecover(signatureHash, v[i], r[i], s[i]);
            
            // In case statistics is signed by validator, who is not registered in SC, ignore him   
            if (activeValidatorExist(chainId, signerAcc) == true) {
                involvedVestingSum += chain.users[signerAcc].validator.vesting;
                involvedSignaturesCount++;
            }
        }
        
        // There must be more than 50% out of total possible vesting involved in signatures
        if (chain.involvedVestingNotaryCond == true) {
            // There must be more than 50% out of total possible vesting involved
            involvedVestingSum = 2*involvedVestingSum;
            require(involvedVestingSum > chain.totalVesting, "Invalid statistics data: involvedVesting <= 50% of chain.totalVesting");
        }
        
        
        // There must be more than 2/3 + 1 out of all active validators unique signatures
        if (chain.participationNotaryCond == true) {
            uint256 actNumOfValidators = chain.validators.list.length;
            
            // min. number of active validators for BFT to work properly is 4
            if (actNumOfValidators >= 4) {
                uint256 minRequiredSignaturesCount = ((2 * actNumOfValidators) / 3) + 1;
                
                require(involvedSignaturesCount >= minRequiredSignaturesCount, "Invalid statistics data: Not enough signatures provided (2/3 + 1 cond)");
            }
            // if there is less than 4 active validators, everyone has to sign statistics
            else {
                require(involvedSignaturesCount == actNumOfValidators, "Invalid statistics data: Not enough signatures provided (involvedSignatures == activeValidatorsCount)");
            }
        }
    }
    
    // Returns list of user's addresses that are allowed to transact - their deposit >= min. required deposit
    function getAllowedToTransact(uint256 chainId, uint256 batch) external view returns (address[100] memory users, uint256 count, bool end) {
        ChainInfo storage chain = chains[chainId];
        
        count = 0;
        uint256 i = batch * 100;
        uint256 usersTotalCount = chain.transactors.list.length;
        address acc;
        while(i < (batch + 1)*100 && i < usersTotalCount) {
            acc = chain.transactors.list[i];
            
            // Get transactors
            if (chain.users[acc].transactor.whitelisted == true) {
                users[count] = acc;
                count++;
            } 
            i++;
        }
        
        if (i == usersTotalCount) {
            end = true;
        }
        else {
            end = false;
        }
    }
    
    // Returns list of addresses of active validators
    function getActiveValidators(uint256 chainId) view external returns (address[MAX_NUM_OF_VALIDATORS_LIMIT] memory validators, uint256 count) {
        ChainInfo storage chain = chains[chainId];
        
        count = 0;
        for (uint256 i = 0; i < chain.validators.list.length; i++) {
            validators[count] = chain.validators.list[i];
            count++;
        }
    }
    
    // Sets mining validator's mining flag to true and emit event so other nodes vote him
    function startMining(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        address acc = msg.sender;
        uint256 validatorVesting = chain.users[acc].validator.vesting;
        
        require(chain.registered == true, "Non-registered chain");
        require(validatorExist(chainId, acc) == true, "Non-existing validator (0 vesting balance)");
        require(vestingIncreaseRequestExist(chainId, acc) == false, "Cannot start mining - there is ongoing vesting request.");
        require(chain.chainValidator.checkVesting(validatorVesting, acc, chain.validators.list.length) == true, "User does not meet chain validator's min.required vesting condition");
        
        if (activeValidatorExist(chainId, acc) == true) {
            // Emit event even if validator is already active - user might want to explicitely emit this event in case something went wrong on the nodes and
            // others did not vote him
            emit AccountMining(chainId, acc, true);
            
            return;
        }
            
        // Upper limit of validators reached
        if (chain.validators.list.length >= chain.maxNumOfValidators) {
            require(validatorVesting > chain.users[chain.lastValidator].validator.vesting, "Upper limit of validators reached. Must vest more than the last validator to replace him.");
            validatorReplace(chainId, acc);
        }
        // There is still empty place for new validator
        else {
            validatorInsert(chainId, acc);
        }
    }
  
    // Sets mining validator's mining flag to false and emit event so other nodes unvote
    function stopMining(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        address acc = msg.sender;
        
        require(chain.registered == true, "Non-registered chain");
        require(validatorExist(chainId, acc) == true, "Non-existing validator (0 vesting balance)");
    
        if (activeValidatorExist(chainId, acc) == false) {
            // Emit event even if validator is already inactive - user might want to explicitely emit this event in case something went wrong on the nodes and
            // others did not unvote him
            emit AccountMining(chainId, acc, false);
            
            return;
        }
        
        validatorRemove(chainId, acc);
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
        uint256 index = map.listIndex[acc];
        require(index > 0 && index <= map.list.length, "RemoveAcc invalid index");
        
        // Move an last element of array into the vacated key slot.
        uint256 foundIndex = index - 1;
        uint256 lastIndex  = map.list.length - 1;
    
        map.listIndex[map.list[lastIndex]] = foundIndex + 1;
        map.list[foundIndex] = map.list[lastIndex];
        map.list.length--;
    
        // Deletes element
        map.listIndex[acc] = 0;
    }
    
    // Returns true, if acc exists in the iterable map, otherwise false
    function existAcc(IterableMap storage map, address acc) internal view returns (bool) {
        return map.listIndex[acc] != 0;
    }
    
    // Inits validator data holder in the users mapping 
    function validatorCreate(uint256 chainId, address acc, uint256 vesting) internal {
        Validator storage validator = chains[chainId].users[acc].validator;
        validator.vesting               = uint96(vesting);
        // Inits previously notary windows as mined so validator does not get removed from the list of actively mining validators right after the creation
        validator.currentNotaryMined    = true;
        validator.prevNotaryMined       = true;
        validator.secondPrevNotaryMined = true;
    }
    
    // Deinits validator data holder in the users mapping 
    function validatorDelete(uint256 chainId, address acc) internal {
        Validator storage validator = chains[chainId].users[acc].validator;
        
        if (activeValidatorExist(chainId, acc) == true) {
            validatorRemove(chainId, acc);
        }
        
        if (validator.vesting               != 0)       validator.vesting = 0;
        if (validator.currentNotaryMined    == true)    validator.currentNotaryMined = false;
        if (validator.prevNotaryMined       == true)    validator.prevNotaryMined = false;
        if (validator.secondPrevNotaryMined == true)    validator.secondPrevNotaryMined = false;
    }
    
    // Inserts validator into the list of actively mining validators
    function validatorInsert(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        Validator storage validator = chain.users[acc].validator;
        
        insertAcc(chain.validators, acc);   
        
        // Updates chain total vesting
        chain.totalVesting = chain.totalVesting.add(validator.vesting);
        
        // Updates lastValidator in case new validator's vesting balance is less
        if (validator.vesting < chain.users[chain.lastValidator].validator.vesting) {
            chain.lastValidator = acc;
        }
        
        emit AccountMining(chainId, acc, true);
    }
    
    // Removes validator from the list of actively mining validators
    function validatorRemove(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        Validator storage validator = chain.users[acc].validator;
        
        removeAcc(chain.validators, acc);   
        
        // Updates chain total vesting
        chain.totalVesting = chain.totalVesting.sub(validator.vesting);
        
        // If there is no active validator left, set chain.active flag to false so others might vest in chain without
        // waiting for the next notary window to be processed
        if (chain.validators.list.length == 0) {
            chain.active = false;
        }
        // There are still some active validators left
        else {
            // If lastValidator is being removed, find a new validator with the smallest vesting balance
            if (chain.lastValidator == acc) {
                resetLastValidator(chainId);
            }
        }
        
        emit AccountMining(chainId, acc, false);
    }
    
    // Replaces lastValidator for the new one in the list of actively mining validators
    function validatorReplace(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        address accToBeReplaced = chain.lastValidator;
        Validator memory validatorToBeReplaced = chain.users[accToBeReplaced].validator;
        Validator memory newValidator = chain.users[acc].validator;
        
        // Updates chain total vesting
        chain.totalVesting = chain.totalVesting.sub(validatorToBeReplaced.vesting);
        chain.totalVesting = chain.totalVesting.add(newValidator.vesting);
        
        // Updates active validarors list
        removeAcc(chain.validators, accToBeReplaced);
        insertAcc(chain.validators, acc);
        
        // Finds a new validator with the smallest vesting balance
        resetLastValidator(chainId);
        
        emit AccountMining(chainId, accToBeReplaced, false);
        emit AccountMining(chainId, acc, true);
    }
    
    // Resets last validator - the one with the smallest vesting balance
    function resetLastValidator(uint256 chainId) internal {
        ChainInfo storage chain = chains[chainId];
        
        address foundLastValidatorAcc = chain.validators.list[0];
        Validator memory foundLastValidator  = chain.users[foundLastValidatorAcc].validator;
        
        address actValidatorAcc;
        Validator memory actValidator;
        for (uint256 i = 1; i < chain.validators.list.length; i++) {
            actValidatorAcc = chain.validators.list[i];
            actValidator    = chain.users[actValidatorAcc].validator;
            
            if (actValidator.vesting < foundLastValidator.vesting) {
                foundLastValidatorAcc = actValidatorAcc;
                foundLastValidator    = actValidator;
            }
        }
        
        chain.lastValidator = foundLastValidatorAcc;
    }
    
    // Returns true, if acc is in the list of actively mining validators, otherwise false
    function activeValidatorExist(uint256 chainId, address acc) internal view returns (bool) {
        return existAcc(chains[chainId].validators, acc);
    }
    
    // Returns true, if acc hase vesting > 0, otherwise false
    function validatorExist(uint256 chainId, address acc) internal view returns (bool) {
        return chains[chainId].users[acc].validator.vesting > 0;
    }
    
    // Inits transactor data holder in the users mapping and inserts it into the list of transactors
    function transactorCreate(uint256 chainId, address acc, uint256 deposit) internal {
        Transactor storage transactor = chains[chainId].users[acc].transactor;
        
        transactor.deposit            = uint96(deposit);
        transactorWhitelist(chainId, acc);
        
        insertAcc(chains[chainId].transactors, acc);
    }
    
    // Deinits transactor data holder in the users mapping and removes it from the list of transactors
    function transactorDelete(uint256 chainId, address acc) internal {
        Transactor storage transactor = chains[chainId].users[acc].transactor;
        
        if (transactor.deposit != 0) transactor.deposit = 0;
        transactorBlacklist(chainId, acc);
        
        removeAcc(chains[chainId].transactors, acc);
    }
    
    // Returns true, if acc is in the list of transactors
    function transactorExist(uint256 chainId, address acc) internal view returns (bool) {
        return existAcc(chains[chainId].transactors, acc);
    }
    
    // Blacklists transactor
    function transactorBlacklist(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        Transactor storage transactor = chain.users[acc].transactor;
        
        if (transactor.whitelisted == true) {
            chain.actNumOfTransactors--;
            
            transactor.whitelisted = false;
            emit AccountWhitelist(chainId, acc, false);
        }
    }
    
    // Whitelists transactor
    function transactorWhitelist(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        Transactor storage transactor = chain.users[acc].transactor;
        
        if (transactor.whitelisted == false) {
            chain.actNumOfTransactors++;
            
            transactor.whitelisted = true;
            emit AccountWhitelist(chainId, acc, true);
        }
    }
    
    /**************************************************************************************************************************/
    /*********************************** Functions related to the vesting/deposit requests ************************************/
    /**************************************************************************************************************************/
    
    // Creates new vesting request
    function vestingIncreaseRequestCreate(uint256 chainId, address acc, uint256 vesting) internal {
        ChainInfo storage chain          = chains[chainId];
        RequestsEntry storage entry      = chain.requests.accounts[acc];
        
        entry.vestingIncreaseRequest.newVesting  = uint96(vesting);
        entry.vestingIncreaseRequest.notaryBlock = chain.lastNotary.block; 
        
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
    // This method should never be called directly, VestingIncreaseRequestDelete & depositWithdrawRequestDelete should be called instead
    function requestsPairDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        address[] storage requestsList = chain.requests.list;
        
        // request_exists(chainId, acc), vestingIncreaseRequestExist(chainId, acc) and deposoti_withdraw_exists(chainId, acc) could be used instead
        uint256 index = chain.requests.accounts[acc].index;
        require(index > 0 && index <= requestsList.length, "RequestsPair delete: invalid index");
    
        // Move an last element of array into the vacated key slot.
        uint256 foundIndex = index - 1;
        uint256 lastIndex = requestsList.length - 1;
    
        chain.requests.accounts[requestsList[lastIndex]].index = foundIndex + 1;
        requestsList[foundIndex] = requestsList[lastIndex];
        requestsList.length--;
    
        delete chain.requests.accounts[acc];
    }
    
    function vestingIncreaseRequestDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no ongoing deposit request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].depositWithdrawRequest.notaryBlock == 0) {
            requestsPairDelete(chainId, acc);
            return;
        } 
        
        // There is ongoing deposit request for this account - only reset vesting request
        VestingIncreaseRequest storage request = chain.requests.accounts[acc].vestingIncreaseRequest;
        request.notaryBlock    = 0;
        request.newVesting     = 0;
    }
    
    function depositWithdrawRequestDelete(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        // There is no ongoing vesting request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].vestingIncreaseRequest.notaryBlock == 0) {
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
    function vestingIncreaseRequestExist(uint256 chainId, address acc) internal view returns (bool) {
        return chains[chainId].requests.accounts[acc].vestingIncreaseRequest.notaryBlock != 0;
    }
    
    // Checks if acc has any ongoing DEPOSIT WITHDRAWAL request
    function depositWithdrawRequestExist(uint256 chainId, address acc) internal view returns (bool) {
        return chains[chainId].requests.accounts[acc].depositWithdrawRequest.notaryBlock != 0;
    }
    
    function _requestVestInChain(uint256 chainId, uint256 vesting, address acc) internal {
        ChainInfo storage chain     = chains[chainId];
        Validator storage validator = chain.users[acc].validator;
        
        uint256 validatorVesting = validator.vesting;
        
        // if newVesting > oldVesting, send tokens first and update internal vesting balance in cofirm method
        if (vesting > validatorVesting) {
            uint256 toVest = vesting - validatorVesting;
            token.transferFrom(acc, address(this), toVest);
            
            // Internally creates new validator
            if (validatorExist(chainId, acc) == false) {
                validatorCreate(chainId, acc, vesting);
            }
            
            // Chain is not active - process vesting immediately so user can start mining and chain can become active
            if (chain.active == false) {
                validator.vesting = uint96(vesting);
                
                emit VestInChain(chainId, acc, vesting, chain.lastNotary.block, true);
            }
            // Chain is active - process vesting in 2 steps
            else {
                vestingIncreaseRequestCreate(chainId, acc, vesting);
                emit VestInChain(chainId, acc, vesting, chain.requests.accounts[acc].vestingIncreaseRequest.notaryBlock, false);
            }
            
            return;
        }
      
        // Full of partial vesting withdrawals are processed immediately
        // Full withdrawal -> delete validator
        uint256 toWithdraw;
        if (vesting == 0) {
            toWithdraw = validatorVesting;
            validatorDelete(chainId, acc);
        }
        // Partial withdrawal
        else {
            toWithdraw = validatorVesting - vesting;
            
            // If validator is actively mining, decrease chain's total vesting
            if (activeValidatorExist(chainId, acc) == true) {
                chain.totalVesting = chain.totalVesting.sub(toWithdraw);
            }
         
            validator.vesting = uint96(vesting);    
        } 
        
        // Transfers tokens
        token.transfer(acc, toWithdraw);
        
        emit VestInChain(chainId, acc, vesting, chain.lastNotary.block, true);
    }
    
    // Basically just transfers the tokens, validator's vesting balance update is always done at the of notary atomatically
    function _confirmVestIncreaseInChain(uint256 chainId, address acc) internal {
        ChainInfo storage chain        = chains[chainId];
        VestingIncreaseRequest storage request = chain.requests.accounts[acc].vestingIncreaseRequest;
        Validator storage validator    = chain.users[acc].validator;
        
        uint256 newVesting = request.newVesting;
        uint256 requestNotaryBlock = request.notaryBlock;
        vestingIncreaseRequestDelete(chainId, acc);
        
        uint256 origVesting = validator.vesting;
        validator.vesting = uint96(newVesting);
        
        if (activeValidatorExist(chainId, acc) == true) {
            chain.totalVesting = chain.totalVesting.add(newVesting - origVesting);
        }
        
        emit VestInChain(chainId, acc, newVesting, requestNotaryBlock, true);
    }
    
    function _requestDepositInChain(uint256 chainId, uint256 deposit, address acc) internal {
        ChainInfo storage chain       = chains[chainId];
        Transactor storage transactor = chain.users[acc].transactor;
        
        // If user wants to withdraw whole deposit
        if (deposit == 0) {
            // Chain is not active and last notary is older than 2*CHAIN_INACTIVITY_TIMEOUT - enable full deposit withdrawal immmediately
            if (chain.active == false && chain.lastNotary.timestamp + 2*CHAIN_INACTIVITY_TIMEOUT < now) {
                uint256 toWithdraw = transactor.deposit;
                transactorDelete(chainId, acc);
                
                // Withdraw whole deposit
                token.transfer(acc, toWithdraw);
                
                emit DepositInChain(chainId, acc, deposit, chain.lastNotary.block, true);
            }
            // Chain is active - create withdrawal request and process full deposit withdrawal in 2 steps
            else {
                depositWithdrawRequestCreate(chainId, acc);
                
                emit DepositInChain(chainId, acc, deposit, chain.requests.accounts[acc].depositWithdrawRequest.notaryBlock, false);  
                transactorBlacklist(chainId, acc);
            }  
          
            return;
        }
      
        // If user wants to deposit in chain, process it immediately
        uint256 actTransactorDeposit = transactor.deposit;
        
        if(actTransactorDeposit > deposit) {
            transactor.deposit = uint96(deposit);
         
            uint256 toWithdraw = actTransactorDeposit - deposit;
            token.transfer(acc, toWithdraw);
        } else {
            uint256 toDeposit = deposit - actTransactorDeposit;
            token.transferFrom(acc, address(this), toDeposit);
         
            // First deposit - create internally new user
            if (transactorExist(chainId, acc) == false) {
                transactorCreate(chainId, acc, deposit);
            }
            else {
                transactor.deposit = uint96(deposit);
                transactorWhitelist(chainId, acc);
            }
        }
        
        emit DepositInChain(chainId, acc, deposit, chain.lastNotary.block, true);
    }
    
    function _confirmDepositWithdrawalFromChain(uint256 chainId, address acc) internal {
        ChainInfo storage chain = chains[chainId];
        
        uint256 toWithdraw = chain.users[acc].transactor.deposit;
        uint256 requestNotaryBlock = chain.requests.accounts[acc].depositWithdrawRequest.notaryBlock;
        
        transactorDelete(chainId, acc);
        depositWithdrawRequestDelete(chainId, acc);
        
        // Withdraw whole deposit
        token.transfer(acc, toWithdraw);
        
        emit DepositInChain(chainId, acc, 0, requestNotaryBlock, true);
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
        Transactor storage transactor = chain.users[acc].transactor;
        transactorDeposit = transactor.deposit;
        
        // This can happen only if there is non-registered transactor(user) in statistics, which means that there is probaly
        // ongoing coordinated attack based on invalid statistics sent to the notary
        if (transactorExist(chainId, acc) == false || userGas[i] == 0) {
            // Ignores non-registred user and let nodes know he is not allowed to transact eventhough they should not let such user to transact
            emit AccountWhitelist(chainId, users[i], false);
            continue;
        }
        
        userCost = (userGas[i] * largestReward) / largestTx;
        
        // This can happen only if user runs out of tokens(which should never happen due to min.required deposit)
        if(userCost > transactorDeposit ) {
            userCost = transactorDeposit;
        
            transactorDelete(chainId, acc);
        }
        else {
            transactorDeposit -= userCost;
            
            // Updates user's stored deposit balance based on his usage
            transactor.deposit = uint96(transactorDeposit);
            
            // Check if user's deposit balance is >= min. required deposit conditions
            if (checkLitionMinDeposit(transactorDeposit) == false || chain.chainValidator.checkDeposit(transactorDeposit, acc, chain.actNumOfTransactors) == false) {
                // If not, do not allow him to transact anymore
                transactorBlacklist(chainId, acc);
            }
        }
        
        // Adds user's cost to the total cost
        // No need for safe math as we internally allow max MAX_NUM_OF_PROCESSED_USERS users to be proceesed
        // max possible userCost is 10^32 * 10^17, so max possible totalCost is MAX_NUM_OF_PROCESSED_USERS(250) * 10^32 * 10^17 = 10^49, which will never overfloww uint256
        totalCost += userCost;  
     }
   }

   // Process miners rewards based on their participation rate(how many blocks they signed) and their vesting balance
   function processMinersRewards(uint256 chainId, uint256 startNotaryBlock, uint256 endNotaryBlock, address[] memory miners, uint32[] memory blocksMined, uint256 litToDistribute) internal {
     ChainInfo storage chain = chains[chainId];
     
     // Min. vesting balance to be a trust node. Trust Nodes haved doubled(virtually) vesting
     uint256 minTrustNodeVesting = 50000*LIT_PRECISION; 
     
     // How many block could validator mined since the last notary in case he did sign every possible block 
     uint256 maxBlocksMined = endNotaryBlock - startNotaryBlock;

     // Total involved vesting 
     uint256 totalInvolvedVesting = 0;
     
     // Selected validator's vesting balance
     uint256 validatorVesting;
     
     address actMinerAcc;
     // Runs through all miners and calculates total involved vesting.
     // in the statistics
     for(uint256 i = 0; i < miners.length; i++) {
        actMinerAcc = miners[i];
        
        // This can happen only if there is non-registered validator in statistics, which means for 99% that the validator withdrawed whole vesting during current notary window
        // or there is ongoing coordinated attack based on invalid statistics sent to the notary
        if (activeValidatorExist(chainId, actMinerAcc) == false || blocksMined[i] == 0) {
            continue;
        }
        
        // Updates validator's mining flags statistics
        Validator storage validator     = chain.users[actMinerAcc].validator;
        validator.secondPrevNotaryMined = validator.prevNotaryMined;
        validator.prevNotaryMined       = validator.currentNotaryMined;
        validator.currentNotaryMined    = true;
        
        validatorVesting = chain.users[actMinerAcc].validator.vesting;
        
        // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
        if (validatorVesting >= minTrustNodeVesting) {
            // Validator's stored vesting is max uint96
            validatorVesting *= 2;
        }

        // No need for safe math
        // max possible (maxBlocksMined / blocksMined[i]) valuse is 10^32, max possible validatorVesting value is 10^96, when virtually doubled it is 10^192, 
        // so max possible totalInvolvedVesting value is 2*chain.maxNumOfValidators(24) * 10^32 * 10^192 = 42*10^224, which cannot overfloww uint256
        totalInvolvedVesting += (maxBlocksMined * validatorVesting) / blocksMined[i];
     }
     
     // In case totalInvolvedVesting == 0, something is wrong and there is no need for notary to continue as rewards cannot be calculated. It might happen
     // as edge case when the last validator stopped mining durint current notary window or there is ongoing coordinated attack based on invalid statistics sent to the notary
     require(totalInvolvedVesting > 0, "totalInvolvedVesting == 0. Invalid statistics or 0 active miners left in the chain");

     
     uint256 minerReward;
     // Runs through all miners and calculates their reward based on:
     //     1. How many blocks out of max_blocks_mined each miner signed
     //     2. How many token each miner vested
     for(uint256 i = 0; i < miners.length; i++) {
        actMinerAcc = miners[i];
        
        // This can happen only if there is non-registered validator in statistics, which means for 99% that the validator withdrawed whole vesting during current notary window
        // or there is ongoing coordinated attack based on invalid statistics sent to the notary
        if (activeValidatorExist(chainId, actMinerAcc) == false || blocksMined[i] == 0) {
            continue;
        }
        
        validatorVesting = chain.users[actMinerAcc].validator.vesting;
        
        // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
        if (validatorVesting >= minTrustNodeVesting) {
            // Validator's stored vesting is max uint96
            validatorVesting *= 2;
        }
        
        // No need for safe math as max value of (maxBlocksMined / blocksMined[i]) is 10^32, max value of (validatorVesting / totalInvolvedVesting) is 1 and 
        // max value of litToDistribute(calculated in processUsersConsumptions) is 10^97, so max possible miner reward is 10^32 * 1 * 10^97 = 10^129
        minerReward = (validatorVesting * maxBlocksMined * litToDistribute) / blocksMined[i] / totalInvolvedVesting;
        token.transfer(miners[i], minerReward);
        
        // No need for safe math as miner reward is calculated as fraction of total litToDistribute and sum of all miners rewards must always be <= litToDistribute
        litToDistribute -= minerReward;
     }

     if(litToDistribute > 0) {
        // Sends the rest(math rounding) to the miner, who called notary function
        token.transfer(msg.sender, litToDistribute);
     }
   }
   
    // Removes validators that did not mine at all during the last 3 notary windows
    function removeInactiveValidators(uint256 chainId) internal {
        ChainInfo storage chain = chains[chainId];
        
        address validatorAcc;
        for (uint256 i = 0; i < chain.validators.list.length; i++) {
            validatorAcc = chain.validators.list[i];
            Validator memory validator = chain.users[validatorAcc].validator;
           
            if (validator.currentNotaryMined || validator.prevNotaryMined || validator.secondPrevNotaryMined) {
                continue;
            }
           
            validatorRemove(chainId, validatorAcc);
        } 
    }
   
    // Checks if chain is active(successfull notary processed during last CHAIN_INACTIVITY_TIMEOUT), if not set it active flag to false
    // If last notary is older than CHAIN_INACTIVITY_TIMEOUT, it means that validators cannot reach consensus or there is no active validator and chain is basically stuck.
    function checkAndSetChainActivity(uint256 chainId) internal {
        ChainInfo storage chain = chains[chainId];
        
        if (chain.active == true && chain.lastNotary.timestamp + CHAIN_INACTIVITY_TIMEOUT < now) {
            chain.active = false;   
        }
    }
}

// SafeMath library. Source: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/math/SafeMath.sol
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