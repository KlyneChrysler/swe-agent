package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type gradeRow struct {
	Mode           string        `json:"mode"`
	Scenario       string        `json:"scenario"`
	Pass           bool          `json:"pass"`
	Empty          bool          `json:"empty_output"`
	Misses         []finding     `json:"misses"`
	FalsePositives []finding     `json:"false_positives"`
	Results        []checkResult `json:"results"`
}

func summarise(resultsDir string) error {
	rows, err := loadGrades(resultsDir)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return fmt.Errorf("no grade.json files under %s", resultsDir)
	}
	fmt.Print(renderSummary(rows))
	return nil
}

func loadGrades(resultsDir string) ([]gradeRow, error) {
	paths, err := filepath.Glob(filepath.Join(resultsDir, "*", "grade.json"))
	if err != nil {
		return nil, fmt.Errorf("glob %s: %w", resultsDir, err)
	}
	rows := []gradeRow{}
	for _, path := range paths {
		row, err := loadGrade(path)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func loadGrade(path string) (gradeRow, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return gradeRow{}, fmt.Errorf("read %s: %w", path, err)
	}
	var row gradeRow
	if err := json.Unmarshal(raw, &row); err != nil {
		return gradeRow{}, fmt.Errorf("parse %s: %w", path, err)
	}
	return row, nil
}

func renderSummary(rows []gradeRow) string {
	var out strings.Builder
	fmt.Fprintf(&out, "# %s: %d/%d passed\n\n", rows[0].Mode, passCount(rows), len(rows))
	out.WriteString(renderFailureTable(rows))
	out.WriteString(renderFailedScenarios(rows))
	return out.String()
}

func passCount(rows []gradeRow) int {
	count := 0
	for _, row := range rows {
		if row.Pass {
			count++
		}
	}
	return count
}

func renderFailureTable(rows []gradeRow) string {
	counts := failureCounts(rows)
	if len(counts) == 0 {
		return ""
	}
	var out strings.Builder
	out.WriteString("| failure | scenarios |\n|---|---|\n")
	for _, key := range sortedKeys(counts) {
		fmt.Fprintf(&out, "| %s | %d |\n", key, counts[key])
	}
	out.WriteString("\n")
	return out.String()
}

func failureCounts(rows []gradeRow) map[string]int {
	counts := map[string]int{}
	for _, row := range rows {
		for _, label := range row.failureLabels() {
			counts[label]++
		}
	}
	return counts
}

func (r gradeRow) failureLabels() []string {
	labels := []string{}
	if r.Empty {
		labels = append(labels, "empty output")
	}
	for _, miss := range r.Misses {
		labels = append(labels, fmt.Sprintf("missed rule %d %s", miss.Rule, miss.Entry))
	}
	for _, falsePositive := range r.FalsePositives {
		labels = append(labels, fmt.Sprintf("false positive rule %d", falsePositive.Rule))
	}
	for _, result := range r.Results {
		if !result.Pass {
			labels = append(labels, "failed "+result.Check)
		}
	}
	return labels
}

func sortedKeys(counts map[string]int) []string {
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]] || counts[keys[i]] == counts[keys[j]] && keys[i] < keys[j]
	})
	return keys
}

func renderFailedScenarios(rows []gradeRow) string {
	var out strings.Builder
	for _, row := range rows {
		if !row.Pass {
			fmt.Fprintf(&out, "- %s: %s\n", row.Scenario, strings.Join(row.failureLabels(), "; "))
		}
	}
	return out.String()
}
