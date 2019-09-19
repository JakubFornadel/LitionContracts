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

type WhitelistAccEventListener struct {
	initialized    bool
	listening      bool
	scClient       *LitionScClient
	eventChannel   chan *LitionScClientWhitelistAccount
	eventSubs      event.Subscription
	filterChainID  []*big.Int
	stopChannel    chan struct{}
	stoppedChannel chan struct{}
	mutex          sync.Mutex
}

func NewWhitelistAccEventListener(scClient *LitionScClient, chainId *big.Int) (*WhitelistAccEventListener, error) {
	p := new(WhitelistAccEventListener)

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

func (listener *WhitelistAccEventListener) Init() error {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == true {
		return nil
	}

	var err error
	listener.eventChannel = make(chan *LitionScClientWhitelistAccount)
	listener.eventSubs, err = listener.scClient.WatchWhitelistAccount(
		&bind.WatchOpts{Context: context.Background(), Start: nil},
		listener.eventChannel,
		listener.filterChainID)

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

func (listener *WhitelistAccEventListener) DeInit() {
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

func (listener *WhitelistAccEventListener) ReInit() error {
	listener.DeInit()
	return listener.Init()
}

func (listener *WhitelistAccEventListener) Start(f func(*LitionScClientWhitelistAccount)) error {
	if listener.initialized == false {
		return errors.New("Trying to Start 'WhitelistAccEventListener' without previous initialization")
	}
	if listener.listening == true {
		log.Warning("Trying to Start 'WhitelistAccEventListener', which is already listening.")
		return nil
	}

	listener.stoppedChannel = make(chan struct{})
	listener.listening = true
	log.Info("WhitelistAccEventListener start listening")

	// close the stoppedchan when this func exits
	defer func() {
		close(listener.stoppedChannel)
		listener.listening = false
	}()

	for {
		select {
		case event := <-listener.eventChannel:
			log.Info("New 'WhitelistAccount' event received.")
			f(event)
		case err := <-listener.eventSubs.Err():
			return err
		case <-listener.stopChannel:
			log.Info("Signal to stop WhitelistAccEventListener received.")
			return nil
		}
	}
}

func (listener *WhitelistAccEventListener) Stop() {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.listening == false {
		return
	}

	close(listener.stopChannel)
	// wait for it to have stopped
	<-listener.stoppedChannel
	listener.stopChannel = make(chan struct{})
	log.Info("WhitelistAccEventListener successfully stopped")
}
