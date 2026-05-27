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

package virtualmem

import (
	"github.com/pawelgaczynski/gain/pkg/pool/sync"
)

const (
	maxVMSize = 64 * 1024 * 1024
)

var builtinPool = NewPool()

func Get(size int) *VirtualMem { _ = "STUB: not implemented"; return nil }

// Put returns the virtual memory to the built-in pool.
func Put(mem *VirtualMem) { _ = "STUB: not implemented"; return }

// Get retrieves a byte slice of the length requested by the caller from pool or allocates a new one.
func (p *Pool) Get(size int) *VirtualMem { _ = "STUB: not implemented"; return nil }

// Put returns the virtual memory to the pool.
func (p *Pool) Put(mem *VirtualMem) { _ = "STUB: not implemented"; return }

// this byte slice is not from Pool.Get(), put it into the previous interval of idx

func index(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

type Pool struct {
	pools [32]sync.Pool[*VirtualMem]
}

func NewPool() Pool { _ = "STUB: not implemented"; return *new(Pool) }
