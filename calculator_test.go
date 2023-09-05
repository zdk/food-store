package main

import (
	"fmt"
	"testing"
)

func TestCalculatePrice(t *testing.T) {
	calc := Calculator{}

	tests := []struct {
		name          string
		orderItems    []OrderItem
		hasMemberCard bool
		expected      float64
	}{
		{
			name: "Test single item without discount",
			orderItems: []OrderItem{
				{Name: "Red", Quantity: 1},
			},
			hasMemberCard: false,
			expected:      50.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.orderItems)
			got := calc.CalculatePrice(tt.orderItems, tt.hasMemberCard)
			if got != tt.expected {
				t.Errorf("Expected %f, but got %f", tt.expected, got)
			}
		})
	}
}
