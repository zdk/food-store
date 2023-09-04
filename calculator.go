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

type Calculator struct{}

func (c *Calculator) CalculatePrice(order map[string]int, hasMemberCard bool) float64 {
	// Calculate the total price
	total := 0.0
	for color, price := range setPrices {
		total += price * float64(order[color])
	}

	// Apply member card discount
	if hasMemberCard {
		total *= memberDiscount
	}

	return total
}
