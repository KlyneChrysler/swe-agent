package fixture

const roundingUnitCents = 100

type Line struct {
	Quantity  int
	UnitCents int
}

func Total(lines []Line) int {
	total := 0
	for _, line := range lines {
		total += line.Quantity * line.UnitCents
	}
	return total
}

func totalRoundedToDollar(lines []Line) int {
	return (Total(lines) + roundingUnitCents/2) / roundingUnitCents * roundingUnitCents
}
