package litionScClient

import (
	"math/big"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/common"
	"gitlab.com/lition/lition/core/types"
	"gitlab.com/lition/lition/ethclient"
)

// ContractClient contains variables needed for communication with lition smart contract
type ContractClient struct {
	ethClient                   *ethclient.Client
	scAddress                   common.Address
	scClient                    *LitionScClient
	chainID                     *big.Int // chainID on top of which all sc calls are made
	accMiningEventListener      *AccMiningEventListener
	accWhitelistedEventListener *AccWhitelistedEventListener
	vestInChainEventListener    *VestInChainEventListener
	depositInChainEventListener *DepositInChainEventListener
	notaryEventListener         *NotaryEventListener
	notaryResetEventListener    *NotaryResetEventListener
}

func NewClient(ethClientURL string, scAddress string, chainID *big.Int) (*ContractClient, error) {
	contractClient := new(ContractClient)
	ethClient, err := ethclient.Dial(ethClientURL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.ethClient = ethClient
	contractClient.scAddress = common.HexToAddress(scAddress)
	contractClient.chainID = chainID

	pScClient, err := NewLitionScClient(contractClient.scAddress, contractClient.ethClient)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.scClient = pScClient

	contractClient.accMiningEventListener = nil
	contractClient.accWhitelistedEventListener = nil
	contractClient.vestInChainEventListener = nil
	contractClient.depositInChainEventListener = nil
	contractClient.notaryEventListener = nil
	contractClient.notaryResetEventListener = nil

	return contractClient, nil
}

func (contractClient *ContractClient) InitAccMiningEventListener() error {
	var err error
	contractClient.accMiningEventListener, err = NewAccMiningEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitAccWhitelistedEventListener() error {
	var err error
	contractClient.accWhitelistedEventListener, err = NewAccWhitelistedEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitVestInChainEventListener() error {
	var err error
	contractClient.vestInChainEventListener, err = NewVestInChainEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitDepositInChainEventListener() error {
	var err error
	contractClient.depositInChainEventListener, err = NewDepositInChainEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitNotaryEventListener() error {
	var err error
	contractClient.notaryEventListener, err = NewNotaryEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitNotaryResetEventListener() error {
	var err error
	contractClient.notaryResetEventListener, err = NewNotaryResetEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) DeInit() {
	if contractClient.accMiningEventListener != nil {
		contractClient.accMiningEventListener.DeInit()
	}
	if contractClient.accWhitelistedEventListener != nil {
		contractClient.accWhitelistedEventListener.DeInit()
	}
	if contractClient.vestInChainEventListener != nil {
		contractClient.vestInChainEventListener.DeInit()
	}
	if contractClient.depositInChainEventListener != nil {
		contractClient.depositInChainEventListener.DeInit()
	}
	if contractClient.notaryEventListener != nil {
		contractClient.notaryEventListener.DeInit()
	}
	if contractClient.notaryResetEventListener != nil {
		contractClient.notaryResetEventListener.DeInit()
	}

	contractClient.chainID = nil
	contractClient.accMiningEventListener = nil
	contractClient.accWhitelistedEventListener = nil
	contractClient.vestInChainEventListener = nil
	contractClient.depositInChainEventListener = nil
	contractClient.notaryEventListener = nil
	contractClient.notaryResetEventListener = nil

	contractClient.ethClient.Close()
}

func (contractClient *ContractClient) Start_accMiningEventListener(f func(*LitionScClientAccountMining)) {
	listener := contractClient.accMiningEventListener
	if listener == nil {
		log.Fatal("Trying to start 'AccountMining' listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	initialized := true
	for {
		if initialized == true {
			retErr := listener.Start(f)
			// Listener was manually stopped, do not try to start it again
			if retErr == nil {
				return
			}
			log.Error("Start accMiningEventListener err: '", retErr, "'. Try to reinit.")
		}

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)

		err := listener.ReInit()
		if err == nil {
			log.Info("Reinit successfull")
			initialized = true
		} else {
			log.Error("Reinit fail")
			initialized = false
		}
	}
}

func (contractClient *ContractClient) Start_accWhitelistedEventListener(f func(*LitionScClientAccountWhitelisted)) {
	listener := contractClient.accWhitelistedEventListener
	if listener == nil {
		log.Fatal("Trying to start 'AccountWhitelist' listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	initialized := true
	for {
		if initialized == true {
			retErr := listener.Start(f)
			// Listener was manually stopped, do not try to start it again
			if retErr == nil {
				return
			}
			log.Error("Start AccountWhitelistEventListener err: '", retErr, "'. Try to reinit.")
		}

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)

		err := listener.ReInit()
		if err == nil {
			log.Info("Reinit successfull")
			initialized = true
		} else {
			log.Error("Reinit fail")
			initialized = false
		}
	}
}

func (contractClient *ContractClient) Start_vestInChainEventListener(f func(*LitionScClientVestInChain)) {
	listener := contractClient.vestInChainEventListener
	if listener == nil {
		log.Fatal("Trying to start 'VestInChain' listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	initialized := true
	for {
		if initialized == true {
			retErr := listener.Start(f)
			// Listener was manually stopped, do not try to start it again
			if retErr == nil {
				return
			}
			log.Error("Start VestInChainEventListener err: '", retErr, "'. Try to reinit.")
		}

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)

		err := listener.ReInit()
		if err == nil {
			log.Info("Reinit successfull")
			initialized = true
		} else {
			log.Error("Reinit fail")
			initialized = false
		}
	}
}

func (contractClient *ContractClient) Start_depositInChainEventListener(f func(*LitionScClientDepositInChain)) {
	listener := contractClient.depositInChainEventListener
	if listener == nil {
		log.Fatal("Trying to start 'DepositInChain' listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	initialized := true
	for {
		if initialized == true {
			retErr := listener.Start(f)
			// Listener was manually stopped, do not try to start it again
			if retErr == nil {
				return
			}
			log.Error("Start DepositInChainEventListener err: '", retErr, "'. Try to reinit.")
		}

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)

		err := listener.ReInit()
		if err == nil {
			log.Info("Reinit successfull")
			initialized = true
		} else {
			log.Error("Reinit fail")
			initialized = false
		}
	}
}

func (contractClient *ContractClient) Start_notaryEventListener(f func(*LitionScClientNotary)) {
	listener := contractClient.notaryEventListener
	if listener == nil {
		log.Fatal("Trying to start 'Notary' listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	initialized := true
	for {
		if initialized == true {
			retErr := listener.Start(f)
			// Listener was manually stopped, do not try to start it again
			if retErr == nil {
				return
			}
			log.Error("Start NotaryEventListener err: '", retErr, "'. Try to reinit.")
		}

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)

		err := listener.ReInit()
		if err == nil {
			log.Info("Reinit successfull")
			initialized = true
		} else {
			log.Error("Reinit fail")
			initialized = false
		}
	}
}

func (contractClient *ContractClient) Start_notaryResetEventListener(f func(*LitionScClientNotaryReset)) {
	listener := contractClient.notaryResetEventListener
	if listener == nil {
		log.Fatal("Trying to start 'NotaryReset' listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	initialized := true
	for {
		if initialized == true {
			retErr := listener.Start(f)
			// Listener was manually stopped, do not try to start it again
			if retErr == nil {
				return
			}
			log.Error("Start NotaryResetEventListener err: '", retErr, "'. Try to reinit.")
		}

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)

		err := listener.ReInit()
		if err == nil {
			log.Info("Reinit successfull")
			initialized = true
		} else {
			log.Error("Reinit fail")
			initialized = false
		}
	}
}

func (contractClient *ContractClient) StartMining(auth *bind.TransactOpts) (*types.Transaction, error) {
	return contractClient.scClient.StartMining(auth, contractClient.chainID)
}

func (contractClient *ContractClient) StopMining(auth *bind.TransactOpts) (*types.Transaction, error) {
	return contractClient.scClient.StopMining(auth, contractClient.chainID)
}

func (contractClient *ContractClient) GetTransactors() ([]common.Address, error) {
	var accountsWhitelist []common.Address
	zeroCount := big.NewInt(0)

	for batchID := big.NewInt(0); ; batchID.Add(batchID, big.NewInt(1)) {
		transactorsList, err := contractClient.scClient.GetTransactors(&bind.CallOpts{}, contractClient.chainID, batchID)
		if err != nil {
			return nil, err
		}

		cmpResult := transactorsList.Count.Cmp(zeroCount)
		if cmpResult == 1 {
			accountsWhitelist = append(accountsWhitelist, transactorsList.Transactors[0:transactorsList.Count.Int64()]...)
		}

		if transactorsList.End == true {
			break
		}
	}

	return accountsWhitelist, nil
}

func (contractClient *ContractClient) GetAllowedToValidate() ([]common.Address, error) {
	var validators []common.Address
	zeroCount := big.NewInt(0)

	for batchID := big.NewInt(0); ; batchID.Add(batchID, big.NewInt(1)) {
		validatorsList, err := contractClient.scClient.GetAllowedToValidate(&bind.CallOpts{}, contractClient.chainID, batchID)
		if err != nil {
			return nil, err
		}

		cmpResult := validatorsList.Count.Cmp(zeroCount)
		if cmpResult == 1 {
			validators = append(validators, validatorsList.Validators[0:validatorsList.Count.Int64()]...)
		}

		if validatorsList.End == true {
			break
		}
	}

	return validators, nil
}

func (contractClient *ContractClient) GetValidators() ([]common.Address, error) {
	var activeValidators []common.Address
	zeroCount := big.NewInt(0)

	for batchID := big.NewInt(0); ; batchID.Add(batchID, big.NewInt(1)) {
		validatorsList, err := contractClient.scClient.GetValidators(&bind.CallOpts{}, contractClient.chainID, batchID)
		if err != nil {
			return nil, err
		}

		cmpResult := validatorsList.Count.Cmp(zeroCount)
		if cmpResult == 1 {
			activeValidators = append(activeValidators, validatorsList.Validators[0:validatorsList.Count.Int64()]...)
		}

		if validatorsList.End == true {
			break
		}
	}

	return activeValidators, nil
}

func (contractClient *ContractClient) Notary(auth *bind.TransactOpts,
	notary_start_block *big.Int,
	notary_end_block *big.Int,
	miners []common.Address,
	blocks_mined []uint32,
	users []common.Address,
	user_gas []uint64,
	largest_tx uint64,
	v []uint8,
	r [][32]byte,
	s [][32]byte) (*types.Transaction, error) {

	return contractClient.scClient.Notary(auth, contractClient.chainID, notary_start_block, notary_end_block, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

func (contractClient *ContractClient) GetChainStaticDetails() (struct {
	Description                string
	Endpoint                   string
	Registered                 bool
	MinRequiredDeposit         *big.Int
	MinRequiredVesting         *big.Int
	RewardBonusRequiredVesting *big.Int
	RewardBonusPercentage      *big.Int
	NotaryPeriod               *big.Int
	MaxNumOfValidators         *big.Int
	MaxNumOfTransactors        *big.Int
	InvolvedVestingNotaryCond  bool
	ParticipationNotaryCond    bool
}, error) {
	return contractClient.scClient.GetChainStaticDetails(&bind.CallOpts{}, contractClient.chainID)
}

func (contractClient *ContractClient) GetChainDynamicDetails() (struct {
	Active               bool
	TotalVesting         *big.Int
	ValidatorsCount      *big.Int
	TransactorsCount     *big.Int
	LastValidatorVesting *big.Int
	LastNotaryBlock      *big.Int
	LastNotaryTimestamp  *big.Int
}, error) {
	return contractClient.scClient.GetChainDynamicDetails(&bind.CallOpts{}, contractClient.chainID)
}

func (contractClient *ContractClient) GetUserDetails(userAddressStr string) (struct {
	Deposit                       *big.Int
	Whitelisted                   bool
	Vesting                       *big.Int
	LastVestingIncreaseTime       *big.Int
	Mining                        bool
	PrevNotaryMined               bool
	VestingReqExist               bool
	VestingReqNotary              *big.Int
	VestingReqValue               *big.Int
	DepositFullWithdrawalReqExist bool
	DepositReqNotary              *big.Int
}, error) {
	userAddress := common.HexToAddress(userAddressStr)

	return contractClient.scClient.GetUserDetails(&bind.CallOpts{}, contractClient.chainID, userAddress)
}
