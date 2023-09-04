package main

import "testing"

func TestCalculatePrice(t *testing.T) {
	calculator := Calculator{}
	order := map[string]int{
			"Red": 1,
			"Green": 1,
	}
	result := calculator.CalculatePrice(order)
	expected := 90.0
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}
