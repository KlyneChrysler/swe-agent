package fixture

import "testing"

func TestDirectoryLooksUpByID(t *testing.T) {
	ada := Customer{ID: "c1", Name: "Ada", Email: "ada@example.com"}
	directory := NewDirectory([]Customer{ada})
	if got := directory.Customer("c1"); got != ada {
		t.Fatalf("Customer = %+v, want %+v", got, ada)
	}
	if got := directory.CustomerInfo("c1"); got != "Ada" {
		t.Fatalf("CustomerInfo = %q, want Ada", got)
	}
	if got := directory.CustomerData("c1"); got != "ada@example.com" {
		t.Fatalf("CustomerData = %q, want ada@example.com", got)
	}
}
