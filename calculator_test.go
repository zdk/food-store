package main

import "testing"

func TestCalculatePrice(t *testing.T) {
	calculator := Calculator{}

	tests := []struct {
		order         map[string]int
		hasMemberCard bool
		expectedPrice float64
	}{
		{map[string]int{"Red": 1, "Green": 1}, false, 90},
		{map[string]int{"Red": 1, "Green": 1}, true, 81},
		{map[string]int{"Orange": 2}, false, 228},
		{map[string]int{"Orange": 3}, false, 342},
		//{map[string]int{"Orange": 2}, true, 205.2},
	}

	for _, test := range tests {
		result := calculator.CalculatePrice(test.order, test.hasMemberCard)
		if result != test.expectedPrice {
			t.Errorf("Expected price: %v, got: %v", test.expectedPrice, result)
		}
	}
}
