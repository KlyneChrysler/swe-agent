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

## License

Apache License 2.0. See [LICENSE](LICENSE).
