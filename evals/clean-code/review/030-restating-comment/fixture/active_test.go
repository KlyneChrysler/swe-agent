package fixture

import "testing"

func TestCountActive(t *testing.T) {
	members := []Member{{Name: "Ada", Active: true}, {Name: "Bo"}, {Name: "Cy", Active: true}}
	if got := CountActive(members); got != 2 {
		t.Fatalf("CountActive = %d, want 2", got)
	}
}
