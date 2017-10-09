# smtpclient
SMTP client library

## Features
- SMTP with TLS connection
- Support attachment
- Support HTML format content
- Support mix content (e.g. HTML + XML + file attachment)
- Support custom MIME content (you can customize headers and string content)

## Example 1 - direct send email
```go
//create email message
message := smtpclient.MailMessage{
    From: "john@gmail.com",
    To: []string{
        "friendA@hotmail.com",
        "friendB@yahoo.com",
    },
    Cc: []string{
        "friendC@mail.com"
    },
    Subject: "This is sample",
    Content: []smtpclient.MailContent{},
}
message.AddHTMLContent("<html><body><h1>Hello</h1></body></html>")
message.AddTextContent("hello again")
message.AddAttachment("/home/john/abc.txt", "customFilename.txt")

//send email
err := smtpclient.SendEmail(
    "smtp.gmail.com", //email server URL
    465, //secure TLS port number
    smtpclient.TCP,
    "john@gmail.com", //your email username
    "your-email-password",
    message
)
```

## Example 2 - send email with helper
```go
//create email message
message := smtpclient.MailMessage{
    From: "john@gmail.com",
    To: []string{
        "friendA@hotmail.com",
        "friendB@yahoo.com",
    },
    Cc: []string{
        "friendC@mail.com"
    },
    Subject: "This is sample",
    Content: []smtpclient.MailContent{},
}
message.AddHTMLContent("<html><body><h1>Hello</h1></body></html>")
message.AddTextContent("hello again")
message.AddAttachment("/home/john/abc.txt", "customFilename.txt")

//create email helper
helper := smtpclient.EmailHelper{
    ServerAddress: "smtp.gmail.com"
	ServerPortNumber: 465
	NetworkProtocol: smtpclient.TCP
	Username: "john@gmail.com"
	Password: "your-email-password"
}

//send email
err := helper.sendMail(&message)
```

## Example 3 - add custom content
```go
//create email message
message := smtpclient.MailMessage{
    From: "john@gmail.com",
    To: []string{
        "friendA@hotmail.com",
        "friendB@yahoo.com",
    },
    Cc: []string{
        "friendC@mail.com"
    },
    Subject: "This is sample",
    Content: []smtpclient.MailContent{},
}

//add custom content
message.Content = append(message.Content, smtpclient.MailContent{
    Header: []string{ //write your own custom header here; you can have more than one header(s)
        `Content-Type: application/json; charset="utf8"`,
        `Content-Transfer-Encoding: utf8`,
    },
    Content: `{"title":"Jungle Booklet", "author":"Chuck Norris"}`
})
```