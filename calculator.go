package main

type Calculator struct{}

func (c *Calculator) CalculatePrice(order map[string]int) float64 {

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
	total := set_prices["Red"]*float64(order["Red"]) + set_prices["Green"]*float64(order["Green"]) +
		set_prices["Blue"]*float64(order["Blue"]) + set_prices["Yellow"]*float64(order["Yellow"]) +
		set_prices["Pink"]*float64(order["Pink"]) + set_prices["Purple"]*float64(order["Purple"]) +
		set_prices["Orange"]*float64(order["Orange"])
	return total
}
