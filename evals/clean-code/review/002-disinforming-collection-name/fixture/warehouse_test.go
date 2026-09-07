package fixture

import "testing"

func TestReceiveAccumulatesOnHand(t *testing.T) {
	warehouse := NewWarehouse()
	warehouse.Receive("EL-00042", 3)
	warehouse.Receive("EL-00042", 2)
	if got := warehouse.OnHand("EL-00042"); got != 5 {
		t.Fatalf("OnHand = %d, want 5", got)
	}
}

func TestOnHandIsZeroForUnknownSKU(t *testing.T) {
	if got := NewWarehouse().OnHand("missing"); got != 0 {
		t.Fatalf("OnHand = %d, want 0", got)
	}
}
