package fixture

import "testing"

func TestApplyUppercasesUSAddresses(t *testing.T) {
	got := NormalizeAddress{Country: unitedStates}.Apply("  12 main st ")
	if got != "12 MAIN ST" {
		t.Fatalf("Apply = %q, want 12 MAIN ST", got)
	}
}

func TestApplyOnlyTrimsElsewhere(t *testing.T) {
	got := NormalizeAddress{Country: "PH"}.Apply("  12 main st ")
	if got != "12 main st" {
		t.Fatalf("Apply = %q, want 12 main st", got)
	}
}
