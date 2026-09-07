package fixture

import (
	"fmt"
	"os"
)

const receiptDir = "/var/receipts"

// Order is a paid purchase ready for its receipt.
type Order struct {
	ID    string
	Email string
	Cents int64
}

type Checkout struct{}

func (Checkout) Complete(order Order) error {
	receipts := NewReceipts(receiptDir)
	mailer := NewWriterMailer(os.Stdout)
	body := receiptText(order)
	if err := receipts.Store(order.ID, body); err != nil {
		return err
	}
	return mailer.Send(order.Email, body)
}

func receiptText(order Order) string {
	return fmt.Sprintf("Order %s: %d cents received", order.ID, order.Cents)
}
