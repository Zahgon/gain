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

	"github.com/alitto/pond"
	"github.com/pawelgaczynski/gain/pkg/queue"
	"github.com/pawelgaczynski/giouring"
	"github.com/rs/zerolog"
)

type readWriteWorker interface {
	worker
	activeConnections() int
}

type readWriteWorkerConfig struct {
	workerConfig
	asyncHandler  bool
	goroutinePool bool
	sendRecvMsg   bool
}

type readWriteWorkerImpl struct {
	*workerImpl
	*writer
	*reader
	ring              *giouring.Ring
	connectionManager *connectionManager
	asyncOpQueue      queue.LockFreeQueue[*connection]
	pool              *pond.WorkerPool
	eventHandler      EventHandler
	asyncHandler      bool
	goroutinePool     bool
	sendRecvMsg       bool
	localAddr         net.Addr
}

func (w *readWriteWorkerImpl) handleAsyncWritesIfEnabled() { _ = "STUB: not implemented"; return }

func (w *readWriteWorkerImpl) handleAsyncWrites() { _ = "STUB: not implemented"; return }

func (w *readWriteWorkerImpl) work(conn *connection, n int) { _ = "STUB: not implemented"; return }

func (w *readWriteWorkerImpl) doAsyncWork(conn *connection, n int) func() {
	_ = "STUB: not implemented"
	return nil
}

func (w *readWriteWorkerImpl) writeData(conn *connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *readWriteWorkerImpl) onRead(cqe *giouring.CompletionQueueEvent, conn *connection) error {
	_ = "STUB: not implemented"
	// https://manpages.debian.org/unstable/manpages-dev/recv.2.en.html
	// These calls return the number of bytes received, or -1 if an error occurred.
	// In the event of an error, errno is set to indicate the error.
	// When a stream socket peer has performed an orderly shutdown,
	// the return value will be 0 (the traditional "end-of-file" return).
	// Datagram sockets in various domains (e.g., the UNIX and Internet domains) permit zero-length datagrams.
	// When such a datagram is received, the return value is 0.
	// The value 0 may also be returned if the requested number of bytes to receive from a stream socket was 0.
	return nil
}

func (w *readWriteWorkerImpl) addNextRequest(conn *connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *readWriteWorkerImpl) closeConn(conn *connection, syscallClose bool, err error) {
	_ = "STUB: not implemented"
	return
}

func newReadWriteWorkerImpl(ring *giouring.Ring, index int, localAddr net.Addr, eventHandler EventHandler,
	connectionManager *connectionManager, config readWriteWorkerConfig, logger zerolog.Logger,
) *readWriteWorkerImpl {
	_ = "STUB: not implemented"
	return nil
}
