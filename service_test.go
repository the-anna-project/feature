package feature

import (
	"reflect"
	"sort"
	"testing"

	"github.com/the-anna-project/context"
)

func Test_Service_Scan(t *testing.T) {
	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	ctx, err := context.New(context.DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	scanConfig := newService.ScanConfig()
	scanConfig.SetMinCount(2)
	scanConfig.SetSequences([]string{
		"This is, a test.",
		"This is, another test.",
	})

	features, err := newService.Scan(ctx, scanConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for _, f := range features {
		// We only want to verify the feature represented by the information
		// sequence ".".
		if f.Sequence() != "." {
			continue
		}

		if f.Count() != 2 {
			t.Fatal("expected", 2, "got", f.Count())
		}
		if f.Sequence() != "." {
			t.Fatal("expected", ".", "got", f.Sequence())
		}
	}
}

func Test_Service_Scan_MinLengthMaxLength(t *testing.T) {
	testCases := []struct {
		MinLength    int
		MaxLength    int
		Sequences    []string
		Expected     []string
		ErrorMatcher func(err error) bool
	}{
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab"},
			Expected:     nil,
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab"},
			Expected:     []string{"a", "b", "ab"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "ab"},
			Expected:     []string{"a", "b", "ab"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc"},
			Expected:     []string{"a", "b", "ab"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"a", "b", "c", "ab", "bc", "abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"a", "b", "c"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    2,
			MaxLength:    2,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"ab", "bc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    3,
			MaxLength:    3,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    2,
			MaxLength:    3,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"ab", "bc", "abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    2,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     []string{"ab", "bc", "abc"},
			ErrorMatcher: nil,
		},
		{
			MinLength:    1,
			MaxLength:    -2,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     nil,
			ErrorMatcher: IsInvalidConfig,
		},
		{
			MinLength:    2,
			MaxLength:    1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     nil,
			ErrorMatcher: IsInvalidConfig,
		},
		{
			MinLength:    0,
			MaxLength:    -1,
			Sequences:    []string{"ab", "ab", "abc", "abc"},
			Expected:     nil,
			ErrorMatcher: IsInvalidConfig,
		},
	}

	for i, testCase := range testCases {
		newService, err := NewService(DefaultServiceConfig())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		ctx, err := context.New(context.DefaultConfig())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		scanConfig := newService.ScanConfig()
		scanConfig.SetMinCount(2)
		scanConfig.SetMaxLength(testCase.MaxLength)
		scanConfig.SetMinLength(testCase.MinLength)
		scanConfig.SetSequences(testCase.Sequences)

		features, err := newService.Scan(ctx, scanConfig)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			var sequences []string
			for _, f := range features {
				sequences = append(sequences, f.Sequence())
			}

			sort.Strings(sequences)
			sort.Strings(testCase.Expected)
			if !reflect.DeepEqual(sequences, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", sequences)
			}
		}
	}
}
