package fixture

import "testing"

func TestTotalAppliesLineDiscounts(t *testing.T) {
	lines := []Line{
		{Quantity: 2, UnitCents: 500, DiscountCents: 100},
		{Quantity: 1, UnitCents: 250},
	}
	if got := Total(lines); got != 1150 {
		t.Fatalf("Total = %d, want 1150", got)
	}
}
