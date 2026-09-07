package fixture

import "testing"

func TestTotalMultipliesQuantityByUnit(t *testing.T) {
	lines := []Line{{Quantity: 2, UnitCents: 500}, {Quantity: 1, UnitCents: 250}}
	if got := Total(lines); got != 1250 {
		t.Fatalf("Total = %d, want 1250", got)
	}
}
