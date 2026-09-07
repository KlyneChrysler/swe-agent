package fixture

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestStoreWritesReceiptUnderDir(t *testing.T) {
	dir := t.TempDir()
	receipts := NewReceipts(dir)

	if err := receipts.Store("A1", "body"); err != nil {
		t.Fatalf("Store returned error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, "A1.txt")); err != nil {
		t.Errorf("receipt file missing: %v", err)
	}
}

func TestWriterMailerAddressesRecipient(t *testing.T) {
	var out bytes.Buffer
	mailer := NewWriterMailer(&out)

	if err := mailer.Send("ada@example.com", "hello"); err != nil {
		t.Fatalf("Send returned error: %v", err)
	}

	if out.String() != "To: ada@example.com\n\nhello\n" {
		t.Errorf("Send wrote %q", out.String())
	}
}

func TestReceiptTextStatesAmount(t *testing.T) {
	text := receiptText(Order{ID: "A1", Cents: 500})

	if text != "Order A1: 500 cents received" {
		t.Errorf("receiptText = %q", text)
	}
}
