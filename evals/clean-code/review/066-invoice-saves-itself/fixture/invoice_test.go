package fixture

import "testing"

func TestTotalAddsTax(t *testing.T) {
	invoice := Invoice{ID: "A1", Customer: "Ada", Cents: 1000, TaxPercent: 20}

	if invoice.Total() != 1200 {
		t.Errorf("Total = %d, want 1200", invoice.Total())
	}
}

func TestTotalWithoutTaxIsUnchanged(t *testing.T) {
	invoice := Invoice{ID: "A1", Customer: "Ada", Cents: 1000}

	if invoice.Total() != 1000 {
		t.Errorf("Total = %d, want 1000", invoice.Total())
	}
}
