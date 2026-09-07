package fixture

import (
	"testing"
	"time"
)

var dueAt = time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)

func TestLateFeeCentsForOverdueUndisputed(t *testing.T) {
	inv := Invoice{TotalCents: 1000, DueAt: dueAt}
	if got := LateFeeCents(inv, dueAt.Add(time.Hour)); got != lateFeeCents {
		t.Fatalf("LateFeeCents = %d, want %d", got, lateFeeCents)
	}
}

func TestLateFeeCentsWaivedWhenDisputed(t *testing.T) {
	inv := Invoice{TotalCents: 1000, DueAt: dueAt, Disputed: true}
	if got := LateFeeCents(inv, dueAt.Add(time.Hour)); got != 0 {
		t.Fatalf("LateFeeCents = %d, want 0", got)
	}
}

func TestStatus(t *testing.T) {
	if got := Status(Invoice{TotalCents: 1000, PaidCents: 1000}); got != "settled" {
		t.Fatalf("Status = %q, want settled", got)
	}
	if got := Status(Invoice{TotalCents: 1000}); got != "open" {
		t.Fatalf("Status = %q, want open", got)
	}
}
