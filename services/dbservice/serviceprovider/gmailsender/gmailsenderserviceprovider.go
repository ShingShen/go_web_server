package gmailsenderserviceprovider

import (
	"fmt"
	"server/utils/gmailsender"
)

type GmailSenderServiceFuncFactory struct{}

func (g *GmailSenderServiceFuncFactory) GetGmailSenderServiceFunc(name string) (IGmailSenderServiceFunc, error) {
	if name == "gmailSenderServiceFunc" {
		return &GmailSenderServiceFunc{}, nil
	}
	return nil, fmt.Errorf("wrong gmail service func type passed")
}

type GmailSenderServiceFunc struct{}

func (g *GmailSenderServiceFunc) GmailSender(recipient string, msg string) error {
	return gmailsender.GmailSender(recipient, msg)
}
