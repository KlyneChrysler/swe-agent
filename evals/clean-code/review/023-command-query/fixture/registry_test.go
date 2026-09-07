package fixture

import "testing"

func TestSetReportsWhetherKeyExisted(t *testing.T) {
	registry := NewRegistry()
	if existed := registry.Set("region", "ph"); existed {
		t.Fatal("first Set reported an existing key")
	}
	if existed := registry.Set("region", "sg"); !existed {
		t.Fatal("second Set did not report the existing key")
	}
}

func TestGetReturnsLatestValue(t *testing.T) {
	registry := NewRegistry()
	registry.Set("region", "ph")
	if value, ok := registry.Get("region"); !ok || value != "ph" {
		t.Fatalf("Get = %q, %v; want ph, true", value, ok)
	}
}
