package fixture

import (
	"errors"
	"testing"
)

var catalog = NewCatalog(map[string]Discount{"EL-1": {Percent: 10}})

func TestDiscountForUnknownSKU(t *testing.T) {
	if _, err := catalog.DiscountFor("EL-9"); !errors.Is(err, ErrNoDiscount) {
		t.Fatalf("DiscountFor = %v, want ErrNoDiscount", err)
	}
}

func TestPriceCents(t *testing.T) {
	if got := catalog.PriceCents("EL-1", 1000); got != 900 {
		t.Fatalf("discounted PriceCents = %d, want 900", got)
	}
	if got := catalog.PriceCents("EL-9", 1000); got != 1000 {
		t.Fatalf("undiscounted PriceCents = %d, want 1000", got)
	}
}
