package feature

import (
	"strings"
	"sync"
)

// ServiceConfig represents the configuration used to create a new service.
type ServiceConfig struct {
}

// DefaultServiceConfig provides a default configuration to create a new service
// by best effort.
func DefaultServiceConfig() ServiceConfig {
	return ServiceConfig{}
}

// NewService creates a new configured service.
func NewService(config ServiceConfig) (Service, error) {
	newService := &service{
		// Internals.
		bootOnce:     sync.Once{},
		closer:       make(chan struct{}, 1),
		shutdownOnce: sync.Once{},
	}

	return newService, nil
}

type service struct {
	// Internals.
	bootOnce     sync.Once
	closer       chan struct{}
	shutdownOnce sync.Once
}

func (s *service) Boot() {
	s.bootOnce.Do(func() {
		// Service specific boot logic goes here.
	})
}

func (s *service) Scan(config ScanConfig) ([]Feature, error) {
	// Validate.
	err := config.Validate()
	if err != nil {
		return nil, maskAny(err)
	}

	// Prepare sequence combinations.
	var allSeqs []string
	{
		for _, sequence := range config.Sequences() {
			for _, seq := range seqCombinations(sequence, config.Separator(), config.MinLength(), config.MaxLength()) {
				if !containsString(allSeqs, seq) {
					allSeqs = append(allSeqs, seq)
				}
			}
		}
	}

	// Find sequence positions.
	positions := map[string][][]float64{}
	{
		for _, sequence := range config.Sequences() {
			for _, seq := range allSeqs {
				if strings.Contains(sequence, seq) {
					if _, ok := positions[seq]; !ok {
						positions[seq] = [][]float64{}
					}
					positions[seq] = append(positions[seq], seqPositions(sequence, seq)...)
				}
			}
		}
	}

	// Create features for each found sequence.
	var newFeatures []Feature
	{
		for seq, ps := range positions {
			if len(ps) < config.MinCount() {
				continue
			}

			featureConfig := DefaultConfig()
			featureConfig.Positions = ps
			featureConfig.Sequence = seq
			newFeature, err := New(featureConfig)
			if err != nil {
				return nil, maskAny(err)
			}

			newFeatures = append(newFeatures, newFeature)
		}
	}

	return newFeatures, nil
}

func (s *service) ScanConfig() ScanConfig {
	return NewScanConfig()
}

func (s *service) Shutdown() {
	s.shutdownOnce.Do(func() {
		close(s.closer)
	})
}
