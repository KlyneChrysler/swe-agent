Tests pass (`go test ./...` → `ok example.com/fixture 0.289s`). Three files, all new. Findings below.

## Correctness / duplication

```
[rule 55 / rule 68 G5] sku.go:5,9 - The category width is stated twice and nothing keeps the two in agreement: the named constant `categoryLength = 2` and the literal `{2}` inside `skuPattern`. Changing a category to three letters requires editing both, and editing only the regex yields a silently truncated prefix from `Category`.
  Fix: derive one from the other — build the pattern with `regexp.MustCompile(fmt.Sprintf(`^[A-Z]{%d}-[0-9]{%d}$`, categoryLength, sequenceLength))` — so a single named constant governs both the match and the slice.
```

```
[rule 68 G25] sku.go:9 - The five-digit sequence width is an unnamed magic literal inside the pattern, while its sibling (the category width) is named; the reader cannot tell why one is a constant and the other is not.
  Fix: introduce `sequenceLength = 5` alongside `categoryLength` and use both in the pattern.
```

## Design

```
[rule 35 / rule 36] sku.go:19-24 - `Category` returns "" as a sentinel for the ill-formed case, the string equivalent of returning null: every caller must defend with an emptiness check, and a caller that forgets gets an empty category flowing into the catalog instead of a failure at the point of the defect. The doc comment at sku.go:17-18 exists only to document this hazard.
  Fix: return `(string, bool)` in the Go idiom, or `(string, error)` with a message naming the operation and the offending input (rule 33), so the ill-formed case cannot be read as a valid category.
```

## Tests

```
[rule 70 T5 / rule 68 G3] sku_test.go:6-11 - The boundary set is incomplete on the side that matters most: input that is too long is never tested. `"EL-000421"`, `"EL-00042 "`, `"EL-00042\n"` (Go's `$` is end-of-text, not end-of-line, so this must be rejected — the test should pin that), and `"xEL-00042"` are all unasserted. A regression to an unanchored pattern would leave this suite green.
  Fix: add the over-long, whitespace-padded, newline-suffixed, and prefixed inputs to the case table, each expecting `false`.
```

```
[rule 70 T1] sku_test.go:19-29 - `Category` is tested on one valid SKU and one obviously-invalid string, so the slice boundary itself is unpinned. A pattern of `[A-Z]{3}` with `categoryLength` left at 2 — exactly the drift the duplication above invites — passes both tests.
  Fix: assert `Category` against a second, distinct category (e.g. `"ZZ-99999"` → `"ZZ"`) and against a near-miss such as `"EL-0042"`.
```

```
[rule 43] sku_test.go:6-16 - The case table is a `map[string]bool`, which Go iterates in random order; the test remains correct but its failure output is order-dependent and the cases cannot be named or run individually.
  Fix: use a slice of structs with a `name` field and `t.Run(name, ...)` subtests, giving deterministic order and one reported concept per case.
```

Nothing to report on names, comments, or formatting: the doc comments on the two exported functions and the regex are the documented exceptions under rule 23, and the file follows the newspaper order with `IsSKU` above its caller-adjacent `Category`.

**Fix first:** collapse the duplicated width literals in `sku.go:5,9` to one named constant each, since that duplication is what makes the thin `Category` tests capable of hiding a real truncation bug.
