package fixture

import "testing"

func TestCollectUnpaidAppendsOnlyUnpaid(t *testing.T) {
	invoices := []Invoice{{Number: "1", Paid: true}, {Number: "2"}, {Number: "3"}}
	var unpaid []Invoice
	CollectUnpaid(invoices, &unpaid)
	if len(unpaid) != 2 || unpaid[0].Number != "2" || unpaid[1].Number != "3" {
		t.Fatalf("unpaid = %+v, want invoices 2 and 3", unpaid)
	}
}
