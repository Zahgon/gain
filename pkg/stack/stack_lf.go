// Copyright 2023 Paweł Gaczyński.
// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed
// by a MIT license that can be found in the LICENSE file.
//
// Original source: https://github.com/golang-design/lockfree/blob/master/stack.go

package stack

import (
	"unsafe"
)

type node[T any] struct {
	value T
	next  unsafe.Pointer
}

type Stack[T any] struct {
	top unsafe.Pointer
	len uint64
}

// NewStack creates a new lock-free queue.
func NewLockFreeStack[T any]() *Stack[T] { _ = "STUB: not implemented"; return nil }

func getZero[T any]() T { _ = "STUB: not implemented"; return *new(T) }

// Pop pops value from the top of the stack.
func (s *Stack[T]) Pop() T { _ = "STUB: not implemented"; return *new(T) }

// Push pushes a value on top of the stack.
func (s *Stack[T]) Push(v T) { _ = "STUB: not implemented"; return }
