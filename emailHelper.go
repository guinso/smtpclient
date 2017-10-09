package smtpclient

//EmailHelper SMTP client helper tool
type EmailHelper struct {
	ServerAddress    string
	ServerPortNumber int
	NetworkProtocol  EmailNetworkProtocol
	Username         string
	Password         string
}

//SendEmail send simple SMTP mail
func (helper *EmailHelper) SendEmail(message *MailMessage) error {
	return SendEmail(helper.ServerAddress, helper.ServerPortNumber,
		helper.NetworkProtocol,
		helper.Username, helper.Password,
		message)
}
