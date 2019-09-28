package litionScClient

import (
	"context"
	"errors"
	"math/big"
	"sync"

	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/event"
)

type AccMiningEventListener struct {
	initialized    bool
	listening      bool
	scClient       *LitionScClient
	eventChannel   chan *LitionScClientAccountMining
	eventSubs      event.Subscription
	filterChainId  []*big.Int
	stopChannel    chan struct{}
	stoppedChannel chan struct{}
	mutex          sync.Mutex
}

func NewAccMiningEventListener(scClient *LitionScClient, chainId *big.Int) (*AccMiningEventListener, error) {
	p := new(AccMiningEventListener)

	p.initialized = false
	p.listening = false
	p.mutex = sync.Mutex{}
	p.scClient = scClient
	p.filterChainId = []*big.Int{chainId}
	err := p.Init()
	if err == nil {
		return p, nil
	}

	return nil, err
}

func (listener *AccMiningEventListener) Init() error {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == true {
		return nil
	}

	var err error
	listener.eventChannel = make(chan *LitionScClientAccountMining)
	listener.eventSubs, err = listener.scClient.WatchAccountMining(&bind.WatchOpts{
		Context: context.Background(), Start: nil},
		listener.eventChannel,
		listener.filterChainId,
		nil)

	if err != nil {
		log.Error(err)
		close(listener.eventChannel)
		return err
	}

	listener.stopChannel = make(chan struct{})
	listener.stoppedChannel = nil // Stopped channel is created and deleted only in start function
	listener.initialized = true

	return nil
}

func (listener *AccMiningEventListener) DeInit() {
	listener.Stop()

	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == false {
		return
	}

	listener.eventSubs.Unsubscribe()
	close(listener.eventChannel)
	close(listener.stopChannel)
	// close(listener.stoppedChannel) // stoppned channel is already closed when stopped listening
	listener.initialized = false
}

func (listener *AccMiningEventListener) ReInit() error {
	listener.DeInit()
	return listener.Init()
}

func (listener *AccMiningEventListener) Start(f func(*LitionScClientAccountMining)) error {
	if listener.initialized == false {
		return errors.New("Trying to Start 'AccMiningEventListener' without previous initialization")
	}
	if listener.listening == true {
		log.Warning("Trying to Start 'AccMiningEventListener', which is already listening.")
		return nil
	}

	listener.stoppedChannel = make(chan struct{})
	listener.listening = true
	log.Info("AccMiningEventListener start listening")

	// close the stoppedchan when this func exits
	defer func() {
		close(listener.stoppedChannel)
		listener.listening = false
	}()

	for {
		select {
		case event := <-listener.eventChannel:
			log.Info("New 'StartMining' event received.")
			f(event)
		case err := <-listener.eventSubs.Err():
			return err
		case <-listener.stopChannel:
			log.Info("Signal to stop AccMiningEventListener received.")
			return nil
		}
	}
}

func (listener *AccMiningEventListener) Stop() {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.listening == false {
		return
	}

	close(listener.stopChannel)
	// wait for it to have stopped
	<-listener.stoppedChannel
	listener.stopChannel = make(chan struct{})
	log.Info("AccMiningEventListener successfully stopped")
}
