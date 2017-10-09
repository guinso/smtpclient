package smtpclient

import "testing"

func TestEmailHelper_SendEmail(t *testing.T) {
	message := MailMessage{
		From: "john@gmail.com",
		To: []string{
			"friedA@yahoo.com",
			"friendB@mail.com",
		},
		Cc: []string{
			"jerry@hotmail.com",
		},
		Subject:  "Sample email from smtpclient library",
		Contents: []MailContent{},
	}
	message.AddHTML("<html><body><h1>Hello</h1><p>world</p></body></html>")
	message.AddAttachment("c:\\log.txt", "log.txt")

	helper := EmailHelper{
		ServerAddress:    "smtp.gmail.com",
		ServerPortNumber: 465,
		NetworkProtocol:  TCP,
		Username:         "john@gmail.com",
		Password:         "your-email-password",
	}

	if err := helper.SendEmail(&message); err != nil {
		t.Error(err)
	}
}
