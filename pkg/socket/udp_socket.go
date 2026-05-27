// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2020 Andy Pan
// Copyright (c) 2017 Max Riveiro
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

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

// GetUDPSockAddr the structured addresses based on the protocol and raw address.
//
//nolint:dupl // dupl marks this incorrectly as duplicate of GetTCPSockAddr
func GetUDPSockAddr(proto, addr string) (unix.Sockaddr, int, *net.UDPAddr, bool, error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), 0, nil, false, nil
}

func determineUDPProto(proto string, addr *net.UDPAddr) (string, error) {
	_ = "STUB: not implemented"
	// If the protocol is set to "udp", we try to determine the actual protocol
	// version from the size of the resolved IP address. Otherwise, we simple use
	// the protocol given to us by the caller.
	return "", nil
}

// udpSocket creates an endpoint for communication and returns a file descriptor that refers to that endpoint.
func udpSocket(proto, addr string, connect bool, sockOpts ...Option) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

// ignore EINPROGRESS for non-blocking socket connect, should be processed by caller

// Allow broadcast.
