package fixture

import (
	"encoding/json"
	"fmt"
)

func LoadSettings(document []byte) (map[string]any, error) {
	var settings map[string]any
	if err := json.Unmarshal(document, &settings); err != nil {
		return nil, fmt.Errorf("decode settings: %w", err)
	}
	return settings, nil
}
