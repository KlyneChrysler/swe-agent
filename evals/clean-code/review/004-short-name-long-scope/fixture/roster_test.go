package fixture

import (
	"testing"
	"time"
)

func shiftAt(day int, startHour, endHour int) Shift {
	start := time.Date(2024, time.March, day, startHour, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.March, day, endHour, 0, 0, 0, time.UTC)
	return Shift{Start: start, End: end}
}

func TestTimeScheduledSumsShifts(t *testing.T) {
	roster := Roster{s: []Shift{shiftAt(1, 9, 12), shiftAt(2, 13, 15)}}
	if got := roster.TimeScheduled(); got != 5*time.Hour {
		t.Fatalf("TimeScheduled = %v, want 5h", got)
	}
}

func TestOverlapsDetectsClash(t *testing.T) {
	roster := Roster{s: []Shift{shiftAt(1, 9, 12)}}
	if !roster.Overlaps(shiftAt(1, 11, 14)) {
		t.Fatal("Overlaps = false, want true")
	}
}

func TestByDayGroupsByDate(t *testing.T) {
	roster := Roster{s: []Shift{shiftAt(1, 9, 12), shiftAt(1, 13, 15), shiftAt(2, 9, 12)}}
	if got := len(roster.ByDay()["2024-03-01"]); got != 2 {
		t.Fatalf("shifts on 2024-03-01 = %d, want 2", got)
	}
}
