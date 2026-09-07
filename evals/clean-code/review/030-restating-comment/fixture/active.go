package fixture

type Member struct {
	Name   string
	Active bool
}

func CountActive(members []Member) int {
	// start the count at zero
	count := 0
	// loop over every member
	for _, member := range members {
		// only count active members
		if member.Active {
			// increment the count
			count++
		}
	}
	// return the count
	return count
}
