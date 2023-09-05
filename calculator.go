package main

const memberDiscount = 0.9

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

var itemsMap = map[string]Item{
	"Red":    {Name: "Red", Price: 50, Discount: 0},
	"Green":  {Name: "Green", Price: 40, Discount: 0.95},
	"Blue":   {Name: "Blue", Price: 30, Discount: 0},
	"Yellow": {Name: "Yellow", Price: 50, Discount: 0},
	"Pink":   {Name: "Pink", Price: 80, Discount: 0.95},
	"Purple": {Name: "Purple", Price: 90, Discount: 0},
	"Orange": {Name: "Orange", Price: 120, Discount: 0.95},
}

// orderItems is now a map of item names to quantities
type OrderItems map[string]int

type Calculator struct{}

func (c *Calculator) CalculatePrice(orderItems OrderItems, hasMemberCard bool) float64 {
	// Calculate the total price
	total := 0.0

	for itemName, quantity := range orderItems {
		if item, exists := itemsMap[itemName]; exists {
			if item.Discount > 0 && quantity > 1 {
				total += item.Price * float64(quantity) * item.Discount
			} else {
				total += item.Price * float64(quantity)
			}
		}
	}

	// Apply member card discount
	if hasMemberCard {
		total *= memberDiscount
	}

	return total
}
