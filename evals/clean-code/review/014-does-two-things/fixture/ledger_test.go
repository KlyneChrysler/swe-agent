package fixture

import "testing"

func TestParseAndRecordAppendsEntry(t *testing.T) {
	var ledger Ledger
	if err := ledger.ParseAndRecord("cash,1250"); err != nil {
		t.Fatalf("ParseAndRecord: %v", err)
	}
	want := Entry{Account: "cash", Cents: 1250}
	if len(ledger.entries) != 1 || ledger.entries[0] != want {
		t.Fatalf("entries = %+v, want [%+v]", ledger.entries, want)
	}
}

func TestParseAndRecordRejectsBadLines(t *testing.T) {
	var ledger Ledger
	for _, line := range []string{"cash", "cash,lots", "a,b,c"} {
		if err := ledger.ParseAndRecord(line); err == nil {
			t.Errorf("ParseAndRecord(%q) accepted a malformed line", line)
		}
	}
}
