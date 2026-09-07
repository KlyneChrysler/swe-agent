package fixture

import "fmt"

type Notifier interface {
	Notify(message string) error
}

type Channel string

const (
	Email Channel = "email"
	SMS   Channel = "sms"
)

type emailNotifier struct {
	address string
}

func (n emailNotifier) Notify(message string) error {
	_, err := fmt.Printf("To: %s\nSubject: notice\n\n%s\n", n.address, message)
	return err
}

type smsNotifier struct {
	number string
}

func (n smsNotifier) Notify(message string) error {
	_, err := fmt.Printf("sms %s: %s\n", n.number, message)
	return err
}

func NewNotifier(channel Channel, target string) (Notifier, error) {
	switch channel {
	case Email:
		return emailNotifier{address: target}, nil
	case SMS:
		return smsNotifier{number: target}, nil
	default:
		return nil, fmt.Errorf("new notifier: unsupported channel %q", channel)
	}
}
