package fixture

import "testing"

func TestAccountAccessorsRoundTrip(t *testing.T) {
	var account Account
	account.SetOwner("Ada")
	account.SetBalanceCents(1250)
	if account.GetOwner() != "Ada" || account.GetBalanceCents() != 1250 {
		t.Fatalf("account = %q %d", account.GetOwner(), account.GetBalanceCents())
	}
}
