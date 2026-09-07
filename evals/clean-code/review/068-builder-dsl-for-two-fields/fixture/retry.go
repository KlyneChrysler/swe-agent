package fixture

import "time"

// RetryPolicy bounds how often and how quickly an operation is retried.
type RetryPolicy struct {
	Attempts int
	Delay    time.Duration
}

type RetryPolicyBuilder struct {
	policy RetryPolicy
}

func NewRetryPolicyBuilder() RetryPolicyBuilder {
	return RetryPolicyBuilder{}
}

func (b RetryPolicyBuilder) WithAttempts(attempts int) RetryPolicyBuilder {
	b.policy.Attempts = attempts
	return b
}

func (b RetryPolicyBuilder) WithDelay(delay time.Duration) RetryPolicyBuilder {
	b.policy.Delay = delay
	return b
}

func (b RetryPolicyBuilder) Build() RetryPolicy {
	return b.policy
}

type Backoff struct {
	base time.Duration
}

func NewBackoff(base time.Duration) Backoff {
	return Backoff{base: base}
}

func (this Backoff) GetDelay(attempt int) time.Duration {
	return this.base << attempt
}
