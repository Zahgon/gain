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
	"golang.org/x/net/bpf"
)

const (
	skfAdOffPlusKSkfAdCPU = 4294963236
	cpuIDSize             = 4
)

type filter []bpf.Instruction

func (f filter) applyTo(fileDescriptor int) error { _ = "STUB: not implemented"; return nil }

// /* A = raw_smp_processor_id(). */
// { BPF_LD  | BPF_W | BPF_ABS, 0, 0, SKF_AD_OFF + SKF_AD_CPU },
// /* Adjust the CPUID to socket group size. */
// { BPF_ALU | BPF_MOD | BPF_K, 0, 0, sock_count },
// /* Return A. */
// { BPF_RET | BPF_A, 0, 0, 0 },
//
//nolint:godot
func newFilter(cpus uint32) filter { _ = "STUB: not implemented"; return *new(filter) }

// return 0xffff bytes (or less) from packet.
