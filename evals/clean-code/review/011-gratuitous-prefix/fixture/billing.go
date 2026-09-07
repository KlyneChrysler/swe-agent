package fixture

const (
	billingTaxPercent  = 12
	percentDenominator = 100
)

type BillingLine struct {
	Description string
	Cents       int
}

type BillingInvoice struct {
	Lines []BillingLine
}

func (inv BillingInvoice) BillingSubtotal() int {
	subtotal := 0
	for _, line := range inv.Lines {
		subtotal += line.Cents
	}
	return subtotal
}

func (inv BillingInvoice) BillingTotal() int {
	subtotal := inv.BillingSubtotal()
	return subtotal + subtotal*billingTaxPercent/percentDenominator
}
