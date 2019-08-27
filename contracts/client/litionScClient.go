package litionScClient

import (
	"math/big"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/common"
	"gitlab.com/lition/lition/ethclient"
)

// ContractClient contains variables needed for communication with lition smart contract
type ContractClient struct {
	ethClient                *ethclient.Client
	scAddress                common.Address
	scClient                 *LitionScClient
	chainID                  *big.Int // chainID on top of which all sc calls are made
	startMiningEventListener *StartMiningEventListener
	stopMiningEventListener  *StopMiningEventListener
	depositEventListener     *DepositEventListener
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

	contractClient.startMiningEventListener = nil
	contractClient.stopMiningEventListener = nil
	contractClient.depositEventListener = nil

	return contractClient, nil
}

func (contractClient *ContractClient) InitStartMiningEventListener() error {
	var err error
	contractClient.startMiningEventListener, err = NewStartMiningEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitStoptMiningEventListener() error {
	var err error
	contractClient.stopMiningEventListener, err = NewStopMiningEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) InitDepositEventListener() error {
	var err error
	contractClient.depositEventListener, err = NewDepositEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) DeInit() {
	if contractClient.startMiningEventListener != nil {
		contractClient.startMiningEventListener.DeInit()
	}
	if contractClient.stopMiningEventListener != nil {
		contractClient.stopMiningEventListener.DeInit()
	}
	if contractClient.depositEventListener != nil {
		contractClient.depositEventListener.DeInit()
	}

	contractClient.chainID = nil
	contractClient.startMiningEventListener = nil
	contractClient.stopMiningEventListener = nil
	contractClient.depositEventListener = nil
	contractClient.ethClient.Close()
}

func (contractClient *ContractClient) Start_StartMiningEventListener(f func(*LitionScClientStartMining)) {
	listener := contractClient.startMiningEventListener
	if listener == nil {
		log.Fatal("Trying to start 'StartMining' listener without previous initialization")
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
			log.Error("Start StartMiningEventListener err: '", retErr, "'. Try to reinit.")
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

func (contractClient *ContractClient) Start_StopMiningEventListener(f func(*LitionScClientStopMining)) {
	listener := contractClient.stopMiningEventListener
	if listener == nil {
		log.Fatal("Trying to start StopMining listener without previous initialization")
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
			log.Error("Start StopMiningEventListener err: '", retErr, "'. Try to reinit.")
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

func (contractClient *ContractClient) Start_DepositEventListener(f func(*LitionScClientDeposit)) {
	listener := contractClient.depositEventListener
	if listener == nil {
		log.Fatal("Trying to start Deposit listener without previous initialization")
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
			log.Error("Start DepositListener err: '", retErr, "'. Try to reinit.")
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

func (contractClient *ContractClient) StartMining(auth *bind.TransactOpts) error {
	tx, err := contractClient.scClient.StartMining(auth, contractClient.chainID)
	if err != nil {
		return err
	}
	log.Info("Transaction 'startMining' TX: ", tx.Hash().String())
	return nil
}

func (contractClient *ContractClient) StopMining(auth *bind.TransactOpts) error {
	tx, err := contractClient.scClient.StopMining(auth, contractClient.chainID)
	if err != nil {
		return err
	}
	log.Info("Transaction 'stopMining' TX: ", tx.Hash().String())
	return nil
}

func (contractClient *ContractClient) AccHasVested(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	hasVested, err := contractClient.scClient.HasVested(&bind.CallOpts{}, contractClient.chainID, userAddress)
	if err != nil {
		return false, err
	}

	return hasVested, nil
}

func (contractClient *ContractClient) AccHasDeposited(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	hasDeposited, err := contractClient.scClient.HasDeposited(&bind.CallOpts{}, contractClient.chainID, userAddress)
	if err != nil {
		return false, err
	}

	return hasDeposited, nil
}

func (contractClient *ContractClient) GetAllowedToTransact() ([]common.Address, error) {
	var accountsWhitelist []common.Address
	zeroCount := big.NewInt(0)

	for batchID := big.NewInt(0); ; batchID.Add(batchID, big.NewInt(1)) {
		accountsList, err := contractClient.scClient.GetAllowedToTransact(&bind.CallOpts{}, contractClient.chainID, batchID)
		if err != nil {
			return nil, err
		}

		log.Info("accountsList.Count: ", accountsList.Count)
		cmpResult := accountsList.Count.Cmp(zeroCount)
		if cmpResult == 0 {
			break
		}

		accountsWhitelist = append(accountsWhitelist, accountsList.Users[0:accountsList.Count.Int64()]...)
	}

	return accountsWhitelist, nil
}

func (contractClient *ContractClient) Notary(auth *bind.TransactOpts, notary_block_no uint32, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32, v []uint8, r [][32]byte, s [][32]byte) error {
	_, err := contractClient.scClient.Notary(auth, contractClient.chainID, notary_block_no, miners, blocks_mined, users, user_gas, largest_tx, v, r, s)
	if err != nil {
		return err
	}
	return nil
}

func (contractClient *ContractClient) GetLastNotary() (int64, error) {
	number, err := contractClient.scClient.GetLastNotary(&bind.CallOpts{}, contractClient.chainID)
	if err != nil {
		return 0, err
	}
	return number.Int64(), nil
}
