package fixture

// Entry is one posting against an account, in cents; debits are negative.
type Entry struct {
	Account string
	Cents   int64
}

type Ledger struct {
	entries []Entry
}

func NewLedger() *Ledger {
	return &Ledger{}
}

func (l *Ledger) Post(entry Entry) {
	l.entries = append(l.entries, entry)
}

func (l *Ledger) Balance(account string) int64 {
	var balance int64
	for _, entry := range l.entries {
		if entry.Account == account {
			balance += entry.Cents
		}
	}
	return balance
}
