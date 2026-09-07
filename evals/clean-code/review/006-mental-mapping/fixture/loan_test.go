package fixture

import (
	"math"
	"testing"
)

func TestMonthlyPaymentMatchesAmortizationTable(t *testing.T) {
	const tolerance = 0.01
	got := MonthlyPayment(100000, 0.06, 360)
	if math.Abs(got-599.55) > tolerance {
		t.Fatalf("MonthlyPayment = %.2f, want 599.55", got)
	}
}
