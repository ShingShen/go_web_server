package gmailsenderservice

import (
	"fmt"
	gmailsenderserviceprovider "server/services/dbservice/serviceprovider/gmailsender"
)

type GmailSenderServiceFactory struct{}

func (g *GmailSenderServiceFactory) GetGmailSenderService(name string) (IGmailSenderService, error) {
	if name == "gmailSender" {
		return &GmailSenderService{&gmailsenderserviceprovider.GmailSenderServiceFuncFactory{}}, nil
	}
	return nil, fmt.Errorf("wrong gmail sender type passed")
}

type GmailSenderService struct {
	gmailSenderServiceFuncFactory gmailsenderserviceprovider.IGmailSenderServiceFactory
}

func (g *GmailSenderService) GmailSender(recipient string, msg string) error {
	gmailSenderServiceFunc, _ := g.gmailSenderServiceFuncFactory.GetGmailSenderServiceFunc("gmailSenderServiceFunc")
	return gmailSenderServiceFunc.GmailSender(recipient, msg)
}
