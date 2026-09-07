package fixture

type Customer struct {
	ID    string
	Name  string
	Email string
}

type Directory struct {
	customers map[string]Customer
}

func NewDirectory(customers []Customer) Directory {
	byID := make(map[string]Customer, len(customers))
	for _, customer := range customers {
		byID[customer.ID] = customer
	}
	return Directory{customers: byID}
}

func (d Directory) Customer(id string) Customer {
	return d.customers[id]
}

func (d Directory) CustomerInfo(id string) string {
	return d.customers[id].Name
}

func (d Directory) CustomerData(id string) string {
	return d.customers[id].Email
}
