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
	"time"

	"github.com/rs/zerolog"
)

const (
	defaultPort           = 8080
	defaultMaxCQEvents    = 16384
	defaultMaxSQEntries   = 16384
	defaultRecvBufferSize = 4096
	defaultSendBufferSize = 4096
)

type Option[T any] func(*T)

type ConfigOption Option[Config]

type ServerArchitecture int

const (
	// Reactor design pattern has one input called Acceptor,
	// which demultiplexes the handling of incoming connections to Consumer workers.
	// The load balancing algorithm can be selected via configuration option.
	Reactor ServerArchitecture = iota
	// The Socket Sharding  allow multiple workers to listen on the same address and port combination.
	// In this case the kernel distributes incoming requests across all the sockets.
	SocketSharding
)

// Config is the configuration for the gain engine.
type Config struct {
	// Architecture indicates one of the two available architectures: Reactor and SocketSharding.
	//
	// The Reactor design pattern has one input called Acceptor,
	// which demultiplexes the handling of incoming connections to Consumer workers.
	// The load balancing algorithm can be selected via configuration option.
	//
	// The Socket Sharding allows multiple workers to listen on the same address and port combination.
	// In this case the kernel distributes incoming requests across all the sockets.
	Architecture ServerArchitecture
	// AsyncHandler indicates whether the engine should run the OnRead EventHandler method in a separate goroutines.
	AsyncHandler bool
	// GoroutinePool indicates use of pool of bounded goroutines for OnRead calls.
	// Important: Valid only if AsyncHandler is true
	GoroutinePool bool
	// CPUAffinity determines whether each engine worker is locked to the one CPU.
	CPUAffinity bool
	// ProcessPriority sets the prority of the process to high (-19). Requires root privileges.
	ProcessPriority bool
	// Workers indicates the number of consumers or shard workers. The default is runtime.NumCPU().
	Workers int
	// CBPFilter uses custom BPF filter to improve the performance of the Socket Sharding architecture.
	CBPFilter bool
	// LoadBalancing indicates the load-balancing algorithm to use when assigning a new connection.
	// Important: valid only for Reactor architecture.
	LoadBalancing LoadBalancing
	// SocketRecvBufferSize sets the maximum socket receive buffer in bytes.
	SocketRecvBufferSize int
	// SocketSendBufferSize sets the maximum socket send buffer in bytes.
	SocketSendBufferSize int
	// TCPKeepAlive sets the TCP keep-alive for the socket.
	TCPKeepAlive time.Duration
	// LoggerLevel indicates the logging level.
	LoggerLevel zerolog.Level
	// PrettyLogger sets the pretty-printing zerolog mode.
	// Important: it is inefficient so should be used only for debugging.
	PrettyLogger bool

	// ==============================
	// io_uring related options
	// ==============================
	// MaxSQEntries sets the maximum number of SQEs that can be submitted in one batch.
	// If the number of SQEs exceeds this value, the io_uring will return a SQE overflow error.
	MaxSQEntries uint
	// MaxCQEvents sets the maximum number of CQEs that can be retrieved in one batch.
	MaxCQEvents uint
}

// WithArchitecture sets the architecture of gain engine.
func WithArchitecture(architecture ServerArchitecture) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithAsyncHandler sets the asynchronous mode for the OnRead callback.
func WithAsyncHandler(asyncHandler bool) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithGoroutinePool sets the goroutine pool for asynchronous handler.
func WithGoroutinePool(goroutinePool bool) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithCPUAffinity sets the CPU affinity option.
func WithCPUAffinity(cpuAffinity bool) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithProcessPriority sets the high process priority. Note: requires root privileges.
func WithProcessPriority(processPriority bool) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithWorkers sets the number of workers.
func WithWorkers(workers int) ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// WithCBPF sets the CBPF filter for the gain engine.
func WithCBPF(cbpf bool) ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// WithLoadBalancing sets the load balancing algorithm.
func WithLoadBalancing(loadBalancing LoadBalancing) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithSocketRecvBufferSize sets the maximum socket receive buffer in bytes.
func WithSocketRecvBufferSize(size int) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithSocketSendBufferSize sets the maximum socket send buffer in bytes.
func WithSocketSendBufferSize(size int) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithTCPKeepAlive sets the TCP keep-alive for the socket.
func WithTCPKeepAlive(tcpKeepAlive time.Duration) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithLoggerLevel sets the logging level.
func WithLoggerLevel(loggerLevel zerolog.Level) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithPrettyLogger sets the pretty-printing zerolog mode.
func WithPrettyLogger(prettyLogger bool) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithMaxSQEntries sets the maximum number of entries in the submission queue.
func WithMaxSQEntries(maxSQEntries uint) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithMaxCQEvents sets the maximum number of entries in the completion queue.
func WithMaxCQEvents(maxCQEvents uint) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

func NewConfig(opts ...ConfigOption) Config { _ = "STUB: not implemented"; return *new(Config) }
