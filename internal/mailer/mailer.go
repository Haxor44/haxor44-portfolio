package mailer

import (
	"fmt"
	"net/smtp"
)

// Mailer sends contact-form notifications over SMTP. It is deliberately
// dependency-free (net/smtp) since a portfolio contact form doesn't need
// more than plain-text notification email.
type Mailer struct {
	Host     string
	Port     string
	User     string
	Password string
	ToEmail  string
}

func New(host, port, user, password, toEmail string) *Mailer {
	return &Mailer{Host: host, Port: port, User: user, Password: password, ToEmail: toEmail}
}

// Configured reports whether enough SMTP settings are present to attempt sending.
func (m *Mailer) Configured() bool {
	return m.Host != "" && m.User != "" && m.Password != "" && m.ToEmail != ""
}

// SendContactNotification emails the site owner about a new contact form
// submission. Callers should check Configured() first if they need to
// distinguish "skipped, SMTP not set up" from "sent"; this always attempts
// to send and only returns an error for an actual send failure.
func (m *Mailer) SendContactNotification(fromName, fromEmail, subject, body string) error {
	if !m.Configured() {
		return fmt.Errorf("mailer: SMTP not configured")
	}

	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)

	if subject == "" {
		subject = "New portfolio contact message"
	}

	msg := fmt.Appendf(nil,
		"To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\nFrom: %s <%s>\r\n\r\n%s\r\n",
		m.ToEmail, m.User, subject, fromName, fromEmail, body,
	)

	if err := smtp.SendMail(addr, auth, m.User, []string{m.ToEmail}, msg); err != nil {
		return fmt.Errorf("mailer: send failed: %w", err)
	}
	return nil
}
