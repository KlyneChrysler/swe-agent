package fixture

const homeCountry = "PH"

type Address struct {
	country string
}

func (a Address) Country() string {
	return a.country
}

type Profile struct {
	billing Address
}

func (p Profile) Billing() Address {
	return p.billing
}

type Customer struct {
	profile Profile
}

func (c Customer) Profile() Profile {
	return c.profile
}

type Order struct {
	customer Customer
}

func (o Order) Customer() Customer {
	return o.customer
}

func IsDomestic(order Order) bool {
	return order.Customer().Profile().Billing().Country() == homeCountry
}
