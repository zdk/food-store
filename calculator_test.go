package main

import (
	"testing"
)

func TestCalculatePrice(t *testing.T) {
	calculator := &Calculator{}

	tests := []struct {
		name          string
		orderItems    OrderItems
		hasMemberCard bool
		expected      float64
	}{
		{
			name: "Single item without discount",
			orderItems: OrderItems{
				"Red": 1,
			},
			hasMemberCard: false,
			expected:      50.0,
		},
		{
			name: "Single item with quantity > 1 and discount",
			orderItems: OrderItems{
				"Green": 2,
			},
			hasMemberCard: false,
			expected:      76.0, // 40 * 2 * 0.95
		},
		{
			name: "Multiple items, some with discount",
			orderItems: OrderItems{
				"Red":    2,
				"Green":  3,
				"Purple": 1,
			},
			hasMemberCard: false,
			expected:      304.0, // (50*2)+((40*3)*0.95)+90
		},
		{
			name: "Multiple items with member card discount",
			orderItems: OrderItems{
				"Red":    2,
				"Green":  3,
				"Purple": 1,
			},
			hasMemberCard: true,
			expected:      273.6, // 304.0 * 0.9
		},
		{
			name: "Item not in itemsMap",
			orderItems: OrderItems{
				"NonExistentColor": 2,
			},
			hasMemberCard: false,
			expected:      0.0, // Not found in itemsMap, hence price is 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calculator.CalculatePrice(tt.orderItems, tt.hasMemberCard)
			if got != tt.expected {
				t.Errorf("Expected %f, but got %f", tt.expected, got)
			}
		})
	}
}

func TestCalculatePrice_ErrorCase(t *testing.T) {
	calculator := &Calculator{}

	// Create a test order with an unknown item
	order := OrderItems{
		"UnknownItem": 2,
	}

	_, err := calculator.CalculatePrice(order, false)
	if err == nil {
		t.Errorf("Expected an error for unknown item, got none")
		return
	}

	expectedError := "item not found: UnknownItem"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %s, got: %s", expectedError, err.Error())
	}
}
