package fixture

import "testing"

func TestListenAddressJoinsHostAndPort(t *testing.T) {
	if got, ok := ListenAddress("localhost:8080"); !ok || got != "localhost:8080" {
		t.Fatalf("ListenAddress = %q, %v", got, ok)
	}
}

func TestListenAddressRejectsBadInput(t *testing.T) {
	for _, text := range []string{"localhost", ":8080", "localhost:http", "localhost:70000"} {
		if _, ok := ListenAddress(text); ok {
			t.Errorf("ListenAddress(%q) accepted", text)
		}
	}
}
