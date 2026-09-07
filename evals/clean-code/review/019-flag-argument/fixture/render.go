package fixture

import "fmt"

type Invoice struct {
	Number     string
	Customer   string
	TotalCents int
}

func Render(inv Invoice, compact bool) string {
	if compact {
		return fmt.Sprintf("%s %d", inv.Number, inv.TotalCents)
	}
	return fmt.Sprintf("Invoice %s\nCustomer: %s\nTotal: %d cents\n", inv.Number, inv.Customer, inv.TotalCents)
}
