package fixture

import "testing"

func TestSaveCopiesBytes(t *testing.T) {
	storage := NewMemoryStorage()
	payload := []byte("hello")
	if err := storage.Save("greeting", payload); err != nil {
		t.Fatalf("Save: %v", err)
	}
	payload[0] = 'j'
	if string(storage.m_entries["greeting"]) != "hello" {
		t.Fatal("Save stored a slice that aliases the caller's bytes")
	}
	if got := storage.Count(); got != 1 {
		t.Fatalf("Count = %d, want 1", got)
	}
}
