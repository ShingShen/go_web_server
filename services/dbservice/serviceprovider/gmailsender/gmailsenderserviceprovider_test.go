package gmailsenderserviceprovider

import (
	"testing"
)

func TestGmailSenderServiceFunc(t *testing.T) {
	gmailSenderServiceFuncFactory := &GmailSenderServiceFuncFactory{}
	const name string = "gmailSenderServiceFunc"
	const errName string = "errName"

	// getGmailSenderServiceFunc
	getGmailSenderServiceFunc, err := gmailSenderServiceFuncFactory.GetGmailSenderServiceFunc(name)
	if err == nil {
		t.Logf("getGmailSenderServiceFunc passed: %v, %v", getGmailSenderServiceFunc, err)
	} else {
		t.Errorf("getGmailSenderServiceFunc failed: %v, %v", getGmailSenderServiceFunc, err)
	}

	// getGmailSenderServiceFunc Name Err
	getGmailSenderServiceFuncNameErr, err := gmailSenderServiceFuncFactory.GetGmailSenderServiceFunc(errName)
	if err != nil {
		t.Logf("getGmailSenderServiceFuncNameErr passed: %v, %v", getGmailSenderServiceFuncNameErr, err)
	} else {
		t.Errorf("getGmailSenderServiceFuncNameErr failed: %v, %v", getGmailSenderServiceFuncNameErr, err)
	}
}
