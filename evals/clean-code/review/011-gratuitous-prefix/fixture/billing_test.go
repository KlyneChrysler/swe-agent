package fixture

import "testing"

func TestBillingTotalAddsTax(t *testing.T) {
	lines := []BillingLine{{Description: "seat", Cents: 600}, {Description: "seat", Cents: 400}}
	invoice := BillingInvoice{Lines: lines}
	if got := invoice.BillingSubtotal(); got != 1000 {
		t.Fatalf("BillingSubtotal = %d, want 1000", got)
	}
	if got := invoice.BillingTotal(); got != 1120 {
		t.Fatalf("BillingTotal = %d, want 1120", got)
	}
}
