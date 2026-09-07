package fixture

import (
	"fmt"
	"io"
)

type Mailer interface {
	Send(to, body string) error
}

type writerMailer struct {
	out io.Writer
}

func NewWriterMailer(out io.Writer) Mailer {
	return writerMailer{out: out}
}

func (m writerMailer) Send(to, body string) error {
	_, err := fmt.Fprintf(m.out, "To: %s\n\n%s\n", to, body)
	return err
}
