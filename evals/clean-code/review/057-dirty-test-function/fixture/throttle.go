package fixture

import "time"

// Policy bounds how many calls a Throttle admits within one window.
type Policy struct {
	Limit  int
	Window time.Duration
}

type Throttle struct {
	policy Policy
	now    func() time.Time
	calls  []time.Time
}

func NewThrottle(policy Policy, now func() time.Time) *Throttle {
	return &Throttle{policy: policy, now: now}
}

func (t *Throttle) Allow() bool {
	t.dropExpired()
	if len(t.calls) >= t.policy.Limit {
		return false
	}
	t.calls = append(t.calls, t.now())
	return true
}

func (t *Throttle) dropExpired() {
	cutoff := t.now().Add(-t.policy.Window)
	kept := t.calls[:0]
	for _, call := range t.calls {
		if call.After(cutoff) {
			kept = append(kept, call)
		}
	}
	t.calls = kept
}
