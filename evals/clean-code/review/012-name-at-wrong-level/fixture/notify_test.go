package fixture

import "testing"

type spyNotifier struct {
	recipient string
	body      string
}

func (s *spyNotifier) SendSMTP(recipient, body string) error {
	s.recipient = recipient
	s.body = body
	return nil
}

func TestNotifyOverdueSendsInvoiceMessage(t *testing.T) {
	spy := &spyNotifier{}
	if err := NotifyOverdue(spy, "ada@example.com", "INV-7"); err != nil {
		t.Fatalf("NotifyOverdue: %v", err)
	}
	if spy.recipient != "ada@example.com" || spy.body != "invoice INV-7 is overdue" {
		t.Fatalf("sent %q to %q", spy.body, spy.recipient)
	}
}
