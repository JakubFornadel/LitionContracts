package main

import (
	"math/big"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"

	litionContractClient "gitlab.com/lition/lition_contracts/contracts/client"
)

func processStartMining(event *litionContractClient.LitionStartMining) {
	log.Info("processStartMining. Acc: ", event.Miner.String())
}

func processStopMining(event *litionContractClient.LitionStopMining) {
	log.Info("processStopMining. Acc: ", event.Miner.String())
}

func main() {
	infuraURL := "wss://ropsten.infura.io/ws"
	contractAddress := "0xF4f9c1c8D66C8c9c09456BaD6a9890C3caa768c3"
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
	litionContractClient, err := litionContractClient.NewClient(infuraURL, contractAddress, big.NewInt(int64(chainID)))
	if err != nil {
		log.Fatal("Unable to init Lition smart contract client")
	}

	// Init Lition Smartcontract event listeners
	err = litionContractClient.InitStartMiningEventListener()
	if err != nil {
		log.Fatal("Unable to init 'StartMining' event listeners")
	}
	err = litionContractClient.InitStoptMiningEventListener()
	if err != nil {
		log.Fatal("Unable to init 'StopMining' event listeners")
	}

	// Start standalone event listeners
	go litionContractClient.Start_StartMiningEventListener(processStartMining)
	go litionContractClient.Start_StopMiningEventListener(processStopMining)

	if privateKeyStr != "" {
		err = litionContractClient.StartMining(auth)
		if err != nil {
			log.Fatal("Unable to send 'StartMining' tx. Err: ", err)
		}
	}

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	if privateKeyStr != "" {
		err = litionContractClient.StopMining(auth)
		if err != nil {
			log.Fatal("Unable to send 'StopMining' tx. Err: ", err)
		}
	}

	// Deinit lition smart contract cliet
	litionContractClient.DeInit()

	log.Info("Test End")
	os.Exit(0)
}
