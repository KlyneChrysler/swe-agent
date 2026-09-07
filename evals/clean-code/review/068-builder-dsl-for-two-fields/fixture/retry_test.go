package fixture

import (
	"testing"
	"time"
)

func TestBuilderSetsAttemptsAndDelay(t *testing.T) {
	policy := NewRetryPolicyBuilder().WithAttempts(3).WithDelay(time.Second).Build()

	if policy != (RetryPolicy{Attempts: 3, Delay: time.Second}) {
		t.Errorf("Build = %+v", policy)
	}
}

func TestBackoffDoublesEachAttempt(t *testing.T) {
	backoff := NewBackoff(time.Second)

	if backoff.GetDelay(2) != 4*time.Second {
		t.Errorf("GetDelay(2) = %v, want 4s", backoff.GetDelay(2))
	}
}
