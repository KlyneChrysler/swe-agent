package fixture

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func writeSettings(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "settings.json")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write settings: %v", err)
	}
	return path
}

func TestLoadSettingsReadsPort(t *testing.T) {
	settings, err := LoadSettings(writeSettings(t, `{"port": 8080}`))
	if err != nil || settings.Port != 8080 {
		t.Fatalf("LoadSettings = %+v, %v", settings, err)
	}
}

func TestLoadSettingsClassifiesFailures(t *testing.T) {
	var decodeErr DecodeError
	if _, err := LoadSettings(writeSettings(t, "not json")); !errors.As(err, &decodeErr) {
		t.Fatalf("bad json gave %v, want DecodeError", err)
	}
	var diskErr DiskReadError
	if _, err := LoadSettings(filepath.Join(t.TempDir(), "missing.json")); !errors.As(err, &diskErr) {
		t.Fatalf("missing file gave %v, want DiskReadError", err)
	}
}

func TestLoadOrDefaultFallsBack(t *testing.T) {
	fallback := Settings{Port: 9000}
	if got := LoadOrDefault(filepath.Join(t.TempDir(), "missing.json"), fallback); got != fallback {
		t.Fatalf("LoadOrDefault = %+v, want fallback", got)
	}
}
