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
	"syscall"

	"github.com/pawelgaczynski/giouring"
)

var waitForArray = []uint32{
	1,
	32,
	64,
	96,
	128,
	256,
	384,
	512,
	768,
	1024,
	1536,
	2048,
	3072,
	4096,
	5120,
	6144,
	7168,
	8192,
	10240,
}

type batchSubmitter struct {
	ring            *giouring.Ring
	timeoutTimeSpec syscall.Timespec
	waitForIndex    uint32
	waitFor         uint32
}

func (s *batchSubmitter) submit() error { _ = "STUB: not implemented"; return nil }

func (s *batchSubmitter) advance(n uint32) { _ = "STUB: not implemented"; return }

func newBatchSubmitter(ring *giouring.Ring) *batchSubmitter { _ = "STUB: not implemented"; return nil }
