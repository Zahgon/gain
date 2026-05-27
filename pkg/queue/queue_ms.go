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

package queue

import (
	"unsafe"
)

type msQueue[T any] struct {
	head      unsafe.Pointer
	tail      unsafe.Pointer
	queueSize int32
}

type node[T any] struct {
	value *T
	next  unsafe.Pointer
}

func NewQueue[T any]() LockFreeQueue[T] { _ = "STUB: not implemented"; return nil }

func NewIntQueue() LockFreeQueue[int] { _ = "STUB: not implemented"; return nil }

func (q *msQueue[T]) Enqueue(value T) { _ = "STUB: not implemented"; return }

func (q *msQueue[T]) Dequeue() T { _ = "STUB: not implemented"; return *new(T) }

func (q *msQueue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *msQueue[T]) Size() int32 { _ = "STUB: not implemented"; return 0 }

func load[T any](p *unsafe.Pointer) *node[T] { _ = "STUB: not implemented"; return nil }

func cas[T any](p *unsafe.Pointer, oldNode, newNode *node[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func getZero[T any]() T { _ = "STUB: not implemented"; return *new(T) }
