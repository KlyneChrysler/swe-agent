package fixture

import (
	"fmt"
	"time"
)

type Invoice struct {
	Number     string
	IssuedAt   time.Time
	TotalCents int
}

func RenderInvoice(inv Invoice) string {
	dollars := inv.TotalCents / centsPerDollar
	cents := inv.TotalCents % centsPerDollar
	return fmt.Sprintf("Invoice %s\nDate: %s\nTotal: $%d.%02d\n",
		inv.Number, inv.IssuedAt.Format(dateLayout), dollars, cents)
}
