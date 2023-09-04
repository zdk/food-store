package main

type Calculator struct{}

func (c *Calculator) CalculatePrice(order map[string]int, hasMemberCard bool) float64 {

	// Define set_prices for each set
	set_prices := map[string]float64{
		"Red":    50,
		"Green":  40,
		"Blue":   30,
		"Yellow": 50,
		"Pink":   80,
		"Purple": 90,
		"Orange": 120,
	}

	// Calculate the total price
	total := 0.0
	for color, price := range set_prices {
		total += price * float64(order[color])
	}

	// Apply member card discount
	if hasMemberCard {
		total *= 0.9
	}

	return total
}
