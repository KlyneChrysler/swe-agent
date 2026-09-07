package fixture

import "testing"

var shared = NewLedger()

func TestPostRecordsEntry(t *testing.T) {
	shared.Post(Entry{Account: "cash", Cents: 500})

	if shared.Balance("cash") != 500 {
		t.Errorf("Balance(cash) = %d, want 500", shared.Balance("cash"))
	}
}

func TestBalanceSumsEntries(t *testing.T) {
	shared.Post(Entry{Account: "cash", Cents: 250})

	if shared.Balance("cash") != 750 {
		t.Errorf("Balance(cash) = %d, want 750", shared.Balance("cash"))
	}
}

func TestBalanceIgnoresOtherAccounts(t *testing.T) {
	ledger := NewLedger()
	ledger.Post(Entry{Account: "cash", Cents: 100})

	if ledger.Balance("bank") != 0 {
		t.Errorf("Balance(bank) = %d, want 0", ledger.Balance("bank"))
	}
}
