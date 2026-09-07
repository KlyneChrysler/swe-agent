I could not write scratch files outside the repo (sandbox denied) and review mode forbids editing this repo, so every finding below is substantiated by reading the code and the test file only; I have marked nothing as empirically executed except `go test ./...`, which passes (`ok example.com/fixture 0.343s`).

Reviewed: `go.mod`, `roster.go`, `roster_test.go` (all newly added, 80 lines).

---

## Correctness

```
[rule 68 G26] roster.go:23 - TimeScheduled sums End.Sub(Start) with no invariant that End follows Start, so a Shift built with End before Start silently subtracts from the total; Shift's fields are public and there is no constructor anywhere that could reject it.
  Fix: build Shift through a factory that returns an error (or panics at the composition root) when End is not after Start, and make Roster hold only validated shifts, so TimeScheduled can never return a duration below zero.
```

```
[rule 68 G2] roster.go:40 - ByDay keys every shift solely by shift.Start, so a shift from 22:00 on the 1st to 06:00 on the 2nd appears only under "2024-03-01" and the 2nd's bucket loses six scheduled hours, which is not what the name ByDay promises to a reader.
  Fix: decide the rule explicitly - either split a spanning shift across the days it covers, or rename to ByStartDay - and pin the decision with a test.
```

```
[rule 68 G26] roster.go:40 - The bucket key is produced by Format on the shift's own time.Time, so the same instant carried in two locations lands in two different days; the grouping silently depends on the location each caller happened to attach.
  Fix: normalize to one location inside Roster (store it as a field set at construction) before deriving the key.
```

```
[rule 37] roster.go:16-18 - Roster accepts a []Shift straight into its field with nothing checked; data arriving from a store or a request is trusted as-is, including a nil-interval or inverted shift.
  Fix: validate on the way in, at the one place a Roster is created.
```

## Design

```
[rule 50] roster.go:16-18 - Roster's only field is unexported and there is no exported constructor, so no caller outside package fixture can build a non-empty Roster; TimeScheduled, Overlaps and ByDay are unreachable in practice [rule 67 F4].
  Fix: add NewRoster(shifts []Shift) (Roster, error) as the single construction point, and copy the slice so the caller cannot mutate the roster afterwards.
```

```
[rule 46] roster_test.go:15,22,29 - All three tests build the object with Roster{s: ...}, reaching into the unexported field, so the tests are welded to the representation: renaming or replacing `s` breaks every test even though no behavior changed.
  Fix: construct through the constructor above; the tests then exercise the same API a caller has.
```

```
[rule 29] roster.go:7-14 - Shift is a hybrid: it exposes its data (public Start and End) and also carries behavior (overlaps), taking the weaknesses of both forms [rule 31].
  Fix: choose one. Since Roster is the object, make Shift a plain data structure and move the interval test to a function over two shifts; or hide Start/End and let Shift answer questions about itself.
```

```
[rule 28] roster.go:37 - ByDay returns map[string][]Shift, a stringly-typed representation whose keys only make sense if the caller knows the private dayLayout constant; roster_test.go:30 proves the leak by hardcoding "2024-03-01" [rule 68 G22, G35].
  Fix: return a day-keyed type the caller can construct (a Date value or a ShiftsOn(day time.Time) []Shift query) instead of forcing every caller to reproduce the format string.
```

## Functions

```
[rule 13] roster.go:37-44 - ByDay mixes two levels: the high-level "group shifts" loop and the low-level detail of how a day key is spelled.
  Fix: extract dayKey(shift Shift) string (or the day-value equivalent) and leave the loop reading as grouping alone.
```

## Tests

```
[rule 70 T5] roster_test.go:21-26 - The only Overlaps test asserts the true case; the false path at roster.go:34 and the false result of Shift.overlaps are never executed, and the decisive boundary - two shifts that touch at exactly 12:00, which the strict Before comparisons treat as no clash - is untested [rule 68 G3].
  Fix: add a test that a 12:00-14:00 candidate does not clash with a 09:00-12:00 shift, and one that a wholly separate day does not clash.
```

```
[rule 70 T1] roster_test.go:14-33 - No test covers the empty roster (zero total, no clash, empty grouping), a zero-length shift, an inverted shift, or a shift spanning midnight; three happy-path tests are the whole suite.
  Fix: one test per boundary, each asserting a single concept.
```

```
[rule 70 T5] roster_test.go:28-33 - ByDay is checked for one bucket's length only; nothing asserts how many days were produced, so a change that dropped the second day's shifts entirely would still pass.
  Fix: assert the day count as well, in its own test.
```

## Names / formatting

```
[rule 4] roster.go:17 - The field `s` is a single letter for a name whose scope is the whole type and every method on it [rule 69 N5].
  Fix: rename to shifts.
```

```
[rule 16] roster_test.go:8,15 - shiftAt takes three positional ints, so call sites read shiftAt(1, 9, 12) and the reader must open the helper to learn which number is the day [rule 68 G16].
  Fix: give the helper a name and shape that reads at the call site, for example shiftOnMarch(day).from(9).to(12), or take a small value type.
```

```
[rule 26] roster_test.go:8 - The parameter list mixes grouping styles in one signature: `day int, startHour, endHour int` [rule 68 G11].
  Fix: write `day, startHour, endHour int`.
```

---

Highest-value fix first: add the validating `NewRoster` constructor and route the tests through it — it closes the negative-duration defect, makes the exported API actually usable, and unhooks the tests from the field name in one change.
