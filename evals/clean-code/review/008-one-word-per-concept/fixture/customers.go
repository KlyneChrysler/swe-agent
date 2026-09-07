package fixture

type Customer struct {
	ID   string
	Name string
}

type CustomerBook struct {
	byID map[string]Customer
}

func (b CustomerBook) FetchCustomer(id string) (Customer, bool) {
	customer, ok := b.byID[id]
	return customer, ok
}
