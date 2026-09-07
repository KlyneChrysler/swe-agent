package fixture

import "time"

const dayLayout = "2006-01-02"

type Shift struct {
	Start time.Time
	End   time.Time
}

func (s Shift) overlaps(other Shift) bool {
	return s.Start.Before(other.End) && other.Start.Before(s.End)
}

type Roster struct {
	s []Shift
}

func (r Roster) TimeScheduled() time.Duration {
	var total time.Duration
	for _, shift := range r.s {
		total += shift.End.Sub(shift.Start)
	}
	return total
}

func (r Roster) Overlaps(candidate Shift) bool {
	for _, shift := range r.s {
		if shift.overlaps(candidate) {
			return true
		}
	}
	return false
}

func (r Roster) ByDay() map[string][]Shift {
	grouped := map[string][]Shift{}
	for _, shift := range r.s {
		key := shift.Start.Format(dayLayout)
		grouped[key] = append(grouped[key], shift)
	}
	return grouped
}
