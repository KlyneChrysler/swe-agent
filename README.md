# swe

Unforgiving senior-engineer standards for Claude Code.

Install it and you get two strict rulesets plus the tooling to enforce
each:

- **The anti-slop standard**: no duplicate functions, one task per unit,
  one kind per file, tiny leaf functions, no swallowed errors, hot-path
  Big-O discipline, plain one-line comments, no speculative abstraction,
  no dead code.
- **The Clean Code standard**: the complete discipline of Robert C.
  Martin's Clean Code, every chapter and the full smells-and-heuristics
  checklist, restated as seventy numbered rules a reviewer can cite.

"It works" is not the bar. "A staff engineer would approve this on the
first read" is the bar.

## What you get

| Component | What it does |
|---|---|
| `unforgiving-standards` skill | The anti-slop standard. Load it before writing, reviewing, or refactoring, and every line is held to it. |
| `swe` agent | A staff-engineer reviewer that reads your diff, hunts every anti-slop rule (duplication first), verifies each finding, and reports ranked, located, fixable violations. It never praises and never invents findings. |
| `/swe` command | Runs the anti-slop review on your current changes before you commit. |
| `clean-code` skill | The Clean Code standard: names, functions, comments, formatting, objects versus data structures, error handling, boundaries, unit tests, classes, systems, emergence, concurrency, successive refinement, and the checklist (C1 to T9). |
| `clean-code` agent | An engineer, not just a reviewer. Given a task it writes the code test-first under all seventy rules, then reviews its own diff against the checklist and fixes every violation before it returns. Given a diff it reports ranked, located, fixable violations citing rule numbers. |
| `/clean-code` command | `/clean-code <task>` builds it as Clean Code. `/clean-code review [target]` reviews your current changes before you commit. |

## Install

In Claude Code:

```
/plugin marketplace add KlyneChrysler/swe-agent
/plugin install swe@swe-agent
```

Restart Claude Code if prompted. That is it.

## Use

- **To build something**: `/clean-code add a rate limiter to the gateway`.
  The agent reads the surrounding code, names the pieces, writes
  test-first in small cycles, refines against the full checklist until a
  pass finds nothing, and reports what it built and how the tests ran.
- **Before you commit**: run `/swe` or `/clean-code review`. Each reviews
  your uncommitted changes and tells you what an unforgiving reviewer
  would reject, and how to fix it. Run both for the strictest bar.
- **While you write**: the standards load as skills, so Claude holds your
  new code to them as it goes.
- **On a specific target**: `/swe path/to/file`,
  `/clean-code review path/to/file`, or either with a commit range.

Reviews are ranked most-severe first. The anti-slop review orders
correctness and duplication, then architecture, then size and
performance, then naming and comment slop. The Clean Code review orders
correctness and duplication (including null), then design, then
functions, then names, comments, and formatting. A clean diff gets a
one-line "clean" and nothing invented to look busy.

The two standards agree wherever they overlap and are safe to load
together. Where they differ in detail (the anti-slop caps are numeric,
Clean Code's are qualitative), the stricter reading wins.

## The rules, in one breath

One task per unit. One kind per file. No duplicate function anywhere.
Orchestrators delegate, leaves stay under 30 lines. Business logic imports
no infrastructure. Inject dependencies, ban globals. Small interfaces at the
consumer. Immutable by default. Errors wrapped, never swallowed. Size
budgets are caps. Hot paths are O(1)/O(log n) and bounded. Tests are part of
done. Names say what. Comments are one plain line with no em dashes. No
speculative abstraction, no dead code. No templated, vibecoded UI.

The full text lives in
[`skills/unforgiving-standards/SKILL.md`](skills/unforgiving-standards/SKILL.md).

## Clean Code, in one breath

Names reveal intent and never lie. Functions are a few lines, do one
thing, sit at one level of abstraction, take at most three arguments,
never a flag, never an output. Exceptions, not return codes; never return
null, never pass null. Comments are a failure to say it in code. Files
read like a newspaper. Objects hide data, data structures expose it,
nothing is a hybrid. Obey the Law of Demeter. Wrap every boundary. Tests
come first and are F.I.R.S.T. Classes have one reason to change.
Construction is separated from use. Run all the tests, remove
duplication, express intent, then and only then minimize. Concurrency
code lives apart. Never ship the first draft. Leave it cleaner than you
found it.

The full text lives in
[`skills/clean-code/SKILL.md`](skills/clean-code/SKILL.md). It is a
restatement of the discipline in original words; it does not reproduce
the book, and you should still read the book.

## Tested

The reviewer was run against three specimens before release:

1. **Planted slop** - one file with 14 deliberate violations (a byte-identical
   duplicate, a swallowed error, six-level nesting, unused parameters, a dead
   computation, a dead method, mixed kinds in one file, a missing injected
   dependency, a restating comment, a TODO stub, an "and...and" type name, a
   vague verb) plus a fully vibecoded hero component. It caught every one,
   cited the right rule, gave a line-located fix, and ranked the swallowed
   error as the highest-value fix.
2. **Clean code** - files written to the standard. It returned "clean, no
   violations found" and verified that verdict by grepping for duplicates
   before clearing them. Zero invented findings.
3. **Hard duplication** - two functions in two files that compute the same
   result by the same method but share no identical line (renamed function,
   renamed parameters and locals, reordered statements, `<` vs `<=` at the
   boundary). It flagged them as one duplicate under rule 8 and explained why
   the differences are cosmetic, including that the boundary check produces
   identical output.

A reviewer has to do both halves: catch real slop and stay silent on clean
code. This one does both. It is proven on clear cases; rule 18 (UI taste) is
its softer edge, catching the blatant tells but leaving subtle design calls
to you.

The `clean-code` reviewer was run against two specimens before release:

1. **Planted violations** - an `OrderManager` and a hybrid `Order` with
   the classic sins: a pricing loop duplicated across two methods, a
   type switch repeated in business code, `return null`, a defensive
   null check, null as the "no coupon" signal, a swallowed exception
   mapped to a return code, a hardcoded production connection built in a
   field initializer, a five-level Demeter chain, a five-argument
   function with a boolean flag and an output list, a hidden cache
   side effect, magic numbers, money in `double`, a dead helper,
   commented-out code, a changelog header, a restating comment, a
   closing-brace comment, single-letter names, a `Manager` class with
   four reasons to change, and no tests. It caught every one, cited the
   rule and checklist entry, gave a line-located fix, and named the
   duplicated switch as the highest-value fix.
2. **Clean code** - a small value class and its two-test suite written to
   the standard. It reported exactly one finding, an unused import the
   author had actually left in, and cleared everything else with a
   one-line justification per rule group. The import was removed; zero
   invented findings.

The `clean-code` implementer was then run cold on an empty Go module with
a real task: an in-memory account ledger in integer cents with deposits,
withdrawals, all-or-nothing transfers, an audit trail, and persistence
behind a caller-supplied interface wired in one place. It worked
test-first in five red-green cycles and delivered eleven small files,
thirty tests, 100% statement coverage, vet and gofmt clean, functions of
a few lines in stepdown order, wrapped errors, and no nulls. Then the
reviewer was turned on its output, and the fix cycle was repeated:

| Cycle | Findings | Correctness among them |
|---|---|---|
| First review | 13 | self-transfer created money |
| Second review | 5 | a stored negative balance was trusted |
| Third review | 4 | a stored record with the wrong id was trusted |

Each fix cycle ran under the writing protocol, and on the second pass the
implementer found an integer-overflow boundary bug on its own that no
review had listed. The final package has 59 tests at 100% coverage. Two
lessons shaped the protocol: a happy-path suite hides boundary defects,
so the implementer now attacks every boundary with a test before it
refines; and self-review is softer than a stranger's review, so it now
runs the full review-mode hunt over its own diff. Run `/clean-code review`
before you commit anyway. An unforgiving reviewer keeps finding edges,
and that is the point.

## License

Apache License 2.0. See [LICENSE](LICENSE).
