// This file was generated by counterfeiter
package mocks

import (
	"sync"

	config_package "github.com/cloudfoundry-incubator/lattice/ltc/config"
	"github.com/cloudfoundry-incubator/lattice/ltc/ssh/command_factory"
)

type FakeSSH struct {
	ConnectAndForwardStub        func(appName string, instanceIndex int, localAddress, remoteAddress string, config *config_package.Config) error
	connectAndForwardMutex       sync.RWMutex
	connectAndForwardArgsForCall []struct {
		appName       string
		instanceIndex int
		localAddress  string
		remoteAddress string
		config        *config_package.Config
	}
	connectAndForwardReturns struct {
		result1 error
	}
	ConnectToShellStub        func(appName string, instanceIndex int, command string, config *config_package.Config) error
	connectToShellMutex       sync.RWMutex
	connectToShellArgsForCall []struct {
		appName       string
		instanceIndex int
		command       string
		config        *config_package.Config
	}
	connectToShellReturns struct {
		result1 error
	}
}

func (fake *FakeSSH) ConnectAndForward(appName string, instanceIndex int, localAddress string, remoteAddress string, config *config_package.Config) error {
	fake.connectAndForwardMutex.Lock()
	fake.connectAndForwardArgsForCall = append(fake.connectAndForwardArgsForCall, struct {
		appName       string
		instanceIndex int
		localAddress  string
		remoteAddress string
		config        *config_package.Config
	}{appName, instanceIndex, localAddress, remoteAddress, config})
	fake.connectAndForwardMutex.Unlock()
	if fake.ConnectAndForwardStub != nil {
		return fake.ConnectAndForwardStub(appName, instanceIndex, localAddress, remoteAddress, config)
	} else {
		return fake.connectAndForwardReturns.result1
	}
}

func (fake *FakeSSH) ConnectAndForwardCallCount() int {
	fake.connectAndForwardMutex.RLock()
	defer fake.connectAndForwardMutex.RUnlock()
	return len(fake.connectAndForwardArgsForCall)
}

func (fake *FakeSSH) ConnectAndForwardArgsForCall(i int) (string, int, string, string, *config_package.Config) {
	fake.connectAndForwardMutex.RLock()
	defer fake.connectAndForwardMutex.RUnlock()
	return fake.connectAndForwardArgsForCall[i].appName, fake.connectAndForwardArgsForCall[i].instanceIndex, fake.connectAndForwardArgsForCall[i].localAddress, fake.connectAndForwardArgsForCall[i].remoteAddress, fake.connectAndForwardArgsForCall[i].config
}

func (fake *FakeSSH) ConnectAndForwardReturns(result1 error) {
	fake.ConnectAndForwardStub = nil
	fake.connectAndForwardReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSH) ConnectToShell(appName string, instanceIndex int, command string, config *config_package.Config) error {
	fake.connectToShellMutex.Lock()
	fake.connectToShellArgsForCall = append(fake.connectToShellArgsForCall, struct {
		appName       string
		instanceIndex int
		command       string
		config        *config_package.Config
	}{appName, instanceIndex, command, config})
	fake.connectToShellMutex.Unlock()
	if fake.ConnectToShellStub != nil {
		return fake.ConnectToShellStub(appName, instanceIndex, command, config)
	} else {
		return fake.connectToShellReturns.result1
	}
}

func (fake *FakeSSH) ConnectToShellCallCount() int {
	fake.connectToShellMutex.RLock()
	defer fake.connectToShellMutex.RUnlock()
	return len(fake.connectToShellArgsForCall)
}

func (fake *FakeSSH) ConnectToShellArgsForCall(i int) (string, int, string, *config_package.Config) {
	fake.connectToShellMutex.RLock()
	defer fake.connectToShellMutex.RUnlock()
	return fake.connectToShellArgsForCall[i].appName, fake.connectToShellArgsForCall[i].instanceIndex, fake.connectToShellArgsForCall[i].command, fake.connectToShellArgsForCall[i].config
}

func (fake *FakeSSH) ConnectToShellReturns(result1 error) {
	fake.ConnectToShellStub = nil
	fake.connectToShellReturns = struct {
		result1 error
	}{result1}
}

var _ command_factory.SSH = new(FakeSSH)