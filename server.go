// Copyright (c) 2023 Paweł Gaczyński
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gain

import (
	"sync"
	"sync/atomic"

	"github.com/rs/zerolog"
)

const (
	inactive uint32 = iota
	starting
	running
	closing
)

type Server interface {
	// Start starts a server that will listen on the specified address and
	// waits for SIGTERM or SIGINT signals to close the server.
	StartAsMainProcess(address string) error
	// Start starts a server that will listen on the specified address.
	Start(address string) error
	// Shutdown closes all connections and shuts down server. It's blocking until the server shuts down.
	Shutdown()
	// AsyncShutdown closes all connections and shuts down server in asynchronous manner.
	// It does not wait for the server shutdown to complete.
	AsyncShutdown()
	// ActiveConnections returns the number of active connections.
	ActiveConnections() int
	// IsRunning returns true if server is running and handling requests.
	IsRunning() bool
}

type engine struct {
	config  Config
	network string
	address string

	closer        func() error
	closeChan     chan closeSignal
	closeBackChan chan bool
	state         atomic.Uint32
	logger        zerolog.Logger

	eventHandler EventHandler

	readWriteWorkers sync.Map
}

type closeSignal int

const (
	system closeSignal = iota
	user
)

func (e *engine) startConsumers(startedWg, doneWg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func (e *engine) handleWorkerStop(
	index int, numberOfWorkers *int32, err error, startedWg *sync.WaitGroup, worker readWriteWorkerImpl,
) {
	_ = "STUB: not implemented"
	return
}

func (e *engine) startReactor(listener *listener, features supportedFeatures) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) startSocketSharding(listeners []*listener, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) start(mainProcess bool, address string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) StartAsMainProcess(address string) error { _ = "STUB: not implemented"; return nil }

func (e *engine) Start(address string) error { _ = "STUB: not implemented"; return nil }

func (e *engine) Shutdown() { _ = "STUB: not implemented"; return }

func (e *engine) AsyncShutdown() { _ = "STUB: not implemented"; return }

func (e *engine) ActiveConnections() int { _ = "STUB: not implemented"; return 0 }

func (e *engine) IsRunning() bool { _ = "STUB: not implemented"; return false }

func NewServer(eventHandler EventHandler, config Config) Server {
	_ = "STUB: not implemented"
	return *new(Server)
}
