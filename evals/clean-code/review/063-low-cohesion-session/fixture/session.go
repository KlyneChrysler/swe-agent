package fixture

import "time"

// Session pairs an authenticated user with the bytes of their in-progress upload.
type Session struct {
	userID    string
	expiresAt time.Time
	buffer    []byte
	cursor    int
}

func NewSession(userID string, expiresAt time.Time) *Session {
	return &Session{userID: userID, expiresAt: expiresAt}
}

func (s *Session) Owner() string {
	return s.userID
}

func (s *Session) IsExpired(now time.Time) bool {
	return now.After(s.expiresAt)
}

func (s *Session) Append(chunk []byte) {
	s.buffer = append(s.buffer, chunk...)
}

func (s *Session) Remaining() int {
	return len(s.buffer) - s.clamp(s.cursor, 0, len(s.buffer))
}

func (s *Session) Advance(count int) {
	s.cursor = s.clamp(s.cursor+count, 0, len(s.buffer))
}

func (s *Session) clamp(value, low, high int) int {
	return min(max(value, low), high)
}
