package fixture

const (
	digitalDiscountPercent = 10
	subscriptionMonths     = 12
	percentDenominator     = 100
)

func PriceCents(item Item) int {
	switch item.Kind {
	case Digital:
		return item.BaseCents - item.BaseCents*digitalDiscountPercent/percentDenominator
	case Subscription:
		return item.BaseCents * subscriptionMonths
	default:
		return item.BaseCents
	}
}
