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

type VestInChainEventListener struct {
	initialized    bool
	listening      bool
	scClient       *LitionScClient
	eventChannel   chan *LitionScClientVestInChain
	eventSubs      event.Subscription
	filterChainID  []*big.Int
	stopChannel    chan struct{}
	stoppedChannel chan struct{}
	mutex          sync.Mutex
}

func NewVestInChainEventListener(scClient *LitionScClient, chainId *big.Int) (*VestInChainEventListener, error) {
	p := new(VestInChainEventListener)

	p.initialized = false
	p.listening = false
	p.mutex = sync.Mutex{}
	p.scClient = scClient
	p.filterChainID = []*big.Int{chainId}
	err := p.Init()
	if err == nil {
		return p, nil
	}

	return nil, err
}

func (listener *VestInChainEventListener) Init() error {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == true {
		return nil
	}

	var err error
	listener.eventChannel = make(chan *LitionScClientVestInChain)
	listener.eventSubs, err = listener.scClient.WatchVestInChain(
		&bind.WatchOpts{Context: context.Background(), Start: nil},
		listener.eventChannel,
		listener.filterChainID,
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

func (listener *VestInChainEventListener) DeInit() {
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

func (listener *VestInChainEventListener) ReInit() error {
	listener.DeInit()
	return listener.Init()
}

func (listener *VestInChainEventListener) Start(f func(*LitionScClientVestInChain)) error {
	if listener.initialized == false {
		return errors.New("Trying to Start 'VestInChainEventListener' without previous initialization")
	}
	if listener.listening == true {
		log.Warning("Trying to Start 'VestInChainEventListener', which is already listening.")
		return nil
	}

	listener.stoppedChannel = make(chan struct{})
	listener.listening = true
	log.Info("VestInChainEventListener start listening")

	// close the stoppedchan when this func exits
	defer func() {
		close(listener.stoppedChannel)
		listener.listening = false
	}()

	for {
		select {
		case event := <-listener.eventChannel:
			log.Info("New 'VestInChain' event received.")
			f(event)
		case err := <-listener.eventSubs.Err():
			return err
		case <-listener.stopChannel:
			log.Info("Signal to stop VestInChainEventListener received.")
			return nil
		}
	}
}

func (listener *VestInChainEventListener) Stop() {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.listening == false {
		return
	}

	close(listener.stopChannel)
	// wait for it to have stopped
	<-listener.stoppedChannel
	listener.stopChannel = make(chan struct{})
	log.Info("VestInChainEventListener successfully stopped")
}
