package fixture

import "testing"

func TestTakeReturnsMostRecentPut(t *testing.T) {
	var pile Pile
	pile.Put("first")
	pile.Put("second")
	if item, ok := pile.Take(); !ok || item != "second" {
		t.Fatalf("Take = %q, %v; want second, true", item, ok)
	}
	if got := pile.Size(); got != 1 {
		t.Fatalf("Size = %d, want 1", got)
	}
}

func TestTakeOnEmptyReportsFalse(t *testing.T) {
	var pile Pile
	if _, ok := pile.Take(); ok {
		t.Fatal("Take on empty = true, want false")
	}
}
