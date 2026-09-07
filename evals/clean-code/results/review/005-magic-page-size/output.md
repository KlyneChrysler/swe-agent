## Review: uncommitted changes (`go.mod`, `page.go`, `page_test.go`)

Baseline: `go test ./...` → `ok example.com/fixture 0.262s`; `go vet ./...` → clean. All findings below are from reading the code; the panic in the first finding is derived from Go slice semantics (`low` index is not clamped and must be `>= 0`), not from an executed test, because verifying it would require adding a file to your tree.

### Correctness / duplication

```
[rule 68 G3] page.go:8-16 - a negative page index panics instead of returning an empty page: with page = -1, page*20 = -20 is not > len(items), so end = -20+20 = 0 and the function evaluates items[-20:0], which panics with "slice bounds out of range [-20:]".
  Fix: compute start := page * pageSize, clamp it with `if start < 0 || start >= len(items) { return Page{} }`, and drop the current > guard.

[rule 70 T5] page_test.go:12-31 - every boundary of the only argument pair is untested: negative page, empty/nil items, a page index landing exactly on len(items) (40 items, page 2), and HasNext on the past-end page (only its length is asserted).
  Fix: add one test per boundary, including a test asserting PageOf(items(25), -1) returns an empty page rather than panicking.

[rule 70 T1] page_test.go:5-11 - the fixture writes the identical string "item" into every element, so no test can observe which window was returned; a body that returned the last 5 items for page 1 (items[len-5:]) would pass every assertion here.
  Fix: fill with distinct values (fmt.Sprintf("item-%d", index)) and assert page.Items[0] == "item-20" on the second page.

[rule 68 G5] page.go:9,11,15 - the page size 20 is written as a bare literal four times and the same length boundary len(items) is compared three times, so changing the page size means finding every copy.
  Fix: const pageSize = 20, used once per expression.
```

### Design

```
[rule 68 G35] page.go:9 - the page size, the one piece of configurable data in this package, is buried inside the lowest-level expression instead of being surfaced as a named top-level constant or a parameter.
  Fix: declare it as a package-level constant at the top of the file, above PageOf.

[rule 68 G33] page.go:9-14 - the boundary condition "where does this page start and stop relative to len(items)" is repeated in three separate comparisons rather than encapsulated in named variables; the early return exists only to hide the panic the general path would otherwise cause.
  Fix: name start and end once (start := min(page*pageSize, len(items)), end := min(start+pageSize, len(items))), after which the special-case return disappears.
```

### Functions

```
[rule 68 G19] page.go:9,11,15 - page*20 is recomputed three times with no name for what it is, so the reader has to re-derive "the index of the first item on this page" at each site.
  Fix: start := page * pageSize, used throughout.
```

### Names / comments / formatting

```
[rule 8] page.go:8 and page_test.go:13,20 - the word "page" carries two concepts: a page number in the PageOf parameter and a Page value in every test local, so `page` in a test and `page` in production mean different things.
  Fix: rename the parameter pageNumber (or pageIndex) and keep `page` for the Page value.

[rule 8] page.go:8 vs page_test.go:5 - "items" also names two concepts: the slice of strings in production and a test function that builds one.
  Fix: rename the helper itemsNumbered or buildItems, which is also a verb phrase as rule 7 requires of a function.

[rule 43] page_test.go:14,21 - each test checks two facts in one compound condition joined by ||, so a failure reports both values and leaves the reader to work out which half broke.
  Fix: assert the item count and HasNext in two separate checks with their own messages.
```

Highest-value fix first: clamp the start index in `PageOf` so a negative page returns an empty page instead of panicking, and add the failing test for it before the fix.
