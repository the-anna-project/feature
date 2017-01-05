package feature

import (
	"reflect"
	"testing"
)

func Test_Feature_New_Error_Positions_Empty(t *testing.T) {
	newConfig := DefaultConfig()
	// Note positions configuration is missing.
	newConfig.Sequence = "test"
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Feature_New_Error_Positions_Dimension(t *testing.T) {
	newConfig := DefaultConfig()
	// Note positions configuration is invalid.
	newConfig.Positions = [][]float64{{0}, {3, 4, 5, 6}}
	newConfig.Sequence = "test"
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Feature_New_Error_Sequence_Empty(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Positions = [][]float64{{0, 1}, {3, 4}}
	// Note sequence configuration is missing.
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Feature_AddPosition(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Positions = [][]float64{{0, 1}}
	newConfig.Sequence = "test"
	newFeature, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	output := newFeature.Positions()
	if !reflect.DeepEqual(output, [][]float64{{0, 1}}) {
		t.Fatal("expected", [][]float64{{0, 1}}, "got", output)
	}

	err = newFeature.AddPosition([]float64{3, 4})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	output = newFeature.Positions()
	if !reflect.DeepEqual(output, [][]float64{{0, 1}, {3, 4}}) {
		t.Fatal("expected", [][]float64{{0, 1}, {3, 4}}, "got", output)
	}

	err = newFeature.AddPosition([]float64{3})
	if !IsInvalidExecution(err) {
		t.Fatal("expected", true, "got", false)
	}
	err = newFeature.AddPosition([]float64{4})
	if !IsInvalidExecution(err) {
		t.Fatal("expected", true, "got", false)
	}
	err = newFeature.AddPosition([]float64{3, 4, 5})
	if !IsInvalidExecution(err) {
		t.Fatal("expected", true, "got", false)
	}
}
