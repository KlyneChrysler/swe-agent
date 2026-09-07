// Command grader scores one run of the clean-code agent against a scenario,
// or summarises a directory of such scores.
//
//	grader review <scenario-dir> <output.md>       score a review transcript
//	grader implement <scenario.json> <workspace>   score an implement workspace
//	grader seed <scenario.json> <workspace>        write a scenario's seed files
//	grader summary <results-dir>                   aggregate grade.json files
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "grader:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 3 && !(len(args) == 2 && args[0] == "summary") {
		return fmt.Errorf("usage: grader review|implement|seed <scenario> <target>, or grader summary <results-dir>")
	}
	switch args[0] {
	case "review":
		return printJSON(gradeReview(args[1], args[2]))
	case "implement":
		return printJSON(gradeImplement(args[1], args[2]))
	case "seed":
		return writeSeed(args[1], args[2])
	case "summary":
		return summarise(args[1])
	}
	return fmt.Errorf("unknown mode %q", args[0])
}

func printJSON(value any, err error) error {
	if err != nil {
		return err
	}
	encoded, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode grade: %w", err)
	}
	fmt.Println(string(encoded))
	return nil
}
