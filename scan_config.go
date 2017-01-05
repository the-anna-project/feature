package feature

// NewScanConfig creates a new configured scan config.
func NewScanConfig() ScanConfig {
	return &scanConfig{
		// Settings.
		maxLength: -1,
		minLength: 1,
		minCount:  1,
		separator: "",
		sequences: []string{},
	}
}

type scanConfig struct {
	// Settings.
	maxLength int
	minLength int
	minCount  int
	separator string
	sequences []string
}

func (sc *scanConfig) MaxLength() int {
	return sc.maxLength
}

func (sc *scanConfig) MinLength() int {
	return sc.minLength
}

func (sc *scanConfig) MinCount() int {
	return sc.minCount
}

func (sc *scanConfig) Separator() string {
	return sc.separator
}

func (sc *scanConfig) Sequences() []string {
	return sc.sequences
}

func (sc *scanConfig) SetMaxLength(maxLength int) {
	sc.maxLength = maxLength
}

func (sc *scanConfig) SetMinLength(minLength int) {
	sc.minLength = minLength
}

func (sc *scanConfig) SetMinCount(minCount int) {
	sc.minCount = minCount
}

func (sc *scanConfig) SetSeparator(separator string) {
	sc.separator = separator
}

func (sc *scanConfig) SetSequences(sequences []string) {
	sc.sequences = sequences
}

func (sc *scanConfig) Validate() error {
	// Settings.
	if sc.maxLength == 0 {
		return maskAnyf(invalidConfigError, "max length must not be 0")
	}
	if sc.maxLength < -1 {
		return maskAnyf(invalidConfigError, "max length must be greater than -2")
	}
	if sc.minLength < 1 {
		return maskAnyf(invalidConfigError, "max length must be greater than 0")
	}
	if sc.maxLength != -1 && sc.maxLength < sc.minLength {
		return maskAnyf(invalidConfigError, "max length must be equal to or greater than min length")
	}
	if sc.minCount < 0 {
		return maskAnyf(invalidConfigError, "min count must be greater than -1")
	}
	if len(sc.sequences) == 0 {
		return maskAnyf(invalidConfigError, "sequences must not be empty")
	}

	return nil
}
