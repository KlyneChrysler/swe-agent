---
description: Run the unforgiving swe review on your current changes to catch AI slop before you commit.
---

Review the current code changes against the unforgiving standard and report
ranked, located, fixable violations.

Steps:

1. Load the `unforgiving-standards` skill if it is not already in context.
2. Determine what to review: uncommitted changes (`git diff HEAD`) by
   default; the last commit if the tree is clean; or whatever the user named
   in their arguments ($ARGUMENTS).
3. Dispatch the `swe` agent to review that diff. It reads the surrounding
   code, hunts every rule (duplication first), verifies each finding, and
   returns a ranked list.
4. Relay the agent's verdict. Do not soften it. If the user asks, apply the
   fixes; otherwise stop at the verdict.

If $ARGUMENTS names files, a PR, or a commit range, review that. Otherwise
review the working tree.
