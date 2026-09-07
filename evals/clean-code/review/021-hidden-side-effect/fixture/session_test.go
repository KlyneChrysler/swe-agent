package fixture

import (
	"testing"
	"time"
)

var opened = time.Date(2024, time.March, 1, 9, 0, 0, 0, time.UTC)

func TestIsActiveWithinTimeout(t *testing.T) {
	session := NewSession(opened)
	if !session.IsActive(opened.Add(10 * time.Minute)) {
		t.Fatal("IsActive = false within the idle timeout")
	}
}

func TestIsActiveAfterTimeout(t *testing.T) {
	session := NewSession(opened)
	if session.IsActive(opened.Add(idleTimeout + time.Minute)) {
		t.Fatal("IsActive = true after the idle timeout")
	}
}
