pragma solidity >= 0.5.11;

interface ChainValidator {
   function validateNewValidator(uint256 vesting, address acc, bool mining, uint256 actNumOfValidators) external returns (bool);
   function validateNewTransactor(uint256 deposit, address acc, uint256 actNumOfTransactors) external returns (bool);
}

contract EnergyChainValidator is ChainValidator {
    /**************************************************************************************************************************/
    /************************************************** Constants *************************************************************/
    /**************************************************************************************************************************/
    
    // Token precision. 1 LIT token = 1*10^18
    uint256 constant LIT_PRECISION               = 10**18;
    
    // How many toknes user must vest to be a trust node (his vesting is virtually doubled when calculatiing rewards)
    uint256 constant TRUST_NODE_VESTING          = 50000*LIT_PRECISION;
    
    // Min deposit value
    uint256 constant MIN_DEPOSIT                 = 5000*LIT_PRECISION;
    
    // Max deposit value
    uint256 constant MAX_DEPOSIT                 = ~uint96(0);
    
    // Min vesting value
    uint256 constant MIN_VESTING                 = ~uint96(0);
    
    // Min vesting value
    uint256 constant MAX_VESTING                 = 500000*LIT_PRECISION;
    
    // Max number of validators for Lition Energy chain
    uint256 MAX_VALIDATORS_COUNT                 = 21;
    
    // Max number of validators reserved for Lition internal accounts
    uint256 MAX_RESERVED_VALIDATORS_COUNT        = 6;
    
    
    /**************************************************************************************************************************/
    /*********************************** Structs and functions related to the list of users ***********************************/
    /**************************************************************************************************************************/
    
    // Iterable map that is used only together with the Users mapping as data holder
    struct IterableMap {
        // map of indexes to the list array
        // indexes are shifted +1 compared to the real indexes of this list, because 0 means non-existing element
        mapping(address => uint256) listIndex;
        // list of addresses 
        address[]                   list;        
    }    
    
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
    
    
    /**************************************************************************************************************************/
    /******************************************* Other functions and functions ************************************************/
    /**************************************************************************************************************************/


    // List of admins - they can add/remove whitelisted validators and users
    IterableMap private admins;
    
    // List of Lition reserved validators
    IterableMap private reservedValidators;
    
    // List of whitelisted users who can deposit
    IterableMap private whitelistedUsers;
    
    constructor() public {
        insertAcc(admins, msg.sender);
    }
    
    // Validates new validator
    function validateNewValidator(uint256 vesting, address acc, bool mining, uint256 actNumOfValidators) external returns (bool) {
        // Reserved account wants to become validator
        if (existAcc(reservedValidators, acc) == true && vesting >= MIN_VESTING && vesting <= MAX_VESTING) {
            return true;
        }
        
        // Non-reserved account wants to become validator
        if (reservedValidators.list.length + actNumOfValidators + 1 > MAX_VALIDATORS_COUNT) {
            return false;
        }
        
        if (vesting < MIN_VESTING || vesting > MAX_VESTING) {
            return false;
        }
        
        return true;
    }
    
    // Validates new transactor
    function validateNewTransactor(uint256 deposit, address acc, uint256 actNumOfTransactors) external returns (bool) {
        if (existAcc(whitelistedUsers, acc) == true && deposit >= MIN_DEPOSIT) {
            return  true;
        }
        
        return false;
    }
    
    // Adds new reserved validators
    function addReservedValidators(address[] calldata accounts) external {
        require(reservedValidators.list.length + accounts.length <= MAX_RESERVED_VALIDATORS_COUNT, "Trying to add too many new reserved validators. There is not enough free spots left. Max(6)");
        
        addUsers(reservedValidators, accounts);
    }
    
    // Removes reserved validators
    function removeReservedValidators(address[] calldata accounts) external {
        require(reservedValidators.list.length > 0, "There are no validators to be removed");
        
        removeUsers(reservedValidators, accounts);
    }
    
    // Adds new whitelisted users
    function whitelistUsers(address[] calldata accounts) external {
        addUsers(whitelistedUsers, accounts);
    }
    
    // Removes existing whitelisted users
    function removeWhitelistedUsers(address[] calldata accounts) external {
        require(whitelistedUsers.list.length > 0, "There are no whitelisted users to be removed");
        
        removeUsers(whitelistedUsers, accounts);
    }

    // Adds new admins
    function addAdmins(address[] calldata accounts) external {
        addUsers(admins, accounts);
    }
    
    // Removes existing admins
    function removeAdmins(address[] calldata accounts) external {
        require(admins.list.length > 0, "There are no admins to be removed");
        
        removeUsers(admins, accounts);
    }
    
    // Returns list of admins
    function getAdmins(uint256 batch) external view returns (address[100] memory accounts, uint256 count, bool end) {
        return getUsers(admins, batch);
    }
    
    // Returns list of whitelisted users
    function getWhitelistedUsers(uint256 batch) external view returns (address[100] memory accounts, uint256 count, bool end) {
        return getUsers(whitelistedUsers, batch);
    }
    
    // Returns list of reserved validators
    function getReservedValidators(uint256 batch) external view returns (address[100] memory accounts, uint256 count, bool end) {
        return getUsers(reservedValidators, batch);
    }
    
    
    // Returns list of suers users
    function getUsers(IterableMap storage internalUsersGroup, uint256 batch) internal view returns (address[100] memory users, uint256 count, bool end) {
        count = 0;
        uint256 usersTotalCount = internalUsersGroup.list.length;
        
        uint256 i;
        for(i = batch * 100; i < (batch + 1)*100 && i < usersTotalCount; i++) {
            users[count] = internalUsersGroup.list[i];
            count++;
        }
        
        if (i >= usersTotalCount) {
            end = true;
        }
        else {
            end = false;
        }
    }
    
    function addUsers(IterableMap storage internalUsersGroup, address[] memory users) internal {
        require(existAcc(admins, msg.sender) == true, "Only admins can do internal changes");
        require(users.length <= 100, "Max number of processed users is 100");
        
        for (uint256 i = 0; i < users.length; i++) {
            if (existAcc(internalUsersGroup, users[i]) == false) {
                insertAcc(internalUsersGroup, users[i]);
            }    
        }
    }
    
    function removeUsers(IterableMap storage internalUsersGroup, address[] memory users) internal {
        require(existAcc(admins, msg.sender) == true, "Only admins can remove whitelisted users");
        require(users.length <= 100, "Max number of processed users is 100");
        
        for (uint256 i = 0; i < users.length; i++) {
            if (existAcc(internalUsersGroup, users[i]) == true) {
                removeAcc(internalUsersGroup, users[i]);
            }    
        }
    }  
}
