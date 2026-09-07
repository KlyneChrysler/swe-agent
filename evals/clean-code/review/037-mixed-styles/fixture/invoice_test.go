package fixture

import "testing"

func TestDiscountAfterAddingLines(t *testing.T) {
	var invoice Invoice
	invoice.AddLine(800)
	invoice.AddLine(200)
	invoice.Apply_Discount(10)
	if invoice.TotalCents != 900 {
		t.Fatalf("TotalCents = %d, want 900", invoice.TotalCents)
	}
}

func TestParseCentsAndPercent(t *testing.T) {
	if cents, err := ParseCents("1250"); err != nil || cents != 1250 {
		t.Fatalf("ParseCents = %d, %v", cents, err)
	}
	if _, err := ParsePercent("ten"); err == nil {
		t.Fatal("ParsePercent accepted a word")
	}
}
