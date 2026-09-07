---
description: Implement a task as unforgiving Clean Code, or review your current changes against it. /clean-code <task> builds it; /clean-code review [target] reviews.
---

Dispatch the `clean-code` agent, the unforgiving Clean Code engineer.

Steps:

1. Load the `clean-code` skill if it is not already in context.
2. Decide the mode from $ARGUMENTS:
   - Starts with `review`: review mode. The target is the rest of the
     arguments (files, a PR, a commit range), or the uncommitted changes
     (`git diff HEAD`) if none, or the last commit if the tree is clean.
   - Anything else: implement mode. The arguments are the task.
   - Empty: review mode on the working tree.
3. Dispatch the `clean-code` agent in that mode. In implement mode it
   reads the surrounding code, names the pieces, writes test-first in
   small cycles, refines against the full checklist until a pass finds
   nothing, and reports what it built and how the tests ran. In review
   mode it hunts every rule, verifies each finding, and returns a ranked
   list.
4. Relay the agent's report. Do not soften a review. If a review is
   followed by "apply the fixes", dispatch the agent again in implement
   mode with the findings as the task.
