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

// Result defines a statistic object that contains the data of a merkhet run
//
// TotalRuns returns the total amount of runs the merkhet instance ran
// at the time this result instance was created
//
// FailedRuns returns the total amount of faild runs the merkhet instance that build
// this result produced
//
// Valid returns if the result was marked valid by the Merkhet instance that build it
type Result interface {
	TotalRuns() uint
	FailedRuns() uint
	Valid() bool
}

// SimpleResult is a small implementation of the Result interface
type SimpleResult struct {
	totalRuns  uint
	totalFails uint
	valid      bool
}

// TotalRuns returns the total amount of runs the merkhet instance ran
// at the time this result instance was created
func (s *SimpleResult) TotalRuns() uint {
	return s.totalRuns
}

// FailedRuns returns the total amount of faild runs the merkhet instance that build
// this result produced
func (s *SimpleResult) FailedRuns() uint {
	return s.totalFails
}

// Valid returns if the result was marked valid by the Merkhet instance that build it
func (s *SimpleResult) Valid() bool {
	return s.valid
}

// NewMerkhetResult creates a new instance of the MerkhetResult interface
func NewMerkhetResult(totalRuns uint, failedRuns uint, valid bool) *SimpleResult {
	return &SimpleResult{
		totalRuns:  totalRuns,
		totalFails: failedRuns,
		valid:      valid,
	}
}
