package fixture

import (
	"testing"
	"time"
)

func TestIssueSetsSubject(t *testing.T) {
	token := Issue("ada", time.Hour)

	if token.Subject != "ada" {
		t.Errorf("Subject = %q, want %q", token.Subject, "ada")
	}
}

func TestFreshTokenIsNotExpired(t *testing.T) {
	token := Issue("ada", time.Hour)

	if token.IsExpired() {
		t.Error("IsExpired = true for a token issued with an hour to live")
	}
}
