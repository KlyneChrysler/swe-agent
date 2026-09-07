package fixture

import (
	"testing"
	"time"
)

func TestIsExpiredAfterDeadline(t *testing.T) {
	deadline := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	session := NewSession("ada", deadline)

	if !session.IsExpired(deadline.Add(time.Second)) {
		t.Error("IsExpired = false one second after the deadline")
	}
}

func TestRemainingShrinksAsCursorAdvances(t *testing.T) {
	session := NewSession("ada", time.Time{})
	session.Append([]byte("hello"))

	session.Advance(2)

	if session.Remaining() != 3 {
		t.Errorf("Remaining = %d, want 3", session.Remaining())
	}
}

func TestAdvanceStopsAtEndOfBuffer(t *testing.T) {
	session := NewSession("ada", time.Time{})
	session.Append([]byte("hi"))

	session.Advance(10)

	if session.Remaining() != 0 {
		t.Errorf("Remaining = %d, want 0", session.Remaining())
	}
}
