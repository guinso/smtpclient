package smtpclient

//SMTPHelper SMTP client helper tool
type SMTPHelper struct {
	ServerAddress    string
	ServerPortNumber int
	NetworkProtocol  EmailNetworkProtocol
	Username         string
	Password         string
}

//SendEmail send simple SMTP mail
func (helper *SMTPHelper) SendEmail(message *MailMessage) error {
	return SendEmail(helper.ServerAddress, helper.ServerPortNumber,
		helper.NetworkProtocol,
		helper.Username, helper.Password,
		message)
}
