package fixture

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPriceAddsTax(t *testing.T) {
	manager := NewOrderManager(20, t.TempDir(), "shop@example.com")

	price := manager.Price(Order{ID: "A1", Cents: 1000})

	if price != 1200 {
		t.Errorf("Price = %d, want 1200", price)
	}
}

func TestSaveWritesReceiptFile(t *testing.T) {
	dir := t.TempDir()
	manager := NewOrderManager(20, dir, "shop@example.com")

	if err := manager.Save(Order{ID: "A1", Cents: 1000}); err != nil {
		t.Fatalf("Save returned error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, "A1.txt")); err != nil {
		t.Errorf("receipt file missing: %v", err)
	}
}

func TestNotificationAddressesCustomer(t *testing.T) {
	manager := NewOrderManager(20, t.TempDir(), "shop@example.com")

	notification := manager.Notification(Order{ID: "A1", Email: "ada@example.com"})

	if notification != "From: shop@example.com\nTo: ada@example.com\n\nOrder A1 received" {
		t.Errorf("Notification = %q", notification)
	}
}
