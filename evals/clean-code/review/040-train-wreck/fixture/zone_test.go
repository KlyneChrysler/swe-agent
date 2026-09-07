package fixture

import "testing"

func orderBilledIn(country string) Order {
	return Order{customer: Customer{profile: Profile{billing: Address{country: country}}}}
}

func TestIsDomestic(t *testing.T) {
	if !IsDomestic(orderBilledIn(homeCountry)) {
		t.Fatal("home-country order reported as foreign")
	}
	if IsDomestic(orderBilledIn("SG")) {
		t.Fatal("foreign order reported as domestic")
	}
}
