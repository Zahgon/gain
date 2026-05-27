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
	"github.com/pawelgaczynski/gain/pkg/pool/sync"
)

var builtinConnectionPool = newConnectionPool()

func getConnection() *connection { _ = "STUB: not implemented"; return nil }

func putConnection(conn *connection) { _ = "STUB: not implemented"; return }

type connectionPool struct {
	internalPool sync.Pool[*connection]
}

func (c *connectionPool) Put(conn *connection) { _ = "STUB: not implemented"; return }

func (c *connectionPool) Get() *connection { _ = "STUB: not implemented"; return nil }

func newConnectionPool() *connectionPool { _ = "STUB: not implemented"; return nil }
