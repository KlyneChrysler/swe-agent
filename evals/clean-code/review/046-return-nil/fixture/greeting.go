package fixture

const greeting = "Hello"

type Customer struct {
	ID   string
	Name string
}

type Directory struct {
	byID map[string]*Customer
}

func NewDirectory(byID map[string]*Customer) Directory {
	return Directory{byID: byID}
}

func (d Directory) Find(id string) *Customer {
	return d.byID[id]
}

func (d Directory) GreetingFor(id string) string {
	customer := d.Find(id)
	if customer == nil {
		return greeting
	}
	return greeting + " " + customer.Name
}
