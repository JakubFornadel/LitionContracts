pragma solidity >= 0.5.11;

interface ChainValidator {
   function validateNewValidator(uint256 vesting, address acc, bool mining, uint256 actNumOfValidators) external returns (bool);
   function validateNewTransactor(uint256 deposit, address acc, uint256 actNumOfTransactors) external returns (bool);
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
    
    // Max num of characters in chain url
    uint256 constant MAX_URL_LENGTH              = 100;
    
    // Max num of characters in chain description
    uint256 constant MAX_DESCRIPTION_LENGTH      = 200;
    
    // Time after which chain becomes inactive in case there was no successfull notary processed
    // Users can then increase their vesting instantly and bypass 2-step process. For instant full deposit withdrawal users must wait 2*CHAIN_INACTIVITY_TIMEOUT
    uint256 constant CHAIN_INACTIVITY_TIMEOUT    = 14 days;
    
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function checkLitionMinVesting(uint256 vesting) private pure returns (bool) {
        return vesting >= 1000*LIT_PRECISION;
    }
    
    // This is lition additional required check for the one from ChainValidator, in which sidechain creator specifies conditions himself
    function checkLitionMinDeposit(uint256 deposit) private pure returns (bool) {
        return deposit >= 1000*LIT_PRECISION;
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
    
    struct VestingRequest {
        // Flag if there is ongoing request for user
        bool                            exist;
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
        // New value of vesting to be set
        uint96                          newVesting;
    }
    
    // Only full deposit withdrawals are saved as deposit requests - other types of deposits do not need to be confirmed
    struct DepositWithdrawalRequest {
        // Flag if there is ongoing request for user
        bool                            exist;
        // Last notary block number when the request was accepted 
        uint256                         notaryBlock;
    }
    
    struct RequestsEntry {
        // index to the requestsList, indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing request
        uint256                         index;
        // Deposit withdrawal request details
        DepositWithdrawalRequest        depositWithdrawalRequest;
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
        // Internal chain ID
        uint256                         id;
        
        // Address of the chain creator
        address                         creator;
        
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
        mapping(address => User)        usersData;
        
        // List of existing users - all users, who have deposited or vested
        IterableMap                     users;
        
        // List of existing validators - only those who vested enough and are actively mining are here
        // Active validator are in separate array because of heavy processing furing notary (no need to filter them out of users thanks tot this)
        IterableMap                     validators;
        
        // Validator with the smallest vesting balance among all existing validators
        // In case someone vests more tokens, he will replace the one with smallest vesting balance
        address                         lastValidator;
        
        // Actual number of whitelisted transactors (their current depost > min.required deposit)
        uint256                         actNumOfTransactors;
        
        // List of users requests for changing their vesting/deposit balance
        Requests                        requests;
        
        // Max number of active validators for chain. There is no limit to how many users can vest.
        // Tthis is limit for active validators (startMining function) 
        // 0  means unlimited
        // It must be some reasonable value together with min. required vesting condition as 
        // too small num of validators limit with too low min. required vesting condition might lead to chain being stuck
        uint256                         maxNumOfValidators;
        
        // Max number of users(transactors) for chain.
        // Tthis is limit for whitelisted transactors (their current depost > min.required deposit)
        // 0  means unlimited
        // It must be some reasonable value together with min. required deposit condition as 
        // too small num of validators limit with too low min. required deposit condition might lead to chain being stuck
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
    //      * 1 notary windows after the confirm method is called - in case new vesting > actual vesting or new vesting == 0
    //      * emmidiately - in case new vesting < actual vesting
    //
    // In case new vesting > actual vesting, user first creates request, tokens are transferred to the sc, then he must confirm this request in the next notary window 
    // and his internal vesting balance is updated 
    function requestVestInChain(uint256 chainId, uint256 vesting) external {
        ChainInfo storage chain = chains[chainId];
        
        // Enable users to vest into registered (but non-active) chain and start minig so it becomes active
        require(chain.registered == true,                                 "Non-registered chain");
        require(transactorExist(chain, msg.sender) == false,              "Validator cannot be transactor at the same time. Withdraw your depoist or use different account");
        require(vestingRequestExist(chain, msg.sender) == false,          "There is already one vesting request being processed for this acc");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chain);
            
        // Full vesting withdrawal
        if (vesting == 0) {
            require(validatorExist(chain, msg.sender) == true,            "Non-existing validator account (0 vesting balance)");
            require(activeValidatorExist(chain, msg.sender) == false,     "StopMinig must be called first");  
        }
        // Vest in chain or withdraw just part of vesting
        else {
            require(vesting <= MAX_VESTING,                                 "vesting is greater than uint96_max_value");
            require(chain.usersData[msg.sender].validator.vesting != vesting,   "Cannot vest the same amount of tokens as you already has vested");
            require(checkLitionMinVesting(vesting) == true,                 "user does not meet Lition's min.required vesting condition");
            
            if (chain.chainValidator != ChainValidator(0)) {
                require(chain.chainValidator.validateNewValidator(vesting, msg.sender, false /* not mining yet */, chain.validators.list.length), "Validator not allowed by external chainvalidator SC");
            }
        }
        
        requestVest(chain, vesting, msg.sender);
    }
    
    // Confirms vesting request
    function confirmVestInChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        
        // Enable users to confirm vesting request into registered (but non-active) chain and start minig so it becomes active
        require(chain.registered == true, "Non-registered chain");
        require(vestingRequestExist(chain, msg.sender) == true, "Non-existing vesting request");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chain);
        
        // Chain is active
        if (chain.active == true) {
            require(chain.lastNotary.block > chain.requests.accounts[msg.sender].vestingRequest.notaryBlock, "Confirm can be called in the next notary window after request was accepted");    
        }
        
        confirmVest(chain, msg.sender);
    }
    
    // Requests deposit in chain. It will be processed and applied to the actual user state after next:
    //      * 1 notary windows after the confirm method is called - in case deposit == 0
    //      * emmidiately                                         - in case deposit != 0
    // Only full deposit withdrawals needs to be handled in 2 steps as it would allow users to send unlimited amount of txs to the sidechain and 
    // withdraw whole deposit right before notary function, in which case he would pay nothing...
    function requestDepositInChain(uint256 chainId, uint256 deposit) external {
        ChainInfo storage chain = chains[chainId];
        
        require(chain.registered == true,                                               "Non-registered chain");
        require(validatorExist(chain, msg.sender) == false,                           "Transactor cannot be validator at the same time. Withdraw your vesting or use different account");
        require(depositWithdrawalRequestExist(chain, msg.sender) == false,            "There is already existing withdrawal request being processed for this acc");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chain);
        
        // Withdraw whole deposit
        if (deposit == 0) {
            require(transactorExist(chain, msg.sender) == true,                       "Non-existing transactor account (0 deposit balance)");
            
            if (chain.active == false) {
                require(chain.lastNotary.timestamp + 2*CHAIN_INACTIVITY_TIMEOUT < now,  "Chain is inactive, for instant full deposit withdrawal wait for 2*CHAIN_INACTIVITY_TIMEOUT since the last notary");
            }
        }
        // Deposit in chain or withdraw just part of deposit
        else {
            require(chain.usersData[msg.sender].transactor.deposit != deposit,          "Cannot deposit the same amount of tokens as you already has deposited");
            require(checkLitionMinDeposit(deposit),                                     "user does not meet Lition's min.required deposit condition");
            require(deposit <= MAX_DEPOSIT,                                             "deposit is greater than uint96_max_value");
            
            if (chain.chainValidator != ChainValidator(0)) {
                require(chain.chainValidator.validateNewTransactor(deposit, msg.sender, chain.actNumOfTransactors), "Transactor not allowed by external chainvalidator SC");
            }
            
            // Upper limit of transactors reached
            if (chain.maxNumOfTransactors != 0 && isAllowedToTransact(chainId, msg.sender) == false) {
                require(chain.actNumOfTransactors <= chain.maxNumOfTransactors, "Upper limit of transactors reached");
            }
        }
                
        requestDeposit(chain, deposit, msg.sender);
    }
    
    // Confirms deposit withdrawal request, token transfer is processed during confirmation
    function confirmDepositWithdrawalFromChain(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];

        require(chain.registered == true, "Non-registered chain");
        require(depositWithdrawalRequestExist(chain, msg.sender) == true, "Non-existing deposit withdrawal request");
        
        // Checks if chain is active, if not set it active flag to false 
        checkAndSetChainActivity(chain);
        
        // Chain is active or it's been inactive for less than 2*CHAIN_INACTIVITY_TIMEOUT
        if (chain.active == true) {
            require(chain.lastNotary.block > chain.requests.accounts[msg.sender].depositWithdrawalRequest.notaryBlock, "Confirm can be called in the next notary window after request was accepted");
        }
        // Chain is inactive
        else {
            require(chain.lastNotary.timestamp + 2*CHAIN_INACTIVITY_TIMEOUT < now, "Chain is inactive, for instant full deposit withdrawal confirm wait for 2*CHAIN_INACTIVITY_TIMEOUT since the last notary");
        }
        
        confirmDepositWithdrawal(chain, msg.sender);
    }
    
    // Internally creates/registers new chain.
    function registerChain(string calldata description, string calldata initEndpoint, ChainValidator chainValidator, uint256 vesting, uint256 maxNumOfValidators,
                           uint256 maxNumOfTransactors, bool involvedVestingNotaryCond, bool participationNotaryCond) external returns (uint256 chainId) {
        require(bytes(description).length > 0 && bytes(description).length <= MAX_DESCRIPTION_LENGTH,   "Chain description length must be: > 0 && <= MAX_DESCRIPTION_LENGTH");
        require(bytes(initEndpoint).length > 0 && bytes(initEndpoint).length <= MAX_URL_LENGTH,         "Chain endpoint length must be: > 0 && <= MAX_URL_LENGTH");
        require(involvedVestingNotaryCond == true || participationNotaryCond == true,                   "At least on notary condition must be specified");
        require(checkLitionMinVesting(vesting) == true,                                                 "Chain creator does not meet Lition's min.required vesting condition");
        require(vesting <= MAX_VESTING,                                                                 "Vesting is greater than uint96_max_value");
    
        chainId                         = nextId;
        ChainInfo storage chain         = chains[chainId];
        
        if (chainValidator != ChainValidator(0)) {
            require(chainValidator.validateNewValidator(vesting, msg.sender, false /* not mining yet */, chain.validators.list.length) == true, "Chain creator not allowed by external chainvalidator SC (vesting)");
            chain.chainValidator = chainValidator;
        }
        
        chain.id                        = chainId;
        chain.description               = description;
        chain.endpoint                  = initEndpoint;
        chain.registered                = true;
        chain.maxNumOfValidators        = maxNumOfValidators;
        chain.maxNumOfTransactors       = maxNumOfTransactors;
        chain.involvedVestingNotaryCond = involvedVestingNotaryCond;
        chain.participationNotaryCond   = participationNotaryCond;
        chain.creator                   = msg.sender;
        
        // Transfers vesting tokens
        token.transferFrom(msg.sender, address(this), vesting);
        
        emit NewChain(chainId, description, initEndpoint);
        emit VestInChain(chainId, msg.sender, vesting, 0 /* zero block */, true);
        
        // Creates validator
        validatorCreate(chain, msg.sender, vesting);
        
        nextId++;
    }
    
    // Changes chain description and endpoint
    function setChainStaticDetails(uint256 chainId, string calldata description, string calldata endpoint) external {
        ChainInfo storage chain = chains[chainId];
        require(msg.sender == chain.creator, "Only chain creator can call this method");
    
        require(bytes(description).length <= MAX_DESCRIPTION_LENGTH,   "Chain description length must be: > 0 && <= MAX_DESCRIPTION_LENGTH");
        require(bytes(endpoint).length <= MAX_URL_LENGTH,              "Chain endpoint length must be: > 0 && <= MAX_URL_LENGTH");
        
        if (bytes(description).length > 0) {
            chain.description = description;
        }
        if (bytes(endpoint).length > 0) {
            chain.endpoint = endpoint;
        }
    }
    
    // Returns true, if user has vested enough tokens to become validator, othervise false
    function isAllowedToValidate(uint256 chainId, address acc) view public returns (bool) {
        // No need to check vesting balance as it cannot be lover than min. required
        return validatorExist(chains[chainId], acc);
    }
    
    // Returns true, if user has vested enough tokens to become validator and is actively mining, othervise false
    function isActiveValidator(uint256 chainId, address acc) view public returns (bool) {
        // No need to check vesting balance as it cannot be lover than min. required
        return activeValidatorExist(chains[chainId], acc);
    }
    
    // Returns true if user's remaining deposit balance >= min. required deposit and is allowed to transact
    function isAllowedToTransact(uint256 chainId, address acc) view public returns (bool) {
        // No need to check deposit balance as whitelisted flag should be alwyas set accordingly
        return chains[chainId].usersData[acc].transactor.whitelisted;
    }
    
    
    // Returns static chain details
    function getChainStaticDetails(uint256 chainId) external view returns (string memory description, string memory endpoint, bool registered, uint256 maxNumOfValidators, uint256 maxNumOfTransactors,
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
    function getChainDynamicDetails(uint256 chainId) public view returns (bool active, uint256 totalVesting, uint256 validatorsCount, uint256 transactorsCount,
                                                                          uint256 lastValidatorVesting, uint256 lastNotaryBlock, uint256 lastNotaryTimestamp) {
        ChainInfo storage chain = chains[chainId];
        
        active               = chain.active;
        totalVesting         = chain.totalVesting;
        validatorsCount      = chain.validators.list.length;
        transactorsCount     = chain.actNumOfTransactors;
        lastValidatorVesting = chain.usersData[chain.lastValidator].validator.vesting;   
        lastNotaryBlock      = chain.lastNotary.block;
        lastNotaryTimestamp  = chain.lastNotary.timestamp;
    }
    
    // Returns user details
    function getUserDetails(uint256 chainId, address acc) external view returns (uint256 deposit, bool whitelisted, uint256 vesting, 
                                                                                 bool mining, bool prevNotaryMined, bool secondPrevNotaryMined,
                                                                                 bool vestingReqExist, uint256 vestingReqNotary, uint256 vestingReqValue,
                                                                                 bool depositFullWithdrawalReqExist, uint256 depositReqNotary) {
        ChainInfo storage chain = chains[chainId];
         
        deposit                 = chain.usersData[acc].transactor.deposit;
        whitelisted             = chain.usersData[acc].transactor.whitelisted;
        vesting                 = chain.usersData[acc].validator.vesting;
        mining                  = activeValidatorExist(chain, acc);
        prevNotaryMined         = chain.usersData[acc].validator.currentNotaryMined;  
        secondPrevNotaryMined   = chain.usersData[acc].validator.prevNotaryMined;  
        
        if (vestingRequestExist(chain, acc)) {
            vestingReqExist    = true;
            vestingReqNotary           = chain.requests.accounts[acc].vestingRequest.notaryBlock;
            vestingReqValue            = chain.requests.accounts[acc].vestingRequest.newVesting;
        }
        
        if (depositWithdrawalRequestExist(chain, acc)) {
            depositFullWithdrawalReqExist  = true;
            depositReqNotary               = chain.requests.accounts[acc].depositWithdrawalRequest.notaryBlock;
        }
    }
    
    // TODO: delete for mainnet
    // Test Notary increases last notayry block and timestamp - testing method to see vesting/deposit changes that need confirmation
    function testNotary(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        require(msg.sender == chain.creator, "Only chain creator can call this method");
        
        emit Notary(chain.id, chain.lastNotary.block + 1, false);
        
        // Remove validators who signed no block during this notary window and have mining flag == true
        removeInactiveValidators(chain);
    
        // Updates info when the last notary was processed 
        chain.lastNotary.block++;
        chain.lastNotary.timestamp = now;
    
        if (chain.active == false) {
            chain.active = true;
        }
    
        emit Notary(chainId, chain.lastNotary.block, true);               
    }
    
    // Notarization function - calculates user consumption as well as validators rewards
    // First, calculate hash from validators, block_mined, users and userGas
    // then, do ec_recover of the signatures to determine signers
    // check if there is enough signers (total vesting of signers > 50% of all vestings) or total num of signes >= 2/3+1 out of all validators
    // then, calculate reward
    function notary(uint256 chainId, uint256 notaryStartBlock, uint256 notaryEndBlock, address[] memory validators, uint32[] memory blocksMined,
                    address[] memory users, uint64[] memory userGas, uint64 largestTx,
                    uint8[] memory v, bytes32[] memory r, bytes32[] memory s) public {
        
        emit Notary(chainId, notaryEndBlock, false);
                  
        ChainInfo storage chain = chains[chainId];
        require(chain.registered    == true,                            "Invalid chain data: Non-registered chain");
        require(isAllowedToValidate(chainId, msg.sender) == true,       "Sender must have vesting balance > 0");
        require(chain.totalVesting  > 0,                                "Current chain total_vesting == 0, there are no active validators");
        
        require(validators.length       > 0,                            "Invalid statistics data: validators.length == 0");
        require(validators.length       == blocksMined.length,          "Invalid statistics data: validators.length != num of block mined");
        if (chain.maxNumOfValidators != 0) {
            require(validators.length   <= chain.maxNumOfValidators,    "Invalid statistics data: validators.length > maxNumOfValidators");
            require(v.length            <= chain.maxNumOfValidators,    "Invalid statistics data: signatures.length > maxNumOfValidators");
        }
        
        if (chain.maxNumOfTransactors != 0) {
            require(users.length    <= chain.maxNumOfTransactors,   "Invalid statistics data: users.length > maxNumOfTransactors");
        }
        require(users.length        > 0,                            "Invalid statistics data: users.length == 0");
        require(users.length        == userGas.length,              "Invalid statistics data: users.length != usersGas.length");
        
        require(v.length            == r.length,                    "Invalid statistics data: v.length != r.length");
        require(v.length            == s.length,                    "Invalid statistics data: v.length != s.length");
        require(notaryStartBlock    >  chain.lastNotary.block,      "Invalid statistics data: notaryBlock_start <= last known notary block");
        require(notaryEndBlock      >  notaryStartBlock,            "Invalid statistics data: notaryEndBlock <= notaryStartBlock");
        require(largestTx           >  0,                           "Invalid statistics data: Largest tx <= 0");
        
        bytes32 signatureHash = keccak256(abi.encodePacked(notaryEndBlock, validators, blocksMined, users, userGas, largestTx));
        
        // Validates notary conditions(involvedVesting && participation) to statistics to be accepted
        validateNotaryConditions(chain, signatureHash, v, r, s);
        
        // Calculates total cost based on user's usage durint current notary window
        uint256 totalCost = processUsersConsumptions(chain, users, userGas, largestTx);
        
        // In case totalCost == 0, something is wrong and there is no need for notary to continue as there is no tokens to be distributed to the validators.
        // There is probably ongoing coordinated attack based on invalid statistics sent to the notary
        require(totalCost > 0, "Invalid statistics data: users totalUsageCost == 0");
        
        // Calculates and process validator's rewards based on their participation rate and vesting balance
        processValidatorsRewards(chain, notaryStartBlock, notaryEndBlock, validators, blocksMined, totalCost);
        
        // Remove validators who signed no block during this notary window and have mining flag == true
        removeInactiveValidators(chain);
        
        // Updates info when the last notary was processed 
        chain.lastNotary.block = notaryEndBlock;
        chain.lastNotary.timestamp = now;
        
        if (chain.active == false) {
            chain.active = true;
        }
        
        emit Notary(chainId, notaryEndBlock, true);
    }
    
    // Returns list of user's addresses that are allowed to transact - their deposit >= min. required deposit
    function getTransactors(uint256 chainId, uint256 batch) external view returns (address[100] memory transactors, uint256 count, bool end) {
        return getUsers(chains[chainId], true, batch);
    }
    
    // Returns list of active and non-active validators
    function getAllowedToValidate(uint256 chainId, uint256 batch) view external returns (address[100] memory validators, uint256 count, bool end) {
        return getUsers(chains[chainId], false, batch);
    }
    
    // Returns list of active validators
    function getValidators(uint256 chainId, uint256 batch) view external returns (address[100] memory validators, uint256 count, bool end) {
        ChainInfo storage chain = chains[chainId];
        
        count = 0;
        uint256 validatorsTotalCount = chain.validators.list.length;
        
        address acc;
        uint256 i;
        for(i = batch * 100; i < (batch + 1)*100 && i < validatorsTotalCount; i++) {
            acc = chain.validators.list[i];
            
            validators[count] = acc;
            count++;
        }
        
        if (i >= validatorsTotalCount) {
            end = true;
        }
        else {
            end = false;
        }
    }
    
    // Returns list of whitelisted transactors in case transactorsFlag == true, otherwise list of validators (active and non-active)
    function getUsers(ChainInfo storage chain, bool transactorsFlag, uint256 batch) internal view returns (address[100] memory users, uint256 count, bool end) {
        count = 0;
        uint256 usersTotalCount = chain.users.list.length;
        
        address acc;
        uint256 i;
        for(i = batch * 100; i < (batch + 1)*100 && i < usersTotalCount; i++) {
            acc = chain.users.list[i];
            
            // Get transactors (only those who are whitelisted - their depist passed min.required conditions)
            if (transactorsFlag == true) {
                if (chain.usersData[acc].transactor.whitelisted == false) {
                    continue;
                } 
            }
            // Get validators (active and non-active)
            else {
                if (chain.usersData[acc].validator.vesting == 0) {
                    continue;
                }
            }
            
            users[count] = acc;
            count++;
        }
        
        if (i >= usersTotalCount) {
            end = true;
        }
        else {
            end = false;
        }
    }
    
    // Sets mining validator's mining flag to true and emit event so other nodes vote him
    function startMining(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        address acc = msg.sender;
        uint256 validatorVesting = chain.usersData[acc].validator.vesting;
        
        require(chain.registered == true,                           "Non-registered chain");
        require(validatorExist(chain, acc) == true,               "Non-existing validator (0 vesting balance)");
        require(vestingRequestExist(chain, acc) == false,         "Cannot start mining - there is ongoing vesting request");
        
        if (chain.chainValidator != ChainValidator(0)) {
            require(chain.chainValidator.validateNewValidator(validatorVesting, acc, true /* mining */, chain.validators.list.length) == true, "Validator not allowed by external chainvalidator SC");
        }
        
        if (activeValidatorExist(chain, acc) == true) {
            // Emit event even if validator is already active - user might want to explicitely emit this event in case something went wrong on the nodes and
            // others did not vote him
            emit AccountMining(chainId, acc, true);
            
            return;
        }
            
        // Upper limit of validators reached
        if (chain.maxNumOfValidators != 0 && chain.validators.list.length >= chain.maxNumOfValidators) {
            require(validatorVesting > chain.usersData[chain.lastValidator].validator.vesting, "Upper limit of validators reached. Must vest more than the last validator to replace him");
            activeValidatorReplace(chain, acc);
        }
        // There is still empty place for new validator
        else {
            activeValidatorInsert(chain, acc);
        }
    }
  
    // Sets mining validator's mining flag to false and emit event so other nodes unvote
    function stopMining(uint256 chainId) external {
        ChainInfo storage chain = chains[chainId];
        address acc = msg.sender;
        
        require(chain.registered == true, "Non-registered chain");
        require(validatorExist(chain, acc) == true, "Non-existing validator (0 vesting balance)");
    
        if (activeValidatorExist(chain, acc) == false) {
            // Emit event even if validator is already inactive - user might want to explicitely emit this event in case something went wrong on the nodes and
            // others did not unvote him
            emit AccountMining(chainId, acc, false);
            
            return;
        }
        
        activeValidatorRemove(chain, acc);
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
    
    // Inits validator data holder in the users mapping and inserts it into the list of users
    function validatorCreate(ChainInfo storage chain, address acc, uint256 vesting) internal {
        Validator storage validator     = chain.usersData[acc].validator;
        validator.vesting               = uint96(vesting);
        
        // Inits previously notary windows as mined so validator does not get removed from the list of actively mining validators right after the creation
        validator.currentNotaryMined    = true;
        validator.prevNotaryMined       = true;
        validator.secondPrevNotaryMined = true;
        
        // No need to check if validatorExist for the same acc as it is not possible to have vesting > 0 & deosit > 0 at the same time
        insertAcc(chain.users, acc);
    }
    
    // Deinits validator data holder in the users mapping and removes it from the list of users
    function validatorDelete(ChainInfo storage chain, address acc) internal {
        Validator storage validator = chain.usersData[acc].validator;
        
        if (activeValidatorExist(chain, acc) == true) {
            activeValidatorRemove(chain, acc);
        }
        
        if (validator.vesting               != 0)       validator.vesting = 0;
        if (validator.currentNotaryMined    == true)    validator.currentNotaryMined = false;
        if (validator.prevNotaryMined       == true)    validator.prevNotaryMined = false;
        if (validator.secondPrevNotaryMined == true)    validator.secondPrevNotaryMined = false;
        
        // No need to check if transactorExist for the same acc as it is not possible to have vesting > 0 & deosit > 0 at the same time
        removeAcc(chain.users, acc);
    }
    
    // Inserts validator into the list of actively mining validators
    function activeValidatorInsert(ChainInfo storage chain, address acc) internal {
        Validator storage validator = chain.usersData[acc].validator;
        
        // Updates lastValidator in case this is first validator or new validator's vesting balance is less
        if (chain.validators.list.length == 0 || validator.vesting < chain.usersData[chain.lastValidator].validator.vesting) {
            chain.lastValidator = acc;
        }
        
        insertAcc(chain.validators, acc);   
        
        // Updates chain total vesting
        chain.totalVesting = chain.totalVesting.add(validator.vesting);
        
        emit AccountMining(chain.id, acc, true);
    }
    
    // Removes validator from the list of actively mining validators
    function activeValidatorRemove(ChainInfo storage chain, address acc) internal {
        Validator storage validator = chain.usersData[acc].validator;
        
        removeAcc(chain.validators, acc);   
        
        // Updates chain total vesting
        chain.totalVesting = chain.totalVesting.sub(validator.vesting);
        
        // If there is no active validator left, set chain.active flag to false so others might vest in chain without
        // waiting for the next notary window to be processed
        if (chain.validators.list.length == 0) {
            chain.active = false;
            chain.lastValidator = address(0x0);
        }
        // There are still some active validators left
        else {
            // If lastValidator is being removed, find a new validator with the smallest vesting balance
            if (chain.lastValidator == acc) {
                resetLastActiveValidator(chain);
            }
        }
        
        emit AccountMining(chain.id, acc, false);
    }
    
    // Replaces lastValidator for the new one in the list of actively mining validators
    function activeValidatorReplace(ChainInfo storage chain, address acc) internal {
        address accToBeReplaced                 = chain.lastValidator;
        Validator memory validatorToBeReplaced  = chain.usersData[accToBeReplaced].validator;
        Validator memory newValidator           = chain.usersData[acc].validator;
        
        // Updates chain total vesting
        chain.totalVesting = chain.totalVesting.sub(validatorToBeReplaced.vesting);
        chain.totalVesting = chain.totalVesting.add(newValidator.vesting);
        
        // Updates active validarors list
        removeAcc(chain.validators, accToBeReplaced);
        insertAcc(chain.validators, acc);
        
        // Finds a new validator with the smallest vesting balance
        resetLastActiveValidator(chain);
        
        emit AccountMining(chain.id, accToBeReplaced, false);
        emit AccountMining(chain.id, acc, true);
    }
    
    // Resets last validator - the one with the smallest vesting balance
    function resetLastActiveValidator(ChainInfo storage chain) internal {
        address foundLastValidatorAcc       = chain.validators.list[0];
        Validator memory foundLastValidator = chain.usersData[foundLastValidatorAcc].validator;
        
        address actValidatorAcc;
        Validator memory actValidator;
        for (uint256 i = 1; i < chain.validators.list.length; i++) {
            actValidatorAcc = chain.validators.list[i];
            actValidator    = chain.usersData[actValidatorAcc].validator;
            
            if (actValidator.vesting < foundLastValidator.vesting) {
                foundLastValidatorAcc = actValidatorAcc;
                foundLastValidator    = actValidator;
            }
        }
        
        chain.lastValidator = foundLastValidatorAcc;
    }
    
    // Returns true, if acc is in the list of actively mining validators, otherwise false
    function activeValidatorExist(ChainInfo storage chain, address acc) internal view returns (bool) {
        return existAcc(chain.validators, acc);
    }
    
    // Returns true, if acc hase vesting > 0, otherwise false
    function validatorExist(ChainInfo storage chain, address acc) internal view returns (bool) {
        return chain.usersData[acc].validator.vesting > 0;
    }
    
    // Inits transactor data holder in the users mapping and inserts it into the list of users
    function transactorCreate(ChainInfo storage chain, address acc, uint256 deposit) internal {
        Transactor storage transactor = chain.usersData[acc].transactor;
        
        transactor.deposit            = uint96(deposit);
        transactorWhitelist(chain, acc);
        
        // No need to check if validatorExist for the same acc as it is not possible to have vesting > 0 & deosit > 0 at the same time
        insertAcc(chain.users, acc);
    }
    
    // Deinits transactor data holder in the users mapping and removes it from the list of users
    function transactorDelete(ChainInfo storage chain, address acc) internal {
        Transactor storage transactor = chain.usersData[acc].transactor;
        
        if (transactor.deposit != 0) transactor.deposit = 0;
        transactorBlacklist(chain, acc);
        
        // No need to check if validatorExist for the same acc as it is not possible to have vesting > 0 & deosit > 0 at the same time
        removeAcc(chain.users, acc);
    }
    
    // Returns true, if acc hase deposit > 0, otherwise false
    function transactorExist(ChainInfo storage chain, address acc) internal view returns (bool) {
        return chain.usersData[acc].transactor.deposit > 0;
    }
    
    // Blacklists transactor
    function transactorBlacklist(ChainInfo storage chain, address acc) internal {
        Transactor storage transactor   = chain.usersData[acc].transactor;
        
        if (transactor.whitelisted == true) {
            chain.actNumOfTransactors--;
            
            transactor.whitelisted = false;
            emit AccountWhitelist(chain.id, acc, false);
        }
    }
    
    // Whitelists transactor
    function transactorWhitelist(ChainInfo storage chain, address acc) internal {
        Transactor storage transactor   = chain.usersData[acc].transactor;
        
        if (transactor.whitelisted == false) {
            chain.actNumOfTransactors++;
            
            transactor.whitelisted = true;
            emit AccountWhitelist(chain.id, acc, true);
        }
    }
    
    /**************************************************************************************************************************/
    /*********************************** Functions related to the vesting/deposit requests ************************************/
    /**************************************************************************************************************************/
    
    // Creates new vesting request
    function vestingRequestCreate(ChainInfo storage chain, address acc, uint256 vesting) internal {
        RequestsEntry storage entry = chain.requests.accounts[acc];
        
        entry.vestingRequest.exist       = true;
        entry.vestingRequest.newVesting  = uint96(vesting);
        entry.vestingRequest.notaryBlock = chain.lastNotary.block; 
        
        // There is no deposit or vesting ongoing request - create new RequestsEntry structure
        if (entry.index == 0) { // anyRequestExists(chain, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chain.requests.list.push(acc);    
            entry.index = chain.requests.list.length; // indexes are stored + 1
        }
    }

    // Creates new deposit withdrawal request
    function depositWithdrawalRequestCreate(ChainInfo storage chain, address acc) internal {
        RequestsEntry storage entry = chain.requests.accounts[acc];
        
        entry.depositWithdrawalRequest.exist       = true;
        entry.depositWithdrawalRequest.notaryBlock = chain.lastNotary.block; 
        
        // There is no deposit or vesting ongoing request - create new RequestsEntry structure
        if (entry.index == 0) { // anyRequestExists(chain, acc) == false could be used instead
            // There is no ongoing deposit request - create new requests pair structure
            chain.requests.list.push(acc);    
            entry.index = chain.requests.list.length; // indexes are stored + 1
        }
    }

    // Deletes existing requests pair(vesting & deposit) from the internal list of requests
    // This method should never be called directly, vestingRequestDelete & depositWithdrawalRequestDelete should be called instead
    function requestsPairDelete(ChainInfo storage chain, address acc) internal {
        address[] storage requestsList  = chain.requests.list;
        
        uint256 index = chain.requests.accounts[acc].index;
        require(index > 0 && index <= requestsList.length, "RequestsPair delete: invalid index");
    
        // Move an last element of array into the vacated key slot.
        uint256 foundIndex = index - 1;
        uint256 lastIndex  = requestsList.length - 1;
    
        chain.requests.accounts[requestsList[lastIndex]].index = foundIndex + 1;
        requestsList[foundIndex] = requestsList[lastIndex];
        requestsList.length--;
    
        delete chain.requests.accounts[acc];
    }
    
    function vestingRequestDelete(ChainInfo storage chain, address acc) internal {
        // There is no ongoing deposit request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].depositWithdrawalRequest.exist == false) {
            requestsPairDelete(chain, acc);
            return;
        } 
        
        // There is ongoing deposit request for this account - only reset vesting request
        VestingRequest storage request = chain.requests.accounts[acc].vestingRequest;
        request.exist          = false;
        request.notaryBlock    = 0;
        request.newVesting     = 0;
    }
    
    function depositWithdrawalRequestDelete(ChainInfo storage chain, address acc) internal {
        // There is no ongoing vesting request for this account - delete whole requests struct 
        if (chain.requests.accounts[acc].vestingRequest.exist == false) {
            requestsPairDelete(chain, acc);
            return;
        } 
        
        // There is ongoing vesting request for this account - only reset vesting request
        DepositWithdrawalRequest storage request = chain.requests.accounts[acc].depositWithdrawalRequest;
        request.exist          = false;
        request.notaryBlock    = 0;
    }
    
    // Checks if acc has any ongoing vesting or deposit request
    function anyRequestExists(ChainInfo storage chain, address acc) internal view returns (bool) {
        return chain.requests.accounts[acc].index != 0;
    }
    
    // Checks if acc has any ongoing vesting request
    function vestingRequestExist(ChainInfo storage chain, address acc) internal view returns (bool) {
        return chain.requests.accounts[acc].vestingRequest.exist;
    }
    
    // Checks if acc has any ongoing DEPOSIT WITHDRAWAL request
    function depositWithdrawalRequestExist(ChainInfo storage chain, address acc) internal view returns (bool) {
        return chain.requests.accounts[acc].depositWithdrawalRequest.exist;
    }
    
    // Full vesting withdrawal  and vesting increase are procesed in 2 steps with confirmaition
    // Immediate full withdrawal is not allowed as validators might already mine some blocks so they can get rewards 
    // based on theirs vesting balance for that
    function requestVest(ChainInfo storage chain, uint256 vesting, address acc) internal {
        Validator storage validator = chain.usersData[acc].validator;
        
        uint256 validatorVesting = validator.vesting;
        
        // Vesting increase - process in 2 steps
        if (vesting > validatorVesting) {
            uint256 toVest = vesting - validatorVesting;
            token.transferFrom(acc, address(this), toVest);
        }
        // Vesting decrease - process immediately
        else if (vesting != 0) {
            uint256 toWithdraw = validatorVesting - vesting;
            
            // If validator is actively mining, decrease chain's total vesting
            if (activeValidatorExist(chain, acc) == true) {
                chain.totalVesting = chain.totalVesting.sub(toWithdraw);
            }
         
            validator.vesting = uint96(vesting);    
            
            // Transfers tokens
            token.transfer(acc, toWithdraw);
            
            emit VestInChain(chain.id, acc, vesting, chain.lastNotary.block, true);
            return;
        }
        
        vestingRequestCreate(chain, acc, vesting);
        emit VestInChain(chain.id, acc, vesting, chain.requests.accounts[acc].vestingRequest.notaryBlock, false);
        
        return;
    }
    
    function confirmVest(ChainInfo storage chain, address acc) internal {
        Validator storage validator             = chain.usersData[acc].validator;
        VestingRequest memory request           = chain.requests.accounts[acc].vestingRequest;
        
        vestingRequestDelete(chain, acc);
        uint256 origVesting = validator.vesting;
        
        // Vesting increase
        if (request.newVesting > origVesting) {
            // Non-existing validator - internally creates new one
            if (validatorExist(chain, acc) == false) {
                validatorCreate(chain, acc, request.newVesting);
            }
            // Existing validator
            else {
                validator.vesting = uint96(request.newVesting);
                
                if (activeValidatorExist(chain, acc) == true) {
                    chain.totalVesting = chain.totalVesting.add(request.newVesting - origVesting);
                }    
            }
        }
        // Full vesting withdrawal - stopMining must be called before
        else {
            uint256 toWithdraw = origVesting;
            validatorDelete(chain, acc);
            
            // Transfers tokens
            token.transfer(acc, toWithdraw);
        }
        
        emit VestInChain(chain.id, acc, request.newVesting, request.notaryBlock, true);
    }
    
    function requestDeposit(ChainInfo storage chain, uint256 deposit, address acc) internal {
        Transactor storage transactor = chain.usersData[acc].transactor;
        
        // If user wants to withdraw whole deposit
        if (deposit == 0) {
            // Chain is not active and last notary is older than 2*CHAIN_INACTIVITY_TIMEOUT - enable full deposit withdrawal immmediately
            if (chain.active == false) {
                uint256 toWithdraw = transactor.deposit;
                transactorDelete(chain, acc);
                
                // Withdraw whole deposit
                token.transfer(acc, toWithdraw);
                
                emit DepositInChain(chain.id, acc, deposit, chain.lastNotary.block, true);
            }
            // Chain is active - create withdrawal request and process full deposit withdrawal in 2 steps
            else {
                depositWithdrawalRequestCreate(chain, acc);
                
                transactorBlacklist(chain, acc);
                emit DepositInChain(chain.id, acc, deposit, chain.requests.accounts[acc].depositWithdrawalRequest.notaryBlock, false);  
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
            if (transactorExist(chain, acc) == false) {
                transactorCreate(chain, acc, deposit);
            }
            else {
                transactor.deposit = uint96(deposit);
                transactorWhitelist(chain, acc);
            }
        }
        
        emit DepositInChain(chain.id, acc, deposit, chain.lastNotary.block, true);
    }
    
    function confirmDepositWithdrawal(ChainInfo storage chain, address acc) internal {
        uint256 toWithdraw = chain.usersData[acc].transactor.deposit;
        uint256 requestNotaryBlock = chain.requests.accounts[acc].depositWithdrawalRequest.notaryBlock;
        
        transactorDelete(chain, acc);
        depositWithdrawalRequestDelete(chain, acc);
        
        // Withdraw whole deposit
        token.transfer(acc, toWithdraw);
        
        emit DepositInChain(chain.id, acc, 0, requestNotaryBlock, true);
    }
    
    /**************************************************************************************************************************/
    /*************************************************** Other functions ******************************************************/
    /**************************************************************************************************************************/

    constructor(ERC20 _token) public {
        token = _token;
    }
  
    // Process users consumption based on their usage
    function processUsersConsumptions(ChainInfo storage chain, address[] memory users, uint64[] memory userGas, uint64 largestTx) internal returns (uint256 totalCost) {
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
            Transactor storage transactor = chain.usersData[acc].transactor;
            transactorDeposit = transactor.deposit;
            
            // This can happen only if there is non-registered transactor(user) in statistics, which means that there is probaly
            // ongoing coordinated attack based on invalid statistics sent to the notary
            // Ignores non-registred user
            if (transactorExist(chain, acc) == false || userGas[i] == 0) {
                // Let nodes know that this user is not allowed to transact only if chain is active - in case it is not and becomes active again 
                // there might be some users that already withdrawed their deposit  
                if (chain.active == true) {
                    emit AccountWhitelist(chain.id, users[i], false);
                }
                continue;
            }
            
            userCost = (userGas[i] * largestReward) / largestTx;
            
            // This can happen only if user runs out of tokens(which should not happen due to min.required deposit)
            if(userCost > transactorDeposit ) {
                userCost = transactorDeposit;
            
                transactorDelete(chain, acc);
            }
            else {
                transactorDeposit -= userCost;
                
                // Updates user's stored deposit balance based on his usage
                transactor.deposit = uint96(transactorDeposit);
                
                // Check if user's deposit balance is >= min. required deposit conditions
                if (checkLitionMinDeposit(transactorDeposit) == false) {
                    transactorBlacklist(chain, acc);
                }
                else if (chain.chainValidator != ChainValidator(0) && chain.chainValidator.validateNewTransactor(transactorDeposit, acc, chain.actNumOfTransactors) == false) {
                    // If not, do not allow him to transact anymore
                    transactorBlacklist(chain, acc);
                } 
            }
            
            // Adds user's cost to the total cost
            // No need for safe math as max possible userCost is 10^64 * 10^17 allows 10^175 users to ovewflow uint256 and that is impossible because of gas 
            totalCost += userCost;  
        }
    }

    // Process validators rewards based on their participation rate(how many blocks they signed) and their vesting balance
    function processValidatorsRewards(ChainInfo storage chain, uint256 startNotaryBlock, uint256 endNotaryBlock, address[] memory validators, uint32[] memory blocksMined, uint256 litToDistribute) internal {
        // Min. vesting balance to be a trust node. Trust Nodes haved doubled(virtually) vesting
        uint256 minTrustNodeVesting = 50000*LIT_PRECISION; 
        
        // How many block could validator mined since the last notary in case he did sign every possible block 
        uint256 maxBlocksMined = endNotaryBlock - startNotaryBlock;
        
        // Total involved vesting 
        uint256 totalInvolvedVesting = 0;
        
        // Selected validator's vesting balance
        uint256 validatorVesting;
        
        address actValidatorAcc;
        // Runs through all validators and calculates total involved vesting.
        // in the statistics
        for(uint256 i = 0; i < validators.length; i++) {
            actValidatorAcc = validators[i];
            
            // This can happen only if there is validator with 0 vesting balance in statistics or there is 0 mined blocks for this validators, which means that
            // there is probably ongoing coordinated attack based on invalid statistics sent to the notary
            if (validatorExist(chain, actValidatorAcc) == false || blocksMined[i] == 0) {
                continue;
            }
            
            // Updates validator's mining flags statistics
            Validator storage validator     = chain.usersData[actValidatorAcc].validator;
            validator.secondPrevNotaryMined = validator.prevNotaryMined;
            validator.prevNotaryMined       = validator.currentNotaryMined;
            validator.currentNotaryMined    = true;
            
            validatorVesting = chain.usersData[actValidatorAcc].validator.vesting;
            
            // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
            if (validatorVesting >= minTrustNodeVesting) {
                // Validator's stored vesting is max uint96
                validatorVesting *= 2;
            }
            
            // No need for safe math
            // max possible (blocksMined[i] / maxBlocksMined) valuse is 10^32, max possible validatorVesting value is 10^96, when virtually doubled it is 10^192, in total 42*10^224
            // so to overflow uint256 there would have to be 10^32 validators, which is impossible because of gas
            totalInvolvedVesting += (blocksMined[i] * validatorVesting) / maxBlocksMined;
        }
        
        // In case totalInvolvedVesting == 0, something is wrong and there is no need for notary to continue as rewards cannot be calculated. It might happen
        // as edge case when the last validator stopped mining durint current notary window or there is ongoing coordinated attack based on invalid statistics sent to the notary
        require(totalInvolvedVesting > 0, "totalInvolvedVesting == 0. Invalid statistics or 0 active validators left in the chain");
        
        
        uint256 validatorReward;
        // Runs through all validators and calculates their reward based on:
        //     1. How many blocks out of max_blocks_mined each validator signed
        //     2. How many token each validator vested
        for(uint256 i = 0; i < validators.length; i++) {
            actValidatorAcc = validators[i];
            
            // This can happen only if there is validator with 0 vesting balance in statistics or there is 0 mined blocks for this validators, which means that
            // there is probably ongoing coordinated attack based on invalid statistics sent to the notary
            if (validatorExist(chain, actValidatorAcc) == false || blocksMined[i] == 0) {
                continue;
            } 
            
            validatorVesting = chain.usersData[actValidatorAcc].validator.vesting;
            
            // In case validator is trust node (his vesting >= 50k LIT tokens) - virtually double his vesting
            if (validatorVesting >= minTrustNodeVesting) {
                // Validator's stored vesting is max uint96
                validatorVesting *= 2;
            }
            
            // No need for safe math as max value of (blocksMined[i] / maxBlocksMined) is 10^32, max value of (validatorVesting / totalInvolvedVesting) is 1 and 
            // max value of litToDistribute(calculated in processUsersConsumptions) is 10^97, so max possible validator reward is 10^32 * 1 * 10^97 = 10^129
            validatorReward = (validatorVesting * blocksMined[i] * litToDistribute) / maxBlocksMined / totalInvolvedVesting;
            token.transfer(validators[i], validatorReward);
            
            // No need for safe math as validator reward is calculated as fraction of total litToDistribute and sum of all validators rewards must always be <= litToDistribute
            litToDistribute -= validatorReward;
        }
        
        if(litToDistribute > 0) {
            // Sends the rest(math rounding) to the validator, who called notary function
            token.transfer(msg.sender, litToDistribute);
        }
    }
   
   // Validates notary conditions(involvedVesting && participation) to statistics to be accepted
    function validateNotaryConditions(ChainInfo storage chain, bytes32 signatureHash, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) internal view {
        uint256 involvedVestingSum = 0;
        uint256 involvedSignaturesCount = 0;
        
        bool[] memory signedValidators = new bool[](chain.validators.list.length); 
        
        address signerAcc;
        for(uint256 i = 0; i < v.length; i++) {
            signerAcc = ecrecover(signatureHash, v[i], r[i], s[i]);
            
            // In case statistics is signed by validator, who is not registered in SC, ignore him   
            if (activeValidatorExist(chain, signerAcc) == false) {
                continue;
            }
            
            uint256 validatorIdx = chain.validators.listIndex[signerAcc] - 1;
            
            // In case there is duplicit signature from the same validator, ignore it
            if (signedValidators[validatorIdx] == true) {
                continue;
            }
            else {
                signedValidators[validatorIdx] = true;
            }
            
            
            involvedVestingSum += chain.usersData[signerAcc].validator.vesting;
            involvedSignaturesCount++;
        }
        
        delete signedValidators;
        
        // There must be more than 50% out of total possible vesting involved in signatures
        if (chain.involvedVestingNotaryCond == true) {
            // There must be more than 50% out of total possible vesting involved
            require(involvedVestingSum*2 > chain.totalVesting, "Invalid statistics data: involvedVesting <= 50% of chain.totalVesting");
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
    
    // Removes validators that did not mine at all during the last 3 notary windows
    function removeInactiveValidators(ChainInfo storage chain) internal {
        address validatorAcc;
        for (uint256 i = 0; i < chain.validators.list.length; i++) {
            validatorAcc = chain.validators.list[i];
            Validator memory validator = chain.usersData[validatorAcc].validator;
           
            if (validator.currentNotaryMined || validator.prevNotaryMined || validator.secondPrevNotaryMined) {
                continue;
            }
           
            activeValidatorRemove(chain, validatorAcc);
        } 
    }
   
    // Checks if chain is active(successfull notary processed during last CHAIN_INACTIVITY_TIMEOUT), if not set it active flag to false
    // If last notary is older than CHAIN_INACTIVITY_TIMEOUT, it means that validators cannot reach consensus or there is no active validator and chain is basically stuck.
    function checkAndSetChainActivity(ChainInfo storage chain) internal {
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
}