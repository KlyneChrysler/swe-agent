package main

import (
	"os"
	"path/filepath"
	"testing"
)

func parseSource(t *testing.T, source string) workspace {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "code.go"), []byte(source), 0o644); err != nil {
		t.Fatal(err)
	}
	sources, err := parseWorkspace(root)
	if err != nil {
		t.Fatal(err)
	}
	return sources
}

func assertProblems(t *testing.T, problems []string, want int) {
	t.Helper()
	if len(problems) != want {
		t.Errorf("problems = %v, want %d", problems, want)
	}
}

func TestBareNilReturnIsFlagged(t *testing.T) {
	sources := parseSource(t, `package p
type T struct{}
func find() *T { return nil }
`)

	assertProblems(t, noNilReturn(sources, ""), 1)
}

func TestNilWithErrorIsAllowed(t *testing.T) {
	sources := parseSource(t, `package p
import "errors"
type T struct{}
func find() (*T, error) { return nil, errors.New("missing") }
`)

	assertProblems(t, noNilReturn(sources, ""), 0)
}

func TestNestingCountsBlocksInsideTheBody(t *testing.T) {
	sources := parseSource(t, `package p
func deep(items []int) {
	for _, item := range items {
		if item > 0 {
			if item > 1 {
				println(item)
			}
		}
	}
}
`)

	assertProblems(t, maxNesting(sources, "2"), 1)
}

func TestBoolParameterIsFlagged(t *testing.T) {
	sources := parseSource(t, `package p
func render(compact bool) {}
`)

	assertProblems(t, noBoolParams(sources, ""), 1)
}

func TestCommentedCodeIsFlagged(t *testing.T) {
	sources := parseSource(t, `package p
// old := compute(x)
// Explains intent, and is not code.
func compute() {}
`)

	assertProblems(t, noCommentedCode(sources, ""), 1)
}

func TestSmellSuffixesAreFlagged(t *testing.T) {
	sources := parseSource(t, `package p
type OrderManager struct{ userData int }
func helper() {}
`)

	assertProblems(t, noSmellNames(sources, ""), 2)
}

func TestSwitchOnOneFieldInTwoFunctionsIsFlagged(t *testing.T) {
	sources := parseSource(t, `package p
type Shape struct{ Kind string }
func area(s Shape) int { switch s.Kind { case "a": return 1 }; return 0 }
func name(s Shape) string { if s.Kind == "a" { return "a" }; return "" }
`)

	assertProblems(t, singleSwitch(sources, "Kind"), 1)
}

func TestErrorAndFuncGlobalsAreAllowed(t *testing.T) {
	sources := parseSource(t, `package p
import "errors"
var ErrMissing = errors.New("missing")
var now = func() int { return 0 }
var registry = map[string]int{}
`)

	assertProblems(t, noGlobalMutable(sources, ""), 1)
}

func TestGetterSetterPairIsFlagged(t *testing.T) {
	sources := parseSource(t, `package p
type Account struct{ balance int }
func (a Account) Balance() int { return a.balance }
func (a *Account) SetBalance(b int) { a.balance = b }
`)

	assertProblems(t, noGetterSetterPairs(sources, ""), 1)
}

func TestSymbolChecks(t *testing.T) {
	sources := parseSource(t, `package p
type Ledger struct{}
func (l *Ledger) Open() {}
`)

	assertProblems(t, symbolExists(sources, "Open"), 0)
	assertProblems(t, symbolAbsent(sources, "Ledger"), 1)
	assertProblems(t, symbolAbsent(sources, "Open"), 1)
}
