package main

import (
	"math/big"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/crypto"

	litionScClient "gitlab.com/lition/lition_contracts/contracts/client"
)

func processAccMining(event *litionScClient.LitionScClientAccountMining) {
	log.Info("processAccMining. Acc: ", event.Account.String(), " miningFlag: ", event.Mining)
}

func processAccWhitelist(event *litionScClient.LitionScClientAccountWhitelisted) {
	log.Info("processAccWhitelist. Acc: ", event.Account.String(), ", Whitelisted: ", event.Whitelisted)
}

func main() {
	infuraURL := "wss://ropsten.infura.io/ws"
	contractAddress := "0xEa1912e78d5aE29cC4a52d8297Cf3aF913aA0187"
	privateKeyStr := ""
	chainID := 0

	var auth *bind.TransactOpts
	if privateKeyStr != "" {
		privateKey, err := crypto.HexToECDSA(privateKeyStr)
		if err != nil {
			log.Fatal(err)
		}
		auth = bind.NewKeyedTransactor(privateKey)
	}

	// Init Lition Smartcontract client
	litionScClient, err := litionScClient.NewClient(infuraURL, contractAddress, big.NewInt(int64(chainID)))
	if err != nil {
		log.Fatal("Unable to init Lition smart contract client")
	}

	chainDynamicDetails, err := litionScClient.GetChainDynamicDetails()
	if err != nil {
		log.Error("err: ", err)
	}

	log.Info("lastNotaryBlock: ", chainDynamicDetails.LastNotaryBlock)
	log.Info("lastNotaryTime: ", chainDynamicDetails.LastNotaryTimestamp)

	// Init Lition Smartcontract event listeners
	err = litionScClient.InitAccMiningEventListener()
	if err != nil {
		log.Fatal("Unable to init 'AccountMining' event listeners")
	}
	err = litionScClient.InitAccWhitelistedEventListener()
	if err != nil {
		log.Fatal("Unable to init 'AccountWHitelist' event listeners")
	}

	// Start standalone event listeners
	go litionScClient.Start_accMiningEventListener(processAccMining)
	go litionScClient.Start_accWhitelistedEventListener(processAccWhitelist)

	if privateKeyStr != "" {
		tx, err := litionScClient.StartMining(auth)
		if err == nil {
			log.Info("'StartMining' tx hash: ", tx.Hash().String())
		} else {
			log.Fatal("Unable to send 'StartMining' tx. Err: ", err)
		}
	}

	transactors, err := litionScClient.GetTransactors()
	if err != nil {
		log.Fatal("Unable to GetTransactors")
	}

	log.Info("GetTransactors: ")
	for _, acc := range transactors {
		log.Info(acc.String())
	}

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	if privateKeyStr != "" {
		tx, err := litionScClient.StopMining(auth)
		if err == nil {
			log.Info("'StopMining' tx hash: ", tx.Hash().String())
		} else {
			log.Fatal("Unable to send 'StopMining' tx. Err: ", err)
		}
	}

	// Deinit lition smart contract cliet
	litionScClient.DeInit()

	log.Info("Demo End")
	os.Exit(0)
}
