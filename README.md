# swe

An unforgiving senior-engineer standard for Claude Code that fights AI slop.

Install it and you unlock one strict ruleset plus the tooling to enforce it:
no duplicate functions, one task per unit, one kind per file, tiny leaf
functions, no swallowed errors, hot-path Big-O discipline, plain one-line
comments, no speculative abstraction, no dead code. "It works" is not the
bar. "A staff engineer would approve this on the first read" is the bar.

## What you get

| Component | What it does |
|---|---|
| `unforgiving-standards` skill | The full standard. Load it before writing, reviewing, or refactoring, and every line is held to it. |
| `swe` agent | A staff-engineer reviewer that reads your diff, hunts every rule (duplication first), verifies each finding, and reports ranked, located, fixable violations. It never praises and never invents findings. |
| `/swe` command | Runs the review on your current changes before you commit. |

## Install

In Claude Code:

```
/plugin marketplace add KlyneChrysler/swe-agent
/plugin install swe@swe-agent
```

Restart Claude Code if prompted. That is it.

## Use

- **Before you commit**: run `/swe`. It reviews your uncommitted changes and
  tells you what an unforgiving reviewer would reject, and how to fix it.
- **While you write**: the standard loads as a skill, so Claude holds your
  new code to it as it goes.
- **On a specific target**: `/swe path/to/file` or `/swe <commit-range>`.

The review is ranked most-severe first: correctness and duplication, then
architecture, then size and performance, then naming and comment slop. A
clean diff gets a one-line "clean" and nothing invented to look busy.

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

## License

Apache License 2.0. See [LICENSE](LICENSE).
