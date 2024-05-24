package response

type EmailMessage struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NewEmailMessage(subject, body string) *EmailMessage {
	return &EmailMessage{
		Subject: subject,
		Body:    body,
	}
}
