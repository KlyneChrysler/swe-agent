## Review: uncommitted changes (3 new files, 51 lines)

Tests currently pass: `go test ./...` → `ok example.com/fixture 0.288s`. That is the only thing verified; everything below is a defect the green suite hides.

### Correctness / duplication

```
[rule 36] directory.go:21-23 - Customer returns Go's zero-value Customer for an unknown id, so a caller
  gets Customer{"", "", ""} and cannot distinguish "not found" from a real record; this is returning
  null wearing a struct. Confirmed by map-index semantics and by the absence of any lookup test for a
  missing id.
  Fix: return a found flag as a second result (`(Customer, bool)`, the Go idiom), or return an
  explicitly designed special-case Customer whose behavior is correct for "absent" (rule 35). Pick one
  and apply it to every lookup.

[rule 20] directory.go:21-31 - three near-identical bodies that differ only by the field selected:
  `return d.customers[id]`, `return d.customers[id].Name`, `return d.customers[id].Email`. Same lookup
  algorithm, three copies, so the missing-id defect above must be fixed in three places.
  Fix: delete CustomerInfo and CustomerData; a caller that wants a name calls `directory.Customer(id)`
  and reads `.Name` (legal, Customer is a data structure - rule 31, and rule 30 exempts chains through
  data structures).

[rule 68 G26] directory.go:14-17 - NewDirectory silently drops customers: two inputs sharing an ID
  leave only the last, and a customer with an empty ID is stored under key "". The constructor assumes
  uniqueness and non-emptiness without asking for it or enforcing it.
  Fix: reject a duplicate or empty ID at construction with an error naming the offending ID (rule 33),
  or document the last-wins rule as deliberate and pin it with a test.

[rule 68 G3] directory_test.go:5-17 - only the happy path is tested. No test for: unknown id, empty
  directory (`NewDirectory(nil)`), duplicate IDs, an empty-string id, a customer whose ID is empty.
  Fix: one test per boundary, before any further work on this file.

[rule 70 T5] directory_test.go:5-17 - same gap stated as a test smell: every boundary condition of the
  only behavior in the package is unasserted.
  Fix: as above.
```

### Design

```
[rule 28] directory.go:25-31 - CustomerInfo and CustomerData are field-level accessors that publish
  Customer's representation through Directory; Directory's interface grows one method per Customer
  field, so adding a Customer field pressures a new Directory method.
  Fix: expose the single operation Directory owns - look a customer up - and let callers read the data
  structure.

[rule 68 G14] directory.go:25-31 - feature envy: both methods manipulate Customer's data (`.Name`,
  `.Email`) and touch nothing of Directory's own beyond the map they index.
  Fix: remove them; the behavior belongs with the caller or with Customer.
```

### Names / comments / formatting

```
[rule 3] directory.go:21,25,29 - `Customer`, `CustomerInfo`, `CustomerData` are distinctions without a
  difference; `Info` and `Data` are noise words and the reader cannot tell from the names which one
  returns a name and which returns an email. This is the book's own getAccount/getAccountInfo/
  getAccountData example verbatim.
  Fix: keep `Customer(id)` alone. If per-field lookups genuinely earn their place later, name them for
  what they return: `CustomerName`, `CustomerEmail`.

[rule 69 N4] directory.go:25,29 - `CustomerData` is ambiguous: it could plausibly mean the whole
  record, and in fact returns the email.
  Fix: as above.

[rule 43] directory_test.go:5-17 - one test function checks three concepts (record lookup, name lookup,
  email lookup) with three Fatalf assertions; the name promises only "looks up by ID".
  Fix: one test per concept, each with one assertion as the target - or, once the pass-through
  accessors are deleted, one test that the lookup returns the stored record.
```

**Fix this first:** delete `CustomerInfo` and `CustomerData` (rules 3, 20, 28, G14 all collapse with them), then change the surviving `Customer(id)` to report a missing id instead of handing back a zero-value struct, and add the unknown-id and duplicate-id tests that prove it.
