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
package logger_test

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/homeport/watchful/pkg/logger"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "watchful pkg logger suite")
}

type PipelineMock struct {
	callback    func(timesCalled int, messages []ChannelMessage)
	timesCalled int
}

// Write formats all passed byte arrays into one final string
func (p *PipelineMock) Write(messages []ChannelMessage) {
	p.callback(p.timesCalled, messages)
	p.timesCalled = p.timesCalled + 1
}

func (p *PipelineMock) Observer(o PipelineObserver) {
	os.Exit(1)
}

// Location returns the location used to determine the date that is passed into the logs
func (p *PipelineMock) Location() *time.Location {
	return time.Local
}

// DevNullLogger is a simple non recoding logger that discards all content
type DevNullLogger struct {
}

func (d *DevNullLogger) Write(p []byte) (n int, err error) {
	return 0, nil
}

func NewDevNullLogger() *DevNullLogger {
	return &DevNullLogger{}
}
