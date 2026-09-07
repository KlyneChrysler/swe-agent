---
name: clean-code
description: Use before writing, reviewing, or refactoring any code. The complete Clean Code discipline (Robert C. Martin) restated as unforgiving numbered rules - intention-revealing names, tiny functions that do one thing at one level of abstraction, comments only where code cannot speak, newspaper formatting, objects versus data structures, exceptions not return codes, wrapped boundaries, F.I.R.S.T. tests, small cohesive classes, separated construction, the four rules of simple design, concurrency discipline, and the full smells-and-heuristics checklist. Load it and hold every line to it.
---

# The Clean Code Standard

This is the discipline of Clean Code, written down as rules you can be held
to. It is unforgiving on purpose. A rule broken is a defect, even when the
code runs, even when the change is small, even when you are in a hurry.
"It works" is not the bar. "A reader understands it on the first pass and
would not want to change it" is the bar.

Two laws sit above every rule below:

- **The Boy Scout Rule.** Leave every file cleaner than you found it. A
  change that touches a file and leaves a violation it could have fixed
  is an incomplete change.
- **The reader is the customer.** Code is read far more than it is
  written. Every rule here exists to make the next read cheaper. Optimize
  for the reader, then for the writer, never the reverse.

Rules are numbered so a review can cite them. Sections follow the
chapters of the book.

## A. Meaningful names

### 1. Names reveal intent

A name answers why the thing exists, what it does, and how it is used. If
a name needs a comment to explain it, the name is wrong. A single letter
or a vague word for anything that outlives one short loop is a defect.
`elapsedDays` says it; `d` does not.

### 2. Names never disinform

Do not call a thing a list unless it is a list. Do not use names that
differ by one character or by a subtle suffix. Do not use names that look
like other names or like literals. A name that suggests a wrong meaning is
worse than a name that suggests nothing.

### 3. Distinctions are meaningful

Two names must differ in meaning, not just in spelling. Number series
(`a1`, `a2`), noise words (`Info`, `Data`, `Object`, `the`), and
near-synonym method names on the same class (`getAccount`,
`getAccountInfo`, `getAccountData`) are defects. If the reader cannot
tell which one to call from the names alone, rename until they can.

### 4. Names are pronounceable and searchable

If you cannot say it out loud, rename it. If you cannot grep for it, it is
too short or it is a magic literal. Name length grows with scope: a
one-letter loop counter in a three-line loop is fine; anything wider gets
a real word. Every constant that means something has a name.

### 5. No encodings

No type prefixes, no member prefixes, no interface prefixes. The reader
should not have to decode. If something must carry a marker, mark the
implementation, never the abstraction the caller sees.

### 6. No mental mapping

A reader must not translate your name into the concept in their head.
Clever single-letter shorthand and private nicknames are defects.
Clarity beats cleverness every time.

### 7. Classes are nouns, methods are verbs

Class names are noun phrases and never verbs. Method names are verb
phrases. Accessors, mutators, and predicates use the get, set, and is
prefixes or the language's idiom. When constructors would be overloaded,
use named static factory methods that say what the argument means.

### 8. One word per concept, one concept per word

Pick one word for one idea across the whole codebase. `fetch`,
`retrieve`, and `get` for the same operation on three classes is a
defect. The reverse holds too: do not reuse one word for two ideas.
`add` that concatenates in one class and inserts in another is a pun.

### 9. Solution-domain names first, problem-domain second

Use computer-science and pattern names where they apply (`Visitor`,
`Queue`, `Factory`). Where none applies, use the problem-domain name so a
domain expert recognizes it. Never invent a third vocabulary.

### 10. Context is meaningful, never gratuitous

A field like `state` means nothing alone; give it context by placing it
in a class or by a prefix that says whose state. The opposite is also a
defect: do not stamp an application prefix on every class and variable.
Short names are fine when the surroundings already say what they are.

## B. Functions

### 11. Small, then smaller

A function is a few lines. Twenty is long. The body of every `if`,
`else`, `while`, or loop is one line, and that line is usually a call to
a function whose name documents the block. Indentation depth never
exceeds two levels. A function that scrolls is a defect.

### 12. A function does one thing

One thing means all its statements are one level below the name. The
test: if you can extract another function from it with a name that is
not a restatement of the original, it was doing more than one thing.
A function with sections is doing more than one thing.

### 13. One level of abstraction per function

Do not mix a high-level concept with a low-level detail in one body.
Code reads top-down like a narrative: each function is followed by the
functions it calls, at the next level down. Mixing levels makes the
reader guess which lines are essential.

### 14. Switch statements appear once, buried, behind polymorphism

A switch or long if-else chain on a type discriminator is allowed exactly
once, at the lowest level, inside a factory that creates polymorphic
objects. The same discriminator switched on in two places is a defect.
Repeated conditionals on the same field are a class trying to be born.

### 15. Descriptive names, however long

A long descriptive name beats a short cryptic one and beats a comment.
Be consistent: use the same phrases, nouns, and verbs across related
functions so the reader can predict the next name.

### 16. Arguments are few, flags and output arguments are banned

Zero arguments is ideal, one is good, two is acceptable, three requires
a strong reason, more is a defect. A boolean argument means the function
does two things; split it. An argument the function writes into is a
defect; return the value instead. When several arguments belong together,
they are a missing object. Name the function so the argument order reads
naturally.

### 17. No side effects, no hidden temporal coupling

A function does what its name says and nothing else. Touching state the
name does not mention, or requiring that it be called in a specific order
that nothing enforces, is a defect. Make ordering explicit by passing the
result of one step into the next.

### 18. Command-query separation

A function either does something or answers something. One that changes
state and returns a status about it is doing both and is a defect. Split
it into a command and a query.

### 19. Exceptions over error codes, error handling is one thing

Return codes force the caller into immediate nested checks; throw instead
and let the logic stay flat. The body of a `try` and the body of a
`catch` are each extracted into their own function. A function that
handles errors does nothing else: `try` is the first word in it and
nothing follows the last `catch` or `finally`. Error enums that every
caller depends on are a defect.

### 20. Do not repeat yourself

Duplication is the root of most smells. The same lines, the same
structure, or the same algorithm in two places is a defect. Extract it.
Every rule in this file is partly a weapon against duplication.

### 21. Structure serves the small function

Single-entry single-exit matters in a large function. In a function that
is a few lines long, multiple returns and early exits are fine when they
read better. `goto` is never fine.

## C. Comments

### 22. A comment is a failure to express yourself in code

Every comment is the author admitting the code did not say it. Before
writing a comment, try to say it with a better name, a smaller function,
or an explanatory variable. Comments rot: they drift from the code and
then lie. Code never lies.

### 23. The few good comments

Allowed, when nothing else will do: a required legal header; a note that
makes a regular expression or a format string readable; an explanation of
intent behind a decision; a clarification of an obscure value from a
library you cannot change; a warning of consequences; a task note that
carries a tracker reference and gets removed on schedule; amplification
of something that looks trivial but is not; documentation on a public
API.

### 24. Everything else is noise, and noise is a defect

Banned: comments that restate the code; comments that are wrong or vague;
comments mandated by policy on every member; change logs and author
bylines inside source; comments that say nothing (`default constructor`);
banner comments and section markers; closing-brace comments; commented-out
code, which version control already remembers; markup inside comments;
comments describing code that lives somewhere else; comments longer than
the code they describe; comments whose connection to the code is unclear;
a header on a short, well-named function; API documentation on code that
is not a public API. Delete them.

## D. Formatting

### 25. Vertical formatting: the newspaper

Files are small. A few hundred lines is normal, and a file near five
hundred is long. A file reads like a newspaper article: the name is the
headline, the top holds the high-level concepts, and detail increases as
you scroll. Blank lines separate concepts; related lines sit tightly
together. Variables are declared as close to their use as possible;
instance variables sit at the top of the class. A function that calls
another sits above it, as close as possible. Things that belong together
conceptually sit together.

### 26. Horizontal formatting: short lines, honest whitespace

Lines are short; past about a hundred and twenty characters is
carelessness. Whitespace shows what goes together: around assignment
operators, after commas, none between a function name and its paren.
Precedence is made visible by spacing. Do not align declarations in
columns; it draws the eye to the wrong thing. Never collapse an `if`,
loop, or empty body onto one line to save space. Indentation is never
broken.

### 27. One team, one style

The team agrees on one set of formatting rules, encodes them in a tool,
and every member follows it. Code that looks like it was written by
several people arguing is a defect regardless of whose style won.

## E. Objects and data structures

### 28. Hide implementation behind abstractions

Exposing every field through a getter and setter is not abstraction; it
is the field with extra steps. Expose an interface that lets the caller
manipulate the essence of the data without knowing its representation.
Think hard about the best way to represent what an object holds, and do
not reflexively add accessors.

### 29. Objects or data structures, never hybrids

An object hides its data and exposes behavior. A data structure exposes
its data and has no meaningful behavior. Each has its strength: with
objects it is easy to add new types; with data structures and procedures
it is easy to add new functions. A hybrid, half object and half struct,
gets the worst of both and is a defect. Pick one deliberately per type.

### 30. The Law of Demeter

A method on an object may call methods on: the object itself, objects
passed as arguments, objects it creates, and objects held in its own
fields. It may not call methods on the objects returned by any of those.
A chain of calls that walks through returned objects is a train wreck
and a defect. Chains through pure data structures are acceptable because
a data structure has no internals to protect. Do not fix a train wreck
by turning an object into a hybrid.

### 31. DTOs and records are data structures

A transfer object, record, or row mapping holds public data and no
business rules. Do not put behavior in it. Do not make it a hybrid by
adding navigational methods that pretend it is an object.

## F. Error handling

### 32. Exceptions, not return codes; write the try first

Use exceptions so the happy path stays clean and error handling stays in
one place. When a function can fail, write the `try` and `catch` before
its body: it defines the scope of what the caller can rely on. Prefer
unchecked exceptions where the language offers the choice; checked
exceptions ripple signature changes through every layer.

### 33. Exceptions carry context

The message says what operation was attempted and what kind of failure
occurred, with enough detail to locate the cause without a debugger. An
exception with no message, or a message that repeats the type name, is a
defect.

### 34. Exception classes serve the caller, and boundaries get wrapped

Define exception types by how callers will handle them, not by where
they come from. Wrap third-party APIs in a class of your own that
translates their many exceptions into the one or two your code cares
about. The wrapper also cuts the dependency on the library.

### 35. Define the normal flow with special cases

When a missing or empty case is expected, return an object that
represents it and behaves correctly, so callers do not need a `catch`
for business as usual. Exceptions are for the exceptional.

### 36. Never return null, never pass null

A returned null is a defect: return an empty collection, a special-case
object, or throw. A null argument is a defect: no function should have to
defend against it, and a caller passing one is asking for a crash later.
Null checks sprinkled through code are the symptom, not the fix.

## G. Boundaries

### 37. Do not pass boundary types around

A generic container or a third-party type is a boundary. Passing it
across your system means every user of it depends on every method it has
and on every change its vendor makes. Wrap it in a class that exposes
only the operations you need, in the terms your system uses.

### 38. Learning tests pin the boundary

Learn a third-party API by writing tests against it that check the
behavior you rely on. Keep them. When the library upgrades, they tell
you exactly what changed before your production code finds out.

### 39. Code that does not exist yet gets the interface you wish you had

When the other side of a boundary is unknown or unfinished, define the
interface your code wants, write against it, and adapt the real thing to
it when it arrives. Never let an unbuilt dependency dictate your design.

### 40. Few places touch the boundary

Third-party code is referenced from as few places as possible. Either
wrap it or adapt it; never spread its types and calls across the code
you own. Your code should depend on what you control.

## H. Unit tests

### 41. The three laws

Write no production code until you have a failing unit test. Write no
more of a test than is needed to fail; not compiling counts as failing.
Write no more production code than is needed to pass. Test and code grow
together in a cycle measured in seconds, not hours.

### 42. Test code is production code

Tests are held to every rule in this file. A dirty test is worse than no
test because it decays, gets abandoned, and takes the safety net with
it. Without tests, no one dares change the production code, and it rots.

### 43. Readable tests: one concept, minimal assertions

Every test reads as build, operate, check. Build a small domain-specific
vocabulary of helpers so the test says what it means in the language of
the problem. Test one concept per test function. Keep assertions per
test to a minimum; one is the target, several on one concept is the
tolerance, several concepts is a defect.

### 44. F.I.R.S.T.

Fast: slow tests do not get run. Independent: no test depends on the
state left by another, and any order passes. Repeatable: same result in
every environment, online or offline. Self-validating: a boolean pass or
fail, never a log to read by hand. Timely: written just before the code
they test, not after.

### 45. The dual standard

A test may spend resources a production system could not, for the sake
of readability. It may never be less clean, and it may never be less
correct.

## I. Classes

### 46. Class organization and the stepdown

Constants first, then static fields, then instance fields, then public
functions, with each private helper directly after the public function
that calls it. The class reads top-down. Loosening visibility for tests
is a last resort, taken only after every other option fails.

### 47. Classes are small, measured in responsibilities

The measure is not lines. A class has exactly one responsibility and
one reason to change. If you cannot describe the class in about
twenty-five words without using "if", "and", "or", or "but", it has too
many. Names ending in `Manager`, `Processor`, `Super`, or `Util` are the
smell of responsibilities piled up.

### 48. Cohesion, or the class trying to get out

Methods use the instance variables. When most methods use most fields,
the class is cohesive. When a subset of methods shares a subset of
fields, that subset is a separate class. When splitting a large function
makes you promote locals to fields so the pieces can share them, a new
class is trying to be born; extract it.

### 49. Organize for change

New behavior arrives by adding a class, not by editing a working one.
Depend on abstractions, not concrete details, so a change in detail does
not reach the code that uses it. A class that must be opened and edited
for every new variant is a defect.

## J. Systems

### 50. Separate construction from use

The application does not build its own collaborators. Construction
happens in one place: the entry point, a factory, or a dependency
injection container. Business logic receives objects ready to use and
never calls a constructor of a collaborator. When the application must
control when something is created but not how, hand it a factory. Lazy
initialization scattered through methods is a defect.

### 51. Cross-cutting concerns are separated, domain objects stay plain

Persistence, transactions, security, and logging are wired around
domain objects, not written into them. Domain objects are plain objects
with no knowledge of the framework that hosts them. Framework
annotations or base classes inside domain logic are a defect.

### 52. Defer decisions, grow the architecture

Choose the smallest thing that works today and keep the ability to
change it. Do not adopt a heavyweight framework up front. Let tests and
clean boundaries make the architecture able to grow. A big design bought
before it is needed is a defect the same as dead code.

### 53. Standards and DSLs only where they pay

Use a standard when it demonstrably adds value, not because it is a
standard. Use a domain-specific language when it lets the code read like
the domain, not to show off.

## K. Emergence: the four rules of simple design

### 54. Runs all the tests

A system that cannot be verified cannot be trusted and must not ship.
Making a system testable pushes it toward small, single-purpose classes
and low coupling.

### 55. Contains no duplication

Duplication of lines, of structure, of algorithms. Remove it at every
scale, including structural duplication that a template method or a
shared abstraction would collapse. Even a few lines repeated are a
defect.

### 56. Expresses the intent of the programmer

Good names, small functions, small classes, standard pattern names where
a pattern is used, and tests that read as documentation. The clearest
expression wins. Care is the difference.

### 57. Minimizes the number of classes and methods

The first three rules taken to an extreme create too many tiny pieces.
Do not build an interface for every class, do not split data from
behavior on principle, do not follow a dogma. Keep the count small, but
this rule is the lowest priority of the four.

## L. Concurrency

### 58. Concurrency is a design decision, kept apart

Concurrency decouples what from when and carries its own design. Code
that manages threads or locks lives in its own small units, separate
from the code that does the work. Mixing them is a defect.

### 59. Limit and copy shared data

Shared mutable data is the source of every concurrency bug. Encapsulate
it severely, keep as little of it as possible, prefer working on copies,
and make threads as independent as the problem allows.

### 60. Know your library and your execution model

Use the thread-safe collections, executors, and non-blocking primitives
your platform provides rather than hand-rolling them. Know the standard
models, producer-consumer, readers-writers, and the dining philosophers,
and the failure modes each invites.

### 61. Synchronized sections are tiny and independent

A lock is held for as few lines as possible and never across I/O. Two
synchronized methods on one shared object that must be called together
are a hidden coupling; make it explicit with one method or one lock.
Locking done by the client of an object rather than inside it is
fragile and a defect.

### 62. Shut down correctly

Graceful shutdown is hard and is planned early. Threads that never
terminate, and shutdown paths that deadlock, are defects found only in
production if they are not designed up front.

### 63. Test threaded code by forcing it to fail

Treat a spurious failure as a real bug, never as a fluke. Get the
non-threaded code working first, in isolation. Make threaded code
pluggable and tunable in thread count and speed. Run with more threads
than cores, on more than one platform. Instrument the code with yields
and sleeps to provoke rare interleavings, and automate it where you can.

## M. Successive refinement

### 64. Never ship the first draft

A working first draft is where the job starts, not where it ends. Clean
it in small steps, running the tests after every step, until every rule
here holds. When a new feature would make the structure creak, refactor
first, then add the feature. No code is exempt, however respected its
author or however long it has worked.

## N. Smells and heuristics: the checklist

Every entry below is a defect when found. Cite the section letter and
the entry.

### 65. Comments

- **C1** Information that belongs in version control or a tracker
  (changelogs, authors, metadata) inside source.
- **C2** Comments that are obsolete or drift from the code.
- **C3** Comments that repeat what the code already says.
- **C4** Comments that are badly written, vague, or rambling.
- **C5** Commented-out code.

### 66. Environment

- **E1** A build that takes more than one step.
- **E2** A test run that takes more than one step.

### 67. Functions

- **F1** Too many arguments.
- **F2** Output arguments.
- **F3** Flag arguments.
- **F4** Functions never called.

### 68. General

- **G1** More than one language in a single source file.
- **G2** Obvious behavior left unimplemented; the least-surprise rule.
- **G3** Wrong behavior at the boundaries; every boundary is tested.
- **G4** Overridden safeties: ignored warnings, disabled tests, silenced
  checks.
- **G5** Duplication in any form: identical code, repeated conditionals
  on one discriminator, similar algorithms with different details.
- **G6** Code at the wrong level of abstraction: details in the base
  class, high-level policy mixed with low-level mechanism.
- **G7** A base class that knows about its derivatives.
- **G8** Too much information: wide interfaces, many public members,
  exposed internals.
- **G9** Dead code: unreachable branches, unused conditions, never-thrown
  catches.
- **G10** Vertical separation: variables and functions far from their
  use.
- **G11** Inconsistency: similar things done in different ways.
- **G12** Clutter: unused variables, functions, and constructors, and
  pointless comments.
- **G13** Artificial coupling between things that have no business
  knowing each other; a constant or helper in a place that has no reason
  to hold it.
- **G14** Feature envy: a method that manipulates another object's data
  more than its own.
- **G15** Selector arguments that pick a behavior; split the function.
- **G16** Obscured intent: dense expressions, magic numbers, terse names.
- **G17** Misplaced responsibility: code where a reader would not look
  for it.
- **G18** Inappropriate static: a method that should be polymorphic made
  static.
- **G19** Missing explanatory variables where an intermediate name would
  make the computation readable.
- **G20** Function names that do not say what the function does.
- **G21** Not understanding the algorithm: it passes, but the author
  cannot explain why.
- **G22** Logical dependencies that are not physical: a module assumes
  a fact about another without asking it.
- **G23** Conditionals where polymorphism belongs; one switch per type
  at most.
- **G24** Ignored conventions: naming, placement, and structure the team
  agreed on.
- **G25** Magic numbers and strings instead of named constants.
- **G26** Imprecision: assuming the first match is the only match, using
  floating point for money, ignoring concurrency on shared updates,
  skipping null and error checks the code depends on.
- **G27** Convention where structure would enforce it.
- **G28** Naked conditionals; extract them into predicate functions.
- **G29** Negative conditionals when a positive reads better.
- **G30** Functions that do more than one thing.
- **G31** Hidden temporal couplings between calls.
- **G32** Arbitrary structure: things placed with no reason a reader can
  infer.
- **G33** Boundary conditions repeated rather than encapsulated in one
  named variable.
- **G34** Functions that drop more than one level of abstraction.
- **G35** Configurable data buried at low levels instead of surfaced at
  the top.
- **G36** Transitive navigation through collaborators; the Law of
  Demeter.

### 69. Names

- **N1** Names that are not descriptive.
- **N2** Names at the wrong level of abstraction for the reader.
- **N3** Standard nomenclature ignored where it exists (pattern names,
  platform conventions).
- **N4** Ambiguous names that could mean two things.
- **N5** Short names for long scopes.
- **N6** Encodings in names.
- **N7** Names that hide side effects (`getX` that creates X).

### 70. Tests

- **T1** Insufficient tests: anything that could break is not tested.
- **T2** No coverage tool in use.
- **T3** Trivial tests skipped.
- **T4** An ignored test that is really an unanswered question.
- **T5** Boundary conditions untested.
- **T6** A bug fixed without tests hunting its neighbors.
- **T7** Patterns of failure ignored instead of read.
- **T8** Patterns of coverage ignored instead of read.
- **T9** Slow tests.

## How to use this

- **Writing code**: hold each unit to every rule as you write it. Write
  the test first. Refactor before you move on. Leave the file cleaner.
- **Reviewing code**: walk the diff section by section. Report each
  violation with the rule number, the location, and the fix. Rank by
  severity: correctness and duplication first, then design (objects,
  boundaries, classes, systems), then functions, then names, comments,
  and formatting.
- **Refactoring**: fix violations without changing behavior, in small
  steps, tests green after each.

The `clean-code` agent writes code under this standard test-first and
refines it against the checklist before returning, and reviews a diff
against it on demand. `/clean-code <task>` builds; `/clean-code review`
reviews your current changes.
