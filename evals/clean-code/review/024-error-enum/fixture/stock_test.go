package fixture

import "testing"

func TestRestockStatuses(t *testing.T) {
	cases := []struct {
		sku, quantity string
		want          Status
	}{
		{"EL-1", "5", OK},
		{"EL-9", "5", NotFound},
		{"EL-1", "five", Malformed},
	}
	for _, c := range cases {
		stock := NewStock(map[string]int{"EL-1": 1})
		if got := stock.Restock(c.sku, c.quantity); got != c.want {
			t.Errorf("Restock(%q, %q) = %v, want %v", c.sku, c.quantity, got, c.want)
		}
	}
}

func TestDescribeMalformed(t *testing.T) {
	if got := Describe(Malformed); got != "quantity is not a number" {
		t.Fatalf("Describe = %q", got)
	}
}
