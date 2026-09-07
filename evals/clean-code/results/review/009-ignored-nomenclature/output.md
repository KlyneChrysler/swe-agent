## Review: uncommitted changes (`go.mod`, `undo.go`, `undo_test.go`)

Verified: `go vet ./...` clean, `go test ./...` passes, `gofmt -l .` clean, coverage 100.0% of statements. Whole codebase is these three files; no duplication exists to find (grep for `Put`/`Take`/`Size` returns only the definitions and the test).

### Correctness / duplication

```
[rule 68 G3] undo.go:11-18 - Take's effect on the pile is never verified beyond the element count, so a wrong removal passes the suite.
  Fix: add a test that Takes twice and asserts "second" then "first".
```
Substantiation: replace line 17 with `p.items = p.items[1:]` (remove the front while still returning the last element). With `["first","second"]`, Take returns `"second"` and leaves `["second"]`, length 1 — both existing assertions still hold and the suite stays green while the stack is broken.

```
[rule 70 T5] undo_test.go:1-22 - The boundaries of a LIFO container are untested: Take called more times than Put, Put after the pile has been emptied, Size on a fresh pile, and Put("") followed by Take (must yield "", true, not the empty-pile ("", false)).
  Fix: one test per boundary, each named for the behavior it pins.
```

```
[rule 70 T8] undo_test.go:1-22 - 100% statement coverage is being read as safety; every statement runs, yet the removal order and the empty-string element are unchecked.
  Fix: judge the suite by behaviors pinned, not by lines executed; add the boundary tests above.
```

### Design

```
[rule 18] undo.go:11 - Take is a command and a query at once: it mutates the pile and returns a status flag about whether the mutation happened, so every caller opens a conditional around it.
  Fix: add the query `func (p *Pile) IsEmpty() bool` and let Take answer only with the item (panicking or returning an error on misuse), or state explicitly that Go's comma-ok idiom is the accepted exception here [rule 69 N3].
```

```
[rule 68 G26] undo.go:17 - Reslicing with p.items[:last] leaves the removed string in the backing array, so a taken item is retained for the life of the pile; a long-lived pile that churns leaks every element it ever held.
  Fix: clear the slot before reslicing - `p.items[last] = ""` then `p.items = p.items[:last]`.
```

```
[rule 57] undo.go:21 - Size has no production caller; its only reference in the repository is the assertion at undo_test.go:12, so it is API built for a caller that does not exist.
  Fix: drop Size unless Pile is a published API, and assert the pile's state through Take instead.
```

### Functions

```
[rule 43] undo_test.go:5-14 - One test asserts two concepts: the item Take returns and the size the pile is left at.
  Fix: split into TestTakeReturnsMostRecentPut and TestTakeRemovesTheItemItReturns.
```

```
[rule 68 G28] undo_test.go:9 - The naked compound condition `!ok || item != "second"` makes the reader work out which of two failures fired from a combined message.
  Fix: assert ok and the item separately, each with its own message.
```

### Names / comments / formatting

```
[rule 2] undo.go:1 - The file is named undo.go but nothing in it concerns undo; the headline of the file disinforms the reader about its contents [rule 25].
  Fix: rename the file to match what it holds (stack.go), or introduce the undo concept the name promises.
```

```
[rule 9] undo.go:3-22 - Pile/Put/Take is a third vocabulary for a structure computer science already names; the standard nomenclature is Stack with Push and Pop, and Take in particular is ambiguous about which end it takes from [rule 69 N3, N4].
  Fix: rename Pile to Stack, Put to Push, Take to Pop.
```

```
[rule 3] go.mod:1, undo.go:1 - `fixture` as module path and package name is a noise word that tells the reader nothing about the domain [rule 69 N1].
  Fix: name the module and package for the domain the pile serves.
```

**Highest-value fix first:** add the two-successive-Takes test — it is the one gap that currently lets a broken removal ship green.
