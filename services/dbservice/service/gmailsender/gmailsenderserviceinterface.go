package gmailsenderservice

type IGmailSenderServiceFactory interface {
	GetGmailSenderService(name string) (IGmailSenderService, error)
}

type IGmailSenderService interface {
	GmailSender(recipient string, msg string) error
}
