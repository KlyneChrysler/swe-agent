# clean-code agent evaluation suite: scenario format

Scenarios exercise the `clean-code` agent (agents/clean-code.md) against the
rules in skills/clean-code/SKILL.md (rules 1-70, checklist entries C1-C5,
E1-E2, F1-F4, G1-G36, N1-N7, T1-T9). Everything is Go, standard library only,
so the grader can use go/ast and `go test`.

## Review scenarios: `review/NNN-slug/`

One directory per scenario, NNN zero-padded from 001.

```
review/017-boolean-flag-argument/
  fixture/            # a Go package named `fixture` that COMPILES (go build ./...)
    report.go
    report_test.go    # optional
  expected.json
```

The fixture is presented to the agent as uncommitted changes in a fresh repo,
so the agent's default `git diff HEAD` shows every fixture file.

`expected.json`:

```json
{
  "title": "Boolean flag argument selects between two renderings",
  "must_find": [
    {"rule": 16, "entry": "F3", "file": "report.go", "symbol": "render",
     "why": "bool parameter `compact` makes render do two things"}
  ],
  "must_not_find": [
    {"rule": 20, "why": "the two branches share no lines; not duplication"}
  ]
}
```

- `must_find`: every violation deliberately planted. `rule` is the rule number
  the agent must cite. `entry` is the checklist entry (C/E/F/G/N/T) when one
  applies, else omit. `file` is relative to `fixture/`. `symbol` is the
  function, type, or method name. `why` is one sentence for the human.
- `must_not_find`: rules a careless reviewer would raise here but that do not
  apply; the grader penalises them. Optional, may be empty.
- A scenario with an empty `must_find` is a CLEAN fixture: the agent must say
  the diff is clean and report nothing. About 8 of every 100 are clean.
- Plant 1 to 3 violations per scenario, never more. Everything else in the
  fixture must be clean by every rule, so a finding outside `must_find` is a
  false positive. Write the fixture as carefully as production code.
- Grading counts a `must_find` as hit when the agent's output contains
  `[rule N]` or `[rule N X]` for that rule (and that entry when given) on the
  same line as the file name or the symbol. Cite from the rule list only.

## Implement scenarios: `implement/NNN-slug.json`

```json
{
  "title": "Parse a duration string into seconds",
  "task": "Implement `ParseDuration(text string) (int, error)` in package `dur` ...",
  "seed": {
    "dur/dur.go": "package dur\n\n// existing code the agent must read and match ...\n"
  },
  "rules_exercised": [16, 19, 36, 41],
  "checks": [
    "tests_exist", "tests_pass", "vet_clean", "gofmt_clean",
    "max_func_lines:20", "max_params:3", "no_bool_params",
    "symbol_exists:ParseDuration"
  ]
}
```

- `task` is the prompt handed to the agent, in the user's voice. It must be
  solvable with the standard library in a single small Go module named
  `example.com/eval` (the runner creates the module). It should tempt a
  specific violation (a flag, a null return, a god function, a comment where
  a name would do, a hybrid, construction inside logic) so the check is
  meaningful.
- `seed` is optional: files that exist before the agent starts. Use it to
  test rule 1 (read and match vocabulary), rule 20 (an existing helper the
  agent must reuse), and the Boy Scout Rule (a fixable violation in a file
  the task must touch). Paths are relative to the module root.
- `rules_exercised`: the rules the task is designed to tempt.
- `checks`: drawn only from this list, applied to every non-test Go file the
  agent leaves in the module unless noted:

  | check | meaning |
  |---|---|
  | `tests_exist` | at least one `_test.go` file with a `Test` function |
  | `tests_pass` | `go test ./...` exits 0 |
  | `vet_clean` | `go vet ./...` exits 0 |
  | `gofmt_clean` | `gofmt -l .` prints nothing |
  | `max_func_lines:N` | no function body longer than N lines |
  | `max_params:N` | no function with more than N parameters |
  | `max_nesting:N` | no block nested deeper than N inside a function |
  | `no_bool_params` | no parameter of type bool |
  | `no_nil_return` | no `return nil` for a pointer, slice, map, or interface result unless an error is also returned non-nil in the same statement |
  | `no_comments_in_funcs` | no comments inside function bodies |
  | `no_banner_comments` | no comment line made mostly of `-`, `=`, `*`, `#` |
  | `no_commented_code` | no comment line ending in `{`, `}`, `;`, or `)` that parses as Go |
  | `no_todo` | no TODO/FIXME/XXX comment |
  | `no_smell_names` | no identifier ending in Manager, Helper, Util, Utils, Processor, Data, Info |
  | `symbol_exists:Name` | a top-level func, type, or method named Name exists |
  | `symbol_absent:Name` | no func, method, or type named Name exists |
  | `max_files:N` | at most N non-test Go files |
  | `no_panic` | no call to panic |
  | `no_global_mutable` | no package-level `var` of non-error, non-function type |
  | `single_switch:Field` | `Field` appears as a switch tag or in `==` comparisons in at most one function |
  | `no_getter_setter_pairs` | no type with both GetX/X and SetX for the same field |

- Every implement scenario includes `tests_exist`, `tests_pass`, `vet_clean`,
  `gofmt_clean`, and `max_func_lines:20`; add the checks the task tempts.

## Distribution

Across each set of 100, every rule 1-70 is exercised at least once, and every
checklist entry at least once. Mix sizes: most scenarios are one file of 20-60
lines; about ten are two or three files (duplication across files, boundaries,
construction versus use, feature envy across types).
