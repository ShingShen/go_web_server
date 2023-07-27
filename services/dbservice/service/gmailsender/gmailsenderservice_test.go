package gmailsenderservice

import (
	"testing"
)

func TestGmailSenderTempService(t *testing.T) {
	gmailSenderServiceFactory := &GmailSenderServiceFactory{}
	const name string = "gmailSender"
	const errName string = "errName"

	getGmailSenderService, err := gmailSenderServiceFactory.GetGmailSenderService(name)
	if err == nil {
		t.Logf("getGmailSenderService passed: %v, %v", getGmailSenderService, err)
	} else {
		t.Errorf("getGmailSenderService failed: %v, %v", getGmailSenderService, err)
	}

	getGmailSenderServiceNameErr, err := gmailSenderServiceFactory.GetGmailSenderService(errName)
	if err != nil {
		t.Logf("getGmailSenderServiceNameErr passed: %v, %v", getGmailSenderServiceNameErr, err)
	} else {
		t.Errorf("getGmailSenderServiceNameErr failed: %v, %v", getGmailSenderServiceNameErr, err)
	}
}
