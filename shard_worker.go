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
	"net"
	"sync/atomic"
	"time"

	"github.com/pawelgaczynski/giouring"
)

type shardWorkerConfig struct {
	readWriteWorkerConfig
	tcpKeepAlive time.Duration
}

type shardWorker struct {
	*acceptor
	*readWriteWorkerImpl
	ring               *giouring.Ring
	connectionManager  *connectionManager
	cpuAffinity        bool
	tcpKeepAlive       time.Duration
	connectionProtocol bool
	accepting          atomic.Bool
}

func (w *shardWorker) onAccept(cqe *giouring.CompletionQueueEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *shardWorker) stopAccept() error { _ = "STUB: not implemented"; return nil }

func (w *shardWorker) activeConnections() int { _ = "STUB: not implemented"; return 0 }

func (w *shardWorker) handleConn(conn *connection, cqe *giouring.CompletionQueueEvent) {
	_ = "STUB: not implemented"
	return
}

func (w *shardWorker) initLoop() { _ = "STUB: not implemented"; return }

// 1 is always index for main socket

func (w *shardWorker) loop(fd int) error { _ = "STUB: not implemented"; return nil }

func (w *shardWorker) closeAllConnsAndRings() { _ = "STUB: not implemented"; return }

func newShardWorker(
	index int, localAddr net.Addr, config shardWorkerConfig, eventHandler EventHandler,
) (*shardWorker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
