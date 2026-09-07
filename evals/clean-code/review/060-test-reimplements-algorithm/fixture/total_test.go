package fixture

import "testing"

func TestTotalMatchesRecomputation(t *testing.T) {
	lines := []Line{
		{Cents: 1000, Quantity: 3, DiscountPercent: 15},
		{Cents: 250, Quantity: 2, DiscountPercent: 0},
	}
	want := 0.0
	for _, line := range lines {
		want += float64(line.Cents*line.Quantity) * (1 - float64(line.DiscountPercent)/100)
	}

	if Total(lines) != int64(want) {
		t.Errorf("Total = %d, want %d", Total(lines), int64(want))
	}
}

func TestZeroQuantityLineRegression(t *testing.T) {
	lines := []Line{{Cents: 500, Quantity: 0, DiscountPercent: 10}}

	if Total(lines) != 0 {
		t.Errorf("Total = %d, want 0", Total(lines))
	}
}
