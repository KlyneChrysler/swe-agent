package fixture

import (
	"testing"
	"time"
)

func TestDaysOverdueCountsWholeDaysPastGrace(t *testing.T) {
	issued := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	inv := Invoice{IssuedAt: issued, Terms: 30 * day}
	now := issued.Add(35 * day)
	if got := DaysOverdue(inv, now); got != 2 {
		t.Fatalf("DaysOverdue = %d, want 2", got)
	}
}

func TestDaysOverdueIsZeroBeforeDue(t *testing.T) {
	issued := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	inv := Invoice{IssuedAt: issued, Terms: 30 * day}
	if got := DaysOverdue(inv, issued.Add(day)); got != 0 {
		t.Fatalf("DaysOverdue = %d, want 0", got)
	}
}
