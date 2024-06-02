package adapter

import (
	"fmt"
	"net/smtp"
)

type EmailCredentials struct {
	host     string
	port     string
	username string
	password string
	from     string
}

func New(username, password, host, port, from string) *EmailCredentials {
	return &EmailCredentials{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (c *EmailCredentials) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", c.username, c.password, c.host)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body))
	addr := fmt.Sprintf("%s:%s", c.host, c.port)
	return smtp.SendMail(addr, auth, c.from, []string{to}, msg)
}
