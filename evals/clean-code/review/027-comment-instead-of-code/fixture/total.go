package fixture

type Line struct {
	Quantity      int
	UnitCents     int
	DiscountCents int
}

func Total(lines []Line) int {
	// sum of quantity times unit price, minus the per-line discount, in cents
	total := 0
	for _, line := range lines {
		total += line.Quantity*line.UnitCents - line.DiscountCents
	}
	return total
}
