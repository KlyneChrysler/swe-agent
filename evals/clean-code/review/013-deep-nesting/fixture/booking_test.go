package fixture

import (
	"testing"
	"time"
)

func bookingAt(room string, startHour, endHour int) Booking {
	start := time.Date(2024, time.March, 1, startHour, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.March, 1, endHour, 0, 0, 0, time.UTC)
	return Booking{Room: room, Start: start, End: end}
}

func TestFirstConflictFindsOverlapInSameRoom(t *testing.T) {
	rooms := []Room{{Name: "A", Bookings: []Booking{bookingAt("A", 9, 11)}}}
	conflict, found := FirstConflict(rooms, bookingAt("A", 10, 12))
	if !found || conflict.Start.Hour() != 9 {
		t.Fatalf("FirstConflict = %+v, %v; want the 9:00 booking", conflict, found)
	}
}

func TestFirstConflictIgnoresOtherRooms(t *testing.T) {
	rooms := []Room{{Name: "A", Bookings: []Booking{bookingAt("A", 9, 11)}}}
	if _, found := FirstConflict(rooms, bookingAt("B", 10, 12)); found {
		t.Fatal("FirstConflict found a clash in a different room")
	}
}
