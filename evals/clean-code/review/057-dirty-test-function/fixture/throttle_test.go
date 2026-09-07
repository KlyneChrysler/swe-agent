package fixture

import (
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	c := time.Now()
	th := NewThrottle(Policy{Limit: 2, Window: time.Second}, func() time.Time { return c })
	th.Allow()
	th.Allow()
	x := th.Allow()
	if x != false {
		t.Errorf("got %v", x)
	}
	c2 := time.Now()
	th2 := NewThrottle(Policy{Limit: 2, Window: time.Second}, func() time.Time { return c2 })
	th2.Allow()
	y := th2.Allow()
	if y != true {
		t.Errorf("got %v", y)
	}
	c3 := time.Now()
	th3 := NewThrottle(Policy{Limit: 1, Window: time.Second}, func() time.Time { return c3 })
	th3.Allow()
	z := th3.Allow()
	if z != false {
		t.Errorf("got %v", z)
	}
}

func TestThrottleAllowsAgainAfterWindow(t *testing.T) {
	throttle := NewThrottle(Policy{Limit: 1, Window: time.Second}, time.Now)
	throttle.Allow()

	time.Sleep(1100 * time.Millisecond)

	if !throttle.Allow() {
		t.Error("Allow refused a call after the window had passed")
	}
}
