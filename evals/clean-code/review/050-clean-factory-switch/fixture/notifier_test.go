package fixture

import "testing"

func TestNewNotifierBuildsEmail(t *testing.T) {
	notifier, err := NewNotifier(Email, "ada@example.com")
	if err != nil {
		t.Fatalf("NewNotifier: %v", err)
	}
	if _, ok := notifier.(emailNotifier); !ok {
		t.Fatalf("NewNotifier built %T, want emailNotifier", notifier)
	}
}

func TestNewNotifierBuildsSMS(t *testing.T) {
	notifier, err := NewNotifier(SMS, "+639170000000")
	if err != nil {
		t.Fatalf("NewNotifier: %v", err)
	}
	if _, ok := notifier.(smsNotifier); !ok {
		t.Fatalf("NewNotifier built %T, want smsNotifier", notifier)
	}
}

func TestNewNotifierRejectsUnknownChannel(t *testing.T) {
	if _, err := NewNotifier("pigeon", "roof"); err == nil {
		t.Fatal("NewNotifier accepted an unknown channel")
	}
}
