package gmailsenderserviceprovider

type IGmailSenderServiceFactory interface {
	GetGmailSenderServiceFunc(name string) (IGmailSenderServiceFunc, error)
}

type IGmailSenderServiceFunc interface {
	GmailSender(recipient string, msg string) error
}
