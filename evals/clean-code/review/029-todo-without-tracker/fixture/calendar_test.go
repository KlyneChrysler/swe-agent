package fixture

import "testing"

func TestDaysBetweenYears(t *testing.T) {
	if got := DaysBetweenYears(2020, 2022); got != 730 {
		t.Fatalf("DaysBetweenYears = %d, want 730", got)
	}
}
