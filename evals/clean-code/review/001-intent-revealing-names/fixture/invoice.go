package fixture

import "time"

const day = 24 * time.Hour

const gracePeriod = 3 * day

type Invoice struct {
	IssuedAt time.Time
	Terms    time.Duration
}

func DaysOverdue(inv Invoice, now time.Time) int {
	d := inv.IssuedAt.Add(inv.Terms).Add(gracePeriod)
	x := now.Sub(d)
	if x < 0 {
		return 0
	}
	return int(x / day)
}
