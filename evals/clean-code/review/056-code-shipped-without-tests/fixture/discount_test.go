package fixture

import "testing"

func TestTierStringNamesGold(t *testing.T) {
	if Gold.String() != "gold" {
		t.Errorf("Gold.String() = %q, want %q", Gold.String(), "gold")
	}
}
