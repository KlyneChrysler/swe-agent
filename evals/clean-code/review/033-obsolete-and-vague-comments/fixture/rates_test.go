package fixture

import "testing"

func TestRateForKnownZone(t *testing.T) {
	table := NewRateTable(map[string]int{"metro": 12})
	if got, err := table.RateFor("metro"); err != nil || got != 12 {
		t.Fatalf("RateFor = %d, %v; want 12, nil", got, err)
	}
}

func TestRateForUnknownZoneErrors(t *testing.T) {
	table := NewRateTable(map[string]int{"metro": 12})
	if _, err := table.RateFor("rural"); err == nil {
		t.Fatal("RateFor accepted an unknown zone")
	}
	if got := table.Zones(); got != 1 {
		t.Fatalf("Zones = %d, want 1", got)
	}
}
