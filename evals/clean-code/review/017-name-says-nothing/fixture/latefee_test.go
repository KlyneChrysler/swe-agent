package fixture

import (
	"testing"
	"time"
)

var due = time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)

func TestUpdateAddsFeeWhenOverdue(t *testing.T) {
	invoice := Invoice{DueAt: due, TotalCents: 10000}
	if got := Update(invoice, due.Add(time.Hour)).TotalCents; got != 12500 {
		t.Fatalf("TotalCents = %d, want 12500", got)
	}
}

func TestUpdateLeavesPaidInvoiceAlone(t *testing.T) {
	invoice := Invoice{DueAt: due, TotalCents: 10000, Paid: true}
	if got := Update(invoice, due.Add(time.Hour)).TotalCents; got != 10000 {
		t.Fatalf("TotalCents = %d, want 10000", got)
	}
}
