package fixture

import (
	"fmt"
	"os"
	"path/filepath"
)

// Order is a purchase of one item, priced in cents.
type Order struct {
	ID    string
	Email string
	Cents int64
}

const (
	percent     = 100
	receiptMode = 0o644
)

type OrderManager struct {
	taxPercent int64
	receiptDir string
	sender     string
}

func NewOrderManager(taxPercent int64, receiptDir, sender string) OrderManager {
	return OrderManager{taxPercent: taxPercent, receiptDir: receiptDir, sender: sender}
}

func (m OrderManager) Price(order Order) int64 {
	return order.Cents + order.Cents*m.taxPercent/percent
}

func (m OrderManager) Save(order Order) error {
	path := filepath.Join(m.receiptDir, order.ID+".txt")
	body := fmt.Sprintf("%s %d", order.ID, order.Cents)
	return os.WriteFile(path, []byte(body), receiptMode)
}

func (m OrderManager) Notification(order Order) string {
	return fmt.Sprintf("From: %s\nTo: %s\n\nOrder %s received", m.sender, order.Email, order.ID)
}
