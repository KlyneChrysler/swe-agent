package fixture

import (
	"errors"
	"testing"
)

func TestConfirmOrderReportsMissingVendorClient(t *testing.T) {
	order := Order{ID: "A1", Phone: "+15550100"}

	err := ConfirmOrder(order)

	if !errors.Is(err, errNoVendorClient) {
		t.Errorf("ConfirmOrder error = %v, want %v", err, errNoVendorClient)
	}
}

func TestConfirmationTextNamesTheOrder(t *testing.T) {
	text := confirmationText(Order{ID: "A1"})

	if text != "Order A1 confirmed" {
		t.Errorf("confirmationText = %q, want %q", text, "Order A1 confirmed")
	}
}
