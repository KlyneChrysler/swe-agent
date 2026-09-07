package fixture

import "fmt"

func formatCents(cents int64) string {
	return fmt.Sprintf("%s%d.%02d", currencySymbol, cents/centsPerUnit, cents%centsPerUnit)
}

// Customer is the billed party.
type Customer struct {
	Name  string
	Email string
}

func (c Customer) Greeting() string {
	return "Dear " + c.Name
}

type Invoice struct {
	customer Customer
	cents    int64
}

func (i Invoice) heading() string {
	return i.customer.Greeting() + ","
}

func NewInvoice(customer Customer, cents int64) Invoice {
	return Invoice{customer: customer, cents: cents}
}

func (i Invoice) Render() string {
	return i.heading() + " you owe " + formatCents(i.cents)
}

const (
	currencySymbol = "$"
	centsPerUnit   = 100
)
