package fixture

import "testing"

func TestQualifiesForFreeShipping(t *testing.T) {
	qualifying := Order{SubtotalCents: 6000, DiscountCents: 500, Destination: "PH", Merchant: "PH", LineCount: 2}
	if !QualifiesForFreeShipping(qualifying) {
		t.Fatal("domestic order over the threshold did not qualify")
	}
	foreign := qualifying
	foreign.Destination = "SG"
	if QualifiesForFreeShipping(foreign) {
		t.Fatal("foreign order qualified")
	}
	discounted := qualifying
	discounted.DiscountCents = 1500
	if QualifiesForFreeShipping(discounted) {
		t.Fatal("order under the net threshold qualified")
	}
}
