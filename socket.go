// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2019 Andy Pan
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

	"github.com/pawelgaczynski/gain/pkg/socket"
)

func parseProtoAddr(addr string) (string, string) { _ = "STUB: not implemented"; return "", "" }

type listener struct {
	fd               int
	addr             net.Addr
	address, network string
	sockOpts         []socket.Option
}

func (ln *listener) normalize() error { _ = "STUB: not implemented"; return nil }

func initListener(network, addr string, config Config) (*listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
