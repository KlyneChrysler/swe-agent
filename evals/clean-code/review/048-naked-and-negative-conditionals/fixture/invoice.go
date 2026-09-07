package fixture

import "time"

const lateFeeCents = 2500

type Invoice struct {
	TotalCents int
	PaidCents  int
	DueAt      time.Time
	Disputed   bool
}

func (inv Invoice) IsNotSettled() bool {
	return inv.PaidCents < inv.TotalCents
}

func LateFeeCents(inv Invoice, now time.Time) int {
	if inv.PaidCents < inv.TotalCents && now.After(inv.DueAt) && !inv.Disputed {
		return lateFeeCents
	}
	return 0
}

func Status(inv Invoice) string {
	if !inv.IsNotSettled() {
		return "settled"
	}
	return "open"
}
