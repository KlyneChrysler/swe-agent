package fixture

import "fmt"

type Notifier interface {
	SendSMTP(recipient, body string) error
}

type ConsoleNotifier struct{}

func (ConsoleNotifier) SendSMTP(recipient, body string) error {
	_, err := fmt.Printf("to %s: %s\n", recipient, body)
	return err
}

func NotifyOverdue(notifier Notifier, recipient, invoiceID string) error {
	return notifier.SendSMTP(recipient, "invoice "+invoiceID+" is overdue")
}
