package fixture

type Priority int

const (
	Low Priority = iota
	Normal
	Urgent
)

const (
	urgentWithinHours = 4
	normalWithinHours = 48
)

func PriorityFor(hoursUntilDue int) Priority {
	if hoursUntilDue <= urgentWithinHours {
		return Urgent
	}
	if hoursUntilDue <= normalWithinHours {
		return Normal
	}
	return Low
}
