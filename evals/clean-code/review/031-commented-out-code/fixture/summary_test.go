package fixture

import "testing"

func TestSummaryFormatsTotal(t *testing.T) {
	if got := Summary([]Sale{{Cents: 1250}, {Cents: 55}}); got != "$13.05" {
		t.Fatalf("Summary = %q, want $13.05", got)
	}
}
