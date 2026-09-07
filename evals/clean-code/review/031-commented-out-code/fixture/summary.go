package fixture

import "fmt"

const centsPerDollar = 100

type Sale struct {
	Cents int
}

func Summary(sales []Sale) string {
	total := 0
	for _, sale := range sales {
		total += sale.Cents
	}
	// if total > 100000 {
	// 	return "large: " + formatCents(total)
	// }
	return formatCents(total)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func formatCents(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/centsPerDollar, cents%centsPerDollar)
}
