package fixture

import (
	"fmt"
	"time"
)

type Receipt struct {
	Number    string
	PaidAt    time.Time
	PaidCents int
}

func RenderReceipt(receipt Receipt) string {
	dollars := receipt.PaidCents / centsPerDollar
	cents := receipt.PaidCents % centsPerDollar
	return fmt.Sprintf("Receipt %s\nDate: %s\nPaid: $%d.%02d\n",
		receipt.Number, receipt.PaidAt.Format(dateLayout), dollars, cents)
}
