package feature

// Feature represents a charactistic within a sequence. During pattern
// recognition it is tried to detect features. Their distributions describe
// location patterns within space.
type Feature interface {
	// AddPosition provides a way to add more positions to the initialized
	// feature. Note positions are vectors in distribution terms.
	AddPosition(position []float64) error
	// Count returns the number of occurrences within analysed sequences. That is,
	// how often this feature was found. Technically spoken,
	// len(Feature.Positions).
	Count() int
	// Positions returns the feature's configured positions.
	Positions() [][]float64
	// Sequence returns the sequence that represents this feature. This is the
	// sub-sequence, the charactistic detected within analysed sequences.
	Sequence() string
}

// ScanConfig represents the configuration used to scan for new features.
type ScanConfig interface {
	// MaxLength returns the length maximum of a sequence detected as feature.
	// E.g. MaxLength set to 3 results in sequences having a length not larger
	// than 3 when detected as features. The value -1 disables any limitiation.
	MaxLength() int
	// MinLength returns the minimum length of a sequence detected as feature.
	// E.g. MinLength set to 3 results in sequences having a length not smaller
	// than 3 when detected as features. The value -1 disables any limitiation.
	MinLength() int
	// MinCount returns the number of occurrences at least required to be detected
	// as feature. E.g. MinCount set to 3 requires a feature to be present within
	// a given input sequence at least 3 times.
	MinCount() int
	// Separator returns the separator used to split sequences into smaller parts.
	// By default this is an empty string resulting in a character split. This can
	// be set to a whitespace to split for words. Note that the concept of words
	// is a feature known to humans based on contextual information humans
	// connected to create reasonable sences. This needs to be achieved by Anna
	// herself. So later this separator needs to be configured by Anna once she is
	// able to recognize improvements in feature detection, resulting in even more
	// awareness of contextual information.
	Separator() string
	// Sequences returns the input sequences being analysed. Out of this
	// information features are detected, if any.
	Sequences() []string
	SetMaxLength(maxLength int)
	SetMinLength(minLength int)
	SetMinCount(minCount int)
	SetSeparator(separate string)
	SetSequences(sequences []string)
	// Validate checks whether ScanConfig is valid for proper execution in
	// Feature.Scan.
	Validate() error
}

// Service represents a service being able to scan for features within
// information sequences.
type Service interface {
	// Boot initializes and starts the whole service like booting a machine. The
	// call to Boot blocks until the service is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()
	// Scan analyses the given sequences to detect patterns. Found patterns are
	// returned in form of a list of features.
	Scan(config ScanConfig) ([]Feature, error)
	// ScanConfig returns a default scan config configured by best effort.
	ScanConfig() ScanConfig
	// Shutdown ends all processes of the service like shutting down a machine.
	// The call to Shutdown blocks until the service is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
