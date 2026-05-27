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
	"sync"
	"sync/atomic"
)

type shutdowner struct {
	shutdowning atomic.Bool
	inProgress  atomic.Bool
	wg          sync.WaitGroup
}

func (s *shutdowner) markShutdownInProgress() { _ = "STUB: not implemented"; return }

func (s *shutdowner) needToShutdown() bool { _ = "STUB: not implemented"; return false }

func (s *shutdowner) notifyFinish() { _ = "STUB: not implemented"; return }

func (s *shutdowner) shutdown() { _ = "STUB: not implemented"; return }

func newShutdowner() *shutdowner { _ = "STUB: not implemented"; return nil }
