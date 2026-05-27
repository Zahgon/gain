// Copyright (c) 2023 Paweł Gaczyński
// Copyright (c) 2019 Andy Pan
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
	"syscall"
)

var ipv4InIPv6Prefix = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff}

const (
	netIPBytesSize        = 16
	sockAddrPortBitOffset = 8
	decimalBase           = 10
	int32size             = 32
)

func RawAnyToSockaddrInet4(rsa *syscall.RawSockaddrAny) (*syscall.SockaddrInet4, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SockaddrToTCPOrUnixAddr converts a Sockaddr to a net.TCPAddr or net.UnixAddr.
// Returns nil if conversion fails.
func SockaddrToTCPOrUnixAddr(sysSockAddr syscall.Sockaddr) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

// SockaddrToUDPAddr converts a Sockaddr to a net.UDPAddr
// Returns nil if conversion fails.
func SockaddrToUDPAddr(sysSockAddr syscall.Sockaddr) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

// sockaddrInet4ToIPAndZone converts a SockaddrInet4 to a net.IP.
// It returns nil if conversion fails.
func sockaddrInet4ToIP(sa *syscall.SockaddrInet4) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

// ipv4InIPv6Prefix

// sockaddrInet6ToIPAndZone converts a SockaddrInet6 to a net.IP with IPv6 Zone.
// It returns nil if conversion fails.
func sockaddrInet6ToIPAndZone(sa *syscall.SockaddrInet6) (net.IP, string) {
	_ = "STUB: not implemented"
	return *new(net.IP), ""
}

// ip6ZoneToString converts an IP6 Zone unix int to a net string
// returns "" if zone is 0.
func ip6ZoneToString(zone int) string { _ = "STUB: not implemented"; return "" }

// BytesToString converts byte slice to a string without memory allocation.
//
// Note it may break if the implementation of string or slice header changes in the future go versions.
func BytesToString(b []byte) string {
	_ = "STUB: not implemented"
	/* #nosec G103 */ return ""
}

// Convert int to decimal string.
func int2decimal(value uint) string { _ = "STUB: not implemented"; return "" }

// Assemble decimal in reverse order.
