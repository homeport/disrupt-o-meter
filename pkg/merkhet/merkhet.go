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
	"github.com/homeport/disrupt-o-meter/pkg/logger"
)

// Merkhet deinfes a runnable mesaurement task that can be executed during the Cloud Foundery maintanance
//
// Install installs the merkhet instance. This method call will be used to setup necessary dependencies of the merkhet
//
// PostConnect is called after DOM successfully authenticated against the cloud foundery instance
//
// PostConnect is called after DOM successfully authenticated against the cloud foundery instance
//
// BuildResult creates a new Result instance containing the
//
// Configuration returns the configuration instances the Merkhet is depending on
//
// Configuration returns the configuration instances the Merkhet is depending on
type Merkhet interface {
	Install()
	PostConnect()
	Execute()
	BuildResult() Result
	Configuration() Configuration
	Logger() logger.Logger
}

// Configuration contains the passed configuration values for a Merkhet instance
//
// Name returns the name provided in the configuration.
//
// ValidRun returns wether the provided total relative to the fails is still concidered a viable run
// The behaviour of this method is heavily reliant on the implementation
type Configuration interface {
	Name() string
	ValidRun(totalRuns uint, failedRuns uint) bool
}

// namedConfiguration is a simple structure containing the name of a MerhetConfiguration
type namedConfiguration struct {
	name string
}

// Name returns the name stored in the named configuration
func (n *namedConfiguration) Name() string {
	return n.name
}

// PercentageConfiguration is a configuration implementation based on a percentage threshhold
type PercentageConfiguration struct {
	namedConfiguration   *namedConfiguration
	percentageThreshhold float64
}

// Name returns the name stored in the configuration delegate
func (p *PercentageConfiguration) Name() string {
	return p.namedConfiguration.Name()
}

// ValidRun returns if the failed runs comapred to the total runs are below the provided percentage threshhold
func (p *PercentageConfiguration) ValidRun(totalRuns uint, failedRuns uint) bool {
	return (float64(failedRuns) / float64(totalRuns)) <= p.percentageThreshhold
}

// FlatConfiguration is an implementation of the Configuration interface that is based on a flat amount of failed runs
// to calucalte viability
type FlatConfiguration struct {
	namedConfiguration *namedConfiguration
	flatThreshhold     uint
}

// Name returns the name stored in the configuration delegate
func (f *FlatConfiguration) Name() string {
	return f.namedConfiguration.Name()
}

// ValidRun returns if the failed runs comapred to the total runs are below the provided percentage threshhold
func (f *FlatConfiguration) ValidRun(totalRuns uint, failedRuns uint) bool {
	return failedRuns <= f.flatThreshhold
}

// NewPercentageConfiguration creates a new configuration intaces that uses a percentage threshold
func NewPercentageConfiguration(name string, percentageTreshhold float64) *PercentageConfiguration {
	return &PercentageConfiguration{
		namedConfiguration: &namedConfiguration{
			name: name,
		},
		percentageThreshhold: percentageTreshhold,
	}
}

// NewFlatConfiguration creates a new configuration intaces that uses a flat threshold
func NewFlatConfiguration(name string, flatThreshhold uint) *FlatConfiguration {
	return &FlatConfiguration{
		namedConfiguration: &namedConfiguration{
			name: name,
		},
		flatThreshhold: flatThreshhold,
	}
}
