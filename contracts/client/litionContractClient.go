package litionContractClient

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// ContractClient contains variables needed for communication with lition smart contract
type ContractClient struct {
	ethClient                *ethclient.Client
	scAddress                common.Address
	scClient                 *Lition
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

	pScClient, err := NewLition(contractClient.scAddress, contractClient.ethClient)
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

func (contractClient *ContractClient) Start_StartMiningEventListener(f func(*LitionStartMining)) {
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

func (contractClient *ContractClient) Start_StopMiningEventListener(f func(*LitionStopMining)) {
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

func (contractClient *ContractClient) Start_DepositEventListener(f func(*LitionDeposit)) {
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
	log.Info("Transaction 'startMining' TX: ", tx.Hash())
	return nil
}

func (contractClient *ContractClient) StopMining(auth *bind.TransactOpts) error {
	tx, err := contractClient.scClient.StopMining(auth, contractClient.chainID)
	if err != nil {
		return err
	}
	log.Info("Transaction 'stopMining' TX: ", tx.Hash())
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
