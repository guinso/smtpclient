package smtpclient

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

//EmailNetworkProtocol network type used to connect to email server
type EmailNetworkProtocol string

const (
	//TCP TCP network
	TCP EmailNetworkProtocol = "tcp"
	//UDP UDP network
	UDP EmailNetworkProtocol = "udp"
)

//SendEmail send SMTP email
func SendEmail(emailServer string, portNumber int,
	networkProtocol EmailNetworkProtocol,
	username string, password string,
	message *MailMessage) error {

	msg, err := message.String()
	if err != nil {
		return err
	}

	return SendBasicEmail(
		emailServer, portNumber, networkProtocol,
		username, password,
		message.From,
		message.To,
		msg)
}

//SendBasicEmail send SMTP email
func SendBasicEmail(emailServer string, portNumber int,
	networkProtocol EmailNetworkProtocol, username string,
	password string,
	from string,
	to []string,
	message string) error {

	tlsConfig := tls.Config{
		ServerName:         emailServer,
		InsecureSkipVerify: true,
	}

	//establish TLS connection
	conn, connErr := tls.Dial(
		string(networkProtocol),
		fmt.Sprintf("%s:%d", emailServer, portNumber),
		&tlsConfig)
	if connErr != nil {
		return connErr
	}

	defer conn.Close()

	//create email client
	client, clientErr := smtp.NewClient(conn, emailServer)
	if clientErr != nil {
		return clientErr
	}

	defer client.Close()

	//create plain authenticate credential
	auth := smtp.PlainAuth("", username, password, emailServer)

	//authenticate email client with credential
	if authErr := client.Auth(auth); authErr != nil {
		return authErr
	}

	var err error

	//set sent from
	if err = client.Mail(from); err != nil {
		return err
	}

	//set to receipain(s)
	for _, rcp := range to {
		if err = client.Rcpt(rcp); err != nil {
			return err
		}
	}

	//start write message
	writter, writterErr := client.Data()
	if writterErr != nil {
		return writterErr
	}

	if _, err = writter.Write([]byte(message)); err != nil {
		writter.Close()
		return err
	}

	if err = writter.Close(); err != nil {
		return err
	}

	return nil
}
