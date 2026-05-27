// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2021 Andy Pan
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

package byteslice

import (
	"unsafe"

	"github.com/pawelgaczynski/gain/pkg/pool/sync"
)

var builtinPool = NewByteSlicePool()

// Pool consists of 32 sync.Pool, representing byte slices of length from 0 to 32 in powers of 2.
type Pool struct {
	pools [32]sync.Pool[unsafe.Pointer]
}

// Get returns a byte slice with given length from the built-in pool.
func Get(size int) []byte { _ = "STUB: not implemented"; return nil }

// Put returns the byte slice to the built-in pool.
func Put(buf []byte) { _ = "STUB: not implemented"; return }

// Get retrieves a byte slice of the length requested by the caller from pool or allocates a new one.
func (p *Pool) Get(size int) []byte { _ = "STUB: not implemented"; return nil }

// Put returns the byte slice to the pool.
func (p *Pool) Put(buf []byte) { _ = "STUB: not implemented"; return }

// this byte slice is not from Pool.Get(), put it into the previous interval of idx

// array pointer

func index(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func NewByteSlicePool() Pool { _ = "STUB: not implemented"; return *new(Pool) }
