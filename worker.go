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
	"github.com/pawelgaczynski/giouring"
	"github.com/rs/zerolog"
)

type startSignal int

const (
	done startSignal = iota
)

type worker interface {
	loop(socket int) error
	setIndex(index int)
	index() int
	shutdown()
	ringFd() int
	started() bool
	close()
}

type workerStartListener func()

type workerConfig struct {
	cpuAffinity     bool
	processPriority bool
	maxSQEntries    uint
	maxCQEvents     int
	loggerLevel     zerolog.Level
	prettyLogger    bool
}

type workerImpl struct {
	*looper
	*connCloser
	*shutdowner
	idx            int
	logger         zerolog.Logger
	startedChan    chan startSignal
	onCloseHandler func()
}

func (w *workerImpl) ringFd() int { _ = "STUB: not implemented"; return 0 }

func (w *workerImpl) setIndex(index int) { _ = "STUB: not implemented"; return }

func (w *workerImpl) index() int { _ = "STUB: not implemented"; return 0 }

func (w *workerImpl) processEvent(cqe *giouring.CompletionQueueEvent,
	skipErrorChecker func(*giouring.CompletionQueueEvent) bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *workerImpl) logDebug() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func (w *workerImpl) logInfo() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func (w *workerImpl) logWarn() *zerolog.Event { _ = "STUB: not implemented"; return nil }

func (w *workerImpl) logError(err error) *zerolog.Event { _ = "STUB: not implemented"; return nil }

func (w *workerImpl) close() { _ = "STUB: not implemented"; return }

func newWorkerImpl(
	ring *giouring.Ring, config workerConfig, index int, logger zerolog.Logger,
) *workerImpl {
	_ = "STUB: not implemented"
	return nil
}
