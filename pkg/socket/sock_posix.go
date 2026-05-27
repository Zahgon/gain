/*
 * Copyright 2009 The Go Authors. All rights reserved.
 * Copyright (c) 2022 Andy Pan.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

func ipToSockaddrInet4(ip net.IP, port int) (unix.SockaddrInet4, error) {
	_ = "STUB: not implemented"
	return *new(unix.SockaddrInet4), nil
}

func ipToSockaddrInet6(ip net.IP, port int, zone string) (unix.SockaddrInet6, error) {
	_ = "STUB: not implemented"
	// In general, an IP wildcard address, which is either
	// "0.0.0.0" or "::", means the entire IP addressing
	// space. For some historical reason, it is used to
	// specify "any available address" on some operations
	// of IP node.
	//
	// When the IP node supports IPv4-mapped IPv6 address,
	// we allow a listener to listen to the wildcard
	// address of both IP addressing spaces by specifying
	// IPv6 wildcard address.
	return *new(unix.SockaddrInet6), nil
}

// We accept any IPv6 address including IPv4-mapped
// IPv6 address.

func ipToSockaddr(family int, ip net.IP, port int, zone string) (unix.Sockaddr, error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), nil
}
