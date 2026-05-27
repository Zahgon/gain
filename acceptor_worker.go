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
	"sync/atomic"
	"time"

	"github.com/pawelgaczynski/giouring"
)

type acceptorWorkerConfig struct {
	workerConfig
	tcpKeepAlive time.Duration
}

type acceptorWorker struct {
	*acceptor
	*workerImpl
	config        acceptorWorkerConfig
	ring          *giouring.Ring
	loadBalancer  loadBalancer
	eventHandler  EventHandler
	addConnMethod func(consumer, int32) error

	accepting atomic.Bool
}

func (a *acceptorWorker) addConnViaRing(worker consumer, fileDescriptor int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptorWorker) addConnViaQueue(worker consumer, fileDescriptor int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptorWorker) registerConsumer(consumer consumer) { _ = "STUB: not implemented"; return }

func (a *acceptorWorker) closeRingAndConsumers() { _ = "STUB: not implemented"; return }

func (a *acceptorWorker) loop(fd int) error { _ = "STUB: not implemented"; return nil }

func newAcceptorWorker(
	config acceptorWorkerConfig, loadBalancer loadBalancer, eventHandler EventHandler, features supportedFeatures,
) (*acceptorWorker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
