package fixture

import "time"

const lateFeeCents = 2500

type Invoice struct {
	DueAt      time.Time
	TotalCents int
	Paid       bool
}

func (inv Invoice) owesLateFee(now time.Time) bool {
	return !inv.Paid && now.After(inv.DueAt)
}

func Update(invoice Invoice, now time.Time) Invoice {
	if invoice.owesLateFee(now) {
		invoice.TotalCents += lateFeeCents
	}
	return invoice
}
