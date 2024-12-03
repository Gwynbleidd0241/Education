package notifications

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

type MailSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (m *MailSender) SendEmail(to string, subject string, body string, attachments []string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	for _, attachment := range attachments {
		if _, err := os.Stat(attachment); os.IsNotExist(err) {
			log.Printf("Attachment file does not exist: %s", attachment)
			return err
		}
		message.Attach(attachment)
	}

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
	dialer.SSL = false
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("Failed to send email: %s", err)
		return err
	}

	log.Printf("Email successfully sent to %s", to)
	return nil
}

func NewMailSender(host string, port int, username, password string) *MailSender {
	return &MailSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}
