Tests and vet run clean (`go test ./...` → `ok example.com/fixture`, coverage **33.3% of statements** — only `NotifyOverdue` is exercised). Findings below are all against the three new files.

## Correctness / duplication

```
[rule 2] notify.go:11 - ConsoleNotifier.SendSMTP sends nothing over SMTP; it calls fmt.Printf to stdout, so the name states a protocol the implementation does not use.
  Fix: rename the interface method to Send(recipient, body string) error; the transport belongs in the implementation type name (ConsoleNotifier, SMTPNotifier), not in the abstraction's vocabulary.

[rule 5] notify.go:6 - The transport mechanism is encoded into the abstraction the caller sees, so every caller and every future implementation (console, queue, webhook) inherits a name that lies about it.
  Fix: same rename; mark the implementation, never the abstraction.

[rule 33] notify.go:17 - The error from the notifier is returned bare, with no mention of the operation attempted; a caller of NotifyOverdue sees only the transport's message (e.g. an io write error) with no invoice, recipient, or operation to locate the cause.
  Fix: wrap it: return fmt.Errorf("notify overdue invoice %s to %s: %w", invoiceID, recipient, err).
```

## Design

```
[rule 16] notify.go:16 - Three arguments, two of which are adjacent untyped strings; NotifyOverdue(n, "INV-7", "ada@example.com") compiles and silently sends to the wrong place. recipient and invoiceID travel together and are a missing object.
  Fix: introduce a small data structure (type OverdueInvoice struct { ID, Recipient string }) or distinct named types, giving NotifyOverdue(notifier, invoice).

[rule 37] notify.go:16-18 - Neither invoiceID nor recipient is checked before use; an empty invoiceID produces the message "invoice  is overdue" and an empty recipient sends it nowhere. Values crossing into this function are trusted unconditionally.
  Fix: reject empty identifiers with an error that names the operation, and test both cases.

[rule 68 G35] notify.go:17 - The message wording is configurable policy hard-coded at the lowest level, mid-expression, via string concatenation.
  Fix: hoist it to a named constant at the top of the file (const overdueMessageFormat = "invoice %s is overdue") and build the body with fmt.Sprintf.
```

## Tests

```
[rule 70 T5] notify_test.go:16-24 - Only the happy path is tested. The failure branch of NotifyOverdue (notifier returns an error) is never exercised, so nothing pins the behavior the error-handling fix above changes.
  Fix: add a test whose spy returns a sentinel error and assert NotifyOverdue's returned error wraps it and names the invoice.

[rule 70 T1] notify.go:11-14 - ConsoleNotifier has no test at all; coverage confirms it (33.3% of statements, the two uncovered ones are its body).
  Fix: make the destination injectable (an io.Writer field instead of fmt.Printf's implicit stdout) and assert the rendered line against a bytes.Buffer.

[rule 43] notify_test.go:21-22 - Two independent facts, recipient and body, are checked in one compound condition, and the failure message prints only what was sent, never what was expected, so a reader must diff by hand to learn which field broke.
  Fix: split into two checks, each printing got and want.

[rule 70 T5] notify_test.go - No boundary cases: empty invoice id, empty recipient, or the same notifier used for two consecutive calls.
  Fix: one test per boundary once the validation above exists.
```

## Names

```
[rule 8] notify.go:6,16 - Two words for one concept: the interface says Send, the only caller says Notify. Nothing tells a reader whether NotifyOverdue and Send are the same operation at two levels or two different ideas.
  Fix: after the Send rename, keep Notify* for the policy layer and Send for the transport layer deliberately, and use that split consistently as more notifications arrive.
```

**Highest-value fix first:** rename `SendSMTP` to `Send` — the abstraction currently names a protocol that the shipped implementation does not speak, and every future caller and implementer will copy that lie.
