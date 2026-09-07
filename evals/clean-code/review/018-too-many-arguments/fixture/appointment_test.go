package fixture

import (
	"testing"
	"time"
)

func TestEndAddsDuration(t *testing.T) {
	start := time.Date(2024, time.March, 1, 9, 0, 0, 0, time.UTC)
	appointment := NewAppointment("Ada", "Dr. Cruz", start, 45, "3B")
	if got := appointment.End(); !got.Equal(start.Add(45 * time.Minute)) {
		t.Fatalf("End = %v, want 09:45", got)
	}
}
