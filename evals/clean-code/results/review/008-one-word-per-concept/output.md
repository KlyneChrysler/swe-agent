## Review: uncommitted changes (5 new files, `example.com/fixture`)

Tests were run as found: `go test ./...` → `ok example.com/fixture 0.300s`. All three tests pass; that result is exactly what makes the findings below dangerous.

### Correctness / duplication

```
[rule 8] customers.go:12, orders.go:12, products.go:14 - one concept ("look up a record by its key") carries three different verbs: FetchCustomer, RetrieveOrder, GetProduct, so a caller cannot predict any method name from the others.
  Fix: pick one verb for the concept across the package - Find (or Get) - and rename to FindCustomer, FindOrder, FindProduct, including the test names on lookup_test.go:5,12,19.
```
```
[rule 20] customers.go:12-15 and products.go:14-17 - identical algorithm, one map index plus a comma-ok pass-through, differing only in the value type and a key normalization; orders.go:12-19 is the same lookup written a third way as a linear scan.
  Fix: collapse into one generic keyed lookup (`type index[K comparable, V any]` with a `find(K) (V, bool)`), give Order a map keyed by Number, and let the three types delegate to it.
```
```
[rule 68 G27] products.go:15 - the case-insensitivity invariant lives only on the read side: the lookup upper-cases the argument, but nothing upper-cases the stored key, so a Catalog whose map holds "el-1" can never return that product, no matter what the caller passes.
  Fix: normalize in one place a constructor owns - `NewCatalog(products []Product)` that stores `strings.ToUpper(product.SKU)` - so the structure enforces the rule instead of a convention nobody can see.
```
```
[rule 68 G26] orders.go:13-16 - the scan assumes the first match is the only match; two entries with Number 7 silently resolve to whichever was appended first, and no test or type constraint says order numbers are unique.
  Fix: key orders by Number in a map built by a constructor that rejects a duplicate number with an exception naming the number.
```
```
[rule 68 G11] customers.go:12 vs products.go:14 - inconsistent behavior for the same operation: customer ids are matched exactly, SKUs case-insensitively, and nothing in either name or signature announces the difference.
  Fix: decide the normalization policy per key type and state it in a named key type (`type SKU string` that normalizes on construction), so the difference is visible at the call site.
```

### Design

```
[rule 50] customers.go:8, orders.go:8, products.go:10 - all three types have only unexported fields and no factory, so no code outside this package can construct a non-empty CustomerBook, OrderLog, or Catalog; the exported lookups are unreachable in practice.
  Fix: add NewCustomerBook, NewOrderLog, NewCatalog taking the records they index, and make them the single place the maps are built.
```
```
[rule 46] lookup_test.go:6,13,20 - the tests reach past the API into the unexported fields byID, entries, and bySKU, so every test is welded to the storage representation and would break on the map/slice change recommended above.
  Fix: build fixtures through the constructors above; the tests then exercise the same path production callers use.
```
```
[rule 47] customers.go:8, orders.go:8, products.go:10 - CustomerBook, OrderLog, Catalog are three unrelated nouns for one role, an indexed collection of records, forcing the reader to learn three vocabularies for one idea.
  Fix: name them by their shared role (CustomerIndex/OrderIndex/ProductIndex, or Customers/Orders/Products) and keep the suffix consistent.
```

### Tests

```
[rule 70 T5] lookup_test.go:5-24 - every boundary is untested: an unknown id, an empty id, a nil or empty collection, a duplicate order number, an already-upper-case SKU, and a stored lower-case SKU (the case that is actually broken, see products.go:15).
  Fix: add a test per boundary, starting with the two that fail today - lookup of a lower-case stored SKU, and lookup of a duplicated order number.
```
```
[rule 43] lookup_test.go:7, 14, 21 - each test checks two concepts in one compound condition (`!ok || customer.Name != "Ada"`), so a failure message cannot tell the reader whether the record was missing or the wrong record came back.
  Fix: split into a found-check and a value-check, each with its own message, or one test per concept.
```
```
[rule 70 T3] lookup_test.go - the miss path (`ok == false`) of all three lookups is never asserted, though it is half of every signature's contract.
  Fix: one "returns not found for an unknown key" test per type.
```

### Names

```
[rule 2] lookup_test.go:13 - the variable `log` shadows the standard library package name, disinforming any reader who expects logging and breaking the file the moment `log` is imported.
  Fix: rename to `orders`.
```
```
[rule 69 N4] orders.go:9 - `entries` names the storage generically while its siblings name the key (`byID`, `bySKU`), so the same field role reads three ways across the package.
  Fix: rename to `byNumber` once orders are keyed by number.
```

**Highest-value fix first:** products.go:15 — normalize the SKU where the catalog is constructed, not only where it is read, and add the failing test for a lower-case stored key; that single defect makes products silently unfindable and the current happy-path test hides it.
