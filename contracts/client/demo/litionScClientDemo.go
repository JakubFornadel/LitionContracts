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

func processStartMining(event *litionScClient.LitionScClientStartMining) {
	log.Info("processStartMining. Acc: ", event.Miner.String())
}

func processStopMining(event *litionScClient.LitionScClientStopMining) {
	log.Info("processStopMining. Acc: ", event.Miner.String())
}

func processDeposit(event *litionScClient.LitionScClientDeposit) {
	log.Info("processDeposit. Acc: ", event.Depositer.String(), "Amount: ", event.Deposit)
}

func main() {
	infuraURL := "wss://ropsten.infura.io/ws"
	contractAddress := "0xD754Dc0AF95a4f8615FC990344D9F7327042E658"
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

	// Init Lition Smartcontract event listeners
	err = litionScClient.InitStartMiningEventListener()
	if err != nil {
		log.Fatal("Unable to init 'StartMining' event listeners")
	}
	err = litionScClient.InitStoptMiningEventListener()
	if err != nil {
		log.Fatal("Unable to init 'StopMining' event listeners")
	}
	err = litionScClient.InitDepositEventListener()
	if err != nil {
		log.Fatal("Unable to init 'Deposit' event listeners")
	}

	// Start standalone event listeners
	go litionScClient.Start_StartMiningEventListener(processStartMining)
	go litionScClient.Start_StopMiningEventListener(processStopMining)
	go litionScClient.Start_DepositEventListener(processDeposit)

	if privateKeyStr != "" {
		err = litionScClient.StartMining(auth)
		if err != nil {
			log.Fatal("Unable to send 'StartMining' tx. Err: ", err)
		}
	}

	accountWhitelist, err := litionScClient.GetAllowedToTransact()
	if err != nil {
		log.Fatal("Unable to GetAllowedToTransact")
	}

	log.Info("GetAllowedToTransact: ")
	for _, acc := range accountWhitelist {
		log.Info(acc.String())
	}

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	if privateKeyStr != "" {
		err = litionScClient.StopMining(auth)
		if err != nil {
			log.Fatal("Unable to send 'StopMining' tx. Err: ", err)
		}
	}

	// Deinit lition smart contract cliet
	litionScClient.DeInit()

	log.Info("Demo End")
	os.Exit(0)
}
