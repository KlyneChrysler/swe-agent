package fixture

import (
	"testing"
	"time"
)

var dueDate = time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)

func TestApplyLateFeeAfterDue(t *testing.T) {
	row := PaymentRow{InvoiceID: "INV-7", DueAt: dueDate, Cents: 10000}
	row.ApplyLateFee(dueDate.Add(time.Hour))
	if row.Cents != 11500 {
		t.Fatalf("Cents = %d, want 11500", row.Cents)
	}
}

func TestTotalCentsSumsRows(t *testing.T) {
	rows := []PaymentRow{{Cents: 100}, {Cents: 250}}
	if got := TotalCents(rows); got != 350 {
		t.Fatalf("TotalCents = %d, want 350", got)
	}
}
