package fixture

import "time"

// Token grants access until ExpiresAt.
type Token struct {
	Subject   string
	ExpiresAt time.Time
}

func Issue(subject string, ttl time.Duration) Token {
	return Token{Subject: subject, ExpiresAt: time.Now().Add(ttl)}
}

func (t Token) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}
