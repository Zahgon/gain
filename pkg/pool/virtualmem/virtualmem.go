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
	"os"
)

var pageSize = os.Getpagesize()

func doubleSize(size int) int {
	_ = "STUB: not implemented"
	//nolint:gomnd // skip gomnd linter since this method is self-documented
	return 0
}

type VirtualMem struct {
	Buf  []byte
	Size int
}

func (m *VirtualMem) Zeroes() { _ = "STUB: not implemented"; return }

func NewVirtualMem(size int) *VirtualMem { _ = "STUB: not implemented"; return nil }

func AdjustBufferSize(size int) int { _ = "STUB: not implemented"; return 0 }

func allocateBuffer(size int) []byte { _ = "STUB: not implemented"; return nil }

//nolint:govet // it is perfectly inteded use

func internalMmap(addr uintptr, length, flags int, fd uintptr) (uintptr, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func internalMunmap(addr uintptr, length int) error { _ = "STUB: not implemented"; return nil }
