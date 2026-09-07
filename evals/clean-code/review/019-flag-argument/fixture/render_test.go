package fixture

import "testing"

var sample = Invoice{Number: "INV-7", Customer: "Ada", TotalCents: 1250}

func TestRenderCompact(t *testing.T) {
	if got := Render(sample, true); got != "INV-7 1250" {
		t.Fatalf("Render compact = %q", got)
	}
}

func TestRenderFull(t *testing.T) {
	want := "Invoice INV-7\nCustomer: Ada\nTotal: 1250 cents\n"
	if got := Render(sample, false); got != want {
		t.Fatalf("Render full = %q, want %q", got, want)
	}
}
