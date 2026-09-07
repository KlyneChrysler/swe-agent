package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type reviewScenario struct {
	Title       string    `json:"title"`
	MustFind    []finding `json:"must_find"`
	MustNotFind []finding `json:"must_not_find"`
}

type finding struct {
	Rule   int    `json:"rule"`
	Entry  string `json:"entry,omitempty"`
	File   string `json:"file,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Why    string `json:"why"`
}

// ReviewGrade is the score of one review transcript.
type ReviewGrade struct {
	Mode           string    `json:"mode"`
	Scenario       string    `json:"scenario"`
	Title          string    `json:"title"`
	Clean          bool      `json:"clean_fixture"`
	Hits           []finding `json:"hits"`
	Misses         []finding `json:"misses"`
	FalsePositives []finding `json:"false_positives"`
	Citations      int       `json:"citations"`
	Empty          bool      `json:"empty_output"`
	Pass           bool      `json:"pass"`
}

var (
	bracketGroup = regexp.MustCompile(`\[[^\]]*\brule \d+[^\]]*\]`)
	ruleNumber   = regexp.MustCompile(`\brule (\d+)`)
	entryCode    = regexp.MustCompile(`\b[A-Z]\d+\b`)
)

func gradeReview(scenarioDir, outputPath string) (ReviewGrade, error) {
	scenario, err := loadReviewScenario(scenarioDir)
	if err != nil {
		return ReviewGrade{}, err
	}
	output, err := os.ReadFile(outputPath)
	if err != nil {
		return ReviewGrade{}, fmt.Errorf("read agent output: %w", err)
	}
	return scoreReview(filepath.Base(scenarioDir), scenario, string(output)), nil
}

func loadReviewScenario(dir string) (reviewScenario, error) {
	var scenario reviewScenario
	if err := readJSON(filepath.Join(dir, "expected.json"), &scenario); err != nil {
		return reviewScenario{}, err
	}
	return scenario, nil
}

func readJSON(path string, into any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	if err := json.Unmarshal(raw, into); err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}
	return nil
}

func scoreReview(name string, scenario reviewScenario, output string) ReviewGrade {
	transcript := newTranscript(output)
	grade := ReviewGrade{
		Mode:      "review",
		Scenario:  name,
		Title:     scenario.Title,
		Clean:     len(scenario.MustFind) == 0,
		Hits:      []finding{},
		Misses:    []finding{},
		Citations: transcript.citationCount(),
		Empty:     strings.TrimSpace(output) == "",
	}
	for _, expected := range scenario.MustFind {
		grade.record(expected, transcript.locates(expected))
	}
	grade.FalsePositives = transcript.citedAmong(scenario.MustNotFind)
	grade.Pass = grade.isPass()
	return grade
}

func (g *ReviewGrade) record(expected finding, found bool) {
	if found {
		g.Hits = append(g.Hits, expected)
		return
	}
	g.Misses = append(g.Misses, expected)
}

func (g ReviewGrade) isPass() bool {
	if g.Empty {
		return false
	}
	if g.Clean {
		return g.Citations == 0
	}
	return len(g.Misses) == 0 && len(g.FalsePositives) == 0
}

type transcript struct {
	lines []string
}

func newTranscript(output string) transcript {
	return transcript{lines: strings.Split(output, "\n")}
}

func (t transcript) citationCount() int {
	return len(bracketGroup.FindAllString(strings.Join(t.lines, "\n"), -1))
}

func (t transcript) citedAmong(candidates []finding) []finding {
	cited := []finding{}
	for _, candidate := range candidates {
		if t.cites(candidate) {
			cited = append(cited, candidate)
		}
	}
	return cited
}

func (t transcript) cites(expected finding) bool {
	for _, line := range t.lines {
		if tagMatches(line, expected) {
			return true
		}
	}
	return false
}

// locates reports whether some line cites the rule and that line or the next
// names the file or the symbol.
func (t transcript) locates(expected finding) bool {
	for index, line := range t.lines {
		if tagMatches(line, expected) && t.namesLocation(index, expected) {
			return true
		}
	}
	return false
}

func (t transcript) namesLocation(index int, expected finding) bool {
	window := t.lines[index]
	if index+1 < len(t.lines) {
		window += "\n" + t.lines[index+1]
	}
	return mentions(window, expected.File) || mentions(window, expected.Symbol)
}

func mentions(text, name string) bool {
	return name != "" && strings.Contains(text, filepath.Base(name))
}

// tagMatches accepts the line's bracketed citations as one, so compound forms
// such as `[rule 16 / rule 67 F3]`, `[rule 70 T1, T5]`, and `[rule 1][rule 69 N1]` all count.
func tagMatches(line string, expected finding) bool {
	citations := strings.Join(bracketGroup.FindAllString(line, -1), " ")
	return citesRule(citations, expected.Rule) && citesEntry(citations, expected.Entry)
}

func citesRule(group string, rule int) bool {
	for _, match := range ruleNumber.FindAllStringSubmatch(group, -1) {
		if match[1] == fmt.Sprint(rule) {
			return true
		}
	}
	return false
}

func citesEntry(group, entry string) bool {
	if entry == "" {
		return true
	}
	for _, code := range entryCode.FindAllString(group, -1) {
		if code == entry {
			return true
		}
	}
	return false
}
