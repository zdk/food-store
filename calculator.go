package main

const memberDiscount = 0.9

var setPrices = map[string]float64{
	"Red":    50,
	"Green":  40,
	"Blue":   30,
	"Yellow": 50,
	"Pink":   80,
	"Purple": 90,
	"Orange": 120,
}

var specialDiscounts = map[string]float64{
	"Orange": 0.95,
	"Pink":   0.95,
	"Green":  0.95,
}

type OrderItem struct {
	Name     string
	Quantity int
}

type Calculator struct{}

func (c *Calculator) CalculatePrice(order map[string]int, hasMemberCard bool) float64 {
	// Calculate the total price
	total := 0.0
	for color, price := range setPrices {
		discount, ok := specialDiscounts[color]
		if ok && order[color] > 1 {
			total += price * float64(order[color]) * discount
		} else {
			total += price * float64(order[color])
		}
	}

	// Apply member card discount
	if hasMemberCard {
		total *= memberDiscount
	}

	return total
}
