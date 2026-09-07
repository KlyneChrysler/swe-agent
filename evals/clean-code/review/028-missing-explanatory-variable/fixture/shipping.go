package fixture

const freeShippingThresholdCents = 5000

type Country string

type Order struct {
	SubtotalCents int
	DiscountCents int
	Destination   Country
	Merchant      Country
	LineCount     int
}

func QualifiesForFreeShipping(order Order) bool {
	return order.SubtotalCents-order.DiscountCents >= freeShippingThresholdCents &&
		order.Destination == order.Merchant && order.LineCount > 0
}
