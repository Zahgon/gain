// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2021 Andy Pan
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
)

// SetNoDelay controls whether the operating system should delay
// packet transmission in hopes of sending fewer packets (Nagle's algorithm).
//
// The default is true (no delay), meaning that data is
// sent as soon as possible after a Write.
func SetNoDelay(fd, noDelay int) error { _ = "STUB: not implemented"; return nil }

// SetRecvBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func SetRecvBuffer(fd, size int) error { _ = "STUB: not implemented"; return nil }

// SetSendBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func SetSendBuffer(fd, size int) error { _ = "STUB: not implemented"; return nil }

// SetReuseport enables SO_REUSEPORT option on socket.
func SetReuseport(fd, reusePort int) error { _ = "STUB: not implemented"; return nil }

// SetReuseAddr enables SO_REUSEADDR option on socket.
func SetReuseAddr(fd, reuseAddr int) error { _ = "STUB: not implemented"; return nil }

// SetIPv6Only restricts a IPv6 socket to only process IPv6 requests or both IPv4 and IPv6 requests.
func SetIPv6Only(fd, ipv6only int) error { _ = "STUB: not implemented"; return nil }

// SetQuickAck controls quickack mode on socket.
// If quickack mode, acks are sent immediately, rather than delayed if needed in accordance to normal TCP operation.
func SetQuickAck(fd, enabled int) error { _ = "STUB: not implemented"; return nil }

// SetFastOpen enables TCP_FASTOPEN on socket which allow to send and accept data in the opening SYN packet.
// https://sysctl-explorer.net/net/ipv4/tcp_fastopen/
func SetFastOpen(fd, enabled int) error { _ = "STUB: not implemented"; return nil }

// SetLinger sets the behavior of Close on a connection which still
// has data waiting to be sent or to be acknowledged.
//
// If sec < 0 (the default), the operating system finishes sending the
// data in the background.
//
// If sec == 0, the operating system discards any unsent or
// unacknowledged data.
//
// If sec > 0, the data is sent in the background as with sec < 0. On
// some operating systems after sec seconds have elapsed any remaining
// unsent data may be discarded.
func SetLinger(fd, sec int) error { _ = "STUB: not implemented"; return nil }

// SetMulticastMembership returns with a socket option function based on the IP
// version. Returns nil when multicast membership cannot be applied.
func SetMulticastMembership(proto string, udpAddr *net.UDPAddr) func(int, int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetIPv4MulticastMembership joins fd to the specified multicast IPv4 address.
// ifIndex is the index of the interface where the multicast datagrams will be
// received. If ifIndex is 0 then the operating system will choose the default,
// it is usually needed when the host has multiple network interfaces configured.
func SetIPv4MulticastMembership(fd int, mcast net.IP, ifIndex int) error {
	_ = "STUB: not implemented"
	// Multicast interfaces are selected by IP address on IPv4 (and by index on IPv6)
	return nil
}

// SetIPv6MulticastMembership joins fd to the specified multicast IPv6 address.
// ifIndex is the index of the interface where the multicast datagrams will be
// received. If ifIndex is 0 then the operating system will choose the default,
// it is usually needed when the host has multiple network interfaces configured.
func SetIPv6MulticastMembership(fd int, mcast net.IP, ifIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

// interfaceFirstIPv4Addr returns the first IPv4 address of the interface.
func interfaceFirstIPv4Addr(ifIndex int) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}
