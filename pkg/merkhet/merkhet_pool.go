// Copyright © 2019 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package merkhet

import (
	"sync"
	"time"
)

// HeartbeatConsumer consumes a single heartbeat
type HeartbeatConsumer func(heartbeat Heartbeat) (err error)

// Pool contains a map of registered merkhets as well as manages them
//
// StartWorker pushes a new Merkhet instance to the Pool.
// This method will also instantly start the worker sub routine
//
// StartHeartbeats starts all the heartbeats registered with the running workers
//
// Size returns the current size of the Pool
//
// ForEach executes the provided function for each Merkhet instance currently managed by the Pool
//
// ForEachHeartbeat executes something for each heartbeat instance
//
// Shutdown shuts the pool and it's heartbeats down
//
// TaskWaitGroup returns the wait group responsible for tasks
type Pool interface {
	StartWorker(m Merkhet, duration time.Duration, heartbeat Consumer)
	StartHeartbeats()
	Size() uint
	BeatingHearts() (heartbeats []Heartbeat)
	ForEach(consumer Consumer) (output ConsumerResult)
	Shutdown()
	TaskWaitGroup() *sync.WaitGroup
}

// SimplePool is a basic implementation of the MerkhetPool interface
type SimplePool struct {
	heartbeats             []Heartbeat
	TaskWaitGroupReference *sync.WaitGroup
}

// StartWorker pushes a new merkhet instance into the Pool and starts the worker
func (s *SimplePool) StartWorker(m Merkhet, duration time.Duration, heartbeat Consumer) {
	worker := NewMerkhetWorker(m)
	go worker.StartWorker() // Start the worker instance in a different go routine

	beat := NewTickedHeartbeat(worker, duration, heartbeat)
	s.heartbeats = append(s.heartbeats, beat)
}

// StartHeartbeats starts all the heartbeats registered with the running workers
func (s *SimplePool) StartHeartbeats() {
	for _, beat := range s.heartbeats {
		if !beat.IsBeating() {
			beat.StartBeating()
		}
	}
}

// Size returns the size of the pool
func (s *SimplePool) Size() uint {
	return uint(len(s.heartbeats))
}

// BeatingHearts returns the currently beating hearts
func (s *SimplePool) BeatingHearts() (heartbeats []Heartbeat) {
	result := make([]Heartbeat, 0)
	for _, h := range s.heartbeats {
		if h.IsBeating() {
			result = append(result, h)
		}
	}
	return result
}

// ForEach executes the provided function for each Merkhet instance currently managed by the pool
func (s *SimplePool) ForEach(c Consumer) (output ConsumerResult) {
	result := NewWaitGroupConsumerResult()

	c.Notify(s.TaskWaitGroup())
	c.Notify(result.WaitGroup)

	for _, beat := range s.heartbeats {
		future := NewFuture()
		result.AddFuture(future)

		c.Consume(beat.Worker().Merkhet(), future)
	}

	return result
}

// Shutdown shuts the pool and it's heartbeats down
func (s *SimplePool) Shutdown() {
	for _, beat := range s.heartbeats {
		beat.StopBeating()
	}
	s.TaskWaitGroup().Wait()
	for _, beat := range s.heartbeats {
		close(beat.Worker().ControllerChannel().C)
	}
}

// TaskWaitGroup returns the wait group responsible for tasks
func (s *SimplePool) TaskWaitGroup() *sync.WaitGroup {
	return s.TaskWaitGroupReference
}

// NewPool returns a fresh empty instance of the go container
func NewPool() *SimplePool {
	return &SimplePool{
		TaskWaitGroupReference: &sync.WaitGroup{},
	}
}
