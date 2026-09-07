Reviewed the three staged files (`go.mod`, `invoice.go`, `invoice_test.go`). Tests pass: `go test ./...` → `ok example.com/fixture 0.307s`. I verified each behavioral claim below with a temporary probe test, which I removed; the working tree is exactly as I found it.

## Correctness / duplication

```
[rule 68 G3][rule 70 T5] invoice_test.go:8-23 - the only boundaries that matter are untested: the exact deadline, the inside-grace window, and a fractional day past the deadline; the suite pins day 35 and day 1 only.
  Fix: add tests for now == IssuedAt+Terms+gracePeriod (verified: returns 0), one nanosecond past it (verified: returns 0), day 31 and 32 inside grace, and 35.5 days (truncates to 2). Name what each pins.
```
```
[rule 37][rule 68 G26] invoice.go:16 - `now.Sub(d)` silently saturates at the max int64 Duration, so an Invoice arriving from a store with an unset IssuedAt returns 106751 (verified) instead of failing, and a negative Terms returns 97 days overdue on the issue date (verified).
  Fix: reject the invalid invoice at the boundary - return an error (or refuse to construct an Invoice) when IssuedAt is the zero time or Terms is negative, rather than reporting a number no caller can distinguish from a real one.
```
```
[rule 2][rule 68 G20] invoice.go:5,20 - the function promises days but counts fixed 24-hour spans; in a DST zone the deadline drifts by an hour (verified: an invoice issued 2024-03-01 00:00 America/New_York has a wall deadline of 2024-04-03 01:00 EDT, so the calendar deadline midnight still reports 0).
  Fix: either count calendar days with time.Time date arithmetic, or rename to say what it measures (e.g. `HoursOverdue`/`OverdueSpans`) and state the UTC assumption in the type.
```
```
[rule 20][rule 55] invoice_test.go:9-10 and invoice_test.go:18-19 - identical two-line build (`issued := time.Date(...)`, `inv := Invoice{IssuedAt: issued, Terms: 30 * day}`) repeated verbatim.
  Fix: one helper in the test's domain vocabulary, e.g. `invoiceIssuedOn(issued)` or `netThirtyInvoice()`, per rule 43.
```

## Design

```
[rule 68 G35] invoice.go:7 - `gracePeriod` is business policy frozen as a package constant at the lowest level; no caller can vary it per customer or contract, and no test can exercise a different one.
  Fix: surface it at the composition root - make it a field on Invoice (or a parameter of the calculation) supplied where invoices are built.
```
```
[rule 13][rule 68 G19] invoice.go:14-21 - one body mixes two levels: the low-level deadline arithmetic and the high-level conversion of an overdue span into whole days.
  Fix: extract `paymentDeadline(inv Invoice) time.Time`, and give the subtraction result a name (`overdue := now.Sub(paymentDeadline(inv))`) so the top level reads as policy.
```

## Names / formatting

```
[rule 1][rule 69 N1] invoice.go:15-16 - `d` and `x` name the two concepts the function exists to compute (the payment deadline and the overdue span) and say nothing about either.
  Fix: `deadline` and `overdue`.
```
```
[rule 68 G11] invoice_test.go:11 vs invoice_test.go:20 - two sibling tests express the same step differently: one binds `now :=`, the other inlines `issued.Add(day)` into the call.
  Fix: pick one shape and use it in both.
```

Highest-value fix first: write the boundary tests (exact deadline, inside grace, fractional day) — they are cheap, they pin the behavior every other fix must preserve, and they are the tests that would have exposed the saturation and DST defects above.
