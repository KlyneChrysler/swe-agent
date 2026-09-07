package fixture

import "time"

type Appointment struct {
	Patient string
	Doctor  string
	Start   time.Time
	Minutes int
	Room    string
}

func NewAppointment(patient, doctor string, start time.Time, minutes int, room string) Appointment {
	return Appointment{
		Patient: patient,
		Doctor:  doctor,
		Start:   start,
		Minutes: minutes,
		Room:    room,
	}
}

func (a Appointment) End() time.Time {
	return a.Start.Add(time.Duration(a.Minutes) * time.Minute)
}
