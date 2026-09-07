Tests run: `go test ./...` → `ok example.com/fixture 0.271s` (1 test). Findings below.

## Correctness / duplication

```
[rule 68 G26] billing.go:27 - Integer division truncates the tax, so any subtotal under 9 cents is taxed zero (8 → 8, not 8.96) and the rounding policy is nowhere stated or chosen; every invoice silently under-collects.
  Fix: decide and encode the policy explicitly, e.g. taxCents := (subtotal*billingTaxPercent + percentDenominator/2) / percentDenominator for round-half-up, and name it in a test.
```
```
[rule 70 T5] billing_test.go:5 - The only test is a happy path with two positive lines; the boundaries are untested: an invoice with no lines, a single 1-cent line (the truncation case above), a negative line (refund/credit), and a zero-cent line.
  Fix: add one test per boundary, each asserting one concept.
```
```
[rule 68 G3] billing.go:25 - Behaviour at the boundary is wrong, not just untested: BillingTotal on a refund-only invoice truncates toward zero, so a -8 cent line yields -8 while an +8 cent line also yields 8, making a charge and its refund not cancel out once amounts differ in sign.
  Fix: apply the same rounding rule to the magnitude, then restore the sign, and pin it with a charge-then-refund test that asserts the pair nets to zero.
```

## Design

```
[rule 29] billing.go:13-15 - BillingInvoice is a hybrid: it exposes the field Lines (a mutable slice any caller can append to or truncate) and also carries the behaviour BillingSubtotal/BillingTotal over that same field.
  Fix: pick one. Either make it an object - unexported lines, a constructor NewInvoice(lines ...Line), no exposed slice - or make it a plain data structure and move the arithmetic into functions that take it.
```
```
[rule 31] billing.go:8-11 - BillingLine is correctly a data structure, but it is being read by a method on another type that sums its Cents; the money arithmetic lives outside the type that owns the amount.
  Fix: if Line stays a data structure this is acceptable, but then Invoice must stop being an object (see rule 29 above); do not leave both halves.
```

## Functions

```
[rule 68 G19] billing.go:27 - The tax amount has no name; the reader has to parse subtotal + subtotal*billingTaxPercent/percentDenominator to see that the second term is the tax.
  Fix: extract taxCents := subtotal * billingTaxPercent / percentDenominator (or a taxOn(subtotal int) int helper placed directly under BillingTotal, rule 46) and return subtotal + taxCents.
```

## Names / comments / formatting

```
[rule 10] billing.go:17,25 - Gratuitous context: the receiver is already a BillingInvoice, so BillingSubtotal and BillingTotal read as invoice.BillingTotal() and stutter at every call site.
  Fix: rename to Subtotal() and Total(). Likewise BillingLine/BillingInvoice → Line/Invoice, with the package name carrying the "billing" context.
```
```
[rule 43] billing_test.go:8-13 - One test function checks two concepts, subtotal accumulation and tax application, with two assertions; the name promises only the tax behaviour, so the first assertion is undocumented by the test name.
  Fix: split into TestSubtotalSumsLines and TestTotalAddsTax, one concept and one target assertion each.
```
```
[rule 69 N5] billing_test.go:8,11 - `got` is fine for a two-line scope, but the literals 1000 and 1120 are unexplained magic numbers repeated in both the condition and the message.
  Fix: name them, e.g. const wantSubtotal = 600 + 400 and derive the expected total from the tax rate so the test states the rule rather than a precomputed answer.
```

Highest-value fix first: settle the rounding policy in `BillingTotal` (billing.go:27) and drive it with the 1-cent and refund boundary tests — everything else is structure, this one is money.
