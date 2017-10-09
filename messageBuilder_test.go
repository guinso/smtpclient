package smtpclient

import "testing"
import "strings"

func TestMailContent_String(t *testing.T) {
	mailContent := MailContent{
		Headers: []string{
			"Content-Type: text/plain; charset=\"utf-8\"",
			"Content-Transfer-Encoding: 7bit",
		},
		Content: "asd qwe",
	}

	str, err := mailContent.String()
	if err != nil {
		t.Error(err)
	}

	expected := "Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"Content-Transfer-Encoding: 7bit\r\n" +
		"\r\nasd qwe\r\n"

	if strings.Compare(str, expected) != 0 {
		t.Errorf("\r\nGenerated message content not tally with expectation:\r\n"+
			"==Expected==\r\n%s\r\n\r\n"+
			"==Actual==\r\n%s", expected, str)
	}
}

func TestMailMessage_String(t *testing.T) {
	mail := MailMessage{
		Subject: "Sample Email",
		From:    "me@mail.com",
		To: []string{
			"friendA@mail.com",
			"friendB@mail.com",
		},
		Cc: []string{
			"friendC@mail.com",
			"friendD@mail.com",
		},
	}

	mail.AddTextContent("abcdefg")
	mail.AddHTMLContent("<h1>hi</h1>")

	mailMsg, err := mail.String()
	if err != nil {
		t.Error(err)
		return
	}

	expected := "From: me@mail.com\r\n" +
		"To: friendA@mail.com;friendB@mail.com\r\n" +
		"Cc: friendC@mail.com;friendD@mail.com\r\n" +
		"Subject: Sample Email\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: multipart/mixed; boundary=\"**meme34375901-abc\"\r\n\r\n" +

		"--**meme34375901-abc\r\n" +
		"Content-Type: text/plain; charset=\"utf8\"\r\n" +
		"Content-Transfer-Encoding: 7bit\r\n\r\n" +

		"abcdefg\r\n\r\n" +

		"--**meme34375901-abc\r\n" +
		"Content-Type: text/html; charset=\"utf8\"\r\n" +
		"Content-Transfer-Encoding: 7bit\r\n\r\n" +

		"<h1>hi</h1>\r\n"

	if strings.Compare(mailMsg, expected) != 0 {
		t.Errorf("\r\nGenerated mail message not tally with expectation:\r\n"+
			"==Expected==\r\n%s\r\n\r\n"+
			"==Actual==\r\n%s", expected, mailMsg)
	}
}
