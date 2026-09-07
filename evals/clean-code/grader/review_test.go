package main

import "testing"

var flagScenario = reviewScenario{
	Title:       "flag argument",
	MustFind:    []finding{{Rule: 16, Entry: "F3", File: "report.go", Symbol: "render"}},
	MustNotFind: []finding{{Rule: 20}},
}

func TestCitationOnTheFileLineIsAHit(t *testing.T) {
	output := "[rule 16 F3] report.go:12 - bool flag.\n  Fix: split render."

	grade := scoreReview("017", flagScenario, output)

	if !grade.Pass || len(grade.Hits) != 1 {
		t.Errorf("grade = %+v, want one hit and pass", grade)
	}
}

func TestSymbolOnTheFixLineCountsAsLocated(t *testing.T) {
	output := "[rule 16 F3] somewhere - bool flag.\n  Fix: split render into two functions."

	grade := scoreReview("017", flagScenario, output)

	if len(grade.Hits) != 1 {
		t.Errorf("hits = %+v, want the finding located via the Fix line", grade.Hits)
	}
}

func TestCompoundCitationIsAHit(t *testing.T) {
	output := "[rule 16 / rule 67 F3] report.go:15 and :23 - Render takes a flag."

	grade := scoreReview("017", flagScenario, output)

	if len(grade.Hits) != 1 {
		t.Errorf("hits = %+v, want the compound citation to count", grade.Hits)
	}
}

func TestAdjacentBracketsCountAsOneCitation(t *testing.T) {
	output := "[rule 16][rule 67 F3] report.go:15 - Render takes a flag."

	grade := scoreReview("017", flagScenario, output)

	if len(grade.Hits) != 1 {
		t.Errorf("hits = %+v, want adjacent brackets to count together", grade.Hits)
	}
}

func TestWrongEntryIsAMiss(t *testing.T) {
	output := "[rule 16 F1] report.go:12 - too many arguments."

	grade := scoreReview("017", flagScenario, output)

	if len(grade.Misses) != 1 || grade.Pass {
		t.Errorf("grade = %+v, want a miss", grade)
	}
}

func TestForbiddenRuleIsAFalsePositive(t *testing.T) {
	output := "[rule 16 F3] report.go:12 - flag.\n[rule 20] report.go:20 - duplication."

	grade := scoreReview("017", flagScenario, output)

	if len(grade.FalsePositives) != 1 || grade.Pass {
		t.Errorf("grade = %+v, want one false positive and fail", grade)
	}
}

func TestEmptyTranscriptNeverPasses(t *testing.T) {
	grade := scoreReview("c", reviewScenario{Title: "clean"}, "  \n")

	if grade.Pass || !grade.Empty {
		t.Errorf("grade = %+v, want empty output to fail", grade)
	}
}

func TestCleanFixturePassesOnlyWithoutCitations(t *testing.T) {
	clean := reviewScenario{Title: "clean"}

	if !scoreReview("c", clean, "The diff is clean.").Pass {
		t.Error("no citations should pass a clean fixture")
	}
	if scoreReview("c", clean, "[rule 3] a.go:1 - noise word.").Pass {
		t.Error("a citation should fail a clean fixture")
	}
}
