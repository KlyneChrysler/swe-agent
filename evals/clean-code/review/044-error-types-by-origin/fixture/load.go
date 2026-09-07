package fixture

import (
	"encoding/json"
	"errors"
	"os"
)

type Settings struct {
	Port int `json:"port"`
}

func LoadSettings(path string) (Settings, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, DiskReadError{Err: err}
	}
	var settings Settings
	if err := json.Unmarshal(data, &settings); err != nil {
		return Settings{}, DecodeError{Err: err}
	}
	return settings, nil
}

func LoadOrDefault(path string, fallback Settings) Settings {
	settings, err := LoadSettings(path)
	var diskErr DiskReadError
	var decodeErr DecodeError
	if errors.As(err, &diskErr) || errors.As(err, &decodeErr) {
		return fallback
	}
	return settings
}
