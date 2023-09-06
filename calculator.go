// Package main provides functionality to calculate the total cost of an order
package main

import "fmt"

// memberDiscount represents the percentage discount for members.
const memberDiscount = 0.9

// itemsMap is a map of available items in the store. The key is the name of the item.
var itemsMap = map[string]Item{
	"Red":    {Price: 50, Discount: 0},
	"Green":  {Price: 40, Discount: 0.95},
	"Blue":   {Price: 30, Discount: 0},
	"Yellow": {Price: 50, Discount: 0},
	"Pink":   {Price: 80, Discount: 0.95},
	"Purple": {Price: 90, Discount: 0},
	"Orange": {Price: 120, Discount: 0.95},
}

// OrderItems is a map of item names to quantities
type OrderItems map[string]int

// Item struct defines properties of a purchasable item.
type Item struct {
	Name     string  // Name of the item, e.g., "Red", "Green", etc.
	Price    float64 // Base price of the item.
	Discount float64 // Discount multiplier, e.g., 0.95 implies a 5% discount. A value of 0 means no discount.
}

// DiscountedPrice calculates the total price of an item after considering any discounts and the quantity.
func (i *Item) DiscountedPrice(quantity int) float64 {
	if i.Discount > 0 && quantity > 1 {
		return i.Price * i.Discount * float64(quantity)
	}
	return i.Price * float64(quantity)
}

// Calculator struct can be used to calculate the price of an order.
type Calculator struct{}

// CalculatePrice computes the total cost of an order, considering any discounts
func (c *Calculator) CalculatePrice(orderItems OrderItems, hasMemberCard bool) (float64, error) {
	total := 0.0
	for itemName, quantity := range orderItems {
		item, exists := itemsMap[itemName]
		if !exists {
			return 0, fmt.Errorf("item not found: %s", itemName)
		}
		total += item.DiscountedPrice(quantity)
	}
	// Apply member card discount if applicable
	if hasMemberCard {
		total *= memberDiscount
	}
	return total, nil
}
