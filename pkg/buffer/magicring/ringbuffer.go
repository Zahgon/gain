// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2019 Chao yuepan, Andy Pan, Allen Xu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE

package magicring

import (
	"io"
	"os"
	"unsafe"

	"github.com/pawelgaczynski/gain/pkg/pool/virtualmem"
)

var (
	DefaultMagicBufferSize = os.Getpagesize()
	MinRead                = 1024
)

// RingBuffer is a circular buffer that implement io.ReaderWriter interface.
type RingBuffer struct {
	vm *virtualmem.VirtualMem

	Size    int
	r       int // next position to read
	w       int // next position to write
	isEmpty bool
}

func (rb *RingBuffer) ReadAddress() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (rb *RingBuffer) WriteAddress() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (rb *RingBuffer) Zeroes() { _ = "STUB: not implemented"; return }

func (rb *RingBuffer) ReleaseBytes() { _ = "STUB: not implemented"; return }

func (rb *RingBuffer) Reset() { _ = "STUB: not implemented"; return }

func (rb *RingBuffer) Read(buffer []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *RingBuffer) Write(buffer []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *RingBuffer) AdvanceWrite(bytesToAdvance int) { _ = "STUB: not implemented"; return }

func (rb *RingBuffer) AdvanceRead(bytesToAdvance int) { _ = "STUB: not implemented"; return }

func (rb *RingBuffer) Grow(newCap int) { _ = "STUB: not implemented"; return }

// Available returns the length of available bytes to write.
func (rb *RingBuffer) Available() int { _ = "STUB: not implemented"; return 0 }

// Buffered returns the length of available bytes to read.
func (rb *RingBuffer) Buffered() int { _ = "STUB: not implemented"; return 0 }

// IsFull tells if this ring-buffer is full.
func (rb *RingBuffer) IsFull() bool { _ = "STUB: not implemented"; return false }

// IsEmpty tells if this ring-buffer is empty.
func (rb *RingBuffer) IsEmpty() bool {
	_ = "STUB: not implemented"

	// Cap returns the size of the underlying buffer.
	return false
}

func (rb *RingBuffer) Cap() int { _ = "STUB: not implemented"; return 0 }

func (rb *RingBuffer) Next(nBytes int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (rb *RingBuffer) Peek(bytesToPeak int) []byte { _ = "STUB: not implemented"; return nil }

// length of ring-buffer

func (rb *RingBuffer) peekAll() []byte { _ = "STUB: not implemented"; return nil }

// Discard skips the next n bytes by advancing the read pointer.
func (rb *RingBuffer) Discard(bytesToDiscard int) int { _ = "STUB: not implemented"; return 0 }

// Bytes returns all available read bytes. It does not move the read pointer and only copy the available data.
func (rb *RingBuffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (rb *RingBuffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

// ReadByte reads and returns the next byte from the input or ErrIsEmpty.
func (rb *RingBuffer) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *RingBuffer) GrowIfUnsufficientFreeSpace() { _ = "STUB: not implemented"; return }

// ReadFrom implements io.ReaderFrom.
func (rb *RingBuffer) ReadFrom(reader io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteTo implements io.WriterTo.
func (rb *RingBuffer) WriteTo(writer io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// New returns a new Buffer whose buffer has the given size.
func NewMagicBuffer(size int) *RingBuffer { _ = "STUB: not implemented"; return nil }
