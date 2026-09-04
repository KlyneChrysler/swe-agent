---
name: unforgiving-standards
description: Use before writing, reviewing, or refactoring any code. The unforgiving senior-engineer standard that fights AI slop - one task per unit, one kind per file, no duplicate functions, tiny leaves, no swallowed errors, hot-path Big-O, plain one-line comments. Load it and hold every line to it.
---

# The Unforgiving Standard

This is a senior-engineer standard written to fight AI slop. It is
unforgiving on purpose. A rule broken is a defect, even when the code runs,
even when it is small, even when you are in a hurry. "It works" is not the
bar. "A staff engineer would approve this on the first read" is the bar.

AI-generated code fails in predictable ways: it over-comments, it invents
abstractions nobody asked for, it copies a helper instead of reusing one, it
writes 80-line functions, it names things `data` and `helper`, it leaves
`TODO` stubs and `// in a real implementation` apologies, it swallows
errors, it restates the code in a comment. Each rule below kills one of
those failure modes. Follow them all, every time.

## The rules

### 1. One task per unit

Every function, method, or component does exactly one thing. Two
responsibilities is two functions. The test: name it without using "and". A
name with "and" in it, or a name so vague it could mean anything
(`handle`, `process`, `manage`, `doWork`), is the smell of a unit doing too
much or understood too little.

### 2. One kind per file

Every declaration has a KIND. A file holds one kind, serving one purpose:

- a data shape (DTO, wire struct, record) is a shape file
- an interface is an interface file
- a mapper or codec (converts between two representations) is its own file
- behavior (handlers, services, the actual work) is a behavior file
- configuration, errors, registries, constants each have their own file

A shape declared inside a behavior file is a defect. A conversion function
buried next to the handler that calls it is a defect - conversion is a kind,
not a stray helper. This holds for private declarations too: there is no
"it's unexported so it doesn't count" exception. If your language forces one
public type per file, that is the floor, not the ceiling.

### 3. Orchestrators delegate, leaves compute

Two kinds of function exist, each with its own budget:

- **Orchestrators** coordinate a task. Every line delegates to a named
  step. They read like prose. No inline computation, no nested branching.
- **Leaves** do the actual work. Target 15 lines, hard cap 30. A leaf that
  outgrows the cap is split, and the split parts are extracted and named.

Do not shred a computation that belongs together into a dozen one-liners.
Cohesion still rules; the cap is 30, not 1.

### 4. Dependencies point one way

Business logic must not import infrastructure. The domain (the rules of your
problem) compiles and unit-tests with zero database, zero HTTP, zero
framework. Infrastructure depends on the domain, never the reverse. One
framework import inside your core logic is a defect, not a shortcut. This is
hexagonal / clean / onion architecture; the name does not matter, the
one-way arrow does.

### 5. Inject dependencies, ban globals

Collaborators arrive through the constructor (or function parameters). No
global mutable state, no singletons, no service locators, no hidden
`init()` side effects. Concrete types are wired together in exactly one
place - the composition root (`main`, the DI container, the app entry). Code
that reaches out to grab its own dependencies cannot be tested and cannot be
reasoned about.

### 6. Small interfaces, defined at the consumer

An interface has 1-3 methods and lives with the code that USES it, not the
code that implements it. A fat interface (5+ methods) is a design failure -
split it. "Accept interfaces, return structs."

### 7. Immutability by default

Values do not mutate. Transformations return new values. Mutation is allowed
only inside a single function's local scope, or where the language idiom
demands it and a comment says why. Shared mutable state is where bugs breed.

### 8. No duplicate function or method, anywhere

Two functions that do the same thing are one function in the wrong number of
places. This is the single most common AI-slop tell, and it is unforgiving:

- **Byte-identical bodies** are a defect. Delete the copy, point callers at
  the one that stays.
- **Same logic, cosmetic differences** (renamed locals, reordered
  independent lines, a hardcoded value that should be an argument) is a
  duplicate. Extract one function; pass the difference as a parameter.
- **Placement** follows the layer rules: shared by two modules goes to the
  shared library; shared inside one module goes to its lowest shared spot;
  used by one type becomes a private method on it. Never copy to dodge an
  import.
- Test helpers copied across test files count too. One definition.

Before writing any function, ask: does one that does this already exist? If
yes, call it. If it almost does, parameterize the difference. Writing the
second copy is never the answer.

### 9. Errors are handled, wrapped, never swallowed

Every error is handled meaningfully or propagated with added context. Empty
catch blocks, ignored return values, and log-and-continue on an error that
matters are defects. Wrap with context at each boundary so a failure reads
as a trail, not a mystery. User-facing surfaces get a friendly message; the
log gets the detail. An intentionally ignored error carries a comment
saying why, every time.

### 10. Size budgets are caps, not suggestions

| Unit | Target | Hard cap |
|---|---|---|
| Leaf function | 15 lines | 30 lines |
| File | 200-400 lines | 800 lines |
| UI component | 100 lines | 150 lines |
| Interface | 1-3 methods | 5 methods |
| Function parameters | 3 | 5 (then pass a struct) |
| Nesting depth | 2 | 4 |

Over a hard cap blocks the change. Over a target means the split is your
next edit, not a someday.

### 11. Performance is a rule, not an afterthought

Distinguish hot paths (run per request, per event, per item, per frame)
from cold paths (startup, admin, one-offs).

- **Hot path**: O(1) or O(log n) per item. Membership is a map/set lookup,
  never a linear scan. No allocation that scales with traffic where it can
  be avoided (pre-size, reuse, compile-once). Everything bounded - queues,
  caches, retries have a cap and a defined overflow policy. No lock held
  across I/O. A quadratic on a hot path is a defect.
- **Cold path**: no accidental quadratic (nested scans over collections that
  grow together). Any super-linear function carries a one-line complexity
  comment stating its class and its bound.

Pick the right data structure first; micro-tuning without a complexity
argument is rejected.

### 12. Tests are part of done

Domain and application logic is tested first (write the test, watch it fail,
make it pass, refactor). Table-driven / parameterized tests over copy-pasted
cases. Hand-written fakes over heavy mocking where the language allows.
Adapters get an integration test against a real (containerized) dependency.
Untested core logic is unfinished, not done.

### 13. Names say what, types say how

Names carry meaning. `verdictStore`, not `db2`. Booleans read as predicates
(`isReady`, `hasExpired`). Packages and modules are nouns, functions are
verbs. No invented abbreviations (industry-standard ID, URL, TLS are fine).
One concept per file, named after it. A name you have to read the body to
understand is the wrong name.

### 14. Comments: one line, plain words

Every comment is one line, high level, in simple words. It says what the
thing is FOR, never restates HOW the code works line by line. No em dashes
anywhere. A comment that needs a second line means the code needs a better
name or a smaller function - fix the code, not the comment. Delete every
comment that narrates the obvious (`// increment i`), every `TODO` with no
ticket, every `// in a real implementation this would...` apology. Working
code has no apologies in it.

### 15. Declarations on one line, air inside methods

Function and method signatures do not wrap parameters onto continuation
lines, however long. A call too long for one line gets named local variables
first, then a one-line call. Inside any method longer than three lines,
blank lines separate setup, action, and result. Dense walls of code and
half-wrapped signatures both hide bugs.

### 16. Configuration from the environment, validated at startup

Anything that varies between deploys (addresses, credentials, ports, flags)
comes from the environment, not a constant and not `if env == "prod"`.
Validate the whole config at startup and crash loudly on anything missing.
Fail at boot, never at the first request that needed the missing value.

### 17. No speculative abstraction, no dead code

Build for the caller you have, not the three you imagine. No interface with
one implementation "in case". No configuration option nobody sets. No
generic framework for a specific job. Delete dead code the moment it is
dead - commented-out blocks, unused parameters, unreachable branches,
scaffolding "kept just in case". YAGNI is a rule here, not a mood.

## How to use this

- **Writing code**: hold each unit to every rule as you write it. The rules
  are cheap up front and expensive to retrofit.
- **Reviewing code**: walk the diff rule by rule. Report each violation with
  the rule number, the location, and the fix. Rank by severity: correctness
  and duplication first, then architecture, then size and naming, then
  comments.
- **Refactoring**: fix violations without changing behavior. Prove behavior
  is unchanged with the tests before and after.

The `swe` reviewer agent applies this standard to a diff on demand. The
`/swe` command runs it against your current changes.
