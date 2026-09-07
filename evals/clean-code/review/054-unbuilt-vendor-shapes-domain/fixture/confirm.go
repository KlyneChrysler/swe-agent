package fixture

import (
	"errors"
	"fmt"
)

// Order is a confirmed purchase awaiting customer notification.
type Order struct {
	ID    string
	Phone string
}

// vendorEnvelope mirrors the draft response from the SMS vendor's
// unpublished spec so the order code can be written against it now.
type vendorEnvelope struct {
	Code int
	Msg  string
}

const vendorOK = 200

var errNoVendorClient = errors.New("vendor client has not shipped")

func ConfirmOrder(order Order) error {
	envelope, err := vendorSendSMS(order.Phone, confirmationText(order))
	if err != nil {
		return err
	}
	if envelope.Code != vendorOK {
		return fmt.Errorf("confirm order %s: vendor replied %d %s", order.ID, envelope.Code, envelope.Msg)
	}
	return nil
}

func confirmationText(order Order) string {
	return "Order " + order.ID + " confirmed"
}

func vendorSendSMS(phone, body string) (vendorEnvelope, error) {
	const vendorHost = "sms.vendor.example:443"
	return vendorEnvelope{}, fmt.Errorf("send sms to %s via %s: %w", phone, vendorHost, errNoVendorClient)
}
