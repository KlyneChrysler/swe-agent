package fixture

type Account struct {
	balanceCents int
	owner        string
}

func (a *Account) GetBalanceCents() int {
	return a.balanceCents
}

func (a *Account) SetBalanceCents(cents int) {
	a.balanceCents = cents
}

func (a *Account) GetOwner() string {
	return a.owner
}

func (a *Account) SetOwner(owner string) {
	a.owner = owner
}
