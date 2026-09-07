package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type implementScenario struct {
	Title          string            `json:"title"`
	Task           string            `json:"task"`
	Seed           map[string]string `json:"seed"`
	RulesExercised []int             `json:"rules_exercised"`
	Checks         []string          `json:"checks"`
}

// ImplementGrade is the score of one implement workspace.
type ImplementGrade struct {
	Mode           string        `json:"mode"`
	Scenario       string        `json:"scenario"`
	Title          string        `json:"title"`
	RulesExercised []int         `json:"rules_exercised"`
	Results        []checkResult `json:"results"`
	Pass           bool          `json:"pass"`
}

type checkResult struct {
	Check    string   `json:"check"`
	Pass     bool     `json:"pass"`
	Problems []string `json:"problems,omitempty"`
}

func gradeImplement(scenarioPath, workspace string) (ImplementGrade, error) {
	var scenario implementScenario
	if err := readJSON(scenarioPath, &scenario); err != nil {
		return ImplementGrade{}, err
	}
	sources, err := parseWorkspace(workspace)
	if err != nil {
		return ImplementGrade{}, err
	}
	return scoreImplement(scenarioName(scenarioPath), scenario, sources), nil
}

func scenarioName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), ".json")
}

func scoreImplement(name string, scenario implementScenario, sources workspace) ImplementGrade {
	grade := ImplementGrade{
		Mode:           "implement",
		Scenario:       name,
		Title:          scenario.Title,
		RulesExercised: scenario.RulesExercised,
		Results:        []checkResult{},
		Pass:           true,
	}
	for _, check := range scenario.Checks {
		grade.add(runCheck(check, sources))
	}
	return grade
}

func (g *ImplementGrade) add(result checkResult) {
	g.Results = append(g.Results, result)
	g.Pass = g.Pass && result.Pass
}

func runCheck(spec string, sources workspace) checkResult {
	name, argument := splitCheck(spec)
	check, known := checks[name]
	if !known {
		return checkResult{Check: spec, Problems: []string{"unknown check"}}
	}
	problems := check(sources, argument)
	return checkResult{Check: spec, Pass: len(problems) == 0, Problems: problems}
}

func splitCheck(spec string) (string, string) {
	name, argument, _ := strings.Cut(spec, ":")
	return name, argument
}

func limit(argument string) int {
	value, err := strconv.Atoi(argument)
	if err != nil {
		return 0
	}
	return value
}

func writeSeed(scenarioPath, workspace string) error {
	var scenario implementScenario
	if err := readJSON(scenarioPath, &scenario); err != nil {
		return err
	}
	for relative, content := range scenario.Seed {
		if err := writeFile(filepath.Join(workspace, relative), content); err != nil {
			return err
		}
	}
	return nil
}

func writeFile(path, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create %s: %w", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
