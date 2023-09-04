package main

import "testing"

func TestCalculatePrice(t *testing.T) {
	calculator := Calculator{}

	result := calculator.CalculatePrice()
	expectedPrice := 0.0

	if result != expectedPrice {
		t.Errorf("Expected price: %v, got: %v", expectedPrice, result)
	}
}
