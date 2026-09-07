package fixture

// Line is a priced quantity of one item with a whole-number discount.
type Line struct {
	Cents           int64
	Quantity        int64
	DiscountPercent int64
}

const percent = 100

func Total(lines []Line) int64 {
	var total int64
	for _, line := range lines {
		total += line.net()
	}
	return total
}

func (l Line) net() int64 {
	gross := l.Cents * l.Quantity
	return gross - gross*l.DiscountPercent/percent
}
