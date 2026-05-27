// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2019 Andy Pan
// Copyright (c) 2016 Aliaksandr Valialkin, VertaMedia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Use of this source code is governed by a MIT license that can be found
// at https://github.com/valyala/bytebufferpool/blob/master/LICENSE
package ringbuffer

import (
	"github.com/pawelgaczynski/gain/pkg/buffer/magicring"
	"github.com/pawelgaczynski/gain/pkg/pool/sync"
)

const (
	minBitSize = 12 // 2**6=64 is a CPU cache line size
	steps      = 15

	minSize = 1 << minBitSize

	calibrateCallsThreshold = 42000
	maxPercentile           = 0.95
)

// RingBuffer is the alias of ring.Buffer.
type RingBuffer = magicring.RingBuffer

// Pool represents ring-buffer pool.
//
// Distinct pools may be used for distinct types of byte buffers.
// Properly determined byte buffer types with their own pools may help to reduce
// memory waste.
type Pool struct {
	calls       [steps]uint64
	calibrating uint64

	defaultSize uint64
	maxSize     uint64

	pool sync.Pool[*RingBuffer]
}

var builtinPool = NewRingBufferPool()

// Get returns an empty byte buffer from the pool.
//
// Got byte buffer may be returned to the pool via Put call.
// This reduces the number of memory allocations required for byte buffer
// management.
func Get() *RingBuffer { _ = "STUB: not implemented"; return nil }

// Get returns new byte buffer with zero length.
//
// The byte buffer may be returned to the pool via Put after the use
// in order to minimize GC overhead.
func (p *Pool) Get() *RingBuffer { _ = "STUB: not implemented"; return nil }

// Put returns byte buffer to the pool.
//
// RingBuffer mustn't be touched after returning it to the pool,
// otherwise, data races will occur.
func Put(b *RingBuffer) { _ = "STUB: not implemented"; return }

// Put releases byte buffer obtained via Get to the pool.
//
// The buffer mustn't be accessed after returning to the pool.
func (p *Pool) Put(buffer *RingBuffer) { _ = "STUB: not implemented"; return }

func (p *Pool) calibrate() { _ = "STUB: not implemented"; return }

func NewRingBufferPool() Pool { _ = "STUB: not implemented"; return *new(Pool) }

type callSize struct {
	calls uint64
	size  uint64
}

type callSizes []callSize

func (ci callSizes) Len() int { _ = "STUB: not implemented"; return 0 }

func (ci callSizes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ci callSizes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func indexRingBufferPool(n int) int { _ = "STUB: not implemented"; return 0 }
