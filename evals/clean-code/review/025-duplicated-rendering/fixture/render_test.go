package fixture

import (
	"testing"
	"time"
)

var issued = time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)

func TestRenderInvoice(t *testing.T) {
	want := "Invoice INV-7\nDate: 2024-03-01\nTotal: $12.05\n"
	if got := RenderInvoice(Invoice{Number: "INV-7", IssuedAt: issued, TotalCents: 1205}); got != want {
		t.Fatalf("RenderInvoice = %q, want %q", got, want)
	}
}

func TestRenderReceipt(t *testing.T) {
	want := "Receipt R-7\nDate: 2024-03-01\nPaid: $12.05\n"
	if got := RenderReceipt(Receipt{Number: "R-7", PaidAt: issued, PaidCents: 1205}); got != want {
		t.Fatalf("RenderReceipt = %q, want %q", got, want)
	}
}
