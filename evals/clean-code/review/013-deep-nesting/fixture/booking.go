package fixture

import "time"

type Booking struct {
	Room  string
	Start time.Time
	End   time.Time
}

func (b Booking) overlaps(other Booking) bool {
	return b.Start.Before(other.End) && other.Start.Before(b.End)
}

type Room struct {
	Name     string
	Bookings []Booking
}

func FirstConflict(rooms []Room, requested Booking) (Booking, bool) {
	for _, room := range rooms {
		if room.Name == requested.Room {
			for _, existing := range room.Bookings {
				if existing.overlaps(requested) {
					return existing, true
				}
			}
		}
	}
	return Booking{}, false
}
