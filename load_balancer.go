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
	"net"
)

type LoadBalancing int

const (
	// RoundRobin forwards accepted connections to dedicated workers sequentially.
	RoundRobin LoadBalancing = iota
	// LeastConnections forwards the next accepted connection to the worker with the least number of active connections.
	LeastConnections
	// SourceAddrHash forwards the next accepted connection to the worker by hashing the remote peer address.
	SourceIPHash
)

type loadBalancer interface {
	register(consumer)
	next(net.Addr) consumer
	forEach(func(consumer) error) error
}

type genericLoadBalancer struct {
	workers []consumer
	size    int
}

func (b *genericLoadBalancer) register(worker consumer) { _ = "STUB: not implemented"; return }

type roundRobinLoadBalancer struct {
	*genericLoadBalancer
	nextWorkerIndex int
}

func (b *roundRobinLoadBalancer) next(_ net.Addr) consumer {
	_ = "STUB: not implemented"
	return *new(consumer)
}

func (b *roundRobinLoadBalancer) forEach(callback func(consumer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func newRoundRobinLoadBalancer() loadBalancer { _ = "STUB: not implemented"; return *new(loadBalancer) }

type leastConnectionsLoadBalancer struct {
	*genericLoadBalancer
}

func (b *leastConnectionsLoadBalancer) next(_ net.Addr) consumer {
	_ = "STUB: not implemented"
	return *new(consumer)
}

func (b *leastConnectionsLoadBalancer) forEach(callback func(consumer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func newLeastConnectionsLoadBalancer() loadBalancer {
	_ = "STUB: not implemented"
	return *new(loadBalancer)
}

type sourceIPHashLoadBalancer struct {
	*genericLoadBalancer
}

func (b *sourceIPHashLoadBalancer) hash(s string) int { _ = "STUB: not implemented"; return 0 }

func (b *sourceIPHashLoadBalancer) next(addr net.Addr) consumer {
	_ = "STUB: not implemented"
	return *new(consumer)
}

func (b *sourceIPHashLoadBalancer) forEach(callback func(consumer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func newSourceIPHashLoadBalancer() loadBalancer {
	_ = "STUB: not implemented"
	return *new(loadBalancer)
}

func createLoadBalancer(loadBalancing LoadBalancing) (loadBalancer, error) {
	_ = "STUB: not implemented"
	return *new(loadBalancer), nil
}
