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

type connCloser struct {
	ring   *giouring.Ring
	logger zerolog.Logger
}

func (c *connCloser) addCloseRequest(fd int) (*giouring.SubmissionQueueEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *connCloser) addCloseConnRequest(conn *connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *connCloser) syscallCloseSocket(fileDescriptor int) error {
	_ = "STUB: not implemented"
	return nil
}

func newConnCloser(ring *giouring.Ring, logger zerolog.Logger) *connCloser {
	_ = "STUB: not implemented"
	return nil
}
