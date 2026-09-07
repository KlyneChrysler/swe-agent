package fixture

import "math"

const monthsPerYear = 12

func MonthlyPayment(p, r float64, n int) float64 {
	q := r / monthsPerYear
	g := math.Pow(1+q, float64(n))
	return p * q * g / (g - 1)
}
