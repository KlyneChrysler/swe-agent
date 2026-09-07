package fixture

import "testing"

func TestPriceCentsByKind(t *testing.T) {
	cases := map[Kind]int{Physical: 1000, Digital: 900, Subscription: 12000}
	for kind, want := range cases {
		if got := PriceCents(Item{Kind: kind, BaseCents: 1000}); got != want {
			t.Errorf("PriceCents(%v) = %d, want %d", kind, got, want)
		}
	}
}

func TestLabelByKind(t *testing.T) {
	cases := map[Kind]string{Physical: "ships in 3-5 days", Digital: "download", Subscription: "yearly plan"}
	for kind, want := range cases {
		if got := Label(Item{Kind: kind}); got != want {
			t.Errorf("Label(%v) = %q, want %q", kind, got, want)
		}
	}
}
