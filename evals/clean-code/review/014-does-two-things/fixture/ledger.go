package fixture

import (
	"fmt"
	"strconv"
	"strings"
)

const entryFieldCount = 2

type Entry struct {
	Account string
	Cents   int
}

type Ledger struct {
	entries []Entry
}

func (l *Ledger) ParseAndRecord(line string) error {
	fields := strings.Split(line, ",")
	if len(fields) != entryFieldCount {
		return fmt.Errorf("parse ledger line %q: want %d fields, got %d", line, entryFieldCount, len(fields))
	}
	cents, err := strconv.Atoi(fields[1])
	if err != nil {
		return fmt.Errorf("parse ledger line %q: amount: %w", line, err)
	}
	l.entries = append(l.entries, Entry{Account: fields[0], Cents: cents})
	return nil
}
