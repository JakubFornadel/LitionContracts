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
	ethClient                 *ethclient.Client
	scAddress                 common.Address
	scClient                  *LitionScClient
	chainID                   *big.Int // chainID on top of which all sc calls are made
	accMiningEventListener    *AccMiningEventListener
	accWhitelistEventListener *AccWhitelistEventListener
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
	contractClient.accWhitelistEventListener = nil

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

func (contractClient *ContractClient) InitAccWhitelistEventListener() error {
	var err error
	contractClient.accWhitelistEventListener, err = NewAccWhitelistEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) DeInit() {
	if contractClient.accMiningEventListener != nil {
		contractClient.accMiningEventListener.DeInit()
	}
	if contractClient.accWhitelistEventListener != nil {
		contractClient.accWhitelistEventListener.DeInit()
	}

	contractClient.chainID = nil
	contractClient.accMiningEventListener = nil
	contractClient.accWhitelistEventListener = nil
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

func (contractClient *ContractClient) Start_accWhitelistEventListener(f func(*LitionScClientAccountWhitelist)) {
	listener := contractClient.accWhitelistEventListener
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
			log.Error("Start accWhitelistEventListener err: '", retErr, "'. Try to reinit.")
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

func (contractClient *ContractClient) IsAllowedToValidate(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	return contractClient.scClient.IsAllowedToValidate(&bind.CallOpts{}, contractClient.chainID, userAddress)
}

func (contractClient *ContractClient) IsActiveValidator(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	return contractClient.scClient.IsActiveValidator(&bind.CallOpts{}, contractClient.chainID, userAddress)
}

func (contractClient *ContractClient) IsAllowedToTransact(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	return contractClient.scClient.IsAllowedToTransact(&bind.CallOpts{}, contractClient.chainID, userAddress)
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
	user_gas []uint32,
	largest_tx uint32,
	v []uint8,
	r [][32]byte,
	s [][32]byte) (*types.Transaction, error) {

	return contractClient.scClient.Notary(auth, contractClient.chainID, notary_start_block, notary_end_block, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
}

func (contractClient *ContractClient) GetChainStaticDetails() (struct {
	Description               string
	Endpoint                  string
	Registered                bool
	MaxNumOfValidators        *big.Int
	MaxNumOfTransactors       *big.Int
	InvolvedVestingNotaryCond bool
	ParticipationNotaryCond   bool
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
