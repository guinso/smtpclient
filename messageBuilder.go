package smtpclient

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	boundary = "**meme34375901-abc"
)

//MailMessage email main message body
type MailMessage struct {
	Subject  string
	From     string
	To       []string
	Cc       []string
	Contents []MailContent
}

//MailContent atomic message content
type MailContent struct {
	Headers []string
	Content string
}

//String generate plain text mail content
func (mailContent *MailContent) String() (string, error) {
	if len(mailContent.Content) < 1 {
		summary := ""
		if len(mailContent.Content) >= 10 {
			summary = mailContent.Content[0:9] + "..."
		} else {
			summary = mailContent.Content
		}

		return "", fmt.Errorf("mail content (%s) must has atleast one item", summary)
	}

	return strings.Join(mailContent.Headers, "\r\n") + "\r\n\r\n" + mailContent.Content + "\r\n", nil
}

//AddText add simple text content into mail body
func (mail *MailMessage) AddText(plainText string) {
	mail.Contents = append(mail.Contents, MailContent{
		Headers: []string{
			"Content-Type: text/plain; charset=\"utf8\"",
			"Content-Transfer-Encoding: 7bit",
		},
		Content: plainText,
	})
}

//AddHTML add HTML format text content into mail body
func (mail *MailMessage) AddHTML(htmlText string) {
	mail.Contents = append(mail.Contents, MailContent{
		Headers: []string{
			"Content-Type: text/html; charset=\"utf8\"",
			"Content-Transfer-Encoding: 7bit",
		},
		Content: htmlText,
	})
}

//AddAttachment add file attachment into mail body
func (mail *MailMessage) AddAttachment(filePath string, filename string) error {

	rawFile, fileErr := ioutil.ReadFile(filePath)
	if fileErr != nil {
		return fileErr
	}

	mail.Contents = append(mail.Contents, MailContent{
		Headers: []string{
			"Content-Type: " + http.DetectContentType(rawFile),
			"Content-Transfer-Encoding: base64",
			"Content-disposition: attachment;filename=\"" + filename + "\"",
		},
		Content: base64.StdEncoding.EncodeToString(rawFile),
	})

	return nil
}

//String convert mail message body to plain text
func (mail *MailMessage) String() (string, error) {
	result := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Cc: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n",
		mail.From,
		strings.Join(mail.To, ";"),
		strings.Join(mail.Cc, ";"),
		mail.Subject)

	contentCnt := len(mail.Contents)
	if contentCnt == 1 {
		tmp, contentErr := mail.Contents[0].String()
		if contentErr != nil {
			return "", contentErr
		}
		result += tmp
	} else if contentCnt > 1 {
		result += "Content-Type: multipart/mixed; boundary=\"" + boundary + "\"\r\n"
		var tmpStr string
		var tmpErr error
		for i := 0; i < contentCnt; i++ {
			tmpStr, tmpErr = mail.Contents[i].String()
			if tmpErr != nil {
				return "", tmpErr
			}

			result += "\r\n--" + boundary + "\r\n" + tmpStr
		}
	} else {
		return "", errors.New("content message must has at least one item")
	}

	return result, nil
}
