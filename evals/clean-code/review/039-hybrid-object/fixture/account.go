package fixture

import "fmt"

type Account struct {
	BalanceCents   int
	OverdraftCents int
}

func (a *Account) Withdraw(cents int) error {
	if cents > a.BalanceCents+a.OverdraftCents {
		return fmt.Errorf("withdraw %d cents: exceeds balance %d plus overdraft %d",
			cents, a.BalanceCents, a.OverdraftCents)
	}
	a.BalanceCents -= cents
	return nil
}
