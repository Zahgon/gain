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
	"syscall"

	"github.com/pawelgaczynski/giouring"
)

type acceptor struct {
	ring              *giouring.Ring
	fd                int
	clientAddr        *syscall.RawSockaddrAny
	clientLenPointer  *uint32
	connectionManager *connectionManager
}

func (a *acceptor) addAcceptRequest() error { _ = "STUB: not implemented"; return nil }

func (a *acceptor) addAcceptConnRequest() error { _ = "STUB: not implemented"; return nil }

func (a *acceptor) lastClientAddr() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func newAcceptor(ring *giouring.Ring, connectionManager *connectionManager) *acceptor {
	_ = "STUB: not implemented"
	return nil
}
