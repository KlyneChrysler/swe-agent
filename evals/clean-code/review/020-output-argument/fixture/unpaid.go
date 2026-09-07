package fixture

type Invoice struct {
	Number string
	Paid   bool
}

func CollectUnpaid(invoices []Invoice, unpaid *[]Invoice) {
	for _, invoice := range invoices {
		if !invoice.Paid {
			*unpaid = append(*unpaid, invoice)
		}
	}
}
