package fixture

import "testing"

type spyStock struct {
	reserved []string
}

func (s *spyStock) Reserve(skus []string) error {
	s.reserved = skus
	return nil
}

type spyCourier struct {
	label string
}

func (c *spyCourier) Dispatch(orderID, label string) error {
	c.label = label
	return nil
}

func TestFulfillReservesThenDispatchesLabel(t *testing.T) {
	stock := &spyStock{}
	courier := &spyCourier{}
	order := Order{ID: "o1", Recipient: " ada ", AddressLines: []string{"12 Main St", "Cebu"}, SKUs: []string{"EL-1"}}
	if err := NewFulfiller(stock, courier).Fulfill(order); err != nil {
		t.Fatalf("Fulfill: %v", err)
	}
	if len(stock.reserved) != 1 || courier.label != "ADA\n12 Main St\nCebu" {
		t.Fatalf("reserved %v, label %q", stock.reserved, courier.label)
	}
}
