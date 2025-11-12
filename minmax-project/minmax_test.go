package main

import (
	"testing"
)

func TestMinmax_NormalCase(t *testing.T) {

// Scenario: Normal values ​​within and outside the range	
	min := 10.0
	max := 20.0
	values := []float64{5.0, 12.0, 15.0, 18.0, 25.0, 10.0, 20.0}


	expected := []float64{12.0, 15.0, 18.0, 10.0, 20.0}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_NormalCase failed: expected %v, got %v", expected, result)
	}
}

func TestMinmax_SingleValue(t *testing.T) {
	// Scenario: Only one value within the range
	min := 5.0
	max := 15.0
	values := []float64{10.0}

	expected := []float64{10.0}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_SingleValue failed: expected %v, got %v", expected, result)
	}
}


func TestMinmax_MinGreaterThanMax(t *testing.T) {
	// Scenario: Minimum greater than maximum (the function should work the same)
	min := 20.0
	max := 10.0
	values := []float64{5.0, 15.0, 25.0}

	expected := []float64{}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_MinGreaterThanMax failed: expected %v, got %v", expected, result)
	}
}

func TestMinmax_NoValuesInRange(t *testing.T) {
	// Scenario: No values within the range
	min := 10.0
	max := 20.0
	values := []float64{5.0, 25.0, 30.0, 1.0}

	expected := []float64{}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_NoValuesInRange failed: expected %v, got %v", expected, result)
	}
}

func TestMinmax_NegativeRange(t *testing.T) {
	// Scenario: Range with negative values
	min := -15.0
	max := -5.0
	values := []float64{-20.0, -12.0, -10.0, -8.0, -3.0, 0.0}

	expected := []float64{-12.0, -10.0, -8.0}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_NegativeRange failed: expected %v, got %v", expected, result)
	}
}

func TestMinmax_MinEqualsMax(t *testing.T) {
	// Scenario: Minimum equals maximum
	min := 10.0
	max := 10.0
	values := []float64{5.0, 10.0, 15.0, 10.0, 9.999, 10.001}

	expected := []float64{10.0, 10.0}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_MinEqualsMax failed: expected %v, got %v", expected, result)
	}
}

func TestMinmax_EmptyValues(t *testing.T) {
	// Scenario: No input values
	min := 10.0
	max := 20.0
	values := []float64{}

	expected := []float64{}
	result := minmax(min, max, values...)

	if !sameSlices(result, expected) {
		t.Errorf("TestMinmax_EmptyValues failed: expected %v, got %v", expected, result)
	}
}


// Auxiliary function to compare slices of float64
func sameSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false 
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
