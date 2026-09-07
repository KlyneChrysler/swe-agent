Reviewed the three new files (`address.go`, `address_test.go`, `go.mod`) plus the whole repository for context — the repo contains nothing else, so no cross-file duplication exists to report. Verified state: `go test ./...` → `ok example.com/fixture 0.276s`; `go vet ./...` clean; `gofmt -l .` clean.

## Findings

**Correctness / duplication**

```
[rule 68 G3 / rule 70 T5] address_test.go:5-17 - The only two tests cover one happy path each; every boundary of Apply is untested: empty string, whitespace-only string, an already-uppercase line, the zero-value NormalizeAddress{} (empty Country), and a line with tabs or newlines rather than spaces.
  Fix: add a test per boundary, including NormalizeAddress{}.Apply("") == "" and NormalizeAddress{Country: unitedStates}.Apply("   ") == "".
```

```
[rule 68 G26] address.go:13 - Exact-string comparison n.Country == unitedStates silently degrades to trim-only for any input whose case or padding differs ("us", "Us", " US "), and country codes arriving from a store or request are not trusted data (rule 37).
  Fix: normalize the discriminator before comparing, e.g. compare strings.EqualFold(strings.TrimSpace(n.Country), unitedStates), and test "us" explicitly.
```

```
[rule 44 / rule 70 T1] address_test.go:6 - The US test builds its input from the unitedStates constant, so it can never detect a wrong constant value; the wire value "US" is asserted nowhere, while the sibling test at line 13 hard-codes "PH".
  Fix: use the literal NormalizeAddress{Country: "US"} in the test, pinning the value the outside world sends.
```

```
[rule 20 / rule 68 G5] address_test.go:5-17 - The two test bodies are structurally identical, differing only by the country literal, the input's expected result, and the message; the input literal "  12 main st " is duplicated, and each expected value is written twice (comparison and Fatalf message).
  Fix: collapse into one table-driven test (Go's standard nomenclature, rule 9 / N3) with fields country, line, want, and a single t.Errorf("Apply(%q) = %q, want %q", ...).
```

**Design**

```
[rule 29 / rule 31] address.go:7-9,11 - NormalizeAddress is a hybrid: it exposes the public field Country and also carries the behavior method Apply, so it is half data structure and half object and gets the weaknesses of both.
  Fix: pick one - make the field unexported (country string) and construct through NewAddressNormalizer(country string), so callers manipulate the essence rather than the representation.
```

```
[rule 28] address.go:8 - Country is raw representation leaked into the type's public surface; nothing prevents a caller from setting it to "", "usa", or " us " after construction.
  Fix: validate and canonicalize the country code once at construction and keep it private thereafter.
```

**Names / comments / formatting**

```
[rule 7 / rule 69 N3] address.go:7 - The type name NormalizeAddress is a verb phrase; class names are noun phrases and verb phrases belong on methods.
  Fix: rename the type to AddressNormalizer and the method to Normalize, giving AddressNormalizer.Normalize(line) instead of NormalizeAddress.Apply(line).
```

```
[rule 15 / rule 68 G20] address.go:11 - Apply does not say what the function does; the reader must open the body to learn that it trims and may uppercase.
  Fix: name it Normalize (with the renamed type above), so the call site reads as the domain reads.
```

```
[rule 1 / rule 69 N1] address.go:11 - The parameter line is a noise word that names the shape of the string rather than its meaning; the type is about addresses.
  Fix: rename to addressLine.
```

```
[rule 70 T2] repository-wide - No coverage tooling is in use, so coverage patterns cannot be read.
  Fix: make `go test -cover ./...` the documented one-step test command (rule 66 E2).
```

Highest-value fix first: add the missing boundary tests (empty, whitespace-only, zero-value `Country`, and a lowercase `"us"`), because the lowercase-country case is a live defect that the current happy-path suite hides.
