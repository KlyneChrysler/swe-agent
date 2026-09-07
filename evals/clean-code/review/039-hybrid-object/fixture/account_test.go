package fixture

import "testing"

func TestWithdrawWithinOverdraft(t *testing.T) {
	account := Account{BalanceCents: 1000, OverdraftCents: 500}
	if err := account.Withdraw(1200); err != nil {
		t.Fatalf("Withdraw: %v", err)
	}
	if account.BalanceCents != -200 {
		t.Fatalf("BalanceCents = %d, want -200", account.BalanceCents)
	}
}

func TestWithdrawBeyondOverdraftFails(t *testing.T) {
	account := Account{BalanceCents: 1000, OverdraftCents: 500}
	if err := account.Withdraw(1600); err == nil {
		t.Fatal("Withdraw exceeded the overdraft without error")
	}
}
