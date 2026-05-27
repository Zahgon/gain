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

import "sync/atomic"

type connectionManager struct {
	connections      map[int]*connection
	connectionsCount atomic.Int32
	closing          bool
	releaseFdSet     map[int]void
	keyPool          *keyPool
}

func (c *connectionManager) fork(conn *connection, write bool) *connection {
	_ = "STUB: not implemented"
	return nil
}

func (c *connectionManager) getFd(fd int) *connection { _ = "STUB: not implemented"; return nil }

func (c *connectionManager) get(key int, fd int) *connection { _ = "STUB: not implemented"; return nil }

func (c *connectionManager) release(key int) { _ = "STUB: not implemented"; return }

func (c *connectionManager) close(callback func(conn *connection) bool, fdSkipped int) {
	_ = "STUB: not implemented"
	return
}

func (c *connectionManager) allClosed() bool { _ = "STUB: not implemented"; return false }

func (c *connectionManager) activeConnections() int { _ = "STUB: not implemented"; return 0 }

func newConnectionManager() *connectionManager { _ = "STUB: not implemented"; return nil }
