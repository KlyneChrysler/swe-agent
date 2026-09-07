package fixture

import "time"

const lateFeeCents = 1500

type PaymentRow struct {
	InvoiceID string
	DueAt     time.Time
	Cents     int
}

func (r *PaymentRow) ApplyLateFee(now time.Time) {
	if now.After(r.DueAt) {
		r.Cents += lateFeeCents
	}
}

func TotalCents(rows []PaymentRow) int {
	total := 0
	for _, row := range rows {
		total += row.Cents
	}
	return total
}
