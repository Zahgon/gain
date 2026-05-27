package gain

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

import (
	"net"
	"sync"

	"github.com/pawelgaczynski/gain/pkg/queue"
	"github.com/pawelgaczynski/giouring"
)

type consumerConfig struct {
	readWriteWorkerConfig
}

type consumer interface {
	readWriteWorker
	addConnToQueue(fd int) error
	setSocketAddr(fd int, addr net.Addr)
}

type consumerWorker struct {
	*readWriteWorkerImpl
	config consumerConfig

	socketAddresses sync.Map
	// used for kernels < 5.18 where OP_MSG_RING is not supported
	connQueue queue.LockFreeQueue[int]
}

func (c *consumerWorker) setSocketAddr(fd int, addr net.Addr) { _ = "STUB: not implemented"; return }

func (c *consumerWorker) addConnToQueue(fd int) error { _ = "STUB: not implemented"; return nil }

func (c *consumerWorker) closeAllConns() { _ = "STUB: not implemented"; return }

func (c *consumerWorker) activeConnections() int { _ = "STUB: not implemented"; return 0 }

func (c *consumerWorker) handleConn(conn *connection, cqe *giouring.CompletionQueueEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *consumerWorker) handleNewConn(fd int) error { _ = "STUB: not implemented"; return nil }

func (c *consumerWorker) getConnsFromQueue() { _ = "STUB: not implemented"; return }

func (c *consumerWorker) handleJobsInQueues() { _ = "STUB: not implemented"; return }

func (c *consumerWorker) loop(_ int) error { _ = "STUB: not implemented"; return nil }

func newConsumerWorker(
	index int, localAddr net.Addr, config consumerConfig, eventHandler EventHandler, features supportedFeatures,
) (*consumerWorker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
