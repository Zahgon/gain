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
	"io"
	"net"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/pawelgaczynski/gain/pkg/buffer/magicring"
)

type connectionState int

const (
	connInvalid connectionState = iota
	connAccept
	connRead
	connWrite
	connClose
)

func (s connectionState) String() string { _ = "STUB: not implemented"; return "" }

const (
	kernelSpace = iota
	userSpace
)

func connModeString(m uint32) string { _ = "STUB: not implemented"; return "" }

const (
	msgControlBufferSize = 64
)

const (
	noOp = iota
	readOp
	writeOp
	closeOp
)

const (
	tcp = iota
	udp
)

type connection struct {
	fd      int
	key     int
	network uint32

	inboundBuffer  *magicring.RingBuffer
	outboundBuffer *magicring.RingBuffer
	state          connectionState
	mode           atomic.Uint32
	closed         atomic.Bool

	msgHdr      *syscall.Msghdr
	rawSockaddr *syscall.RawSockaddrAny

	localAddr  net.Addr
	remoteAddr net.Addr

	ctx interface{}

	nextAsyncOp int
}

func (c *connection) outboundReadAddress() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (c *connection) inboundWriteAddress() unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (c *connection) setKernelSpace() { _ = "STUB: not implemented"; return }

func (c *connection) setUserSpace() { _ = "STUB: not implemented"; return }

func (c *connection) Context() interface{} { _ = "STUB: not implemented"; return nil }

func (c *connection) SetContext(ctx interface{}) { _ = "STUB: not implemented"; return }

func (c *connection) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *connection) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *connection) Fd() int { _ = "STUB: not implemented"; return 0 }

func (c *connection) userOpAllowed(name string) error { _ = "STUB: not implemented"; return nil }

func (c *connection) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck

func (c *connection) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck

func (c *connection) SetLinger(sec int) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck

func (c *connection) SetNoDelay(noDelay bool) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck

func (c *connection) SetKeepAlivePeriod(period time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:wrapcheck

func (c *connection) onKernelRead(n int) { _ = "STUB: not implemented"; return }

func (c *connection) onKernelWrite(n int) { _ = "STUB: not implemented"; return }

func (c *connection) isClosed() bool { _ = "STUB: not implemented"; return false }

func (c *connection) Close() error { _ = "STUB: not implemented"; return nil }

func (c *connection) Next(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:wrapcheck

func (c *connection) Discard(n int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *connection) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *connection) ReadFrom(reader io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:wrapcheck

func (c *connection) WriteTo(writer io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:wrapcheck

func (c *connection) Read(buffer []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func (c *connection) Write(buffer []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck

func (c *connection) OutboundBuffered() int { _ = "STUB: not implemented"; return 0 }

func (c *connection) InboundBuffered() int { _ = "STUB: not implemented"; return 0 }

func (c *connection) setMsgHeaderWrite() { _ = "STUB: not implemented"; return }

func (c *connection) initMsgHeader() { _ = "STUB: not implemented"; return }

func (c *connection) fork(newConn *connection, key int, write bool) *connection {
	_ = "STUB: not implemented"
	return nil
}

func newConnection() *connection { _ = "STUB: not implemented"; return nil }
