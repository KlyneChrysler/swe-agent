package fixture

import "time"

const idleTimeout = 30 * time.Minute

type Session struct {
	lastSeen time.Time
}

func NewSession(now time.Time) *Session {
	return &Session{lastSeen: now}
}

func (s *Session) IsActive(now time.Time) bool {
	if now.Sub(s.lastSeen) > idleTimeout {
		return false
	}
	s.lastSeen = now
	return true
}
