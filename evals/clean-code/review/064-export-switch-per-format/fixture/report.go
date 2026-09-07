package fixture

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Format names a serialisation the report can be exported as.
type Format string

const (
	CSV  Format = "csv"
	JSON Format = "json"
)

type Report struct {
	Rows [][]string
}

func (r Report) Export(format Format) (string, error) {
	switch format {
	case CSV:
		return r.csv(), nil
	case JSON:
		return r.json()
	}
	return "", fmt.Errorf("export report: unsupported format %q", format)
}

func (r Report) FileName(base string, format Format) string {
	switch format {
	case CSV:
		return base + ".csv"
	case JSON:
		return base + ".json"
	}
	return base
}

func (r Report) csv() string {
	lines := make([]string, 0, len(r.Rows))
	for _, row := range r.Rows {
		lines = append(lines, strings.Join(row, ","))
	}
	return strings.Join(lines, "\n")
}

func (r Report) json() (string, error) {
	encoded, err := json.Marshal(r.Rows)
	if err != nil {
		return "", fmt.Errorf("encode report rows: %w", err)
	}
	return string(encoded), nil
}
