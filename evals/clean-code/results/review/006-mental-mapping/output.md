Reviewed the three added files (`go.mod`, `loan.go`, `loan_test.go`). Verification notes: `go test ./...` passes (`ok example.com/fixture 0.288s`). The boundary values below were derived by hand from the exact expressions in `loan.go:8-10`; I was not permitted to write a scratch program outside the repo to execute them, so they are stated as arithmetic, not as observed output.

## Correctness / duplication

```
[rule 68 G3] loan.go:10 - A zero interest rate returns NaN instead of a payment: r=0 makes q=0 and g=math.Pow(1+0,n)=1, so the expression is p*0*1/(1-1) = 0/0 = NaN, returned silently to the caller.
  Fix: handle the zero-rate case as its own named path (principal divided by the term) before the amortization formula, and add a test asserting MonthlyPayment(120000, 0, 360) == 333.33.
```

```
[rule 68 G3] loan.go:10 - A zero term divides by zero: n=0 makes g=(1+q)^0=1, so the denominator g-1 is 0 and the numerator p*q*g is non-zero, yielding +Inf.
  Fix: reject a non-positive term at entry with an error that names the operation and the offending value.
```

```
[rule 68 G26] loan.go:7-10 - No input is checked, so hostile values return plausible-looking numbers rather than a reported failure: a negative term makes g<1, the denominator negative, and the function returns a negative "payment"; a negative principal does the same.
  Fix: return (float64, error) and validate principal, rate, and term, with messages such as "monthly payment for term -360 months: term must be positive".
```

```
[rule 19] loan.go:7 - The function has no failure channel at all, so every invalid input is expressed as NaN, Inf, or a wrong-signed float that propagates into caller arithmetic undetected.
  Fix: make failure explicit in the signature (Go's error return is the local idiom for the exception this rule requires); no caller should have to test the result with math.IsNaN.
```

```
[rule 20] loan_test.go:11-12 - The expected value 599.55 is written twice, once in the comparison and once in the failure message; change one and the message lies about what was wanted.
  Fix: const want = 599.55, then compare against want and print want.
```

## Design

```
[rule 16] loan.go:7 - Three arguments, and principal, annual rate, and term are one concept (the terms of a loan) passed as loose scalars.
  Fix: introduce a LoanTerms data structure holding Principal, AnnualRate, TermMonths and take it as the single argument.
```

```
[rule 68 G26] loan.go:7 - Money is carried as float64, so results are inexact by construction and the test must compare with a 0.01 tolerance to pass at all.
  Fix: keep float64 only inside the formula and round the result to whole cents at the boundary of the function, so callers receive an exact money value.
```

## Functions

```
[rule 68 G16] loan.go:10 - The return expression p * q * g / (g - 1) is dense and unexplained; nothing in the line says it is the standard amortization formula.
  Fix: extract named intermediates, e.g. growth := compoundGrowth(monthlyRate, termMonths) and return principal * monthlyRate * growth / (growth - 1) with the operands named.
```

## Names / comments / formatting

```
[rule 1] loan.go:7 - The parameters p, r, and n are single letters covering the whole function; none reveals intent.
  Fix: rename to principal, annualRate, termMonths.
```

```
[rule 2] loan.go:7 - r and n disinform about their units: r must be an annual rate (the body divides it by monthsPerYear) and n must be a count of months (the test passes 360), but a caller reading only the signature cannot tell annual from monthly or months from years, and passing 30 for a thirty-year loan compiles and returns nonsense.
  Fix: put the unit in the name: annualRate and termMonths.
```

```
[rule 4] loan.go:8-9 - The locals q and g are unpronounceable and unsearchable.
  Fix: rename to monthlyRate and growthFactor.
```

```
[rule 68 G25] loan_test.go:10-11 - The literals 100000, 0.06, 360, and 599.55 appear bare, so the reader must infer which is principal, which is rate, and which is term.
  Fix: declare named constants principal, annualRate, termMonths, and want in the test body.
```

```
[rule 70 T1] loan_test.go:8-13 - One happy-path test is the entire suite; nothing covers zero rate, zero or negative term, negative principal, a one-month term, or a very large term.
  Fix: add one test per boundary listed above, each asserting a single concept.
```

```
[rule 70 T2] repository - No coverage tool is configured, so the gap above is invisible to the next reader.
  Fix: record go test -cover as the standard one-step test command.
```

Highest-value fix first: `loan.go:10` returns NaN for a zero-interest loan — special-case the zero rate and cover it with a test before anything else in this list.
