package fixture

type Order struct {
	Number     int
	CustomerID string
}

type OrderLog struct {
	entries []Order
}

func (l OrderLog) RetrieveOrder(number int) (Order, bool) {
	for _, order := range l.entries {
		if order.Number == number {
			return order, true
		}
	}
	return Order{}, false
}
