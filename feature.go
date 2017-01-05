// Package feature implements a service for detecting patterns in information
// sequences. A information sequence can be any string.
package feature

import (
	"sync"
)

// Config represents the configuration used to create a new feature.
type Config struct {
	// Settings.

	// Positions represents the index locations of a detected feature.
	Positions [][]float64
	// Sequence represents the input sequence being detected as feature. That
	// means, the sequence of a feature is the actual conceptual feature itself.
	Sequence string
}

// DefaultConfig provides a default configuration to create a new feature
// by best effort.
func DefaultConfig() Config {
	return Config{
		// Settings.
		Positions: nil,
		Sequence:  "",
	}
}

// New creates a new configured feature.
func New(config Config) (Feature, error) {
	if len(config.Positions) == 0 {
		return nil, maskAnyf(invalidConfigError, "max length must not be 0")
	}
	for _, p := range config.Positions {
		if len(p) != 2 {
			return nil, maskAnyf(invalidConfigError, "positions must have 2 dimensions")
		}
	}
	if config.Sequence == "" {
		return nil, maskAnyf(invalidConfigError, "sequence must not be empty")
	}

	newFeature := &feature{
		// Internals.
		mutex: sync.Mutex{},

		// Settings.
		positions: config.Positions,
		sequence:  config.Sequence,
	}

	return newFeature, nil
}

type feature struct {
	// Internals.
	mutex sync.Mutex

	// Settings.
	positions [][]float64
	sequence  string
}

func (f *feature) AddPosition(position []float64) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(position) != 2 {
		return maskAnyf(invalidExecutionError, "positions must have 2 dimensions")
	}

	f.positions = append(f.positions, position)

	return nil
}

func (f *feature) Count() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.positions)
}

func (f *feature) Positions() [][]float64 {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.positions
}

func (f *feature) Sequence() string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.sequence
}
